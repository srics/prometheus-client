package db

import (
	"os"
	"testing"
)

func TestCreateNode(t *testing.T) {
	url := os.Getenv("PROMURL")
	//namespace := "logging"
	// TODO
	CreateNode(url)

}
