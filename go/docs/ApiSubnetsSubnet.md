# ApiSubnetsSubnet

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to [**ApiSubnetsSubnetType**](_api_subnets_subnet_type.md) |  | [optional] 
**Config** | Pointer to **map[string]interface{}** | Configuration object. Settings vary by type. | [optional] 
**Tenants** | Pointer to [**[]ApiBlueprintsIdUpdatePermissionsResourcePermissionSites**](ApiBlueprintsIdUpdatePermissionsResourcePermissionSites.md) | Array of tenant account ID objects that are allowed access | [optional] 
**Visibility** | Pointer to **string** | private or public | [optional] [default to "private"]
**Labels** | Pointer to **[]string** | Array of label strings, can be used for filtering. | [optional] 

## Methods

### NewApiSubnetsSubnet

`func NewApiSubnetsSubnet() *ApiSubnetsSubnet`

NewApiSubnetsSubnet instantiates a new ApiSubnetsSubnet object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiSubnetsSubnetWithDefaults

`func NewApiSubnetsSubnetWithDefaults() *ApiSubnetsSubnet`

NewApiSubnetsSubnetWithDefaults instantiates a new ApiSubnetsSubnet object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *ApiSubnetsSubnet) GetType() ApiSubnetsSubnetType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ApiSubnetsSubnet) GetTypeOk() (*ApiSubnetsSubnetType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ApiSubnetsSubnet) SetType(v ApiSubnetsSubnetType)`

SetType sets Type field to given value.

### HasType

`func (o *ApiSubnetsSubnet) HasType() bool`

HasType returns a boolean if a field has been set.

### GetConfig

`func (o *ApiSubnetsSubnet) GetConfig() map[string]interface{}`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *ApiSubnetsSubnet) GetConfigOk() (*map[string]interface{}, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *ApiSubnetsSubnet) SetConfig(v map[string]interface{})`

SetConfig sets Config field to given value.

### HasConfig

`func (o *ApiSubnetsSubnet) HasConfig() bool`

HasConfig returns a boolean if a field has been set.

### GetTenants

`func (o *ApiSubnetsSubnet) GetTenants() []ApiBlueprintsIdUpdatePermissionsResourcePermissionSites`

GetTenants returns the Tenants field if non-nil, zero value otherwise.

### GetTenantsOk

`func (o *ApiSubnetsSubnet) GetTenantsOk() (*[]ApiBlueprintsIdUpdatePermissionsResourcePermissionSites, bool)`

GetTenantsOk returns a tuple with the Tenants field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenants

`func (o *ApiSubnetsSubnet) SetTenants(v []ApiBlueprintsIdUpdatePermissionsResourcePermissionSites)`

SetTenants sets Tenants field to given value.

### HasTenants

`func (o *ApiSubnetsSubnet) HasTenants() bool`

HasTenants returns a boolean if a field has been set.

### GetVisibility

`func (o *ApiSubnetsSubnet) GetVisibility() string`

GetVisibility returns the Visibility field if non-nil, zero value otherwise.

### GetVisibilityOk

`func (o *ApiSubnetsSubnet) GetVisibilityOk() (*string, bool)`

GetVisibilityOk returns a tuple with the Visibility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisibility

`func (o *ApiSubnetsSubnet) SetVisibility(v string)`

SetVisibility sets Visibility field to given value.

### HasVisibility

`func (o *ApiSubnetsSubnet) HasVisibility() bool`

HasVisibility returns a boolean if a field has been set.

### GetLabels

`func (o *ApiSubnetsSubnet) GetLabels() []string`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *ApiSubnetsSubnet) GetLabelsOk() (*[]string, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *ApiSubnetsSubnet) SetLabels(v []string)`

SetLabels sets Labels field to given value.

### HasLabels

`func (o *ApiSubnetsSubnet) HasLabels() bool`

HasLabels returns a boolean if a field has been set.

### SetLabelsNil

`func (o *ApiSubnetsSubnet) SetLabelsNil(b bool)`

 SetLabelsNil sets the value for Labels to be an explicit nil

### UnsetLabels
`func (o *ApiSubnetsSubnet) UnsetLabels()`

UnsetLabels ensures that no value is present for Labels, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


