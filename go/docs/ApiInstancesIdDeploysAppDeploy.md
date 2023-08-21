# ApiInstancesIdDeploysAppDeploy

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DeploymentId** | Pointer to **int64** | Deployment ID. | [optional] 
**Version** | Pointer to **string** | Deployment Version number identifier (userVersion). Can be passed along with deploymentId to identify the version | [optional] 
**VersionId** | Pointer to **int64** | Deployment Version ID. This can be passed instead of deploymentId and version. | [optional] 
**Config** | Pointer to **map[string]interface{}** | Map of configuration properties that vary by instance type. | [optional] 
**StageOnly** | Pointer to **bool** | Stage Only, do not run the deploy right away and instead set status to staged so it can be deployed later on. | [optional] [default to false]

## Methods

### NewApiInstancesIdDeploysAppDeploy

`func NewApiInstancesIdDeploysAppDeploy() *ApiInstancesIdDeploysAppDeploy`

NewApiInstancesIdDeploysAppDeploy instantiates a new ApiInstancesIdDeploysAppDeploy object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiInstancesIdDeploysAppDeployWithDefaults

`func NewApiInstancesIdDeploysAppDeployWithDefaults() *ApiInstancesIdDeploysAppDeploy`

NewApiInstancesIdDeploysAppDeployWithDefaults instantiates a new ApiInstancesIdDeploysAppDeploy object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDeploymentId

`func (o *ApiInstancesIdDeploysAppDeploy) GetDeploymentId() int64`

GetDeploymentId returns the DeploymentId field if non-nil, zero value otherwise.

### GetDeploymentIdOk

`func (o *ApiInstancesIdDeploysAppDeploy) GetDeploymentIdOk() (*int64, bool)`

GetDeploymentIdOk returns a tuple with the DeploymentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeploymentId

`func (o *ApiInstancesIdDeploysAppDeploy) SetDeploymentId(v int64)`

SetDeploymentId sets DeploymentId field to given value.

### HasDeploymentId

`func (o *ApiInstancesIdDeploysAppDeploy) HasDeploymentId() bool`

HasDeploymentId returns a boolean if a field has been set.

### GetVersion

`func (o *ApiInstancesIdDeploysAppDeploy) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *ApiInstancesIdDeploysAppDeploy) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *ApiInstancesIdDeploysAppDeploy) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *ApiInstancesIdDeploysAppDeploy) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetVersionId

`func (o *ApiInstancesIdDeploysAppDeploy) GetVersionId() int64`

GetVersionId returns the VersionId field if non-nil, zero value otherwise.

### GetVersionIdOk

`func (o *ApiInstancesIdDeploysAppDeploy) GetVersionIdOk() (*int64, bool)`

GetVersionIdOk returns a tuple with the VersionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersionId

`func (o *ApiInstancesIdDeploysAppDeploy) SetVersionId(v int64)`

SetVersionId sets VersionId field to given value.

### HasVersionId

`func (o *ApiInstancesIdDeploysAppDeploy) HasVersionId() bool`

HasVersionId returns a boolean if a field has been set.

### GetConfig

`func (o *ApiInstancesIdDeploysAppDeploy) GetConfig() map[string]interface{}`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *ApiInstancesIdDeploysAppDeploy) GetConfigOk() (*map[string]interface{}, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *ApiInstancesIdDeploysAppDeploy) SetConfig(v map[string]interface{})`

SetConfig sets Config field to given value.

### HasConfig

`func (o *ApiInstancesIdDeploysAppDeploy) HasConfig() bool`

HasConfig returns a boolean if a field has been set.

### GetStageOnly

`func (o *ApiInstancesIdDeploysAppDeploy) GetStageOnly() bool`

GetStageOnly returns the StageOnly field if non-nil, zero value otherwise.

### GetStageOnlyOk

`func (o *ApiInstancesIdDeploysAppDeploy) GetStageOnlyOk() (*bool, bool)`

GetStageOnlyOk returns a tuple with the StageOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStageOnly

`func (o *ApiInstancesIdDeploysAppDeploy) SetStageOnly(v bool)`

SetStageOnly sets StageOnly field to given value.

### HasStageOnly

`func (o *ApiInstancesIdDeploysAppDeploy) HasStageOnly() bool`

HasStageOnly returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


