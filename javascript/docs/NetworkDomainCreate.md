# MorpheusApi.NetworkDomainCreate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**name** | **String** | Name | [optional] 
**description** | **String** | Description | [optional] 
**displayName** | **String** | Overrides displayed name in domain selection components. Useful if using many OU Paths. | [optional] 
**publicZone** | **Boolean** | Public Zone | [optional] [default to false]
**taskSetId** | **Number** | Workflow ID. Workflows can be applied to an instance when associated with a domain. Useful for custom domain related scripting. (Important if wanting to ensure the computer is removed from the domain using teardown phased workflows.)  | [optional] 
**active** | **Boolean** | Active | [optional] 
**domainController** | **Boolean** | Join Domain Controller | [optional] [default to true]
**domainUsername** | **String** | Domain Username | [optional] 
**domainPassword** | **String** | Domain Password | [optional] 
**dcServer** | **String** | DC Server. If specified, must be the server name and not an IP Address | [optional] 
**ouPath** | **String** | OU Path. (i.e. &#39;OU&#x3D;staging,DC&#x3D;ad,DC&#x3D;yourdomain,DC&#x3D;com&#39;) | [optional] 
**guestUsername** | **String** | Guest Username. If set, will change the instances RPC Service User after joining a Domain. | [optional] 
**guestPassword** | **String** | Guest Password | [optional] 


