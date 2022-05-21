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

// PublishersAllOfData struct for PublishersAllOfData
type PublishersAllOfData struct {
	Count      int32                `json:"count"`
	Publishers map[string]Publisher `json:"publishers"`
}

// NewPublishersAllOfData instantiates a new PublishersAllOfData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPublishersAllOfData(count int32, publishers map[string]Publisher) *PublishersAllOfData {
	this := PublishersAllOfData{}
	this.Count = count
	this.Publishers = publishers
	return &this
}

// NewPublishersAllOfDataWithDefaults instantiates a new PublishersAllOfData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPublishersAllOfDataWithDefaults() *PublishersAllOfData {
	this := PublishersAllOfData{}
	return &this
}

// GetCount returns the Count field value
func (o *PublishersAllOfData) GetCount() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Count
}

// GetCountOk returns a tuple with the Count field value
// and a boolean to check if the value has been set.
func (o *PublishersAllOfData) GetCountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Count, true
}

// SetCount sets field value
func (o *PublishersAllOfData) SetCount(v int32) {
	o.Count = v
}

// GetPublishers returns the Publishers field value
func (o *PublishersAllOfData) GetPublishers() map[string]Publisher {
	if o == nil {
		var ret map[string]Publisher
		return ret
	}

	return o.Publishers
}

// GetPublishersOk returns a tuple with the Publishers field value
// and a boolean to check if the value has been set.
func (o *PublishersAllOfData) GetPublishersOk() (*map[string]Publisher, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Publishers, true
}

// SetPublishers sets field value
func (o *PublishersAllOfData) SetPublishers(v map[string]Publisher) {
	o.Publishers = v
}

func (o PublishersAllOfData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["count"] = o.Count
	}
	if true {
		toSerialize["publishers"] = o.Publishers
	}
	return json.Marshal(toSerialize)
}

type NullablePublishersAllOfData struct {
	value *PublishersAllOfData
	isSet bool
}

func (v NullablePublishersAllOfData) Get() *PublishersAllOfData {
	return v.value
}

func (v *NullablePublishersAllOfData) Set(val *PublishersAllOfData) {
	v.value = val
	v.isSet = true
}

func (v NullablePublishersAllOfData) IsSet() bool {
	return v.isSet
}

func (v *NullablePublishersAllOfData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePublishersAllOfData(val *PublishersAllOfData) *NullablePublishersAllOfData {
	return &NullablePublishersAllOfData{value: val, isSet: true}
}

func (v NullablePublishersAllOfData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePublishersAllOfData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
