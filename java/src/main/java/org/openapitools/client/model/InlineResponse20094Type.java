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
 * InlineResponse20094Type
 */
@javax.annotation.Generated(value = "org.openapitools.codegen.languages.JavaClientCodegen", date = "2023-08-21T11:26:47.028Z[GMT]")
public class InlineResponse20094Type {
  public static final String SERIALIZED_NAME_ID = "id";
  @SerializedName(SERIALIZED_NAME_ID)
  private Long id;

  public static final String SERIALIZED_NAME_CODE = "code";
  @SerializedName(SERIALIZED_NAME_CODE)
  private String code;

  public static final String SERIALIZED_NAME_NAME = "name";
  @SerializedName(SERIALIZED_NAME_NAME)
  private String name;

  public static final String SERIALIZED_NAME_DESCRIPTION = "description";
  @SerializedName(SERIALIZED_NAME_DESCRIPTION)
  private String description;

  public static final String SERIALIZED_NAME_ENABLED = "enabled";
  @SerializedName(SERIALIZED_NAME_ENABLED)
  private Boolean enabled;

  public static final String SERIALIZED_NAME_CREATABLE = "creatable";
  @SerializedName(SERIALIZED_NAME_CREATABLE)
  private Boolean creatable;

  public static final String SERIALIZED_NAME_SELECTABLE = "selectable";
  @SerializedName(SERIALIZED_NAME_SELECTABLE)
  private Boolean selectable;

  public static final String SERIALIZED_NAME_HAS_FIREWALL = "hasFirewall";
  @SerializedName(SERIALIZED_NAME_HAS_FIREWALL)
  private Boolean hasFirewall;

  public static final String SERIALIZED_NAME_HAS_DHCP = "hasDhcp";
  @SerializedName(SERIALIZED_NAME_HAS_DHCP)
  private Boolean hasDhcp;

  public static final String SERIALIZED_NAME_HAS_ROUTING = "hasRouting";
  @SerializedName(SERIALIZED_NAME_HAS_ROUTING)
  private Boolean hasRouting;

  public static final String SERIALIZED_NAME_HAS_NETWORK_SERVER = "hasNetworkServer";
  @SerializedName(SERIALIZED_NAME_HAS_NETWORK_SERVER)
  private Boolean hasNetworkServer;

  public static final String SERIALIZED_NAME_OPTION_TYPES = "optionTypes";
  @SerializedName(SERIALIZED_NAME_OPTION_TYPES)
  private List<Object> optionTypes = null;

  public static final String SERIALIZED_NAME_RULE_OPTION_TYPES = "ruleOptionTypes";
  @SerializedName(SERIALIZED_NAME_RULE_OPTION_TYPES)
  private List<Object> ruleOptionTypes = null;


