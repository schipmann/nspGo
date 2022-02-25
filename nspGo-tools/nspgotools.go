package nspgotools

import (
	"fmt"

	"github.com/noirbizarre/gonja"
)

type Tools struct {
	JinjaTemplate      string
	JinjaParameterFill string
}

func (tool *Tools) LoadTemplateJinja(template string, parameter string) {
	tpl, err := gonja.FromString(template)

	if err != nil {
		panic(err)
	}
	// Now you can render the template with the given
	// pongo2.Context how often you want to.
	out, err := tpl.Execute(gonja.Context{"name": "florian"})
	if err != nil {
		panic(err)
	}
	fmt.Println(out)
	// Output: Hello Florian!
}
