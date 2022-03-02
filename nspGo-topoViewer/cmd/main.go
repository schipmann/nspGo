package main

import (
	"fmt"
	"math/big"
	"time"

	nspgosession "local.com/nspgo/nspGo-session"
	nsptopoviewer "local.com/nspgo/nspGo-topoViewer"
)

func main() {
	session := nspgosession.Session{}
	session.LoadConfig()
	session.EncodeUserName()
	session.GetRestToken()

	req := nsptopoviewer.Requester{}

	req.GetNetworkIetf(session.IpAdressIprc, session.Token, session.Proxy.Enable, session.Proxy.ProxyAddress)
	ietf := req.ResponseData
	req.GetL3Networks(session.IpAdressIprc, session.Token, session.Proxy.Enable, session.Proxy.ProxyAddress)
	l3 := req.ResponseData
	req.GetServiceLayer(session.IpAdressIprc, session.Token, session.Proxy.Enable, session.Proxy.ProxyAddress)
	sl := req.ResponseData

	// measure runtime of JSON to Struct - Struct to JSON

	start := time.Now()
	r := new(big.Int)
	fmt.Println(r.Binomial(1000, 10))

	process := nsptopoviewer.CyGraph{}
	process.UnmarshalToCyGraph(ietf, l3, sl)
	process.MarshalToCyto()

	elapsed := time.Since(start)
	fmt.Printf("Binomial took %s", elapsed)
	//measure runtime of JSON to Struct - Struct to JSON >
}