  public InlineResponse20094Type id(Long id) {
    
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


  public InlineResponse20094Type code(String code) {
    
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


  public InlineResponse20094Type name(String name) {
    
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


  public InlineResponse20094Type description(String description) {
    
    this.description = description;
    return this;
  }

   /**
   * Get description
   * @return description
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public String getDescription() {
    return description;
  }


  public void setDescription(String description) {
    this.description = description;
  }


  public InlineResponse20094Type enabled(Boolean enabled) {
    
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


  public InlineResponse20094Type creatable(Boolean creatable) {
    
    this.creatable = creatable;
    return this;
  }

   /**
   * Get creatable
   * @return creatable
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public Boolean getCreatable() {
    return creatable;
  }


  public void setCreatable(Boolean creatable) {
    this.creatable = creatable;
  }


  public InlineResponse20094Type selectable(Boolean selectable) {
    
    this.selectable = selectable;
    return this;
  }

   /**
   * Get selectable
   * @return selectable
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public Boolean getSelectable() {
    return selectable;
  }


  public void setSelectable(Boolean selectable) {
    this.selectable = selectable;
  }


  public InlineResponse20094Type hasFirewall(Boolean hasFirewall) {
    
    this.hasFirewall = hasFirewall;
    return this;
  }

   /**
   * Get hasFirewall
   * @return hasFirewall
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public Boolean getHasFirewall() {
    return hasFirewall;
  }


  public void setHasFirewall(Boolean hasFirewall) {
    this.hasFirewall = hasFirewall;
  }


  public InlineResponse20094Type hasDhcp(Boolean hasDhcp) {
    
    this.hasDhcp = hasDhcp;
    return this;
  }

   /**
   * Get hasDhcp
   * @return hasDhcp
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public Boolean getHasDhcp() {
    return hasDhcp;
  }


  public void setHasDhcp(Boolean hasDhcp) {
    this.hasDhcp = hasDhcp;
  }


  public InlineResponse20094Type hasRouting(Boolean hasRouting) {
    
    this.hasRouting = hasRouting;
    return this;
  }

   /**
   * Get hasRouting
   * @return hasRouting
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public Boolean getHasRouting() {
    return hasRouting;
  }


  public void setHasRouting(Boolean hasRouting) {
    this.hasRouting = hasRouting;
  }


  public InlineResponse20094Type hasNetworkServer(Boolean hasNetworkServer) {
    
    this.hasNetworkServer = hasNetworkServer;
    return this;
  }

   /**
   * Get hasNetworkServer
   * @return hasNetworkServer
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public Boolean getHasNetworkServer() {
    return hasNetworkServer;
  }


  public void setHasNetworkServer(Boolean hasNetworkServer) {
    this.hasNetworkServer = hasNetworkServer;
  }


  public InlineResponse20094Type optionTypes(List<Object> optionTypes) {
    
    this.optionTypes = optionTypes;
    return this;
  }

  public InlineResponse20094Type addOptionTypesItem(Object optionTypesItem) {
    if (this.optionTypes == null) {
      this.optionTypes = new ArrayList<Object>();
    }
    this.optionTypes.add(optionTypesItem);
    return this;
  }

   /**
   * Get optionTypes
   * @return optionTypes
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public List<Object> getOptionTypes() {
    return optionTypes;
  }


  public void setOptionTypes(List<Object> optionTypes) {
    this.optionTypes = optionTypes;
  }


  public InlineResponse20094Type ruleOptionTypes(List<Object> ruleOptionTypes) {
    
    this.ruleOptionTypes = ruleOptionTypes;
    return this;
  }

  public InlineResponse20094Type addRuleOptionTypesItem(Object ruleOptionTypesItem) {
    if (this.ruleOptionTypes == null) {
      this.ruleOptionTypes = new ArrayList<Object>();
    }
    this.ruleOptionTypes.add(ruleOptionTypesItem);
    return this;
  }

   /**
   * Get ruleOptionTypes
   * @return ruleOptionTypes
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public List<Object> getRuleOptionTypes() {
    return ruleOptionTypes;
  }


  public void setRuleOptionTypes(List<Object> ruleOptionTypes) {
    this.ruleOptionTypes = ruleOptionTypes;
  }


  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (o == null || getClass() != o.getClass()) {
      return false;
    }
    InlineResponse20094Type inlineResponse20094Type = (InlineResponse20094Type) o;
    return Objects.equals(this.id, inlineResponse20094Type.id) &&
        Objects.equals(this.code, inlineResponse20094Type.code) &&
        Objects.equals(this.name, inlineResponse20094Type.name) &&
        Objects.equals(this.description, inlineResponse20094Type.description) &&
        Objects.equals(this.enabled, inlineResponse20094Type.enabled) &&
        Objects.equals(this.creatable, inlineResponse20094Type.creatable) &&
        Objects.equals(this.selectable, inlineResponse20094Type.selectable) &&
        Objects.equals(this.hasFirewall, inlineResponse20094Type.hasFirewall) &&
        Objects.equals(this.hasDhcp, inlineResponse20094Type.hasDhcp) &&
        Objects.equals(this.hasRouting, inlineResponse20094Type.hasRouting) &&
        Objects.equals(this.hasNetworkServer, inlineResponse20094Type.hasNetworkServer) &&
        Objects.equals(this.optionTypes, inlineResponse20094Type.optionTypes) &&
        Objects.equals(this.ruleOptionTypes, inlineResponse20094Type.ruleOptionTypes);
  }

  @Override
  public int hashCode() {
    return Objects.hash(id, code, name, description, enabled, creatable, selectable, hasFirewall, hasDhcp, hasRouting, hasNetworkServer, optionTypes, ruleOptionTypes);
  }


  @Override
  public String toString() {
    StringBuilder sb = new StringBuilder();
    sb.append("class InlineResponse20094Type {\n");
    sb.append("    id: ").append(toIndentedString(id)).append("\n");
    sb.append("    code: ").append(toIndentedString(code)).append("\n");
    sb.append("    name: ").append(toIndentedString(name)).append("\n");
    sb.append("    description: ").append(toIndentedString(description)).append("\n");
    sb.append("    enabled: ").append(toIndentedString(enabled)).append("\n");
    sb.append("    creatable: ").append(toIndentedString(creatable)).append("\n");
    sb.append("    selectable: ").append(toIndentedString(selectable)).append("\n");
    sb.append("    hasFirewall: ").append(toIndentedString(hasFirewall)).append("\n");
    sb.append("    hasDhcp: ").append(toIndentedString(hasDhcp)).append("\n");
    sb.append("    hasRouting: ").append(toIndentedString(hasRouting)).append("\n");
    sb.append("    hasNetworkServer: ").append(toIndentedString(hasNetworkServer)).append("\n");
    sb.append("    optionTypes: ").append(toIndentedString(optionTypes)).append("\n");
    sb.append("    ruleOptionTypes: ").append(toIndentedString(ruleOptionTypes)).append("\n");
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

