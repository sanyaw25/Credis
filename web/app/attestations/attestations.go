package attestations

import (
	"Credis/web/app/db"
	"Credis/web/app/models"
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func CreateAttestation(c *gin.Context) {
	var attestation models.Attestation
	if err := c.ShouldBindJSON(&attestation); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	fileURL := c.DefaultPostForm("file_url", "")
	if fileURL == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "File URL is required"})
		return
	}

	attestation.Timestamp = time.Now().Unix()
	attestation.FileURL = fileURL

	collection := db.Client.Database("attestations").Collection("attestation")

	_, err := collection.InsertOne(context.Background(), attestation)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to insert attestation into MongoDB"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Attestation created successfully", "attestation": attestation})
}
