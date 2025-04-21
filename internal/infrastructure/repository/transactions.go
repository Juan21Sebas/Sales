package repository

import (
	"database/sql"
	"fmt"

	"github.com/clubhub/sales/internal/pkg/entity/model"
	"github.com/clubhub/sales/internal/pkg/entity/schema"

	"github.com/gin-gonic/gin"
)

// func (p *BDRepository) CreateCandidate(ctx *gin.Context, request *model.CreateCandidate) (int64, error) {
// 	p.mu.Lock()
// 	defer p.mu.Unlock()

// 	now := time.Now().UTC()

// 	query := `
//         INSERT INTO candidates (
//             name,
//             email,
//             gender,
//             salary_expected,
//             phone,
//             status,
//             experience_years,
//             created_at,
//             updated_at
//         ) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?)
//     `

// 	result, err := p.db.Exec(
// 		query,
// 		request.Name,
// 		request.Email,
// 		request.Gender,
// 		request.SalaryExpected,
// 		request.Phone,
// 		"PENDING",
// 		request.ExperienceYears,
// 		now,
// 		now,
// 	)
// 	if err != nil {
// 		return 0, fmt.Errorf("error creating candidate: %w", err)
// 	}

// 	id, err := result.LastInsertId()
// 	if err != nil {
// 		return 0, fmt.Errorf("error getting inserted id: %w", err)
// 	}

// 	return id, nil
// }

func (p *BDRepository) SelectCandidate(ctx *gin.Context, request *model.GetCandidate) (*schema.CandidatesGetResponse, error) {
	p.mu.Lock()
	defer p.mu.Unlock()

	query := `
        SELECT 
            name,
            email,
            gender,
            salary_expected,
            phone,
            experience_years
        FROM candidates
        WHERE id = ?
    `

	var candidate schema.CandidatesGetResponse
	err := p.db.QueryRow(query, request.Id).Scan(
		&candidate.Name,
		&candidate.Email,
		&candidate.Gender,
		&candidate.SalaryExpected,
		&candidate.Phone,
		&candidate.ExperienceYears,
	)

	if err == sql.ErrNoRows {
		return nil, fmt.Errorf("candidate not found with id: %s", request.Id)
	}
	if err != nil {
		return nil, fmt.Errorf("error querying candidate: %w", err)
	}

	return &candidate, nil
}

// func (p *BDRepository) UpdateCandidate(ctx *gin.Context, request *model.UpdateCandidate) (*schema.CandidatesUpdateResponse, error) {
// 	p.mu.Lock()
// 	defer p.mu.Unlock()

// 	now := time.Now().UTC()

// 	query := `
// 		UPDATE candidates
// 		SET
// 			name = ?,
// 			email = ?,
// 			gender = ?,
// 			salary_expected = ?,
// 			phone = ?,
// 			experience_years = ?,
// 			updated_at = ?
// 		WHERE id = ?
// 	`

// 	result, err := p.db.Exec(
// 		query,
// 		request.Name,
// 		request.Email,
// 		request.Gender,
// 		request.SalaryExpected,
// 		request.Phone,
// 		request.ExperienceYears,
// 		now,
// 		request.Id,
// 	)
// 	if err != nil {
// 		return nil, fmt.Errorf("error updating candidate: %w", err)
// 	}

// 	rowsAffected, err := result.RowsAffected()
// 	if err != nil {
// 		return nil, fmt.Errorf("error getting rows affected: %w", err)
// 	}

// 	if rowsAffected == 0 {
// 		return nil, fmt.Errorf("candidate not found with id: %s", request.Id)
// 	}

// 	updatedCandidate := &schema.CandidatesUpdateResponse{
// 		Name:            request.Name,
// 		Email:           request.Email,
// 		Gender:          request.Gender,
// 		SalaryExpected:  request.SalaryExpected,
// 		Phone:           request.Phone,
// 		ExperienceYears: request.ExperienceYears,
// 	}

// 	return updatedCandidate, nil
// }

// func (p *BDRepository) DeleteCandidate(ctx *gin.Context, request *model.DeleteCandidate) error {
// 	p.mu.Lock()
// 	defer p.mu.Unlock()

// 	query := `DELETE FROM candidates WHERE id = ?`

// 	result, err := p.db.Exec(query, request.Id)
// 	if err != nil {
// 		return fmt.Errorf("error deleting candidate: %w", err)
// 	}

// 	rowsAffected, err := result.RowsAffected()
// 	if err != nil {
// 		return fmt.Errorf("error getting rows affected: %w", err)
// 	}

// 	if rowsAffected == 0 {
// 		return fmt.Errorf("candidate not found with id: %s", request.Id)
// 	}

// 	return nil
// }
