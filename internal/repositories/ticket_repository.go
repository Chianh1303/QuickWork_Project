package repositories

import (
	"time"

	"QuickWork/internal/models"

	"gorm.io/gorm"
)

type TicketRepository interface {
	CreateTicket(ticket *models.Ticket) error
	GetUserTickets(userID uint) ([]models.Ticket, error)
	GetTicketByApplicationAndReporter(appID uint, reporterID uint) (*models.Ticket, error)
	ReappealTicket(ticketID uint, userID uint, reason string, evidence string) (int64, error)
}

type ticketRepository struct {
	db *gorm.DB
}

func NewTicketRepository(db *gorm.DB) TicketRepository {
	return &ticketRepository{db: db}
}

func (r *ticketRepository) CreateTicket(ticket *models.Ticket) error {
	return r.db.Create(ticket).Error
}

func (r *ticketRepository) GetUserTickets(userID uint) ([]models.Ticket, error) {
	var tickets []models.Ticket
	err := r.db.Preload("Application").Preload("Application.Job").
		Where("reporter_id = ? OR target_id = ?", userID, userID).
		Order("created_at desc").Find(&tickets).Error
	return tickets, err
}

func (r *ticketRepository) GetTicketByApplicationAndReporter(appID uint, reporterID uint) (*models.Ticket, error) {
	var ticket models.Ticket
	err := r.db.Where("application_id = ? AND reporter_id = ?", appID, reporterID).First(&ticket).Error
	if err != nil {
		return nil, err
	}
	return &ticket, nil
}

func (r *ticketRepository) ReappealTicket(ticketID uint, userID uint, reason string, evidence string) (int64, error) {
	now := time.Now()
	result := r.db.Model(&models.Ticket{}).
		Where("id = ? AND (reporter_id = ? OR target_id = ?) AND (status = 'resolved' OR status = 'rejected')", ticketID, userID, userID).
		Updates(map[string]interface{}{
			"status":            "pending",
			"is_reappealed":     true,
			"reappeal_reason":   reason,
			"reappeal_evidence": evidence,
			"reappealed_at":     now,
		})
	return result.RowsAffected, result.Error
}
