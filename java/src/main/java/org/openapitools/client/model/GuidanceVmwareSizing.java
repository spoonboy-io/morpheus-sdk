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
import org.openapitools.client.model.GuidanceVmwareSizingConfig;
import org.openapitools.client.model.GuidanceVmwareSizingPlanAfterAction;
import org.openapitools.client.model.GuidanceVmwareSizingPlanBeforeAction;
import org.openapitools.client.model.GuidanceVmwareSizingResource;
import org.openapitools.client.model.GuidanceVmwareSizingSavings;
import org.openapitools.client.model.GuidanceVmwareSizingType;
import org.openapitools.client.model.GuidanceVmwareSizingZone;
import org.threeten.bp.OffsetDateTime;

/**
 * GuidanceVmwareSizing
 */
@javax.annotation.Generated(value = "org.openapitools.codegen.languages.JavaClientCodegen", date = "2023-08-21T11:26:47.028Z[GMT]")
public class GuidanceVmwareSizing {
  public static final String SERIALIZED_NAME_ID = "id";
  @SerializedName(SERIALIZED_NAME_ID)
  private Long id;

  public static final String SERIALIZED_NAME_DATE_CREATED = "dateCreated";
  @SerializedName(SERIALIZED_NAME_DATE_CREATED)
  private OffsetDateTime dateCreated;

  public static final String SERIALIZED_NAME_LAST_UPDATED = "lastUpdated";
  @SerializedName(SERIALIZED_NAME_LAST_UPDATED)
  private OffsetDateTime lastUpdated;

  public static final String SERIALIZED_NAME_ACTION_CATEGORY = "actionCategory";
  @SerializedName(SERIALIZED_NAME_ACTION_CATEGORY)
  private String actionCategory;

  public static final String SERIALIZED_NAME_ACTION_MESSAGE = "actionMessage";
  @SerializedName(SERIALIZED_NAME_ACTION_MESSAGE)
  private String actionMessage;

  public static final String SERIALIZED_NAME_ACTION_TITLE = "actionTitle";
  @SerializedName(SERIALIZED_NAME_ACTION_TITLE)
  private String actionTitle;

  public static final String SERIALIZED_NAME_ACTION_TYPE = "actionType";
  @SerializedName(SERIALIZED_NAME_ACTION_TYPE)
  private String actionType;

  public static final String SERIALIZED_NAME_ACTION_VALUE = "actionValue";
  @SerializedName(SERIALIZED_NAME_ACTION_VALUE)
  private String actionValue;

  public static final String SERIALIZED_NAME_ACTION_VALUE_TYPE = "actionValueType";
  @SerializedName(SERIALIZED_NAME_ACTION_VALUE_TYPE)
  private String actionValueType;

  public static final String SERIALIZED_NAME_ACTION_PLAN_ID = "actionPlanId";
  @SerializedName(SERIALIZED_NAME_ACTION_PLAN_ID)
  private Long actionPlanId;

  public static final String SERIALIZED_NAME_STATUS_MESSAGE = "statusMessage";
  @SerializedName(SERIALIZED_NAME_STATUS_MESSAGE)
  private String statusMessage;

  public static final String SERIALIZED_NAME_ACCOUNT_ID = "accountId";
  @SerializedName(SERIALIZED_NAME_ACCOUNT_ID)
  private Long accountId;

  public static final String SERIALIZED_NAME_USER_ID = "userId";
  @SerializedName(SERIALIZED_NAME_USER_ID)
  private String userId;

  public static final String SERIALIZED_NAME_SITE_ID = "siteId";
  @SerializedName(SERIALIZED_NAME_SITE_ID)
  private Long siteId;

  public static final String SERIALIZED_NAME_ZONE = "zone";
  @SerializedName(SERIALIZED_NAME_ZONE)
  private GuidanceVmwareSizingZone zone;

  public static final String SERIALIZED_NAME_STATE = "state";
  @SerializedName(SERIALIZED_NAME_STATE)
  private String state;

  public static final String SERIALIZED_NAME_STATE_MESSAGE = "stateMessage";
  @SerializedName(SERIALIZED_NAME_STATE_MESSAGE)
  private String stateMessage;

