package handler

import (
	"encoding/json"
	"net/http"
	"shiori-server/service"

	"github.com/gin-gonic/gin"
)

type TopHandler struct {
	service service.TopService
}

func NewTopHandler(s service.TopService) *TopHandler {
	return &TopHandler{service: s}
}

// GetTopPage トップページ情報の取得
// @Summary トップページ取得
// @Description トップ画像と挨拶文を取得。
// @Tags top
// @Accept json
// @Produce json
// @Success 200 {object} dto.TopPageResource
// @Failure 500 {object} map[string]string
// @Router /top [get]
func (h *TopHandler) GetTopPage(c *gin.Context) {
	response, err := h.service.GetTopPage()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"topページ情報取得サービスが呼び出せませんでした。": err.Error()})
		return
	}

	c.Writer.Header().Set("Content-Type", "application/json")
	c.Writer.WriteHeader(http.StatusOK)

	encoder := json.NewEncoder(c.Writer)
	encoder.SetEscapeHTML(false)

	if err := encoder.Encode(response); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"JSONデータのエンコードに失敗しました。": err.Error()})

	}
}
