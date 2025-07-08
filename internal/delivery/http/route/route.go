package route

import (
	"encoding/json"
	"os"

	"github.com/januaruwanda/go-project-layout.git/internal/delivery/http"

	"github.com/gofiber/contrib/fiberzerolog"
	jwtware "github.com/gofiber/contrib/jwt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/swagger"
	"github.com/rs/zerolog"
	"github.com/spf13/viper"
)

type RouteConfig struct {
	App                   *fiber.App
	HealthCheckController *http.HealthCheckController
}

func NewRouteConfig() *RouteConfig {
	logger := zerolog.New(os.Stderr).With().Timestamp().Logger()

	app := fiber.New(fiber.Config{
		Prefork:     false,
		JSONEncoder: json.Marshal,
		JSONDecoder: json.Unmarshal,
		BodyLimit:   100 * 1024 * 1024,
	})

	app.Use(fiberzerolog.New(fiberzerolog.Config{
		Logger: &logger,
	}))

	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
	}))

	app.Get("/swagger/*", swagger.HandlerDefault)

	healthCheckController := http.NewHealthCheckController()

	routeConfig := RouteConfig{
		App:                   app,
		HealthCheckController: healthCheckController,
	}

	routeConfig.SetupNonAuthGroup()
	routeConfig.SetupAuthJWT()
	routeConfig.SetupAuthGroup()
	return &routeConfig
}

func (rc *RouteConfig) SetupNonAuthGroup() {
	rc.SetupHealthCheck()
}

func (rc *RouteConfig) SetupAuthGroup() {
	// You can write route with auth here
}

func (rc *RouteConfig) SetupAuthJWT() {
	rc.App.Use(jwtware.New(jwtware.Config{
		SigningKey: jwtware.SigningKey{Key: []byte(viper.GetString("encryption.jwt_private_key"))},
	}))
}

func (rc *RouteConfig) SetupHealthCheck() {
	rc.App.Get("/api/healthcheck", rc.HealthCheckController.HealthCheck)
}

func (rc *RouteConfig) Listen(address string) {
	rc.App.Listen(address)
}
