/*
3gpp-as-session-with-qos

API for setting us an AS session with required QoS. © 2021, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved. 

API version: 1.1.3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package assessionwithqos

import (
	"encoding/json"
)

// QosMonitoringReport struct for QosMonitoringReport
type QosMonitoringReport struct {
	UlDelays *[]int32 `json:"ulDelays,omitempty"`
	DlDelays *[]int32 `json:"dlDelays,omitempty"`
	RtDelays *[]int32 `json:"rtDelays,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _QosMonitoringReport QosMonitoringReport

// NewQosMonitoringReport instantiates a new QosMonitoringReport object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewQosMonitoringReport() *QosMonitoringReport {
	this := QosMonitoringReport{}
	return &this
}

// NewQosMonitoringReportWithDefaults instantiates a new QosMonitoringReport object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewQosMonitoringReportWithDefaults() *QosMonitoringReport {
	this := QosMonitoringReport{}
	return &this
}

// GetUlDelays returns the UlDelays field value if set, zero value otherwise.
func (o *QosMonitoringReport) GetUlDelays() []int32 {
	if o == nil || o.UlDelays == nil {
		var ret []int32
		return ret
	}
	return *o.UlDelays
}

// GetUlDelaysOk returns a tuple with the UlDelays field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QosMonitoringReport) GetUlDelaysOk() (*[]int32, bool) {
	if o == nil || o.UlDelays == nil {
		return nil, false
	}
	return o.UlDelays, true
}

// HasUlDelays returns a boolean if a field has been set.
func (o *QosMonitoringReport) HasUlDelays() bool {
	if o != nil && o.UlDelays != nil {
		return true
	}

	return false
}

// SetUlDelays gets a reference to the given []int32 and assigns it to the UlDelays field.
func (o *QosMonitoringReport) SetUlDelays(v []int32) {
	o.UlDelays = &v
}

// GetDlDelays returns the DlDelays field value if set, zero value otherwise.
func (o *QosMonitoringReport) GetDlDelays() []int32 {
	if o == nil || o.DlDelays == nil {
		var ret []int32
		return ret
	}
	return *o.DlDelays
}

// GetDlDelaysOk returns a tuple with the DlDelays field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QosMonitoringReport) GetDlDelaysOk() (*[]int32, bool) {
	if o == nil || o.DlDelays == nil {
		return nil, false
	}
	return o.DlDelays, true
}

// HasDlDelays returns a boolean if a field has been set.
func (o *QosMonitoringReport) HasDlDelays() bool {
	if o != nil && o.DlDelays != nil {
		return true
	}

	return false
}

// SetDlDelays gets a reference to the given []int32 and assigns it to the DlDelays field.
func (o *QosMonitoringReport) SetDlDelays(v []int32) {
	o.DlDelays = &v
}

// GetRtDelays returns the RtDelays field value if set, zero value otherwise.
func (o *QosMonitoringReport) GetRtDelays() []int32 {
	if o == nil || o.RtDelays == nil {
		var ret []int32
		return ret
	}
	return *o.RtDelays
}

// GetRtDelaysOk returns a tuple with the RtDelays field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QosMonitoringReport) GetRtDelaysOk() (*[]int32, bool) {
	if o == nil || o.RtDelays == nil {
		return nil, false
	}
	return o.RtDelays, true
}

// HasRtDelays returns a boolean if a field has been set.
func (o *QosMonitoringReport) HasRtDelays() bool {
	if o != nil && o.RtDelays != nil {
		return true
	}

	return false
}

// SetRtDelays gets a reference to the given []int32 and assigns it to the RtDelays field.
func (o *QosMonitoringReport) SetRtDelays(v []int32) {
	o.RtDelays = &v
}

func (o QosMonitoringReport) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UlDelays != nil {
		toSerialize["ulDelays"] = o.UlDelays
	}
	if o.DlDelays != nil {
		toSerialize["dlDelays"] = o.DlDelays
	}
	if o.RtDelays != nil {
		toSerialize["rtDelays"] = o.RtDelays
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *QosMonitoringReport) UnmarshalJSON(bytes []byte) (err error) {
	varQosMonitoringReport := _QosMonitoringReport{}

	if err = json.Unmarshal(bytes, &varQosMonitoringReport); err == nil {
		*o = QosMonitoringReport(varQosMonitoringReport)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ulDelays")
		delete(additionalProperties, "dlDelays")
		delete(additionalProperties, "rtDelays")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableQosMonitoringReport struct {
	value *QosMonitoringReport
	isSet bool
}

func (v NullableQosMonitoringReport) Get() *QosMonitoringReport {
	return v.value
}

func (v *NullableQosMonitoringReport) Set(val *QosMonitoringReport) {
	v.value = val
	v.isSet = true
}

func (v NullableQosMonitoringReport) IsSet() bool {
	return v.isSet
}

func (v *NullableQosMonitoringReport) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableQosMonitoringReport(val *QosMonitoringReport) *NullableQosMonitoringReport {
	return &NullableQosMonitoringReport{value: val, isSet: true}
}

func (v NullableQosMonitoringReport) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableQosMonitoringReport) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


