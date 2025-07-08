package main

import (
	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"

	"github.com/januaruwanda/go-project-layout.git/internal/delivery/http/route"
	"github.com/januaruwanda/go-project-layout.git/internal/utils"
	"github.com/januaruwanda/go-project-layout.git/pkg/database"
)

func init() {
	viper.SetConfigFile(`config.json`)
	err := viper.ReadInConfig()
	if err != nil {
		log.Panic().Msg("please set config.json file")
	}

	if viper.GetBool(`debug`) {
		log.Info().Msg("Service RUN on DEBUG mode")
	}

	utils.InitCrypto()
	database.NewConnection()

}

// @title BigVision Operation API
// @version 1.0
// @description BigVision Operation
// @termsOfService http://swagger.io/terms/
// @contact.name API Support
// @contact.email 975316@telkom.co.id
// @license.name MIT
// @license.url https://opensource.org/license/mit
// @host localhost:8080
// @BasePath /

// @securityDefinitions.apiKey BearerAuth
// @type apiKey
// @name Authorization
// @in header
// @description "JWT Bearer token required for authentication"

// @security BearerAuth
func main() {
	host := viper.GetString(`server.host`)
	port := viper.GetString(`server.port`)
	router := route.NewRouteConfig()
	router.Listen(host + ":" + port)
}
