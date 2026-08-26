package models

import "time"

type CVEvaluation struct {
	ID                       uint      `gorm:"primaryKey" json:"id"`
	UserID                   uint      `gorm:"not null;index" json:"user_id"`
	CvURL                    string    `gorm:"type:varchar(255)" json:"cv_url"`
	ATSScore                 int       `json:"ats_score"`
	SkillsScore              int       `json:"skills_score"`
	ExperienceScore          int       `json:"experience_score"`
	EducationScore           int       `json:"education_score"`
	STARSituation            int       `json:"star_situation"`
	STARTask                 int       `json:"star_task"`
	STARAction               int       `json:"star_action"`
	STARResult               int       `json:"star_result"`
	Strengths                string    `gorm:"type:text" json:"strengths"`  // JSON array text
	Weaknesses               string    `gorm:"type:text" json:"weaknesses"` // JSON array text
	RecommendationDecision   string    `gorm:"type:varchar(50)" json:"recommendation_decision"`
	RecommendationConfidence float64   `json:"recommendation_confidence"`
	EvaluationSource         string    `gorm:"type:varchar(20)" json:"evaluation_source"`
	CreatedAt                time.Time `json:"created_at"`
}
