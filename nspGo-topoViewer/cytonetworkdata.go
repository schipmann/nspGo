package nsptopoviewer

import (
	"encoding/json"
	"io/ioutil"
	"reflect"
	"strconv"

	log "github.com/sirupsen/logrus"

	"github.com/buger/jsonparser"
)

type CyGraph struct {
	Networks     []networkStruct `json:"network"`
	ServiceLayer l3NetworkStruct `json:"servicelayer"`
	GraphNodes   []cyNode        `json:"nodes"`
	GraphLinks   []cyLink        `json:"links"`
}

type networkStruct struct {
	Key         string `json:"key`
	ID          string `json:"id"`
	Area        string `json:"area"`
	IsisLevel   string `json:"isislevel"`
	DisplayName string `json:"displayname"`
	Protocol    string `json:"protocol"`
}

type l3NetworkStruct struct {
	ServiceName string `json:"servicename"`
	SiteA       string `json:"sitea"`
	SiteB       string `json:"siteb"`
}

type cyNode struct {
	Data struct {
		ID            string `json:"id"`
		IdInt         string `json:"idInt"`
		Name          string `json:"name"`
		SystemAddress string `json:"systemAddress"`
		RouterType    string `json:"routerType"`
		Network       string `json:"network"`
	} `json:"data"`
	Group      string   `json:"group"`
	Pos        position `json:"position"`
	Removed    bool     `json:"removed"`
	Selected   bool     `json:"selected"`
	Selectable bool     `json:"selectable"`
	Locked     bool     `json:"locked"`
	Grabbed    bool     `json:"grabbed"`
	Grabbable  bool     `json:"grabbable"`
	Classes    string   `json:"classes"`
}

type cyLink struct {
	Data struct {
		ID     string `json:"id"`
		Source string `json:"source,omitempty"`
		Target string `json:"target,omitempty"`
		Weight string `json:"weight"`
		Group  string `json:"group"`
		Name   string `json:"name"`
	} `json:"data"`
	Pos        position `json:"position"`
	Removed    bool     `json:"removed"`
	Group      string   `json:"group"`
	Selected   bool     `json:"selected"`
	Selectable bool     `json:"selectable"`
	Locked     bool     `json:"locked"`
	Grabbed    bool     `json:"grabbed"`
	Grabbable  bool     `json:"grabbable"`
	Classes    string   `json:"classes"`
}

type position struct {
	X float64 `json:"x"`
	Y float64 `json:"y"`
}

func (cyGraph *CyGraph) AppendCytoNode(id string, idInt string, hostName string, systemAddress string, routerType string, network string, classes string) (result CyGraph) {
	var cyNode = cyNode{
		Data: struct {
			ID            string "json:\"id\""
			IdInt         string "json:\"idInt\""
			Name          string "json:\"name\""
			SystemAddress string "json:\"systemAddress\""
			RouterType    string "json:\"routerType\""
			Network       string "json:\"network\""
		}{ID: id, IdInt: idInt, Name: hostName, SystemAddress: systemAddress, RouterType: routerType, Network: network},
		Group:      "nodes",
		Removed:    false,
		Selected:   false,
		Selectable: true,
		Locked:     false,
		Grabbed:    false,
		Grabbable:  true,
		Classes:    classes,
	}
	if !cyGraph.checkForDouble(cyNode) {
		cyGraph.GraphNodes = append(cyGraph.GraphNodes, cyNode)
	}
	return
}

func (cyGraph *CyGraph) AppendCytoLink(id string, source string, target string, weight string, group string, name string) (result CyGraph) {
	var cyLink = cyLink{
		Data: struct {
			ID     string "json:\"id\""
			Source string `json:"source,omitempty"`
			Target string `json:"target,omitempty"`
			Weight string `json:"weight"`
			Group  string `json:"group"`
			Name   string `json:"name"`
		}{ID: id, Source: source, Target: target, Weight: weight, Group: group, Name: name},
		Group:      "edges",
		Removed:    false,
		Selected:   false,
		Selectable: true,
		Locked:     false,
		Grabbed:    false,
		Grabbable:  true,
		Classes:    "asad-henry",
	}
	cyGraph.GraphLinks = append(cyGraph.GraphLinks, cyLink)
	return
}

