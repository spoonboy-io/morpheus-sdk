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
 * TaskLibraryScriptConfigTaskOptions
 */
@javax.annotation.Generated(value = "org.openapitools.codegen.languages.JavaClientCodegen", date = "2023-08-21T11:26:47.028Z[GMT]")
public class TaskLibraryScriptConfigTaskOptions {
  public static final String SERIALIZED_NAME_HOST = "host";
  @SerializedName(SERIALIZED_NAME_HOST)
  private String host;

  public static final String SERIALIZED_NAME_USERNAME = "username";
  @SerializedName(SERIALIZED_NAME_USERNAME)
  private String username;

  public static final String SERIALIZED_NAME_PASSWORD = "password";
  @SerializedName(SERIALIZED_NAME_PASSWORD)
  private String password;

  public static final String SERIALIZED_NAME_PASSWORD_HASH = "passwordHash";
  @SerializedName(SERIALIZED_NAME_PASSWORD_HASH)
  private String passwordHash;

  public static final String SERIALIZED_NAME_LOCAL_SCRIPT_GIT_REF = "localScriptGitRef";
  @SerializedName(SERIALIZED_NAME_LOCAL_SCRIPT_GIT_REF)
  private String localScriptGitRef;

  public static final String SERIALIZED_NAME_LOCAL_SCRIPT_GIT_ID = "localScriptGitId";
  @SerializedName(SERIALIZED_NAME_LOCAL_SCRIPT_GIT_ID)
  private String localScriptGitId;

  public static final String SERIALIZED_NAME_CONTAINER_SCRIPT_ID = "containerScriptId";
  @SerializedName(SERIALIZED_NAME_CONTAINER_SCRIPT_ID)
  private String containerScriptId;

  public static final String SERIALIZED_NAME_SSH_KEY = "sshKey";
  @SerializedName(SERIALIZED_NAME_SSH_KEY)
  private String sshKey;

  public static final String SERIALIZED_NAME_CONTAINER_SCRIPT = "containerScript";
  @SerializedName(SERIALIZED_NAME_CONTAINER_SCRIPT)
  private String containerScript;

  public static final String SERIALIZED_NAME_PORT = "port";
  @SerializedName(SERIALIZED_NAME_PORT)
  private String port;


  public TaskLibraryScriptConfigTaskOptions host(String host) {
    
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


  public TaskLibraryScriptConfigTaskOptions username(String username) {
    
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


  public TaskLibraryScriptConfigTaskOptions password(String password) {
    
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


  public TaskLibraryScriptConfigTaskOptions passwordHash(String passwordHash) {
    
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


  public TaskLibraryScriptConfigTaskOptions localScriptGitRef(String localScriptGitRef) {
    
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


  public TaskLibraryScriptConfigTaskOptions localScriptGitId(String localScriptGitId) {
    
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


  public TaskLibraryScriptConfigTaskOptions containerScriptId(String containerScriptId) {
    
    this.containerScriptId = containerScriptId;
    return this;
  }

   /**
   * Get containerScriptId
   * @return containerScriptId
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public String getContainerScriptId() {
    return containerScriptId;
  }


  public void setContainerScriptId(String containerScriptId) {
    this.containerScriptId = containerScriptId;
  }


  public TaskLibraryScriptConfigTaskOptions sshKey(String sshKey) {
    
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


  public TaskLibraryScriptConfigTaskOptions containerScript(String containerScript) {
    
    this.containerScript = containerScript;
    return this;
  }

   /**
   * Get containerScript
   * @return containerScript
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public String getContainerScript() {
    return containerScript;
  }


  public void setContainerScript(String containerScript) {
    this.containerScript = containerScript;
  }


  public TaskLibraryScriptConfigTaskOptions port(String port) {
    
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


  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (o == null || getClass() != o.getClass()) {
      return false;
    }
    TaskLibraryScriptConfigTaskOptions taskLibraryScriptConfigTaskOptions = (TaskLibraryScriptConfigTaskOptions) o;
    return Objects.equals(this.host, taskLibraryScriptConfigTaskOptions.host) &&
        Objects.equals(this.username, taskLibraryScriptConfigTaskOptions.username) &&
        Objects.equals(this.password, taskLibraryScriptConfigTaskOptions.password) &&
        Objects.equals(this.passwordHash, taskLibraryScriptConfigTaskOptions.passwordHash) &&
        Objects.equals(this.localScriptGitRef, taskLibraryScriptConfigTaskOptions.localScriptGitRef) &&
        Objects.equals(this.localScriptGitId, taskLibraryScriptConfigTaskOptions.localScriptGitId) &&
        Objects.equals(this.containerScriptId, taskLibraryScriptConfigTaskOptions.containerScriptId) &&
        Objects.equals(this.sshKey, taskLibraryScriptConfigTaskOptions.sshKey) &&
        Objects.equals(this.containerScript, taskLibraryScriptConfigTaskOptions.containerScript) &&
        Objects.equals(this.port, taskLibraryScriptConfigTaskOptions.port);
  }

  @Override
  public int hashCode() {
    return Objects.hash(host, username, password, passwordHash, localScriptGitRef, localScriptGitId, containerScriptId, sshKey, containerScript, port);
  }


  @Override
  public String toString() {
    StringBuilder sb = new StringBuilder();
    sb.append("class TaskLibraryScriptConfigTaskOptions {\n");
    sb.append("    host: ").append(toIndentedString(host)).append("\n");
    sb.append("    username: ").append(toIndentedString(username)).append("\n");
    sb.append("    password: ").append(toIndentedString(password)).append("\n");
    sb.append("    passwordHash: ").append(toIndentedString(passwordHash)).append("\n");
    sb.append("    localScriptGitRef: ").append(toIndentedString(localScriptGitRef)).append("\n");
    sb.append("    localScriptGitId: ").append(toIndentedString(localScriptGitId)).append("\n");
    sb.append("    containerScriptId: ").append(toIndentedString(containerScriptId)).append("\n");
    sb.append("    sshKey: ").append(toIndentedString(sshKey)).append("\n");
    sb.append("    containerScript: ").append(toIndentedString(containerScript)).append("\n");
    sb.append("    port: ").append(toIndentedString(port)).append("\n");
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

