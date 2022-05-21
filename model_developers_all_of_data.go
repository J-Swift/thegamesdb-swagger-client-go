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

// DevelopersAllOfData struct for DevelopersAllOfData
type DevelopersAllOfData struct {
	Count      int32                `json:"count"`
	Developers map[string]Developer `json:"developers"`
}

// NewDevelopersAllOfData instantiates a new DevelopersAllOfData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDevelopersAllOfData(count int32, developers map[string]Developer) *DevelopersAllOfData {
	this := DevelopersAllOfData{}
	this.Count = count
	this.Developers = developers
	return &this
}

// NewDevelopersAllOfDataWithDefaults instantiates a new DevelopersAllOfData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDevelopersAllOfDataWithDefaults() *DevelopersAllOfData {
	this := DevelopersAllOfData{}
	return &this
}

// GetCount returns the Count field value
func (o *DevelopersAllOfData) GetCount() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Count
}

// GetCountOk returns a tuple with the Count field value
// and a boolean to check if the value has been set.
func (o *DevelopersAllOfData) GetCountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Count, true
}

// SetCount sets field value
func (o *DevelopersAllOfData) SetCount(v int32) {
	o.Count = v
}

// GetDevelopers returns the Developers field value
func (o *DevelopersAllOfData) GetDevelopers() map[string]Developer {
	if o == nil {
		var ret map[string]Developer
		return ret
	}

	return o.Developers
}

// GetDevelopersOk returns a tuple with the Developers field value
// and a boolean to check if the value has been set.
func (o *DevelopersAllOfData) GetDevelopersOk() (*map[string]Developer, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Developers, true
}

// SetDevelopers sets field value
func (o *DevelopersAllOfData) SetDevelopers(v map[string]Developer) {
	o.Developers = v
}

func (o DevelopersAllOfData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["count"] = o.Count
	}
	if true {
		toSerialize["developers"] = o.Developers
	}
	return json.Marshal(toSerialize)
}

type NullableDevelopersAllOfData struct {
	value *DevelopersAllOfData
	isSet bool
}

func (v NullableDevelopersAllOfData) Get() *DevelopersAllOfData {
	return v.value
}

func (v *NullableDevelopersAllOfData) Set(val *DevelopersAllOfData) {
	v.value = val
	v.isSet = true
}

func (v NullableDevelopersAllOfData) IsSet() bool {
	return v.isSet
}

func (v *NullableDevelopersAllOfData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDevelopersAllOfData(val *DevelopersAllOfData) *NullableDevelopersAllOfData {
	return &NullableDevelopersAllOfData{value: val, isSet: true}
}

func (v NullableDevelopersAllOfData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDevelopersAllOfData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