//throws JSON raw into Struct "RequestStruct" then selects the needed Data for the CytoScape Graph and saves it in the Struct "CyGraph"
func (cyGraph *CyGraph) UnmarshalToCyGraph(ietfNetwork, networkNames, serviceLayer []byte) {
	var basket IETFNetworkStruct

	error2 := json.Unmarshal(ietfNetwork, &basket)
	if error2 != nil {
		log.Error(error2)
	}

	m := make(map[string]string)

	jsonPath := &basket.Response.Data.Network //to increase readablity, pointer makes sure there is no dublication in the memory

	for index, networks := range *jsonPath {
		for nindex, node := range networks.Node {
			nodeId := node.NodeID
			sourceNodeName := node.TeNodeAugment.Te.TeNodeID.DottedQuad.String
			m[nodeId] = sourceNodeName
			cyGraph.AppendCytoNode(sourceNodeName, strconv.Itoa(index)+"-"+strconv.Itoa(nindex), sourceNodeName, sourceNodeName, "7750", networks.NetworkID, "classes string") //strconv.Itoa formats Int64 to String
		}
		for lindex, link := range networks.Link {
			source := link.Source.SourceNode
			target := link.Destination.DestNode
			if m[source] != "" && m[target] != "" {
				cyGraph.AppendCytoLink(strconv.Itoa(index)+"-"+strconv.Itoa(lindex), m[source], m[target], "2", "ip", "link-"+m[source]+"-"+m[target])
			}
		}
	}

	log.Info("Anzahl der Elemente in GraphNodes: ", len(cyGraph.GraphNodes))
	log.Info("Anzahl der Elemente in GraphLinks: ", len(cyGraph.GraphLinks))

	cyGraph.addNetworkNames(networkNames)
	cyGraph.addServiceLayer(serviceLayer)
}

//Because the NetworkName API Returns the JSON with the Networks as Keys, it has to be possible to select the keys without knowing them,
//thats what the first jsonParser iteration does, the second one is only to go one level deeper (maybe there is a prettier way of doing this)
func (network *CyGraph) addNetworkNames(networkNames []byte) {
	jsonparser.ObjectEach(networkNames, func(mainkey []byte, value []byte, dataType jsonparser.ValueType, offset int) error {
		dataHelper, _, _, _ := jsonparser.Get(value, "nspNetwork", "networkState", "networkParams", "area")
		jsonparser.ObjectEach(dataHelper, func(key []byte, value []byte, dataType jsonparser.ValueType, offset int) error {
			network.appendToNetwork(mainkey, returnHelper(value, "areaId"), returnHelper(value, "area"), returnHelper(value, "isisLevel"), returnHelper(value, "displayName"), returnHelper(value, "protocol"))
			return nil
		})
		return nil
	}, "response", "data", "network")
}

func (cyGraph *CyGraph) appendToNetwork(key, areaId, area, isisLevel, displayName, protocol []byte) {
	var network = networkStruct{Key: string(key), ID: string(areaId), Area: string(area), IsisLevel: string(isisLevel), DisplayName: string(displayName), Protocol: string(protocol)}
	cyGraph.Networks = append(cyGraph.Networks, network)
}

func returnHelper(yes []byte, param string) []byte {
	result, _, _, _ := jsonparser.Get(yes, param)
	return result
}

func (sL *CyGraph) addServiceLayer(serviceLayer []byte) {
	var basket ServiceLayerStruct

	err := json.Unmarshal(serviceLayer, &basket)
	if err != nil {
		panic(err)
	}

	sL.ServiceLayer.ServiceName = string(basket.NspServiceIntentIntent[0].ServiceName)
	sL.ServiceLayer.SiteA = string(basket.NspServiceIntentIntent[0].IntentSpecificData.EvpnEpipeEvpnEpipe.SiteA.DeviceID)
	sL.ServiceLayer.SiteB = string(basket.NspServiceIntentIntent[0].IntentSpecificData.EvpnEpipeEvpnEpipe.SiteB.DeviceID)
}

//Marshals the Struct "CytoGraph" to a JSON-File that Cytoscape can handle it
func (cyGraph *CyGraph) MarshalToCyto() {

	file, err := json.MarshalIndent(cyGraph, "", "")
	if err != nil {
		log.Error(err)
	}

	fileerror := ioutil.WriteFile("../../vis-library/cytostruct.json", file, 0644)
	if fileerror != nil {
		log.Error(fileerror)
	}
}

func (cyNode *CyGraph) checkForDouble(a cyNode) bool {
	for _, v := range cyNode.GraphNodes {
		if reflect.DeepEqual(a.Data, v.Data) {
			return true
		}
	}
	return false
}
