package repository

// import (
// 	"database/sql"
// 	"errors"
// 	"net/http/httptest"
// 	"testing"

// 	"github.com/clubhub/sales/internal/pkg/entity/model"
// 	"github.com/clubhub/sales/internal/pkg/entity/schema"

// 	"github.com/DATA-DOG/go-sqlmock"
// 	"github.com/gin-gonic/gin"
// 	"github.com/stretchr/testify/assert"
// )

// type MockResult struct {
// 	lastID    int64
// 	rowsAff   int64
// 	lastIdErr error
// }

// func (m MockResult) LastInsertId() (int64, error) {
// 	return m.lastID, m.lastIdErr
// }

// func (m MockResult) RowsAffected() (int64, error) {
// 	return m.rowsAff, nil
// }

// type MockResult2 struct {
// 	rowsAff int64
// 	rowsErr error
// }

// func (m MockResult2) LastInsertId() (int64, error) {
// 	return 0, nil
// }

// func (m MockResult2) RowsAffected() (int64, error) {
// 	return m.rowsAff, m.rowsErr
// }

// func TestCreateCandidate(t *testing.T) {
// 	db, mock, err := sqlmock.New()
// 	if err != nil {
// 		t.Fatalf("Error creating mock db: %v", err)
// 	}
// 	defer db.Close()

// 	repo := &BDRepository{
// 		db: db,
// 	}

// 	t.Run("success case", func(t *testing.T) {
// 		request := &model.CreateCandidate{
// 			Name:            "John Doe",
// 			Email:           "john@example.com",
// 			Gender:          "Male",
// 			SalaryExpected:  50000.00,
// 			Phone:           "+1234567890",
// 			ExperienceYears: 5,
// 		}

// 		mock.ExpectExec("INSERT INTO candidates").
// 			WithArgs(
// 				request.Name,
// 				request.Email,
// 				request.Gender,
// 				request.SalaryExpected,
// 				request.Phone,
// 				"PENDING",
// 				request.ExperienceYears,
// 				sqlmock.AnyArg(),
// 				sqlmock.AnyArg(),
// 			).
// 			WillReturnResult(sqlmock.NewResult(123, 1))

// 		c, _ := gin.CreateTestContext(httptest.NewRecorder())
// 		id, err := repo.CreateCandidate(c, request)

// 		assert.NoError(t, err)
// 		assert.Equal(t, int64(123), id)
// 	})

// 	t.Run("error case - db error", func(t *testing.T) {
// 		request := &model.CreateCandidate{
// 			Name: "John Doe",
// 		}

// 		mock.ExpectExec("INSERT INTO candidates").
// 			WillReturnError(errors.New("db error"))

// 		c, _ := gin.CreateTestContext(httptest.NewRecorder())
// 		_, err := repo.CreateCandidate(c, request)

// 		assert.Error(t, err)
// 		assert.Contains(t, err.Error(), "error creating candidate")
// 	})

// 	t.Run("error case - LastInsertId error", func(t *testing.T) {
// 		request := &model.CreateCandidate{
// 			Name:            "John Doe",
// 			Email:           "john@example.com",
// 			Gender:          "Male",
// 			SalaryExpected:  50000.00,
// 			Phone:           "+1234567890",
// 			ExperienceYears: 5,
// 		}

// 		mockResult := MockResult{
// 			lastID:    0,
// 			lastIdErr: errors.New("last insert id error"),
// 		}

// 		mock.ExpectExec("INSERT INTO candidates").
// 			WithArgs(
// 				request.Name,
// 				request.Email,
// 				request.Gender,
// 				request.SalaryExpected,
// 				request.Phone,
// 				"PENDING",
// 				request.ExperienceYears,
// 				sqlmock.AnyArg(),
// 				sqlmock.AnyArg(),
// 			).
// 			WillReturnResult(mockResult)

// 		c, _ := gin.CreateTestContext(httptest.NewRecorder())
// 		_, err := repo.CreateCandidate(c, request)

// 		assert.Error(t, err)
// 		assert.Contains(t, err.Error(), "error getting inserted id")
// 	})
// }

// func TestUpdateCandidate(t *testing.T) {
// 	db, mock, err := sqlmock.New()
// 	if err != nil {
// 		t.Fatalf("Error creating mock db: %v", err)
// 	}
// 	defer db.Close()

// 	repo := &BDRepository{
// 		db: db,
// 	}

// 	t.Run("success case", func(t *testing.T) {
// 		request := &model.UpdateCandidate{
// 			Id:              "123",
// 			Name:            "John Updated",
// 			Email:           "john.updated@example.com",
// 			Gender:          "Male",
// 			SalaryExpected:  60000.00,
// 			Phone:           "+1234567890",
// 			ExperienceYears: 6,
// 		}

