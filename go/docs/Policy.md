# Policy

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **NullableString** |  | [optional] 
**PolicyType** | Pointer to [**InlineResponse20079LoadBalancerMonitorLoadBalancerType**](inline_response_200_79_loadBalancerMonitor_loadBalancer_type.md) |  | [optional] 
**Zone** | Pointer to [**NullableInlineResponse20082LoadBalancerInstanceSslCert**](inline_response_200_82_loadBalancerInstance_sslCert.md) |  | [optional] 
**Site** | Pointer to [**NullableInlineResponse20082LoadBalancerInstanceSslCert**](inline_response_200_82_loadBalancerInstance_sslCert.md) |  | [optional] 
**User** | Pointer to [**NullableInlineResponse20083LoadBalancerNodeCreatedBy**](inline_response_200_83_loadBalancerNode_createdBy.md) |  | [optional] 
**Role** | Pointer to [**NullablePolicyRole**](policy_role.md) |  | [optional] 
**RefType** | Pointer to **NullableString** |  | [optional] 
**RefId** | Pointer to **NullableString** |  | [optional] 
**EachUser** | Pointer to **NullableBool** |  | [optional] 
**Config** | Pointer to [**OneOfobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobject**](oneOf&lt;object,object,object,object,object,object,object,object,object,object,object,object,object,object,object,object,object,object,object,object,object,object,object,object,object,object,object,object,object&gt;.md) |  | [optional] 
**Enabled** | Pointer to **bool** |  | [optional] 
**Owner** | Pointer to [**InlineResponse20040AppDeployInstance**](inline_response_200_40_appDeploy_instance.md) |  | [optional] 
**Accounts** | Pointer to [**[]InlineResponse20040AppDeployInstance**](InlineResponse20040AppDeployInstance.md) |  | [optional] 

## Methods

### NewPolicy

`func NewPolicy() *Policy`

NewPolicy instantiates a new Policy object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPolicyWithDefaults

`func NewPolicyWithDefaults() *Policy`

NewPolicyWithDefaults instantiates a new Policy object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Policy) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Policy) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Policy) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *Policy) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *Policy) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Policy) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Policy) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Policy) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *Policy) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Policy) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Policy) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *Policy) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *Policy) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *Policy) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetPolicyType

`func (o *Policy) GetPolicyType() InlineResponse20079LoadBalancerMonitorLoadBalancerType`

GetPolicyType returns the PolicyType field if non-nil, zero value otherwise.

### GetPolicyTypeOk

`func (o *Policy) GetPolicyTypeOk() (*InlineResponse20079LoadBalancerMonitorLoadBalancerType, bool)`

GetPolicyTypeOk returns a tuple with the PolicyType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyType

`func (o *Policy) SetPolicyType(v InlineResponse20079LoadBalancerMonitorLoadBalancerType)`

SetPolicyType sets PolicyType field to given value.

### HasPolicyType

`func (o *Policy) HasPolicyType() bool`

HasPolicyType returns a boolean if a field has been set.

### GetZone

`func (o *Policy) GetZone() InlineResponse20082LoadBalancerInstanceSslCert`

GetZone returns the Zone field if non-nil, zero value otherwise.

### GetZoneOk

`func (o *Policy) GetZoneOk() (*InlineResponse20082LoadBalancerInstanceSslCert, bool)`

GetZoneOk returns a tuple with the Zone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZone

`func (o *Policy) SetZone(v InlineResponse20082LoadBalancerInstanceSslCert)`

SetZone sets Zone field to given value.

### HasZone

`func (o *Policy) HasZone() bool`

HasZone returns a boolean if a field has been set.

### SetZoneNil

`func (o *Policy) SetZoneNil(b bool)`

 SetZoneNil sets the value for Zone to be an explicit nil

### UnsetZone
`func (o *Policy) UnsetZone()`

UnsetZone ensures that no value is present for Zone, not even an explicit nil
### GetSite

`func (o *Policy) GetSite() InlineResponse20082LoadBalancerInstanceSslCert`

GetSite returns the Site field if non-nil, zero value otherwise.

### GetSiteOk

`func (o *Policy) GetSiteOk() (*InlineResponse20082LoadBalancerInstanceSslCert, bool)`

GetSiteOk returns a tuple with the Site field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSite

`func (o *Policy) SetSite(v InlineResponse20082LoadBalancerInstanceSslCert)`

SetSite sets Site field to given value.

### HasSite

`func (o *Policy) HasSite() bool`

HasSite returns a boolean if a field has been set.

### SetSiteNil

`func (o *Policy) SetSiteNil(b bool)`

 SetSiteNil sets the value for Site to be an explicit nil

### UnsetSite
`func (o *Policy) UnsetSite()`

UnsetSite ensures that no value is present for Site, not even an explicit nil
### GetUser

