# AddVDIPoolsRequestVdiPoolOneOf


## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**name** | **str** | Virtual Desktop name | 
**max_pool_size** | **float** | Max limit on number of allocations and instances within the pool.  | 
**instance_config** | **str** | Instance Config JSON. Passing as a string will preserve property order.  See &#x60;config&#x60; object for required values. | 
**description** | **str** | Virtual Desktop description | [optional] 
**owner** | **int** | Owner (User) ID | [optional] 
**min_idle** | **float** | Min Idle - Sets the minimum number of idle instances on standby in the pool. The pool will always try to maintain this number of available instances on standby.  | [optional] 
**initial_pool_size** | **float** | The initial size of the pool to be allocated on creation | [optional] 
**max_idle** | **float** | Sets the maximum number of idle instances on standby in the pool. If the number of idle instances supersedes this, the pool will start removing instances  | [optional] 
**allocation_timeout_minutes** | **float** | Time (in minutes) after a user disconnects before an allocation is recycled or shutdown depending on persistence.  | [optional] 
**persistent_user** | **bool** | Persistent Desktop Pool | [optional]  if omitted the server will use the default value of False
**recyclable** | **bool** | Recyclable VDI Pools only work with cloud types that support snapshot management (i.e. Vmware, Nutanix, VCD) | [optional]  if omitted the server will use the default value of False
**allow_copy** | **bool** | Allow copy from desktop | [optional]  if omitted the server will use the default value of False
**allow_printer** | **bool** | Allow local printers from Desktop | [optional]  if omitted the server will use the default value of False
**allow_fileshare** | **bool** | Allow File Share | [optional]  if omitted the server will use the default value of False
**allow_hypervisor_console** | **bool** | Allow hypervisor console | [optional]  if omitted the server will use the default value of False
**auto_create_local_user_on_reservation** | **bool** | Auto create local user upon reservation | [optional]  if omitted the server will use the default value of False
**enabled** | **bool** | Can be used to enable or disable the VDI pool | [optional]  if omitted the server will use the default value of True
**icon_path** | **str** | The relative location of an icon image | [optional] 
**apps** | **[int]** | Array of VDI App IDs | [optional] 
**gateway** | **int** | VDI Gateway ID | [optional] 
**config** | [**AddVDIPoolsRequestVdiPoolOneOfConfig**](AddVDIPoolsRequestVdiPoolOneOfConfig.md) |  | [optional] 
**guest_console_jump_host** | **str** | Guest Console Jump Host | [optional] 
**guest_console_jump_port** | **int** | Guest Console Jump Port | [optional] 
**guest_console_jump_username** | **str** | Guest Console Jump Username | [optional] 
**guest_console_jump_password** | **str** | Guest Console Jump Password | [optional] 
**guest_console_jump_keypair** | **int** | Guest Console Jump Key Pair. see &#x60;Key Pair&#x60; | [optional] 
**any string name** | **bool, date, datetime, dict, float, int, list, str, none_type** | any string name can be used but the value must be the correct type | [optional]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


