package handler

import (
	"errors"
	"log"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
	"github.com/kironono/pinkie/model"
	"github.com/kironono/pinkie/registry"
	"github.com/kironono/pinkie/usecase"
)

type UserHandler interface {
	List(http.ResponseWriter, *http.Request)
	Show(http.ResponseWriter, *http.Request)
}

type user struct {
	uc usecase.User
}

func NewUser(repo registry.Repository) UserHandler {
	return &user{
		uc: usecase.NewUser(repo.NewUser()),
	}
}

func (u *user) List(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	page, err := strconv.Atoi(r.URL.Query().Get("page"))
	if err != nil || page < 1 {
		page = 1
	}
	per, err := strconv.Atoi(r.URL.Query().Get("per"))
	if err != nil || per < 1 {
		per = DEFAULT_PER_PAGE_NUM
	}
	order := "created_at desc"

	users, err := u.uc.List(ctx, model.PageNum(page), model.PerPageNum(per), model.Order(order))
	if err != nil {
		RespondJSON(ctx, w, &ErrResponse{
			Message: err.Error(),
		}, http.StatusInternalServerError)
		return
	}
	RespondJSON(ctx, w, users, http.StatusOK)
}

func (u *user) Show(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	id, _ := strconv.Atoi(chi.URLParam(r, "id"))

	user, err := u.uc.Show(ctx, model.UserID(id))
	if err != nil {
		switch {
		case errors.Is(err, model.ErrRecordNotFound):
			RespondJSON(ctx, w, &ErrResponse{
				Message: "Not Found",
			}, http.StatusNotFound)
		default:
			log.Printf("%s\n", err)
			RespondJSON(ctx, w, &ErrResponse{
				Message: err.Error(),
			}, http.StatusInternalServerError)
		}
		return
	}
	RespondJSON(ctx, w, user, http.StatusOK)
}
