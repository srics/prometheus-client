package utils

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"prometheus/models"
)

var (
	node      models.NodeMetric
	namespace models.NamespaceMetric
)

// HTTPGetReq function shall return instance:values
func HTTPGetReq(url string) (values map[string]interface{}) {

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatalf("NewRequest construct error : %d", err)
	}

	req.SetBasicAuth(os.Getenv("USERNAME"), os.Getenv("PASSWORD"))

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatalf("GET request error on URL specified : %d", err)

	}

	defer resp.Body.Close()

	respData, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("Reading Body Error : %d", err)
	}

	unmarshal := json.Unmarshal(respData, &node)
	if unmarshal != nil {
		log.Fatalf("Error in unmarshaling request %d", err)

	}

	// values represents instance:value
	values = make(map[string]interface{})

	// range over the values
	for i := range node.Data.Result {
		values[node.Data.Result[i].Metric.Instance] = node.Data.Result[i].Value[1]
	}

	return values

}

//HTTPGetNamespaceReq unmarshalling to Namespace struct
// TODO: Break func till unmarshal
func HTTPGetNamespaceReq(url string) (values map[string]interface{}) {

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatalf("NewRequest construct error : %d", err)
	}

	req.SetBasicAuth(os.Getenv("USERNAME"), os.Getenv("PASSWORD"))

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatalf("GET request error on URL specified : %d", err)

	}

	defer resp.Body.Close()

	respData, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("Reading Body Error : %d", err)
	}

	unmarshal := json.Unmarshal(respData, &namespace)
	if unmarshal != nil {
		log.Fatalf("Error in unmarshaling request %d", err)

	}

	// values represents instance:value
	values = make(map[string]interface{})

	// range over the values
	for i := range namespace.Data.Result {
		values[namespace.Data.Result[i].Metric.Namespace] = namespace.Data.Result[i].Value[1]
	}

	return values

}
