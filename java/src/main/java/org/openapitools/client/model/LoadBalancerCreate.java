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
 * LoadBalancerCreate
 */
@javax.annotation.Generated(value = "org.openapitools.codegen.languages.JavaClientCodegen", date = "2023-08-21T11:26:47.028Z[GMT]")
public class LoadBalancerCreate {
  public static final String SERIALIZED_NAME_NAME = "name";
  @SerializedName(SERIALIZED_NAME_NAME)
  private String name;

  public static final String SERIALIZED_NAME_DESCRIPTION = "description";
  @SerializedName(SERIALIZED_NAME_DESCRIPTION)
  private String description;

  public static final String SERIALIZED_NAME_NETWORK_SERVER_ID = "networkServerId";
  @SerializedName(SERIALIZED_NAME_NETWORK_SERVER_ID)
  private Long networkServerId;

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


  public LoadBalancerCreate name(String name) {
    
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


  public LoadBalancerCreate description(String description) {
    
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


  public LoadBalancerCreate networkServerId(Long networkServerId) {
    
    this.networkServerId = networkServerId;
    return this;
  }

   /**
   * Network Server ID
   * @return networkServerId
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "Network Server ID")

  public Long getNetworkServerId() {
    return networkServerId;
  }


  public void setNetworkServerId(Long networkServerId) {
    this.networkServerId = networkServerId;
  }


  public LoadBalancerCreate config(Object config) {
    
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


  public LoadBalancerCreate visibility(String visibility) {
    
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


  public LoadBalancerCreate tenants(List<LoadBalancerCreateTenants> tenants) {
    
    this.tenants = tenants;
    return this;
  }

  public LoadBalancerCreate addTenantsItem(LoadBalancerCreateTenants tenantsItem) {
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


  public LoadBalancerCreate resourcePermission(LoadBalancerCreateResourcePermission resourcePermission) {
    
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
    LoadBalancerCreate loadBalancerCreate = (LoadBalancerCreate) o;
    return Objects.equals(this.name, loadBalancerCreate.name) &&
        Objects.equals(this.description, loadBalancerCreate.description) &&
        Objects.equals(this.networkServerId, loadBalancerCreate.networkServerId) &&
        Objects.equals(this.config, loadBalancerCreate.config) &&
        Objects.equals(this.visibility, loadBalancerCreate.visibility) &&
        Objects.equals(this.tenants, loadBalancerCreate.tenants) &&
        Objects.equals(this.resourcePermission, loadBalancerCreate.resourcePermission);
  }

  @Override
  public int hashCode() {
    return Objects.hash(name, description, networkServerId, config, visibility, tenants, resourcePermission);
  }


  @Override
  public String toString() {
    StringBuilder sb = new StringBuilder();
    sb.append("class LoadBalancerCreate {\n");
    sb.append("    name: ").append(toIndentedString(name)).append("\n");
    sb.append("    description: ").append(toIndentedString(description)).append("\n");
    sb.append("    networkServerId: ").append(toIndentedString(networkServerId)).append("\n");
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

