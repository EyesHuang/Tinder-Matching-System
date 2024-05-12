package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"

	person "tinder"
	rHttp "tinder/http"
	repo "tinder/repo"
	rService "tinder/service"

	"github.com/go-playground/validator"
)

func main() {
	port := flag.Int("port", 8080, "listen port")
	flag.Parse()
	person.Validate = validator.New()

	if err := run(*port); err != nil {
		log.Fatal(err)
	}
}

func run(port int) error {
	repo := repo.NewMemoryRepo()
	matcher := rService.NewMatcherService(repo)
	srv := rHttp.NewServer(matcher)
	return http.ListenAndServe(fmt.Sprintf(":%d", port), srv)
}
