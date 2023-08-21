# CatalogOrderCreateItems

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to [**CatalogOrderCreateType**](catalogOrderCreate_type.md) |  | [optional] 
**Config** | **map[string]interface{}** | Config Object, required options depend on the catalog item type&#39;s associated option types. The values passed in here are injected into the instance config or app spec or workflow script(s) defined by the type.  | 
**Context** | Pointer to **string** | Context Type for running the workflow, determines if a target resource must be selected. &#x60;instance&#x60;, &#x60;server&#x60;, or &#x60;appliance&#x60;. This may only be passed if the type allows it, usually the type determines the context for the user. Only applies to type &#x60;workflow&#x60;.  | [optional] 
**Target** | Pointer to **int64** | Resource (Instance or Server) ID for context when running the &#x60;workflow&#x60;. Only applies to type &#x60;workflow&#x60; and only required when context is &#x60;instance&#x60; or &#x60;server&#x60;.  | [optional] 

## Methods

### NewCatalogOrderCreateItems

`func NewCatalogOrderCreateItems(config map[string]interface{}, ) *CatalogOrderCreateItems`

NewCatalogOrderCreateItems instantiates a new CatalogOrderCreateItems object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCatalogOrderCreateItemsWithDefaults

`func NewCatalogOrderCreateItemsWithDefaults() *CatalogOrderCreateItems`

NewCatalogOrderCreateItemsWithDefaults instantiates a new CatalogOrderCreateItems object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *CatalogOrderCreateItems) GetType() CatalogOrderCreateType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CatalogOrderCreateItems) GetTypeOk() (*CatalogOrderCreateType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CatalogOrderCreateItems) SetType(v CatalogOrderCreateType)`

SetType sets Type field to given value.

### HasType

`func (o *CatalogOrderCreateItems) HasType() bool`

HasType returns a boolean if a field has been set.

### GetConfig

`func (o *CatalogOrderCreateItems) GetConfig() map[string]interface{}`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *CatalogOrderCreateItems) GetConfigOk() (*map[string]interface{}, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *CatalogOrderCreateItems) SetConfig(v map[string]interface{})`

SetConfig sets Config field to given value.


### GetContext

`func (o *CatalogOrderCreateItems) GetContext() string`

GetContext returns the Context field if non-nil, zero value otherwise.

### GetContextOk

`func (o *CatalogOrderCreateItems) GetContextOk() (*string, bool)`

GetContextOk returns a tuple with the Context field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContext

`func (o *CatalogOrderCreateItems) SetContext(v string)`

SetContext sets Context field to given value.

### HasContext

`func (o *CatalogOrderCreateItems) HasContext() bool`

HasContext returns a boolean if a field has been set.

### GetTarget

`func (o *CatalogOrderCreateItems) GetTarget() int64`

GetTarget returns the Target field if non-nil, zero value otherwise.

### GetTargetOk

`func (o *CatalogOrderCreateItems) GetTargetOk() (*int64, bool)`

GetTargetOk returns a tuple with the Target field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTarget

`func (o *CatalogOrderCreateItems) SetTarget(v int64)`

SetTarget sets Target field to given value.

### HasTarget

`func (o *CatalogOrderCreateItems) HasTarget() bool`

HasTarget returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


