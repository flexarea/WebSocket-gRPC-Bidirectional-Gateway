  ## Without Mesh

  █ TOTAL RESULTS

    checks_total.......: 118946  11877.953943/s
    checks_succeeded...: 100.00% 118946 out of 118946
    checks_failed......: 0.00%   0 out of 118946

    ✓ latency < 50ms

    CUSTOM
    msg_latency...........: avg=4.16ms min=257.79µs med=3.58ms max=29.64ms p(90)=7.12ms p(95)=8.93ms

    EXECUTION
    iteration_duration....: avg=10s    min=10s      med=10s    max=10.01s  p(90)=10s    p(95)=10s
    iterations............: 50     4.993003/s
    vus...................: 50     min=50         max=50
    vus_max...............: 50     min=50         max=50

    NETWORK
    data_received.........: 18 MB  1.8 MB/s
    data_sent.............: 14 MB  1.4 MB/s

    GRPC
    grpc_latency..........: avg=2.04ms min=86.32µs  med=1.8ms  max=16.69ms p(90)=3.62ms p(95)=4.45ms

    WEBSOCKET
    ws_connecting.........: avg=2.44ms min=721.22µs med=1.78ms max=12.35ms p(90)=3.85ms p(95)=4.31ms
    ws_msgs_received......: 118946 11877.953943/s
    ws_msgs_sent..........: 118996 11882.946946/s
    ws_session_duration...: avg=10s    min=10s      med=10s    max=10.01s  p(90)=10s    p(95)=10s
    ws_sessions...........: 50     4.993003/s

 ## With Mesh (side proxy)

 █ TOTAL RESULTS

    checks_total.......: 39127  3903.690542/s
    checks_succeeded...: 99.87% 39077 out of 39127
    checks_failed......: 0.12%  50 out of 39127

    ✗ latency < 50ms
      ↳  99% — ✓ 39077 / ✗ 50

    CUSTOM
    msg_latency...........: avg=12.73ms min=1.87ms med=12.26ms max=63.9ms  p(90)=17.84ms p(95)=19.63ms

    EXECUTION
    iteration_duration....: avg=10.01s  min=10.01s med=10.01s  max=10.02s  p(90)=10.01s  p(95)=10.01s
    iterations............: 50     4.988487/s
    vus...................: 50     min=50        max=50
    vus_max...............: 50     min=50        max=50

    NETWORK
    data_received.........: 5.9 MB 585 kB/s
    data_sent.............: 4.5 MB 451 kB/s

    GRPC
    grpc_latency..........: avg=9.5ms   min=1.38ms med=9.2ms   max=60.45ms p(90)=13.17ms p(95)=14.63ms

    WEBSOCKET
    ws_connecting.........: avg=15.36ms min=9.29ms med=15.26ms max=21.09ms p(90)=19.25ms p(95)=19.55ms
    ws_msgs_received......: 39127  3903.690542/s
    ws_msgs_sent..........: 39177  3908.679029/s
    ws_session_duration...: avg=10.01s  min=10.01s med=10.01s  max=10.02s  p(90)=10.01s  p(95)=10.01s
    ws_sessions...........: 50     4.988487/s

