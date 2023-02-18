# AsSessionWithQoSSubscription

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Self** | Pointer to **string** | string formatted according to IETF RFC 3986 identifying a referenced resource. | [optional] 
**SupportedFeatures** | Pointer to **string** |  | [optional] 
**NotificationDestination** | **string** | string formatted according to IETF RFC 3986 identifying a referenced resource. | 
**FlowInfo** | Pointer to [**[]FlowInfo**](FlowInfo.md) | Describe the data flow which requires QoS. | [optional] 
**EthFlowInfo** | Pointer to [**[]EthFlowDescription**](EthFlowDescription.md) | Identifies Ethernet packet flows. | [optional] 
**QosReference** | Pointer to **string** | Identifies a pre-defined QoS information | [optional] 
**AltQoSReferences** | Pointer to **[]string** | Identifies an ordered list of pre-defined QoS information. The lower the index of the array for a given entry, the higher the priority. | [optional] 
**DisUeNotif** | Pointer to **bool** |  | [optional] 
**UeIpv4Addr** | Pointer to **string** | string identifying a Ipv4 address formatted in the \&quot;dotted decimal\&quot; notation as defined in IETF RFC 1166. | [optional] 
**IpDomain** | Pointer to **string** |  | [optional] 
**UeIpv6Addr** | Pointer to **string** | string identifying a Ipv6 address formatted according to clause 4 in IETF RFC 5952. The mixed Ipv4 Ipv6 notation according to clause 5 of IETF RFC 5952 shall not be used. | [optional] 
**MacAddr** | Pointer to **string** |  | [optional] 
**UsageThreshold** | Pointer to [**UsageThreshold**](UsageThreshold.md) |  | [optional] 
**SponsorInfo** | Pointer to [**SponsorInformation**](SponsorInformation.md) |  | [optional] 
**QosMonInfo** | Pointer to [**QosMonitoringInformation**](QosMonitoringInformation.md) |  | [optional] 
**RequestTestNotification** | Pointer to **bool** | Set to true by the SCS/AS to request the SCEF to send a test notification as defined in subclause 5.2.5.3. Set to false or omitted otherwise. | [optional] 
**WebsockNotifConfig** | Pointer to [**WebsockNotifConfig**](WebsockNotifConfig.md) |  | [optional] 

## Methods

### NewAsSessionWithQoSSubscription

`func NewAsSessionWithQoSSubscription(notificationDestination string, ) *AsSessionWithQoSSubscription`

NewAsSessionWithQoSSubscription instantiates a new AsSessionWithQoSSubscription object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAsSessionWithQoSSubscriptionWithDefaults

`func NewAsSessionWithQoSSubscriptionWithDefaults() *AsSessionWithQoSSubscription`

NewAsSessionWithQoSSubscriptionWithDefaults instantiates a new AsSessionWithQoSSubscription object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSelf

`func (o *AsSessionWithQoSSubscription) GetSelf() string`

GetSelf returns the Self field if non-nil, zero value otherwise.

### GetSelfOk

`func (o *AsSessionWithQoSSubscription) GetSelfOk() (*string, bool)`

GetSelfOk returns a tuple with the Self field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelf

`func (o *AsSessionWithQoSSubscription) SetSelf(v string)`

SetSelf sets Self field to given value.

### HasSelf

`func (o *AsSessionWithQoSSubscription) HasSelf() bool`

HasSelf returns a boolean if a field has been set.

### GetSupportedFeatures

`func (o *AsSessionWithQoSSubscription) GetSupportedFeatures() string`

GetSupportedFeatures returns the SupportedFeatures field if non-nil, zero value otherwise.

### GetSupportedFeaturesOk

`func (o *AsSessionWithQoSSubscription) GetSupportedFeaturesOk() (*string, bool)`

GetSupportedFeaturesOk returns a tuple with the SupportedFeatures field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportedFeatures

