## gRPC+Websocket Test via K6 

```
 █ TOTAL RESULTS

    checks_total.......: 210876  21066.198533/s
    checks_succeeded...: 100.00% 210876 out of 210876
    checks_failed......: 0.00%   0 out of 210876

    ✓ latency < 50ms

    CUSTOM
    msg_latency...........: avg=549.08µs min=186.11µs med=368.64µs max=10.89ms p(90)=1.13ms  p(95)=1.63ms

    EXECUTION
    iteration_duration....: avg=10s      min=10s      med=10s      max=10s     p(90)=10s     p(95)=10s
    iterations............: 50     4.994926/s
    vus...................: 50     min=50         max=50
    vus_max...............: 50     min=50         max=50

    NETWORK
    data_received.........: 23 MB  2.3 MB/s
    data_sent.............: 24 MB  2.4 MB/s

    GRPC
    grpc_latency..........: avg=288.86µs min=66.24µs  med=198.25µs max=5.4ms   p(90)=493.4µs p(95)=967.7µs

    WEBSOCKET
    ws_connecting.........: avg=2.34ms   min=440.32µs med=2.18ms   max=4.94ms  p(90)=3.99ms  p(95)=4.33ms
    ws_msgs_received......: 210876 21066.198533/s
    ws_msgs_sent..........: 210926 21071.193458/s
    ws_session_duration...: avg=10s      min=10s      med=10s      max=10s     p(90)=10s     p(95)=10s
    ws_sessions...........: 50     4.994926/s
```

## Bidi Baseline Performance

| Metric | WS+gRPC total | gRPC layer | WS overhead |
|--------|--------------|------------|-------------|
| req/sec | 21,066 | — | — |
| p50 | 0.369ms | 0.198ms | 0.171ms |
| p90 | 1.13ms | 0.493ms | 0.637ms |
| p95 | 1.63ms | 0.968ms | 0.662ms |
