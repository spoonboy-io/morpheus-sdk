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
import org.openapitools.client.model.BackupJobBackups;
import org.openapitools.client.model.InlineResponse200108NetworkFloatingIpCreatedBy;
import org.threeten.bp.OffsetDateTime;

/**
 * BackupResult
 */
@javax.annotation.Generated(value = "org.openapitools.codegen.languages.JavaClientCodegen", date = "2023-08-21T11:26:47.028Z[GMT]")
public class BackupResult {
  public static final String SERIALIZED_NAME_ID = "id";
  @SerializedName(SERIALIZED_NAME_ID)
  private Long id;

  public static final String SERIALIZED_NAME_BACKUP = "backup";
  @SerializedName(SERIALIZED_NAME_BACKUP)
  private BackupJobBackups backup;

  public static final String SERIALIZED_NAME_BACKUP_SET_ID = "backupSetId";
  @SerializedName(SERIALIZED_NAME_BACKUP_SET_ID)
  private String backupSetId;

  public static final String SERIALIZED_NAME_INSTANCE_ID = "instanceId";
  @SerializedName(SERIALIZED_NAME_INSTANCE_ID)
  private Long instanceId;

  public static final String SERIALIZED_NAME_CONTAINER_ID = "containerId";
  @SerializedName(SERIALIZED_NAME_CONTAINER_ID)
  private Long containerId;

  public static final String SERIALIZED_NAME_SERVER_ID = "serverId";
  @SerializedName(SERIALIZED_NAME_SERVER_ID)
  private Long serverId;

  public static final String SERIALIZED_NAME_STATUS = "status";
  @SerializedName(SERIALIZED_NAME_STATUS)
  private String status;

  public static final String SERIALIZED_NAME_ERROR_MESSAGE = "errorMessage";
  @SerializedName(SERIALIZED_NAME_ERROR_MESSAGE)
  private String errorMessage;

  public static final String SERIALIZED_NAME_START_DATE = "startDate";
  @SerializedName(SERIALIZED_NAME_START_DATE)
  private OffsetDateTime startDate;

  public static final String SERIALIZED_NAME_END_DATE = "endDate";
  @SerializedName(SERIALIZED_NAME_END_DATE)
  private OffsetDateTime endDate;

  public static final String SERIALIZED_NAME_DURATION_MILLIS = "durationMillis";
  @SerializedName(SERIALIZED_NAME_DURATION_MILLIS)
  private Long durationMillis;

  public static final String SERIALIZED_NAME_SIZE_IN_BYTES = "sizeInBytes";
  @SerializedName(SERIALIZED_NAME_SIZE_IN_BYTES)
  private Long sizeInBytes;

  public static final String SERIALIZED_NAME_SIZE_IN_MB = "sizeInMb";
  @SerializedName(SERIALIZED_NAME_SIZE_IN_MB)
  private Long sizeInMb;

  public static final String SERIALIZED_NAME_VOLUME_PATH = "volumePath";
  @SerializedName(SERIALIZED_NAME_VOLUME_PATH)
  private String volumePath;

  public static final String SERIALIZED_NAME_RESULT_ARCHIVE = "resultArchive";
  @SerializedName(SERIALIZED_NAME_RESULT_ARCHIVE)
  private String resultArchive;

  public static final String SERIALIZED_NAME_RESULT_PATH = "resultPath";
  @SerializedName(SERIALIZED_NAME_RESULT_PATH)
  private String resultPath;

  public static final String SERIALIZED_NAME_EXTERNAL_ID = "externalId";
  @SerializedName(SERIALIZED_NAME_EXTERNAL_ID)
  private String externalId;

  public static final String SERIALIZED_NAME_SNAPSHOT_ID = "snapshotId";
  @SerializedName(SERIALIZED_NAME_SNAPSHOT_ID)
  private String snapshotId;

  public static final String SERIALIZED_NAME_SNAPSHOT_EXTERNAL_ID = "snapshotExternalId";
  @SerializedName(SERIALIZED_NAME_SNAPSHOT_EXTERNAL_ID)
  private String snapshotExternalId;

  public static final String SERIALIZED_NAME_CREATED_BY = "createdBy";
  @SerializedName(SERIALIZED_NAME_CREATED_BY)
  private InlineResponse200108NetworkFloatingIpCreatedBy createdBy;

  public static final String SERIALIZED_NAME_DATE_CREATED = "dateCreated";
  @SerializedName(SERIALIZED_NAME_DATE_CREATED)
  private OffsetDateTime dateCreated;

  public static final String SERIALIZED_NAME_LAST_UPDATED = "lastUpdated";
  @SerializedName(SERIALIZED_NAME_LAST_UPDATED)
  private OffsetDateTime lastUpdated;


  public BackupResult id(Long id) {
    
    this.id = id;
    return this;
  }

