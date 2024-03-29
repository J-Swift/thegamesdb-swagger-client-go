/*
TheGamesDB API

API Documentation

API version: 2.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package gamesdb

import (
	"encoding/json"
)

// GamesByGameIDV1 struct for GamesByGameIDV1
type GamesByGameIDV1 struct {
	Code                      int32                          `json:"code"`
	Status                    string                         `json:"status"`
	RemainingMonthlyAllowance int32                          `json:"remaining_monthly_allowance"`
	ExtraAllowance            int32                          `json:"extra_allowance"`
	Pages                     PaginatedApiResponseAllOfPages `json:"pages"`
	Data                      GamesByGameIDAllOfData         `json:"data"`
	Include                   GamesByGameIDV1AllOfInclude    `json:"include"`
}

// NewGamesByGameIDV1 instantiates a new GamesByGameIDV1 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGamesByGameIDV1(code int32, status string, remainingMonthlyAllowance int32, extraAllowance int32, pages PaginatedApiResponseAllOfPages, data GamesByGameIDAllOfData, include GamesByGameIDV1AllOfInclude) *GamesByGameIDV1 {
	this := GamesByGameIDV1{}
	this.Code = code
	this.Status = status
	this.RemainingMonthlyAllowance = remainingMonthlyAllowance
	this.ExtraAllowance = extraAllowance
	this.Pages = pages
	this.Data = data
	this.Include = include
	return &this
}

// NewGamesByGameIDV1WithDefaults instantiates a new GamesByGameIDV1 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGamesByGameIDV1WithDefaults() *GamesByGameIDV1 {
	this := GamesByGameIDV1{}
	return &this
}

// GetCode returns the Code field value
func (o *GamesByGameIDV1) GetCode() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Code
}

// GetCodeOk returns a tuple with the Code field value
// and a boolean to check if the value has been set.
func (o *GamesByGameIDV1) GetCodeOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Code, true
}

// SetCode sets field value
func (o *GamesByGameIDV1) SetCode(v int32) {
	o.Code = v
}

// GetStatus returns the Status field value
func (o *GamesByGameIDV1) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *GamesByGameIDV1) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *GamesByGameIDV1) SetStatus(v string) {
	o.Status = v
}

// GetRemainingMonthlyAllowance returns the RemainingMonthlyAllowance field value
func (o *GamesByGameIDV1) GetRemainingMonthlyAllowance() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.RemainingMonthlyAllowance
}

// GetRemainingMonthlyAllowanceOk returns a tuple with the RemainingMonthlyAllowance field value
// and a boolean to check if the value has been set.
func (o *GamesByGameIDV1) GetRemainingMonthlyAllowanceOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RemainingMonthlyAllowance, true
}

// SetRemainingMonthlyAllowance sets field value
func (o *GamesByGameIDV1) SetRemainingMonthlyAllowance(v int32) {
	o.RemainingMonthlyAllowance = v
}

// GetExtraAllowance returns the ExtraAllowance field value
func (o *GamesByGameIDV1) GetExtraAllowance() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.ExtraAllowance
}

// GetExtraAllowanceOk returns a tuple with the ExtraAllowance field value
// and a boolean to check if the value has been set.
func (o *GamesByGameIDV1) GetExtraAllowanceOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ExtraAllowance, true
}

// SetExtraAllowance sets field value
func (o *GamesByGameIDV1) SetExtraAllowance(v int32) {
	o.ExtraAllowance = v
}

// GetPages returns the Pages field value
func (o *GamesByGameIDV1) GetPages() PaginatedApiResponseAllOfPages {
	if o == nil {
		var ret PaginatedApiResponseAllOfPages
		return ret
	}

	return o.Pages
}

// GetPagesOk returns a tuple with the Pages field value
// and a boolean to check if the value has been set.
func (o *GamesByGameIDV1) GetPagesOk() (*PaginatedApiResponseAllOfPages, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Pages, true
}

// SetPages sets field value
func (o *GamesByGameIDV1) SetPages(v PaginatedApiResponseAllOfPages) {
	o.Pages = v
}

// GetData returns the Data field value
func (o *GamesByGameIDV1) GetData() GamesByGameIDAllOfData {
	if o == nil {
		var ret GamesByGameIDAllOfData
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *GamesByGameIDV1) GetDataOk() (*GamesByGameIDAllOfData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *GamesByGameIDV1) SetData(v GamesByGameIDAllOfData) {
	o.Data = v
}

// GetInclude returns the Include field value
func (o *GamesByGameIDV1) GetInclude() GamesByGameIDV1AllOfInclude {
	if o == nil {
		var ret GamesByGameIDV1AllOfInclude
		return ret
	}

	return o.Include
}

// GetIncludeOk returns a tuple with the Include field value
// and a boolean to check if the value has been set.
func (o *GamesByGameIDV1) GetIncludeOk() (*GamesByGameIDV1AllOfInclude, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Include, true
}

// SetInclude sets field value
func (o *GamesByGameIDV1) SetInclude(v GamesByGameIDV1AllOfInclude) {
	o.Include = v
}

func (o GamesByGameIDV1) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["code"] = o.Code
	}
	if true {
		toSerialize["status"] = o.Status
	}
	if true {
		toSerialize["remaining_monthly_allowance"] = o.RemainingMonthlyAllowance
	}
	if true {
		toSerialize["extra_allowance"] = o.ExtraAllowance
	}
	if true {
		toSerialize["pages"] = o.Pages
	}
	if true {
		toSerialize["data"] = o.Data
	}
	if true {
		toSerialize["include"] = o.Include
	}
	return json.Marshal(toSerialize)
}

type NullableGamesByGameIDV1 struct {
	value *GamesByGameIDV1
	isSet bool
}

func (v NullableGamesByGameIDV1) Get() *GamesByGameIDV1 {
	return v.value
}

func (v *NullableGamesByGameIDV1) Set(val *GamesByGameIDV1) {
	v.value = val
	v.isSet = true
}

func (v NullableGamesByGameIDV1) IsSet() bool {
	return v.isSet
}

func (v *NullableGamesByGameIDV1) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGamesByGameIDV1(val *GamesByGameIDV1) *NullableGamesByGameIDV1 {
	return &NullableGamesByGameIDV1{value: val, isSet: true}
}

func (v NullableGamesByGameIDV1) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGamesByGameIDV1) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
