package queries

import (
	u "net/url"
	"prometheus/utils"
	//	"log"
)

const (
	// query path
	query = "/api/v1/query"
	// irate, avg used to calculate node percentage of all nodes
	// being monitored by node exporters which shall run as daemon sets
	// all nodes.
	cpuNode = "100 - (avg by (instance) (irate(node_cpu_seconds_total{job='node-exporter',mode='idle'}[5m])) * 100)"

	// % of memory using free, buffered and cached divided by total memory ranging over 5m
	memoryNode = "100 * (1 - ((avg_over_time(node_memory_MemFree_bytes[5m]) + avg_over_time(node_memory_Cached_bytes[5m]) + avg_over_time(node_memory_Buffers_bytes[5m])) / avg_over_time(node_memory_MemTotal_bytes[5m])))"

	// CPU Consumption per namespace

	cpuNamespace = "sum(rate(container_cpu_usage_seconds_total{image!='',namespace='logging'}[5m])) by (namespace)"

	// Memory Consumption per namespace
	memNamespace = "sum(container_memory_working_set_bytes{namespace='logging'}) by (namespace)"
)

// CPUNode shall return CPU Metrics of a node
func CPUNode(url, username, password string) (values map[string]interface{}) {
	queryurl := url + query + "?" + "query=" + u.QueryEscape(cpuNode)
	values = utils.HTTPGetReq(queryurl, username, password)
	return values
}

// MEMNode shall return MEM Metrics of a node
func MEMNode(url, username, password string) (values map[string]interface{}) {
	queryurl := url + query + "?" + "query=" + u.QueryEscape(memoryNode)
	values = utils.HTTPGetReq(queryurl, username, password)
	return values
}

// CPUNamespace shall return the cpu consumption of the namespace
func CPUNamespace(url, username, password string) (values map[string]interface{}) {
	queryurl := url + query + "?" + "query=" + u.QueryEscape(cpuNamespace)
	values = utils.HTTPGetNamespaceReq(queryurl, username, password)
	return values
}

// MEMNamespace shall return the mem consumption of the namespace
func MEMNamespace(url, username, password string) (values map[string]interface{}) {
	queryurl := url + query + "?" + "query=" + u.QueryEscape(memNamespace)
	values = utils.HTTPGetNamespaceReq(queryurl, username, password)
	return values
}
