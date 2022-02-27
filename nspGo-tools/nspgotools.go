package nspgotools

import (
	"fmt"
	"io"
	"os"

	"github.com/noirbizarre/gonja"
	"github.com/scrapli/scrapligo/driver/base"
	"github.com/scrapli/scrapligo/netconf"
	"github.com/scrapli/scrapligo/transport"
	log "github.com/sirupsen/logrus"
	"gopkg.in/natefinch/lumberjack.v2"
)

type Tools struct {
	JinjaTemplate      string
	JinjaParameterFill string
	LogFileName        string
	NetconfDriver      *netconf.Driver
}

func (tool *Tools) LoadTemplateJinja(template string, parameter string) {
	tpl, err := gonja.FromString(template)

	if err != nil {
		panic(err)
	}
	// Now you can render the template with the given
	// pongo2.Context how often you want to.
	out, err := tpl.Execute(gonja.Context{"name": "florian"})
	if err != nil {
		panic(err)
	}
	fmt.Println(out)
	// Output: Hello Florian!
}

func (tool *Tools) InitLogger(filePath string) {
	mw := io.MultiWriter(os.Stdout, &lumberjack.Logger{
		Filename:   filePath,
		MaxSize:    10, // megabytes
		MaxBackups: 3,
		MaxAge:     28,   //days
		Compress:   true, // disabled by default
	})
	log.SetOutput(mw)
	log.SetFormatter(&log.TextFormatter{
		DisableColors: false,
		FullTimestamp: true})
}

// follow this guide
// https://github.com/scrapli/scrapligo/blob/v0.1.2/examples/netconf/main.go

func (tool *Tools) NetconfClient(neIp, username string, password string, operation string, payload string) {

	switch operation {
	case "get":
		d, _ := netconf.NewNetconfDriver(
			neIp,
			// base.WithPort(21830),
			base.WithAuthStrictKey(false),
			base.WithAuthUsername(username),
			base.WithAuthPassword(password),
			base.WithTransportType(transport.StandardTransportName),
		)

		err := d.Open()
		if err != nil {
			fmt.Printf("failed to open driver; error: %+v\n", err)
			return
		}
		defer d.Close()

		// filter := `
		// <state xmlns="urn:nokia.com:sros:ns:yang:sr:state">
		// <system><version><version-number/></version></system>
		// </state>`

		// filter := `
		// <state xmlns="urn:nokia.com:sros:ns:yang:sr:state">
		// </state>`

		r, err := d.Get(netconf.WithNetconfFilter(payload))
		if err != nil {
			fmt.Printf("failed to get with filter; error: %+v\n", err)
			return
		}
		fmt.Println(string(r.ChannelInput))
		fmt.Printf("Get Response: %s\n", r.Result)

	case "edit":
		d, _ := netconf.NewNetconfDriver(
			neIp,
			// base.WithPort(21830),
			base.WithAuthStrictKey(false),
			base.WithAuthUsername(username),
			base.WithAuthPassword(password),
			base.WithTransportType(transport.StandardTransportName),
		)

		err := d.Open()
		if err != nil {
			fmt.Printf("failed to open driver; error: %+v\n", err)
			return
		}
		defer d.Close()

		edit := ` 
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
		`

		r, err := d.EditConfig("candidate", edit)
		if err != nil {
			fmt.Printf("failed to edit config; error: %+v\n", err)
			return
		}

		fmt.Println(string(r.ChannelInput))

		fmt.Printf("Edit Config Response: %s\n", r.Result)

		_, _ = d.Commit()
	case "ada":
		fmt.Println("three")
	}

}
