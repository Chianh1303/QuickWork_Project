package dto

// ============================================================
// EVALUATE CV DTO
// ============================================================

type EvaluateCVRequest struct {
	FullName string   `json:"full_name"`
	Phone    string   `json:"phone"`
	Gender   string   `json:"gender"`
	Skills   []string `json:"skills"`
	Bio      string   `json:"bio"`
	CvURL    string   `json:"cv_url"`
}

type EvaluateCVResponse struct {
	Score              int                `json:"score"`
	Recommendation     Recommendation     `json:"recommendation"`
	ATS                ATSReview          `json:"ats"`
	Skills             SkillsReview       `json:"skills"`
	Experience         ExperienceReview   `json:"experience"`
	Education          EducationReview    `json:"education"`
	Strengths          []string           `json:"strengths"`
	Weaknesses         []string           `json:"weaknesses"`
	MissingInformation []string           `json:"missing_information"`
	STAR               STARReview         `json:"star_analysis"`
	SuggestedSummary   string             `json:"suggested_summary"`
	ActionableTips     []string           `json:"actionable_tips"`
	EvaluationSource   string             `json:"evaluation_source"`
}

type Recommendation struct {
	Decision   string  `json:"decision"`
	Confidence float64 `json:"confidence"`
	Reason     string  `json:"reason"`
}

type ATSReview struct {
	Score     int      `json:"score"`
	Strengths []string `json:"strengths"`
	Issues    []string `json:"issues"`
}

type SkillsReview struct {
	Score     int      `json:"score"`
	Technical []string `json:"technical"`
	Soft      []string `json:"soft"`
}

type ExperienceReview struct {
	Score    int             `json:"score"`
	Projects []ProjectReview `json:"projects"`
}

type ProjectReview struct {
	Name         string   `json:"name"`
	Technologies []string `json:"technologies"`
	Description  string   `json:"description"`
	Impact       string   `json:"impact"`
}

type EducationReview struct {
	Score int             `json:"score"`
	Items []EducationItem `json:"items"`
}

type EducationItem struct {
	Institution    string `json:"institution"`
	Degree         string `json:"degree"`
	GraduationYear string `json:"graduation_year"`
}

type STARReview struct {
	Situation int `json:"situation"`
	Task      int `json:"task"`
	Action    int `json:"action"`
	Result    int `json:"result"`
}

// ============================================================
// MATCH JOB DTO
// ============================================================

type MatchJobRequest struct {
	JobID uint `json:"job_id"`
}

type MatchJobResponse struct {
	MatchScore     int      `json:"match_score"`
	MatchingReason []string `json:"matching_reasons"`
	MissingSkills  []string `json:"missing_skills"`
}

// ============================================================
// GENERATE JOB DTO
// ============================================================

type GenerateJobRequest struct {
	Title string `json:"title"`
}

type GenerateJobResponse struct {
	Description     string `json:"description"`
	Requirements    string `json:"requirements"`
	Benefits        string `json:"benefits"`
	SuggestedSalary string `json:"suggested_salary"`
}
