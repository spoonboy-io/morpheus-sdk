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
 * For a full list of available firewall rule options, see ruleOptionTypes in the specific Network Router Type detail
 */
@ApiModel(description = "For a full list of available firewall rule options, see ruleOptionTypes in the specific Network Router Type detail")
@javax.annotation.Generated(value = "org.openapitools.codegen.languages.JavaClientCodegen", date = "2023-08-21T11:26:47.028Z[GMT]")
public class NetworkRouterFirewallRuleCreate {
  public static final String SERIALIZED_NAME_NAME = "name";
  @SerializedName(SERIALIZED_NAME_NAME)
  private String name;

  public static final String SERIALIZED_NAME_ENABLED = "enabled";
  @SerializedName(SERIALIZED_NAME_ENABLED)
  private Boolean enabled = true;

  public static final String SERIALIZED_NAME_PRIORITY = "priority";
  @SerializedName(SERIALIZED_NAME_PRIORITY)
  private Long priority;


  public NetworkRouterFirewallRuleCreate name(String name) {
    
    this.name = name;
    return this;
  }

   /**
   * Firewall rule name
   * @return name
  **/
  @ApiModelProperty(required = true, value = "Firewall rule name")

  public String getName() {
    return name;
  }


  public void setName(String name) {
    this.name = name;
  }


  public NetworkRouterFirewallRuleCreate enabled(Boolean enabled) {
    
    this.enabled = enabled;
    return this;
  }

   /**
   * Can be used to enable / disable the rule (true, false). Default is on
   * @return enabled
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "Can be used to enable / disable the rule (true, false). Default is on")

  public Boolean getEnabled() {
    return enabled;
  }


  public void setEnabled(Boolean enabled) {
    this.enabled = enabled;
  }


  public NetworkRouterFirewallRuleCreate priority(Long priority) {
    
    this.priority = priority;
    return this;
  }

   /**
   * Priority for rule
   * @return priority
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "Priority for rule")

  public Long getPriority() {
    return priority;
  }


  public void setPriority(Long priority) {
    this.priority = priority;
  }


  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (o == null || getClass() != o.getClass()) {
      return false;
    }
    NetworkRouterFirewallRuleCreate networkRouterFirewallRuleCreate = (NetworkRouterFirewallRuleCreate) o;
    return Objects.equals(this.name, networkRouterFirewallRuleCreate.name) &&
        Objects.equals(this.enabled, networkRouterFirewallRuleCreate.enabled) &&
        Objects.equals(this.priority, networkRouterFirewallRuleCreate.priority);
  }

  @Override
  public int hashCode() {
    return Objects.hash(name, enabled, priority);
  }


  @Override
  public String toString() {
    StringBuilder sb = new StringBuilder();
    sb.append("class NetworkRouterFirewallRuleCreate {\n");
    sb.append("    name: ").append(toIndentedString(name)).append("\n");
    sb.append("    enabled: ").append(toIndentedString(enabled)).append("\n");
    sb.append("    priority: ").append(toIndentedString(priority)).append("\n");
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
