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
import org.openapitools.client.model.AnyOfinstancesConfigAzureinstancesConfigVMWareinstancesConfigGCPinstancesConfigAWSobject;
import org.openapitools.client.model.ApiServersIdMakeManagedServerTags;
import org.openapitools.client.model.CatalogItemTypeInstanceScribeCloud;
import org.openapitools.client.model.CatalogItemTypeInstanceScribeGroup;
import org.openapitools.client.model.CatalogItemTypeInstanceScribeLayout;
import org.openapitools.client.model.CatalogItemTypeInstanceScribePlan;
import org.openapitools.client.model.CatalogItemTypeInstanceScribePorts;
import org.openapitools.client.model.CatalogItemTypeInstanceScribeSecurityGroups;
import org.openapitools.client.model.InstanceCreateNetwork;
import org.openapitools.client.model.InstanceCreateVolume;
import org.openapitools.client.model.ServicePlanOptions;

/**
 * CatalogItemTypeInstanceScribe
 */
@javax.annotation.Generated(value = "org.openapitools.codegen.languages.JavaClientCodegen", date = "2023-08-21T11:26:47.028Z[GMT]")
public class CatalogItemTypeInstanceScribe {
  public static final String SERIALIZED_NAME_GROUP = "group";
  @SerializedName(SERIALIZED_NAME_GROUP)
  private CatalogItemTypeInstanceScribeGroup group;

  public static final String SERIALIZED_NAME_CLOUD = "cloud";
  @SerializedName(SERIALIZED_NAME_CLOUD)
  private CatalogItemTypeInstanceScribeCloud cloud;

  public static final String SERIALIZED_NAME_TYPE = "type";
  @SerializedName(SERIALIZED_NAME_TYPE)
  private String type;

  public static final String SERIALIZED_NAME_NAME = "name";
  @SerializedName(SERIALIZED_NAME_NAME)
  private String name;

  public static final String SERIALIZED_NAME_CONFIG = "config";
  @SerializedName(SERIALIZED_NAME_CONFIG)
  private AnyOfinstancesConfigAzureinstancesConfigVMWareinstancesConfigGCPinstancesConfigAWSobject config = null;

  public static final String SERIALIZED_NAME_VOLUMES = "volumes";
  @SerializedName(SERIALIZED_NAME_VOLUMES)
  private List<InstanceCreateVolume> volumes = new ArrayList<InstanceCreateVolume>();

  public static final String SERIALIZED_NAME_HOST_NAME = "hostName";
  @SerializedName(SERIALIZED_NAME_HOST_NAME)
  private String hostName;

  public static final String SERIALIZED_NAME_ENVIRONMENT = "environment";
  @SerializedName(SERIALIZED_NAME_ENVIRONMENT)
  private String environment;

  public static final String SERIALIZED_NAME_LAYOUT = "layout";
  @SerializedName(SERIALIZED_NAME_LAYOUT)
  private CatalogItemTypeInstanceScribeLayout layout;

  public static final String SERIALIZED_NAME_PLAN = "plan";
  @SerializedName(SERIALIZED_NAME_PLAN)
  private CatalogItemTypeInstanceScribePlan plan;

  public static final String SERIALIZED_NAME_VERSION = "version";
  @SerializedName(SERIALIZED_NAME_VERSION)
  private String version;

  public static final String SERIALIZED_NAME_EVARS = "evars";
  @SerializedName(SERIALIZED_NAME_EVARS)
  private List<ApiServersIdMakeManagedServerTags> evars = null;

  public static final String SERIALIZED_NAME_SERVICE_PLAN_OPTIONS = "servicePlanOptions";
  @SerializedName(SERIALIZED_NAME_SERVICE_PLAN_OPTIONS)
  private ServicePlanOptions servicePlanOptions;

  public static final String SERIALIZED_NAME_SECURITY_GROUPS = "securityGroups";
  @SerializedName(SERIALIZED_NAME_SECURITY_GROUPS)
  private List<CatalogItemTypeInstanceScribeSecurityGroups> securityGroups = null;

  public static final String SERIALIZED_NAME_NETWORK_INTERFACES = "networkInterfaces";
  @SerializedName(SERIALIZED_NAME_NETWORK_INTERFACES)
  private List<InstanceCreateNetwork> networkInterfaces = null;

  public static final String SERIALIZED_NAME_LABELS = "labels";
  @SerializedName(SERIALIZED_NAME_LABELS)
  private List<String> labels = null;

  public static final String SERIALIZED_NAME_TAGS = "tags";
  @SerializedName(SERIALIZED_NAME_TAGS)
  private List<ApiServersIdMakeManagedServerTags> tags = null;

