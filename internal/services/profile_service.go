package services

import (
	"errors"

	"QuickWork/internal/models"
	"QuickWork/internal/repositories"

	"gorm.io/gorm"
)

var (
	ErrStudentProfileNotFound  = errors.New("Không tìm thấy hồ sơ sinh viên")
	ErrBusinessProfileNotFound = errors.New("Không tìm thấy hồ sơ doanh nghiệp")
	ErrInvalidPDFFormat        = errors.New("Hồ sơ đính kèm bắt buộc phải là định dạng file PDF")
)

type ProfileService interface {
	GetStudentProfile(userID uint) (*models.Student, error)
	UpdateStudentProfile(userID uint, fullName, phone, gender, skills, avatarUrl, cvUrl string) (*models.Student, error)
	GetBusinessProfile(userID uint) (*models.Business, error)
	UpdateBusinessProfile(userID uint, companyName, phone, address, logoUrl string) (*models.Business, error)
}

type profileService struct {
	profileRepo repositories.ProfileRepository
}

func NewProfileService(profileRepo repositories.ProfileRepository) ProfileService {
	return &profileService{profileRepo: profileRepo}
}

func (s *profileService) GetStudentProfile(userID uint) (*models.Student, error) {
	student, err := s.profileRepo.GetStudentByUserID(userID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrStudentProfileNotFound
		}
		return nil, err
	}
	return student, nil
}

func (s *profileService) UpdateStudentProfile(userID uint, fullName, phone, gender, skills, avatarUrl, cvUrl string) (*models.Student, error) {
	student, err := s.profileRepo.GetStudentByUserID(userID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrStudentProfileNotFound
		}
		return nil, err
	}

	if fullName != "" {
		student.FullName = fullName
	}
	if phone != "" {
		student.Phone = phone
	}
	if gender != "" {
		student.Gender = gender
	}
	if skills != "" {
		student.Skills = skills
	}
	if avatarUrl != "" {
		student.AvatarUrl = avatarUrl
	}
	if cvUrl != "" {
		student.CvUrl = cvUrl
	}

	if err := s.profileRepo.SaveStudent(student); err != nil {
		return nil, err
	}
	return student, nil
}

func (s *profileService) GetBusinessProfile(userID uint) (*models.Business, error) {
	business, err := s.profileRepo.GetBusinessByUserID(userID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrBusinessProfileNotFound
		}
		return nil, err
	}
	return business, nil
}

func (s *profileService) UpdateBusinessProfile(userID uint, companyName, phone, address, logoUrl string) (*models.Business, error) {
	business, err := s.profileRepo.GetBusinessByUserID(userID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrBusinessProfileNotFound
		}
		return nil, err
	}

	if companyName != "" {
		business.CompanyName = companyName
	}
	if phone != "" {
		business.Phone = phone
	}
	if address != "" {
		business.Address = address
	}
	if logoUrl != "" {
		business.LogoUrl = logoUrl
	}

	if err := s.profileRepo.SaveBusiness(business); err != nil {
		return nil, err
	}
	return business, nil
}
