/*
 * TheGamesDB API
 *
 * API Documentations
 *
 * API version: 1.1.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package gamesdb

// PlatformsAllOfData struct for PlatformsAllOfData
type PlatformsAllOfData struct {
	Count     int32               `json:"count"`
	Platforms map[string]Platform `json:"platforms"`
}
