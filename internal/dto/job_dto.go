package dto

type CreateJobInput struct {
	Title       string  `json:"title"`
	Description string  `json:"description"`
	Location    string  `json:"location"`
	Salary      float64 `json:"salary"`
	Slots       int     `json:"slots"`
	WorkingDate string  `json:"working_date"`
	Category    string  `json:"category"`
	JobType     string  `json:"job_type"`
}

type JobFilterQuery struct {
	Search    string
	Location  string
	Category  string
	JobType   string
	MaxSalary string
}
