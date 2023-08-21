# MorpheusApi.NetworkPoolServerUpdateSolarWinds

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**name** | **String** | Name | [optional] 
**enabled** | **Boolean** | Can be used to enable / disable the network pool server. | [optional] [default to true]
**serviceUrl** | **String** | URL | [optional] 
**serviceUsername** | **String** | Username | [optional] 
**servicePassword** | **String** | Password | [optional] 
**serviceThrottleRate** | **Number** | Throttle Rate | [optional] [default to 0]
**ignoreSsl** | **Boolean** | Disable SSL SNI Verification | [optional] 
**config** | [**NetworkPoolServerCreateBluecatConfig**](NetworkPoolServerCreateBluecatConfig.md) |  | [optional] 
**credential** | [**NetworkPoolServerCreateBluecatCredential**](NetworkPoolServerCreateBluecatCredential.md) |  | [optional] 


