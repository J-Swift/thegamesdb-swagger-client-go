/*
 * TheGamesDB API
 *
 * API Documentation
 *
 * API version: 2.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package gamesdb

// PlatformsImages struct for PlatformsImages
type PlatformsImages struct {
	Code                      int32                          `json:"code"`
	Status                    string                         `json:"status"`
	RemainingMonthlyAllowance int32                          `json:"remaining_monthly_allowance"`
	ExtraAllowance            int32                          `json:"extra_allowance"`
	Pages                     PaginatedApiResponseAllOfPages `json:"pages"`
	Data                      PlatformsImagesAllOfData       `json:"data"`
}
