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
 * ZoneAwsConfig
 */
@javax.annotation.Generated(value = "org.openapitools.codegen.languages.JavaClientCodegen", date = "2023-08-21T11:26:47.028Z[GMT]")
public class ZoneAwsConfig {
  public static final String SERIALIZED_NAME_ENDPOINT = "endpoint";
  @SerializedName(SERIALIZED_NAME_ENDPOINT)
  private String endpoint;

  public static final String SERIALIZED_NAME_ACCESS_KEY = "accessKey";
  @SerializedName(SERIALIZED_NAME_ACCESS_KEY)
  private String accessKey;

  public static final String SERIALIZED_NAME_SECRET_KEY = "secretKey";
  @SerializedName(SERIALIZED_NAME_SECRET_KEY)
  private String secretKey;

  public static final String SERIALIZED_NAME_USE_HOST_CREDENTIALS = "_useHostCredentials";
  @SerializedName(SERIALIZED_NAME_USE_HOST_CREDENTIALS)
  private String useHostCredentials;

  public static final String SERIALIZED_NAME_STS_ASSUME_ROLE = "stsAssumeRole";
  @SerializedName(SERIALIZED_NAME_STS_ASSUME_ROLE)
  private String stsAssumeRole;

  public static final String SERIALIZED_NAME_IS_VPC = "isVpc";
  @SerializedName(SERIALIZED_NAME_IS_VPC)
  private String isVpc;

  public static final String SERIALIZED_NAME_VPC = "vpc";
  @SerializedName(SERIALIZED_NAME_VPC)
  private String vpc;

  public static final String SERIALIZED_NAME_IMAGE_STORE_ID = "imageStoreId";
  @SerializedName(SERIALIZED_NAME_IMAGE_STORE_ID)
  private String imageStoreId;

  public static final String SERIALIZED_NAME_EBS_ENCRYPTION = "ebsEncryption";
  @SerializedName(SERIALIZED_NAME_EBS_ENCRYPTION)
  private String ebsEncryption;

  public static final String SERIALIZED_NAME_COSTING_REPORT = "costingReport";
  @SerializedName(SERIALIZED_NAME_COSTING_REPORT)
  private String costingReport;

  public static final String SERIALIZED_NAME_COSTING_FOLDER = "costingFolder";
  @SerializedName(SERIALIZED_NAME_COSTING_FOLDER)
  private String costingFolder;

  public static final String SERIALIZED_NAME_COSTING_BUCKET = "costingBucket";
  @SerializedName(SERIALIZED_NAME_COSTING_BUCKET)
  private String costingBucket;

  public static final String SERIALIZED_NAME_COSTING_BUCKET_NAME = "costingBucketName";
  @SerializedName(SERIALIZED_NAME_COSTING_BUCKET_NAME)
  private String costingBucketName;

  public static final String SERIALIZED_NAME_COSTING_REGION = "costingRegion";
  @SerializedName(SERIALIZED_NAME_COSTING_REGION)
  private String costingRegion;

  public static final String SERIALIZED_NAME_COSTING_ACCESS_KEY = "costingAccessKey";
  @SerializedName(SERIALIZED_NAME_COSTING_ACCESS_KEY)
  private String costingAccessKey;

  public static final String SERIALIZED_NAME_COSTING_SECRET_KEY = "costingSecretKey";
  @SerializedName(SERIALIZED_NAME_COSTING_SECRET_KEY)
  private String costingSecretKey;

  public static final String SERIALIZED_NAME_COSTING_REPORT_NAME = "costingReportName";
  @SerializedName(SERIALIZED_NAME_COSTING_REPORT_NAME)
  private String costingReportName;

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

  public static final String SERIALIZED_NAME_SECURITY_SERVER = "securityServer";
  @SerializedName(SERIALIZED_NAME_SECURITY_SERVER)
  private String securityServer;

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

  public static final String SERIALIZED_NAME_SERVICE_REGISTRY_ID = "serviceRegistryId";
  @SerializedName(SERIALIZED_NAME_SERVICE_REGISTRY_ID)
  private String serviceRegistryId;

