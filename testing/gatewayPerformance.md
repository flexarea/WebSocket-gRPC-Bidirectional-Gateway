## Test

```
  █ TOTAL RESULTS

    checks_total.......: 178857  17869.163599/s
    checks_succeeded...: 100.00% 178857 out of 178857
    checks_failed......: 0.00%   0 out of 178857

    ✓ latency < 50ms

    CUSTOM
    msg_latency...........: avg=2.74ms min=0s     med=3ms    max=9ms    p(90)=4ms    p(95)=4ms

    EXECUTION
    iteration_duration....: avg=10s    min=10s    med=10s    max=10s    p(90)=10s    p(95)=10s
    iterations............: 50     4.995377/s
    vus...................: 50     min=50         max=50
    vus_max...............: 50     min=50         max=50

    NETWORK
    data_received.........: 22 MB  2.2 MB/s
    data_sent.............: 15 MB  1.5 MB/s

    WEBSOCKET
    ws_connecting.........: avg=3.37ms min=1.17ms med=2.45ms max=7.99ms p(90)=5.07ms p(95)=6.4ms
    ws_msgs_received......: 178857 17869.163599/s
    ws_msgs_sent..........: 178907 17874.158976/s
    ws_session_duration...: avg=10s    min=10s    med=10s    max=10s    p(90)=10s    p(95)=10s
    ws_sessions...........: 50     4.995377/s
```

| Metric | gRPC only (ghz) | WebSocket+gRPC (k6) | Delta |
|--------|----------------|---------------------|-------|
| req/sec | 34,970 | 17,869 | -17,101 |
| p50 | 0.93ms | 3ms | +2.07ms |
| p95 | 2.17ms | 4ms | +1.83ms |
| p99 | 2.56ms | ~6ms | +~3.44ms |
