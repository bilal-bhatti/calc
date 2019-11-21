package parser

type Listener interface {
	OnNumber(tok *Token)
	OnFraction(tok *Token)
	OnMixed(tok *Token)
	OnMul()
	OnDiv()
	OnAdd()
	OnSub()
}
