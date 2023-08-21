# InlineResponse20080LoadBalancerPool

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** |  | [optional] 
**LoadBalancer** | Pointer to [**InlineResponse20079LoadBalancerMonitorLoadBalancer**](inline_response_200_79_loadBalancerMonitor_loadBalancer.md) |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Category** | Pointer to **NullableString** |  | [optional] 
**Visibility** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **NullableString** |  | [optional] 
**InternalId** | Pointer to **string** |  | [optional] 
**ExternalId** | Pointer to **string** |  | [optional] 
**Enabled** | Pointer to **bool** |  | [optional] 
**VipSticky** | Pointer to **NullableString** |  | [optional] 
**VipBalance** | Pointer to **string** |  | [optional] 
**AllowNat** | Pointer to **NullableString** |  | [optional] 
**AllowSnat** | Pointer to **NullableString** |  | [optional] 
**VipClientIpMode** | Pointer to **NullableString** |  | [optional] 
**VipServerIpMode** | Pointer to **NullableString** |  | [optional] 
**MinActive** | Pointer to **int64** |  | [optional] 
**MinInService** | Pointer to **NullableString** |  | [optional] 
**MinUpMonitor** | Pointer to **NullableString** |  | [optional] 
**MinUpAction** | Pointer to **NullableString** |  | [optional] 
**MaxQueueDepth** | Pointer to **NullableString** |  | [optional] 
**MaxQueueTime** | Pointer to **NullableString** |  | [optional] 
**NumberActive** | Pointer to **int64** |  | [optional] 
**NumberInService** | Pointer to **int64** |  | [optional] 
**HealthScore** | Pointer to **int64** |  | [optional] 
**PerformanceScore** | Pointer to **int64** |  | [optional] 
**HealthPenalty** | Pointer to **int64** |  | [optional] 
**SecurityPenalty** | Pointer to **int64** |  | [optional] 
**ErrorPenalty** | Pointer to **int64** |  | [optional] 
**DownAction** | Pointer to **NullableString** |  | [optional] 
**RampTime** | Pointer to **NullableString** |  | [optional] 
**Port** | Pointer to **NullableString** |  | [optional] 
**PortType** | Pointer to **NullableString** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**Nodes** | Pointer to [**[]InlineResponse20040AppDeployInstance**](InlineResponse20040AppDeployInstance.md) |  | [optional] 
**Monitors** | Pointer to [**[]InlineResponse20040AppDeployInstance**](InlineResponse20040AppDeployInstance.md) |  | [optional] 
**Members** | Pointer to **[]map[string]interface{}** |  | [optional] 
**Config** | Pointer to **map[string]interface{}** |  | [optional] 
**CreatedBy** | Pointer to **NullableString** |  | [optional] 
**DateCreated** | Pointer to **time.Time** |  | [optional] 
**LastUpdated** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewInlineResponse20080LoadBalancerPool

`func NewInlineResponse20080LoadBalancerPool() *InlineResponse20080LoadBalancerPool`

NewInlineResponse20080LoadBalancerPool instantiates a new InlineResponse20080LoadBalancerPool object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse20080LoadBalancerPoolWithDefaults

`func NewInlineResponse20080LoadBalancerPoolWithDefaults() *InlineResponse20080LoadBalancerPool`

NewInlineResponse20080LoadBalancerPoolWithDefaults instantiates a new InlineResponse20080LoadBalancerPool object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *InlineResponse20080LoadBalancerPool) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *InlineResponse20080LoadBalancerPool) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *InlineResponse20080LoadBalancerPool) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *InlineResponse20080LoadBalancerPool) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLoadBalancer

`func (o *InlineResponse20080LoadBalancerPool) GetLoadBalancer() InlineResponse20079LoadBalancerMonitorLoadBalancer`

GetLoadBalancer returns the LoadBalancer field if non-nil, zero value otherwise.

### GetLoadBalancerOk

`func (o *InlineResponse20080LoadBalancerPool) GetLoadBalancerOk() (*InlineResponse20079LoadBalancerMonitorLoadBalancer, bool)`

GetLoadBalancerOk returns a tuple with the LoadBalancer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoadBalancer

`func (o *InlineResponse20080LoadBalancerPool) SetLoadBalancer(v InlineResponse20079LoadBalancerMonitorLoadBalancer)`

SetLoadBalancer sets LoadBalancer field to given value.

### HasLoadBalancer

