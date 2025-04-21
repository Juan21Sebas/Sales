package model

// type CreateCandidate struct {
// 	Name            string  `json:"name" binding:"required"`
// 	Email           string  `json:"email" binding:"required,email"`
// 	Gender          string  `json:"gender" binding:"required,oneof=Male Female Other"`
// 	SalaryExpected  float64 `json:"salary_expected" binding:"required"`
// 	Phone           string  `json:"phone" binding:"required"`
// 	ExperienceYears int     `json:"experience_years" binding:"required"`
// }

// type DeleteCandidate struct {
// 	Id string `json:"id"`
// }

type GetCandidate struct {
	Id string `json:"id"`
}

// type UpdateCandidate struct {
// 	Id              string  `json:"id"`
// 	Name            string  `json:"name"`
// 	Email           string  `json:"email"`
// 	Gender          string  `json:"gender"`
// 	SalaryExpected  float64 `json:"salary_expected"`
// 	Phone           string  `json:"phone"`
// 	ExperienceYears int     `json:"experience_years"`
// }
