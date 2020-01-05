/*
 * TheGamesDB API
 *
 * API Documentations
 *
 * API version: 1.1.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package gamesdb

// Developers struct for Developers
type Developers struct {
	Code                      int32               `json:"code"`
	Status                    string              `json:"status"`
	RemainingMonthlyAllowance int32               `json:"remaining_monthly_allowance"`
	ExtraAllowance            int32               `json:"extra_allowance"`
	Data                      DevelopersAllOfData `json:"data"`
}
