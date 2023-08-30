# IntegrationConfigIntegration


## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**name** | **str** | Name, a unique identifier for the integration | 
**type** | **str** | Integration Type Code | 
**enabled** | **bool** | Set &#x60;true&#x60; to enable integration | [optional] 
**refresh** | **bool** | Pass &#x60;false&#x60; to skip refresh.  By default, refresh is done on update, when it is supported by the integration type.  | [optional]  if omitted the server will use the default value of True
**credential** | [**UpdateTasksRequestTaskCredential**](UpdateTasksRequestTaskCredential.md) |  | [optional] 
**any string name** | **bool, date, datetime, dict, float, int, list, str, none_type** | any string name can be used but the value must be the correct type | [optional]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


