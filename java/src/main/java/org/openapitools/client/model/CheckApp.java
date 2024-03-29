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
import org.openapitools.client.model.InlineResponse200107NetworkPoolCreatedBy;
import org.openapitools.client.model.InlineResponse20082LoadBalancerInstanceSslCert;
import org.threeten.bp.OffsetDateTime;

/**
 * CheckApp
 */
@javax.annotation.Generated(value = "org.openapitools.codegen.languages.JavaClientCodegen", date = "2023-08-21T11:26:47.028Z[GMT]")
public class CheckApp {
  public static final String SERIALIZED_NAME_ID = "id";
  @SerializedName(SERIALIZED_NAME_ID)
  private Long id;

  public static final String SERIALIZED_NAME_ACCOUNT = "account";
  @SerializedName(SERIALIZED_NAME_ACCOUNT)
  private ApiBlueprintsIdUpdatePermissionsResourcePermissionSites account;

  public static final String SERIALIZED_NAME_ACTIVE = "active";
  @SerializedName(SERIALIZED_NAME_ACTIVE)
  private Boolean active;

  public static final String SERIALIZED_NAME_APP = "app";
  @SerializedName(SERIALIZED_NAME_APP)
  private InlineResponse20082LoadBalancerInstanceSslCert app;

  public static final String SERIALIZED_NAME_NAME = "name";
  @SerializedName(SERIALIZED_NAME_NAME)
  private String name;

  public static final String SERIALIZED_NAME_DESCRIPTION = "description";
  @SerializedName(SERIALIZED_NAME_DESCRIPTION)
  private String description;

  public static final String SERIALIZED_NAME_IN_UPTIME = "inUptime";
  @SerializedName(SERIALIZED_NAME_IN_UPTIME)
  private Boolean inUptime;

  public static final String SERIALIZED_NAME_LAST_CHECK_STATUS = "lastCheckStatus";
  @SerializedName(SERIALIZED_NAME_LAST_CHECK_STATUS)
  private String lastCheckStatus;

  public static final String SERIALIZED_NAME_LAST_WARNING_DATE = "lastWarningDate";
  @SerializedName(SERIALIZED_NAME_LAST_WARNING_DATE)
  private OffsetDateTime lastWarningDate;

  public static final String SERIALIZED_NAME_LAST_ERROR_DATE = "lastErrorDate";
  @SerializedName(SERIALIZED_NAME_LAST_ERROR_DATE)
  private OffsetDateTime lastErrorDate;

  public static final String SERIALIZED_NAME_LAST_SUCCESS_DATE = "lastSuccessDate";
  @SerializedName(SERIALIZED_NAME_LAST_SUCCESS_DATE)
  private OffsetDateTime lastSuccessDate;

  public static final String SERIALIZED_NAME_LAST_RUN_DATE = "lastRunDate";
  @SerializedName(SERIALIZED_NAME_LAST_RUN_DATE)
  private OffsetDateTime lastRunDate;

  public static final String SERIALIZED_NAME_LAST_ERROR = "lastError";
  @SerializedName(SERIALIZED_NAME_LAST_ERROR)
  private String lastError;

  public static final String SERIALIZED_NAME_LAST_TIMER = "lastTimer";
  @SerializedName(SERIALIZED_NAME_LAST_TIMER)
  private Long lastTimer;

  public static final String SERIALIZED_NAME_HEALTH = "health";
  @SerializedName(SERIALIZED_NAME_HEALTH)
  private Long health;

  public static final String SERIALIZED_NAME_HISTORY = "history";
  @SerializedName(SERIALIZED_NAME_HISTORY)
  private String history;

  public static final String SERIALIZED_NAME_SEVERITY = "severity";
  @SerializedName(SERIALIZED_NAME_SEVERITY)
  private String severity;

  public static final String SERIALIZED_NAME_CREATE_INCIDENT = "createIncident";
  @SerializedName(SERIALIZED_NAME_CREATE_INCIDENT)
  private Boolean createIncident;

