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
import org.openapitools.client.model.ClusterContainersAvailableActions;
import org.openapitools.client.model.ClusterHistoryCreatedBy;
import org.threeten.bp.OffsetDateTime;

/**
 * ClusterHistoryEventItem
 */
@javax.annotation.Generated(value = "org.openapitools.codegen.languages.JavaClientCodegen", date = "2023-08-21T11:26:47.028Z[GMT]")
public class ClusterHistoryEventItem {
  public static final String SERIALIZED_NAME_ID = "id";
  @SerializedName(SERIALIZED_NAME_ID)
  private Long id;

  public static final String SERIALIZED_NAME_PROCESS_ID = "processId";
  @SerializedName(SERIALIZED_NAME_PROCESS_ID)
  private Long processId;

  public static final String SERIALIZED_NAME_ACCOUNT_ID = "accountId";
  @SerializedName(SERIALIZED_NAME_ACCOUNT_ID)
  private Long accountId;

  public static final String SERIALIZED_NAME_UNIQUE_ID = "uniqueId";
  @SerializedName(SERIALIZED_NAME_UNIQUE_ID)
  private String uniqueId;

  public static final String SERIALIZED_NAME_PROCESS_TYPE = "processType";
  @SerializedName(SERIALIZED_NAME_PROCESS_TYPE)
  private ClusterContainersAvailableActions processType;

  public static final String SERIALIZED_NAME_DESCRIPTION = "description";
  @SerializedName(SERIALIZED_NAME_DESCRIPTION)
  private String description;

  public static final String SERIALIZED_NAME_REF_TYPE = "refType";
  @SerializedName(SERIALIZED_NAME_REF_TYPE)
  private String refType;

  public static final String SERIALIZED_NAME_REF_ID = "refId";
  @SerializedName(SERIALIZED_NAME_REF_ID)
  private Long refId;

  public static final String SERIALIZED_NAME_SUB_TYPE = "subType";
  @SerializedName(SERIALIZED_NAME_SUB_TYPE)
  private String subType;

  public static final String SERIALIZED_NAME_SUB_ID = "subId";
  @SerializedName(SERIALIZED_NAME_SUB_ID)
  private String subId;

  public static final String SERIALIZED_NAME_ZONE_ID = "zoneId";
  @SerializedName(SERIALIZED_NAME_ZONE_ID)
  private String zoneId;

  public static final String SERIALIZED_NAME_INTEGRATION_ID = "integrationId";
  @SerializedName(SERIALIZED_NAME_INTEGRATION_ID)
  private String integrationId;

  public static final String SERIALIZED_NAME_INSTANCE_ID = "instanceId";
  @SerializedName(SERIALIZED_NAME_INSTANCE_ID)
  private String instanceId;

  public static final String SERIALIZED_NAME_CONTAINER_ID = "containerId";
  @SerializedName(SERIALIZED_NAME_CONTAINER_ID)
  private String containerId;

  public static final String SERIALIZED_NAME_SERVER_ID = "serverId";
  @SerializedName(SERIALIZED_NAME_SERVER_ID)
  private Long serverId;

  public static final String SERIALIZED_NAME_CONTAINER_NAME = "containerName";
  @SerializedName(SERIALIZED_NAME_CONTAINER_NAME)
  private String containerName;

  public static final String SERIALIZED_NAME_DISPLAY_NAME = "displayName";
  @SerializedName(SERIALIZED_NAME_DISPLAY_NAME)
  private String displayName;

  public static final String SERIALIZED_NAME_STATUS = "status";
  @SerializedName(SERIALIZED_NAME_STATUS)
  private String status;

  public static final String SERIALIZED_NAME_REASON = "reason";
  @SerializedName(SERIALIZED_NAME_REASON)
  private String reason;

  public static final String SERIALIZED_NAME_PERCENT = "percent";
  @SerializedName(SERIALIZED_NAME_PERCENT)
  private Long percent;

  public static final String SERIALIZED_NAME_STATUS_ETA = "statusEta";
  @SerializedName(SERIALIZED_NAME_STATUS_ETA)
  private Long statusEta;

