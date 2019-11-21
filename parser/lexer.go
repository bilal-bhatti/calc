package parser

import (
	"bufio"
	"errors"
	"fmt"
	"regexp"
	"strings"
)

type TokenType int

const (
	NUM TokenType = iota
	FRA
	MIX
	MUL
	DIV
	ADD
	SUB
)

var precedence = map[TokenType]int{
	MUL: 2,
	DIV: 2,
	ADD: 1,
	SUB: 1,
}

type Token struct {
	TokenType                     TokenType
	Whole, Numerator, Denominator string
}

func (t *Token) IsHigherPrecedence(other *Token) bool {
	return precedence[t.TokenType] >= precedence[other.TokenType]
}

func (t *Token) String() string {
	if t.TokenType > MIX {
		return fmt.Sprintf("{%d}", t.TokenType)
	}
	return fmt.Sprintf("{%s, %s, %s}", t.Whole, t.Numerator, t.Denominator)
}

func Scan(input string) ([]*Token, error) {
	var number = regexp.MustCompile(`^[0-9]+$`)
	var fraction = regexp.MustCompile(`^(?P<n>[0-9]+)/(?P<d>[0-9]+)?$`)
	var mixed = regexp.MustCompile(`^(?P<w>[0-9]+)_(?P<n>[0-9]+)/(?P<d>[0-9]+)$`)
	var op = regexp.MustCompile(`^[*\-+/]$`)

	var tokens = make([]*Token, 0)

	scanner := bufio.NewScanner(strings.NewReader(input))
	scanner.Split(bufio.ScanWords)

	for scanner.Scan() {
		value := strings.TrimSpace(scanner.Text())

		if len(value) > 0 {
			// match whole numbers
			if number.MatchString(value) {
				tokens = append(tokens, &Token{
					TokenType: NUM,
					Numerator: value,
				})
				continue
			}

			// Match fractions
			if fraction.MatchString(value) {
				match := fraction.FindStringSubmatch(value)

				result := make(map[string]string)
				for i, name := range fraction.SubexpNames() {
					if i != 0 && name != "" {
						result[name] = match[i]
					}
				}

				tokens = append(tokens, &Token{
					TokenType:   FRA,
					Numerator:   result["n"],
					Denominator: result["d"],
				})
				continue
			}

			// Match mixed fraction
			if mixed.MatchString(value) {
				match := mixed.FindStringSubmatch(value)

				result := make(map[string]string)
				for i, name := range mixed.SubexpNames() {
					if i != 0 && name != "" {
						result[name] = match[i]
					}
				}

				tokens = append(tokens, &Token{
					TokenType:   MIX,
					Whole:       result["w"],
					Numerator:   result["n"],
					Denominator: result["d"],
				})
				continue
			}

			if op.MatchString(value) {
				switch value {
				case "+":
					tokens = append(tokens, &Token{TokenType: ADD})
				case "-":
					tokens = append(tokens, &Token{TokenType: SUB})
				case "*":
					tokens = append(tokens, &Token{TokenType: MUL})
				case "/":
					tokens = append(tokens, &Token{TokenType: DIV})
				}

				continue
			}

			return nil, errors.New("Invalid input: " + value)
		}
	}
	return tokens, nil
}
