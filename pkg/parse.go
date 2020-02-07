package gfoo

import (
	"bufio"
	"bytes"
	"io"
	"unicode"
)

func skipSpace(in *bufio.Reader, pos *Position) error {
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

func (self *GFoo) parseForm(in *bufio.Reader, pos *Position) (Form, error) {
	c, _, err := in.ReadRune()
	
	if err != nil {
		return nil, err
	}

	switch c {
	case '\'':
		pos.column++
		var f Form

		if f, err = self.parseIdentifier(in,  0, pos); err != nil {
			return nil, err
		}
		
		return f.Quote().Literal(), nil
	case '"':
		pos.column++
		return self.parseString(in, pos)
	case '(':
		pos.column++
		return self.parseSlice(in, pos)
	default:
		if unicode.IsDigit(c) {
			pos.column++
			return self.parseNumber(in, c, pos)
		}

		if unicode.IsGraphic(c) {
			pos.column++
			return self.parseIdentifier(in, c, pos)
		}
	}

	return nil, self.Errorf(*pos, "Unexpected input: %v", c)
}

func (self *GFoo) parseSlice(in *bufio.Reader, pos *Position) (Form, error) {
	var forms []Form
	var f Form
	
	for {
		if err := skipSpace(in, pos); err != nil {
			return nil, err
		}
		
		c, _, err := in.ReadRune()
		
		if err != nil {
			return nil, err
		}

		if c == ')' {
			break
		}

		if err = in.UnreadRune(); err != nil {
			return nil, err
		}

		if f, err = self.parseForm(in, pos); err != nil {
			return nil, err
		}

		forms = append(forms, f)
	}

	return NewSliceForm(forms), nil
}

func (self *GFoo) parseIdentifier(in *bufio.Reader, c rune, pos *Position) (Form, error) {
	var buffer bytes.Buffer
	var err error
	
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
	}
	
	return NewIdentifier(buffer.String()), nil
}

func (self *GFoo) parseNumber(in *bufio.Reader, c rune, pos *Position) (Form, error) {
	v := int64(0)
	base := int64(10)
	var err error

	if c == 0 {
		if c, _, err = in.ReadRune(); err != nil {
			return nil, err
		}

		pos.column++

		if !unicode.IsDigit(c) {
			return nil, self.Errorf(*pos, "Expected number: %v", c)
		}
	}
	
	if c == '0' {
		if c, _, err = in.ReadRune(); err != nil {
			if err == io.EOF {
				return NewLiteral(&Int64, v), nil
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

		if !unicode.IsDigit(c) {
			if err = in.UnreadRune(); err != nil {
				return nil, err
			}
			
			break
		}
	}
	
	return NewLiteral(&Int64, v), nil
}

func (self *GFoo) parseString(in *bufio.Reader, pos *Position) (Form, error) {
	var buffer bytes.Buffer

	for {
		c, _, err := in.ReadRune()
		
		if err != nil {
			return nil, err
		}

		if c == '"' {
			break
		}

		if _, err = buffer.WriteRune(c); err != nil {
			return nil, err
		}
	}
	
	return NewLiteral(&String, buffer.String()), nil
}
