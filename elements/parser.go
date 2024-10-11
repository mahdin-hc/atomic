package elements

import (
	"errors"
	"fmt"
	"strconv"
	"unicode"
)

// Token Types
type TokenType int

const (
	TOKEN_ELEMENT TokenType = iota
	TOKEN_NUMBER
	TOKEN_PLUS
	TOKEN_MINUS
	TOKEN_LPAREN
	TOKEN_RPAREN
	TOKEN_LBRACE
	TOKEN_RBRACE
	TOKEN_END
)

type Token struct {
	typ   TokenType
	value string
}

// ParserError represents an error that occurs during parsing.
type ParserError struct {
	Message string
	Pos     int
}

func (e *ParserError) Error() string {
	return fmt.Sprintf("Error at position %d: %s", e.Pos, e.Message)
}

// Parser for the chemical formula
type Parser struct {
	input []rune
	pos   int
	token Token
}

// NewParser initializes a new Parser with the input string
func NewParser(input string) *Parser {
	return &Parser{input: []rune(input)}
}

// Peek the next token without advancing the current position
func (p *Parser) peek() (Token, error) {
	oldPos := p.pos
	oldToken := p.token
	if err := p.nextToken(); err != nil {
		return Token{}, err
	}
	peekedToken := p.token
	p.pos = oldPos
	p.token = oldToken
	return peekedToken, nil
}

// Advance to the next token
func (p *Parser) nextToken() error {
	for p.pos < len(p.input) && unicode.IsSpace(p.input[p.pos]) {
		p.pos++
	}

	if p.pos >= len(p.input) {
		p.token = Token{typ: TOKEN_END}
		return nil
	}

	switch ch := p.input[p.pos]; {
	case unicode.IsUpper(ch):
		start := p.pos
		p.pos++
		for p.pos < len(p.input) && unicode.IsLower(p.input[p.pos]) {
			p.pos++
		}
		p.token = Token{typ: TOKEN_ELEMENT, value: string(p.input[start:p.pos])}
	case unicode.IsDigit(ch):
		start := p.pos
		for p.pos < len(p.input) && unicode.IsDigit(p.input[p.pos]) {
			p.pos++
		}
		p.token = Token{typ: TOKEN_NUMBER, value: string(p.input[start:p.pos])}
	case ch == '(':
		p.pos++
		p.token = Token{typ: TOKEN_LPAREN, value: "("}
	case ch == ')':
		p.pos++
		p.token = Token{typ: TOKEN_RPAREN, value: ")"}
	case ch == '[':
		p.pos++
		p.token = Token{typ: TOKEN_LBRACE, value: "["}
	case ch == ']':
		p.pos++
		p.token = Token{typ: TOKEN_RBRACE, value: "]"}
	case ch == '+':
		p.pos++
		p.token = Token{typ: TOKEN_PLUS, value: "+"}
	case ch == '-':
		p.pos++
		p.token = Token{typ: TOKEN_MINUS, value: "-"}
	default:
		return &ParserError{Message: fmt.Sprintf("Unknown character: %c", ch), Pos: p.pos}
	}
	return nil
}

// Expect a specific token type, return an error if not matching
func (p *Parser) expect(typ TokenType) error {
	if p.token.typ != typ {
		return &ParserError{Message: fmt.Sprintf("Expected token type %v, got %v", typ, p.token.typ), Pos: p.pos}
	}
	return nil
}

// Parse a compound (may contain elements or nested groups)
func (p *Parser) ParseCompound() (Compound, error) {
	compound := Compound{}
	err := p.nextToken()
	if err != nil {
		return compound, err
	}
	for p.token.typ == TOKEN_ELEMENT || p.token.typ == TOKEN_LPAREN || p.token.typ == TOKEN_LBRACE {
		if p.token.typ == TOKEN_ELEMENT {
			elements, err := p.parseElements()
			if err != nil {
				return compound, err
			}
			compound.Molecules = append(compound.Molecules, Molecule{Elements: elements})
		} else if p.token.typ == TOKEN_LPAREN || p.token.typ == TOKEN_LBRACE {
			groupMolecule, err := p.parseMolecules()
			if err != nil {
				return compound, err
			}
			compound.Molecules = append(compound.Molecules, groupMolecule...)
		}
	}
	if p.token.typ == TOKEN_PLUS || p.token.typ == TOKEN_MINUS {
		charge, err := p.parseCharge()
		if err != nil {
			return compound, err
		}
		compound.Charge = charge
	}
	
	return compound, nil
}