`func (o *InlineResponse20080LoadBalancerPool) HasLoadBalancer() bool`

HasLoadBalancer returns a boolean if a field has been set.

### GetName

`func (o *InlineResponse20080LoadBalancerPool) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *InlineResponse20080LoadBalancerPool) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *InlineResponse20080LoadBalancerPool) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *InlineResponse20080LoadBalancerPool) HasName() bool`

HasName returns a boolean if a field has been set.

### GetCategory

`func (o *InlineResponse20080LoadBalancerPool) GetCategory() string`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *InlineResponse20080LoadBalancerPool) GetCategoryOk() (*string, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *InlineResponse20080LoadBalancerPool) SetCategory(v string)`

SetCategory sets Category field to given value.

### HasCategory

`func (o *InlineResponse20080LoadBalancerPool) HasCategory() bool`

HasCategory returns a boolean if a field has been set.

### SetCategoryNil

`func (o *InlineResponse20080LoadBalancerPool) SetCategoryNil(b bool)`

 SetCategoryNil sets the value for Category to be an explicit nil

### UnsetCategory
`func (o *InlineResponse20080LoadBalancerPool) UnsetCategory()`

UnsetCategory ensures that no value is present for Category, not even an explicit nil
### GetVisibility

`func (o *InlineResponse20080LoadBalancerPool) GetVisibility() string`

GetVisibility returns the Visibility field if non-nil, zero value otherwise.

### GetVisibilityOk

`func (o *InlineResponse20080LoadBalancerPool) GetVisibilityOk() (*string, bool)`

GetVisibilityOk returns a tuple with the Visibility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisibility

`func (o *InlineResponse20080LoadBalancerPool) SetVisibility(v string)`

SetVisibility sets Visibility field to given value.

### HasVisibility

`func (o *InlineResponse20080LoadBalancerPool) HasVisibility() bool`

HasVisibility returns a boolean if a field has been set.

### GetDescription

`func (o *InlineResponse20080LoadBalancerPool) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *InlineResponse20080LoadBalancerPool) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *InlineResponse20080LoadBalancerPool) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *InlineResponse20080LoadBalancerPool) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *InlineResponse20080LoadBalancerPool) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *InlineResponse20080LoadBalancerPool) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetInternalId

`func (o *InlineResponse20080LoadBalancerPool) GetInternalId() string`

GetInternalId returns the InternalId field if non-nil, zero value otherwise.

### GetInternalIdOk

`func (o *InlineResponse20080LoadBalancerPool) GetInternalIdOk() (*string, bool)`

GetInternalIdOk returns a tuple with the InternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInternalId

`func (o *InlineResponse20080LoadBalancerPool) SetInternalId(v string)`

SetInternalId sets InternalId field to given value.

### HasInternalId

`func (o *InlineResponse20080LoadBalancerPool) HasInternalId() bool`

HasInternalId returns a boolean if a field has been set.

### GetExternalId

`func (o *InlineResponse20080LoadBalancerPool) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *InlineResponse20080LoadBalancerPool) GetExternalIdOk() (*string, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *InlineResponse20080LoadBalancerPool) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.

### HasExternalId

`func (o *InlineResponse20080LoadBalancerPool) HasExternalId() bool`

HasExternalId returns a boolean if a field has been set.

### GetEnabled

`func (o *InlineResponse20080LoadBalancerPool) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *InlineResponse20080LoadBalancerPool) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *InlineResponse20080LoadBalancerPool) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *InlineResponse20080LoadBalancerPool) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetVipSticky

`func (o *InlineResponse20080LoadBalancerPool) GetVipSticky() string`

GetVipSticky returns the VipSticky field if non-nil, zero value otherwise.

### GetVipStickyOk

`func (o *InlineResponse20080LoadBalancerPool) GetVipStickyOk() (*string, bool)`

GetVipStickyOk returns a tuple with the VipSticky field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVipSticky

`func (o *InlineResponse20080LoadBalancerPool) SetVipSticky(v string)`

SetVipSticky sets VipSticky field to given value.

### HasVipSticky

`func (o *InlineResponse20080LoadBalancerPool) HasVipSticky() bool`

HasVipSticky returns a boolean if a field has been set.

### SetVipStickyNil

`func (o *InlineResponse20080LoadBalancerPool) SetVipStickyNil(b bool)`

 SetVipStickyNil sets the value for VipSticky to be an explicit nil

