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
import org.openapitools.client.model.ClusterLayoutComputeServers;
import org.openapitools.client.model.ClusterLayoutSpecTemplates;
import org.openapitools.client.model.InlineResponse20079LoadBalancerMonitorLoadBalancerType;
import org.openapitools.client.model.InlineResponse20094Network;
import org.openapitools.client.model.OptionType;
import org.threeten.bp.OffsetDateTime;

/**
 * ClusterLayout
 */
@javax.annotation.Generated(value = "org.openapitools.codegen.languages.JavaClientCodegen", date = "2023-08-21T11:26:47.028Z[GMT]")
public class ClusterLayout {
  public static final String SERIALIZED_NAME_ID = "id";
  @SerializedName(SERIALIZED_NAME_ID)
  private Long id;

  public static final String SERIALIZED_NAME_INTERNAL_ID = "internalId";
  @SerializedName(SERIALIZED_NAME_INTERNAL_ID)
  private String internalId;

  public static final String SERIALIZED_NAME_SERVER_COUNT = "serverCount";
  @SerializedName(SERIALIZED_NAME_SERVER_COUNT)
  private Long serverCount;

  public static final String SERIALIZED_NAME_DATE_CREATED = "dateCreated";
  @SerializedName(SERIALIZED_NAME_DATE_CREATED)
  private OffsetDateTime dateCreated;

  public static final String SERIALIZED_NAME_CODE = "code";
  @SerializedName(SERIALIZED_NAME_CODE)
  private String code;

  public static final String SERIALIZED_NAME_LAST_UPDATED = "lastUpdated";
  @SerializedName(SERIALIZED_NAME_LAST_UPDATED)
  private OffsetDateTime lastUpdated;

  public static final String SERIALIZED_NAME_HAS_AUTO_SCALE = "hasAutoScale";
  @SerializedName(SERIALIZED_NAME_HAS_AUTO_SCALE)
  private Boolean hasAutoScale;

  public static final String SERIALIZED_NAME_MEMORY_REQUIREMENT = "memoryRequirement";
  @SerializedName(SERIALIZED_NAME_MEMORY_REQUIREMENT)
  private Long memoryRequirement;

  public static final String SERIALIZED_NAME_CLUSTER_VERSION = "clusterVersion";
  @SerializedName(SERIALIZED_NAME_CLUSTER_VERSION)
  private String clusterVersion;

  public static final String SERIALIZED_NAME_COMPUTE_VERSION = "computeVersion";
  @SerializedName(SERIALIZED_NAME_COMPUTE_VERSION)
  private String computeVersion;

  public static final String SERIALIZED_NAME_HAS_SETTINGS = "hasSettings";
  @SerializedName(SERIALIZED_NAME_HAS_SETTINGS)
  private Boolean hasSettings;

  public static final String SERIALIZED_NAME_SORT_ORDER = "sortOrder";
  @SerializedName(SERIALIZED_NAME_SORT_ORDER)
  private Long sortOrder;

  public static final String SERIALIZED_NAME_HAS_CONFIG = "hasConfig";
  @SerializedName(SERIALIZED_NAME_HAS_CONFIG)
  private Boolean hasConfig;

  public static final String SERIALIZED_NAME_NAME = "name";
  @SerializedName(SERIALIZED_NAME_NAME)
  private String name;

  public static final String SERIALIZED_NAME_CREATABLE = "creatable";
  @SerializedName(SERIALIZED_NAME_CREATABLE)
  private Boolean creatable;

  public static final String SERIALIZED_NAME_ENABLED = "enabled";
  @SerializedName(SERIALIZED_NAME_ENABLED)
  private Boolean enabled;

  public static final String SERIALIZED_NAME_DESCRIPTION = "description";
  @SerializedName(SERIALIZED_NAME_DESCRIPTION)
  private String description;

  public static final String SERIALIZED_NAME_GROUP_TYPE = "groupType";
  @SerializedName(SERIALIZED_NAME_GROUP_TYPE)
  private InlineResponse20079LoadBalancerMonitorLoadBalancerType groupType;

  public static final String SERIALIZED_NAME_LABELS = "labels";
  @SerializedName(SERIALIZED_NAME_LABELS)
  private List<String> labels = null;

  public static final String SERIALIZED_NAME_ENVIRONMENT_VARIABLES = "environmentVariables";
  @SerializedName(SERIALIZED_NAME_ENVIRONMENT_VARIABLES)
  private List<Object> environmentVariables = null;

