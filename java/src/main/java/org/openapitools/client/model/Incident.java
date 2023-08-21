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
import java.util.ArrayList;
import java.util.List;
import org.openapitools.client.model.ApiBlueprintsIdUpdatePermissionsResourcePermissionSites;
import org.openapitools.client.model.Check;
import org.openapitools.client.model.InlineResponse20040AppDeployInstance;
import org.threeten.bp.OffsetDateTime;

/**
 * Incident
 */
@javax.annotation.Generated(value = "org.openapitools.codegen.languages.JavaClientCodegen", date = "2023-08-21T11:26:47.028Z[GMT]")
public class Incident {
  public static final String SERIALIZED_NAME_ID = "id";
  @SerializedName(SERIALIZED_NAME_ID)
  private Long id;

  public static final String SERIALIZED_NAME_ACCOUNT = "account";
  @SerializedName(SERIALIZED_NAME_ACCOUNT)
  private ApiBlueprintsIdUpdatePermissionsResourcePermissionSites account;

  public static final String SERIALIZED_NAME_APP = "app";
  @SerializedName(SERIALIZED_NAME_APP)
  private String app;

  public static final String SERIALIZED_NAME_AUTO_CLOSE = "autoClose";
  @SerializedName(SERIALIZED_NAME_AUTO_CLOSE)
  private Boolean autoClose;

  public static final String SERIALIZED_NAME_CHANNEL_ID = "channelId";
  @SerializedName(SERIALIZED_NAME_CHANNEL_ID)
  private String channelId;

  public static final String SERIALIZED_NAME_CHECK_GROUPS = "checkGroups";
  @SerializedName(SERIALIZED_NAME_CHECK_GROUPS)
  private List<InlineResponse20040AppDeployInstance> checkGroups = null;

  public static final String SERIALIZED_NAME_CHECKS = "checks";
  @SerializedName(SERIALIZED_NAME_CHECKS)
  private List<Check> checks = null;

  public static final String SERIALIZED_NAME_COMMENT = "comment";
  @SerializedName(SERIALIZED_NAME_COMMENT)
  private String comment;

  public static final String SERIALIZED_NAME_DISPLAY_NAME = "displayName";
  @SerializedName(SERIALIZED_NAME_DISPLAY_NAME)
  private String displayName;

  public static final String SERIALIZED_NAME_DURATION = "duration";
  @SerializedName(SERIALIZED_NAME_DURATION)
  private String duration;

  public static final String SERIALIZED_NAME_END_DATE = "endDate";
  @SerializedName(SERIALIZED_NAME_END_DATE)
  private OffsetDateTime endDate;

  public static final String SERIALIZED_NAME_IN_UPTIME = "inUptime";
  @SerializedName(SERIALIZED_NAME_IN_UPTIME)
  private Boolean inUptime;

  public static final String SERIALIZED_NAME_MUTED = "muted";
  @SerializedName(SERIALIZED_NAME_MUTED)
  private Boolean muted;

  public static final String SERIALIZED_NAME_LAST_CHECK_TIME = "lastCheckTime";
  @SerializedName(SERIALIZED_NAME_LAST_CHECK_TIME)
  private OffsetDateTime lastCheckTime;

  public static final String SERIALIZED_NAME_LAST_ERROR = "lastError";
  @SerializedName(SERIALIZED_NAME_LAST_ERROR)
  private String lastError;

  public static final String SERIALIZED_NAME_LAST_MESSAGE = "lastMessage";
  @SerializedName(SERIALIZED_NAME_LAST_MESSAGE)
  private String lastMessage;

  public static final String SERIALIZED_NAME_NAME = "name";
  @SerializedName(SERIALIZED_NAME_NAME)
  private String name;

  public static final String SERIALIZED_NAME_RESOLUTION = "resolution";
  @SerializedName(SERIALIZED_NAME_RESOLUTION)
  private String resolution;

  public static final String SERIALIZED_NAME_SEVERITY = "severity";
  @SerializedName(SERIALIZED_NAME_SEVERITY)
  private String severity;

  public static final String SERIALIZED_NAME_SEVERITY_ID = "severityId";
  @SerializedName(SERIALIZED_NAME_SEVERITY_ID)
  private Long severityId;

  public static final String SERIALIZED_NAME_START_DATE = "startDate";
  @SerializedName(SERIALIZED_NAME_START_DATE)
  private OffsetDateTime startDate;

  public static final String SERIALIZED_NAME_STATUS = "status";
  @SerializedName(SERIALIZED_NAME_STATUS)
  private String status;

  public static final String SERIALIZED_NAME_VISIBILITY = "visibility";
  @SerializedName(SERIALIZED_NAME_VISIBILITY)
  private String visibility;


