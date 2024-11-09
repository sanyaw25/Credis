package attestations

import (
	"Credis/web/app/db"
	"Credis/web/app/models"
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// CreateAttestation handles the creation of an attestation with a file reference.
func CreateAttestation(c *gin.Context) {
	// Bind the form data (attestation details)
	var attestation models.Attestation
	if err := c.ShouldBindJSON(&attestation); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Attach the uploaded file URL/path to the attestation
	fileURL := c.DefaultPostForm("file_url", "")
	if fileURL == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "File URL is required"})
		return
	}

	// Set the timestamp to the current time
	attestation.Timestamp = time.Now().Unix()
	attestation.FileURL = fileURL // Save the file URL/path in the attestation

	// Get the MongoDB collection
	collection := db.Client.Database("attestations").Collection("attestation")

	// Insert the attestation document into the collection
	_, err := collection.InsertOne(context.Background(), attestation)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to insert attestation into MongoDB"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Attestation created successfully", "attestation": attestation})
}
