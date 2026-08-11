package dto

type CreateJobInput struct {
	Title       string  `json:"title"`
	Description string  `json:"description"`
	Location    string  `json:"location"`
	Salary      float64 `json:"salary"`
	Slots       int     `json:"slots"`
	WorkingDate string  `json:"working_date"`
}

type JobFilterQuery struct {
	Search    string
	Location  string
	Category  string
	JobType   string
	MaxSalary string
}
