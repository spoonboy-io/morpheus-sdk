# ApiMonitoringContactsContact

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Unique name scoped to your account for the contact | 
**EmailAddress** | Pointer to **string** | Email notification address | [optional] 
**SmsAddress** | Pointer to **string** | SMS notification address | [optional] 
**SlackHook** | Pointer to **string** | Slack Hook | [optional] 

## Methods

### NewApiMonitoringContactsContact

`func NewApiMonitoringContactsContact(name string, ) *ApiMonitoringContactsContact`

NewApiMonitoringContactsContact instantiates a new ApiMonitoringContactsContact object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiMonitoringContactsContactWithDefaults

`func NewApiMonitoringContactsContactWithDefaults() *ApiMonitoringContactsContact`

NewApiMonitoringContactsContactWithDefaults instantiates a new ApiMonitoringContactsContact object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *ApiMonitoringContactsContact) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ApiMonitoringContactsContact) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ApiMonitoringContactsContact) SetName(v string)`

SetName sets Name field to given value.


### GetEmailAddress

`func (o *ApiMonitoringContactsContact) GetEmailAddress() string`

GetEmailAddress returns the EmailAddress field if non-nil, zero value otherwise.

### GetEmailAddressOk

`func (o *ApiMonitoringContactsContact) GetEmailAddressOk() (*string, bool)`

GetEmailAddressOk returns a tuple with the EmailAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailAddress

`func (o *ApiMonitoringContactsContact) SetEmailAddress(v string)`

SetEmailAddress sets EmailAddress field to given value.

### HasEmailAddress

`func (o *ApiMonitoringContactsContact) HasEmailAddress() bool`

HasEmailAddress returns a boolean if a field has been set.

### GetSmsAddress

`func (o *ApiMonitoringContactsContact) GetSmsAddress() string`

GetSmsAddress returns the SmsAddress field if non-nil, zero value otherwise.

### GetSmsAddressOk

`func (o *ApiMonitoringContactsContact) GetSmsAddressOk() (*string, bool)`

GetSmsAddressOk returns a tuple with the SmsAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmsAddress

`func (o *ApiMonitoringContactsContact) SetSmsAddress(v string)`

SetSmsAddress sets SmsAddress field to given value.

### HasSmsAddress

`func (o *ApiMonitoringContactsContact) HasSmsAddress() bool`

HasSmsAddress returns a boolean if a field has been set.

### GetSlackHook

`func (o *ApiMonitoringContactsContact) GetSlackHook() string`

GetSlackHook returns the SlackHook field if non-nil, zero value otherwise.

### GetSlackHookOk

`func (o *ApiMonitoringContactsContact) GetSlackHookOk() (*string, bool)`

GetSlackHookOk returns a tuple with the SlackHook field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlackHook

`func (o *ApiMonitoringContactsContact) SetSlackHook(v string)`

SetSlackHook sets SlackHook field to given value.

### HasSlackHook

`func (o *ApiMonitoringContactsContact) HasSlackHook() bool`

HasSlackHook returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


