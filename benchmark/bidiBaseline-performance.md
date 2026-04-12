## gRPC Test

```
Summary:
  Count:	277929
  Total:	10.00 s
  Slowest:	11.11 ms
  Fastest:	0.28 ms
  Average:	1.51 ms
  Requests/sec:	27794.24

Response time histogram:
  0.284  [1]      |
  1.366  [134691] |∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎
  2.449  [122611] |∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎
  3.531  [19728]  |∎∎∎∎∎∎
  4.614  [780]    |
  5.696  [53]     |
  6.779  [21]     |
  7.861  [3]      |
  8.944  [1]      |
  10.026 [3]      |
  11.108 [1]      |

Latency distribution:
  10 % in 0.80 ms
  25 % in 0.96 ms
  50 % in 1.40 ms
  75 % in 1.99 ms
  90 % in 2.35 ms
  95 % in 2.58 ms
  99 % in 3.06 ms

Status code distribution:
  [OK]         277893 responses
  [Canceled]   36 responses

Error distribution:
  [36]   rpc error: code = Canceled desc = grpc: the client connection is closing
```

## Gateway Test
```

 █ TOTAL RESULTS

    checks_total.......: 186469  18625.359948/s
    checks_succeeded...: 100.00% 186469 out of 186469
    checks_failed......: 0.00%   0 out of 186469

    ✓ latency < 50ms

    CUSTOM
    msg_latency...........: avg=2.63ms min=0s     med=2ms    max=20ms    p(90)=5ms    p(95)=5ms

    EXECUTION
    iteration_duration....: avg=10s    min=10s    med=10s    max=10.01s  p(90)=10.01s p(95)=10.01s
    iterations............: 50     4.994224/s
    vus...................: 50     min=50         max=50
    vus_max...............: 50     min=50         max=50

    NETWORK
    data_received.........: 16 MB  1.6 MB/s
    data_sent.............: 16 MB  1.6 MB/s

    WEBSOCKET
    ws_connecting.........: avg=8.55ms min=6.35ms med=8.53ms max=10.85ms p(90)=9.88ms p(95)=10.34ms
    ws_msgs_received......: 186469 18625.359948/s
    ws_msgs_sent..........: 186519 18630.354172/s
    ws_session_duration...: avg=10s    min=10s    med=10s    max=10.01s  p(90)=10.01s p(95)=10.01s
    ws_sessions...........: 50     4.994224/s


running (10.0s), 00/50 VUs, 50 complete and 0 interrupted iterations
```

| Metric | Unary gRPC | Unary WS+gRPC | Delta | Bidi gRPC | Bidi WS+gRPC | Delta |
|--------|------------|---------------|-------|-----------|--------------|-------|
| req/sec | 34,970 | 17,869 | -17,101 | 27,794 | 18,625 | -9,169 |
| p50 | 0.93ms | 3ms | +2.07ms | 1.40ms | 2ms | +0.60ms |
| p95 | 2.17ms | 4ms | +1.83ms | 2.58ms | 5ms | +2.42ms |
| p99 | 2.56ms | ~6ms | +~3.44ms | 3.06ms | ~7ms | +~3.94ms |

