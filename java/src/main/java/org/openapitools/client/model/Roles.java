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
import org.openapitools.client.model.InlineResponse20040AppDeployInstance;
import org.threeten.bp.OffsetDateTime;

/**
 * Roles
 */
@javax.annotation.Generated(value = "org.openapitools.codegen.languages.JavaClientCodegen", date = "2023-08-21T11:26:47.028Z[GMT]")
public class Roles {
  public static final String SERIALIZED_NAME_ID = "id";
  @SerializedName(SERIALIZED_NAME_ID)
  private Long id;

  public static final String SERIALIZED_NAME_NAME = "name";
  @SerializedName(SERIALIZED_NAME_NAME)
  private String name;

  public static final String SERIALIZED_NAME_AUTHORITY = "authority";
  @SerializedName(SERIALIZED_NAME_AUTHORITY)
  private String authority;

  public static final String SERIALIZED_NAME_DESCRIPTION = "description";
  @SerializedName(SERIALIZED_NAME_DESCRIPTION)
  private String description;

  public static final String SERIALIZED_NAME_SCOPE = "scope";
  @SerializedName(SERIALIZED_NAME_SCOPE)
  private String scope;

  public static final String SERIALIZED_NAME_ROLE_TYPE = "roleType";
  @SerializedName(SERIALIZED_NAME_ROLE_TYPE)
  private String roleType;

  public static final String SERIALIZED_NAME_MULTITENANT = "multitenant";
  @SerializedName(SERIALIZED_NAME_MULTITENANT)
  private Boolean multitenant;

  public static final String SERIALIZED_NAME_MULTITENANT_LOCKED = "multitenantLocked";
  @SerializedName(SERIALIZED_NAME_MULTITENANT_LOCKED)
  private Boolean multitenantLocked;

  public static final String SERIALIZED_NAME_PARENT_ROLE_ID = "parentRoleId";
  @SerializedName(SERIALIZED_NAME_PARENT_ROLE_ID)
  private String parentRoleId;

  public static final String SERIALIZED_NAME_DIVERGED = "diverged";
  @SerializedName(SERIALIZED_NAME_DIVERGED)
  private Boolean diverged;

  public static final String SERIALIZED_NAME_OWNER_ID = "ownerId";
  @SerializedName(SERIALIZED_NAME_OWNER_ID)
  private Long ownerId;

  public static final String SERIALIZED_NAME_OWNER = "owner";
  @SerializedName(SERIALIZED_NAME_OWNER)
  private InlineResponse20040AppDeployInstance owner;

  public static final String SERIALIZED_NAME_DEFAULT_PERSONA = "defaultPersona";
  @SerializedName(SERIALIZED_NAME_DEFAULT_PERSONA)
  private String defaultPersona;

  public static final String SERIALIZED_NAME_DATE_CREATED = "dateCreated";
  @SerializedName(SERIALIZED_NAME_DATE_CREATED)
  private OffsetDateTime dateCreated;

  public static final String SERIALIZED_NAME_LAST_UPDATED = "lastUpdated";
  @SerializedName(SERIALIZED_NAME_LAST_UPDATED)
  private OffsetDateTime lastUpdated;


