package main

import (
	"errors"
	"log"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/swagger"
	"rccsilva.com/template-go/domain"
)

type Handlers struct {
	app *domain.App
}

func newHandler(app *domain.App) *Handlers {
	return &Handlers{app: app}
}

type BadRequestResponse struct {
	Message string `json:"message"`
}

type HealthResponse struct {
	Status string `json:"status"`
}

// @Description Get system health.
// @Summary Get system health
// @Tags System
// @Accept json
// @Produce json
// @Router /api/health [get]
// @Success 200 {object} HealthResponse
func (Handlers) GetHealth(ctx *fiber.Ctx) error {
	return ctx.JSON(HealthResponse{Status: "ok"})
}

// @Description Creates an user
// @Summary Creates an user
// @Tags User
// @Accept json
// @Produce json
// @Router /api/v1/user [post]
// @Success 201 {object} domain.CreateUserResponse
// @Param request body domain.CreateUserRequest true "create user request"
func (h Handlers) CreateUser(ctx *fiber.Ctx) error {
	p := new(domain.CreateUserRequest)
	ctx.BodyParser(&p)
	data, err := h.app.CreateUser(p)
	return parseResult(ctx, fiber.StatusCreated, data, err)
}

// @Description Gets an user
// @Summary Gets an user
// @Tags User
// @Accept json
// @Produce json
// @Router /api/v1/user/{id} [get]
// @Success 200 {object} domain.CreateUserResponse
// @Param        id   path      int  true  "user id"
func (h Handlers) GetUser(ctx *fiber.Ctx) error {
	id, err := strconv.Atoi(ctx.Params("id"))
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(&BadRequestResponse{Message: "id must be an integer"})
	}
	data, err := h.app.GetUser(id)
	return parseResult(ctx, fiber.StatusOK, data, err)
}

func parseResult(ctx *fiber.Ctx, defaultStatus int, data any, err error) error {
	if err == nil {
		return ctx.Status(defaultStatus).JSON(data)
	}
	var re *domain.ResultError

	if errors.As(err, &re) {
		return ctx.Status(re.StatusCode).JSON(re)
	}

	log.Printf("%+v", err)

	return ctx.SendStatus(500)
}

func (h *Handlers) configRoutes(router fiber.Router) {
	router.Use(recover.New(recover.Config{StackTraceHandler: func(c *fiber.Ctx, e any) {
		log.Printf("%+v", e)
		c.Next()
	}}))
	router.Get("/swagger/*", swagger.HandlerDefault)

	api := router.Group("api")
	{
		api.Get("/health", h.GetHealth)

		v1 := api.Group("v1")
		{
			v1.Post("user", h.CreateUser)
			v1.Get("user/:id", h.GetUser)
		}
	}
}
