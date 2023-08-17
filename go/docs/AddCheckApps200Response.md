# AddCheckApps200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Success** | Pointer to **bool** |  | [optional] 
**CheckApp** | Pointer to [**CheckApp**](CheckApp.md) |  | [optional] 

## Methods

### NewAddCheckApps200Response

`func NewAddCheckApps200Response() *AddCheckApps200Response`

NewAddCheckApps200Response instantiates a new AddCheckApps200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddCheckApps200ResponseWithDefaults

`func NewAddCheckApps200ResponseWithDefaults() *AddCheckApps200Response`

NewAddCheckApps200ResponseWithDefaults instantiates a new AddCheckApps200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSuccess

`func (o *AddCheckApps200Response) GetSuccess() bool`

GetSuccess returns the Success field if non-nil, zero value otherwise.

### GetSuccessOk

`func (o *AddCheckApps200Response) GetSuccessOk() (*bool, bool)`

GetSuccessOk returns a tuple with the Success field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccess

`func (o *AddCheckApps200Response) SetSuccess(v bool)`

SetSuccess sets Success field to given value.

### HasSuccess

`func (o *AddCheckApps200Response) HasSuccess() bool`

HasSuccess returns a boolean if a field has been set.

### GetCheckApp

`func (o *AddCheckApps200Response) GetCheckApp() CheckApp`

GetCheckApp returns the CheckApp field if non-nil, zero value otherwise.

### GetCheckAppOk

`func (o *AddCheckApps200Response) GetCheckAppOk() (*CheckApp, bool)`

GetCheckAppOk returns a tuple with the CheckApp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCheckApp

`func (o *AddCheckApps200Response) SetCheckApp(v CheckApp)`

SetCheckApp sets CheckApp field to given value.

### HasCheckApp

`func (o *AddCheckApps200Response) HasCheckApp() bool`

HasCheckApp returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