`func (o *AsSessionWithQoSSubscription) SetSupportedFeatures(v string)`

SetSupportedFeatures sets SupportedFeatures field to given value.

### HasSupportedFeatures

`func (o *AsSessionWithQoSSubscription) HasSupportedFeatures() bool`

HasSupportedFeatures returns a boolean if a field has been set.

### GetNotificationDestination

`func (o *AsSessionWithQoSSubscription) GetNotificationDestination() string`

GetNotificationDestination returns the NotificationDestination field if non-nil, zero value otherwise.

### GetNotificationDestinationOk

`func (o *AsSessionWithQoSSubscription) GetNotificationDestinationOk() (*string, bool)`

GetNotificationDestinationOk returns a tuple with the NotificationDestination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotificationDestination

`func (o *AsSessionWithQoSSubscription) SetNotificationDestination(v string)`

SetNotificationDestination sets NotificationDestination field to given value.


### GetFlowInfo

`func (o *AsSessionWithQoSSubscription) GetFlowInfo() []FlowInfo`

GetFlowInfo returns the FlowInfo field if non-nil, zero value otherwise.

### GetFlowInfoOk

`func (o *AsSessionWithQoSSubscription) GetFlowInfoOk() (*[]FlowInfo, bool)`

GetFlowInfoOk returns a tuple with the FlowInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlowInfo

`func (o *AsSessionWithQoSSubscription) SetFlowInfo(v []FlowInfo)`

SetFlowInfo sets FlowInfo field to given value.

### HasFlowInfo

`func (o *AsSessionWithQoSSubscription) HasFlowInfo() bool`

HasFlowInfo returns a boolean if a field has been set.

### GetEthFlowInfo

`func (o *AsSessionWithQoSSubscription) GetEthFlowInfo() []EthFlowDescription`

GetEthFlowInfo returns the EthFlowInfo field if non-nil, zero value otherwise.

### GetEthFlowInfoOk

`func (o *AsSessionWithQoSSubscription) GetEthFlowInfoOk() (*[]EthFlowDescription, bool)`

GetEthFlowInfoOk returns a tuple with the EthFlowInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEthFlowInfo

`func (o *AsSessionWithQoSSubscription) SetEthFlowInfo(v []EthFlowDescription)`

SetEthFlowInfo sets EthFlowInfo field to given value.

### HasEthFlowInfo

`func (o *AsSessionWithQoSSubscription) HasEthFlowInfo() bool`

HasEthFlowInfo returns a boolean if a field has been set.

### GetQosReference

`func (o *AsSessionWithQoSSubscription) GetQosReference() string`

GetQosReference returns the QosReference field if non-nil, zero value otherwise.

### GetQosReferenceOk

`func (o *AsSessionWithQoSSubscription) GetQosReferenceOk() (*string, bool)`

GetQosReferenceOk returns a tuple with the QosReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQosReference

`func (o *AsSessionWithQoSSubscription) SetQosReference(v string)`

SetQosReference sets QosReference field to given value.

### HasQosReference

`func (o *AsSessionWithQoSSubscription) HasQosReference() bool`

HasQosReference returns a boolean if a field has been set.

### GetAltQoSReferences

`func (o *AsSessionWithQoSSubscription) GetAltQoSReferences() []string`

GetAltQoSReferences returns the AltQoSReferences field if non-nil, zero value otherwise.

### GetAltQoSReferencesOk

`func (o *AsSessionWithQoSSubscription) GetAltQoSReferencesOk() (*[]string, bool)`

GetAltQoSReferencesOk returns a tuple with the AltQoSReferences field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAltQoSReferences

`func (o *AsSessionWithQoSSubscription) SetAltQoSReferences(v []string)`

SetAltQoSReferences sets AltQoSReferences field to given value.

### HasAltQoSReferences

`func (o *AsSessionWithQoSSubscription) HasAltQoSReferences() bool`

HasAltQoSReferences returns a boolean if a field has been set.

