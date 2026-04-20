## gRPC+Websocket Test via K6 

```
  █ TOTAL RESULTS

    checks_total.......: 225434  22523.620288/s
    checks_succeeded...: 100.00% 225434 out of 225434
    checks_failed......: 0.00%   0 out of 225434

    ✓ latency < 50ms

    CUSTOM
    msg_latency...........: avg=536.66µs min=0s      med=0s       max=13ms   p(90)=1ms      p(95)=1.99ms

    EXECUTION
    iteration_duration....: avg=10s      min=10s     med=10s      max=10s    p(90)=10s      p(95)=10s
    iterations............: 50     4.995613/s
    vus...................: 50     min=50         max=50
    vus_max...............: 50     min=50         max=50

    NETWORK
    data_received.........: 25 MB  2.5 MB/s
    data_sent.............: 26 MB  2.6 MB/s

    GRPC
    grpc_latency..........: avg=288.77µs min=72.59µs med=196.79µs max=6.04ms p(90)=492.84µs p(95)=988.33µs

    WEBSOCKET
    ws_connecting.........: avg=4.09ms   min=2.27ms  med=4.05ms   max=5.96ms p(90)=5.17ms   p(95)=5.56ms
    ws_msgs_received......: 225434 22523.620288/s
    ws_msgs_sent..........: 225484 22528.615901/s
    ws_session_duration...: avg=10s      min=10s     med=10s      max=10s    p(90)=10s      p(95)=10s
    ws_sessions...........: 50     4.995613/s
```

| Metric | Unary gRPC (ghz) | Unary WS+gRPC (k6) | Unary gRPC layer | Unary WS overhead | Unary Delta | Bidi gRPC (ghz) | Bidi WS+gRPC (k6) | Bidi gRPC layer | Bidi WS overhead | Bidi Delta |
|--------|-----------------|-------------------|-----------------|-------------------|-------------|-----------------|------------------|-----------------|------------------|------------|
| req/sec | 34,970 | 14,016 | — | — | -20,954 | 27,794 | 22,524 | — | — | -5,270 |
| p50 | 0.93ms | 3ms | 2.70ms | 0.30ms | +2.07ms | 1.40ms | ~0ms | 0.197ms | ~0ms | -1.40ms |
| p90 | — | 4.99ms | 3.88ms | 1.11ms | — | — | 1ms | 0.493ms | 0.507ms | — |
| p95 | 2.17ms | 5ms | 4.30ms | 0.70ms | +2.83ms | 2.58ms | 1.99ms | 0.988ms | 1.002ms | -0.59ms |

> Note: msg_latency p50 for bidi WS+gRPC reports 0ms due to Date.now() 
> millisecond resolution limitations in k6. Actual latency is sub-millisecond, 
> consistent with grpc_latency p50 of 0.197ms measured at the gateway layer 
> using Go's time.Now().UnixMilli().
