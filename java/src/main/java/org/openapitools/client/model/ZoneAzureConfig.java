/*
 * Morpheus API
 * Morpheus is a powerful cloud management tool that provides provisioning, monitoring, logging, backups, and application deployment strategies.  This document describes the Morpheus API protocol and the available endpoints. Sections are organized in the same manner as they appear in the Morpheus UI.
 *
 * The version of the OpenAPI document: 6.2.1
 * Contact: dev@morpheusdata.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */


package org.openapitools.client.model;

import java.util.Objects;
import java.util.Arrays;
import com.google.gson.TypeAdapter;
import com.google.gson.annotations.JsonAdapter;
import com.google.gson.annotations.SerializedName;
import com.google.gson.stream.JsonReader;
import com.google.gson.stream.JsonWriter;
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;
import java.io.IOException;
import org.openapitools.client.model.ZoneVcenterConfigNetworkServer;

/**
 * ZoneAzureConfig
 */
@javax.annotation.Generated(value = "org.openapitools.codegen.languages.JavaClientCodegen", date = "2023-08-21T11:26:47.028Z[GMT]")
public class ZoneAzureConfig {
  public static final String SERIALIZED_NAME_SUBSCRIBER_ID = "subscriberId";
  @SerializedName(SERIALIZED_NAME_SUBSCRIBER_ID)
  private String subscriberId;

  public static final String SERIALIZED_NAME_TENANT_ID = "tenantId";
  @SerializedName(SERIALIZED_NAME_TENANT_ID)
  private String tenantId;

  public static final String SERIALIZED_NAME_CLIENT_ID = "clientId";
  @SerializedName(SERIALIZED_NAME_CLIENT_ID)
  private String clientId;

  public static final String SERIALIZED_NAME_CLIENT_SECRET = "clientSecret";
  @SerializedName(SERIALIZED_NAME_CLIENT_SECRET)
  private String clientSecret;

  public static final String SERIALIZED_NAME_RESOURCE_GROUP = "resourceGroup";
  @SerializedName(SERIALIZED_NAME_RESOURCE_GROUP)
  private String resourceGroup;

  public static final String SERIALIZED_NAME_IMPORT_EXISTING = "_importExisting";
  @SerializedName(SERIALIZED_NAME_IMPORT_EXISTING)
  private String importExisting;

  public static final String SERIALIZED_NAME_IMPORT_EXISTING = "importExisting";
  @SerializedName(SERIALIZED_NAME_IMPORT_EXISTING)
  private String importExisting;

  public static final String SERIALIZED_NAME_INVENTORY_LEVEL = "inventoryLevel";
  @SerializedName(SERIALIZED_NAME_INVENTORY_LEVEL)
  private String inventoryLevel;

  public static final String SERIALIZED_NAME_APPLIANCE_URL = "applianceUrl";
  @SerializedName(SERIALIZED_NAME_APPLIANCE_URL)
  private String applianceUrl;

  public static final String SERIALIZED_NAME_DATACENTER_NAME = "datacenterName";
  @SerializedName(SERIALIZED_NAME_DATACENTER_NAME)
  private String datacenterName;

  public static final String SERIALIZED_NAME_NETWORK_SERVER_ID = "networkServer.id";
  @SerializedName(SERIALIZED_NAME_NETWORK_SERVER_ID)
  private String networkServerId;

  public static final String SERIALIZED_NAME_NETWORK_SERVER = "networkServer";
  @SerializedName(SERIALIZED_NAME_NETWORK_SERVER)
  private ZoneVcenterConfigNetworkServer networkServer;

  public static final String SERIALIZED_NAME_SECURITY_MODE = "securityMode";
  @SerializedName(SERIALIZED_NAME_SECURITY_MODE)
  private String securityMode;

  public static final String SERIALIZED_NAME_CERTIFICATE_PROVIDER = "certificateProvider";
  @SerializedName(SERIALIZED_NAME_CERTIFICATE_PROVIDER)
  private String certificateProvider;

  public static final String SERIALIZED_NAME_BACKUP_MODE = "backupMode";
  @SerializedName(SERIALIZED_NAME_BACKUP_MODE)
  private String backupMode;

