# InlineResponse20083LoadBalancerNode

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Visibility** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **NullableString** |  | [optional] 
**IpAddress** | Pointer to **string** |  | [optional] 
**Port** | Pointer to **int64** |  | [optional] 
**PortType** | Pointer to **NullableString** |  | [optional] 
**MonitorPort** | Pointer to **NullableString** |  | [optional] 
**Weight** | Pointer to **NullableInt64** |  | [optional] 
**NodeState** | Pointer to **NullableString** |  | [optional] 
**InternalId** | Pointer to **NullableString** |  | [optional] 
**ExternalId** | Pointer to **NullableString** |  | [optional] 
**Enabled** | Pointer to **bool** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**StatusMessage** | Pointer to **NullableString** |  | [optional] 
**StatusDate** | Pointer to **NullableTime** |  | [optional] 
**Server** | Pointer to [**NullableInlineResponse20082LoadBalancerInstanceSslCert**](inline_response_200_82_loadBalancerInstance_sslCert.md) |  | [optional] 
**InstanceId** | Pointer to **NullableInt64** |  | [optional] 
**ContainerId** | Pointer to **NullableInt64** |  | [optional] 
**NodeSource** | Pointer to **NullableString** |  | [optional] 
**Monitor** | Pointer to [**NullableInlineResponse20082LoadBalancerInstanceSslCert**](inline_response_200_82_loadBalancerInstance_sslCert.md) |  | [optional] 
**MaxConnections** | Pointer to **NullableInt64** |  | [optional] 
**ExternalRefType** | Pointer to **NullableString** |  | [optional] 
**ExternalRefId** | Pointer to **NullableString** |  | [optional] 
**ExternalRefName** | Pointer to **NullableString** |  | [optional] 
**CreatedBy** | Pointer to [**NullableInlineResponse20083LoadBalancerNodeCreatedBy**](inline_response_200_83_loadBalancerNode_createdBy.md) |  | [optional] 
**DateCreated** | Pointer to **time.Time** |  | [optional] 
**LastUpdated** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewInlineResponse20083LoadBalancerNode

`func NewInlineResponse20083LoadBalancerNode() *InlineResponse20083LoadBalancerNode`

NewInlineResponse20083LoadBalancerNode instantiates a new InlineResponse20083LoadBalancerNode object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse20083LoadBalancerNodeWithDefaults

`func NewInlineResponse20083LoadBalancerNodeWithDefaults() *InlineResponse20083LoadBalancerNode`

NewInlineResponse20083LoadBalancerNodeWithDefaults instantiates a new InlineResponse20083LoadBalancerNode object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *InlineResponse20083LoadBalancerNode) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *InlineResponse20083LoadBalancerNode) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *InlineResponse20083LoadBalancerNode) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *InlineResponse20083LoadBalancerNode) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *InlineResponse20083LoadBalancerNode) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *InlineResponse20083LoadBalancerNode) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *InlineResponse20083LoadBalancerNode) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *InlineResponse20083LoadBalancerNode) HasName() bool`

HasName returns a boolean if a field has been set.

### GetVisibility

`func (o *InlineResponse20083LoadBalancerNode) GetVisibility() string`

GetVisibility returns the Visibility field if non-nil, zero value otherwise.

### GetVisibilityOk

`func (o *InlineResponse20083LoadBalancerNode) GetVisibilityOk() (*string, bool)`

GetVisibilityOk returns a tuple with the Visibility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisibility

`func (o *InlineResponse20083LoadBalancerNode) SetVisibility(v string)`

SetVisibility sets Visibility field to given value.

### HasVisibility

`func (o *InlineResponse20083LoadBalancerNode) HasVisibility() bool`

HasVisibility returns a boolean if a field has been set.

### GetDescription

`func (o *InlineResponse20083LoadBalancerNode) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *InlineResponse20083LoadBalancerNode) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *InlineResponse20083LoadBalancerNode) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *InlineResponse20083LoadBalancerNode) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *InlineResponse20083LoadBalancerNode) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *InlineResponse20083LoadBalancerNode) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetIpAddress

`func (o *InlineResponse20083LoadBalancerNode) GetIpAddress() string`

GetIpAddress returns the IpAddress field if non-nil, zero value otherwise.

### GetIpAddressOk

`func (o *InlineResponse20083LoadBalancerNode) GetIpAddressOk() (*string, bool)`

