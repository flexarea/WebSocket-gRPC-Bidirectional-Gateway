```
     execution: local
        script: gatewayBenchMark.js
        output: -

     scenarios: (100.00%) 1 scenario, 50 max VUs, 40s max duration (incl. graceful stop):
              * default: 50 looping VUs for 10s (gracefulStop: 30s)



  █ TOTAL RESULTS

    checks_total.......: 47162   4711.069055/s
    checks_succeeded...: 100.00% 47162 out of 47162
    checks_failed......: 0.00%   0 out of 47162

    ✓ latency < 50ms

    CUSTOM
    msg_latency...........: avg=3.08ms min=0s       med=3ms    max=11ms  p(90)=4ms    p(95)=5ms

    EXECUTION
    iteration_duration....: avg=10s    min=10s      med=10s    max=10s   p(90)=10s    p(95)=10s
    iterations............: 50     4.99456/s
    vus...................: 50     min=50        max=50
    vus_max...............: 50     min=50        max=50

    NETWORK
    data_received.........: 5.6 MB 557 kB/s
    data_sent.............: 4.0 MB 397 kB/s

    WEBSOCKET
    ws_connecting.........: avg=2.96ms min=549.64µs med=3.04ms max=4.1ms p(90)=3.74ms p(95)=3.86ms
    ws_msgs_received......: 47162  4711.069055/s
    ws_msgs_sent..........: 47173  4712.167858/s
    ws_session_duration...: avg=10s    min=10s      med=10s    max=10s   p(90)=10s    p(95)=10s
    ws_sessions...........: 50     4.99456/s

```
