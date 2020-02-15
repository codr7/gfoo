package gfoo

import (
	"bufio"
	//"fmt"
	"io"
	"strings"
	"unicode"
)

func isId(c rune) bool {
	return unicode.IsGraphic(c) && !unicode.IsSpace(c) &&
		c != '(' && c != ')' && c != '{' && c != '}' && c != '[' && c != ']' &&
		c != '\'' && c != '"'
}

func (self *Scope) parseForm(in *bufio.Reader, pos *Pos) (Form, error) {
	c, _, err := in.ReadRune()
	
	if err != nil {
		return nil, err
	}

	switch c {
	case '\'':
		fpos := *pos
		pos.column++
		var f Form

		if f, err = self.parseForm(in, pos); err != nil {
			return nil, err
		}
		
		return NewQuote(f, fpos), nil
	case '"':
		return self.parseString(in, pos)
	case '(':
		return self.parseGroup(in, pos)
	case '{':
		return self.parseScope(in, pos)
	case '[':
		return self.parseSlice(in, pos)
	default:
		if unicode.IsDigit(c) {
			return self.parseNumber(in, c, pos)
		}

		if isId(c) {
			return self.parseId(in, c, pos)
		}
	}

	return nil, self.Error(*pos, "Unexpected input: %v", c)
}

func (self *Scope) parseForms(in *bufio.Reader, pos *Pos, end rune) ([]Form, error) {
	var out []Form
	
	for {
		if err := skipSpace(in, pos); err != nil {
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

		if f, err = self.parseForm(in, pos); err != nil {
			return nil, err
		}

		out = append(out, f)
	}

	return out, nil
}

func (self *Scope) parseGroup(in *bufio.Reader, pos *Pos) (Form, error) {
	fpos := *pos
	pos.column++
	forms, err := self.parseForms(in, pos, ')')

	if err != nil {
		return nil, err
	}

	return NewGroup(forms, fpos), nil
}

func (self *Scope) parseId(in *bufio.Reader, c rune, pos *Pos) (Form, error) {
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

		if !isId(c) || (c == '_' && pc != 0) || (c == '.' && pc != 0 && pc != '.') {
			if err = in.UnreadRune(); err != nil {
				return nil, err
			}
			
			break
		}
		
		pos.column++
	}

	c, _, err = in.ReadRune()

	if err != nil && err != io.EOF {
		return nil, err
	}

	if err == nil {
		if c == '(' {
			var f Form

			if f, err = self.parseGroup(in, pos); err != nil {
				return nil, err
			}

			f.(*Group).Push(NewId(buffer.String(), fpos))
			return f, nil
		} else if err = in.UnreadRune(); err != nil {
			return nil, err
		}
	}
		
	return NewId(buffer.String(), fpos), nil
}

func (self *Scope) parseNumber(in *bufio.Reader, c rune, pos *Pos) (Form, error) {
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
				return NewLiteral(NewVal(&TInt64, v), fpos), nil
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
	
	return NewLiteral(NewVal(&TInt64, v), fpos), nil
}

func (self *Scope) parseScope(in *bufio.Reader, pos *Pos) (Form, error) {
	fpos := *pos
	pos.column++
	forms, err := self.parseForms(in, pos, '}')

	if err != nil {
		return nil, err
	}
	
	return NewScopeForm(forms, fpos), nil
}

func (self *Scope) parseSlice(in *bufio.Reader, pos *Pos) (Form, error) {
	fpos := *pos
	pos.column++
	forms, err := self.parseForms(in, pos, ']')

	if err != nil {
		return nil, err
	}
	
	return NewSliceForm(forms, fpos), nil
}

func (self *Scope) parseString(in *bufio.Reader, pos *Pos) (Form, error) {
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

func skipSpace(in *bufio.Reader, pos *Pos) error {
	for {
		c, _, err := in.ReadRune()
		
		if err != nil {
			return err
		}

		switch c {
		case ' ':
			pos.column++
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
