# InlineResponse20040AppDeploy

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** |  | [optional] 
**InstanceId** | Pointer to **int64** |  | [optional] 
**Instance** | Pointer to [**InlineResponse20040AppDeployInstance**](inline_response_200_40_appDeploy_instance.md) |  | [optional] 
**Deployment** | Pointer to [**InlineResponse20040AppDeployDeployment**](inline_response_200_40_appDeploy_deployment.md) |  | [optional] 
**DeploymentVersionId** | Pointer to **int64** |  | [optional] 
**DeploymentVersion** | Pointer to [**InlineResponse20040AppDeployDeploymentVersion**](inline_response_200_40_appDeploy_deploymentVersion.md) |  | [optional] 
**Config** | Pointer to **map[string]interface{}** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**DeployDate** | Pointer to **time.Time** |  | [optional] 
**DateCreated** | Pointer to **time.Time** |  | [optional] 
**LastUpdated** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewInlineResponse20040AppDeploy

`func NewInlineResponse20040AppDeploy() *InlineResponse20040AppDeploy`

NewInlineResponse20040AppDeploy instantiates a new InlineResponse20040AppDeploy object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse20040AppDeployWithDefaults

`func NewInlineResponse20040AppDeployWithDefaults() *InlineResponse20040AppDeploy`

NewInlineResponse20040AppDeployWithDefaults instantiates a new InlineResponse20040AppDeploy object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *InlineResponse20040AppDeploy) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *InlineResponse20040AppDeploy) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *InlineResponse20040AppDeploy) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *InlineResponse20040AppDeploy) HasId() bool`

HasId returns a boolean if a field has been set.

### GetInstanceId

`func (o *InlineResponse20040AppDeploy) GetInstanceId() int64`

GetInstanceId returns the InstanceId field if non-nil, zero value otherwise.

### GetInstanceIdOk

`func (o *InlineResponse20040AppDeploy) GetInstanceIdOk() (*int64, bool)`

GetInstanceIdOk returns a tuple with the InstanceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceId

`func (o *InlineResponse20040AppDeploy) SetInstanceId(v int64)`

SetInstanceId sets InstanceId field to given value.

### HasInstanceId

`func (o *InlineResponse20040AppDeploy) HasInstanceId() bool`

HasInstanceId returns a boolean if a field has been set.

### GetInstance

`func (o *InlineResponse20040AppDeploy) GetInstance() InlineResponse20040AppDeployInstance`

GetInstance returns the Instance field if non-nil, zero value otherwise.

### GetInstanceOk

`func (o *InlineResponse20040AppDeploy) GetInstanceOk() (*InlineResponse20040AppDeployInstance, bool)`

GetInstanceOk returns a tuple with the Instance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstance

`func (o *InlineResponse20040AppDeploy) SetInstance(v InlineResponse20040AppDeployInstance)`

SetInstance sets Instance field to given value.

### HasInstance

`func (o *InlineResponse20040AppDeploy) HasInstance() bool`

HasInstance returns a boolean if a field has been set.

### GetDeployment

`func (o *InlineResponse20040AppDeploy) GetDeployment() InlineResponse20040AppDeployDeployment`

GetDeployment returns the Deployment field if non-nil, zero value otherwise.

### GetDeploymentOk

`func (o *InlineResponse20040AppDeploy) GetDeploymentOk() (*InlineResponse20040AppDeployDeployment, bool)`

GetDeploymentOk returns a tuple with the Deployment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeployment

`func (o *InlineResponse20040AppDeploy) SetDeployment(v InlineResponse20040AppDeployDeployment)`

SetDeployment sets Deployment field to given value.

### HasDeployment

`func (o *InlineResponse20040AppDeploy) HasDeployment() bool`

HasDeployment returns a boolean if a field has been set.

### GetDeploymentVersionId

`func (o *InlineResponse20040AppDeploy) GetDeploymentVersionId() int64`

GetDeploymentVersionId returns the DeploymentVersionId field if non-nil, zero value otherwise.

### GetDeploymentVersionIdOk

`func (o *InlineResponse20040AppDeploy) GetDeploymentVersionIdOk() (*int64, bool)`

GetDeploymentVersionIdOk returns a tuple with the DeploymentVersionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeploymentVersionId

`func (o *InlineResponse20040AppDeploy) SetDeploymentVersionId(v int64)`

SetDeploymentVersionId sets DeploymentVersionId field to given value.

### HasDeploymentVersionId

`func (o *InlineResponse20040AppDeploy) HasDeploymentVersionId() bool`

HasDeploymentVersionId returns a boolean if a field has been set.

### GetDeploymentVersion

`func (o *InlineResponse20040AppDeploy) GetDeploymentVersion() InlineResponse20040AppDeployDeploymentVersion`

GetDeploymentVersion returns the DeploymentVersion field if non-nil, zero value otherwise.

### GetDeploymentVersionOk

`func (o *InlineResponse20040AppDeploy) GetDeploymentVersionOk() (*InlineResponse20040AppDeployDeploymentVersion, bool)`

GetDeploymentVersionOk returns a tuple with the DeploymentVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeploymentVersion

`func (o *InlineResponse20040AppDeploy) SetDeploymentVersion(v InlineResponse20040AppDeployDeploymentVersion)`

SetDeploymentVersion sets DeploymentVersion field to given value.

### HasDeploymentVersion

`func (o *InlineResponse20040AppDeploy) HasDeploymentVersion() bool`

HasDeploymentVersion returns a boolean if a field has been set.

### GetConfig

`func (o *InlineResponse20040AppDeploy) GetConfig() map[string]interface{}`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *InlineResponse20040AppDeploy) GetConfigOk() (*map[string]interface{}, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *InlineResponse20040AppDeploy) SetConfig(v map[string]interface{})`

