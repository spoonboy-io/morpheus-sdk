# InstanceCreateInstance

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Name of the instance to be created. | 
**Site** | [**InstanceCreateInstanceSite**](instanceCreate_instance_site.md) |  | 
**InstanceType** | [**InstanceCreateInstanceInstanceType**](instanceCreate_instance_instanceType.md) |  | 
**Layout** | [**InstanceCreateInstanceLayout**](instanceCreate_instance_layout.md) |  | 
**Plan** | [**InstanceCreateInstancePlan**](instanceCreate_instance_plan.md) |  | 
**InstanceContext** | Pointer to **string** | Environment | [optional] 

## Methods

### NewInstanceCreateInstance

`func NewInstanceCreateInstance(name string, site InstanceCreateInstanceSite, instanceType InstanceCreateInstanceInstanceType, layout InstanceCreateInstanceLayout, plan InstanceCreateInstancePlan, ) *InstanceCreateInstance`

NewInstanceCreateInstance instantiates a new InstanceCreateInstance object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInstanceCreateInstanceWithDefaults

`func NewInstanceCreateInstanceWithDefaults() *InstanceCreateInstance`

NewInstanceCreateInstanceWithDefaults instantiates a new InstanceCreateInstance object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *InstanceCreateInstance) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *InstanceCreateInstance) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *InstanceCreateInstance) SetName(v string)`

SetName sets Name field to given value.


### GetSite

`func (o *InstanceCreateInstance) GetSite() InstanceCreateInstanceSite`

GetSite returns the Site field if non-nil, zero value otherwise.

### GetSiteOk

`func (o *InstanceCreateInstance) GetSiteOk() (*InstanceCreateInstanceSite, bool)`

GetSiteOk returns a tuple with the Site field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSite

`func (o *InstanceCreateInstance) SetSite(v InstanceCreateInstanceSite)`

SetSite sets Site field to given value.


### GetInstanceType

`func (o *InstanceCreateInstance) GetInstanceType() InstanceCreateInstanceInstanceType`

GetInstanceType returns the InstanceType field if non-nil, zero value otherwise.

### GetInstanceTypeOk

`func (o *InstanceCreateInstance) GetInstanceTypeOk() (*InstanceCreateInstanceInstanceType, bool)`

GetInstanceTypeOk returns a tuple with the InstanceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceType

`func (o *InstanceCreateInstance) SetInstanceType(v InstanceCreateInstanceInstanceType)`

SetInstanceType sets InstanceType field to given value.


### GetLayout

`func (o *InstanceCreateInstance) GetLayout() InstanceCreateInstanceLayout`

GetLayout returns the Layout field if non-nil, zero value otherwise.

### GetLayoutOk

`func (o *InstanceCreateInstance) GetLayoutOk() (*InstanceCreateInstanceLayout, bool)`

GetLayoutOk returns a tuple with the Layout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLayout

`func (o *InstanceCreateInstance) SetLayout(v InstanceCreateInstanceLayout)`

SetLayout sets Layout field to given value.


### GetPlan

`func (o *InstanceCreateInstance) GetPlan() InstanceCreateInstancePlan`

GetPlan returns the Plan field if non-nil, zero value otherwise.

### GetPlanOk

`func (o *InstanceCreateInstance) GetPlanOk() (*InstanceCreateInstancePlan, bool)`

GetPlanOk returns a tuple with the Plan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlan

`func (o *InstanceCreateInstance) SetPlan(v InstanceCreateInstancePlan)`

SetPlan sets Plan field to given value.


### GetInstanceContext

`func (o *InstanceCreateInstance) GetInstanceContext() string`

GetInstanceContext returns the InstanceContext field if non-nil, zero value otherwise.

### GetInstanceContextOk

`func (o *InstanceCreateInstance) GetInstanceContextOk() (*string, bool)`

GetInstanceContextOk returns a tuple with the InstanceContext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceContext

`func (o *InstanceCreateInstance) SetInstanceContext(v string)`

SetInstanceContext sets InstanceContext field to given value.

### HasInstanceContext

`func (o *InstanceCreateInstance) HasInstanceContext() bool`

HasInstanceContext returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


