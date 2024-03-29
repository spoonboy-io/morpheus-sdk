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

/**
 * NetworkDhcpRelayCreate
 */
@javax.annotation.Generated(value = "org.openapitools.codegen.languages.JavaClientCodegen", date = "2023-08-21T11:26:47.028Z[GMT]")
public class NetworkDhcpRelayCreate {
  public static final String SERIALIZED_NAME_NAME = "name";
  @SerializedName(SERIALIZED_NAME_NAME)
  private String name;

  public static final String SERIALIZED_NAME_SERVER_IP_ADDRESSES = "serverIpAddresses";
  @SerializedName(SERIALIZED_NAME_SERVER_IP_ADDRESSES)
  private List<String> serverIpAddresses = null;


  public NetworkDhcpRelayCreate name(String name) {
    
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


  public NetworkDhcpRelayCreate serverIpAddresses(List<String> serverIpAddresses) {
    
    this.serverIpAddresses = serverIpAddresses;
    return this;
  }

  public NetworkDhcpRelayCreate addServerIpAddressesItem(String serverIpAddressesItem) {
    if (this.serverIpAddresses == null) {
      this.serverIpAddresses = new ArrayList<String>();
    }
    this.serverIpAddresses.add(serverIpAddressesItem);
    return this;
  }

   /**
   * Get serverIpAddresses
   * @return serverIpAddresses
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public List<String> getServerIpAddresses() {
    return serverIpAddresses;
  }


  public void setServerIpAddresses(List<String> serverIpAddresses) {
    this.serverIpAddresses = serverIpAddresses;
  }


  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (o == null || getClass() != o.getClass()) {
      return false;
    }
    NetworkDhcpRelayCreate networkDhcpRelayCreate = (NetworkDhcpRelayCreate) o;
    return Objects.equals(this.name, networkDhcpRelayCreate.name) &&
        Objects.equals(this.serverIpAddresses, networkDhcpRelayCreate.serverIpAddresses);
  }

  @Override
  public int hashCode() {
    return Objects.hash(name, serverIpAddresses);
  }


  @Override
  public String toString() {
    StringBuilder sb = new StringBuilder();
    sb.append("class NetworkDhcpRelayCreate {\n");
    sb.append("    name: ").append(toIndentedString(name)).append("\n");
    sb.append("    serverIpAddresses: ").append(toIndentedString(serverIpAddresses)).append("\n");
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

