# IntegrationInventory

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **NullableString** |  | [optional] 
**ExternalId** | Pointer to **NullableString** |  | [optional] 
**Owner** | Pointer to [**InlineResponse20040AppDeployInstance**](inline_response_200_40_appDeploy_instance.md) |  | [optional] 
**Tenants** | Pointer to [**[]InlineResponse20040AppDeployInstance**](InlineResponse20040AppDeployInstance.md) |  | [optional] 

## Methods

### NewIntegrationInventory

`func NewIntegrationInventory() *IntegrationInventory`

NewIntegrationInventory instantiates a new IntegrationInventory object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIntegrationInventoryWithDefaults

`func NewIntegrationInventoryWithDefaults() *IntegrationInventory`

NewIntegrationInventoryWithDefaults instantiates a new IntegrationInventory object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *IntegrationInventory) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *IntegrationInventory) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *IntegrationInventory) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *IntegrationInventory) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *IntegrationInventory) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *IntegrationInventory) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *IntegrationInventory) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *IntegrationInventory) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *IntegrationInventory) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *IntegrationInventory) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *IntegrationInventory) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *IntegrationInventory) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *IntegrationInventory) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *IntegrationInventory) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetExternalId

`func (o *IntegrationInventory) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *IntegrationInventory) GetExternalIdOk() (*string, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *IntegrationInventory) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.

### HasExternalId

`func (o *IntegrationInventory) HasExternalId() bool`

HasExternalId returns a boolean if a field has been set.

### SetExternalIdNil

`func (o *IntegrationInventory) SetExternalIdNil(b bool)`

 SetExternalIdNil sets the value for ExternalId to be an explicit nil

### UnsetExternalId
`func (o *IntegrationInventory) UnsetExternalId()`

UnsetExternalId ensures that no value is present for ExternalId, not even an explicit nil
### GetOwner

`func (o *IntegrationInventory) GetOwner() InlineResponse20040AppDeployInstance`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *IntegrationInventory) GetOwnerOk() (*InlineResponse20040AppDeployInstance, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *IntegrationInventory) SetOwner(v InlineResponse20040AppDeployInstance)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *IntegrationInventory) HasOwner() bool`

HasOwner returns a boolean if a field has been set.

### GetTenants

`func (o *IntegrationInventory) GetTenants() []InlineResponse20040AppDeployInstance`

GetTenants returns the Tenants field if non-nil, zero value otherwise.

### GetTenantsOk

`func (o *IntegrationInventory) GetTenantsOk() (*[]InlineResponse20040AppDeployInstance, bool)`

GetTenantsOk returns a tuple with the Tenants field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenants

`func (o *IntegrationInventory) SetTenants(v []InlineResponse20040AppDeployInstance)`

SetTenants sets Tenants field to given value.

### HasTenants

`func (o *IntegrationInventory) HasTenants() bool`

HasTenants returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


