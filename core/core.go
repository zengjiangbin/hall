package core

import "fmt"

type hallCore struct{}

func NewService() *hallCore {
	fmt.Println("new service")
	return new(hallCore)
}
