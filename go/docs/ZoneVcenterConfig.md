# ZoneVcenterConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiUrl** | Pointer to **string** |  | [optional] 
**Username** | Pointer to **string** |  | [optional] 
**Password** | Pointer to **NullableString** |  | [optional] 
**Datacenter** | Pointer to **string** |  | [optional] 
**Cluster** | Pointer to **string** |  | [optional] 
**ResourcePoolId** | Pointer to **string** |  | [optional] 
**ResourcePool** | Pointer to **string** |  | [optional] 
**RpcMode** | Pointer to **string** |  | [optional] 
**HideHostSelection** | Pointer to **string** |  | [optional] 
**HideHostSelection** | Pointer to **string** |  | [optional] 
**ImportExisting** | Pointer to **string** |  | [optional] 
**ImportExisting** | Pointer to **string** |  | [optional] 
**EnableVnc** | Pointer to **string** |  | [optional] 
**EnableVnc** | Pointer to **string** |  | [optional] 
**EnableDiskTypeSelection** | Pointer to **string** |  | [optional] 
**EnableNetworkTypeSelection** | Pointer to **string** |  | [optional] 
**DiskStorageType** | Pointer to **NullableString** |  | [optional] 
**ApplianceUrl** | Pointer to **NullableString** |  | [optional] 
**DatacenterName** | Pointer to **NullableString** |  | [optional] 
**NetworkServerId** | Pointer to **string** |  | [optional] 
**NetworkServer** | Pointer to [**ZoneVcenterConfigNetworkServer**](zoneVcenterConfig_networkServer.md) |  | [optional] 
**SecurityMode** | Pointer to **string** |  | [optional] 
**CertificateProvider** | Pointer to **NullableString** |  | [optional] 
**BackupMode** | Pointer to **NullableString** |  | [optional] 
**ReplicationMode** | Pointer to **NullableString** |  | [optional] 
**DnsIntegrationId** | Pointer to **NullableString** |  | [optional] 
**ConfigCmdbId** | Pointer to **NullableString** |  | [optional] 
**ConfigManagementId** | Pointer to **NullableString** |  | [optional] 
**ConfigCmId** | Pointer to **NullableString** |  | [optional] 
**SecurityServer** | Pointer to **NullableString** |  | [optional] 
**ServiceRegistryId** | Pointer to **NullableString** |  | [optional] 
**EnableDiskTypeSelection** | Pointer to **NullableString** |  | [optional] 
**KubeUrl** | Pointer to **NullableString** |  | [optional] 
**ApiVersion** | Pointer to **NullableString** |  | [optional] 
**DatacenterId** | Pointer to **NullableString** |  | [optional] 
**ConfigCmdbDiscovery** | Pointer to **bool** |  | [optional] 
**DistributedWorkerId** | Pointer to **NullableString** |  | [optional] 
**PasswordHash** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewZoneVcenterConfig

`func NewZoneVcenterConfig() *ZoneVcenterConfig`

NewZoneVcenterConfig instantiates a new ZoneVcenterConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewZoneVcenterConfigWithDefaults

`func NewZoneVcenterConfigWithDefaults() *ZoneVcenterConfig`

NewZoneVcenterConfigWithDefaults instantiates a new ZoneVcenterConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiUrl

`func (o *ZoneVcenterConfig) GetApiUrl() string`

GetApiUrl returns the ApiUrl field if non-nil, zero value otherwise.

### GetApiUrlOk

`func (o *ZoneVcenterConfig) GetApiUrlOk() (*string, bool)`

GetApiUrlOk returns a tuple with the ApiUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiUrl

`func (o *ZoneVcenterConfig) SetApiUrl(v string)`

SetApiUrl sets ApiUrl field to given value.

### HasApiUrl

`func (o *ZoneVcenterConfig) HasApiUrl() bool`

HasApiUrl returns a boolean if a field has been set.

### GetUsername

`func (o *ZoneVcenterConfig) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *ZoneVcenterConfig) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *ZoneVcenterConfig) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *ZoneVcenterConfig) HasUsername() bool`

HasUsername returns a boolean if a field has been set.

### GetPassword

