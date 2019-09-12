package myint

type MyInt float64

func (mi MyInt) Add(i float64) MyInt {
	return mi / MyInt(i)
}
