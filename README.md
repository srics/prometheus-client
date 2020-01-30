# Prometheus

- CPU and Memory consumptions calculated for each node running in a cluster
```
BNG-188814-C02YH02XJGH8:queries adheipsingh$ go test -v
=== RUN   TestCPUNode
--- PASS: TestCPUNode (1.09s)
    query_test.go:13: CPU Usage Values 
        : Instance[ip-172-20-59-4.us-west-2.compute.internal] : Value[6.633333334078387]
    query_test.go:13: CPU Usage Values 
        : Instance[ip-172-20-49-247.us-west-2.compute.internal] : Value[21.133333332836628]
    query_test.go:13: CPU Usage Values 
        : Instance[ip-172-20-50-189.us-west-2.compute.internal] : Value[10.88333333532016]
=== RUN   TestMEMNode
--- PASS: TestMEMNode (0.71s)
    query_test.go:23: Memory Usage Values 
        : Instance[ip-172-20-49-247.us-west-2.compute.internal] : Value[49.367007878180125]
    query_test.go:23: Memory Usage Values 
        : Instance[ip-172-20-50-189.us-west-2.compute.internal] : Value[17.670936324898356]
    query_test.go:23: Memory Usage Values 
        : Instance[ip-172-20-59-4.us-west-2.compute.internal] : Value[24.211021352856243]
=== RUN   TestCPUNamespace
--- PASS: TestCPUNamespace (1.44s)
    query_test.go:36: CPU Usage Values 
        : Namespace[kube-system] : Value[map[:3.329674950594584]]
    query_test.go:36: CPU Usage Values 
        : Namespace[logging] : Value[map[]]
=== RUN   TestMEMNamespace
--- PASS: TestMEMNamespace (1.40s)
    query_test.go:48: CPU Usage Values 
        : Namespace[kube-system] : Value[map[:3.354086176746536]]
    query_test.go:48: CPU Usage Values 
        : Namespace[cnox] : Value[map[:0.0030606717777775182]]
PASS
ok      prometheus/queries      4.651s
```