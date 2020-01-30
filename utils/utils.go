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
	generic models.GenericMetric
	node    models.NodeMetric
)

// HTTPGetMetric returns metric value
func HTTPGetMetric(url string) (values interface{}) {

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

	unmarshal := json.Unmarshal(respData, &generic)
	if unmarshal != nil {
		log.Fatalf("Error in unmarshaling request %d", err)

	}

	// values represents instance:value
	//values = make(map[string]interface{})

	// range over the values
	for i := range generic.Data.Result {
		return generic.Data.Result[i].Value[1]
	}

	return values

}

func HTTPGetNodeMetric(url string) (values map[string]interface{}) {

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
