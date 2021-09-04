/*
 * Ory Kratos API
 *
 * Documentation for all public and administrative Ory Kratos APIs. Public and administrative APIs are exposed on different ports. Public APIs can face the public internet without any protection while administrative APIs should never be exposed without prior authorization. To protect the administative API port you should use something like Nginx, Ory Oathkeeper, or any other technology capable of authorizing incoming requests. 
 *
 * API version: v0.7.3-alpha.8
 * Contact: hi@ory.sh
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// VolumeUsageData VolumeUsageData Usage details about the volume. This information is used by the `GET /system/df` endpoint, and omitted in other endpoints.
type VolumeUsageData struct {
	// The number of containers referencing this volume. This field is set to `-1` if the reference-count is not available.
	RefCount int64 `json:"RefCount"`
	// Amount of disk space used by the volume (in bytes). This information is only available for volumes created with the `\"local\"` volume driver. For volumes created with other volume drivers, this field is set to `-1` (\"not available\")
	Size int64 `json:"Size"`
}

// NewVolumeUsageData instantiates a new VolumeUsageData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVolumeUsageData(refCount int64, size int64) *VolumeUsageData {
	this := VolumeUsageData{}
	this.RefCount = refCount
	this.Size = size
	return &this
}

// NewVolumeUsageDataWithDefaults instantiates a new VolumeUsageData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVolumeUsageDataWithDefaults() *VolumeUsageData {
	this := VolumeUsageData{}
	return &this
}

// GetRefCount returns the RefCount field value
func (o *VolumeUsageData) GetRefCount() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.RefCount
}

// GetRefCountOk returns a tuple with the RefCount field value
// and a boolean to check if the value has been set.
func (o *VolumeUsageData) GetRefCountOk() (*int64, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.RefCount, true
}

// SetRefCount sets field value
func (o *VolumeUsageData) SetRefCount(v int64) {
	o.RefCount = v
}

// GetSize returns the Size field value
func (o *VolumeUsageData) GetSize() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.Size
}

// GetSizeOk returns a tuple with the Size field value
// and a boolean to check if the value has been set.
func (o *VolumeUsageData) GetSizeOk() (*int64, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Size, true
}

// SetSize sets field value
func (o *VolumeUsageData) SetSize(v int64) {
	o.Size = v
}

func (o VolumeUsageData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["RefCount"] = o.RefCount
	}
	if true {
		toSerialize["Size"] = o.Size
	}
	return json.Marshal(toSerialize)
}

type NullableVolumeUsageData struct {
	value *VolumeUsageData
	isSet bool
}

func (v NullableVolumeUsageData) Get() *VolumeUsageData {
	return v.value
}

func (v *NullableVolumeUsageData) Set(val *VolumeUsageData) {
	v.value = val
	v.isSet = true
}

func (v NullableVolumeUsageData) IsSet() bool {
	return v.isSet
}

func (v *NullableVolumeUsageData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVolumeUsageData(val *VolumeUsageData) *NullableVolumeUsageData {
	return &NullableVolumeUsageData{value: val, isSet: true}
}

func (v NullableVolumeUsageData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVolumeUsageData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