### UnsetVipSticky
`func (o *InlineResponse20080LoadBalancerPool) UnsetVipSticky()`

UnsetVipSticky ensures that no value is present for VipSticky, not even an explicit nil
### GetVipBalance

`func (o *InlineResponse20080LoadBalancerPool) GetVipBalance() string`

GetVipBalance returns the VipBalance field if non-nil, zero value otherwise.

### GetVipBalanceOk

`func (o *InlineResponse20080LoadBalancerPool) GetVipBalanceOk() (*string, bool)`

GetVipBalanceOk returns a tuple with the VipBalance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVipBalance

`func (o *InlineResponse20080LoadBalancerPool) SetVipBalance(v string)`

SetVipBalance sets VipBalance field to given value.

### HasVipBalance

`func (o *InlineResponse20080LoadBalancerPool) HasVipBalance() bool`

HasVipBalance returns a boolean if a field has been set.

### GetAllowNat

`func (o *InlineResponse20080LoadBalancerPool) GetAllowNat() string`

GetAllowNat returns the AllowNat field if non-nil, zero value otherwise.

### GetAllowNatOk

`func (o *InlineResponse20080LoadBalancerPool) GetAllowNatOk() (*string, bool)`

GetAllowNatOk returns a tuple with the AllowNat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowNat

`func (o *InlineResponse20080LoadBalancerPool) SetAllowNat(v string)`

SetAllowNat sets AllowNat field to given value.

### HasAllowNat

`func (o *InlineResponse20080LoadBalancerPool) HasAllowNat() bool`

HasAllowNat returns a boolean if a field has been set.

### SetAllowNatNil

`func (o *InlineResponse20080LoadBalancerPool) SetAllowNatNil(b bool)`

 SetAllowNatNil sets the value for AllowNat to be an explicit nil

### UnsetAllowNat
`func (o *InlineResponse20080LoadBalancerPool) UnsetAllowNat()`

UnsetAllowNat ensures that no value is present for AllowNat, not even an explicit nil
### GetAllowSnat

`func (o *InlineResponse20080LoadBalancerPool) GetAllowSnat() string`

GetAllowSnat returns the AllowSnat field if non-nil, zero value otherwise.

### GetAllowSnatOk

`func (o *InlineResponse20080LoadBalancerPool) GetAllowSnatOk() (*string, bool)`

GetAllowSnatOk returns a tuple with the AllowSnat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowSnat

`func (o *InlineResponse20080LoadBalancerPool) SetAllowSnat(v string)`

SetAllowSnat sets AllowSnat field to given value.

### HasAllowSnat

`func (o *InlineResponse20080LoadBalancerPool) HasAllowSnat() bool`

HasAllowSnat returns a boolean if a field has been set.

### SetAllowSnatNil

`func (o *InlineResponse20080LoadBalancerPool) SetAllowSnatNil(b bool)`

 SetAllowSnatNil sets the value for AllowSnat to be an explicit nil

### UnsetAllowSnat
`func (o *InlineResponse20080LoadBalancerPool) UnsetAllowSnat()`

UnsetAllowSnat ensures that no value is present for AllowSnat, not even an explicit nil
### GetVipClientIpMode

`func (o *InlineResponse20080LoadBalancerPool) GetVipClientIpMode() string`

GetVipClientIpMode returns the VipClientIpMode field if non-nil, zero value otherwise.

### GetVipClientIpModeOk

`func (o *InlineResponse20080LoadBalancerPool) GetVipClientIpModeOk() (*string, bool)`

GetVipClientIpModeOk returns a tuple with the VipClientIpMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVipClientIpMode

`func (o *InlineResponse20080LoadBalancerPool) SetVipClientIpMode(v string)`

SetVipClientIpMode sets VipClientIpMode field to given value.

### HasVipClientIpMode

`func (o *InlineResponse20080LoadBalancerPool) HasVipClientIpMode() bool`

HasVipClientIpMode returns a boolean if a field has been set.

### SetVipClientIpModeNil

`func (o *InlineResponse20080LoadBalancerPool) SetVipClientIpModeNil(b bool)`

 SetVipClientIpModeNil sets the value for VipClientIpMode to be an explicit nil

### UnsetVipClientIpMode
`func (o *InlineResponse20080LoadBalancerPool) UnsetVipClientIpMode()`

UnsetVipClientIpMode ensures that no value is present for VipClientIpMode, not even an explicit nil
### GetVipServerIpMode

