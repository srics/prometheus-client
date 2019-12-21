package queries

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	u "net/url"
	"prometheus/models"
)

const (
	// query path
	query = "/api/v1/query"
	// irate, avg used to calculate node percentage of all nodes
	// being monitored by node exporters which shall run as daemon sets
	// all nodes.
	cpu = "100 - (avg by (instance) (irate(node_cpu_seconds_total{job='node-exporter',mode='idle'}[5m])) * 100)"

	// % of memory using free, buffered and cached divided by total memory ranging over 5m
	memory = "100 * (1 - ((avg_over_time(node_memory_MemFree_bytes[5m]) + avg_over_time(node_memory_Cached_bytes[5m]) + avg_over_time(node_memory_Buffers_bytes[5m])) / avg_over_time(node_memory_MemTotal_bytes[5m])))"
)

// CPU function returns a response of type instance:value
func CPU(url, username, password string) (values map[string]interface{}) {

	// Construct url
	queryurl := url + query + "?" + "query=" + u.QueryEscape(cpu)

	req, err := http.NewRequest("GET", queryurl, nil)
	if err != nil {
		log.Fatalf("NewRequest construct error : %d", err)
	}

	req.SetBasicAuth(username, password)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatalf("GET request error on URL specified : %d", err)

	}

	defer resp.Body.Close()

	respData, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("Reading Body Error : %d", err)
	}

	var id models.Prometheus

	unmarshal := json.Unmarshal(respData, &id)
	if unmarshal != nil {
		log.Fatalf("Error in unmarshaling request %d", err)

	}

	// values represents instance:value
	values = make(map[string]interface{})

	// range over the values
	for i := range id.Data.Result {
		values[id.Data.Result[i].Metric.Instance] = id.Data.Result[i].Value[1]
	}

	return values
}

// Memory Function Returns Memory instance:value
func Memory(url, username, password string) (values map[string]interface{}) {

	queryurl := url + query + "?" + "query=" + u.QueryEscape(memory)

	req, err := http.NewRequest("GET", queryurl, nil)
	if err != nil {
		log.Fatalf("NewRequest construct error : %d", err)
	}

	req.SetBasicAuth(username, password)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatalf("GET request error on URL specified : %d", err)

	}

	defer resp.Body.Close()

	respData, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("Reading Body Error : %d", err)
	}

	var id models.Prometheus

	unmarshal := json.Unmarshal(respData, &id)
	if unmarshal != nil {
		log.Fatalf("Error in unmarshaling request %d", err)

	}

	// values represents instance:value
	values = make(map[string]interface{})

	// range over the values
	for i := range id.Data.Result {
		values[id.Data.Result[i].Metric.Instance] = id.Data.Result[i].Value[1]
	}

	return values
}
