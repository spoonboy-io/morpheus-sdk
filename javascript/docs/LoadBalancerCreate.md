# MorpheusApi.LoadBalancerCreate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**name** | **String** | Name | [optional] 
**description** | **String** | Description | [optional] 
**networkServerId** | **Number** | Network Server ID | [optional] 
**config** | **Object** | Configuration object with parameters that vary by load balancer type. | [optional] 
**visibility** | **String** | private or public | [optional] [default to &#39;public&#39;]
**tenants** | [**[LoadBalancerCreateTenants]**](LoadBalancerCreateTenants.md) | Array of tenant account ids that are allowed access | [optional] 
**resourcePermission** | [**LoadBalancerCreateResourcePermission**](LoadBalancerCreateResourcePermission.md) |  | [optional] 


