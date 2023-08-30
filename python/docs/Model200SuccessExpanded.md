# Model200SuccessExpanded

A response object

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**success** | **bool** | Success indicator, true when the request succeeded and false when an error occurred | [optional]  if omitted the server will use the default value of True
**msg** | **str, none_type** | Message containing a description of the result, usually a message about the error that occurred | [optional] 
**errors** | **{str: (bool, date, datetime, dict, float, int, list, str, none_type)}, none_type** | Validation errors, with a key for Object containing error messages for each invalid parameter (key) | [optional] 
**any string name** | **bool, date, datetime, dict, float, int, list, str, none_type** | any string name can be used but the value must be the correct type | [optional]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


