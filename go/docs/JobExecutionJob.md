# JobExecutionJob

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **NullableString** |  | [optional] 
**Type** | Pointer to [**InlineResponse20079LoadBalancerMonitorLoadBalancerType**](inline_response_200_79_loadBalancerMonitor_loadBalancer_type.md) |  | [optional] 

## Methods

### NewJobExecutionJob

`func NewJobExecutionJob() *JobExecutionJob`

NewJobExecutionJob instantiates a new JobExecutionJob object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewJobExecutionJobWithDefaults

`func NewJobExecutionJobWithDefaults() *JobExecutionJob`

NewJobExecutionJobWithDefaults instantiates a new JobExecutionJob object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *JobExecutionJob) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *JobExecutionJob) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *JobExecutionJob) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *JobExecutionJob) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *JobExecutionJob) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *JobExecutionJob) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *JobExecutionJob) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *JobExecutionJob) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *JobExecutionJob) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *JobExecutionJob) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *JobExecutionJob) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *JobExecutionJob) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *JobExecutionJob) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *JobExecutionJob) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetType

`func (o *JobExecutionJob) GetType() InlineResponse20079LoadBalancerMonitorLoadBalancerType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *JobExecutionJob) GetTypeOk() (*InlineResponse20079LoadBalancerMonitorLoadBalancerType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *JobExecutionJob) SetType(v InlineResponse20079LoadBalancerMonitorLoadBalancerType)`

SetType sets Type field to given value.

### HasType

`func (o *JobExecutionJob) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


