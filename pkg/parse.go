package gfoo

import (
	"bufio"
	//"fmt"
	"io"
	"math/big"
	"strings"
	"unicode"
)

func IsId(c rune) bool {
	return unicode.IsGraphic(c) && !unicode.IsSpace(c) &&
		c != '\'' && c != '@' && c != '"' &&
		c != '(' && c != ')' && c != '{' && c != '}' && c != '[' && c != ']'
}

func (self *Scope) ParseBody(in *bufio.Reader, end rune, pos *Pos) ([]Form, error) {
	var out []Form
	
	for {
		if err := SkipSpace(in, pos); err != nil {
			return nil, err
		}
		
		c, _, err := in.ReadRune()
		
		if err != nil {
			return nil, err
		}

		if c == end {
			break
		}

		if err = in.UnreadRune(); err != nil {
			return nil, err
		}

		var f Form

		if f, err = self.ParseForm(in, pos); err != nil {
			return nil, err
		}

		out = append(out, f)
	}

	return out, nil
}

func (self *Scope) ParseForm(in *bufio.Reader, pos *Pos) (Form, error) {
	c, _, err := in.ReadRune()
	
	if err != nil {
		return nil, err
	}

	switch c {
	case '\'':
		return self.ParseQuote(in, pos)
	case '@':
		return self.ParseUnquote(in, pos)
	case '"':
		return self.ParseString(in, pos)
	case '(':
		return self.ParseGroup(in, pos)
	case '{':
		return self.ParseScope(in, pos)
	case '[':
		return self.ParseSlice(in, pos)
	default:
		if unicode.IsDigit(c) {
			return self.ParseNumber(in, c, pos)
		}

		if IsId(c) {
			return self.ParseId(in, c, pos)
		}
	}

	return nil, self.Error(*pos, "Unexpected input: %v", c)
}

func (self *Scope) ParseGroup(in *bufio.Reader, pos *Pos) (Form, error) {
	fpos := *pos
	pos.column++
	body, err := self.ParseBody(in, ')', pos)

	if err != nil {
		return nil, err
	}

	return NewGroup(body, fpos), nil
}

func (self *Scope) ParseId(in *bufio.Reader, c rune, pos *Pos) (Form, error) {
	var buffer strings.Builder
	var err error
	fpos := *pos
	var pc rune
	
	if c > 0 {
		pos.column++
	}
	
	for {
		if c > 0 {			
			if _, err = buffer.WriteRune(c); err != nil {
				return nil, err
			}
		}

		pc = c
		c, _, err = in.ReadRune()
		
		if err == io.EOF {
			break
		}
		
		if err != nil {
			return nil, err
		}

		if !IsId(c) ||
			(c == ',' && pc != 0) ||
 			(c == '_' && pc != 0) ||
			(c == '.' && pc != 0 && pc != '.') {
			if err = in.UnreadRune(); err != nil {
				return nil, err
			}
			
			break
		}
		
		pos.column++
	}

	if pc != ':' {
		c, _, err = in.ReadRune()
		
		if err != nil && err != io.EOF {
			return nil, err
		}
		
		if err == nil {
			if c == '(' {
				var f Form
				
				if f, err = self.ParseGroup(in, pos); err != nil {
					return nil, err
				}
				
				f.(*Group).Push(NewId(buffer.String(), fpos))
				return f, nil
			} else if err = in.UnreadRune(); err != nil {
				return nil, err
			}
		}
	}
		
	return NewId(buffer.String(), fpos), nil
}

func (self *Scope) ParseNumber(in *bufio.Reader, c rune, pos *Pos) (Form, error) {
	v := int64(0)
	base := int64(10)
	var err error
	fpos := *pos
	
	if c == 0 {
		if c, _, err = in.ReadRune(); err != nil {
			return nil, err
		}

		pos.column++

		if !unicode.IsDigit(c) {
			return nil, self.Error(*pos, "Expected number: %v", c)
		}
	} else {
		pos.column++
	}
	
	if c == '0' {
		if c, _, err = in.ReadRune(); err != nil {
			if err == io.EOF {
				return NewLiteral(NewVal(&TInt, big.NewInt(v)), fpos), nil
			}
			
			return nil, err
		}
		
		switch c {
		case 'b':
			pos.column++
			base = 2
			c = 0
		case 'x':
			pos.column++
			base = 16
			c = 0
		default:			
			if err = in.UnreadRune(); err != nil {
				return nil, err
			}

			c = '0'
		}
	}

	for {
		if c > 0 {
			var dv int64
			
			if base == 16 && c >= 'a' && c <= 'f' {
				dv = 10 + int64(c - 'a')
			} else {
				dv = int64(c - '0')
			}

			v = v * base + dv
		}
		
		c, _, err = in.ReadRune()

		if err == io.EOF {
			break
		}
		
		if err != nil {
			return nil, err
		}

		if !unicode.IsDigit(c) && (base != 16 || c < 'a' || c > 'f') {
			if err = in.UnreadRune(); err != nil {
				return nil, err
			}
			
			break
		}

		pos.column++
	}
	
	return NewLiteral(NewVal(&TInt, big.NewInt(v)), fpos), nil
}

func (self *Scope) ParseQuote(in *bufio.Reader, pos *Pos) (Form, error) {
	fpos := *pos
	pos.column++
	var f Form
	var err error
	
	if f, err = self.ParseForm(in, pos); err != nil {
		return nil, err
	}
	
	return NewQuote(f, fpos), nil
}

func (self *Scope) ParseScope(in *bufio.Reader, pos *Pos) (Form, error) {
	fpos := *pos
	pos.column++
	body, err := self.ParseBody(in, '}', pos)

	if err != nil {
		return nil, err
	}
	
	return NewScopeForm(body, fpos), nil
}

func (self *Scope) ParseSlice(in *bufio.Reader, pos *Pos) (Form, error) {
	fpos := *pos
	pos.column++
	body, err := self.ParseBody(in, ']', pos)

	if err != nil {
		return nil, err
	}
	
	return NewSliceForm(body, fpos), nil
}

func (self *Scope) ParseString(in *bufio.Reader, pos *Pos) (Form, error) {
	var buffer strings.Builder
	fpos := *pos
	pos.column++
	
	for {
		c, _, err := in.ReadRune()

		if err == io.EOF {
			break
		}
		
		if err != nil {
			return nil, err
		}

		pos.column++

		if c == '"' {
			break
		}

		if _, err = buffer.WriteRune(c); err != nil {
			return nil, err
		}
	}
	
	return NewLiteral(NewVal(&TString, buffer.String()), fpos), nil
}

func (self *Scope) ParseUnquote(in *bufio.Reader, pos *Pos) (Form, error) {
	fpos := *pos
	pos.column++
	var f Form
	var err error
	
	if f, err = self.ParseForm(in, pos); err != nil {
		return nil, err
	}
	
	return NewUnquote(f, fpos), nil
}

func SkipSpace(in *bufio.Reader, pos *Pos) error {
	for {
		c, _, err := in.ReadRune()
		
		if err != nil {
			return err
		}

		switch c {
		case ' ':
			pos.column++
		case '\r':
		case '\n':
			pos.line++
			pos.column = MIN_COLUMN
		default:
			if err = in.UnreadRune(); err != nil {
				return err
			}

			return nil
		}
	}
}
