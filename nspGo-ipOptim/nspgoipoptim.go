package nspgoipoptim

import (
	"crypto/tls"
	"time"

	"github.com/go-resty/resty/v2"
	log "github.com/sirupsen/logrus"
	"github.com/tidwall/pretty"
	nspgoconstants "local.com/nspgo/nspGo-constants"
)

type IpOptim struct {
	PathProfileData string // plachoder for yaml path-profile-data
	ResponseData    []byte
}

func (ipO *IpOptim) GetNetworkIetf(urlHost string, token string, proxyEnable string, proxyAddress string) (result string) {
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

	log.Info("get ietf network is success: ", resp.IsSuccess())
	log.Debug("get ietf network debug ", resp)

	//ipO.ResponseData = resp.Body()
	//result = string(pretty.Pretty(resp.Body()))

	if err != nil {
		log.Error("get ietf network is unsuccesful: ", err)
		return
	}
	return string(pretty.Pretty(resp.Body()))
}
