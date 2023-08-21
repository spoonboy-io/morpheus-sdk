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
import org.openapitools.client.model.InlineResponse20095NetworkRouterFirewallRuleGroups;

/**
 * InlineResponse20095NetworkRouterFirewall
 */
@javax.annotation.Generated(value = "org.openapitools.codegen.languages.JavaClientCodegen", date = "2023-08-21T11:26:47.028Z[GMT]")
public class InlineResponse20095NetworkRouterFirewall {
  public static final String SERIALIZED_NAME_ENABLED = "enabled";
  @SerializedName(SERIALIZED_NAME_ENABLED)
  private Boolean enabled;

  public static final String SERIALIZED_NAME_VERSION = "version";
  @SerializedName(SERIALIZED_NAME_VERSION)
  private String version;

  public static final String SERIALIZED_NAME_DEFAULT_POLICY = "defaultPolicy";
  @SerializedName(SERIALIZED_NAME_DEFAULT_POLICY)
  private String defaultPolicy;

  public static final String SERIALIZED_NAME_GLOBAL = "global";
  @SerializedName(SERIALIZED_NAME_GLOBAL)
  private String global;

  public static final String SERIALIZED_NAME_RULE_GROUPS = "ruleGroups";
  @SerializedName(SERIALIZED_NAME_RULE_GROUPS)
  private List<InlineResponse20095NetworkRouterFirewallRuleGroups> ruleGroups = null;


  public InlineResponse20095NetworkRouterFirewall enabled(Boolean enabled) {
    
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


  public InlineResponse20095NetworkRouterFirewall version(String version) {
    
    this.version = version;
    return this;
  }

   /**
   * Get version
   * @return version
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public String getVersion() {
    return version;
  }


  public void setVersion(String version) {
    this.version = version;
  }


  public InlineResponse20095NetworkRouterFirewall defaultPolicy(String defaultPolicy) {
    
    this.defaultPolicy = defaultPolicy;
    return this;
  }

   /**
   * Get defaultPolicy
   * @return defaultPolicy
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public String getDefaultPolicy() {
    return defaultPolicy;
  }


  public void setDefaultPolicy(String defaultPolicy) {
    this.defaultPolicy = defaultPolicy;
  }


  public InlineResponse20095NetworkRouterFirewall global(String global) {
    
    this.global = global;
    return this;
  }

   /**
   * Get global
   * @return global
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public String getGlobal() {
    return global;
  }


  public void setGlobal(String global) {
    this.global = global;
  }


  public InlineResponse20095NetworkRouterFirewall ruleGroups(List<InlineResponse20095NetworkRouterFirewallRuleGroups> ruleGroups) {
    
    this.ruleGroups = ruleGroups;
    return this;
  }

  public InlineResponse20095NetworkRouterFirewall addRuleGroupsItem(InlineResponse20095NetworkRouterFirewallRuleGroups ruleGroupsItem) {
    if (this.ruleGroups == null) {
      this.ruleGroups = new ArrayList<InlineResponse20095NetworkRouterFirewallRuleGroups>();
    }
    this.ruleGroups.add(ruleGroupsItem);
    return this;
  }

   /**
   * Get ruleGroups
   * @return ruleGroups
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public List<InlineResponse20095NetworkRouterFirewallRuleGroups> getRuleGroups() {
    return ruleGroups;
  }


  public void setRuleGroups(List<InlineResponse20095NetworkRouterFirewallRuleGroups> ruleGroups) {
    this.ruleGroups = ruleGroups;
  }


  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (o == null || getClass() != o.getClass()) {
      return false;
    }
    InlineResponse20095NetworkRouterFirewall inlineResponse20095NetworkRouterFirewall = (InlineResponse20095NetworkRouterFirewall) o;
    return Objects.equals(this.enabled, inlineResponse20095NetworkRouterFirewall.enabled) &&
        Objects.equals(this.version, inlineResponse20095NetworkRouterFirewall.version) &&
        Objects.equals(this.defaultPolicy, inlineResponse20095NetworkRouterFirewall.defaultPolicy) &&
        Objects.equals(this.global, inlineResponse20095NetworkRouterFirewall.global) &&
        Objects.equals(this.ruleGroups, inlineResponse20095NetworkRouterFirewall.ruleGroups);
  }

  @Override
  public int hashCode() {
    return Objects.hash(enabled, version, defaultPolicy, global, ruleGroups);
  }


  @Override
  public String toString() {
    StringBuilder sb = new StringBuilder();
    sb.append("class InlineResponse20095NetworkRouterFirewall {\n");
    sb.append("    enabled: ").append(toIndentedString(enabled)).append("\n");
    sb.append("    version: ").append(toIndentedString(version)).append("\n");
    sb.append("    defaultPolicy: ").append(toIndentedString(defaultPolicy)).append("\n");
    sb.append("    global: ").append(toIndentedString(global)).append("\n");
    sb.append("    ruleGroups: ").append(toIndentedString(ruleGroups)).append("\n");
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

