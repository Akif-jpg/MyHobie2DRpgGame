package engine

import (
	"time"

	"github.com/Akif-jpg/MyHobieMMORPGGame/services/engine/entities"
)

type Engine struct {
	// Add fields for game state, such as players, NPCs, items, etc.
	entitiyList []entities.Entity
}

func NewEngine() *Engine {
	return &Engine{
		make([]entities.Entity, 0),
	}
}

func (e *Engine) AddEntity(entity entities.Entity) {
	e.entitiyList = append(e.entitiyList, entity)
}

func (e *Engine) Start() {
	go e.update()
}

// Update the game state based on player actions and other events
func (e *Engine) update() {
	// Implement game logic here
	lastUpdateTime := time.Now()

	for {
		now := time.Now()
		deltaTime := now.Sub(lastUpdateTime).Seconds()
		lastUpdateTime = now

		// Process player actions and other events here

		// Example: Update NPCs, handle combat, etc.
		//
		// Make entitiy list pre updates
		for _, entity := range e.entitiyList {
			entity.PreUpdate(deltaTime)
		}

		// Make entitiy list updates
		for _, entity := range e.entitiyList {
			entity.Update(deltaTime)
		}

		// Make entitiy list last updates
		for _, entity := range e.entitiyList {
			entity.LastUpdate(deltaTime)
			if !entity.IsAlive() {
				// Handle entity death, such as removing it from the game world, dropping loot, etc.
			}
		}
	}

	// Update game state here
}
