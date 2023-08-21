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
 * TaskPuppetConfigTaskOptions
 */
@javax.annotation.Generated(value = "org.openapitools.codegen.languages.JavaClientCodegen", date = "2023-08-21T11:26:47.028Z[GMT]")
public class TaskPuppetConfigTaskOptions {
  public static final String SERIALIZED_NAME_PORT = "port";
  @SerializedName(SERIALIZED_NAME_PORT)
  private String port;

  public static final String SERIALIZED_NAME_HOST = "host";
  @SerializedName(SERIALIZED_NAME_HOST)
  private String host;

  public static final String SERIALIZED_NAME_USERNAME = "username";
  @SerializedName(SERIALIZED_NAME_USERNAME)
  private String username;

  public static final String SERIALIZED_NAME_PUPPET_ENVIRONMENT = "puppetEnvironment";
  @SerializedName(SERIALIZED_NAME_PUPPET_ENVIRONMENT)
  private String puppetEnvironment;

  public static final String SERIALIZED_NAME_PUPPET_NODE_NAME = "puppetNodeName";
  @SerializedName(SERIALIZED_NAME_PUPPET_NODE_NAME)
  private String puppetNodeName;

  public static final String SERIALIZED_NAME_SSH_KEY = "sshKey";
  @SerializedName(SERIALIZED_NAME_SSH_KEY)
  private String sshKey;

  public static final String SERIALIZED_NAME_LOCAL_SCRIPT_GIT_ID = "localScriptGitId";
  @SerializedName(SERIALIZED_NAME_LOCAL_SCRIPT_GIT_ID)
  private String localScriptGitId;

  public static final String SERIALIZED_NAME_LOCAL_SCRIPT_GIT_REF = "localScriptGitRef";
  @SerializedName(SERIALIZED_NAME_LOCAL_SCRIPT_GIT_REF)
  private String localScriptGitRef;

  public static final String SERIALIZED_NAME_PASSWORD = "password";
  @SerializedName(SERIALIZED_NAME_PASSWORD)
  private String password;

  public static final String SERIALIZED_NAME_PASSWORD_HASH = "passwordHash";
  @SerializedName(SERIALIZED_NAME_PASSWORD_HASH)
  private String passwordHash;


  public TaskPuppetConfigTaskOptions port(String port) {
    
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


  public TaskPuppetConfigTaskOptions host(String host) {
    
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


  public TaskPuppetConfigTaskOptions username(String username) {
    
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


  public TaskPuppetConfigTaskOptions puppetEnvironment(String puppetEnvironment) {
    
    this.puppetEnvironment = puppetEnvironment;
    return this;
  }

   /**
   * Get puppetEnvironment
   * @return puppetEnvironment
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public String getPuppetEnvironment() {
    return puppetEnvironment;
  }


  public void setPuppetEnvironment(String puppetEnvironment) {
    this.puppetEnvironment = puppetEnvironment;
  }


  public TaskPuppetConfigTaskOptions puppetNodeName(String puppetNodeName) {
    
    this.puppetNodeName = puppetNodeName;
    return this;
  }

   /**
   * Get puppetNodeName
   * @return puppetNodeName
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public String getPuppetNodeName() {
    return puppetNodeName;
  }


  public void setPuppetNodeName(String puppetNodeName) {
    this.puppetNodeName = puppetNodeName;
  }


  public TaskPuppetConfigTaskOptions sshKey(String sshKey) {
    
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


  public TaskPuppetConfigTaskOptions localScriptGitId(String localScriptGitId) {
    
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


  public TaskPuppetConfigTaskOptions localScriptGitRef(String localScriptGitRef) {
    
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


  public TaskPuppetConfigTaskOptions password(String password) {
    
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


  public TaskPuppetConfigTaskOptions passwordHash(String passwordHash) {
    
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


  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (o == null || getClass() != o.getClass()) {
      return false;
    }
    TaskPuppetConfigTaskOptions taskPuppetConfigTaskOptions = (TaskPuppetConfigTaskOptions) o;
    return Objects.equals(this.port, taskPuppetConfigTaskOptions.port) &&
        Objects.equals(this.host, taskPuppetConfigTaskOptions.host) &&
        Objects.equals(this.username, taskPuppetConfigTaskOptions.username) &&
        Objects.equals(this.puppetEnvironment, taskPuppetConfigTaskOptions.puppetEnvironment) &&
        Objects.equals(this.puppetNodeName, taskPuppetConfigTaskOptions.puppetNodeName) &&
        Objects.equals(this.sshKey, taskPuppetConfigTaskOptions.sshKey) &&
        Objects.equals(this.localScriptGitId, taskPuppetConfigTaskOptions.localScriptGitId) &&
        Objects.equals(this.localScriptGitRef, taskPuppetConfigTaskOptions.localScriptGitRef) &&
        Objects.equals(this.password, taskPuppetConfigTaskOptions.password) &&
        Objects.equals(this.passwordHash, taskPuppetConfigTaskOptions.passwordHash);
  }

  @Override
  public int hashCode() {
    return Objects.hash(port, host, username, puppetEnvironment, puppetNodeName, sshKey, localScriptGitId, localScriptGitRef, password, passwordHash);
  }


  @Override
  public String toString() {
    StringBuilder sb = new StringBuilder();
    sb.append("class TaskPuppetConfigTaskOptions {\n");
    sb.append("    port: ").append(toIndentedString(port)).append("\n");
    sb.append("    host: ").append(toIndentedString(host)).append("\n");
    sb.append("    username: ").append(toIndentedString(username)).append("\n");
    sb.append("    puppetEnvironment: ").append(toIndentedString(puppetEnvironment)).append("\n");
    sb.append("    puppetNodeName: ").append(toIndentedString(puppetNodeName)).append("\n");
    sb.append("    sshKey: ").append(toIndentedString(sshKey)).append("\n");
    sb.append("    localScriptGitId: ").append(toIndentedString(localScriptGitId)).append("\n");
    sb.append("    localScriptGitRef: ").append(toIndentedString(localScriptGitRef)).append("\n");
    sb.append("    password: ").append(toIndentedString(password)).append("\n");
    sb.append("    passwordHash: ").append(toIndentedString(passwordHash)).append("\n");
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

