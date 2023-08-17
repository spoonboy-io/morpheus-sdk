/*
 * Morpheus API
 * Morpheus is a powerful cloud management tool that provides provisioning, monitoring, logging, backups, and application deployment strategies.  This document describes the Morpheus API protocol and the available endpoints. Sections are organized in the same manner as they appear in the Morpheus UI.
 *
 * The version of the OpenAPI document: 6.1.1
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
import java.io.IOException;
import java.math.BigDecimal;
import java.time.OffsetDateTime;
import org.openapitools.client.model.CheckCheckType;
import org.openapitools.client.model.CheckConfig;
import org.openapitools.client.model.CheckContainer;
import org.openapitools.client.model.CheckCreatedBy;
import org.openapitools.client.model.UpdateBlueprintPermissionsRequestResourcePermissionSitesInner;
import org.openapitools.jackson.nullable.JsonNullable;

import com.google.gson.Gson;
import com.google.gson.GsonBuilder;
import com.google.gson.JsonArray;
import com.google.gson.JsonDeserializationContext;
import com.google.gson.JsonDeserializer;
import com.google.gson.JsonElement;
import com.google.gson.JsonObject;
import com.google.gson.JsonParseException;
import com.google.gson.TypeAdapterFactory;
import com.google.gson.reflect.TypeToken;
import com.google.gson.TypeAdapter;
import com.google.gson.stream.JsonReader;
import com.google.gson.stream.JsonWriter;
import java.io.IOException;

import java.lang.reflect.Type;
import java.util.HashMap;
import java.util.HashSet;
import java.util.List;
import java.util.Map;
import java.util.Map.Entry;
import java.util.Set;

import org.openapitools.client.JSON;

/**
 * Check
 */
@javax.annotation.Generated(value = "org.openapitools.codegen.languages.JavaClientCodegen", date = "2023-08-17T13:37:08.123540Z[Etc/UTC]")
public class Check {
  public static final String SERIALIZED_NAME_ID = "id";
  @SerializedName(SERIALIZED_NAME_ID)
  private Long id;

  public static final String SERIALIZED_NAME_ACCOUNT = "account";
  @SerializedName(SERIALIZED_NAME_ACCOUNT)
  private UpdateBlueprintPermissionsRequestResourcePermissionSitesInner account;

  public static final String SERIALIZED_NAME_ACTIVE = "active";
  @SerializedName(SERIALIZED_NAME_ACTIVE)
  private Boolean active;

  public static final String SERIALIZED_NAME_API_KEY = "apiKey";
  @SerializedName(SERIALIZED_NAME_API_KEY)
  private String apiKey;

  public static final String SERIALIZED_NAME_AVAILABILITY = "availability";
  @SerializedName(SERIALIZED_NAME_AVAILABILITY)
  private BigDecimal availability;

  public static final String SERIALIZED_NAME_CHECK_AGENT = "checkAgent";
  @SerializedName(SERIALIZED_NAME_CHECK_AGENT)
  private String checkAgent;

  public static final String SERIALIZED_NAME_CHECK_INTERVAL = "checkInterval";
  @SerializedName(SERIALIZED_NAME_CHECK_INTERVAL)
  private Long checkInterval;

  public static final String SERIALIZED_NAME_CHECK_SPEC = "checkSpec";
  @SerializedName(SERIALIZED_NAME_CHECK_SPEC)
  private String checkSpec;

  public static final String SERIALIZED_NAME_CHECK_TYPE = "checkType";
  @SerializedName(SERIALIZED_NAME_CHECK_TYPE)
  private CheckCheckType checkType;

  public static final String SERIALIZED_NAME_CONFIG = "config";
  @SerializedName(SERIALIZED_NAME_CONFIG)
  private CheckConfig config;

  public static final String SERIALIZED_NAME_CONTAINER = "container";
  @SerializedName(SERIALIZED_NAME_CONTAINER)
  private CheckContainer container;

  public static final String SERIALIZED_NAME_CREATE_INCIDENT = "createIncident";
  @SerializedName(SERIALIZED_NAME_CREATE_INCIDENT)
  private Boolean createIncident;

  public static final String SERIALIZED_NAME_MUTED = "muted";
  @SerializedName(SERIALIZED_NAME_MUTED)
  private Boolean muted;

  public static final String SERIALIZED_NAME_CREATED_BY = "createdBy";
  @SerializedName(SERIALIZED_NAME_CREATED_BY)
  private CheckCreatedBy createdBy;

  public static final String SERIALIZED_NAME_DATE_CREATED = "dateCreated";
  @SerializedName(SERIALIZED_NAME_DATE_CREATED)
  private OffsetDateTime dateCreated;

  public static final String SERIALIZED_NAME_DESCRIPTION = "description";
  @SerializedName(SERIALIZED_NAME_DESCRIPTION)
  private String description;

