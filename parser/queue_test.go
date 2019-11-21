package parser

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestQueue(t *testing.T) {
	vals := []string{"1", "2", "3"}
	s := Queue{}

	for _, v := range vals {
		s.Enqueue(&Token{Numerator: v, TokenType: NUM})
	}

	v, _ := s.Dequeue()

	assert.Equal(t, "1", v.Numerator, "Should be 1")

	v, _ = s.Dequeue()
	v, _ = s.Dequeue()

	assert.Equal(t, "3", v.Numerator, "Should be 3")

	_, err := s.Dequeue()

	assert.EqualError(t, err, "Empty queue")
}