GetIpAddressOk returns a tuple with the IpAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpAddress

`func (o *InlineResponse20083LoadBalancerNode) SetIpAddress(v string)`

SetIpAddress sets IpAddress field to given value.

### HasIpAddress

`func (o *InlineResponse20083LoadBalancerNode) HasIpAddress() bool`

HasIpAddress returns a boolean if a field has been set.

### GetPort

`func (o *InlineResponse20083LoadBalancerNode) GetPort() int64`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *InlineResponse20083LoadBalancerNode) GetPortOk() (*int64, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *InlineResponse20083LoadBalancerNode) SetPort(v int64)`

SetPort sets Port field to given value.

### HasPort

`func (o *InlineResponse20083LoadBalancerNode) HasPort() bool`

HasPort returns a boolean if a field has been set.

### GetPortType

`func (o *InlineResponse20083LoadBalancerNode) GetPortType() string`

GetPortType returns the PortType field if non-nil, zero value otherwise.

### GetPortTypeOk

`func (o *InlineResponse20083LoadBalancerNode) GetPortTypeOk() (*string, bool)`

GetPortTypeOk returns a tuple with the PortType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortType

`func (o *InlineResponse20083LoadBalancerNode) SetPortType(v string)`

SetPortType sets PortType field to given value.

### HasPortType

`func (o *InlineResponse20083LoadBalancerNode) HasPortType() bool`

HasPortType returns a boolean if a field has been set.

### SetPortTypeNil

`func (o *InlineResponse20083LoadBalancerNode) SetPortTypeNil(b bool)`

 SetPortTypeNil sets the value for PortType to be an explicit nil

### UnsetPortType
`func (o *InlineResponse20083LoadBalancerNode) UnsetPortType()`

UnsetPortType ensures that no value is present for PortType, not even an explicit nil
### GetMonitorPort

`func (o *InlineResponse20083LoadBalancerNode) GetMonitorPort() string`

GetMonitorPort returns the MonitorPort field if non-nil, zero value otherwise.

### GetMonitorPortOk

`func (o *InlineResponse20083LoadBalancerNode) GetMonitorPortOk() (*string, bool)`

GetMonitorPortOk returns a tuple with the MonitorPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonitorPort

`func (o *InlineResponse20083LoadBalancerNode) SetMonitorPort(v string)`

SetMonitorPort sets MonitorPort field to given value.

### HasMonitorPort

`func (o *InlineResponse20083LoadBalancerNode) HasMonitorPort() bool`

HasMonitorPort returns a boolean if a field has been set.

### SetMonitorPortNil

`func (o *InlineResponse20083LoadBalancerNode) SetMonitorPortNil(b bool)`

 SetMonitorPortNil sets the value for MonitorPort to be an explicit nil

### UnsetMonitorPort
`func (o *InlineResponse20083LoadBalancerNode) UnsetMonitorPort()`

UnsetMonitorPort ensures that no value is present for MonitorPort, not even an explicit nil
### GetWeight

`func (o *InlineResponse20083LoadBalancerNode) GetWeight() int64`

GetWeight returns the Weight field if non-nil, zero value otherwise.

### GetWeightOk

`func (o *InlineResponse20083LoadBalancerNode) GetWeightOk() (*int64, bool)`

GetWeightOk returns a tuple with the Weight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWeight

`func (o *InlineResponse20083LoadBalancerNode) SetWeight(v int64)`

SetWeight sets Weight field to given value.

### HasWeight

`func (o *InlineResponse20083LoadBalancerNode) HasWeight() bool`

HasWeight returns a boolean if a field has been set.

### SetWeightNil

`func (o *InlineResponse20083LoadBalancerNode) SetWeightNil(b bool)`

 SetWeightNil sets the value for Weight to be an explicit nil

### UnsetWeight
`func (o *InlineResponse20083LoadBalancerNode) UnsetWeight()`

UnsetWeight ensures that no value is present for Weight, not even an explicit nil
### GetNodeState

`func (o *InlineResponse20083LoadBalancerNode) GetNodeState() string`

GetNodeState returns the NodeState field if non-nil, zero value otherwise.

### GetNodeStateOk

`func (o *InlineResponse20083LoadBalancerNode) GetNodeStateOk() (*string, bool)`

GetNodeStateOk returns a tuple with the NodeState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeState

`func (o *InlineResponse20083LoadBalancerNode) SetNodeState(v string)`

SetNodeState sets NodeState field to given value.

### HasNodeState

`func (o *InlineResponse20083LoadBalancerNode) HasNodeState() bool`

HasNodeState returns a boolean if a field has been set.

### SetNodeStateNil

`func (o *InlineResponse20083LoadBalancerNode) SetNodeStateNil(b bool)`

 SetNodeStateNil sets the value for NodeState to be an explicit nil

### UnsetNodeState
`func (o *InlineResponse20083LoadBalancerNode) UnsetNodeState()`

UnsetNodeState ensures that no value is present for NodeState, not even an explicit nil
### GetInternalId

`func (o *InlineResponse20083LoadBalancerNode) GetInternalId() string`

GetInternalId returns the InternalId field if non-nil, zero value otherwise.

### GetInternalIdOk

`func (o *InlineResponse20083LoadBalancerNode) GetInternalIdOk() (*string, bool)`

GetInternalIdOk returns a tuple with the InternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInternalId

`func (o *InlineResponse20083LoadBalancerNode) SetInternalId(v string)`

SetInternalId sets InternalId field to given value.

### HasInternalId

`func (o *InlineResponse20083LoadBalancerNode) HasInternalId() bool`

HasInternalId returns a boolean if a field has been set.

### SetInternalIdNil

`func (o *InlineResponse20083LoadBalancerNode) SetInternalIdNil(b bool)`

 SetInternalIdNil sets the value for InternalId to be an explicit nil

### UnsetInternalId
`func (o *InlineResponse20083LoadBalancerNode) UnsetInternalId()`

UnsetInternalId ensures that no value is present for InternalId, not even an explicit nil
### GetExternalId

`func (o *InlineResponse20083LoadBalancerNode) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *InlineResponse20083LoadBalancerNode) GetExternalIdOk() (*string, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *InlineResponse20083LoadBalancerNode) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.

