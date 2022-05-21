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

// Platform struct for Platform
type Platform struct {
	Id         int32   `json:"id"`
	Name       string  `json:"name"`
	Alias      string  `json:"alias"`
	Icon       string  `json:"icon"`
	Console    string  `json:"console"`
	Controller string  `json:"controller"`
	Developer  string  `json:"developer"`
	Overview   *string `json:"overview,omitempty"`
}

// NewPlatform instantiates a new Platform object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPlatform(id int32, name string, alias string, icon string, console string, controller string, developer string) *Platform {
	this := Platform{}
	this.Id = id
	this.Name = name
	this.Alias = alias
	this.Icon = icon
	this.Console = console
	this.Controller = controller
	this.Developer = developer
	return &this
}

// NewPlatformWithDefaults instantiates a new Platform object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPlatformWithDefaults() *Platform {
	this := Platform{}
	return &this
}

// GetId returns the Id field value
func (o *Platform) GetId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *Platform) GetIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *Platform) SetId(v int32) {
	o.Id = v
}

// GetName returns the Name field value
func (o *Platform) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *Platform) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *Platform) SetName(v string) {
	o.Name = v
}

// GetAlias returns the Alias field value
func (o *Platform) GetAlias() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Alias
}

// GetAliasOk returns a tuple with the Alias field value
// and a boolean to check if the value has been set.
func (o *Platform) GetAliasOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Alias, true
}

// SetAlias sets field value
func (o *Platform) SetAlias(v string) {
	o.Alias = v
}

// GetIcon returns the Icon field value
func (o *Platform) GetIcon() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Icon
}

// GetIconOk returns a tuple with the Icon field value
// and a boolean to check if the value has been set.
func (o *Platform) GetIconOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Icon, true
}

// SetIcon sets field value
func (o *Platform) SetIcon(v string) {
	o.Icon = v
}

// GetConsole returns the Console field value
func (o *Platform) GetConsole() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Console
}

// GetConsoleOk returns a tuple with the Console field value
// and a boolean to check if the value has been set.
func (o *Platform) GetConsoleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Console, true
}

// SetConsole sets field value
func (o *Platform) SetConsole(v string) {
	o.Console = v
}

// GetController returns the Controller field value
func (o *Platform) GetController() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Controller
}

// GetControllerOk returns a tuple with the Controller field value
// and a boolean to check if the value has been set.
func (o *Platform) GetControllerOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Controller, true
}

// SetController sets field value
func (o *Platform) SetController(v string) {
	o.Controller = v
}

// GetDeveloper returns the Developer field value
func (o *Platform) GetDeveloper() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Developer
}

// GetDeveloperOk returns a tuple with the Developer field value
// and a boolean to check if the value has been set.
func (o *Platform) GetDeveloperOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Developer, true
}

// SetDeveloper sets field value
func (o *Platform) SetDeveloper(v string) {
	o.Developer = v
}

// GetOverview returns the Overview field value if set, zero value otherwise.
func (o *Platform) GetOverview() string {
	if o == nil || o.Overview == nil {
		var ret string
		return ret
	}
	return *o.Overview
}

// GetOverviewOk returns a tuple with the Overview field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Platform) GetOverviewOk() (*string, bool) {
	if o == nil || o.Overview == nil {
		return nil, false
	}
	return o.Overview, true
}

// HasOverview returns a boolean if a field has been set.
func (o *Platform) HasOverview() bool {
	if o != nil && o.Overview != nil {
		return true
	}

	return false
}

// SetOverview gets a reference to the given string and assigns it to the Overview field.
func (o *Platform) SetOverview(v string) {
	o.Overview = &v
}

func (o Platform) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["id"] = o.Id
	}
	if true {
		toSerialize["name"] = o.Name
	}
	if true {
		toSerialize["alias"] = o.Alias
	}
	if true {
		toSerialize["icon"] = o.Icon
	}
	if true {
		toSerialize["console"] = o.Console
	}
	if true {
		toSerialize["controller"] = o.Controller
	}
	if true {
		toSerialize["developer"] = o.Developer
	}
	if o.Overview != nil {
		toSerialize["overview"] = o.Overview
	}
	return json.Marshal(toSerialize)
}

type NullablePlatform struct {
	value *Platform
	isSet bool
}

func (v NullablePlatform) Get() *Platform {
	return v.value
}

func (v *NullablePlatform) Set(val *Platform) {
	v.value = val
	v.isSet = true
}

func (v NullablePlatform) IsSet() bool {
	return v.isSet
}

func (v *NullablePlatform) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePlatform(val *Platform) *NullablePlatform {
	return &NullablePlatform{value: val, isSet: true}
}

func (v NullablePlatform) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePlatform) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
