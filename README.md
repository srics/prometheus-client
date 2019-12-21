# Prometheus

- CPU and Memory consumptions calculated for each node running in a cluster
```
BNG-188814-C02YH02XJGH8:queries adheipsingh$ go test -v
=== RUN   TestCPUNode
--- PASS: TestCPUNode (1.13s)
    query_test.go:15: CPU Usage Values 
        : Instance[10.20.32.97:9100] : Value[9.499999999825377]
    query_test.go:15: CPU Usage Values 
        : Instance[10.20.24.29:9100] : Value[5]
    query_test.go:15: CPU Usage Values 
        : Instance[10.20.24.51:9100] : Value[4.299999999930151]
    query_test.go:15: CPU Usage Values 
        : Instance[10.20.28.6:9100] : Value[7.900000000372529]
    query_test.go:15: CPU Usage Values 
        : Instance[10.20.32.57:9100] : Value[9.40000000060536]
    query_test.go:15: CPU Usage Values 
        : Instance[10.20.32.8:9100] : Value[6.500000003725276]
=== RUN   TestMEMNode
--- PASS: TestMEMNode (0.21s)
    query_test.go:27: Memory Usage Values 
        : Instance[10.20.24.29:9100] : Value[37.97103793505576]
    query_test.go:27: Memory Usage Values 
        : Instance[10.20.24.51:9100] : Value[36.39307793490655]
    query_test.go:27: Memory Usage Values 
        : Instance[10.20.28.6:9100] : Value[38.15633116206012]
    query_test.go:27: Memory Usage Values 
        : Instance[10.20.32.57:9100] : Value[62.19700051873025]
    query_test.go:27: Memory Usage Values 
        : Instance[10.20.32.8:9100] : Value[20.64542924078917]
    query_test.go:27: Memory Usage Values 
        : Instance[10.20.32.97:9100] : Value[36.392933816350904]
=== RUN   TestCPUNamespace
--- PASS: TestCPUNamespace (0.21s)
    query_test.go:38: CPU Usage Values 
        : Namespace[logging] : Value[0.1317236145729926]
=== RUN   TestMEMNamespace
--- PASS: TestMEMNamespace (0.21s)
    query_test.go:49: Memory Usage Values 
        : Namespace[logging] : Value[9049116672]
=== RUN   TestQueryCPU
--- PASS: TestQueryCPU (0.41s)
    query_test.go:62: Metric Namespace CPU 
        : map[logging:0.12730917942034184]
    query_test.go:65: Metric Node CPU 
```