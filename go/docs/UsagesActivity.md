# UsagesActivity

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Success** | Pointer to **bool** |  | [optional] 
**ActivityType** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Message** | Pointer to **string** |  | [optional] 
**ObjectType** | Pointer to **string** |  | [optional] 
**ObjectId** | Pointer to **int64** |  | [optional] 
**User** | Pointer to [**NullableInlineResponse20083LoadBalancerNodeCreatedBy**](inline_response_200_83_loadBalancerNode_createdBy.md) |  | [optional] 
**Ts** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewUsagesActivity

`func NewUsagesActivity() *UsagesActivity`

NewUsagesActivity instantiates a new UsagesActivity object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUsagesActivityWithDefaults

`func NewUsagesActivityWithDefaults() *UsagesActivity`

NewUsagesActivityWithDefaults instantiates a new UsagesActivity object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *UsagesActivity) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *UsagesActivity) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *UsagesActivity) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *UsagesActivity) HasId() bool`

HasId returns a boolean if a field has been set.

### GetSuccess

`func (o *UsagesActivity) GetSuccess() bool`

GetSuccess returns the Success field if non-nil, zero value otherwise.

### GetSuccessOk

`func (o *UsagesActivity) GetSuccessOk() (*bool, bool)`

GetSuccessOk returns a tuple with the Success field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccess

`func (o *UsagesActivity) SetSuccess(v bool)`

SetSuccess sets Success field to given value.

### HasSuccess

`func (o *UsagesActivity) HasSuccess() bool`

HasSuccess returns a boolean if a field has been set.

### GetActivityType

`func (o *UsagesActivity) GetActivityType() string`

GetActivityType returns the ActivityType field if non-nil, zero value otherwise.

### GetActivityTypeOk

`func (o *UsagesActivity) GetActivityTypeOk() (*string, bool)`

GetActivityTypeOk returns a tuple with the ActivityType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActivityType

`func (o *UsagesActivity) SetActivityType(v string)`

SetActivityType sets ActivityType field to given value.

### HasActivityType

`func (o *UsagesActivity) HasActivityType() bool`

HasActivityType returns a boolean if a field has been set.

### GetName

`func (o *UsagesActivity) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UsagesActivity) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UsagesActivity) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *UsagesActivity) HasName() bool`

HasName returns a boolean if a field has been set.

### GetMessage

`func (o *UsagesActivity) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *UsagesActivity) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *UsagesActivity) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *UsagesActivity) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetObjectType

`func (o *UsagesActivity) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *UsagesActivity) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *UsagesActivity) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.

### HasObjectType

`func (o *UsagesActivity) HasObjectType() bool`

HasObjectType returns a boolean if a field has been set.

### GetObjectId

`func (o *UsagesActivity) GetObjectId() int64`

GetObjectId returns the ObjectId field if non-nil, zero value otherwise.

### GetObjectIdOk

`func (o *UsagesActivity) GetObjectIdOk() (*int64, bool)`

GetObjectIdOk returns a tuple with the ObjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectId

`func (o *UsagesActivity) SetObjectId(v int64)`

SetObjectId sets ObjectId field to given value.

### HasObjectId

`func (o *UsagesActivity) HasObjectId() bool`

HasObjectId returns a boolean if a field has been set.

### GetUser

`func (o *UsagesActivity) GetUser() InlineResponse20083LoadBalancerNodeCreatedBy`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *UsagesActivity) GetUserOk() (*InlineResponse20083LoadBalancerNodeCreatedBy, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *UsagesActivity) SetUser(v InlineResponse20083LoadBalancerNodeCreatedBy)`

SetUser sets User field to given value.

### HasUser

`func (o *UsagesActivity) HasUser() bool`

HasUser returns a boolean if a field has been set.

### SetUserNil

`func (o *UsagesActivity) SetUserNil(b bool)`

 SetUserNil sets the value for User to be an explicit nil

### UnsetUser
`func (o *UsagesActivity) UnsetUser()`

UnsetUser ensures that no value is present for User, not even an explicit nil
### GetTs

`func (o *UsagesActivity) GetTs() time.Time`

GetTs returns the Ts field if non-nil, zero value otherwise.

### GetTsOk

`func (o *UsagesActivity) GetTsOk() (*time.Time, bool)`

GetTsOk returns a tuple with the Ts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTs

`func (o *UsagesActivity) SetTs(v time.Time)`

SetTs sets Ts field to given value.

### HasTs

`func (o *UsagesActivity) HasTs() bool`

HasTs returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