  public static final String SERIALIZED_NAME_OPTION_TYPES = "optionTypes";
  @SerializedName(SERIALIZED_NAME_OPTION_TYPES)
  private List<OptionType> optionTypes = null;

  public static final String SERIALIZED_NAME_ACTIONS = "actions";
  @SerializedName(SERIALIZED_NAME_ACTIONS)
  private List<Object> actions = null;

  public static final String SERIALIZED_NAME_COMPUTE_SERVERS = "computeServers";
  @SerializedName(SERIALIZED_NAME_COMPUTE_SERVERS)
  private List<ClusterLayoutComputeServers> computeServers = null;

  public static final String SERIALIZED_NAME_INSTALL_CONTAINER_RUNTIME = "installContainerRuntime";
  @SerializedName(SERIALIZED_NAME_INSTALL_CONTAINER_RUNTIME)
  private Boolean installContainerRuntime;

  public static final String SERIALIZED_NAME_PROVISION_TYPE = "provisionType";
  @SerializedName(SERIALIZED_NAME_PROVISION_TYPE)
  private InlineResponse20094Network provisionType;

  public static final String SERIALIZED_NAME_SPEC_TEMPLATES = "specTemplates";
  @SerializedName(SERIALIZED_NAME_SPEC_TEMPLATES)
  private List<ClusterLayoutSpecTemplates> specTemplates = null;

  public static final String SERIALIZED_NAME_TASK_SETS = "taskSets";
  @SerializedName(SERIALIZED_NAME_TASK_SETS)
  private List<Object> taskSets = null;

  public static final String SERIALIZED_NAME_TYPE = "type";
  @SerializedName(SERIALIZED_NAME_TYPE)
  private InlineResponse20079LoadBalancerMonitorLoadBalancerType type;


