package main

import (
	nspgotools "local.com/nspgo/nspGo-tools"
)

func main() {
	// init class
	//
	t := nspgotools.Tools{}

	// 	template := (`{
	//     "ietf-yang-patch:yang-patch": {
	//         "patch-id": "as",
	//         "edit": [
	//             {
	//                 "edit-id": "port-01",
	//                 "operation": "merge",
	//                 "target": "nokia-conf:configure/port",
	//                 "value": {
	//                     "nokia-conf:port": [
	//                         {
	//                         "port-id": "1/1/c1/1",
	//                         "ethernet": {
	//                             "mode": "hybrid"
	//                         }
	//                     },
	//                     {
	//                         "port-id": "1/1/c2/1",
	//                         "ethernet": {
	//                             "mode": "hybrid"
	//                         }
	//                     }
	//                     ]
	//                 }
	//             },
	//             {
	//                 "edit-id": "svc-01",
	//                 "operation": "merge",
	//                 "target": "nokia-conf:configure/service",
	//                 "value": {
	//                     "nokia-conf:service": {
	//                         "epipe": [
	//                           {% for n in range(1,4091) %}
	//                             {
	//                                 "service-name": "{{n}}",
	//                                 "service-id": "{{n}}",
	//                                 "customer": "1",
	//                                 "sap": [
	//                                     {
	//                                         "sap-id": "1/1/c1/1:{{n}}"
	//                                     },
	//                                     {
	//                                         "sap-id": "1/1/c2/1:{{n}}"
	//                                     }
	//                                 ]
	//                             },
	//                           {% endfor %}
	//                         ]
	//                     }
	//                 }
	//             }
	//         ]
	//     }
	// }`)

	template := (`
	{
		"ietf-yang-patch:yang-patch": {
		  "patch-id": "as",
		  "edit": [
			{
			  "edit-id": "svc-01",
			  "operation": "merge",
			  "target": "nokia-conf:configure",
			  "value": {
				"nokia-conf:configure": {
				  "port": [
					{
					  "port-id": "1/1/c1/1",
					  "ethernet": [
						{
						  "mode": "hybrid"
						}
					  ]
					},
					{
					  "port-id": "1/1/c2/1",
					  "ethernet": [
						{
						  "mode": "hybrid"
						}
					  ]
					}
				  ],
				  "service": {
					"sdp": [
					  {
						"sdp-id": "1",
						"admin-state": "enable",
						"delivery-type": "mpls",
						"ldp": "true",
						"far-end": {
						  "ip-address": "99.99.99.1"
						}
					  }
					],
					"vpls":[
					   {% for n in range(1,4090) %}
					   {
						"service-name": "service{{n}}",
						"service-id": {{n}},
						"description": "{{mplsParam02}}",
						"customer": "1",
						"spoke-sdp": [
						  {
							"sdp-bind-id": "1:{{n}}",
							"description": "This Is MPLS-Params-01 PlaceHolder-{{n}}-{{mplsParam01}}"
						  }
						],
						"sap": [
						  {
							"sap-id": "1/1/c2/1:{{n}}",
							"description": "This Is MPLS-Params-02 PlaceHolder-{{n}}-{{pwLabels}}"
						  }
						]
					  },
					{% endfor %}
					]
				  }
				}
			  }
			}
		  ]
		}
	  }`)

	parameter := (`{"name": "florian"}`)

	t.LoadTemplateJinja(template, parameter)
}
