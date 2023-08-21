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
import org.threeten.bp.OffsetDateTime;

/**
 * ExecutionRequest
 */
@javax.annotation.Generated(value = "org.openapitools.codegen.languages.JavaClientCodegen", date = "2023-08-21T11:26:47.028Z[GMT]")
public class ExecutionRequest {
  public static final String SERIALIZED_NAME_ID = "id";
  @SerializedName(SERIALIZED_NAME_ID)
  private Long id;

  public static final String SERIALIZED_NAME_UNIQUE_ID = "uniqueId";
  @SerializedName(SERIALIZED_NAME_UNIQUE_ID)
  private String uniqueId;

  public static final String SERIALIZED_NAME_CONTAINER_ID = "containerId";
  @SerializedName(SERIALIZED_NAME_CONTAINER_ID)
  private String containerId;

  public static final String SERIALIZED_NAME_SERVER_ID = "serverId";
  @SerializedName(SERIALIZED_NAME_SERVER_ID)
  private String serverId;

  public static final String SERIALIZED_NAME_INSTANCE_ID = "instanceId";
  @SerializedName(SERIALIZED_NAME_INSTANCE_ID)
  private Long instanceId;

  public static final String SERIALIZED_NAME_RESOURCE_ID = "resourceId";
  @SerializedName(SERIALIZED_NAME_RESOURCE_ID)
  private String resourceId;

  public static final String SERIALIZED_NAME_APP_ID = "appId";
  @SerializedName(SERIALIZED_NAME_APP_ID)
  private String appId;

  public static final String SERIALIZED_NAME_STD_OUT = "stdOut";
  @SerializedName(SERIALIZED_NAME_STD_OUT)
  private String stdOut;

  public static final String SERIALIZED_NAME_STD_ERR = "stdErr";
  @SerializedName(SERIALIZED_NAME_STD_ERR)
  private String stdErr;

  public static final String SERIALIZED_NAME_EXIT_CODE = "exitCode";
  @SerializedName(SERIALIZED_NAME_EXIT_CODE)
  private Long exitCode;

  public static final String SERIALIZED_NAME_STATUS = "status";
  @SerializedName(SERIALIZED_NAME_STATUS)
  private String status;

  public static final String SERIALIZED_NAME_EXPIRES_AT = "expiresAt";
  @SerializedName(SERIALIZED_NAME_EXPIRES_AT)
  private OffsetDateTime expiresAt;

  public static final String SERIALIZED_NAME_CREATED_BY_ID = "createdById";
  @SerializedName(SERIALIZED_NAME_CREATED_BY_ID)
  private Long createdById;

  public static final String SERIALIZED_NAME_STATUS_MESSAGE = "statusMessage";
  @SerializedName(SERIALIZED_NAME_STATUS_MESSAGE)
  private String statusMessage;

  public static final String SERIALIZED_NAME_ERROR_MESSAGE = "errorMessage";
  @SerializedName(SERIALIZED_NAME_ERROR_MESSAGE)
  private String errorMessage;

  public static final String SERIALIZED_NAME_CONFIG = "config";
  @SerializedName(SERIALIZED_NAME_CONFIG)
  private Object config;

  public static final String SERIALIZED_NAME_RAW_DATA = "rawData";
  @SerializedName(SERIALIZED_NAME_RAW_DATA)
  private String rawData;


  public ExecutionRequest id(Long id) {
    
    this.id = id;
    return this;
  }

