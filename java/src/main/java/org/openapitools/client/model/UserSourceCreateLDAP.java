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
 * LDAP Configuration
 */
@ApiModel(description = "LDAP Configuration")
@javax.annotation.Generated(value = "org.openapitools.codegen.languages.JavaClientCodegen", date = "2023-08-21T11:26:47.028Z[GMT]")
public class UserSourceCreateLDAP {
  public static final String SERIALIZED_NAME_URL = "url";
  @SerializedName(SERIALIZED_NAME_URL)
  private String url;

  public static final String SERIALIZED_NAME_BINDING_USERNAME = "bindingUsername";
  @SerializedName(SERIALIZED_NAME_BINDING_USERNAME)
  private String bindingUsername;

  public static final String SERIALIZED_NAME_BINDING_PASSWORD = "bindingPassword";
  @SerializedName(SERIALIZED_NAME_BINDING_PASSWORD)
  private String bindingPassword;

  public static final String SERIALIZED_NAME_REQUIRED_GROUP = "requiredGroup";
  @SerializedName(SERIALIZED_NAME_REQUIRED_GROUP)
  private String requiredGroup;


  public UserSourceCreateLDAP url(String url) {
    
    this.url = url;
    return this;
  }

   /**
   * URL of Endpoint
   * @return url
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(example = "my.fqdn.com", value = "URL of Endpoint")

  public String getUrl() {
    return url;
  }


  public void setUrl(String url) {
    this.url = url;
  }


  public UserSourceCreateLDAP bindingUsername(String bindingUsername) {
    
    this.bindingUsername = bindingUsername;
    return this;
  }

   /**
   * Binding Username
   * @return bindingUsername
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(example = "jsmith", value = "Binding Username")

  public String getBindingUsername() {
    return bindingUsername;
  }


  public void setBindingUsername(String bindingUsername) {
    this.bindingUsername = bindingUsername;
  }


  public UserSourceCreateLDAP bindingPassword(String bindingPassword) {
    
    this.bindingPassword = bindingPassword;
    return this;
  }

   /**
   * Binding Password
   * @return bindingPassword
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(example = "MyPassword", value = "Binding Password")

  public String getBindingPassword() {
    return bindingPassword;
  }


  public void setBindingPassword(String bindingPassword) {
    this.bindingPassword = bindingPassword;
  }


  public UserSourceCreateLDAP requiredGroup(String requiredGroup) {
    
    this.requiredGroup = requiredGroup;
    return this;
  }

   /**
   * User DN Expression
   * @return requiredGroup
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(example = "uid=$username,cn=users,cn=account,dc=example,dc=veritas,dc=com", value = "User DN Expression")

  public String getRequiredGroup() {
    return requiredGroup;
  }


  public void setRequiredGroup(String requiredGroup) {
    this.requiredGroup = requiredGroup;
  }


  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (o == null || getClass() != o.getClass()) {
      return false;
    }
    UserSourceCreateLDAP userSourceCreateLDAP = (UserSourceCreateLDAP) o;
    return Objects.equals(this.url, userSourceCreateLDAP.url) &&
        Objects.equals(this.bindingUsername, userSourceCreateLDAP.bindingUsername) &&
        Objects.equals(this.bindingPassword, userSourceCreateLDAP.bindingPassword) &&
        Objects.equals(this.requiredGroup, userSourceCreateLDAP.requiredGroup);
  }

  @Override
  public int hashCode() {
    return Objects.hash(url, bindingUsername, bindingPassword, requiredGroup);
  }


  @Override
  public String toString() {
    StringBuilder sb = new StringBuilder();
    sb.append("class UserSourceCreateLDAP {\n");
    sb.append("    url: ").append(toIndentedString(url)).append("\n");
    sb.append("    bindingUsername: ").append(toIndentedString(bindingUsername)).append("\n");
    sb.append("    bindingPassword: ").append(toIndentedString(bindingPassword)).append("\n");
    sb.append("    requiredGroup: ").append(toIndentedString(requiredGroup)).append("\n");
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
