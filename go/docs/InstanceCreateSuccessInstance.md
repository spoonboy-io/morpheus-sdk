# InstanceCreateSuccessInstance

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Name of the instance to be created. | 
**Site** | [**InstanceCreateInstanceSite**](instanceCreate_instance_site.md) |  | 
**InstanceType** | [**InstanceCreateInstanceInstanceType**](instanceCreate_instance_instanceType.md) |  | 
**Layout** | [**InstanceCreateInstanceLayout**](instanceCreate_instance_layout.md) |  | 
**Plan** | [**InstanceCreateInstancePlan**](instanceCreate_instance_plan.md) |  | 

## Methods

### NewInstanceCreateSuccessInstance

`func NewInstanceCreateSuccessInstance(name string, site InstanceCreateInstanceSite, instanceType InstanceCreateInstanceInstanceType, layout InstanceCreateInstanceLayout, plan InstanceCreateInstancePlan, ) *InstanceCreateSuccessInstance`

NewInstanceCreateSuccessInstance instantiates a new InstanceCreateSuccessInstance object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInstanceCreateSuccessInstanceWithDefaults

`func NewInstanceCreateSuccessInstanceWithDefaults() *InstanceCreateSuccessInstance`

NewInstanceCreateSuccessInstanceWithDefaults instantiates a new InstanceCreateSuccessInstance object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *InstanceCreateSuccessInstance) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *InstanceCreateSuccessInstance) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *InstanceCreateSuccessInstance) SetName(v string)`

SetName sets Name field to given value.


### GetSite

`func (o *InstanceCreateSuccessInstance) GetSite() InstanceCreateInstanceSite`

GetSite returns the Site field if non-nil, zero value otherwise.

### GetSiteOk

`func (o *InstanceCreateSuccessInstance) GetSiteOk() (*InstanceCreateInstanceSite, bool)`

GetSiteOk returns a tuple with the Site field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSite

`func (o *InstanceCreateSuccessInstance) SetSite(v InstanceCreateInstanceSite)`

SetSite sets Site field to given value.


### GetInstanceType

`func (o *InstanceCreateSuccessInstance) GetInstanceType() InstanceCreateInstanceInstanceType`

GetInstanceType returns the InstanceType field if non-nil, zero value otherwise.

### GetInstanceTypeOk

`func (o *InstanceCreateSuccessInstance) GetInstanceTypeOk() (*InstanceCreateInstanceInstanceType, bool)`

GetInstanceTypeOk returns a tuple with the InstanceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceType

`func (o *InstanceCreateSuccessInstance) SetInstanceType(v InstanceCreateInstanceInstanceType)`

SetInstanceType sets InstanceType field to given value.


### GetLayout

`func (o *InstanceCreateSuccessInstance) GetLayout() InstanceCreateInstanceLayout`

GetLayout returns the Layout field if non-nil, zero value otherwise.

### GetLayoutOk

`func (o *InstanceCreateSuccessInstance) GetLayoutOk() (*InstanceCreateInstanceLayout, bool)`

GetLayoutOk returns a tuple with the Layout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLayout

`func (o *InstanceCreateSuccessInstance) SetLayout(v InstanceCreateInstanceLayout)`

SetLayout sets Layout field to given value.


### GetPlan

`func (o *InstanceCreateSuccessInstance) GetPlan() InstanceCreateInstancePlan`

GetPlan returns the Plan field if non-nil, zero value otherwise.

### GetPlanOk

`func (o *InstanceCreateSuccessInstance) GetPlanOk() (*InstanceCreateInstancePlan, bool)`

GetPlanOk returns a tuple with the Plan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlan

`func (o *InstanceCreateSuccessInstance) SetPlan(v InstanceCreateInstancePlan)`

SetPlan sets Plan field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


