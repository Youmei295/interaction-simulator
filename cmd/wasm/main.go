//go:build js && wasm
package main

import (
	"encoding/json"
	"syscall/js"

	"interaction-simulator/internal/core"
	"interaction-simulator/internal/simulator/v2_evo"
)

var sim core.Simulator

func getState(this js.Value, args []js.Value) any {
	state := sim.GetState()
	bytes, err := json.Marshal(state)
	if err != nil {
		return `{"error": "failed to marshal state"}`
	}
	return string(bytes)
}

func advanceTick(this js.Value, args []js.Value) any {
	sim.AdvanceTick()
	state := sim.GetState()
	bytes, err := json.Marshal(state)
	if err != nil {
		return `{"error": "failed to marshal state"}`
	}
	return string(bytes)
}

func resetSim(this js.Value, args []js.Value) any {
	if len(args) == 0 {
		return `{"error": "missing config"}`
	}
	
	configStr := args[0].String()
	var config core.SimConfig
	if err := json.Unmarshal([]byte(configStr), &config); err != nil {
		return `{"error": "invalid config"}`
	}

	// Apply defaults
	if config.Distribution == nil {
		config.Distribution = map[string]float64{"AlwaysCooperator": 1.0}
	}
	if config.Size <= 0 {
		config.Size = 12
	}

	sim.Reset(config)
	state := sim.GetState()
	bytes, err := json.Marshal(state)
	if err != nil {
		return `{"error": "failed to marshal state"}`
	}
	return string(bytes)
}

func main() {
	// Create a channel to block the main function
	// Go WebAssembly requires the main function to never return
	c := make(chan struct{}, 0)

	// Initialize the simulation with default parameters
	sim = v2_evo.NewEngine()
	sim.Reset(core.SimConfig{
		Topology: "Ring",
		Size:     12,
		Distribution: map[string]float64{
			"AlwaysCooperator": 0.33,
			"AlwaysCheater":    0.33,
			"Copycat":          0.34,
		},
	})

	// Expose functions to JavaScript
	js.Global().Set("sim_getState", js.FuncOf(getState))
	js.Global().Set("sim_advanceTick", js.FuncOf(advanceTick))
	js.Global().Set("sim_reset", js.FuncOf(resetSim))

	<-c
}
