# MonitoringSettingsServiceNow

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enabled** | Pointer to **bool** | Enabled | [optional] 
**Integration** | Pointer to [**MonitoringSettingsServiceNowIntegration**](monitoringSettings_serviceNow_integration.md) |  | [optional] 
**NewIncidentAction** | Pointer to **NullableString** | New Incident Action | [optional] 
**CloseIncidentAction** | Pointer to **NullableString** | Close Incident Action | [optional] 
**InfoMapping** | Pointer to **NullableString** | Info Mapping | [optional] 
**WarningMapping** | Pointer to **NullableString** | Warning Mapping | [optional] 
**CriticalMapping** | Pointer to **NullableString** | Critical Mapping | [optional] 

## Methods

### NewMonitoringSettingsServiceNow

`func NewMonitoringSettingsServiceNow() *MonitoringSettingsServiceNow`

NewMonitoringSettingsServiceNow instantiates a new MonitoringSettingsServiceNow object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMonitoringSettingsServiceNowWithDefaults

`func NewMonitoringSettingsServiceNowWithDefaults() *MonitoringSettingsServiceNow`

NewMonitoringSettingsServiceNowWithDefaults instantiates a new MonitoringSettingsServiceNow object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnabled

`func (o *MonitoringSettingsServiceNow) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *MonitoringSettingsServiceNow) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *MonitoringSettingsServiceNow) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *MonitoringSettingsServiceNow) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetIntegration

`func (o *MonitoringSettingsServiceNow) GetIntegration() MonitoringSettingsServiceNowIntegration`

GetIntegration returns the Integration field if non-nil, zero value otherwise.

### GetIntegrationOk

`func (o *MonitoringSettingsServiceNow) GetIntegrationOk() (*MonitoringSettingsServiceNowIntegration, bool)`

GetIntegrationOk returns a tuple with the Integration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntegration

`func (o *MonitoringSettingsServiceNow) SetIntegration(v MonitoringSettingsServiceNowIntegration)`

SetIntegration sets Integration field to given value.

### HasIntegration

`func (o *MonitoringSettingsServiceNow) HasIntegration() bool`

HasIntegration returns a boolean if a field has been set.

### GetNewIncidentAction

`func (o *MonitoringSettingsServiceNow) GetNewIncidentAction() string`

GetNewIncidentAction returns the NewIncidentAction field if non-nil, zero value otherwise.

### GetNewIncidentActionOk

`func (o *MonitoringSettingsServiceNow) GetNewIncidentActionOk() (*string, bool)`

GetNewIncidentActionOk returns a tuple with the NewIncidentAction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewIncidentAction

`func (o *MonitoringSettingsServiceNow) SetNewIncidentAction(v string)`

SetNewIncidentAction sets NewIncidentAction field to given value.

### HasNewIncidentAction

`func (o *MonitoringSettingsServiceNow) HasNewIncidentAction() bool`

HasNewIncidentAction returns a boolean if a field has been set.

### SetNewIncidentActionNil

`func (o *MonitoringSettingsServiceNow) SetNewIncidentActionNil(b bool)`

 SetNewIncidentActionNil sets the value for NewIncidentAction to be an explicit nil

### UnsetNewIncidentAction
`func (o *MonitoringSettingsServiceNow) UnsetNewIncidentAction()`

UnsetNewIncidentAction ensures that no value is present for NewIncidentAction, not even an explicit nil
### GetCloseIncidentAction

`func (o *MonitoringSettingsServiceNow) GetCloseIncidentAction() string`

GetCloseIncidentAction returns the CloseIncidentAction field if non-nil, zero value otherwise.

### GetCloseIncidentActionOk

`func (o *MonitoringSettingsServiceNow) GetCloseIncidentActionOk() (*string, bool)`

GetCloseIncidentActionOk returns a tuple with the CloseIncidentAction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloseIncidentAction

`func (o *MonitoringSettingsServiceNow) SetCloseIncidentAction(v string)`

SetCloseIncidentAction sets CloseIncidentAction field to given value.

### HasCloseIncidentAction

`func (o *MonitoringSettingsServiceNow) HasCloseIncidentAction() bool`

HasCloseIncidentAction returns a boolean if a field has been set.

### SetCloseIncidentActionNil

`func (o *MonitoringSettingsServiceNow) SetCloseIncidentActionNil(b bool)`

 SetCloseIncidentActionNil sets the value for CloseIncidentAction to be an explicit nil

### UnsetCloseIncidentAction
`func (o *MonitoringSettingsServiceNow) UnsetCloseIncidentAction()`

UnsetCloseIncidentAction ensures that no value is present for CloseIncidentAction, not even an explicit nil
### GetInfoMapping

`func (o *MonitoringSettingsServiceNow) GetInfoMapping() string`

GetInfoMapping returns the InfoMapping field if non-nil, zero value otherwise.

### GetInfoMappingOk

`func (o *MonitoringSettingsServiceNow) GetInfoMappingOk() (*string, bool)`

GetInfoMappingOk returns a tuple with the InfoMapping field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfoMapping

`func (o *MonitoringSettingsServiceNow) SetInfoMapping(v string)`

SetInfoMapping sets InfoMapping field to given value.

### HasInfoMapping

`func (o *MonitoringSettingsServiceNow) HasInfoMapping() bool`

HasInfoMapping returns a boolean if a field has been set.

### SetInfoMappingNil

`func (o *MonitoringSettingsServiceNow) SetInfoMappingNil(b bool)`

 SetInfoMappingNil sets the value for InfoMapping to be an explicit nil

### UnsetInfoMapping
`func (o *MonitoringSettingsServiceNow) UnsetInfoMapping()`

UnsetInfoMapping ensures that no value is present for InfoMapping, not even an explicit nil
### GetWarningMapping

`func (o *MonitoringSettingsServiceNow) GetWarningMapping() string`

GetWarningMapping returns the WarningMapping field if non-nil, zero value otherwise.

### GetWarningMappingOk

`func (o *MonitoringSettingsServiceNow) GetWarningMappingOk() (*string, bool)`

GetWarningMappingOk returns a tuple with the WarningMapping field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarningMapping

`func (o *MonitoringSettingsServiceNow) SetWarningMapping(v string)`

SetWarningMapping sets WarningMapping field to given value.

### HasWarningMapping

`func (o *MonitoringSettingsServiceNow) HasWarningMapping() bool`

HasWarningMapping returns a boolean if a field has been set.

### SetWarningMappingNil

`func (o *MonitoringSettingsServiceNow) SetWarningMappingNil(b bool)`

 SetWarningMappingNil sets the value for WarningMapping to be an explicit nil

### UnsetWarningMapping
`func (o *MonitoringSettingsServiceNow) UnsetWarningMapping()`

UnsetWarningMapping ensures that no value is present for WarningMapping, not even an explicit nil
### GetCriticalMapping

`func (o *MonitoringSettingsServiceNow) GetCriticalMapping() string`

GetCriticalMapping returns the CriticalMapping field if non-nil, zero value otherwise.

### GetCriticalMappingOk

`func (o *MonitoringSettingsServiceNow) GetCriticalMappingOk() (*string, bool)`

GetCriticalMappingOk returns a tuple with the CriticalMapping field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCriticalMapping

`func (o *MonitoringSettingsServiceNow) SetCriticalMapping(v string)`

SetCriticalMapping sets CriticalMapping field to given value.

### HasCriticalMapping

`func (o *MonitoringSettingsServiceNow) HasCriticalMapping() bool`

HasCriticalMapping returns a boolean if a field has been set.

### SetCriticalMappingNil

`func (o *MonitoringSettingsServiceNow) SetCriticalMappingNil(b bool)`

 SetCriticalMappingNil sets the value for CriticalMapping to be an explicit nil

### UnsetCriticalMapping
`func (o *MonitoringSettingsServiceNow) UnsetCriticalMapping()`

UnsetCriticalMapping ensures that no value is present for CriticalMapping, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