`func (o *InlineResponse20080LoadBalancerPool) GetVipServerIpMode() string`

GetVipServerIpMode returns the VipServerIpMode field if non-nil, zero value otherwise.

### GetVipServerIpModeOk

`func (o *InlineResponse20080LoadBalancerPool) GetVipServerIpModeOk() (*string, bool)`

GetVipServerIpModeOk returns a tuple with the VipServerIpMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVipServerIpMode

`func (o *InlineResponse20080LoadBalancerPool) SetVipServerIpMode(v string)`

SetVipServerIpMode sets VipServerIpMode field to given value.

### HasVipServerIpMode

`func (o *InlineResponse20080LoadBalancerPool) HasVipServerIpMode() bool`

HasVipServerIpMode returns a boolean if a field has been set.

### SetVipServerIpModeNil

`func (o *InlineResponse20080LoadBalancerPool) SetVipServerIpModeNil(b bool)`

 SetVipServerIpModeNil sets the value for VipServerIpMode to be an explicit nil

### UnsetVipServerIpMode
`func (o *InlineResponse20080LoadBalancerPool) UnsetVipServerIpMode()`

UnsetVipServerIpMode ensures that no value is present for VipServerIpMode, not even an explicit nil
### GetMinActive

`func (o *InlineResponse20080LoadBalancerPool) GetMinActive() int64`

GetMinActive returns the MinActive field if non-nil, zero value otherwise.

### GetMinActiveOk

`func (o *InlineResponse20080LoadBalancerPool) GetMinActiveOk() (*int64, bool)`

GetMinActiveOk returns a tuple with the MinActive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinActive

`func (o *InlineResponse20080LoadBalancerPool) SetMinActive(v int64)`

SetMinActive sets MinActive field to given value.

### HasMinActive

`func (o *InlineResponse20080LoadBalancerPool) HasMinActive() bool`

HasMinActive returns a boolean if a field has been set.

### GetMinInService

`func (o *InlineResponse20080LoadBalancerPool) GetMinInService() string`

GetMinInService returns the MinInService field if non-nil, zero value otherwise.

### GetMinInServiceOk

`func (o *InlineResponse20080LoadBalancerPool) GetMinInServiceOk() (*string, bool)`

GetMinInServiceOk returns a tuple with the MinInService field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinInService

`func (o *InlineResponse20080LoadBalancerPool) SetMinInService(v string)`

SetMinInService sets MinInService field to given value.

### HasMinInService

`func (o *InlineResponse20080LoadBalancerPool) HasMinInService() bool`

HasMinInService returns a boolean if a field has been set.

### SetMinInServiceNil

`func (o *InlineResponse20080LoadBalancerPool) SetMinInServiceNil(b bool)`

 SetMinInServiceNil sets the value for MinInService to be an explicit nil

### UnsetMinInService
`func (o *InlineResponse20080LoadBalancerPool) UnsetMinInService()`

UnsetMinInService ensures that no value is present for MinInService, not even an explicit nil
### GetMinUpMonitor

`func (o *InlineResponse20080LoadBalancerPool) GetMinUpMonitor() string`

GetMinUpMonitor returns the MinUpMonitor field if non-nil, zero value otherwise.

### GetMinUpMonitorOk

`func (o *InlineResponse20080LoadBalancerPool) GetMinUpMonitorOk() (*string, bool)`

GetMinUpMonitorOk returns a tuple with the MinUpMonitor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinUpMonitor

`func (o *InlineResponse20080LoadBalancerPool) SetMinUpMonitor(v string)`

SetMinUpMonitor sets MinUpMonitor field to given value.

### HasMinUpMonitor

`func (o *InlineResponse20080LoadBalancerPool) HasMinUpMonitor() bool`

HasMinUpMonitor returns a boolean if a field has been set.

### SetMinUpMonitorNil

`func (o *InlineResponse20080LoadBalancerPool) SetMinUpMonitorNil(b bool)`

 SetMinUpMonitorNil sets the value for MinUpMonitor to be an explicit nil

### UnsetMinUpMonitor
`func (o *InlineResponse20080LoadBalancerPool) UnsetMinUpMonitor()`

UnsetMinUpMonitor ensures that no value is present for MinUpMonitor, not even an explicit nil
### GetMinUpAction

`func (o *InlineResponse20080LoadBalancerPool) GetMinUpAction() string`