### HasExternalId

`func (o *InlineResponse20083LoadBalancerNode) HasExternalId() bool`

HasExternalId returns a boolean if a field has been set.

### SetExternalIdNil

`func (o *InlineResponse20083LoadBalancerNode) SetExternalIdNil(b bool)`

 SetExternalIdNil sets the value for ExternalId to be an explicit nil

### UnsetExternalId
`func (o *InlineResponse20083LoadBalancerNode) UnsetExternalId()`

UnsetExternalId ensures that no value is present for ExternalId, not even an explicit nil
### GetEnabled

`func (o *InlineResponse20083LoadBalancerNode) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *InlineResponse20083LoadBalancerNode) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *InlineResponse20083LoadBalancerNode) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *InlineResponse20083LoadBalancerNode) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetStatus

`func (o *InlineResponse20083LoadBalancerNode) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *InlineResponse20083LoadBalancerNode) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *InlineResponse20083LoadBalancerNode) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *InlineResponse20083LoadBalancerNode) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetStatusMessage

`func (o *InlineResponse20083LoadBalancerNode) GetStatusMessage() string`

GetStatusMessage returns the StatusMessage field if non-nil, zero value otherwise.

### GetStatusMessageOk

`func (o *InlineResponse20083LoadBalancerNode) GetStatusMessageOk() (*string, bool)`

GetStatusMessageOk returns a tuple with the StatusMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusMessage

`func (o *InlineResponse20083LoadBalancerNode) SetStatusMessage(v string)`

SetStatusMessage sets StatusMessage field to given value.

### HasStatusMessage

`func (o *InlineResponse20083LoadBalancerNode) HasStatusMessage() bool`

HasStatusMessage returns a boolean if a field has been set.

### SetStatusMessageNil

`func (o *InlineResponse20083LoadBalancerNode) SetStatusMessageNil(b bool)`

 SetStatusMessageNil sets the value for StatusMessage to be an explicit nil

### UnsetStatusMessage
`func (o *InlineResponse20083LoadBalancerNode) UnsetStatusMessage()`

UnsetStatusMessage ensures that no value is present for StatusMessage, not even an explicit nil
### GetStatusDate

`func (o *InlineResponse20083LoadBalancerNode) GetStatusDate() time.Time`

GetStatusDate returns the StatusDate field if non-nil, zero value otherwise.

### GetStatusDateOk

`func (o *InlineResponse20083LoadBalancerNode) GetStatusDateOk() (*time.Time, bool)`

