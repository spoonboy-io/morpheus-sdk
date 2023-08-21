# CatalogCartItemCreateItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to [**CatalogCartItemCreateItemType**](catalogCartItemCreate_item_type.md) |  | [optional] 
**Quantity** | Pointer to **int64** | Quantity for this catalog item. Will be overridden to 1 if quantity not allowed by the item type.  | [optional] 
**Config** | Pointer to **map[string]interface{}** | Config Object, required options depend on the catalog item type&#39;s associated option types. The values passed in here are injected into the instance config or app spec or workflow script(s) defined by the type.  | [optional] 
**Context** | Pointer to **string** | Context Type for running the workflow, determines if a target resource must be selected. &#x60;instance&#x60;, &#x60;server&#x60;, or &#x60;appliance&#x60;. This may only be passed if the type allows it, usually the type determines the context for the user. Only applies to type &#x60;workflow&#x60;.  | [optional] 
**Target** | Pointer to **int64** | Resource (Instance or Server) ID for context when running the &#x60;workflow&#x60;. Only applies to type &#x60;workflow&#x60; and only required when context is &#x60;instance&#x60; or &#x60;server&#x60;.  | [optional] 

## Methods

### NewCatalogCartItemCreateItem

`func NewCatalogCartItemCreateItem() *CatalogCartItemCreateItem`

NewCatalogCartItemCreateItem instantiates a new CatalogCartItemCreateItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCatalogCartItemCreateItemWithDefaults

`func NewCatalogCartItemCreateItemWithDefaults() *CatalogCartItemCreateItem`

NewCatalogCartItemCreateItemWithDefaults instantiates a new CatalogCartItemCreateItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *CatalogCartItemCreateItem) GetType() CatalogCartItemCreateItemType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CatalogCartItemCreateItem) GetTypeOk() (*CatalogCartItemCreateItemType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CatalogCartItemCreateItem) SetType(v CatalogCartItemCreateItemType)`

SetType sets Type field to given value.

### HasType

`func (o *CatalogCartItemCreateItem) HasType() bool`

HasType returns a boolean if a field has been set.

### GetQuantity

`func (o *CatalogCartItemCreateItem) GetQuantity() int64`

GetQuantity returns the Quantity field if non-nil, zero value otherwise.

### GetQuantityOk

`func (o *CatalogCartItemCreateItem) GetQuantityOk() (*int64, bool)`

GetQuantityOk returns a tuple with the Quantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantity

`func (o *CatalogCartItemCreateItem) SetQuantity(v int64)`

SetQuantity sets Quantity field to given value.

### HasQuantity

`func (o *CatalogCartItemCreateItem) HasQuantity() bool`

HasQuantity returns a boolean if a field has been set.

### GetConfig

`func (o *CatalogCartItemCreateItem) GetConfig() map[string]interface{}`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *CatalogCartItemCreateItem) GetConfigOk() (*map[string]interface{}, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *CatalogCartItemCreateItem) SetConfig(v map[string]interface{})`

SetConfig sets Config field to given value.

### HasConfig

`func (o *CatalogCartItemCreateItem) HasConfig() bool`

HasConfig returns a boolean if a field has been set.

### GetContext

`func (o *CatalogCartItemCreateItem) GetContext() string`

GetContext returns the Context field if non-nil, zero value otherwise.

### GetContextOk

`func (o *CatalogCartItemCreateItem) GetContextOk() (*string, bool)`

GetContextOk returns a tuple with the Context field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContext

`func (o *CatalogCartItemCreateItem) SetContext(v string)`

SetContext sets Context field to given value.

### HasContext

`func (o *CatalogCartItemCreateItem) HasContext() bool`

HasContext returns a boolean if a field has been set.

### GetTarget

`func (o *CatalogCartItemCreateItem) GetTarget() int64`

GetTarget returns the Target field if non-nil, zero value otherwise.

### GetTargetOk

`func (o *CatalogCartItemCreateItem) GetTargetOk() (*int64, bool)`

GetTargetOk returns a tuple with the Target field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTarget

`func (o *CatalogCartItemCreateItem) SetTarget(v int64)`

SetTarget sets Target field to given value.

### HasTarget

`func (o *CatalogCartItemCreateItem) HasTarget() bool`

HasTarget returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


