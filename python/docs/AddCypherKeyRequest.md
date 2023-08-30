# AddCypherKeyRequest

The following parameters are available under the root context of the JSON body. The secret mount is capable of storing the entire JSON object as key=value pairs, or just a single string. To store a string instead, use the value query parameter instead of JSON, or pass type=string. There are a couple of special keys that the API will look for in the payload. The ttl key is a special key that if present in the payload will be parsed and used as the ttl parameter (lease duration in seconds). The value key is a special key that if present in the payload will be parsed and used as the secret data (instead of the entire payload). This is true when type=string. 

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ttl** | [**AddCypherKeyRequestTtl**](AddCypherKeyRequestTtl.md) |  | [optional] 
**value** | **str** | The secret value to be stored. This is only parsed if type is passed as &#x60;string&#x60;. | [optional] 
**any string name** | **bool, date, datetime, dict, float, int, list, str, none_type** | any string name can be used but the value must be the correct type | [optional]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


