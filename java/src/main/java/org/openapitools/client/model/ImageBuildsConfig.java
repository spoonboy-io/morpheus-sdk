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
import org.openapitools.client.model.ImageBuildsConfigConfig;
import org.openapitools.client.model.ImageBuildsConfigInstance;
import org.openapitools.client.model.ImageBuildsConfigNetworkInterfaces;
import org.openapitools.client.model.ImageBuildsConfigPlan;
import org.openapitools.client.model.ImageBuildsConfigVolumes;

/**
 * ImageBuildsConfig
 */
@javax.annotation.Generated(value = "org.openapitools.codegen.languages.JavaClientCodegen", date = "2023-08-21T11:26:47.028Z[GMT]")
public class ImageBuildsConfig {
  public static final String SERIALIZED_NAME_INSTANCE = "instance";
  @SerializedName(SERIALIZED_NAME_INSTANCE)
  private ImageBuildsConfigInstance instance;

  public static final String SERIALIZED_NAME_NETWORK_INTERFACES = "networkInterfaces";
  @SerializedName(SERIALIZED_NAME_NETWORK_INTERFACES)
  private List<ImageBuildsConfigNetworkInterfaces> networkInterfaces = null;

  public static final String SERIALIZED_NAME_VOLUMES = "volumes";
  @SerializedName(SERIALIZED_NAME_VOLUMES)
  private List<ImageBuildsConfigVolumes> volumes = null;

  public static final String SERIALIZED_NAME_STORAGE_CONTROLLERS = "storageControllers";
  @SerializedName(SERIALIZED_NAME_STORAGE_CONTROLLERS)
  private List<Object> storageControllers = null;

  public static final String SERIALIZED_NAME_ZONE_ID = "zoneId";
  @SerializedName(SERIALIZED_NAME_ZONE_ID)
  private Long zoneId;

  public static final String SERIALIZED_NAME_CONFIG = "config";
  @SerializedName(SERIALIZED_NAME_CONFIG)
  private ImageBuildsConfigConfig config;

  public static final String SERIALIZED_NAME_PLAN = "plan";
  @SerializedName(SERIALIZED_NAME_PLAN)
  private ImageBuildsConfigPlan plan;


  public ImageBuildsConfig instance(ImageBuildsConfigInstance instance) {
    
    this.instance = instance;
    return this;
  }

   /**
   * Get instance
   * @return instance
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public ImageBuildsConfigInstance getInstance() {
    return instance;
  }


  public void setInstance(ImageBuildsConfigInstance instance) {
    this.instance = instance;
  }


  public ImageBuildsConfig networkInterfaces(List<ImageBuildsConfigNetworkInterfaces> networkInterfaces) {
    
    this.networkInterfaces = networkInterfaces;
    return this;
  }

  public ImageBuildsConfig addNetworkInterfacesItem(ImageBuildsConfigNetworkInterfaces networkInterfacesItem) {
    if (this.networkInterfaces == null) {
      this.networkInterfaces = new ArrayList<ImageBuildsConfigNetworkInterfaces>();
    }
    this.networkInterfaces.add(networkInterfacesItem);
    return this;
  }

   /**
   * Get networkInterfaces
   * @return networkInterfaces
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public List<ImageBuildsConfigNetworkInterfaces> getNetworkInterfaces() {
    return networkInterfaces;
  }


  public void setNetworkInterfaces(List<ImageBuildsConfigNetworkInterfaces> networkInterfaces) {
    this.networkInterfaces = networkInterfaces;
  }


  public ImageBuildsConfig volumes(List<ImageBuildsConfigVolumes> volumes) {
    
    this.volumes = volumes;
    return this;
  }

  public ImageBuildsConfig addVolumesItem(ImageBuildsConfigVolumes volumesItem) {
    if (this.volumes == null) {
      this.volumes = new ArrayList<ImageBuildsConfigVolumes>();
    }
    this.volumes.add(volumesItem);
    return this;
  }

   /**
   * Get volumes
   * @return volumes
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public List<ImageBuildsConfigVolumes> getVolumes() {
    return volumes;
  }


  public void setVolumes(List<ImageBuildsConfigVolumes> volumes) {
    this.volumes = volumes;
  }


  public ImageBuildsConfig storageControllers(List<Object> storageControllers) {
    
    this.storageControllers = storageControllers;
    return this;
  }

  public ImageBuildsConfig addStorageControllersItem(Object storageControllersItem) {
    if (this.storageControllers == null) {
      this.storageControllers = new ArrayList<Object>();
    }
    this.storageControllers.add(storageControllersItem);
    return this;
  }

   /**
   * Get storageControllers
   * @return storageControllers
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public List<Object> getStorageControllers() {
    return storageControllers;
  }


  public void setStorageControllers(List<Object> storageControllers) {
    this.storageControllers = storageControllers;
  }


  public ImageBuildsConfig zoneId(Long zoneId) {
    
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


  public ImageBuildsConfig config(ImageBuildsConfigConfig config) {
    
    this.config = config;
    return this;
  }

   /**
   * Get config
   * @return config
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public ImageBuildsConfigConfig getConfig() {
    return config;
  }


  public void setConfig(ImageBuildsConfigConfig config) {
    this.config = config;
  }


  public ImageBuildsConfig plan(ImageBuildsConfigPlan plan) {
    
    this.plan = plan;
    return this;
  }

   /**
   * Get plan
   * @return plan
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public ImageBuildsConfigPlan getPlan() {
    return plan;
  }


  public void setPlan(ImageBuildsConfigPlan plan) {
    this.plan = plan;
  }


  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (o == null || getClass() != o.getClass()) {
      return false;
    }
    ImageBuildsConfig imageBuildsConfig = (ImageBuildsConfig) o;
    return Objects.equals(this.instance, imageBuildsConfig.instance) &&
        Objects.equals(this.networkInterfaces, imageBuildsConfig.networkInterfaces) &&
        Objects.equals(this.volumes, imageBuildsConfig.volumes) &&
        Objects.equals(this.storageControllers, imageBuildsConfig.storageControllers) &&
        Objects.equals(this.zoneId, imageBuildsConfig.zoneId) &&
        Objects.equals(this.config, imageBuildsConfig.config) &&
        Objects.equals(this.plan, imageBuildsConfig.plan);
  }

  @Override
  public int hashCode() {
    return Objects.hash(instance, networkInterfaces, volumes, storageControllers, zoneId, config, plan);
  }


  @Override
  public String toString() {
    StringBuilder sb = new StringBuilder();
    sb.append("class ImageBuildsConfig {\n");
    sb.append("    instance: ").append(toIndentedString(instance)).append("\n");
    sb.append("    networkInterfaces: ").append(toIndentedString(networkInterfaces)).append("\n");
    sb.append("    volumes: ").append(toIndentedString(volumes)).append("\n");
    sb.append("    storageControllers: ").append(toIndentedString(storageControllers)).append("\n");
    sb.append("    zoneId: ").append(toIndentedString(zoneId)).append("\n");
    sb.append("    config: ").append(toIndentedString(config)).append("\n");
    sb.append("    plan: ").append(toIndentedString(plan)).append("\n");
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
