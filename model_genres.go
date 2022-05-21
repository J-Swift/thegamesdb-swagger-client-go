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

// Genres struct for Genres
type Genres struct {
	Code                      int32           `json:"code"`
	Status                    string          `json:"status"`
	RemainingMonthlyAllowance int32           `json:"remaining_monthly_allowance"`
	ExtraAllowance            int32           `json:"extra_allowance"`
	Data                      GenresAllOfData `json:"data"`
}

// NewGenres instantiates a new Genres object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGenres(code int32, status string, remainingMonthlyAllowance int32, extraAllowance int32, data GenresAllOfData) *Genres {
	this := Genres{}
	this.Code = code
	this.Status = status
	this.RemainingMonthlyAllowance = remainingMonthlyAllowance
	this.ExtraAllowance = extraAllowance
	this.Data = data
	return &this
}

// NewGenresWithDefaults instantiates a new Genres object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGenresWithDefaults() *Genres {
	this := Genres{}
	return &this
}

// GetCode returns the Code field value
func (o *Genres) GetCode() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Code
}

// GetCodeOk returns a tuple with the Code field value
// and a boolean to check if the value has been set.
func (o *Genres) GetCodeOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Code, true
}

// SetCode sets field value
func (o *Genres) SetCode(v int32) {
	o.Code = v
}

// GetStatus returns the Status field value
func (o *Genres) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *Genres) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *Genres) SetStatus(v string) {
	o.Status = v
}

// GetRemainingMonthlyAllowance returns the RemainingMonthlyAllowance field value
func (o *Genres) GetRemainingMonthlyAllowance() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.RemainingMonthlyAllowance
}

// GetRemainingMonthlyAllowanceOk returns a tuple with the RemainingMonthlyAllowance field value
// and a boolean to check if the value has been set.
func (o *Genres) GetRemainingMonthlyAllowanceOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RemainingMonthlyAllowance, true
}

// SetRemainingMonthlyAllowance sets field value
func (o *Genres) SetRemainingMonthlyAllowance(v int32) {
	o.RemainingMonthlyAllowance = v
}

// GetExtraAllowance returns the ExtraAllowance field value
func (o *Genres) GetExtraAllowance() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.ExtraAllowance
}

// GetExtraAllowanceOk returns a tuple with the ExtraAllowance field value
// and a boolean to check if the value has been set.
func (o *Genres) GetExtraAllowanceOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ExtraAllowance, true
}

// SetExtraAllowance sets field value
func (o *Genres) SetExtraAllowance(v int32) {
	o.ExtraAllowance = v
}

// GetData returns the Data field value
func (o *Genres) GetData() GenresAllOfData {
	if o == nil {
		var ret GenresAllOfData
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *Genres) GetDataOk() (*GenresAllOfData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *Genres) SetData(v GenresAllOfData) {
	o.Data = v
}

func (o Genres) MarshalJSON() ([]byte, error) {
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
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullableGenres struct {
	value *Genres
	isSet bool
}

func (v NullableGenres) Get() *Genres {
	return v.value
}

func (v *NullableGenres) Set(val *Genres) {
	v.value = val
	v.isSet = true
}

func (v NullableGenres) IsSet() bool {
	return v.isSet
}

func (v *NullableGenres) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGenres(val *Genres) *NullableGenres {
	return &NullableGenres{value: val, isSet: true}
}

func (v NullableGenres) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGenres) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
