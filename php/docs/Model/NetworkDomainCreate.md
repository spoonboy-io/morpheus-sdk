# # NetworkDomainCreate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**name** | **string** | Name | [optional]
**description** | **string** | Description | [optional]
**display_name** | **string** | Overrides displayed name in domain selection components. Useful if using many OU Paths. | [optional]
**public_zone** | **bool** | Public Zone | [optional] [default to false]
**task_set_id** | **int** | Workflow ID. Workflows can be applied to an instance when associated with a domain. Useful for custom domain related scripting. (Important if wanting to ensure the computer is removed from the domain using teardown phased workflows.) | [optional]
**active** | **bool** | Active | [optional]
**domain_controller** | **bool** | Join Domain Controller | [optional] [default to true]
**domain_username** | **string** | Domain Username | [optional]
**domain_password** | **string** | Domain Password | [optional]
**dc_server** | **string** | DC Server. If specified, must be the server name and not an IP Address | [optional]
**ou_path** | **string** | OU Path. (i.e. &#39;OU&#x3D;staging,DC&#x3D;ad,DC&#x3D;yourdomain,DC&#x3D;com&#39;) | [optional]
**guest_username** | **string** | Guest Username. If set, will change the instances RPC Service User after joining a Domain. | [optional]
**guest_password** | **string** | Guest Password | [optional]

[[Back to Model list]](../../README.md#models) [[Back to API list]](../../README.md#endpoints) [[Back to README]](../../README.md)
