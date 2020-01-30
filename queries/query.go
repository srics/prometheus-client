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

	values = utils.HTTPGetNodeMetric(queryurl)

	return values
}

// MEMNode shall return MEM Metrics of a node
func MEMNode(url string) (values interface{}) {
	queryurl := url + path + "?" + "query=" + u.QueryEscape(memoryNode)
	values = utils.HTTPGetNodeMetric(queryurl)
	return values
}

// CPUNamespace shall return the cpu consumption of the namespace
func CPUNamespace(url string) (values interface{}) {

	// cpu consumption per namespace
	// ORG QUERY: sum(rate(container_cpu_usage_seconds_total{image!='',namespace='scanner'}[30m])) /  sum (machine_cpu_cores) * 100 + sum(rate(container_cpu_usage_seconds_total{image!='',namespace='cnox'}[30m])) /  sum (machine_cpu_cores) * 100 + sum(rate(container_cpu_usage_seconds_total{image!='',namespace='scanner'}[30m])) /  sum (machine_cpu_cores) * 100
	// Contruct URL to keep namespace dynamic

	//q := "sum(irate(container_cpu_usage_seconds_total{image!='',"
	escapeQ := "sum(rate(container_cpu_usage_seconds_total{image!='',namespace='scanner'}[30m])) /  sum (machine_cpu_cores) * 100 + sum(rate(container_cpu_usage_seconds_total{image!='',namespace='cnox'}[30m])) /  sum (machine_cpu_cores) * 100 + sum(rate(container_cpu_usage_seconds_total{image!='',namespace='scanner'}[30m])) /  sum (machine_cpu_cores) * 100"
	queryurl := url + path + "?" + "query=" + u.QueryEscape(escapeQ)
	values = utils.HTTPGetMetric(queryurl)
	return values
}

// MEMNamespace shall return the mem consumption of the namespace
func MEMNamespace(url string) (values interface{}) {

	// memory consumption per namespace
	// ORG QUERY: sum (container_memory_usage_bytes{image!="",namespace="scanner"}) / sum(node_memory_MemTotal_bytes) * 100 + sum (container_memory_usage_bytes{image!="",namespace="cnox"}) / sum(node_memory_MemTotal_bytes) * 100 + sum (container_memory_usage_bytes{image!="",namespace="monitoring"}) / sum(node_memory_MemTotal_bytes) * 100
	// Contruct URL to keep namespace dynamic
	escapeQ := "sum (container_memory_usage_bytes{image!='',namespace='scanner'}) / sum(node_memory_MemTotal_bytes) * 100 + sum (container_memory_usage_bytes{image!='',namespace='cnox'}) / sum(node_memory_MemTotal_bytes) * 100 +	sum (container_memory_usage_bytes{image!='',namespace='monitoring'}) / sum(node_memory_MemTotal_bytes) * 100"

	queryurl := url + path + "?" + "query=" + u.QueryEscape(escapeQ)
	values = utils.HTTPGetMetric(queryurl)
	return values
}

// GetNodeCount shall return the current ready nodes
func GetNodeCount(url string) (values interface{}) {
	// node count with state ready
	// ORG QUERY: count(kube_node_status_condition{condition="Ready",status="true"})
	q := "count(kube_node_status_condition{condition='Ready',status='true'})"
	queryurl := url + path + "?" + "query=" + u.QueryEscape(q)

	values = utils.HTTPGetMetric(queryurl)

	return values

}

// GetCnoxPodCount shall return the number of pods in namespaces
func GetCnoxPodCount(url string) (values interface{}) {
	// pod count in namespaces
	// ORG QUERY: count(kube_pod_info{namespace='scanner'})  + count(kube_pod_info{namespace='cnox'}) + count(kube_pod_info{namespace='monitoring'})
	q := "count(kube_pod_info{namespace='scanner'})  + count(kube_pod_info{namespace='cnox'}) + count(kube_pod_info{namespace='monitoring'})"
	queryurl := url + path + "?" + "query=" + u.QueryEscape(q)

	values = utils.HTTPGetMetric(queryurl)

	return values
}

// GetPodCount shall return the number of pods running in cluster
func GetPodCount(url string) (values interface{}) {
	// pod count in namespaces
	// ORG QUERY: count(kube_pod_info{namespace!=''}
	q := "count(kube_pod_info{namespace!=''})"
	queryurl := url + path + "?" + "query=" + u.QueryEscape(q)

	values = utils.HTTPGetMetric(queryurl)

	return values
}

// GetSvcCount shall return the number of svc running in cluster
func GetSvcCount(url string) (values interface{}) {
	// pod count in namespaces
	// ORG QUERY: count(kube_service_info{namespace!=''})
	q := "count(kube_service_info{namespace!=''})"
	queryurl := url + path + "?" + "query=" + u.QueryEscape(q)

	values = utils.HTTPGetMetric(queryurl)

	return values
}

// GetCnoxSvcCount shall return the number of svc running in namespaces
func GetCnoxSvcCount(url string) (values interface{}) {
	// pod count in namespaces
	// ORG QUERY: count(kube_service_info{namespace='cnox'}) + count(kube_service_info{namespace='scanner'}) + count(kube_service_info{namespace='monitoring'})
	q := "count(kube_service_info{namespace='cnox'}) + count(kube_service_info{namespace='scanner'}) + count(kube_service_info{namespace='monitoring'})"
	queryurl := url + path + "?" + "query=" + u.QueryEscape(q)

	values = utils.HTTPGetMetric(queryurl)

	return values
}
