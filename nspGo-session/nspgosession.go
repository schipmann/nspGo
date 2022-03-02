package nspgosession

import (
	"crypto/tls"
	"encoding/base64"
	"strings"
	"time"

	nspgoconstants "local.com/nspgo/nspGo-constants"

	"github.com/buger/jsonparser"
	"github.com/go-resty/resty/v2"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

type Session struct {
	IpAdressNspOs string
	IpAdressIprc  string
	Username      string
	Password      string
	Base64Str     string
	Token         string
	Proxy         Proxy
}

type Proxy struct {
	Enable       string
	ProxyAddress string
}

func (s *Session) LoadConfig() {
	log.SetFormatter(&log.TextFormatter{
		ForceColors:   true,
		FullTimestamp: true,
	})

	/*viper.SetConfigName("nspConfig") // config file name without extension
	viper.SetConfigType("yaml")
	//viper.AddConfigPath(".")
	//viper.AddConfigPath("./nspGo-session/cmd") // config file path
	viper.AutomaticEnv() // read value ENV variable

	err := viper.ReadInConfig()
	if err != nil {
		fmt.Println("fatal error config file: default \n", err)
		//os.Exit(1)
	}
	// Set default value
	viper.SetDefault("nsp.linetoken", "DefaultLineTokenValue")

	// pass to struct data
	//env := viper.GetString("nsp.env")
	s.IpAdressNspOs = viper.GetString("nsp.nspOsIP")
	s.IpAdressIprc = viper.GetString("nsp.nspIprcIP")
	s.Username = viper.GetString("nsp.username")
	s.Password = viper.GetString("nsp.Password")
	s.Token = viper.GetString("nsp.linetoken")
	s.Proxy.Enable = viper.GetString("nsp.proxy.enable")
	s.Proxy.ProxyAddress = viper.GetString("nsp.proxy.proxyAddress")*/

	s.IpAdressNspOs = "147.75.102.178"
	s.IpAdressIprc = "147.75.102.178"
	s.Username = "admin"
	s.Password = "zkw3VSs6"
	s.Token = viper.GetString("nsp.linetoken")
	s.Proxy.Enable = "false"

	log.Info("Session loaded Succesfully.")
}

func (s *Session) EncodeUserName() string {
	var plainCredentials strings.Builder
	plainCredentials.WriteString(s.Username)
	plainCredentials.WriteString(":")
	plainCredentials.WriteString(s.Password)

	var base64Credentials string = base64.StdEncoding.EncodeToString([]byte(plainCredentials.String()))

	s.Base64Str = base64Credentials
	return s.Base64Str
}

func (s *Session) GetRestToken() {
	client := resty.New()
	client.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	if s.Proxy.Enable == "true" {
		client.SetProxy(s.Proxy.ProxyAddress)
	}

	// POST JSON string
	// No need to set content type, if you have client level setting
	//
	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("authorization", "Basic "+s.Base64Str).
		SetBody(`{ "grant_type": "client_credentials" }`).
		Post("https://" + s.IpAdressNspOs + nspgoconstants.GLBL_NSP_SESSION_URL_TOKEN)

	log.Info("Get token is success: ", resp.IsSuccess(), ".")

	if err != nil {
		log.Warn("token revoke not succesful" + err.Error())
	}

	s.Token, _ = jsonparser.GetString([]byte(string(resp.Body())), "access_token")
}

func (s *Session) RevokeRestToken() {
	client := resty.New()
	client.SetTimeout(20 * time.Second)
	client.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	if s.Proxy.Enable == "true" {
		client.SetProxy(s.Proxy.ProxyAddress)
	}

	// POST JSON string
	// No need to set content type, if you have client level setting
	//
	resp, err := client.R().
		SetHeader("Content-Type", "application/x-www-form-urlencoded").
		SetHeader("authorization", "Basic "+s.Base64Str).
		SetBody("token=" + s.Base64Str + "&token_type_hint=token").
		Post("https://" + s.IpAdressNspOs + nspgoconstants.GLBL_NSP_SESSION_URL_TOKEN_REVOKE)

	log.Info("revoke token is success: ", resp.IsSuccess())

	if err != nil {
		log.Error("token revoke not succesful", err)
		return
	}
}