  public static final String SERIALIZED_NAME_SEVERITY = "severity";
  @SerializedName(SERIALIZED_NAME_SEVERITY)
  private String severity;

  public static final String SERIALIZED_NAME_RESOLVED = "resolved";
  @SerializedName(SERIALIZED_NAME_RESOLVED)
  private Boolean resolved;

  public static final String SERIALIZED_NAME_RESOLVED_MESSAGE = "resolvedMessage";
  @SerializedName(SERIALIZED_NAME_RESOLVED_MESSAGE)
  private String resolvedMessage;

  public static final String SERIALIZED_NAME_REF_TYPE = "refType";
  @SerializedName(SERIALIZED_NAME_REF_TYPE)
  private String refType;

  public static final String SERIALIZED_NAME_REF_ID = "refId";
  @SerializedName(SERIALIZED_NAME_REF_ID)
  private Long refId;

  public static final String SERIALIZED_NAME_REF_NAME = "refName";
  @SerializedName(SERIALIZED_NAME_REF_NAME)
  private String refName;

  public static final String SERIALIZED_NAME_TYPE = "type";
  @SerializedName(SERIALIZED_NAME_TYPE)
  private GuidanceVmwareSizingType type;

  public static final String SERIALIZED_NAME_SAVINGS = "savings";
  @SerializedName(SERIALIZED_NAME_SAVINGS)
  private GuidanceVmwareSizingSavings savings;

  public static final String SERIALIZED_NAME_RESOURCE = "resource";
  @SerializedName(SERIALIZED_NAME_RESOURCE)
  private GuidanceVmwareSizingResource resource;

  public static final String SERIALIZED_NAME_PLAN_BEFORE_ACTION = "planBeforeAction";
  @SerializedName(SERIALIZED_NAME_PLAN_BEFORE_ACTION)
  private GuidanceVmwareSizingPlanBeforeAction planBeforeAction;

  public static final String SERIALIZED_NAME_PLAN_AFTER_ACTION = "planAfterAction";
  @SerializedName(SERIALIZED_NAME_PLAN_AFTER_ACTION)
  private GuidanceVmwareSizingPlanAfterAction planAfterAction;

  public static final String SERIALIZED_NAME_CONFIG = "config";
  @SerializedName(SERIALIZED_NAME_CONFIG)
  private GuidanceVmwareSizingConfig config;


