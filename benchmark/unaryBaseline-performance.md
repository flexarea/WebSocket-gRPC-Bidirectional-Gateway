## gRPC+Websocket Test via K6 

```
  █ TOTAL RESULTS

    checks_total.......: 139556  13939.807531/s
    checks_succeeded...: 100.00% 139556 out of 139556
    checks_failed......: 0.00%   0 out of 139556

    ✓ latency < 50ms

    CUSTOM
    msg_latency...........: avg=3.52ms min=423.42µs med=3.34ms max=21.07ms p(90)=4.89ms p(95)=5.44ms

    EXECUTION
    iteration_duration....: avg=10s    min=10s      med=10s    max=10.01s  p(90)=10s    p(95)=10.01s
    iterations............: 50     4.994342/s
    vus...................: 50     min=50         max=50
    vus_max...............: 50     min=50         max=50

    NETWORK
    data_received.........: 21 MB  2.1 MB/s
    data_sent.............: 16 MB  1.6 MB/s

    GRPC
    grpc_latency..........: avg=2.8ms  min=155.13µs med=2.66ms max=17.59ms p(90)=3.89ms p(95)=4.33ms

    WEBSOCKET
    ws_connecting.........: avg=4.28ms min=795.35µs med=2.85ms max=10.1ms  p(90)=8.98ms p(95)=9.73ms
    ws_msgs_received......: 139556 13939.807531/s
    ws_msgs_sent..........: 139606 13944.801873/s
    ws_session_duration...: avg=10s    min=10s      med=10s    max=10.01s  p(90)=10s    p(95)=10.01s
    ws_sessions...........: 50     4.994342/s
```



## Unary Baseline Performance

| Metric | WS+gRPC total | gRPC layer | WS overhead |
|--------|--------------|------------|-------------|
| req/sec | 13,940 | — | — |
| p50 | 3.34ms | 2.66ms | 0.68ms |
| p90 | 4.89ms | 3.89ms | 1.00ms |
| p95 | 5.44ms | 4.33ms | 1.11ms |
