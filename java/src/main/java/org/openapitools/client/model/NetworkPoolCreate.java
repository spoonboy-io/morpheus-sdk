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
import org.openapitools.client.model.NetworkPoolCreateIpRanges;
import org.openapitools.client.model.NetworkPoolCreateType;

/**
 * NetworkPoolCreate
 */
@javax.annotation.Generated(value = "org.openapitools.codegen.languages.JavaClientCodegen", date = "2023-08-21T11:26:47.028Z[GMT]")
public class NetworkPoolCreate {
  public static final String SERIALIZED_NAME_NAME = "name";
  @SerializedName(SERIALIZED_NAME_NAME)
  private String name;

  public static final String SERIALIZED_NAME_TYPE = "type";
  @SerializedName(SERIALIZED_NAME_TYPE)
  private NetworkPoolCreateType type;

  public static final String SERIALIZED_NAME_IP_RANGES = "ipRanges";
  @SerializedName(SERIALIZED_NAME_IP_RANGES)
  private List<NetworkPoolCreateIpRanges> ipRanges = null;

  public static final String SERIALIZED_NAME_CONFIG = "config";
  @SerializedName(SERIALIZED_NAME_CONFIG)
  private Object config;


  public NetworkPoolCreate name(String name) {
    
    this.name = name;
    return this;
  }

   /**
   * Name
   * @return name
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "Name")

  public String getName() {
    return name;
  }


  public void setName(String name) {
    this.name = name;
  }


  public NetworkPoolCreate type(NetworkPoolCreateType type) {
    
    this.type = type;
    return this;
  }

   /**
   * Get type
   * @return type
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public NetworkPoolCreateType getType() {
    return type;
  }


  public void setType(NetworkPoolCreateType type) {
    this.type = type;
  }


  public NetworkPoolCreate ipRanges(List<NetworkPoolCreateIpRanges> ipRanges) {
    
    this.ipRanges = ipRanges;
    return this;
  }

  public NetworkPoolCreate addIpRangesItem(NetworkPoolCreateIpRanges ipRangesItem) {
    if (this.ipRanges == null) {
      this.ipRanges = new ArrayList<NetworkPoolCreateIpRanges>();
    }
    this.ipRanges.add(ipRangesItem);
    return this;
  }

   /**
   * Array of IP range objects. Type &#39;morpheus&#39; expects startAddress and endAddress. Type &#39;morpheusipv6&#39; expects a cidrIPv6.
   * @return ipRanges
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "Array of IP range objects. Type 'morpheus' expects startAddress and endAddress. Type 'morpheusipv6' expects a cidrIPv6.")

  public List<NetworkPoolCreateIpRanges> getIpRanges() {
    return ipRanges;
  }


  public void setIpRanges(List<NetworkPoolCreateIpRanges> ipRanges) {
    this.ipRanges = ipRanges;
  }


  public NetworkPoolCreate config(Object config) {
    
    this.config = config;
    return this;
  }

   /**
   * Configuration object with parameters that vary by pool type.
   * @return config
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "Configuration object with parameters that vary by pool type.")

  public Object getConfig() {
    return config;
  }


  public void setConfig(Object config) {
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
    NetworkPoolCreate networkPoolCreate = (NetworkPoolCreate) o;
    return Objects.equals(this.name, networkPoolCreate.name) &&
        Objects.equals(this.type, networkPoolCreate.type) &&
        Objects.equals(this.ipRanges, networkPoolCreate.ipRanges) &&
        Objects.equals(this.config, networkPoolCreate.config);
  }

  @Override
  public int hashCode() {
    return Objects.hash(name, type, ipRanges, config);
  }


  @Override
  public String toString() {
    StringBuilder sb = new StringBuilder();
    sb.append("class NetworkPoolCreate {\n");
    sb.append("    name: ").append(toIndentedString(name)).append("\n");
    sb.append("    type: ").append(toIndentedString(type)).append("\n");
    sb.append("    ipRanges: ").append(toIndentedString(ipRanges)).append("\n");
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

