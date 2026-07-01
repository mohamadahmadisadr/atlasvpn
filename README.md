# AtlasVPN

An educational Go project for exploring network programming and the fundamentals of VPNs. The codebase currently implements a minimal UDP-based client-server transport system, serving as a foundation for building a virtual private network.


## Getting Started

### Prerequisites
- Go 1.26+

### Run the Server
Start the server to listen for incoming UDP packets:
```bash
go run cmd/server
```
By default, the server binds to `127.0.0.1:3000`.

### Run the Client
In a separate terminal, run the client to send a test message and receive the server's confirmation:
```bash
go run cmd/client
```
