package repositories

import (
	"QuickWork/internal/models"
	"gorm.io/gorm"
)

type AttendanceRepository interface {
	GetStudentByUserID(userID uint) (*models.Student, error)
	GetActiveAttendance(studentID uint) (*models.Attendance, error)
	GetApplicationByStudentAndJob(studentID, jobID uint) (*models.Application, error)
	CreateAttendance(attendance *models.Attendance) error
	GetWorkingAttendance(studentID, jobID uint) (*models.Attendance, error)
	SaveAttendance(attendance *models.Attendance) error
}

type attendanceRepository struct {
	db *gorm.DB
}

func NewAttendanceRepository(db *gorm.DB) AttendanceRepository {
	return &attendanceRepository{db: db}
}

func (r *attendanceRepository) GetStudentByUserID(userID uint) (*models.Student, error) {
	var student models.Student
	if err := r.db.Where("user_id = ?", userID).First(&student).Error; err != nil {
		return nil, err
	}
	return &student, nil
}

func (r *attendanceRepository) GetActiveAttendance(studentID uint) (*models.Attendance, error) {
	var attendance models.Attendance
	if err := r.db.Where("student_id = ? AND status = ?", studentID, "working").First(&attendance).Error; err != nil {
		return nil, err
	}
	return &attendance, nil
}

func (r *attendanceRepository) GetApplicationByStudentAndJob(studentID, jobID uint) (*models.Application, error) {
	var app models.Application
	if err := r.db.Where("student_id = ? AND job_id = ?", studentID, jobID).First(&app).Error; err != nil {
		return nil, err
	}
	return &app, nil
}

func (r *attendanceRepository) CreateAttendance(attendance *models.Attendance) error {
	return r.db.Create(attendance).Error
}

func (r *attendanceRepository) GetWorkingAttendance(studentID, jobID uint) (*models.Attendance, error) {
	var attendance models.Attendance
	if err := r.db.Where("student_id = ? AND job_id = ? AND status = ?", studentID, jobID, "working").First(&attendance).Error; err != nil {
		return nil, err
	}
	return &attendance, nil
}

func (r *attendanceRepository) SaveAttendance(attendance *models.Attendance) error {
	return r.db.Save(attendance).Error
}
