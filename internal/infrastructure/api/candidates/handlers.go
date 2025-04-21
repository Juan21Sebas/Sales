package candidates

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/clubhub/sales/internal/pkg/entity/model"
	"github.com/clubhub/sales/internal/pkg/ports"
)

type candidatesHandler struct {
	Service    ports.CommunicationServices
	Repository ports.DBRepository
}

func newHandler(service ports.CommunicationServices, repository ports.DBRepository) *candidatesHandler {
	return &candidatesHandler{
		Service:    service,
		Repository: repository,
	}
}

func (o *candidatesHandler) getCandidates() gin.HandlerFunc {
	return func(c *gin.Context) {
		var Candidate model.GetCandidate
		Candidate.Id = c.Param("id")
		if err := c.ShouldBindQuery(&Candidate); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid Request "})
			return
		}
		entityResponse, err := o.Service.SelectCandidate(c, &Candidate)
		if err != nil {
			c.JSON(http.StatusNotFound, err.Error())
			return
		}

		c.Set("entityResponse", *entityResponse)
		c.JSON(http.StatusOK, entityResponse)
	}
}

// func (o *candidatesHandler) postCanditates() gin.HandlerFunc {
// 	return func(c *gin.Context) {
// 		var Candidate model.CreateCandidate
// 		if err := c.BindJSON(&Candidate); err != nil {
// 			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid Request"})
// 			return
// 		}
// 		entityResponse, err := o.Service.CreateCandidate(c, &Candidate)
// 		if err != nil {
// 			c.JSON(http.StatusNotFound, err.Error())
// 			return
// 		}

// 		c.Set("entityResponse", *entityResponse)
// 		c.JSON(http.StatusOK, entityResponse)
// 	}
// }

// func (o *candidatesHandler) putCandidates() gin.HandlerFunc {
// 	return func(c *gin.Context) {
// 		var Candidates model.UpdateCandidate
// 		Candidates.Id = c.Param("id")
// 		if err := c.ShouldBindQuery(&Candidates); err != nil {
// 			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid Request "})
// 			return
// 		}
// 		if err := c.BindJSON(&Candidates); err != nil {
// 			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid Request "})
// 			return
// 		}
// 		entityResponse, err := o.Service.UpdateCandidate(c, &Candidates)
// 		if err != nil {
// 			c.JSON(http.StatusNotFound, err.Error())
// 			return
// 		}

// 		c.Set("entityResponse", *entityResponse)
// 		c.JSON(http.StatusOK, entityResponse)
// 	}
// }

// func (o *candidatesHandler) deleteCandidates() gin.HandlerFunc {
// 	return func(c *gin.Context) {
// 		var Candidates model.DeleteCandidate
// 		Candidates.Id = c.Param("id")
// 		if err := c.ShouldBindQuery(&Candidates); err != nil {
// 			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid Request "})
// 			return
// 		}
// 		entityResponse, err := o.Service.DeleteCandidate(c, &Candidates)
// 		if err != nil {
// 			c.JSON(http.StatusNotFound, err.Error())
// 			return
// 		}

// 		c.Set("entityResponse", *entityResponse)
// 		c.JSON(http.StatusOK, entityResponse)
// 	}
// }