// 		mock.ExpectExec("UPDATE candidates").
// 			WithArgs(
// 				request.Name,
// 				request.Email,
// 				request.Gender,
// 				request.SalaryExpected,
// 				request.Phone,
// 				request.ExperienceYears,
// 				sqlmock.AnyArg(),
// 				request.Id,
// 			).
// 			WillReturnResult(sqlmock.NewResult(0, 1))

// 		c, _ := gin.CreateTestContext(httptest.NewRecorder())
// 		result, err := repo.UpdateCandidate(c, request)

// 		assert.NoError(t, err)
// 		assert.Equal(t, request.Name, result.Name)
// 		assert.Equal(t, request.Email, result.Email)
// 	})

// 	t.Run("error case - update error", func(t *testing.T) {
// 		request := &model.UpdateCandidate{
// 			Id:              "123",
// 			Name:            "John Updated",
// 			Email:           "john.updated@example.com",
// 			Gender:          "Male",
// 			SalaryExpected:  60000.00,
// 			Phone:           "+1234567890",
// 			ExperienceYears: 6,
// 		}

// 		mock.ExpectExec("UPDATE candidates").
// 			WithArgs(
// 				request.Name,
// 				request.Email,
// 				request.Gender,
// 				request.SalaryExpected,
// 				request.Phone,
// 				request.ExperienceYears,
// 				sqlmock.AnyArg(),
// 				request.Id,
// 			).
// 			WillReturnError(errors.New("db error"))

// 		c, _ := gin.CreateTestContext(httptest.NewRecorder())
// 		_, err := repo.UpdateCandidate(c, request)

// 		assert.Error(t, err)
// 		assert.Contains(t, err.Error(), "error updating candidate")
// 	})

// 	t.Run("error case - rows affected error", func(t *testing.T) {
// 		request := &model.UpdateCandidate{
// 			Id:              "123",
// 			Name:            "John Updated",
// 			Email:           "john.updated@example.com",
// 			Gender:          "Male",
// 			SalaryExpected:  60000.00,
// 			Phone:           "+1234567890",
// 			ExperienceYears: 6,
// 		}

// 		mockResult := MockResult2{
// 			rowsAff: 0,
// 			rowsErr: errors.New("rows affected error"),
// 		}

// 		mock.ExpectExec("UPDATE candidates").
// 			WithArgs(
// 				request.Name,
// 				request.Email,
// 				request.Gender,
// 				request.SalaryExpected,
// 				request.Phone,
// 				request.ExperienceYears,
// 				sqlmock.AnyArg(),
// 				request.Id,
// 			).
// 			WillReturnResult(mockResult)

// 		c, _ := gin.CreateTestContext(httptest.NewRecorder())
// 		_, err := repo.UpdateCandidate(c, request)

// 		assert.Error(t, err)
// 		assert.Contains(t, err.Error(), "error getting rows affected")
// 	})

// 	t.Run("not found case", func(t *testing.T) {
// 		request := &model.UpdateCandidate{Id: "123"}

// 		mock.ExpectExec("UPDATE candidates").
// 			WillReturnResult(sqlmock.NewResult(0, 0))

// 		c, _ := gin.CreateTestContext(httptest.NewRecorder())
// 		_, err := repo.UpdateCandidate(c, request)

// 		assert.Error(t, err)
// 		assert.Contains(t, err.Error(), "candidate not found")
// 	})
// }
// func TestDeleteCandidate(t *testing.T) {
// 	db, mock, err := sqlmock.New()
// 	if err != nil {
// 		t.Fatalf("Error creating mock db: %v", err)
// 	}
// 	defer db.Close()

// 	repo := &BDRepository{
// 		db: db,
// 	}

// 	t.Run("success case", func(t *testing.T) {
// 		mock.ExpectExec("DELETE FROM candidates").
// 			WithArgs("123").
// 			WillReturnResult(sqlmock.NewResult(0, 1))

// 		c, _ := gin.CreateTestContext(httptest.NewRecorder())
// 		err := repo.DeleteCandidate(c, &model.DeleteCandidate{Id: "123"})

// 		assert.NoError(t, err)
// 	})

// 	t.Run("error case - delete error", func(t *testing.T) {
// 		mock.ExpectExec("DELETE FROM candidates").
// 			WithArgs("123").
// 			WillReturnError(errors.New("db error"))

// 		c, _ := gin.CreateTestContext(httptest.NewRecorder())
// 		err := repo.DeleteCandidate(c, &model.DeleteCandidate{Id: "123"})