  public Roles id(Long id) {
    
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


  public Roles name(String name) {
    
    this.name = name;
    return this;
  }

   /**
   * a unique name of the role
   * @return name
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "a unique name of the role")

  public String getName() {
    return name;
  }


  public void setName(String name) {
    this.name = name;
  }


  public Roles authority(String authority) {
    
    this.authority = authority;
    return this;
  }

   /**
   * Alias for name
   * @return authority
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "Alias for name")

  public String getAuthority() {
    return authority;
  }


  public void setAuthority(String authority) {
    this.authority = authority;
  }


  public Roles description(String description) {
    
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


  public Roles scope(String scope) {
    
    this.scope = scope;
    return this;
  }

   /**
   * Get scope
   * @return scope
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public String getScope() {
    return scope;
  }


  public void setScope(String scope) {
    this.scope = scope;
  }


  public Roles roleType(String roleType) {
    
    this.roleType = roleType;
    return this;
  }

   /**
   * Get roleType
   * @return roleType
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public String getRoleType() {
    return roleType;
  }


  public void setRoleType(String roleType) {
    this.roleType = roleType;
  }


  public Roles multitenant(Boolean multitenant) {
    
    this.multitenant = multitenant;
    return this;
  }

   /**
   * Get multitenant
   * @return multitenant
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public Boolean getMultitenant() {
    return multitenant;
  }


  public void setMultitenant(Boolean multitenant) {
    this.multitenant = multitenant;
  }


  public Roles multitenantLocked(Boolean multitenantLocked) {
    
    this.multitenantLocked = multitenantLocked;
    return this;
  }

   /**
   * Get multitenantLocked
   * @return multitenantLocked
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public Boolean getMultitenantLocked() {
    return multitenantLocked;
  }


  public void setMultitenantLocked(Boolean multitenantLocked) {
    this.multitenantLocked = multitenantLocked;
  }


  public Roles parentRoleId(String parentRoleId) {
    
    this.parentRoleId = parentRoleId;
    return this;
  }

   /**
   * Get parentRoleId
   * @return parentRoleId
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public String getParentRoleId() {
    return parentRoleId;
  }


  public void setParentRoleId(String parentRoleId) {
    this.parentRoleId = parentRoleId;
  }


  public Roles diverged(Boolean diverged) {
    
    this.diverged = diverged;
    return this;
  }

   /**
   * Get diverged
   * @return diverged
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public Boolean getDiverged() {
    return diverged;
  }


  public void setDiverged(Boolean diverged) {
    this.diverged = diverged;
  }


  public Roles ownerId(Long ownerId) {
    
    this.ownerId = ownerId;
    return this;
  }

   /**
   * Get ownerId
   * @return ownerId
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public Long getOwnerId() {
    return ownerId;
  }


  public void setOwnerId(Long ownerId) {
    this.ownerId = ownerId;
  }


  public Roles owner(InlineResponse20040AppDeployInstance owner) {
    
    this.owner = owner;
    return this;
  }

   /**
   * Get owner
   * @return owner
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public InlineResponse20040AppDeployInstance getOwner() {
    return owner;
  }


  public void setOwner(InlineResponse20040AppDeployInstance owner) {
    this.owner = owner;
  }


  public Roles defaultPersona(String defaultPersona) {
    
    this.defaultPersona = defaultPersona;
    return this;
  }

   /**
   * Get defaultPersona
   * @return defaultPersona
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public String getDefaultPersona() {
    return defaultPersona;
  }


  public void setDefaultPersona(String defaultPersona) {
    this.defaultPersona = defaultPersona;
  }


  public Roles dateCreated(OffsetDateTime dateCreated) {
    
    this.dateCreated = dateCreated;
    return this;
  }

   /**
   * Get dateCreated
   * @return dateCreated
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public OffsetDateTime getDateCreated() {
    return dateCreated;
  }


  public void setDateCreated(OffsetDateTime dateCreated) {
    this.dateCreated = dateCreated;
  }


  public Roles lastUpdated(OffsetDateTime lastUpdated) {
    
    this.lastUpdated = lastUpdated;
    return this;
  }

   /**
   * Get lastUpdated
   * @return lastUpdated
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public OffsetDateTime getLastUpdated() {
    return lastUpdated;
  }


  public void setLastUpdated(OffsetDateTime lastUpdated) {
    this.lastUpdated = lastUpdated;
  }


  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (o == null || getClass() != o.getClass()) {
      return false;
    }
    Roles roles = (Roles) o;
    return Objects.equals(this.id, roles.id) &&
        Objects.equals(this.name, roles.name) &&
        Objects.equals(this.authority, roles.authority) &&
        Objects.equals(this.description, roles.description) &&
        Objects.equals(this.scope, roles.scope) &&
        Objects.equals(this.roleType, roles.roleType) &&
        Objects.equals(this.multitenant, roles.multitenant) &&
        Objects.equals(this.multitenantLocked, roles.multitenantLocked) &&
        Objects.equals(this.parentRoleId, roles.parentRoleId) &&
        Objects.equals(this.diverged, roles.diverged) &&
        Objects.equals(this.ownerId, roles.ownerId) &&
        Objects.equals(this.owner, roles.owner) &&
        Objects.equals(this.defaultPersona, roles.defaultPersona) &&
        Objects.equals(this.dateCreated, roles.dateCreated) &&
        Objects.equals(this.lastUpdated, roles.lastUpdated);
  }

  @Override
  public int hashCode() {
    return Objects.hash(id, name, authority, description, scope, roleType, multitenant, multitenantLocked, parentRoleId, diverged, ownerId, owner, defaultPersona, dateCreated, lastUpdated);
  }


  @Override
  public String toString() {
    StringBuilder sb = new StringBuilder();
    sb.append("class Roles {\n");
    sb.append("    id: ").append(toIndentedString(id)).append("\n");
    sb.append("    name: ").append(toIndentedString(name)).append("\n");
    sb.append("    authority: ").append(toIndentedString(authority)).append("\n");
    sb.append("    description: ").append(toIndentedString(description)).append("\n");
    sb.append("    scope: ").append(toIndentedString(scope)).append("\n");
    sb.append("    roleType: ").append(toIndentedString(roleType)).append("\n");
    sb.append("    multitenant: ").append(toIndentedString(multitenant)).append("\n");
    sb.append("    multitenantLocked: ").append(toIndentedString(multitenantLocked)).append("\n");
    sb.append("    parentRoleId: ").append(toIndentedString(parentRoleId)).append("\n");
    sb.append("    diverged: ").append(toIndentedString(diverged)).append("\n");
    sb.append("    ownerId: ").append(toIndentedString(ownerId)).append("\n");
    sb.append("    owner: ").append(toIndentedString(owner)).append("\n");
    sb.append("    defaultPersona: ").append(toIndentedString(defaultPersona)).append("\n");
    sb.append("    dateCreated: ").append(toIndentedString(dateCreated)).append("\n");
    sb.append("    lastUpdated: ").append(toIndentedString(lastUpdated)).append("\n");
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

