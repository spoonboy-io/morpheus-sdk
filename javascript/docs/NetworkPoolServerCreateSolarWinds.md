# MorpheusApi.NetworkPoolServerCreateSolarWinds

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**type** | **String** | Type Code (SolarWinds) | 
**name** | **String** | Name | 
**enabled** | **Boolean** | Can be used to enable / disable the network pool server. | [optional] [default to true]
**serviceUrl** | **String** | URL | 
**serviceUsername** | **String** | Username | [optional] 
**servicePassword** | **String** | Password | [optional] 
**serviceThrottleRate** | **Number** | Throttle Rate | [optional] [default to 0]
**ignoreSsl** | **Boolean** | Disable SSL SNI Verification | [optional] [default to true]
**config** | [**NetworkPoolServerCreateBluecatConfig**](NetworkPoolServerCreateBluecatConfig.md) |  | [optional] 
**credential** | [**NetworkPoolServerCreateBluecatCredential**](NetworkPoolServerCreateBluecatCredential.md) |  | [optional] 



## Enum: TypeEnum


* `solarwindsipam` (value: `"solarwindsipam"`)




