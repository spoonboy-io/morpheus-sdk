# GetAlerts200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Alert** | Pointer to [**Alert**](Alert.md) |  | [optional] 
**Checks** | Pointer to [**[]Check**](Check.md) |  | [optional] 
**CheckGroups** | Pointer to [**[]CheckGroup**](CheckGroup.md) |  | [optional] 
**Apps** | Pointer to [**[]CheckApp**](CheckApp.md) |  | [optional] 

## Methods

### NewGetAlerts200Response

`func NewGetAlerts200Response() *GetAlerts200Response`

NewGetAlerts200Response instantiates a new GetAlerts200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetAlerts200ResponseWithDefaults

`func NewGetAlerts200ResponseWithDefaults() *GetAlerts200Response`

NewGetAlerts200ResponseWithDefaults instantiates a new GetAlerts200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAlert

`func (o *GetAlerts200Response) GetAlert() Alert`

GetAlert returns the Alert field if non-nil, zero value otherwise.

### GetAlertOk

`func (o *GetAlerts200Response) GetAlertOk() (*Alert, bool)`

GetAlertOk returns a tuple with the Alert field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlert

`func (o *GetAlerts200Response) SetAlert(v Alert)`

SetAlert sets Alert field to given value.

### HasAlert

`func (o *GetAlerts200Response) HasAlert() bool`

HasAlert returns a boolean if a field has been set.

### GetChecks

`func (o *GetAlerts200Response) GetChecks() []Check`

GetChecks returns the Checks field if non-nil, zero value otherwise.

### GetChecksOk

`func (o *GetAlerts200Response) GetChecksOk() (*[]Check, bool)`

GetChecksOk returns a tuple with the Checks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChecks

`func (o *GetAlerts200Response) SetChecks(v []Check)`

SetChecks sets Checks field to given value.

### HasChecks

`func (o *GetAlerts200Response) HasChecks() bool`

HasChecks returns a boolean if a field has been set.

### GetCheckGroups

`func (o *GetAlerts200Response) GetCheckGroups() []CheckGroup`

GetCheckGroups returns the CheckGroups field if non-nil, zero value otherwise.

### GetCheckGroupsOk

`func (o *GetAlerts200Response) GetCheckGroupsOk() (*[]CheckGroup, bool)`

GetCheckGroupsOk returns a tuple with the CheckGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCheckGroups

`func (o *GetAlerts200Response) SetCheckGroups(v []CheckGroup)`

SetCheckGroups sets CheckGroups field to given value.

### HasCheckGroups

`func (o *GetAlerts200Response) HasCheckGroups() bool`

HasCheckGroups returns a boolean if a field has been set.

### GetApps

`func (o *GetAlerts200Response) GetApps() []CheckApp`

GetApps returns the Apps field if non-nil, zero value otherwise.

### GetAppsOk

`func (o *GetAlerts200Response) GetAppsOk() (*[]CheckApp, bool)`

GetAppsOk returns a tuple with the Apps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApps

`func (o *GetAlerts200Response) SetApps(v []CheckApp)`

SetApps sets Apps field to given value.

### HasApps

`func (o *GetAlerts200Response) HasApps() bool`

HasApps returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


