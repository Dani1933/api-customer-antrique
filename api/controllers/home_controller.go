package controllers

import (
	"net/http"

	"github.com/Dani1933/api-customer-antrique/api/responses"
)

func (server *Server) Home(w http.ResponseWriter, r *http.Request) {
	responses.JSON(w, http.StatusOK, "Welcome To This Awesome API")

}