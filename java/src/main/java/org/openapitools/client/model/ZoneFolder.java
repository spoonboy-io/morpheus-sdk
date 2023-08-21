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
import org.openapitools.client.model.InlineResponse20040AppDeployInstance;
import org.openapitools.client.model.InlineResponse20082LoadBalancerInstanceSslCert;
import org.openapitools.client.model.ResourcePermissions;
import org.openapitools.client.model.ZoneDatastoreTenants;

/**
 * ZoneFolder
 */
@javax.annotation.Generated(value = "org.openapitools.codegen.languages.JavaClientCodegen", date = "2023-08-21T11:26:47.028Z[GMT]")
public class ZoneFolder {
  public static final String SERIALIZED_NAME_ID = "id";
  @SerializedName(SERIALIZED_NAME_ID)
  private Long id;

  public static final String SERIALIZED_NAME_NAME = "name";
  @SerializedName(SERIALIZED_NAME_NAME)
  private String name;

  public static final String SERIALIZED_NAME_ZONE = "zone";
  @SerializedName(SERIALIZED_NAME_ZONE)
  private InlineResponse20040AppDeployInstance zone;

  public static final String SERIALIZED_NAME_PARENT = "parent";
  @SerializedName(SERIALIZED_NAME_PARENT)
  private InlineResponse20082LoadBalancerInstanceSslCert parent;

  public static final String SERIALIZED_NAME_TYPE = "type";
  @SerializedName(SERIALIZED_NAME_TYPE)
  private String type;

  public static final String SERIALIZED_NAME_EXTERNAL_ID = "externalId";
  @SerializedName(SERIALIZED_NAME_EXTERNAL_ID)
  private String externalId;

  public static final String SERIALIZED_NAME_VISIBILITY = "visibility";
  @SerializedName(SERIALIZED_NAME_VISIBILITY)
  private String visibility;

  public static final String SERIALIZED_NAME_READ_ONLY = "readOnly";
  @SerializedName(SERIALIZED_NAME_READ_ONLY)
  private Boolean readOnly;

  public static final String SERIALIZED_NAME_DEFAULT_FOLDER = "defaultFolder";
  @SerializedName(SERIALIZED_NAME_DEFAULT_FOLDER)
  private Boolean defaultFolder;

  public static final String SERIALIZED_NAME_DEFAULT_STORE = "defaultStore";
  @SerializedName(SERIALIZED_NAME_DEFAULT_STORE)
  private Boolean defaultStore;

  public static final String SERIALIZED_NAME_ACTIVE = "active";
  @SerializedName(SERIALIZED_NAME_ACTIVE)
  private Boolean active;

  public static final String SERIALIZED_NAME_TENANTS = "tenants";
  @SerializedName(SERIALIZED_NAME_TENANTS)
  private List<ZoneDatastoreTenants> tenants = null;

  public static final String SERIALIZED_NAME_RESOURCE_PERMISSIONS = "resourcePermissions";
  @SerializedName(SERIALIZED_NAME_RESOURCE_PERMISSIONS)
  private ResourcePermissions resourcePermissions;

  public static final String SERIALIZED_NAME_DEPTH = "depth";
  @SerializedName(SERIALIZED_NAME_DEPTH)
  private Long depth;


