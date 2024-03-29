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
 * IdentitySourcesOktaConfigConfig
 */
@javax.annotation.Generated(value = "org.openapitools.codegen.languages.JavaClientCodegen", date = "2023-08-21T11:26:47.028Z[GMT]")
public class IdentitySourcesOktaConfigConfig {
  public static final String SERIALIZED_NAME_URL = "url";
  @SerializedName(SERIALIZED_NAME_URL)
  private String url;

  public static final String SERIALIZED_NAME_ADMINISTRATOR_A_P_I_TOKEN = "administratorAPIToken";
  @SerializedName(SERIALIZED_NAME_ADMINISTRATOR_A_P_I_TOKEN)
  private String administratorAPIToken;

  public static final String SERIALIZED_NAME_REQUIRED_GROUP = "requiredGroup";
  @SerializedName(SERIALIZED_NAME_REQUIRED_GROUP)
  private String requiredGroup;

  public static final String SERIALIZED_NAME_REQUIRED_GROUP_ID = "requiredGroupId";
  @SerializedName(SERIALIZED_NAME_REQUIRED_GROUP_ID)
  private String requiredGroupId;

  public static final String SERIALIZED_NAME_ADMINISTRATOR_A_P_I_TOKEN_HASH = "administratorAPITokenHash";
  @SerializedName(SERIALIZED_NAME_ADMINISTRATOR_A_P_I_TOKEN_HASH)
  private String administratorAPITokenHash;


  public IdentitySourcesOktaConfigConfig url(String url) {
    
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


  public IdentitySourcesOktaConfigConfig administratorAPIToken(String administratorAPIToken) {
    
    this.administratorAPIToken = administratorAPIToken;
    return this;
  }

   /**
   * Get administratorAPIToken
   * @return administratorAPIToken
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public String getAdministratorAPIToken() {
    return administratorAPIToken;
  }


  public void setAdministratorAPIToken(String administratorAPIToken) {
    this.administratorAPIToken = administratorAPIToken;
  }


  public IdentitySourcesOktaConfigConfig requiredGroup(String requiredGroup) {
    
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


  public IdentitySourcesOktaConfigConfig requiredGroupId(String requiredGroupId) {
    
    this.requiredGroupId = requiredGroupId;
    return this;
  }

   /**
   * Get requiredGroupId
   * @return requiredGroupId
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public String getRequiredGroupId() {
    return requiredGroupId;
  }


  public void setRequiredGroupId(String requiredGroupId) {
    this.requiredGroupId = requiredGroupId;
  }


  public IdentitySourcesOktaConfigConfig administratorAPITokenHash(String administratorAPITokenHash) {
    
    this.administratorAPITokenHash = administratorAPITokenHash;
    return this;
  }

   /**
   * Get administratorAPITokenHash
   * @return administratorAPITokenHash
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public String getAdministratorAPITokenHash() {
    return administratorAPITokenHash;
  }


  public void setAdministratorAPITokenHash(String administratorAPITokenHash) {
    this.administratorAPITokenHash = administratorAPITokenHash;
  }


  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (o == null || getClass() != o.getClass()) {
      return false;
    }
    IdentitySourcesOktaConfigConfig identitySourcesOktaConfigConfig = (IdentitySourcesOktaConfigConfig) o;
    return Objects.equals(this.url, identitySourcesOktaConfigConfig.url) &&
        Objects.equals(this.administratorAPIToken, identitySourcesOktaConfigConfig.administratorAPIToken) &&
        Objects.equals(this.requiredGroup, identitySourcesOktaConfigConfig.requiredGroup) &&
        Objects.equals(this.requiredGroupId, identitySourcesOktaConfigConfig.requiredGroupId) &&
        Objects.equals(this.administratorAPITokenHash, identitySourcesOktaConfigConfig.administratorAPITokenHash);
  }

  @Override
  public int hashCode() {
    return Objects.hash(url, administratorAPIToken, requiredGroup, requiredGroupId, administratorAPITokenHash);
  }


  @Override
  public String toString() {
    StringBuilder sb = new StringBuilder();
    sb.append("class IdentitySourcesOktaConfigConfig {\n");
    sb.append("    url: ").append(toIndentedString(url)).append("\n");
    sb.append("    administratorAPIToken: ").append(toIndentedString(administratorAPIToken)).append("\n");
    sb.append("    requiredGroup: ").append(toIndentedString(requiredGroup)).append("\n");
    sb.append("    requiredGroupId: ").append(toIndentedString(requiredGroupId)).append("\n");
    sb.append("    administratorAPITokenHash: ").append(toIndentedString(administratorAPITokenHash)).append("\n");
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

