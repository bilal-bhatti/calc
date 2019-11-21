package parser

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStack(t *testing.T) {
	vals := []string{"1", "2", "3"}
	s := Stack{}

	for _, v := range vals {
		s.Push(&Token{Numerator: v, TokenType: NUM})
	}

	v, _ := s.Pop()

	assert.Equal(t, "3", v.Numerator, "Should be 3")

	_, _ = s.Pop()

	v, _ = s.Pop()
	assert.Equal(t, "1", v.Numerator, "Should be 1")

	_, err := s.Pop()

	assert.EqualError(t, err, "Empty stack")
}