### GetDisUeNotif

`func (o *AsSessionWithQoSSubscription) GetDisUeNotif() bool`

GetDisUeNotif returns the DisUeNotif field if non-nil, zero value otherwise.

### GetDisUeNotifOk

`func (o *AsSessionWithQoSSubscription) GetDisUeNotifOk() (*bool, bool)`

GetDisUeNotifOk returns a tuple with the DisUeNotif field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisUeNotif

`func (o *AsSessionWithQoSSubscription) SetDisUeNotif(v bool)`

SetDisUeNotif sets DisUeNotif field to given value.

### HasDisUeNotif

`func (o *AsSessionWithQoSSubscription) HasDisUeNotif() bool`

HasDisUeNotif returns a boolean if a field has been set.

### GetUeIpv4Addr

`func (o *AsSessionWithQoSSubscription) GetUeIpv4Addr() string`

GetUeIpv4Addr returns the UeIpv4Addr field if non-nil, zero value otherwise.

### GetUeIpv4AddrOk

`func (o *AsSessionWithQoSSubscription) GetUeIpv4AddrOk() (*string, bool)`

GetUeIpv4AddrOk returns a tuple with the UeIpv4Addr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUeIpv4Addr

`func (o *AsSessionWithQoSSubscription) SetUeIpv4Addr(v string)`

SetUeIpv4Addr sets UeIpv4Addr field to given value.

### HasUeIpv4Addr

`func (o *AsSessionWithQoSSubscription) HasUeIpv4Addr() bool`

HasUeIpv4Addr returns a boolean if a field has been set.

### GetIpDomain

`func (o *AsSessionWithQoSSubscription) GetIpDomain() string`

GetIpDomain returns the IpDomain field if non-nil, zero value otherwise.

### GetIpDomainOk

`func (o *AsSessionWithQoSSubscription) GetIpDomainOk() (*string, bool)`

GetIpDomainOk returns a tuple with the IpDomain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpDomain

`func (o *AsSessionWithQoSSubscription) SetIpDomain(v string)`

SetIpDomain sets IpDomain field to given value.

### HasIpDomain

`func (o *AsSessionWithQoSSubscription) HasIpDomain() bool`

HasIpDomain returns a boolean if a field has been set.

### GetUeIpv6Addr

`func (o *AsSessionWithQoSSubscription) GetUeIpv6Addr() string`

GetUeIpv6Addr returns the UeIpv6Addr field if non-nil, zero value otherwise.

### GetUeIpv6AddrOk

`func (o *AsSessionWithQoSSubscription) GetUeIpv6AddrOk() (*string, bool)`

GetUeIpv6AddrOk returns a tuple with the UeIpv6Addr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUeIpv6Addr

`func (o *AsSessionWithQoSSubscription) SetUeIpv6Addr(v string)`

SetUeIpv6Addr sets UeIpv6Addr field to given value.

### HasUeIpv6Addr

`func (o *AsSessionWithQoSSubscription) HasUeIpv6Addr() bool`

HasUeIpv6Addr returns a boolean if a field has been set.

### GetMacAddr

`func (o *AsSessionWithQoSSubscription) GetMacAddr() string`

GetMacAddr returns the MacAddr field if non-nil, zero value otherwise.

### GetMacAddrOk

`func (o *AsSessionWithQoSSubscription) GetMacAddrOk() (*string, bool)`

GetMacAddrOk returns a tuple with the MacAddr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMacAddr

`func (o *AsSessionWithQoSSubscription) SetMacAddr(v string)`

SetMacAddr sets MacAddr field to given value.

### HasMacAddr

`func (o *AsSessionWithQoSSubscription) HasMacAddr() bool`

HasMacAddr returns a boolean if a field has been set.

### GetUsageThreshold

`func (o *AsSessionWithQoSSubscription) GetUsageThreshold() UsageThreshold`

