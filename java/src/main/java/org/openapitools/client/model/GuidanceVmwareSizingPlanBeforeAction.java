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
import org.openapitools.client.model.GuidanceVmwareSizingPlanBeforeActionConfig;
import org.openapitools.client.model.GuidanceVmwareSizingPlanBeforeActionPriceSets;
import org.openapitools.client.model.GuidanceVmwareSizingPlanBeforeActionProvisionType;
import org.threeten.bp.OffsetDateTime;

/**
 * GuidanceVmwareSizingPlanBeforeAction
 */
@javax.annotation.Generated(value = "org.openapitools.codegen.languages.JavaClientCodegen", date = "2023-08-21T11:26:47.028Z[GMT]")
public class GuidanceVmwareSizingPlanBeforeAction {
  public static final String SERIALIZED_NAME_ID = "id";
  @SerializedName(SERIALIZED_NAME_ID)
  private Long id;

  public static final String SERIALIZED_NAME_NAME = "name";
  @SerializedName(SERIALIZED_NAME_NAME)
  private String name;

  public static final String SERIALIZED_NAME_CODE = "code";
  @SerializedName(SERIALIZED_NAME_CODE)
  private String code;

  public static final String SERIALIZED_NAME_ACTIVE = "active";
  @SerializedName(SERIALIZED_NAME_ACTIVE)
  private Boolean active;

  public static final String SERIALIZED_NAME_SORT_ORDER = "sortOrder";
  @SerializedName(SERIALIZED_NAME_SORT_ORDER)
  private Long sortOrder;

  public static final String SERIALIZED_NAME_DESCRIPTION = "description";
  @SerializedName(SERIALIZED_NAME_DESCRIPTION)
  private String description;

  public static final String SERIALIZED_NAME_MAX_STORAGE = "maxStorage";
  @SerializedName(SERIALIZED_NAME_MAX_STORAGE)
  private Long maxStorage;

  public static final String SERIALIZED_NAME_MAX_MEMORY = "maxMemory";
  @SerializedName(SERIALIZED_NAME_MAX_MEMORY)
  private Long maxMemory;

  public static final String SERIALIZED_NAME_MAX_CPU = "maxCpu";
  @SerializedName(SERIALIZED_NAME_MAX_CPU)
  private String maxCpu;

  public static final String SERIALIZED_NAME_MAX_CORES = "maxCores";
  @SerializedName(SERIALIZED_NAME_MAX_CORES)
  private Long maxCores;

  public static final String SERIALIZED_NAME_MAX_DISKS = "maxDisks";
  @SerializedName(SERIALIZED_NAME_MAX_DISKS)
  private String maxDisks;

  public static final String SERIALIZED_NAME_CORES_PER_SOCKET = "coresPerSocket";
  @SerializedName(SERIALIZED_NAME_CORES_PER_SOCKET)
  private Long coresPerSocket;

  public static final String SERIALIZED_NAME_CUSTOM_CPU = "customCpu";
  @SerializedName(SERIALIZED_NAME_CUSTOM_CPU)
  private Boolean customCpu;

  public static final String SERIALIZED_NAME_CUSTOM_CORES = "customCores";
  @SerializedName(SERIALIZED_NAME_CUSTOM_CORES)
  private Boolean customCores;

  public static final String SERIALIZED_NAME_CUSTOM_MAX_STORAGE = "customMaxStorage";
  @SerializedName(SERIALIZED_NAME_CUSTOM_MAX_STORAGE)
  private Boolean customMaxStorage;

  public static final String SERIALIZED_NAME_CUSTOM_MAX_DATA_STORAGE = "customMaxDataStorage";
  @SerializedName(SERIALIZED_NAME_CUSTOM_MAX_DATA_STORAGE)
  private Boolean customMaxDataStorage;

  public static final String SERIALIZED_NAME_CUSTOM_MAX_MEMORY = "customMaxMemory";
  @SerializedName(SERIALIZED_NAME_CUSTOM_MAX_MEMORY)
  private Boolean customMaxMemory;

