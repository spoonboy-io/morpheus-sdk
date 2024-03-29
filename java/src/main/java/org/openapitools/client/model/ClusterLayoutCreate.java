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
import java.util.ArrayList;
import java.util.List;
import org.openapitools.client.model.ApiBlueprintsIdUpdatePermissionsResourcePermissionSites;
import org.openapitools.client.model.ApiServicePlansServicePlanProvisionType;
import org.openapitools.client.model.ClusterLayoutCreateEnvironmentVariables;
import org.openapitools.client.model.ClusterLayoutCreateGroupType;
import org.openapitools.client.model.ClusterLayoutCreateMasters;

/**
 * ClusterLayoutCreate
 */
@javax.annotation.Generated(value = "org.openapitools.codegen.languages.JavaClientCodegen", date = "2023-08-21T11:26:47.028Z[GMT]")
public class ClusterLayoutCreate {
  public static final String SERIALIZED_NAME_NAME = "name";
  @SerializedName(SERIALIZED_NAME_NAME)
  private String name;

  public static final String SERIALIZED_NAME_DESCRIPTION = "description";
  @SerializedName(SERIALIZED_NAME_DESCRIPTION)
  private String description;

  public static final String SERIALIZED_NAME_LABELS = "labels";
  @SerializedName(SERIALIZED_NAME_LABELS)
  private List<String> labels = null;

  public static final String SERIALIZED_NAME_COMPUTE_VERSION = "computeVersion";
  @SerializedName(SERIALIZED_NAME_COMPUTE_VERSION)
  private String computeVersion;

  public static final String SERIALIZED_NAME_CREATABLE = "creatable";
  @SerializedName(SERIALIZED_NAME_CREATABLE)
  private Boolean creatable = true;

  public static final String SERIALIZED_NAME_HAS_AUTO_SCALE = "hasAutoScale";
  @SerializedName(SERIALIZED_NAME_HAS_AUTO_SCALE)
  private Boolean hasAutoScale = false;

  public static final String SERIALIZED_NAME_INSTALL_CONTAINER_RUNTIME = "installContainerRuntime";
  @SerializedName(SERIALIZED_NAME_INSTALL_CONTAINER_RUNTIME)
  private Boolean installContainerRuntime = false;

  public static final String SERIALIZED_NAME_MEMORY_REQUIREMENT = "memoryRequirement";
  @SerializedName(SERIALIZED_NAME_MEMORY_REQUIREMENT)
  private Long memoryRequirement;

  public static final String SERIALIZED_NAME_GROUP_TYPE = "groupType";
  @SerializedName(SERIALIZED_NAME_GROUP_TYPE)
  private ClusterLayoutCreateGroupType groupType;

  public static final String SERIALIZED_NAME_PROVISION_TYPE = "provisionType";
  @SerializedName(SERIALIZED_NAME_PROVISION_TYPE)
  private ApiServicePlansServicePlanProvisionType provisionType;

  public static final String SERIALIZED_NAME_OPTION_TYPES = "optionTypes";
  @SerializedName(SERIALIZED_NAME_OPTION_TYPES)
  private List<ApiBlueprintsIdUpdatePermissionsResourcePermissionSites> optionTypes = null;

  public static final String SERIALIZED_NAME_TASK_SETS = "taskSets";
  @SerializedName(SERIALIZED_NAME_TASK_SETS)
  private List<ApiBlueprintsIdUpdatePermissionsResourcePermissionSites> taskSets = null;

  public static final String SERIALIZED_NAME_ENVIRONMENT_VARIABLES = "environmentVariables";
  @SerializedName(SERIALIZED_NAME_ENVIRONMENT_VARIABLES)
  private List<ClusterLayoutCreateEnvironmentVariables> environmentVariables = null;

  public static final String SERIALIZED_NAME_MASTERS = "masters";
  @SerializedName(SERIALIZED_NAME_MASTERS)
  private List<ClusterLayoutCreateMasters> masters = null;

  public static final String SERIALIZED_NAME_WORKERS = "workers";
  @SerializedName(SERIALIZED_NAME_WORKERS)
  private List<ClusterLayoutCreateMasters> workers = null;


