/*
Ory Kratos API

Documentation for all public and administrative Ory Kratos APIs. Public and administrative APIs are exposed on different ports. Public APIs can face the public internet without any protection while administrative APIs should never be exposed without prior authorization. To protect the administative API port you should use something like Nginx, Ory Oathkeeper, or any other technology capable of authorizing incoming requests. 

API version: v0.11.0-fork
Contact: hi@ory.sh
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
	"time"
)

// Message struct for Message
type Message struct {
	Body *string `json:"body,omitempty"`
	// CreatedAt is a helper struct field for gobuffalo.pop.
	CreatedAt *time.Time `json:"created_at,omitempty"`
	Id *string `json:"id,omitempty"`
	Recipient *string `json:"recipient,omitempty"`
	SendCount *int64 `json:"send_count,omitempty"`
	Status *CourierMessageStatus `json:"status,omitempty"`
	Subject *string `json:"subject,omitempty"`
	TemplateType *string `json:"template_type,omitempty"`
	Type *CourierMessageType `json:"type,omitempty"`
	// UpdatedAt is a helper struct field for gobuffalo.pop.
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
}

// NewMessage instantiates a new Message object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMessage() *Message {
	this := Message{}
	return &this
}

// NewMessageWithDefaults instantiates a new Message object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMessageWithDefaults() *Message {
	this := Message{}
	return &this
}

// GetBody returns the Body field value if set, zero value otherwise.
func (o *Message) GetBody() string {
	if o == nil || o.Body == nil {
		var ret string
		return ret
	}
	return *o.Body
}

// GetBodyOk returns a tuple with the Body field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Message) GetBodyOk() (*string, bool) {
	if o == nil || o.Body == nil {
		return nil, false
	}
	return o.Body, true
}

// HasBody returns a boolean if a field has been set.
func (o *Message) HasBody() bool {
	if o != nil && o.Body != nil {
		return true
	}

	return false
}

// SetBody gets a reference to the given string and assigns it to the Body field.
func (o *Message) SetBody(v string) {
	o.Body = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *Message) GetCreatedAt() time.Time {
	if o == nil || o.CreatedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Message) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || o.CreatedAt == nil {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *Message) HasCreatedAt() bool {
	if o != nil && o.CreatedAt != nil {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *Message) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *Message) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Message) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *Message) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *Message) SetId(v string) {
	o.Id = &v
}

// GetRecipient returns the Recipient field value if set, zero value otherwise.
func (o *Message) GetRecipient() string {
	if o == nil || o.Recipient == nil {
		var ret string
		return ret
	}
	return *o.Recipient
}

// GetRecipientOk returns a tuple with the Recipient field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Message) GetRecipientOk() (*string, bool) {
	if o == nil || o.Recipient == nil {
		return nil, false
	}
	return o.Recipient, true
}

// HasRecipient returns a boolean if a field has been set.
func (o *Message) HasRecipient() bool {
	if o != nil && o.Recipient != nil {
		return true
	}

	return false
}

// SetRecipient gets a reference to the given string and assigns it to the Recipient field.
func (o *Message) SetRecipient(v string) {
	o.Recipient = &v
}

// GetSendCount returns the SendCount field value if set, zero value otherwise.
func (o *Message) GetSendCount() int64 {
	if o == nil || o.SendCount == nil {
		var ret int64
		return ret
	}
	return *o.SendCount
}

// GetSendCountOk returns a tuple with the SendCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Message) GetSendCountOk() (*int64, bool) {
	if o == nil || o.SendCount == nil {
		return nil, false
	}
	return o.SendCount, true
}

// HasSendCount returns a boolean if a field has been set.
func (o *Message) HasSendCount() bool {
	if o != nil && o.SendCount != nil {
		return true
	}

	return false
}

// SetSendCount gets a reference to the given int64 and assigns it to the SendCount field.
func (o *Message) SetSendCount(v int64) {
	o.SendCount = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *Message) GetStatus() CourierMessageStatus {
	if o == nil || o.Status == nil {
		var ret CourierMessageStatus
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Message) GetStatusOk() (*CourierMessageStatus, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *Message) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given CourierMessageStatus and assigns it to the Status field.
func (o *Message) SetStatus(v CourierMessageStatus) {
	o.Status = &v
}

// GetSubject returns the Subject field value if set, zero value otherwise.
func (o *Message) GetSubject() string {
	if o == nil || o.Subject == nil {
		var ret string
		return ret
	}
	return *o.Subject
}

// GetSubjectOk returns a tuple with the Subject field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Message) GetSubjectOk() (*string, bool) {
	if o == nil || o.Subject == nil {
		return nil, false
	}
	return o.Subject, true
}

// HasSubject returns a boolean if a field has been set.
func (o *Message) HasSubject() bool {
	if o != nil && o.Subject != nil {
		return true
	}

	return false
}

// SetSubject gets a reference to the given string and assigns it to the Subject field.
func (o *Message) SetSubject(v string) {
	o.Subject = &v
}

// GetTemplateType returns the TemplateType field value if set, zero value otherwise.
func (o *Message) GetTemplateType() string {
	if o == nil || o.TemplateType == nil {
		var ret string
		return ret
	}
	return *o.TemplateType
}

// GetTemplateTypeOk returns a tuple with the TemplateType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Message) GetTemplateTypeOk() (*string, bool) {
	if o == nil || o.TemplateType == nil {
		return nil, false
	}
	return o.TemplateType, true
}

// HasTemplateType returns a boolean if a field has been set.
func (o *Message) HasTemplateType() bool {
	if o != nil && o.TemplateType != nil {
		return true
	}

	return false
}

// SetTemplateType gets a reference to the given string and assigns it to the TemplateType field.
func (o *Message) SetTemplateType(v string) {
	o.TemplateType = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *Message) GetType() CourierMessageType {
	if o == nil || o.Type == nil {
		var ret CourierMessageType
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Message) GetTypeOk() (*CourierMessageType, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *Message) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given CourierMessageType and assigns it to the Type field.
func (o *Message) SetType(v CourierMessageType) {
	o.Type = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *Message) GetUpdatedAt() time.Time {
	if o == nil || o.UpdatedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Message) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil || o.UpdatedAt == nil {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *Message) HasUpdatedAt() bool {
	if o != nil && o.UpdatedAt != nil {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given time.Time and assigns it to the UpdatedAt field.
func (o *Message) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = &v
}

func (o Message) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Body != nil {
		toSerialize["body"] = o.Body
	}
	if o.CreatedAt != nil {
		toSerialize["created_at"] = o.CreatedAt
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Recipient != nil {
		toSerialize["recipient"] = o.Recipient
	}
	if o.SendCount != nil {
		toSerialize["send_count"] = o.SendCount
	}
	if o.Status != nil {
		toSerialize["status"] = o.Status
	}
	if o.Subject != nil {
		toSerialize["subject"] = o.Subject
	}
	if o.TemplateType != nil {
		toSerialize["template_type"] = o.TemplateType
	}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	if o.UpdatedAt != nil {
		toSerialize["updated_at"] = o.UpdatedAt
	}
	return json.Marshal(toSerialize)
}

type NullableMessage struct {
	value *Message
	isSet bool
}

func (v NullableMessage) Get() *Message {
	return v.value
}

func (v *NullableMessage) Set(val *Message) {
	v.value = val
	v.isSet = true
}

func (v NullableMessage) IsSet() bool {
	return v.isSet
}

func (v *NullableMessage) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMessage(val *Message) *NullableMessage {
	return &NullableMessage{value: val, isSet: true}
}

func (v NullableMessage) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMessage) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


