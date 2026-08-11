package services

import (
	"errors"
	"strings"
	"time"

	"QuickWork/internal/models"
	"QuickWork/internal/repositories"

	"gorm.io/gorm"
)

var (
	ErrAttendanceStudentNotFoundCheckIn  = errors.New("Tài khoản của bạn chưa cập nhật hồ sơ sinh viên!")
	ErrAttendanceStudentNotFoundCheckOut = errors.New("Không tìm thấy hồ sơ sinh viên")
	ErrActiveShiftExists                 = errors.New("Bạn đang có một ca làm việc chưa kết thúc. Vui lòng Check-out ca cũ trước!")
	ErrNotAppliedToJob                   = errors.New("Bạn chưa ứng tuyển vào công việc này!")
	ErrOfferNotAcceptedForShift          = errors.New("Công việc này chưa được trúng tuyển hoặc chưa kích hoạt ca làm!")
	ErrNoOpenShiftToCheckout             = errors.New("Không tìm thấy ca làm việc nào đang mở cần Check-out")
)

type AttendanceService interface {
	CheckIn(userID, jobID uint) (*models.Attendance, error)
	CheckOut(userID, jobID uint) (*models.Attendance, error)
}

type attendanceService struct {
	attendanceRepo repositories.AttendanceRepository
}

func NewAttendanceService(attendanceRepo repositories.AttendanceRepository) AttendanceService {
	return &attendanceService{attendanceRepo: attendanceRepo}
}

func (s *attendanceService) CheckIn(userID, jobID uint) (*models.Attendance, error) {
	student, err := s.attendanceRepo.GetStudentByUserID(userID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrAttendanceStudentNotFoundCheckIn
		}
		return nil, err
	}

	existingAttendance, err := s.attendanceRepo.GetActiveAttendance(student.ID)
	if err == nil && existingAttendance != nil {
		return nil, ErrActiveShiftExists
	}

	app, err := s.attendanceRepo.GetApplicationByStudentAndJob(student.ID, jobID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrNotAppliedToJob
		}
		return nil, err
	}

	statusNorm := strings.ToLower(strings.ReplaceAll(app.Status, "_", " "))
	if statusNorm != "offer accepted" {
		return nil, ErrOfferNotAcceptedForShift
	}

	now := time.Now()
	attendance := &models.Attendance{
		StudentID:   student.ID,
		JobID:       jobID,
		CheckInTime: &now,
		Status:      "working",
	}

	if err := s.attendanceRepo.CreateAttendance(attendance); err != nil {
		return nil, err
	}

	return attendance, nil
}

func (s *attendanceService) CheckOut(userID, jobID uint) (*models.Attendance, error) {
	student, err := s.attendanceRepo.GetStudentByUserID(userID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrAttendanceStudentNotFoundCheckOut
		}
		return nil, err
	}

	attendance, err := s.attendanceRepo.GetWorkingAttendance(student.ID, jobID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrNoOpenShiftToCheckout
		}
		return nil, err
	}

	now := time.Now()
	attendance.CheckOutTime = &now
	attendance.Status = "completed"

	if err := s.attendanceRepo.SaveAttendance(attendance); err != nil {
		return nil, err
	}

	return attendance, nil
}