  public ClusterLayoutCreate name(String name) {
    
    this.name = name;
    return this;
  }

   /**
   * Cluster layout name
   * @return name
  **/
  @ApiModelProperty(required = true, value = "Cluster layout name")

  public String getName() {
    return name;
  }


  public void setName(String name) {
    this.name = name;
  }


  public ClusterLayoutCreate description(String description) {
    
    this.description = description;
    return this;
  }

   /**
   * Cluster layout description
   * @return description
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "Cluster layout description")

  public String getDescription() {
    return description;
  }


  public void setDescription(String description) {
    this.description = description;
  }


  public ClusterLayoutCreate labels(List<String> labels) {
    
    this.labels = labels;
    return this;
  }

  public ClusterLayoutCreate addLabelsItem(String labelsItem) {
    if (this.labels == null) {
      this.labels = new ArrayList<String>();
    }
    this.labels.add(labelsItem);
    return this;
  }

   /**
   * Get labels
   * @return labels
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public List<String> getLabels() {
    return labels;
  }


  public void setLabels(List<String> labels) {
    this.labels = labels;
  }


  public ClusterLayoutCreate computeVersion(String computeVersion) {
    
    this.computeVersion = computeVersion;
    return this;
  }

   /**
   * Version of the cluster layout
   * @return computeVersion
  **/
  @ApiModelProperty(required = true, value = "Version of the cluster layout")

  public String getComputeVersion() {
    return computeVersion;
  }


  public void setComputeVersion(String computeVersion) {
    this.computeVersion = computeVersion;
  }


  public ClusterLayoutCreate creatable(Boolean creatable) {
    
    this.creatable = creatable;
    return this;
  }

   /**
   * Can be used to enable / disable the creatability of the cluster layout.
   * @return creatable
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "Can be used to enable / disable the creatability of the cluster layout.")

  public Boolean getCreatable() {
    return creatable;
  }


  public void setCreatable(Boolean creatable) {
    this.creatable = creatable;
  }


  public ClusterLayoutCreate hasAutoScale(Boolean hasAutoScale) {
    
    this.hasAutoScale = hasAutoScale;
    return this;
  }

   /**
   * Can be used to enable / disable the horizontal scaling.
   * @return hasAutoScale
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "Can be used to enable / disable the horizontal scaling.")

  public Boolean getHasAutoScale() {
    return hasAutoScale;
  }


  public void setHasAutoScale(Boolean hasAutoScale) {
    this.hasAutoScale = hasAutoScale;
  }


  public ClusterLayoutCreate installContainerRuntime(Boolean installContainerRuntime) {
    
    this.installContainerRuntime = installContainerRuntime;
    return this;
  }

   /**
   * Install Docker (container runtime).
   * @return installContainerRuntime
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "Install Docker (container runtime).")

  public Boolean getInstallContainerRuntime() {
    return installContainerRuntime;
  }


  public void setInstallContainerRuntime(Boolean installContainerRuntime) {
    this.installContainerRuntime = installContainerRuntime;
  }


  public ClusterLayoutCreate memoryRequirement(Long memoryRequirement) {
    
    this.memoryRequirement = memoryRequirement;
    return this;
  }

   /**
   * Memory requirement in bytes
   * @return memoryRequirement
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "Memory requirement in bytes")

  public Long getMemoryRequirement() {
    return memoryRequirement;
  }


  public void setMemoryRequirement(Long memoryRequirement) {
    this.memoryRequirement = memoryRequirement;
  }


  public ClusterLayoutCreate groupType(ClusterLayoutCreateGroupType groupType) {
    
    this.groupType = groupType;
    return this;
  }

   /**
   * Get groupType
   * @return groupType
  **/
  @ApiModelProperty(required = true, value = "")

  public ClusterLayoutCreateGroupType getGroupType() {
    return groupType;
  }


  public void setGroupType(ClusterLayoutCreateGroupType groupType) {
    this.groupType = groupType;
  }