`func (o *ZoneVcenterConfig) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *ZoneVcenterConfig) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *ZoneVcenterConfig) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *ZoneVcenterConfig) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### SetPasswordNil

`func (o *ZoneVcenterConfig) SetPasswordNil(b bool)`

 SetPasswordNil sets the value for Password to be an explicit nil

### UnsetPassword
`func (o *ZoneVcenterConfig) UnsetPassword()`

UnsetPassword ensures that no value is present for Password, not even an explicit nil
### GetDatacenter

`func (o *ZoneVcenterConfig) GetDatacenter() string`

GetDatacenter returns the Datacenter field if non-nil, zero value otherwise.

### GetDatacenterOk

`func (o *ZoneVcenterConfig) GetDatacenterOk() (*string, bool)`

GetDatacenterOk returns a tuple with the Datacenter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatacenter

`func (o *ZoneVcenterConfig) SetDatacenter(v string)`

SetDatacenter sets Datacenter field to given value.

### HasDatacenter

`func (o *ZoneVcenterConfig) HasDatacenter() bool`

HasDatacenter returns a boolean if a field has been set.

### GetCluster

`func (o *ZoneVcenterConfig) GetCluster() string`

GetCluster returns the Cluster field if non-nil, zero value otherwise.

### GetClusterOk

`func (o *ZoneVcenterConfig) GetClusterOk() (*string, bool)`

GetClusterOk returns a tuple with the Cluster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCluster

`func (o *ZoneVcenterConfig) SetCluster(v string)`

SetCluster sets Cluster field to given value.

### HasCluster

`func (o *ZoneVcenterConfig) HasCluster() bool`

HasCluster returns a boolean if a field has been set.

### GetResourcePoolId

`func (o *ZoneVcenterConfig) GetResourcePoolId() string`

GetResourcePoolId returns the ResourcePoolId field if non-nil, zero value otherwise.

### GetResourcePoolIdOk

`func (o *ZoneVcenterConfig) GetResourcePoolIdOk() (*string, bool)`

GetResourcePoolIdOk returns a tuple with the ResourcePoolId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourcePoolId

`func (o *ZoneVcenterConfig) SetResourcePoolId(v string)`

SetResourcePoolId sets ResourcePoolId field to given value.

### HasResourcePoolId

`func (o *ZoneVcenterConfig) HasResourcePoolId() bool`

HasResourcePoolId returns a boolean if a field has been set.

### GetResourcePool

`func (o *ZoneVcenterConfig) GetResourcePool() string`

GetResourcePool returns the ResourcePool field if non-nil, zero value otherwise.

### GetResourcePoolOk

`func (o *ZoneVcenterConfig) GetResourcePoolOk() (*string, bool)`

GetResourcePoolOk returns a tuple with the ResourcePool field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourcePool

`func (o *ZoneVcenterConfig) SetResourcePool(v string)`

SetResourcePool sets ResourcePool field to given value.

### HasResourcePool

`func (o *ZoneVcenterConfig) HasResourcePool() bool`

HasResourcePool returns a boolean if a field has been set.

### GetRpcMode

`func (o *ZoneVcenterConfig) GetRpcMode() string`

GetRpcMode returns the RpcMode field if non-nil, zero value otherwise.

### GetRpcModeOk

`func (o *ZoneVcenterConfig) GetRpcModeOk() (*string, bool)`

GetRpcModeOk returns a tuple with the RpcMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRpcMode

`func (o *ZoneVcenterConfig) SetRpcMode(v string)`

SetRpcMode sets RpcMode field to given value.

### HasRpcMode

`func (o *ZoneVcenterConfig) HasRpcMode() bool`

HasRpcMode returns a boolean if a field has been set.

### GetHideHostSelection

`func (o *ZoneVcenterConfig) GetHideHostSelection() string`

GetHideHostSelection returns the HideHostSelection field if non-nil, zero value otherwise.

### GetHideHostSelectionOk

`func (o *ZoneVcenterConfig) GetHideHostSelectionOk() (*string, bool)`

GetHideHostSelectionOk returns a tuple with the HideHostSelection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHideHostSelection

`func (o *ZoneVcenterConfig) SetHideHostSelection(v string)`

SetHideHostSelection sets HideHostSelection field to given value.

### HasHideHostSelection

`func (o *ZoneVcenterConfig) HasHideHostSelection() bool`

HasHideHostSelection returns a boolean if a field has been set.

### GetHideHostSelection

`func (o *ZoneVcenterConfig) GetHideHostSelection() string`

GetHideHostSelection returns the HideHostSelection field if non-nil, zero value otherwise.

### GetHideHostSelectionOk

`func (o *ZoneVcenterConfig) GetHideHostSelectionOk() (*string, bool)`

GetHideHostSelectionOk returns a tuple with the HideHostSelection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHideHostSelection

`func (o *ZoneVcenterConfig) SetHideHostSelection(v string)`

SetHideHostSelection sets HideHostSelection field to given value.

### HasHideHostSelection

`func (o *ZoneVcenterConfig) HasHideHostSelection() bool`

HasHideHostSelection returns a boolean if a field has been set.

### GetImportExisting

`func (o *ZoneVcenterConfig) GetImportExisting() string`

GetImportExisting returns the ImportExisting field if non-nil, zero value otherwise.

### GetImportExistingOk

`func (o *ZoneVcenterConfig) GetImportExistingOk() (*string, bool)`

GetImportExistingOk returns a tuple with the ImportExisting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImportExisting

`func (o *ZoneVcenterConfig) SetImportExisting(v string)`

SetImportExisting sets ImportExisting field to given value.

### HasImportExisting

`func (o *ZoneVcenterConfig) HasImportExisting() bool`

HasImportExisting returns a boolean if a field has been set.

### GetImportExisting

`func (o *ZoneVcenterConfig) GetImportExisting() string`

GetImportExisting returns the ImportExisting field if non-nil, zero value otherwise.

### GetImportExistingOk

`func (o *ZoneVcenterConfig) GetImportExistingOk() (*string, bool)`

GetImportExistingOk returns a tuple with the ImportExisting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImportExisting

`func (o *ZoneVcenterConfig) SetImportExisting(v string)`

SetImportExisting sets ImportExisting field to given value.

### HasImportExisting

`func (o *ZoneVcenterConfig) HasImportExisting() bool`

HasImportExisting returns a boolean if a field has been set.

### GetEnableVnc

`func (o *ZoneVcenterConfig) GetEnableVnc() string`

GetEnableVnc returns the EnableVnc field if non-nil, zero value otherwise.

### GetEnableVncOk

`func (o *ZoneVcenterConfig) GetEnableVncOk() (*string, bool)`

GetEnableVncOk returns a tuple with the EnableVnc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableVnc

`func (o *ZoneVcenterConfig) SetEnableVnc(v string)`

SetEnableVnc sets EnableVnc field to given value.

### HasEnableVnc

`func (o *ZoneVcenterConfig) HasEnableVnc() bool`

HasEnableVnc returns a boolean if a field has been set.

### GetEnableVnc

`func (o *ZoneVcenterConfig) GetEnableVnc() string`

GetEnableVnc returns the EnableVnc field if non-nil, zero value otherwise.

### GetEnableVncOk

`func (o *ZoneVcenterConfig) GetEnableVncOk() (*string, bool)`

GetEnableVncOk returns a tuple with the EnableVnc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableVnc

`func (o *ZoneVcenterConfig) SetEnableVnc(v string)`

SetEnableVnc sets EnableVnc field to given value.

### HasEnableVnc

`func (o *ZoneVcenterConfig) HasEnableVnc() bool`

HasEnableVnc returns a boolean if a field has been set.

### GetEnableDiskTypeSelection

`func (o *ZoneVcenterConfig) GetEnableDiskTypeSelection() string`

GetEnableDiskTypeSelection returns the EnableDiskTypeSelection field if non-nil, zero value otherwise.

### GetEnableDiskTypeSelectionOk

`func (o *ZoneVcenterConfig) GetEnableDiskTypeSelectionOk() (*string, bool)`

GetEnableDiskTypeSelectionOk returns a tuple with the EnableDiskTypeSelection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableDiskTypeSelection

`func (o *ZoneVcenterConfig) SetEnableDiskTypeSelection(v string)`

SetEnableDiskTypeSelection sets EnableDiskTypeSelection field to given value.

### HasEnableDiskTypeSelection

`func (o *ZoneVcenterConfig) HasEnableDiskTypeSelection() bool`

HasEnableDiskTypeSelection returns a boolean if a field has been set.

### GetEnableNetworkTypeSelection

`func (o *ZoneVcenterConfig) GetEnableNetworkTypeSelection() string`

GetEnableNetworkTypeSelection returns the EnableNetworkTypeSelection field if non-nil, zero value otherwise.

### GetEnableNetworkTypeSelectionOk

`func (o *ZoneVcenterConfig) GetEnableNetworkTypeSelectionOk() (*string, bool)`

GetEnableNetworkTypeSelectionOk returns a tuple with the EnableNetworkTypeSelection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableNetworkTypeSelection

`func (o *ZoneVcenterConfig) SetEnableNetworkTypeSelection(v string)`

SetEnableNetworkTypeSelection sets EnableNetworkTypeSelection field to given value.

### HasEnableNetworkTypeSelection

`func (o *ZoneVcenterConfig) HasEnableNetworkTypeSelection() bool`

HasEnableNetworkTypeSelection returns a boolean if a field has been set.

### GetDiskStorageType

`func (o *ZoneVcenterConfig) GetDiskStorageType() string`

GetDiskStorageType returns the DiskStorageType field if non-nil, zero value otherwise.

### GetDiskStorageTypeOk

`func (o *ZoneVcenterConfig) GetDiskStorageTypeOk() (*string, bool)`

GetDiskStorageTypeOk returns a tuple with the DiskStorageType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiskStorageType

`func (o *ZoneVcenterConfig) SetDiskStorageType(v string)`

SetDiskStorageType sets DiskStorageType field to given value.

### HasDiskStorageType

`func (o *ZoneVcenterConfig) HasDiskStorageType() bool`

HasDiskStorageType returns a boolean if a field has been set.

### SetDiskStorageTypeNil

`func (o *ZoneVcenterConfig) SetDiskStorageTypeNil(b bool)`

 SetDiskStorageTypeNil sets the value for DiskStorageType to be an explicit nil

### UnsetDiskStorageType
`func (o *ZoneVcenterConfig) UnsetDiskStorageType()`

UnsetDiskStorageType ensures that no value is present for DiskStorageType, not even an explicit nil
### GetApplianceUrl

`func (o *ZoneVcenterConfig) GetApplianceUrl() string`

GetApplianceUrl returns the ApplianceUrl field if non-nil, zero value otherwise.

### GetApplianceUrlOk

`func (o *ZoneVcenterConfig) GetApplianceUrlOk() (*string, bool)`

GetApplianceUrlOk returns a tuple with the ApplianceUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplianceUrl

`func (o *ZoneVcenterConfig) SetApplianceUrl(v string)`

SetApplianceUrl sets ApplianceUrl field to given value.

### HasApplianceUrl

`func (o *ZoneVcenterConfig) HasApplianceUrl() bool`

HasApplianceUrl returns a boolean if a field has been set.

### SetApplianceUrlNil

`func (o *ZoneVcenterConfig) SetApplianceUrlNil(b bool)`

 SetApplianceUrlNil sets the value for ApplianceUrl to be an explicit nil

### UnsetApplianceUrl
`func (o *ZoneVcenterConfig) UnsetApplianceUrl()`

UnsetApplianceUrl ensures that no value is present for ApplianceUrl, not even an explicit nil
### GetDatacenterName

`func (o *ZoneVcenterConfig) GetDatacenterName() string`

GetDatacenterName returns the DatacenterName field if non-nil, zero value otherwise.

### GetDatacenterNameOk

`func (o *ZoneVcenterConfig) GetDatacenterNameOk() (*string, bool)`

GetDatacenterNameOk returns a tuple with the DatacenterName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatacenterName

`func (o *ZoneVcenterConfig) SetDatacenterName(v string)`

SetDatacenterName sets DatacenterName field to given value.

### HasDatacenterName

`func (o *ZoneVcenterConfig) HasDatacenterName() bool`

HasDatacenterName returns a boolean if a field has been set.

### SetDatacenterNameNil

`func (o *ZoneVcenterConfig) SetDatacenterNameNil(b bool)`

 SetDatacenterNameNil sets the value for DatacenterName to be an explicit nil

### UnsetDatacenterName
`func (o *ZoneVcenterConfig) UnsetDatacenterName()`

UnsetDatacenterName ensures that no value is present for DatacenterName, not even an explicit nil
### GetNetworkServerId

`func (o *ZoneVcenterConfig) GetNetworkServerId() string`

GetNetworkServerId returns the NetworkServerId field if non-nil, zero value otherwise.

### GetNetworkServerIdOk

`func (o *ZoneVcenterConfig) GetNetworkServerIdOk() (*string, bool)`

GetNetworkServerIdOk returns a tuple with the NetworkServerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkServerId

`func (o *ZoneVcenterConfig) SetNetworkServerId(v string)`

SetNetworkServerId sets NetworkServerId field to given value.

### HasNetworkServerId

`func (o *ZoneVcenterConfig) HasNetworkServerId() bool`

HasNetworkServerId returns a boolean if a field has been set.

### GetNetworkServer

`func (o *ZoneVcenterConfig) GetNetworkServer() ZoneVcenterConfigNetworkServer`

GetNetworkServer returns the NetworkServer field if non-nil, zero value otherwise.

### GetNetworkServerOk

`func (o *ZoneVcenterConfig) GetNetworkServerOk() (*ZoneVcenterConfigNetworkServer, bool)`

GetNetworkServerOk returns a tuple with the NetworkServer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkServer

`func (o *ZoneVcenterConfig) SetNetworkServer(v ZoneVcenterConfigNetworkServer)`

SetNetworkServer sets NetworkServer field to given value.

### HasNetworkServer

`func (o *ZoneVcenterConfig) HasNetworkServer() bool`

HasNetworkServer returns a boolean if a field has been set.

### GetSecurityMode

`func (o *ZoneVcenterConfig) GetSecurityMode() string`

GetSecurityMode returns the SecurityMode field if non-nil, zero value otherwise.

### GetSecurityModeOk

`func (o *ZoneVcenterConfig) GetSecurityModeOk() (*string, bool)`

GetSecurityModeOk returns a tuple with the SecurityMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurityMode

`func (o *ZoneVcenterConfig) SetSecurityMode(v string)`

SetSecurityMode sets SecurityMode field to given value.

### HasSecurityMode

`func (o *ZoneVcenterConfig) HasSecurityMode() bool`

HasSecurityMode returns a boolean if a field has been set.

### GetCertificateProvider

`func (o *ZoneVcenterConfig) GetCertificateProvider() string`

GetCertificateProvider returns the CertificateProvider field if non-nil, zero value otherwise.

### GetCertificateProviderOk

`func (o *ZoneVcenterConfig) GetCertificateProviderOk() (*string, bool)`

GetCertificateProviderOk returns a tuple with the CertificateProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificateProvider

`func (o *ZoneVcenterConfig) SetCertificateProvider(v string)`

SetCertificateProvider sets CertificateProvider field to given value.

### HasCertificateProvider

`func (o *ZoneVcenterConfig) HasCertificateProvider() bool`

HasCertificateProvider returns a boolean if a field has been set.

### SetCertificateProviderNil

`func (o *ZoneVcenterConfig) SetCertificateProviderNil(b bool)`

 SetCertificateProviderNil sets the value for CertificateProvider to be an explicit nil

### UnsetCertificateProvider
`func (o *ZoneVcenterConfig) UnsetCertificateProvider()`

UnsetCertificateProvider ensures that no value is present for CertificateProvider, not even an explicit nil
### GetBackupMode

`func (o *ZoneVcenterConfig) GetBackupMode() string`

GetBackupMode returns the BackupMode field if non-nil, zero value otherwise.

### GetBackupModeOk

`func (o *ZoneVcenterConfig) GetBackupModeOk() (*string, bool)`

GetBackupModeOk returns a tuple with the BackupMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackupMode

`func (o *ZoneVcenterConfig) SetBackupMode(v string)`

SetBackupMode sets BackupMode field to given value.

### HasBackupMode

`func (o *ZoneVcenterConfig) HasBackupMode() bool`

HasBackupMode returns a boolean if a field has been set.

### SetBackupModeNil

`func (o *ZoneVcenterConfig) SetBackupModeNil(b bool)`

 SetBackupModeNil sets the value for BackupMode to be an explicit nil

### UnsetBackupMode
`func (o *ZoneVcenterConfig) UnsetBackupMode()`

UnsetBackupMode ensures that no value is present for BackupMode, not even an explicit nil
### GetReplicationMode

`func (o *ZoneVcenterConfig) GetReplicationMode() string`

GetReplicationMode returns the ReplicationMode field if non-nil, zero value otherwise.

### GetReplicationModeOk

`func (o *ZoneVcenterConfig) GetReplicationModeOk() (*string, bool)`

GetReplicationModeOk returns a tuple with the ReplicationMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplicationMode

`func (o *ZoneVcenterConfig) SetReplicationMode(v string)`

SetReplicationMode sets ReplicationMode field to given value.

### HasReplicationMode

`func (o *ZoneVcenterConfig) HasReplicationMode() bool`

HasReplicationMode returns a boolean if a field has been set.

### SetReplicationModeNil

`func (o *ZoneVcenterConfig) SetReplicationModeNil(b bool)`

 SetReplicationModeNil sets the value for ReplicationMode to be an explicit nil

### UnsetReplicationMode
`func (o *ZoneVcenterConfig) UnsetReplicationMode()`

UnsetReplicationMode ensures that no value is present for ReplicationMode, not even an explicit nil
### GetDnsIntegrationId

`func (o *ZoneVcenterConfig) GetDnsIntegrationId() string`

GetDnsIntegrationId returns the DnsIntegrationId field if non-nil, zero value otherwise.

### GetDnsIntegrationIdOk

`func (o *ZoneVcenterConfig) GetDnsIntegrationIdOk() (*string, bool)`

GetDnsIntegrationIdOk returns a tuple with the DnsIntegrationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsIntegrationId

`func (o *ZoneVcenterConfig) SetDnsIntegrationId(v string)`

SetDnsIntegrationId sets DnsIntegrationId field to given value.

### HasDnsIntegrationId

`func (o *ZoneVcenterConfig) HasDnsIntegrationId() bool`

HasDnsIntegrationId returns a boolean if a field has been set.

### SetDnsIntegrationIdNil

`func (o *ZoneVcenterConfig) SetDnsIntegrationIdNil(b bool)`

 SetDnsIntegrationIdNil sets the value for DnsIntegrationId to be an explicit nil

### UnsetDnsIntegrationId
`func (o *ZoneVcenterConfig) UnsetDnsIntegrationId()`

UnsetDnsIntegrationId ensures that no value is present for DnsIntegrationId, not even an explicit nil
### GetConfigCmdbId

`func (o *ZoneVcenterConfig) GetConfigCmdbId() string`

GetConfigCmdbId returns the ConfigCmdbId field if non-nil, zero value otherwise.

### GetConfigCmdbIdOk

`func (o *ZoneVcenterConfig) GetConfigCmdbIdOk() (*string, bool)`

GetConfigCmdbIdOk returns a tuple with the ConfigCmdbId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigCmdbId

`func (o *ZoneVcenterConfig) SetConfigCmdbId(v string)`

SetConfigCmdbId sets ConfigCmdbId field to given value.

### HasConfigCmdbId

`func (o *ZoneVcenterConfig) HasConfigCmdbId() bool`

HasConfigCmdbId returns a boolean if a field has been set.

### SetConfigCmdbIdNil

`func (o *ZoneVcenterConfig) SetConfigCmdbIdNil(b bool)`

 SetConfigCmdbIdNil sets the value for ConfigCmdbId to be an explicit nil

### UnsetConfigCmdbId
`func (o *ZoneVcenterConfig) UnsetConfigCmdbId()`

UnsetConfigCmdbId ensures that no value is present for ConfigCmdbId, not even an explicit nil
### GetConfigManagementId

`func (o *ZoneVcenterConfig) GetConfigManagementId() string`

GetConfigManagementId returns the ConfigManagementId field if non-nil, zero value otherwise.

### GetConfigManagementIdOk

`func (o *ZoneVcenterConfig) GetConfigManagementIdOk() (*string, bool)`

GetConfigManagementIdOk returns a tuple with the ConfigManagementId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigManagementId

`func (o *ZoneVcenterConfig) SetConfigManagementId(v string)`

SetConfigManagementId sets ConfigManagementId field to given value.

### HasConfigManagementId

`func (o *ZoneVcenterConfig) HasConfigManagementId() bool`

HasConfigManagementId returns a boolean if a field has been set.

### SetConfigManagementIdNil

`func (o *ZoneVcenterConfig) SetConfigManagementIdNil(b bool)`

 SetConfigManagementIdNil sets the value for ConfigManagementId to be an explicit nil

### UnsetConfigManagementId
`func (o *ZoneVcenterConfig) UnsetConfigManagementId()`

UnsetConfigManagementId ensures that no value is present for ConfigManagementId, not even an explicit nil
### GetConfigCmId

`func (o *ZoneVcenterConfig) GetConfigCmId() string`

GetConfigCmId returns the ConfigCmId field if non-nil, zero value otherwise.

### GetConfigCmIdOk

`func (o *ZoneVcenterConfig) GetConfigCmIdOk() (*string, bool)`

GetConfigCmIdOk returns a tuple with the ConfigCmId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigCmId

`func (o *ZoneVcenterConfig) SetConfigCmId(v string)`

SetConfigCmId sets ConfigCmId field to given value.

### HasConfigCmId

`func (o *ZoneVcenterConfig) HasConfigCmId() bool`

HasConfigCmId returns a boolean if a field has been set.

### SetConfigCmIdNil

`func (o *ZoneVcenterConfig) SetConfigCmIdNil(b bool)`

 SetConfigCmIdNil sets the value for ConfigCmId to be an explicit nil

### UnsetConfigCmId
`func (o *ZoneVcenterConfig) UnsetConfigCmId()`

UnsetConfigCmId ensures that no value is present for ConfigCmId, not even an explicit nil
### GetSecurityServer

`func (o *ZoneVcenterConfig) GetSecurityServer() string`

GetSecurityServer returns the SecurityServer field if non-nil, zero value otherwise.

### GetSecurityServerOk

`func (o *ZoneVcenterConfig) GetSecurityServerOk() (*string, bool)`

GetSecurityServerOk returns a tuple with the SecurityServer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurityServer

`func (o *ZoneVcenterConfig) SetSecurityServer(v string)`

SetSecurityServer sets SecurityServer field to given value.

### HasSecurityServer

`func (o *ZoneVcenterConfig) HasSecurityServer() bool`

HasSecurityServer returns a boolean if a field has been set.

### SetSecurityServerNil

`func (o *ZoneVcenterConfig) SetSecurityServerNil(b bool)`

 SetSecurityServerNil sets the value for SecurityServer to be an explicit nil

### UnsetSecurityServer
`func (o *ZoneVcenterConfig) UnsetSecurityServer()`

UnsetSecurityServer ensures that no value is present for SecurityServer, not even an explicit nil
### GetServiceRegistryId

`func (o *ZoneVcenterConfig) GetServiceRegistryId() string`

GetServiceRegistryId returns the ServiceRegistryId field if non-nil, zero value otherwise.

### GetServiceRegistryIdOk

`func (o *ZoneVcenterConfig) GetServiceRegistryIdOk() (*string, bool)`

GetServiceRegistryIdOk returns a tuple with the ServiceRegistryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceRegistryId

`func (o *ZoneVcenterConfig) SetServiceRegistryId(v string)`

SetServiceRegistryId sets ServiceRegistryId field to given value.

### HasServiceRegistryId

`func (o *ZoneVcenterConfig) HasServiceRegistryId() bool`

HasServiceRegistryId returns a boolean if a field has been set.

### SetServiceRegistryIdNil

`func (o *ZoneVcenterConfig) SetServiceRegistryIdNil(b bool)`

 SetServiceRegistryIdNil sets the value for ServiceRegistryId to be an explicit nil

### UnsetServiceRegistryId
`func (o *ZoneVcenterConfig) UnsetServiceRegistryId()`

UnsetServiceRegistryId ensures that no value is present for ServiceRegistryId, not even an explicit nil
### GetEnableDiskTypeSelection

`func (o *ZoneVcenterConfig) GetEnableDiskTypeSelection() string`

GetEnableDiskTypeSelection returns the EnableDiskTypeSelection field if non-nil, zero value otherwise.

### GetEnableDiskTypeSelectionOk

`func (o *ZoneVcenterConfig) GetEnableDiskTypeSelectionOk() (*string, bool)`

GetEnableDiskTypeSelectionOk returns a tuple with the EnableDiskTypeSelection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableDiskTypeSelection

`func (o *ZoneVcenterConfig) SetEnableDiskTypeSelection(v string)`

SetEnableDiskTypeSelection sets EnableDiskTypeSelection field to given value.

### HasEnableDiskTypeSelection

`func (o *ZoneVcenterConfig) HasEnableDiskTypeSelection() bool`

HasEnableDiskTypeSelection returns a boolean if a field has been set.

### SetEnableDiskTypeSelectionNil

`func (o *ZoneVcenterConfig) SetEnableDiskTypeSelectionNil(b bool)`

 SetEnableDiskTypeSelectionNil sets the value for EnableDiskTypeSelection to be an explicit nil

### UnsetEnableDiskTypeSelection
`func (o *ZoneVcenterConfig) UnsetEnableDiskTypeSelection()`

UnsetEnableDiskTypeSelection ensures that no value is present for EnableDiskTypeSelection, not even an explicit nil
### GetKubeUrl

`func (o *ZoneVcenterConfig) GetKubeUrl() string`

GetKubeUrl returns the KubeUrl field if non-nil, zero value otherwise.

### GetKubeUrlOk

`func (o *ZoneVcenterConfig) GetKubeUrlOk() (*string, bool)`

GetKubeUrlOk returns a tuple with the KubeUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKubeUrl

`func (o *ZoneVcenterConfig) SetKubeUrl(v string)`

SetKubeUrl sets KubeUrl field to given value.

### HasKubeUrl

`func (o *ZoneVcenterConfig) HasKubeUrl() bool`

HasKubeUrl returns a boolean if a field has been set.

### SetKubeUrlNil

`func (o *ZoneVcenterConfig) SetKubeUrlNil(b bool)`

 SetKubeUrlNil sets the value for KubeUrl to be an explicit nil

### UnsetKubeUrl
`func (o *ZoneVcenterConfig) UnsetKubeUrl()`

UnsetKubeUrl ensures that no value is present for KubeUrl, not even an explicit nil
### GetApiVersion

`func (o *ZoneVcenterConfig) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *ZoneVcenterConfig) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *ZoneVcenterConfig) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.

