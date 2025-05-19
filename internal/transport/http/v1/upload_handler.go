package v1

import (
	"context"
	minioclient "dg-server/internal/minioClient"
	"fmt"
	"net/http"

	minioSDK "github.com/minio/minio-go/v7"

	"github.com/labstack/echo/v4"
)

func UploadEvidence(c echo.Context) error {
	// Parse the multipart form
	fileHeader, err := c.FormFile("file")
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "File not found in request"})
	}

	// Open the uploaded file
	file, err := fileHeader.Open()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to open uploaded file"})
	}
	defer file.Close()

	// Upload to MinIO
	objectName := fileHeader.Filename
	bucket := "evidence-tasks"
	contentType := fileHeader.Header.Get("Content-Type")

	upInfo, err := minioclient.Client.PutObject(context.Background(), bucket, objectName, file, fileHeader.Size, minioSDK.PutObjectOptions{
		ContentType: contentType,
	})

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to upload to MinIO"})
	}
	fmt.Println(upInfo)
	return c.JSON(http.StatusOK, map[string]string{"message": "File uploaded successfully", "object": objectName, "location": upInfo.Location})
}
