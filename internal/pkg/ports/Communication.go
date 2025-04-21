package ports

import (
	"github.com/clubhub/sales/internal/pkg/entity/model"
	"github.com/clubhub/sales/internal/pkg/entity/schema"

	"github.com/gin-gonic/gin"
)

type CommunicationServices interface {
	//CreateCandidate(ctx *gin.Context, request *model.CreateCandidate) (*schema.Response, error)
	SelectCandidate(ctx *gin.Context, request *model.GetCandidate) (*schema.Response, error)
	// UpdateCandidate(ctx *gin.Context, request *model.UpdateCandidate) (*schema.Response, error)
	// DeleteCandidate(ctx *gin.Context, request *model.DeleteCandidate) (*schema.Response, error)
}

type DBRepository interface {
	//CreateCandidate(ctx *gin.Context, request *model.CreateCandidate) (int64, error)
	SelectCandidate(ctx *gin.Context, request *model.GetCandidate) (*schema.CandidatesGetResponse, error)
	// UpdateCandidate(ctx *gin.Context, request *model.UpdateCandidate) (*schema.CandidatesUpdateResponse, error)
	// DeleteCandidate(ctx *gin.Context, request *model.DeleteCandidate) error
}
