package main

import (
	nspgotools "local.com/nspgo/nspGo-tools"
)

func main() {
	// init class
	//
	t := nspgotools.Tools{}

	template := (`

	{
		"ietf-yang-patch:yang-patch": {
			"patch-id": "as",
			"edit": [{
				"edit-id": "svc-01",
				"operation": "merge",
				"target": "nokia-conf:configure",
				"value": {
					"nokia-conf:configure": {
						"port": [{
								"port-id": "1/1/c1/1",
								"ethernet": [{
									"mode": "hybrid"
								}]
							},
							{
								"port-id": "1/1/c2/1",
								"ethernet": [{
									"mode": "hybrid"
								}]
							}
						],
						"service": {
							"sdp": [{
								"sdp-id": "1",
								"admin-state": "enable",
								"delivery-type": "mpls",
								"ldp": "true",
								"far-end": {
									"ip-address": "99.99.99.1"
								}
							}],
							"vpls": [{% for n in range(1,700) %}{  
								"service-name": "service{{n}}",
								"service-id": "n",
								"description": "<% mplsParam01 %>",
								"customer": "1",
								"spoke-sdp": [{
									"sdp-bind-id": "1:{{n}}",
									"description": "This Is MPLS-Params-01 PlaceHolder-{{n}}-<% mplsParam02 %>"
								}],
								"sap": [{
									"sap-id": "1/1/c2/1:{{n}}",
									"description": "This Is MPLS-Params-02 PlaceHolder-{{n}}-<% pwLabels %>"
								}]
							}{%- if not loop.last -%} , {% endif %}
							{% endfor %}]
						}
					}
				}
			}]
		}
	}


`)

	parameter := (`{"name": "florian"}`)

	t.LoadTemplateJinja(template, parameter)
}
