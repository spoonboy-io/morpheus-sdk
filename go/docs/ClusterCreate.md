# ClusterCreate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | [**OneOfstringobject**](oneOf&lt;string,object&gt;.md) |  | 
**Name** | **string** | Name of the cluster to be created | 
**Description** | Pointer to **string** | Description of the cluster to be created | [optional] 
**Labels** | Pointer to **[]string** | Array of strings (keywords). This will override labels passed under the &#x60;server&#x60; object. | [optional] 
**Group** | [**ClusterCreateGroup**](clusterCreate_group.md) |  | 
**Cloud** | [**ClusterCreateCloud**](clusterCreate_cloud.md) |  | 
**Layout** | [**ClusterCreateLayout**](clusterCreate_layout.md) |  | 
**Server** | [**ClusterServerCreate**](clusterServerCreate.md) |  | 

## Methods

### NewClusterCreate

`func NewClusterCreate(type_ OneOfstringobject, name string, group ClusterCreateGroup, cloud ClusterCreateCloud, layout ClusterCreateLayout, server ClusterServerCreate, ) *ClusterCreate`

NewClusterCreate instantiates a new ClusterCreate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClusterCreateWithDefaults

`func NewClusterCreateWithDefaults() *ClusterCreate`

NewClusterCreateWithDefaults instantiates a new ClusterCreate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *ClusterCreate) GetType() OneOfstringobject`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ClusterCreate) GetTypeOk() (*OneOfstringobject, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ClusterCreate) SetType(v OneOfstringobject)`

SetType sets Type field to given value.


### GetName

`func (o *ClusterCreate) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ClusterCreate) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ClusterCreate) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *ClusterCreate) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ClusterCreate) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ClusterCreate) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ClusterCreate) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetLabels

`func (o *ClusterCreate) GetLabels() []string`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *ClusterCreate) GetLabelsOk() (*[]string, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *ClusterCreate) SetLabels(v []string)`

SetLabels sets Labels field to given value.

### HasLabels

`func (o *ClusterCreate) HasLabels() bool`

HasLabels returns a boolean if a field has been set.

### GetGroup

`func (o *ClusterCreate) GetGroup() ClusterCreateGroup`

GetGroup returns the Group field if non-nil, zero value otherwise.

### GetGroupOk

`func (o *ClusterCreate) GetGroupOk() (*ClusterCreateGroup, bool)`

GetGroupOk returns a tuple with the Group field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroup

`func (o *ClusterCreate) SetGroup(v ClusterCreateGroup)`

SetGroup sets Group field to given value.


### GetCloud

`func (o *ClusterCreate) GetCloud() ClusterCreateCloud`

GetCloud returns the Cloud field if non-nil, zero value otherwise.

### GetCloudOk

`func (o *ClusterCreate) GetCloudOk() (*ClusterCreateCloud, bool)`

GetCloudOk returns a tuple with the Cloud field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloud

`func (o *ClusterCreate) SetCloud(v ClusterCreateCloud)`

SetCloud sets Cloud field to given value.


### GetLayout

`func (o *ClusterCreate) GetLayout() ClusterCreateLayout`

GetLayout returns the Layout field if non-nil, zero value otherwise.

### GetLayoutOk

`func (o *ClusterCreate) GetLayoutOk() (*ClusterCreateLayout, bool)`

GetLayoutOk returns a tuple with the Layout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLayout

`func (o *ClusterCreate) SetLayout(v ClusterCreateLayout)`

SetLayout sets Layout field to given value.


### GetServer

`func (o *ClusterCreate) GetServer() ClusterServerCreate`

GetServer returns the Server field if non-nil, zero value otherwise.

### GetServerOk

`func (o *ClusterCreate) GetServerOk() (*ClusterServerCreate, bool)`

GetServerOk returns a tuple with the Server field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServer

`func (o *ClusterCreate) SetServer(v ClusterServerCreate)`

SetServer sets Server field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


