/*
 * TheGamesDB API
 *
 * API Documentations
 *
 * API version: 1.1.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package gamesdb

// GamesImagesAllOfData struct for GamesImagesAllOfData
type GamesImagesAllOfData struct {
	Count   int32                  `json:"count"`
	BaseUrl ImageBaseUrlMeta       `json:"base_url"`
	Images  map[string][]GameImage `json:"images"`
}
