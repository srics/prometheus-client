package queries

import (
	"os"
	"testing"
)

func TestCPUNode(t *testing.T) {
	url := os.Getenv("PROMURL")
	username := os.Getenv("USERNAME")
	password := os.Getenv("PASSWORD")

	values := CPUNode(url, username, password)
	for i, c := range values {
		t.Logf("CPU Usage Values \n: Instance[%s] : Value[%s]", i, c)
	}

}

func TestMEMNode(t *testing.T) {
	url := os.Getenv("PROMURL")
	username := os.Getenv("USERNAME")
	password := os.Getenv("PASSWORD")

	values := MEMNode(url, username, password)
	for i, c := range values {
		t.Logf("Memory Usage Values \n: Instance[%s] : Value[%s]", i, c)
	}
}

func TestCPUNamespace(t *testing.T) {
	url := os.Getenv("PROMURL")
	username := os.Getenv("USERNAME")
	password := os.Getenv("PASSWORD")
	// use any namespace as required, keep on populating
	// this ns slice
	ns := []string{"kube-system", "logging"}

	for _, c := range ns {
		values := CPUNamespace(url, username, password, c)
		t.Logf("CPU Usage Values \n: Namespace[%s] : Value[%s]", c, values)
	}

}

func TestMEMNamespace(t *testing.T) {
	url := os.Getenv("PROMURL")
	username := os.Getenv("USERNAME")
	password := os.Getenv("PASSWORD")
	ns := []string{"kube-system", "logging"}

	for _, c := range ns {
		values := CPUNamespace(url, username, password, c)
		t.Logf("CPU Usage Values \n: Namespace[%s] : Value[%s]", c, values)
	}
}

func TestQueryNamespace(t *testing.T) {
	url := os.Getenv("PROMURL")
	username := os.Getenv("USERNAME")
	password := os.Getenv("PASSWORD")

	ns := []string{"kube-system", "logging"}

	metric := make(chan map[string]interface{}, 100)
	for _, c := range ns {
		go QueryNamespace(metric, url, username, password, c)
		m1 := <-metric
		t.Logf("Metric Namespace CPU \n: %s", m1)

		m2 := <-metric
		t.Logf("Metric Namespace Memory \n: %s", m2)
	}

}
