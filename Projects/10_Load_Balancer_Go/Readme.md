# Load Balancer built using Go Programming Language

This project demonstrates a simple **Round-Robin Load Balancer** implemented in **Go**. The load balancer proxies incoming HTTP requests to multiple backend servers.

## Features
- **Round-robin load balancing**: Distributes incoming requests evenly across available servers.
- **Reverse Proxy**: Forwards requests to backend servers using a reverse proxy.

## Round-Robin Load Balancing

- The load balancer uses a **round-robin** strategy to distribute incoming requests evenly across the available servers.
- The `getNextAvailableServer()` function ensures that requests are sent to the next server in the list, wrapping around once all servers have been used.

## Example Code Structure

### `main.go`
- The entry point for the load balancer.
- Defines the `Server` interface and `myServer` struct that represents backend servers.
- The `LoadBalancer` struct handles incoming requests and forwards them to backend servers using a round-robin method.

### `utils.go` (optional)
- Utility functions like `HandleError` for error handling. (If you implement custom error handling in the `utils` package.)

## Example Server Configuration

You can modify the list of backend servers in the `main.go` file. Just replace the example URLs with your own backend server URLs:
