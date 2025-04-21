package service

// import (
// 	"fmt"
// 	"net/http/httptest"
// 	"testing"

// 	"github.com/clubhub/sales/internal/pkg/entity/model"

// 	entity "github.com/clubhub/sales/internal/pkg/entity/schema"
// 	mockRepository "github.com/clubhub/sales/internal/pkg/ports/mocks"

// 	"github.com/gin-gonic/gin"
// 	"github.com/stretchr/testify/assert"
// 	"github.com/stretchr/testify/mock"
// )

// func TestNewService(t *testing.T) {
// 	mockRepo := mockRepository.NewDBRepository(t)
// 	service := NewService(mockRepo)
// 	assert.NotNil(t, service, "El servicio no debe ser nil")
// }

// func TestCreateUser(t *testing.T) {
// 	t.Run("success case", func(t *testing.T) {
// 		mockRepo := mockRepository.NewDBRepository(t)
// 		svc := &Repository{repo: mockRepo}

// 		mockRepo.On("CreateCandidate", mock.Anything, mock.Anything).Return(int64(123), nil)

// 		c, _ := gin.CreateTestContext(httptest.NewRecorder())
// 		req := &model.CreateCandidate{
// 			Name:            "John Doe",
// 			Email:           "john@example.com",
// 			Gender:          "Male",
// 			SalaryExpected:  50000.00,
// 			Phone:           "+1234567890",
// 			ExperienceYears: 5,
// 		}

// 		expectedResp := &entity.Response{
// 			Data: int64(123),
// 			Result: entity.Result{
// 				Details: []entity.Detail{
// 					{InternalCode: "200", Message: "OK", Detail: "Registro Creado"},
// 				},
// 				Source: "Create Candidate",
// 			},
// 		}

// 		response, err := svc.CreateCandidate(c, req)
// 		assert.NoError(t, err)
// 		assert.Equal(t, expectedResp, response)
// 	})

// 	t.Run("error case", func(t *testing.T) {
// 		mockRepo := mockRepository.NewDBRepository(t)
// 		svc := &Repository{repo: mockRepo}

// 		expectedError := fmt.Errorf("database error")
// 		mockRepo.On("CreateCandidate", mock.Anything, mock.Anything).Return(int64(0), expectedError)

// 		c, _ := gin.CreateTestContext(httptest.NewRecorder())
// 		req := &model.CreateCandidate{
// 			Name:            "John Doe",
// 			Email:           "john@example.com",
// 			Gender:          "Male",
// 			SalaryExpected:  50000.00,
// 			Phone:           "+1234567890",
// 			ExperienceYears: 5,
// 		}

// 		response, err := svc.CreateCandidate(c, req)
// 		assert.Error(t, err)
// 		assert.Equal(t, expectedError, err)
// 		assert.Nil(t, response)
// 	})
// }

// func TestSelectCandidate(t *testing.T) {
// 	t.Run("success case", func(t *testing.T) {
// 		mockRepo := mockRepository.NewDBRepository(t)
// 		svc := &Repository{repo: mockRepo}

// 		expectedCandidate := &entity.CandidatesGetResponse{
// 			Name:            "John Doe",
// 			Email:           "john@example.com",
// 			Gender:          "Male",
// 			SalaryExpected:  50000.00,
// 			Phone:           "+1234567890",
// 			ExperienceYears: 5,
// 		}

// 		mockRepo.On("SelectCandidate", mock.Anything, mock.Anything).Return(expectedCandidate, nil)

// 		c, _ := gin.CreateTestContext(httptest.NewRecorder())
// 		req := &model.GetCandidate{Id: "123"}

// 		expectedResp := &entity.Response{
// 			Data: expectedCandidate,
// 			Result: entity.Result{
// 				Details: []entity.Detail{
// 					{InternalCode: "200", Message: "OK", Detail: "Registro Seleccionado"},
// 				},
// 				Source: "Select Candidate",
// 			},
// 		}

// 		response, err := svc.SelectCandidate(c, req)
// 		assert.NoError(t, err)
// 		assert.Equal(t, expectedResp, response)
// 	})

// 	t.Run("error case", func(t *testing.T) {
// 		mockRepo := mockRepository.NewDBRepository(t)
// 		svc := &Repository{repo: mockRepo}

