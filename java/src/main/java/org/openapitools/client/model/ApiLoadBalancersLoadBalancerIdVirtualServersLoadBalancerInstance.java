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
 * ApiLoadBalancersLoadBalancerIdVirtualServersLoadBalancerInstance
 */
@javax.annotation.Generated(value = "org.openapitools.codegen.languages.JavaClientCodegen", date = "2023-08-21T11:26:47.028Z[GMT]")
public class ApiLoadBalancersLoadBalancerIdVirtualServersLoadBalancerInstance {
  public static final String SERIALIZED_NAME_VIP_NAME = "vipName";
  @SerializedName(SERIALIZED_NAME_VIP_NAME)
  private String vipName;

  public static final String SERIALIZED_NAME_DESCRIPTION = "description";
  @SerializedName(SERIALIZED_NAME_DESCRIPTION)
  private String description;

  public static final String SERIALIZED_NAME_VIP_ADDRESS = "vipAddress";
  @SerializedName(SERIALIZED_NAME_VIP_ADDRESS)
  private String vipAddress;

  public static final String SERIALIZED_NAME_VIP_PORT = "vipPort";
  @SerializedName(SERIALIZED_NAME_VIP_PORT)
  private String vipPort;

  public static final String SERIALIZED_NAME_VIP_PROTOCOL = "vipProtocol";
  @SerializedName(SERIALIZED_NAME_VIP_PROTOCOL)
  private String vipProtocol;

  public static final String SERIALIZED_NAME_VIP_HOSTNAME = "vipHostname";
  @SerializedName(SERIALIZED_NAME_VIP_HOSTNAME)
  private String vipHostname;

  public static final String SERIALIZED_NAME_SSL_CERT = "sslCert";
  @SerializedName(SERIALIZED_NAME_SSL_CERT)
  private Long sslCert;

  public static final String SERIALIZED_NAME_SSL_SERVER_CERT = "sslServerCert";
  @SerializedName(SERIALIZED_NAME_SSL_SERVER_CERT)
  private Long sslServerCert;

  public static final String SERIALIZED_NAME_CONFIG = "config";
  @SerializedName(SERIALIZED_NAME_CONFIG)
  private Object config;


  public ApiLoadBalancersLoadBalancerIdVirtualServersLoadBalancerInstance vipName(String vipName) {
    
    this.vipName = vipName;
    return this;
  }

