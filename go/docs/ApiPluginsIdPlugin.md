# ApiPluginsIdPlugin

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enabled** | Pointer to **bool** | Can be used to disable the plugin | [optional] 
**Config** | Pointer to **map[string]interface{}** | Configuration object that contains settings for the applicable option types. | [optional] 

## Methods

### NewApiPluginsIdPlugin

`func NewApiPluginsIdPlugin() *ApiPluginsIdPlugin`

NewApiPluginsIdPlugin instantiates a new ApiPluginsIdPlugin object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiPluginsIdPluginWithDefaults

`func NewApiPluginsIdPluginWithDefaults() *ApiPluginsIdPlugin`

NewApiPluginsIdPluginWithDefaults instantiates a new ApiPluginsIdPlugin object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnabled

`func (o *ApiPluginsIdPlugin) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *ApiPluginsIdPlugin) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *ApiPluginsIdPlugin) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *ApiPluginsIdPlugin) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetConfig

`func (o *ApiPluginsIdPlugin) GetConfig() map[string]interface{}`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *ApiPluginsIdPlugin) GetConfigOk() (*map[string]interface{}, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *ApiPluginsIdPlugin) SetConfig(v map[string]interface{})`

SetConfig sets Config field to given value.

### HasConfig

`func (o *ApiPluginsIdPlugin) HasConfig() bool`

HasConfig returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


