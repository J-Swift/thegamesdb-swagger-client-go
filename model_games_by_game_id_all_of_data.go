/*
 * TheGamesDB API
 *
 * API Documentation
 *
 * API version: 2.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package gamesdb

// GamesByGameIdAllOfData struct for GamesByGameIdAllOfData
type GamesByGameIdAllOfData struct {
	Count int32  `json:"count"`
	Games []Game `json:"games"`
}
