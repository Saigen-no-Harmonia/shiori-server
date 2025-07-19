package controllers

import (
	"net/http"
	"shiori-server/database"
	"shiori-server/models"
	"shiori-server/models/resource"
	"shiori-server/util"
	"strconv"

	"github.com/gin-gonic/gin"
)

// GetGalleryPhotos ギャラリー画像の一覧を取得
// @Summary ギャラリー画像取得
// @Description ギャラリー画像の一覧をlimitとoffsetで取得
// @Tags gallery
// @Accept json
// @Produce json
// @Param limit query int false "取得件数" default(20)
// @Param offset query int false "開始位置" default(0)
// @Success 200 {object} resource.GalleryPhotoPageResource
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /gallery [get]
func GetGalleryPhotos(c *gin.Context) {
	limitStr := c.DefaultQuery("limit", "20")
	offsetStr := c.DefaultQuery("offset", "0")

	// int に変換
	limit, err := strconv.Atoi(limitStr)
	if err != nil || limit < 1 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid limit"})
		return
	}

	offset, err := strconv.Atoi(offsetStr)
	if err != nil || offset < 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid offset"})
		return
	}

	// DBから画像取得
	rows, err := database.DB.Query("SELECT id, s3_object_name, display_number "+
		"FROM M_GALLERY_PHOTO "+
		"WHERE delete_flag = 0 "+
		"ORDER BY display_number ASC "+
		"LIMIT ? OFFSET ?",
		limit,
		offset,
	)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer rows.Close()

	// 取得したデータをモデルに格納
	var galleryPhotos []models.GalleryPhoto
	for rows.Next() {
		var p models.GalleryPhoto
		if err := rows.Scan(
			&p.Id,
			&p.S3ObjectName,
			&p.DisplayNumber,
		); err != nil {
			c.JSON(
				http.StatusInternalServerError,
				gin.H{"error": err.Error()},
			)
		}

		// S3アクセス用URLを取得・格納
		p.PhotoUrl = util.GetS3AccessUrl(p.S3ObjectName)

		galleryPhotos = append(galleryPhotos, p)
	}

	// Resourceに変換
	var galleryPhotoPageResources []resource.GalleryPhotoResource
	for index := 0; index < len(galleryPhotos); index++ {
		src := galleryPhotos[index]
		dest := resource.MapToGalleryResource(src)
		galleryPhotoPageResources = append(galleryPhotoPageResources, *dest)
	}

	// レスポンスを生成
	response := resource.NewGalleryPhotoPageResource(galleryPhotoPageResources)

	c.JSON(http.StatusOK, response)
}