  public Incident id(Long id) {
    
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


  public Incident account(ApiBlueprintsIdUpdatePermissionsResourcePermissionSites account) {
    
    this.account = account;
    return this;
  }

   /**
   * Get account
   * @return account
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public ApiBlueprintsIdUpdatePermissionsResourcePermissionSites getAccount() {
    return account;
  }


  public void setAccount(ApiBlueprintsIdUpdatePermissionsResourcePermissionSites account) {
    this.account = account;
  }


  public Incident app(String app) {
    
    this.app = app;
    return this;
  }

   /**
   * Get app
   * @return app
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public String getApp() {
    return app;
  }


  public void setApp(String app) {
    this.app = app;
  }


  public Incident autoClose(Boolean autoClose) {
    
    this.autoClose = autoClose;
    return this;
  }

   /**
   * Get autoClose
   * @return autoClose
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public Boolean getAutoClose() {
    return autoClose;
  }


  public void setAutoClose(Boolean autoClose) {
    this.autoClose = autoClose;
  }


  public Incident channelId(String channelId) {
    
    this.channelId = channelId;
    return this;
  }

   /**
   * Get channelId
   * @return channelId
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public String getChannelId() {
    return channelId;
  }


  public void setChannelId(String channelId) {
    this.channelId = channelId;
  }


  public Incident checkGroups(List<InlineResponse20040AppDeployInstance> checkGroups) {
    
    this.checkGroups = checkGroups;
    return this;
  }

  public Incident addCheckGroupsItem(InlineResponse20040AppDeployInstance checkGroupsItem) {
    if (this.checkGroups == null) {
      this.checkGroups = new ArrayList<InlineResponse20040AppDeployInstance>();
    }
    this.checkGroups.add(checkGroupsItem);
    return this;
  }

   /**
   * Get checkGroups
   * @return checkGroups
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public List<InlineResponse20040AppDeployInstance> getCheckGroups() {
    return checkGroups;
  }


  public void setCheckGroups(List<InlineResponse20040AppDeployInstance> checkGroups) {
    this.checkGroups = checkGroups;
  }


  public Incident checks(List<Check> checks) {
    
    this.checks = checks;
    return this;
  }

  public Incident addChecksItem(Check checksItem) {
    if (this.checks == null) {
      this.checks = new ArrayList<Check>();
    }
    this.checks.add(checksItem);
    return this;
  }

   /**
   * Get checks
   * @return checks
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public List<Check> getChecks() {
    return checks;
  }


  public void setChecks(List<Check> checks) {
    this.checks = checks;
  }


  public Incident comment(String comment) {
    
    this.comment = comment;
    return this;
  }

   /**
   * Get comment
   * @return comment
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public String getComment() {
    return comment;
  }


  public void setComment(String comment) {
    this.comment = comment;
  }


  public Incident displayName(String displayName) {
    
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


  public Incident duration(String duration) {
    
    this.duration = duration;
    return this;
  }

   /**
   * Get duration
   * @return duration
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public String getDuration() {
    return duration;
  }


  public void setDuration(String duration) {
    this.duration = duration;
  }


  public Incident endDate(OffsetDateTime endDate) {
    
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


  public Incident inUptime(Boolean inUptime) {
    
    this.inUptime = inUptime;
    return this;
  }

   /**
   * Get inUptime
   * @return inUptime
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public Boolean getInUptime() {
    return inUptime;
  }


  public void setInUptime(Boolean inUptime) {
    this.inUptime = inUptime;
  }


  public Incident muted(Boolean muted) {
    
    this.muted = muted;
    return this;
  }

   /**
   * Get muted
   * @return muted
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public Boolean getMuted() {
    return muted;
  }


  public void setMuted(Boolean muted) {
    this.muted = muted;
  }


  public Incident lastCheckTime(OffsetDateTime lastCheckTime) {
    
    this.lastCheckTime = lastCheckTime;
    return this;
  }

   /**
   * Get lastCheckTime
   * @return lastCheckTime
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public OffsetDateTime getLastCheckTime() {
    return lastCheckTime;
  }


  public void setLastCheckTime(OffsetDateTime lastCheckTime) {
    this.lastCheckTime = lastCheckTime;
  }


  public Incident lastError(String lastError) {
    
    this.lastError = lastError;
    return this;
  }

   /**
   * Get lastError
   * @return lastError
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public String getLastError() {
    return lastError;
  }


  public void setLastError(String lastError) {
    this.lastError = lastError;
  }


  public Incident lastMessage(String lastMessage) {
    
    this.lastMessage = lastMessage;
    return this;
  }

   /**
   * Get lastMessage
   * @return lastMessage
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public String getLastMessage() {
    return lastMessage;
  }


  public void setLastMessage(String lastMessage) {
    this.lastMessage = lastMessage;
  }


  public Incident name(String name) {
    
    this.name = name;
    return this;
  }

   /**
   * Get name
   * @return name
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public String getName() {
    return name;
  }


  public void setName(String name) {
    this.name = name;
  }


  public Incident resolution(String resolution) {
    
    this.resolution = resolution;
    return this;
  }

   /**
   * Get resolution
   * @return resolution
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public String getResolution() {
    return resolution;
  }


  public void setResolution(String resolution) {
    this.resolution = resolution;
  }


  public Incident severity(String severity) {
    
    this.severity = severity;
    return this;
  }

   /**
   * Get severity
   * @return severity
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public String getSeverity() {
    return severity;
  }


  public void setSeverity(String severity) {
    this.severity = severity;
  }


  public Incident severityId(Long severityId) {
    
    this.severityId = severityId;
    return this;
  }

   /**
   * Get severityId
   * @return severityId
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public Long getSeverityId() {
    return severityId;
  }


  public void setSeverityId(Long severityId) {
    this.severityId = severityId;
  }


  public Incident startDate(OffsetDateTime startDate) {
    
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


  public Incident status(String status) {
    
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


  public Incident visibility(String visibility) {
    
    this.visibility = visibility;
    return this;
  }

   /**
   * Get visibility
   * @return visibility
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public String getVisibility() {
    return visibility;
  }


  public void setVisibility(String visibility) {
    this.visibility = visibility;
  }


  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (o == null || getClass() != o.getClass()) {
      return false;
    }
    Incident incident = (Incident) o;
    return Objects.equals(this.id, incident.id) &&
        Objects.equals(this.account, incident.account) &&
        Objects.equals(this.app, incident.app) &&
        Objects.equals(this.autoClose, incident.autoClose) &&
        Objects.equals(this.channelId, incident.channelId) &&
        Objects.equals(this.checkGroups, incident.checkGroups) &&
        Objects.equals(this.checks, incident.checks) &&
        Objects.equals(this.comment, incident.comment) &&
        Objects.equals(this.displayName, incident.displayName) &&
        Objects.equals(this.duration, incident.duration) &&
        Objects.equals(this.endDate, incident.endDate) &&
        Objects.equals(this.inUptime, incident.inUptime) &&
        Objects.equals(this.muted, incident.muted) &&
        Objects.equals(this.lastCheckTime, incident.lastCheckTime) &&
        Objects.equals(this.lastError, incident.lastError) &&
        Objects.equals(this.lastMessage, incident.lastMessage) &&
        Objects.equals(this.name, incident.name) &&
        Objects.equals(this.resolution, incident.resolution) &&
        Objects.equals(this.severity, incident.severity) &&
        Objects.equals(this.severityId, incident.severityId) &&
        Objects.equals(this.startDate, incident.startDate) &&
        Objects.equals(this.status, incident.status) &&
        Objects.equals(this.visibility, incident.visibility);
  }

  @Override
  public int hashCode() {
    return Objects.hash(id, account, app, autoClose, channelId, checkGroups, checks, comment, displayName, duration, endDate, inUptime, muted, lastCheckTime, lastError, lastMessage, name, resolution, severity, severityId, startDate, status, visibility);
  }


  @Override
  public String toString() {
    StringBuilder sb = new StringBuilder();
    sb.append("class Incident {\n");
    sb.append("    id: ").append(toIndentedString(id)).append("\n");
    sb.append("    account: ").append(toIndentedString(account)).append("\n");
    sb.append("    app: ").append(toIndentedString(app)).append("\n");
    sb.append("    autoClose: ").append(toIndentedString(autoClose)).append("\n");
    sb.append("    channelId: ").append(toIndentedString(channelId)).append("\n");
    sb.append("    checkGroups: ").append(toIndentedString(checkGroups)).append("\n");
    sb.append("    checks: ").append(toIndentedString(checks)).append("\n");
    sb.append("    comment: ").append(toIndentedString(comment)).append("\n");
    sb.append("    displayName: ").append(toIndentedString(displayName)).append("\n");
    sb.append("    duration: ").append(toIndentedString(duration)).append("\n");
    sb.append("    endDate: ").append(toIndentedString(endDate)).append("\n");
    sb.append("    inUptime: ").append(toIndentedString(inUptime)).append("\n");
    sb.append("    muted: ").append(toIndentedString(muted)).append("\n");
    sb.append("    lastCheckTime: ").append(toIndentedString(lastCheckTime)).append("\n");
    sb.append("    lastError: ").append(toIndentedString(lastError)).append("\n");
    sb.append("    lastMessage: ").append(toIndentedString(lastMessage)).append("\n");
    sb.append("    name: ").append(toIndentedString(name)).append("\n");
    sb.append("    resolution: ").append(toIndentedString(resolution)).append("\n");
    sb.append("    severity: ").append(toIndentedString(severity)).append("\n");
    sb.append("    severityId: ").append(toIndentedString(severityId)).append("\n");
    sb.append("    startDate: ").append(toIndentedString(startDate)).append("\n");
    sb.append("    status: ").append(toIndentedString(status)).append("\n");
    sb.append("    visibility: ").append(toIndentedString(visibility)).append("\n");
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
