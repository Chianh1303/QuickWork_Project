package controllers

import (
	"QuickWork/internal/services"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type SavedJobController struct {
	savedJobService services.SavedJobService
}

func NewSavedJobController(savedJobService services.SavedJobService) *SavedJobController {
	return &SavedJobController{savedJobService: savedJobService}
}

func (ctrl *SavedJobController) ToggleSaveJob(c *fiber.Ctx) error {
	userID, ok := getUserIDFromContext(c)
	if !ok || userID == 0 {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Unauthorized"})
	}

	jobIDParam := c.Params("jobId")
	jobID, err := strconv.ParseUint(jobIDParam, 10, 32)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid job ID"})
	}

	isSaved, err := ctrl.savedJobService.ToggleSaveJob(userID, uint(jobID))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	message := "❤️ Đã lưu công việc vào mục Yêu thích!"
	if !isSaved {
		message = "Đã bỏ lưu công việc khỏi mục Yêu thích."
	}

	return c.JSON(fiber.Map{
		"message":  message,
		"is_saved": isSaved,
		"job_id":   jobID,
	})
}

func (ctrl *SavedJobController) GetSavedJobs(c *fiber.Ctx) error {
	userID, ok := getUserIDFromContext(c)
	if !ok || userID == 0 {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Unauthorized"})
	}

	jobs, err := ctrl.savedJobService.GetSavedJobs(userID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	ids, _ := ctrl.savedJobService.GetSavedJobIDs(userID)

	return c.JSON(fiber.Map{
		"jobs":      jobs,
		"saved_ids": ids,
	})
}
