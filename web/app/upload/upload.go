package upload

import (
	"Credis/web/app/db"
	"Credis/web/app/models"
	"fmt"
	"net/http"
	"os/exec"
	"path/filepath"
	"time"

	"github.com/gin-gonic/gin"
)

// UploadFile handles the file upload, creates an attestation, and executes a Python script.
func UploadFile(c *gin.Context) {
	// Get user info (e.g., nickname from the context or token)
	nickname := c.GetString("nickname") // Assuming middleware sets the nickname

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
	// Create an attestation document
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

	go runScript()

	// Respond with the file URL, attestation details, and Python script output
	c.JSON(http.StatusOK, gin.H{
		"message":     "File uploaded, attestation created",
		"file_url":    uploadPath,
		"attestation": attestation,
	})

}

func runScript() {
	cmd := exec.Command("python3", "/mnt/Disk_2/hack_cbs/project/Credis/model/model.py", "CS 101", "Introduction to Computer Science")
	err := cmd.Run()
	if err != nil {
		fmt.Println(err)
	}
}
