package main

import "github.com/bketelsen/factorish"

func init() {
	factorish.RegisterComponent("todo", todoTmpl)
}

var todoTmpl = `<div>
	<p>Task: {{ .Name }}</p>
	<p>Description: {{ .Description }}</p>
</div>`

type Todo struct {
	*factorish.Component
	Name        string
	Description string
}
