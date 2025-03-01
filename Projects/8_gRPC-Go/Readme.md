# gRPC-Go 

This repository demonstrates different types of gRPC streaming using Go.

## 🚀 Project Overview

This project includes:
- **Unary RPC**: Client sends a single request, server returns a single response.
- **Server Streaming RPC**: Client sends a single request, server streams multiple responses.
- **Client Streaming RPC**: Client streams multiple requests, server returns a single response.
- **Bidirectional Streaming RPC**: Both client and server stream data simultaneously.

---

## 📌 Prerequisites

Ensure you have the following installed:

- **Go 1.23+**
- **Protocol Buffers Compiler (`protoc`)**
- **Go gRPC Plugins (`protoc-gen-go` & `protoc-gen-go-grpc`)**

### Install Protocol Buffers Compiler (protoc)
Download and install `protoc` from:  
👉 [https://github.com/protocolbuffers/protobuf/releases](https://github.com/protocolbuffers/protobuf/releases)

### Install Go gRPC Plugins
```
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
```
