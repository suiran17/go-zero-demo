# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Project Overview

This is a go-zero framework demo project containing examples of API services, RPC services, and microservice architectures. The project is structured to demonstrate different service patterns using the go-zero framework.

## Architecture

### Project Structure
- `api-demo/` - HTTP API service examples using go-zero's API framework
- `rpc-demo/` - gRPC service examples using go-zero's RPC framework  
- `service-demo/` - Complete service examples including both single-service and micro-service patterns
- `测试api/` - Test API examples (Chinese directory)

### Key Components

**API Services**: Use `.api` files to define HTTP endpoints and generate code with `goctl api`
- Structure: `internal/handler/` (HTTP handlers), `internal/logic/` (business logic), `internal/types/` (request/response structs)
- Configuration: YAML files in `etc/` directory (e.g., `greet-api.yaml`)

**RPC Services**: Use `.proto` files to define gRPC services and generate code with `goctl rpc`
- Structure: `internal/server/` (gRPC servers), `internal/logic/` (business logic), `pb.go` files for protobuf
- Configuration: YAML files in `etc/` directory with etcd configuration

**Service Context**: All services use `internal/svc/servicecontext.go` for dependency injection

## Common Commands

### API Development
```bash
# Generate API code from .api file
goctl api go -api greet.api -dir . -style gozero

# Format API files
goctl api format
```

### RPC Development
```bash
# Generate RPC code from .proto file
goctl rpc protoc helllo.proto --go_out=. --go-grpc_out=. --zrpc_out=.
```

### Running Services
```bash
# Run API service (typically on port 8888)
go run greet.go

# Run RPC service (typically on port 8080)
go run helllo.go
```

### Testing
```bash
# Test API endpoints
curl localhost:8888/from/you

# Use .http files for testing (present in most service directories)
```

## Configuration

- API services use `Host` and `Port` in YAML config
- RPC services use `ListenOn` and `Etcd` configuration for service discovery
- Default ports: API services on 8888, RPC services on 8080

## Go Workspace

The `service-demo/` directory uses Go workspaces with `go.work` file including:
- `./micro-service`
- `./single-service` 
- `./study`

## Framework Dependencies

Primary dependency: `github.com/zeromicro/go-zero v1.6.1`

The project demonstrates go-zero's code generation capabilities, service discovery with etcd, and standard microservice patterns.