GetStatusDateOk returns a tuple with the StatusDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusDate

`func (o *InlineResponse20083LoadBalancerNode) SetStatusDate(v time.Time)`

SetStatusDate sets StatusDate field to given value.

### HasStatusDate

`func (o *InlineResponse20083LoadBalancerNode) HasStatusDate() bool`

HasStatusDate returns a boolean if a field has been set.

### SetStatusDateNil

`func (o *InlineResponse20083LoadBalancerNode) SetStatusDateNil(b bool)`

 SetStatusDateNil sets the value for StatusDate to be an explicit nil

### UnsetStatusDate
`func (o *InlineResponse20083LoadBalancerNode) UnsetStatusDate()`

UnsetStatusDate ensures that no value is present for StatusDate, not even an explicit nil
### GetServer

`func (o *InlineResponse20083LoadBalancerNode) GetServer() InlineResponse20082LoadBalancerInstanceSslCert`

GetServer returns the Server field if non-nil, zero value otherwise.

### GetServerOk

`func (o *InlineResponse20083LoadBalancerNode) GetServerOk() (*InlineResponse20082LoadBalancerInstanceSslCert, bool)`

GetServerOk returns a tuple with the Server field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServer

`func (o *InlineResponse20083LoadBalancerNode) SetServer(v InlineResponse20082LoadBalancerInstanceSslCert)`

SetServer sets Server field to given value.

### HasServer

`func (o *InlineResponse20083LoadBalancerNode) HasServer() bool`

HasServer returns a boolean if a field has been set.

### SetServerNil

`func (o *InlineResponse20083LoadBalancerNode) SetServerNil(b bool)`

 SetServerNil sets the value for Server to be an explicit nil

### UnsetServer
`func (o *InlineResponse20083LoadBalancerNode) UnsetServer()`

UnsetServer ensures that no value is present for Server, not even an explicit nil
### GetInstanceId

`func (o *InlineResponse20083LoadBalancerNode) GetInstanceId() int64`

GetInstanceId returns the InstanceId field if non-nil, zero value otherwise.

### GetInstanceIdOk

`func (o *InlineResponse20083LoadBalancerNode) GetInstanceIdOk() (*int64, bool)`

GetInstanceIdOk returns a tuple with the InstanceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceId

`func (o *InlineResponse20083LoadBalancerNode) SetInstanceId(v int64)`

SetInstanceId sets InstanceId field to given value.

### HasInstanceId

`func (o *InlineResponse20083LoadBalancerNode) HasInstanceId() bool`

HasInstanceId returns a boolean if a field has been set.

### SetInstanceIdNil

`func (o *InlineResponse20083LoadBalancerNode) SetInstanceIdNil(b bool)`

 SetInstanceIdNil sets the value for InstanceId to be an explicit nil

### UnsetInstanceId
`func (o *InlineResponse20083LoadBalancerNode) UnsetInstanceId()`

UnsetInstanceId ensures that no value is present for InstanceId, not even an explicit nil
### GetContainerId

`func (o *InlineResponse20083LoadBalancerNode) GetContainerId() int64`

GetContainerId returns the ContainerId field if non-nil, zero value otherwise.

### GetContainerIdOk

`func (o *InlineResponse20083LoadBalancerNode) GetContainerIdOk() (*int64, bool)`

GetContainerIdOk returns a tuple with the ContainerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainerId

`func (o *InlineResponse20083LoadBalancerNode) SetContainerId(v int64)`

SetContainerId sets ContainerId field to given value.

### HasContainerId

`func (o *InlineResponse20083LoadBalancerNode) HasContainerId() bool`

HasContainerId returns a boolean if a field has been set.

### SetContainerIdNil

`func (o *InlineResponse20083LoadBalancerNode) SetContainerIdNil(b bool)`

 SetContainerIdNil sets the value for ContainerId to be an explicit nil

### UnsetContainerId
`func (o *InlineResponse20083LoadBalancerNode) UnsetContainerId()`

UnsetContainerId ensures that no value is present for ContainerId, not even an explicit nil
### GetNodeSource

`func (o *InlineResponse20083LoadBalancerNode) GetNodeSource() string`

GetNodeSource returns the NodeSource field if non-nil, zero value otherwise.

### GetNodeSourceOk

`func (o *InlineResponse20083LoadBalancerNode) GetNodeSourceOk() (*string, bool)`

