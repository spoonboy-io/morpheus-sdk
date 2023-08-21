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

/**
 * ServerConfig
 */
@javax.annotation.Generated(value = "org.openapitools.codegen.languages.JavaClientCodegen", date = "2023-08-21T11:26:47.028Z[GMT]")
public class ServerConfig {
  public static final String SERIALIZED_NAME_POOL_PROVIDER_TYPE = "poolProviderType";
  @SerializedName(SERIALIZED_NAME_POOL_PROVIDER_TYPE)
  private String poolProviderType;

  public static final String SERIALIZED_NAME_IS_VPC_SELECTABLE = "isVpcSelectable";
  @SerializedName(SERIALIZED_NAME_IS_VPC_SELECTABLE)
  private Boolean isVpcSelectable;

  public static final String SERIALIZED_NAME_SMBIOS_ASSET_TAG = "smbiosAssetTag";
  @SerializedName(SERIALIZED_NAME_SMBIOS_ASSET_TAG)
  private String smbiosAssetTag;

  public static final String SERIALIZED_NAME_IS_E_C2 = "isEC2";
  @SerializedName(SERIALIZED_NAME_IS_E_C2)
  private Boolean isEC2;

  public static final String SERIALIZED_NAME_RESOURCE_POOL_ID = "resourcePoolId";
  @SerializedName(SERIALIZED_NAME_RESOURCE_POOL_ID)
  private Long resourcePoolId;

  public static final String SERIALIZED_NAME_HOST_ID = "hostId";
  @SerializedName(SERIALIZED_NAME_HOST_ID)
  private Long hostId;

  public static final String SERIALIZED_NAME_CREATE_USER = "createUser";
  @SerializedName(SERIALIZED_NAME_CREATE_USER)
  private Boolean createUser;

  public static final String SERIALIZED_NAME_NESTED_VIRTUALIZATION = "nestedVirtualization";
  @SerializedName(SERIALIZED_NAME_NESTED_VIRTUALIZATION)
  private String nestedVirtualization;

  public static final String SERIALIZED_NAME_VMWARE_FOLDER_ID = "vmwareFolderId";
  @SerializedName(SERIALIZED_NAME_VMWARE_FOLDER_ID)
  private String vmwareFolderId;

  public static final String SERIALIZED_NAME_NO_AGENT = "noAgent";
  @SerializedName(SERIALIZED_NAME_NO_AGENT)
  private Boolean noAgent;

  public static final String SERIALIZED_NAME_POWER_SCHEDULE_TYPE = "powerScheduleType";
  @SerializedName(SERIALIZED_NAME_POWER_SCHEDULE_TYPE)
  private Long powerScheduleType;


  public ServerConfig poolProviderType(String poolProviderType) {
    
    this.poolProviderType = poolProviderType;
    return this;
  }

