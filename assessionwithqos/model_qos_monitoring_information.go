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

// QosMonitoringInformation struct for QosMonitoringInformation
type QosMonitoringInformation struct {
	ReqQosMonParams []RequestedQosMonitoringParameter `json:"reqQosMonParams"`
	RepFreqs []ReportingFrequency `json:"repFreqs"`
	RepThreshDl *int32 `json:"repThreshDl,omitempty"`
	RepThreshUl *int32 `json:"repThreshUl,omitempty"`
	RepThreshRp *int32 `json:"repThreshRp,omitempty"`
	WaitTime *int32 `json:"waitTime,omitempty"`
	RepPeriod *int32 `json:"repPeriod,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _QosMonitoringInformation QosMonitoringInformation

// NewQosMonitoringInformation instantiates a new QosMonitoringInformation object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewQosMonitoringInformation(reqQosMonParams []RequestedQosMonitoringParameter, repFreqs []ReportingFrequency) *QosMonitoringInformation {
	this := QosMonitoringInformation{}
	this.ReqQosMonParams = reqQosMonParams
	this.RepFreqs = repFreqs
	return &this
}

// NewQosMonitoringInformationWithDefaults instantiates a new QosMonitoringInformation object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewQosMonitoringInformationWithDefaults() *QosMonitoringInformation {
	this := QosMonitoringInformation{}
	return &this
}

// GetReqQosMonParams returns the ReqQosMonParams field value
func (o *QosMonitoringInformation) GetReqQosMonParams() []RequestedQosMonitoringParameter {
	if o == nil {
		var ret []RequestedQosMonitoringParameter
		return ret
	}

	return o.ReqQosMonParams
}

// GetReqQosMonParamsOk returns a tuple with the ReqQosMonParams field value
// and a boolean to check if the value has been set.
func (o *QosMonitoringInformation) GetReqQosMonParamsOk() (*[]RequestedQosMonitoringParameter, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ReqQosMonParams, true
}

// SetReqQosMonParams sets field value
func (o *QosMonitoringInformation) SetReqQosMonParams(v []RequestedQosMonitoringParameter) {
	o.ReqQosMonParams = v
}

// GetRepFreqs returns the RepFreqs field value
func (o *QosMonitoringInformation) GetRepFreqs() []ReportingFrequency {
	if o == nil {
		var ret []ReportingFrequency
		return ret
	}

	return o.RepFreqs
}

// GetRepFreqsOk returns a tuple with the RepFreqs field value
// and a boolean to check if the value has been set.
func (o *QosMonitoringInformation) GetRepFreqsOk() (*[]ReportingFrequency, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.RepFreqs, true
}

// SetRepFreqs sets field value
func (o *QosMonitoringInformation) SetRepFreqs(v []ReportingFrequency) {
	o.RepFreqs = v
}

// GetRepThreshDl returns the RepThreshDl field value if set, zero value otherwise.
func (o *QosMonitoringInformation) GetRepThreshDl() int32 {
	if o == nil || o.RepThreshDl == nil {
		var ret int32
		return ret
	}
	return *o.RepThreshDl
}

// GetRepThreshDlOk returns a tuple with the RepThreshDl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QosMonitoringInformation) GetRepThreshDlOk() (*int32, bool) {
	if o == nil || o.RepThreshDl == nil {
		return nil, false
	}
	return o.RepThreshDl, true
}

// HasRepThreshDl returns a boolean if a field has been set.
func (o *QosMonitoringInformation) HasRepThreshDl() bool {
	if o != nil && o.RepThreshDl != nil {
		return true
	}

	return false
}

// SetRepThreshDl gets a reference to the given int32 and assigns it to the RepThreshDl field.
func (o *QosMonitoringInformation) SetRepThreshDl(v int32) {
	o.RepThreshDl = &v
}

// GetRepThreshUl returns the RepThreshUl field value if set, zero value otherwise.
func (o *QosMonitoringInformation) GetRepThreshUl() int32 {
	if o == nil || o.RepThreshUl == nil {
		var ret int32
		return ret
	}
	return *o.RepThreshUl
}

// GetRepThreshUlOk returns a tuple with the RepThreshUl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QosMonitoringInformation) GetRepThreshUlOk() (*int32, bool) {
	if o == nil || o.RepThreshUl == nil {
		return nil, false
	}
	return o.RepThreshUl, true
}

// HasRepThreshUl returns a boolean if a field has been set.
func (o *QosMonitoringInformation) HasRepThreshUl() bool {
	if o != nil && o.RepThreshUl != nil {
		return true
	}

	return false
}

// SetRepThreshUl gets a reference to the given int32 and assigns it to the RepThreshUl field.
func (o *QosMonitoringInformation) SetRepThreshUl(v int32) {
	o.RepThreshUl = &v
}

// GetRepThreshRp returns the RepThreshRp field value if set, zero value otherwise.
func (o *QosMonitoringInformation) GetRepThreshRp() int32 {
	if o == nil || o.RepThreshRp == nil {
		var ret int32
		return ret
	}
	return *o.RepThreshRp
}

// GetRepThreshRpOk returns a tuple with the RepThreshRp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QosMonitoringInformation) GetRepThreshRpOk() (*int32, bool) {
	if o == nil || o.RepThreshRp == nil {
		return nil, false
	}
	return o.RepThreshRp, true
}

// HasRepThreshRp returns a boolean if a field has been set.
func (o *QosMonitoringInformation) HasRepThreshRp() bool {
	if o != nil && o.RepThreshRp != nil {
		return true
	}

	return false
}

// SetRepThreshRp gets a reference to the given int32 and assigns it to the RepThreshRp field.
func (o *QosMonitoringInformation) SetRepThreshRp(v int32) {
	o.RepThreshRp = &v
}

// GetWaitTime returns the WaitTime field value if set, zero value otherwise.
func (o *QosMonitoringInformation) GetWaitTime() int32 {
	if o == nil || o.WaitTime == nil {
		var ret int32
		return ret
	}
	return *o.WaitTime
}

// GetWaitTimeOk returns a tuple with the WaitTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QosMonitoringInformation) GetWaitTimeOk() (*int32, bool) {
	if o == nil || o.WaitTime == nil {
		return nil, false
	}
	return o.WaitTime, true
}

// HasWaitTime returns a boolean if a field has been set.
func (o *QosMonitoringInformation) HasWaitTime() bool {
	if o != nil && o.WaitTime != nil {
		return true
	}

	return false
}

// SetWaitTime gets a reference to the given int32 and assigns it to the WaitTime field.
func (o *QosMonitoringInformation) SetWaitTime(v int32) {
	o.WaitTime = &v
}

// GetRepPeriod returns the RepPeriod field value if set, zero value otherwise.
func (o *QosMonitoringInformation) GetRepPeriod() int32 {
	if o == nil || o.RepPeriod == nil {
		var ret int32
		return ret
	}
	return *o.RepPeriod
}

// GetRepPeriodOk returns a tuple with the RepPeriod field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QosMonitoringInformation) GetRepPeriodOk() (*int32, bool) {
	if o == nil || o.RepPeriod == nil {
		return nil, false
	}
	return o.RepPeriod, true
}

// HasRepPeriod returns a boolean if a field has been set.
func (o *QosMonitoringInformation) HasRepPeriod() bool {
	if o != nil && o.RepPeriod != nil {
		return true
	}

	return false
}

// SetRepPeriod gets a reference to the given int32 and assigns it to the RepPeriod field.
func (o *QosMonitoringInformation) SetRepPeriod(v int32) {
	o.RepPeriod = &v
}

func (o QosMonitoringInformation) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["reqQosMonParams"] = o.ReqQosMonParams
	}
	if true {
		toSerialize["repFreqs"] = o.RepFreqs
	}
	if o.RepThreshDl != nil {
		toSerialize["repThreshDl"] = o.RepThreshDl
	}
	if o.RepThreshUl != nil {
		toSerialize["repThreshUl"] = o.RepThreshUl
	}
	if o.RepThreshRp != nil {
		toSerialize["repThreshRp"] = o.RepThreshRp
	}
	if o.WaitTime != nil {
		toSerialize["waitTime"] = o.WaitTime
	}
	if o.RepPeriod != nil {
		toSerialize["repPeriod"] = o.RepPeriod
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *QosMonitoringInformation) UnmarshalJSON(bytes []byte) (err error) {
	varQosMonitoringInformation := _QosMonitoringInformation{}

	if err = json.Unmarshal(bytes, &varQosMonitoringInformation); err == nil {
		*o = QosMonitoringInformation(varQosMonitoringInformation)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "reqQosMonParams")
		delete(additionalProperties, "repFreqs")
		delete(additionalProperties, "repThreshDl")
		delete(additionalProperties, "repThreshUl")
		delete(additionalProperties, "repThreshRp")
		delete(additionalProperties, "waitTime")
		delete(additionalProperties, "repPeriod")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableQosMonitoringInformation struct {
	value *QosMonitoringInformation
	isSet bool
}

func (v NullableQosMonitoringInformation) Get() *QosMonitoringInformation {
	return v.value
}

func (v *NullableQosMonitoringInformation) Set(val *QosMonitoringInformation) {
	v.value = val
	v.isSet = true
}

func (v NullableQosMonitoringInformation) IsSet() bool {
	return v.isSet
}

func (v *NullableQosMonitoringInformation) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableQosMonitoringInformation(val *QosMonitoringInformation) *NullableQosMonitoringInformation {
	return &NullableQosMonitoringInformation{value: val, isSet: true}
}

func (v NullableQosMonitoringInformation) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableQosMonitoringInformation) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

