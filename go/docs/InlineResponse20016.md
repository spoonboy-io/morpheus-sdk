# InlineResponse20016

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MonitorApp** | Pointer to [**CheckApp**](checkApp.md) |  | [optional] 
**CheckGroups** | Pointer to [**[]CheckGroup**](CheckGroup.md) |  | [optional] 
**Checks** | Pointer to [**[]Check**](Check.md) |  | [optional] 
**OpenIncidents** | Pointer to [**[]Incident**](Incident.md) |  | [optional] 

## Methods

### NewInlineResponse20016

`func NewInlineResponse20016() *InlineResponse20016`

NewInlineResponse20016 instantiates a new InlineResponse20016 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse20016WithDefaults

`func NewInlineResponse20016WithDefaults() *InlineResponse20016`

NewInlineResponse20016WithDefaults instantiates a new InlineResponse20016 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMonitorApp

`func (o *InlineResponse20016) GetMonitorApp() CheckApp`

GetMonitorApp returns the MonitorApp field if non-nil, zero value otherwise.

### GetMonitorAppOk

`func (o *InlineResponse20016) GetMonitorAppOk() (*CheckApp, bool)`

GetMonitorAppOk returns a tuple with the MonitorApp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonitorApp

`func (o *InlineResponse20016) SetMonitorApp(v CheckApp)`

SetMonitorApp sets MonitorApp field to given value.

### HasMonitorApp

`func (o *InlineResponse20016) HasMonitorApp() bool`

HasMonitorApp returns a boolean if a field has been set.

### GetCheckGroups

`func (o *InlineResponse20016) GetCheckGroups() []CheckGroup`

GetCheckGroups returns the CheckGroups field if non-nil, zero value otherwise.

### GetCheckGroupsOk

`func (o *InlineResponse20016) GetCheckGroupsOk() (*[]CheckGroup, bool)`

GetCheckGroupsOk returns a tuple with the CheckGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCheckGroups

`func (o *InlineResponse20016) SetCheckGroups(v []CheckGroup)`

SetCheckGroups sets CheckGroups field to given value.

### HasCheckGroups

`func (o *InlineResponse20016) HasCheckGroups() bool`

HasCheckGroups returns a boolean if a field has been set.

### GetChecks

`func (o *InlineResponse20016) GetChecks() []Check`

GetChecks returns the Checks field if non-nil, zero value otherwise.

### GetChecksOk

`func (o *InlineResponse20016) GetChecksOk() (*[]Check, bool)`

GetChecksOk returns a tuple with the Checks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChecks

`func (o *InlineResponse20016) SetChecks(v []Check)`

SetChecks sets Checks field to given value.

### HasChecks

`func (o *InlineResponse20016) HasChecks() bool`

HasChecks returns a boolean if a field has been set.

### GetOpenIncidents

`func (o *InlineResponse20016) GetOpenIncidents() []Incident`

GetOpenIncidents returns the OpenIncidents field if non-nil, zero value otherwise.

### GetOpenIncidentsOk

`func (o *InlineResponse20016) GetOpenIncidentsOk() (*[]Incident, bool)`

GetOpenIncidentsOk returns a tuple with the OpenIncidents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpenIncidents

`func (o *InlineResponse20016) SetOpenIncidents(v []Incident)`

SetOpenIncidents sets OpenIncidents field to given value.

### HasOpenIncidents

`func (o *InlineResponse20016) HasOpenIncidents() bool`

HasOpenIncidents returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