GetNodeSourceOk returns a tuple with the NodeSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeSource

`func (o *InlineResponse20083LoadBalancerNode) SetNodeSource(v string)`

SetNodeSource sets NodeSource field to given value.

### HasNodeSource

`func (o *InlineResponse20083LoadBalancerNode) HasNodeSource() bool`

HasNodeSource returns a boolean if a field has been set.

### SetNodeSourceNil

`func (o *InlineResponse20083LoadBalancerNode) SetNodeSourceNil(b bool)`

 SetNodeSourceNil sets the value for NodeSource to be an explicit nil

### UnsetNodeSource
`func (o *InlineResponse20083LoadBalancerNode) UnsetNodeSource()`

UnsetNodeSource ensures that no value is present for NodeSource, not even an explicit nil
### GetMonitor

`func (o *InlineResponse20083LoadBalancerNode) GetMonitor() InlineResponse20082LoadBalancerInstanceSslCert`

GetMonitor returns the Monitor field if non-nil, zero value otherwise.

### GetMonitorOk

`func (o *InlineResponse20083LoadBalancerNode) GetMonitorOk() (*InlineResponse20082LoadBalancerInstanceSslCert, bool)`

GetMonitorOk returns a tuple with the Monitor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonitor

`func (o *InlineResponse20083LoadBalancerNode) SetMonitor(v InlineResponse20082LoadBalancerInstanceSslCert)`

SetMonitor sets Monitor field to given value.

### HasMonitor

`func (o *InlineResponse20083LoadBalancerNode) HasMonitor() bool`

HasMonitor returns a boolean if a field has been set.

### SetMonitorNil

`func (o *InlineResponse20083LoadBalancerNode) SetMonitorNil(b bool)`

 SetMonitorNil sets the value for Monitor to be an explicit nil

### UnsetMonitor
`func (o *InlineResponse20083LoadBalancerNode) UnsetMonitor()`

UnsetMonitor ensures that no value is present for Monitor, not even an explicit nil
### GetMaxConnections

`func (o *InlineResponse20083LoadBalancerNode) GetMaxConnections() int64`

GetMaxConnections returns the MaxConnections field if non-nil, zero value otherwise.

### GetMaxConnectionsOk

`func (o *InlineResponse20083LoadBalancerNode) GetMaxConnectionsOk() (*int64, bool)`

GetMaxConnectionsOk returns a tuple with the MaxConnections field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxConnections

`func (o *InlineResponse20083LoadBalancerNode) SetMaxConnections(v int64)`

SetMaxConnections sets MaxConnections field to given value.

### HasMaxConnections

`func (o *InlineResponse20083LoadBalancerNode) HasMaxConnections() bool`

HasMaxConnections returns a boolean if a field has been set.

### SetMaxConnectionsNil

`func (o *InlineResponse20083LoadBalancerNode) SetMaxConnectionsNil(b bool)`

 SetMaxConnectionsNil sets the value for MaxConnections to be an explicit nil

### UnsetMaxConnections
`func (o *InlineResponse20083LoadBalancerNode) UnsetMaxConnections()`

UnsetMaxConnections ensures that no value is present for MaxConnections, not even an explicit nil
### GetExternalRefType

`func (o *InlineResponse20083LoadBalancerNode) GetExternalRefType() string`

GetExternalRefType returns the ExternalRefType field if non-nil, zero value otherwise.

### GetExternalRefTypeOk

`func (o *InlineResponse20083LoadBalancerNode) GetExternalRefTypeOk() (*string, bool)`

GetExternalRefTypeOk returns a tuple with the ExternalRefType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalRefType

`func (o *InlineResponse20083LoadBalancerNode) SetExternalRefType(v string)`

SetExternalRefType sets ExternalRefType field to given value.

### HasExternalRefType

`func (o *InlineResponse20083LoadBalancerNode) HasExternalRefType() bool`

HasExternalRefType returns a boolean if a field has been set.

### SetExternalRefTypeNil

`func (o *InlineResponse20083LoadBalancerNode) SetExternalRefTypeNil(b bool)`

 SetExternalRefTypeNil sets the value for ExternalRefType to be an explicit nil

### UnsetExternalRefType
`func (o *InlineResponse20083LoadBalancerNode) UnsetExternalRefType()`