  public static final String SERIALIZED_NAME_END_DATE = "endDate";
  @SerializedName(SERIALIZED_NAME_END_DATE)
  private OffsetDateTime endDate;

  public static final String SERIALIZED_NAME_HEALTH = "health";
  @SerializedName(SERIALIZED_NAME_HEALTH)
  private Long health;

  public static final String SERIALIZED_NAME_IN_UPTIME = "inUptime";
  @SerializedName(SERIALIZED_NAME_IN_UPTIME)
  private Boolean inUptime;

  public static final String SERIALIZED_NAME_LAST_BOX_STATS = "lastBoxStats";
  @SerializedName(SERIALIZED_NAME_LAST_BOX_STATS)
  private String lastBoxStats;

  public static final String SERIALIZED_NAME_LAST_CHECK_STATUS = "lastCheckStatus";
  @SerializedName(SERIALIZED_NAME_LAST_CHECK_STATUS)
  private String lastCheckStatus;

  public static final String SERIALIZED_NAME_LAST_ERROR = "lastError";
  @SerializedName(SERIALIZED_NAME_LAST_ERROR)
  private String lastError;

  public static final String SERIALIZED_NAME_LAST_ERROR_DATE = "lastErrorDate";
  @SerializedName(SERIALIZED_NAME_LAST_ERROR_DATE)
  private OffsetDateTime lastErrorDate;

  public static final String SERIALIZED_NAME_LAST_MESSAGE = "lastMessage";
  @SerializedName(SERIALIZED_NAME_LAST_MESSAGE)
  private String lastMessage;

  public static final String SERIALIZED_NAME_LAST_METRIC = "lastMetric";
  @SerializedName(SERIALIZED_NAME_LAST_METRIC)
  private String lastMetric;

  public static final String SERIALIZED_NAME_LAST_RUN_DATE = "lastRunDate";
  @SerializedName(SERIALIZED_NAME_LAST_RUN_DATE)
  private OffsetDateTime lastRunDate;

  public static final String SERIALIZED_NAME_LAST_STATS = "lastStats";
  @SerializedName(SERIALIZED_NAME_LAST_STATS)
  private String lastStats;

  public static final String SERIALIZED_NAME_LAST_SUCCESS_DATE = "lastSuccessDate";
  @SerializedName(SERIALIZED_NAME_LAST_SUCCESS_DATE)
  private OffsetDateTime lastSuccessDate;

  public static final String SERIALIZED_NAME_LAST_TIMER = "lastTimer";
  @SerializedName(SERIALIZED_NAME_LAST_TIMER)
  private Long lastTimer;

  public static final String SERIALIZED_NAME_LAST_UPDATED = "lastUpdated";
  @SerializedName(SERIALIZED_NAME_LAST_UPDATED)
  private OffsetDateTime lastUpdated;

  public static final String SERIALIZED_NAME_LAST_WARNING_DATE = "lastWarningDate";
  @SerializedName(SERIALIZED_NAME_LAST_WARNING_DATE)
  private OffsetDateTime lastWarningDate;

  public static final String SERIALIZED_NAME_NAME = "name";
  @SerializedName(SERIALIZED_NAME_NAME)
  private String name;

  public static final String SERIALIZED_NAME_NEXT_RUN_DATE = "nextRunDate";
  @SerializedName(SERIALIZED_NAME_NEXT_RUN_DATE)
  private OffsetDateTime nextRunDate;

  public static final String SERIALIZED_NAME_OUTAGE_TIME = "outageTime";
  @SerializedName(SERIALIZED_NAME_OUTAGE_TIME)
  private Long outageTime;

  public static final String SERIALIZED_NAME_SEVERITY = "severity";
  @SerializedName(SERIALIZED_NAME_SEVERITY)
  private String severity;

  public static final String SERIALIZED_NAME_START_DATE = "startDate";
  @SerializedName(SERIALIZED_NAME_START_DATE)
  private OffsetDateTime startDate;

  public static final String SERIALIZED_NAME_DELETED = "deleted";
  @SerializedName(SERIALIZED_NAME_DELETED)
  private Boolean deleted;

  public Check() {
  }

  public Check id(Long id) {
    
    this.id = id;
    return this;
  }

   /**
   * Get id
   * @return id
  **/
  @javax.annotation.Nullable
  public Long getId() {
    return id;
  }


  public void setId(Long id) {
    this.id = id;
  }


  public Check account(UpdateBlueprintPermissionsRequestResourcePermissionSitesInner account) {
    
    this.account = account;
    return this;
  }

   /**
   * Get account
   * @return account
  **/
  @javax.annotation.Nullable
  public UpdateBlueprintPermissionsRequestResourcePermissionSitesInner getAccount() {
    return account;
  }


  public void setAccount(UpdateBlueprintPermissionsRequestResourcePermissionSitesInner account) {
    this.account = account;
  }


  public Check active(Boolean active) {
    
    this.active = active;
    return this;
  }

