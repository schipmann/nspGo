package nspgotools

import (
	"fmt"

	"github.com/noirbizarre/gonja"
)

type Tools struct {
	TemplateJinja string
	ParameterFill string
}

func (tool *Tools) LoadTemplateJinja(template string, parameter string) {

	// Compile the template first (i. e. creating the AST)
	// tpl, err := gonja.FromString(`{
	//       "ietf-yang-patch:yang-patch": {
	//         "patch-id": "{{ name }}",
	//         "edit": [
	//           {
	//             "edit-id": "{{ name }}",
	//             {% for n in range(1,10) %}{{n}}
	//             "operation": "create",
	//             "{% endfor %}"
	//             "target": "/interface=test-yang-patch-{{ name }}",
	//             "value": {
	//               "nokia-conf:interface": [
	//                 {
	//                   "interface-name": "test-yang-patch-{{ name }}"
	//                 }
	//               ]
	//             }
	//           }
	//         ]
	//       }
	//     }`)
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
