// Package gameapp
package gameapp

import (
	"fmt"
	"sync"

	gameconfig "github.com/Akif-jpg/MyHobieMMORPGGame/config"
	"github.com/gofiber/fiber/v3"
)

type RuntimeEnvEnum int

const (
	DEVELOPER RuntimeEnvEnum = iota
	PRODUCT
)

type GameApp struct {
	App        *fiber.App
	GameConfig *gameconfig.GameConfig
	initOnce   sync.Once // Init'in bir kez çalışmasını garantiler
}

func New() *GameApp {
	return &GameApp{}
}

func (g *GameApp) Init(ree RuntimeEnvEnum) {
	g.initOnce.Do(func() {
		switch ree {
		case DEVELOPER:
			g.GameConfig = (&gameconfig.GameConfig{}).GetConfig(gameconfig.DEVELOPER)
			fmt.Println("Developer game config initialized")
		case PRODUCT:
			g.GameConfig = (&gameconfig.GameConfig{}).GetConfig(gameconfig.PRODUCT)
			fmt.Println("Product game config initialized")
		default:
			g.GameConfig = (&gameconfig.GameConfig{}).GetConfig(gameconfig.DEVELOPER)
			fmt.Println("Default game config initialized")
		}

		g.App = fiber.New()
		g.App.Get("health", func(c fiber.Ctx) error {
			return c.SendString("OK")
		})

		fmt.Println("Port:", g.GameConfig.Port)
	})
}

func (g *GameApp) Start() error {
	if g.App == nil || g.GameConfig == nil {
		return fmt.Errorf("GameApp is not initialized, call Init() first")
	}
	return g.App.Listen(":" + g.GameConfig.Port)
}
