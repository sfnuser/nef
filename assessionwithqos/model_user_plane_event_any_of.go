/*
3gpp-as-session-with-qos

API for setting us an AS session with required QoS. © 2021, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved. 

API version: 1.1.3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package assessionwithqos

import (
	"encoding/json"
	"fmt"
)

// UserPlaneEventAnyOf the model 'UserPlaneEventAnyOf'
type UserPlaneEventAnyOf string

// List of UserPlaneEvent_anyOf
const (
	SESSION_TERMINATION UserPlaneEventAnyOf = "SESSION_TERMINATION"
	LOSS_OF_BEARER UserPlaneEventAnyOf = "LOSS_OF_BEARER"
	RECOVERY_OF_BEARER UserPlaneEventAnyOf = "RECOVERY_OF_BEARER"
	RELEASE_OF_BEARER UserPlaneEventAnyOf = "RELEASE_OF_BEARER"
	USAGE_REPORT UserPlaneEventAnyOf = "USAGE_REPORT"
	FAILED_RESOURCES_ALLOCATION UserPlaneEventAnyOf = "FAILED_RESOURCES_ALLOCATION"
	QOS_GUARANTEED UserPlaneEventAnyOf = "QOS_GUARANTEED"
	QOS_NOT_GUARANTEED UserPlaneEventAnyOf = "QOS_NOT_GUARANTEED"
	QOS_MONITORING UserPlaneEventAnyOf = "QOS_MONITORING"
	SUCCESSFUL_RESOURCES_ALLOCATION UserPlaneEventAnyOf = "SUCCESSFUL_RESOURCES_ALLOCATION"
)

var allowedUserPlaneEventAnyOfEnumValues = []UserPlaneEventAnyOf{
	"SESSION_TERMINATION",
	"LOSS_OF_BEARER",
	"RECOVERY_OF_BEARER",
	"RELEASE_OF_BEARER",
	"USAGE_REPORT",
	"FAILED_RESOURCES_ALLOCATION",
	"QOS_GUARANTEED",
	"QOS_NOT_GUARANTEED",
	"QOS_MONITORING",
	"SUCCESSFUL_RESOURCES_ALLOCATION",
}

func (v *UserPlaneEventAnyOf) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := UserPlaneEventAnyOf(value)
	for _, existing := range allowedUserPlaneEventAnyOfEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid UserPlaneEventAnyOf", value)
}

// NewUserPlaneEventAnyOfFromValue returns a pointer to a valid UserPlaneEventAnyOf
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewUserPlaneEventAnyOfFromValue(v string) (*UserPlaneEventAnyOf, error) {
	ev := UserPlaneEventAnyOf(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for UserPlaneEventAnyOf: valid values are %v", v, allowedUserPlaneEventAnyOfEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v UserPlaneEventAnyOf) IsValid() bool {
	for _, existing := range allowedUserPlaneEventAnyOfEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to UserPlaneEvent_anyOf value
func (v UserPlaneEventAnyOf) Ptr() *UserPlaneEventAnyOf {
	return &v
}

type NullableUserPlaneEventAnyOf struct {
	value *UserPlaneEventAnyOf
	isSet bool
}

func (v NullableUserPlaneEventAnyOf) Get() *UserPlaneEventAnyOf {
	return v.value
}

func (v *NullableUserPlaneEventAnyOf) Set(val *UserPlaneEventAnyOf) {
	v.value = val
	v.isSet = true
}

func (v NullableUserPlaneEventAnyOf) IsSet() bool {
	return v.isSet
}

func (v *NullableUserPlaneEventAnyOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserPlaneEventAnyOf(val *UserPlaneEventAnyOf) *NullableUserPlaneEventAnyOf {
	return &NullableUserPlaneEventAnyOf{value: val, isSet: true}
}

func (v NullableUserPlaneEventAnyOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserPlaneEventAnyOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