  public static final String SERIALIZED_NAME_MESSAGE = "message";
  @SerializedName(SERIALIZED_NAME_MESSAGE)
  private String message;

  public static final String SERIALIZED_NAME_OUTPUT = "output";
  @SerializedName(SERIALIZED_NAME_OUTPUT)
  private String output;

  public static final String SERIALIZED_NAME_ERROR = "error";
  @SerializedName(SERIALIZED_NAME_ERROR)
  private String error;

  public static final String SERIALIZED_NAME_START_DATE = "startDate";
  @SerializedName(SERIALIZED_NAME_START_DATE)
  private OffsetDateTime startDate;

  public static final String SERIALIZED_NAME_END_DATE = "endDate";
  @SerializedName(SERIALIZED_NAME_END_DATE)
  private OffsetDateTime endDate;

  public static final String SERIALIZED_NAME_DURATION = "duration";
  @SerializedName(SERIALIZED_NAME_DURATION)
  private Long duration;

  public static final String SERIALIZED_NAME_DATE_CREATED = "dateCreated";
  @SerializedName(SERIALIZED_NAME_DATE_CREATED)
  private OffsetDateTime dateCreated;

  public static final String SERIALIZED_NAME_LAST_UPDATED = "lastUpdated";
  @SerializedName(SERIALIZED_NAME_LAST_UPDATED)
  private OffsetDateTime lastUpdated;

  public static final String SERIALIZED_NAME_CREATED_BY = "createdBy";
  @SerializedName(SERIALIZED_NAME_CREATED_BY)
  private ClusterHistoryCreatedBy createdBy;

  public static final String SERIALIZED_NAME_UPDATED_BY = "updatedBy";
  @SerializedName(SERIALIZED_NAME_UPDATED_BY)
  private ClusterHistoryCreatedBy updatedBy;


