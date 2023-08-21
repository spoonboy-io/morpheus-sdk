# MorpheusApi.NetworkUpdate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**displayName** | **String** | Display Name | [optional] 
**labels** | **[String]** | Array of label strings, can be used for filtering. | [optional] 
**description** | **String** | Description | [optional] 
**cidr** | **String** | CIDR Network | [optional] 
**gateway** | **String** | Network Gateway | [optional] 
**dnsPrimary** | **String** | Primary DNS Server | [optional] 
**dnsSecondary** | **String** | Secondary DNS Server | [optional] 
**vlanId** | **Number** |  | [optional] 
**pool** | **Number** | Network Pool ID | [optional] 
**allowStaticOverride** | **Boolean** | Allow IP Override | [optional] 
**assignPublicIp** | **Boolean** | Assign Public IP | [optional] 
**active** | **Boolean** | Activate (true) or disable (false) the network | [optional] 
**dhcpServer** | **Boolean** | DHCP Server enabled network | [optional] 
**networkDomain** | [**NetworkCreateNetworkDomain**](NetworkCreateNetworkDomain.md) |  | [optional] 
**searchDomains** | **String** | Search Domains | [optional] 
**networkProxy** | [**NetworkCreateNetworkProxy**](NetworkCreateNetworkProxy.md) |  | [optional] 
**applianceUrlProxyBypass** | **Boolean** | Bypass Proxy for Appliance URL | [optional] 
**noProxy** | **String** | Comma-separated list of ip addresses or name servers to exclude proxy traversal for. Typically locally routable servers are excluded. | [optional] 
**visibility** | **String** | Visibility, private or public. | [optional] [default to &#39;private&#39;]
**config** | **Object** | Configuration object. Settings vary by type. | [optional] 
**tenants** | [**[ApiBlueprintsIdUpdatePermissionsResourcePermissionSites]**](ApiBlueprintsIdUpdatePermissionsResourcePermissionSites.md) | Array of tenant account ids that are allowed access | [optional] 
**resourcePermissions** | [**NetworkCreateResourcePermissions**](NetworkCreateResourcePermissions.md) |  | [optional] 



## Enum: VisibilityEnum


* `private` (value: `"private"`)

* `public` (value: `"public"`)