GetMinUpAction returns the MinUpAction field if non-nil, zero value otherwise.

### GetMinUpActionOk

`func (o *InlineResponse20080LoadBalancerPool) GetMinUpActionOk() (*string, bool)`

GetMinUpActionOk returns a tuple with the MinUpAction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinUpAction

`func (o *InlineResponse20080LoadBalancerPool) SetMinUpAction(v string)`

SetMinUpAction sets MinUpAction field to given value.

### HasMinUpAction

`func (o *InlineResponse20080LoadBalancerPool) HasMinUpAction() bool`

HasMinUpAction returns a boolean if a field has been set.

### SetMinUpActionNil

`func (o *InlineResponse20080LoadBalancerPool) SetMinUpActionNil(b bool)`

 SetMinUpActionNil sets the value for MinUpAction to be an explicit nil

### UnsetMinUpAction
`func (o *InlineResponse20080LoadBalancerPool) UnsetMinUpAction()`

UnsetMinUpAction ensures that no value is present for MinUpAction, not even an explicit nil
### GetMaxQueueDepth

`func (o *InlineResponse20080LoadBalancerPool) GetMaxQueueDepth() string`

GetMaxQueueDepth returns the MaxQueueDepth field if non-nil, zero value otherwise.

### GetMaxQueueDepthOk

`func (o *InlineResponse20080LoadBalancerPool) GetMaxQueueDepthOk() (*string, bool)`

GetMaxQueueDepthOk returns a tuple with the MaxQueueDepth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxQueueDepth

`func (o *InlineResponse20080LoadBalancerPool) SetMaxQueueDepth(v string)`

SetMaxQueueDepth sets MaxQueueDepth field to given value.

### HasMaxQueueDepth

`func (o *InlineResponse20080LoadBalancerPool) HasMaxQueueDepth() bool`

HasMaxQueueDepth returns a boolean if a field has been set.

### SetMaxQueueDepthNil

`func (o *InlineResponse20080LoadBalancerPool) SetMaxQueueDepthNil(b bool)`

 SetMaxQueueDepthNil sets the value for MaxQueueDepth to be an explicit nil

### UnsetMaxQueueDepth
`func (o *InlineResponse20080LoadBalancerPool) UnsetMaxQueueDepth()`

UnsetMaxQueueDepth ensures that no value is present for MaxQueueDepth, not even an explicit nil
### GetMaxQueueTime

`func (o *InlineResponse20080LoadBalancerPool) GetMaxQueueTime() string`

GetMaxQueueTime returns the MaxQueueTime field if non-nil, zero value otherwise.

### GetMaxQueueTimeOk

`func (o *InlineResponse20080LoadBalancerPool) GetMaxQueueTimeOk() (*string, bool)`

GetMaxQueueTimeOk returns a tuple with the MaxQueueTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxQueueTime

`func (o *InlineResponse20080LoadBalancerPool) SetMaxQueueTime(v string)`

SetMaxQueueTime sets MaxQueueTime field to given value.

### HasMaxQueueTime

`func (o *InlineResponse20080LoadBalancerPool) HasMaxQueueTime() bool`

HasMaxQueueTime returns a boolean if a field has been set.

### SetMaxQueueTimeNil

`func (o *InlineResponse20080LoadBalancerPool) SetMaxQueueTimeNil(b bool)`

 SetMaxQueueTimeNil sets the value for MaxQueueTime to be an explicit nil

### UnsetMaxQueueTime
`func (o *InlineResponse20080LoadBalancerPool) UnsetMaxQueueTime()`

UnsetMaxQueueTime ensures that no value is present for MaxQueueTime, not even an explicit nil
### GetNumberActive

`func (o *InlineResponse20080LoadBalancerPool) GetNumberActive() int64`

GetNumberActive returns the NumberActive field if non-nil, zero value otherwise.

### GetNumberActiveOk

`func (o *InlineResponse20080LoadBalancerPool) GetNumberActiveOk() (*int64, bool)`

GetNumberActiveOk returns a tuple with the NumberActive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumberActive

`func (o *InlineResponse20080LoadBalancerPool) SetNumberActive(v int64)`

SetNumberActive sets NumberActive field to given value.

### HasNumberActive

`func (o *InlineResponse20080LoadBalancerPool) HasNumberActive() bool`

HasNumberActive returns a boolean if a field has been set.

### GetNumberInService

