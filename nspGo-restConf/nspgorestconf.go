package nspgorestconf

import (
	"crypto/tls"
	"io/ioutil"
	"time"

	"github.com/go-resty/resty/v2"
	log "github.com/sirupsen/logrus"
	nspgoconstants "local.com/nspgo/nspGo-constants"
)

type RestConf struct {
	Payload      []byte
	ResponseData []byte
}

func (rConf *RestConf) ReadRestConfPayload(file string) {
	body, err := ioutil.ReadFile(file)
	if err != nil {
		log.Fatalf("unable to read file: %v", err)
	}
	rConf.Payload = body
	//fmt.Println(string(body))
}

func (rConf *RestConf) PatchRestConfNetworkDevice(urlHost string, token string, proxyEnable string, proxyAddress string, neId string, xPath string, payload []byte, asycn bool) (result string) {
	client := resty.New()
	client.SetTimeout(60 * time.Second)
	client.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	if proxyEnable == "true" {
		client.SetProxy(proxyAddress)
	}

	if asycn == true {
		url := ("https://" + urlHost + nspgoconstants.GLBL_NSP_RESTCONF_BASE_URL + "data/network-device-mgr:network-devices/network-device=" + neId + xPath + "?async=true")
		resp, err := client.R().
			SetHeader("Accept", "application/yang-data+json").
			SetHeader("Content-Type", "application/yang-patch+json").
			SetHeader("authorization", "Bearer "+token).
			SetBody(payload).
			Patch(url)

		// log.Info("PatchRestConfNetworkDevice is success: ", resp.IsSuccess())
		// log.Debug("PatchRestConfNetworkDevice debug ", resp)
		log.Info("PatchRestConfNetworkDevice Response: ", resp.String())
		//fmt.Println(result)
		if err != nil {
			log.Error("PatchRestConfNetworkDevice is unsuccesful: ", err)
			return
		}
		return resp.String()
	} else { //asycn == false
		url := ("https://" + urlHost + nspgoconstants.GLBL_NSP_RESTCONF_BASE_URL + "data/network-device-mgr:network-devices/network-device=" + neId + xPath)
		resp, err := client.R().
			SetHeader("Accept", "application/yang-data+json").
			SetHeader("Content-Type", "application/yang-patch+json").
			SetHeader("authorization", "Bearer "+token).
			SetBody(payload).
			Patch(url)

		// log.Info("PatchRestConfNetworkDevice is success: ", resp.IsSuccess())
		// log.Debug("PatchRestConfNetworkDevice debug ", resp)
		log.Info("PatchRestConfNetworkDevice Response: ", resp.String())
		//fmt.Println(result)
		if err != nil {
			log.Error("PatchRestConfNetworkDevice is unsuccesful: ", err)
			return
		}
	}
	return
}
