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
import org.openapitools.client.model.LoadBalancerCreateResourcePermission;
import org.openapitools.client.model.LoadBalancerCreateTenants;

/**
 * LoadBalancerUpdate
 */
@javax.annotation.Generated(value = "org.openapitools.codegen.languages.JavaClientCodegen", date = "2023-08-21T11:26:47.028Z[GMT]")
public class LoadBalancerUpdate {
  public static final String SERIALIZED_NAME_NAME = "name";
  @SerializedName(SERIALIZED_NAME_NAME)
  private String name;

  public static final String SERIALIZED_NAME_DESCRIPTION = "description";
  @SerializedName(SERIALIZED_NAME_DESCRIPTION)
  private String description;

  public static final String SERIALIZED_NAME_ENABLED = "enabled";
  @SerializedName(SERIALIZED_NAME_ENABLED)
  private Boolean enabled;

  public static final String SERIALIZED_NAME_CONFIG = "config";
  @SerializedName(SERIALIZED_NAME_CONFIG)
  private Object config;

  public static final String SERIALIZED_NAME_VISIBILITY = "visibility";
  @SerializedName(SERIALIZED_NAME_VISIBILITY)
  private String visibility = "public";

  public static final String SERIALIZED_NAME_TENANTS = "tenants";
  @SerializedName(SERIALIZED_NAME_TENANTS)
  private List<LoadBalancerCreateTenants> tenants = null;

  public static final String SERIALIZED_NAME_RESOURCE_PERMISSION = "resourcePermission";
  @SerializedName(SERIALIZED_NAME_RESOURCE_PERMISSION)
  private LoadBalancerCreateResourcePermission resourcePermission;


  public LoadBalancerUpdate name(String name) {
    
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


  public LoadBalancerUpdate description(String description) {
    
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


  public LoadBalancerUpdate enabled(Boolean enabled) {
    
    this.enabled = enabled;
    return this;
  }

   /**
   * Activate (true) or disable (false)
   * @return enabled
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "Activate (true) or disable (false)")

  public Boolean getEnabled() {
    return enabled;
  }


  public void setEnabled(Boolean enabled) {
    this.enabled = enabled;
  }


  public LoadBalancerUpdate config(Object config) {
    
    this.config = config;
    return this;
  }

   /**
   * Configuration object with parameters that vary by load balancer type.
   * @return config
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "Configuration object with parameters that vary by load balancer type.")

  public Object getConfig() {
    return config;
  }


  public void setConfig(Object config) {
    this.config = config;
  }


  public LoadBalancerUpdate visibility(String visibility) {
    
    this.visibility = visibility;
    return this;
  }

   /**
   * private or public
   * @return visibility
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "private or public")

  public String getVisibility() {
    return visibility;
  }


  public void setVisibility(String visibility) {
    this.visibility = visibility;
  }


  public LoadBalancerUpdate tenants(List<LoadBalancerCreateTenants> tenants) {
    
    this.tenants = tenants;
    return this;
  }

  public LoadBalancerUpdate addTenantsItem(LoadBalancerCreateTenants tenantsItem) {
    if (this.tenants == null) {
      this.tenants = new ArrayList<LoadBalancerCreateTenants>();
    }
    this.tenants.add(tenantsItem);
    return this;
  }

   /**
   * Array of tenant account ids that are allowed access
   * @return tenants
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "Array of tenant account ids that are allowed access")

  public List<LoadBalancerCreateTenants> getTenants() {
    return tenants;
  }


  public void setTenants(List<LoadBalancerCreateTenants> tenants) {
    this.tenants = tenants;
  }


  public LoadBalancerUpdate resourcePermission(LoadBalancerCreateResourcePermission resourcePermission) {
    
    this.resourcePermission = resourcePermission;
    return this;
  }

   /**
   * Get resourcePermission
   * @return resourcePermission
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public LoadBalancerCreateResourcePermission getResourcePermission() {
    return resourcePermission;
  }


  public void setResourcePermission(LoadBalancerCreateResourcePermission resourcePermission) {
    this.resourcePermission = resourcePermission;
  }


  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (o == null || getClass() != o.getClass()) {
      return false;
    }
    LoadBalancerUpdate loadBalancerUpdate = (LoadBalancerUpdate) o;
    return Objects.equals(this.name, loadBalancerUpdate.name) &&
        Objects.equals(this.description, loadBalancerUpdate.description) &&
        Objects.equals(this.enabled, loadBalancerUpdate.enabled) &&
        Objects.equals(this.config, loadBalancerUpdate.config) &&
        Objects.equals(this.visibility, loadBalancerUpdate.visibility) &&
        Objects.equals(this.tenants, loadBalancerUpdate.tenants) &&
        Objects.equals(this.resourcePermission, loadBalancerUpdate.resourcePermission);
  }

  @Override
  public int hashCode() {
    return Objects.hash(name, description, enabled, config, visibility, tenants, resourcePermission);
  }


  @Override
  public String toString() {
    StringBuilder sb = new StringBuilder();
    sb.append("class LoadBalancerUpdate {\n");
    sb.append("    name: ").append(toIndentedString(name)).append("\n");
    sb.append("    description: ").append(toIndentedString(description)).append("\n");
    sb.append("    enabled: ").append(toIndentedString(enabled)).append("\n");
    sb.append("    config: ").append(toIndentedString(config)).append("\n");
    sb.append("    visibility: ").append(toIndentedString(visibility)).append("\n");
    sb.append("    tenants: ").append(toIndentedString(tenants)).append("\n");
    sb.append("    resourcePermission: ").append(toIndentedString(resourcePermission)).append("\n");
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