  public static final String SERIALIZED_NAME_METADATA = "metadata";
  @SerializedName(SERIALIZED_NAME_METADATA)
  private List<ApiServersIdMakeManagedServerTags> metadata = null;

  public static final String SERIALIZED_NAME_PORTS = "ports";
  @SerializedName(SERIALIZED_NAME_PORTS)
  private List<CatalogItemTypeInstanceScribePorts> ports = null;

  public static final String SERIALIZED_NAME_TASK_SET_ID = "taskSetId";
  @SerializedName(SERIALIZED_NAME_TASK_SET_ID)
  private Long taskSetId;

  public static final String SERIALIZED_NAME_TASK_SET_NAME = "taskSetName";
  @SerializedName(SERIALIZED_NAME_TASK_SET_NAME)
  private String taskSetName;


  public CatalogItemTypeInstanceScribe group(CatalogItemTypeInstanceScribeGroup group) {
    
    this.group = group;
    return this;
  }

   /**
   * Get group
   * @return group
  **/
  @ApiModelProperty(required = true, value = "")

  public CatalogItemTypeInstanceScribeGroup getGroup() {
    return group;
  }


  public void setGroup(CatalogItemTypeInstanceScribeGroup group) {
    this.group = group;
  }


  public CatalogItemTypeInstanceScribe cloud(CatalogItemTypeInstanceScribeCloud cloud) {
    
    this.cloud = cloud;
    return this;
  }

   /**
   * Get cloud
   * @return cloud
  **/
  @ApiModelProperty(required = true, value = "")

  public CatalogItemTypeInstanceScribeCloud getCloud() {
    return cloud;
  }


  public void setCloud(CatalogItemTypeInstanceScribeCloud cloud) {
    this.cloud = cloud;
  }


  public CatalogItemTypeInstanceScribe type(String type) {
    
    this.type = type;
    return this;
  }

   /**
   * The type of instance by code we want to fetch.
   * @return type
  **/
  @ApiModelProperty(example = "ubuntu", required = true, value = "The type of instance by code we want to fetch.")

  public String getType() {
    return type;
  }


  public void setType(String type) {
    this.type = type;
  }


  public CatalogItemTypeInstanceScribe name(String name) {
    
    this.name = name;
    return this;
  }

   /**
   * Name of the instance to be created.
   * @return name
  **/
  @ApiModelProperty(example = "myinstance", required = true, value = "Name of the instance to be created.")

  public String getName() {
    return name;
  }


  public void setName(String name) {
    this.name = name;
  }


  public CatalogItemTypeInstanceScribe config(AnyOfinstancesConfigAzureinstancesConfigVMWareinstancesConfigGCPinstancesConfigAWSobject config) {
    
    this.config = config;
    return this;
  }

   /**
   * Get config
   * @return config
  **/
  @ApiModelProperty(required = true, value = "")

  public AnyOfinstancesConfigAzureinstancesConfigVMWareinstancesConfigGCPinstancesConfigAWSobject getConfig() {
    return config;
  }


  public void setConfig(AnyOfinstancesConfigAzureinstancesConfigVMWareinstancesConfigGCPinstancesConfigAWSobject config) {
    this.config = config;
  }


  public CatalogItemTypeInstanceScribe volumes(List<InstanceCreateVolume> volumes) {
    
    this.volumes = volumes;
    return this;
  }

  public CatalogItemTypeInstanceScribe addVolumesItem(InstanceCreateVolume volumesItem) {
    this.volumes.add(volumesItem);
    return this;
  }

   /**
   * The (optional) volumes parameter is for LV configuration, can create additional LVs at provision It should be passed as an array of
   * @return volumes
  **/
  @ApiModelProperty(required = true, value = "The (optional) volumes parameter is for LV configuration, can create additional LVs at provision It should be passed as an array of")

  public List<InstanceCreateVolume> getVolumes() {
    return volumes;
  }


  public void setVolumes(List<InstanceCreateVolume> volumes) {
    this.volumes = volumes;
  }


  public CatalogItemTypeInstanceScribe hostName(String hostName) {
    
    this.hostName = hostName;
    return this;
  }

