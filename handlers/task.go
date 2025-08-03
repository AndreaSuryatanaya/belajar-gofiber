package handlers

import (
	"time"

	"belajar-go-fiber/database"
	"belajar-go-fiber/models"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type CreateTaskRequest struct {
	Todo      string    `json:"todo" validate:"required"`
	StartDate time.Time `json:"start_date" validate:"required"`
	EndDate   time.Time `json:"end_date" validate:"required"`
}

type UpdateTaskRequest struct {
	Todo      string    `json:"todo"`
	StartDate time.Time `json:"start_date"`
	EndDate   time.Time `json:"end_date"`
}

// GetTasks retrieves all tasks (no user filtering)
func GetTasks(c *fiber.Ctx) error {
	var tasks []models.Task
	if err := database.DB.Find(&tasks).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to retrieve tasks",
		})
	}

	return c.JSON(fiber.Map{
		"tasks": tasks,
	})
}

// GetTask retrieves a specific task by ID (open to all users)
func GetTask(c *fiber.Ctx) error {
	taskID := c.Params("id")

	taskUUID, err := uuid.Parse(taskID)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid task ID format",
		})
	}

	var task models.Task
	if err := database.DB.First(&task, "id = ?", taskUUID).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Task not found",
		})
	}

	return c.JSON(task)
}

// CreateTask creates a new task for the authenticated user
func CreateTask(c *fiber.Ctx) error {
	user := c.Locals("user").(models.User)

	var req CreateTaskRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	// Validate end date is after start date
	if req.EndDate.Before(req.StartDate) {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "End date must be after start date",
		})
	}

	task := models.Task{
		UserID:    user.ID,
		Todo:      req.Todo,
		StartDate: req.StartDate,
		EndDate:   req.EndDate,
	}

	if err := database.DB.Create(&task).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to create task",
		})
	}

	return c.Status(fiber.StatusCreated).JSON(task)
}

// UpdateTask updates an existing task (open to all users)
func UpdateTask(c *fiber.Ctx) error {
	taskID := c.Params("id")

	taskUUID, err := uuid.Parse(taskID)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid task ID format",
		})
	}

	var task models.Task
	if err := database.DB.Where("id = ?", taskUUID).First(&task).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Task not found",
		})
	}

	var req UpdateTaskRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	// Update fields if provided
	if req.Todo != "" {
		task.Todo = req.Todo
	}
	if !req.StartDate.IsZero() {
		task.StartDate = req.StartDate
	}
	if !req.EndDate.IsZero() {
		task.EndDate = req.EndDate
	}

	// Validate end date is after start date
	if task.EndDate.Before(task.StartDate) {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "End date must be after start date",
		})
	}

	if err := database.DB.Save(&task).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to update task",
		})
	}

	return c.JSON(task)
}

// DeleteTask deletes a task (open to all users)
func DeleteTask(c *fiber.Ctx) error {
	taskID := c.Params("id")

	taskUUID, err := uuid.Parse(taskID)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid task ID format",
		})
	}

	var task models.Task
	if err := database.DB.Where("id = ?", taskUUID).First(&task).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Task not found",
		})
	}

	if err := database.DB.Delete(&task).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to delete task",
		})
	}

	return c.JSON(fiber.Map{
		"message": "Task deleted successfully",
	})
}