UnsetExternalRefType ensures that no value is present for ExternalRefType, not even an explicit nil
### GetExternalRefId

`func (o *InlineResponse20083LoadBalancerNode) GetExternalRefId() string`

GetExternalRefId returns the ExternalRefId field if non-nil, zero value otherwise.

### GetExternalRefIdOk

`func (o *InlineResponse20083LoadBalancerNode) GetExternalRefIdOk() (*string, bool)`

GetExternalRefIdOk returns a tuple with the ExternalRefId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalRefId

`func (o *InlineResponse20083LoadBalancerNode) SetExternalRefId(v string)`

SetExternalRefId sets ExternalRefId field to given value.

### HasExternalRefId

`func (o *InlineResponse20083LoadBalancerNode) HasExternalRefId() bool`

HasExternalRefId returns a boolean if a field has been set.

### SetExternalRefIdNil

`func (o *InlineResponse20083LoadBalancerNode) SetExternalRefIdNil(b bool)`

 SetExternalRefIdNil sets the value for ExternalRefId to be an explicit nil

### UnsetExternalRefId
`func (o *InlineResponse20083LoadBalancerNode) UnsetExternalRefId()`

UnsetExternalRefId ensures that no value is present for ExternalRefId, not even an explicit nil
### GetExternalRefName

`func (o *InlineResponse20083LoadBalancerNode) GetExternalRefName() string`

GetExternalRefName returns the ExternalRefName field if non-nil, zero value otherwise.

### GetExternalRefNameOk

`func (o *InlineResponse20083LoadBalancerNode) GetExternalRefNameOk() (*string, bool)`

GetExternalRefNameOk returns a tuple with the ExternalRefName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalRefName

`func (o *InlineResponse20083LoadBalancerNode) SetExternalRefName(v string)`

SetExternalRefName sets ExternalRefName field to given value.

### HasExternalRefName

`func (o *InlineResponse20083LoadBalancerNode) HasExternalRefName() bool`

HasExternalRefName returns a boolean if a field has been set.

### SetExternalRefNameNil

`func (o *InlineResponse20083LoadBalancerNode) SetExternalRefNameNil(b bool)`

 SetExternalRefNameNil sets the value for ExternalRefName to be an explicit nil

### UnsetExternalRefName
`func (o *InlineResponse20083LoadBalancerNode) UnsetExternalRefName()`

UnsetExternalRefName ensures that no value is present for ExternalRefName, not even an explicit nil
### GetCreatedBy

`func (o *InlineResponse20083LoadBalancerNode) GetCreatedBy() InlineResponse20083LoadBalancerNodeCreatedBy`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *InlineResponse20083LoadBalancerNode) GetCreatedByOk() (*InlineResponse20083LoadBalancerNodeCreatedBy, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *InlineResponse20083LoadBalancerNode) SetCreatedBy(v InlineResponse20083LoadBalancerNodeCreatedBy)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *InlineResponse20083LoadBalancerNode) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### SetCreatedByNil

`func (o *InlineResponse20083LoadBalancerNode) SetCreatedByNil(b bool)`

 SetCreatedByNil sets the value for CreatedBy to be an explicit nil

### UnsetCreatedBy
`func (o *InlineResponse20083LoadBalancerNode) UnsetCreatedBy()`

UnsetCreatedBy ensures that no value is present for CreatedBy, not even an explicit nil
### GetDateCreated

`func (o *InlineResponse20083LoadBalancerNode) GetDateCreated() time.Time`

GetDateCreated returns the DateCreated field if non-nil, zero value otherwise.

### GetDateCreatedOk

`func (o *InlineResponse20083LoadBalancerNode) GetDateCreatedOk() (*time.Time, bool)`

GetDateCreatedOk returns a tuple with the DateCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateCreated

`func (o *InlineResponse20083LoadBalancerNode) SetDateCreated(v time.Time)`

SetDateCreated sets DateCreated field to given value.

### HasDateCreated

`func (o *InlineResponse20083LoadBalancerNode) HasDateCreated() bool`

HasDateCreated returns a boolean if a field has been set.

### GetLastUpdated

`func (o *InlineResponse20083LoadBalancerNode) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *InlineResponse20083LoadBalancerNode) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *InlineResponse20083LoadBalancerNode) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.

### HasLastUpdated

`func (o *InlineResponse20083LoadBalancerNode) HasLastUpdated() bool`

HasLastUpdated returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


