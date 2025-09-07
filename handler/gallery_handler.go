package handler

import (
	"encoding/json"
	"net/http"
	"shiori-server/dto"
	"shiori-server/service"
	"strconv"

	"github.com/gin-gonic/gin"
)

type GalleryHandler struct {
	service service.GalleryService
}

func NewGalleryHandler(s service.GalleryService) *GalleryHandler {
	return &GalleryHandler{service: s}
}

// GetGalleryPhotos ギャラリー画像の一覧を取得
// @Summary ギャラリー画像取得
// @Description ギャラリー画像の一覧をlimitとoffsetで取得
// @Tags gallery
// @Accept json
// @Produce json
// @Param limit query int false "取得件数" default(20)
// @Param offset query int false "開始位置" default(0)
// @Success 200 {object} dto.GalleryPhotoPageResource
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /gallery [get]
func (h *GalleryHandler) GetGalleryPhotos(c *gin.Context) {
	limitStr := c.DefaultQuery("limit", "20")
	offsetStr := c.DefaultQuery("offset", "0")

	// パラメータをint型に変換
	limit, offset, err := parseParamToInt(limitStr, offsetStr)
	if err != nil {
		c.JSON(
			http.StatusInternalServerError,
			gin.H{"failed to parse limit or offset to integer.": err.Error()},
		)
		return
	}

	galleryPhotos, err := h.service.GetGalleryPhotos(limit, offset)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"ギャラリー画像の取得に失敗しました": err.Error()})
		return
	}

	response := dto.NewGalleryPhotoPageResource(galleryPhotos)

	// S3画像アクセス用URLのエスケープ処理
	c.Writer.Header().Set("Content-Type", "application/json")
	c.Writer.WriteHeader(http.StatusOK)

	encoder := json.NewEncoder(c.Writer)
	encoder.SetEscapeHTML(false)

	if err := encoder.Encode(response); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
}

func parseParamToInt(limitStr string, offsetStr string) (int, int, error) {
	var limit int
	var offset int
	var err error

	// int に変換
	limit, err = strconv.Atoi(limitStr)
	if err != nil || limit < 1 {
		return limit, offset, err
	}

	offset, err = strconv.Atoi(offsetStr)
	if err != nil || offset < 0 {
		return limit, offset, err
	}

	return limit, offset, nil
}
