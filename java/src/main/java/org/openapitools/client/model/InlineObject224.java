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
import org.openapitools.client.model.ApiServersIdResizeServer;
import org.openapitools.client.model.ApiServersIdResizeServicePlanOptions;
import org.openapitools.client.model.InstanceCreateNetwork;
import org.openapitools.client.model.InstanceCreateVolume;

/**
 * InlineObject224
 */
@javax.annotation.Generated(value = "org.openapitools.codegen.languages.JavaClientCodegen", date = "2023-08-21T11:26:47.028Z[GMT]")
public class InlineObject224 {
  public static final String SERIALIZED_NAME_SERVER = "server";
  @SerializedName(SERIALIZED_NAME_SERVER)
  private ApiServersIdResizeServer server;

  public static final String SERIALIZED_NAME_SERVICE_PLAN_OPTIONS = "servicePlanOptions";
  @SerializedName(SERIALIZED_NAME_SERVICE_PLAN_OPTIONS)
  private ApiServersIdResizeServicePlanOptions servicePlanOptions;

  public static final String SERIALIZED_NAME_VOLUMES = "volumes";
  @SerializedName(SERIALIZED_NAME_VOLUMES)
  private List<InstanceCreateVolume> volumes = null;

  public static final String SERIALIZED_NAME_DELETE_ORIGINAL_VOLUMES = "deleteOriginalVolumes";
  @SerializedName(SERIALIZED_NAME_DELETE_ORIGINAL_VOLUMES)
  private Boolean deleteOriginalVolumes = false;

  public static final String SERIALIZED_NAME_NETWORK_INTERFACES = "networkInterfaces";
  @SerializedName(SERIALIZED_NAME_NETWORK_INTERFACES)
  private List<InstanceCreateNetwork> networkInterfaces = null;


  public InlineObject224 server(ApiServersIdResizeServer server) {
    
    this.server = server;
    return this;
  }

   /**
   * Get server
   * @return server
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public ApiServersIdResizeServer getServer() {
    return server;
  }


  public void setServer(ApiServersIdResizeServer server) {
    this.server = server;
  }


  public InlineObject224 servicePlanOptions(ApiServersIdResizeServicePlanOptions servicePlanOptions) {
    
    this.servicePlanOptions = servicePlanOptions;
    return this;
  }

   /**
   * Get servicePlanOptions
   * @return servicePlanOptions
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public ApiServersIdResizeServicePlanOptions getServicePlanOptions() {
    return servicePlanOptions;
  }


  public void setServicePlanOptions(ApiServersIdResizeServicePlanOptions servicePlanOptions) {
    this.servicePlanOptions = servicePlanOptions;
  }


  public InlineObject224 volumes(List<InstanceCreateVolume> volumes) {
    
    this.volumes = volumes;
    return this;
  }

  public InlineObject224 addVolumesItem(InstanceCreateVolume volumesItem) {
    if (this.volumes == null) {
      this.volumes = new ArrayList<InstanceCreateVolume>();
    }
    this.volumes.add(volumesItem);
    return this;
  }

   /**
   * List of volumes with their new sizes.
   * @return volumes
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "List of volumes with their new sizes.")

  public List<InstanceCreateVolume> getVolumes() {
    return volumes;
  }


  public void setVolumes(List<InstanceCreateVolume> volumes) {
    this.volumes = volumes;
  }


  public InlineObject224 deleteOriginalVolumes(Boolean deleteOriginalVolumes) {
    
    this.deleteOriginalVolumes = deleteOriginalVolumes;
    return this;
  }

   /**
   * Delete the original volumes after resizing. (Amazon only)
   * @return deleteOriginalVolumes
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "Delete the original volumes after resizing. (Amazon only)")

  public Boolean getDeleteOriginalVolumes() {
    return deleteOriginalVolumes;
  }


  public void setDeleteOriginalVolumes(Boolean deleteOriginalVolumes) {
    this.deleteOriginalVolumes = deleteOriginalVolumes;
  }


  public InlineObject224 networkInterfaces(List<InstanceCreateNetwork> networkInterfaces) {
    
    this.networkInterfaces = networkInterfaces;
    return this;
  }

  public InlineObject224 addNetworkInterfacesItem(InstanceCreateNetwork networkInterfacesItem) {
    if (this.networkInterfaces == null) {
      this.networkInterfaces = new ArrayList<InstanceCreateNetwork>();
    }
    this.networkInterfaces.add(networkInterfacesItem);
    return this;
  }

   /**
   * Key for network configurations. Include id to update an existing interface.
   * @return networkInterfaces
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "Key for network configurations. Include id to update an existing interface.")

  public List<InstanceCreateNetwork> getNetworkInterfaces() {
    return networkInterfaces;
  }


  public void setNetworkInterfaces(List<InstanceCreateNetwork> networkInterfaces) {
    this.networkInterfaces = networkInterfaces;
  }


  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (o == null || getClass() != o.getClass()) {
      return false;
    }
    InlineObject224 inlineObject224 = (InlineObject224) o;
    return Objects.equals(this.server, inlineObject224.server) &&
        Objects.equals(this.servicePlanOptions, inlineObject224.servicePlanOptions) &&
        Objects.equals(this.volumes, inlineObject224.volumes) &&
        Objects.equals(this.deleteOriginalVolumes, inlineObject224.deleteOriginalVolumes) &&
        Objects.equals(this.networkInterfaces, inlineObject224.networkInterfaces);
  }

  @Override
  public int hashCode() {
    return Objects.hash(server, servicePlanOptions, volumes, deleteOriginalVolumes, networkInterfaces);
  }


  @Override
  public String toString() {
    StringBuilder sb = new StringBuilder();
    sb.append("class InlineObject224 {\n");
    sb.append("    server: ").append(toIndentedString(server)).append("\n");
    sb.append("    servicePlanOptions: ").append(toIndentedString(servicePlanOptions)).append("\n");
    sb.append("    volumes: ").append(toIndentedString(volumes)).append("\n");
    sb.append("    deleteOriginalVolumes: ").append(toIndentedString(deleteOriginalVolumes)).append("\n");
    sb.append("    networkInterfaces: ").append(toIndentedString(networkInterfaces)).append("\n");
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

