package main

type EntityStatus uint8

const (
	StatusNone EntityStatus = iota
	Active
)

type EntityKind uint8

const (
	Pod EntityKind = iota
	StatefulSet
	Deployment
)

type EntityID uint64

type Position struct {
	x float64
	y float64
	z float64
}