// 		assert.Error(t, err)
// 		assert.Contains(t, err.Error(), "error deleting candidate")
// 	})

// 	t.Run("error case - rows affected error", func(t *testing.T) {
// 		mockResult := MockResult2{
// 			rowsAff: 0,
// 			rowsErr: errors.New("rows affected error"),
// 		}

// 		mock.ExpectExec("DELETE FROM candidates").
// 			WithArgs("123").
// 			WillReturnResult(mockResult)

// 		c, _ := gin.CreateTestContext(httptest.NewRecorder())
// 		err := repo.DeleteCandidate(c, &model.DeleteCandidate{Id: "123"})

// 		assert.Error(t, err)
// 		assert.Contains(t, err.Error(), "error getting rows affected")
// 	})

// 	t.Run("not found case", func(t *testing.T) {
// 		mock.ExpectExec("DELETE FROM candidates").
// 			WithArgs("123").
// 			WillReturnResult(sqlmock.NewResult(0, 0))

// 		c, _ := gin.CreateTestContext(httptest.NewRecorder())
// 		err := repo.DeleteCandidate(c, &model.DeleteCandidate{Id: "123"})

// 		assert.Error(t, err)
// 		assert.Contains(t, err.Error(), "candidate not found")
// 	})
// }

// func TestSelectCandidate(t *testing.T) {
// 	db, mock, err := sqlmock.New()
// 	if err != nil {
// 		t.Fatalf("Error creating mock db: %v", err)
// 	}
// 	defer db.Close()

// 	repo := &BDRepository{
// 		db: db,
// 	}

// 	t.Run("success case", func(t *testing.T) {
// 		expectedCandidate := &schema.CandidatesGetResponse{
// 			Name:            "John Doe",
// 			Email:           "john@example.com",
// 			Gender:          "Male",
// 			SalaryExpected:  50000.00,
// 			Phone:           "+1234567890",
// 			ExperienceYears: 5,
// 		}

// 		rows := sqlmock.NewRows([]string{"name", "email", "gender", "salary_expected", "phone", "experience_years"}).
// 			AddRow(expectedCandidate.Name, expectedCandidate.Email, expectedCandidate.Gender,
// 				expectedCandidate.SalaryExpected, expectedCandidate.Phone, expectedCandidate.ExperienceYears)

// 		mock.ExpectQuery("SELECT (.+) FROM candidates").
// 			WithArgs("123").
// 			WillReturnRows(rows)

// 		c, _ := gin.CreateTestContext(httptest.NewRecorder())
// 		result, err := repo.SelectCandidate(c, &model.GetCandidate{Id: "123"})

// 		assert.NoError(t, err)
// 		assert.Equal(t, expectedCandidate, result)
// 	})

// 	t.Run("not found case", func(t *testing.T) {
// 		mock.ExpectQuery("SELECT (.+) FROM candidates").
// 			WithArgs("123").
// 			WillReturnError(sql.ErrNoRows)

// 		c, _ := gin.CreateTestContext(httptest.NewRecorder())
// 		result, err := repo.SelectCandidate(c, &model.GetCandidate{Id: "123"})

// 		assert.Error(t, err)
// 		assert.Nil(t, result)
// 		assert.Contains(t, err.Error(), "candidate not found with id: 123")
// 	})

// 	t.Run("query error case", func(t *testing.T) {
// 		mock.ExpectQuery("SELECT (.+) FROM candidates").
// 			WithArgs("123").
// 			WillReturnError(errors.New("db error"))

// 		c, _ := gin.CreateTestContext(httptest.NewRecorder())
// 		result, err := repo.SelectCandidate(c, &model.GetCandidate{Id: "123"})

// 		assert.Error(t, err)
// 		assert.Nil(t, result)
// 		assert.Contains(t, err.Error(), "error querying candidate")
// 	})

// 	t.Run("scan error case", func(t *testing.T) {
// 		// Devolver columnas incorrectas para provocar error de scan
// 		rows := sqlmock.NewRows([]string{"name", "email"}).
// 			AddRow("John Doe", "john@example.com")

// 		mock.ExpectQuery("SELECT (.+) FROM candidates").
// 			WithArgs("123").
// 			WillReturnRows(rows)

// 		c, _ := gin.CreateTestContext(httptest.NewRecorder())
// 		result, err := repo.SelectCandidate(c, &model.GetCandidate{Id: "123"})

// 		assert.Error(t, err)
// 		assert.Nil(t, result)
// 		assert.Contains(t, err.Error(), "error querying candidate")
// 	})
// }
