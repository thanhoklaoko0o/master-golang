package api

import (
	"fmt"
	"log"
	"net/http"

	"github.com/alfredomendozacap/jwt-auth/auto"
	"github.com/alfredomendozacap/jwt-auth/config"

	"github.com/alfredomendozacap/jwt-auth/api/router"
)

func init() {
	config.Load()
	auto.Load()
}

// Run message
func Run() {
	fmt.Printf("\n\tListening [::]:%d", config.PORT)
	listen(config.PORT)
}

func listen(port int) {
	r := router.New()
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), r))
}
