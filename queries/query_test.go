package queries

import (
	"os"
	"testing"
)

func TestCPU(t *testing.T) {
	url := os.Getenv("PROMURL")
	username := os.Getenv("USERNAME")
	password := os.Getenv("PASSWORD")

	values := CPU(url, username, password)
	for i, c := range values {
		t.Logf("CPU Usage Values \n: Instance[%s] : Value[%s]", i, c)
	}

}

func TestMemory(t *testing.T) {
	url := os.Getenv("PROMURL")
	username := os.Getenv("USERNAME")
	password := os.Getenv("PASSWORD")

	values := Memory(url, username, password)
	for i, c := range values {
		t.Logf("Memory Usage Values \n: Instance[%s] : Value[%s]", i, c)
	}
}
