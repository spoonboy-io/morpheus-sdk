# BillingInstancesVolumes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Size** | Pointer to **int64** |  | [optional] 
**TypeCode** | Pointer to **string** |  | [optional] 
**Datastore** | Pointer to [**BillingInstancesDatastore**](billingInstances_datastore.md) |  | [optional] 

## Methods

### NewBillingInstancesVolumes

`func NewBillingInstancesVolumes() *BillingInstancesVolumes`

NewBillingInstancesVolumes instantiates a new BillingInstancesVolumes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBillingInstancesVolumesWithDefaults

`func NewBillingInstancesVolumesWithDefaults() *BillingInstancesVolumes`

NewBillingInstancesVolumesWithDefaults instantiates a new BillingInstancesVolumes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSize

`func (o *BillingInstancesVolumes) GetSize() int64`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *BillingInstancesVolumes) GetSizeOk() (*int64, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *BillingInstancesVolumes) SetSize(v int64)`

SetSize sets Size field to given value.

### HasSize

`func (o *BillingInstancesVolumes) HasSize() bool`

HasSize returns a boolean if a field has been set.

### GetTypeCode

`func (o *BillingInstancesVolumes) GetTypeCode() string`

GetTypeCode returns the TypeCode field if non-nil, zero value otherwise.

### GetTypeCodeOk

`func (o *BillingInstancesVolumes) GetTypeCodeOk() (*string, bool)`

GetTypeCodeOk returns a tuple with the TypeCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTypeCode

`func (o *BillingInstancesVolumes) SetTypeCode(v string)`

SetTypeCode sets TypeCode field to given value.

### HasTypeCode

`func (o *BillingInstancesVolumes) HasTypeCode() bool`

HasTypeCode returns a boolean if a field has been set.

### GetDatastore

`func (o *BillingInstancesVolumes) GetDatastore() BillingInstancesDatastore`

GetDatastore returns the Datastore field if non-nil, zero value otherwise.

### GetDatastoreOk

`func (o *BillingInstancesVolumes) GetDatastoreOk() (*BillingInstancesDatastore, bool)`

GetDatastoreOk returns a tuple with the Datastore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatastore

`func (o *BillingInstancesVolumes) SetDatastore(v BillingInstancesDatastore)`

SetDatastore sets Datastore field to given value.

### HasDatastore

`func (o *BillingInstancesVolumes) HasDatastore() bool`

HasDatastore returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


