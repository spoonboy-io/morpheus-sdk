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
 * IdentitySourcesADConfigConfig
 */
@javax.annotation.Generated(value = "org.openapitools.codegen.languages.JavaClientCodegen", date = "2023-08-21T11:26:47.028Z[GMT]")
public class IdentitySourcesADConfigConfig {
  public static final String SERIALIZED_NAME_URL = "url";
  @SerializedName(SERIALIZED_NAME_URL)
  private String url;

  public static final String SERIALIZED_NAME_DOMAIN = "domain";
  @SerializedName(SERIALIZED_NAME_DOMAIN)
  private String domain;

  public static final String SERIALIZED_NAME_USE_S_S_L = "useSSL";
  @SerializedName(SERIALIZED_NAME_USE_S_S_L)
  private String useSSL;

  public static final String SERIALIZED_NAME_BINDING_USERNAME = "bindingUsername";
  @SerializedName(SERIALIZED_NAME_BINDING_USERNAME)
  private String bindingUsername;

  public static final String SERIALIZED_NAME_BINDING_PASSWORD = "bindingPassword";
  @SerializedName(SERIALIZED_NAME_BINDING_PASSWORD)
  private String bindingPassword;

  public static final String SERIALIZED_NAME_REQUIRED_GROUP = "requiredGroup";
  @SerializedName(SERIALIZED_NAME_REQUIRED_GROUP)
  private String requiredGroup;

  public static final String SERIALIZED_NAME_REQUIRED_GROUP_D_N = "requiredGroupDN";
  @SerializedName(SERIALIZED_NAME_REQUIRED_GROUP_D_N)
  private String requiredGroupDN;

  public static final String SERIALIZED_NAME_SEARCH_MEMBER_GROUPS = "searchMemberGroups";
  @SerializedName(SERIALIZED_NAME_SEARCH_MEMBER_GROUPS)
  private Boolean searchMemberGroups;

  public static final String SERIALIZED_NAME_BINDING_PASSWORD_HASH = "bindingPasswordHash";
  @SerializedName(SERIALIZED_NAME_BINDING_PASSWORD_HASH)
  private String bindingPasswordHash;


  public IdentitySourcesADConfigConfig url(String url) {
    
    this.url = url;
    return this;
  }

   /**
   * Get url
   * @return url
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public String getUrl() {
    return url;
  }


  public void setUrl(String url) {
    this.url = url;
  }


  public IdentitySourcesADConfigConfig domain(String domain) {
    
    this.domain = domain;
    return this;
  }

   /**
   * Get domain
   * @return domain
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public String getDomain() {
    return domain;
  }


  public void setDomain(String domain) {
    this.domain = domain;
  }


  public IdentitySourcesADConfigConfig useSSL(String useSSL) {
    
    this.useSSL = useSSL;
    return this;
  }

   /**
   * Get useSSL
   * @return useSSL
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public String getUseSSL() {
    return useSSL;
  }


  public void setUseSSL(String useSSL) {
    this.useSSL = useSSL;
  }


  public IdentitySourcesADConfigConfig bindingUsername(String bindingUsername) {
    
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


  public IdentitySourcesADConfigConfig bindingPassword(String bindingPassword) {
    
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


  public IdentitySourcesADConfigConfig requiredGroup(String requiredGroup) {
    
    this.requiredGroup = requiredGroup;
    return this;
  }

   /**
   * Get requiredGroup
   * @return requiredGroup
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public String getRequiredGroup() {
    return requiredGroup;
  }


  public void setRequiredGroup(String requiredGroup) {
    this.requiredGroup = requiredGroup;
  }


  public IdentitySourcesADConfigConfig requiredGroupDN(String requiredGroupDN) {
    
    this.requiredGroupDN = requiredGroupDN;
    return this;
  }

   /**
   * Get requiredGroupDN
   * @return requiredGroupDN
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public String getRequiredGroupDN() {
    return requiredGroupDN;
  }


  public void setRequiredGroupDN(String requiredGroupDN) {
    this.requiredGroupDN = requiredGroupDN;
  }


  public IdentitySourcesADConfigConfig searchMemberGroups(Boolean searchMemberGroups) {
    
    this.searchMemberGroups = searchMemberGroups;
    return this;
  }

   /**
   * Get searchMemberGroups
   * @return searchMemberGroups
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public Boolean getSearchMemberGroups() {
    return searchMemberGroups;
  }


  public void setSearchMemberGroups(Boolean searchMemberGroups) {
    this.searchMemberGroups = searchMemberGroups;
  }


  public IdentitySourcesADConfigConfig bindingPasswordHash(String bindingPasswordHash) {
    
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
    IdentitySourcesADConfigConfig identitySourcesADConfigConfig = (IdentitySourcesADConfigConfig) o;
    return Objects.equals(this.url, identitySourcesADConfigConfig.url) &&
        Objects.equals(this.domain, identitySourcesADConfigConfig.domain) &&
        Objects.equals(this.useSSL, identitySourcesADConfigConfig.useSSL) &&
        Objects.equals(this.bindingUsername, identitySourcesADConfigConfig.bindingUsername) &&
        Objects.equals(this.bindingPassword, identitySourcesADConfigConfig.bindingPassword) &&
        Objects.equals(this.requiredGroup, identitySourcesADConfigConfig.requiredGroup) &&
        Objects.equals(this.requiredGroupDN, identitySourcesADConfigConfig.requiredGroupDN) &&
        Objects.equals(this.searchMemberGroups, identitySourcesADConfigConfig.searchMemberGroups) &&
        Objects.equals(this.bindingPasswordHash, identitySourcesADConfigConfig.bindingPasswordHash);
  }

  @Override
  public int hashCode() {
    return Objects.hash(url, domain, useSSL, bindingUsername, bindingPassword, requiredGroup, requiredGroupDN, searchMemberGroups, bindingPasswordHash);
  }


  @Override
  public String toString() {
    StringBuilder sb = new StringBuilder();
    sb.append("class IdentitySourcesADConfigConfig {\n");
    sb.append("    url: ").append(toIndentedString(url)).append("\n");
    sb.append("    domain: ").append(toIndentedString(domain)).append("\n");
    sb.append("    useSSL: ").append(toIndentedString(useSSL)).append("\n");
    sb.append("    bindingUsername: ").append(toIndentedString(bindingUsername)).append("\n");
    sb.append("    bindingPassword: ").append(toIndentedString(bindingPassword)).append("\n");
    sb.append("    requiredGroup: ").append(toIndentedString(requiredGroup)).append("\n");
    sb.append("    requiredGroupDN: ").append(toIndentedString(requiredGroupDN)).append("\n");
    sb.append("    searchMemberGroups: ").append(toIndentedString(searchMemberGroups)).append("\n");
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
