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
 * ClusterServerCreateVolumes
 */
@javax.annotation.Generated(value = "org.openapitools.codegen.languages.JavaClientCodegen", date = "2023-08-21T11:26:47.028Z[GMT]")
public class ClusterServerCreateVolumes {
  public static final String SERIALIZED_NAME_ID = "id";
  @SerializedName(SERIALIZED_NAME_ID)
  private Long id = -1l;

  public static final String SERIALIZED_NAME_ROOT_VOLUME = "rootVolume";
  @SerializedName(SERIALIZED_NAME_ROOT_VOLUME)
  private Boolean rootVolume = true;

  public static final String SERIALIZED_NAME_NAME = "name";
  @SerializedName(SERIALIZED_NAME_NAME)
  private String name = "root";

  public static final String SERIALIZED_NAME_SIZE = "size";
  @SerializedName(SERIALIZED_NAME_SIZE)
  private Long size;

  public static final String SERIALIZED_NAME_SIZE_ID = "sizeId";
  @SerializedName(SERIALIZED_NAME_SIZE_ID)
  private String sizeId;

  public static final String SERIALIZED_NAME_STORAGE_TYPE = "storageType";
  @SerializedName(SERIALIZED_NAME_STORAGE_TYPE)
  private Long storageType;

  public static final String SERIALIZED_NAME_DATASTORE_ID = "datastoreId";
  @SerializedName(SERIALIZED_NAME_DATASTORE_ID)
  private String datastoreId;


  public ClusterServerCreateVolumes id(Long id) {
    
    this.id = id;
    return this;
  }

   /**
   * The id for the LV configuration being created
   * @return id
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "The id for the LV configuration being created")

  public Long getId() {
    return id;
  }


  public void setId(Long id) {
    this.id = id;
  }


  public ClusterServerCreateVolumes rootVolume(Boolean rootVolume) {
    
    this.rootVolume = rootVolume;
    return this;
  }

   /**
   * If set to false then a non-root LV will be created
   * @return rootVolume
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "If set to false then a non-root LV will be created")

  public Boolean getRootVolume() {
    return rootVolume;
  }


  public void setRootVolume(Boolean rootVolume) {
    this.rootVolume = rootVolume;
  }


  public ClusterServerCreateVolumes name(String name) {
    
    this.name = name;
    return this;
  }

   /**
   * Name/type of the LV being created
   * @return name
  **/
  @ApiModelProperty(required = true, value = "Name/type of the LV being created")

  public String getName() {
    return name;
  }


  public void setName(String name) {
    this.name = name;
  }


  public ClusterServerCreateVolumes size(Long size) {
    
    this.size = size;
    return this;
  }

   /**
   * Size of the LV to be created in GBs  Default is from the service plan 
   * @return size
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "Size of the LV to be created in GBs  Default is from the service plan ")

  public Long getSize() {
    return size;
  }


  public void setSize(Long size) {
    this.size = size;
  }


  public ClusterServerCreateVolumes sizeId(String sizeId) {
    
    this.sizeId = sizeId;
    return this;
  }

   /**
   * Can be used to select pre-existing LV choices from Morpheus
   * @return sizeId
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "Can be used to select pre-existing LV choices from Morpheus")

  public String getSizeId() {
    return sizeId;
  }


  public void setSizeId(String sizeId) {
    this.sizeId = sizeId;
  }


  public ClusterServerCreateVolumes storageType(Long storageType) {
    
    this.storageType = storageType;
    return this;
  }

   /**
   * Identifier for LV type
   * @return storageType
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "Identifier for LV type")

  public Long getStorageType() {
    return storageType;
  }


  public void setStorageType(Long storageType) {
    this.storageType = storageType;
  }


  public ClusterServerCreateVolumes datastoreId(String datastoreId) {
    
    this.datastoreId = datastoreId;
    return this;
  }

   /**
   * The ID of the specific datastore. Auto selection can be specified as auto or &#x60;autoCluster&#x60; (for clusters).
   * @return datastoreId
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(required = true, value = "The ID of the specific datastore. Auto selection can be specified as auto or `autoCluster` (for clusters).")

  public String getDatastoreId() {
    return datastoreId;
  }


  public void setDatastoreId(String datastoreId) {
    this.datastoreId = datastoreId;
  }


  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (o == null || getClass() != o.getClass()) {
      return false;
    }
    ClusterServerCreateVolumes clusterServerCreateVolumes = (ClusterServerCreateVolumes) o;
    return Objects.equals(this.id, clusterServerCreateVolumes.id) &&
        Objects.equals(this.rootVolume, clusterServerCreateVolumes.rootVolume) &&
        Objects.equals(this.name, clusterServerCreateVolumes.name) &&
        Objects.equals(this.size, clusterServerCreateVolumes.size) &&
        Objects.equals(this.sizeId, clusterServerCreateVolumes.sizeId) &&
        Objects.equals(this.storageType, clusterServerCreateVolumes.storageType) &&
        Objects.equals(this.datastoreId, clusterServerCreateVolumes.datastoreId);
  }

  @Override
  public int hashCode() {
    return Objects.hash(id, rootVolume, name, size, sizeId, storageType, datastoreId);
  }


  @Override
  public String toString() {
    StringBuilder sb = new StringBuilder();
    sb.append("class ClusterServerCreateVolumes {\n");
    sb.append("    id: ").append(toIndentedString(id)).append("\n");
    sb.append("    rootVolume: ").append(toIndentedString(rootVolume)).append("\n");
    sb.append("    name: ").append(toIndentedString(name)).append("\n");
    sb.append("    size: ").append(toIndentedString(size)).append("\n");
    sb.append("    sizeId: ").append(toIndentedString(sizeId)).append("\n");
    sb.append("    storageType: ").append(toIndentedString(storageType)).append("\n");
    sb.append("    datastoreId: ").append(toIndentedString(datastoreId)).append("\n");
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