   /**
   * Get poolProviderType
   * @return poolProviderType
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public String getPoolProviderType() {
    return poolProviderType;
  }


  public void setPoolProviderType(String poolProviderType) {
    this.poolProviderType = poolProviderType;
  }


  public ServerConfig isVpcSelectable(Boolean isVpcSelectable) {
    
    this.isVpcSelectable = isVpcSelectable;
    return this;
  }

   /**
   * Get isVpcSelectable
   * @return isVpcSelectable
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public Boolean getIsVpcSelectable() {
    return isVpcSelectable;
  }


  public void setIsVpcSelectable(Boolean isVpcSelectable) {
    this.isVpcSelectable = isVpcSelectable;
  }


  public ServerConfig smbiosAssetTag(String smbiosAssetTag) {
    
    this.smbiosAssetTag = smbiosAssetTag;
    return this;
  }

   /**
   * Get smbiosAssetTag
   * @return smbiosAssetTag
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public String getSmbiosAssetTag() {
    return smbiosAssetTag;
  }


  public void setSmbiosAssetTag(String smbiosAssetTag) {
    this.smbiosAssetTag = smbiosAssetTag;
  }


  public ServerConfig isEC2(Boolean isEC2) {
    
    this.isEC2 = isEC2;
    return this;
  }

   /**
   * Get isEC2
   * @return isEC2
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public Boolean getIsEC2() {
    return isEC2;
  }


  public void setIsEC2(Boolean isEC2) {
    this.isEC2 = isEC2;
  }


  public ServerConfig resourcePoolId(Long resourcePoolId) {
    
    this.resourcePoolId = resourcePoolId;
    return this;
  }

   /**
   * Get resourcePoolId
   * @return resourcePoolId
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public Long getResourcePoolId() {
    return resourcePoolId;
  }


  public void setResourcePoolId(Long resourcePoolId) {
    this.resourcePoolId = resourcePoolId;
  }


  public ServerConfig hostId(Long hostId) {
    
    this.hostId = hostId;
    return this;
  }

   /**
   * Get hostId
   * @return hostId
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public Long getHostId() {
    return hostId;
  }


  public void setHostId(Long hostId) {
    this.hostId = hostId;
  }


  public ServerConfig createUser(Boolean createUser) {
    
    this.createUser = createUser;
    return this;
  }

   /**
   * Get createUser
   * @return createUser
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public Boolean getCreateUser() {
    return createUser;
  }


  public void setCreateUser(Boolean createUser) {
    this.createUser = createUser;
  }


  public ServerConfig nestedVirtualization(String nestedVirtualization) {
    
    this.nestedVirtualization = nestedVirtualization;
    return this;
  }

   /**
   * Get nestedVirtualization
   * @return nestedVirtualization
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public String getNestedVirtualization() {
    return nestedVirtualization;
  }


  public void setNestedVirtualization(String nestedVirtualization) {
    this.nestedVirtualization = nestedVirtualization;
  }


  public ServerConfig vmwareFolderId(String vmwareFolderId) {
    
    this.vmwareFolderId = vmwareFolderId;
    return this;
  }

   /**
   * Get vmwareFolderId
   * @return vmwareFolderId
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public String getVmwareFolderId() {
    return vmwareFolderId;
  }


  public void setVmwareFolderId(String vmwareFolderId) {
    this.vmwareFolderId = vmwareFolderId;
  }


  public ServerConfig noAgent(Boolean noAgent) {
    
    this.noAgent = noAgent;
    return this;
  }

   /**
   * Get noAgent
   * @return noAgent
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public Boolean getNoAgent() {
    return noAgent;
  }


  public void setNoAgent(Boolean noAgent) {
    this.noAgent = noAgent;
  }


  public ServerConfig powerScheduleType(Long powerScheduleType) {
    
    this.powerScheduleType = powerScheduleType;
    return this;
  }

   /**
   * Get powerScheduleType
   * @return powerScheduleType
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public Long getPowerScheduleType() {
    return powerScheduleType;
  }


  public void setPowerScheduleType(Long powerScheduleType) {
    this.powerScheduleType = powerScheduleType;
  }


  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (o == null || getClass() != o.getClass()) {
      return false;
    }
    ServerConfig serverConfig = (ServerConfig) o;
    return Objects.equals(this.poolProviderType, serverConfig.poolProviderType) &&
        Objects.equals(this.isVpcSelectable, serverConfig.isVpcSelectable) &&
        Objects.equals(this.smbiosAssetTag, serverConfig.smbiosAssetTag) &&
        Objects.equals(this.isEC2, serverConfig.isEC2) &&
        Objects.equals(this.resourcePoolId, serverConfig.resourcePoolId) &&
        Objects.equals(this.hostId, serverConfig.hostId) &&
        Objects.equals(this.createUser, serverConfig.createUser) &&
        Objects.equals(this.nestedVirtualization, serverConfig.nestedVirtualization) &&
        Objects.equals(this.vmwareFolderId, serverConfig.vmwareFolderId) &&
        Objects.equals(this.noAgent, serverConfig.noAgent) &&
        Objects.equals(this.powerScheduleType, serverConfig.powerScheduleType);
  }

  @Override
  public int hashCode() {
    return Objects.hash(poolProviderType, isVpcSelectable, smbiosAssetTag, isEC2, resourcePoolId, hostId, createUser, nestedVirtualization, vmwareFolderId, noAgent, powerScheduleType);
  }


  @Override
  public String toString() {
    StringBuilder sb = new StringBuilder();
    sb.append("class ServerConfig {\n");
    sb.append("    poolProviderType: ").append(toIndentedString(poolProviderType)).append("\n");
    sb.append("    isVpcSelectable: ").append(toIndentedString(isVpcSelectable)).append("\n");
    sb.append("    smbiosAssetTag: ").append(toIndentedString(smbiosAssetTag)).append("\n");
    sb.append("    isEC2: ").append(toIndentedString(isEC2)).append("\n");
    sb.append("    resourcePoolId: ").append(toIndentedString(resourcePoolId)).append("\n");
    sb.append("    hostId: ").append(toIndentedString(hostId)).append("\n");
    sb.append("    createUser: ").append(toIndentedString(createUser)).append("\n");
    sb.append("    nestedVirtualization: ").append(toIndentedString(nestedVirtualization)).append("\n");
    sb.append("    vmwareFolderId: ").append(toIndentedString(vmwareFolderId)).append("\n");
    sb.append("    noAgent: ").append(toIndentedString(noAgent)).append("\n");
    sb.append("    powerScheduleType: ").append(toIndentedString(powerScheduleType)).append("\n");
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

