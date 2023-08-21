# MorpheusApi.NetworkPoolServerCreateInfoblox

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**type** | **String** | Type Code (Infoblox) | 
**name** | **String** | Name | 
**enabled** | **Boolean** | Can be used to enable / disable the network pool server. | [optional] [default to true]
**serviceUrl** | **String** | URL | 
**serviceUsername** | **String** | Username | [optional] 
**servicePassword** | **String** | Password | [optional] 
**serviceThrottleRate** | **Number** | Throttle Rate | [optional] [default to 0]
**ignoreSsl** | **Boolean** | Disable SSL SNI Verification | [optional] [default to true]
**networkFilter** | **String** | Network Filter | [optional] 
**zoneFilter** | **String** | Zone Filter | [optional] 
**tenantMatch** | **String** | Tenant Match | [optional] 
**serviceMode** | **String** | IP Mode | [optional] [default to &#39;static&#39;]
**config** | [**NetworkPoolServerCreateInfobloxConfig**](NetworkPoolServerCreateInfobloxConfig.md) |  | [optional] 
**credential** | [**NetworkPoolServerCreateBluecatCredential**](NetworkPoolServerCreateBluecatCredential.md) |  | [optional] 



## Enum: TypeEnum


* `infoblox` (value: `"infoblox"`)





## Enum: ServiceModeEnum


* `static` (value: `"static"`)

* `dhcp` (value: `"dhcp"`)