### HasApiVersion

`func (o *ZoneVcenterConfig) HasApiVersion() bool`

HasApiVersion returns a boolean if a field has been set.

### SetApiVersionNil

`func (o *ZoneVcenterConfig) SetApiVersionNil(b bool)`

 SetApiVersionNil sets the value for ApiVersion to be an explicit nil

### UnsetApiVersion
`func (o *ZoneVcenterConfig) UnsetApiVersion()`

UnsetApiVersion ensures that no value is present for ApiVersion, not even an explicit nil
### GetDatacenterId

`func (o *ZoneVcenterConfig) GetDatacenterId() string`

GetDatacenterId returns the DatacenterId field if non-nil, zero value otherwise.

### GetDatacenterIdOk

`func (o *ZoneVcenterConfig) GetDatacenterIdOk() (*string, bool)`

GetDatacenterIdOk returns a tuple with the DatacenterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatacenterId

`func (o *ZoneVcenterConfig) SetDatacenterId(v string)`

SetDatacenterId sets DatacenterId field to given value.

### HasDatacenterId

`func (o *ZoneVcenterConfig) HasDatacenterId() bool`

HasDatacenterId returns a boolean if a field has been set.

### SetDatacenterIdNil

`func (o *ZoneVcenterConfig) SetDatacenterIdNil(b bool)`

 SetDatacenterIdNil sets the value for DatacenterId to be an explicit nil