   /**
   * Get active
   * @return active
  **/
  @javax.annotation.Nullable
  public Boolean getActive() {
    return active;
  }


  public void setActive(Boolean active) {
    this.active = active;
  }


  public Check apiKey(String apiKey) {
    
    this.apiKey = apiKey;
    return this;
  }

   /**
   * Get apiKey
   * @return apiKey
  **/
  @javax.annotation.Nullable
  public String getApiKey() {
    return apiKey;
  }


  public void setApiKey(String apiKey) {
    this.apiKey = apiKey;
  }


  public Check availability(BigDecimal availability) {
    
    this.availability = availability;
    return this;
  }

   /**
   * Get availability
   * @return availability
  **/
  @javax.annotation.Nullable
  public BigDecimal getAvailability() {
    return availability;
  }


  public void setAvailability(BigDecimal availability) {
    this.availability = availability;
  }


  public Check checkAgent(String checkAgent) {
    
    this.checkAgent = checkAgent;
    return this;
  }

   /**
   * Get checkAgent
   * @return checkAgent
  **/
  @javax.annotation.Nullable
  public String getCheckAgent() {
    return checkAgent;
  }


  public void setCheckAgent(String checkAgent) {
    this.checkAgent = checkAgent;
  }


  public Check checkInterval(Long checkInterval) {
    
    this.checkInterval = checkInterval;
    return this;
  }

   /**
   * Get checkInterval
   * @return checkInterval
  **/
  @javax.annotation.Nullable
  public Long getCheckInterval() {
    return checkInterval;
  }


  public void setCheckInterval(Long checkInterval) {
    this.checkInterval = checkInterval;
  }


  public Check checkSpec(String checkSpec) {
    
    this.checkSpec = checkSpec;
    return this;
  }

   /**
   * Get checkSpec
   * @return checkSpec
  **/
  @javax.annotation.Nullable
  public String getCheckSpec() {
    return checkSpec;
  }


  public void setCheckSpec(String checkSpec) {
    this.checkSpec = checkSpec;
  }


  public Check checkType(CheckCheckType checkType) {
    
    this.checkType = checkType;
    return this;
  }

   /**
   * Get checkType
   * @return checkType
  **/
  @javax.annotation.Nullable
  public CheckCheckType getCheckType() {
    return checkType;
  }


  public void setCheckType(CheckCheckType checkType) {
    this.checkType = checkType;
  }


  public Check config(CheckConfig config) {
    
    this.config = config;
    return this;
  }

   /**
   * Get config
   * @return config
  **/
  @javax.annotation.Nullable
  public CheckConfig getConfig() {
    return config;
  }


  public void setConfig(CheckConfig config) {
    this.config = config;
  }


  public Check container(CheckContainer container) {
    
    this.container = container;
    return this;
  }

   /**
   * Get container
   * @return container
  **/
  @javax.annotation.Nullable
  public CheckContainer getContainer() {
    return container;
  }


  public void setContainer(CheckContainer container) {
    this.container = container;
  }


  public Check createIncident(Boolean createIncident) {
    
    this.createIncident = createIncident;
    return this;
  }

   /**
   * Get createIncident
   * @return createIncident
  **/
  @javax.annotation.Nullable
  public Boolean getCreateIncident() {
    return createIncident;
  }


  public void setCreateIncident(Boolean createIncident) {
    this.createIncident = createIncident;
  }


  public Check muted(Boolean muted) {
    
    this.muted = muted;
    return this;
  }

   /**
   * Get muted
   * @return muted
  **/
  @javax.annotation.Nullable
  public Boolean getMuted() {
    return muted;
  }


  public void setMuted(Boolean muted) {
    this.muted = muted;
  }


  public Check createdBy(CheckCreatedBy createdBy) {
    
    this.createdBy = createdBy;
    return this;
  }

   /**
   * Get createdBy
   * @return createdBy
  **/
  @javax.annotation.Nullable
  public CheckCreatedBy getCreatedBy() {
    return createdBy;
  }


  public void setCreatedBy(CheckCreatedBy createdBy) {
    this.createdBy = createdBy;
  }


  public Check dateCreated(OffsetDateTime dateCreated) {
    
    this.dateCreated = dateCreated;
    return this;
  }

   /**
   * Get dateCreated
   * @return dateCreated
  **/
  @javax.annotation.Nullable
  public OffsetDateTime getDateCreated() {
    return dateCreated;
  }


  public void setDateCreated(OffsetDateTime dateCreated) {
    this.dateCreated = dateCreated;
  }


  public Check description(String description) {
    
    this.description = description;
    return this;
  }

   /**
   * Get description
   * @return description
  **/
  @javax.annotation.Nullable
  public String getDescription() {
    return description;
  }


  public void setDescription(String description) {
    this.description = description;
  }


  public Check endDate(OffsetDateTime endDate) {
    
    this.endDate = endDate;
    return this;
  }