  public static final String SERIALIZED_NAME_REPLICATION_MODE = "replicationMode";
  @SerializedName(SERIALIZED_NAME_REPLICATION_MODE)
  private String replicationMode;

  public static final String SERIALIZED_NAME_DNS_INTEGRATION_ID = "dnsIntegrationId";
  @SerializedName(SERIALIZED_NAME_DNS_INTEGRATION_ID)
  private String dnsIntegrationId;

  public static final String SERIALIZED_NAME_CONFIG_MANAGEMENT_ID = "configManagementId";
  @SerializedName(SERIALIZED_NAME_CONFIG_MANAGEMENT_ID)
  private String configManagementId;

  public static final String SERIALIZED_NAME_CONFIG_CMDB_ID = "configCmdbId";
  @SerializedName(SERIALIZED_NAME_CONFIG_CMDB_ID)
  private String configCmdbId;

  public static final String SERIALIZED_NAME_SECURITY_SERVER = "securityServer";
  @SerializedName(SERIALIZED_NAME_SECURITY_SERVER)
  private String securityServer;

  public static final String SERIALIZED_NAME_ACCOUNT_TYPE = "accountType";
  @SerializedName(SERIALIZED_NAME_ACCOUNT_TYPE)
  private String accountType;

  public static final String SERIALIZED_NAME_SERVICE_REGISTRY_ID = "serviceRegistryId";
  @SerializedName(SERIALIZED_NAME_SERVICE_REGISTRY_ID)
  private String serviceRegistryId;

  public static final String SERIALIZED_NAME_CLOUD_TYPE = "cloudType";
  @SerializedName(SERIALIZED_NAME_CLOUD_TYPE)
  private String cloudType;

  public static final String SERIALIZED_NAME_RPC_MODE = "rpcMode";
  @SerializedName(SERIALIZED_NAME_RPC_MODE)
  private String rpcMode;

  public static final String SERIALIZED_NAME_DISK_ENCRYPTION = "diskEncryption";
  @SerializedName(SERIALIZED_NAME_DISK_ENCRYPTION)
  private String diskEncryption;

  public static final String SERIALIZED_NAME_ENCRYPTION_SET = "encryptionSet";
  @SerializedName(SERIALIZED_NAME_ENCRYPTION_SET)
  private String encryptionSet;

  public static final String SERIALIZED_NAME_CSP_TENANT_ID = "cspTenantId";
  @SerializedName(SERIALIZED_NAME_CSP_TENANT_ID)
  private String cspTenantId;

  public static final String SERIALIZED_NAME_CSP_CLIENT_ID = "cspClientId";
  @SerializedName(SERIALIZED_NAME_CSP_CLIENT_ID)
  private String cspClientId;

  public static final String SERIALIZED_NAME_CSP_CLIENT_SECRET = "cspClientSecret";
  @SerializedName(SERIALIZED_NAME_CSP_CLIENT_SECRET)
  private String cspClientSecret;

  public static final String SERIALIZED_NAME_CSP_CUSTOMER = "cspCustomer";
  @SerializedName(SERIALIZED_NAME_CSP_CUSTOMER)
  private String cspCustomer;

  public static final String SERIALIZED_NAME_CONFIG_CMDB_DISCOVERY = "configCmdbDiscovery";
  @SerializedName(SERIALIZED_NAME_CONFIG_CMDB_DISCOVERY)
  private Boolean configCmdbDiscovery;

  public static final String SERIALIZED_NAME_AZURE_COSTING_MODE = "azureCostingMode";
  @SerializedName(SERIALIZED_NAME_AZURE_COSTING_MODE)
  private String azureCostingMode;

  public static final String SERIALIZED_NAME_CLIENT_SECRET_HASH = "clientSecretHash";
  @SerializedName(SERIALIZED_NAME_CLIENT_SECRET_HASH)
  private String clientSecretHash;

  public static final String SERIALIZED_NAME_CSP_CLIENT_SECRET_HASH = "cspClientSecretHash";
  @SerializedName(SERIALIZED_NAME_CSP_CLIENT_SECRET_HASH)
  private String cspClientSecretHash;


