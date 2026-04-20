## gRPC+Websocket Test via K6 

```
  █ TOTAL RESULTS

    checks_total.......: 140295  14015.847002/s
    checks_succeeded...: 100.00% 140295 out of 140295
    checks_failed......: 0.00%   0 out of 140295

    ✓ latency < 50ms

    CUSTOM
    msg_latency...........: avg=3.51ms min=0s       med=3ms    max=33ms    p(90)=4.99ms p(95)=5ms

    EXECUTION
    iteration_duration....: avg=10s    min=10s      med=10s    max=10s     p(90)=10s    p(95)=10s
    iterations............: 50     4.995134/s
    vus...................: 50     min=50         max=50
    vus_max...............: 50     min=50         max=50

    NETWORK
    data_received.........: 21 MB  2.1 MB/s
    data_sent.............: 16 MB  1.6 MB/s

    GRPC
    grpc_latency..........: avg=2.84ms min=160.26µs med=2.7ms  max=15.53ms p(90)=3.88ms p(95)=4.3ms

    WEBSOCKET
    ws_connecting.........: avg=4.2ms  min=2.35ms   med=3.48ms max=7.95ms  p(90)=6.8ms  p(95)=7.09ms
    ws_msgs_received......: 140295 14015.847002/s
    ws_msgs_sent..........: 140345 14020.842136/s
    ws_session_duration...: avg=10s    min=10s      med=10s    max=10s     p(90)=10s    p(95)=10s
    ws_sessions...........: 50     4.995134/s
```


| Metric | WS+gRPC total | gRPC layer | WS overhead |
|--------|--------------|------------|-------------|
| req/sec | 14,016 | — | — |
| p50 | 3ms | 2.70ms | 0.30ms |
| p90 | 4.99ms | 3.88ms | 1.11ms |
| p95 | 5ms | 4.30ms | 0.70ms |