`func (o *InlineResponse20080LoadBalancerPool) GetNumberInService() int64`

GetNumberInService returns the NumberInService field if non-nil, zero value otherwise.

### GetNumberInServiceOk

`func (o *InlineResponse20080LoadBalancerPool) GetNumberInServiceOk() (*int64, bool)`

GetNumberInServiceOk returns a tuple with the NumberInService field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumberInService

`func (o *InlineResponse20080LoadBalancerPool) SetNumberInService(v int64)`

SetNumberInService sets NumberInService field to given value.

### HasNumberInService

`func (o *InlineResponse20080LoadBalancerPool) HasNumberInService() bool`

HasNumberInService returns a boolean if a field has been set.

### GetHealthScore

`func (o *InlineResponse20080LoadBalancerPool) GetHealthScore() int64`

GetHealthScore returns the HealthScore field if non-nil, zero value otherwise.

### GetHealthScoreOk

`func (o *InlineResponse20080LoadBalancerPool) GetHealthScoreOk() (*int64, bool)`

GetHealthScoreOk returns a tuple with the HealthScore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHealthScore

`func (o *InlineResponse20080LoadBalancerPool) SetHealthScore(v int64)`

SetHealthScore sets HealthScore field to given value.

### HasHealthScore

`func (o *InlineResponse20080LoadBalancerPool) HasHealthScore() bool`

HasHealthScore returns a boolean if a field has been set.

### GetPerformanceScore

`func (o *InlineResponse20080LoadBalancerPool) GetPerformanceScore() int64`

GetPerformanceScore returns the PerformanceScore field if non-nil, zero value otherwise.

### GetPerformanceScoreOk

`func (o *InlineResponse20080LoadBalancerPool) GetPerformanceScoreOk() (*int64, bool)`

GetPerformanceScoreOk returns a tuple with the PerformanceScore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPerformanceScore

`func (o *InlineResponse20080LoadBalancerPool) SetPerformanceScore(v int64)`

SetPerformanceScore sets PerformanceScore field to given value.

### HasPerformanceScore

`func (o *InlineResponse20080LoadBalancerPool) HasPerformanceScore() bool`

HasPerformanceScore returns a boolean if a field has been set.

### GetHealthPenalty

`func (o *InlineResponse20080LoadBalancerPool) GetHealthPenalty() int64`

GetHealthPenalty returns the HealthPenalty field if non-nil, zero value otherwise.

### GetHealthPenaltyOk

`func (o *InlineResponse20080LoadBalancerPool) GetHealthPenaltyOk() (*int64, bool)`

GetHealthPenaltyOk returns a tuple with the HealthPenalty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHealthPenalty

`func (o *InlineResponse20080LoadBalancerPool) SetHealthPenalty(v int64)`

SetHealthPenalty sets HealthPenalty field to given value.

### HasHealthPenalty

`func (o *InlineResponse20080LoadBalancerPool) HasHealthPenalty() bool`

HasHealthPenalty returns a boolean if a field has been set.

### GetSecurityPenalty

`func (o *InlineResponse20080LoadBalancerPool) GetSecurityPenalty() int64`

GetSecurityPenalty returns the SecurityPenalty field if non-nil, zero value otherwise.

### GetSecurityPenaltyOk

`func (o *InlineResponse20080LoadBalancerPool) GetSecurityPenaltyOk() (*int64, bool)`

GetSecurityPenaltyOk returns a tuple with the SecurityPenalty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurityPenalty

`func (o *InlineResponse20080LoadBalancerPool) SetSecurityPenalty(v int64)`

SetSecurityPenalty sets SecurityPenalty field to given value.

### HasSecurityPenalty

`func (o *InlineResponse20080LoadBalancerPool) HasSecurityPenalty() bool`

HasSecurityPenalty returns a boolean if a field has been set.

### GetErrorPenalty

`func (o *InlineResponse20080LoadBalancerPool) GetErrorPenalty() int64`

GetErrorPenalty returns the ErrorPenalty field if non-nil, zero value otherwise.

### GetErrorPenaltyOk

`func (o *InlineResponse20080LoadBalancerPool) GetErrorPenaltyOk() (*int64, bool)`

GetErrorPenaltyOk returns a tuple with the ErrorPenalty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorPenalty

`func (o *InlineResponse20080LoadBalancerPool) SetErrorPenalty(v int64)`

SetErrorPenalty sets ErrorPenalty field to given value.

