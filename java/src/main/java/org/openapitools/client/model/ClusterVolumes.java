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
import org.openapitools.client.model.ApiBlueprintsIdUpdatePermissionsResourcePermissionSites;

/**
 * ClusterVolumes
 */
@javax.annotation.Generated(value = "org.openapitools.codegen.languages.JavaClientCodegen", date = "2023-08-21T11:26:47.028Z[GMT]")
public class ClusterVolumes {
  public static final String SERIALIZED_NAME_ID = "id";
  @SerializedName(SERIALIZED_NAME_ID)
  private Long id;

  public static final String SERIALIZED_NAME_INTERNAL_ID = "internalId";
  @SerializedName(SERIALIZED_NAME_INTERNAL_ID)
  private String internalId;

  public static final String SERIALIZED_NAME_DISPLAY_ORDER = "displayOrder";
  @SerializedName(SERIALIZED_NAME_DISPLAY_ORDER)
  private Long displayOrder;

  public static final String SERIALIZED_NAME_ACTIVE = "active";
  @SerializedName(SERIALIZED_NAME_ACTIVE)
  private Boolean active;

  public static final String SERIALIZED_NAME_USED_STORAGE = "usedStorage";
  @SerializedName(SERIALIZED_NAME_USED_STORAGE)
  private Long usedStorage;

  public static final String SERIALIZED_NAME_PROVISION_TYPE = "provisionType";
  @SerializedName(SERIALIZED_NAME_PROVISION_TYPE)
  private String provisionType;

  public static final String SERIALIZED_NAME_RESIZEABLE = "resizeable";
  @SerializedName(SERIALIZED_NAME_RESIZEABLE)
  private Boolean resizeable;

  public static final String SERIALIZED_NAME_ONLINE = "online";
  @SerializedName(SERIALIZED_NAME_ONLINE)
  private Boolean online;

  public static final String SERIALIZED_NAME_DEVICE_DISPLAY_NAME = "deviceDisplayName";
  @SerializedName(SERIALIZED_NAME_DEVICE_DISPLAY_NAME)
  private String deviceDisplayName;

  public static final String SERIALIZED_NAME_REF_TYPE = "refType";
  @SerializedName(SERIALIZED_NAME_REF_TYPE)
  private String refType;

  public static final String SERIALIZED_NAME_NAME = "name";
  @SerializedName(SERIALIZED_NAME_NAME)
  private String name;

  public static final String SERIALIZED_NAME_EXTERNAL_ID = "externalId";
  @SerializedName(SERIALIZED_NAME_EXTERNAL_ID)
  private String externalId;

  public static final String SERIALIZED_NAME_DATASTORE_OPTION = "datastoreOption";
  @SerializedName(SERIALIZED_NAME_DATASTORE_OPTION)
  private String datastoreOption;

  public static final String SERIALIZED_NAME_CLAIM_NAME = "claimName";
  @SerializedName(SERIALIZED_NAME_CLAIM_NAME)
  private String claimName;

  public static final String SERIALIZED_NAME_VOLUME_TYPE = "volumeType";
  @SerializedName(SERIALIZED_NAME_VOLUME_TYPE)
  private String volumeType;

  public static final String SERIALIZED_NAME_DEVICE_NAME = "deviceName";
  @SerializedName(SERIALIZED_NAME_DEVICE_NAME)
  private String deviceName;

  public static final String SERIALIZED_NAME_REMOVABLE = "removable";
  @SerializedName(SERIALIZED_NAME_REMOVABLE)
  private Boolean removable;

  public static final String SERIALIZED_NAME_POOL_NAME = "poolName";
  @SerializedName(SERIALIZED_NAME_POOL_NAME)
  private String poolName;

  public static final String SERIALIZED_NAME_READ_ONLY = "readOnly";
  @SerializedName(SERIALIZED_NAME_READ_ONLY)
  private Boolean readOnly;

