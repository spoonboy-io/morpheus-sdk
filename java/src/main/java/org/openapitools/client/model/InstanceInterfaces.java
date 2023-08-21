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
import org.openapitools.client.model.InstanceNetwork;

/**
 * InstanceInterfaces
 */
@javax.annotation.Generated(value = "org.openapitools.codegen.languages.JavaClientCodegen", date = "2023-08-21T11:26:47.028Z[GMT]")
public class InstanceInterfaces {
  public static final String SERIALIZED_NAME_ID = "id";
  @SerializedName(SERIALIZED_NAME_ID)
  private String id;

  public static final String SERIALIZED_NAME_NETWORK = "network";
  @SerializedName(SERIALIZED_NAME_NETWORK)
  private InstanceNetwork network;

  public static final String SERIALIZED_NAME_IP_ADDRESS = "ipAddress";
  @SerializedName(SERIALIZED_NAME_IP_ADDRESS)
  private String ipAddress;

  public static final String SERIALIZED_NAME_NETWORK_INTERFACE_TYPE_ID = "networkInterfaceTypeId";
  @SerializedName(SERIALIZED_NAME_NETWORK_INTERFACE_TYPE_ID)
  private Long networkInterfaceTypeId;

  public static final String SERIALIZED_NAME_IP_MODE = "ipMode";
  @SerializedName(SERIALIZED_NAME_IP_MODE)
  private String ipMode;


  public InstanceInterfaces id(String id) {
    
    this.id = id;
    return this;
  }

   /**
   * Get id
   * @return id
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public String getId() {
    return id;
  }


  public void setId(String id) {
    this.id = id;
  }


  public InstanceInterfaces network(InstanceNetwork network) {
    
    this.network = network;
    return this;
  }

   /**
   * Get network
   * @return network
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public InstanceNetwork getNetwork() {
    return network;
  }


  public void setNetwork(InstanceNetwork network) {
    this.network = network;
  }


  public InstanceInterfaces ipAddress(String ipAddress) {
    
    this.ipAddress = ipAddress;
    return this;
  }

   /**
   * Get ipAddress
   * @return ipAddress
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public String getIpAddress() {
    return ipAddress;
  }


  public void setIpAddress(String ipAddress) {
    this.ipAddress = ipAddress;
  }


  public InstanceInterfaces networkInterfaceTypeId(Long networkInterfaceTypeId) {
    
    this.networkInterfaceTypeId = networkInterfaceTypeId;
    return this;
  }

   /**
   * Get networkInterfaceTypeId
   * @return networkInterfaceTypeId
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public Long getNetworkInterfaceTypeId() {
    return networkInterfaceTypeId;
  }


  public void setNetworkInterfaceTypeId(Long networkInterfaceTypeId) {
    this.networkInterfaceTypeId = networkInterfaceTypeId;
  }


  public InstanceInterfaces ipMode(String ipMode) {
    
    this.ipMode = ipMode;
    return this;
  }

   /**
   * Get ipMode
   * @return ipMode
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public String getIpMode() {
    return ipMode;
  }


  public void setIpMode(String ipMode) {
    this.ipMode = ipMode;
  }


  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (o == null || getClass() != o.getClass()) {
      return false;
    }
    InstanceInterfaces instanceInterfaces = (InstanceInterfaces) o;
    return Objects.equals(this.id, instanceInterfaces.id) &&
        Objects.equals(this.network, instanceInterfaces.network) &&
        Objects.equals(this.ipAddress, instanceInterfaces.ipAddress) &&
        Objects.equals(this.networkInterfaceTypeId, instanceInterfaces.networkInterfaceTypeId) &&
        Objects.equals(this.ipMode, instanceInterfaces.ipMode);
  }

  @Override
  public int hashCode() {
    return Objects.hash(id, network, ipAddress, networkInterfaceTypeId, ipMode);
  }


  @Override
  public String toString() {
    StringBuilder sb = new StringBuilder();
    sb.append("class InstanceInterfaces {\n");
    sb.append("    id: ").append(toIndentedString(id)).append("\n");
    sb.append("    network: ").append(toIndentedString(network)).append("\n");
    sb.append("    ipAddress: ").append(toIndentedString(ipAddress)).append("\n");
    sb.append("    networkInterfaceTypeId: ").append(toIndentedString(networkInterfaceTypeId)).append("\n");
    sb.append("    ipMode: ").append(toIndentedString(ipMode)).append("\n");
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
