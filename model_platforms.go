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

// Platforms struct for Platforms
type Platforms struct {
	Code                      int32              `json:"code"`
	Status                    string             `json:"status"`
	RemainingMonthlyAllowance int32              `json:"remaining_monthly_allowance"`
	ExtraAllowance            int32              `json:"extra_allowance"`
	Data                      PlatformsAllOfData `json:"data"`
}

// NewPlatforms instantiates a new Platforms object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPlatforms(code int32, status string, remainingMonthlyAllowance int32, extraAllowance int32, data PlatformsAllOfData) *Platforms {
	this := Platforms{}
	this.Code = code
	this.Status = status
	this.RemainingMonthlyAllowance = remainingMonthlyAllowance
	this.ExtraAllowance = extraAllowance
	this.Data = data
	return &this
}

// NewPlatformsWithDefaults instantiates a new Platforms object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPlatformsWithDefaults() *Platforms {
	this := Platforms{}
	return &this
}

// GetCode returns the Code field value
func (o *Platforms) GetCode() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Code
}

// GetCodeOk returns a tuple with the Code field value
// and a boolean to check if the value has been set.
func (o *Platforms) GetCodeOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Code, true
}

// SetCode sets field value
func (o *Platforms) SetCode(v int32) {
	o.Code = v
}

// GetStatus returns the Status field value
func (o *Platforms) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *Platforms) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *Platforms) SetStatus(v string) {
	o.Status = v
}

// GetRemainingMonthlyAllowance returns the RemainingMonthlyAllowance field value
func (o *Platforms) GetRemainingMonthlyAllowance() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.RemainingMonthlyAllowance
}

// GetRemainingMonthlyAllowanceOk returns a tuple with the RemainingMonthlyAllowance field value
// and a boolean to check if the value has been set.
func (o *Platforms) GetRemainingMonthlyAllowanceOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RemainingMonthlyAllowance, true
}

// SetRemainingMonthlyAllowance sets field value
func (o *Platforms) SetRemainingMonthlyAllowance(v int32) {
	o.RemainingMonthlyAllowance = v
}

// GetExtraAllowance returns the ExtraAllowance field value
func (o *Platforms) GetExtraAllowance() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.ExtraAllowance
}

// GetExtraAllowanceOk returns a tuple with the ExtraAllowance field value
// and a boolean to check if the value has been set.
func (o *Platforms) GetExtraAllowanceOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ExtraAllowance, true
}

// SetExtraAllowance sets field value
func (o *Platforms) SetExtraAllowance(v int32) {
	o.ExtraAllowance = v
}

// GetData returns the Data field value
func (o *Platforms) GetData() PlatformsAllOfData {
	if o == nil {
		var ret PlatformsAllOfData
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *Platforms) GetDataOk() (*PlatformsAllOfData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *Platforms) SetData(v PlatformsAllOfData) {
	o.Data = v
}

func (o Platforms) MarshalJSON() ([]byte, error) {
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

type NullablePlatforms struct {
	value *Platforms
	isSet bool
}

func (v NullablePlatforms) Get() *Platforms {
	return v.value
}

func (v *NullablePlatforms) Set(val *Platforms) {
	v.value = val
	v.isSet = true
}

func (v NullablePlatforms) IsSet() bool {
	return v.isSet
}

func (v *NullablePlatforms) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePlatforms(val *Platforms) *NullablePlatforms {
	return &NullablePlatforms{value: val, isSet: true}
}

func (v NullablePlatforms) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePlatforms) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
