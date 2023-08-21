# MorpheusApi.LoadBalancerUpdate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**name** | **String** | Name | [optional] 
**description** | **String** | Description | [optional] 
**enabled** | **Boolean** | Activate (true) or disable (false) | [optional] 
**config** | **Object** | Configuration object with parameters that vary by load balancer type. | [optional] 
**visibility** | **String** | private or public | [optional] [default to &#39;public&#39;]
**tenants** | [**[LoadBalancerCreateTenants]**](LoadBalancerCreateTenants.md) | Array of tenant account ids that are allowed access | [optional] 
**resourcePermission** | [**LoadBalancerCreateResourcePermission**](LoadBalancerCreateResourcePermission.md) |  | [optional] 


