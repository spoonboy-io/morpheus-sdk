# ListCheckApps200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Meta** | Pointer to [**MetaMeta**](MetaMeta.md) |  | [optional] 
**MonitorApps** | Pointer to [**[]CheckApp**](CheckApp.md) |  | [optional] 

## Methods

### NewListCheckApps200Response

`func NewListCheckApps200Response() *ListCheckApps200Response`

NewListCheckApps200Response instantiates a new ListCheckApps200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListCheckApps200ResponseWithDefaults

`func NewListCheckApps200ResponseWithDefaults() *ListCheckApps200Response`

NewListCheckApps200ResponseWithDefaults instantiates a new ListCheckApps200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMeta

`func (o *ListCheckApps200Response) GetMeta() MetaMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *ListCheckApps200Response) GetMetaOk() (*MetaMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *ListCheckApps200Response) SetMeta(v MetaMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *ListCheckApps200Response) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetMonitorApps

`func (o *ListCheckApps200Response) GetMonitorApps() []CheckApp`

GetMonitorApps returns the MonitorApps field if non-nil, zero value otherwise.

### GetMonitorAppsOk

`func (o *ListCheckApps200Response) GetMonitorAppsOk() (*[]CheckApp, bool)`

GetMonitorAppsOk returns a tuple with the MonitorApps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonitorApps

`func (o *ListCheckApps200Response) SetMonitorApps(v []CheckApp)`

SetMonitorApps sets MonitorApps field to given value.

### HasMonitorApps

`func (o *ListCheckApps200Response) HasMonitorApps() bool`

HasMonitorApps returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


