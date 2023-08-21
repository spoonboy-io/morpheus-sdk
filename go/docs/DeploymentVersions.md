# DeploymentVersions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** |  | [optional] 
**DeployType** | Pointer to **string** |  | [optional] 
**FetchUrl** | Pointer to **NullableString** |  | [optional] 
**GitUrl** | Pointer to **NullableString** |  | [optional] 
**GitRef** | Pointer to **NullableString** |  | [optional] 
**UserVersion** | Pointer to **string** |  | [optional] 
**Version** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**DateCreated** | Pointer to **time.Time** |  | [optional] 
**LastUpdated** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewDeploymentVersions

`func NewDeploymentVersions() *DeploymentVersions`

NewDeploymentVersions instantiates a new DeploymentVersions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeploymentVersionsWithDefaults

`func NewDeploymentVersionsWithDefaults() *DeploymentVersions`

NewDeploymentVersionsWithDefaults instantiates a new DeploymentVersions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *DeploymentVersions) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DeploymentVersions) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DeploymentVersions) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *DeploymentVersions) HasId() bool`

HasId returns a boolean if a field has been set.

### GetDeployType

`func (o *DeploymentVersions) GetDeployType() string`

GetDeployType returns the DeployType field if non-nil, zero value otherwise.

### GetDeployTypeOk

`func (o *DeploymentVersions) GetDeployTypeOk() (*string, bool)`

GetDeployTypeOk returns a tuple with the DeployType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeployType

`func (o *DeploymentVersions) SetDeployType(v string)`

SetDeployType sets DeployType field to given value.

### HasDeployType

`func (o *DeploymentVersions) HasDeployType() bool`

HasDeployType returns a boolean if a field has been set.

### GetFetchUrl

`func (o *DeploymentVersions) GetFetchUrl() string`

GetFetchUrl returns the FetchUrl field if non-nil, zero value otherwise.

### GetFetchUrlOk

`func (o *DeploymentVersions) GetFetchUrlOk() (*string, bool)`

GetFetchUrlOk returns a tuple with the FetchUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFetchUrl

`func (o *DeploymentVersions) SetFetchUrl(v string)`

SetFetchUrl sets FetchUrl field to given value.

### HasFetchUrl

`func (o *DeploymentVersions) HasFetchUrl() bool`

HasFetchUrl returns a boolean if a field has been set.

### SetFetchUrlNil

`func (o *DeploymentVersions) SetFetchUrlNil(b bool)`

 SetFetchUrlNil sets the value for FetchUrl to be an explicit nil

### UnsetFetchUrl
`func (o *DeploymentVersions) UnsetFetchUrl()`

UnsetFetchUrl ensures that no value is present for FetchUrl, not even an explicit nil
### GetGitUrl

`func (o *DeploymentVersions) GetGitUrl() string`

GetGitUrl returns the GitUrl field if non-nil, zero value otherwise.

### GetGitUrlOk

`func (o *DeploymentVersions) GetGitUrlOk() (*string, bool)`

GetGitUrlOk returns a tuple with the GitUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGitUrl

`func (o *DeploymentVersions) SetGitUrl(v string)`

SetGitUrl sets GitUrl field to given value.

### HasGitUrl

`func (o *DeploymentVersions) HasGitUrl() bool`

HasGitUrl returns a boolean if a field has been set.

### SetGitUrlNil

`func (o *DeploymentVersions) SetGitUrlNil(b bool)`

 SetGitUrlNil sets the value for GitUrl to be an explicit nil

### UnsetGitUrl
`func (o *DeploymentVersions) UnsetGitUrl()`

UnsetGitUrl ensures that no value is present for GitUrl, not even an explicit nil
### GetGitRef

`func (o *DeploymentVersions) GetGitRef() string`

GetGitRef returns the GitRef field if non-nil, zero value otherwise.

### GetGitRefOk

`func (o *DeploymentVersions) GetGitRefOk() (*string, bool)`

GetGitRefOk returns a tuple with the GitRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGitRef

`func (o *DeploymentVersions) SetGitRef(v string)`

SetGitRef sets GitRef field to given value.

### HasGitRef

`func (o *DeploymentVersions) HasGitRef() bool`

HasGitRef returns a boolean if a field has been set.

### SetGitRefNil

`func (o *DeploymentVersions) SetGitRefNil(b bool)`

 SetGitRefNil sets the value for GitRef to be an explicit nil

### UnsetGitRef
`func (o *DeploymentVersions) UnsetGitRef()`

UnsetGitRef ensures that no value is present for GitRef, not even an explicit nil
### GetUserVersion

`func (o *DeploymentVersions) GetUserVersion() string`

GetUserVersion returns the UserVersion field if non-nil, zero value otherwise.

### GetUserVersionOk

`func (o *DeploymentVersions) GetUserVersionOk() (*string, bool)`

GetUserVersionOk returns a tuple with the UserVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserVersion

`func (o *DeploymentVersions) SetUserVersion(v string)`

SetUserVersion sets UserVersion field to given value.

### HasUserVersion

`func (o *DeploymentVersions) HasUserVersion() bool`

HasUserVersion returns a boolean if a field has been set.

### GetVersion

`func (o *DeploymentVersions) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *DeploymentVersions) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *DeploymentVersions) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *DeploymentVersions) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetStatus

`func (o *DeploymentVersions) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *DeploymentVersions) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *DeploymentVersions) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *DeploymentVersions) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetDateCreated

`func (o *DeploymentVersions) GetDateCreated() time.Time`

GetDateCreated returns the DateCreated field if non-nil, zero value otherwise.

### GetDateCreatedOk

`func (o *DeploymentVersions) GetDateCreatedOk() (*time.Time, bool)`

GetDateCreatedOk returns a tuple with the DateCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateCreated

`func (o *DeploymentVersions) SetDateCreated(v time.Time)`

SetDateCreated sets DateCreated field to given value.

### HasDateCreated

`func (o *DeploymentVersions) HasDateCreated() bool`

HasDateCreated returns a boolean if a field has been set.

### GetLastUpdated

`func (o *DeploymentVersions) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *DeploymentVersions) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *DeploymentVersions) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.

### HasLastUpdated

`func (o *DeploymentVersions) HasLastUpdated() bool`

HasLastUpdated returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


