package handlers

import (
	"bitcoin_checker_api/config"
	"bitcoin_checker_api/internal/models"
	"bitcoin_checker_api/internal/repositories"
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
)

type Handler struct {
	repository repositories.Repository
	cfg        *config.Config
}

func NewHandler(cfg *config.Config, repository repositories.Repository) *Handler {
	return &Handler{
		cfg:        cfg,
		repository: repository,
	}
}

func (that *Handler) Rate(c *gin.Context) {
	fmt.Println("In rate")
	converter := models.NewConverter()
	requestURL := fmt.Sprintf("%s%s", that.cfg.Converter.Endpoint, converter.GetQueryParams())

	fmt.Println(requestURL)
	res, err := http.Get(requestURL)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, "Invalid status value")
		return
	}
	body, _ := io.ReadAll(res.Body)
	res.Body.Close()
	c.IndentedJSON(http.StatusOK, fmt.Sprintf("%s", body))
}

func (that *Handler) Subscription(c *gin.Context) {

	c.IndentedJSON(http.StatusOK, that.repository)
}

func (that *Handler) SendMail(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, that.repository)
}