   /**
   * Hostname of the instance to be created.  Can be the same as the instance name.
   * @return hostName
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(example = "myinstance", value = "Hostname of the instance to be created.  Can be the same as the instance name.")

  public String getHostName() {
    return hostName;
  }


  public void setHostName(String hostName) {
    this.hostName = hostName;
  }


  public CatalogItemTypeInstanceScribe environment(String environment) {
    
    this.environment = environment;
    return this;
  }

   /**
   * Environment code
   * @return environment
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "Environment code")

  public String getEnvironment() {
    return environment;
  }


  public void setEnvironment(String environment) {
    this.environment = environment;
  }


  public CatalogItemTypeInstanceScribe layout(CatalogItemTypeInstanceScribeLayout layout) {
    
    this.layout = layout;
    return this;
  }

   /**
   * Get layout
   * @return layout
  **/
  @ApiModelProperty(required = true, value = "")

  public CatalogItemTypeInstanceScribeLayout getLayout() {
    return layout;
  }


  public void setLayout(CatalogItemTypeInstanceScribeLayout layout) {
    this.layout = layout;
  }


  public CatalogItemTypeInstanceScribe plan(CatalogItemTypeInstanceScribePlan plan) {
    
    this.plan = plan;
    return this;
  }

   /**
   * Get plan
   * @return plan
  **/
  @ApiModelProperty(required = true, value = "")

  public CatalogItemTypeInstanceScribePlan getPlan() {
    return plan;
  }


  public void setPlan(CatalogItemTypeInstanceScribePlan plan) {
    this.plan = plan;
  }


  public CatalogItemTypeInstanceScribe version(String version) {
    
    this.version = version;
    return this;
  }

