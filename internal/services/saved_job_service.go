package services

import (
	"QuickWork/internal/models"
	"QuickWork/internal/repositories"
	"errors"

	"gorm.io/gorm"
)

type SavedJobService interface {
	ToggleSaveJob(userID uint, jobID uint) (bool, error)
	GetSavedJobs(userID uint) ([]models.Job, error)
	GetSavedJobIDs(userID uint) ([]uint, error)
}

type savedJobService struct {
	savedJobRepo repositories.SavedJobRepository
}

func NewSavedJobService(savedJobRepo repositories.SavedJobRepository) SavedJobService {
	return &savedJobService{
		savedJobRepo: savedJobRepo,
	}
}

func (s *savedJobService) ToggleSaveJob(userID uint, jobID uint) (bool, error) {
	student, err := s.savedJobRepo.GetStudentByUserID(userID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return false, ErrStudentNotFound
		}
		return false, err
	}

	isSaved, err := s.savedJobRepo.IsJobSaved(student.ID, jobID)
	if err != nil {
		return false, err
	}

	if isSaved {
		if err := s.savedJobRepo.UnsaveJob(student.ID, jobID); err != nil {
			return false, err
		}
		return false, nil
	} else {
		if err := s.savedJobRepo.SaveJob(student.ID, jobID); err != nil {
			return false, err
		}
		return true, nil
	}
}

func (s *savedJobService) GetSavedJobs(userID uint) ([]models.Job, error) {
	student, err := s.savedJobRepo.GetStudentByUserID(userID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return []models.Job{}, nil
		}
		return nil, err
	}

	savedList, err := s.savedJobRepo.GetSavedJobsByStudentID(student.ID)
	if err != nil {
		return nil, err
	}

	var jobs []models.Job
	for _, sj := range savedList {
		if sj.Job.ID > 0 {
			jobs = append(jobs, sj.Job)
		}
	}

	return jobs, nil
}

func (s *savedJobService) GetSavedJobIDs(userID uint) ([]uint, error) {
	student, err := s.savedJobRepo.GetStudentByUserID(userID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return []uint{}, nil
		}
		return nil, err
	}

	savedList, err := s.savedJobRepo.GetSavedJobsByStudentID(student.ID)
	if err != nil {
		return nil, err
	}

	var ids []uint
	for _, sj := range savedList {
		ids = append(ids, sj.JobID)
	}

	return ids, nil
}
