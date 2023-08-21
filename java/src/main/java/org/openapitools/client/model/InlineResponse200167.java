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
import org.openapitools.client.model.InlineResponse200167Appliance;
import org.openapitools.client.model.InlineResponse200167Permissions;
import org.openapitools.client.model.User;

/**
 * InlineResponse200167
 */
@javax.annotation.Generated(value = "org.openapitools.codegen.languages.JavaClientCodegen", date = "2023-08-21T11:26:47.028Z[GMT]")
public class InlineResponse200167 {
  public static final String SERIALIZED_NAME_USER = "user";
  @SerializedName(SERIALIZED_NAME_USER)
  private User user;

  public static final String SERIALIZED_NAME_IS_MASTER_ACCOUNT = "isMasterAccount";
  @SerializedName(SERIALIZED_NAME_IS_MASTER_ACCOUNT)
  private Boolean isMasterAccount;

  public static final String SERIALIZED_NAME_PERMISSIONS = "permissions";
  @SerializedName(SERIALIZED_NAME_PERMISSIONS)
  private List<InlineResponse200167Permissions> permissions = null;

  public static final String SERIALIZED_NAME_APPLIANCE = "appliance";
  @SerializedName(SERIALIZED_NAME_APPLIANCE)
  private InlineResponse200167Appliance appliance;


  public InlineResponse200167 user(User user) {
    
    this.user = user;
    return this;
  }

   /**
   * Get user
   * @return user
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public User getUser() {
    return user;
  }


  public void setUser(User user) {
    this.user = user;
  }


  public InlineResponse200167 isMasterAccount(Boolean isMasterAccount) {
    
    this.isMasterAccount = isMasterAccount;
    return this;
  }

   /**
   * Get isMasterAccount
   * @return isMasterAccount
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public Boolean getIsMasterAccount() {
    return isMasterAccount;
  }


  public void setIsMasterAccount(Boolean isMasterAccount) {
    this.isMasterAccount = isMasterAccount;
  }


  public InlineResponse200167 permissions(List<InlineResponse200167Permissions> permissions) {
    
    this.permissions = permissions;
    return this;
  }

  public InlineResponse200167 addPermissionsItem(InlineResponse200167Permissions permissionsItem) {
    if (this.permissions == null) {
      this.permissions = new ArrayList<InlineResponse200167Permissions>();
    }
    this.permissions.add(permissionsItem);
    return this;
  }

   /**
   * Get permissions
   * @return permissions
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public List<InlineResponse200167Permissions> getPermissions() {
    return permissions;
  }


  public void setPermissions(List<InlineResponse200167Permissions> permissions) {
    this.permissions = permissions;
  }


  public InlineResponse200167 appliance(InlineResponse200167Appliance appliance) {
    
    this.appliance = appliance;
    return this;
  }

   /**
   * Get appliance
   * @return appliance
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public InlineResponse200167Appliance getAppliance() {
    return appliance;
  }


  public void setAppliance(InlineResponse200167Appliance appliance) {
    this.appliance = appliance;
  }


  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (o == null || getClass() != o.getClass()) {
      return false;
    }
    InlineResponse200167 inlineResponse200167 = (InlineResponse200167) o;
    return Objects.equals(this.user, inlineResponse200167.user) &&
        Objects.equals(this.isMasterAccount, inlineResponse200167.isMasterAccount) &&
        Objects.equals(this.permissions, inlineResponse200167.permissions) &&
        Objects.equals(this.appliance, inlineResponse200167.appliance);
  }

  @Override
  public int hashCode() {
    return Objects.hash(user, isMasterAccount, permissions, appliance);
  }


  @Override
  public String toString() {
    StringBuilder sb = new StringBuilder();
    sb.append("class InlineResponse200167 {\n");
    sb.append("    user: ").append(toIndentedString(user)).append("\n");
    sb.append("    isMasterAccount: ").append(toIndentedString(isMasterAccount)).append("\n");
    sb.append("    permissions: ").append(toIndentedString(permissions)).append("\n");
    sb.append("    appliance: ").append(toIndentedString(appliance)).append("\n");
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

