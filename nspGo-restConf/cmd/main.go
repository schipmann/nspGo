package main

import (
	"fmt"
	"strconv"
	"sync"
	"time"

	log "github.com/sirupsen/logrus"

	nspgorestconf "local.com/nspgo/nspGo-restConf"
	nspgosession "local.com/nspgo/nspGo-session"
)

func main() {
	// init class session
	//
	p := nspgosession.Session{}
	p.LoadConfig()

	fmt.Println("nsp.nspOsIP :", p.IpAdressNspOs)
	fmt.Println("nsp.nspIprcIP :", p.IpAdressIprc)
	fmt.Println("nsp.Username :", p.Username)
	fmt.Println("nsp.Password :", p.Password)
	fmt.Println("nsp.linetoken :", p.Token)

	p.EncodeUserName()
	fmt.Println(p.EncodeUserName())

	p.GetRestToken()
	fmt.Println("nsp.linetoken_NEW :", p.Token)

	//neId := "10.6.31.98"

	// init class RestConf
	rc := nspgorestconf.RestConf{}
	//rc.PatchRestConfNetworkDevice(p.IpAdressIprc, p.Token, p.Proxy.Enable, p.Proxy.ProxyAddress, neId, xPath)

	rc.ReadRestConfPayload("./nspGo-restConf/resconf-payload.json")
	restconfPayload := rc.Payload
	// fmt.Println("Print PayloadByte")
	// fmt.Println(restconfPayload)

	// fmt.Println("Print PayloadString")
	// fmt.Println(string(restconfPayload))

	//RestConf xPath
	// xPath := "/root/nokia-conf:configure/router=Base"
	//xPath := "/root/nokia-conf:configure"

	// RestConf Payload // This payload is for MDC
	// 10 mb satelite
	// 155
	// two size of payload.

	// payload := (`{
	// 	"ietf-yang-patch:yang-patch": {
	// 	  "patch-id": "as",
	// 	  "edit": [
	// 		{
	// 		  "edit-id": "as",
	// 		  "operation": "create",
	// 		  "target": "/interface=test-yang-patch-27",
	// 		  "value": {
	// 			"nokia-conf:interface": [
	// 			  {
	// 				"interface-name": "test-yang-patch-27"
	// 			  }
	// 			]
	// 		  }
	// 		}
	// 	  ]
	// 	}
	//   }`)

	//measure the start time

	start := time.Now()
	log.Info("Start Time: ", start)

	// subnet 10.6.30.0 has 185 Nodes
	var wg_30 sync.WaitGroup
	wg_30.Add(185)

	for i := 1; i <= 185; i++ {
		go func(i int) {
			//fmt.Println("10.6.30." + strconv.Itoa(i))
			neId := "10.6.30." + strconv.Itoa(i)

			rc.PatchRestConfNetworkDevice(p.IpAdressIprc, p.Token, p.Proxy.Enable, p.Proxy.ProxyAddress, neId, "/root", restconfPayload, true)

			wg_30.Done()
		}(i)
	}
	wg_30.Wait()

	// subnet 10.6.31.0 has 138 Nodes
	// var wg_31 sync.WaitGroup
	// wg_31.Add(138)

	// for j := 1; j <= 138; j++ {
	// 	go func(j int) {
	// 		//fmt.Println("10.6.30." + strconv.Itoa(j))
	// 		neId := "10.6.31." + strconv.Itoa(j)

	// 		rc.PatchRestConfNetworkDevice(p.IpAdressIprc, p.Token, p.Proxy.Enable, p.Proxy.ProxyAddress, neId, "/root", restconfPayload, true)

	// 		wg_31.Done()
	// 	}(j)
	// }
	// wg_31.Wait()

	// log.Info("Finished")
	// log.Info("Elapsed Time: ", time.Since(start))

	p.RevokeRestToken()

}
