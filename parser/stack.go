package parser

import (
	"errors"
	"fmt"
)

type Stack struct {
	stack []*Token
}

func (s *Stack) Push(tok *Token) {
	s.stack = append(s.stack, tok)
}

func (s *Stack) Pop() (*Token, error) {
	length := len(s.stack)

	if length == 0 {
		return nil, errors.New("Empty stack")
	}

	tok := s.stack[length-1]
	s.stack = s.stack[:length-1]
	return tok, nil
}

func (s *Stack) Peek() (*Token, error) {
	length := len(s.stack)

	if length == 0 {
		return nil, errors.New("Empty stack")
	}

	tok := s.stack[length-1]
	return tok, nil
}

func (s *Stack) String() string {
	return fmt.Sprintf("%v", s.stack)
}
