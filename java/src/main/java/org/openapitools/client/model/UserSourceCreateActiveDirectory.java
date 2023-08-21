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
 * Active Directory Configuration
 */
@ApiModel(description = "Active Directory Configuration")
@javax.annotation.Generated(value = "org.openapitools.codegen.languages.JavaClientCodegen", date = "2023-08-21T11:26:47.028Z[GMT]")
public class UserSourceCreateActiveDirectory {
  public static final String SERIALIZED_NAME_URL = "url";
  @SerializedName(SERIALIZED_NAME_URL)
  private String url;

  public static final String SERIALIZED_NAME_DOMAIN = "domain";
  @SerializedName(SERIALIZED_NAME_DOMAIN)
  private String domain;

  public static final String SERIALIZED_NAME_USE_S_S_L = "useSSL";
  @SerializedName(SERIALIZED_NAME_USE_S_S_L)
  private Boolean useSSL = false;

  public static final String SERIALIZED_NAME_BINDING_USERNAME = "bindingUsername";
  @SerializedName(SERIALIZED_NAME_BINDING_USERNAME)
  private String bindingUsername;

  public static final String SERIALIZED_NAME_BINDING_PASSWORD = "bindingPassword";
  @SerializedName(SERIALIZED_NAME_BINDING_PASSWORD)
  private String bindingPassword;

  public static final String SERIALIZED_NAME_REQUIRED_GROUP = "requiredGroup";
  @SerializedName(SERIALIZED_NAME_REQUIRED_GROUP)
  private String requiredGroup;

  public static final String SERIALIZED_NAME_SEARCH_MEMBER_GROUPS = "searchMemberGroups";
  @SerializedName(SERIALIZED_NAME_SEARCH_MEMBER_GROUPS)
  private Boolean searchMemberGroups = false;


  public UserSourceCreateActiveDirectory url(String url) {
    
    this.url = url;
    return this;
  }

   /**
   * AD Server IP/FQDN
   * @return url
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(example = "my.fqdn.com", value = "AD Server IP/FQDN")

  public String getUrl() {
    return url;
  }


  public void setUrl(String url) {
    this.url = url;
  }


  public UserSourceCreateActiveDirectory domain(String domain) {
    
    this.domain = domain;
    return this;
  }

   /**
   * Domain
   * @return domain
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(example = "fqdn.com", value = "Domain")

  public String getDomain() {
    return domain;
  }


  public void setDomain(String domain) {
    this.domain = domain;
  }


  public UserSourceCreateActiveDirectory useSSL(Boolean useSSL) {
    
    this.useSSL = useSSL;
    return this;
  }

   /**
   * Use SSL
   * @return useSSL
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "Use SSL")

  public Boolean getUseSSL() {
    return useSSL;
  }


  public void setUseSSL(Boolean useSSL) {
    this.useSSL = useSSL;
  }


  public UserSourceCreateActiveDirectory bindingUsername(String bindingUsername) {
    
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


  public UserSourceCreateActiveDirectory bindingPassword(String bindingPassword) {
    
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


  public UserSourceCreateActiveDirectory requiredGroup(String requiredGroup) {
    
    this.requiredGroup = requiredGroup;
    return this;
  }

   /**
   * Required Group
   * @return requiredGroup
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(example = "Administrators", value = "Required Group")

  public String getRequiredGroup() {
    return requiredGroup;
  }


  public void setRequiredGroup(String requiredGroup) {
    this.requiredGroup = requiredGroup;
  }


  public UserSourceCreateActiveDirectory searchMemberGroups(Boolean searchMemberGroups) {
    
    this.searchMemberGroups = searchMemberGroups;
    return this;
  }

   /**
   * Include Member Groups
   * @return searchMemberGroups
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "Include Member Groups")

  public Boolean getSearchMemberGroups() {
    return searchMemberGroups;
  }


  public void setSearchMemberGroups(Boolean searchMemberGroups) {
    this.searchMemberGroups = searchMemberGroups;
  }


  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (o == null || getClass() != o.getClass()) {
      return false;
    }
    UserSourceCreateActiveDirectory userSourceCreateActiveDirectory = (UserSourceCreateActiveDirectory) o;
    return Objects.equals(this.url, userSourceCreateActiveDirectory.url) &&
        Objects.equals(this.domain, userSourceCreateActiveDirectory.domain) &&
        Objects.equals(this.useSSL, userSourceCreateActiveDirectory.useSSL) &&
        Objects.equals(this.bindingUsername, userSourceCreateActiveDirectory.bindingUsername) &&
        Objects.equals(this.bindingPassword, userSourceCreateActiveDirectory.bindingPassword) &&
        Objects.equals(this.requiredGroup, userSourceCreateActiveDirectory.requiredGroup) &&
        Objects.equals(this.searchMemberGroups, userSourceCreateActiveDirectory.searchMemberGroups);
  }

  @Override
  public int hashCode() {
    return Objects.hash(url, domain, useSSL, bindingUsername, bindingPassword, requiredGroup, searchMemberGroups);
  }


  @Override
  public String toString() {
    StringBuilder sb = new StringBuilder();
    sb.append("class UserSourceCreateActiveDirectory {\n");
    sb.append("    url: ").append(toIndentedString(url)).append("\n");
    sb.append("    domain: ").append(toIndentedString(domain)).append("\n");
    sb.append("    useSSL: ").append(toIndentedString(useSSL)).append("\n");
    sb.append("    bindingUsername: ").append(toIndentedString(bindingUsername)).append("\n");
    sb.append("    bindingPassword: ").append(toIndentedString(bindingPassword)).append("\n");
    sb.append("    requiredGroup: ").append(toIndentedString(requiredGroup)).append("\n");
    sb.append("    searchMemberGroups: ").append(toIndentedString(searchMemberGroups)).append("\n");
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

