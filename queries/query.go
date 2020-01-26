package queries

import (
	u "net/url"
	"prometheus/utils"
)

const (
	// query path
	path = "/api/v1/query"
	// irate, avg used to calculate node percentage of all nodes
	// being monitored by node exporters which shall run as daemon sets
	// all nodes.
	cpuNode = "100 - (avg by (instance) (irate(node_cpu_seconds_total{job='node-exporter',mode='idle'}[5m])) * 100)"

	// % of memory using free, buffered and cached divided by total memory ranging over 5m
	memoryNode = "100 * (1 - ((avg_over_time(node_memory_MemFree_bytes[5m]) + avg_over_time(node_memory_Cached_bytes[5m]) + avg_over_time(node_memory_Buffers_bytes[5m])) / avg_over_time(node_memory_MemTotal_bytes[5m])))"
)

// CPUNode shall return CPU Metrics of a node
func CPUNode(url string) (values map[string]interface{}) {
	queryurl := url + path + "?" + "query=" + u.QueryEscape(cpuNode)

	values = utils.HTTPGetReq(queryurl)
	return values
}

// MEMNode shall return MEM Metrics of a node
func MEMNode(url string) (values map[string]interface{}) {
	queryurl := url + path + "?" + "query=" + u.QueryEscape(memoryNode)
	values = utils.HTTPGetReq(queryurl)
	return values
}

// CPUNamespace shall return the cpu consumption of the namespace
func CPUNamespace(url, namespace string) (values map[string]interface{}) {

	// cpu consumption per namespace
	// ORG QUERY: sum(rate(container_cpu_usage_seconds_total{image!='',namespace='cnox'}[5m])) /  sum (machine_cpu_cores) * 100
	// Contruct URL to keep namespace dynamic

	q := "sum(rate(container_cpu_usage_seconds_total{image!='',"
	escapeQ := q + "namespace='" + namespace + "'}[5m]))" + "/ sum (machine_cpu_cores) * 100"
	queryurl := url + path + "?" + "query=" + u.QueryEscape(escapeQ)
	values = utils.HTTPGetNamespaceReq(queryurl)
	return values
}

// MEMNamespace shall return the mem consumption of the namespace
func MEMNamespace(url, namespace string) (values map[string]interface{}) {

	// memory consumption per namespace
	// ORG QUERY: sum(rate(container_memory_working_set_bytes{image!='',namespace='cnox'}[5m])) / sum(node_memory_MemAvailable_bytes) * 100
	// Contruct URL to keep namespace dynamic
	q := "sum(container_memory_working_set_bytes{image!='',"
	escapeQ := q + "namespace='" + namespace + "'}[5m]))" + " / sum(node_memory_MemTotal_bytes) * 100"

	queryurl := url + path + "?" + "query=" + u.QueryEscape(escapeQ)
	values = utils.HTTPGetNamespaceReq(queryurl)
	return values
}
