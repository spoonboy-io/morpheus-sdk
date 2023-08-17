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
import java.util.ArrayList;
import java.util.Arrays;
import java.util.List;
import org.openapitools.client.model.ActivityActivityInnerUser;
import org.openapitools.client.model.CheckCheckType;
import org.openapitools.client.model.CheckGroupInstance;
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
 * CheckGroup
 */
@javax.annotation.Generated(value = "org.openapitools.codegen.languages.JavaClientCodegen", date = "2023-08-17T13:37:08.123540Z[Etc/UTC]")
public class CheckGroup {
  public static final String SERIALIZED_NAME_ID = "id";
  @SerializedName(SERIALIZED_NAME_ID)
  private Long id;

  public static final String SERIALIZED_NAME_ACCOUNT = "account";
  @SerializedName(SERIALIZED_NAME_ACCOUNT)
  private UpdateBlueprintPermissionsRequestResourcePermissionSitesInner account;

  public static final String SERIALIZED_NAME_INSTANCE = "instance";
  @SerializedName(SERIALIZED_NAME_INSTANCE)
  private CheckGroupInstance instance;

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

  public static final String SERIALIZED_NAME_OUTAGE_TIME = "outageTime";
  @SerializedName(SERIALIZED_NAME_OUTAGE_TIME)
  private Long outageTime;

  public static final String SERIALIZED_NAME_LAST_TIMER = "lastTimer";
  @SerializedName(SERIALIZED_NAME_LAST_TIMER)
  private Long lastTimer;

  public static final String SERIALIZED_NAME_HEALTH = "health";
  @SerializedName(SERIALIZED_NAME_HEALTH)
  private Long health;

  public static final String SERIALIZED_NAME_HISTORY = "history";
  @SerializedName(SERIALIZED_NAME_HISTORY)
  private String history;

  public static final String SERIALIZED_NAME_MIN_HAPPY = "minHappy";
  @SerializedName(SERIALIZED_NAME_MIN_HAPPY)
  private Long minHappy;

  public static final String SERIALIZED_NAME_LAST_METRIC = "lastMetric";
  @SerializedName(SERIALIZED_NAME_LAST_METRIC)
  private String lastMetric;

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
  private ActivityActivityInnerUser createdBy;

  public static final String SERIALIZED_NAME_DATE_CREATED = "dateCreated";
  @SerializedName(SERIALIZED_NAME_DATE_CREATED)
  private OffsetDateTime dateCreated;

  public static final String SERIALIZED_NAME_LAST_UPDATED = "lastUpdated";
  @SerializedName(SERIALIZED_NAME_LAST_UPDATED)
  private OffsetDateTime lastUpdated;

  public static final String SERIALIZED_NAME_AVAILABILITY = "availability";
  @SerializedName(SERIALIZED_NAME_AVAILABILITY)
  private BigDecimal availability;

  public static final String SERIALIZED_NAME_CHECK_TYPE = "checkType";
  @SerializedName(SERIALIZED_NAME_CHECK_TYPE)
  private CheckCheckType checkType;

  public static final String SERIALIZED_NAME_CHECKS = "checks";
  @SerializedName(SERIALIZED_NAME_CHECKS)
  private List<Long> checks;

  public CheckGroup() {
  }

