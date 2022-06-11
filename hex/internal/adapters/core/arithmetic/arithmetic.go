package arithmetic

type Adapter struct {
}

func NewAdapter() *Adapter {
	return &Adapter{}
}

func (arith Adapter) Addition(a int32, b int32) (int32, error) {
	return a + b, nil
}

func (arith Adapter) Substraction(a int32, b int32) (int32, error) {
	return a - b, nil
}

func (arith Adapter) Multipulication(a int32, b int32) (int32, error) {
	return a * b, nil
}

func (arith Adapter) Division(a int32, b int32) (int32, error) {
	return a / b, nil
}