  public static final String SERIALIZED_NAME_ADD_VOLUMES = "addVolumes";
  @SerializedName(SERIALIZED_NAME_ADD_VOLUMES)
  private Boolean addVolumes;

  public static final String SERIALIZED_NAME_MEMORY_OPTION_SOURCE = "memoryOptionSource";
  @SerializedName(SERIALIZED_NAME_MEMORY_OPTION_SOURCE)
  private String memoryOptionSource;

  public static final String SERIALIZED_NAME_CPU_OPTION_SOURCE = "cpuOptionSource";
  @SerializedName(SERIALIZED_NAME_CPU_OPTION_SOURCE)
  private String cpuOptionSource;

  public static final String SERIALIZED_NAME_DATE_CREATED = "dateCreated";
  @SerializedName(SERIALIZED_NAME_DATE_CREATED)
  private OffsetDateTime dateCreated;

  public static final String SERIALIZED_NAME_LAST_UPDATED = "lastUpdated";
  @SerializedName(SERIALIZED_NAME_LAST_UPDATED)
  private OffsetDateTime lastUpdated;

  public static final String SERIALIZED_NAME_REGION_CODE = "regionCode";
  @SerializedName(SERIALIZED_NAME_REGION_CODE)
  private String regionCode;

  public static final String SERIALIZED_NAME_VISIBILITY = "visibility";
  @SerializedName(SERIALIZED_NAME_VISIBILITY)
  private String visibility;

  public static final String SERIALIZED_NAME_EDITABLE = "editable";
  @SerializedName(SERIALIZED_NAME_EDITABLE)
  private Boolean editable;

  public static final String SERIALIZED_NAME_PROVISION_TYPE = "provisionType";
  @SerializedName(SERIALIZED_NAME_PROVISION_TYPE)
  private GuidanceVmwareSizingPlanBeforeActionProvisionType provisionType;

  public static final String SERIALIZED_NAME_TENANTS = "tenants";
  @SerializedName(SERIALIZED_NAME_TENANTS)
  private String tenants;

  public static final String SERIALIZED_NAME_PRICE_SETS = "priceSets";
  @SerializedName(SERIALIZED_NAME_PRICE_SETS)
  private List<GuidanceVmwareSizingPlanBeforeActionPriceSets> priceSets = null;

  public static final String SERIALIZED_NAME_CONFIG = "config";
  @SerializedName(SERIALIZED_NAME_CONFIG)
  private GuidanceVmwareSizingPlanBeforeActionConfig config;


