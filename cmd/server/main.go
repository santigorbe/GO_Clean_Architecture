package main

import (
	"github.com/santigorbe/clean_arq/internal/delivery/http"
	"github.com/santigorbe/clean_arq/internal/repository"
	"github.com/santigorbe/clean_arq/internal/usecase"
	"github.com/santigorbe/clean_arq/pkg/server"
)

func main() {
	repo := repository.NewInMemoryUserRepo()
	useCase := usecase.NewUserUseCase(repo)
	r := server.NewRouter()
	http.NewUserHandler(r, useCase)

	server.StartServer(r)
}
