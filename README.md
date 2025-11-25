# Tech Test API

A simple Go HTTP API using standard library for basic mathematical operations.

## Prerequisites

- Golang 1.24

## How to Run

2. Run the server:
   ```bash
   go run cmd/main.go
   ```

   The server will start on port 8080.

## API Endpoints

### Ping
Check if the server is running:
```bash
curl http://localhost:8080/ping
```
Returns: `pong`

### Addition
Add two numbers together:
```bash
curl "http://localhost:8080/add?a=10&b=5"
```
Returns: `15.00`

### Subtraction
Subtract b from a:
```bash
curl "http://localhost:8080/sub?a=10&b=3"
```
Returns: `7.00`

## Example Usage

```bash
# Start the server
go run cmd/main.go

# Test addition
curl "http://localhost:8080/add?a=15&b=25"
# Returns: 40.00

# Test subtraction
curl "http://localhost:8080/sub?a=20&b=8"
# Returns: 12.00

# Test with decimals
curl "http://localhost:8080/add?a=3.5&b=2.1"
# Returns: 5.60
```