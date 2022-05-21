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

// PlatformImage struct for PlatformImage
type PlatformImage struct {
	Id       *int32  `json:"id,omitempty"`
	Type     *string `json:"type,omitempty"`
	Filename *string `json:"filename,omitempty"`
}

// NewPlatformImage instantiates a new PlatformImage object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPlatformImage() *PlatformImage {
	this := PlatformImage{}
	return &this
}

// NewPlatformImageWithDefaults instantiates a new PlatformImage object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPlatformImageWithDefaults() *PlatformImage {
	this := PlatformImage{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *PlatformImage) GetId() int32 {
	if o == nil || o.Id == nil {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlatformImage) GetIdOk() (*int32, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *PlatformImage) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *PlatformImage) SetId(v int32) {
	o.Id = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *PlatformImage) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlatformImage) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *PlatformImage) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *PlatformImage) SetType(v string) {
	o.Type = &v
}

// GetFilename returns the Filename field value if set, zero value otherwise.
func (o *PlatformImage) GetFilename() string {
	if o == nil || o.Filename == nil {
		var ret string
		return ret
	}
	return *o.Filename
}

// GetFilenameOk returns a tuple with the Filename field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlatformImage) GetFilenameOk() (*string, bool) {
	if o == nil || o.Filename == nil {
		return nil, false
	}
	return o.Filename, true
}

// HasFilename returns a boolean if a field has been set.
func (o *PlatformImage) HasFilename() bool {
	if o != nil && o.Filename != nil {
		return true
	}

	return false
}

// SetFilename gets a reference to the given string and assigns it to the Filename field.
func (o *PlatformImage) SetFilename(v string) {
	o.Filename = &v
}

func (o PlatformImage) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	if o.Filename != nil {
		toSerialize["filename"] = o.Filename
	}
	return json.Marshal(toSerialize)
}

type NullablePlatformImage struct {
	value *PlatformImage
	isSet bool
}

func (v NullablePlatformImage) Get() *PlatformImage {
	return v.value
}

func (v *NullablePlatformImage) Set(val *PlatformImage) {
	v.value = val
	v.isSet = true
}

func (v NullablePlatformImage) IsSet() bool {
	return v.isSet
}

func (v *NullablePlatformImage) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePlatformImage(val *PlatformImage) *NullablePlatformImage {
	return &NullablePlatformImage{value: val, isSet: true}
}

func (v NullablePlatformImage) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePlatformImage) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
