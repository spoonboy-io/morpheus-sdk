# ApiMonitoringIncidentsIncident

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Resolution** | Pointer to **string** | Description of the resolution to this incident | [optional] 
**Comment** | Pointer to **string** | Comment on this incident, updates summary field | [optional] 
**Status** | Pointer to **string** | Set status | [optional] 
**Severity** | Pointer to **string** | Set severity | [optional] 
**Name** | **string** | Set display name | 
**StartDate** | Pointer to **time.Time** | Set start time | [optional] 
**EndDate** | Pointer to **time.Time** | Set start time | [optional] 
**InUptime** | Pointer to **bool** | Set &#39;In Availability&#39; | [optional] 

## Methods

### NewApiMonitoringIncidentsIncident

`func NewApiMonitoringIncidentsIncident(name string, ) *ApiMonitoringIncidentsIncident`

NewApiMonitoringIncidentsIncident instantiates a new ApiMonitoringIncidentsIncident object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiMonitoringIncidentsIncidentWithDefaults

`func NewApiMonitoringIncidentsIncidentWithDefaults() *ApiMonitoringIncidentsIncident`

NewApiMonitoringIncidentsIncidentWithDefaults instantiates a new ApiMonitoringIncidentsIncident object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResolution

`func (o *ApiMonitoringIncidentsIncident) GetResolution() string`

GetResolution returns the Resolution field if non-nil, zero value otherwise.

### GetResolutionOk

`func (o *ApiMonitoringIncidentsIncident) GetResolutionOk() (*string, bool)`

GetResolutionOk returns a tuple with the Resolution field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResolution

`func (o *ApiMonitoringIncidentsIncident) SetResolution(v string)`

SetResolution sets Resolution field to given value.

### HasResolution

`func (o *ApiMonitoringIncidentsIncident) HasResolution() bool`

HasResolution returns a boolean if a field has been set.

### GetComment

`func (o *ApiMonitoringIncidentsIncident) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *ApiMonitoringIncidentsIncident) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *ApiMonitoringIncidentsIncident) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *ApiMonitoringIncidentsIncident) HasComment() bool`

HasComment returns a boolean if a field has been set.

### GetStatus

`func (o *ApiMonitoringIncidentsIncident) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ApiMonitoringIncidentsIncident) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ApiMonitoringIncidentsIncident) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ApiMonitoringIncidentsIncident) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetSeverity

`func (o *ApiMonitoringIncidentsIncident) GetSeverity() string`

GetSeverity returns the Severity field if non-nil, zero value otherwise.

### GetSeverityOk

`func (o *ApiMonitoringIncidentsIncident) GetSeverityOk() (*string, bool)`

GetSeverityOk returns a tuple with the Severity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeverity

`func (o *ApiMonitoringIncidentsIncident) SetSeverity(v string)`

SetSeverity sets Severity field to given value.

### HasSeverity

`func (o *ApiMonitoringIncidentsIncident) HasSeverity() bool`

HasSeverity returns a boolean if a field has been set.

### GetName

`func (o *ApiMonitoringIncidentsIncident) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ApiMonitoringIncidentsIncident) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ApiMonitoringIncidentsIncident) SetName(v string)`

SetName sets Name field to given value.


### GetStartDate

`func (o *ApiMonitoringIncidentsIncident) GetStartDate() time.Time`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *ApiMonitoringIncidentsIncident) GetStartDateOk() (*time.Time, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *ApiMonitoringIncidentsIncident) SetStartDate(v time.Time)`

SetStartDate sets StartDate field to given value.

### HasStartDate

`func (o *ApiMonitoringIncidentsIncident) HasStartDate() bool`

HasStartDate returns a boolean if a field has been set.

### GetEndDate

`func (o *ApiMonitoringIncidentsIncident) GetEndDate() time.Time`

GetEndDate returns the EndDate field if non-nil, zero value otherwise.

### GetEndDateOk

`func (o *ApiMonitoringIncidentsIncident) GetEndDateOk() (*time.Time, bool)`

GetEndDateOk returns a tuple with the EndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndDate

`func (o *ApiMonitoringIncidentsIncident) SetEndDate(v time.Time)`

SetEndDate sets EndDate field to given value.

### HasEndDate

`func (o *ApiMonitoringIncidentsIncident) HasEndDate() bool`

HasEndDate returns a boolean if a field has been set.

### GetInUptime

`func (o *ApiMonitoringIncidentsIncident) GetInUptime() bool`

GetInUptime returns the InUptime field if non-nil, zero value otherwise.

### GetInUptimeOk

`func (o *ApiMonitoringIncidentsIncident) GetInUptimeOk() (*bool, bool)`

GetInUptimeOk returns a tuple with the InUptime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInUptime

`func (o *ApiMonitoringIncidentsIncident) SetInUptime(v bool)`

SetInUptime sets InUptime field to given value.

### HasInUptime

`func (o *ApiMonitoringIncidentsIncident) HasInUptime() bool`

HasInUptime returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


