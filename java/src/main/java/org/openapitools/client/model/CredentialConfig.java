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
 * CredentialConfig
 */
@javax.annotation.Generated(value = "org.openapitools.codegen.languages.JavaClientCodegen", date = "2023-08-21T11:26:47.028Z[GMT]")
public class CredentialConfig {
  public static final String SERIALIZED_NAME_CLIENT_SECRET = "clientSecret";
  @SerializedName(SERIALIZED_NAME_CLIENT_SECRET)
  private String clientSecret;

  public static final String SERIALIZED_NAME_CLIENT_ID = "clientId";
  @SerializedName(SERIALIZED_NAME_CLIENT_ID)
  private String clientId;

  public static final String SERIALIZED_NAME_CLIENT_AUTH = "clientAuth";
  @SerializedName(SERIALIZED_NAME_CLIENT_AUTH)
  private String clientAuth;

  public static final String SERIALIZED_NAME_SCOPE = "scope";
  @SerializedName(SERIALIZED_NAME_SCOPE)
  private String scope;

  public static final String SERIALIZED_NAME_GRANT_TYPE = "grantType";
  @SerializedName(SERIALIZED_NAME_GRANT_TYPE)
  private String grantType;

  public static final String SERIALIZED_NAME_ACCESS_TOKEN_URL = "accessTokenUrl";
  @SerializedName(SERIALIZED_NAME_ACCESS_TOKEN_URL)
  private String accessTokenUrl;

  public static final String SERIALIZED_NAME_CLIENT_SECRET_HASH = "clientSecretHash";
  @SerializedName(SERIALIZED_NAME_CLIENT_SECRET_HASH)
  private String clientSecretHash;


  public CredentialConfig clientSecret(String clientSecret) {
    
    this.clientSecret = clientSecret;
    return this;
  }

   /**
   * Get clientSecret
   * @return clientSecret
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public String getClientSecret() {
    return clientSecret;
  }


  public void setClientSecret(String clientSecret) {
    this.clientSecret = clientSecret;
  }


  public CredentialConfig clientId(String clientId) {
    
    this.clientId = clientId;
    return this;
  }

   /**
   * Get clientId
   * @return clientId
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public String getClientId() {
    return clientId;
  }


  public void setClientId(String clientId) {
    this.clientId = clientId;
  }


  public CredentialConfig clientAuth(String clientAuth) {
    
    this.clientAuth = clientAuth;
    return this;
  }

   /**
   * Get clientAuth
   * @return clientAuth
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public String getClientAuth() {
    return clientAuth;
  }


  public void setClientAuth(String clientAuth) {
    this.clientAuth = clientAuth;
  }


  public CredentialConfig scope(String scope) {
    
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


  public CredentialConfig grantType(String grantType) {
    
    this.grantType = grantType;
    return this;
  }

   /**
   * Get grantType
   * @return grantType
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public String getGrantType() {
    return grantType;
  }


  public void setGrantType(String grantType) {
    this.grantType = grantType;
  }


  public CredentialConfig accessTokenUrl(String accessTokenUrl) {
    
    this.accessTokenUrl = accessTokenUrl;
    return this;
  }

   /**
   * Get accessTokenUrl
   * @return accessTokenUrl
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public String getAccessTokenUrl() {
    return accessTokenUrl;
  }


  public void setAccessTokenUrl(String accessTokenUrl) {
    this.accessTokenUrl = accessTokenUrl;
  }


  public CredentialConfig clientSecretHash(String clientSecretHash) {
    
    this.clientSecretHash = clientSecretHash;
    return this;
  }

   /**
   * Get clientSecretHash
   * @return clientSecretHash
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public String getClientSecretHash() {
    return clientSecretHash;
  }


  public void setClientSecretHash(String clientSecretHash) {
    this.clientSecretHash = clientSecretHash;
  }


  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (o == null || getClass() != o.getClass()) {
      return false;
    }
    CredentialConfig credentialConfig = (CredentialConfig) o;
    return Objects.equals(this.clientSecret, credentialConfig.clientSecret) &&
        Objects.equals(this.clientId, credentialConfig.clientId) &&
        Objects.equals(this.clientAuth, credentialConfig.clientAuth) &&
        Objects.equals(this.scope, credentialConfig.scope) &&
        Objects.equals(this.grantType, credentialConfig.grantType) &&
        Objects.equals(this.accessTokenUrl, credentialConfig.accessTokenUrl) &&
        Objects.equals(this.clientSecretHash, credentialConfig.clientSecretHash);
  }

  @Override
  public int hashCode() {
    return Objects.hash(clientSecret, clientId, clientAuth, scope, grantType, accessTokenUrl, clientSecretHash);
  }


  @Override
  public String toString() {
    StringBuilder sb = new StringBuilder();
    sb.append("class CredentialConfig {\n");
    sb.append("    clientSecret: ").append(toIndentedString(clientSecret)).append("\n");
    sb.append("    clientId: ").append(toIndentedString(clientId)).append("\n");
    sb.append("    clientAuth: ").append(toIndentedString(clientAuth)).append("\n");
    sb.append("    scope: ").append(toIndentedString(scope)).append("\n");
    sb.append("    grantType: ").append(toIndentedString(grantType)).append("\n");
    sb.append("    accessTokenUrl: ").append(toIndentedString(accessTokenUrl)).append("\n");
    sb.append("    clientSecretHash: ").append(toIndentedString(clientSecretHash)).append("\n");
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