### UnsetDatacenterId
`func (o *ZoneVcenterConfig) UnsetDatacenterId()`

UnsetDatacenterId ensures that no value is present for DatacenterId, not even an explicit nil
### GetConfigCmdbDiscovery

`func (o *ZoneVcenterConfig) GetConfigCmdbDiscovery() bool`

GetConfigCmdbDiscovery returns the ConfigCmdbDiscovery field if non-nil, zero value otherwise.

### GetConfigCmdbDiscoveryOk

`func (o *ZoneVcenterConfig) GetConfigCmdbDiscoveryOk() (*bool, bool)`

GetConfigCmdbDiscoveryOk returns a tuple with the ConfigCmdbDiscovery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigCmdbDiscovery

`func (o *ZoneVcenterConfig) SetConfigCmdbDiscovery(v bool)`

SetConfigCmdbDiscovery sets ConfigCmdbDiscovery field to given value.

### HasConfigCmdbDiscovery

`func (o *ZoneVcenterConfig) HasConfigCmdbDiscovery() bool`

HasConfigCmdbDiscovery returns a boolean if a field has been set.

### GetDistributedWorkerId

`func (o *ZoneVcenterConfig) GetDistributedWorkerId() string`

GetDistributedWorkerId returns the DistributedWorkerId field if non-nil, zero value otherwise.

