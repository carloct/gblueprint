package main

// Health ...
type Health struct {
	Version string `json:"version"`
}

// CurrentVersion ...
var CurrentVersion = Health{"0.1.1"}