  public ClusterHistoryEventItem id(Long id) {
    
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


  public ClusterHistoryEventItem processId(Long processId) {
    
    this.processId = processId;
    return this;
  }

   /**
   * Get processId
   * @return processId
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public Long getProcessId() {
    return processId;
  }


  public void setProcessId(Long processId) {
    this.processId = processId;
  }


  public ClusterHistoryEventItem accountId(Long accountId) {
    
    this.accountId = accountId;
    return this;
  }

   /**
   * Get accountId
   * @return accountId
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public Long getAccountId() {
    return accountId;
  }


  public void setAccountId(Long accountId) {
    this.accountId = accountId;
  }


  public ClusterHistoryEventItem uniqueId(String uniqueId) {
    
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


  public ClusterHistoryEventItem processType(ClusterContainersAvailableActions processType) {
    
    this.processType = processType;
    return this;
  }

   /**
   * Get processType
   * @return processType
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public ClusterContainersAvailableActions getProcessType() {
    return processType;
  }


  public void setProcessType(ClusterContainersAvailableActions processType) {
    this.processType = processType;
  }


  public ClusterHistoryEventItem description(String description) {
    
    this.description = description;
    return this;
  }

   /**
   * Get description
   * @return description
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public String getDescription() {
    return description;
  }


  public void setDescription(String description) {
    this.description = description;
  }


  public ClusterHistoryEventItem refType(String refType) {
    
    this.refType = refType;
    return this;
  }

   /**
   * Get refType
   * @return refType
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public String getRefType() {
    return refType;
  }


  public void setRefType(String refType) {
    this.refType = refType;
  }


  public ClusterHistoryEventItem refId(Long refId) {
    
    this.refId = refId;
    return this;
  }

   /**
   * Get refId
   * @return refId
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public Long getRefId() {
    return refId;
  }


  public void setRefId(Long refId) {
    this.refId = refId;
  }


  public ClusterHistoryEventItem subType(String subType) {
    
    this.subType = subType;
    return this;
  }

   /**
   * Get subType
   * @return subType
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public String getSubType() {
    return subType;
  }


  public void setSubType(String subType) {
    this.subType = subType;
  }


  public ClusterHistoryEventItem subId(String subId) {
    
    this.subId = subId;
    return this;
  }

   /**
   * Get subId
   * @return subId
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public String getSubId() {
    return subId;
  }


  public void setSubId(String subId) {
    this.subId = subId;
  }


  public ClusterHistoryEventItem zoneId(String zoneId) {
    
    this.zoneId = zoneId;
    return this;
  }

   /**
   * Get zoneId
   * @return zoneId
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public String getZoneId() {
    return zoneId;
  }


  public void setZoneId(String zoneId) {
    this.zoneId = zoneId;
  }


  public ClusterHistoryEventItem integrationId(String integrationId) {
    
    this.integrationId = integrationId;
    return this;
  }

   /**
   * Get integrationId
   * @return integrationId
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public String getIntegrationId() {
    return integrationId;
  }


  public void setIntegrationId(String integrationId) {
    this.integrationId = integrationId;
  }


  public ClusterHistoryEventItem instanceId(String instanceId) {
    
    this.instanceId = instanceId;
    return this;
  }

   /**
   * Get instanceId
   * @return instanceId
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public String getInstanceId() {
    return instanceId;
  }


  public void setInstanceId(String instanceId) {
    this.instanceId = instanceId;
  }


  public ClusterHistoryEventItem containerId(String containerId) {
    
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


  public ClusterHistoryEventItem serverId(Long serverId) {
    
    this.serverId = serverId;
    return this;
  }

   /**
   * Get serverId
   * @return serverId
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public Long getServerId() {
    return serverId;
  }


  public void setServerId(Long serverId) {
    this.serverId = serverId;
  }


  public ClusterHistoryEventItem containerName(String containerName) {
    
    this.containerName = containerName;
    return this;
  }

   /**
   * Get containerName
   * @return containerName
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public String getContainerName() {
    return containerName;
  }


  public void setContainerName(String containerName) {
    this.containerName = containerName;
  }


  public ClusterHistoryEventItem displayName(String displayName) {
    
    this.displayName = displayName;
    return this;
  }

   /**
   * Get displayName
   * @return displayName
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public String getDisplayName() {
    return displayName;
  }


  public void setDisplayName(String displayName) {
    this.displayName = displayName;
  }


  public ClusterHistoryEventItem status(String status) {
    
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


  public ClusterHistoryEventItem reason(String reason) {
    
    this.reason = reason;
    return this;
  }

   /**
   * Get reason
   * @return reason
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public String getReason() {
    return reason;
  }


  public void setReason(String reason) {
    this.reason = reason;
  }


  public ClusterHistoryEventItem percent(Long percent) {
    
    this.percent = percent;
    return this;
  }

   /**
   * Get percent
   * @return percent
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public Long getPercent() {
    return percent;
  }


  public void setPercent(Long percent) {
    this.percent = percent;
  }


  public ClusterHistoryEventItem statusEta(Long statusEta) {
    
    this.statusEta = statusEta;
    return this;
  }

   /**
   * Get statusEta
   * @return statusEta
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public Long getStatusEta() {
    return statusEta;
  }


  public void setStatusEta(Long statusEta) {
    this.statusEta = statusEta;
  }


  public ClusterHistoryEventItem message(String message) {
    
    this.message = message;
    return this;
  }

   /**
   * Get message
   * @return message
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public String getMessage() {
    return message;
  }


  public void setMessage(String message) {
    this.message = message;
  }


  public ClusterHistoryEventItem output(String output) {
    
    this.output = output;
    return this;
  }

   /**
   * Get output
   * @return output
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public String getOutput() {
    return output;
  }


  public void setOutput(String output) {
    this.output = output;
  }


  public ClusterHistoryEventItem error(String error) {
    
    this.error = error;
    return this;
  }

   /**
   * Get error
   * @return error
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public String getError() {
    return error;
  }


  public void setError(String error) {
    this.error = error;
  }


  public ClusterHistoryEventItem startDate(OffsetDateTime startDate) {
    
    this.startDate = startDate;
    return this;
  }

   /**
   * Get startDate
   * @return startDate
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public OffsetDateTime getStartDate() {
    return startDate;
  }


  public void setStartDate(OffsetDateTime startDate) {
    this.startDate = startDate;
  }


  public ClusterHistoryEventItem endDate(OffsetDateTime endDate) {
    
    this.endDate = endDate;
    return this;
  }

   /**
   * Get endDate
   * @return endDate
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public OffsetDateTime getEndDate() {
    return endDate;
  }


  public void setEndDate(OffsetDateTime endDate) {
    this.endDate = endDate;
  }


  public ClusterHistoryEventItem duration(Long duration) {
    
    this.duration = duration;
    return this;
  }

   /**
   * Get duration
   * @return duration
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public Long getDuration() {
    return duration;
  }


  public void setDuration(Long duration) {
    this.duration = duration;
  }


  public ClusterHistoryEventItem dateCreated(OffsetDateTime dateCreated) {
    
    this.dateCreated = dateCreated;
    return this;
  }

   /**
   * Get dateCreated
   * @return dateCreated
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public OffsetDateTime getDateCreated() {
    return dateCreated;
  }


  public void setDateCreated(OffsetDateTime dateCreated) {
    this.dateCreated = dateCreated;
  }


  public ClusterHistoryEventItem lastUpdated(OffsetDateTime lastUpdated) {
    
    this.lastUpdated = lastUpdated;
    return this;
  }

   /**
   * Get lastUpdated
   * @return lastUpdated
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public OffsetDateTime getLastUpdated() {
    return lastUpdated;
  }


  public void setLastUpdated(OffsetDateTime lastUpdated) {
    this.lastUpdated = lastUpdated;
  }


  public ClusterHistoryEventItem createdBy(ClusterHistoryCreatedBy createdBy) {
    
    this.createdBy = createdBy;
    return this;
  }

   /**
   * Get createdBy
   * @return createdBy
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public ClusterHistoryCreatedBy getCreatedBy() {
    return createdBy;
  }


  public void setCreatedBy(ClusterHistoryCreatedBy createdBy) {
    this.createdBy = createdBy;
  }


  public ClusterHistoryEventItem updatedBy(ClusterHistoryCreatedBy updatedBy) {
    
    this.updatedBy = updatedBy;
    return this;
  }

   /**
   * Get updatedBy
   * @return updatedBy
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public ClusterHistoryCreatedBy getUpdatedBy() {
    return updatedBy;
  }


  public void setUpdatedBy(ClusterHistoryCreatedBy updatedBy) {
    this.updatedBy = updatedBy;
  }


  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (o == null || getClass() != o.getClass()) {
      return false;
    }
    ClusterHistoryEventItem clusterHistoryEventItem = (ClusterHistoryEventItem) o;
    return Objects.equals(this.id, clusterHistoryEventItem.id) &&
        Objects.equals(this.processId, clusterHistoryEventItem.processId) &&
        Objects.equals(this.accountId, clusterHistoryEventItem.accountId) &&
        Objects.equals(this.uniqueId, clusterHistoryEventItem.uniqueId) &&
        Objects.equals(this.processType, clusterHistoryEventItem.processType) &&
        Objects.equals(this.description, clusterHistoryEventItem.description) &&
        Objects.equals(this.refType, clusterHistoryEventItem.refType) &&
        Objects.equals(this.refId, clusterHistoryEventItem.refId) &&
        Objects.equals(this.subType, clusterHistoryEventItem.subType) &&
        Objects.equals(this.subId, clusterHistoryEventItem.subId) &&
        Objects.equals(this.zoneId, clusterHistoryEventItem.zoneId) &&
        Objects.equals(this.integrationId, clusterHistoryEventItem.integrationId) &&
        Objects.equals(this.instanceId, clusterHistoryEventItem.instanceId) &&
        Objects.equals(this.containerId, clusterHistoryEventItem.containerId) &&
        Objects.equals(this.serverId, clusterHistoryEventItem.serverId) &&
        Objects.equals(this.containerName, clusterHistoryEventItem.containerName) &&
        Objects.equals(this.displayName, clusterHistoryEventItem.displayName) &&
        Objects.equals(this.status, clusterHistoryEventItem.status) &&
        Objects.equals(this.reason, clusterHistoryEventItem.reason) &&
        Objects.equals(this.percent, clusterHistoryEventItem.percent) &&
        Objects.equals(this.statusEta, clusterHistoryEventItem.statusEta) &&
        Objects.equals(this.message, clusterHistoryEventItem.message) &&
        Objects.equals(this.output, clusterHistoryEventItem.output) &&
        Objects.equals(this.error, clusterHistoryEventItem.error) &&
        Objects.equals(this.startDate, clusterHistoryEventItem.startDate) &&
        Objects.equals(this.endDate, clusterHistoryEventItem.endDate) &&
        Objects.equals(this.duration, clusterHistoryEventItem.duration) &&
        Objects.equals(this.dateCreated, clusterHistoryEventItem.dateCreated) &&
        Objects.equals(this.lastUpdated, clusterHistoryEventItem.lastUpdated) &&
        Objects.equals(this.createdBy, clusterHistoryEventItem.createdBy) &&
        Objects.equals(this.updatedBy, clusterHistoryEventItem.updatedBy);
  }

  @Override
  public int hashCode() {
    return Objects.hash(id, processId, accountId, uniqueId, processType, description, refType, refId, subType, subId, zoneId, integrationId, instanceId, containerId, serverId, containerName, displayName, status, reason, percent, statusEta, message, output, error, startDate, endDate, duration, dateCreated, lastUpdated, createdBy, updatedBy);
  }


  @Override
  public String toString() {
    StringBuilder sb = new StringBuilder();
    sb.append("class ClusterHistoryEventItem {\n");
    sb.append("    id: ").append(toIndentedString(id)).append("\n");
    sb.append("    processId: ").append(toIndentedString(processId)).append("\n");
    sb.append("    accountId: ").append(toIndentedString(accountId)).append("\n");
    sb.append("    uniqueId: ").append(toIndentedString(uniqueId)).append("\n");
    sb.append("    processType: ").append(toIndentedString(processType)).append("\n");
    sb.append("    description: ").append(toIndentedString(description)).append("\n");
    sb.append("    refType: ").append(toIndentedString(refType)).append("\n");
    sb.append("    refId: ").append(toIndentedString(refId)).append("\n");
    sb.append("    subType: ").append(toIndentedString(subType)).append("\n");
    sb.append("    subId: ").append(toIndentedString(subId)).append("\n");
    sb.append("    zoneId: ").append(toIndentedString(zoneId)).append("\n");
    sb.append("    integrationId: ").append(toIndentedString(integrationId)).append("\n");
    sb.append("    instanceId: ").append(toIndentedString(instanceId)).append("\n");
    sb.append("    containerId: ").append(toIndentedString(containerId)).append("\n");
    sb.append("    serverId: ").append(toIndentedString(serverId)).append("\n");
    sb.append("    containerName: ").append(toIndentedString(containerName)).append("\n");
    sb.append("    displayName: ").append(toIndentedString(displayName)).append("\n");
    sb.append("    status: ").append(toIndentedString(status)).append("\n");
    sb.append("    reason: ").append(toIndentedString(reason)).append("\n");
    sb.append("    percent: ").append(toIndentedString(percent)).append("\n");
    sb.append("    statusEta: ").append(toIndentedString(statusEta)).append("\n");
    sb.append("    message: ").append(toIndentedString(message)).append("\n");
    sb.append("    output: ").append(toIndentedString(output)).append("\n");
    sb.append("    error: ").append(toIndentedString(error)).append("\n");
    sb.append("    startDate: ").append(toIndentedString(startDate)).append("\n");
    sb.append("    endDate: ").append(toIndentedString(endDate)).append("\n");
    sb.append("    duration: ").append(toIndentedString(duration)).append("\n");
    sb.append("    dateCreated: ").append(toIndentedString(dateCreated)).append("\n");
    sb.append("    lastUpdated: ").append(toIndentedString(lastUpdated)).append("\n");
    sb.append("    createdBy: ").append(toIndentedString(createdBy)).append("\n");
    sb.append("    updatedBy: ").append(toIndentedString(updatedBy)).append("\n");
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

