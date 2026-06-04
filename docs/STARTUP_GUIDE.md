# Interaction Simulator: Startup Guide

Welcome to the Interaction Simulator! This guide will walk you through the steps to get the project up and running on your local machine.

## Prerequisites

Before you begin, ensure you have the following installed on your system:
- **Go (Golang)**: The backend is built with Go. You can download and install it from [golang.org](https://go.dev/). (Version 1.18+ recommended)
- **Git**: To clone the repository (if you haven't already).
- **A Modern Web Browser**: To view the frontend interface (Chrome, Firefox, Edge, Safari, etc.).

## Installation & Setup

1. **Open your terminal or command prompt.**
2. **Navigate to the project root directory:**
   ```bash
   cd path/to/interaction-simulator
   ```
3. **Ensure Go modules are tidy and dependencies are downloaded:**
   ```bash
   go mod tidy
   ```

## Running the Simulator

The simulator consists of a Go backend that serves both the API and the static frontend files.

1. **Navigate to the server entry point directory:**
   ```bash
   cd cmd/server
   ```
2. **Run the Go server:**
   ```bash
   go run main.go
   ```
   *Alternatively, you can build it into an executable and run it:*
   ```bash
   go build -o simulator.exe .
   ./simulator.exe
   ```
3. **Wait for the startup message:**
   You should see output in your terminal indicating the server has started successfully:
   ```text
   Server starting on http://localhost:8081
   ```

## Using the Web Interface

Once the server is running, you can access the simulator's UI:

1. **Open your web browser** and navigate to: [http://localhost:8082](http://localhost:8082)
2. **The Interface:**
   - **Main Canvas:** You will see a force-directed graph representing the network of agents (nodes). Each node is colored according to its strategy.
   - **Control Panel (Right Side):**
     - **Advance 1 Tick:** Click this button to run one cycle of the simulation. All connected nodes will play the Prisoner's Dilemma against their neighbors, updating their scores and memories.
     - **Graph Topology:** Use the dropdown to select between different network shapes (e.g., Ring Network, Fully Connected).
     - **Population Size:** Adjust the number of nodes in the network (e.g., 12).
     - **Reset & Generate:** Click this to apply your topology and size settings, completely resetting the simulation state.
   - **Legend:** Shows which color corresponds to which strategy (e.g., Green for Always Cooperate, Red for Always Cheat).
3. **Interactivity:**
   - You can drag nodes around the canvas to inspect the network visually.
   - Hover over any node to see a tooltip detailing its specific ID, its current strategy, and its accumulated score.

## Troubleshooting

- **Port Conflict:** If you get an error saying the port is already in use, make sure you don't have another instance of the server running. You can change the port in `cmd/server/main.go` if necessary.
- **UI Not Loading:** Ensure you are running the `go run main.go` command from exactly the `cmd/server` directory. The server relies on a relative path to serve the `static/` directory (it looks for `../../static`).

Happy Simulating!