   /**
   * Backup Result ID
   * @return id
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "Backup Result ID")

  public Long getId() {
    return id;
  }


  public void setId(Long id) {
    this.id = id;
  }


  public BackupResult backup(BackupJobBackups backup) {
    
    this.backup = backup;
    return this;
  }

   /**
   * Get backup
   * @return backup
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public BackupJobBackups getBackup() {
    return backup;
  }


  public void setBackup(BackupJobBackups backup) {
    this.backup = backup;
  }


  public BackupResult backupSetId(String backupSetId) {
    
    this.backupSetId = backupSetId;
    return this;
  }

   /**
   * Get backupSetId
   * @return backupSetId
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public String getBackupSetId() {
    return backupSetId;
  }


  public void setBackupSetId(String backupSetId) {
    this.backupSetId = backupSetId;
  }


  public BackupResult instanceId(Long instanceId) {
    
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


  public BackupResult containerId(Long containerId) {
    
    this.containerId = containerId;
    return this;
  }

   /**
   * Get containerId
   * @return containerId
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public Long getContainerId() {
    return containerId;
  }


  public void setContainerId(Long containerId) {
    this.containerId = containerId;
  }


  public BackupResult serverId(Long serverId) {
    
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


  public BackupResult status(String status) {
    
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


  public BackupResult errorMessage(String errorMessage) {
    
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


  public BackupResult startDate(OffsetDateTime startDate) {
    
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


  public BackupResult endDate(OffsetDateTime endDate) {
    
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


  public BackupResult durationMillis(Long durationMillis) {
    
    this.durationMillis = durationMillis;
    return this;
  }

   /**
   * Get durationMillis
   * @return durationMillis
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public Long getDurationMillis() {
    return durationMillis;
  }


  public void setDurationMillis(Long durationMillis) {
    this.durationMillis = durationMillis;
  }


  public BackupResult sizeInBytes(Long sizeInBytes) {
    
    this.sizeInBytes = sizeInBytes;
    return this;
  }

   /**
   * Get sizeInBytes
   * @return sizeInBytes
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public Long getSizeInBytes() {
    return sizeInBytes;
  }


  public void setSizeInBytes(Long sizeInBytes) {
    this.sizeInBytes = sizeInBytes;
  }


  public BackupResult sizeInMb(Long sizeInMb) {
    
    this.sizeInMb = sizeInMb;
    return this;
  }

   /**
   * Get sizeInMb
   * @return sizeInMb
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public Long getSizeInMb() {
    return sizeInMb;
  }


  public void setSizeInMb(Long sizeInMb) {
    this.sizeInMb = sizeInMb;
  }


  public BackupResult volumePath(String volumePath) {
    
    this.volumePath = volumePath;
    return this;
  }

   /**
   * Get volumePath
   * @return volumePath
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public String getVolumePath() {
    return volumePath;
  }


  public void setVolumePath(String volumePath) {
    this.volumePath = volumePath;
  }


  public BackupResult resultArchive(String resultArchive) {
    
    this.resultArchive = resultArchive;
    return this;
  }

   /**
   * Get resultArchive
   * @return resultArchive
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public String getResultArchive() {
    return resultArchive;
  }


  public void setResultArchive(String resultArchive) {
    this.resultArchive = resultArchive;
  }


  public BackupResult resultPath(String resultPath) {
    
    this.resultPath = resultPath;
    return this;
  }

   /**
   * Get resultPath
   * @return resultPath
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public String getResultPath() {
    return resultPath;
  }


  public void setResultPath(String resultPath) {
    this.resultPath = resultPath;
  }


  public BackupResult externalId(String externalId) {
    
    this.externalId = externalId;
    return this;
  }

   /**
   * Get externalId
   * @return externalId
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public String getExternalId() {
    return externalId;
  }


  public void setExternalId(String externalId) {
    this.externalId = externalId;
  }


  public BackupResult snapshotId(String snapshotId) {
    
    this.snapshotId = snapshotId;
    return this;
  }

   /**
   * Get snapshotId
   * @return snapshotId
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public String getSnapshotId() {
    return snapshotId;
  }


  public void setSnapshotId(String snapshotId) {
    this.snapshotId = snapshotId;
  }


  public BackupResult snapshotExternalId(String snapshotExternalId) {
    
    this.snapshotExternalId = snapshotExternalId;
    return this;
  }

   /**
   * Get snapshotExternalId
   * @return snapshotExternalId
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public String getSnapshotExternalId() {
    return snapshotExternalId;
  }


  public void setSnapshotExternalId(String snapshotExternalId) {
    this.snapshotExternalId = snapshotExternalId;
  }


  public BackupResult createdBy(InlineResponse200108NetworkFloatingIpCreatedBy createdBy) {
    
    this.createdBy = createdBy;
    return this;
  }

   /**
   * Get createdBy
   * @return createdBy
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public InlineResponse200108NetworkFloatingIpCreatedBy getCreatedBy() {
    return createdBy;
  }


  public void setCreatedBy(InlineResponse200108NetworkFloatingIpCreatedBy createdBy) {
    this.createdBy = createdBy;
  }


  public BackupResult dateCreated(OffsetDateTime dateCreated) {
    
    this.dateCreated = dateCreated;
    return this;
  }

   /**
   * Date Created
   * @return dateCreated
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "Date Created")

  public OffsetDateTime getDateCreated() {
    return dateCreated;
  }


  public void setDateCreated(OffsetDateTime dateCreated) {
    this.dateCreated = dateCreated;
  }


  public BackupResult lastUpdated(OffsetDateTime lastUpdated) {
    
    this.lastUpdated = lastUpdated;
    return this;
  }

   /**
   * Last Updated
   * @return lastUpdated
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "Last Updated")

  public OffsetDateTime getLastUpdated() {
    return lastUpdated;
  }


  public void setLastUpdated(OffsetDateTime lastUpdated) {
    this.lastUpdated = lastUpdated;
  }


  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (o == null || getClass() != o.getClass()) {
      return false;
    }
    BackupResult backupResult = (BackupResult) o;
    return Objects.equals(this.id, backupResult.id) &&
        Objects.equals(this.backup, backupResult.backup) &&
        Objects.equals(this.backupSetId, backupResult.backupSetId) &&
        Objects.equals(this.instanceId, backupResult.instanceId) &&
        Objects.equals(this.containerId, backupResult.containerId) &&
        Objects.equals(this.serverId, backupResult.serverId) &&
        Objects.equals(this.status, backupResult.status) &&
        Objects.equals(this.errorMessage, backupResult.errorMessage) &&
        Objects.equals(this.startDate, backupResult.startDate) &&
        Objects.equals(this.endDate, backupResult.endDate) &&
        Objects.equals(this.durationMillis, backupResult.durationMillis) &&
        Objects.equals(this.sizeInBytes, backupResult.sizeInBytes) &&
        Objects.equals(this.sizeInMb, backupResult.sizeInMb) &&
        Objects.equals(this.volumePath, backupResult.volumePath) &&
        Objects.equals(this.resultArchive, backupResult.resultArchive) &&
        Objects.equals(this.resultPath, backupResult.resultPath) &&
        Objects.equals(this.externalId, backupResult.externalId) &&
        Objects.equals(this.snapshotId, backupResult.snapshotId) &&
        Objects.equals(this.snapshotExternalId, backupResult.snapshotExternalId) &&
        Objects.equals(this.createdBy, backupResult.createdBy) &&
        Objects.equals(this.dateCreated, backupResult.dateCreated) &&
        Objects.equals(this.lastUpdated, backupResult.lastUpdated);
  }

  @Override
  public int hashCode() {
    return Objects.hash(id, backup, backupSetId, instanceId, containerId, serverId, status, errorMessage, startDate, endDate, durationMillis, sizeInBytes, sizeInMb, volumePath, resultArchive, resultPath, externalId, snapshotId, snapshotExternalId, createdBy, dateCreated, lastUpdated);
  }


  @Override
  public String toString() {
    StringBuilder sb = new StringBuilder();
    sb.append("class BackupResult {\n");
    sb.append("    id: ").append(toIndentedString(id)).append("\n");
    sb.append("    backup: ").append(toIndentedString(backup)).append("\n");
    sb.append("    backupSetId: ").append(toIndentedString(backupSetId)).append("\n");
    sb.append("    instanceId: ").append(toIndentedString(instanceId)).append("\n");
    sb.append("    containerId: ").append(toIndentedString(containerId)).append("\n");
    sb.append("    serverId: ").append(toIndentedString(serverId)).append("\n");
    sb.append("    status: ").append(toIndentedString(status)).append("\n");
    sb.append("    errorMessage: ").append(toIndentedString(errorMessage)).append("\n");
    sb.append("    startDate: ").append(toIndentedString(startDate)).append("\n");
    sb.append("    endDate: ").append(toIndentedString(endDate)).append("\n");
    sb.append("    durationMillis: ").append(toIndentedString(durationMillis)).append("\n");
    sb.append("    sizeInBytes: ").append(toIndentedString(sizeInBytes)).append("\n");
    sb.append("    sizeInMb: ").append(toIndentedString(sizeInMb)).append("\n");
    sb.append("    volumePath: ").append(toIndentedString(volumePath)).append("\n");
    sb.append("    resultArchive: ").append(toIndentedString(resultArchive)).append("\n");
    sb.append("    resultPath: ").append(toIndentedString(resultPath)).append("\n");
    sb.append("    externalId: ").append(toIndentedString(externalId)).append("\n");
    sb.append("    snapshotId: ").append(toIndentedString(snapshotId)).append("\n");
    sb.append("    snapshotExternalId: ").append(toIndentedString(snapshotExternalId)).append("\n");
    sb.append("    createdBy: ").append(toIndentedString(createdBy)).append("\n");
    sb.append("    dateCreated: ").append(toIndentedString(dateCreated)).append("\n");
    sb.append("    lastUpdated: ").append(toIndentedString(lastUpdated)).append("\n");
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
