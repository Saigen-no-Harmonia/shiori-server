package handler

import (
	"encoding/json"
	"net/http"
	"shiori-server/dto"
	"shiori-server/service"

	"github.com/gin-gonic/gin"
)

type FamiliesHandler struct {
	service service.FamiliesService
}

func NewFamiliesHandler(s service.FamiliesService) *FamiliesHandler {
	return &FamiliesHandler{service: s}
}

// GetFamilies 家族情報ページの情報を取得
// @Summary 家族情報取得
// @Description 主催者・参加者・猫プロフィール情報を取得
// @Tags families
// @Accept json
// @Produce json
// @Success 200 {object} dto.FamiliesPageResource
// @Failure 500 {object} map[string]string
// @Router /families [get]
func (h *FamiliesHandler) GetFamilies(c *gin.Context) {
	// familiesプロフィール情報を取得
	families, err := h.service.GetFamiliesPage()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Familiesページ情報が取得できませんでした。": err.Error()})
		return
	}

	// Resourceに詰め替え
	response := dto.NewFamiliesPageResponse(families)

	// S3画像アクセス用URLのエスケープ処理
	c.Writer.Header().Set("Content-Type", "application/json")
	c.Writer.WriteHeader(http.StatusOK)
	encoder := json.NewEncoder(c.Writer)
	encoder.SetEscapeHTML(false)
	if err := encoder.Encode(response); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
}
