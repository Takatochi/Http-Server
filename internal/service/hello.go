package service

import "strconv"

type Hello struct {
}

func (h *Hello) Hello(name string) string {
	if name == "" {
		name = "World"
	}
	return "Hello " + name
}
func (h *Hello) TypeOf(number string) int {

	typeOf, err := strconv.Atoi(number)
	if err != nil {
		typeOf = 0
	}
	typeOf = typeOf * 3
	return typeOf
}
