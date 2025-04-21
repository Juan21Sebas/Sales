package candidates

import (
	"database/sql"

	"github.com/gin-gonic/gin"

	repository "github.com/clubhub/sales/internal/infrastructure/repository"
	services "github.com/clubhub/sales/internal/pkg/service"
)

func RegisterRoutes(e *gin.Engine, db *sql.DB) {
	Repository := repository.NewBdRepository(db)
	Service := services.NewService(Repository)
	Handler := newHandler(Service, Repository)
	// Registra las rutas candidates
	//e.POST("/candidates/", Handler.postCanditates())
	e.GET("/responses/:id", Handler.getCandidates())
	//e.PUT("/candidates/:id", Handler.putCandidates())
	//e.DELETE("/candidates/:id", Handler.deleteCandidates())

}
