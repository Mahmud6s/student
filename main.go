package main

import (
	"WEB-NEW-WDB/handler"
	"WEB-NEW-WDB/storage/postgres"
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/alexedwards/scs/v2"
	"github.com/go-playground/form"
	_ "github.com/lib/pq"
	"github.com/pressly/goose"
	"github.com/spf13/viper"
)

var sessionManager *scs.SessionManager

// .....................................................................

func main() {
	// .........Viper-for-Config..................................
	config := viper.NewWithOptions(
		viper.EnvKeyReplacer(
			strings.NewReplacer(".", "_"),
		),
	)
	config.SetConfigFile("env/config")
	config.SetConfigType("ini")
	config.AutomaticEnv()
	if err := config.ReadInConfig(); err != nil {
		log.Fatalf("error loading configuration: %v", err)
	}
	// ...................................................................

	decoder := form.NewDecoder() //..........Useing-Playground-for-Decode..

	//..........DATABASECONNECTION............................
	PostGstorage, err := postgres.PostGresStorageCON(config)
	if err != nil {
		log.Fatalln(err)
	}
	//.........DB CONNECTION END............................................

	if err := goose.SetDialect("postgres"); err != nil {
		log.Fatalln(err)
	}

	if err := goose.Up(PostGstorage.DB.DB, "migrations"); err != nil {
		log.Fatalln(err)
	}
	//.............SASSIONS......................................
	lt := config.GetDuration("session.lifetime")
	it := config.GetDuration("session.idletime")
	sessionManager = scs.New()
	sessionManager.Lifetime = lt * time.Hour
	sessionManager.IdleTimeout = it * time.Minute
	sessionManager.Cookie.Name = "web-session"
	sessionManager.Cookie.HttpOnly = true
	sessionManager.Cookie.Secure = true

	//...................................................................

	chi := handler.NewHandler(decoder, sessionManager, PostGstorage)
	p := config.GetInt("server.port")

	if err := http.ListenAndServe(fmt.Sprintf(":%d", p), chi); err != nil {
		log.Fatalf("%#v", err)
	}
}