  public GuidanceVmwareSizing id(Long id) {
    
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


  public GuidanceVmwareSizing dateCreated(OffsetDateTime dateCreated) {
    
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


  public GuidanceVmwareSizing lastUpdated(OffsetDateTime lastUpdated) {
    
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


  public GuidanceVmwareSizing actionCategory(String actionCategory) {
    
    this.actionCategory = actionCategory;
    return this;
  }

   /**
   * Get actionCategory
   * @return actionCategory
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public String getActionCategory() {
    return actionCategory;
  }


  public void setActionCategory(String actionCategory) {
    this.actionCategory = actionCategory;
  }


  public GuidanceVmwareSizing actionMessage(String actionMessage) {
    
    this.actionMessage = actionMessage;
    return this;
  }

   /**
   * Get actionMessage
   * @return actionMessage
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public String getActionMessage() {
    return actionMessage;
  }


  public void setActionMessage(String actionMessage) {
    this.actionMessage = actionMessage;
  }


  public GuidanceVmwareSizing actionTitle(String actionTitle) {
    
    this.actionTitle = actionTitle;
    return this;
  }

   /**
   * Get actionTitle
   * @return actionTitle
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public String getActionTitle() {
    return actionTitle;
  }


  public void setActionTitle(String actionTitle) {
    this.actionTitle = actionTitle;
  }


  public GuidanceVmwareSizing actionType(String actionType) {
    
    this.actionType = actionType;
    return this;
  }

   /**
   * Get actionType
   * @return actionType
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public String getActionType() {
    return actionType;
  }


  public void setActionType(String actionType) {
    this.actionType = actionType;
  }


  public GuidanceVmwareSizing actionValue(String actionValue) {
    
    this.actionValue = actionValue;
    return this;
  }

   /**
   * Get actionValue
   * @return actionValue
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public String getActionValue() {
    return actionValue;
  }


  public void setActionValue(String actionValue) {
    this.actionValue = actionValue;
  }


  public GuidanceVmwareSizing actionValueType(String actionValueType) {
    
    this.actionValueType = actionValueType;
    return this;
  }

   /**
   * Get actionValueType
   * @return actionValueType
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public String getActionValueType() {
    return actionValueType;
  }


  public void setActionValueType(String actionValueType) {
    this.actionValueType = actionValueType;
  }


  public GuidanceVmwareSizing actionPlanId(Long actionPlanId) {
    
    this.actionPlanId = actionPlanId;
    return this;
  }

   /**
   * Get actionPlanId
   * @return actionPlanId
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public Long getActionPlanId() {
    return actionPlanId;
  }


  public void setActionPlanId(Long actionPlanId) {
    this.actionPlanId = actionPlanId;
  }


  public GuidanceVmwareSizing statusMessage(String statusMessage) {
    
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


  public GuidanceVmwareSizing accountId(Long accountId) {
    
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


  public GuidanceVmwareSizing userId(String userId) {
    
    this.userId = userId;
    return this;
  }

   /**
   * Get userId
   * @return userId
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public String getUserId() {
    return userId;
  }


  public void setUserId(String userId) {
    this.userId = userId;
  }


  public GuidanceVmwareSizing siteId(Long siteId) {
    
    this.siteId = siteId;
    return this;
  }

   /**
   * Get siteId
   * @return siteId
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public Long getSiteId() {
    return siteId;
  }


  public void setSiteId(Long siteId) {
    this.siteId = siteId;
  }


  public GuidanceVmwareSizing zone(GuidanceVmwareSizingZone zone) {
    
    this.zone = zone;
    return this;
  }

   /**
   * Get zone
   * @return zone
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public GuidanceVmwareSizingZone getZone() {
    return zone;
  }


  public void setZone(GuidanceVmwareSizingZone zone) {
    this.zone = zone;
  }


  public GuidanceVmwareSizing state(String state) {
    
    this.state = state;
    return this;
  }

   /**
   * Get state
   * @return state
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public String getState() {
    return state;
  }


  public void setState(String state) {
    this.state = state;
  }


  public GuidanceVmwareSizing stateMessage(String stateMessage) {
    
    this.stateMessage = stateMessage;
    return this;
  }

   /**
   * Get stateMessage
   * @return stateMessage
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public String getStateMessage() {
    return stateMessage;
  }


  public void setStateMessage(String stateMessage) {
    this.stateMessage = stateMessage;
  }


  public GuidanceVmwareSizing severity(String severity) {
    
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


  public GuidanceVmwareSizing resolved(Boolean resolved) {
    
    this.resolved = resolved;
    return this;
  }

   /**
   * Get resolved
   * @return resolved
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public Boolean getResolved() {
    return resolved;
  }


  public void setResolved(Boolean resolved) {
    this.resolved = resolved;
  }


  public GuidanceVmwareSizing resolvedMessage(String resolvedMessage) {
    
    this.resolvedMessage = resolvedMessage;
    return this;
  }

   /**
   * Get resolvedMessage
   * @return resolvedMessage
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public String getResolvedMessage() {
    return resolvedMessage;
  }


  public void setResolvedMessage(String resolvedMessage) {
    this.resolvedMessage = resolvedMessage;
  }


  public GuidanceVmwareSizing refType(String refType) {
    
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


  public GuidanceVmwareSizing refId(Long refId) {
    
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


  public GuidanceVmwareSizing refName(String refName) {
    
    this.refName = refName;
    return this;
  }

   /**
   * Get refName
   * @return refName
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public String getRefName() {
    return refName;
  }


  public void setRefName(String refName) {
    this.refName = refName;
  }


  public GuidanceVmwareSizing type(GuidanceVmwareSizingType type) {
    
    this.type = type;
    return this;
  }

   /**
   * Get type
   * @return type
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public GuidanceVmwareSizingType getType() {
    return type;
  }


  public void setType(GuidanceVmwareSizingType type) {
    this.type = type;
  }


  public GuidanceVmwareSizing savings(GuidanceVmwareSizingSavings savings) {
    
    this.savings = savings;
    return this;
  }

   /**
   * Get savings
   * @return savings
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public GuidanceVmwareSizingSavings getSavings() {
    return savings;
  }


  public void setSavings(GuidanceVmwareSizingSavings savings) {
    this.savings = savings;
  }


  public GuidanceVmwareSizing resource(GuidanceVmwareSizingResource resource) {
    
    this.resource = resource;
    return this;
  }

   /**
   * Get resource
   * @return resource
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public GuidanceVmwareSizingResource getResource() {
    return resource;
  }


  public void setResource(GuidanceVmwareSizingResource resource) {
    this.resource = resource;
  }


  public GuidanceVmwareSizing planBeforeAction(GuidanceVmwareSizingPlanBeforeAction planBeforeAction) {
    
    this.planBeforeAction = planBeforeAction;
    return this;
  }

   /**
   * Get planBeforeAction
   * @return planBeforeAction
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public GuidanceVmwareSizingPlanBeforeAction getPlanBeforeAction() {
    return planBeforeAction;
  }


  public void setPlanBeforeAction(GuidanceVmwareSizingPlanBeforeAction planBeforeAction) {
    this.planBeforeAction = planBeforeAction;
  }


  public GuidanceVmwareSizing planAfterAction(GuidanceVmwareSizingPlanAfterAction planAfterAction) {
    
    this.planAfterAction = planAfterAction;
    return this;
  }

   /**
   * Get planAfterAction
   * @return planAfterAction
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public GuidanceVmwareSizingPlanAfterAction getPlanAfterAction() {
    return planAfterAction;
  }


  public void setPlanAfterAction(GuidanceVmwareSizingPlanAfterAction planAfterAction) {
    this.planAfterAction = planAfterAction;
  }


  public GuidanceVmwareSizing config(GuidanceVmwareSizingConfig config) {
    
    this.config = config;
    return this;
  }

   /**
   * Get config
   * @return config
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public GuidanceVmwareSizingConfig getConfig() {
    return config;
  }


  public void setConfig(GuidanceVmwareSizingConfig config) {
    this.config = config;
  }


  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (o == null || getClass() != o.getClass()) {
      return false;
    }
    GuidanceVmwareSizing guidanceVmwareSizing = (GuidanceVmwareSizing) o;
    return Objects.equals(this.id, guidanceVmwareSizing.id) &&
        Objects.equals(this.dateCreated, guidanceVmwareSizing.dateCreated) &&
        Objects.equals(this.lastUpdated, guidanceVmwareSizing.lastUpdated) &&
        Objects.equals(this.actionCategory, guidanceVmwareSizing.actionCategory) &&
        Objects.equals(this.actionMessage, guidanceVmwareSizing.actionMessage) &&
        Objects.equals(this.actionTitle, guidanceVmwareSizing.actionTitle) &&
        Objects.equals(this.actionType, guidanceVmwareSizing.actionType) &&
        Objects.equals(this.actionValue, guidanceVmwareSizing.actionValue) &&
        Objects.equals(this.actionValueType, guidanceVmwareSizing.actionValueType) &&
        Objects.equals(this.actionPlanId, guidanceVmwareSizing.actionPlanId) &&
        Objects.equals(this.statusMessage, guidanceVmwareSizing.statusMessage) &&
        Objects.equals(this.accountId, guidanceVmwareSizing.accountId) &&
        Objects.equals(this.userId, guidanceVmwareSizing.userId) &&
        Objects.equals(this.siteId, guidanceVmwareSizing.siteId) &&
        Objects.equals(this.zone, guidanceVmwareSizing.zone) &&
        Objects.equals(this.state, guidanceVmwareSizing.state) &&
        Objects.equals(this.stateMessage, guidanceVmwareSizing.stateMessage) &&
        Objects.equals(this.severity, guidanceVmwareSizing.severity) &&
        Objects.equals(this.resolved, guidanceVmwareSizing.resolved) &&
        Objects.equals(this.resolvedMessage, guidanceVmwareSizing.resolvedMessage) &&
        Objects.equals(this.refType, guidanceVmwareSizing.refType) &&
        Objects.equals(this.refId, guidanceVmwareSizing.refId) &&
        Objects.equals(this.refName, guidanceVmwareSizing.refName) &&
        Objects.equals(this.type, guidanceVmwareSizing.type) &&
        Objects.equals(this.savings, guidanceVmwareSizing.savings) &&
        Objects.equals(this.resource, guidanceVmwareSizing.resource) &&
        Objects.equals(this.planBeforeAction, guidanceVmwareSizing.planBeforeAction) &&
        Objects.equals(this.planAfterAction, guidanceVmwareSizing.planAfterAction) &&
        Objects.equals(this.config, guidanceVmwareSizing.config);
  }

  @Override
  public int hashCode() {
    return Objects.hash(id, dateCreated, lastUpdated, actionCategory, actionMessage, actionTitle, actionType, actionValue, actionValueType, actionPlanId, statusMessage, accountId, userId, siteId, zone, state, stateMessage, severity, resolved, resolvedMessage, refType, refId, refName, type, savings, resource, planBeforeAction, planAfterAction, config);
  }


  @Override
  public String toString() {
    StringBuilder sb = new StringBuilder();
    sb.append("class GuidanceVmwareSizing {\n");
    sb.append("    id: ").append(toIndentedString(id)).append("\n");
    sb.append("    dateCreated: ").append(toIndentedString(dateCreated)).append("\n");
    sb.append("    lastUpdated: ").append(toIndentedString(lastUpdated)).append("\n");
    sb.append("    actionCategory: ").append(toIndentedString(actionCategory)).append("\n");
    sb.append("    actionMessage: ").append(toIndentedString(actionMessage)).append("\n");
    sb.append("    actionTitle: ").append(toIndentedString(actionTitle)).append("\n");
    sb.append("    actionType: ").append(toIndentedString(actionType)).append("\n");
    sb.append("    actionValue: ").append(toIndentedString(actionValue)).append("\n");
    sb.append("    actionValueType: ").append(toIndentedString(actionValueType)).append("\n");
    sb.append("    actionPlanId: ").append(toIndentedString(actionPlanId)).append("\n");
    sb.append("    statusMessage: ").append(toIndentedString(statusMessage)).append("\n");
    sb.append("    accountId: ").append(toIndentedString(accountId)).append("\n");
    sb.append("    userId: ").append(toIndentedString(userId)).append("\n");
    sb.append("    siteId: ").append(toIndentedString(siteId)).append("\n");
    sb.append("    zone: ").append(toIndentedString(zone)).append("\n");
    sb.append("    state: ").append(toIndentedString(state)).append("\n");
    sb.append("    stateMessage: ").append(toIndentedString(stateMessage)).append("\n");
    sb.append("    severity: ").append(toIndentedString(severity)).append("\n");
    sb.append("    resolved: ").append(toIndentedString(resolved)).append("\n");
    sb.append("    resolvedMessage: ").append(toIndentedString(resolvedMessage)).append("\n");
    sb.append("    refType: ").append(toIndentedString(refType)).append("\n");
    sb.append("    refId: ").append(toIndentedString(refId)).append("\n");
    sb.append("    refName: ").append(toIndentedString(refName)).append("\n");
    sb.append("    type: ").append(toIndentedString(type)).append("\n");
    sb.append("    savings: ").append(toIndentedString(savings)).append("\n");
    sb.append("    resource: ").append(toIndentedString(resource)).append("\n");
    sb.append("    planBeforeAction: ").append(toIndentedString(planBeforeAction)).append("\n");
    sb.append("    planAfterAction: ").append(toIndentedString(planAfterAction)).append("\n");
    sb.append("    config: ").append(toIndentedString(config)).append("\n");
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

