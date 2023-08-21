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
 * The parameters for creating a network firewall rule group is type dependent. The following lists the common parameters. See get a specific type to list available options for the network server type. 
 */
@ApiModel(description = "The parameters for creating a network firewall rule group is type dependent. The following lists the common parameters. See get a specific type to list available options for the network server type. ")
@javax.annotation.Generated(value = "org.openapitools.codegen.languages.JavaClientCodegen", date = "2023-08-21T11:26:47.028Z[GMT]")
public class NetworkFirewallRuleGroupCreate {
  public static final String SERIALIZED_NAME_NAME = "name";
  @SerializedName(SERIALIZED_NAME_NAME)
  private String name;

  public static final String SERIALIZED_NAME_DESCRIPTION = "description";
  @SerializedName(SERIALIZED_NAME_DESCRIPTION)
  private String description;

  public static final String SERIALIZED_NAME_PRIORITY = "priority";
  @SerializedName(SERIALIZED_NAME_PRIORITY)
  private Long priority;

  public static final String SERIALIZED_NAME_EXTERNAL_TYPE = "externalType";
  @SerializedName(SERIALIZED_NAME_EXTERNAL_TYPE)
  private String externalType;

  public static final String SERIALIZED_NAME_GROUP_LAYER = "groupLayer";
  @SerializedName(SERIALIZED_NAME_GROUP_LAYER)
  private String groupLayer;


  public NetworkFirewallRuleGroupCreate name(String name) {
    
    this.name = name;
    return this;
  }

   /**
   * Network firewall rule group name
   * @return name
  **/
  @ApiModelProperty(required = true, value = "Network firewall rule group name")

  public String getName() {
    return name;
  }


  public void setName(String name) {
    this.name = name;
  }


  public NetworkFirewallRuleGroupCreate description(String description) {
    
    this.description = description;
    return this;
  }

   /**
   * Network firewall rule group description
   * @return description
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "Network firewall rule group description")

  public String getDescription() {
    return description;
  }


  public void setDescription(String description) {
    this.description = description;
  }


  public NetworkFirewallRuleGroupCreate priority(Long priority) {
    
    this.priority = priority;
    return this;
  }

   /**
   * Network firewall rule group priority
   * @return priority
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "Network firewall rule group priority")

  public Long getPriority() {
    return priority;
  }


  public void setPriority(Long priority) {
    this.priority = priority;
  }


  public NetworkFirewallRuleGroupCreate externalType(String externalType) {
    
    this.externalType = externalType;
    return this;
  }

   /**
   * Use SecurityPolicy
   * @return externalType
  **/
  @ApiModelProperty(required = true, value = "Use SecurityPolicy")

  public String getExternalType() {
    return externalType;
  }


  public void setExternalType(String externalType) {
    this.externalType = externalType;
  }


  public NetworkFirewallRuleGroupCreate groupLayer(String groupLayer) {
    
    this.groupLayer = groupLayer;
    return this;
  }

   /**
   * Get groupLayer
   * @return groupLayer
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public String getGroupLayer() {
    return groupLayer;
  }


  public void setGroupLayer(String groupLayer) {
    this.groupLayer = groupLayer;
  }


  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (o == null || getClass() != o.getClass()) {
      return false;
    }
    NetworkFirewallRuleGroupCreate networkFirewallRuleGroupCreate = (NetworkFirewallRuleGroupCreate) o;
    return Objects.equals(this.name, networkFirewallRuleGroupCreate.name) &&
        Objects.equals(this.description, networkFirewallRuleGroupCreate.description) &&
        Objects.equals(this.priority, networkFirewallRuleGroupCreate.priority) &&
        Objects.equals(this.externalType, networkFirewallRuleGroupCreate.externalType) &&
        Objects.equals(this.groupLayer, networkFirewallRuleGroupCreate.groupLayer);
  }

  @Override
  public int hashCode() {
    return Objects.hash(name, description, priority, externalType, groupLayer);
  }


  @Override
  public String toString() {
    StringBuilder sb = new StringBuilder();
    sb.append("class NetworkFirewallRuleGroupCreate {\n");
    sb.append("    name: ").append(toIndentedString(name)).append("\n");
    sb.append("    description: ").append(toIndentedString(description)).append("\n");
    sb.append("    priority: ").append(toIndentedString(priority)).append("\n");
    sb.append("    externalType: ").append(toIndentedString(externalType)).append("\n");
    sb.append("    groupLayer: ").append(toIndentedString(groupLayer)).append("\n");
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
