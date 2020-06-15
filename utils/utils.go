package utils

import (
	"bytes"
	"crypto/tls"
	"crypto/x509"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/pkg/errors"

	"prometheus/models"
)

var (
	generic models.GenericMetric
	node    models.NodeMetric
)

/*
 * http client connection to CNOX SaaS Cloud
 */
func getCNOXSaaSClient(cloudSaaSCertBundle string) (*http.Client, error) {
	var transport http.Transport
	if cloudSaaSCertBundle != "" {
		cert, err := ioutil.ReadFile(cloudSaaSCertBundle)
		if err != nil {
			return nil, errors.Wrapf(err, "Reading api-server tls cert bundle")
		}
		caCertPool := x509.NewCertPool()
		caCertPool.AppendCertsFromPEM(cert)
		transport.TLSClientConfig = &tls.Config{RootCAs: caCertPool}
	}
	client := http.Client{
		Transport: &transport,
		Timeout: time.Second * 120,
	}

	return &client, nil
}

/*
 * http client connection to prometheus
 */
func getPromHTTPClient() (*http.Client) {
	client := http.Client{
		Timeout: time.Second * 120,
	}
	return &client
}

// HTTPGetMetric returns metric value
func HTTPGetMetric(url string) (values interface{}) {

	client := getPromHTTPClient()
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatalf("NewRequest construct error : %d", err)
	}

	req.SetBasicAuth(os.Getenv("USERNAME"), os.Getenv("PASSWORD"))

	resp, err := client.Do(req)
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

	client := getPromHTTPClient()
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatalf("NewRequest construct error : %d", err)
	}

	req.SetBasicAuth(os.Getenv("USERNAME"), os.Getenv("PASSWORD"))

	resp, err := client.Do(req)
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

// HTTPPost to endpoint
func HTTPPost(ep string, cloudSaaSCertBundle string, json []byte) (resp *http.Response, err error) {

	client, err := getCNOXSaaSClient(cloudSaaSCertBundle)
	if err != nil {
		log.Fatalf("getCNOXSaaSClient error : %s", err)
	}
	req, err := http.NewRequest("POST", ep, bytes.NewBuffer(json))
	req.Header.Set("Content-type", "application/json")
	if err != nil {
		log.Fatalf("NewRequest construct error : %d", err)
	}

	req.SetBasicAuth(os.Getenv("USERNAME"), os.Getenv("PASSWORD"))

	resp, err = client.Do(req)
	if err != nil {
		log.Fatalf("POST request error on URL specified : %s", err)

	}

	defer resp.Body.Close()

	return resp, err

}
