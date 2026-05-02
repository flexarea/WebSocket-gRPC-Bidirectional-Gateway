## With Mesh (sidecar proxy)

```

  █ TOTAL RESULTS

    checks_total.......: 70687  7047.449274/s
    checks_succeeded...: 99.97% 70672 out of 70687
    checks_failed......: 0.02%  15 out of 70687

    ✗ latency < 50ms
      ↳  99% — ✓ 70672 / ✗ 15

    CUSTOM
    msg_latency...........: avg=6.78ms  min=314.62µs med=6.2ms   max=52.13ms p(90)=10.67ms p(95)=12.43ms

    EXECUTION
    iteration_duration....: avg=10.02s  min=10.01s   med=10.02s  max=10.02s  p(90)=10.02s  p(95)=10.02s
    iterations............: 50     4.984968/s
    vus...................: 50     min=50        max=50
    vus_max...............: 50     min=50        max=50

    NETWORK
    data_received.........: 7.9 MB 784 kB/s
    data_sent.............: 8.1 MB 812 kB/s

    GRPC
    grpc_latency..........: avg=4.15ms  min=102.18µs med=3.76ms  max=37.35ms p(90)=6.85ms  p(95)=8.2ms

    WEBSOCKET
    ws_connecting.........: avg=20.03ms min=15.18ms  med=20.31ms max=23.96ms p(90)=21.38ms p(95)=21.65ms
    ws_msgs_received......: 70687  7047.449274/s
    ws_msgs_sent..........: 70737  7052.434242/s
    ws_session_duration...: avg=10.02s  min=10.01s   med=10.02s  max=10.02s  p(90)=10.02s  p(95)=10.02s
    ws_sessions...........: 50     4.984968/s
```