// Parse a group, which is either inside parentheses or square brackets
func (p *Parser) parseMolecules() ([]Molecule, error) {
	var groupOpen, groupClose TokenType
	if p.token.typ == TOKEN_LPAREN {
		groupOpen, groupClose = TOKEN_LPAREN, TOKEN_RPAREN
	} else if p.token.typ == TOKEN_LBRACE {
		groupOpen, groupClose = TOKEN_LBRACE, TOKEN_RBRACE
	}

	if err := p.expect(groupOpen); err != nil {
		return []Molecule{}, err
	}
	if err := p.nextToken(); err != nil {
		return []Molecule{}, err
	}

	molecules := []Molecule{}
	molecule := Molecule{}
	for p.token.typ == TOKEN_ELEMENT || p.token.typ == TOKEN_LPAREN || p.token.typ == TOKEN_LBRACE {
		if p.token.typ == TOKEN_ELEMENT {
			elements, err := p.parseElements()
			if err != nil {
				return molecules, err
			}
			molecule.Elements = append(molecule.Elements, elements...)
		} else if p.token.typ == TOKEN_LPAREN || p.token.typ == TOKEN_LBRACE {
			nestedGroup, err := p.parseMolecules()
			if err != nil {
				return molecules, err
			}
			molecules = append(molecules, nestedGroup...)
		}
	}

	if err := p.expect(groupClose); err != nil {
		return molecules, err
	}
	if err := p.nextToken(); err != nil {
		return molecules, err
	}

	// Check for repetition like H2 or (H2O)3
	if p.token.typ == TOKEN_NUMBER {
		count, err := p.parseNumber()
		if err != nil {
			return molecules, err
		}
		for i := 0; i < count; i++ {
			molecules = append(molecules, molecule)
		}
	} else {
		molecules = append(molecules, molecule)
	}

	if p.token.typ == TOKEN_PLUS || p.token.typ == TOKEN_MINUS {
		charge, err := p.parseCharge()
		if err != nil {
			return molecules, err
		}
		for i := range molecules {
			molecules[i].Charge = charge
		}
	}

	return molecules, nil
}

// Parse an element with optional charge
func (p *Parser) parseElements() ([]Element, error) {
	element, exists := ElementTable[p.token.value]
	if !exists {
		element = Element{Symbol: p.token.value} 
	}

	elements := []Element{element}
	if err := p.nextToken(); err != nil {
		return nil, err
	}

	// Optional number as subscript (e.g. H2)
	if p.token.typ == TOKEN_NUMBER {
		count, err := strconv.Atoi(p.token.value)
		if err != nil {
			return nil, err
		}
		for i := 1; i < count; i++ {
			elements = append(elements, element)
		}
		if err := p.nextToken(); err != nil {
			return nil, err
		}
	}

	if p.token.typ == TOKEN_PLUS || p.token.typ == TOKEN_MINUS {
		charge, err := p.parseCharge()
		if err != nil {
			return nil, err
		}
		for i := range elements {
			elements[i].Charge = charge
		}
	}

	return elements, nil
}

// Parse a number
func (p *Parser) parseNumber() (int, error) {
	if p.token.typ != TOKEN_NUMBER {
		return 0, errors.New("expected a number")
	}
	num, err := strconv.Atoi(p.token.value)
	if err != nil {
		return 0, err
	}
	if err := p.nextToken(); err != nil {
		return 0, err
	}
	return num, nil
}

// Parse a charge
func (p *Parser) parseCharge() (int, error) {
	if p.token.typ == TOKEN_PLUS {
		if err := p.nextToken(); err != nil {
			return 0, err
		}
		if p.token.typ == TOKEN_NUMBER {
			charge, _ := p.parseNumber()
			return charge, nil
		}
		return 1, nil
	}
	if p.token.typ == TOKEN_MINUS {
		if err := p.nextToken(); err != nil {
			return 0, err
		}
		if p.token.typ == TOKEN_NUMBER {
			charge, _ := p.parseNumber()
			return -charge, nil
		}
		return -1, nil
	}
	return 0, nil
}
