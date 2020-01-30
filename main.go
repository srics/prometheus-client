package main

import (
	"encoding/json"
	"log"
	"os"
	"prometheus/models"
	"prometheus/queries"
	"prometheus/utils"
	"time"
)

var url = os.Getenv("PROMURL")
var ep = os.Getenv("ENDPOINT")
var clusterName = os.Getenv("CLUSTERNAME")
var customerID = os.Getenv("CUSTOMERID")

func payload() {
	var f = func(url string) (a models.Payload) {
		cnoxMem := queries.MEMNamespace(url)
		cnoxCPU := queries.CPUNamespace(url)
		nodeCPU := queries.CPUNode(url)
		nodeMem := queries.MEMNode(url)
		node := queries.GetNodeCount(url)
		cnoxpods := queries.GetCnoxPodCount(url)
		pods := queries.GetPodCount(url)
		svc := queries.GetSvcCount(url)
		cnoxsvc := queries.GetCnoxSvcCount(url)
		timestamp := time.Now()
		a = models.Payload{
			ClusterName:  clusterName,
			CustomerID:   customerID,
			CnoxMem:      cnoxMem,
			CnoxCPU:      cnoxCPU,
			Nodes:        node,
			CnoxPods:     cnoxpods,
			Timestamp:    timestamp,
			Pods:         pods,
			Services:     svc,
			CnoxServices: cnoxsvc,
			NodeSummaryJSON: models.NodeSummaryJSON{
				CPU: nodeCPU,
				Mem: nodeMem,
			},
		}
		return a
	}

	b, err := json.Marshal(f(url))
	if err != nil {
		panic(err)
	}

	p, err := utils.HTTPPost(ep, b)
	if err != nil {
		log.Printf("%s", p.Status)
	}
	log.Printf("Response:[%s],Method:[%s]", p.Status, p.Request.Method)

}

func main() {

	payload()
}
