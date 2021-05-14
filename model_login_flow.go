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
	"time"
)

// LoginFlow This object represents a login flow. A login flow is initiated at the \"Initiate Login API / Browser Flow\" endpoint by a client.  Once a login flow is completed successfully, a session cookie or session token will be issued.
type LoginFlow struct {
	// and so on.
	Active *string `json:"active,omitempty"`
	// ExpiresAt is the time (UTC) when the flow expires. If the user still wishes to log in, a new flow has to be initiated.
	ExpiresAt time.Time `json:"expires_at"`
	// Forced stores whether this login flow should enforce re-authentication.
	Forced *bool `json:"forced,omitempty"`
	Id string `json:"id"`
	// IssuedAt is the time (UTC) when the flow started.
	IssuedAt time.Time `json:"issued_at"`
	// RequestURL is the initial URL that was requested from Ory Kratos. It can be used to forward information contained in the URL's path or query for example.
	RequestUrl string `json:"request_url"`
	// The flow type can either be `api` or `browser`.
	Type string `json:"type"`
	Ui UiContainer `json:"ui"`
}

// NewLoginFlow instantiates a new LoginFlow object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLoginFlow(expiresAt time.Time, id string, issuedAt time.Time, requestUrl string, type_ string, ui UiContainer) *LoginFlow {
	this := LoginFlow{}
	this.ExpiresAt = expiresAt
	this.Id = id
	this.IssuedAt = issuedAt
	this.RequestUrl = requestUrl
	this.Type = type_
	this.Ui = ui
	return &this
}

// NewLoginFlowWithDefaults instantiates a new LoginFlow object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLoginFlowWithDefaults() *LoginFlow {
	this := LoginFlow{}
	return &this
}

// GetActive returns the Active field value if set, zero value otherwise.
func (o *LoginFlow) GetActive() string {
	if o == nil || o.Active == nil {
		var ret string
		return ret
	}
	return *o.Active
}

// GetActiveOk returns a tuple with the Active field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LoginFlow) GetActiveOk() (*string, bool) {
	if o == nil || o.Active == nil {
		return nil, false
	}
	return o.Active, true
}

// HasActive returns a boolean if a field has been set.
func (o *LoginFlow) HasActive() bool {
	if o != nil && o.Active != nil {
		return true
	}

	return false
}

// SetActive gets a reference to the given string and assigns it to the Active field.
func (o *LoginFlow) SetActive(v string) {
	o.Active = &v
}

// GetExpiresAt returns the ExpiresAt field value
func (o *LoginFlow) GetExpiresAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.ExpiresAt
}

// GetExpiresAtOk returns a tuple with the ExpiresAt field value
// and a boolean to check if the value has been set.
func (o *LoginFlow) GetExpiresAtOk() (*time.Time, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ExpiresAt, true
}

// SetExpiresAt sets field value
func (o *LoginFlow) SetExpiresAt(v time.Time) {
	o.ExpiresAt = v
}

// GetForced returns the Forced field value if set, zero value otherwise.
func (o *LoginFlow) GetForced() bool {
	if o == nil || o.Forced == nil {
		var ret bool
		return ret
	}
	return *o.Forced
}

// GetForcedOk returns a tuple with the Forced field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LoginFlow) GetForcedOk() (*bool, bool) {
	if o == nil || o.Forced == nil {
		return nil, false
	}
	return o.Forced, true
}

// HasForced returns a boolean if a field has been set.
func (o *LoginFlow) HasForced() bool {
	if o != nil && o.Forced != nil {
		return true
	}

	return false
}

// SetForced gets a reference to the given bool and assigns it to the Forced field.
func (o *LoginFlow) SetForced(v bool) {
	o.Forced = &v
}

// GetId returns the Id field value
func (o *LoginFlow) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *LoginFlow) GetIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *LoginFlow) SetId(v string) {
	o.Id = v
}

// GetIssuedAt returns the IssuedAt field value
func (o *LoginFlow) GetIssuedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.IssuedAt
}

// GetIssuedAtOk returns a tuple with the IssuedAt field value
// and a boolean to check if the value has been set.
func (o *LoginFlow) GetIssuedAtOk() (*time.Time, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.IssuedAt, true
}

// SetIssuedAt sets field value
func (o *LoginFlow) SetIssuedAt(v time.Time) {
	o.IssuedAt = v
}

// GetRequestUrl returns the RequestUrl field value
func (o *LoginFlow) GetRequestUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RequestUrl
}

// GetRequestUrlOk returns a tuple with the RequestUrl field value
// and a boolean to check if the value has been set.
func (o *LoginFlow) GetRequestUrlOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.RequestUrl, true
}

// SetRequestUrl sets field value
func (o *LoginFlow) SetRequestUrl(v string) {
	o.RequestUrl = v
}

// GetType returns the Type field value
func (o *LoginFlow) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *LoginFlow) GetTypeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *LoginFlow) SetType(v string) {
	o.Type = v
}

// GetUi returns the Ui field value
func (o *LoginFlow) GetUi() UiContainer {
	if o == nil {
		var ret UiContainer
		return ret
	}

	return o.Ui
}

// GetUiOk returns a tuple with the Ui field value
// and a boolean to check if the value has been set.
func (o *LoginFlow) GetUiOk() (*UiContainer, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Ui, true
}

// SetUi sets field value
func (o *LoginFlow) SetUi(v UiContainer) {
	o.Ui = v
}

func (o LoginFlow) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Active != nil {
		toSerialize["active"] = o.Active
	}
	if true {
		toSerialize["expires_at"] = o.ExpiresAt
	}
	if o.Forced != nil {
		toSerialize["forced"] = o.Forced
	}
	if true {
		toSerialize["id"] = o.Id
	}
	if true {
		toSerialize["issued_at"] = o.IssuedAt
	}
	if true {
		toSerialize["request_url"] = o.RequestUrl
	}
	if true {
		toSerialize["type"] = o.Type
	}
	if true {
		toSerialize["ui"] = o.Ui
	}
	return json.Marshal(toSerialize)
}

type NullableLoginFlow struct {
	value *LoginFlow
	isSet bool
}

func (v NullableLoginFlow) Get() *LoginFlow {
	return v.value
}

func (v *NullableLoginFlow) Set(val *LoginFlow) {
	v.value = val
	v.isSet = true
}

func (v NullableLoginFlow) IsSet() bool {
	return v.isSet
}

func (v *NullableLoginFlow) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLoginFlow(val *LoginFlow) *NullableLoginFlow {
	return &NullableLoginFlow{value: val, isSet: true}
}

func (v NullableLoginFlow) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLoginFlow) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


