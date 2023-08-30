# UpdateHostCloudRequest


## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**cloud_id** | **int** | The cloud/zone ID we are moving the set of servers to | [optional] 
**servers** | [**[UpdateHostCloudRequestServersInner]**](UpdateHostCloudRequestServersInner.md) | A JSON array of source: and target: server ids to be moved. If the target is blank Morpheus will automatically try to match by the servers unique or externalId | [optional] 
**any string name** | **bool, date, datetime, dict, float, int, list, str, none_type** | any string name can be used but the value must be the correct type | [optional]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


