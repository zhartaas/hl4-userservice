package handler

import (
	"github.com/go-chi/chi/v5"
	httpSwagger "github.com/swaggo/http-swagger/v2"
	"hl4-user_service/pkg/server/response"
	"hl4-user_service/service"
	"net/http"

	"hl4-user_service/docs"
	"hl4-user_service/pkg/server/router"
)

type Handler struct {
	HTTP    *chi.Mux
	service *service.Service
}

func New(service *service.Service) (h *Handler) {
	h = &Handler{
		HTTP:    router.New(),
		service: service,
	}
	docs.SwaggerInfo.BasePath = "/"
	h.HTTP.Get("/swagger/*", httpSwagger.WrapHandler)

	h.HTTP.Route("/", func(r chi.Router) {
		r.Get("/users", h.list)
	})

	return
}

// @Summary	list all users
// @Tags		user
// @Accept		json
// @Produce	json
// @Success	200 {object} response.Object
// @Failure	500	{object}	response.Object
// @Router		/user [get]
func (h *Handler) list(w http.ResponseWriter, r *http.Request) {
	res, err := h.service.GetAllUsers()
	if err != nil {
		response.InternalServerError(w, r, err)
		return
	}

	response.OK(w, r, res)
}
