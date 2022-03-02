package nsptopoviewer

import (
	"crypto/tls"
	"io/ioutil"
	"time"

	nspgoconstants "local.com/nspgo/nspGo-constants"

	"github.com/go-resty/resty/v2"
	log "github.com/sirupsen/logrus"
	"github.com/tidwall/pretty"
)

type Requester struct {
	ResponseData []byte
}

func (ipO *Requester) GetNetworkIetf(urlHost string, token string, proxyEnable string, proxyAddress string) (result []byte) {
	client := resty.New()
	client.SetTimeout(60 * time.Second)
	client.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	if proxyEnable == "true" {
		client.SetProxy(proxyAddress)
	}

	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("authorization", "Bearer "+token).
		SetBody(`{ "grant_type": "client_credentials" }`).
		Get("https://" + urlHost + nspgoconstants.GLBL_NSP_IP_OPTIM_BASE_URL + "ietf/te/networks")

	log.Info("Get IETFNetwork is success: ", resp.IsSuccess(), ".")

	ipO.ResponseData = resp.Body()
	result = pretty.Pretty(resp.Body())

	if err != nil {
		log.Error("Get ietf network is unsuccesful: ", err)
		return
	}

	ioutil.WriteFile("ietfNetworks.json", result, 0644)
	return
}

func (req *Requester) GetServiceLayer(urlHost string, token string, proxyEnable string, proxyAddress string) (result []byte) {
	client := resty.New()
	client.SetTimeout(60 * time.Second)
	client.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	if proxyEnable == "true" {
		client.SetProxy(proxyAddress)
	}

	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("authorization", "Bearer "+token).
		SetBody(`{ "grant_type": "client_credentials" }`).
		Get("https://" + urlHost + nspgoconstants.GLBL_NSP_SERVICE_SESSION)

	log.Info("Get Service Layer is success: ", resp.IsSuccess(), ".")

	req.ResponseData = resp.Body()
	if err != nil {
		log.Error("Get service Layer is unsuccesful: ", err)
		return
	}

	ioutil.WriteFile("serviceLayer.json", result, 0644)
	return
}

func (req *Requester) GetL3Networks(urlHost string, token string, proxyEnable string, proxyAddress string) (result []byte) {
	client := resty.New()
	client.SetTimeout(60 * time.Second)
	client.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	if proxyEnable == "true" {
		client.SetProxy(proxyAddress)
	}

	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("authorization", "Bearer "+token).
		SetBody(`{ "grant_type": "client_credentials" }`).
		Get("https://" + urlHost + nspgoconstants.GLBL_NSP_IP_OPTIM_BASE_URL + "nsp/net/l3/networks")

	log.Info("Get L3 network is success: ", resp.IsSuccess(), ".")

	req.ResponseData = resp.Body()
	result = pretty.Pretty(resp.Body())

	if err != nil {
		log.Error("Get L3 network is unsuccesful: ", err, ".")
		return
	}

	ioutil.WriteFile("l3Network.json", result, 0644)

	return
}
