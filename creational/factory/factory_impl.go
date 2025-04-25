package factory

import "strconv"

type iInt struct{}

func (n iInt) Parse(s string) (any, error) {
	return strconv.Atoi(s)
}

type iFloat struct{}

func (n iFloat) Parse(s string) (any, error) {
	return strconv.ParseFloat(s, 64)
}