  public static final String SERIALIZED_NAME_SOURCE_ID = "sourceId";
  @SerializedName(SERIALIZED_NAME_SOURCE_ID)
  private String sourceId;

  public static final String SERIALIZED_NAME_ZONE_ID = "zoneId";
  @SerializedName(SERIALIZED_NAME_ZONE_ID)
  private Long zoneId;

  public static final String SERIALIZED_NAME_ROOT_VOLUME = "rootVolume";
  @SerializedName(SERIALIZED_NAME_ROOT_VOLUME)
  private Boolean rootVolume;

  public static final String SERIALIZED_NAME_REF_ID = "refId";
  @SerializedName(SERIALIZED_NAME_REF_ID)
  private Long refId;

  public static final String SERIALIZED_NAME_CATEGORY = "category";
  @SerializedName(SERIALIZED_NAME_CATEGORY)
  private String category;

  public static final String SERIALIZED_NAME_STATUS = "status";
  @SerializedName(SERIALIZED_NAME_STATUS)
  private String status;

  public static final String SERIALIZED_NAME_RAW_DATA = "rawData";
  @SerializedName(SERIALIZED_NAME_RAW_DATA)
  private String rawData;

  public static final String SERIALIZED_NAME_MAX_STORAGE = "maxStorage";
  @SerializedName(SERIALIZED_NAME_MAX_STORAGE)
  private Long maxStorage;

  public static final String SERIALIZED_NAME_ACCOUNT = "account";
  @SerializedName(SERIALIZED_NAME_ACCOUNT)
  private ApiBlueprintsIdUpdatePermissionsResourcePermissionSites account;

  public static final String SERIALIZED_NAME_TYPE = "type";
  @SerializedName(SERIALIZED_NAME_TYPE)
  private ApiBlueprintsIdUpdatePermissionsResourcePermissionSites type;

  public static final String SERIALIZED_NAME_DATASTORE = "datastore";
  @SerializedName(SERIALIZED_NAME_DATASTORE)
  private ApiBlueprintsIdUpdatePermissionsResourcePermissionSites datastore;


