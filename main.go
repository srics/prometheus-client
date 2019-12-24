package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"prometheus/queries"

	"github.com/gorilla/mux"
)

// Add namespaces to this slice
var namespace = []string{"logging", "kube-system"}

// getnamespacecpu returns CPU consumed in a namespace
func getnamespacecpu(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	for _, c := range namespace {
		fmt.Fprintf(w, "CPU consumption in namespace: %v\n", c)
		x := queries.CPUNamespace(os.Getenv("PROMURL"), c)
		var _ = json.NewEncoder(w).Encode(x)
	}
}

// getnamespacemem returns MEM consumed in a namespace
func getnamespacemem(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	for _, c := range namespace {
		fmt.Fprintf(w, "Memory consumption in namespace: %v\n", c)
		x := queries.MEMNamespace(os.Getenv("PROMURL"), c)
		var _ = json.NewEncoder(w).Encode(x)
	}
}

// getnodecpu returns CPU consumed in nodes
func getnodecpu(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	x := queries.CPUNode(os.Getenv("PROMURL"))
	fmt.Fprintf(w, "CPU Consumption in Nodes\n")
	for i, c := range x {
		fmt.Fprintf(w, "[%s] : [%s] \n", i, c)
	}

}

// getnodemem returns MEM consumed in nodes
func getnodemem(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	x := queries.MEMNode(os.Getenv("PROMURL"))
	fmt.Fprintf(w, "Memory Consumption in Nodes\n")
	for i, c := range x {
		fmt.Fprintf(w, "[%s] : [%s] \n", i, c)
	}

}
func main() {
	r := mux.NewRouter()
	r.HandleFunc("/metrics/namespacecpu", getnamespacecpu).Methods("GET")
	r.HandleFunc("/metrics/namespacemem", getnamespacemem).Methods("GET")
	r.HandleFunc("/metrics/nodecpu", getnodecpu).Methods("GET")
	r.HandleFunc("/metrics/nodemem", getnodemem).Methods("GET")
	log.Fatal(http.ListenAndServe(":8000", r))
}
