## gRPC Test

```
Summary:
  Count:	349681
  Total:	10.00 s
  Slowest:	9.48 ms
  Fastest:	0.27 ms
  Average:	1.17 ms
  Requests/sec:	34970.10

Response time histogram:
  0.270 [1]      |
  1.191 [220967] |∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎
  2.112 [107311] |∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎
  3.033 [20691]  |∎∎∎∎
  3.954 [556]    |
  4.875 [84]     |
  5.796 [17]     |
  6.717 [5]      |
  7.638 [2]      |
  8.559 [5]      |
  9.480 [3]      |

Latency distribution:
  10 % in 0.66 ms 
  25 % in 0.75 ms 
  50 % in 0.93 ms 
  75 % in 1.58 ms 
  90 % in 1.97 ms 
  95 % in 2.17 ms 
  99 % in 2.56 ms 

Status code distribution:
  [OK]            349642 responses   
  [Canceled]      4 responses        
  [Unavailable]   35 responses       

Error distribution:
  [4]    rpc error: code = Canceled desc = grpc: the client connection is closing                                                                      
  [35]   rpc error: code = Unavailable desc = error reading from server: read tcp 127.0.0.1:51532->127.0.0.1:50051: use of closed network connection   
```
## Gateway Test

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
