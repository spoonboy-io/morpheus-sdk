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
import org.openapitools.client.model.ApiBlueprintsIdUpdatePermissionsResourcePermissionSites;

/**
 * NetworkTransportZoneCreate
 */
@javax.annotation.Generated(value = "org.openapitools.codegen.languages.JavaClientCodegen", date = "2023-08-21T11:26:47.028Z[GMT]")
public class NetworkTransportZoneCreate {
  public static final String SERIALIZED_NAME_NAME = "name";
  @SerializedName(SERIALIZED_NAME_NAME)
  private String name;

  public static final String SERIALIZED_NAME_DESCRIPTION = "description";
  @SerializedName(SERIALIZED_NAME_DESCRIPTION)
  private String description;

  public static final String SERIALIZED_NAME_VISIBILITY = "visibility";
  @SerializedName(SERIALIZED_NAME_VISIBILITY)
  private String visibility;

  public static final String SERIALIZED_NAME_TENANTS = "tenants";
  @SerializedName(SERIALIZED_NAME_TENANTS)
  private List<ApiBlueprintsIdUpdatePermissionsResourcePermissionSites> tenants = null;


  public NetworkTransportZoneCreate name(String name) {
    
    this.name = name;
    return this;
  }

   /**
   * Network transport zone name
   * @return name
  **/
  @ApiModelProperty(required = true, value = "Network transport zone name")

  public String getName() {
    return name;
  }


  public void setName(String name) {
    this.name = name;
  }


  public NetworkTransportZoneCreate description(String description) {
    
    this.description = description;
    return this;
  }

   /**
   * Network transport zone description
   * @return description
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "Network transport zone description")

  public String getDescription() {
    return description;
  }


  public void setDescription(String description) {
    this.description = description;
  }


  public NetworkTransportZoneCreate visibility(String visibility) {
    
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


  public NetworkTransportZoneCreate tenants(List<ApiBlueprintsIdUpdatePermissionsResourcePermissionSites> tenants) {
    
    this.tenants = tenants;
    return this;
  }

  public NetworkTransportZoneCreate addTenantsItem(ApiBlueprintsIdUpdatePermissionsResourcePermissionSites tenantsItem) {
    if (this.tenants == null) {
      this.tenants = new ArrayList<ApiBlueprintsIdUpdatePermissionsResourcePermissionSites>();
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

  public List<ApiBlueprintsIdUpdatePermissionsResourcePermissionSites> getTenants() {
    return tenants;
  }


  public void setTenants(List<ApiBlueprintsIdUpdatePermissionsResourcePermissionSites> tenants) {
    this.tenants = tenants;
  }


  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (o == null || getClass() != o.getClass()) {
      return false;
    }
    NetworkTransportZoneCreate networkTransportZoneCreate = (NetworkTransportZoneCreate) o;
    return Objects.equals(this.name, networkTransportZoneCreate.name) &&
        Objects.equals(this.description, networkTransportZoneCreate.description) &&
        Objects.equals(this.visibility, networkTransportZoneCreate.visibility) &&
        Objects.equals(this.tenants, networkTransportZoneCreate.tenants);
  }

  @Override
  public int hashCode() {
    return Objects.hash(name, description, visibility, tenants);
  }


  @Override
  public String toString() {
    StringBuilder sb = new StringBuilder();
    sb.append("class NetworkTransportZoneCreate {\n");
    sb.append("    name: ").append(toIndentedString(name)).append("\n");
    sb.append("    description: ").append(toIndentedString(description)).append("\n");
    sb.append("    visibility: ").append(toIndentedString(visibility)).append("\n");
    sb.append("    tenants: ").append(toIndentedString(tenants)).append("\n");
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

