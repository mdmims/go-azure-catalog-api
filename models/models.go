package models

import (
	"database/sql"
	"errors"
)

type AssetType struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

// getAssetType retrieves a specific assetType record
func (p *AssetType) getAssetType(db *sql.DB) error {
	return errors.New("Not implemented")
}

// updateAssetType updates an existing assetType record
func (p *AssetType) updateAssetType(db *sql.DB) error {
	return errors.New("Not implemented")
}

// deleteAssetType deletes an existing assetType record
func (p *AssetType) deleteAssetType(db *sql.DB) error {
	return errors.New("Not implemented")
}

// createAssetType creates a new assetType record
func (p *AssetType) createAssetType(db *sql.DB) error {
	return errors.New("Not implemented")
}

// getAssetTypes retrieves all assetType records
func getAssetTypes(db *sql.DB, start, count int) ([]AssetType, error) {
	return nil, errors.New("Not implemented")
}
