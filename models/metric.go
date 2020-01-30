package models

import "time"

// NodeMetric represents a node metric object.
type NodeMetric struct {
	Status string `json:"status"`
	Data   struct {
		ResultType string `json:"resultType"`
		Result     []struct {
			Metric struct {
				Instance string `json:"instance"`
			} `json:"metric"`
			Value []interface{} `json:"value"`
		} `json:"result"`
	} `json:"data"`
}

// GenericMetric respresents a generic metric object
type GenericMetric struct {
	Data struct {
		Result []struct {
			Metric struct {
				Value string `json:"value"`
			} `json:"metric"`
			Value []interface{} `json:"value"`
		} `json:"result"`
		ResultType string `json:"resultType"`
	} `json:"data"`
	Status string `json:"status"`
}

// Payload object to post
type Payload struct {
	ClusterName     string          `json:"cluster_name"`
	CnoxCPU         interface{}     `json:"cnox_cpu"`
	CnoxMem         interface{}     `json:"cnox_mem"`
	CnoxPods        interface{}     `json:"cnox_pods"`
	CnoxServices    interface{}     `json:"cnox_services"`
	CustomerID      string          `json:"customer_id"`
	Nodes           interface{}     `json:"nodes"`
	Pods            interface{}     `json:"pods"`
	Services        interface{}     `json:"services"`
	Timestamp       time.Time       `json:"timestamp"`
	NodeSummaryJSON NodeSummaryJSON `json:"nodeSummaryJSON"`
}

type NodeSummaryJSON struct {
	CPU map[string]interface{} `json:"cpu"`
	Mem interface{}            `json:"mem"`
}
