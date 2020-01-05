/*
 * TheGamesDB API
 *
 * API Documentations
 *
 * API version: 1.1.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package gamesdb

// GamesByGameIdV1 struct for GamesByGameIdV1
type GamesByGameIdV1 struct {
	Code                      int32                          `json:"code"`
	Status                    string                         `json:"status"`
	RemainingMonthlyAllowance int32                          `json:"remaining_monthly_allowance"`
	ExtraAllowance            int32                          `json:"extra_allowance"`
	Pages                     PaginatedApiResponseAllOfPages `json:"pages"`
	Data                      GamesByGameIdAllOfData         `json:"data"`
	Include                   GamesByGameIdV1AllOfInclude    `json:"include"`
}