### HasErrorPenalty

`func (o *InlineResponse20080LoadBalancerPool) HasErrorPenalty() bool`

HasErrorPenalty returns a boolean if a field has been set.

### GetDownAction

`func (o *InlineResponse20080LoadBalancerPool) GetDownAction() string`

GetDownAction returns the DownAction field if non-nil, zero value otherwise.

### GetDownActionOk

`func (o *InlineResponse20080LoadBalancerPool) GetDownActionOk() (*string, bool)`

GetDownActionOk returns a tuple with the DownAction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDownAction

`func (o *InlineResponse20080LoadBalancerPool) SetDownAction(v string)`

SetDownAction sets DownAction field to given value.

### HasDownAction

`func (o *InlineResponse20080LoadBalancerPool) HasDownAction() bool`

HasDownAction returns a boolean if a field has been set.

### SetDownActionNil

`func (o *InlineResponse20080LoadBalancerPool) SetDownActionNil(b bool)`

 SetDownActionNil sets the value for DownAction to be an explicit nil

### UnsetDownAction
`func (o *InlineResponse20080LoadBalancerPool) UnsetDownAction()`

UnsetDownAction ensures that no value is present for DownAction, not even an explicit nil
### GetRampTime

`func (o *InlineResponse20080LoadBalancerPool) GetRampTime() string`

GetRampTime returns the RampTime field if non-nil, zero value otherwise.

### GetRampTimeOk

`func (o *InlineResponse20080LoadBalancerPool) GetRampTimeOk() (*string, bool)`

GetRampTimeOk returns a tuple with the RampTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRampTime

`func (o *InlineResponse20080LoadBalancerPool) SetRampTime(v string)`

SetRampTime sets RampTime field to given value.

### HasRampTime

`func (o *InlineResponse20080LoadBalancerPool) HasRampTime() bool`

HasRampTime returns a boolean if a field has been set.

### SetRampTimeNil

`func (o *InlineResponse20080LoadBalancerPool) SetRampTimeNil(b bool)`

 SetRampTimeNil sets the value for RampTime to be an explicit nil

### UnsetRampTime
`func (o *InlineResponse20080LoadBalancerPool) UnsetRampTime()`

UnsetRampTime ensures that no value is present for RampTime, not even an explicit nil
### GetPort

`func (o *InlineResponse20080LoadBalancerPool) GetPort() string`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *InlineResponse20080LoadBalancerPool) GetPortOk() (*string, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *InlineResponse20080LoadBalancerPool) SetPort(v string)`

SetPort sets Port field to given value.

### HasPort

`func (o *InlineResponse20080LoadBalancerPool) HasPort() bool`

HasPort returns a boolean if a field has been set.

### SetPortNil

`func (o *InlineResponse20080LoadBalancerPool) SetPortNil(b bool)`

 SetPortNil sets the value for Port to be an explicit nil

### UnsetPort
`func (o *InlineResponse20080LoadBalancerPool) UnsetPort()`

UnsetPort ensures that no value is present for Port, not even an explicit nil
### GetPortType

`func (o *InlineResponse20080LoadBalancerPool) GetPortType() string`

GetPortType returns the PortType field if non-nil, zero value otherwise.

### GetPortTypeOk

`func (o *InlineResponse20080LoadBalancerPool) GetPortTypeOk() (*string, bool)`

GetPortTypeOk returns a tuple with the PortType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortType

`func (o *InlineResponse20080LoadBalancerPool) SetPortType(v string)`

SetPortType sets PortType field to given value.

### HasPortType

`func (o *InlineResponse20080LoadBalancerPool) HasPortType() bool`

HasPortType returns a boolean if a field has been set.

### SetPortTypeNil

`func (o *InlineResponse20080LoadBalancerPool) SetPortTypeNil(b bool)`

 SetPortTypeNil sets the value for PortType to be an explicit nil

### UnsetPortType
`func (o *InlineResponse20080LoadBalancerPool) UnsetPortType()`

UnsetPortType ensures that no value is present for PortType, not even an explicit nil
### GetStatus

`func (o *InlineResponse20080LoadBalancerPool) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *InlineResponse20080LoadBalancerPool) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *InlineResponse20080LoadBalancerPool) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *InlineResponse20080LoadBalancerPool) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetNodes

`func (o *InlineResponse20080LoadBalancerPool) GetNodes() []InlineResponse20040AppDeployInstance`