  public static final String SERIALIZED_NAME_MUTED = "muted";
  @SerializedName(SERIALIZED_NAME_MUTED)
  private Boolean muted;

  public static final String SERIALIZED_NAME_CREATED_BY = "createdBy";
  @SerializedName(SERIALIZED_NAME_CREATED_BY)
  private InlineResponse200107NetworkPoolCreatedBy createdBy;

  public static final String SERIALIZED_NAME_DATE_CREATED = "dateCreated";
  @SerializedName(SERIALIZED_NAME_DATE_CREATED)
  private OffsetDateTime dateCreated;

  public static final String SERIALIZED_NAME_LAST_UPDATED = "lastUpdated";
  @SerializedName(SERIALIZED_NAME_LAST_UPDATED)
  private OffsetDateTime lastUpdated;

  public static final String SERIALIZED_NAME_AVAILABILITY = "availability";
  @SerializedName(SERIALIZED_NAME_AVAILABILITY)
  private String availability;

  public static final String SERIALIZED_NAME_CHECKS = "checks";
  @SerializedName(SERIALIZED_NAME_CHECKS)
  private List<Long> checks = null;

  public static final String SERIALIZED_NAME_CHECK_GROUPS = "checkGroups";
  @SerializedName(SERIALIZED_NAME_CHECK_GROUPS)
  private List<Long> checkGroups = null;


