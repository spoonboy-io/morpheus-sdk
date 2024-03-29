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
 * TaskWriteAttributesConfigTaskOptions
 */
@javax.annotation.Generated(value = "org.openapitools.codegen.languages.JavaClientCodegen", date = "2023-08-21T11:26:47.028Z[GMT]")
public class TaskWriteAttributesConfigTaskOptions {
  public static final String SERIALIZED_NAME_HOST = "host";
  @SerializedName(SERIALIZED_NAME_HOST)
  private String host;

  public static final String SERIALIZED_NAME_SSH_KEY = "sshKey";
  @SerializedName(SERIALIZED_NAME_SSH_KEY)
  private String sshKey;

  public static final String SERIALIZED_NAME_LOCAL_SCRIPT_GIT_REF = "localScriptGitRef";
  @SerializedName(SERIALIZED_NAME_LOCAL_SCRIPT_GIT_REF)
  private String localScriptGitRef;

  public static final String SERIALIZED_NAME_PORT = "port";
  @SerializedName(SERIALIZED_NAME_PORT)
  private String port;

  public static final String SERIALIZED_NAME_LOCAL_SCRIPT_GIT_ID = "localScriptGitId";
  @SerializedName(SERIALIZED_NAME_LOCAL_SCRIPT_GIT_ID)
  private String localScriptGitId;

  public static final String SERIALIZED_NAME_PASSWORD = "password";
  @SerializedName(SERIALIZED_NAME_PASSWORD)
  private String password;

  public static final String SERIALIZED_NAME_PASSWORD_HASH = "passwordHash";
  @SerializedName(SERIALIZED_NAME_PASSWORD_HASH)
  private String passwordHash;

  public static final String SERIALIZED_NAME_WRITE_ATTRIBUTES_ATTRIBUTES = "writeAttributes.attributes";
  @SerializedName(SERIALIZED_NAME_WRITE_ATTRIBUTES_ATTRIBUTES)
  private String writeAttributesAttributes;

  public static final String SERIALIZED_NAME_USERNAME = "username";
  @SerializedName(SERIALIZED_NAME_USERNAME)
  private String username;


  public TaskWriteAttributesConfigTaskOptions host(String host) {
    
    this.host = host;
    return this;
  }

   /**
   * Get host
   * @return host
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public String getHost() {
    return host;
  }


  public void setHost(String host) {
    this.host = host;
  }


  public TaskWriteAttributesConfigTaskOptions sshKey(String sshKey) {
    
    this.sshKey = sshKey;
    return this;
  }

   /**
   * Get sshKey
   * @return sshKey
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public String getSshKey() {
    return sshKey;
  }


  public void setSshKey(String sshKey) {
    this.sshKey = sshKey;
  }


  public TaskWriteAttributesConfigTaskOptions localScriptGitRef(String localScriptGitRef) {
    
    this.localScriptGitRef = localScriptGitRef;
    return this;
  }

   /**
   * Get localScriptGitRef
   * @return localScriptGitRef
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public String getLocalScriptGitRef() {
    return localScriptGitRef;
  }


  public void setLocalScriptGitRef(String localScriptGitRef) {
    this.localScriptGitRef = localScriptGitRef;
  }


  public TaskWriteAttributesConfigTaskOptions port(String port) {
    
    this.port = port;
    return this;
  }

   /**
   * Get port
   * @return port
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public String getPort() {
    return port;
  }


  public void setPort(String port) {
    this.port = port;
  }


  public TaskWriteAttributesConfigTaskOptions localScriptGitId(String localScriptGitId) {
    
    this.localScriptGitId = localScriptGitId;
    return this;
  }

   /**
   * Get localScriptGitId
   * @return localScriptGitId
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public String getLocalScriptGitId() {
    return localScriptGitId;
  }


  public void setLocalScriptGitId(String localScriptGitId) {
    this.localScriptGitId = localScriptGitId;
  }


  public TaskWriteAttributesConfigTaskOptions password(String password) {
    
    this.password = password;
    return this;
  }

   /**
   * Get password
   * @return password
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public String getPassword() {
    return password;
  }


  public void setPassword(String password) {
    this.password = password;
  }


  public TaskWriteAttributesConfigTaskOptions passwordHash(String passwordHash) {
    
    this.passwordHash = passwordHash;
    return this;
  }

   /**
   * Get passwordHash
   * @return passwordHash
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public String getPasswordHash() {
    return passwordHash;
  }


  public void setPasswordHash(String passwordHash) {
    this.passwordHash = passwordHash;
  }


  public TaskWriteAttributesConfigTaskOptions writeAttributesAttributes(String writeAttributesAttributes) {
    
    this.writeAttributesAttributes = writeAttributesAttributes;
    return this;
  }

   /**
   * Get writeAttributesAttributes
   * @return writeAttributesAttributes
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public String getWriteAttributesAttributes() {
    return writeAttributesAttributes;
  }


  public void setWriteAttributesAttributes(String writeAttributesAttributes) {
    this.writeAttributesAttributes = writeAttributesAttributes;
  }


  public TaskWriteAttributesConfigTaskOptions username(String username) {
    
    this.username = username;
    return this;
  }

   /**
   * Get username
   * @return username
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public String getUsername() {
    return username;
  }


  public void setUsername(String username) {
    this.username = username;
  }


  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (o == null || getClass() != o.getClass()) {
      return false;
    }
    TaskWriteAttributesConfigTaskOptions taskWriteAttributesConfigTaskOptions = (TaskWriteAttributesConfigTaskOptions) o;
    return Objects.equals(this.host, taskWriteAttributesConfigTaskOptions.host) &&
        Objects.equals(this.sshKey, taskWriteAttributesConfigTaskOptions.sshKey) &&
        Objects.equals(this.localScriptGitRef, taskWriteAttributesConfigTaskOptions.localScriptGitRef) &&
        Objects.equals(this.port, taskWriteAttributesConfigTaskOptions.port) &&
        Objects.equals(this.localScriptGitId, taskWriteAttributesConfigTaskOptions.localScriptGitId) &&
        Objects.equals(this.password, taskWriteAttributesConfigTaskOptions.password) &&
        Objects.equals(this.passwordHash, taskWriteAttributesConfigTaskOptions.passwordHash) &&
        Objects.equals(this.writeAttributesAttributes, taskWriteAttributesConfigTaskOptions.writeAttributesAttributes) &&
        Objects.equals(this.username, taskWriteAttributesConfigTaskOptions.username);
  }

  @Override
  public int hashCode() {
    return Objects.hash(host, sshKey, localScriptGitRef, port, localScriptGitId, password, passwordHash, writeAttributesAttributes, username);
  }


  @Override
  public String toString() {
    StringBuilder sb = new StringBuilder();
    sb.append("class TaskWriteAttributesConfigTaskOptions {\n");
    sb.append("    host: ").append(toIndentedString(host)).append("\n");
    sb.append("    sshKey: ").append(toIndentedString(sshKey)).append("\n");
    sb.append("    localScriptGitRef: ").append(toIndentedString(localScriptGitRef)).append("\n");
    sb.append("    port: ").append(toIndentedString(port)).append("\n");
    sb.append("    localScriptGitId: ").append(toIndentedString(localScriptGitId)).append("\n");
    sb.append("    password: ").append(toIndentedString(password)).append("\n");
    sb.append("    passwordHash: ").append(toIndentedString(passwordHash)).append("\n");
    sb.append("    writeAttributesAttributes: ").append(toIndentedString(writeAttributesAttributes)).append("\n");
    sb.append("    username: ").append(toIndentedString(username)).append("\n");
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

