# Rate Limiter

A high-performance rate limiter service that controls the number of client requests allowed over a specified period using the token bucket algorithm. Excess requests are blocked to prevent abuse and ensure fair resource usage.

## Features

- **Token Bucket Algorithm**: Implements rate limiting with configurable rate and capacity.
- **Dual Protocol Support**: Provides both gRPC and HTTP APIs for flexibility.
- **Middleware Integration**: Includes middleware for seamless integration into gRPC and HTTP servers.
- **Configurable**: Environment-based configuration using YAML files.
- **Docker Support**: Containerized deployment with Docker and Docker Compose.
- **Thread-Safe**: Concurrent request handling with mutex-protected token bucket.

## Tech Stack

- **Language**: Go 1.21.4
- **Protocols**: gRPC, HTTP
- **Libraries**:
  - Gorilla Mux (HTTP routing)
  - gRPC Ecosystem Middleware
  - Viper (Configuration management)
  - Protobuf (gRPC message definitions)
- **Containerization**: Docker, Docker Compose

## Installation Steps

1. **Clone the repository**:
   ```bash
   git clone https://github.com/shashankbiet/rate-limiter.git
   cd rate-limiter
   ```

2. **Install Go dependencies**:
   ```bash
   go mod tidy
   ```

3. **Install Protocol Buffers compiler** (if not already installed):
   - For macOS: `brew install protobuf`
   - For other systems, follow [official instructions](https://grpc.io/docs/protoc-installation/).

4. **Generate gRPC code** (if needed):
   ```bash
   protoc -I=proto --go_out=. --go_opt=module=github.com/shashankbiet/rate-limiter --go-grpc_out=require_unimplemented_servers=false:. --go-grpc_opt=module=github.com/shashankbiet/rate-limiter proto/**/*.proto
   ```

## Environment Variables

Configure the application using environment variables or YAML files in the `conf/` directory. Key variables include:

- `ENVIRONMENT`: Environment type (e.g., "local", "dev")
- `APP_NAME`: Application name (default: "rate-limiter")
- `HTTP_SERVER.PORT`: HTTP server port (default: 3001)
- `GRPC_SERVER.PORT`: gRPC server port (default: 9001)

Example `conf/dev.yaml`:
```yaml
ENVIRONMENT: "dev"
APP_NAME: "rate-limiter"
HTTP_SERVER:
  PORT: 3001
GRPC_SERVER:
  PORT: 9001
```

## Running the Application

### Using Docker Compose (Recommended)

1. Build and run the service:
   ```bash
   docker-compose up --build
   ```

2. The service will be available at:
   - HTTP: `http://localhost:3001`
   - gRPC: `localhost:9001`

### Using Go Directly

1. Run the application:
   ```bash
   go run main.go
   ```

2. Ensure configuration files are in the `conf/` directory.

## Usage Examples

### HTTP API

Get user details:
```bash
curl http://localhost:3001/api/user/122
```

Response (success):
```json
user: 122
```

### gRPC API

Use a gRPC client (e.g., grpcurl or a Go client) to call the ProductService.

Example using grpcurl:
```bash
grpcurl -plaintext localhost:9001 productpb.ProductService.GetProduct
```

Request payload:
```json
{
    "id": "244"
}
```

Response:
```json
{
    "product": {
        "id": "0",
        "name": "Product: 244",
        "description": "",
        "brand": "",
        "sku": 0,
        "price": 0,
        "weight": 0
    }
}
```

Note: Rate limiting is applied per request. If the rate limit is exceeded, requests will be blocked.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
