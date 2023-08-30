# NetworkDomainCreate


## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**name** | **str** | Name | [optional] 
**description** | **str** | Description | [optional] 
**display_name** | **str** | Overrides displayed name in domain selection components. Useful if using many OU Paths. | [optional] 
**public_zone** | **bool** | Public Zone | [optional]  if omitted the server will use the default value of False
**task_set_id** | **int** | Workflow ID. Workflows can be applied to an instance when associated with a domain. Useful for custom domain related scripting. (Important if wanting to ensure the computer is removed from the domain using teardown phased workflows.)  | [optional] 
**active** | **bool** | Active | [optional] 
**domain_controller** | **bool** | Join Domain Controller | [optional]  if omitted the server will use the default value of True
**domain_username** | **str** | Domain Username | [optional] 
**domain_password** | **str** | Domain Password | [optional] 
**dc_server** | **str** | DC Server. If specified, must be the server name and not an IP Address | [optional] 
**ou_path** | **str** | OU Path. (i.e. &#39;OU&#x3D;staging,DC&#x3D;ad,DC&#x3D;yourdomain,DC&#x3D;com&#39;) | [optional] 
**guest_username** | **str** | Guest Username. If set, will change the instances RPC Service User after joining a Domain. | [optional] 
**guest_password** | **str** | Guest Password | [optional] 
**any string name** | **bool, date, datetime, dict, float, int, list, str, none_type** | any string name can be used but the value must be the correct type | [optional]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


