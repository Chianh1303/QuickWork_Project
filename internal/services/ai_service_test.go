package services_test

import (
	"QuickWork/internal/dto"
	"QuickWork/internal/engine"
	"QuickWork/internal/models"
	"QuickWork/internal/parsers"
	"QuickWork/internal/repositories"
	"QuickWork/internal/services"
	"errors"
	"testing"

	"gorm.io/gorm"
)

type mockAIRepository struct {
	repositories.AIRepository
	students map[uint]*models.Student
	jobs     map[uint]*models.Job
}

func (m *mockAIRepository) GetStudentByUserID(userID uint) (*models.Student, error) {
	if s, found := m.students[userID]; found {
		return s, nil
	}
	return nil, gorm.ErrRecordNotFound
}

func (m *mockAIRepository) GetJobByID(jobID uint) (*models.Job, error) {
	if j, found := m.jobs[jobID]; found {
		return j, nil
	}
	return nil, gorm.ErrRecordNotFound
}

type mockPDFParser struct {
	parsers.PDFParser
}

func (m *mockPDFParser) GetCVFileBytesAndInfo(cvURL string) ([]byte, string, int64, bool) {
	return nil, "", 0, false
}

func (m *mockPDFParser) ParsePDFTextContent(pdfBytes []byte) parsers.ParsedCVData {
	return parsers.ParsedCVData{}
}

func TestMatchJob_StudentNotFound(t *testing.T) {
	repo := &mockAIRepository{
		students: map[uint]*models.Student{},
		jobs:     map[uint]*models.Job{},
	}
	svc := services.NewAIService(repo, nil, nil, &mockPDFParser{}, engine.NewMatchingEngine())

	_, err := svc.MatchJob(999, dto.MatchJobRequest{JobID: 1})
	if !errors.Is(err, services.ErrStudentNotFound) {
		t.Fatalf("expected ErrStudentNotFound, got %v", err)
	}
}

func TestMatchJob_JobNotFound(t *testing.T) {
	repo := &mockAIRepository{
		students: map[uint]*models.Student{
			1: {ID: 1, UserID: 1, FullName: "Test Student", Skills: `["go", "git"]`},
		},
		jobs: map[uint]*models.Job{},
	}
	svc := services.NewAIService(repo, nil, nil, &mockPDFParser{}, engine.NewMatchingEngine())

	_, err := svc.MatchJob(1, dto.MatchJobRequest{JobID: 999})
	if !errors.Is(err, services.ErrJobNotFound) {
		t.Fatalf("expected ErrJobNotFound, got %v", err)
	}
}

func TestMatchJob_FullMatch(t *testing.T) {
	repo := &mockAIRepository{
		students: map[uint]*models.Student{
			1: {ID: 1, UserID: 1, FullName: "Test Student", Skills: `["golang", "git"]`},
		},
		jobs: map[uint]*models.Job{
			10: {ID: 10, Title: "Golang Developer", Description: "Requires Go and Git.", Category: "IT", Salary: 10000000},
		},
	}
	svc := services.NewAIService(repo, nil, nil, &mockPDFParser{}, engine.NewMatchingEngine())

	res, err := svc.MatchJob(1, dto.MatchJobRequest{JobID: 10})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if res.MatchScore != 100 {
		t.Fatalf("expected MatchScore = 100, got %d", res.MatchScore)
	}
	if len(res.MissingSkills) != 0 {
		t.Fatalf("expected 0 MissingSkills, got %v", res.MissingSkills)
	}
}

func TestMatchJob_PartialMatch(t *testing.T) {
	repo := &mockAIRepository{
		students: map[uint]*models.Student{
			1: {ID: 1, UserID: 1, FullName: "Test Student", Skills: "git"},
		},
		jobs: map[uint]*models.Job{
			10: {ID: 10, Title: "Golang Developer", Description: "Requires Go and Git.", Category: "IT", Salary: 10000000},
		},
	}
	svc := services.NewAIService(repo, nil, nil, &mockPDFParser{}, engine.NewMatchingEngine())

	res, err := svc.MatchJob(1, dto.MatchJobRequest{JobID: 10})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if res.MatchScore >= 100 || res.MatchScore <= 0 {
		t.Fatalf("expected partial MatchScore between 1 and 99, got %d", res.MatchScore)
	}
	if len(res.MissingSkills) == 0 {
		t.Fatalf("expected missing skills, got empty")
	}
}

func TestMatchJob_CommaSeparatedSkills(t *testing.T) {
	repo := &mockAIRepository{
		students: map[uint]*models.Student{
			1: {ID: 1, UserID: 1, FullName: "Test Student", Skills: "golang, git"},
		},
		jobs: map[uint]*models.Job{
			10: {ID: 10, Title: "Golang Developer", Description: "Requires Go and Git.", Category: "IT", Salary: 10000000},
		},
	}
	svc := services.NewAIService(repo, nil, nil, &mockPDFParser{}, engine.NewMatchingEngine())

	res, err := svc.MatchJob(1, dto.MatchJobRequest{JobID: 10})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if res.MatchScore != 100 {
		t.Fatalf("expected MatchScore = 100 for comma separated skills, got %d", res.MatchScore)
	}
}