  public ClusterVolumes id(Long id) {
    
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


  public ClusterVolumes internalId(String internalId) {
    
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


  public ClusterVolumes displayOrder(Long displayOrder) {
    
    this.displayOrder = displayOrder;
    return this;
  }

   /**
   * Get displayOrder
   * @return displayOrder
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public Long getDisplayOrder() {
    return displayOrder;
  }


  public void setDisplayOrder(Long displayOrder) {
    this.displayOrder = displayOrder;
  }


  public ClusterVolumes active(Boolean active) {
    
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


  public ClusterVolumes usedStorage(Long usedStorage) {
    
    this.usedStorage = usedStorage;
    return this;
  }

   /**
   * Get usedStorage
   * @return usedStorage
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public Long getUsedStorage() {
    return usedStorage;
  }


  public void setUsedStorage(Long usedStorage) {
    this.usedStorage = usedStorage;
  }


  public ClusterVolumes provisionType(String provisionType) {
    
    this.provisionType = provisionType;
    return this;
  }

   /**
   * Get provisionType
   * @return provisionType
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public String getProvisionType() {
    return provisionType;
  }


  public void setProvisionType(String provisionType) {
    this.provisionType = provisionType;
  }


  public ClusterVolumes resizeable(Boolean resizeable) {
    
    this.resizeable = resizeable;
    return this;
  }

   /**
   * Get resizeable
   * @return resizeable
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public Boolean getResizeable() {
    return resizeable;
  }


  public void setResizeable(Boolean resizeable) {
    this.resizeable = resizeable;
  }


  public ClusterVolumes online(Boolean online) {
    
    this.online = online;
    return this;
  }

   /**
   * Get online
   * @return online
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public Boolean getOnline() {
    return online;
  }


  public void setOnline(Boolean online) {
    this.online = online;
  }


  public ClusterVolumes deviceDisplayName(String deviceDisplayName) {
    
    this.deviceDisplayName = deviceDisplayName;
    return this;
  }

   /**
   * Get deviceDisplayName
   * @return deviceDisplayName
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public String getDeviceDisplayName() {
    return deviceDisplayName;
  }


  public void setDeviceDisplayName(String deviceDisplayName) {
    this.deviceDisplayName = deviceDisplayName;
  }


  public ClusterVolumes refType(String refType) {
    
    this.refType = refType;
    return this;
  }

   /**
   * Get refType
   * @return refType
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public String getRefType() {
    return refType;
  }


  public void setRefType(String refType) {
    this.refType = refType;
  }


  public ClusterVolumes name(String name) {
    
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


  public ClusterVolumes externalId(String externalId) {
    
    this.externalId = externalId;
    return this;
  }

   /**
   * Get externalId
   * @return externalId
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public String getExternalId() {
    return externalId;
  }


  public void setExternalId(String externalId) {
    this.externalId = externalId;
  }


  public ClusterVolumes datastoreOption(String datastoreOption) {
    
    this.datastoreOption = datastoreOption;
    return this;
  }

   /**
   * Get datastoreOption
   * @return datastoreOption
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public String getDatastoreOption() {
    return datastoreOption;
  }


  public void setDatastoreOption(String datastoreOption) {
    this.datastoreOption = datastoreOption;
  }


  public ClusterVolumes claimName(String claimName) {
    
    this.claimName = claimName;
    return this;
  }

   /**
   * Get claimName
   * @return claimName
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public String getClaimName() {
    return claimName;
  }


  public void setClaimName(String claimName) {
    this.claimName = claimName;
  }


  public ClusterVolumes volumeType(String volumeType) {
    
    this.volumeType = volumeType;
    return this;
  }

   /**
   * Get volumeType
   * @return volumeType
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public String getVolumeType() {
    return volumeType;
  }


  public void setVolumeType(String volumeType) {
    this.volumeType = volumeType;
  }


  public ClusterVolumes deviceName(String deviceName) {
    
    this.deviceName = deviceName;
    return this;
  }

   /**
   * Get deviceName
   * @return deviceName
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public String getDeviceName() {
    return deviceName;
  }


  public void setDeviceName(String deviceName) {
    this.deviceName = deviceName;
  }


  public ClusterVolumes removable(Boolean removable) {
    
    this.removable = removable;
    return this;
  }

   /**
   * Get removable
   * @return removable
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public Boolean getRemovable() {
    return removable;
  }


  public void setRemovable(Boolean removable) {
    this.removable = removable;
  }


  public ClusterVolumes poolName(String poolName) {
    
    this.poolName = poolName;
    return this;
  }

   /**
   * Get poolName
   * @return poolName
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public String getPoolName() {
    return poolName;
  }


  public void setPoolName(String poolName) {
    this.poolName = poolName;
  }


  public ClusterVolumes readOnly(Boolean readOnly) {
    
    this.readOnly = readOnly;
    return this;
  }

   /**
   * Get readOnly
   * @return readOnly
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public Boolean getReadOnly() {
    return readOnly;
  }


  public void setReadOnly(Boolean readOnly) {
    this.readOnly = readOnly;
  }


  public ClusterVolumes sourceId(String sourceId) {
    
    this.sourceId = sourceId;
    return this;
  }

   /**
   * Get sourceId
   * @return sourceId
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public String getSourceId() {
    return sourceId;
  }


  public void setSourceId(String sourceId) {
    this.sourceId = sourceId;
  }


  public ClusterVolumes zoneId(Long zoneId) {
    
    this.zoneId = zoneId;
    return this;
  }

   /**
   * Get zoneId
   * @return zoneId
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public Long getZoneId() {
    return zoneId;
  }


  public void setZoneId(Long zoneId) {
    this.zoneId = zoneId;
  }


  public ClusterVolumes rootVolume(Boolean rootVolume) {
    
    this.rootVolume = rootVolume;
    return this;
  }

   /**
   * Get rootVolume
   * @return rootVolume
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public Boolean getRootVolume() {
    return rootVolume;
  }


  public void setRootVolume(Boolean rootVolume) {
    this.rootVolume = rootVolume;
  }


  public ClusterVolumes refId(Long refId) {
    
    this.refId = refId;
    return this;
  }

   /**
   * Get refId
   * @return refId
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public Long getRefId() {
    return refId;
  }


  public void setRefId(Long refId) {
    this.refId = refId;
  }


  public ClusterVolumes category(String category) {
    
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


  public ClusterVolumes status(String status) {
    
    this.status = status;
    return this;
  }

   /**
   * Get status
   * @return status
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public String getStatus() {
    return status;
  }


  public void setStatus(String status) {
    this.status = status;
  }


  public ClusterVolumes rawData(String rawData) {
    
    this.rawData = rawData;
    return this;
  }

   /**
   * Get rawData
   * @return rawData
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public String getRawData() {
    return rawData;
  }


  public void setRawData(String rawData) {
    this.rawData = rawData;
  }


  public ClusterVolumes maxStorage(Long maxStorage) {
    
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


  public ClusterVolumes account(ApiBlueprintsIdUpdatePermissionsResourcePermissionSites account) {
    
    this.account = account;
    return this;
  }

   /**
   * Get account
   * @return account
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public ApiBlueprintsIdUpdatePermissionsResourcePermissionSites getAccount() {
    return account;
  }


  public void setAccount(ApiBlueprintsIdUpdatePermissionsResourcePermissionSites account) {
    this.account = account;
  }


  public ClusterVolumes type(ApiBlueprintsIdUpdatePermissionsResourcePermissionSites type) {
    
    this.type = type;
    return this;
  }

   /**
   * Get type
   * @return type
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public ApiBlueprintsIdUpdatePermissionsResourcePermissionSites getType() {
    return type;
  }


  public void setType(ApiBlueprintsIdUpdatePermissionsResourcePermissionSites type) {
    this.type = type;
  }


  public ClusterVolumes datastore(ApiBlueprintsIdUpdatePermissionsResourcePermissionSites datastore) {
    
    this.datastore = datastore;
    return this;
  }

   /**
   * Get datastore
   * @return datastore
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public ApiBlueprintsIdUpdatePermissionsResourcePermissionSites getDatastore() {
    return datastore;
  }


  public void setDatastore(ApiBlueprintsIdUpdatePermissionsResourcePermissionSites datastore) {
    this.datastore = datastore;
  }


  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (o == null || getClass() != o.getClass()) {
      return false;
    }
    ClusterVolumes clusterVolumes = (ClusterVolumes) o;
    return Objects.equals(this.id, clusterVolumes.id) &&
        Objects.equals(this.internalId, clusterVolumes.internalId) &&
        Objects.equals(this.displayOrder, clusterVolumes.displayOrder) &&
        Objects.equals(this.active, clusterVolumes.active) &&
        Objects.equals(this.usedStorage, clusterVolumes.usedStorage) &&
        Objects.equals(this.provisionType, clusterVolumes.provisionType) &&
        Objects.equals(this.resizeable, clusterVolumes.resizeable) &&
        Objects.equals(this.online, clusterVolumes.online) &&
        Objects.equals(this.deviceDisplayName, clusterVolumes.deviceDisplayName) &&
        Objects.equals(this.refType, clusterVolumes.refType) &&
        Objects.equals(this.name, clusterVolumes.name) &&
        Objects.equals(this.externalId, clusterVolumes.externalId) &&
        Objects.equals(this.datastoreOption, clusterVolumes.datastoreOption) &&
        Objects.equals(this.claimName, clusterVolumes.claimName) &&
        Objects.equals(this.volumeType, clusterVolumes.volumeType) &&
        Objects.equals(this.deviceName, clusterVolumes.deviceName) &&
        Objects.equals(this.removable, clusterVolumes.removable) &&
        Objects.equals(this.poolName, clusterVolumes.poolName) &&
        Objects.equals(this.readOnly, clusterVolumes.readOnly) &&
        Objects.equals(this.sourceId, clusterVolumes.sourceId) &&
        Objects.equals(this.zoneId, clusterVolumes.zoneId) &&
        Objects.equals(this.rootVolume, clusterVolumes.rootVolume) &&
        Objects.equals(this.refId, clusterVolumes.refId) &&
        Objects.equals(this.category, clusterVolumes.category) &&
        Objects.equals(this.status, clusterVolumes.status) &&
        Objects.equals(this.rawData, clusterVolumes.rawData) &&
        Objects.equals(this.maxStorage, clusterVolumes.maxStorage) &&
        Objects.equals(this.account, clusterVolumes.account) &&
        Objects.equals(this.type, clusterVolumes.type) &&
        Objects.equals(this.datastore, clusterVolumes.datastore);
  }

  @Override
  public int hashCode() {
    return Objects.hash(id, internalId, displayOrder, active, usedStorage, provisionType, resizeable, online, deviceDisplayName, refType, name, externalId, datastoreOption, claimName, volumeType, deviceName, removable, poolName, readOnly, sourceId, zoneId, rootVolume, refId, category, status, rawData, maxStorage, account, type, datastore);
  }


  @Override
  public String toString() {
    StringBuilder sb = new StringBuilder();
    sb.append("class ClusterVolumes {\n");
    sb.append("    id: ").append(toIndentedString(id)).append("\n");
    sb.append("    internalId: ").append(toIndentedString(internalId)).append("\n");
    sb.append("    displayOrder: ").append(toIndentedString(displayOrder)).append("\n");
    sb.append("    active: ").append(toIndentedString(active)).append("\n");
    sb.append("    usedStorage: ").append(toIndentedString(usedStorage)).append("\n");
    sb.append("    provisionType: ").append(toIndentedString(provisionType)).append("\n");
    sb.append("    resizeable: ").append(toIndentedString(resizeable)).append("\n");
    sb.append("    online: ").append(toIndentedString(online)).append("\n");
    sb.append("    deviceDisplayName: ").append(toIndentedString(deviceDisplayName)).append("\n");
    sb.append("    refType: ").append(toIndentedString(refType)).append("\n");
    sb.append("    name: ").append(toIndentedString(name)).append("\n");
    sb.append("    externalId: ").append(toIndentedString(externalId)).append("\n");
    sb.append("    datastoreOption: ").append(toIndentedString(datastoreOption)).append("\n");
    sb.append("    claimName: ").append(toIndentedString(claimName)).append("\n");
    sb.append("    volumeType: ").append(toIndentedString(volumeType)).append("\n");
    sb.append("    deviceName: ").append(toIndentedString(deviceName)).append("\n");
    sb.append("    removable: ").append(toIndentedString(removable)).append("\n");
    sb.append("    poolName: ").append(toIndentedString(poolName)).append("\n");
    sb.append("    readOnly: ").append(toIndentedString(readOnly)).append("\n");
    sb.append("    sourceId: ").append(toIndentedString(sourceId)).append("\n");
    sb.append("    zoneId: ").append(toIndentedString(zoneId)).append("\n");
    sb.append("    rootVolume: ").append(toIndentedString(rootVolume)).append("\n");
    sb.append("    refId: ").append(toIndentedString(refId)).append("\n");
    sb.append("    category: ").append(toIndentedString(category)).append("\n");
    sb.append("    status: ").append(toIndentedString(status)).append("\n");
    sb.append("    rawData: ").append(toIndentedString(rawData)).append("\n");
    sb.append("    maxStorage: ").append(toIndentedString(maxStorage)).append("\n");
    sb.append("    account: ").append(toIndentedString(account)).append("\n");
    sb.append("    type: ").append(toIndentedString(type)).append("\n");
    sb.append("    datastore: ").append(toIndentedString(datastore)).append("\n");
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

