# Prometheus

- CPU and Memory consumptions calculated for each node running in a cluster
```
=== RUN   TestCPU
--- PASS: TestCPU (0.99s)
    query_test.go:15: CPU Usage Values 
        : Instance[10.20.24.51:9100] : Value[4.800000000395812]
    query_test.go:15: CPU Usage Values 
        : Instance[10.20.28.6:9100] : Value[5.2999999999883585]
    query_test.go:15: CPU Usage Values 
        : Instance[10.20.32.57:9100] : Value[8.300000000162981]
    query_test.go:15: CPU Usage Values 
        : Instance[10.20.32.8:9100] : Value[6.200000001117587]
    query_test.go:15: CPU Usage Values 
        : Instance[10.20.32.97:9100] : Value[4.599999999918509]
    query_test.go:15: CPU Usage Values 
        : Instance[10.20.24.29:9100] : Value[4.600000000209548]
=== RUN   TestMemory
--- PASS: TestMemory (0.21s)
    query_test.go:27: Memory Usage Values 
        : Instance[10.20.24.29:9100] : Value[36.894112027931214]
    query_test.go:27: Memory Usage Values 
        : Instance[10.20.24.51:9100] : Value[36.52546015359307]
    query_test.go:27: Memory Usage Values 
        : Instance[10.20.28.6:9100] : Value[37.002542590424206]
    query_test.go:27: Memory Usage Values 
        : Instance[10.20.32.57:9100] : Value[61.25722408128807]
    query_test.go:27: Memory Usage Values 
        : Instance[10.20.32.8:9100] : Value[20.73735552685021]
    query_test.go:27: Memory Usage Values 
        : Instance[10.20.32.97:9100] : Value[35.97621246940871]
PASS
ok      prometheus/queries      1.209s
```