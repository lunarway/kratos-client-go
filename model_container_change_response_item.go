/*
 * Ory Kratos API
 *
 * Documentation for all public and administrative Ory Kratos APIs. Public and administrative APIs are exposed on different ports. Public APIs can face the public internet without any protection while administrative APIs should never be exposed without prior authorization. To protect the administative API port you should use something like Nginx, Ory Oathkeeper, or any other technology capable of authorizing incoming requests. 
 *
 * API version: v0.6.0-alpha.2
 * Contact: hi@ory.sh
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// ContainerChangeResponseItem ContainerChangeResponseItem change item in response to ContainerChanges operation
type ContainerChangeResponseItem struct {
	// Kind of change
	Kind int32 `json:"Kind"`
	// Path to file that has changed
	Path string `json:"Path"`
}

// NewContainerChangeResponseItem instantiates a new ContainerChangeResponseItem object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewContainerChangeResponseItem(kind int32, path string) *ContainerChangeResponseItem {
	this := ContainerChangeResponseItem{}
	this.Kind = kind
	this.Path = path
	return &this
}

// NewContainerChangeResponseItemWithDefaults instantiates a new ContainerChangeResponseItem object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewContainerChangeResponseItemWithDefaults() *ContainerChangeResponseItem {
	this := ContainerChangeResponseItem{}
	return &this
}

// GetKind returns the Kind field value
func (o *ContainerChangeResponseItem) GetKind() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Kind
}

// GetKindOk returns a tuple with the Kind field value
// and a boolean to check if the value has been set.
func (o *ContainerChangeResponseItem) GetKindOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Kind, true
}

// SetKind sets field value
func (o *ContainerChangeResponseItem) SetKind(v int32) {
	o.Kind = v
}

// GetPath returns the Path field value
func (o *ContainerChangeResponseItem) GetPath() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Path
}

// GetPathOk returns a tuple with the Path field value
// and a boolean to check if the value has been set.
func (o *ContainerChangeResponseItem) GetPathOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Path, true
}

// SetPath sets field value
func (o *ContainerChangeResponseItem) SetPath(v string) {
	o.Path = v
}

func (o ContainerChangeResponseItem) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["Kind"] = o.Kind
	}
	if true {
		toSerialize["Path"] = o.Path
	}
	return json.Marshal(toSerialize)
}

type NullableContainerChangeResponseItem struct {
	value *ContainerChangeResponseItem
	isSet bool
}

func (v NullableContainerChangeResponseItem) Get() *ContainerChangeResponseItem {
	return v.value
}

func (v *NullableContainerChangeResponseItem) Set(val *ContainerChangeResponseItem) {
	v.value = val
	v.isSet = true
}

func (v NullableContainerChangeResponseItem) IsSet() bool {
	return v.isSet
}

func (v *NullableContainerChangeResponseItem) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableContainerChangeResponseItem(val *ContainerChangeResponseItem) *NullableContainerChangeResponseItem {
	return &NullableContainerChangeResponseItem{value: val, isSet: true}
}

func (v NullableContainerChangeResponseItem) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableContainerChangeResponseItem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


