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

const (
	cloudSaaSCertBundleFile = "/etc/ssl/certs/bundle.crt"
)

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
	log.Printf("%s", b)

	var certBundle string
	if _, err := os.Stat(cloudSaaSCertBundleFile); err == nil {
		certBundle = cloudSaaSCertBundleFile
	}

	p, err := utils.HTTPPost(ep, certBundle, b)
	if err != nil {
		log.Printf("Endpoing %s, %s", ep, p.Status)
	}
	log.Printf("Endpoint %s, Response:[%s],Method:[%s]", ep, p.Status, p.Request.Method)

}

func main() {

	payload()
}
