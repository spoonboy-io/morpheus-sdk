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
import org.openapitools.client.model.InlineResponse20094Network;

/**
 * InlineResponse20094Interfaces
 */
@javax.annotation.Generated(value = "org.openapitools.codegen.languages.JavaClientCodegen", date = "2023-08-21T11:26:47.028Z[GMT]")
public class InlineResponse20094Interfaces {
  public static final String SERIALIZED_NAME_ID = "id";
  @SerializedName(SERIALIZED_NAME_ID)
  private Long id;

  public static final String SERIALIZED_NAME_NAME = "name";
  @SerializedName(SERIALIZED_NAME_NAME)
  private String name;

  public static final String SERIALIZED_NAME_CODE = "code";
  @SerializedName(SERIALIZED_NAME_CODE)
  private String code;

  public static final String SERIALIZED_NAME_INTERFACE_TYPE = "interfaceType";
  @SerializedName(SERIALIZED_NAME_INTERFACE_TYPE)
  private String interfaceType;

  public static final String SERIALIZED_NAME_NETWORK_POSITION = "networkPosition";
  @SerializedName(SERIALIZED_NAME_NETWORK_POSITION)
  private String networkPosition;

  public static final String SERIALIZED_NAME_IP_ADDRESS = "ipAddress";
  @SerializedName(SERIALIZED_NAME_IP_ADDRESS)
  private String ipAddress;

  public static final String SERIALIZED_NAME_CIDR = "cidr";
  @SerializedName(SERIALIZED_NAME_CIDR)
  private String cidr;

  public static final String SERIALIZED_NAME_EXTERNAL_LINK = "externalLink";
  @SerializedName(SERIALIZED_NAME_EXTERNAL_LINK)
  private String externalLink;

  public static final String SERIALIZED_NAME_ENABLED = "enabled";
  @SerializedName(SERIALIZED_NAME_ENABLED)
  private Boolean enabled;

  public static final String SERIALIZED_NAME_NETWORK = "network";
  @SerializedName(SERIALIZED_NAME_NETWORK)
  private InlineResponse20094Network network;


  public InlineResponse20094Interfaces id(Long id) {
    
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


  public InlineResponse20094Interfaces name(String name) {
    
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


  public InlineResponse20094Interfaces code(String code) {
    
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


  public InlineResponse20094Interfaces interfaceType(String interfaceType) {
    
    this.interfaceType = interfaceType;
    return this;
  }

   /**
   * Get interfaceType
   * @return interfaceType
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public String getInterfaceType() {
    return interfaceType;
  }


  public void setInterfaceType(String interfaceType) {
    this.interfaceType = interfaceType;
  }


  public InlineResponse20094Interfaces networkPosition(String networkPosition) {
    
    this.networkPosition = networkPosition;
    return this;
  }

   /**
   * Get networkPosition
   * @return networkPosition
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public String getNetworkPosition() {
    return networkPosition;
  }


  public void setNetworkPosition(String networkPosition) {
    this.networkPosition = networkPosition;
  }


  public InlineResponse20094Interfaces ipAddress(String ipAddress) {
    
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


  public InlineResponse20094Interfaces cidr(String cidr) {
    
    this.cidr = cidr;
    return this;
  }

   /**
   * Get cidr
   * @return cidr
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public String getCidr() {
    return cidr;
  }


  public void setCidr(String cidr) {
    this.cidr = cidr;
  }


  public InlineResponse20094Interfaces externalLink(String externalLink) {
    
    this.externalLink = externalLink;
    return this;
  }

   /**
   * Get externalLink
   * @return externalLink
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public String getExternalLink() {
    return externalLink;
  }


  public void setExternalLink(String externalLink) {
    this.externalLink = externalLink;
  }


  public InlineResponse20094Interfaces enabled(Boolean enabled) {
    
    this.enabled = enabled;
    return this;
  }

   /**
   * Get enabled
   * @return enabled
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public Boolean getEnabled() {
    return enabled;
  }


  public void setEnabled(Boolean enabled) {
    this.enabled = enabled;
  }


  public InlineResponse20094Interfaces network(InlineResponse20094Network network) {
    
    this.network = network;
    return this;
  }

   /**
   * Get network
   * @return network
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public InlineResponse20094Network getNetwork() {
    return network;
  }


  public void setNetwork(InlineResponse20094Network network) {
    this.network = network;
  }


  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (o == null || getClass() != o.getClass()) {
      return false;
    }
    InlineResponse20094Interfaces inlineResponse20094Interfaces = (InlineResponse20094Interfaces) o;
    return Objects.equals(this.id, inlineResponse20094Interfaces.id) &&
        Objects.equals(this.name, inlineResponse20094Interfaces.name) &&
        Objects.equals(this.code, inlineResponse20094Interfaces.code) &&
        Objects.equals(this.interfaceType, inlineResponse20094Interfaces.interfaceType) &&
        Objects.equals(this.networkPosition, inlineResponse20094Interfaces.networkPosition) &&
        Objects.equals(this.ipAddress, inlineResponse20094Interfaces.ipAddress) &&
        Objects.equals(this.cidr, inlineResponse20094Interfaces.cidr) &&
        Objects.equals(this.externalLink, inlineResponse20094Interfaces.externalLink) &&
        Objects.equals(this.enabled, inlineResponse20094Interfaces.enabled) &&
        Objects.equals(this.network, inlineResponse20094Interfaces.network);
  }

  @Override
  public int hashCode() {
    return Objects.hash(id, name, code, interfaceType, networkPosition, ipAddress, cidr, externalLink, enabled, network);
  }


  @Override
  public String toString() {
    StringBuilder sb = new StringBuilder();
    sb.append("class InlineResponse20094Interfaces {\n");
    sb.append("    id: ").append(toIndentedString(id)).append("\n");
    sb.append("    name: ").append(toIndentedString(name)).append("\n");
    sb.append("    code: ").append(toIndentedString(code)).append("\n");
    sb.append("    interfaceType: ").append(toIndentedString(interfaceType)).append("\n");
    sb.append("    networkPosition: ").append(toIndentedString(networkPosition)).append("\n");
    sb.append("    ipAddress: ").append(toIndentedString(ipAddress)).append("\n");
    sb.append("    cidr: ").append(toIndentedString(cidr)).append("\n");
    sb.append("    externalLink: ").append(toIndentedString(externalLink)).append("\n");
    sb.append("    enabled: ").append(toIndentedString(enabled)).append("\n");
    sb.append("    network: ").append(toIndentedString(network)).append("\n");
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

