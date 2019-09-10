package myint

type MyType int

func (mi MyType) Divide(n int) MyType{
	return mi / MyType(n)
}
func (mi MyType) Add(n int) MyType{
	return mi + MyType(n)
}
func (mi MyType) Sub(n int) MyType{
	return mi - MyType(n)
}
func (mi MyType) Multiply(n int) MyType{
	return mi * MyType(n)
}