  public ClusterLayout id(Long id) {
    
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


  public ClusterLayout internalId(String internalId) {
    
    this.internalId = internalId;
    return this;
  }

   /**
   * Get internalId
   * @return internalId
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public String getInternalId() {
    return internalId;
  }


  public void setInternalId(String internalId) {
    this.internalId = internalId;
  }


  public ClusterLayout serverCount(Long serverCount) {
    
    this.serverCount = serverCount;
    return this;
  }

   /**
   * Get serverCount
   * @return serverCount
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public Long getServerCount() {
    return serverCount;
  }


  public void setServerCount(Long serverCount) {
    this.serverCount = serverCount;
  }


  public ClusterLayout dateCreated(OffsetDateTime dateCreated) {
    
    this.dateCreated = dateCreated;
    return this;
  }

   /**
   * Get dateCreated
   * @return dateCreated
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public OffsetDateTime getDateCreated() {
    return dateCreated;
  }


  public void setDateCreated(OffsetDateTime dateCreated) {
    this.dateCreated = dateCreated;
  }


  public ClusterLayout code(String code) {
    
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


  public ClusterLayout lastUpdated(OffsetDateTime lastUpdated) {
    
    this.lastUpdated = lastUpdated;
    return this;
  }

   /**
   * Get lastUpdated
   * @return lastUpdated
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public OffsetDateTime getLastUpdated() {
    return lastUpdated;
  }


  public void setLastUpdated(OffsetDateTime lastUpdated) {
    this.lastUpdated = lastUpdated;
  }


  public ClusterLayout hasAutoScale(Boolean hasAutoScale) {
    
    this.hasAutoScale = hasAutoScale;
    return this;
  }

   /**
   * Get hasAutoScale
   * @return hasAutoScale
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public Boolean getHasAutoScale() {
    return hasAutoScale;
  }


  public void setHasAutoScale(Boolean hasAutoScale) {
    this.hasAutoScale = hasAutoScale;
  }


  public ClusterLayout memoryRequirement(Long memoryRequirement) {
    
    this.memoryRequirement = memoryRequirement;
    return this;
  }

   /**
   * Get memoryRequirement
   * @return memoryRequirement
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public Long getMemoryRequirement() {
    return memoryRequirement;
  }


  public void setMemoryRequirement(Long memoryRequirement) {
    this.memoryRequirement = memoryRequirement;
  }


  public ClusterLayout clusterVersion(String clusterVersion) {
    
    this.clusterVersion = clusterVersion;
    return this;
  }

   /**
   * Get clusterVersion
   * @return clusterVersion
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public String getClusterVersion() {
    return clusterVersion;
  }


  public void setClusterVersion(String clusterVersion) {
    this.clusterVersion = clusterVersion;
  }


  public ClusterLayout computeVersion(String computeVersion) {
    
    this.computeVersion = computeVersion;
    return this;
  }

   /**
   * Get computeVersion
   * @return computeVersion
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public String getComputeVersion() {
    return computeVersion;
  }


  public void setComputeVersion(String computeVersion) {
    this.computeVersion = computeVersion;
  }


  public ClusterLayout hasSettings(Boolean hasSettings) {
    
    this.hasSettings = hasSettings;
    return this;
  }

   /**
   * Get hasSettings
   * @return hasSettings
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public Boolean getHasSettings() {
    return hasSettings;
  }


  public void setHasSettings(Boolean hasSettings) {
    this.hasSettings = hasSettings;
  }


  public ClusterLayout sortOrder(Long sortOrder) {
    
    this.sortOrder = sortOrder;
    return this;
  }

   /**
   * Get sortOrder
   * @return sortOrder
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public Long getSortOrder() {
    return sortOrder;
  }


  public void setSortOrder(Long sortOrder) {
    this.sortOrder = sortOrder;
  }


  public ClusterLayout hasConfig(Boolean hasConfig) {
    
    this.hasConfig = hasConfig;
    return this;
  }

   /**
   * Get hasConfig
   * @return hasConfig
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public Boolean getHasConfig() {
    return hasConfig;
  }


  public void setHasConfig(Boolean hasConfig) {
    this.hasConfig = hasConfig;
  }


  public ClusterLayout name(String name) {
    
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


  public ClusterLayout creatable(Boolean creatable) {
    
    this.creatable = creatable;
    return this;
  }

   /**
   * Get creatable
   * @return creatable
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public Boolean getCreatable() {
    return creatable;
  }


  public void setCreatable(Boolean creatable) {
    this.creatable = creatable;
  }


  public ClusterLayout enabled(Boolean enabled) {
    
    this.enabled = enabled;
    return this;
  }

   /**
   * Get enabled
   * @return enabled
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public Boolean getEnabled() {
    return enabled;
  }


  public void setEnabled(Boolean enabled) {
    this.enabled = enabled;
  }


  public ClusterLayout description(String description) {
    
    this.description = description;
    return this;
  }

   /**
   * Get description
   * @return description
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public String getDescription() {
    return description;
  }


  public void setDescription(String description) {
    this.description = description;
  }


  public ClusterLayout groupType(InlineResponse20079LoadBalancerMonitorLoadBalancerType groupType) {
    
    this.groupType = groupType;
    return this;
  }

   /**
   * Get groupType
   * @return groupType
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public InlineResponse20079LoadBalancerMonitorLoadBalancerType getGroupType() {
    return groupType;
  }


  public void setGroupType(InlineResponse20079LoadBalancerMonitorLoadBalancerType groupType) {
    this.groupType = groupType;
  }


  public ClusterLayout labels(List<String> labels) {
    
    this.labels = labels;
    return this;
  }

  public ClusterLayout addLabelsItem(String labelsItem) {
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


  public ClusterLayout environmentVariables(List<Object> environmentVariables) {
    
    this.environmentVariables = environmentVariables;
    return this;
  }

  public ClusterLayout addEnvironmentVariablesItem(Object environmentVariablesItem) {
    if (this.environmentVariables == null) {
      this.environmentVariables = new ArrayList<Object>();
    }
    this.environmentVariables.add(environmentVariablesItem);
    return this;
  }

   /**
   * Get environmentVariables
   * @return environmentVariables
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public List<Object> getEnvironmentVariables() {
    return environmentVariables;
  }


  public void setEnvironmentVariables(List<Object> environmentVariables) {
    this.environmentVariables = environmentVariables;
  }


  public ClusterLayout optionTypes(List<OptionType> optionTypes) {
    
    this.optionTypes = optionTypes;
    return this;
  }

  public ClusterLayout addOptionTypesItem(OptionType optionTypesItem) {
    if (this.optionTypes == null) {
      this.optionTypes = new ArrayList<OptionType>();
    }
    this.optionTypes.add(optionTypesItem);
    return this;
  }

   /**
   * Get optionTypes
   * @return optionTypes
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public List<OptionType> getOptionTypes() {
    return optionTypes;
  }


  public void setOptionTypes(List<OptionType> optionTypes) {
    this.optionTypes = optionTypes;
  }


  public ClusterLayout actions(List<Object> actions) {
    
    this.actions = actions;
    return this;
  }

  public ClusterLayout addActionsItem(Object actionsItem) {
    if (this.actions == null) {
      this.actions = new ArrayList<Object>();
    }
    this.actions.add(actionsItem);
    return this;
  }

   /**
   * Get actions
   * @return actions
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public List<Object> getActions() {
    return actions;
  }


  public void setActions(List<Object> actions) {
    this.actions = actions;
  }


  public ClusterLayout computeServers(List<ClusterLayoutComputeServers> computeServers) {
    
    this.computeServers = computeServers;
    return this;
  }

  public ClusterLayout addComputeServersItem(ClusterLayoutComputeServers computeServersItem) {
    if (this.computeServers == null) {
      this.computeServers = new ArrayList<ClusterLayoutComputeServers>();
    }
    this.computeServers.add(computeServersItem);
    return this;
  }

   /**
   * Get computeServers
   * @return computeServers
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public List<ClusterLayoutComputeServers> getComputeServers() {
    return computeServers;
  }


  public void setComputeServers(List<ClusterLayoutComputeServers> computeServers) {
    this.computeServers = computeServers;
  }


  public ClusterLayout installContainerRuntime(Boolean installContainerRuntime) {
    
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


  public ClusterLayout provisionType(InlineResponse20094Network provisionType) {
    
    this.provisionType = provisionType;
    return this;
  }

   /**
   * Get provisionType
   * @return provisionType
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public InlineResponse20094Network getProvisionType() {
    return provisionType;
  }


  public void setProvisionType(InlineResponse20094Network provisionType) {
    this.provisionType = provisionType;
  }


  public ClusterLayout specTemplates(List<ClusterLayoutSpecTemplates> specTemplates) {
    
    this.specTemplates = specTemplates;
    return this;
  }

  public ClusterLayout addSpecTemplatesItem(ClusterLayoutSpecTemplates specTemplatesItem) {
    if (this.specTemplates == null) {
      this.specTemplates = new ArrayList<ClusterLayoutSpecTemplates>();
    }
    this.specTemplates.add(specTemplatesItem);
    return this;
  }

   /**
   * Get specTemplates
   * @return specTemplates
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public List<ClusterLayoutSpecTemplates> getSpecTemplates() {
    return specTemplates;
  }


  public void setSpecTemplates(List<ClusterLayoutSpecTemplates> specTemplates) {
    this.specTemplates = specTemplates;
  }


  public ClusterLayout taskSets(List<Object> taskSets) {
    
    this.taskSets = taskSets;
    return this;
  }

  public ClusterLayout addTaskSetsItem(Object taskSetsItem) {
    if (this.taskSets == null) {
      this.taskSets = new ArrayList<Object>();
    }
    this.taskSets.add(taskSetsItem);
    return this;
  }

   /**
   * Get taskSets
   * @return taskSets
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public List<Object> getTaskSets() {
    return taskSets;
  }


  public void setTaskSets(List<Object> taskSets) {
    this.taskSets = taskSets;
  }


  public ClusterLayout type(InlineResponse20079LoadBalancerMonitorLoadBalancerType type) {
    
    this.type = type;
    return this;
  }

   /**
   * Get type
   * @return type
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public InlineResponse20079LoadBalancerMonitorLoadBalancerType getType() {
    return type;
  }


  public void setType(InlineResponse20079LoadBalancerMonitorLoadBalancerType type) {
    this.type = type;
  }


  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (o == null || getClass() != o.getClass()) {
      return false;
    }
    ClusterLayout clusterLayout = (ClusterLayout) o;
    return Objects.equals(this.id, clusterLayout.id) &&
        Objects.equals(this.internalId, clusterLayout.internalId) &&
        Objects.equals(this.serverCount, clusterLayout.serverCount) &&
        Objects.equals(this.dateCreated, clusterLayout.dateCreated) &&
        Objects.equals(this.code, clusterLayout.code) &&
        Objects.equals(this.lastUpdated, clusterLayout.lastUpdated) &&
        Objects.equals(this.hasAutoScale, clusterLayout.hasAutoScale) &&
        Objects.equals(this.memoryRequirement, clusterLayout.memoryRequirement) &&
        Objects.equals(this.clusterVersion, clusterLayout.clusterVersion) &&
        Objects.equals(this.computeVersion, clusterLayout.computeVersion) &&
        Objects.equals(this.hasSettings, clusterLayout.hasSettings) &&
        Objects.equals(this.sortOrder, clusterLayout.sortOrder) &&
        Objects.equals(this.hasConfig, clusterLayout.hasConfig) &&
        Objects.equals(this.name, clusterLayout.name) &&
        Objects.equals(this.creatable, clusterLayout.creatable) &&
        Objects.equals(this.enabled, clusterLayout.enabled) &&
        Objects.equals(this.description, clusterLayout.description) &&
        Objects.equals(this.groupType, clusterLayout.groupType) &&
        Objects.equals(this.labels, clusterLayout.labels) &&
        Objects.equals(this.environmentVariables, clusterLayout.environmentVariables) &&
        Objects.equals(this.optionTypes, clusterLayout.optionTypes) &&
        Objects.equals(this.actions, clusterLayout.actions) &&
        Objects.equals(this.computeServers, clusterLayout.computeServers) &&
        Objects.equals(this.installContainerRuntime, clusterLayout.installContainerRuntime) &&
        Objects.equals(this.provisionType, clusterLayout.provisionType) &&
        Objects.equals(this.specTemplates, clusterLayout.specTemplates) &&
        Objects.equals(this.taskSets, clusterLayout.taskSets) &&
        Objects.equals(this.type, clusterLayout.type);
  }

  @Override
  public int hashCode() {
    return Objects.hash(id, internalId, serverCount, dateCreated, code, lastUpdated, hasAutoScale, memoryRequirement, clusterVersion, computeVersion, hasSettings, sortOrder, hasConfig, name, creatable, enabled, description, groupType, labels, environmentVariables, optionTypes, actions, computeServers, installContainerRuntime, provisionType, specTemplates, taskSets, type);
  }


  @Override
  public String toString() {
    StringBuilder sb = new StringBuilder();
    sb.append("class ClusterLayout {\n");
    sb.append("    id: ").append(toIndentedString(id)).append("\n");
    sb.append("    internalId: ").append(toIndentedString(internalId)).append("\n");
    sb.append("    serverCount: ").append(toIndentedString(serverCount)).append("\n");
    sb.append("    dateCreated: ").append(toIndentedString(dateCreated)).append("\n");
    sb.append("    code: ").append(toIndentedString(code)).append("\n");
    sb.append("    lastUpdated: ").append(toIndentedString(lastUpdated)).append("\n");
    sb.append("    hasAutoScale: ").append(toIndentedString(hasAutoScale)).append("\n");
    sb.append("    memoryRequirement: ").append(toIndentedString(memoryRequirement)).append("\n");
    sb.append("    clusterVersion: ").append(toIndentedString(clusterVersion)).append("\n");
    sb.append("    computeVersion: ").append(toIndentedString(computeVersion)).append("\n");
    sb.append("    hasSettings: ").append(toIndentedString(hasSettings)).append("\n");
    sb.append("    sortOrder: ").append(toIndentedString(sortOrder)).append("\n");
    sb.append("    hasConfig: ").append(toIndentedString(hasConfig)).append("\n");
    sb.append("    name: ").append(toIndentedString(name)).append("\n");
    sb.append("    creatable: ").append(toIndentedString(creatable)).append("\n");
    sb.append("    enabled: ").append(toIndentedString(enabled)).append("\n");
    sb.append("    description: ").append(toIndentedString(description)).append("\n");
    sb.append("    groupType: ").append(toIndentedString(groupType)).append("\n");
    sb.append("    labels: ").append(toIndentedString(labels)).append("\n");
    sb.append("    environmentVariables: ").append(toIndentedString(environmentVariables)).append("\n");
    sb.append("    optionTypes: ").append(toIndentedString(optionTypes)).append("\n");
    sb.append("    actions: ").append(toIndentedString(actions)).append("\n");
    sb.append("    computeServers: ").append(toIndentedString(computeServers)).append("\n");
    sb.append("    installContainerRuntime: ").append(toIndentedString(installContainerRuntime)).append("\n");
    sb.append("    provisionType: ").append(toIndentedString(provisionType)).append("\n");
    sb.append("    specTemplates: ").append(toIndentedString(specTemplates)).append("\n");
    sb.append("    taskSets: ").append(toIndentedString(taskSets)).append("\n");
    sb.append("    type: ").append(toIndentedString(type)).append("\n");
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