   /**
   * VIP Name
   * @return vipName
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "VIP Name")

  public String getVipName() {
    return vipName;
  }


  public void setVipName(String vipName) {
    this.vipName = vipName;
  }


  public ApiLoadBalancersLoadBalancerIdVirtualServersLoadBalancerInstance description(String description) {
    
    this.description = description;
    return this;
  }

   /**
   * Description
   * @return description
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "Description")

  public String getDescription() {
    return description;
  }


  public void setDescription(String description) {
    this.description = description;
  }


  public ApiLoadBalancersLoadBalancerIdVirtualServersLoadBalancerInstance vipAddress(String vipAddress) {
    
    this.vipAddress = vipAddress;
    return this;
  }

   /**
   * VIP Address
   * @return vipAddress
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "VIP Address")

  public String getVipAddress() {
    return vipAddress;
  }


  public void setVipAddress(String vipAddress) {
    this.vipAddress = vipAddress;
  }


  public ApiLoadBalancersLoadBalancerIdVirtualServersLoadBalancerInstance vipPort(String vipPort) {
    
    this.vipPort = vipPort;
    return this;
  }

   /**
   * VIP Port
   * @return vipPort
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "VIP Port")

  public String getVipPort() {
    return vipPort;
  }


  public void setVipPort(String vipPort) {
    this.vipPort = vipPort;
  }


  public ApiLoadBalancersLoadBalancerIdVirtualServersLoadBalancerInstance vipProtocol(String vipProtocol) {
    
    this.vipProtocol = vipProtocol;
    return this;
  }

   /**
   * VIP Protocol
   * @return vipProtocol
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "VIP Protocol")

  public String getVipProtocol() {
    return vipProtocol;
  }


  public void setVipProtocol(String vipProtocol) {
    this.vipProtocol = vipProtocol;
  }


  public ApiLoadBalancersLoadBalancerIdVirtualServersLoadBalancerInstance vipHostname(String vipHostname) {
    
    this.vipHostname = vipHostname;
    return this;
  }

   /**
   * VIP Hostname
   * @return vipHostname
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "VIP Hostname")

  public String getVipHostname() {
    return vipHostname;
  }


  public void setVipHostname(String vipHostname) {
    this.vipHostname = vipHostname;
  }


  public ApiLoadBalancersLoadBalancerIdVirtualServersLoadBalancerInstance sslCert(Long sslCert) {
    
    this.sslCert = sslCert;
    return this;
  }

   /**
   * SSL Client Certificate ID
   * @return sslCert
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "SSL Client Certificate ID")

  public Long getSslCert() {
    return sslCert;
  }


  public void setSslCert(Long sslCert) {
    this.sslCert = sslCert;
  }


  public ApiLoadBalancersLoadBalancerIdVirtualServersLoadBalancerInstance sslServerCert(Long sslServerCert) {
    
    this.sslServerCert = sslServerCert;
    return this;
  }

   /**
   * SSL Server Certificate ID
   * @return sslServerCert
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "SSL Server Certificate ID")

  public Long getSslServerCert() {
    return sslServerCert;
  }


  public void setSslServerCert(Long sslServerCert) {
    this.sslServerCert = sslServerCert;
  }


  public ApiLoadBalancersLoadBalancerIdVirtualServersLoadBalancerInstance config(Object config) {
    
    this.config = config;
    return this;
  }

   /**
   * Configuration object with parameters that vary by type.
   * @return config
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "Configuration object with parameters that vary by type.")

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
    ApiLoadBalancersLoadBalancerIdVirtualServersLoadBalancerInstance apiLoadBalancersLoadBalancerIdVirtualServersLoadBalancerInstance = (ApiLoadBalancersLoadBalancerIdVirtualServersLoadBalancerInstance) o;
    return Objects.equals(this.vipName, apiLoadBalancersLoadBalancerIdVirtualServersLoadBalancerInstance.vipName) &&
        Objects.equals(this.description, apiLoadBalancersLoadBalancerIdVirtualServersLoadBalancerInstance.description) &&
        Objects.equals(this.vipAddress, apiLoadBalancersLoadBalancerIdVirtualServersLoadBalancerInstance.vipAddress) &&
        Objects.equals(this.vipPort, apiLoadBalancersLoadBalancerIdVirtualServersLoadBalancerInstance.vipPort) &&
        Objects.equals(this.vipProtocol, apiLoadBalancersLoadBalancerIdVirtualServersLoadBalancerInstance.vipProtocol) &&
        Objects.equals(this.vipHostname, apiLoadBalancersLoadBalancerIdVirtualServersLoadBalancerInstance.vipHostname) &&
        Objects.equals(this.sslCert, apiLoadBalancersLoadBalancerIdVirtualServersLoadBalancerInstance.sslCert) &&
        Objects.equals(this.sslServerCert, apiLoadBalancersLoadBalancerIdVirtualServersLoadBalancerInstance.sslServerCert) &&
        Objects.equals(this.config, apiLoadBalancersLoadBalancerIdVirtualServersLoadBalancerInstance.config);
  }

  @Override
  public int hashCode() {
    return Objects.hash(vipName, description, vipAddress, vipPort, vipProtocol, vipHostname, sslCert, sslServerCert, config);
  }


  @Override
  public String toString() {
    StringBuilder sb = new StringBuilder();
    sb.append("class ApiLoadBalancersLoadBalancerIdVirtualServersLoadBalancerInstance {\n");
    sb.append("    vipName: ").append(toIndentedString(vipName)).append("\n");
    sb.append("    description: ").append(toIndentedString(description)).append("\n");
    sb.append("    vipAddress: ").append(toIndentedString(vipAddress)).append("\n");
    sb.append("    vipPort: ").append(toIndentedString(vipPort)).append("\n");
    sb.append("    vipProtocol: ").append(toIndentedString(vipProtocol)).append("\n");
    sb.append("    vipHostname: ").append(toIndentedString(vipHostname)).append("\n");
    sb.append("    sslCert: ").append(toIndentedString(sslCert)).append("\n");
    sb.append("    sslServerCert: ").append(toIndentedString(sslServerCert)).append("\n");
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
