# GetStaticRoutes200ResponseNetworkRoutesInner
## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **Int32** |  | [optional] 
**Name** | **String** |  | [optional] 
**Code** | **String** |  | [optional] 
**Description** | **String** |  | [optional] 
**Priority** | **String** |  | [optional] 
**RouteType** | **String** |  | [optional] 
**Source** | **String** |  | [optional] 
**SourceType** | **String** |  | [optional] 
**Destination** | **String** |  | [optional] 
**DestinationType** | **String** |  | [optional] 
**DefaultRoute** | **Boolean** |  | [optional] 
**NetworkMtu** | **String** |  | [optional] 
**ExternalInterface** | **String** |  | [optional] 
**InternalId** | **String** |  | [optional] 
**UniqueId** | **String** |  | [optional] 
**ExternalType** | **String** |  | [optional] 
**Enabled** | **Boolean** |  | [optional] 
**Visible** | **Boolean** |  | [optional] 

## Examples

- Prepare the resource
```powershell
$GetStaticRoutes200ResponseNetworkRoutesInner = Initialize-PSOpenAPIToolsGetStaticRoutes200ResponseNetworkRoutesInner  -Id null `
 -Name null `
 -Code null `
 -Description null `
 -Priority null `
 -RouteType null `
 -Source null `
 -SourceType null `
 -Destination null `
 -DestinationType null `
 -DefaultRoute null `
 -NetworkMtu null `
 -ExternalInterface null `
 -InternalId null `
 -UniqueId null `
 -ExternalType null `
 -Enabled null `
 -Visible null
```

- Convert the resource to JSON
```powershell
$GetStaticRoutes200ResponseNetworkRoutesInner | ConvertTo-JSON
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

