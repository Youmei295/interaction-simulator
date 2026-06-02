# Specification: Base Graph-Network Prisoner's Dilemma Simulator (MVP)

## 1. System Overview
This is a lightweight Agent-Based Simulation (ABS) running the Iterated Prisoner's Dilemma over a static graph network. Nodes are fixed in place, edges are unweighted, and every connected pair interacts exactly once per simulation cycle.

---

## 2. Core Entities & Data Structures

### 2.1 Node (Agent)
A Node represents an independent worker operating within the network.
* **Fields:**
  * `ID`: Unique identifier string (e.g., `"node_01"`).
  * `Strategy`: The algorithmic logic type controlling game choices.
  * `Score`: A floating-point number tracking accumulated utility.
  * `Memory`: A simple dictionary/map tracking the *last action* played by immediate neighbors.
* **Memory Schema:**
  * `Key`: Neighbor Node ID (`string`)
  * `Value`: The last Action the neighbor executed (`"Cooperate"` or `"Defect"`).

### 2.2 Edge (Connection)
An Edge represents a static, bidirectional communication line between two nodes.
* **Fields:**
  * `NodeA`: Identifier string of the first node.
  * `NodeB`: Identifier string of the second node.

---

## 3. The Execution Lifecycle (The Tick Loop)

The simulator runs inside a sequential loop. Every simulation `tick` must execute the following distinct runtime phases:

### Phase 1: Queue Generation
1. Initialize an empty dynamic array/slice of Node Pairs: `TransactionQueue = []`.
2. Iterate through all active `Edges` in the system.
3. For each edge, append the pair `(NodeA, NodeB)` into the `TransactionQueue` exactly **once**.

### Phase 2: Game Execution
1. Shuffle the `TransactionQueue` to randomize execution order (preventing sequence bias).
2. For each pair popped from the queue, evaluate their chosen actions.
3. Node A and Node B query their respective `Strategy` functions to output `"Cooperate"` or `"Defect"`.
4. Actions must be evaluated simultaneously. Neither node can see the other's current choice beforehand.

### Phase 3: State Resolution & Payoff Matrix
Calculate the utility output based on the standard Prisoner's Dilemma matrix:

| Node A Action | Node B Action | Payoff A | Payoff B | Outcome Condition |
| :--- | :--- | :--- | :--- | :--- |
| **Cooperate** | **Cooperate** | +2 | +2 | Mutual Cooperation (Win-Win) |
| **Defect** | **Cooperate** | +3 | -1 | Temptation Reward / Sucker Payoff |
| **Cooperate** | **Defect** | -1 | +3 | Sucker Payoff / Temptation Reward |
| **Defect** | **Defect** | 0 | 0 | Mutual Punishment (Lose-Lose) |

1. Increment/Decrement each node's `Score` property by their respective payoff value.
2. Update each node's `Memory` map with the action the opponent just played (overwriting the previous value).

---

## 4. Initial Behavioral Strategies
The engine must initially support three baseline algorithms for its nodes:

* **Always Cooperator (`AlwaysCooperator`):** * Returns `"Cooperate"` under all conditions.
* **Always Cheater (`AlwaysCheater`):** * Returns `"Defect"` under all conditions.
* **Copycat / Tit-for-Tat (`Copycat`):**
  * Checks the `Memory` map for the opponent's ID.
  * If no history exists (first interaction) $\rightarrow$ Returns `"Cooperate"`.
  * If history exists $\rightarrow$ Returns the exact action the opponent played in the previous tick.

---

## 5. Agent Instructions for Implementation
When initiating code generation for this MVP, adhere to the following constraints:
1. **Modular Architecture:** Build the engine with a clean separation of concerns. Use the Strategy Registry pattern to allow adding new strategies without modifying core logic. Separate core interfaces (`internal/core`), strategies (`internal/strategy`), simulation logic (`internal/simulator`), and API handlers (`internal/api`).
2. **Concurrency Safety:** Even in this MVP, assume the execution phase might be parallelized in the future. State mutations (updating a Node's `Score` and `Memory/Meta`) must be handled safely using mutexes to avoid race conditions.
3. **Graph Initialization:** Provide a helper function in a dedicated `topology` package to easily generate standard graph topologies (e.g., a fully connected graph, a ring graph) to facilitate immediate testing.