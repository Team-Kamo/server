package main

import (
	"OctaneServer/config"
	"OctaneServer/data"
	"OctaneServer/handlers"
	"net/http"
	"os"
	"time"

	"github.com/goccy/go-json"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/etag"
	"github.com/gofiber/fiber/v2/middleware/filesystem"
	"github.com/gofiber/fiber/v2/middleware/limiter"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/monitor"
	"github.com/gofiber/fiber/v2/middleware/pprof"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func init() {
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stdout})
	if config.CurrentConfig.Debug {
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
	} else {
		zerolog.SetGlobalLevel(zerolog.InfoLevel)
	}
	log.Info().Msg(config.Msg[config.CurrentConfig.Lang].Console.Starting)
}

func main() {
	app := fiber.New(fiber.Config{
		Prefork:     true,
		JSONEncoder: json.Marshal,
		JSONDecoder: json.Unmarshal,
	})

	if config.CurrentConfig.Debug {
		app.Use(pprof.New())
		app.Get("/metrics", monitor.New(monitor.Config{Title: "Metrics"}))
	}
	app.Use(limiter.New())
	app.Use(compress.New())
	app.Use(recover.New())
	app.Use(etag.New())
	app.Use(logger.New(logger.Config{Output: log.Logger, Format: config.CurrentConfig.LogFormat}))
	app.Use(config.CurrentConfig.Root+"uploads", filesystem.New(filesystem.Config{
		Root: http.Dir("./uploads"),
	}))
	app.Use(limiter.New(limiter.Config{
		Max:                    1,
		Expiration:             1 * time.Second,
		SkipSuccessfulRequests: true,
	}))
	app.Get(config.CurrentConfig.Root+handlers.Health, handlers.HealthGet)
	app.Post(config.CurrentConfig.Root+handlers.Room, handlers.RoomPost)
	app.Get(config.CurrentConfig.Root+handlers.Status, handlers.StatusGet)
	app.Put(config.CurrentConfig.Root+handlers.Status, handlers.StatusPut)
	app.Delete(config.CurrentConfig.Root+handlers.Status, handlers.StatusDelete)
	app.Get(config.CurrentConfig.Root+handlers.Content, handlers.ContentGet)
	app.Put(config.CurrentConfig.Root+handlers.Content, handlers.ContentPut)
	app.Delete(config.CurrentConfig.Root+handlers.Content, handlers.ContentDelete)
	app.Get(config.CurrentConfig.Root+handlers.RoomID, handlers.RoomIDGet)
	app.Post(config.CurrentConfig.Root+handlers.RoomID, handlers.RoomIdPost)
	app.Delete(config.CurrentConfig.Root+handlers.RoomID, handlers.RoomIDDelete)

	defer func() {
		data.Close()
	}()

	log.Fatal().Err(app.Listen(":" + config.CurrentConfig.Port)).Msg(config.Msg[config.CurrentConfig.Lang].Console.HandlerExit)
}
