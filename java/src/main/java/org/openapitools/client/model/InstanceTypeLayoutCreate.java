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
import org.openapitools.client.model.ApiLibraryLayoutsIdPermissionsInstanceTypeLayoutPermissions;
import org.openapitools.client.model.ClusterLayoutCreateEnvironmentVariables;
import org.openapitools.client.model.InstanceTypeCreatePriceSets;

/**
 * InstanceTypeLayoutCreate
 */
@javax.annotation.Generated(value = "org.openapitools.codegen.languages.JavaClientCodegen", date = "2023-08-21T11:26:47.028Z[GMT]")
public class InstanceTypeLayoutCreate {
  public static final String SERIALIZED_NAME_NAME = "name";
  @SerializedName(SERIALIZED_NAME_NAME)
  private String name;

  public static final String SERIALIZED_NAME_LABELS = "labels";
  @SerializedName(SERIALIZED_NAME_LABELS)
  private List<String> labels = null;

  public static final String SERIALIZED_NAME_INSTANCE_VERSION = "instanceVersion";
  @SerializedName(SERIALIZED_NAME_INSTANCE_VERSION)
  private String instanceVersion;

  public static final String SERIALIZED_NAME_DESCRIPTION = "description";
  @SerializedName(SERIALIZED_NAME_DESCRIPTION)
  private String description;

  public static final String SERIALIZED_NAME_CREATABLE = "creatable";
  @SerializedName(SERIALIZED_NAME_CREATABLE)
  private Boolean creatable = true;

  public static final String SERIALIZED_NAME_PROVISION_TYPE_CODE = "provisionTypeCode";
  @SerializedName(SERIALIZED_NAME_PROVISION_TYPE_CODE)
  private String provisionTypeCode;

  public static final String SERIALIZED_NAME_MEMORY_REQUIREMENT = "memoryRequirement";
  @SerializedName(SERIALIZED_NAME_MEMORY_REQUIREMENT)
  private String memoryRequirement;

  public static final String SERIALIZED_NAME_HAS_AUTO_SCALE = "hasAutoScale";
  @SerializedName(SERIALIZED_NAME_HAS_AUTO_SCALE)
  private Boolean hasAutoScale = false;

  public static final String SERIALIZED_NAME_SUPPORTS_CONVERT_TO_MANAGED = "supportsConvertToManaged";
  @SerializedName(SERIALIZED_NAME_SUPPORTS_CONVERT_TO_MANAGED)
  private Boolean supportsConvertToManaged = false;

  public static final String SERIALIZED_NAME_CONTAINER_TYPES = "containerTypes";
  @SerializedName(SERIALIZED_NAME_CONTAINER_TYPES)
  private List<Long> containerTypes = null;

  public static final String SERIALIZED_NAME_OPTION_TYPES = "optionTypes";
  @SerializedName(SERIALIZED_NAME_OPTION_TYPES)
  private List<Long> optionTypes = null;

  public static final String SERIALIZED_NAME_SPEC_TEMPLATES = "specTemplates";
  @SerializedName(SERIALIZED_NAME_SPEC_TEMPLATES)
  private List<Long> specTemplates = null;

  public static final String SERIALIZED_NAME_ENVIRONMENT_VARIABLES = "environmentVariables";
  @SerializedName(SERIALIZED_NAME_ENVIRONMENT_VARIABLES)
  private List<ClusterLayoutCreateEnvironmentVariables> environmentVariables = null;

  public static final String SERIALIZED_NAME_PRICE_SETS = "priceSets";
  @SerializedName(SERIALIZED_NAME_PRICE_SETS)
  private List<InstanceTypeCreatePriceSets> priceSets = null;

  public static final String SERIALIZED_NAME_PERMISSIONS = "permissions";
  @SerializedName(SERIALIZED_NAME_PERMISSIONS)
  private ApiLibraryLayoutsIdPermissionsInstanceTypeLayoutPermissions permissions;


  public InstanceTypeLayoutCreate name(String name) {
    
    this.name = name;
    return this;
  }

   /**
   * Layout name
   * @return name
  **/
  @ApiModelProperty(required = true, value = "Layout name")

  public String getName() {
    return name;
  }


  public void setName(String name) {
    this.name = name;
  }


  public InstanceTypeLayoutCreate labels(List<String> labels) {
    
    this.labels = labels;
    return this;
  }

