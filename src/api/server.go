package api

import (
	"blogos/src/api/auto"
	"blogos/src/api/router"
	"blogos/src/config"
	"fmt"
	"log"
	"net/http"
)

func Run() {
	config.Load()

	auto.Load()

	fmt.Printf("Listening [::]:%d\n", config.PORT)

	listen(config.PORT)
}

func listen(port int) {
	r := router.New()

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), r))
}