   /**
   * Get endDate
   * @return endDate
  **/
  @javax.annotation.Nullable
  public OffsetDateTime getEndDate() {
    return endDate;
  }


  public void setEndDate(OffsetDateTime endDate) {
    this.endDate = endDate;
  }


  public Check health(Long health) {
    
    this.health = health;
    return this;
  }

   /**
   * Get health
   * @return health
  **/
  @javax.annotation.Nullable
  public Long getHealth() {
    return health;
  }


  public void setHealth(Long health) {
    this.health = health;
  }


  public Check inUptime(Boolean inUptime) {
    
    this.inUptime = inUptime;
    return this;
  }

   /**
   * Get inUptime
   * @return inUptime
  **/
  @javax.annotation.Nullable
  public Boolean getInUptime() {
    return inUptime;
  }


  public void setInUptime(Boolean inUptime) {
    this.inUptime = inUptime;
  }


  public Check lastBoxStats(String lastBoxStats) {
    
    this.lastBoxStats = lastBoxStats;
    return this;
  }

   /**
   * Get lastBoxStats
   * @return lastBoxStats
  **/
  @javax.annotation.Nullable
  public String getLastBoxStats() {
    return lastBoxStats;
  }


  public void setLastBoxStats(String lastBoxStats) {
    this.lastBoxStats = lastBoxStats;
  }


  public Check lastCheckStatus(String lastCheckStatus) {
    
    this.lastCheckStatus = lastCheckStatus;
    return this;
  }

   /**
   * Get lastCheckStatus
   * @return lastCheckStatus
  **/
  @javax.annotation.Nullable
  public String getLastCheckStatus() {
    return lastCheckStatus;
  }


  public void setLastCheckStatus(String lastCheckStatus) {
    this.lastCheckStatus = lastCheckStatus;
  }


  public Check lastError(String lastError) {
    
    this.lastError = lastError;
    return this;
  }

   /**
   * Get lastError
   * @return lastError
  **/
  @javax.annotation.Nullable
  public String getLastError() {
    return lastError;
  }


  public void setLastError(String lastError) {
    this.lastError = lastError;
  }


  public Check lastErrorDate(OffsetDateTime lastErrorDate) {
    
    this.lastErrorDate = lastErrorDate;
    return this;
  }

   /**
   * Get lastErrorDate
   * @return lastErrorDate
  **/
  @javax.annotation.Nullable
  public OffsetDateTime getLastErrorDate() {
    return lastErrorDate;
  }


  public void setLastErrorDate(OffsetDateTime lastErrorDate) {
    this.lastErrorDate = lastErrorDate;
  }


  public Check lastMessage(String lastMessage) {
    
    this.lastMessage = lastMessage;
    return this;
  }

   /**
   * Get lastMessage
   * @return lastMessage
  **/
  @javax.annotation.Nullable
  public String getLastMessage() {
    return lastMessage;
  }


  public void setLastMessage(String lastMessage) {
    this.lastMessage = lastMessage;
  }


  public Check lastMetric(String lastMetric) {
    
    this.lastMetric = lastMetric;
    return this;
  }

   /**
   * Get lastMetric
   * @return lastMetric
  **/
  @javax.annotation.Nullable
  public String getLastMetric() {
    return lastMetric;
  }


  public void setLastMetric(String lastMetric) {
    this.lastMetric = lastMetric;
  }


  public Check lastRunDate(OffsetDateTime lastRunDate) {
    
    this.lastRunDate = lastRunDate;
    return this;
  }

   /**
   * Get lastRunDate
   * @return lastRunDate
  **/
  @javax.annotation.Nullable
  public OffsetDateTime getLastRunDate() {
    return lastRunDate;
  }


  public void setLastRunDate(OffsetDateTime lastRunDate) {
    this.lastRunDate = lastRunDate;
  }


  public Check lastStats(String lastStats) {
    
    this.lastStats = lastStats;
    return this;
  }

   /**
   * Get lastStats
   * @return lastStats
  **/
  @javax.annotation.Nullable
  public String getLastStats() {
    return lastStats;
  }


  public void setLastStats(String lastStats) {
    this.lastStats = lastStats;
  }


  public Check lastSuccessDate(OffsetDateTime lastSuccessDate) {
    
    this.lastSuccessDate = lastSuccessDate;
    return this;
  }

   /**
   * Get lastSuccessDate
   * @return lastSuccessDate
  **/
  @javax.annotation.Nullable
  public OffsetDateTime getLastSuccessDate() {
    return lastSuccessDate;
  }


  public void setLastSuccessDate(OffsetDateTime lastSuccessDate) {
    this.lastSuccessDate = lastSuccessDate;
  }


  public Check lastTimer(Long lastTimer) {
    
    this.lastTimer = lastTimer;
    return this;
  }

   /**
   * Get lastTimer
   * @return lastTimer
  **/
  @javax.annotation.Nullable
  public Long getLastTimer() {
    return lastTimer;
  }


