package main

import (
	"context"
	"fmt"
	"job-portal-api/config"
	"job-portal-api/internal/auth"
	"job-portal-api/internal/cache"
	"job-portal-api/internal/database"
	"job-portal-api/internal/handlers"
	"job-portal-api/internal/repository"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/rs/zerolog/log"
)

func main() {
	err := startApp() //calling the start app
	if err != nil {
		log.Panic().Err(err).Send()
	}
	log.Info().Msg("hello this is job portal app")
}

func startApp() error {
	cfg := config.GetConfig()
	log.Info().Msg("main : Started : Initializing authentication support")
	//message to the developer
	privatePEM := cfg.KeyConfig.Private_Key
	// privatePEM, err := os.ReadFile("private.pem")
	// //reading the file and returning the byte format
	// // if err != nil {
	// // 	log.Error().Err(err).Msg("reading auth private key")
	// // 	return err
	// // }
	privateKey, err := jwt.ParseRSAPrivateKeyFromPEM([]byte(privatePEM)) //converting the byte format and saving the address of privatekey
	if err != nil {
		log.Error().Err(err).Msg("parsing auth private key")
		return err
	}
	publicPEM := cfg.KeyConfig.Public_Key
	// publicPEM, err := os.ReadFile("pubkey.pem") //reading the file and returning the byte format
	// if err != nil {
	// 	log.Error().Err(err).Msg("reading auth public key")
	// 	return err
	// }

	publicKey, err := jwt.ParseRSAPublicKeyFromPEM([]byte(publicPEM))
	if err != nil {
		log.Error().Err(err).Msg("parsing auth public key")
		return err
	}

	a, err := auth.NewAuth(privateKey, publicKey)
	if err != nil {
		log.Error().Err(err).Msg("constructing authentication")
		return err
	}
	//connection of DB
	log.Info().Msg("main : Started : Initializing db support")
	db, err := database.Open()
	if err != nil {
		log.Error().Err(err).Msg("connecting to db")
		return err
	}
	pg, err := db.DB()
	if err != nil {
		log.Error().Err(err).Msg("Getting database connection")
		return err
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	err = pg.PingContext(ctx) //verifies the database connection is there are not
	if err != nil {
		log.Error().Err(err).Msg("Database is not connected")
		return err
	}
	// redis database connection
	rdb := database.RedisClient()
	log.Info().Msg("main : Started : Initializing redis connection")

	redisLayer := cache.NewRDBLayer(rdb)
	//initialize conn layer support
	ms, err := repository.NewRepo(db)
	if err != nil {
		log.Error().Err(err).Msg("database not connected to repository layer")
		return err
	}
	// svc, err := services.NewService(ms, redisLayer)
	// if err != nil {
	// 	return err
	// }

	// err = AutoMigrate()
	// if err != nil {
	// 	return err
	// }

	api := http.Server{
		Addr:         fmt.Sprintf("%s:%s", cfg.AppConfig.AppHost, cfg.AppConfig.Port),
		ReadTimeout:  time.Duration(cfg.AppConfig.ReadTimeout) * time.Second,
		WriteTimeout: time.Duration(cfg.AppConfig.WriteTimeout) * time.Second,
		IdleTimeout:  time.Duration(cfg.AppConfig.IdleTimeout) * time.Second,
		Handler:      handlers.API(a, ms, redisLayer),
	}
	// channel to store any errors while setting up the service
	serverErrors := make(chan error, 1)
	go func() {
		log.Info().Str("port", api.Addr).Msg("main: API listening")
		serverErrors <- api.ListenAndServe()
	}()
	shutdown := make(chan os.Signal, 1)
	//shutdown is just an empty channel
	signal.Notify(shutdown, os.Interrupt)
	//notify gives the value to the shutdown channel if interrupt occur(interrupt occurs when we click ctrl+C)
	select {
	case err := <-serverErrors:
		log.Error().Err(err).Msg("server error")
		return err
	case sig := <-shutdown:
		log.Info().Msgf("main: Start shutdown %s", sig)
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		//ctx is the input of context
		defer cancel()
		//after timeout cancel function wil work
		err := api.Shutdown(ctx)
		//api.Shutdown is graceful shutdown
		//shutdown is taking context
		if err != nil {
			err = api.Close() // forcing shutdown
			log.Error().Err(err).Msg("could not stop server gracefully")
			return err
		}
	}
	return nil
}
