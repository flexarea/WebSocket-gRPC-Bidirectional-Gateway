## Test 1

```
     scenarios: (100.00%) 1 scenario, 50 max VUs, 40s max duration (incl. graceful stop):
              * default: 50 looping VUs for 10s (gracefulStop: 30s)



  █ TOTAL RESULTS

    checks_total.......: 102405  10236.296104/s
    checks_succeeded...: 100.00% 102405 out of 102405
    checks_failed......: 0.00%   0 out of 102405

    ✓ response received

    EXECUTION
    iteration_duration....: avg=4.85ms min=519.65µs med=4.19ms max=66.96ms p(90)=8.46ms p(95)=10.19ms
    iterations............: 102405 10236.296104/s
    vus...................: 50     min=50         max=50
    vus_max...............: 50     min=50         max=50

    NETWORK
    data_received.........: 65 MB  6.5 MB/s
    data_sent.............: 26 MB  2.6 MB/s

    WEBSOCKET
    ws_connecting.........: avg=3.03ms min=242.31µs med=2.3ms  max=26.05ms p(90)=6.3ms  p(95)=7.98ms
    ws_msgs_received......: 102405 10236.296104/s
    ws_msgs_sent..........: 102405 10236.296104/s
    ws_session_duration...: avg=4.8ms  min=489.95µs med=4.13ms max=66.86ms p(90)=8.4ms  p(95)=10.13ms
    ws_sessions...........: 102405 10236.296104/s

```



## Test 2

```
 █ TOTAL RESULTS

    checks_total.......: 2160096 182486.090225/s
    checks_succeeded...: 100.00% 2160096 out of 2160096
    checks_failed......: 0.00%   0 out of 2160096

    ✓ response received

    EXECUTION
    iteration_duration....: avg=5.85s  min=5.36s    med=5.8s   max=6.46s  p(90)=6.31s  p(95)=6.35s
    iterations............: 100     8.448055/s
    vus...................: 50      min=50          max=50
    vus_max...............: 50      min=50          max=50

    NETWORK
    data_received.........: 212 MB  18 MB/s
    data_sent.............: 58 MB   4.9 MB/s

    WEBSOCKET
    ws_connecting.........: avg=2.04ms min=424.32µs med=1.71ms max=9.78ms p(90)=3.52ms p(95)=4.27ms
    ws_msgs_received......: 2160096 182486.090225/s
    ws_msgs_sent..........: 1000000 84480.546339/s
    ws_session_duration...: avg=5.85s  min=5.36s    med=5.8s   max=6.46s  p(90)=6.31s  p(95)=6.35s
    ws_sessions...........: 100     8.448055/s
```
