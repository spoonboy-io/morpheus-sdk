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
import org.openapitools.client.model.InstanceServicePlanAutoOptions;
import org.openapitools.client.model.ServerServicePlansDatastores;

/**
 * ServerServicePlans
 */
@javax.annotation.Generated(value = "org.openapitools.codegen.languages.JavaClientCodegen", date = "2023-08-21T11:26:47.028Z[GMT]")
public class ServerServicePlans {
  public static final String SERIALIZED_NAME_ID = "id";
  @SerializedName(SERIALIZED_NAME_ID)
  private Long id;

  public static final String SERIALIZED_NAME_NAME = "name";
  @SerializedName(SERIALIZED_NAME_NAME)
  private String name;

  public static final String SERIALIZED_NAME_VALUE = "value";
  @SerializedName(SERIALIZED_NAME_VALUE)
  private Long value;

  public static final String SERIALIZED_NAME_CODE = "code";
  @SerializedName(SERIALIZED_NAME_CODE)
  private String code;

  public static final String SERIALIZED_NAME_MAX_STORAGE = "maxStorage";
  @SerializedName(SERIALIZED_NAME_MAX_STORAGE)
  private Long maxStorage;

  public static final String SERIALIZED_NAME_MAX_MEMORY = "maxMemory";
  @SerializedName(SERIALIZED_NAME_MAX_MEMORY)
  private Long maxMemory;

  public static final String SERIALIZED_NAME_MAX_CPU = "maxCpu";
  @SerializedName(SERIALIZED_NAME_MAX_CPU)
  private Long maxCpu;

  public static final String SERIALIZED_NAME_MAX_CORES = "maxCores";
  @SerializedName(SERIALIZED_NAME_MAX_CORES)
  private Long maxCores;

  public static final String SERIALIZED_NAME_MAX_DATA_STORAGE = "maxDataStorage";
  @SerializedName(SERIALIZED_NAME_MAX_DATA_STORAGE)
  private Long maxDataStorage;

  public static final String SERIALIZED_NAME_CUSTOM_CPU = "customCpu";
  @SerializedName(SERIALIZED_NAME_CUSTOM_CPU)
  private Boolean customCpu;

  public static final String SERIALIZED_NAME_CUSTOM_MAX_MEMORY = "customMaxMemory";
  @SerializedName(SERIALIZED_NAME_CUSTOM_MAX_MEMORY)
  private Boolean customMaxMemory;

  public static final String SERIALIZED_NAME_CUSTOM_MAX_STORAGE = "customMaxStorage";
  @SerializedName(SERIALIZED_NAME_CUSTOM_MAX_STORAGE)
  private Boolean customMaxStorage;

  public static final String SERIALIZED_NAME_CUSTOM_MAX_DATA_STORAGE = "customMaxDataStorage";
  @SerializedName(SERIALIZED_NAME_CUSTOM_MAX_DATA_STORAGE)
  private Boolean customMaxDataStorage;

  public static final String SERIALIZED_NAME_CUSTOM_CORES_PER_SOCKET = "customCoresPerSocket";
  @SerializedName(SERIALIZED_NAME_CUSTOM_CORES_PER_SOCKET)
  private Boolean customCoresPerSocket;

  public static final String SERIALIZED_NAME_CORES_PER_SOCKET = "coresPerSocket";
  @SerializedName(SERIALIZED_NAME_CORES_PER_SOCKET)
  private Long coresPerSocket;

  public static final String SERIALIZED_NAME_STORAGE_TYPES = "storageTypes";
  @SerializedName(SERIALIZED_NAME_STORAGE_TYPES)
  private List<Object> storageTypes = null;

  public static final String SERIALIZED_NAME_ROOT_STORAGE_TYPES = "rootStorageTypes";
  @SerializedName(SERIALIZED_NAME_ROOT_STORAGE_TYPES)
  private List<Object> rootStorageTypes = null;

  public static final String SERIALIZED_NAME_ADD_VOLUMES = "addVolumes";
  @SerializedName(SERIALIZED_NAME_ADD_VOLUMES)
  private Boolean addVolumes;

  public static final String SERIALIZED_NAME_CUSTOMIZE_VOLUME = "customizeVolume";
  @SerializedName(SERIALIZED_NAME_CUSTOMIZE_VOLUME)
  private Boolean customizeVolume;