  public ClusterLayoutCreate provisionType(ApiServicePlansServicePlanProvisionType provisionType) {
    
    this.provisionType = provisionType;
    return this;
  }

   /**
   * Get provisionType
   * @return provisionType
  **/
  @ApiModelProperty(required = true, value = "")

  public ApiServicePlansServicePlanProvisionType getProvisionType() {
    return provisionType;
  }


  public void setProvisionType(ApiServicePlansServicePlanProvisionType provisionType) {
    this.provisionType = provisionType;
  }


  public ClusterLayoutCreate optionTypes(List<ApiBlueprintsIdUpdatePermissionsResourcePermissionSites> optionTypes) {
    
    this.optionTypes = optionTypes;
    return this;
  }

  public ClusterLayoutCreate addOptionTypesItem(ApiBlueprintsIdUpdatePermissionsResourcePermissionSites optionTypesItem) {
    if (this.optionTypes == null) {
      this.optionTypes = new ArrayList<ApiBlueprintsIdUpdatePermissionsResourcePermissionSites>();
    }
    this.optionTypes.add(optionTypesItem);
    return this;
  }

   /**
   * Array of cluster layout option types
   * @return optionTypes
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "Array of cluster layout option types")

  public List<ApiBlueprintsIdUpdatePermissionsResourcePermissionSites> getOptionTypes() {
    return optionTypes;
  }


  public void setOptionTypes(List<ApiBlueprintsIdUpdatePermissionsResourcePermissionSites> optionTypes) {
    this.optionTypes = optionTypes;
  }


  public ClusterLayoutCreate taskSets(List<ApiBlueprintsIdUpdatePermissionsResourcePermissionSites> taskSets) {
    
    this.taskSets = taskSets;
    return this;
  }

  public ClusterLayoutCreate addTaskSetsItem(ApiBlueprintsIdUpdatePermissionsResourcePermissionSites taskSetsItem) {
    if (this.taskSets == null) {
      this.taskSets = new ArrayList<ApiBlueprintsIdUpdatePermissionsResourcePermissionSites>();
    }
    this.taskSets.add(taskSetsItem);
    return this;
  }

   /**
   * Array of cluster layout task sets
   * @return taskSets
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "Array of cluster layout task sets")

  public List<ApiBlueprintsIdUpdatePermissionsResourcePermissionSites> getTaskSets() {
    return taskSets;
  }


  public void setTaskSets(List<ApiBlueprintsIdUpdatePermissionsResourcePermissionSites> taskSets) {
    this.taskSets = taskSets;
  }


  public ClusterLayoutCreate environmentVariables(List<ClusterLayoutCreateEnvironmentVariables> environmentVariables) {
    
    this.environmentVariables = environmentVariables;
    return this;
  }

  public ClusterLayoutCreate addEnvironmentVariablesItem(ClusterLayoutCreateEnvironmentVariables environmentVariablesItem) {
    if (this.environmentVariables == null) {
      this.environmentVariables = new ArrayList<ClusterLayoutCreateEnvironmentVariables>();
    }
    this.environmentVariables.add(environmentVariablesItem);
    return this;
  }

   /**
   * Array of cluster layout env variables
   * @return environmentVariables
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "Array of cluster layout env variables")

  public List<ClusterLayoutCreateEnvironmentVariables> getEnvironmentVariables() {
    return environmentVariables;
  }


  public void setEnvironmentVariables(List<ClusterLayoutCreateEnvironmentVariables> environmentVariables) {
    this.environmentVariables = environmentVariables;
  }


  public ClusterLayoutCreate masters(List<ClusterLayoutCreateMasters> masters) {
    
    this.masters = masters;
    return this;
  }

  public ClusterLayoutCreate addMastersItem(ClusterLayoutCreateMasters mastersItem) {
    if (this.masters == null) {
      this.masters = new ArrayList<ClusterLayoutCreateMasters>();
    }
    this.masters.add(mastersItem);
    return this;
  }

   /**
   * Array of cluster layout master nodes
   * @return masters
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "Array of cluster layout master nodes")

  public List<ClusterLayoutCreateMasters> getMasters() {
    return masters;
  }


  public void setMasters(List<ClusterLayoutCreateMasters> masters) {
    this.masters = masters;
  }


  public ClusterLayoutCreate workers(List<ClusterLayoutCreateMasters> workers) {
    
    this.workers = workers;
    return this;
  }

  public ClusterLayoutCreate addWorkersItem(ClusterLayoutCreateMasters workersItem) {
    if (this.workers == null) {
      this.workers = new ArrayList<ClusterLayoutCreateMasters>();
    }
    this.workers.add(workersItem);
    return this;
  }

   /**
   * Array of cluster layout worker nodes
   * @return workers
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "Array of cluster layout worker nodes")

  public List<ClusterLayoutCreateMasters> getWorkers() {
    return workers;
  }


  public void setWorkers(List<ClusterLayoutCreateMasters> workers) {
    this.workers = workers;
  }


  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (o == null || getClass() != o.getClass()) {
      return false;
    }
    ClusterLayoutCreate clusterLayoutCreate = (ClusterLayoutCreate) o;
    return Objects.equals(this.name, clusterLayoutCreate.name) &&
        Objects.equals(this.description, clusterLayoutCreate.description) &&
        Objects.equals(this.labels, clusterLayoutCreate.labels) &&
        Objects.equals(this.computeVersion, clusterLayoutCreate.computeVersion) &&
        Objects.equals(this.creatable, clusterLayoutCreate.creatable) &&
        Objects.equals(this.hasAutoScale, clusterLayoutCreate.hasAutoScale) &&
        Objects.equals(this.installContainerRuntime, clusterLayoutCreate.installContainerRuntime) &&
        Objects.equals(this.memoryRequirement, clusterLayoutCreate.memoryRequirement) &&
        Objects.equals(this.groupType, clusterLayoutCreate.groupType) &&
        Objects.equals(this.provisionType, clusterLayoutCreate.provisionType) &&
        Objects.equals(this.optionTypes, clusterLayoutCreate.optionTypes) &&
        Objects.equals(this.taskSets, clusterLayoutCreate.taskSets) &&
        Objects.equals(this.environmentVariables, clusterLayoutCreate.environmentVariables) &&
        Objects.equals(this.masters, clusterLayoutCreate.masters) &&
        Objects.equals(this.workers, clusterLayoutCreate.workers);
  }

  @Override
  public int hashCode() {
    return Objects.hash(name, description, labels, computeVersion, creatable, hasAutoScale, installContainerRuntime, memoryRequirement, groupType, provisionType, optionTypes, taskSets, environmentVariables, masters, workers);
  }


  @Override
  public String toString() {
    StringBuilder sb = new StringBuilder();
    sb.append("class ClusterLayoutCreate {\n");
    sb.append("    name: ").append(toIndentedString(name)).append("\n");
    sb.append("    description: ").append(toIndentedString(description)).append("\n");
    sb.append("    labels: ").append(toIndentedString(labels)).append("\n");
    sb.append("    computeVersion: ").append(toIndentedString(computeVersion)).append("\n");
    sb.append("    creatable: ").append(toIndentedString(creatable)).append("\n");
    sb.append("    hasAutoScale: ").append(toIndentedString(hasAutoScale)).append("\n");
    sb.append("    installContainerRuntime: ").append(toIndentedString(installContainerRuntime)).append("\n");
    sb.append("    memoryRequirement: ").append(toIndentedString(memoryRequirement)).append("\n");
    sb.append("    groupType: ").append(toIndentedString(groupType)).append("\n");
    sb.append("    provisionType: ").append(toIndentedString(provisionType)).append("\n");
    sb.append("    optionTypes: ").append(toIndentedString(optionTypes)).append("\n");
    sb.append("    taskSets: ").append(toIndentedString(taskSets)).append("\n");
    sb.append("    environmentVariables: ").append(toIndentedString(environmentVariables)).append("\n");
    sb.append("    masters: ").append(toIndentedString(masters)).append("\n");
    sb.append("    workers: ").append(toIndentedString(workers)).append("\n");
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

