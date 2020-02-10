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

func (self *VM) parseForm(in *bufio.Reader, pos *Pos) (Form, error) {
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
		
		return NewQuote(fpos, f), nil
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

func (self *VM) parseForms(in *bufio.Reader, pos *Pos, end rune) ([]Form, error) {
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

func (self *VM) parseGroup(in *bufio.Reader, pos *Pos) (Form, error) {
	fpos := *pos
	pos.column++
	forms, err := self.parseForms(in, pos, ')')

	if err != nil {
		return nil, err
	}

	return NewGroup(fpos, forms), nil
}

func (self *VM) parseId(in *bufio.Reader, c rune, pos *Pos) (Form, error) {
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

			f.(*Group).AddForm(NewId(fpos, buffer.String()))
			return f, nil
		} else if err = in.UnreadRune(); err != nil {
			return nil, err
		}
	}
		
	return NewId(fpos, buffer.String()), nil
}

func (self *VM) parseNumber(in *bufio.Reader, c rune, pos *Pos) (Form, error) {
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
				return NewLiteral(fpos, NewVal(&TInt64, v)), nil
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
	
	return NewLiteral(fpos, NewVal(&TInt64, v)), nil
}

func (self *VM) parseScope(in *bufio.Reader, pos *Pos) (Form, error) {
	fpos := *pos
	pos.column++
	forms, err := self.parseForms(in, pos, '}')

	if err != nil {
		return nil, err
	}
	
	return NewScopeForm(fpos, forms), nil
}

func (self *VM) parseSlice(in *bufio.Reader, pos *Pos) (Form, error) {
	fpos := *pos
	pos.column++
	forms, err := self.parseForms(in, pos, ']')

	if err != nil {
		return nil, err
	}
	
	return NewSliceForm(fpos, forms), nil
}

func (self *VM) parseString(in *bufio.Reader, pos *Pos) (Form, error) {
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
	
	return NewLiteral(fpos, NewVal(&TString, buffer.String())), nil
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