  public GuidanceVmwareSizingPlanBeforeAction id(Long id) {
    
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


  public GuidanceVmwareSizingPlanBeforeAction name(String name) {
    
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


  public GuidanceVmwareSizingPlanBeforeAction code(String code) {
    
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


  public GuidanceVmwareSizingPlanBeforeAction active(Boolean active) {
    
    this.active = active;
    return this;
  }

   /**
   * Get active
   * @return active
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public Boolean getActive() {
    return active;
  }


  public void setActive(Boolean active) {
    this.active = active;
  }


  public GuidanceVmwareSizingPlanBeforeAction sortOrder(Long sortOrder) {
    
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


  public GuidanceVmwareSizingPlanBeforeAction description(String description) {
    
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


  public GuidanceVmwareSizingPlanBeforeAction maxStorage(Long maxStorage) {
    
    this.maxStorage = maxStorage;
    return this;
  }

   /**
   * Get maxStorage
   * @return maxStorage
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public Long getMaxStorage() {
    return maxStorage;
  }


  public void setMaxStorage(Long maxStorage) {
    this.maxStorage = maxStorage;
  }


  public GuidanceVmwareSizingPlanBeforeAction maxMemory(Long maxMemory) {
    
    this.maxMemory = maxMemory;
    return this;
  }

   /**
   * Get maxMemory
   * @return maxMemory
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public Long getMaxMemory() {
    return maxMemory;
  }


  public void setMaxMemory(Long maxMemory) {
    this.maxMemory = maxMemory;
  }


  public GuidanceVmwareSizingPlanBeforeAction maxCpu(String maxCpu) {
    
    this.maxCpu = maxCpu;
    return this;
  }

   /**
   * Get maxCpu
   * @return maxCpu
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public String getMaxCpu() {
    return maxCpu;
  }


  public void setMaxCpu(String maxCpu) {
    this.maxCpu = maxCpu;
  }


  public GuidanceVmwareSizingPlanBeforeAction maxCores(Long maxCores) {
    
    this.maxCores = maxCores;
    return this;
  }

   /**
   * Get maxCores
   * @return maxCores
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public Long getMaxCores() {
    return maxCores;
  }


  public void setMaxCores(Long maxCores) {
    this.maxCores = maxCores;
  }


  public GuidanceVmwareSizingPlanBeforeAction maxDisks(String maxDisks) {
    
    this.maxDisks = maxDisks;
    return this;
  }

   /**
   * Get maxDisks
   * @return maxDisks
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public String getMaxDisks() {
    return maxDisks;
  }


  public void setMaxDisks(String maxDisks) {
    this.maxDisks = maxDisks;
  }


  public GuidanceVmwareSizingPlanBeforeAction coresPerSocket(Long coresPerSocket) {
    
    this.coresPerSocket = coresPerSocket;
    return this;
  }

   /**
   * Get coresPerSocket
   * @return coresPerSocket
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public Long getCoresPerSocket() {
    return coresPerSocket;
  }


  public void setCoresPerSocket(Long coresPerSocket) {
    this.coresPerSocket = coresPerSocket;
  }


  public GuidanceVmwareSizingPlanBeforeAction customCpu(Boolean customCpu) {
    
    this.customCpu = customCpu;
    return this;
  }

   /**
   * Get customCpu
   * @return customCpu
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public Boolean getCustomCpu() {
    return customCpu;
  }


  public void setCustomCpu(Boolean customCpu) {
    this.customCpu = customCpu;
  }


  public GuidanceVmwareSizingPlanBeforeAction customCores(Boolean customCores) {
    
    this.customCores = customCores;
    return this;
  }

   /**
   * Get customCores
   * @return customCores
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public Boolean getCustomCores() {
    return customCores;
  }


  public void setCustomCores(Boolean customCores) {
    this.customCores = customCores;
  }


  public GuidanceVmwareSizingPlanBeforeAction customMaxStorage(Boolean customMaxStorage) {
    
    this.customMaxStorage = customMaxStorage;
    return this;
  }

   /**
   * Get customMaxStorage
   * @return customMaxStorage
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public Boolean getCustomMaxStorage() {
    return customMaxStorage;
  }


  public void setCustomMaxStorage(Boolean customMaxStorage) {
    this.customMaxStorage = customMaxStorage;
  }


  public GuidanceVmwareSizingPlanBeforeAction customMaxDataStorage(Boolean customMaxDataStorage) {
    
    this.customMaxDataStorage = customMaxDataStorage;
    return this;
  }

   /**
   * Get customMaxDataStorage
   * @return customMaxDataStorage
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public Boolean getCustomMaxDataStorage() {
    return customMaxDataStorage;
  }


  public void setCustomMaxDataStorage(Boolean customMaxDataStorage) {
    this.customMaxDataStorage = customMaxDataStorage;
  }


  public GuidanceVmwareSizingPlanBeforeAction customMaxMemory(Boolean customMaxMemory) {
    
    this.customMaxMemory = customMaxMemory;
    return this;
  }

   /**
   * Get customMaxMemory
   * @return customMaxMemory
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public Boolean getCustomMaxMemory() {
    return customMaxMemory;
  }


  public void setCustomMaxMemory(Boolean customMaxMemory) {
    this.customMaxMemory = customMaxMemory;
  }


  public GuidanceVmwareSizingPlanBeforeAction addVolumes(Boolean addVolumes) {
    
    this.addVolumes = addVolumes;
    return this;
  }

   /**
   * Get addVolumes
   * @return addVolumes
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public Boolean getAddVolumes() {
    return addVolumes;
  }


  public void setAddVolumes(Boolean addVolumes) {
    this.addVolumes = addVolumes;
  }


  public GuidanceVmwareSizingPlanBeforeAction memoryOptionSource(String memoryOptionSource) {
    
    this.memoryOptionSource = memoryOptionSource;
    return this;
  }

   /**
   * Get memoryOptionSource
   * @return memoryOptionSource
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public String getMemoryOptionSource() {
    return memoryOptionSource;
  }


  public void setMemoryOptionSource(String memoryOptionSource) {
    this.memoryOptionSource = memoryOptionSource;
  }


  public GuidanceVmwareSizingPlanBeforeAction cpuOptionSource(String cpuOptionSource) {
    
    this.cpuOptionSource = cpuOptionSource;
    return this;
  }

   /**
   * Get cpuOptionSource
   * @return cpuOptionSource
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public String getCpuOptionSource() {
    return cpuOptionSource;
  }


  public void setCpuOptionSource(String cpuOptionSource) {
    this.cpuOptionSource = cpuOptionSource;
  }


  public GuidanceVmwareSizingPlanBeforeAction dateCreated(OffsetDateTime dateCreated) {
    
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


  public GuidanceVmwareSizingPlanBeforeAction lastUpdated(OffsetDateTime lastUpdated) {
    
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


  public GuidanceVmwareSizingPlanBeforeAction regionCode(String regionCode) {
    
    this.regionCode = regionCode;
    return this;
  }

   /**
   * Get regionCode
   * @return regionCode
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public String getRegionCode() {
    return regionCode;
  }


  public void setRegionCode(String regionCode) {
    this.regionCode = regionCode;
  }


  public GuidanceVmwareSizingPlanBeforeAction visibility(String visibility) {
    
    this.visibility = visibility;
    return this;
  }

   /**
   * Get visibility
   * @return visibility
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public String getVisibility() {
    return visibility;
  }


  public void setVisibility(String visibility) {
    this.visibility = visibility;
  }


  public GuidanceVmwareSizingPlanBeforeAction editable(Boolean editable) {
    
    this.editable = editable;
    return this;
  }

   /**
   * Get editable
   * @return editable
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public Boolean getEditable() {
    return editable;
  }


  public void setEditable(Boolean editable) {
    this.editable = editable;
  }


  public GuidanceVmwareSizingPlanBeforeAction provisionType(GuidanceVmwareSizingPlanBeforeActionProvisionType provisionType) {
    
    this.provisionType = provisionType;
    return this;
  }

   /**
   * Get provisionType
   * @return provisionType
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public GuidanceVmwareSizingPlanBeforeActionProvisionType getProvisionType() {
    return provisionType;
  }


  public void setProvisionType(GuidanceVmwareSizingPlanBeforeActionProvisionType provisionType) {
    this.provisionType = provisionType;
  }


  public GuidanceVmwareSizingPlanBeforeAction tenants(String tenants) {
    
    this.tenants = tenants;
    return this;
  }

   /**
   * Get tenants
   * @return tenants
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public String getTenants() {
    return tenants;
  }


  public void setTenants(String tenants) {
    this.tenants = tenants;
  }


  public GuidanceVmwareSizingPlanBeforeAction priceSets(List<GuidanceVmwareSizingPlanBeforeActionPriceSets> priceSets) {
    
    this.priceSets = priceSets;
    return this;
  }

  public GuidanceVmwareSizingPlanBeforeAction addPriceSetsItem(GuidanceVmwareSizingPlanBeforeActionPriceSets priceSetsItem) {
    if (this.priceSets == null) {
      this.priceSets = new ArrayList<GuidanceVmwareSizingPlanBeforeActionPriceSets>();
    }
    this.priceSets.add(priceSetsItem);
    return this;
  }

   /**
   * Get priceSets
   * @return priceSets
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public List<GuidanceVmwareSizingPlanBeforeActionPriceSets> getPriceSets() {
    return priceSets;
  }


  public void setPriceSets(List<GuidanceVmwareSizingPlanBeforeActionPriceSets> priceSets) {
    this.priceSets = priceSets;
  }


  public GuidanceVmwareSizingPlanBeforeAction config(GuidanceVmwareSizingPlanBeforeActionConfig config) {
    
    this.config = config;
    return this;
  }

   /**
   * Get config
   * @return config
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public GuidanceVmwareSizingPlanBeforeActionConfig getConfig() {
    return config;
  }


  public void setConfig(GuidanceVmwareSizingPlanBeforeActionConfig config) {
    this.config = config;
  }


  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (o == null || getClass() != o.getClass()) {
      return false;
    }
    GuidanceVmwareSizingPlanBeforeAction guidanceVmwareSizingPlanBeforeAction = (GuidanceVmwareSizingPlanBeforeAction) o;
    return Objects.equals(this.id, guidanceVmwareSizingPlanBeforeAction.id) &&
        Objects.equals(this.name, guidanceVmwareSizingPlanBeforeAction.name) &&
        Objects.equals(this.code, guidanceVmwareSizingPlanBeforeAction.code) &&
        Objects.equals(this.active, guidanceVmwareSizingPlanBeforeAction.active) &&
        Objects.equals(this.sortOrder, guidanceVmwareSizingPlanBeforeAction.sortOrder) &&
        Objects.equals(this.description, guidanceVmwareSizingPlanBeforeAction.description) &&
        Objects.equals(this.maxStorage, guidanceVmwareSizingPlanBeforeAction.maxStorage) &&
        Objects.equals(this.maxMemory, guidanceVmwareSizingPlanBeforeAction.maxMemory) &&
        Objects.equals(this.maxCpu, guidanceVmwareSizingPlanBeforeAction.maxCpu) &&
        Objects.equals(this.maxCores, guidanceVmwareSizingPlanBeforeAction.maxCores) &&
        Objects.equals(this.maxDisks, guidanceVmwareSizingPlanBeforeAction.maxDisks) &&
        Objects.equals(this.coresPerSocket, guidanceVmwareSizingPlanBeforeAction.coresPerSocket) &&
        Objects.equals(this.customCpu, guidanceVmwareSizingPlanBeforeAction.customCpu) &&
        Objects.equals(this.customCores, guidanceVmwareSizingPlanBeforeAction.customCores) &&
        Objects.equals(this.customMaxStorage, guidanceVmwareSizingPlanBeforeAction.customMaxStorage) &&
        Objects.equals(this.customMaxDataStorage, guidanceVmwareSizingPlanBeforeAction.customMaxDataStorage) &&
        Objects.equals(this.customMaxMemory, guidanceVmwareSizingPlanBeforeAction.customMaxMemory) &&
        Objects.equals(this.addVolumes, guidanceVmwareSizingPlanBeforeAction.addVolumes) &&
        Objects.equals(this.memoryOptionSource, guidanceVmwareSizingPlanBeforeAction.memoryOptionSource) &&
        Objects.equals(this.cpuOptionSource, guidanceVmwareSizingPlanBeforeAction.cpuOptionSource) &&
        Objects.equals(this.dateCreated, guidanceVmwareSizingPlanBeforeAction.dateCreated) &&
        Objects.equals(this.lastUpdated, guidanceVmwareSizingPlanBeforeAction.lastUpdated) &&
        Objects.equals(this.regionCode, guidanceVmwareSizingPlanBeforeAction.regionCode) &&
        Objects.equals(this.visibility, guidanceVmwareSizingPlanBeforeAction.visibility) &&
        Objects.equals(this.editable, guidanceVmwareSizingPlanBeforeAction.editable) &&
        Objects.equals(this.provisionType, guidanceVmwareSizingPlanBeforeAction.provisionType) &&
        Objects.equals(this.tenants, guidanceVmwareSizingPlanBeforeAction.tenants) &&
        Objects.equals(this.priceSets, guidanceVmwareSizingPlanBeforeAction.priceSets) &&
        Objects.equals(this.config, guidanceVmwareSizingPlanBeforeAction.config);
  }

  @Override
  public int hashCode() {
    return Objects.hash(id, name, code, active, sortOrder, description, maxStorage, maxMemory, maxCpu, maxCores, maxDisks, coresPerSocket, customCpu, customCores, customMaxStorage, customMaxDataStorage, customMaxMemory, addVolumes, memoryOptionSource, cpuOptionSource, dateCreated, lastUpdated, regionCode, visibility, editable, provisionType, tenants, priceSets, config);
  }


  @Override
  public String toString() {
    StringBuilder sb = new StringBuilder();
    sb.append("class GuidanceVmwareSizingPlanBeforeAction {\n");
    sb.append("    id: ").append(toIndentedString(id)).append("\n");
    sb.append("    name: ").append(toIndentedString(name)).append("\n");
    sb.append("    code: ").append(toIndentedString(code)).append("\n");
    sb.append("    active: ").append(toIndentedString(active)).append("\n");
    sb.append("    sortOrder: ").append(toIndentedString(sortOrder)).append("\n");
    sb.append("    description: ").append(toIndentedString(description)).append("\n");
    sb.append("    maxStorage: ").append(toIndentedString(maxStorage)).append("\n");
    sb.append("    maxMemory: ").append(toIndentedString(maxMemory)).append("\n");
    sb.append("    maxCpu: ").append(toIndentedString(maxCpu)).append("\n");
    sb.append("    maxCores: ").append(toIndentedString(maxCores)).append("\n");
    sb.append("    maxDisks: ").append(toIndentedString(maxDisks)).append("\n");
    sb.append("    coresPerSocket: ").append(toIndentedString(coresPerSocket)).append("\n");
    sb.append("    customCpu: ").append(toIndentedString(customCpu)).append("\n");
    sb.append("    customCores: ").append(toIndentedString(customCores)).append("\n");
    sb.append("    customMaxStorage: ").append(toIndentedString(customMaxStorage)).append("\n");
    sb.append("    customMaxDataStorage: ").append(toIndentedString(customMaxDataStorage)).append("\n");
    sb.append("    customMaxMemory: ").append(toIndentedString(customMaxMemory)).append("\n");
    sb.append("    addVolumes: ").append(toIndentedString(addVolumes)).append("\n");
    sb.append("    memoryOptionSource: ").append(toIndentedString(memoryOptionSource)).append("\n");
    sb.append("    cpuOptionSource: ").append(toIndentedString(cpuOptionSource)).append("\n");
    sb.append("    dateCreated: ").append(toIndentedString(dateCreated)).append("\n");
    sb.append("    lastUpdated: ").append(toIndentedString(lastUpdated)).append("\n");
    sb.append("    regionCode: ").append(toIndentedString(regionCode)).append("\n");
    sb.append("    visibility: ").append(toIndentedString(visibility)).append("\n");
    sb.append("    editable: ").append(toIndentedString(editable)).append("\n");
    sb.append("    provisionType: ").append(toIndentedString(provisionType)).append("\n");
    sb.append("    tenants: ").append(toIndentedString(tenants)).append("\n");
    sb.append("    priceSets: ").append(toIndentedString(priceSets)).append("\n");
    sb.append("    config: ").append(toIndentedString(config)).append("\n");
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
