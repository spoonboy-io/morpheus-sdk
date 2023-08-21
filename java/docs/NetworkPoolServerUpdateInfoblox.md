

# NetworkPoolServerUpdateInfoblox

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**name** | **String** | Name |  [optional]
**enabled** | **Boolean** | Can be used to enable / disable the network pool server. |  [optional]
**serviceUrl** | **String** | URL |  [optional]
**serviceUsername** | **String** | Username |  [optional]
**servicePassword** | **String** | Password |  [optional]
**serviceThrottleRate** | **Long** | Throttle Rate |  [optional]
**ignoreSsl** | **Boolean** | Disable SSL SNI Verification |  [optional]
**networkFilter** | **String** | Network Filter |  [optional]
**zoneFilter** | **String** | Zone Filter |  [optional]
**tenantMatch** | **String** | Tenant Match |  [optional]
**serviceMode** | [**ServiceModeEnum**](#ServiceModeEnum) | IP Mode |  [optional]
**config** | [**NetworkPoolServerCreateInfobloxConfig**](NetworkPoolServerCreateInfobloxConfig.md) |  |  [optional]
**credential** | [**NetworkPoolServerCreateBluecatCredential**](NetworkPoolServerCreateBluecatCredential.md) |  |  [optional]



## Enum: ServiceModeEnum

Name | Value
---- | -----
STATIC | &quot;static&quot;
DHCP | &quot;dhcp&quot;



