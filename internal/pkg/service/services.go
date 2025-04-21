package service

import (
	"net/http"
	"strconv"

	"github.com/clubhub/sales/internal/pkg/ports"

	entity "github.com/clubhub/sales/internal/pkg/entity/schema"

	"github.com/clubhub/sales/internal/pkg/entity/model"

	"github.com/gin-gonic/gin"
)

type Repository struct {
	repo ports.DBRepository
}

func NewService(repo ports.DBRepository) *Repository {
	return &Repository{
		repo: repo,
	}
}

// func (r *Repository) CreateCandidate(ctx *gin.Context, request *model.CreateCandidate) (*entity.Response, error) {

// 	resp, err := r.repo.CreateCandidate(ctx, request)
// 	if err != nil {
// 		return nil, err
// 	}

// 	return &entity.Response{
// 		Data: resp,
// 		Result: entity.Result{
// 			Details: []entity.Detail{
// 				{
// 					InternalCode: strconv.Itoa(http.StatusOK),
// 					Message:      http.StatusText(http.StatusOK),
// 					Detail:       "Registro Creado",
// 				},
// 			},
// 			Source: "Create Candidate",
// 		},
// 	}, nil

// }

func (r *Repository) SelectCandidate(ctx *gin.Context, request *model.GetCandidate) (*entity.Response, error) {

	resp, err := r.repo.SelectCandidate(ctx, request)
	if err != nil {
		return nil, err
	}

	return &entity.Response{
		Data: resp,
		Result: entity.Result{
			Details: []entity.Detail{
				{
					InternalCode: strconv.Itoa(http.StatusOK),
					Message:      http.StatusText(http.StatusOK),
					Detail:       "Registro Seleccionado",
				},
			},
			Source: "Select Candidate",
		},
	}, nil

}

// func (r *Repository) UpdateCandidate(ctx *gin.Context, request *model.UpdateCandidate) (*entity.Response, error) {

// 	resp, err := r.repo.UpdateCandidate(ctx, request)
// 	if err != nil {
// 		return nil, err
// 	}

// 	return &entity.Response{
// 		Data: resp,
// 		Result: entity.Result{
// 			Details: []entity.Detail{
// 				{
// 					InternalCode: strconv.Itoa(http.StatusOK),
// 					Message:      http.StatusText(http.StatusOK),
// 					Detail:       "Registro Actualizado",
// 				},
// 			},
// 			Source: "Update Candidate",
// 		},
// 	}, nil

// }

// func (r *Repository) DeleteCandidate(ctx *gin.Context, request *model.DeleteCandidate) (*entity.Response, error) {

// 	err := r.repo.DeleteCandidate(ctx, request)
// 	if err != nil {
// 		return nil, err
// 	}

// 	return &entity.Response{
// 		Result: entity.Result{
// 			Details: []entity.Detail{
// 				{
// 					InternalCode: strconv.Itoa(http.StatusOK),
// 					Message:      http.StatusText(http.StatusOK),
// 					Detail:       "Registro Eliminado",
// 				},
// 			},
// 			Source: "Delete Candidate",
// 		},
// 	}, nil

// }
