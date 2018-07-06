package main

import "github.com/bketelsen/factorish"

var helloTmpl = `<div>
	<p>Hello World! Love {{ .Name }}</p>
</div>`

func init() {
	factorish.RegisterComponent("hello", helloTmpl)
}

type Hello struct {
	Name string
}
