# AsSessionWithQoSSubscriptionPatch

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FlowInfo** | Pointer to [**[]FlowInfo**](FlowInfo.md) | Describe the data flow which requires QoS. | [optional] 
**EthFlowInfo** | Pointer to [**[]EthFlowDescription**](EthFlowDescription.md) | Identifies Ethernet packet flows. | [optional] 
**QosReference** | Pointer to **string** | Pre-defined QoS reference | [optional] 
**AltQoSReferences** | Pointer to **[]string** | Identifies an ordered list of pre-defined QoS information. The lower the index of the array for a given entry, the higher the priority. | [optional] 
**DisUeNotif** | Pointer to **bool** |  | [optional] 
**UsageThreshold** | Pointer to [**NullableUsageThresholdRm**](UsageThresholdRm.md) |  | [optional] 
**QosMonInfo** | Pointer to [**QosMonitoringInformationRm**](QosMonitoringInformationRm.md) |  | [optional] 

## Methods

### NewAsSessionWithQoSSubscriptionPatch

`func NewAsSessionWithQoSSubscriptionPatch() *AsSessionWithQoSSubscriptionPatch`

NewAsSessionWithQoSSubscriptionPatch instantiates a new AsSessionWithQoSSubscriptionPatch object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAsSessionWithQoSSubscriptionPatchWithDefaults

`func NewAsSessionWithQoSSubscriptionPatchWithDefaults() *AsSessionWithQoSSubscriptionPatch`

NewAsSessionWithQoSSubscriptionPatchWithDefaults instantiates a new AsSessionWithQoSSubscriptionPatch object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFlowInfo

`func (o *AsSessionWithQoSSubscriptionPatch) GetFlowInfo() []FlowInfo`

GetFlowInfo returns the FlowInfo field if non-nil, zero value otherwise.

### GetFlowInfoOk

`func (o *AsSessionWithQoSSubscriptionPatch) GetFlowInfoOk() (*[]FlowInfo, bool)`

GetFlowInfoOk returns a tuple with the FlowInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlowInfo

`func (o *AsSessionWithQoSSubscriptionPatch) SetFlowInfo(v []FlowInfo)`

SetFlowInfo sets FlowInfo field to given value.

### HasFlowInfo

`func (o *AsSessionWithQoSSubscriptionPatch) HasFlowInfo() bool`

HasFlowInfo returns a boolean if a field has been set.

### GetEthFlowInfo

`func (o *AsSessionWithQoSSubscriptionPatch) GetEthFlowInfo() []EthFlowDescription`

GetEthFlowInfo returns the EthFlowInfo field if non-nil, zero value otherwise.

### GetEthFlowInfoOk

`func (o *AsSessionWithQoSSubscriptionPatch) GetEthFlowInfoOk() (*[]EthFlowDescription, bool)`

GetEthFlowInfoOk returns a tuple with the EthFlowInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEthFlowInfo

`func (o *AsSessionWithQoSSubscriptionPatch) SetEthFlowInfo(v []EthFlowDescription)`

SetEthFlowInfo sets EthFlowInfo field to given value.

### HasEthFlowInfo

`func (o *AsSessionWithQoSSubscriptionPatch) HasEthFlowInfo() bool`

HasEthFlowInfo returns a boolean if a field has been set.

### GetQosReference

`func (o *AsSessionWithQoSSubscriptionPatch) GetQosReference() string`

GetQosReference returns the QosReference field if non-nil, zero value otherwise.

### GetQosReferenceOk

`func (o *AsSessionWithQoSSubscriptionPatch) GetQosReferenceOk() (*string, bool)`

GetQosReferenceOk returns a tuple with the QosReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQosReference

`func (o *AsSessionWithQoSSubscriptionPatch) SetQosReference(v string)`

SetQosReference sets QosReference field to given value.

### HasQosReference

`func (o *AsSessionWithQoSSubscriptionPatch) HasQosReference() bool`

