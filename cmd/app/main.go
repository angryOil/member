package main

import (
	"github.com/gorilla/mux"
	"member/cmd/app/handler"
	"member/internal/controller"
	"member/internal/repository"
	"member/internal/repository/infla"
	"member/internal/service"
	"net/http"
)

func main() {
	r := mux.NewRouter()
	h := getHandler()

	r.PathPrefix("/members").Handler(h)
	http.ListenAndServe(":8084", r)
}

func getHandler() http.Handler {
	return handler.NewHandler(controller.NewMemberController(service.NewMemberService(repository.NewMemberRepository(infla.NewDB()))))
}
