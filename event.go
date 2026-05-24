package main

type EventType uint8

const (
	None EventType = iota
	Damage
	Spawn
	Displacement
)

type EventID uint64
