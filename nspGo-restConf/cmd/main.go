package main

import (
	"fmt"
	"os"
	"sync"
	"time"

	log "github.com/sirupsen/logrus"
	"github.com/tidwall/gjson"

	nspgorestconf "local.com/nspgo/nspGo-restConf"
	nspgosession "local.com/nspgo/nspGo-session"
)

func main() {
	// init class session
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

	// init class RestConf
	rc := nspgorestconf.RestConf{}

	//measure the start time
	start := time.Now()
	log.Info("Start Time: ", start)

	//get list NE
	rc.ReadRestConfPayload("./nspGo-restConf/resconf-inventory-payload.json")
	restconfPayloadInventory := rc.Payload
	inventoryJson := rc.NspRestconfInventory(p.IpAdressIprc, p.Token, p.Proxy.Enable, p.Proxy.ProxyAddress, "operations/nsp-inventory:find", restconfPayloadInventory)
	value := gjson.Get(inventoryJson, "nsp-inventory:output.data.#.ne-id")
	// log.Info(value.String())

	listOfNeId := []string{}
	value.ForEach(func(key, value gjson.Result) bool {
		//println(value.String())
		listOfNeId = append(listOfNeId, value.String())
		return true // keep iterating
	})
	// log.Info("listOfNeId: ", listOfNeId)

	// get payload
	pathToPayload := "./nspGo-restConf/resconf-payload.json"
	file, err := os.Stat(pathToPayload)
	if err != nil {
		return
	}
	log.Info("Payload Size(bytes): ", file.Size())

	rc.ReadRestConfPayload(pathToPayload)
	restconfPayloadCreate := rc.Payload

	// Running Restconf Request In Sequence
	log.Info("Running Sequence: ", len(listOfNeId))
	value.ForEach(func(key, value gjson.Result) bool {
		//println(value.String())
		rc.PatchRestConfMdc(p.IpAdressIprc, p.Token, p.Proxy.Enable, p.Proxy.ProxyAddress, value.String(), "/root", restconfPayloadCreate, false)
		return true // keep iterating
	})

	// Running Restconf Request In Concurrent
	log.Info("Running Concurrency: ", len(listOfNeId))
	var waitingGroupNeList sync.WaitGroup
	waitingGroupNeList.Add(len(listOfNeId))

	for j := 0; j < len(listOfNeId); j++ {
		go func(j int) {
			// fmt.Println(listOfNeId[j])
			rc.PatchRestConfMdc(p.IpAdressIprc, p.Token, p.Proxy.Enable, p.Proxy.ProxyAddress, listOfNeId[j], "/root", restconfPayloadCreate, false)
			waitingGroupNeList.Done()
		}(j)
	}
	waitingGroupNeList.Wait()

	log.Info("Finished")
	log.Info("Elapsed Time: ", time.Since(start))

	p.RevokeRestToken()
}
