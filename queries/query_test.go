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

	values := CPUNamespace(url, username, password)
	for i, c := range values {
		t.Logf("CPU Usage Values \n: Namespace[%s] : Value[%s]", i, c)
	}
}

func TestMEMNamespace(t *testing.T) {
	url := os.Getenv("PROMURL")
	username := os.Getenv("USERNAME")
	password := os.Getenv("PASSWORD")

	values := MEMNamespace(url, username, password)
	for i, c := range values {
		t.Logf("Memory Usage Values \n: Namespace[%s] : Value[%s]", i, c)
	}
}
