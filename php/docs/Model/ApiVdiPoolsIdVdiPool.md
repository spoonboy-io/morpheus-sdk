# # ApiVdiPoolsIdVdiPool

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**name** | **string** | Virtual Desktop name | [optional]
**description** | **string** | Virtual Desktop description | [optional]
**owner** | **int** | Owner (User) ID | [optional]
**min_idle** | **float** | Min Idle - Sets the minimum number of idle instances on standby in the pool. The pool will always try to maintain this number of available instances on standby. | [optional]
**initial_pool_size** | **float** | The initial size of the pool to be allocated on creation | [optional]
**max_idle** | **float** | Sets the maximum number of idle instances on standby in the pool. If the number of idle instances supersedes this, the pool will start removing instances | [optional]
**max_pool_size** | **float** | Max limit on number of allocations and instances within the pool. | [optional]
**allocation_timeout_minutes** | **float** | Time (in minutes) after a user disconnects before an allocation is recycled or shutdown depending on persistence. | [optional]
**persistent_user** | **bool** | Persistent Desktop Pool | [optional] [default to false]
**recyclable** | **bool** | Recyclable VDI Pools only work with cloud types that support snapshot management (i.e. Vmware, Nutanix, VCD) | [optional] [default to false]
**allow_copy** | **bool** | Allow copy from desktop | [optional] [default to false]
**allow_printer** | **bool** | Allow local printers from Desktop | [optional] [default to false]
**allow_fileshare** | **bool** | Allow File Share | [optional] [default to false]
**allow_hypervisor_console** | **bool** | Allow hypervisor console | [optional] [default to false]
**auto_create_local_user_on_reservation** | **bool** | Auto create local user upon reservation | [optional] [default to false]
**enabled** | **bool** | Can be used to enable or disable the VDI pool | [optional] [default to true]
**icon_path** | **string** | The relative location of an icon image | [optional]
**apps** | **int[]** | Array of VDI App IDs | [optional]
**gateway** | **int** | VDI Gateway ID | [optional]
**instance_config** | **string** | Instance Config JSON. Passing as a string will preserve property order.  See &#x60;config&#x60; object for required values. | [optional]
**config** | [**\OpenAPI\Client\Model\ApiVdiPoolsIdVdiPoolConfig**](ApiVdiPoolsIdVdiPoolConfig.md) |  | [optional]
**guest_console_jump_host** | **string** | Guest Console Jump Host | [optional]
**guest_console_jump_port** | **int** | Guest Console Jump Port | [optional]
**guest_console_jump_username** | **string** | Guest Console Jump Username | [optional]
**guest_console_jump_password** | **string** | Guest Console Jump Password | [optional]
**guest_console_jump_keypair** | **int** | Guest Console Jump Key Pair. see &#x60;Key Pair&#x60; | [optional]

[[Back to Model list]](../../README.md#models) [[Back to API list]](../../README.md#endpoints) [[Back to README]](../../README.md)
