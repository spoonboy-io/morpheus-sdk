# MetaObject

Metadata about the returned array of results, provides pagination information.

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**offset** | **int** | Offset records, the number of records to skip, can be used to paginate over results. | [optional]  if omitted the server will use the default value of 0
**max** | **int** | Max size, the maximum number of records to include in the response. | [optional]  if omitted the server will use the default value of 25
**size** | **int** | Number of records returned in the response | [optional]  if omitted the server will use the default value of 0
**total** | **int** | Total number of records found | [optional]  if omitted the server will use the default value of 0
**any string name** | **bool, date, datetime, dict, float, int, list, str, none_type** | any string name can be used but the value must be the correct type | [optional]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