  public static final String SERIALIZED_NAME_CONFIG_MANAGEMENT_ID = "configManagementId";
  @SerializedName(SERIALIZED_NAME_CONFIG_MANAGEMENT_ID)
  private String configManagementId;

  public static final String SERIALIZED_NAME_CONFIG_CMDB_DISCOVERY = "configCmdbDiscovery";
  @SerializedName(SERIALIZED_NAME_CONFIG_CMDB_DISCOVERY)
  private Boolean configCmdbDiscovery;

  public static final String SERIALIZED_NAME_SECRET_KEY_HASH = "secretKeyHash";
  @SerializedName(SERIALIZED_NAME_SECRET_KEY_HASH)
  private String secretKeyHash;

  public static final String SERIALIZED_NAME_COSTING_SECRET_KEY_HASH = "costingSecretKeyHash";
  @SerializedName(SERIALIZED_NAME_COSTING_SECRET_KEY_HASH)
  private String costingSecretKeyHash;


  public ZoneAwsConfig endpoint(String endpoint) {
    
    this.endpoint = endpoint;
    return this;
  }

   /**
   * Get endpoint
   * @return endpoint
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public String getEndpoint() {
    return endpoint;
  }


  public void setEndpoint(String endpoint) {
    this.endpoint = endpoint;
  }


  public ZoneAwsConfig accessKey(String accessKey) {
    
    this.accessKey = accessKey;
    return this;
  }

   /**
   * Get accessKey
   * @return accessKey
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public String getAccessKey() {
    return accessKey;
  }


  public void setAccessKey(String accessKey) {
    this.accessKey = accessKey;
  }


  public ZoneAwsConfig secretKey(String secretKey) {
    
    this.secretKey = secretKey;
    return this;
  }

   /**
   * Get secretKey
   * @return secretKey
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public String getSecretKey() {
    return secretKey;
  }


  public void setSecretKey(String secretKey) {
    this.secretKey = secretKey;
  }


  public ZoneAwsConfig useHostCredentials(String useHostCredentials) {
    
    this.useHostCredentials = useHostCredentials;
    return this;
  }

   /**
   * Get useHostCredentials
   * @return useHostCredentials
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public String getUseHostCredentials() {
    return useHostCredentials;
  }


  public void setUseHostCredentials(String useHostCredentials) {
    this.useHostCredentials = useHostCredentials;
  }


  public ZoneAwsConfig stsAssumeRole(String stsAssumeRole) {
    
    this.stsAssumeRole = stsAssumeRole;
    return this;
  }

   /**
   * Get stsAssumeRole
   * @return stsAssumeRole
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public String getStsAssumeRole() {
    return stsAssumeRole;
  }


  public void setStsAssumeRole(String stsAssumeRole) {
    this.stsAssumeRole = stsAssumeRole;
  }


  public ZoneAwsConfig isVpc(String isVpc) {
    
    this.isVpc = isVpc;
    return this;
  }

   /**
   * Get isVpc
   * @return isVpc
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public String getIsVpc() {
    return isVpc;
  }


  public void setIsVpc(String isVpc) {
    this.isVpc = isVpc;
  }


  public ZoneAwsConfig vpc(String vpc) {
    
    this.vpc = vpc;
    return this;
  }

   /**
   * Get vpc
   * @return vpc
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public String getVpc() {
    return vpc;
  }


  public void setVpc(String vpc) {
    this.vpc = vpc;
  }


  public ZoneAwsConfig imageStoreId(String imageStoreId) {
    
    this.imageStoreId = imageStoreId;
    return this;
  }

   /**
   * Get imageStoreId
   * @return imageStoreId
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public String getImageStoreId() {
    return imageStoreId;
  }


  public void setImageStoreId(String imageStoreId) {
    this.imageStoreId = imageStoreId;
  }


  public ZoneAwsConfig ebsEncryption(String ebsEncryption) {
    
    this.ebsEncryption = ebsEncryption;
    return this;
  }

   /**
   * Get ebsEncryption
   * @return ebsEncryption
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public String getEbsEncryption() {
    return ebsEncryption;
  }


  public void setEbsEncryption(String ebsEncryption) {
    this.ebsEncryption = ebsEncryption;
  }


  public ZoneAwsConfig costingReport(String costingReport) {
    
    this.costingReport = costingReport;
    return this;
  }

   /**
   * Get costingReport
   * @return costingReport
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public String getCostingReport() {
    return costingReport;
  }


  public void setCostingReport(String costingReport) {
    this.costingReport = costingReport;
  }


  public ZoneAwsConfig costingFolder(String costingFolder) {
    
    this.costingFolder = costingFolder;
    return this;
  }

   /**
   * Get costingFolder
   * @return costingFolder
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public String getCostingFolder() {
    return costingFolder;
  }


  public void setCostingFolder(String costingFolder) {
    this.costingFolder = costingFolder;
  }


  public ZoneAwsConfig costingBucket(String costingBucket) {
    
    this.costingBucket = costingBucket;
    return this;
  }

   /**
   * Get costingBucket
   * @return costingBucket
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public String getCostingBucket() {
    return costingBucket;
  }


  public void setCostingBucket(String costingBucket) {
    this.costingBucket = costingBucket;
  }


  public ZoneAwsConfig costingBucketName(String costingBucketName) {
    
    this.costingBucketName = costingBucketName;
    return this;
  }

   /**
   * Get costingBucketName
   * @return costingBucketName
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public String getCostingBucketName() {
    return costingBucketName;
  }


  public void setCostingBucketName(String costingBucketName) {
    this.costingBucketName = costingBucketName;
  }


  public ZoneAwsConfig costingRegion(String costingRegion) {
    
    this.costingRegion = costingRegion;
    return this;
  }

   /**
   * Get costingRegion
   * @return costingRegion
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public String getCostingRegion() {
    return costingRegion;
  }


  public void setCostingRegion(String costingRegion) {
    this.costingRegion = costingRegion;
  }


  public ZoneAwsConfig costingAccessKey(String costingAccessKey) {
    
    this.costingAccessKey = costingAccessKey;
    return this;
  }

   /**
   * Get costingAccessKey
   * @return costingAccessKey
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public String getCostingAccessKey() {
    return costingAccessKey;
  }


  public void setCostingAccessKey(String costingAccessKey) {
    this.costingAccessKey = costingAccessKey;
  }


  public ZoneAwsConfig costingSecretKey(String costingSecretKey) {
    
    this.costingSecretKey = costingSecretKey;
    return this;
  }

   /**
   * Get costingSecretKey
   * @return costingSecretKey
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public String getCostingSecretKey() {
    return costingSecretKey;
  }


  public void setCostingSecretKey(String costingSecretKey) {
    this.costingSecretKey = costingSecretKey;
  }


  public ZoneAwsConfig costingReportName(String costingReportName) {
    
    this.costingReportName = costingReportName;
    return this;
  }

   /**
   * Get costingReportName
   * @return costingReportName
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public String getCostingReportName() {
    return costingReportName;
  }


  public void setCostingReportName(String costingReportName) {
    this.costingReportName = costingReportName;
  }


  public ZoneAwsConfig applianceUrl(String applianceUrl) {
    
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


  public ZoneAwsConfig datacenterName(String datacenterName) {
    
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


  public ZoneAwsConfig networkServerId(String networkServerId) {
    
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


  public ZoneAwsConfig networkServer(ZoneVcenterConfigNetworkServer networkServer) {
    
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


  public ZoneAwsConfig securityServer(String securityServer) {
    
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


  public ZoneAwsConfig certificateProvider(String certificateProvider) {
    
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


  public ZoneAwsConfig backupMode(String backupMode) {
    
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


  public ZoneAwsConfig replicationMode(String replicationMode) {
    
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


  public ZoneAwsConfig dnsIntegrationId(String dnsIntegrationId) {
    
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


  public ZoneAwsConfig serviceRegistryId(String serviceRegistryId) {
    
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


  public ZoneAwsConfig configManagementId(String configManagementId) {
    
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


  public ZoneAwsConfig configCmdbDiscovery(Boolean configCmdbDiscovery) {
    
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


  public ZoneAwsConfig secretKeyHash(String secretKeyHash) {
    
    this.secretKeyHash = secretKeyHash;
    return this;
  }

   /**
   * Get secretKeyHash
   * @return secretKeyHash
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public String getSecretKeyHash() {
    return secretKeyHash;
  }


  public void setSecretKeyHash(String secretKeyHash) {
    this.secretKeyHash = secretKeyHash;
  }


  public ZoneAwsConfig costingSecretKeyHash(String costingSecretKeyHash) {
    
    this.costingSecretKeyHash = costingSecretKeyHash;
    return this;
  }

   /**
   * Get costingSecretKeyHash
   * @return costingSecretKeyHash
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public String getCostingSecretKeyHash() {
    return costingSecretKeyHash;
  }


  public void setCostingSecretKeyHash(String costingSecretKeyHash) {
    this.costingSecretKeyHash = costingSecretKeyHash;
  }


  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (o == null || getClass() != o.getClass()) {
      return false;
    }
    ZoneAwsConfig zoneAwsConfig = (ZoneAwsConfig) o;
    return Objects.equals(this.endpoint, zoneAwsConfig.endpoint) &&
        Objects.equals(this.accessKey, zoneAwsConfig.accessKey) &&
        Objects.equals(this.secretKey, zoneAwsConfig.secretKey) &&
        Objects.equals(this.useHostCredentials, zoneAwsConfig.useHostCredentials) &&
        Objects.equals(this.stsAssumeRole, zoneAwsConfig.stsAssumeRole) &&
        Objects.equals(this.isVpc, zoneAwsConfig.isVpc) &&
        Objects.equals(this.vpc, zoneAwsConfig.vpc) &&
        Objects.equals(this.imageStoreId, zoneAwsConfig.imageStoreId) &&
        Objects.equals(this.ebsEncryption, zoneAwsConfig.ebsEncryption) &&
        Objects.equals(this.costingReport, zoneAwsConfig.costingReport) &&
        Objects.equals(this.costingFolder, zoneAwsConfig.costingFolder) &&
        Objects.equals(this.costingBucket, zoneAwsConfig.costingBucket) &&
        Objects.equals(this.costingBucketName, zoneAwsConfig.costingBucketName) &&
        Objects.equals(this.costingRegion, zoneAwsConfig.costingRegion) &&
        Objects.equals(this.costingAccessKey, zoneAwsConfig.costingAccessKey) &&
        Objects.equals(this.costingSecretKey, zoneAwsConfig.costingSecretKey) &&
        Objects.equals(this.costingReportName, zoneAwsConfig.costingReportName) &&
        Objects.equals(this.applianceUrl, zoneAwsConfig.applianceUrl) &&
        Objects.equals(this.datacenterName, zoneAwsConfig.datacenterName) &&
        Objects.equals(this.networkServerId, zoneAwsConfig.networkServerId) &&
        Objects.equals(this.networkServer, zoneAwsConfig.networkServer) &&
        Objects.equals(this.securityServer, zoneAwsConfig.securityServer) &&
        Objects.equals(this.certificateProvider, zoneAwsConfig.certificateProvider) &&
        Objects.equals(this.backupMode, zoneAwsConfig.backupMode) &&
        Objects.equals(this.replicationMode, zoneAwsConfig.replicationMode) &&
        Objects.equals(this.dnsIntegrationId, zoneAwsConfig.dnsIntegrationId) &&
        Objects.equals(this.serviceRegistryId, zoneAwsConfig.serviceRegistryId) &&
        Objects.equals(this.configManagementId, zoneAwsConfig.configManagementId) &&
        Objects.equals(this.configCmdbDiscovery, zoneAwsConfig.configCmdbDiscovery) &&
        Objects.equals(this.secretKeyHash, zoneAwsConfig.secretKeyHash) &&
        Objects.equals(this.costingSecretKeyHash, zoneAwsConfig.costingSecretKeyHash);
  }

  @Override
  public int hashCode() {
    return Objects.hash(endpoint, accessKey, secretKey, useHostCredentials, stsAssumeRole, isVpc, vpc, imageStoreId, ebsEncryption, costingReport, costingFolder, costingBucket, costingBucketName, costingRegion, costingAccessKey, costingSecretKey, costingReportName, applianceUrl, datacenterName, networkServerId, networkServer, securityServer, certificateProvider, backupMode, replicationMode, dnsIntegrationId, serviceRegistryId, configManagementId, configCmdbDiscovery, secretKeyHash, costingSecretKeyHash);
  }


  @Override
  public String toString() {
    StringBuilder sb = new StringBuilder();
    sb.append("class ZoneAwsConfig {\n");
    sb.append("    endpoint: ").append(toIndentedString(endpoint)).append("\n");
    sb.append("    accessKey: ").append(toIndentedString(accessKey)).append("\n");
    sb.append("    secretKey: ").append(toIndentedString(secretKey)).append("\n");
    sb.append("    useHostCredentials: ").append(toIndentedString(useHostCredentials)).append("\n");
    sb.append("    stsAssumeRole: ").append(toIndentedString(stsAssumeRole)).append("\n");
    sb.append("    isVpc: ").append(toIndentedString(isVpc)).append("\n");
    sb.append("    vpc: ").append(toIndentedString(vpc)).append("\n");
    sb.append("    imageStoreId: ").append(toIndentedString(imageStoreId)).append("\n");
    sb.append("    ebsEncryption: ").append(toIndentedString(ebsEncryption)).append("\n");
    sb.append("    costingReport: ").append(toIndentedString(costingReport)).append("\n");
    sb.append("    costingFolder: ").append(toIndentedString(costingFolder)).append("\n");
    sb.append("    costingBucket: ").append(toIndentedString(costingBucket)).append("\n");
    sb.append("    costingBucketName: ").append(toIndentedString(costingBucketName)).append("\n");
    sb.append("    costingRegion: ").append(toIndentedString(costingRegion)).append("\n");
    sb.append("    costingAccessKey: ").append(toIndentedString(costingAccessKey)).append("\n");
    sb.append("    costingSecretKey: ").append(toIndentedString(costingSecretKey)).append("\n");
    sb.append("    costingReportName: ").append(toIndentedString(costingReportName)).append("\n");
    sb.append("    applianceUrl: ").append(toIndentedString(applianceUrl)).append("\n");
    sb.append("    datacenterName: ").append(toIndentedString(datacenterName)).append("\n");
    sb.append("    networkServerId: ").append(toIndentedString(networkServerId)).append("\n");
    sb.append("    networkServer: ").append(toIndentedString(networkServer)).append("\n");
    sb.append("    securityServer: ").append(toIndentedString(securityServer)).append("\n");
    sb.append("    certificateProvider: ").append(toIndentedString(certificateProvider)).append("\n");
    sb.append("    backupMode: ").append(toIndentedString(backupMode)).append("\n");
    sb.append("    replicationMode: ").append(toIndentedString(replicationMode)).append("\n");
    sb.append("    dnsIntegrationId: ").append(toIndentedString(dnsIntegrationId)).append("\n");
    sb.append("    serviceRegistryId: ").append(toIndentedString(serviceRegistryId)).append("\n");
    sb.append("    configManagementId: ").append(toIndentedString(configManagementId)).append("\n");
    sb.append("    configCmdbDiscovery: ").append(toIndentedString(configCmdbDiscovery)).append("\n");
    sb.append("    secretKeyHash: ").append(toIndentedString(secretKeyHash)).append("\n");
    sb.append("    costingSecretKeyHash: ").append(toIndentedString(costingSecretKeyHash)).append("\n");
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