  public ZoneFolder id(Long id) {
    
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


  public ZoneFolder name(String name) {
    
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


  public ZoneFolder zone(InlineResponse20040AppDeployInstance zone) {
    
    this.zone = zone;
    return this;
  }

   /**
   * Get zone
   * @return zone
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public InlineResponse20040AppDeployInstance getZone() {
    return zone;
  }


  public void setZone(InlineResponse20040AppDeployInstance zone) {
    this.zone = zone;
  }


  public ZoneFolder parent(InlineResponse20082LoadBalancerInstanceSslCert parent) {
    
    this.parent = parent;
    return this;
  }

   /**
   * Get parent
   * @return parent
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public InlineResponse20082LoadBalancerInstanceSslCert getParent() {
    return parent;
  }


  public void setParent(InlineResponse20082LoadBalancerInstanceSslCert parent) {
    this.parent = parent;
  }


  public ZoneFolder type(String type) {
    
    this.type = type;
    return this;
  }

   /**
   * Get type
   * @return type
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public String getType() {
    return type;
  }


  public void setType(String type) {
    this.type = type;
  }


  public ZoneFolder externalId(String externalId) {
    
    this.externalId = externalId;
    return this;
  }

   /**
   * Get externalId
   * @return externalId
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public String getExternalId() {
    return externalId;
  }


  public void setExternalId(String externalId) {
    this.externalId = externalId;
  }


  public ZoneFolder visibility(String visibility) {
    
    this.visibility = visibility;
    return this;
  }

   /**
   * Get visibility
   * @return visibility
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public String getVisibility() {
    return visibility;
  }


  public void setVisibility(String visibility) {
    this.visibility = visibility;
  }


  public ZoneFolder readOnly(Boolean readOnly) {
    
    this.readOnly = readOnly;
    return this;
  }

   /**
   * Get readOnly
   * @return readOnly
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public Boolean getReadOnly() {
    return readOnly;
  }


  public void setReadOnly(Boolean readOnly) {
    this.readOnly = readOnly;
  }


  public ZoneFolder defaultFolder(Boolean defaultFolder) {
    
    this.defaultFolder = defaultFolder;
    return this;
  }

   /**
   * Get defaultFolder
   * @return defaultFolder
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public Boolean getDefaultFolder() {
    return defaultFolder;
  }


  public void setDefaultFolder(Boolean defaultFolder) {
    this.defaultFolder = defaultFolder;
  }


  public ZoneFolder defaultStore(Boolean defaultStore) {
    
    this.defaultStore = defaultStore;
    return this;
  }

   /**
   * Get defaultStore
   * @return defaultStore
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public Boolean getDefaultStore() {
    return defaultStore;
  }


  public void setDefaultStore(Boolean defaultStore) {
    this.defaultStore = defaultStore;
  }


  public ZoneFolder active(Boolean active) {
    
    this.active = active;
    return this;
  }

   /**
   * Get active
   * @return active
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public Boolean getActive() {
    return active;
  }


  public void setActive(Boolean active) {
    this.active = active;
  }


  public ZoneFolder tenants(List<ZoneDatastoreTenants> tenants) {
    
    this.tenants = tenants;
    return this;
  }

  public ZoneFolder addTenantsItem(ZoneDatastoreTenants tenantsItem) {
    if (this.tenants == null) {
      this.tenants = new ArrayList<ZoneDatastoreTenants>();
    }
    this.tenants.add(tenantsItem);
    return this;
  }

   /**
   * Get tenants
   * @return tenants
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public List<ZoneDatastoreTenants> getTenants() {
    return tenants;
  }


  public void setTenants(List<ZoneDatastoreTenants> tenants) {
    this.tenants = tenants;
  }


  public ZoneFolder resourcePermissions(ResourcePermissions resourcePermissions) {
    
    this.resourcePermissions = resourcePermissions;
    return this;
  }

   /**
   * Get resourcePermissions
   * @return resourcePermissions
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public ResourcePermissions getResourcePermissions() {
    return resourcePermissions;
  }


  public void setResourcePermissions(ResourcePermissions resourcePermissions) {
    this.resourcePermissions = resourcePermissions;
  }


  public ZoneFolder depth(Long depth) {
    
    this.depth = depth;
    return this;
  }

   /**
   * Get depth
   * @return depth
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public Long getDepth() {
    return depth;
  }


  public void setDepth(Long depth) {
    this.depth = depth;
  }


  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (o == null || getClass() != o.getClass()) {
      return false;
    }
    ZoneFolder zoneFolder = (ZoneFolder) o;
    return Objects.equals(this.id, zoneFolder.id) &&
        Objects.equals(this.name, zoneFolder.name) &&
        Objects.equals(this.zone, zoneFolder.zone) &&
        Objects.equals(this.parent, zoneFolder.parent) &&
        Objects.equals(this.type, zoneFolder.type) &&
        Objects.equals(this.externalId, zoneFolder.externalId) &&
        Objects.equals(this.visibility, zoneFolder.visibility) &&
        Objects.equals(this.readOnly, zoneFolder.readOnly) &&
        Objects.equals(this.defaultFolder, zoneFolder.defaultFolder) &&
        Objects.equals(this.defaultStore, zoneFolder.defaultStore) &&
        Objects.equals(this.active, zoneFolder.active) &&
        Objects.equals(this.tenants, zoneFolder.tenants) &&
        Objects.equals(this.resourcePermissions, zoneFolder.resourcePermissions) &&
        Objects.equals(this.depth, zoneFolder.depth);
  }

  @Override
  public int hashCode() {
    return Objects.hash(id, name, zone, parent, type, externalId, visibility, readOnly, defaultFolder, defaultStore, active, tenants, resourcePermissions, depth);
  }


  @Override
  public String toString() {
    StringBuilder sb = new StringBuilder();
    sb.append("class ZoneFolder {\n");
    sb.append("    id: ").append(toIndentedString(id)).append("\n");
    sb.append("    name: ").append(toIndentedString(name)).append("\n");
    sb.append("    zone: ").append(toIndentedString(zone)).append("\n");
    sb.append("    parent: ").append(toIndentedString(parent)).append("\n");
    sb.append("    type: ").append(toIndentedString(type)).append("\n");
    sb.append("    externalId: ").append(toIndentedString(externalId)).append("\n");
    sb.append("    visibility: ").append(toIndentedString(visibility)).append("\n");
    sb.append("    readOnly: ").append(toIndentedString(readOnly)).append("\n");
    sb.append("    defaultFolder: ").append(toIndentedString(defaultFolder)).append("\n");
    sb.append("    defaultStore: ").append(toIndentedString(defaultStore)).append("\n");
    sb.append("    active: ").append(toIndentedString(active)).append("\n");
    sb.append("    tenants: ").append(toIndentedString(tenants)).append("\n");
    sb.append("    resourcePermissions: ").append(toIndentedString(resourcePermissions)).append("\n");
    sb.append("    depth: ").append(toIndentedString(depth)).append("\n");
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

