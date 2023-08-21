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
import org.openapitools.client.model.ClusterLayoutComputeServerType;
import org.openapitools.client.model.ClusterLayoutContainerType;

/**
 * ClusterLayoutComputeServers
 */
@javax.annotation.Generated(value = "org.openapitools.codegen.languages.JavaClientCodegen", date = "2023-08-21T11:26:47.028Z[GMT]")
public class ClusterLayoutComputeServers {
  public static final String SERIALIZED_NAME_ID = "id";
  @SerializedName(SERIALIZED_NAME_ID)
  private Long id;

  public static final String SERIALIZED_NAME_PRIORITY_ORDER = "priorityOrder";
  @SerializedName(SERIALIZED_NAME_PRIORITY_ORDER)
  private Long priorityOrder;

  public static final String SERIALIZED_NAME_NODE_COUNT = "nodeCount";
  @SerializedName(SERIALIZED_NAME_NODE_COUNT)
  private Long nodeCount;

  public static final String SERIALIZED_NAME_NODE_TYPE = "nodeType";
  @SerializedName(SERIALIZED_NAME_NODE_TYPE)
  private String nodeType;

  public static final String SERIALIZED_NAME_MIN_NODE_COUNT = "minNodeCount";
  @SerializedName(SERIALIZED_NAME_MIN_NODE_COUNT)
  private Long minNodeCount;

  public static final String SERIALIZED_NAME_MAX_NODE_COUNT = "maxNodeCount";
  @SerializedName(SERIALIZED_NAME_MAX_NODE_COUNT)
  private String maxNodeCount;

  public static final String SERIALIZED_NAME_DYNAMIC_COUNT = "dynamicCount";
  @SerializedName(SERIALIZED_NAME_DYNAMIC_COUNT)
  private Boolean dynamicCount;

  public static final String SERIALIZED_NAME_INSTALL_CONTAINER_RUNTIME = "installContainerRuntime";
  @SerializedName(SERIALIZED_NAME_INSTALL_CONTAINER_RUNTIME)
  private Boolean installContainerRuntime;

  public static final String SERIALIZED_NAME_INSTALL_STORAGE_RUNTIME = "installStorageRuntime";
  @SerializedName(SERIALIZED_NAME_INSTALL_STORAGE_RUNTIME)
  private Boolean installStorageRuntime;

  public static final String SERIALIZED_NAME_NAME = "name";
  @SerializedName(SERIALIZED_NAME_NAME)
  private String name;

  public static final String SERIALIZED_NAME_CODE = "code";
  @SerializedName(SERIALIZED_NAME_CODE)
  private String code;

  public static final String SERIALIZED_NAME_CATEGORY = "category";
  @SerializedName(SERIALIZED_NAME_CATEGORY)
  private String category;

  public static final String SERIALIZED_NAME_CONFIG = "config";
  @SerializedName(SERIALIZED_NAME_CONFIG)
  private String config;

  public static final String SERIALIZED_NAME_CONTAINER_TYPE = "containerType";
  @SerializedName(SERIALIZED_NAME_CONTAINER_TYPE)
  private ClusterLayoutContainerType containerType;

  public static final String SERIALIZED_NAME_COMPUTE_SERVER_TYPE = "computeServerType";
  @SerializedName(SERIALIZED_NAME_COMPUTE_SERVER_TYPE)
  private ClusterLayoutComputeServerType computeServerType;

  public static final String SERIALIZED_NAME_PROVISION_SERVICE = "provisionService";
  @SerializedName(SERIALIZED_NAME_PROVISION_SERVICE)
  private String provisionService;

  public static final String SERIALIZED_NAME_PLAN_CATEGORY = "planCategory";
  @SerializedName(SERIALIZED_NAME_PLAN_CATEGORY)
  private String planCategory;

  public static final String SERIALIZED_NAME_NAME_PREFIX = "namePrefix";
  @SerializedName(SERIALIZED_NAME_NAME_PREFIX)
  private String namePrefix;

  public static final String SERIALIZED_NAME_NAME_SUFFIX = "nameSuffix";
  @SerializedName(SERIALIZED_NAME_NAME_SUFFIX)
  private String nameSuffix;

