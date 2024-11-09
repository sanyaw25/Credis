package upload

import (
	"Credis/web/app/db"
	"Credis/web/app/models"
	"fmt"
	"net/http"
	"path/filepath"
	"time"

	"github.com/gin-gonic/gin"
)

// UploadFile handles the file upload and creates an attestation.
func UploadFile(c *gin.Context) {
	// Get user info (e.g., nickname from the context or token)
	// For now, assuming the nickname is passed as a query parameter or header (adjust as per your auth system)
	nickname := c.GetString("nickname") // Assuming you have middleware to set the nickname

	// Upload the file using form-data
	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "File upload failed"})
		return
	}

	// Save the file to the server's local storage
	uploadPath := filepath.Join("uploads", time.Now().Format("2006-01-02_15-04-05")+"_"+file.Filename)
	if err := c.SaveUploadedFile(file, uploadPath); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not save file"})
		return
	}

	// Create an attestation document using the file name and the user's nickname
	attestation := models.Attestation{
		UserID:          nickname,
		AttestationType: "File Upload",
		Content:         fmt.Sprintf("Attestation for file: %s uploaded by %s", file.Filename, nickname),
		Timestamp:       time.Now().Unix(),
		FileURL:         uploadPath,
	}

	// Save the attestation to MongoDB
	collection := db.Client.Database("attestations").Collection("attestation")
	_, err = collection.InsertOne(c, attestation)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to insert attestation into MongoDB"})
		return
	}

	// Respond with the file URL and confirmation
	c.JSON(http.StatusOK, gin.H{
		"message":     "File uploaded and attestation created successfully!",
		"file_url":    uploadPath,
		"attestation": attestation,
	})
}