  public CheckGroup id(Long id) {
    
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


  public CheckGroup account(UpdateBlueprintPermissionsRequestResourcePermissionSitesInner account) {
    
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


  public CheckGroup instance(CheckGroupInstance instance) {
    
    this.instance = instance;
    return this;
  }

   /**
   * Get instance
   * @return instance
  **/
  @javax.annotation.Nullable
  public CheckGroupInstance getInstance() {
    return instance;
  }


  public void setInstance(CheckGroupInstance instance) {
    this.instance = instance;
  }


  public CheckGroup name(String name) {
    
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


  public CheckGroup description(String description) {
    
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


  public CheckGroup inUptime(Boolean inUptime) {
    
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


  public CheckGroup lastCheckStatus(String lastCheckStatus) {
    
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


  public CheckGroup lastWarningDate(OffsetDateTime lastWarningDate) {
    
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


  public CheckGroup lastErrorDate(OffsetDateTime lastErrorDate) {
    
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


  public CheckGroup lastSuccessDate(OffsetDateTime lastSuccessDate) {
    
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


  public CheckGroup lastRunDate(OffsetDateTime lastRunDate) {
    
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


  public CheckGroup lastError(String lastError) {
    
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


  public CheckGroup outageTime(Long outageTime) {
    
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


  public CheckGroup lastTimer(Long lastTimer) {
    
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


  public CheckGroup health(Long health) {
    
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


  public CheckGroup history(String history) {
    
    this.history = history;
    return this;
  }

   /**
   * Get history
   * @return history
  **/
  @javax.annotation.Nullable
  public String getHistory() {
    return history;
  }


  public void setHistory(String history) {
    this.history = history;
  }


  public CheckGroup minHappy(Long minHappy) {
    
    this.minHappy = minHappy;
    return this;
  }

   /**
   * Get minHappy
   * @return minHappy
  **/
  @javax.annotation.Nullable
  public Long getMinHappy() {
    return minHappy;
  }


  public void setMinHappy(Long minHappy) {
    this.minHappy = minHappy;
  }


  public CheckGroup lastMetric(String lastMetric) {
    
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


  public CheckGroup severity(String severity) {
    
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


  public CheckGroup createIncident(Boolean createIncident) {
    
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


  public CheckGroup muted(Boolean muted) {
    
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


  public CheckGroup createdBy(ActivityActivityInnerUser createdBy) {
    
    this.createdBy = createdBy;
    return this;
  }

   /**
   * Get createdBy
   * @return createdBy
  **/
  @javax.annotation.Nullable
  public ActivityActivityInnerUser getCreatedBy() {
    return createdBy;
  }


  public void setCreatedBy(ActivityActivityInnerUser createdBy) {
    this.createdBy = createdBy;
  }


  public CheckGroup dateCreated(OffsetDateTime dateCreated) {
    
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


  public CheckGroup lastUpdated(OffsetDateTime lastUpdated) {
    
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


  public CheckGroup availability(BigDecimal availability) {
    
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


  public CheckGroup checkType(CheckCheckType checkType) {
    
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


  public CheckGroup checks(List<Long> checks) {
    
    this.checks = checks;
    return this;
  }

  public CheckGroup addChecksItem(Long checksItem) {
    if (this.checks == null) {
      this.checks = new ArrayList<>();
    }
    this.checks.add(checksItem);
    return this;
  }

   /**
   * Get checks
   * @return checks
  **/
  @javax.annotation.Nullable
  public List<Long> getChecks() {
    return checks;
  }


  public void setChecks(List<Long> checks) {
    this.checks = checks;
  }



  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (o == null || getClass() != o.getClass()) {
      return false;
    }
    CheckGroup checkGroup = (CheckGroup) o;
    return Objects.equals(this.id, checkGroup.id) &&
        Objects.equals(this.account, checkGroup.account) &&
        Objects.equals(this.instance, checkGroup.instance) &&
        Objects.equals(this.name, checkGroup.name) &&
        Objects.equals(this.description, checkGroup.description) &&
        Objects.equals(this.inUptime, checkGroup.inUptime) &&
        Objects.equals(this.lastCheckStatus, checkGroup.lastCheckStatus) &&
        Objects.equals(this.lastWarningDate, checkGroup.lastWarningDate) &&
        Objects.equals(this.lastErrorDate, checkGroup.lastErrorDate) &&
        Objects.equals(this.lastSuccessDate, checkGroup.lastSuccessDate) &&
        Objects.equals(this.lastRunDate, checkGroup.lastRunDate) &&
        Objects.equals(this.lastError, checkGroup.lastError) &&
        Objects.equals(this.outageTime, checkGroup.outageTime) &&
        Objects.equals(this.lastTimer, checkGroup.lastTimer) &&
        Objects.equals(this.health, checkGroup.health) &&
        Objects.equals(this.history, checkGroup.history) &&
        Objects.equals(this.minHappy, checkGroup.minHappy) &&
        Objects.equals(this.lastMetric, checkGroup.lastMetric) &&
        Objects.equals(this.severity, checkGroup.severity) &&
        Objects.equals(this.createIncident, checkGroup.createIncident) &&
        Objects.equals(this.muted, checkGroup.muted) &&
        Objects.equals(this.createdBy, checkGroup.createdBy) &&
        Objects.equals(this.dateCreated, checkGroup.dateCreated) &&
        Objects.equals(this.lastUpdated, checkGroup.lastUpdated) &&
        Objects.equals(this.availability, checkGroup.availability) &&
        Objects.equals(this.checkType, checkGroup.checkType) &&
        Objects.equals(this.checks, checkGroup.checks);
  }

  private static <T> boolean equalsNullable(JsonNullable<T> a, JsonNullable<T> b) {
    return a == b || (a != null && b != null && a.isPresent() && b.isPresent() && Objects.deepEquals(a.get(), b.get()));
  }

  @Override
  public int hashCode() {
    return Objects.hash(id, account, instance, name, description, inUptime, lastCheckStatus, lastWarningDate, lastErrorDate, lastSuccessDate, lastRunDate, lastError, outageTime, lastTimer, health, history, minHappy, lastMetric, severity, createIncident, muted, createdBy, dateCreated, lastUpdated, availability, checkType, checks);
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
    sb.append("class CheckGroup {\n");
    sb.append("    id: ").append(toIndentedString(id)).append("\n");
    sb.append("    account: ").append(toIndentedString(account)).append("\n");
    sb.append("    instance: ").append(toIndentedString(instance)).append("\n");
    sb.append("    name: ").append(toIndentedString(name)).append("\n");
    sb.append("    description: ").append(toIndentedString(description)).append("\n");
    sb.append("    inUptime: ").append(toIndentedString(inUptime)).append("\n");
    sb.append("    lastCheckStatus: ").append(toIndentedString(lastCheckStatus)).append("\n");
    sb.append("    lastWarningDate: ").append(toIndentedString(lastWarningDate)).append("\n");
    sb.append("    lastErrorDate: ").append(toIndentedString(lastErrorDate)).append("\n");
    sb.append("    lastSuccessDate: ").append(toIndentedString(lastSuccessDate)).append("\n");
    sb.append("    lastRunDate: ").append(toIndentedString(lastRunDate)).append("\n");
    sb.append("    lastError: ").append(toIndentedString(lastError)).append("\n");
    sb.append("    outageTime: ").append(toIndentedString(outageTime)).append("\n");
    sb.append("    lastTimer: ").append(toIndentedString(lastTimer)).append("\n");
    sb.append("    health: ").append(toIndentedString(health)).append("\n");
    sb.append("    history: ").append(toIndentedString(history)).append("\n");
    sb.append("    minHappy: ").append(toIndentedString(minHappy)).append("\n");
    sb.append("    lastMetric: ").append(toIndentedString(lastMetric)).append("\n");
    sb.append("    severity: ").append(toIndentedString(severity)).append("\n");
    sb.append("    createIncident: ").append(toIndentedString(createIncident)).append("\n");
    sb.append("    muted: ").append(toIndentedString(muted)).append("\n");
    sb.append("    createdBy: ").append(toIndentedString(createdBy)).append("\n");
    sb.append("    dateCreated: ").append(toIndentedString(dateCreated)).append("\n");
    sb.append("    lastUpdated: ").append(toIndentedString(lastUpdated)).append("\n");
    sb.append("    availability: ").append(toIndentedString(availability)).append("\n");
    sb.append("    checkType: ").append(toIndentedString(checkType)).append("\n");
    sb.append("    checks: ").append(toIndentedString(checks)).append("\n");
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
    openapiFields.add("instance");
    openapiFields.add("name");
    openapiFields.add("description");
    openapiFields.add("inUptime");
    openapiFields.add("lastCheckStatus");
    openapiFields.add("lastWarningDate");
    openapiFields.add("lastErrorDate");
    openapiFields.add("lastSuccessDate");
    openapiFields.add("lastRunDate");
    openapiFields.add("lastError");
    openapiFields.add("outageTime");
    openapiFields.add("lastTimer");
    openapiFields.add("health");
    openapiFields.add("history");
    openapiFields.add("minHappy");
    openapiFields.add("lastMetric");
    openapiFields.add("severity");
    openapiFields.add("createIncident");
    openapiFields.add("muted");
    openapiFields.add("createdBy");
    openapiFields.add("dateCreated");
    openapiFields.add("lastUpdated");
    openapiFields.add("availability");
    openapiFields.add("checkType");
    openapiFields.add("checks");

    // a set of required properties/fields (JSON key names)
    openapiRequiredFields = new HashSet<String>();
  }

 /**
  * Validates the JSON Element and throws an exception if issues found
  *
  * @param jsonElement JSON Element
  * @throws IOException if the JSON Element is invalid with respect to CheckGroup
  */
  public static void validateJsonElement(JsonElement jsonElement) throws IOException {
      if (jsonElement == null) {
        if (!CheckGroup.openapiRequiredFields.isEmpty()) { // has required fields but JSON element is null
          throw new IllegalArgumentException(String.format("The required field(s) %s in CheckGroup is not found in the empty JSON string", CheckGroup.openapiRequiredFields.toString()));
        }
      }

      Set<Entry<String, JsonElement>> entries = jsonElement.getAsJsonObject().entrySet();
      // check to see if the JSON string contains additional fields
      for (Entry<String, JsonElement> entry : entries) {
        if (!CheckGroup.openapiFields.contains(entry.getKey())) {
          throw new IllegalArgumentException(String.format("The field `%s` in the JSON string is not defined in the `CheckGroup` properties. JSON: %s", entry.getKey(), jsonElement.toString()));
        }
      }
        JsonObject jsonObj = jsonElement.getAsJsonObject();
      // validate the optional field `account`
      if (jsonObj.get("account") != null && !jsonObj.get("account").isJsonNull()) {
        UpdateBlueprintPermissionsRequestResourcePermissionSitesInner.validateJsonElement(jsonObj.get("account"));
      }
      // validate the optional field `instance`
      if (jsonObj.get("instance") != null && !jsonObj.get("instance").isJsonNull()) {
        CheckGroupInstance.validateJsonElement(jsonObj.get("instance"));
      }
      if ((jsonObj.get("name") != null && !jsonObj.get("name").isJsonNull()) && !jsonObj.get("name").isJsonPrimitive()) {
        throw new IllegalArgumentException(String.format("Expected the field `name` to be a primitive type in the JSON string but got `%s`", jsonObj.get("name").toString()));
      }
      if ((jsonObj.get("description") != null && !jsonObj.get("description").isJsonNull()) && !jsonObj.get("description").isJsonPrimitive()) {
        throw new IllegalArgumentException(String.format("Expected the field `description` to be a primitive type in the JSON string but got `%s`", jsonObj.get("description").toString()));
      }
      if ((jsonObj.get("lastCheckStatus") != null && !jsonObj.get("lastCheckStatus").isJsonNull()) && !jsonObj.get("lastCheckStatus").isJsonPrimitive()) {
        throw new IllegalArgumentException(String.format("Expected the field `lastCheckStatus` to be a primitive type in the JSON string but got `%s`", jsonObj.get("lastCheckStatus").toString()));
      }
      if ((jsonObj.get("lastError") != null && !jsonObj.get("lastError").isJsonNull()) && !jsonObj.get("lastError").isJsonPrimitive()) {
        throw new IllegalArgumentException(String.format("Expected the field `lastError` to be a primitive type in the JSON string but got `%s`", jsonObj.get("lastError").toString()));
      }
      if ((jsonObj.get("history") != null && !jsonObj.get("history").isJsonNull()) && !jsonObj.get("history").isJsonPrimitive()) {
        throw new IllegalArgumentException(String.format("Expected the field `history` to be a primitive type in the JSON string but got `%s`", jsonObj.get("history").toString()));
      }
      if ((jsonObj.get("lastMetric") != null && !jsonObj.get("lastMetric").isJsonNull()) && !jsonObj.get("lastMetric").isJsonPrimitive()) {
        throw new IllegalArgumentException(String.format("Expected the field `lastMetric` to be a primitive type in the JSON string but got `%s`", jsonObj.get("lastMetric").toString()));
      }
      if ((jsonObj.get("severity") != null && !jsonObj.get("severity").isJsonNull()) && !jsonObj.get("severity").isJsonPrimitive()) {
        throw new IllegalArgumentException(String.format("Expected the field `severity` to be a primitive type in the JSON string but got `%s`", jsonObj.get("severity").toString()));
      }
      // validate the optional field `createdBy`
      if (jsonObj.get("createdBy") != null && !jsonObj.get("createdBy").isJsonNull()) {
        ActivityActivityInnerUser.validateJsonElement(jsonObj.get("createdBy"));
      }
      // validate the optional field `checkType`
      if (jsonObj.get("checkType") != null && !jsonObj.get("checkType").isJsonNull()) {
        CheckCheckType.validateJsonElement(jsonObj.get("checkType"));
      }
      // ensure the optional json data is an array if present
      if (jsonObj.get("checks") != null && !jsonObj.get("checks").isJsonNull() && !jsonObj.get("checks").isJsonArray()) {
        throw new IllegalArgumentException(String.format("Expected the field `checks` to be an array in the JSON string but got `%s`", jsonObj.get("checks").toString()));
      }
  }

  public static class CustomTypeAdapterFactory implements TypeAdapterFactory {
    @SuppressWarnings("unchecked")
    @Override
    public <T> TypeAdapter<T> create(Gson gson, TypeToken<T> type) {
       if (!CheckGroup.class.isAssignableFrom(type.getRawType())) {
         return null; // this class only serializes 'CheckGroup' and its subtypes
       }
       final TypeAdapter<JsonElement> elementAdapter = gson.getAdapter(JsonElement.class);
       final TypeAdapter<CheckGroup> thisAdapter
                        = gson.getDelegateAdapter(this, TypeToken.get(CheckGroup.class));

       return (TypeAdapter<T>) new TypeAdapter<CheckGroup>() {
           @Override
           public void write(JsonWriter out, CheckGroup value) throws IOException {
             JsonObject obj = thisAdapter.toJsonTree(value).getAsJsonObject();
             elementAdapter.write(out, obj);
           }

           @Override
           public CheckGroup read(JsonReader in) throws IOException {
             JsonElement jsonElement = elementAdapter.read(in);
             validateJsonElement(jsonElement);
             return thisAdapter.fromJsonTree(jsonElement);
           }

       }.nullSafe();
    }
  }

 /**
  * Create an instance of CheckGroup given an JSON string
  *
  * @param jsonString JSON string
  * @return An instance of CheckGroup
  * @throws IOException if the JSON string is invalid with respect to CheckGroup
  */
  public static CheckGroup fromJson(String jsonString) throws IOException {
    return JSON.getGson().fromJson(jsonString, CheckGroup.class);
  }

 /**
  * Convert an instance of CheckGroup to an JSON string
  *
  * @return JSON string
  */
  public String toJson() {
    return JSON.getGson().toJson(this);
  }
}

