package handler

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/oganes5796/simple-blog/models"
)

func (h *Handler) CreateArticle(c *fiber.Ctx) error {
	var req struct {
		Title   string `json:"title"`
		Content string `json:"content"`
	}
	if err := c.BodyParser(&req); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "invalid request"})
	}

	userId := c.Locals("userId").(int)

	article, err := h.services.Article.CreateArticle(userId, req.Title, req.Content)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(article)
}

func (h *Handler) GetArticle(c *fiber.Ctx) error {
	articleId, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "invalid article id"})
	}

	article, err := h.services.Article.GetArticleByID(articleId)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "post not found"})
	}

	return c.JSON(article)
}

func (h *Handler) GetArticles(c *fiber.Ctx) error {
	article, err := h.services.Article.GetAllArticle()
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(article)
}

func (h *Handler) UpdateArticle(c *fiber.Ctx) error {
	var req struct {
		Title   string `json:"title"`
		Content string `json:"content"`
	}
	if err := c.BodyParser(&req); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "invalid request"})
	}

	articleId, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "invalid article id"})
	}

	userId := c.Locals("userId").(int)

	article := models.Article{
		ID:       articleId,
		AuthorID: userId,
		Title:    req.Title,
		Content:  req.Content,
	}

	if err := h.services.Article.UpdateArticle(article); err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "article not found or access denied"})
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{"message": "post update"})
}

func (h *Handler) DeleteArticle(c *fiber.Ctx) error {
	articleId, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "invalid article id"})
	}

	userId := c.Locals("userId").(int)

	if err := h.services.Article.DeleteArticle(articleId, userId); err != nil {
		return c.Status(http.StatusNotFound).JSON(fiber.Map{"error": "article not found or access denied"})
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{"message": "post deleted"})
}
