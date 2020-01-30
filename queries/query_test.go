package queries

import (
	"os"
	"testing"
)

var url = os.Getenv("PROMURL")

func TestCPUNode(t *testing.T) {
	values := CPUNode(url)
	for i, c := range values {
		t.Logf("CPU Usage Values \n: Instance[%s] : Value[%s]", i, c)
	}

}

func TestMEMNode(t *testing.T) {
	url := os.Getenv("PROMURL")

	values := MEMNode(url)

	t.Logf("Memory Usage Values \n: %s", values)

}

func TestCPUNamespace(t *testing.T) {
	url := os.Getenv("PROMURL")
	values := CPUNamespace(url)
	t.Logf("CPU Usage Values \n: : %s", values)

}

func TestMEMNamespace(t *testing.T) {
	url := os.Getenv("PROMURL")

	values := CPUNamespace(url)
	t.Logf("CPU Usage Values \n : %s", values)

}

/*
func TestGetNodeCount(t *testing.T) {
	url := os.Getenv("PROMURL")
	values := GetNodeCount(url)

	for _, c := range values {
		t.Logf("Node Count [%s]", c)
	}
}
*/

/*
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
*/
