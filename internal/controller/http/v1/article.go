package v1

import (
	"encoding/json"
	"github.com/gofiber/fiber/v2"
	"strconv"

	"fiber_blog/internal/entity"
	"fiber_blog/internal/usecase"
	"fiber_blog/pkg/logger"
)

var articleUseCase usecase.ArticleRepo
var articleLog logger.Interface

func newArticleRoutes(h fiber.Router, a usecase.ArticleRepo, l logger.Interface) {
	articleUseCase = a
	articleLog = l

	h.Get("/articles/:id", retrieveArticle)
	h.Put("/articles/:id", changeArticle)
	h.Post("/articles", createArticle)
	h.Delete("/articles/:id", deleteArticle)
}

type retrieveArticleResponse struct {
	entity.Article
}

func retrieveArticle(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return err
	}
	article, err := articleUseCase.Retrieve(entity.ArticleID(id))
	if err != nil {
		return err
	}
	response, err := json.Marshal(&retrieveArticleResponse{Article: article})
	if err != nil {
		return err
	}
	c.Send(response)
	return nil
}

type changeArticleRequest struct {
	entity.Article
}

func changeArticle(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return err
	}

	changeArticleRequest := changeArticleRequest{}
	err = json.Unmarshal(c.Body(), &changeArticleRequest)
	if err != nil {
		return err
	}

	changeArticleRequest.Article.ID = entity.ArticleID(id) // TODO зачем, если в теле можно передавать?

	err = articleUseCase.Update(changeArticleRequest.Article)
	if err != nil {
		return err
	}
	return nil
}

type createArticleRequest struct {
	entity.Article
}

type createArticleResponse struct {
	entity.ArticleID
}

func createArticle(c *fiber.Ctx) error {
	createArticleRequest := createArticleRequest{}
	err := json.Unmarshal(c.Body(), &createArticleRequest)
	if err != nil {
		return err
	}

	articleID, err := articleUseCase.Add(createArticleRequest.Article)
	if err != nil {
		return err
	}
	response, err := json.Marshal(&createArticleResponse{ArticleID: articleID})
	if err != nil {
		return err
	}
	c.Send(response)
	return nil
}

func deleteArticle(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return err
	}

	err = articleUseCase.Delete(entity.ArticleID(id))
	if err != nil {
		return err
	}
	return nil
}
