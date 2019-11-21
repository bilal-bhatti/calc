package calc

import (
	"calc/parser"
	"errors"
	"fmt"
)

type CalcStack struct {
	stack []*Numby
}

func (s *CalcStack) Push(i *Numby) {
	s.stack = append(s.stack, i)
}

func (s *CalcStack) Pop() (*Numby, error) {
	length := len(s.stack)

	if length == 0 {
		return nil, errors.New("Empty stack")
	}

	tok := s.stack[length-1]
	s.stack = s.stack[:length-1]
	return tok, nil
}

func (s *CalcStack) String() string {
	return fmt.Sprintf("%v", s.stack)
}

func (s *CalcStack) OnNumber(tok *parser.Token) {
	s.Push(NewFromNumber(tok.Numerator))
}

func (s *CalcStack) OnFraction(tok *parser.Token) {
	s.Push(NewFromFraction(tok.Numerator, tok.Denominator))
}

func (s *CalcStack) OnMixed(tok *parser.Token) {
	s.Push(NewFromMixed(tok.Whole, tok.Numerator, tok.Denominator))
}

func (s *CalcStack) OnMul() {
	op2, _ := s.Pop()
	op1, _ := s.Pop()

	s.Push(op1.Mul(op2))
}

func (s *CalcStack) OnDiv() {
	op2, _ := s.Pop()
	op1, _ := s.Pop()

	s.Push(op1.Div(op2))
}

func (s *CalcStack) OnAdd() {
	op2, _ := s.Pop()
	op1, _ := s.Pop()

	s.Push(op1.Add(op2))
}

func (s *CalcStack) OnSub() {
	op2, _ := s.Pop()
	op1, _ := s.Pop()

	s.Push(op1.Sub(op2))
}
