/*
 * Ory Kratos API
 *
 * Documentation for all public and administrative Ory Kratos APIs. Public and administrative APIs are exposed on different ports. Public APIs can face the public internet without any protection while administrative APIs should never be exposed without prior authorization. To protect the administative API port you should use something like Nginx, Ory Oathkeeper, or any other technology capable of authorizing incoming requests. 
 *
 * API version: v0.6.2-alpha.1
 * Contact: hi@ory.sh
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// RevokeSession struct for RevokeSession
type RevokeSession struct {
	// The Session Token  Invalidate this session token.
	SessionToken string `json:"session_token"`
}

// NewRevokeSession instantiates a new RevokeSession object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRevokeSession(sessionToken string) *RevokeSession {
	this := RevokeSession{}
	this.SessionToken = sessionToken
	return &this
}

// NewRevokeSessionWithDefaults instantiates a new RevokeSession object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRevokeSessionWithDefaults() *RevokeSession {
	this := RevokeSession{}
	return &this
}

// GetSessionToken returns the SessionToken field value
func (o *RevokeSession) GetSessionToken() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SessionToken
}

// GetSessionTokenOk returns a tuple with the SessionToken field value
// and a boolean to check if the value has been set.
func (o *RevokeSession) GetSessionTokenOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.SessionToken, true
}

// SetSessionToken sets field value
func (o *RevokeSession) SetSessionToken(v string) {
	o.SessionToken = v
}

func (o RevokeSession) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["session_token"] = o.SessionToken
	}
	return json.Marshal(toSerialize)
}

type NullableRevokeSession struct {
	value *RevokeSession
	isSet bool
}

func (v NullableRevokeSession) Get() *RevokeSession {
	return v.value
}

func (v *NullableRevokeSession) Set(val *RevokeSession) {
	v.value = val
	v.isSet = true
}

func (v NullableRevokeSession) IsSet() bool {
	return v.isSet
}

func (v *NullableRevokeSession) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRevokeSession(val *RevokeSession) *NullableRevokeSession {
	return &NullableRevokeSession{value: val, isSet: true}
}

func (v NullableRevokeSession) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRevokeSession) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


