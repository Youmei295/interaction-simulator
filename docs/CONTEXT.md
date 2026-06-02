# Interaction Simulator Context

## What is this project?
The Interaction Simulator is an Agent-Based Simulation (ABS) platform that models the Iterated Prisoner's Dilemma over various graph networks. It is deeply inspired by "The Evolution of Trust" game, aiming to explore how different strategies (like cooperating, cheating, or retaliating) perform and evolve when agents interact repeatedly.

## Core Concepts
- **Nodes (Agents):** Independent actors in the simulation. Each has a specific `Strategy` that dictates its behavior.
- **Edges (Network):** Communication lines between nodes. Nodes only interact with their immediate neighbors.
- **Ticks/Generations:** The simulation runs in discrete steps. In every tick, each connected pair of nodes plays exactly one round of the Prisoner's Dilemma.
- **Payoff Matrix:** Standard Prisoner's Dilemma payoffs (Temptation > Reward > Punishment > Sucker).

## Current Architecture
The project has recently been refactored into a highly modular architecture to support future expansions (like evolutionary mechanics, new topologies, and dozens of strategies).

- `cmd/server/main.go`: The main entry point that wires the application together.
- `internal/core/`: Contains the fundamental data structures (`Node`, `Edge`, `Graph`, `Action`).
- `internal/strategy/`: Implements the Strategy Registry pattern. New strategies are added here as isolated files, preventing the need to modify core logic.
- `internal/simulator/`: Contains the actual simulation engines. Currently holds `v1_static` (a basic engine where nodes just accumulate points). Future engines (like `v2_evo` with reproduction/elimination) will live here alongside it.
- `internal/topology/`: Helpers for generating different network shapes (Ring, Fully Connected, etc.).
- `internal/api/`: HTTP REST handlers for the frontend.
- `static/`: A custom-built, glassmorphic frontend utilizing D3.js for physics-based force-directed graph rendering.

## Important Documentation
- `docs/BLUEPRINT.md`: The original MVP specification.
- `docs/RULES_AND_VERSIONS.md`: The living document detailing the exact rules of the current version and a roadmap of planned future variations.

## How to Run
```bash
cd cmd/server
go run main.go
```
Then navigate to `http://localhost:8081` in your browser.
