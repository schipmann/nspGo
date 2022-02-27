package main

import (
	"fmt"
	"log"

	"github.com/Juniper/go-netconf/netconf"
	"golang.org/x/crypto/ssh"

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

	getFilter := `
	<state xmlns="urn:nokia.com:sros:ns:yang:sr:state">
	<system><version><version-number/></version></system>
	</state>`

	t.NetconfClient("172.23.160.41", "admin", "NokiaNsp1!", "get", getFilter)

	t.NetconfClient("172.23.160.41", "admin", "NokiaNsp1!", "edit", getFilter)

	// d, _ := netconf.NewNetconfDriver(
	// 	"172.23.160.41",
	// 	// base.WithPort(21830),
	// 	base.WithAuthStrictKey(false),
	// 	base.WithAuthUsername("admin"),
	// 	base.WithAuthPassword("NokiaNsp1!"),
	// 	base.WithTransportType(transport.StandardTransportName),
	// )

	// err := d.Open()
	// if err != nil {
	// 	fmt.Printf("failed to open driver; error: %+v\n", err)
	// 	return
	// }
	// defer d.Close()

	// r, err := d.GetConfig("running")
	// if err != nil {
	// 	fmt.Printf("failed to get config; error: %+v\n", err)
	// 	return
	// }
	// fmt.Printf("Get Config Response:\n%s\n", r.Result)

	// filter := `
	// <state xmlns="urn:nokia.com:sros:ns:yang:sr:state">
	// <system><version><version-number/></version></system>
	// </state>`

	// r, err = d.Get(netconf.WithNetconfFilter(filter))
	// if err != nil {
	// 	fmt.Printf("failed to get with filter; error: %+v\n", err)
	// 	return
	// }

	// fmt.Printf("Get Response: %s\n", r.Result)

	sshConfig := &ssh.ClientConfig{
		User:            "admin",
		Auth:            []ssh.AuthMethod{ssh.Password("NokiaNsp1!")},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}

	s, err := netconf.DialSSH("172.16.240.189", sshConfig)

	if err != nil {
		log.Fatal(err)
	}

	defer s.Close()

	fmt.Println(s.ServerCapabilities)
	fmt.Println(s.SessionID)

	// Sends raw XML
	reply, err := s.Exec(netconf.RawMethod(`
	<?xml version="1.0" encoding="UTF-8"?>
	<rpc xmlns="urn:ietf:params:xml:ns:netconf:base:1.0" message-id="101">
		<edit-config>
			<target>
				<running/>
			</target>
			<configure xmlns="urn:nokia.com:sros:ns:yang:sr:conf">
				<service>
					<sdp>
						<sdp-id>1</sdp-id>
						<admin-state>enable</admin-state>
						<delivery-type>mpls</delivery-type>
						<ldp>true</ldp>
						<far-end>
							<ip-address>99.99.99.1</ip-address>
						</far-end>
					</sdp>
				</service>
			</configure>
		</edit-config>
	</rpc>`))
	if err != nil {
		panic(err)
	}
	fmt.Printf("Reply: %+v", reply)

}
