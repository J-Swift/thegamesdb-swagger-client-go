/*
 * TheGamesDB API
 *
 * API Documentations
 *
 * API version: 1.1.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package gamesdb

// GamesUpdatesAllOfData struct for GamesUpdatesAllOfData
type GamesUpdatesAllOfData struct {
	Count   int32         `json:"count"`
	Updates []UpdateModel `json:"updates"`
}