  public static final String SERIALIZED_NAME_FORCE_NAME_INDEX = "forceNameIndex";
  @SerializedName(SERIALIZED_NAME_FORCE_NAME_INDEX)
  private Boolean forceNameIndex;

  public static final String SERIALIZED_NAME_LOAD_BALANCE = "loadBalance";
  @SerializedName(SERIALIZED_NAME_LOAD_BALANCE)
  private Boolean loadBalance;


  public ClusterLayoutComputeServers id(Long id) {
    
    this.id = id;
    return this;
  }

   /**
   * Get id
   * @return id
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public Long getId() {
    return id;
  }


  public void setId(Long id) {
    this.id = id;
  }


  public ClusterLayoutComputeServers priorityOrder(Long priorityOrder) {
    
    this.priorityOrder = priorityOrder;
    return this;
  }

   /**
   * Get priorityOrder
   * @return priorityOrder
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public Long getPriorityOrder() {
    return priorityOrder;
  }


  public void setPriorityOrder(Long priorityOrder) {
    this.priorityOrder = priorityOrder;
  }


  public ClusterLayoutComputeServers nodeCount(Long nodeCount) {
    
    this.nodeCount = nodeCount;
    return this;
  }

   /**
   * Get nodeCount
   * @return nodeCount
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public Long getNodeCount() {
    return nodeCount;
  }


  public void setNodeCount(Long nodeCount) {
    this.nodeCount = nodeCount;
  }


  public ClusterLayoutComputeServers nodeType(String nodeType) {
    
    this.nodeType = nodeType;
    return this;
  }

   /**
   * Get nodeType
   * @return nodeType
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public String getNodeType() {
    return nodeType;
  }


  public void setNodeType(String nodeType) {
    this.nodeType = nodeType;
  }


  public ClusterLayoutComputeServers minNodeCount(Long minNodeCount) {
    
    this.minNodeCount = minNodeCount;
    return this;
  }

   /**
   * Get minNodeCount
   * @return minNodeCount
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public Long getMinNodeCount() {
    return minNodeCount;
  }


  public void setMinNodeCount(Long minNodeCount) {
    this.minNodeCount = minNodeCount;
  }


  public ClusterLayoutComputeServers maxNodeCount(String maxNodeCount) {
    
    this.maxNodeCount = maxNodeCount;
    return this;
  }

   /**
   * Get maxNodeCount
   * @return maxNodeCount
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public String getMaxNodeCount() {
    return maxNodeCount;
  }


  public void setMaxNodeCount(String maxNodeCount) {
    this.maxNodeCount = maxNodeCount;
  }


  public ClusterLayoutComputeServers dynamicCount(Boolean dynamicCount) {
    
    this.dynamicCount = dynamicCount;
    return this;
  }

   /**
   * Get dynamicCount
   * @return dynamicCount
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public Boolean getDynamicCount() {
    return dynamicCount;
  }


  public void setDynamicCount(Boolean dynamicCount) {
    this.dynamicCount = dynamicCount;
  }


  public ClusterLayoutComputeServers installContainerRuntime(Boolean installContainerRuntime) {
    
    this.installContainerRuntime = installContainerRuntime;
    return this;
  }

   /**
   * Get installContainerRuntime
   * @return installContainerRuntime
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public Boolean getInstallContainerRuntime() {
    return installContainerRuntime;
  }


  public void setInstallContainerRuntime(Boolean installContainerRuntime) {
    this.installContainerRuntime = installContainerRuntime;
  }


  public ClusterLayoutComputeServers installStorageRuntime(Boolean installStorageRuntime) {
    
    this.installStorageRuntime = installStorageRuntime;
    return this;
  }

   /**
   * Get installStorageRuntime
   * @return installStorageRuntime
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public Boolean getInstallStorageRuntime() {
    return installStorageRuntime;
  }


  public void setInstallStorageRuntime(Boolean installStorageRuntime) {
    this.installStorageRuntime = installStorageRuntime;
  }


  public ClusterLayoutComputeServers name(String name) {
    
    this.name = name;
    return this;
  }

   /**
   * Get name
   * @return name
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public String getName() {
    return name;
  }


  public void setName(String name) {
    this.name = name;
  }


  public ClusterLayoutComputeServers code(String code) {
    
    this.code = code;
    return this;
  }

   /**
   * Get code
   * @return code
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public String getCode() {
    return code;
  }


  public void setCode(String code) {
    this.code = code;
  }


  public ClusterLayoutComputeServers category(String category) {
    
    this.category = category;
    return this;
  }

   /**
   * Get category
   * @return category
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public String getCategory() {
    return category;
  }


  public void setCategory(String category) {
    this.category = category;
  }


  public ClusterLayoutComputeServers config(String config) {
    
    this.config = config;
    return this;
  }

   /**
   * Get config
   * @return config
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public String getConfig() {
    return config;
  }


  public void setConfig(String config) {
    this.config = config;
  }


  public ClusterLayoutComputeServers containerType(ClusterLayoutContainerType containerType) {
    
    this.containerType = containerType;
    return this;
  }

   /**
   * Get containerType
   * @return containerType
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public ClusterLayoutContainerType getContainerType() {
    return containerType;
  }


  public void setContainerType(ClusterLayoutContainerType containerType) {
    this.containerType = containerType;
  }


  public ClusterLayoutComputeServers computeServerType(ClusterLayoutComputeServerType computeServerType) {
    
    this.computeServerType = computeServerType;
    return this;
  }

   /**
   * Get computeServerType
   * @return computeServerType
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public ClusterLayoutComputeServerType getComputeServerType() {
    return computeServerType;
  }


  public void setComputeServerType(ClusterLayoutComputeServerType computeServerType) {
    this.computeServerType = computeServerType;
  }


  public ClusterLayoutComputeServers provisionService(String provisionService) {
    
    this.provisionService = provisionService;
    return this;
  }

   /**
   * Get provisionService
   * @return provisionService
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public String getProvisionService() {
    return provisionService;
  }


  public void setProvisionService(String provisionService) {
    this.provisionService = provisionService;
  }


  public ClusterLayoutComputeServers planCategory(String planCategory) {
    
    this.planCategory = planCategory;
    return this;
  }

   /**
   * Get planCategory
   * @return planCategory
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public String getPlanCategory() {
    return planCategory;
  }


  public void setPlanCategory(String planCategory) {
    this.planCategory = planCategory;
  }


  public ClusterLayoutComputeServers namePrefix(String namePrefix) {
    
    this.namePrefix = namePrefix;
    return this;
  }

   /**
   * Get namePrefix
   * @return namePrefix
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public String getNamePrefix() {
    return namePrefix;
  }


  public void setNamePrefix(String namePrefix) {
    this.namePrefix = namePrefix;
  }


  public ClusterLayoutComputeServers nameSuffix(String nameSuffix) {
    
    this.nameSuffix = nameSuffix;
    return this;
  }

   /**
   * Get nameSuffix
   * @return nameSuffix
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public String getNameSuffix() {
    return nameSuffix;
  }


  public void setNameSuffix(String nameSuffix) {
    this.nameSuffix = nameSuffix;
  }


  public ClusterLayoutComputeServers forceNameIndex(Boolean forceNameIndex) {
    
    this.forceNameIndex = forceNameIndex;
    return this;
  }

   /**
   * Get forceNameIndex
   * @return forceNameIndex
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public Boolean getForceNameIndex() {
    return forceNameIndex;
  }


  public void setForceNameIndex(Boolean forceNameIndex) {
    this.forceNameIndex = forceNameIndex;
  }


  public ClusterLayoutComputeServers loadBalance(Boolean loadBalance) {
    
    this.loadBalance = loadBalance;
    return this;
  }

   /**
   * Get loadBalance
   * @return loadBalance
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public Boolean getLoadBalance() {
    return loadBalance;
  }


  public void setLoadBalance(Boolean loadBalance) {
    this.loadBalance = loadBalance;
  }


  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (o == null || getClass() != o.getClass()) {
      return false;
    }
    ClusterLayoutComputeServers clusterLayoutComputeServers = (ClusterLayoutComputeServers) o;
    return Objects.equals(this.id, clusterLayoutComputeServers.id) &&
        Objects.equals(this.priorityOrder, clusterLayoutComputeServers.priorityOrder) &&
        Objects.equals(this.nodeCount, clusterLayoutComputeServers.nodeCount) &&
        Objects.equals(this.nodeType, clusterLayoutComputeServers.nodeType) &&
        Objects.equals(this.minNodeCount, clusterLayoutComputeServers.minNodeCount) &&
        Objects.equals(this.maxNodeCount, clusterLayoutComputeServers.maxNodeCount) &&
        Objects.equals(this.dynamicCount, clusterLayoutComputeServers.dynamicCount) &&
        Objects.equals(this.installContainerRuntime, clusterLayoutComputeServers.installContainerRuntime) &&
        Objects.equals(this.installStorageRuntime, clusterLayoutComputeServers.installStorageRuntime) &&
        Objects.equals(this.name, clusterLayoutComputeServers.name) &&
        Objects.equals(this.code, clusterLayoutComputeServers.code) &&
        Objects.equals(this.category, clusterLayoutComputeServers.category) &&
        Objects.equals(this.config, clusterLayoutComputeServers.config) &&
        Objects.equals(this.containerType, clusterLayoutComputeServers.containerType) &&
        Objects.equals(this.computeServerType, clusterLayoutComputeServers.computeServerType) &&
        Objects.equals(this.provisionService, clusterLayoutComputeServers.provisionService) &&
        Objects.equals(this.planCategory, clusterLayoutComputeServers.planCategory) &&
        Objects.equals(this.namePrefix, clusterLayoutComputeServers.namePrefix) &&
        Objects.equals(this.nameSuffix, clusterLayoutComputeServers.nameSuffix) &&
        Objects.equals(this.forceNameIndex, clusterLayoutComputeServers.forceNameIndex) &&
        Objects.equals(this.loadBalance, clusterLayoutComputeServers.loadBalance);
  }

  @Override
  public int hashCode() {
    return Objects.hash(id, priorityOrder, nodeCount, nodeType, minNodeCount, maxNodeCount, dynamicCount, installContainerRuntime, installStorageRuntime, name, code, category, config, containerType, computeServerType, provisionService, planCategory, namePrefix, nameSuffix, forceNameIndex, loadBalance);
  }


  @Override
  public String toString() {
    StringBuilder sb = new StringBuilder();
    sb.append("class ClusterLayoutComputeServers {\n");
    sb.append("    id: ").append(toIndentedString(id)).append("\n");
    sb.append("    priorityOrder: ").append(toIndentedString(priorityOrder)).append("\n");
    sb.append("    nodeCount: ").append(toIndentedString(nodeCount)).append("\n");
    sb.append("    nodeType: ").append(toIndentedString(nodeType)).append("\n");
    sb.append("    minNodeCount: ").append(toIndentedString(minNodeCount)).append("\n");
    sb.append("    maxNodeCount: ").append(toIndentedString(maxNodeCount)).append("\n");
    sb.append("    dynamicCount: ").append(toIndentedString(dynamicCount)).append("\n");
    sb.append("    installContainerRuntime: ").append(toIndentedString(installContainerRuntime)).append("\n");
    sb.append("    installStorageRuntime: ").append(toIndentedString(installStorageRuntime)).append("\n");
    sb.append("    name: ").append(toIndentedString(name)).append("\n");
    sb.append("    code: ").append(toIndentedString(code)).append("\n");
    sb.append("    category: ").append(toIndentedString(category)).append("\n");
    sb.append("    config: ").append(toIndentedString(config)).append("\n");
    sb.append("    containerType: ").append(toIndentedString(containerType)).append("\n");
    sb.append("    computeServerType: ").append(toIndentedString(computeServerType)).append("\n");
    sb.append("    provisionService: ").append(toIndentedString(provisionService)).append("\n");
    sb.append("    planCategory: ").append(toIndentedString(planCategory)).append("\n");
    sb.append("    namePrefix: ").append(toIndentedString(namePrefix)).append("\n");
    sb.append("    nameSuffix: ").append(toIndentedString(nameSuffix)).append("\n");
    sb.append("    forceNameIndex: ").append(toIndentedString(forceNameIndex)).append("\n");
    sb.append("    loadBalance: ").append(toIndentedString(loadBalance)).append("\n");
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
