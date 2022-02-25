package nspgotools

import (
	"fmt"
	"io"
	"os"

	"github.com/noirbizarre/gonja"
	log "github.com/sirupsen/logrus"
	"gopkg.in/natefinch/lumberjack.v2"
)

type Tools struct {
	JinjaTemplate      string
	JinjaParameterFill string
	LogFileName        string
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
		Filename:   "./logs/nspGo-restconf.log",
		MaxSize:    500, // megabytes
		MaxBackups: 3,
		MaxAge:     28,   //days
		Compress:   true, // disabled by default
	})
	log.SetOutput(mw)
	log.SetFormatter(&log.TextFormatter{
		DisableColors: false,
		FullTimestamp: true})
}