`func (o *Policy) GetUser() InlineResponse20083LoadBalancerNodeCreatedBy`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *Policy) GetUserOk() (*InlineResponse20083LoadBalancerNodeCreatedBy, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *Policy) SetUser(v InlineResponse20083LoadBalancerNodeCreatedBy)`

SetUser sets User field to given value.

### HasUser

`func (o *Policy) HasUser() bool`

HasUser returns a boolean if a field has been set.

### SetUserNil

`func (o *Policy) SetUserNil(b bool)`

 SetUserNil sets the value for User to be an explicit nil

### UnsetUser
`func (o *Policy) UnsetUser()`

UnsetUser ensures that no value is present for User, not even an explicit nil
### GetRole

`func (o *Policy) GetRole() PolicyRole`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *Policy) GetRoleOk() (*PolicyRole, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *Policy) SetRole(v PolicyRole)`

SetRole sets Role field to given value.

### HasRole

`func (o *Policy) HasRole() bool`

HasRole returns a boolean if a field has been set.

### SetRoleNil

`func (o *Policy) SetRoleNil(b bool)`

 SetRoleNil sets the value for Role to be an explicit nil

### UnsetRole
`func (o *Policy) UnsetRole()`

UnsetRole ensures that no value is present for Role, not even an explicit nil
### GetRefType

`func (o *Policy) GetRefType() string`

GetRefType returns the RefType field if non-nil, zero value otherwise.

### GetRefTypeOk

`func (o *Policy) GetRefTypeOk() (*string, bool)`

GetRefTypeOk returns a tuple with the RefType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefType

`func (o *Policy) SetRefType(v string)`

SetRefType sets RefType field to given value.

### HasRefType

`func (o *Policy) HasRefType() bool`

HasRefType returns a boolean if a field has been set.

### SetRefTypeNil

`func (o *Policy) SetRefTypeNil(b bool)`

 SetRefTypeNil sets the value for RefType to be an explicit nil

### UnsetRefType
`func (o *Policy) UnsetRefType()`

UnsetRefType ensures that no value is present for RefType, not even an explicit nil
### GetRefId

`func (o *Policy) GetRefId() string`

GetRefId returns the RefId field if non-nil, zero value otherwise.

### GetRefIdOk

`func (o *Policy) GetRefIdOk() (*string, bool)`

GetRefIdOk returns a tuple with the RefId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefId

`func (o *Policy) SetRefId(v string)`

SetRefId sets RefId field to given value.

### HasRefId

`func (o *Policy) HasRefId() bool`

HasRefId returns a boolean if a field has been set.

### SetRefIdNil

`func (o *Policy) SetRefIdNil(b bool)`

 SetRefIdNil sets the value for RefId to be an explicit nil

### UnsetRefId
`func (o *Policy) UnsetRefId()`

UnsetRefId ensures that no value is present for RefId, not even an explicit nil
### GetEachUser

`func (o *Policy) GetEachUser() bool`

GetEachUser returns the EachUser field if non-nil, zero value otherwise.

### GetEachUserOk

`func (o *Policy) GetEachUserOk() (*bool, bool)`

GetEachUserOk returns a tuple with the EachUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEachUser

`func (o *Policy) SetEachUser(v bool)`

SetEachUser sets EachUser field to given value.

### HasEachUser

`func (o *Policy) HasEachUser() bool`

HasEachUser returns a boolean if a field has been set.

### SetEachUserNil

`func (o *Policy) SetEachUserNil(b bool)`

 SetEachUserNil sets the value for EachUser to be an explicit nil

### UnsetEachUser
`func (o *Policy) UnsetEachUser()`

UnsetEachUser ensures that no value is present for EachUser, not even an explicit nil
### GetConfig

`func (o *Policy) GetConfig() OneOfobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobject`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *Policy) GetConfigOk() (*OneOfobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobject, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *Policy) SetConfig(v OneOfobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobject)`

SetConfig sets Config field to given value.

### HasConfig

`func (o *Policy) HasConfig() bool`

HasConfig returns a boolean if a field has been set.

### GetEnabled

`func (o *Policy) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *Policy) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *Policy) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *Policy) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetOwner

`func (o *Policy) GetOwner() InlineResponse20040AppDeployInstance`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *Policy) GetOwnerOk() (*InlineResponse20040AppDeployInstance, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *Policy) SetOwner(v InlineResponse20040AppDeployInstance)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *Policy) HasOwner() bool`

HasOwner returns a boolean if a field has been set.

### GetAccounts

`func (o *Policy) GetAccounts() []InlineResponse20040AppDeployInstance`

GetAccounts returns the Accounts field if non-nil, zero value otherwise.

### GetAccountsOk

`func (o *Policy) GetAccountsOk() (*[]InlineResponse20040AppDeployInstance, bool)`

GetAccountsOk returns a tuple with the Accounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccounts

`func (o *Policy) SetAccounts(v []InlineResponse20040AppDeployInstance)`

SetAccounts sets Accounts field to given value.

### HasAccounts

`func (o *Policy) HasAccounts() bool`

HasAccounts returns a boolean if a field has been set.

### SetAccountsNil

`func (o *Policy) SetAccountsNil(b bool)`

 SetAccountsNil sets the value for Accounts to be an explicit nil

### UnsetAccounts
`func (o *Policy) UnsetAccounts()`

UnsetAccounts ensures that no value is present for Accounts, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