  public CheckApp id(Long id) {
    
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


  public CheckApp account(ApiBlueprintsIdUpdatePermissionsResourcePermissionSites account) {
    
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


  public CheckApp active(Boolean active) {
    
    this.active = active;
    return this;
  }

   /**
   * Get active
   * @return active
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public Boolean getActive() {
    return active;
  }


  public void setActive(Boolean active) {
    this.active = active;
  }


  public CheckApp app(InlineResponse20082LoadBalancerInstanceSslCert app) {
    
    this.app = app;
    return this;
  }

   /**
   * Get app
   * @return app
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public InlineResponse20082LoadBalancerInstanceSslCert getApp() {
    return app;
  }


  public void setApp(InlineResponse20082LoadBalancerInstanceSslCert app) {
    this.app = app;
  }


  public CheckApp name(String name) {
    
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


  public CheckApp description(String description) {
    
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


  public CheckApp inUptime(Boolean inUptime) {
    
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


  public CheckApp lastCheckStatus(String lastCheckStatus) {
    
    this.lastCheckStatus = lastCheckStatus;
    return this;
  }

   /**
   * Get lastCheckStatus
   * @return lastCheckStatus
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public String getLastCheckStatus() {
    return lastCheckStatus;
  }


  public void setLastCheckStatus(String lastCheckStatus) {
    this.lastCheckStatus = lastCheckStatus;
  }


  public CheckApp lastWarningDate(OffsetDateTime lastWarningDate) {
    
    this.lastWarningDate = lastWarningDate;
    return this;
  }

   /**
   * Get lastWarningDate
   * @return lastWarningDate
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public OffsetDateTime getLastWarningDate() {
    return lastWarningDate;
  }


  public void setLastWarningDate(OffsetDateTime lastWarningDate) {
    this.lastWarningDate = lastWarningDate;
  }


  public CheckApp lastErrorDate(OffsetDateTime lastErrorDate) {
    
    this.lastErrorDate = lastErrorDate;
    return this;
  }

   /**
   * Get lastErrorDate
   * @return lastErrorDate
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public OffsetDateTime getLastErrorDate() {
    return lastErrorDate;
  }


  public void setLastErrorDate(OffsetDateTime lastErrorDate) {
    this.lastErrorDate = lastErrorDate;
  }


  public CheckApp lastSuccessDate(OffsetDateTime lastSuccessDate) {
    
    this.lastSuccessDate = lastSuccessDate;
    return this;
  }

   /**
   * Get lastSuccessDate
   * @return lastSuccessDate
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public OffsetDateTime getLastSuccessDate() {
    return lastSuccessDate;
  }


  public void setLastSuccessDate(OffsetDateTime lastSuccessDate) {
    this.lastSuccessDate = lastSuccessDate;
  }


  public CheckApp lastRunDate(OffsetDateTime lastRunDate) {
    
    this.lastRunDate = lastRunDate;
    return this;
  }

   /**
   * Get lastRunDate
   * @return lastRunDate
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public OffsetDateTime getLastRunDate() {
    return lastRunDate;
  }


  public void setLastRunDate(OffsetDateTime lastRunDate) {
    this.lastRunDate = lastRunDate;
  }


  public CheckApp lastError(String lastError) {
    
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


  public CheckApp lastTimer(Long lastTimer) {
    
    this.lastTimer = lastTimer;
    return this;
  }

   /**
   * Get lastTimer
   * @return lastTimer
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public Long getLastTimer() {
    return lastTimer;
  }


  public void setLastTimer(Long lastTimer) {
    this.lastTimer = lastTimer;
  }


  public CheckApp health(Long health) {
    
    this.health = health;
    return this;
  }

   /**
   * Get health
   * @return health
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public Long getHealth() {
    return health;
  }


  public void setHealth(Long health) {
    this.health = health;
  }


  public CheckApp history(String history) {
    
    this.history = history;
    return this;
  }

   /**
   * Get history
   * @return history
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public String getHistory() {
    return history;
  }


  public void setHistory(String history) {
    this.history = history;
  }


  public CheckApp severity(String severity) {
    
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


  public CheckApp createIncident(Boolean createIncident) {
    
    this.createIncident = createIncident;
    return this;
  }

   /**
   * Get createIncident
   * @return createIncident
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public Boolean getCreateIncident() {
    return createIncident;
  }


  public void setCreateIncident(Boolean createIncident) {
    this.createIncident = createIncident;
  }


  public CheckApp muted(Boolean muted) {
    
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


  public CheckApp createdBy(InlineResponse200107NetworkPoolCreatedBy createdBy) {
    
    this.createdBy = createdBy;
    return this;
  }

   /**
   * Get createdBy
   * @return createdBy
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public InlineResponse200107NetworkPoolCreatedBy getCreatedBy() {
    return createdBy;
  }


  public void setCreatedBy(InlineResponse200107NetworkPoolCreatedBy createdBy) {
    this.createdBy = createdBy;
  }


  public CheckApp dateCreated(OffsetDateTime dateCreated) {
    
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


  public CheckApp lastUpdated(OffsetDateTime lastUpdated) {
    
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


  public CheckApp availability(String availability) {
    
    this.availability = availability;
    return this;
  }

   /**
   * Get availability
   * @return availability
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public String getAvailability() {
    return availability;
  }


  public void setAvailability(String availability) {
    this.availability = availability;
  }


  public CheckApp checks(List<Long> checks) {
    
    this.checks = checks;
    return this;
  }

  public CheckApp addChecksItem(Long checksItem) {
    if (this.checks == null) {
      this.checks = new ArrayList<Long>();
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

  public List<Long> getChecks() {
    return checks;
  }


  public void setChecks(List<Long> checks) {
    this.checks = checks;
  }


  public CheckApp checkGroups(List<Long> checkGroups) {
    
    this.checkGroups = checkGroups;
    return this;
  }

  public CheckApp addCheckGroupsItem(Long checkGroupsItem) {
    if (this.checkGroups == null) {
      this.checkGroups = new ArrayList<Long>();
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

  public List<Long> getCheckGroups() {
    return checkGroups;
  }


  public void setCheckGroups(List<Long> checkGroups) {
    this.checkGroups = checkGroups;
  }


  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (o == null || getClass() != o.getClass()) {
      return false;
    }
    CheckApp checkApp = (CheckApp) o;
    return Objects.equals(this.id, checkApp.id) &&
        Objects.equals(this.account, checkApp.account) &&
        Objects.equals(this.active, checkApp.active) &&
        Objects.equals(this.app, checkApp.app) &&
        Objects.equals(this.name, checkApp.name) &&
        Objects.equals(this.description, checkApp.description) &&
        Objects.equals(this.inUptime, checkApp.inUptime) &&
        Objects.equals(this.lastCheckStatus, checkApp.lastCheckStatus) &&
        Objects.equals(this.lastWarningDate, checkApp.lastWarningDate) &&
        Objects.equals(this.lastErrorDate, checkApp.lastErrorDate) &&
        Objects.equals(this.lastSuccessDate, checkApp.lastSuccessDate) &&
        Objects.equals(this.lastRunDate, checkApp.lastRunDate) &&
        Objects.equals(this.lastError, checkApp.lastError) &&
        Objects.equals(this.lastTimer, checkApp.lastTimer) &&
        Objects.equals(this.health, checkApp.health) &&
        Objects.equals(this.history, checkApp.history) &&
        Objects.equals(this.severity, checkApp.severity) &&
        Objects.equals(this.createIncident, checkApp.createIncident) &&
        Objects.equals(this.muted, checkApp.muted) &&
        Objects.equals(this.createdBy, checkApp.createdBy) &&
        Objects.equals(this.dateCreated, checkApp.dateCreated) &&
        Objects.equals(this.lastUpdated, checkApp.lastUpdated) &&
        Objects.equals(this.availability, checkApp.availability) &&
        Objects.equals(this.checks, checkApp.checks) &&
        Objects.equals(this.checkGroups, checkApp.checkGroups);
  }

  @Override
  public int hashCode() {
    return Objects.hash(id, account, active, app, name, description, inUptime, lastCheckStatus, lastWarningDate, lastErrorDate, lastSuccessDate, lastRunDate, lastError, lastTimer, health, history, severity, createIncident, muted, createdBy, dateCreated, lastUpdated, availability, checks, checkGroups);
  }


  @Override
  public String toString() {
    StringBuilder sb = new StringBuilder();
    sb.append("class CheckApp {\n");
    sb.append("    id: ").append(toIndentedString(id)).append("\n");
    sb.append("    account: ").append(toIndentedString(account)).append("\n");
    sb.append("    active: ").append(toIndentedString(active)).append("\n");
    sb.append("    app: ").append(toIndentedString(app)).append("\n");
    sb.append("    name: ").append(toIndentedString(name)).append("\n");
    sb.append("    description: ").append(toIndentedString(description)).append("\n");
    sb.append("    inUptime: ").append(toIndentedString(inUptime)).append("\n");
    sb.append("    lastCheckStatus: ").append(toIndentedString(lastCheckStatus)).append("\n");
    sb.append("    lastWarningDate: ").append(toIndentedString(lastWarningDate)).append("\n");
    sb.append("    lastErrorDate: ").append(toIndentedString(lastErrorDate)).append("\n");
    sb.append("    lastSuccessDate: ").append(toIndentedString(lastSuccessDate)).append("\n");
    sb.append("    lastRunDate: ").append(toIndentedString(lastRunDate)).append("\n");
    sb.append("    lastError: ").append(toIndentedString(lastError)).append("\n");
    sb.append("    lastTimer: ").append(toIndentedString(lastTimer)).append("\n");
    sb.append("    health: ").append(toIndentedString(health)).append("\n");
    sb.append("    history: ").append(toIndentedString(history)).append("\n");
    sb.append("    severity: ").append(toIndentedString(severity)).append("\n");
    sb.append("    createIncident: ").append(toIndentedString(createIncident)).append("\n");
    sb.append("    muted: ").append(toIndentedString(muted)).append("\n");
    sb.append("    createdBy: ").append(toIndentedString(createdBy)).append("\n");
    sb.append("    dateCreated: ").append(toIndentedString(dateCreated)).append("\n");
    sb.append("    lastUpdated: ").append(toIndentedString(lastUpdated)).append("\n");
    sb.append("    availability: ").append(toIndentedString(availability)).append("\n");
    sb.append("    checks: ").append(toIndentedString(checks)).append("\n");
    sb.append("    checkGroups: ").append(toIndentedString(checkGroups)).append("\n");
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