   /**
   * Version of the layout to create.
   * @return version
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "Version of the layout to create.")

  public String getVersion() {
    return version;
  }


  public void setVersion(String version) {
    this.version = version;
  }


  public CatalogItemTypeInstanceScribe evars(List<ApiServersIdMakeManagedServerTags> evars) {
    
    this.evars = evars;
    return this;
  }

  public CatalogItemTypeInstanceScribe addEvarsItem(ApiServersIdMakeManagedServerTags evarsItem) {
    if (this.evars == null) {
      this.evars = new ArrayList<ApiServersIdMakeManagedServerTags>();
    }
    this.evars.add(evarsItem);
    return this;
  }

   /**
   * Environment Variables, an array of objects that have name and value.
   * @return evars
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "Environment Variables, an array of objects that have name and value.")

  public List<ApiServersIdMakeManagedServerTags> getEvars() {
    return evars;
  }


  public void setEvars(List<ApiServersIdMakeManagedServerTags> evars) {
    this.evars = evars;
  }


  public CatalogItemTypeInstanceScribe servicePlanOptions(ServicePlanOptions servicePlanOptions) {
    
    this.servicePlanOptions = servicePlanOptions;
    return this;
  }

   /**
   * Get servicePlanOptions
   * @return servicePlanOptions
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public ServicePlanOptions getServicePlanOptions() {
    return servicePlanOptions;
  }


  public void setServicePlanOptions(ServicePlanOptions servicePlanOptions) {
    this.servicePlanOptions = servicePlanOptions;
  }


  public CatalogItemTypeInstanceScribe securityGroups(List<CatalogItemTypeInstanceScribeSecurityGroups> securityGroups) {
    
    this.securityGroups = securityGroups;
    return this;
  }

  public CatalogItemTypeInstanceScribe addSecurityGroupsItem(CatalogItemTypeInstanceScribeSecurityGroups securityGroupsItem) {
    if (this.securityGroups == null) {
      this.securityGroups = new ArrayList<CatalogItemTypeInstanceScribeSecurityGroups>();
    }
    this.securityGroups.add(securityGroupsItem);
    return this;
  }

   /**
   * Key for security group configuration. It should be passed as an array of objects containing the id of the security group to assign the instance to.
   * @return securityGroups
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "Key for security group configuration. It should be passed as an array of objects containing the id of the security group to assign the instance to.")

  public List<CatalogItemTypeInstanceScribeSecurityGroups> getSecurityGroups() {
    return securityGroups;
  }


  public void setSecurityGroups(List<CatalogItemTypeInstanceScribeSecurityGroups> securityGroups) {
    this.securityGroups = securityGroups;
  }


  public CatalogItemTypeInstanceScribe networkInterfaces(List<InstanceCreateNetwork> networkInterfaces) {
    
    this.networkInterfaces = networkInterfaces;
    return this;
  }

  public CatalogItemTypeInstanceScribe addNetworkInterfacesItem(InstanceCreateNetwork networkInterfacesItem) {
    if (this.networkInterfaces == null) {
      this.networkInterfaces = new ArrayList<InstanceCreateNetwork>();
    }
    this.networkInterfaces.add(networkInterfacesItem);
    return this;
  }

   /**
   * The networkInterfaces parameter is for network configuration.  The Options API &#x60;/api/options/zoneNetworkOptions?zoneId&#x3D;5&amp;provisionTypeId&#x3D;10&#x60; can be used to see which options are available. 
   * @return networkInterfaces
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "The networkInterfaces parameter is for network configuration.  The Options API `/api/options/zoneNetworkOptions?zoneId=5&provisionTypeId=10` can be used to see which options are available. ")

  public List<InstanceCreateNetwork> getNetworkInterfaces() {
    return networkInterfaces;
  }


  public void setNetworkInterfaces(List<InstanceCreateNetwork> networkInterfaces) {
    this.networkInterfaces = networkInterfaces;
  }


  public CatalogItemTypeInstanceScribe labels(List<String> labels) {
    
    this.labels = labels;
    return this;
  }

  public CatalogItemTypeInstanceScribe addLabelsItem(String labelsItem) {
    if (this.labels == null) {
      this.labels = new ArrayList<String>();
    }
    this.labels.add(labelsItem);
    return this;
  }

   /**
   * Array of strings (keywords).
   * @return labels
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "Array of strings (keywords).")

  public List<String> getLabels() {
    return labels;
  }


  public void setLabels(List<String> labels) {
    this.labels = labels;
  }


  public CatalogItemTypeInstanceScribe tags(List<ApiServersIdMakeManagedServerTags> tags) {
    
    this.tags = tags;
    return this;
  }

  public CatalogItemTypeInstanceScribe addTagsItem(ApiServersIdMakeManagedServerTags tagsItem) {
    if (this.tags == null) {
      this.tags = new ArrayList<ApiServersIdMakeManagedServerTags>();
    }
    this.tags.add(tagsItem);
    return this;
  }

   /**
   * Metadata tags, Array of objects having a name and value.
   * @return tags
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "Metadata tags, Array of objects having a name and value.")

  public List<ApiServersIdMakeManagedServerTags> getTags() {
    return tags;
  }


  public void setTags(List<ApiServersIdMakeManagedServerTags> tags) {
    this.tags = tags;
  }


  public CatalogItemTypeInstanceScribe metadata(List<ApiServersIdMakeManagedServerTags> metadata) {
    
    this.metadata = metadata;
    return this;
  }

  public CatalogItemTypeInstanceScribe addMetadataItem(ApiServersIdMakeManagedServerTags metadataItem) {
    if (this.metadata == null) {
      this.metadata = new ArrayList<ApiServersIdMakeManagedServerTags>();
    }
    this.metadata.add(metadataItem);
    return this;
  }

   /**
   * Alias for &#x60;tags&#x60;.
   * @return metadata
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "Alias for `tags`.")

  public List<ApiServersIdMakeManagedServerTags> getMetadata() {
    return metadata;
  }


  public void setMetadata(List<ApiServersIdMakeManagedServerTags> metadata) {
    this.metadata = metadata;
  }


  public CatalogItemTypeInstanceScribe ports(List<CatalogItemTypeInstanceScribePorts> ports) {
    
    this.ports = ports;
    return this;
  }

  public CatalogItemTypeInstanceScribe addPortsItem(CatalogItemTypeInstanceScribePorts portsItem) {
    if (this.ports == null) {
      this.ports = new ArrayList<CatalogItemTypeInstanceScribePorts>();
    }
    this.ports.add(portsItem);
    return this;
  }

   /**
   * The ports parameter is for port configuration.  The layout may have default ports, which are defined in node types, that are always configured. This parameter will be for additional custom ports to be opened. 
   * @return ports
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "The ports parameter is for port configuration.  The layout may have default ports, which are defined in node types, that are always configured. This parameter will be for additional custom ports to be opened. ")

  public List<CatalogItemTypeInstanceScribePorts> getPorts() {
    return ports;
  }


  public void setPorts(List<CatalogItemTypeInstanceScribePorts> ports) {
    this.ports = ports;
  }


  public CatalogItemTypeInstanceScribe taskSetId(Long taskSetId) {
    
    this.taskSetId = taskSetId;
    return this;
  }

   /**
   * The Workflow ID to execute.
   * @return taskSetId
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "The Workflow ID to execute.")

  public Long getTaskSetId() {
    return taskSetId;
  }


  public void setTaskSetId(Long taskSetId) {
    this.taskSetId = taskSetId;
  }


  public CatalogItemTypeInstanceScribe taskSetName(String taskSetName) {
    
    this.taskSetName = taskSetName;
    return this;
  }

   /**
   * The Workflow Name to execute.
   * @return taskSetName
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "The Workflow Name to execute.")

  public String getTaskSetName() {
    return taskSetName;
  }


  public void setTaskSetName(String taskSetName) {
    this.taskSetName = taskSetName;
  }


  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (o == null || getClass() != o.getClass()) {
      return false;
    }
    CatalogItemTypeInstanceScribe catalogItemTypeInstanceScribe = (CatalogItemTypeInstanceScribe) o;
    return Objects.equals(this.group, catalogItemTypeInstanceScribe.group) &&
        Objects.equals(this.cloud, catalogItemTypeInstanceScribe.cloud) &&
        Objects.equals(this.type, catalogItemTypeInstanceScribe.type) &&
        Objects.equals(this.name, catalogItemTypeInstanceScribe.name) &&
        Objects.equals(this.config, catalogItemTypeInstanceScribe.config) &&
        Objects.equals(this.volumes, catalogItemTypeInstanceScribe.volumes) &&
        Objects.equals(this.hostName, catalogItemTypeInstanceScribe.hostName) &&
        Objects.equals(this.environment, catalogItemTypeInstanceScribe.environment) &&
        Objects.equals(this.layout, catalogItemTypeInstanceScribe.layout) &&
        Objects.equals(this.plan, catalogItemTypeInstanceScribe.plan) &&
        Objects.equals(this.version, catalogItemTypeInstanceScribe.version) &&
        Objects.equals(this.evars, catalogItemTypeInstanceScribe.evars) &&
        Objects.equals(this.servicePlanOptions, catalogItemTypeInstanceScribe.servicePlanOptions) &&
        Objects.equals(this.securityGroups, catalogItemTypeInstanceScribe.securityGroups) &&
        Objects.equals(this.networkInterfaces, catalogItemTypeInstanceScribe.networkInterfaces) &&
        Objects.equals(this.labels, catalogItemTypeInstanceScribe.labels) &&
        Objects.equals(this.tags, catalogItemTypeInstanceScribe.tags) &&
        Objects.equals(this.metadata, catalogItemTypeInstanceScribe.metadata) &&
        Objects.equals(this.ports, catalogItemTypeInstanceScribe.ports) &&
        Objects.equals(this.taskSetId, catalogItemTypeInstanceScribe.taskSetId) &&
        Objects.equals(this.taskSetName, catalogItemTypeInstanceScribe.taskSetName);
  }

  @Override
  public int hashCode() {
    return Objects.hash(group, cloud, type, name, config, volumes, hostName, environment, layout, plan, version, evars, servicePlanOptions, securityGroups, networkInterfaces, labels, tags, metadata, ports, taskSetId, taskSetName);
  }


  @Override
  public String toString() {
    StringBuilder sb = new StringBuilder();
    sb.append("class CatalogItemTypeInstanceScribe {\n");
    sb.append("    group: ").append(toIndentedString(group)).append("\n");
    sb.append("    cloud: ").append(toIndentedString(cloud)).append("\n");
    sb.append("    type: ").append(toIndentedString(type)).append("\n");
    sb.append("    name: ").append(toIndentedString(name)).append("\n");
    sb.append("    config: ").append(toIndentedString(config)).append("\n");
    sb.append("    volumes: ").append(toIndentedString(volumes)).append("\n");
    sb.append("    hostName: ").append(toIndentedString(hostName)).append("\n");
    sb.append("    environment: ").append(toIndentedString(environment)).append("\n");
    sb.append("    layout: ").append(toIndentedString(layout)).append("\n");
    sb.append("    plan: ").append(toIndentedString(plan)).append("\n");
    sb.append("    version: ").append(toIndentedString(version)).append("\n");
    sb.append("    evars: ").append(toIndentedString(evars)).append("\n");
    sb.append("    servicePlanOptions: ").append(toIndentedString(servicePlanOptions)).append("\n");
    sb.append("    securityGroups: ").append(toIndentedString(securityGroups)).append("\n");
    sb.append("    networkInterfaces: ").append(toIndentedString(networkInterfaces)).append("\n");
    sb.append("    labels: ").append(toIndentedString(labels)).append("\n");
    sb.append("    tags: ").append(toIndentedString(tags)).append("\n");
    sb.append("    metadata: ").append(toIndentedString(metadata)).append("\n");
    sb.append("    ports: ").append(toIndentedString(ports)).append("\n");
    sb.append("    taskSetId: ").append(toIndentedString(taskSetId)).append("\n");
    sb.append("    taskSetName: ").append(toIndentedString(taskSetName)).append("\n");
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
