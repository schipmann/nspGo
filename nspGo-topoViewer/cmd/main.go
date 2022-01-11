package main

import (
	"io/ioutil"
	"log"
	"os"

	nsptopoviewer "local.com/nspgo/nspGo-topoViewer"
)

func main() {
	graph1 := nsptopoviewer.CyGraph{}

	content, err := ioutil.ReadFile("../../ietfNetwork.json")
	if err != nil {
		log.Fatal("Error when opening file: ", err)
	}
	filePath, _ := os.Getwd()
	filePath = (filePath + "../../../vis-library/colajs-asad-graph/data-cytoMarshall.json")
	graph1.DumpIetfNetworkToCyGraph(content, nsptopoviewer.IetfNetworkStruct{}, filePath)
}