  public void setLastTimer(Long lastTimer) {
    this.lastTimer = lastTimer;
  }


  public Check lastUpdated(OffsetDateTime lastUpdated) {
    
    this.lastUpdated = lastUpdated;
    return this;
  }

   /**
   * Get lastUpdated
   * @return lastUpdated
  **/
  @javax.annotation.Nullable
  public OffsetDateTime getLastUpdated() {
    return lastUpdated;
  }


  public void setLastUpdated(OffsetDateTime lastUpdated) {
    this.lastUpdated = lastUpdated;
  }


  public Check lastWarningDate(OffsetDateTime lastWarningDate) {
    
    this.lastWarningDate = lastWarningDate;
    return this;
  }

   /**
   * Get lastWarningDate
   * @return lastWarningDate
  **/
  @javax.annotation.Nullable
  public OffsetDateTime getLastWarningDate() {
    return lastWarningDate;
  }


  public void setLastWarningDate(OffsetDateTime lastWarningDate) {
    this.lastWarningDate = lastWarningDate;
  }


  public Check name(String name) {
    
    this.name = name;
    return this;
  }

   /**
   * Get name
   * @return name
  **/
  @javax.annotation.Nullable
  public String getName() {
    return name;
  }


  public void setName(String name) {
    this.name = name;
  }


  public Check nextRunDate(OffsetDateTime nextRunDate) {
    
    this.nextRunDate = nextRunDate;
    return this;
  }

   /**
   * Get nextRunDate
   * @return nextRunDate
  **/
  @javax.annotation.Nullable
  public OffsetDateTime getNextRunDate() {
    return nextRunDate;
  }


  public void setNextRunDate(OffsetDateTime nextRunDate) {
    this.nextRunDate = nextRunDate;
  }


  public Check outageTime(Long outageTime) {
    
    this.outageTime = outageTime;
    return this;
  }

   /**
   * Get outageTime
   * @return outageTime
  **/
  @javax.annotation.Nullable
  public Long getOutageTime() {
    return outageTime;
  }


  public void setOutageTime(Long outageTime) {
    this.outageTime = outageTime;
  }


  public Check severity(String severity) {
    
    this.severity = severity;
    return this;
  }

   /**
   * Get severity
   * @return severity
  **/
  @javax.annotation.Nullable
  public String getSeverity() {
    return severity;
  }


  public void setSeverity(String severity) {
    this.severity = severity;
  }


  public Check startDate(OffsetDateTime startDate) {
    
    this.startDate = startDate;
    return this;
  }

   /**
   * Get startDate
   * @return startDate
  **/
  @javax.annotation.Nullable
  public OffsetDateTime getStartDate() {
    return startDate;
  }


  public void setStartDate(OffsetDateTime startDate) {
    this.startDate = startDate;
  }


  public Check deleted(Boolean deleted) {
    
    this.deleted = deleted;
    return this;
  }

   /**
   * Get deleted
   * @return deleted
  **/
  @javax.annotation.Nullable
  public Boolean getDeleted() {
    return deleted;
  }


