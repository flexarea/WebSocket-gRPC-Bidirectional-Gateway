## Full Benchmark Results (k3s environment)

### Unary

| Metric | No Mesh | Sidecar | Overhead |
|--------|---------|---------|----------|
| req/sec | 11,878 | 3,904 | -7,974 (-67%) |
| p50 | 3.58ms | 12.26ms | +8.68ms (+242%) |
| p90 | 7.12ms | 17.84ms | +10.72ms (+151%) |
| p95 | 8.93ms | 19.63ms | +10.70ms (+120%) |
| gRPC layer p50 | 1.80ms | 9.20ms | +7.40ms (+411%) |
| WS connecting | 2.44ms | 15.36ms | +12.92ms (+530%) |

### Bidi

| Metric | No Mesh | Sidecar | Overhead |
|--------|---------|---------|----------|
| req/sec | 12,573 | 7,047 | -5,526 (-44%) |
| p50 | 2.32ms | 6.20ms | +3.88ms (+167%) |
| p90 | 6.20ms | 10.67ms | +4.47ms (+72%) |
| p95 | 8.27ms | 12.43ms | +4.16ms (+50%) |
| gRPC layer p50 | 0.400ms | 3.76ms | +3.36ms (+840%) |
| WS connecting | 3.33ms | 20.03ms | +16.70ms (+501%) |
