# AppStateSpecs

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Template** | Pointer to [**AppStateTemplate**](appState_template.md) |  | [optional] 
**Isolated** | Pointer to **bool** |  | [optional] 

## Methods

### NewAppStateSpecs

`func NewAppStateSpecs() *AppStateSpecs`

NewAppStateSpecs instantiates a new AppStateSpecs object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAppStateSpecsWithDefaults

`func NewAppStateSpecsWithDefaults() *AppStateSpecs`

NewAppStateSpecsWithDefaults instantiates a new AppStateSpecs object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AppStateSpecs) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AppStateSpecs) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AppStateSpecs) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *AppStateSpecs) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *AppStateSpecs) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AppStateSpecs) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AppStateSpecs) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *AppStateSpecs) HasName() bool`

HasName returns a boolean if a field has been set.

### GetTemplate

`func (o *AppStateSpecs) GetTemplate() AppStateTemplate`

GetTemplate returns the Template field if non-nil, zero value otherwise.

### GetTemplateOk

`func (o *AppStateSpecs) GetTemplateOk() (*AppStateTemplate, bool)`

GetTemplateOk returns a tuple with the Template field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplate

`func (o *AppStateSpecs) SetTemplate(v AppStateTemplate)`

SetTemplate sets Template field to given value.

### HasTemplate

`func (o *AppStateSpecs) HasTemplate() bool`

HasTemplate returns a boolean if a field has been set.

### GetIsolated

`func (o *AppStateSpecs) GetIsolated() bool`

GetIsolated returns the Isolated field if non-nil, zero value otherwise.

### GetIsolatedOk

`func (o *AppStateSpecs) GetIsolatedOk() (*bool, bool)`

GetIsolatedOk returns a tuple with the Isolated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsolated

`func (o *AppStateSpecs) SetIsolated(v bool)`

SetIsolated sets Isolated field to given value.

### HasIsolated

`func (o *AppStateSpecs) HasIsolated() bool`

HasIsolated returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