  public void setDeleted(Boolean deleted) {
    this.deleted = deleted;
  }



  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (o == null || getClass() != o.getClass()) {
      return false;
    }
    Check check = (Check) o;
    return Objects.equals(this.id, check.id) &&
        Objects.equals(this.account, check.account) &&
        Objects.equals(this.active, check.active) &&
        Objects.equals(this.apiKey, check.apiKey) &&
        Objects.equals(this.availability, check.availability) &&
        Objects.equals(this.checkAgent, check.checkAgent) &&
        Objects.equals(this.checkInterval, check.checkInterval) &&
        Objects.equals(this.checkSpec, check.checkSpec) &&
        Objects.equals(this.checkType, check.checkType) &&
        Objects.equals(this.config, check.config) &&
        Objects.equals(this.container, check.container) &&
        Objects.equals(this.createIncident, check.createIncident) &&
        Objects.equals(this.muted, check.muted) &&
        Objects.equals(this.createdBy, check.createdBy) &&
        Objects.equals(this.dateCreated, check.dateCreated) &&
        Objects.equals(this.description, check.description) &&
        Objects.equals(this.endDate, check.endDate) &&
        Objects.equals(this.health, check.health) &&
        Objects.equals(this.inUptime, check.inUptime) &&
        Objects.equals(this.lastBoxStats, check.lastBoxStats) &&
        Objects.equals(this.lastCheckStatus, check.lastCheckStatus) &&
        Objects.equals(this.lastError, check.lastError) &&
        Objects.equals(this.lastErrorDate, check.lastErrorDate) &&
        Objects.equals(this.lastMessage, check.lastMessage) &&
        Objects.equals(this.lastMetric, check.lastMetric) &&
        Objects.equals(this.lastRunDate, check.lastRunDate) &&
        Objects.equals(this.lastStats, check.lastStats) &&
        Objects.equals(this.lastSuccessDate, check.lastSuccessDate) &&
        Objects.equals(this.lastTimer, check.lastTimer) &&
        Objects.equals(this.lastUpdated, check.lastUpdated) &&
        Objects.equals(this.lastWarningDate, check.lastWarningDate) &&
        Objects.equals(this.name, check.name) &&
        Objects.equals(this.nextRunDate, check.nextRunDate) &&
        Objects.equals(this.outageTime, check.outageTime) &&
        Objects.equals(this.severity, check.severity) &&
        Objects.equals(this.startDate, check.startDate) &&
        Objects.equals(this.deleted, check.deleted);
  }

  private static <T> boolean equalsNullable(JsonNullable<T> a, JsonNullable<T> b) {
    return a == b || (a != null && b != null && a.isPresent() && b.isPresent() && Objects.deepEquals(a.get(), b.get()));
  }

  @Override
  public int hashCode() {
    return Objects.hash(id, account, active, apiKey, availability, checkAgent, checkInterval, checkSpec, checkType, config, container, createIncident, muted, createdBy, dateCreated, description, endDate, health, inUptime, lastBoxStats, lastCheckStatus, lastError, lastErrorDate, lastMessage, lastMetric, lastRunDate, lastStats, lastSuccessDate, lastTimer, lastUpdated, lastWarningDate, name, nextRunDate, outageTime, severity, startDate, deleted);
  }

  private static <T> int hashCodeNullable(JsonNullable<T> a) {
    if (a == null) {
      return 1;
    }
    return a.isPresent() ? Arrays.deepHashCode(new Object[]{a.get()}) : 31;
  }

  @Override
  public String toString() {
    StringBuilder sb = new StringBuilder();
    sb.append("class Check {\n");
    sb.append("    id: ").append(toIndentedString(id)).append("\n");
    sb.append("    account: ").append(toIndentedString(account)).append("\n");
    sb.append("    active: ").append(toIndentedString(active)).append("\n");
    sb.append("    apiKey: ").append(toIndentedString(apiKey)).append("\n");
    sb.append("    availability: ").append(toIndentedString(availability)).append("\n");
    sb.append("    checkAgent: ").append(toIndentedString(checkAgent)).append("\n");
    sb.append("    checkInterval: ").append(toIndentedString(checkInterval)).append("\n");
    sb.append("    checkSpec: ").append(toIndentedString(checkSpec)).append("\n");
    sb.append("    checkType: ").append(toIndentedString(checkType)).append("\n");
    sb.append("    config: ").append(toIndentedString(config)).append("\n");
    sb.append("    container: ").append(toIndentedString(container)).append("\n");
    sb.append("    createIncident: ").append(toIndentedString(createIncident)).append("\n");
    sb.append("    muted: ").append(toIndentedString(muted)).append("\n");
    sb.append("    createdBy: ").append(toIndentedString(createdBy)).append("\n");
    sb.append("    dateCreated: ").append(toIndentedString(dateCreated)).append("\n");
    sb.append("    description: ").append(toIndentedString(description)).append("\n");
    sb.append("    endDate: ").append(toIndentedString(endDate)).append("\n");
    sb.append("    health: ").append(toIndentedString(health)).append("\n");
    sb.append("    inUptime: ").append(toIndentedString(inUptime)).append("\n");
    sb.append("    lastBoxStats: ").append(toIndentedString(lastBoxStats)).append("\n");
    sb.append("    lastCheckStatus: ").append(toIndentedString(lastCheckStatus)).append("\n");
    sb.append("    lastError: ").append(toIndentedString(lastError)).append("\n");
    sb.append("    lastErrorDate: ").append(toIndentedString(lastErrorDate)).append("\n");
    sb.append("    lastMessage: ").append(toIndentedString(lastMessage)).append("\n");
    sb.append("    lastMetric: ").append(toIndentedString(lastMetric)).append("\n");
    sb.append("    lastRunDate: ").append(toIndentedString(lastRunDate)).append("\n");
    sb.append("    lastStats: ").append(toIndentedString(lastStats)).append("\n");
    sb.append("    lastSuccessDate: ").append(toIndentedString(lastSuccessDate)).append("\n");
    sb.append("    lastTimer: ").append(toIndentedString(lastTimer)).append("\n");
    sb.append("    lastUpdated: ").append(toIndentedString(lastUpdated)).append("\n");
    sb.append("    lastWarningDate: ").append(toIndentedString(lastWarningDate)).append("\n");
    sb.append("    name: ").append(toIndentedString(name)).append("\n");
    sb.append("    nextRunDate: ").append(toIndentedString(nextRunDate)).append("\n");
    sb.append("    outageTime: ").append(toIndentedString(outageTime)).append("\n");
    sb.append("    severity: ").append(toIndentedString(severity)).append("\n");
    sb.append("    startDate: ").append(toIndentedString(startDate)).append("\n");
    sb.append("    deleted: ").append(toIndentedString(deleted)).append("\n");
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


  public static HashSet<String> openapiFields;
  public static HashSet<String> openapiRequiredFields;

  static {
    // a set of all properties/fields (JSON key names)
    openapiFields = new HashSet<String>();
    openapiFields.add("id");
    openapiFields.add("account");
    openapiFields.add("active");
    openapiFields.add("apiKey");
    openapiFields.add("availability");
    openapiFields.add("checkAgent");
    openapiFields.add("checkInterval");
    openapiFields.add("checkSpec");
    openapiFields.add("checkType");
    openapiFields.add("config");
    openapiFields.add("container");
    openapiFields.add("createIncident");
    openapiFields.add("muted");
    openapiFields.add("createdBy");
    openapiFields.add("dateCreated");
    openapiFields.add("description");
    openapiFields.add("endDate");
    openapiFields.add("health");
    openapiFields.add("inUptime");
    openapiFields.add("lastBoxStats");
    openapiFields.add("lastCheckStatus");
    openapiFields.add("lastError");
    openapiFields.add("lastErrorDate");
    openapiFields.add("lastMessage");
    openapiFields.add("lastMetric");
    openapiFields.add("lastRunDate");
    openapiFields.add("lastStats");
    openapiFields.add("lastSuccessDate");
    openapiFields.add("lastTimer");
    openapiFields.add("lastUpdated");
    openapiFields.add("lastWarningDate");
    openapiFields.add("name");
    openapiFields.add("nextRunDate");
    openapiFields.add("outageTime");
    openapiFields.add("severity");
    openapiFields.add("startDate");
    openapiFields.add("deleted");

    // a set of required properties/fields (JSON key names)
    openapiRequiredFields = new HashSet<String>();
  }

 /**
  * Validates the JSON Element and throws an exception if issues found
  *
  * @param jsonElement JSON Element
  * @throws IOException if the JSON Element is invalid with respect to Check
  */
  public static void validateJsonElement(JsonElement jsonElement) throws IOException {
      if (jsonElement == null) {
        if (!Check.openapiRequiredFields.isEmpty()) { // has required fields but JSON element is null
          throw new IllegalArgumentException(String.format("The required field(s) %s in Check is not found in the empty JSON string", Check.openapiRequiredFields.toString()));
        }
      }

      Set<Entry<String, JsonElement>> entries = jsonElement.getAsJsonObject().entrySet();
      // check to see if the JSON string contains additional fields
      for (Entry<String, JsonElement> entry : entries) {
        if (!Check.openapiFields.contains(entry.getKey())) {
          throw new IllegalArgumentException(String.format("The field `%s` in the JSON string is not defined in the `Check` properties. JSON: %s", entry.getKey(), jsonElement.toString()));
        }
      }
        JsonObject jsonObj = jsonElement.getAsJsonObject();
      // validate the optional field `account`
      if (jsonObj.get("account") != null && !jsonObj.get("account").isJsonNull()) {
        UpdateBlueprintPermissionsRequestResourcePermissionSitesInner.validateJsonElement(jsonObj.get("account"));
      }
      if ((jsonObj.get("apiKey") != null && !jsonObj.get("apiKey").isJsonNull()) && !jsonObj.get("apiKey").isJsonPrimitive()) {
        throw new IllegalArgumentException(String.format("Expected the field `apiKey` to be a primitive type in the JSON string but got `%s`", jsonObj.get("apiKey").toString()));
      }
      if ((jsonObj.get("checkAgent") != null && !jsonObj.get("checkAgent").isJsonNull()) && !jsonObj.get("checkAgent").isJsonPrimitive()) {
        throw new IllegalArgumentException(String.format("Expected the field `checkAgent` to be a primitive type in the JSON string but got `%s`", jsonObj.get("checkAgent").toString()));
      }
      if ((jsonObj.get("checkSpec") != null && !jsonObj.get("checkSpec").isJsonNull()) && !jsonObj.get("checkSpec").isJsonPrimitive()) {
        throw new IllegalArgumentException(String.format("Expected the field `checkSpec` to be a primitive type in the JSON string but got `%s`", jsonObj.get("checkSpec").toString()));
      }
      // validate the optional field `checkType`
      if (jsonObj.get("checkType") != null && !jsonObj.get("checkType").isJsonNull()) {
        CheckCheckType.validateJsonElement(jsonObj.get("checkType"));
      }
      // validate the optional field `config`
      if (jsonObj.get("config") != null && !jsonObj.get("config").isJsonNull()) {
        CheckConfig.validateJsonElement(jsonObj.get("config"));
      }
      // validate the optional field `container`
      if (jsonObj.get("container") != null && !jsonObj.get("container").isJsonNull()) {
        CheckContainer.validateJsonElement(jsonObj.get("container"));
      }
      // validate the optional field `createdBy`
      if (jsonObj.get("createdBy") != null && !jsonObj.get("createdBy").isJsonNull()) {
        CheckCreatedBy.validateJsonElement(jsonObj.get("createdBy"));
      }
      if ((jsonObj.get("description") != null && !jsonObj.get("description").isJsonNull()) && !jsonObj.get("description").isJsonPrimitive()) {
        throw new IllegalArgumentException(String.format("Expected the field `description` to be a primitive type in the JSON string but got `%s`", jsonObj.get("description").toString()));
      }
      if ((jsonObj.get("lastBoxStats") != null && !jsonObj.get("lastBoxStats").isJsonNull()) && !jsonObj.get("lastBoxStats").isJsonPrimitive()) {
        throw new IllegalArgumentException(String.format("Expected the field `lastBoxStats` to be a primitive type in the JSON string but got `%s`", jsonObj.get("lastBoxStats").toString()));
      }
      if ((jsonObj.get("lastCheckStatus") != null && !jsonObj.get("lastCheckStatus").isJsonNull()) && !jsonObj.get("lastCheckStatus").isJsonPrimitive()) {
        throw new IllegalArgumentException(String.format("Expected the field `lastCheckStatus` to be a primitive type in the JSON string but got `%s`", jsonObj.get("lastCheckStatus").toString()));
      }
      if ((jsonObj.get("lastError") != null && !jsonObj.get("lastError").isJsonNull()) && !jsonObj.get("lastError").isJsonPrimitive()) {
        throw new IllegalArgumentException(String.format("Expected the field `lastError` to be a primitive type in the JSON string but got `%s`", jsonObj.get("lastError").toString()));
      }
      if ((jsonObj.get("lastMessage") != null && !jsonObj.get("lastMessage").isJsonNull()) && !jsonObj.get("lastMessage").isJsonPrimitive()) {
        throw new IllegalArgumentException(String.format("Expected the field `lastMessage` to be a primitive type in the JSON string but got `%s`", jsonObj.get("lastMessage").toString()));
      }
      if ((jsonObj.get("lastMetric") != null && !jsonObj.get("lastMetric").isJsonNull()) && !jsonObj.get("lastMetric").isJsonPrimitive()) {
        throw new IllegalArgumentException(String.format("Expected the field `lastMetric` to be a primitive type in the JSON string but got `%s`", jsonObj.get("lastMetric").toString()));
      }
      if ((jsonObj.get("lastStats") != null && !jsonObj.get("lastStats").isJsonNull()) && !jsonObj.get("lastStats").isJsonPrimitive()) {
        throw new IllegalArgumentException(String.format("Expected the field `lastStats` to be a primitive type in the JSON string but got `%s`", jsonObj.get("lastStats").toString()));
      }
      if ((jsonObj.get("name") != null && !jsonObj.get("name").isJsonNull()) && !jsonObj.get("name").isJsonPrimitive()) {
        throw new IllegalArgumentException(String.format("Expected the field `name` to be a primitive type in the JSON string but got `%s`", jsonObj.get("name").toString()));
      }
      if ((jsonObj.get("severity") != null && !jsonObj.get("severity").isJsonNull()) && !jsonObj.get("severity").isJsonPrimitive()) {
        throw new IllegalArgumentException(String.format("Expected the field `severity` to be a primitive type in the JSON string but got `%s`", jsonObj.get("severity").toString()));
      }
  }

  public static class CustomTypeAdapterFactory implements TypeAdapterFactory {
    @SuppressWarnings("unchecked")
    @Override
    public <T> TypeAdapter<T> create(Gson gson, TypeToken<T> type) {
       if (!Check.class.isAssignableFrom(type.getRawType())) {
         return null; // this class only serializes 'Check' and its subtypes
       }
       final TypeAdapter<JsonElement> elementAdapter = gson.getAdapter(JsonElement.class);
       final TypeAdapter<Check> thisAdapter
                        = gson.getDelegateAdapter(this, TypeToken.get(Check.class));

       return (TypeAdapter<T>) new TypeAdapter<Check>() {
           @Override
           public void write(JsonWriter out, Check value) throws IOException {
             JsonObject obj = thisAdapter.toJsonTree(value).getAsJsonObject();
             elementAdapter.write(out, obj);
           }

           @Override
           public Check read(JsonReader in) throws IOException {
             JsonElement jsonElement = elementAdapter.read(in);
             validateJsonElement(jsonElement);
             return thisAdapter.fromJsonTree(jsonElement);
           }

       }.nullSafe();
    }
  }

 /**
  * Create an instance of Check given an JSON string
  *
  * @param jsonString JSON string
  * @return An instance of Check
  * @throws IOException if the JSON string is invalid with respect to Check
  */
  public static Check fromJson(String jsonString) throws IOException {
    return JSON.getGson().fromJson(jsonString, Check.class);
  }

 /**
  * Convert an instance of Check to an JSON string
  *
  * @return JSON string
  */
  public String toJson() {
    return JSON.getGson().toJson(this);
  }
}

