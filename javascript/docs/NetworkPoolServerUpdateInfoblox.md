# MorpheusApi.NetworkPoolServerUpdateInfoblox

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
**networkFilter** | **String** | Network Filter | [optional] 
**zoneFilter** | **String** | Zone Filter | [optional] 
**tenantMatch** | **String** | Tenant Match | [optional] 
**serviceMode** | **String** | IP Mode | [optional] [default to &#39;static&#39;]
**config** | [**NetworkPoolServerCreateInfobloxConfig**](NetworkPoolServerCreateInfobloxConfig.md) |  | [optional] 
**credential** | [**NetworkPoolServerCreateBluecatCredential**](NetworkPoolServerCreateBluecatCredential.md) |  | [optional] 



## Enum: ServiceModeEnum


* `static` (value: `"static"`)

* `dhcp` (value: `"dhcp"`)




