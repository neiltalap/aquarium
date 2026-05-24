package main

import (
	"fmt"
	"time"
)

// tracking positions of all of the entities
type PositionStore struct {
	entityToIndex map[EntityID]uint64
	entities      map[uint64]EntityID
	positions     []Position
}

// tracking relationships between events and entities
type EventRelationshipStore struct {
	eventToIndex map[EventID]uint64
	events       map[uint64]EventID
	subject      []EntityID
	object       []EntityID
}

// tracking events
type EventStore struct {
	indexToEvent map[EventID]uint64
	events       map[uint64]EventID
	eventType    []EventType
	occurredAt   []time.Time
}

type EntityStatusStore struct {
	entityToIndex  map[EntityID]uint64
	entities       map[uint64]EntityID
	entityStatuses []EntityStatus
}

type EntityKindStore struct {
	entityToIndex map[EntityID]uint64
	entities      map[uint64]EntityID
	EntityKind    []EntityKind
}

type World struct {
	sizeX                  uint64
	sizeY                  uint64
	sizeZ                  uint64
	currentTime            time.Time
	positionStore          PositionStore
	eventRelationshipStore EventRelationshipStore
	eventStore             EventStore
}

func main() {
	fmt.Println("Hello!")

	// main loop
	for {

	}
}