  public ZoneAzureConfig subscriberId(String subscriberId) {
    
    this.subscriberId = subscriberId;
    return this;
  }

   /**
   * Get subscriberId
   * @return subscriberId
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public String getSubscriberId() {
    return subscriberId;
  }


  public void setSubscriberId(String subscriberId) {
    this.subscriberId = subscriberId;
  }


  public ZoneAzureConfig tenantId(String tenantId) {
    
    this.tenantId = tenantId;
    return this;
  }

   /**
   * Get tenantId
   * @return tenantId
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public String getTenantId() {
    return tenantId;
  }


  public void setTenantId(String tenantId) {
    this.tenantId = tenantId;
  }


  public ZoneAzureConfig clientId(String clientId) {
    
    this.clientId = clientId;
    return this;
  }

   /**
   * Get clientId
   * @return clientId
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public String getClientId() {
    return clientId;
  }


  public void setClientId(String clientId) {
    this.clientId = clientId;
  }


  public ZoneAzureConfig clientSecret(String clientSecret) {
    
    this.clientSecret = clientSecret;
    return this;
  }

   /**
   * Get clientSecret
   * @return clientSecret
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public String getClientSecret() {
    return clientSecret;
  }


  public void setClientSecret(String clientSecret) {
    this.clientSecret = clientSecret;
  }


  public ZoneAzureConfig resourceGroup(String resourceGroup) {
    
    this.resourceGroup = resourceGroup;
    return this;
  }

   /**
   * Get resourceGroup
   * @return resourceGroup
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public String getResourceGroup() {
    return resourceGroup;
  }


  public void setResourceGroup(String resourceGroup) {
    this.resourceGroup = resourceGroup;
  }


  public ZoneAzureConfig importExisting(String importExisting) {
    
    this.importExisting = importExisting;
    return this;
  }

   /**
   * Get importExisting
   * @return importExisting
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public String getImportExisting() {
    return importExisting;
  }


  public void setImportExisting(String importExisting) {
    this.importExisting = importExisting;
  }


  public ZoneAzureConfig importExisting(String importExisting) {
    
    this.importExisting = importExisting;
    return this;
  }

   /**
   * Get importExisting
   * @return importExisting
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public String getImportExisting() {
    return importExisting;
  }


  public void setImportExisting(String importExisting) {
    this.importExisting = importExisting;
  }


  public ZoneAzureConfig inventoryLevel(String inventoryLevel) {
    
    this.inventoryLevel = inventoryLevel;
    return this;
  }

   /**
   * Get inventoryLevel
   * @return inventoryLevel
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public String getInventoryLevel() {
    return inventoryLevel;
  }


  public void setInventoryLevel(String inventoryLevel) {
    this.inventoryLevel = inventoryLevel;
  }


  public ZoneAzureConfig applianceUrl(String applianceUrl) {
    
    this.applianceUrl = applianceUrl;
    return this;
  }

   /**
   * Get applianceUrl
   * @return applianceUrl
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public String getApplianceUrl() {
    return applianceUrl;
  }


  public void setApplianceUrl(String applianceUrl) {
    this.applianceUrl = applianceUrl;
  }


  public ZoneAzureConfig datacenterName(String datacenterName) {
    
    this.datacenterName = datacenterName;
    return this;
  }

   /**
   * Get datacenterName
   * @return datacenterName
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public String getDatacenterName() {
    return datacenterName;
  }


  public void setDatacenterName(String datacenterName) {
    this.datacenterName = datacenterName;
  }


  public ZoneAzureConfig networkServerId(String networkServerId) {
    
    this.networkServerId = networkServerId;
    return this;
  }

   /**
   * Get networkServerId
   * @return networkServerId
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public String getNetworkServerId() {
    return networkServerId;
  }


  public void setNetworkServerId(String networkServerId) {
    this.networkServerId = networkServerId;
  }


  public ZoneAzureConfig networkServer(ZoneVcenterConfigNetworkServer networkServer) {
    
    this.networkServer = networkServer;
    return this;
  }

   /**
   * Get networkServer
   * @return networkServer
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public ZoneVcenterConfigNetworkServer getNetworkServer() {
    return networkServer;
  }


  public void setNetworkServer(ZoneVcenterConfigNetworkServer networkServer) {
    this.networkServer = networkServer;
  }


  public ZoneAzureConfig securityMode(String securityMode) {
    
    this.securityMode = securityMode;
    return this;
  }

   /**
   * Get securityMode
   * @return securityMode
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public String getSecurityMode() {
    return securityMode;
  }


  public void setSecurityMode(String securityMode) {
    this.securityMode = securityMode;
  }


  public ZoneAzureConfig certificateProvider(String certificateProvider) {
    
    this.certificateProvider = certificateProvider;
    return this;
  }

   /**
   * Get certificateProvider
   * @return certificateProvider
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public String getCertificateProvider() {
    return certificateProvider;
  }


  public void setCertificateProvider(String certificateProvider) {
    this.certificateProvider = certificateProvider;
  }


  public ZoneAzureConfig backupMode(String backupMode) {
    
    this.backupMode = backupMode;
    return this;
  }

   /**
   * Get backupMode
   * @return backupMode
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public String getBackupMode() {
    return backupMode;
  }


  public void setBackupMode(String backupMode) {
    this.backupMode = backupMode;
  }


  public ZoneAzureConfig replicationMode(String replicationMode) {
    
    this.replicationMode = replicationMode;
    return this;
  }

   /**
   * Get replicationMode
   * @return replicationMode
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public String getReplicationMode() {
    return replicationMode;
  }


  public void setReplicationMode(String replicationMode) {
    this.replicationMode = replicationMode;
  }


  public ZoneAzureConfig dnsIntegrationId(String dnsIntegrationId) {
    
    this.dnsIntegrationId = dnsIntegrationId;
    return this;
  }

   /**
   * Get dnsIntegrationId
   * @return dnsIntegrationId
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public String getDnsIntegrationId() {
    return dnsIntegrationId;
  }


  public void setDnsIntegrationId(String dnsIntegrationId) {
    this.dnsIntegrationId = dnsIntegrationId;
  }


  public ZoneAzureConfig configManagementId(String configManagementId) {
    
    this.configManagementId = configManagementId;
    return this;
  }

   /**
   * Get configManagementId
   * @return configManagementId
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public String getConfigManagementId() {
    return configManagementId;
  }


  public void setConfigManagementId(String configManagementId) {
    this.configManagementId = configManagementId;
  }


  public ZoneAzureConfig configCmdbId(String configCmdbId) {
    
    this.configCmdbId = configCmdbId;
    return this;
  }

   /**
   * Get configCmdbId
   * @return configCmdbId
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public String getConfigCmdbId() {
    return configCmdbId;
  }


  public void setConfigCmdbId(String configCmdbId) {
    this.configCmdbId = configCmdbId;
  }


  public ZoneAzureConfig securityServer(String securityServer) {
    
    this.securityServer = securityServer;
    return this;
  }

   /**
   * Get securityServer
   * @return securityServer
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public String getSecurityServer() {
    return securityServer;
  }


  public void setSecurityServer(String securityServer) {
    this.securityServer = securityServer;
  }


  public ZoneAzureConfig accountType(String accountType) {
    
    this.accountType = accountType;
    return this;
  }

   /**
   * Get accountType
   * @return accountType
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public String getAccountType() {
    return accountType;
  }


  public void setAccountType(String accountType) {
    this.accountType = accountType;
  }


  public ZoneAzureConfig serviceRegistryId(String serviceRegistryId) {
    
    this.serviceRegistryId = serviceRegistryId;
    return this;
  }

   /**
   * Get serviceRegistryId
   * @return serviceRegistryId
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public String getServiceRegistryId() {
    return serviceRegistryId;
  }


  public void setServiceRegistryId(String serviceRegistryId) {
    this.serviceRegistryId = serviceRegistryId;
  }


  public ZoneAzureConfig cloudType(String cloudType) {
    
    this.cloudType = cloudType;
    return this;
  }

   /**
   * Get cloudType
   * @return cloudType
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public String getCloudType() {
    return cloudType;
  }


  public void setCloudType(String cloudType) {
    this.cloudType = cloudType;
  }


  public ZoneAzureConfig rpcMode(String rpcMode) {
    
    this.rpcMode = rpcMode;
    return this;
  }

   /**
   * Get rpcMode
   * @return rpcMode
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public String getRpcMode() {
    return rpcMode;
  }


  public void setRpcMode(String rpcMode) {
    this.rpcMode = rpcMode;
  }


  public ZoneAzureConfig diskEncryption(String diskEncryption) {
    
    this.diskEncryption = diskEncryption;
    return this;
  }

   /**
   * Get diskEncryption
   * @return diskEncryption
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public String getDiskEncryption() {
    return diskEncryption;
  }


  public void setDiskEncryption(String diskEncryption) {
    this.diskEncryption = diskEncryption;
  }


  public ZoneAzureConfig encryptionSet(String encryptionSet) {
    
    this.encryptionSet = encryptionSet;
    return this;
  }

   /**
   * Get encryptionSet
   * @return encryptionSet
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public String getEncryptionSet() {
    return encryptionSet;
  }


  public void setEncryptionSet(String encryptionSet) {
    this.encryptionSet = encryptionSet;
  }


  public ZoneAzureConfig cspTenantId(String cspTenantId) {
    
    this.cspTenantId = cspTenantId;
    return this;
  }

   /**
   * Get cspTenantId
   * @return cspTenantId
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public String getCspTenantId() {
    return cspTenantId;
  }


  public void setCspTenantId(String cspTenantId) {
    this.cspTenantId = cspTenantId;
  }


  public ZoneAzureConfig cspClientId(String cspClientId) {
    
    this.cspClientId = cspClientId;
    return this;
  }

   /**
   * Get cspClientId
   * @return cspClientId
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public String getCspClientId() {
    return cspClientId;
  }


  public void setCspClientId(String cspClientId) {
    this.cspClientId = cspClientId;
  }


  public ZoneAzureConfig cspClientSecret(String cspClientSecret) {
    
    this.cspClientSecret = cspClientSecret;
    return this;
  }

   /**
   * Get cspClientSecret
   * @return cspClientSecret
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public String getCspClientSecret() {
    return cspClientSecret;
  }


  public void setCspClientSecret(String cspClientSecret) {
    this.cspClientSecret = cspClientSecret;
  }


  public ZoneAzureConfig cspCustomer(String cspCustomer) {
    
    this.cspCustomer = cspCustomer;
    return this;
  }

   /**
   * Get cspCustomer
   * @return cspCustomer
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public String getCspCustomer() {
    return cspCustomer;
  }


  public void setCspCustomer(String cspCustomer) {
    this.cspCustomer = cspCustomer;
  }


  public ZoneAzureConfig configCmdbDiscovery(Boolean configCmdbDiscovery) {
    
    this.configCmdbDiscovery = configCmdbDiscovery;
    return this;
  }

   /**
   * Get configCmdbDiscovery
   * @return configCmdbDiscovery
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public Boolean getConfigCmdbDiscovery() {
    return configCmdbDiscovery;
  }


  public void setConfigCmdbDiscovery(Boolean configCmdbDiscovery) {
    this.configCmdbDiscovery = configCmdbDiscovery;
  }


  public ZoneAzureConfig azureCostingMode(String azureCostingMode) {
    
    this.azureCostingMode = azureCostingMode;
    return this;
  }

   /**
   * Get azureCostingMode
   * @return azureCostingMode
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public String getAzureCostingMode() {
    return azureCostingMode;
  }


  public void setAzureCostingMode(String azureCostingMode) {
    this.azureCostingMode = azureCostingMode;
  }


  public ZoneAzureConfig clientSecretHash(String clientSecretHash) {
    
    this.clientSecretHash = clientSecretHash;
    return this;
  }

   /**
   * Get clientSecretHash
   * @return clientSecretHash
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public String getClientSecretHash() {
    return clientSecretHash;
  }


  public void setClientSecretHash(String clientSecretHash) {
    this.clientSecretHash = clientSecretHash;
  }


  public ZoneAzureConfig cspClientSecretHash(String cspClientSecretHash) {
    
    this.cspClientSecretHash = cspClientSecretHash;
    return this;
  }

   /**
   * Get cspClientSecretHash
   * @return cspClientSecretHash
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public String getCspClientSecretHash() {
    return cspClientSecretHash;
  }


  public void setCspClientSecretHash(String cspClientSecretHash) {
    this.cspClientSecretHash = cspClientSecretHash;
  }


  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (o == null || getClass() != o.getClass()) {
      return false;
    }
    ZoneAzureConfig zoneAzureConfig = (ZoneAzureConfig) o;
    return Objects.equals(this.subscriberId, zoneAzureConfig.subscriberId) &&
        Objects.equals(this.tenantId, zoneAzureConfig.tenantId) &&
        Objects.equals(this.clientId, zoneAzureConfig.clientId) &&
        Objects.equals(this.clientSecret, zoneAzureConfig.clientSecret) &&
        Objects.equals(this.resourceGroup, zoneAzureConfig.resourceGroup) &&
        Objects.equals(this.importExisting, zoneAzureConfig.importExisting) &&
        Objects.equals(this.importExisting, zoneAzureConfig.importExisting) &&
        Objects.equals(this.inventoryLevel, zoneAzureConfig.inventoryLevel) &&
        Objects.equals(this.applianceUrl, zoneAzureConfig.applianceUrl) &&
        Objects.equals(this.datacenterName, zoneAzureConfig.datacenterName) &&
        Objects.equals(this.networkServerId, zoneAzureConfig.networkServerId) &&
        Objects.equals(this.networkServer, zoneAzureConfig.networkServer) &&
        Objects.equals(this.securityMode, zoneAzureConfig.securityMode) &&
        Objects.equals(this.certificateProvider, zoneAzureConfig.certificateProvider) &&
        Objects.equals(this.backupMode, zoneAzureConfig.backupMode) &&
        Objects.equals(this.replicationMode, zoneAzureConfig.replicationMode) &&
        Objects.equals(this.dnsIntegrationId, zoneAzureConfig.dnsIntegrationId) &&
        Objects.equals(this.configManagementId, zoneAzureConfig.configManagementId) &&
        Objects.equals(this.configCmdbId, zoneAzureConfig.configCmdbId) &&
        Objects.equals(this.securityServer, zoneAzureConfig.securityServer) &&
        Objects.equals(this.accountType, zoneAzureConfig.accountType) &&
        Objects.equals(this.serviceRegistryId, zoneAzureConfig.serviceRegistryId) &&
        Objects.equals(this.cloudType, zoneAzureConfig.cloudType) &&
        Objects.equals(this.rpcMode, zoneAzureConfig.rpcMode) &&
        Objects.equals(this.diskEncryption, zoneAzureConfig.diskEncryption) &&
        Objects.equals(this.encryptionSet, zoneAzureConfig.encryptionSet) &&
        Objects.equals(this.cspTenantId, zoneAzureConfig.cspTenantId) &&
        Objects.equals(this.cspClientId, zoneAzureConfig.cspClientId) &&
        Objects.equals(this.cspClientSecret, zoneAzureConfig.cspClientSecret) &&
        Objects.equals(this.cspCustomer, zoneAzureConfig.cspCustomer) &&
        Objects.equals(this.configCmdbDiscovery, zoneAzureConfig.configCmdbDiscovery) &&
        Objects.equals(this.azureCostingMode, zoneAzureConfig.azureCostingMode) &&
        Objects.equals(this.clientSecretHash, zoneAzureConfig.clientSecretHash) &&
        Objects.equals(this.cspClientSecretHash, zoneAzureConfig.cspClientSecretHash);
  }

  @Override
  public int hashCode() {
    return Objects.hash(subscriberId, tenantId, clientId, clientSecret, resourceGroup, importExisting, importExisting, inventoryLevel, applianceUrl, datacenterName, networkServerId, networkServer, securityMode, certificateProvider, backupMode, replicationMode, dnsIntegrationId, configManagementId, configCmdbId, securityServer, accountType, serviceRegistryId, cloudType, rpcMode, diskEncryption, encryptionSet, cspTenantId, cspClientId, cspClientSecret, cspCustomer, configCmdbDiscovery, azureCostingMode, clientSecretHash, cspClientSecretHash);
  }


  @Override
  public String toString() {
    StringBuilder sb = new StringBuilder();
    sb.append("class ZoneAzureConfig {\n");
    sb.append("    subscriberId: ").append(toIndentedString(subscriberId)).append("\n");
    sb.append("    tenantId: ").append(toIndentedString(tenantId)).append("\n");
    sb.append("    clientId: ").append(toIndentedString(clientId)).append("\n");
    sb.append("    clientSecret: ").append(toIndentedString(clientSecret)).append("\n");
    sb.append("    resourceGroup: ").append(toIndentedString(resourceGroup)).append("\n");
    sb.append("    importExisting: ").append(toIndentedString(importExisting)).append("\n");
    sb.append("    importExisting: ").append(toIndentedString(importExisting)).append("\n");
    sb.append("    inventoryLevel: ").append(toIndentedString(inventoryLevel)).append("\n");
    sb.append("    applianceUrl: ").append(toIndentedString(applianceUrl)).append("\n");
    sb.append("    datacenterName: ").append(toIndentedString(datacenterName)).append("\n");
    sb.append("    networkServerId: ").append(toIndentedString(networkServerId)).append("\n");
    sb.append("    networkServer: ").append(toIndentedString(networkServer)).append("\n");
    sb.append("    securityMode: ").append(toIndentedString(securityMode)).append("\n");
    sb.append("    certificateProvider: ").append(toIndentedString(certificateProvider)).append("\n");
    sb.append("    backupMode: ").append(toIndentedString(backupMode)).append("\n");
    sb.append("    replicationMode: ").append(toIndentedString(replicationMode)).append("\n");
    sb.append("    dnsIntegrationId: ").append(toIndentedString(dnsIntegrationId)).append("\n");
    sb.append("    configManagementId: ").append(toIndentedString(configManagementId)).append("\n");
    sb.append("    configCmdbId: ").append(toIndentedString(configCmdbId)).append("\n");
    sb.append("    securityServer: ").append(toIndentedString(securityServer)).append("\n");
    sb.append("    accountType: ").append(toIndentedString(accountType)).append("\n");
    sb.append("    serviceRegistryId: ").append(toIndentedString(serviceRegistryId)).append("\n");
    sb.append("    cloudType: ").append(toIndentedString(cloudType)).append("\n");
    sb.append("    rpcMode: ").append(toIndentedString(rpcMode)).append("\n");
    sb.append("    diskEncryption: ").append(toIndentedString(diskEncryption)).append("\n");
    sb.append("    encryptionSet: ").append(toIndentedString(encryptionSet)).append("\n");
    sb.append("    cspTenantId: ").append(toIndentedString(cspTenantId)).append("\n");
    sb.append("    cspClientId: ").append(toIndentedString(cspClientId)).append("\n");
    sb.append("    cspClientSecret: ").append(toIndentedString(cspClientSecret)).append("\n");
    sb.append("    cspCustomer: ").append(toIndentedString(cspCustomer)).append("\n");
    sb.append("    configCmdbDiscovery: ").append(toIndentedString(configCmdbDiscovery)).append("\n");
    sb.append("    azureCostingMode: ").append(toIndentedString(azureCostingMode)).append("\n");
    sb.append("    clientSecretHash: ").append(toIndentedString(clientSecretHash)).append("\n");
    sb.append("    cspClientSecretHash: ").append(toIndentedString(cspClientSecretHash)).append("\n");
    sb.append("}");
    return sb.toString();
  }

  /**
   * Convert the given object to string with each line indented by 4 spaces
   * (except the first line).
   */
  private String toIndentedString(Object o) {
    if (o == null) {
      return "null";
    }
    return o.toString().replace("\n", "\n    ");
  }

}

