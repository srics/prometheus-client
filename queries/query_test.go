package queries

import (
	"os"
	"testing"
)

func TestCPUNode(t *testing.T) {
	url := os.Getenv("PROMURL")

	values := CPUNode(url)
	for i, c := range values {
		t.Logf("CPU Usage Values \n: Instance[%s] : Value[%s]", i, c)
	}

}

func TestMEMNode(t *testing.T) {
	url := os.Getenv("PROMURL")

	values := MEMNode(url)
	for i, c := range values {
		t.Logf("Memory Usage Values \n: Instance[%s] : Value[%s]", i, c)
	}
}

func TestCPUNamespace(t *testing.T) {
	url := os.Getenv("PROMURL")

	// use any namespace as required, keep on populating
	// this ns slice
	ns := []string{"kube-system", "logging"}

	for _, c := range ns {
		values := CPUNamespace(url, c)
		t.Logf("CPU Usage Values \n: Namespace[%s] : Value[%s]", c, values)
	}

}

func TestMEMNamespace(t *testing.T) {
	url := os.Getenv("PROMURL")

	ns := []string{"kube-system", "logging"}

	for _, c := range ns {
		values := CPUNamespace(url, c)
		t.Logf("CPU Usage Values \n: Namespace[%s] : Value[%s]", c, values)
	}
}

func TestQueryNamespace(t *testing.T) {
	url := os.Getenv("PROMURL")

	ns := []string{"kube-system", "logging"}

	metric := make(chan map[string]interface{}, 100)
	for _, c := range ns {
		go QueryNamespace(metric, url, c)
		m1 := <-metric
		t.Logf("Metric Namespace CPU \n: %s", m1)

		m2 := <-metric
		t.Logf("Metric Namespace Memory \n: %s", m2)
	}

}
