

# InlineObject66

The following parameters are available under the root context of the JSON body. The secret mount is capable of storing the entire JSON object as key=value pairs, or just a single string. To store a string instead, use the value query parameter instead of JSON, or pass type=string. There are a couple of special keys that the API will look for in the payload. The ttl key is a special key that if present in the payload will be parsed and used as the ttl parameter (lease duration in seconds). The value key is a special key that if present in the payload will be parsed and used as the secret data (instead of the entire payload). This is true when type=string. 
## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ttl** | [**OneOfstringlong**](OneOfstringlong.md) | Time to Live. The lease duration in seconds, or a human readable format eg. 15m, 8h, 7d. The default is 0 meaning Never expires. This only is applied if the cypher does not yet exist and is created.  |  [optional]
**value** | **String** | The secret value to be stored. This is only parsed if type is passed as &#x60;string&#x60;. |  [optional]