  public static final String SERIALIZED_NAME_ROOT_DISK_CUSTOMIZABLE = "rootDiskCustomizable";
  @SerializedName(SERIALIZED_NAME_ROOT_DISK_CUSTOMIZABLE)
  private Boolean rootDiskCustomizable;

  public static final String SERIALIZED_NAME_HOST_DISK_MODE = "hostDiskMode";
  @SerializedName(SERIALIZED_NAME_HOST_DISK_MODE)
  private String hostDiskMode;

  public static final String SERIALIZED_NAME_HAS_DATASTORE = "hasDatastore";
  @SerializedName(SERIALIZED_NAME_HAS_DATASTORE)
  private String hasDatastore;

  public static final String SERIALIZED_NAME_LVM_SUPPORTED = "lvmSupported";
  @SerializedName(SERIALIZED_NAME_LVM_SUPPORTED)
  private String lvmSupported;

  public static final String SERIALIZED_NAME_MIN_DISK = "minDisk";
  @SerializedName(SERIALIZED_NAME_MIN_DISK)
  private String minDisk;

  public static final String SERIALIZED_NAME_MAX_DISK = "maxDisk";
  @SerializedName(SERIALIZED_NAME_MAX_DISK)
  private String maxDisk;

  public static final String SERIALIZED_NAME_DATASTORES = "datastores";
  @SerializedName(SERIALIZED_NAME_DATASTORES)
  private ServerServicePlansDatastores datastores;

  public static final String SERIALIZED_NAME_SUPPORTS_AUTO_DATASTORE = "supportsAutoDatastore";
  @SerializedName(SERIALIZED_NAME_SUPPORTS_AUTO_DATASTORE)
  private Boolean supportsAutoDatastore;

  public static final String SERIALIZED_NAME_AUTO_OPTIONS = "autoOptions";
  @SerializedName(SERIALIZED_NAME_AUTO_OPTIONS)
  private List<InstanceServicePlanAutoOptions> autoOptions = null;

  public static final String SERIALIZED_NAME_CPU_OPTIONS = "cpuOptions";
  @SerializedName(SERIALIZED_NAME_CPU_OPTIONS)
  private List<Object> cpuOptions = null;

  public static final String SERIALIZED_NAME_MEMORY_OPTIONS = "memoryOptions";
  @SerializedName(SERIALIZED_NAME_MEMORY_OPTIONS)
  private List<Object> memoryOptions = null;

  public static final String SERIALIZED_NAME_ROOT_CUSTOM_SIZE_OPTIONS = "rootCustomSizeOptions";
  @SerializedName(SERIALIZED_NAME_ROOT_CUSTOM_SIZE_OPTIONS)
  private Object rootCustomSizeOptions;

  public static final String SERIALIZED_NAME_CUSTOM_SIZE_OPTIONS = "customSizeOptions";
  @SerializedName(SERIALIZED_NAME_CUSTOM_SIZE_OPTIONS)
  private Object customSizeOptions;

  public static final String SERIALIZED_NAME_CUSTOM_CORES = "customCores";
  @SerializedName(SERIALIZED_NAME_CUSTOM_CORES)
  private Boolean customCores;

  public static final String SERIALIZED_NAME_MAX_DISKS = "maxDisks";
  @SerializedName(SERIALIZED_NAME_MAX_DISKS)
  private String maxDisks;

  public static final String SERIALIZED_NAME_MEMORY_SIZE_TYPE = "memorySizeType";
  @SerializedName(SERIALIZED_NAME_MEMORY_SIZE_TYPE)
  private String memorySizeType;


