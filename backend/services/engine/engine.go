package engine

import (
	"context"
	"time"

	"github.com/Akif-jpg/MyHobieMMORPGGame/services/engine/entities"
)

type Engine struct {
	entitiyList []entities.Entity
	ctx         context.Context
	cancel      context.CancelCauseFunc
}

func NewEngine(parent context.Context) *Engine {
	ctx, cancel := context.WithCancelCause(parent)

	return &Engine{
		entitiyList: make([]entities.Entity, 0),
		ctx:         ctx,
		cancel:      cancel,
	}
}

func (e *Engine) AddEntity(entity entities.Entity) {
	e.entitiyList = append(e.entitiyList, entity)
}

func (e *Engine) Start() {
	go e.update()
}

func (e *Engine) update() {
	lastUpdateTime := time.Now()

	for {
		select {
		case <-e.ctx.Done():
			// Engine düzgün şekilde durdurulur
			return
		default:
		}

		now := time.Now()
		deltaTime := now.Sub(lastUpdateTime).Seconds()
		lastUpdateTime = now

		for _, entity := range e.entitiyList {
			entity.PreUpdate(e.ctx, deltaTime)
		}

		for _, entity := range e.entitiyList {
			entity.Update(e.ctx, deltaTime)
		}

		for _, entity := range e.entitiyList {
			entity.LastUpdate(e.ctx, deltaTime)
			if !entity.IsAlive() {
				// cleanup / removal
			}
		}
	}
}
