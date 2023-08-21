

# NetworkPoolServerCreateBluecat

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**type** | [**TypeEnum**](#TypeEnum) | Type Code (Bluecat) | 
**name** | **String** | Name | 
**enabled** | **Boolean** | Can be used to enable / disable the network pool server. |  [optional]
**serviceUrl** | **String** | URL | 
**serviceUsername** | **String** | Username |  [optional]
**servicePassword** | **String** | Password |  [optional]
**serviceThrottleRate** | **Long** | Throttle Rate |  [optional]
**ignoreSsl** | **Boolean** | Disable SSL SNI Verification |  [optional]
**networkFilter** | **String** | Network Filter |  [optional]
**config** | [**NetworkPoolServerCreateBluecatConfig**](NetworkPoolServerCreateBluecatConfig.md) |  |  [optional]
**credential** | [**NetworkPoolServerCreateBluecatCredential**](NetworkPoolServerCreateBluecatCredential.md) |  |  [optional]



## Enum: TypeEnum

Name | Value
---- | -----
BLUECAT | &quot;bluecat&quot;