  public ServerServicePlans id(Long id) {
    
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


  public ServerServicePlans name(String name) {
    
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


  public ServerServicePlans value(Long value) {
    
    this.value = value;
    return this;
  }

   /**
   * Get value
   * @return value
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public Long getValue() {
    return value;
  }


  public void setValue(Long value) {
    this.value = value;
  }


  public ServerServicePlans code(String code) {
    
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


  public ServerServicePlans maxStorage(Long maxStorage) {
    
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


  public ServerServicePlans maxMemory(Long maxMemory) {
    
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


  public ServerServicePlans maxCpu(Long maxCpu) {
    
    this.maxCpu = maxCpu;
    return this;
  }

   /**
   * Get maxCpu
   * @return maxCpu
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public Long getMaxCpu() {
    return maxCpu;
  }


  public void setMaxCpu(Long maxCpu) {
    this.maxCpu = maxCpu;
  }


  public ServerServicePlans maxCores(Long maxCores) {
    
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


  public ServerServicePlans maxDataStorage(Long maxDataStorage) {
    
    this.maxDataStorage = maxDataStorage;
    return this;
  }

   /**
   * Get maxDataStorage
   * @return maxDataStorage
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public Long getMaxDataStorage() {
    return maxDataStorage;
  }


  public void setMaxDataStorage(Long maxDataStorage) {
    this.maxDataStorage = maxDataStorage;
  }


  public ServerServicePlans customCpu(Boolean customCpu) {
    
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


  public ServerServicePlans customMaxMemory(Boolean customMaxMemory) {
    
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


  public ServerServicePlans customMaxStorage(Boolean customMaxStorage) {
    
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


  public ServerServicePlans customMaxDataStorage(Boolean customMaxDataStorage) {
    
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


  public ServerServicePlans customCoresPerSocket(Boolean customCoresPerSocket) {
    
    this.customCoresPerSocket = customCoresPerSocket;
    return this;
  }

   /**
   * Get customCoresPerSocket
   * @return customCoresPerSocket
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public Boolean getCustomCoresPerSocket() {
    return customCoresPerSocket;
  }


  public void setCustomCoresPerSocket(Boolean customCoresPerSocket) {
    this.customCoresPerSocket = customCoresPerSocket;
  }


  public ServerServicePlans coresPerSocket(Long coresPerSocket) {
    
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


  public ServerServicePlans storageTypes(List<Object> storageTypes) {
    
    this.storageTypes = storageTypes;
    return this;
  }

  public ServerServicePlans addStorageTypesItem(Object storageTypesItem) {
    if (this.storageTypes == null) {
      this.storageTypes = new ArrayList<Object>();
    }
    this.storageTypes.add(storageTypesItem);
    return this;
  }

   /**
   * Get storageTypes
   * @return storageTypes
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public List<Object> getStorageTypes() {
    return storageTypes;
  }


  public void setStorageTypes(List<Object> storageTypes) {
    this.storageTypes = storageTypes;
  }


  public ServerServicePlans rootStorageTypes(List<Object> rootStorageTypes) {
    
    this.rootStorageTypes = rootStorageTypes;
    return this;
  }

  public ServerServicePlans addRootStorageTypesItem(Object rootStorageTypesItem) {
    if (this.rootStorageTypes == null) {
      this.rootStorageTypes = new ArrayList<Object>();
    }
    this.rootStorageTypes.add(rootStorageTypesItem);
    return this;
  }

   /**
   * Get rootStorageTypes
   * @return rootStorageTypes
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public List<Object> getRootStorageTypes() {
    return rootStorageTypes;
  }


  public void setRootStorageTypes(List<Object> rootStorageTypes) {
    this.rootStorageTypes = rootStorageTypes;
  }


  public ServerServicePlans addVolumes(Boolean addVolumes) {
    
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


  public ServerServicePlans customizeVolume(Boolean customizeVolume) {
    
    this.customizeVolume = customizeVolume;
    return this;
  }

   /**
   * Get customizeVolume
   * @return customizeVolume
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public Boolean getCustomizeVolume() {
    return customizeVolume;
  }


  public void setCustomizeVolume(Boolean customizeVolume) {
    this.customizeVolume = customizeVolume;
  }


  public ServerServicePlans rootDiskCustomizable(Boolean rootDiskCustomizable) {
    
    this.rootDiskCustomizable = rootDiskCustomizable;
    return this;
  }

   /**
   * Get rootDiskCustomizable
   * @return rootDiskCustomizable
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public Boolean getRootDiskCustomizable() {
    return rootDiskCustomizable;
  }


  public void setRootDiskCustomizable(Boolean rootDiskCustomizable) {
    this.rootDiskCustomizable = rootDiskCustomizable;
  }


  public ServerServicePlans hostDiskMode(String hostDiskMode) {
    
    this.hostDiskMode = hostDiskMode;
    return this;
  }

   /**
   * Get hostDiskMode
   * @return hostDiskMode
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public String getHostDiskMode() {
    return hostDiskMode;
  }


  public void setHostDiskMode(String hostDiskMode) {
    this.hostDiskMode = hostDiskMode;
  }


  public ServerServicePlans hasDatastore(String hasDatastore) {
    
    this.hasDatastore = hasDatastore;
    return this;
  }

   /**
   * Get hasDatastore
   * @return hasDatastore
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public String getHasDatastore() {
    return hasDatastore;
  }


  public void setHasDatastore(String hasDatastore) {
    this.hasDatastore = hasDatastore;
  }


  public ServerServicePlans lvmSupported(String lvmSupported) {
    
    this.lvmSupported = lvmSupported;
    return this;
  }

   /**
   * Get lvmSupported
   * @return lvmSupported
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public String getLvmSupported() {
    return lvmSupported;
  }


  public void setLvmSupported(String lvmSupported) {
    this.lvmSupported = lvmSupported;
  }


  public ServerServicePlans minDisk(String minDisk) {
    
    this.minDisk = minDisk;
    return this;
  }

   /**
   * Get minDisk
   * @return minDisk
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public String getMinDisk() {
    return minDisk;
  }


  public void setMinDisk(String minDisk) {
    this.minDisk = minDisk;
  }


  public ServerServicePlans maxDisk(String maxDisk) {
    
    this.maxDisk = maxDisk;
    return this;
  }

   /**
   * Get maxDisk
   * @return maxDisk
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public String getMaxDisk() {
    return maxDisk;
  }


  public void setMaxDisk(String maxDisk) {
    this.maxDisk = maxDisk;
  }


  public ServerServicePlans datastores(ServerServicePlansDatastores datastores) {
    
    this.datastores = datastores;
    return this;
  }

   /**
   * Get datastores
   * @return datastores
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public ServerServicePlansDatastores getDatastores() {
    return datastores;
  }


  public void setDatastores(ServerServicePlansDatastores datastores) {
    this.datastores = datastores;
  }


  public ServerServicePlans supportsAutoDatastore(Boolean supportsAutoDatastore) {
    
    this.supportsAutoDatastore = supportsAutoDatastore;
    return this;
  }

   /**
   * Get supportsAutoDatastore
   * @return supportsAutoDatastore
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public Boolean getSupportsAutoDatastore() {
    return supportsAutoDatastore;
  }


  public void setSupportsAutoDatastore(Boolean supportsAutoDatastore) {
    this.supportsAutoDatastore = supportsAutoDatastore;
  }


  public ServerServicePlans autoOptions(List<InstanceServicePlanAutoOptions> autoOptions) {
    
    this.autoOptions = autoOptions;
    return this;
  }

  public ServerServicePlans addAutoOptionsItem(InstanceServicePlanAutoOptions autoOptionsItem) {
    if (this.autoOptions == null) {
      this.autoOptions = new ArrayList<InstanceServicePlanAutoOptions>();
    }
    this.autoOptions.add(autoOptionsItem);
    return this;
  }

   /**
   * Get autoOptions
   * @return autoOptions
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public List<InstanceServicePlanAutoOptions> getAutoOptions() {
    return autoOptions;
  }


  public void setAutoOptions(List<InstanceServicePlanAutoOptions> autoOptions) {
    this.autoOptions = autoOptions;
  }


  public ServerServicePlans cpuOptions(List<Object> cpuOptions) {
    
    this.cpuOptions = cpuOptions;
    return this;
  }

  public ServerServicePlans addCpuOptionsItem(Object cpuOptionsItem) {
    if (this.cpuOptions == null) {
      this.cpuOptions = new ArrayList<Object>();
    }
    this.cpuOptions.add(cpuOptionsItem);
    return this;
  }

   /**
   * Get cpuOptions
   * @return cpuOptions
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public List<Object> getCpuOptions() {
    return cpuOptions;
  }


  public void setCpuOptions(List<Object> cpuOptions) {
    this.cpuOptions = cpuOptions;
  }


  public ServerServicePlans memoryOptions(List<Object> memoryOptions) {
    
    this.memoryOptions = memoryOptions;
    return this;
  }

  public ServerServicePlans addMemoryOptionsItem(Object memoryOptionsItem) {
    if (this.memoryOptions == null) {
      this.memoryOptions = new ArrayList<Object>();
    }
    this.memoryOptions.add(memoryOptionsItem);
    return this;
  }

   /**
   * Get memoryOptions
   * @return memoryOptions
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public List<Object> getMemoryOptions() {
    return memoryOptions;
  }


  public void setMemoryOptions(List<Object> memoryOptions) {
    this.memoryOptions = memoryOptions;
  }


  public ServerServicePlans rootCustomSizeOptions(Object rootCustomSizeOptions) {
    
    this.rootCustomSizeOptions = rootCustomSizeOptions;
    return this;
  }

   /**
   * Get rootCustomSizeOptions
   * @return rootCustomSizeOptions
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public Object getRootCustomSizeOptions() {
    return rootCustomSizeOptions;
  }


  public void setRootCustomSizeOptions(Object rootCustomSizeOptions) {
    this.rootCustomSizeOptions = rootCustomSizeOptions;
  }


  public ServerServicePlans customSizeOptions(Object customSizeOptions) {
    
    this.customSizeOptions = customSizeOptions;
    return this;
  }

   /**
   * Get customSizeOptions
   * @return customSizeOptions
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public Object getCustomSizeOptions() {
    return customSizeOptions;
  }


  public void setCustomSizeOptions(Object customSizeOptions) {
    this.customSizeOptions = customSizeOptions;
  }


  public ServerServicePlans customCores(Boolean customCores) {
    
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


  public ServerServicePlans maxDisks(String maxDisks) {
    
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


  public ServerServicePlans memorySizeType(String memorySizeType) {
    
    this.memorySizeType = memorySizeType;
    return this;
  }

   /**
   * Get memorySizeType
   * @return memorySizeType
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public String getMemorySizeType() {
    return memorySizeType;
  }


  public void setMemorySizeType(String memorySizeType) {
    this.memorySizeType = memorySizeType;
  }


  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (o == null || getClass() != o.getClass()) {
      return false;
    }
    ServerServicePlans serverServicePlans = (ServerServicePlans) o;
    return Objects.equals(this.id, serverServicePlans.id) &&
        Objects.equals(this.name, serverServicePlans.name) &&
        Objects.equals(this.value, serverServicePlans.value) &&
        Objects.equals(this.code, serverServicePlans.code) &&
        Objects.equals(this.maxStorage, serverServicePlans.maxStorage) &&
        Objects.equals(this.maxMemory, serverServicePlans.maxMemory) &&
        Objects.equals(this.maxCpu, serverServicePlans.maxCpu) &&
        Objects.equals(this.maxCores, serverServicePlans.maxCores) &&
        Objects.equals(this.maxDataStorage, serverServicePlans.maxDataStorage) &&
        Objects.equals(this.customCpu, serverServicePlans.customCpu) &&
        Objects.equals(this.customMaxMemory, serverServicePlans.customMaxMemory) &&
        Objects.equals(this.customMaxStorage, serverServicePlans.customMaxStorage) &&
        Objects.equals(this.customMaxDataStorage, serverServicePlans.customMaxDataStorage) &&
        Objects.equals(this.customCoresPerSocket, serverServicePlans.customCoresPerSocket) &&
        Objects.equals(this.coresPerSocket, serverServicePlans.coresPerSocket) &&
        Objects.equals(this.storageTypes, serverServicePlans.storageTypes) &&
        Objects.equals(this.rootStorageTypes, serverServicePlans.rootStorageTypes) &&
        Objects.equals(this.addVolumes, serverServicePlans.addVolumes) &&
        Objects.equals(this.customizeVolume, serverServicePlans.customizeVolume) &&
        Objects.equals(this.rootDiskCustomizable, serverServicePlans.rootDiskCustomizable) &&
        Objects.equals(this.hostDiskMode, serverServicePlans.hostDiskMode) &&
        Objects.equals(this.hasDatastore, serverServicePlans.hasDatastore) &&
        Objects.equals(this.lvmSupported, serverServicePlans.lvmSupported) &&
        Objects.equals(this.minDisk, serverServicePlans.minDisk) &&
        Objects.equals(this.maxDisk, serverServicePlans.maxDisk) &&
        Objects.equals(this.datastores, serverServicePlans.datastores) &&
        Objects.equals(this.supportsAutoDatastore, serverServicePlans.supportsAutoDatastore) &&
        Objects.equals(this.autoOptions, serverServicePlans.autoOptions) &&
        Objects.equals(this.cpuOptions, serverServicePlans.cpuOptions) &&
        Objects.equals(this.memoryOptions, serverServicePlans.memoryOptions) &&
        Objects.equals(this.rootCustomSizeOptions, serverServicePlans.rootCustomSizeOptions) &&
        Objects.equals(this.customSizeOptions, serverServicePlans.customSizeOptions) &&
        Objects.equals(this.customCores, serverServicePlans.customCores) &&
        Objects.equals(this.maxDisks, serverServicePlans.maxDisks) &&
        Objects.equals(this.memorySizeType, serverServicePlans.memorySizeType);
  }

  @Override
  public int hashCode() {
    return Objects.hash(id, name, value, code, maxStorage, maxMemory, maxCpu, maxCores, maxDataStorage, customCpu, customMaxMemory, customMaxStorage, customMaxDataStorage, customCoresPerSocket, coresPerSocket, storageTypes, rootStorageTypes, addVolumes, customizeVolume, rootDiskCustomizable, hostDiskMode, hasDatastore, lvmSupported, minDisk, maxDisk, datastores, supportsAutoDatastore, autoOptions, cpuOptions, memoryOptions, rootCustomSizeOptions, customSizeOptions, customCores, maxDisks, memorySizeType);
  }


  @Override
  public String toString() {
    StringBuilder sb = new StringBuilder();
    sb.append("class ServerServicePlans {\n");
    sb.append("    id: ").append(toIndentedString(id)).append("\n");
    sb.append("    name: ").append(toIndentedString(name)).append("\n");
    sb.append("    value: ").append(toIndentedString(value)).append("\n");
    sb.append("    code: ").append(toIndentedString(code)).append("\n");
    sb.append("    maxStorage: ").append(toIndentedString(maxStorage)).append("\n");
    sb.append("    maxMemory: ").append(toIndentedString(maxMemory)).append("\n");
    sb.append("    maxCpu: ").append(toIndentedString(maxCpu)).append("\n");
    sb.append("    maxCores: ").append(toIndentedString(maxCores)).append("\n");
    sb.append("    maxDataStorage: ").append(toIndentedString(maxDataStorage)).append("\n");
    sb.append("    customCpu: ").append(toIndentedString(customCpu)).append("\n");
    sb.append("    customMaxMemory: ").append(toIndentedString(customMaxMemory)).append("\n");
    sb.append("    customMaxStorage: ").append(toIndentedString(customMaxStorage)).append("\n");
    sb.append("    customMaxDataStorage: ").append(toIndentedString(customMaxDataStorage)).append("\n");
    sb.append("    customCoresPerSocket: ").append(toIndentedString(customCoresPerSocket)).append("\n");
    sb.append("    coresPerSocket: ").append(toIndentedString(coresPerSocket)).append("\n");
    sb.append("    storageTypes: ").append(toIndentedString(storageTypes)).append("\n");
    sb.append("    rootStorageTypes: ").append(toIndentedString(rootStorageTypes)).append("\n");
    sb.append("    addVolumes: ").append(toIndentedString(addVolumes)).append("\n");
    sb.append("    customizeVolume: ").append(toIndentedString(customizeVolume)).append("\n");
    sb.append("    rootDiskCustomizable: ").append(toIndentedString(rootDiskCustomizable)).append("\n");
    sb.append("    hostDiskMode: ").append(toIndentedString(hostDiskMode)).append("\n");
    sb.append("    hasDatastore: ").append(toIndentedString(hasDatastore)).append("\n");
    sb.append("    lvmSupported: ").append(toIndentedString(lvmSupported)).append("\n");
    sb.append("    minDisk: ").append(toIndentedString(minDisk)).append("\n");
    sb.append("    maxDisk: ").append(toIndentedString(maxDisk)).append("\n");
    sb.append("    datastores: ").append(toIndentedString(datastores)).append("\n");
    sb.append("    supportsAutoDatastore: ").append(toIndentedString(supportsAutoDatastore)).append("\n");
    sb.append("    autoOptions: ").append(toIndentedString(autoOptions)).append("\n");
    sb.append("    cpuOptions: ").append(toIndentedString(cpuOptions)).append("\n");
    sb.append("    memoryOptions: ").append(toIndentedString(memoryOptions)).append("\n");
    sb.append("    rootCustomSizeOptions: ").append(toIndentedString(rootCustomSizeOptions)).append("\n");
    sb.append("    customSizeOptions: ").append(toIndentedString(customSizeOptions)).append("\n");
    sb.append("    customCores: ").append(toIndentedString(customCores)).append("\n");
    sb.append("    maxDisks: ").append(toIndentedString(maxDisks)).append("\n");
    sb.append("    memorySizeType: ").append(toIndentedString(memorySizeType)).append("\n");
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
