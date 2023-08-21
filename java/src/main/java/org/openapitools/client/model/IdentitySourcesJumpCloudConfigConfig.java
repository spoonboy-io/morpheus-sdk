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
 * IdentitySourcesJumpCloudConfigConfig
 */
@javax.annotation.Generated(value = "org.openapitools.codegen.languages.JavaClientCodegen", date = "2023-08-21T11:26:47.028Z[GMT]")
public class IdentitySourcesJumpCloudConfigConfig {
  public static final String SERIALIZED_NAME_ORGANIZATION_ID = "organizationId";
  @SerializedName(SERIALIZED_NAME_ORGANIZATION_ID)
  private String organizationId;

  public static final String SERIALIZED_NAME_BINDING_USERNAME = "bindingUsername";
  @SerializedName(SERIALIZED_NAME_BINDING_USERNAME)
  private String bindingUsername;

  public static final String SERIALIZED_NAME_BINDING_PASSWORD = "bindingPassword";
  @SerializedName(SERIALIZED_NAME_BINDING_PASSWORD)
  private String bindingPassword;

  public static final String SERIALIZED_NAME_REQUIRED_ROLE = "requiredRole";
  @SerializedName(SERIALIZED_NAME_REQUIRED_ROLE)
  private String requiredRole;

  public static final String SERIALIZED_NAME_BINDING_PASSWORD_HASH = "bindingPasswordHash";
  @SerializedName(SERIALIZED_NAME_BINDING_PASSWORD_HASH)
  private String bindingPasswordHash;


  public IdentitySourcesJumpCloudConfigConfig organizationId(String organizationId) {
    
    this.organizationId = organizationId;
    return this;
  }

   /**
   * Get organizationId
   * @return organizationId
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public String getOrganizationId() {
    return organizationId;
  }


  public void setOrganizationId(String organizationId) {
    this.organizationId = organizationId;
  }


  public IdentitySourcesJumpCloudConfigConfig bindingUsername(String bindingUsername) {
    
    this.bindingUsername = bindingUsername;
    return this;
  }

   /**
   * Get bindingUsername
   * @return bindingUsername
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public String getBindingUsername() {
    return bindingUsername;
  }


  public void setBindingUsername(String bindingUsername) {
    this.bindingUsername = bindingUsername;
  }


  public IdentitySourcesJumpCloudConfigConfig bindingPassword(String bindingPassword) {
    
    this.bindingPassword = bindingPassword;
    return this;
  }

   /**
   * Get bindingPassword
   * @return bindingPassword
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public String getBindingPassword() {
    return bindingPassword;
  }


  public void setBindingPassword(String bindingPassword) {
    this.bindingPassword = bindingPassword;
  }


  public IdentitySourcesJumpCloudConfigConfig requiredRole(String requiredRole) {
    
    this.requiredRole = requiredRole;
    return this;
  }

   /**
   * Get requiredRole
   * @return requiredRole
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public String getRequiredRole() {
    return requiredRole;
  }


  public void setRequiredRole(String requiredRole) {
    this.requiredRole = requiredRole;
  }


  public IdentitySourcesJumpCloudConfigConfig bindingPasswordHash(String bindingPasswordHash) {
    
    this.bindingPasswordHash = bindingPasswordHash;
    return this;
  }

   /**
   * Get bindingPasswordHash
   * @return bindingPasswordHash
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public String getBindingPasswordHash() {
    return bindingPasswordHash;
  }


  public void setBindingPasswordHash(String bindingPasswordHash) {
    this.bindingPasswordHash = bindingPasswordHash;
  }


  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (o == null || getClass() != o.getClass()) {
      return false;
    }
    IdentitySourcesJumpCloudConfigConfig identitySourcesJumpCloudConfigConfig = (IdentitySourcesJumpCloudConfigConfig) o;
    return Objects.equals(this.organizationId, identitySourcesJumpCloudConfigConfig.organizationId) &&
        Objects.equals(this.bindingUsername, identitySourcesJumpCloudConfigConfig.bindingUsername) &&
        Objects.equals(this.bindingPassword, identitySourcesJumpCloudConfigConfig.bindingPassword) &&
        Objects.equals(this.requiredRole, identitySourcesJumpCloudConfigConfig.requiredRole) &&
        Objects.equals(this.bindingPasswordHash, identitySourcesJumpCloudConfigConfig.bindingPasswordHash);
  }

  @Override
  public int hashCode() {
    return Objects.hash(organizationId, bindingUsername, bindingPassword, requiredRole, bindingPasswordHash);
  }


  @Override
  public String toString() {
    StringBuilder sb = new StringBuilder();
    sb.append("class IdentitySourcesJumpCloudConfigConfig {\n");
    sb.append("    organizationId: ").append(toIndentedString(organizationId)).append("\n");
    sb.append("    bindingUsername: ").append(toIndentedString(bindingUsername)).append("\n");
    sb.append("    bindingPassword: ").append(toIndentedString(bindingPassword)).append("\n");
    sb.append("    requiredRole: ").append(toIndentedString(requiredRole)).append("\n");
    sb.append("    bindingPasswordHash: ").append(toIndentedString(bindingPasswordHash)).append("\n");
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