GetNodes returns the Nodes field if non-nil, zero value otherwise.

### GetNodesOk

`func (o *InlineResponse20080LoadBalancerPool) GetNodesOk() (*[]InlineResponse20040AppDeployInstance, bool)`

GetNodesOk returns a tuple with the Nodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodes

`func (o *InlineResponse20080LoadBalancerPool) SetNodes(v []InlineResponse20040AppDeployInstance)`

SetNodes sets Nodes field to given value.

### HasNodes

`func (o *InlineResponse20080LoadBalancerPool) HasNodes() bool`

HasNodes returns a boolean if a field has been set.

### GetMonitors

`func (o *InlineResponse20080LoadBalancerPool) GetMonitors() []InlineResponse20040AppDeployInstance`

GetMonitors returns the Monitors field if non-nil, zero value otherwise.

### GetMonitorsOk

`func (o *InlineResponse20080LoadBalancerPool) GetMonitorsOk() (*[]InlineResponse20040AppDeployInstance, bool)`

GetMonitorsOk returns a tuple with the Monitors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonitors

`func (o *InlineResponse20080LoadBalancerPool) SetMonitors(v []InlineResponse20040AppDeployInstance)`

SetMonitors sets Monitors field to given value.

### HasMonitors

`func (o *InlineResponse20080LoadBalancerPool) HasMonitors() bool`

HasMonitors returns a boolean if a field has been set.

### GetMembers

`func (o *InlineResponse20080LoadBalancerPool) GetMembers() []map[string]interface{}`

GetMembers returns the Members field if non-nil, zero value otherwise.

### GetMembersOk

`func (o *InlineResponse20080LoadBalancerPool) GetMembersOk() (*[]map[string]interface{}, bool)`

GetMembersOk returns a tuple with the Members field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMembers

`func (o *InlineResponse20080LoadBalancerPool) SetMembers(v []map[string]interface{})`

SetMembers sets Members field to given value.

### HasMembers

`func (o *InlineResponse20080LoadBalancerPool) HasMembers() bool`

HasMembers returns a boolean if a field has been set.

### GetConfig

`func (o *InlineResponse20080LoadBalancerPool) GetConfig() map[string]interface{}`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *InlineResponse20080LoadBalancerPool) GetConfigOk() (*map[string]interface{}, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *InlineResponse20080LoadBalancerPool) SetConfig(v map[string]interface{})`

SetConfig sets Config field to given value.

### HasConfig

`func (o *InlineResponse20080LoadBalancerPool) HasConfig() bool`

HasConfig returns a boolean if a field has been set.

### GetCreatedBy

`func (o *InlineResponse20080LoadBalancerPool) GetCreatedBy() string`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *InlineResponse20080LoadBalancerPool) GetCreatedByOk() (*string, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *InlineResponse20080LoadBalancerPool) SetCreatedBy(v string)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *InlineResponse20080LoadBalancerPool) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### SetCreatedByNil

`func (o *InlineResponse20080LoadBalancerPool) SetCreatedByNil(b bool)`

 SetCreatedByNil sets the value for CreatedBy to be an explicit nil

### UnsetCreatedBy
`func (o *InlineResponse20080LoadBalancerPool) UnsetCreatedBy()`

UnsetCreatedBy ensures that no value is present for CreatedBy, not even an explicit nil
### GetDateCreated

`func (o *InlineResponse20080LoadBalancerPool) GetDateCreated() time.Time`

GetDateCreated returns the DateCreated field if non-nil, zero value otherwise.

### GetDateCreatedOk

`func (o *InlineResponse20080LoadBalancerPool) GetDateCreatedOk() (*time.Time, bool)`

GetDateCreatedOk returns a tuple with the DateCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateCreated

`func (o *InlineResponse20080LoadBalancerPool) SetDateCreated(v time.Time)`

SetDateCreated sets DateCreated field to given value.

### HasDateCreated

`func (o *InlineResponse20080LoadBalancerPool) HasDateCreated() bool`

HasDateCreated returns a boolean if a field has been set.

### GetLastUpdated

`func (o *InlineResponse20080LoadBalancerPool) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *InlineResponse20080LoadBalancerPool) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *InlineResponse20080LoadBalancerPool) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.

### HasLastUpdated

`func (o *InlineResponse20080LoadBalancerPool) HasLastUpdated() bool`

HasLastUpdated returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