   /**
   * Get id
   * @return id
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public Long getId() {
    return id;
  }


  public void setId(Long id) {
    this.id = id;
  }


  public ExecutionRequest uniqueId(String uniqueId) {
    
    this.uniqueId = uniqueId;
    return this;
  }

   /**
   * Get uniqueId
   * @return uniqueId
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public String getUniqueId() {
    return uniqueId;
  }


  public void setUniqueId(String uniqueId) {
    this.uniqueId = uniqueId;
  }


  public ExecutionRequest containerId(String containerId) {
    
    this.containerId = containerId;
    return this;
  }

   /**
   * Get containerId
   * @return containerId
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public String getContainerId() {
    return containerId;
  }


  public void setContainerId(String containerId) {
    this.containerId = containerId;
  }


  public ExecutionRequest serverId(String serverId) {
    
    this.serverId = serverId;
    return this;
  }

   /**
   * Get serverId
   * @return serverId
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public String getServerId() {
    return serverId;
  }


  public void setServerId(String serverId) {
    this.serverId = serverId;
  }


  public ExecutionRequest instanceId(Long instanceId) {
    
    this.instanceId = instanceId;
    return this;
  }

   /**
   * Get instanceId
   * @return instanceId
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public Long getInstanceId() {
    return instanceId;
  }


  public void setInstanceId(Long instanceId) {
    this.instanceId = instanceId;
  }


  public ExecutionRequest resourceId(String resourceId) {
    
    this.resourceId = resourceId;
    return this;
  }

   /**
   * Get resourceId
   * @return resourceId
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public String getResourceId() {
    return resourceId;
  }


  public void setResourceId(String resourceId) {
    this.resourceId = resourceId;
  }


  public ExecutionRequest appId(String appId) {
    
    this.appId = appId;
    return this;
  }

   /**
   * Get appId
   * @return appId
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public String getAppId() {
    return appId;
  }


  public void setAppId(String appId) {
    this.appId = appId;
  }


  public ExecutionRequest stdOut(String stdOut) {
    
    this.stdOut = stdOut;
    return this;
  }

   /**
   * Get stdOut
   * @return stdOut
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public String getStdOut() {
    return stdOut;
  }


  public void setStdOut(String stdOut) {
    this.stdOut = stdOut;
  }


  public ExecutionRequest stdErr(String stdErr) {
    
    this.stdErr = stdErr;
    return this;
  }

   /**
   * Get stdErr
   * @return stdErr
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public String getStdErr() {
    return stdErr;
  }


  public void setStdErr(String stdErr) {
    this.stdErr = stdErr;
  }


  public ExecutionRequest exitCode(Long exitCode) {
    
    this.exitCode = exitCode;
    return this;
  }

   /**
   * Get exitCode
   * @return exitCode
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public Long getExitCode() {
    return exitCode;
  }


  public void setExitCode(Long exitCode) {
    this.exitCode = exitCode;
  }


  public ExecutionRequest status(String status) {
    
    this.status = status;
    return this;
  }

   /**
   * Get status
   * @return status
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public String getStatus() {
    return status;
  }


  public void setStatus(String status) {
    this.status = status;
  }


  public ExecutionRequest expiresAt(OffsetDateTime expiresAt) {
    
    this.expiresAt = expiresAt;
    return this;
  }

   /**
   * Get expiresAt
   * @return expiresAt
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public OffsetDateTime getExpiresAt() {
    return expiresAt;
  }


  public void setExpiresAt(OffsetDateTime expiresAt) {
    this.expiresAt = expiresAt;
  }


  public ExecutionRequest createdById(Long createdById) {
    
    this.createdById = createdById;
    return this;
  }

   /**
   * Get createdById
   * @return createdById
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public Long getCreatedById() {
    return createdById;
  }


  public void setCreatedById(Long createdById) {
    this.createdById = createdById;
  }


  public ExecutionRequest statusMessage(String statusMessage) {
    
    this.statusMessage = statusMessage;
    return this;
  }

   /**
   * Get statusMessage
   * @return statusMessage
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public String getStatusMessage() {
    return statusMessage;
  }


  public void setStatusMessage(String statusMessage) {
    this.statusMessage = statusMessage;
  }


  public ExecutionRequest errorMessage(String errorMessage) {
    
    this.errorMessage = errorMessage;
    return this;
  }

   /**
   * Get errorMessage
   * @return errorMessage
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public String getErrorMessage() {
    return errorMessage;
  }


  public void setErrorMessage(String errorMessage) {
    this.errorMessage = errorMessage;
  }


  public ExecutionRequest config(Object config) {
    
    this.config = config;
    return this;
  }

   /**
   * Get config
   * @return config
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public Object getConfig() {
    return config;
  }


  public void setConfig(Object config) {
    this.config = config;
  }


  public ExecutionRequest rawData(String rawData) {
    
    this.rawData = rawData;
    return this;
  }

   /**
   * Get rawData
   * @return rawData
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public String getRawData() {
    return rawData;
  }


  public void setRawData(String rawData) {
    this.rawData = rawData;
  }


  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (o == null || getClass() != o.getClass()) {
      return false;
    }
    ExecutionRequest executionRequest = (ExecutionRequest) o;
    return Objects.equals(this.id, executionRequest.id) &&
        Objects.equals(this.uniqueId, executionRequest.uniqueId) &&
        Objects.equals(this.containerId, executionRequest.containerId) &&
        Objects.equals(this.serverId, executionRequest.serverId) &&
        Objects.equals(this.instanceId, executionRequest.instanceId) &&
        Objects.equals(this.resourceId, executionRequest.resourceId) &&
        Objects.equals(this.appId, executionRequest.appId) &&
        Objects.equals(this.stdOut, executionRequest.stdOut) &&
        Objects.equals(this.stdErr, executionRequest.stdErr) &&
        Objects.equals(this.exitCode, executionRequest.exitCode) &&
        Objects.equals(this.status, executionRequest.status) &&
        Objects.equals(this.expiresAt, executionRequest.expiresAt) &&
        Objects.equals(this.createdById, executionRequest.createdById) &&
        Objects.equals(this.statusMessage, executionRequest.statusMessage) &&
        Objects.equals(this.errorMessage, executionRequest.errorMessage) &&
        Objects.equals(this.config, executionRequest.config) &&
        Objects.equals(this.rawData, executionRequest.rawData);
  }

  @Override
  public int hashCode() {
    return Objects.hash(id, uniqueId, containerId, serverId, instanceId, resourceId, appId, stdOut, stdErr, exitCode, status, expiresAt, createdById, statusMessage, errorMessage, config, rawData);
  }


  @Override
  public String toString() {
    StringBuilder sb = new StringBuilder();
    sb.append("class ExecutionRequest {\n");
    sb.append("    id: ").append(toIndentedString(id)).append("\n");
    sb.append("    uniqueId: ").append(toIndentedString(uniqueId)).append("\n");
    sb.append("    containerId: ").append(toIndentedString(containerId)).append("\n");
    sb.append("    serverId: ").append(toIndentedString(serverId)).append("\n");
    sb.append("    instanceId: ").append(toIndentedString(instanceId)).append("\n");
    sb.append("    resourceId: ").append(toIndentedString(resourceId)).append("\n");
    sb.append("    appId: ").append(toIndentedString(appId)).append("\n");
    sb.append("    stdOut: ").append(toIndentedString(stdOut)).append("\n");
    sb.append("    stdErr: ").append(toIndentedString(stdErr)).append("\n");
    sb.append("    exitCode: ").append(toIndentedString(exitCode)).append("\n");
    sb.append("    status: ").append(toIndentedString(status)).append("\n");
    sb.append("    expiresAt: ").append(toIndentedString(expiresAt)).append("\n");
    sb.append("    createdById: ").append(toIndentedString(createdById)).append("\n");
    sb.append("    statusMessage: ").append(toIndentedString(statusMessage)).append("\n");
    sb.append("    errorMessage: ").append(toIndentedString(errorMessage)).append("\n");
    sb.append("    config: ").append(toIndentedString(config)).append("\n");
    sb.append("    rawData: ").append(toIndentedString(rawData)).append("\n");
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

