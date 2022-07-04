package rpcobjects

type Num struct {
	A, B int
}

func (n *Num) Add(*res) error {
	*res = n.A * n.B
	return nil
}
