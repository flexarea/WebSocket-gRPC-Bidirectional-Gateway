## Introduction

Real-time data streaming is a crucial component of modern network applications. However, maintaining persistent,low-latency communication between web clients and distributed backend services introduces complex state and concurrency challenges. This study presents a bidirectional messaging gateway architecture that bridges client-facing WebSockets with backend gRPC streams, implemented in Go and deployed on Kubernetes using k3s. This study quantifies the latency overhead introduced at both the WebSocket and gRP layers under unary and bidirectional streaming patterns, with and without a traffic-handling proxy for container applications (Envoy)
intercepting inter-service traffic.

## Gateway & Protocols


<img width="2139" height="1048" alt="bidi" src="https://github.com/user-attachments/assets/f5cffdc1-ed9f-4fe0-a2d5-206bb67e4792" />
