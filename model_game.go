/*
 * TheGamesDB API
 *
 * API Documentation
 *
 * API version: 2.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package gamesdb

// Game struct for Game
type Game struct {
	Id          int32    `json:"id,omitempty"`
	GameTitle   string   `json:"game_title,omitempty"`
	ReleaseDate string   `json:"release_date,omitempty"`
	Platform    int32    `json:"platform,omitempty"`
	Players     int32    `json:"players,omitempty"`
	Overview    string   `json:"overview,omitempty"`
	LastUpdated string   `json:"last_updated,omitempty"`
	Rating      string   `json:"rating,omitempty"`
	Coop        string   `json:"coop,omitempty"`
	Youtube     string   `json:"youtube,omitempty"`
	Os          string   `json:"os,omitempty"`
	Processor   string   `json:"processor,omitempty"`
	Ram         string   `json:"ram,omitempty"`
	Hdd         string   `json:"hdd,omitempty"`
	Video       string   `json:"video,omitempty"`
	Sound       string   `json:"sound,omitempty"`
	Developers  []int32  `json:"developers,omitempty"`
	Genres      []int32  `json:"genres,omitempty"`
	Publishers  []int32  `json:"publishers,omitempty"`
	Alternates  []string `json:"alternates,omitempty"`
}
