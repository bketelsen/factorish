package main

import (
	"fmt"

	"github.com/bketelsen/factorish"
)

func init() {

}
func main() {

	data := map[string]interface{}{
		"Name":        "Bills",
		"Description": "Pay the Mortgage",
	}
	todoComponent, err := factorish.ComponentByName("todo")
	if err != nil {
		panic(err)
	}
	tree, err := todoComponent.Render(data)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(tree.HTML()))

	todo := &Todo{
		Name:        "Lawn",
		Description: "Mow the Lawn",
	}
	tree2, err := todoComponent.Render(todo)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(tree2.HTML()))

	hello := &Hello{
		Name: "Lawn",
	}
	helloComponent, err := factorish.ComponentByName("hello")
	if err != nil {
		panic(err)
	}
	tree3, err := helloComponent.Render(hello)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(tree3.HTML()))
}
