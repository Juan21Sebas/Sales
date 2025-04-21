package schema

type Response struct {
	Data   any    `json:"data,omitempty" mask:"struct"`
	Result Result `json:"result"`
}

type Result struct {
	Details []Detail `json:"details"`
	Source  string   `json:"source"`
}

type Detail struct {
	InternalCode string `json:"internalCode"`
	Message      string `json:"message"`
	Detail       string `json:"detail"`
}

type CandidatesGetResponse struct {
	Name            string  `json:"name" binding:"required"`
	Email           string  `json:"email" binding:"required,email"`
	Gender          string  `json:"gender" binding:"required,oneof=Male Female Other"`
	SalaryExpected  float64 `json:"salary_expected" binding:"required"`
	Phone           string  `json:"phone" binding:"required"`
	ExperienceYears int     `json:"experience_years" binding:"required"`
}

type CandidatesUpdateResponse struct {
	Name            string  `json:"name" binding:"required"`
	Email           string  `json:"email" binding:"required,email"`
	Gender          string  `json:"gender" binding:"required,oneof=Male Female Other"`
	SalaryExpected  float64 `json:"salary_expected" binding:"required"`
	Phone           string  `json:"phone" binding:"required"`
	ExperienceYears int     `json:"experience_years" binding:"required"`
}
