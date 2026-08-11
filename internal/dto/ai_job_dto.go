package dto

// ============================================================
// RECOMMENDED JOBS
// ============================================================

// RecommendedJobsResponse là response của API
// GET /api/ai/recommended-jobs
type RecommendedJobsResponse struct {
	Jobs []RecommendedJobItem `json:"jobs"`
}

// RecommendedJobItem đại diện cho một Job
// phù hợp với Student.
type RecommendedJobItem struct {
	JobID          uint     `json:"job_id"`
	JobTitle       string   `json:"job_title"`
	Company        string   `json:"company"`
	MatchScore     int      `json:"match_score"`
	MatchingSkills []string `json:"matching_skills"`
	MissingSkills  []string `json:"missing_skills"`
}

// ============================================================
// CV ADVISOR
// ============================================================

// CVAdviceResponse là response của API
// GET /api/ai/cv-advice/:jobId
type CVAdviceResponse struct {
	Job        CVAdviceJob    `json:"job"`
	MatchScore int            `json:"match_score"`
	CVAdvice   CVAdviceDetail `json:"cv_advice"`
}

// CVAdviceJob chứa thông tin Job mà Student đang xem.
type CVAdviceJob struct {
	JobID   uint   `json:"job_id"`
	Title   string `json:"title"`
	Company string `json:"company"`
}

// CVAdviceDetail chứa toàn bộ tư vấn CV.
type CVAdviceDetail struct {
	Summary             string                 `json:"summary"`
	Skills              CVSkillsAdvice         `json:"skills"`
	Projects            []CVProjectAdvice      `json:"projects"`
	KeywordsToHighlight []string               `json:"keywords_to_highlight"`
	Improvements        []string               `json:"improvements"`
}

// CVSkillsAdvice phân tích cách xử lý Skills
// khi ứng tuyển Job này.
type CVSkillsAdvice struct {
	Keep      []string `json:"keep"`
	Highlight []string `json:"highlight"`
	Missing   []string `json:"missing"`
}

// CVProjectAdvice đưa ra lời khuyên cho từng Project.
type CVProjectAdvice struct {
	Project string `json:"project"`
	Advice  string `json:"advice"`
	Impact  string `json:"impact"`
}