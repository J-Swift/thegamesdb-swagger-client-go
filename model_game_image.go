/*
 * TheGamesDB API
 *
 * API Documentation
 *
 * API version: 2.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package gamesdb

// GameImage struct for GameImage
type GameImage struct {
	Id         int32  `json:"id,omitempty"`
	Type       string `json:"type,omitempty"`
	Side       string `json:"side,omitempty"`
	Filename   string `json:"filename,omitempty"`
	Resolution string `json:"resolution,omitempty"`
}
