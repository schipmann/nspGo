package main

import (
	nspgotools "local.com/nspgo/nspGo-tools"
)

func main() {
	// init class
	//
	t := nspgotools.Tools{}

	// template := (`{
	//       "ietf-yang-patch:yang-patch": {
	//         "patch-id": "{{ name }}",
	//         "edit": [
	//           {
	//             "edit-id": "as",
	//             "operation": "create",
	//             "target": "/interface=test-yang-patch-{{ name }}",
	//             "value": {
	//               "nokia-conf:interface": [
	//                 {
	//                   "interface-name": "test-yang-patch-27"
	//                 }
	//               ]
	//             }
	//           }
	//         ]
	//       }
	//     }`)

	template := (` {
    "ietf-yang-patch:yang-patch": {
      "patch-id": "as",
      "edit": [
        {% for n in range(1,4091) %}
        {
          "edit-id": "svc-{{n}}",
          "operation": "merge",
          "target":"nokia-conf:/configure/service/epipe=00{{n}}",
          "value":{"nokia-conf:epipe": {"service-name": "00{{n}}", "service-id": "00{{n}}", "customer": "1"}}
        },
        {
          "edit-id": "sap01-{{n}}",
          "operation": "merge",
          "target":"nokia-conf:/configure/service/epipe=00{{n}}/sap=1%2F1%2Fc1%2F1%3A{{n}}",
          "value":{"nokia-conf:sap": {"sap-id": "1/1/c1/1:{{n}}"}}
        },
        {
          "edit-id": "sap02-{{n}}",
          "operation": "merge",
          "target":"nokia-conf:/configure/service/epipe=00{{n}}/sap=1%2F1%2Fc2%2F1%3A{{n}}",
          "value":{"nokia-conf:sap": {"sap-id": "1/1/c2/1:{{n}}"}}
        },{% endfor %}
      ]
    }
  }`)

	parameter := (`{"name": "florian"}`)

	t.LoadTemplateJinja(template, parameter)

}
