package gfoo

import (
	"bufio"
	"io"
	"strings"
	"unicode"
)

func (self *GFoo) parseForm(in *bufio.Reader, pos *Pos) (Form, error) {
	c, _, err := in.ReadRune()
	
	if err != nil {
		return nil, err
	}

	switch c {
	case '\'':
		fpos := *pos
		pos.column++
		var f Form

		if f, err = self.parseId(in,  0, pos); err != nil {
			return nil, err
		}
		
		return f.Quote().Literal(fpos), nil
	case '"':
		return self.parseString(in, pos)
	case '(':
		return self.parseGroup(in, pos)
	case '[':
		return self.parseSlice(in, pos)
	default:
		if unicode.IsDigit(c) {
			return self.parseNumber(in, c, pos)
		}

		if unicode.IsGraphic(c) {
			return self.parseId(in, c, pos)
		}
	}

	return nil, self.Error(*pos, "Unexpected input: %v", c)
}

func (self *GFoo) parseForms(in *bufio.Reader, pos *Pos, end rune) ([]Form, error) {
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

func (self *GFoo) parseId(in *bufio.Reader, c rune, pos *Pos) (Form, error) {
	var buffer strings.Builder
	var err error
	fpos := *pos

	if c > 0 {
		pos.column++
	}
	
	for {
		if c > 0 {
			if !unicode.IsGraphic(c) || c == '(' || c == ')' || c == '\'' || c == '"' {
				if err = in.UnreadRune(); err != nil {
					return nil, err
				}
				
				break
			}
			
			if _, err = buffer.WriteRune(c); err != nil {
				return nil, err
			}
		}

		c, _, err = in.ReadRune()
		
		if err == io.EOF {
			break
		}
		
		if err != nil {
			return nil, err
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

func (self *GFoo) parseNumber(in *bufio.Reader, c rune, pos *Pos) (Form, error) {
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
				return NewLiteral(fpos, &TInt64, v), nil
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
	
	return NewLiteral(fpos, &TInt64, v), nil
}

func (self *GFoo) parseGroup(in *bufio.Reader, pos *Pos) (Form, error) {
	fpos := *pos
	pos.column++
	forms, err := self.parseForms(in, pos, ')')

	if err != nil {
		return nil, err
	}

	return NewGroup(fpos, forms), nil
}

func (self *GFoo) parseSlice(in *bufio.Reader, pos *Pos) (Form, error) {
	fpos := *pos
	pos.column++
	forms, err := self.parseForms(in, pos, ']')

	if err != nil {
		return nil, err
	}
	
	return NewSlice(fpos, forms), nil
}

func (self *GFoo) parseString(in *bufio.Reader, pos *Pos) (Form, error) {
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
	
	return NewLiteral(fpos, &String, buffer.String()), nil
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
