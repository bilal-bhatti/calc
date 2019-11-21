package parser

// Parse ...
func Parse(l Listener, exp string) {
	tokens, err := Scan(exp)
	if err != nil {
		panic(err)
	}

	queue := Shunt(tokens)

	for {
		token, err := queue.Dequeue()

		if err != nil {
			break
		}

		switch token.TokenType {
		case NUM:
			l.OnNumber(token)
		case FRA:
			l.OnFraction(token)
		case MIX:
			l.OnMixed(token)
		case MUL:
			l.OnMul()
		case DIV:
			l.OnDiv()
		case ADD:
			l.OnAdd()
		case SUB:
			l.OnSub()
		}
	}
}