// 		expectedError := fmt.Errorf("database error")
// 		mockRepo.On("SelectCandidate", mock.Anything, mock.Anything).Return(nil, expectedError)

// 		c, _ := gin.CreateTestContext(httptest.NewRecorder())
// 		req := &model.GetCandidate{Id: "123"}

// 		response, err := svc.SelectCandidate(c, req)
// 		assert.Error(t, err)
// 		assert.Equal(t, expectedError, err)
// 		assert.Nil(t, response)
// 	})
// }

// func TestUpdateCandidate(t *testing.T) {
// 	t.Run("success case", func(t *testing.T) {
// 		mockRepo := mockRepository.NewDBRepository(t)
// 		svc := &Repository{repo: mockRepo}

// 		expectedUpdate := &entity.CandidatesUpdateResponse{
// 			Name:            "John Doe Updated",
// 			Email:           "john.updated@example.com",
// 			Gender:          "Male",
// 			SalaryExpected:  60000.00,
// 			Phone:           "+1234567890",
// 			ExperienceYears: 6,
// 		}

// 		mockRepo.On("UpdateCandidate", mock.Anything, mock.Anything).Return(expectedUpdate, nil)

// 		c, _ := gin.CreateTestContext(httptest.NewRecorder())
// 		req := &model.UpdateCandidate{
// 			Id:              "123",
// 			Name:            "John Doe Updated",
// 			Email:           "john.updated@example.com",
// 			Gender:          "Male",
// 			SalaryExpected:  60000.00,
// 			Phone:           "+1234567890",
// 			ExperienceYears: 6,
// 		}

// 		expectedResp := &entity.Response{
// 			Data: expectedUpdate,
// 			Result: entity.Result{
// 				Details: []entity.Detail{
// 					{InternalCode: "200", Message: "OK", Detail: "Registro Actualizado"},
// 				},
// 				Source: "Update Candidate",
// 			},
// 		}

// 		response, err := svc.UpdateCandidate(c, req)
// 		assert.NoError(t, err)
// 		assert.Equal(t, expectedResp, response)
// 	})

// 	t.Run("error case", func(t *testing.T) {
// 		mockRepo := mockRepository.NewDBRepository(t)
// 		svc := &Repository{repo: mockRepo}

// 		expectedError := fmt.Errorf("database error")
// 		mockRepo.On("UpdateCandidate", mock.Anything, mock.Anything).Return(nil, expectedError)

// 		c, _ := gin.CreateTestContext(httptest.NewRecorder())
// 		req := &model.UpdateCandidate{Id: "123"}

// 		response, err := svc.UpdateCandidate(c, req)
// 		assert.Error(t, err)
// 		assert.Equal(t, expectedError, err)
// 		assert.Nil(t, response)
// 	})
// }

// func TestDeleteCandidate(t *testing.T) {
// 	t.Run("success case", func(t *testing.T) {
// 		mockRepo := mockRepository.NewDBRepository(t)
// 		svc := &Repository{repo: mockRepo}

// 		mockRepo.On("DeleteCandidate", mock.Anything, mock.Anything).Return(nil)

// 		c, _ := gin.CreateTestContext(httptest.NewRecorder())
// 		req := &model.DeleteCandidate{Id: "123"}

// 		expectedResp := &entity.Response{
// 			Result: entity.Result{
// 				Details: []entity.Detail{
// 					{InternalCode: "200", Message: "OK", Detail: "Registro Eliminado"},
// 				},
// 				Source: "Delete Candidate",
// 			},
// 		}

// 		response, err := svc.DeleteCandidate(c, req)
// 		assert.NoError(t, err)
// 		assert.Equal(t, expectedResp, response)
// 	})

// 	t.Run("error case", func(t *testing.T) {
// 		mockRepo := mockRepository.NewDBRepository(t)
// 		svc := &Repository{repo: mockRepo}

// 		expectedError := fmt.Errorf("database error")
// 		mockRepo.On("DeleteCandidate", mock.Anything, mock.Anything).Return(expectedError)

// 		c, _ := gin.CreateTestContext(httptest.NewRecorder())
// 		req := &model.DeleteCandidate{Id: "123"}

// 		response, err := svc.DeleteCandidate(c, req)
// 		assert.Error(t, err)
// 		assert.Equal(t, expectedError, err)
// 		assert.Nil(t, response)
// 	})
// }
