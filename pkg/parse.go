package gfoo

import (
	"bufio"
	"io"
	"strings"
	"unicode"
)

func IsId(c rune) bool {
	return unicode.IsGraphic(c) && !unicode.IsSpace(c) &&
		c != ',' && c != ';' && c != '\'' && c != '@' && c != '\\' && c != '"' &&
		c != '(' && c != ')' && c != '{' && c != '}' && c != '[' && c != ']'
}

func (self *Scope) ParseBody(
	in *bufio.Reader,
	end rune,
	create func([]Form, Pos) Form,
	depth int,
	pos *Pos) (Form, error) {
	fpos := *pos
	pos.column++
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
			if depth == 0 {
				pos.column++
			} else {
				if err = in.UnreadRune(); err != nil {
					return nil, err
				}
			}
			
			break
				
		}

		var f Form

		if c == ';' {
			if f, err = self.ParseBody(in, end, create, depth+1, pos); err != nil {
				return nil, err
			}
		} else {
			if err = in.UnreadRune(); err != nil {
				return nil, err
			}
			
			if f, err = self.ParseForm(in, pos); err != nil {
				return nil, err
			}
		}

		out = append(out, f)
	}

	return create(out, fpos), nil
}

func (self *Scope) ParseChar(in *bufio.Reader, pos *Pos) (Form, error) {
	fpos := *pos
	pos.column++
	c, _, err := in.ReadRune()
	
	if err != nil {
		return nil, Error(fpos, "Invalid character literal: %v", err) 
	}

	pos.column++
	var out rune
	
	switch c {
	case '\'':
		out, _, err = in.ReadRune()
		
		if err != nil {
			return nil, Error(fpos, "Invalid character literal: %v", err)
		}
	case 'n':
		out = '\n'
	default:
		return nil, Error(fpos, "Invalid character literal: %v", c)	
	}
	
	return NewLiteral(NewVal(&TChar, out), fpos), nil
}

func (self *Scope) ParseForm(in *bufio.Reader, pos *Pos) (Form, error) {
	c, _, err := in.ReadRune()
	
	if err != nil {
		return nil, err
	}

	switch c {
	case '!':
		return self.ParseNegate(in, pos)
	case ',':
		return self.ParsePair(in, pos)
	case '\'':
		return self.ParseQuote(in, pos)
	case '@':
		return self.ParseUnquote(in, pos)
	case '\\':
		return self.ParseChar(in, pos)
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

	return nil, Error(*pos, "Unexpected input: %v", c)
}

func (self *Scope) ParseGroup(in *bufio.Reader, pos *Pos) (Form, error) {
	return self.ParseBody(in, ')', func(body []Form, pos Pos) Form { return NewGroup(body, pos) }, 0, pos)
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

	return NewId(buffer.String(), fpos), nil
}

func (self *Scope) ParseNumber(in *bufio.Reader, c rune, pos *Pos) (Form, error) {
	var v Int
	base := Int(10)
	var err error
	fpos := *pos
	
	if c == 0 {
		if c, _, err = in.ReadRune(); err != nil {
			return nil, err
		}

		pos.column++

		if !unicode.IsDigit(c) {
			return nil, Error(*pos, "Expected number: %v", c)
		}
	} else {
		pos.column++
	}
	
	if c == '0' {
		if c, _, err = in.ReadRune(); err != nil {
			if err == io.EOF {
				return NewLiteral(NewVal(&TInt, v), fpos), nil
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
			var dv Int
			
			if base == 16 && c >= 'a' && c <= 'f' {
				dv = 10 + Int(c - 'a')
			} else {
				dv = Int(c - '0')
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
	
	return NewLiteral(NewVal(&TInt, v), fpos), nil
}

func (self *Scope) ParseNegate(in *bufio.Reader, pos *Pos) (Form, error) {
	fpos := *pos
	pos.column++
	var f Form
	var err error
	
	if f, err = self.ParseForm(in, pos); err != nil {
		return nil, err
	}

	return NewNegateForm(f, fpos), nil
}

func (self *Scope) ParsePair(in *bufio.Reader, pos *Pos) (Form, error) {
	fpos := *pos
	pos.column++
	var l, r Form
	var err error
	
	if l, err = self.ParseForm(in, pos); err != nil {
		return nil, err
	}

	if err := SkipSpace(in, pos); err != nil {
		return nil, err
	}

	if r, err = self.ParseForm(in, pos); err != nil {
		return nil, err
	}

	return NewPairForm(l, r, fpos), nil
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
	return self.ParseBody(in, '}', func(body []Form, pos Pos) Form { return NewScopeForm(body, pos) }, 0, pos)
}

func (self *Scope) ParseSlice(in *bufio.Reader, pos *Pos) (Form, error) {
	return self.ParseBody(in, ']', func(body []Form, pos Pos) Form { return NewSliceForm(body, pos) }, 0, pos)
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
