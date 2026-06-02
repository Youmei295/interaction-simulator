# Game Rules and Version History

## Introduction
This document outlines the core rules of the Interaction Simulator and tracks the evolution of the game's mechanics across different versions. It serves as a living document to track current implementation details and planned future variations.

---

## Version 0.1.0: Base Graph-Network Prisoner's Dilemma (Current MVP)

### Core Mechanics
- **Environment**: A static graph network. Nodes represent agents, and edges represent communication lines.
- **Interaction Loop**: In a single simulation cycle (tick), every connected pair of nodes interacts exactly once.
- **Decision Making**: Simultaneous execution. Neither node knows the other's action before making their own choice.
- **Memory Constraint**: Nodes only remember the *last action* played by each immediate neighbor.

### Payoff Matrix
The game uses a standard Iterated Prisoner's Dilemma payoff structure:

| Node A Action | Node B Action | Payoff A | Payoff B | Outcome Condition |
| :--- | :--- | :--- | :--- | :--- |
| **Cooperate** | **Cooperate** | +2 | +2 | Mutual Cooperation (Win-Win) |
| **Defect** | **Cooperate** | +3 | -1 | Temptation Reward / Sucker Payoff |
| **Cooperate** | **Defect** | -1 | +3 | Sucker Payoff / Temptation Reward |
| **Defect** | **Defect** | 0 | 0 | Mutual Punishment (Lose-Lose) |

### Active Strategies
1. **AlwaysCooperator**: Unconditionally returns "Cooperate".
2. **AlwaysCheater**: Unconditionally returns "Defect".
3. **Copycat (Tit-for-Tat)**: Cooperates on the first interaction with a neighbor. Thereafter, it copies the exact action the neighbor played in the previous tick.
4. **Grudger (Grim Trigger)**: Cooperates initially, but if the opponent ever defects, it will defect forever against that specific opponent.
5. **Detective**: Probes the opponent with a fixed sequence (Cooperate, Cooperate, Defect, Cooperate). If the opponent never fights back during the probe, it exploits them by always defecting. If the opponent fights back, it switches to acting like a Copycat.

---

## Planned Future Variations

### Version 0.2.x: Evolutionary Mechanics & Generations
* **Generations**: The concept of continuous ticks is grouped into discrete generations (e.g., a generation = 10 ticks).
* **Reproduction (Fitness)**: At the end of a generation, nodes with higher cumulative scores reproduce or spread their strategy to neighboring nodes.
* **Elimination**: Nodes with the lowest scores have their strategies replaced by the highest-performing strategies in their local neighborhood or the global population.

### Version 0.3.x: Noise and Miscommunication
* **Execution Errors**: Introducing a small probability error rate (e.g., 5%) where an intended action is accidentally flipped (intending to Cooperate but accidentally Defecting).
* **Purpose**: This variation tests the fragility of strict strategies like `Copycat` (which can fall into endless retaliation loops) and necessitates the introduction of forgiving strategies.

### Version 0.4.x: Expanded Strategy Roster
* **Forgiving Copycat (Tit-for-Two-Tats)**: Only retaliates with a defection if the opponent defects twice in a row.
* **Random**: Plays Cooperate or Defect with a 50/50 probability, acting as unpredictable noise in the ecosystem.
* **Pavlov (Win-Stay, Lose-Shift)**: Repeats its previous action if the previous payoff was good (+2 or +3), but switches its action if the payoff was bad (0 or -1).

### Version 0.5.x: Dynamic Topologies & Spatial Mobility
* **Edge Rewiring (Ostracism)**: Nodes gain the ability to proactively sever ties with chronic defectors and randomly form new edges with other nodes.
* **Spatial Grids**: Moving away from abstract graphs to a 2D grid/plane where nodes wander and only interact when their radii overlap.
* **Resource Scarcity**: Modifying the payoff matrix to scale dynamically based on the local density of nodes (e.g., overcrowding reduces the rewards of mutual cooperation).