HasQosReference returns a boolean if a field has been set.

### GetAltQoSReferences

`func (o *AsSessionWithQoSSubscriptionPatch) GetAltQoSReferences() []string`

GetAltQoSReferences returns the AltQoSReferences field if non-nil, zero value otherwise.

### GetAltQoSReferencesOk

`func (o *AsSessionWithQoSSubscriptionPatch) GetAltQoSReferencesOk() (*[]string, bool)`

GetAltQoSReferencesOk returns a tuple with the AltQoSReferences field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAltQoSReferences

`func (o *AsSessionWithQoSSubscriptionPatch) SetAltQoSReferences(v []string)`

SetAltQoSReferences sets AltQoSReferences field to given value.

### HasAltQoSReferences

`func (o *AsSessionWithQoSSubscriptionPatch) HasAltQoSReferences() bool`

HasAltQoSReferences returns a boolean if a field has been set.

### GetDisUeNotif

`func (o *AsSessionWithQoSSubscriptionPatch) GetDisUeNotif() bool`

GetDisUeNotif returns the DisUeNotif field if non-nil, zero value otherwise.

### GetDisUeNotifOk

`func (o *AsSessionWithQoSSubscriptionPatch) GetDisUeNotifOk() (*bool, bool)`

GetDisUeNotifOk returns a tuple with the DisUeNotif field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisUeNotif

`func (o *AsSessionWithQoSSubscriptionPatch) SetDisUeNotif(v bool)`

SetDisUeNotif sets DisUeNotif field to given value.

### HasDisUeNotif

`func (o *AsSessionWithQoSSubscriptionPatch) HasDisUeNotif() bool`

HasDisUeNotif returns a boolean if a field has been set.

### GetUsageThreshold

`func (o *AsSessionWithQoSSubscriptionPatch) GetUsageThreshold() UsageThresholdRm`

GetUsageThreshold returns the UsageThreshold field if non-nil, zero value otherwise.

### GetUsageThresholdOk

`func (o *AsSessionWithQoSSubscriptionPatch) GetUsageThresholdOk() (*UsageThresholdRm, bool)`

GetUsageThresholdOk returns a tuple with the UsageThreshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsageThreshold

`func (o *AsSessionWithQoSSubscriptionPatch) SetUsageThreshold(v UsageThresholdRm)`

SetUsageThreshold sets UsageThreshold field to given value.

### HasUsageThreshold

`func (o *AsSessionWithQoSSubscriptionPatch) HasUsageThreshold() bool`

HasUsageThreshold returns a boolean if a field has been set.

### SetUsageThresholdNil

`func (o *AsSessionWithQoSSubscriptionPatch) SetUsageThresholdNil(b bool)`

 SetUsageThresholdNil sets the value for UsageThreshold to be an explicit nil

### UnsetUsageThreshold
`func (o *AsSessionWithQoSSubscriptionPatch) UnsetUsageThreshold()`

UnsetUsageThreshold ensures that no value is present for UsageThreshold, not even an explicit nil
### GetQosMonInfo

`func (o *AsSessionWithQoSSubscriptionPatch) GetQosMonInfo() QosMonitoringInformationRm`

GetQosMonInfo returns the QosMonInfo field if non-nil, zero value otherwise.

### GetQosMonInfoOk

`func (o *AsSessionWithQoSSubscriptionPatch) GetQosMonInfoOk() (*QosMonitoringInformationRm, bool)`

GetQosMonInfoOk returns a tuple with the QosMonInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQosMonInfo

`func (o *AsSessionWithQoSSubscriptionPatch) SetQosMonInfo(v QosMonitoringInformationRm)`

SetQosMonInfo sets QosMonInfo field to given value.

### HasQosMonInfo

`func (o *AsSessionWithQoSSubscriptionPatch) HasQosMonInfo() bool`

HasQosMonInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


