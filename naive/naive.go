package main

// Request
type LookupRequest struct {
	Key  string
	Hops uint
}

// Local lookup
type LocalMap struct {
	Lookup    map[string]string
	Neighbors map[uint64]uint64
}

// ...
type Hash interface {
	get(string) string
	put(string) string
}

// ...
type P2PNetwork interface {
	ping(uint64) bool
	neighbors() []uint64
}
