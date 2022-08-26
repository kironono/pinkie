package handler

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
	"github.com/go-playground/validator"
	"github.com/kironono/pinkie/model"
	"github.com/kironono/pinkie/registry"
	"github.com/kironono/pinkie/usecase"
)

type JobHandler interface {
	List(http.ResponseWriter, *http.Request)
	Show(http.ResponseWriter, *http.Request)
	Create(http.ResponseWriter, *http.Request)
	Update(http.ResponseWriter, *http.Request)
	Delete(http.ResponseWriter, *http.Request)
}

type job struct {
	uc       usecase.Job
	validate *validator.Validate
}

func NewJob(repo registry.Repository) JobHandler {
	uc := usecase.NewJob(repo.NewJob())
	validate := validator.New()
	return &job{
		uc:       uc,
		validate: validate,
	}
}

func (j *job) Show(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	id, _ := strconv.Atoi(chi.URLParam(r, "id"))

	job, err := j.uc.Show(ctx, model.JobID(id))
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
	RespondJSON(ctx, w, job, http.StatusOK)
}

func (j *job) List(w http.ResponseWriter, r *http.Request) {
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

	jobs, err := j.uc.List(ctx, model.PageNum(page), model.PerPageNum(per), model.Order(order))
	if err != nil {
		RespondJSON(ctx, w, &ErrResponse{
			Message: err.Error(),
		}, http.StatusInternalServerError)
		return
	}
	RespondJSON(ctx, w, jobs, http.StatusOK)
}

func (j *job) Create(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var b struct {
		Name string `json:"name" validate:"required"`
		Slug string `json:"slug" validate:"required"`
	}

	if err := json.NewDecoder(r.Body).Decode(&b); err != nil {
		RespondJSON(ctx, w, &ErrResponse{
			Message: err.Error(),
		}, http.StatusBadRequest)
		return
	}

	if err := j.validate.Struct(b); err != nil {
		RespondJSON(ctx, w, &ErrResponse{
			Message: err.Error(),
		}, http.StatusBadRequest)
		return
	}

	job, err := j.uc.Create(ctx, b.Name, b.Slug)
	if err != nil {
		RespondJSON(ctx, w, &ErrResponse{
			Message: err.Error(),
		}, http.StatusInternalServerError)
		return
	}
	RespondJSON(ctx, w, job, http.StatusCreated)
}

func (j *job) Update(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	id, _ := strconv.Atoi(chi.URLParam(r, "id"))

	var b struct {
		Name string `json:"name" validate:"required"`
		Slug string `json:"slug" validate:"required"`
	}

	if err := json.NewDecoder(r.Body).Decode(&b); err != nil {
		RespondJSON(ctx, w, &ErrResponse{
			Message: err.Error(),
		}, http.StatusBadRequest)
		return
	}

	if err := j.validate.Struct(b); err != nil {
		RespondJSON(ctx, w, &ErrResponse{
			Message: err.Error(),
		}, http.StatusBadRequest)
		return
	}

	job, err := j.uc.Update(ctx, model.JobID(id), b.Name, b.Slug)
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
	RespondJSON(ctx, w, job, http.StatusOK)
}

func (j *job) Delete(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	id, _ := strconv.Atoi(chi.URLParam(r, "id"))

	if err := j.uc.Delete(ctx, model.JobID(id)); err != nil {
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
	RespondJSON(ctx, w, nil, http.StatusOK)
}