GetUsageThreshold returns the UsageThreshold field if non-nil, zero value otherwise.

### GetUsageThresholdOk

`func (o *AsSessionWithQoSSubscription) GetUsageThresholdOk() (*UsageThreshold, bool)`

GetUsageThresholdOk returns a tuple with the UsageThreshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsageThreshold

`func (o *AsSessionWithQoSSubscription) SetUsageThreshold(v UsageThreshold)`

SetUsageThreshold sets UsageThreshold field to given value.

### HasUsageThreshold

`func (o *AsSessionWithQoSSubscription) HasUsageThreshold() bool`

HasUsageThreshold returns a boolean if a field has been set.

### GetSponsorInfo

`func (o *AsSessionWithQoSSubscription) GetSponsorInfo() SponsorInformation`

GetSponsorInfo returns the SponsorInfo field if non-nil, zero value otherwise.

### GetSponsorInfoOk

`func (o *AsSessionWithQoSSubscription) GetSponsorInfoOk() (*SponsorInformation, bool)`

GetSponsorInfoOk returns a tuple with the SponsorInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSponsorInfo

`func (o *AsSessionWithQoSSubscription) SetSponsorInfo(v SponsorInformation)`

SetSponsorInfo sets SponsorInfo field to given value.

### HasSponsorInfo

`func (o *AsSessionWithQoSSubscription) HasSponsorInfo() bool`

HasSponsorInfo returns a boolean if a field has been set.

### GetQosMonInfo

`func (o *AsSessionWithQoSSubscription) GetQosMonInfo() QosMonitoringInformation`

GetQosMonInfo returns the QosMonInfo field if non-nil, zero value otherwise.

### GetQosMonInfoOk

`func (o *AsSessionWithQoSSubscription) GetQosMonInfoOk() (*QosMonitoringInformation, bool)`

GetQosMonInfoOk returns a tuple with the QosMonInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQosMonInfo

`func (o *AsSessionWithQoSSubscription) SetQosMonInfo(v QosMonitoringInformation)`

SetQosMonInfo sets QosMonInfo field to given value.

### HasQosMonInfo

`func (o *AsSessionWithQoSSubscription) HasQosMonInfo() bool`

HasQosMonInfo returns a boolean if a field has been set.

### GetRequestTestNotification

`func (o *AsSessionWithQoSSubscription) GetRequestTestNotification() bool`

GetRequestTestNotification returns the RequestTestNotification field if non-nil, zero value otherwise.

### GetRequestTestNotificationOk

`func (o *AsSessionWithQoSSubscription) GetRequestTestNotificationOk() (*bool, bool)`

GetRequestTestNotificationOk returns a tuple with the RequestTestNotification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestTestNotification

`func (o *AsSessionWithQoSSubscription) SetRequestTestNotification(v bool)`

SetRequestTestNotification sets RequestTestNotification field to given value.

### HasRequestTestNotification

`func (o *AsSessionWithQoSSubscription) HasRequestTestNotification() bool`

HasRequestTestNotification returns a boolean if a field has been set.

### GetWebsockNotifConfig

`func (o *AsSessionWithQoSSubscription) GetWebsockNotifConfig() WebsockNotifConfig`

GetWebsockNotifConfig returns the WebsockNotifConfig field if non-nil, zero value otherwise.

### GetWebsockNotifConfigOk

`func (o *AsSessionWithQoSSubscription) GetWebsockNotifConfigOk() (*WebsockNotifConfig, bool)`

GetWebsockNotifConfigOk returns a tuple with the WebsockNotifConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebsockNotifConfig

`func (o *AsSessionWithQoSSubscription) SetWebsockNotifConfig(v WebsockNotifConfig)`

SetWebsockNotifConfig sets WebsockNotifConfig field to given value.

### HasWebsockNotifConfig

`func (o *AsSessionWithQoSSubscription) HasWebsockNotifConfig() bool`

HasWebsockNotifConfig returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


