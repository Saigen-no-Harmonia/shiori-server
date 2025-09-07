package handler

import (
	"net/http"
	"shiori-server/service"

	"github.com/gin-gonic/gin"
)

type AccessHandler struct {
	service service.AccessService
}

func NewAccessHandler(s service.AccessService) *AccessHandler {
	return &AccessHandler{service: s}
}

// GetAccessPage アクセスページ情報の取得
// @Summary アクセスページ取得
// @Description アクセス情報を返却。
// @Tags access
// @Accept json
// @Produce json
// @Success 200 {object} dto.AccessPageResource
// @Failure 500 {object} map[string]string
// @Router /access [get]
func (h *AccessHandler) GetAccessPage(c *gin.Context) {
	response, err := h.service.GetAccessPage()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"accessページ情報取得に失敗しました。": err.Error()})
		return
	}

	c.JSON(http.StatusOK, response)
}
