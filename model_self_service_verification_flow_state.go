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
	"fmt"
)

// SelfServiceVerificationFlowState The state represents the state of the verification flow.  choose_method: ask the user to choose a method (e.g. recover account via email) sent_email: the email has been sent to the user passed_challenge: the request was successful and the recovery challenge was passed.
type SelfServiceVerificationFlowState string

// List of selfServiceVerificationFlowState
const (
	SELFSERVICEVERIFICATIONFLOWSTATE_CHOOSE_METHOD SelfServiceVerificationFlowState = "choose_method"
	SELFSERVICEVERIFICATIONFLOWSTATE_SENT_EMAIL SelfServiceVerificationFlowState = "sent_email"
	SELFSERVICEVERIFICATIONFLOWSTATE_PASSED_CHALLENGE SelfServiceVerificationFlowState = "passed_challenge"
)

// All allowed values of SelfServiceVerificationFlowState enum
var AllowedSelfServiceVerificationFlowStateEnumValues = []SelfServiceVerificationFlowState{
	"choose_method",
	"sent_email",
	"passed_challenge",
}

func (v *SelfServiceVerificationFlowState) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := SelfServiceVerificationFlowState(value)
	for _, existing := range AllowedSelfServiceVerificationFlowStateEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid SelfServiceVerificationFlowState", value)
}

// NewSelfServiceVerificationFlowStateFromValue returns a pointer to a valid SelfServiceVerificationFlowState
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewSelfServiceVerificationFlowStateFromValue(v string) (*SelfServiceVerificationFlowState, error) {
	ev := SelfServiceVerificationFlowState(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for SelfServiceVerificationFlowState: valid values are %v", v, AllowedSelfServiceVerificationFlowStateEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v SelfServiceVerificationFlowState) IsValid() bool {
	for _, existing := range AllowedSelfServiceVerificationFlowStateEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to selfServiceVerificationFlowState value
func (v SelfServiceVerificationFlowState) Ptr() *SelfServiceVerificationFlowState {
	return &v
}

type NullableSelfServiceVerificationFlowState struct {
	value *SelfServiceVerificationFlowState
	isSet bool
}

func (v NullableSelfServiceVerificationFlowState) Get() *SelfServiceVerificationFlowState {
	return v.value
}

func (v *NullableSelfServiceVerificationFlowState) Set(val *SelfServiceVerificationFlowState) {
	v.value = val
	v.isSet = true
}

func (v NullableSelfServiceVerificationFlowState) IsSet() bool {
	return v.isSet
}

func (v *NullableSelfServiceVerificationFlowState) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSelfServiceVerificationFlowState(val *SelfServiceVerificationFlowState) *NullableSelfServiceVerificationFlowState {
	return &NullableSelfServiceVerificationFlowState{value: val, isSet: true}
}

func (v NullableSelfServiceVerificationFlowState) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSelfServiceVerificationFlowState) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

