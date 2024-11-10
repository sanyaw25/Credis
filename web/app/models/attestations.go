package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Attestation struct {
	ID              primitive.ObjectID `bson:"_id,omitempty"`
	UserID          string             `bson:"user_id"`
	AttestationType string             `bson:"attestation_type"`
	Content         string             `bson:"content"`
	Timestamp       int64              `bson:"timestamp"`
	FileURL         string             `bson:"file_url"`
}
