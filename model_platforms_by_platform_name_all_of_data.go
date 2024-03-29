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

// PlatformsByPlatformNameAllOfData struct for PlatformsByPlatformNameAllOfData
type PlatformsByPlatformNameAllOfData struct {
	Count     int32      `json:"count"`
	Platforms []Platform `json:"platforms"`
}

// NewPlatformsByPlatformNameAllOfData instantiates a new PlatformsByPlatformNameAllOfData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPlatformsByPlatformNameAllOfData(count int32, platforms []Platform) *PlatformsByPlatformNameAllOfData {
	this := PlatformsByPlatformNameAllOfData{}
	this.Count = count
	this.Platforms = platforms
	return &this
}

// NewPlatformsByPlatformNameAllOfDataWithDefaults instantiates a new PlatformsByPlatformNameAllOfData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPlatformsByPlatformNameAllOfDataWithDefaults() *PlatformsByPlatformNameAllOfData {
	this := PlatformsByPlatformNameAllOfData{}
	return &this
}

// GetCount returns the Count field value
func (o *PlatformsByPlatformNameAllOfData) GetCount() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Count
}

// GetCountOk returns a tuple with the Count field value
// and a boolean to check if the value has been set.
func (o *PlatformsByPlatformNameAllOfData) GetCountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Count, true
}

// SetCount sets field value
func (o *PlatformsByPlatformNameAllOfData) SetCount(v int32) {
	o.Count = v
}

// GetPlatforms returns the Platforms field value
func (o *PlatformsByPlatformNameAllOfData) GetPlatforms() []Platform {
	if o == nil {
		var ret []Platform
		return ret
	}

	return o.Platforms
}

// GetPlatformsOk returns a tuple with the Platforms field value
// and a boolean to check if the value has been set.
func (o *PlatformsByPlatformNameAllOfData) GetPlatformsOk() ([]Platform, bool) {
	if o == nil {
		return nil, false
	}
	return o.Platforms, true
}

// SetPlatforms sets field value
func (o *PlatformsByPlatformNameAllOfData) SetPlatforms(v []Platform) {
	o.Platforms = v
}

func (o PlatformsByPlatformNameAllOfData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["count"] = o.Count
	}
	if true {
		toSerialize["platforms"] = o.Platforms
	}
	return json.Marshal(toSerialize)
}

type NullablePlatformsByPlatformNameAllOfData struct {
	value *PlatformsByPlatformNameAllOfData
	isSet bool
}

func (v NullablePlatformsByPlatformNameAllOfData) Get() *PlatformsByPlatformNameAllOfData {
	return v.value
}

func (v *NullablePlatformsByPlatformNameAllOfData) Set(val *PlatformsByPlatformNameAllOfData) {
	v.value = val
	v.isSet = true
}

func (v NullablePlatformsByPlatformNameAllOfData) IsSet() bool {
	return v.isSet
}

func (v *NullablePlatformsByPlatformNameAllOfData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePlatformsByPlatformNameAllOfData(val *PlatformsByPlatformNameAllOfData) *NullablePlatformsByPlatformNameAllOfData {
	return &NullablePlatformsByPlatformNameAllOfData{value: val, isSet: true}
}

func (v NullablePlatformsByPlatformNameAllOfData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePlatformsByPlatformNameAllOfData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
