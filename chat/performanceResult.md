### Test 1: Concurrency Pressure

--concurrency 50 \
--total 10000 \

```
Summary:
  Count:	10000
  Total:	1.12 s
  Slowest:	15.40 ms
  Fastest:	0.31 ms
  Average:	3.37 ms
  Requests/sec:	8967.55

Response time histogram:
  0.314  [1]    |
  1.822  [1032] |∎∎∎∎∎∎∎∎∎∎
  3.330  [4128] |∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎
  4.838  [3656] |∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎
  6.347  [965]  |∎∎∎∎∎∎∎∎∎
  7.855  [167]  |∎∎
  9.363  [30]   |
  10.871 [14]   |
  12.379 [6]    |
  13.888 [0]    |
  15.396 [1]    |

Latency distribution:
  10 % in 1.80 ms
  25 % in 2.44 ms
  50 % in 3.27 ms
  75 % in 4.20 ms
  90 % in 4.99 ms
  95 % in 5.55 ms
  99 % in 7.04 ms

Status code distribution:
  [OK]   10000 responses

```

### Test 2: Concurrency Pressure part II

--concurrency 200 \
--total 10000 \

```
Summary:
  Count:	10000
  Total:	893.96 ms
  Slowest:	38.82 ms
  Fastest:	1.82 ms
  Average:	15.73 ms
  Requests/sec:	11186.14

Response time histogram:
  1.824  [1]    |
  5.524  [52]   |
  9.223  [33]   |
  12.923 [308]  |∎∎
  16.622 [6616] |∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎
  20.322 [2935] |∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎
  24.021 [24]   |
  27.721 [11]   |
  31.420 [11]   |
  35.119 [6]    |
  38.819 [3]    |

Latency distribution:
  10 % in 13.72 ms
  25 % in 14.70 ms
  50 % in 15.82 ms
  75 % in 16.82 ms
  90 % in 17.77 ms
  95 % in 18.40 ms
  99 % in 19.74 ms

Status code distribution:
  [OK]   10000 responses
```

### Test 3: Request Pressure
--concurrency 50 \
--total 100000 \

```
Summary:
  Count:	100000
  Total:	4.49 s
  Slowest:	11.67 ms
  Fastest:	0.22 ms
  Average:	1.61 ms
  Requests/sec:	22276.21

Response time histogram:
  0.216  [1]     |
  1.362  [32344] |∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎
  2.507  [59370] |∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎
  3.652  [7958]  |∎∎∎∎∎
  4.797  [246]   |
  5.943  [39]    |
  7.088  [19]    |
  8.233  [11]    |
  9.378  [9]     |
  10.524 [1]     |
  11.669 [2]     |

Latency distribution:
  10 % in 0.76 ms
  25 % in 1.22 ms
  50 % in 1.60 ms
  75 % in 1.94 ms
  90 % in 2.40 ms
  95 % in 2.73 ms
  99 % in 3.23 ms

Status code distribution:
  [OK]   100000 responses
```