### GetDistributedWorkerIdOk

`func (o *ZoneVcenterConfig) GetDistributedWorkerIdOk() (*string, bool)`

GetDistributedWorkerIdOk returns a tuple with the DistributedWorkerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDistributedWorkerId

`func (o *ZoneVcenterConfig) SetDistributedWorkerId(v string)`

SetDistributedWorkerId sets DistributedWorkerId field to given value.

### HasDistributedWorkerId

`func (o *ZoneVcenterConfig) HasDistributedWorkerId() bool`

HasDistributedWorkerId returns a boolean if a field has been set.

### SetDistributedWorkerIdNil

`func (o *ZoneVcenterConfig) SetDistributedWorkerIdNil(b bool)`

 SetDistributedWorkerIdNil sets the value for DistributedWorkerId to be an explicit nil

### UnsetDistributedWorkerId
`func (o *ZoneVcenterConfig) UnsetDistributedWorkerId()`

UnsetDistributedWorkerId ensures that no value is present for DistributedWorkerId, not even an explicit nil
### GetPasswordHash

`func (o *ZoneVcenterConfig) GetPasswordHash() string`

GetPasswordHash returns the PasswordHash field if non-nil, zero value otherwise.

### GetPasswordHashOk

`func (o *ZoneVcenterConfig) GetPasswordHashOk() (*string, bool)`

GetPasswordHashOk returns a tuple with the PasswordHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPasswordHash

`func (o *ZoneVcenterConfig) SetPasswordHash(v string)`

SetPasswordHash sets PasswordHash field to given value.

### HasPasswordHash

`func (o *ZoneVcenterConfig) HasPasswordHash() bool`

HasPasswordHash returns a boolean if a field has been set.

### SetPasswordHashNil

`func (o *ZoneVcenterConfig) SetPasswordHashNil(b bool)`

 SetPasswordHashNil sets the value for PasswordHash to be an explicit nil

### UnsetPasswordHash
`func (o *ZoneVcenterConfig) UnsetPasswordHash()`

UnsetPasswordHash ensures that no value is present for PasswordHash, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


