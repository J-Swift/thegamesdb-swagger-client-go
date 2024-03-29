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

// Developer struct for Developer
type Developer struct {
	Id   int32  `json:"id"`
	Name string `json:"name"`
}

// NewDeveloper instantiates a new Developer object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeveloper(id int32, name string) *Developer {
	this := Developer{}
	this.Id = id
	this.Name = name
	return &this
}

// NewDeveloperWithDefaults instantiates a new Developer object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeveloperWithDefaults() *Developer {
	this := Developer{}
	return &this
}

// GetId returns the Id field value
func (o *Developer) GetId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *Developer) GetIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *Developer) SetId(v int32) {
	o.Id = v
}

// GetName returns the Name field value
func (o *Developer) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *Developer) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *Developer) SetName(v string) {
	o.Name = v
}

func (o Developer) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["id"] = o.Id
	}
	if true {
		toSerialize["name"] = o.Name
	}
	return json.Marshal(toSerialize)
}

type NullableDeveloper struct {
	value *Developer
	isSet bool
}

func (v NullableDeveloper) Get() *Developer {
	return v.value
}

func (v *NullableDeveloper) Set(val *Developer) {
	v.value = val
	v.isSet = true
}

func (v NullableDeveloper) IsSet() bool {
	return v.isSet
}

func (v *NullableDeveloper) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeveloper(val *Developer) *NullableDeveloper {
	return &NullableDeveloper{value: val, isSet: true}
}

func (v NullableDeveloper) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeveloper) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
