package parser

func Shunt(tokens []*Token) *Queue {
	outQueue := &Queue{}
	opStack := &Stack{}

	for _, t := range tokens {
		switch t.TokenType {
		case NUM:
			outQueue.Enqueue(t)
		case FRA:
			outQueue.Enqueue(t)
		case MIX:
			outQueue.Enqueue(t)
		default:
			for {
				peeked, err := opStack.Peek()
				if err != nil {
					break
				} else {
					if peeked.IsHigherPrecedence(t) {
						popped, _ := opStack.Pop()
						outQueue.Enqueue(popped)
					} else {
						break
					}
				}
			}
			opStack.Push(t)
		}
	}

	// move from stack to queue
	for {
		popped, err := opStack.Pop()
		if err != nil {
			break
		}

		outQueue.Enqueue(popped)
	}

	return outQueue
}
