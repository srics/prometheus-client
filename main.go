package main

import (
	"encoding/json"
	"log"
	"os"
	"prometheus/models"
	"prometheus/queries"
	"time"
)

// Add namespaces to this slice

var url = os.Getenv("PROMURL")

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

}

func main() {

	payload()
}
