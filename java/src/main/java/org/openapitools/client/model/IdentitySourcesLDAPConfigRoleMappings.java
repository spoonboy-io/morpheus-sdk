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
import org.openapitools.client.model.IdentitySourcesLDAPConfigDefaultAccountRole;

/**
 * IdentitySourcesLDAPConfigRoleMappings
 */
@javax.annotation.Generated(value = "org.openapitools.codegen.languages.JavaClientCodegen", date = "2023-08-21T11:26:47.028Z[GMT]")
public class IdentitySourcesLDAPConfigRoleMappings {
  public static final String SERIALIZED_NAME_SOURCE_ROLE_NAME = "sourceRoleName";
  @SerializedName(SERIALIZED_NAME_SOURCE_ROLE_NAME)
  private String sourceRoleName;

  public static final String SERIALIZED_NAME_SOURCE_ROLE_FQN = "sourceRoleFqn";
  @SerializedName(SERIALIZED_NAME_SOURCE_ROLE_FQN)
  private String sourceRoleFqn;

  public static final String SERIALIZED_NAME_MAPPED_ROLE = "mappedRole";
  @SerializedName(SERIALIZED_NAME_MAPPED_ROLE)
  private IdentitySourcesLDAPConfigDefaultAccountRole mappedRole;


  public IdentitySourcesLDAPConfigRoleMappings sourceRoleName(String sourceRoleName) {
    
    this.sourceRoleName = sourceRoleName;
    return this;
  }

   /**
   * Get sourceRoleName
   * @return sourceRoleName
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public String getSourceRoleName() {
    return sourceRoleName;
  }


  public void setSourceRoleName(String sourceRoleName) {
    this.sourceRoleName = sourceRoleName;
  }


  public IdentitySourcesLDAPConfigRoleMappings sourceRoleFqn(String sourceRoleFqn) {
    
    this.sourceRoleFqn = sourceRoleFqn;
    return this;
  }

   /**
   * Get sourceRoleFqn
   * @return sourceRoleFqn
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public String getSourceRoleFqn() {
    return sourceRoleFqn;
  }


  public void setSourceRoleFqn(String sourceRoleFqn) {
    this.sourceRoleFqn = sourceRoleFqn;
  }


  public IdentitySourcesLDAPConfigRoleMappings mappedRole(IdentitySourcesLDAPConfigDefaultAccountRole mappedRole) {
    
    this.mappedRole = mappedRole;
    return this;
  }

   /**
   * Get mappedRole
   * @return mappedRole
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public IdentitySourcesLDAPConfigDefaultAccountRole getMappedRole() {
    return mappedRole;
  }


  public void setMappedRole(IdentitySourcesLDAPConfigDefaultAccountRole mappedRole) {
    this.mappedRole = mappedRole;
  }


  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (o == null || getClass() != o.getClass()) {
      return false;
    }
    IdentitySourcesLDAPConfigRoleMappings identitySourcesLDAPConfigRoleMappings = (IdentitySourcesLDAPConfigRoleMappings) o;
    return Objects.equals(this.sourceRoleName, identitySourcesLDAPConfigRoleMappings.sourceRoleName) &&
        Objects.equals(this.sourceRoleFqn, identitySourcesLDAPConfigRoleMappings.sourceRoleFqn) &&
        Objects.equals(this.mappedRole, identitySourcesLDAPConfigRoleMappings.mappedRole);
  }

  @Override
  public int hashCode() {
    return Objects.hash(sourceRoleName, sourceRoleFqn, mappedRole);
  }


  @Override
  public String toString() {
    StringBuilder sb = new StringBuilder();
    sb.append("class IdentitySourcesLDAPConfigRoleMappings {\n");
    sb.append("    sourceRoleName: ").append(toIndentedString(sourceRoleName)).append("\n");
    sb.append("    sourceRoleFqn: ").append(toIndentedString(sourceRoleFqn)).append("\n");
    sb.append("    mappedRole: ").append(toIndentedString(mappedRole)).append("\n");
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

