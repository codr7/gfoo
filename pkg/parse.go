package gfoo

import (
	"bufio"
	"bytes"
	"unicode"
)

func (gfoo *GFoo) parseForm(in *bufio.Reader, pos *Position) (Form, error) {
	c, _, err := in.ReadRune()
	
	if err != nil {
		return nil, err
	}

	switch c {
	case '"':
		pos.column++
		return gfoo.parseString(in, pos)
	case '(':
		pos.column++
		return gfoo.parseGroup(in, pos)
	default:
		if unicode.IsDigit(c) {
			pos.column++
			return gfoo.parseNumber(in, c, pos)
		}

		if unicode.IsGraphic(c) {
			pos.column++
			return gfoo.parseIdentifier(in, c, pos)
		}
	}

	return nil, gfoo.Errorf(*pos, "Unexpected input: %v", c)
}

func (gfoo *GFoo) parseGroup(in *bufio.Reader, pos *Position) (Form, error) {
	var forms []Form
	var f Form
	
	for {
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

		if f, err = gfoo.parseForm(in, pos); err != nil {
			return nil, err
		}

		forms = append(forms, f)
	}
	
	return NewGroup(forms), nil
}

func (gfoo *GFoo) parseIdentifier(in *bufio.Reader, c rune, pos *Position) (Form, error) {
	return nil, nil
}

func (gfoo *GFoo) parseNumber(in *bufio.Reader, c rune, pos *Position) (Form, error) {
	return nil, nil
}

func (gfoo *GFoo) parseString(in *bufio.Reader, pos *Position) (Form, error) {
	var buffer bytes.Buffer

	for {
		c, _, err := in.ReadRune()
		
		if err != nil {
			return nil, err
		}

		if c == '"' {
			break
		}
	}
	
	return NewLiteral(&String, buffer.String()), nil
}
