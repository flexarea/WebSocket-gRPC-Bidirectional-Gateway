
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