SetConfig sets Config field to given value.

### HasConfig

`func (o *InlineResponse20040AppDeploy) HasConfig() bool`

HasConfig returns a boolean if a field has been set.

### GetStatus

`func (o *InlineResponse20040AppDeploy) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *InlineResponse20040AppDeploy) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *InlineResponse20040AppDeploy) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *InlineResponse20040AppDeploy) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetDeployDate

`func (o *InlineResponse20040AppDeploy) GetDeployDate() time.Time`

GetDeployDate returns the DeployDate field if non-nil, zero value otherwise.

### GetDeployDateOk

`func (o *InlineResponse20040AppDeploy) GetDeployDateOk() (*time.Time, bool)`

GetDeployDateOk returns a tuple with the DeployDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeployDate

`func (o *InlineResponse20040AppDeploy) SetDeployDate(v time.Time)`

SetDeployDate sets DeployDate field to given value.

### HasDeployDate

`func (o *InlineResponse20040AppDeploy) HasDeployDate() bool`

HasDeployDate returns a boolean if a field has been set.

### GetDateCreated

`func (o *InlineResponse20040AppDeploy) GetDateCreated() time.Time`

GetDateCreated returns the DateCreated field if non-nil, zero value otherwise.

### GetDateCreatedOk

`func (o *InlineResponse20040AppDeploy) GetDateCreatedOk() (*time.Time, bool)`

GetDateCreatedOk returns a tuple with the DateCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateCreated

`func (o *InlineResponse20040AppDeploy) SetDateCreated(v time.Time)`

SetDateCreated sets DateCreated field to given value.

### HasDateCreated

`func (o *InlineResponse20040AppDeploy) HasDateCreated() bool`

HasDateCreated returns a boolean if a field has been set.

### GetLastUpdated

`func (o *InlineResponse20040AppDeploy) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *InlineResponse20040AppDeploy) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *InlineResponse20040AppDeploy) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.

### HasLastUpdated

`func (o *InlineResponse20040AppDeploy) HasLastUpdated() bool`

HasLastUpdated returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


