# ListAlerts200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Meta** | Pointer to [**MetaMeta**](MetaMeta.md) |  | [optional] 
**Alerts** | Pointer to [**[]Alert**](Alert.md) |  | [optional] 

## Methods

### NewListAlerts200Response

`func NewListAlerts200Response() *ListAlerts200Response`

NewListAlerts200Response instantiates a new ListAlerts200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListAlerts200ResponseWithDefaults

`func NewListAlerts200ResponseWithDefaults() *ListAlerts200Response`

NewListAlerts200ResponseWithDefaults instantiates a new ListAlerts200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMeta

`func (o *ListAlerts200Response) GetMeta() MetaMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *ListAlerts200Response) GetMetaOk() (*MetaMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *ListAlerts200Response) SetMeta(v MetaMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *ListAlerts200Response) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetAlerts

`func (o *ListAlerts200Response) GetAlerts() []Alert`

GetAlerts returns the Alerts field if non-nil, zero value otherwise.

### GetAlertsOk

`func (o *ListAlerts200Response) GetAlertsOk() (*[]Alert, bool)`

GetAlertsOk returns a tuple with the Alerts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlerts

`func (o *ListAlerts200Response) SetAlerts(v []Alert)`

SetAlerts sets Alerts field to given value.

### HasAlerts

`func (o *ListAlerts200Response) HasAlerts() bool`

HasAlerts returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


