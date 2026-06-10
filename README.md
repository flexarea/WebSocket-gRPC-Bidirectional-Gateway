## Introduction

Real-time data streaming is a crucial component of modern network applications. However, maintaining persistent, low-latency communication between web clients and distributed backend services introduces complex state and concurrency challenges. This independent project presents a bidirectional messaging gateway architecture that bridges client-facing WebSockets with backend gRPC streams, implemented in Go and deployed on Kubernetes using k3s. This study quantifies the latency overhead introduced at both the WebSocket and gRP layers under unary and bidirectional streaming patterns, with and without a traffic-handling proxy for container applications (Envoy)
intercepting inter-service traffic.

![Gateway image](./assets/ws-rpc.png)

## Gateway & Protocols

WebSocket (WS) is a browser-native HTTP/1.1 protocol for full-duplex communication, while gRPC is an HTTP/2 RPC framework. The gateway bridges these incompatible protocols, allowing WS to use gRPC services.

Benchmarks were conducted using Grafana k6 with 50 concurrent virtual users (VUs) over a total duration of 10 seconds. Both baseline architectures were implemented and evaluated on the Middlebury College Ada cluster, while all K3s-related components were deployed and tested on Amazon EC2.

## Baseline Architecture

![Baseline architecture](./assets/baseline_architecture.png)

<small> A hub manager registers/unregister a new WS client. During a client registration, two goroutines (read/write pump) are created to handle concurrent send/receive. For Bidi mode, the architecture is similar  but includes a new goroutine for receiving streams. </small>

## Observed Latency

* The bidirectional streaming architecture demonstrates a 74.85% reduction in gateway overhead and an 88.9% reduction in end-to-end ping-pong latency compared to the unary implementation at median (p50).
* We observe a higher throughput of 21,066/s with the bidirectional architecture, compared to 13,939/s for the unary approach, an increase of approximately 51.2%.

This performance gap is likely attributed to bidirectional streaming's single connection establishment with one-time HTTP/2 header frame negotiation, followed by length-prefixed message transmission terminated by a single end-of-stream (EOS) flag. In contrast, unary calls incur per-request header frame overhead and an individual EOS for every transmitted message [1]. However, further testing would be required to confirm the extent to which these factors contribute to the observed difference.