  public InstanceTypeLayoutCreate addLabelsItem(String labelsItem) {
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


  public InstanceTypeLayoutCreate instanceVersion(String instanceVersion) {
    
    this.instanceVersion = instanceVersion;
    return this;
  }

   /**
   * Version of the layout
   * @return instanceVersion
  **/
  @ApiModelProperty(required = true, value = "Version of the layout")

  public String getInstanceVersion() {
    return instanceVersion;
  }


  public void setInstanceVersion(String instanceVersion) {
    this.instanceVersion = instanceVersion;
  }


  public InstanceTypeLayoutCreate description(String description) {
    
    this.description = description;
    return this;
  }

   /**
   * Layout description
   * @return description
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "Layout description")

  public String getDescription() {
    return description;
  }


  public void setDescription(String description) {
    this.description = description;
  }


  public InstanceTypeLayoutCreate creatable(Boolean creatable) {
    
    this.creatable = creatable;
    return this;
  }

   /**
   * Can be used to enable / disable the creatability of the layout.
   * @return creatable
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "Can be used to enable / disable the creatability of the layout.")

  public Boolean getCreatable() {
    return creatable;
  }


  public void setCreatable(Boolean creatable) {
    this.creatable = creatable;
  }


  public InstanceTypeLayoutCreate provisionTypeCode(String provisionTypeCode) {
    
    this.provisionTypeCode = provisionTypeCode;
    return this;
  }

   /**
   * Provision type code
   * @return provisionTypeCode
  **/
  @ApiModelProperty(required = true, value = "Provision type code")

  public String getProvisionTypeCode() {
    return provisionTypeCode;
  }


  public void setProvisionTypeCode(String provisionTypeCode) {
    this.provisionTypeCode = provisionTypeCode;
  }


  public InstanceTypeLayoutCreate memoryRequirement(String memoryRequirement) {
    
    this.memoryRequirement = memoryRequirement;
    return this;
  }

   /**
   * Memory requirement in megabytes
   * @return memoryRequirement
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "Memory requirement in megabytes")

  public String getMemoryRequirement() {
    return memoryRequirement;
  }


  public void setMemoryRequirement(String memoryRequirement) {
    this.memoryRequirement = memoryRequirement;
  }


  public InstanceTypeLayoutCreate hasAutoScale(Boolean hasAutoScale) {
    
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


  public InstanceTypeLayoutCreate supportsConvertToManaged(Boolean supportsConvertToManaged) {
    
    this.supportsConvertToManaged = supportsConvertToManaged;
    return this;
  }

   /**
   * Can be used to enable / disable the supports convert to managed.
   * @return supportsConvertToManaged
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "Can be used to enable / disable the supports convert to managed.")

  public Boolean getSupportsConvertToManaged() {
    return supportsConvertToManaged;
  }


  public void setSupportsConvertToManaged(Boolean supportsConvertToManaged) {
    this.supportsConvertToManaged = supportsConvertToManaged;
  }


  public InstanceTypeLayoutCreate containerTypes(List<Long> containerTypes) {
    
    this.containerTypes = containerTypes;
    return this;
  }

  public InstanceTypeLayoutCreate addContainerTypesItem(Long containerTypesItem) {
    if (this.containerTypes == null) {
      this.containerTypes = new ArrayList<Long>();
    }
    this.containerTypes.add(containerTypesItem);
    return this;
  }

   /**
   * Array of layout node type IDs
   * @return containerTypes
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "Array of layout node type IDs")

  public List<Long> getContainerTypes() {
    return containerTypes;
  }


  public void setContainerTypes(List<Long> containerTypes) {
    this.containerTypes = containerTypes;
  }


  public InstanceTypeLayoutCreate optionTypes(List<Long> optionTypes) {
    
    this.optionTypes = optionTypes;
    return this;
  }

  public InstanceTypeLayoutCreate addOptionTypesItem(Long optionTypesItem) {
    if (this.optionTypes == null) {
      this.optionTypes = new ArrayList<Long>();
    }
    this.optionTypes.add(optionTypesItem);
    return this;
  }

   /**
   * Array of layout option type IDs
   * @return optionTypes
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "Array of layout option type IDs")

  public List<Long> getOptionTypes() {
    return optionTypes;
  }


  public void setOptionTypes(List<Long> optionTypes) {
    this.optionTypes = optionTypes;
  }


  public InstanceTypeLayoutCreate specTemplates(List<Long> specTemplates) {
    
    this.specTemplates = specTemplates;
    return this;
  }

  public InstanceTypeLayoutCreate addSpecTemplatesItem(Long specTemplatesItem) {
    if (this.specTemplates == null) {
      this.specTemplates = new ArrayList<Long>();
    }
    this.specTemplates.add(specTemplatesItem);
    return this;
  }

   /**
   * Array of layout spec template IDs
   * @return specTemplates
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "Array of layout spec template IDs")

  public List<Long> getSpecTemplates() {
    return specTemplates;
  }


  public void setSpecTemplates(List<Long> specTemplates) {
    this.specTemplates = specTemplates;
  }


  public InstanceTypeLayoutCreate environmentVariables(List<ClusterLayoutCreateEnvironmentVariables> environmentVariables) {
    
    this.environmentVariables = environmentVariables;
    return this;
  }

  public InstanceTypeLayoutCreate addEnvironmentVariablesItem(ClusterLayoutCreateEnvironmentVariables environmentVariablesItem) {
    if (this.environmentVariables == null) {
      this.environmentVariables = new ArrayList<ClusterLayoutCreateEnvironmentVariables>();
    }
    this.environmentVariables.add(environmentVariablesItem);
    return this;
  }

   /**
   * The environmentVariables parameter is array of env objects
   * @return environmentVariables
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "The environmentVariables parameter is array of env objects")

  public List<ClusterLayoutCreateEnvironmentVariables> getEnvironmentVariables() {
    return environmentVariables;
  }


  public void setEnvironmentVariables(List<ClusterLayoutCreateEnvironmentVariables> environmentVariables) {
    this.environmentVariables = environmentVariables;
  }


  public InstanceTypeLayoutCreate priceSets(List<InstanceTypeCreatePriceSets> priceSets) {
    
    this.priceSets = priceSets;
    return this;
  }

  public InstanceTypeLayoutCreate addPriceSetsItem(InstanceTypeCreatePriceSets priceSetsItem) {
    if (this.priceSets == null) {
      this.priceSets = new ArrayList<InstanceTypeCreatePriceSets>();
    }
    this.priceSets.add(priceSetsItem);
    return this;
  }

   /**
   * Array of price set objects
   * @return priceSets
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "Array of price set objects")

  public List<InstanceTypeCreatePriceSets> getPriceSets() {
    return priceSets;
  }


  public void setPriceSets(List<InstanceTypeCreatePriceSets> priceSets) {
    this.priceSets = priceSets;
  }


  public InstanceTypeLayoutCreate permissions(ApiLibraryLayoutsIdPermissionsInstanceTypeLayoutPermissions permissions) {
    
    this.permissions = permissions;
    return this;
  }

   /**
   * Get permissions
   * @return permissions
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public ApiLibraryLayoutsIdPermissionsInstanceTypeLayoutPermissions getPermissions() {
    return permissions;
  }


  public void setPermissions(ApiLibraryLayoutsIdPermissionsInstanceTypeLayoutPermissions permissions) {
    this.permissions = permissions;
  }


  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (o == null || getClass() != o.getClass()) {
      return false;
    }
    InstanceTypeLayoutCreate instanceTypeLayoutCreate = (InstanceTypeLayoutCreate) o;
    return Objects.equals(this.name, instanceTypeLayoutCreate.name) &&
        Objects.equals(this.labels, instanceTypeLayoutCreate.labels) &&
        Objects.equals(this.instanceVersion, instanceTypeLayoutCreate.instanceVersion) &&
        Objects.equals(this.description, instanceTypeLayoutCreate.description) &&
        Objects.equals(this.creatable, instanceTypeLayoutCreate.creatable) &&
        Objects.equals(this.provisionTypeCode, instanceTypeLayoutCreate.provisionTypeCode) &&
        Objects.equals(this.memoryRequirement, instanceTypeLayoutCreate.memoryRequirement) &&
        Objects.equals(this.hasAutoScale, instanceTypeLayoutCreate.hasAutoScale) &&
        Objects.equals(this.supportsConvertToManaged, instanceTypeLayoutCreate.supportsConvertToManaged) &&
        Objects.equals(this.containerTypes, instanceTypeLayoutCreate.containerTypes) &&
        Objects.equals(this.optionTypes, instanceTypeLayoutCreate.optionTypes) &&
        Objects.equals(this.specTemplates, instanceTypeLayoutCreate.specTemplates) &&
        Objects.equals(this.environmentVariables, instanceTypeLayoutCreate.environmentVariables) &&
        Objects.equals(this.priceSets, instanceTypeLayoutCreate.priceSets) &&
        Objects.equals(this.permissions, instanceTypeLayoutCreate.permissions);
  }

  @Override
  public int hashCode() {
    return Objects.hash(name, labels, instanceVersion, description, creatable, provisionTypeCode, memoryRequirement, hasAutoScale, supportsConvertToManaged, containerTypes, optionTypes, specTemplates, environmentVariables, priceSets, permissions);
  }


  @Override
  public String toString() {
    StringBuilder sb = new StringBuilder();
    sb.append("class InstanceTypeLayoutCreate {\n");
    sb.append("    name: ").append(toIndentedString(name)).append("\n");
    sb.append("    labels: ").append(toIndentedString(labels)).append("\n");
    sb.append("    instanceVersion: ").append(toIndentedString(instanceVersion)).append("\n");
    sb.append("    description: ").append(toIndentedString(description)).append("\n");
    sb.append("    creatable: ").append(toIndentedString(creatable)).append("\n");
    sb.append("    provisionTypeCode: ").append(toIndentedString(provisionTypeCode)).append("\n");
    sb.append("    memoryRequirement: ").append(toIndentedString(memoryRequirement)).append("\n");
    sb.append("    hasAutoScale: ").append(toIndentedString(hasAutoScale)).append("\n");
    sb.append("    supportsConvertToManaged: ").append(toIndentedString(supportsConvertToManaged)).append("\n");
    sb.append("    containerTypes: ").append(toIndentedString(containerTypes)).append("\n");
    sb.append("    optionTypes: ").append(toIndentedString(optionTypes)).append("\n");
    sb.append("    specTemplates: ").append(toIndentedString(specTemplates)).append("\n");
    sb.append("    environmentVariables: ").append(toIndentedString(environmentVariables)).append("\n");
    sb.append("    priceSets: ").append(toIndentedString(priceSets)).append("\n");
    sb.append("    permissions: ").append(toIndentedString(permissions)).append("\n");
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

