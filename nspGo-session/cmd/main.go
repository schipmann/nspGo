package main

import (
	"fmt"

	nspgosession "local.com/nspgo/nspGo-session"
)

func main() {
	p := nspgosession.Session{}
	p.LoadConfig()

	fmt.Println("nsp.nspOsIP :", p.IpAdressNspOs)
	fmt.Println("nsp.nspIprcIP :", p.IpAdressIprc)
	fmt.Println("nsp.Username :", p.Username)
	fmt.Println("nsp.Password :", p.Password)
	fmt.Println("nsp.linetoken :", p.Token)
	fmt.Println("nsp.proxyEnable :", p.Proxy.Enable)
	fmt.Println("nsp.proxyAddress :", p.Proxy.ProxyAddress)

	p.EncodeUserName()
	fmt.Println(p.EncodeUserName())

	p.GetRestToken()
	fmt.Println("nsp.linetoken_NEW :", p.Token)

	p.RevokeRestToken()
}
