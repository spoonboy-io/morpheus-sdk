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
import org.openapitools.client.model.InlineResponse20083LoadBalancerNodeCreatedBy;
import org.openapitools.client.model.InlineResponse20094Network;
import org.threeten.bp.OffsetDateTime;

/**
 * ClusterJobs
 */
@javax.annotation.Generated(value = "org.openapitools.codegen.languages.JavaClientCodegen", date = "2023-08-21T11:26:47.028Z[GMT]")
public class ClusterJobs {
  public static final String SERIALIZED_NAME_ID = "id";
  @SerializedName(SERIALIZED_NAME_ID)
  private Long id;

  public static final String SERIALIZED_NAME_NAME = "name";
  @SerializedName(SERIALIZED_NAME_NAME)
  private String name;

  public static final String SERIALIZED_NAME_LABELS = "labels";
  @SerializedName(SERIALIZED_NAME_LABELS)
  private List<String> labels = null;

  public static final String SERIALIZED_NAME_TYPE = "type";
  @SerializedName(SERIALIZED_NAME_TYPE)
  private InlineResponse20094Network type;

  public static final String SERIALIZED_NAME_JOB_SUMMARY = "jobSummary";
  @SerializedName(SERIALIZED_NAME_JOB_SUMMARY)
  private String jobSummary;

  public static final String SERIALIZED_NAME_SCHEDULE_MODE = "scheduleMode";
  @SerializedName(SERIALIZED_NAME_SCHEDULE_MODE)
  private String scheduleMode;

  public static final String SERIALIZED_NAME_DATE_TIME = "dateTime";
  @SerializedName(SERIALIZED_NAME_DATE_TIME)
  private OffsetDateTime dateTime;

  public static final String SERIALIZED_NAME_STATUS = "status";
  @SerializedName(SERIALIZED_NAME_STATUS)
  private String status;

  public static final String SERIALIZED_NAME_NAMESPACE = "namespace";
  @SerializedName(SERIALIZED_NAME_NAMESPACE)
  private String namespace;

  public static final String SERIALIZED_NAME_CATEGORY = "category";
  @SerializedName(SERIALIZED_NAME_CATEGORY)
  private String category;

  public static final String SERIALIZED_NAME_DESCRIPTION = "description";
  @SerializedName(SERIALIZED_NAME_DESCRIPTION)
  private String description;

  public static final String SERIALIZED_NAME_ENABLED = "enabled";
  @SerializedName(SERIALIZED_NAME_ENABLED)
  private Boolean enabled;

  public static final String SERIALIZED_NAME_DATE_CREATED = "dateCreated";
  @SerializedName(SERIALIZED_NAME_DATE_CREATED)
  private OffsetDateTime dateCreated;

  public static final String SERIALIZED_NAME_LAST_UPDATED = "lastUpdated";
  @SerializedName(SERIALIZED_NAME_LAST_UPDATED)
  private OffsetDateTime lastUpdated;

  public static final String SERIALIZED_NAME_LAST_RUN = "lastRun";
  @SerializedName(SERIALIZED_NAME_LAST_RUN)
  private OffsetDateTime lastRun;

  public static final String SERIALIZED_NAME_LAST_RESULT = "lastResult";
  @SerializedName(SERIALIZED_NAME_LAST_RESULT)
  private String lastResult;

  public static final String SERIALIZED_NAME_CREATED_BY = "createdBy";
  @SerializedName(SERIALIZED_NAME_CREATED_BY)
  private InlineResponse20083LoadBalancerNodeCreatedBy createdBy;

  public static final String SERIALIZED_NAME_TARGET_TYPE = "targetType";
  @SerializedName(SERIALIZED_NAME_TARGET_TYPE)
  private String targetType;

  public static final String SERIALIZED_NAME_TARGETS = "targets";
  @SerializedName(SERIALIZED_NAME_TARGETS)
  private List<Object> targets = null;

  public static final String SERIALIZED_NAME_CUSTOM_CONFIG = "customConfig";
  @SerializedName(SERIALIZED_NAME_CUSTOM_CONFIG)
  private Object customConfig;

  public static final String SERIALIZED_NAME_CUSTOM_OPTIONS = "customOptions";
  @SerializedName(SERIALIZED_NAME_CUSTOM_OPTIONS)
  private Object customOptions;


  public ClusterJobs id(Long id) {
    
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


  public ClusterJobs name(String name) {
    
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


  public ClusterJobs labels(List<String> labels) {
    
    this.labels = labels;
    return this;
  }

  public ClusterJobs addLabelsItem(String labelsItem) {
    if (this.labels == null) {
      this.labels = new ArrayList<String>();
    }
    this.labels.add(labelsItem);
    return this;
  }

   /**
   * Get labels
   * @return labels
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public List<String> getLabels() {
    return labels;
  }


  public void setLabels(List<String> labels) {
    this.labels = labels;
  }


  public ClusterJobs type(InlineResponse20094Network type) {
    
    this.type = type;
    return this;
  }

   /**
   * Get type
   * @return type
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public InlineResponse20094Network getType() {
    return type;
  }


  public void setType(InlineResponse20094Network type) {
    this.type = type;
  }


  public ClusterJobs jobSummary(String jobSummary) {
    
    this.jobSummary = jobSummary;
    return this;
  }

   /**
   * Get jobSummary
   * @return jobSummary
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public String getJobSummary() {
    return jobSummary;
  }


  public void setJobSummary(String jobSummary) {
    this.jobSummary = jobSummary;
  }


  public ClusterJobs scheduleMode(String scheduleMode) {
    
    this.scheduleMode = scheduleMode;
    return this;
  }

   /**
   * Get scheduleMode
   * @return scheduleMode
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public String getScheduleMode() {
    return scheduleMode;
  }


  public void setScheduleMode(String scheduleMode) {
    this.scheduleMode = scheduleMode;
  }


  public ClusterJobs dateTime(OffsetDateTime dateTime) {
    
    this.dateTime = dateTime;
    return this;
  }

   /**
   * Get dateTime
   * @return dateTime
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public OffsetDateTime getDateTime() {
    return dateTime;
  }


  public void setDateTime(OffsetDateTime dateTime) {
    this.dateTime = dateTime;
  }


  public ClusterJobs status(String status) {
    
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


  public ClusterJobs namespace(String namespace) {
    
    this.namespace = namespace;
    return this;
  }

   /**
   * Get namespace
   * @return namespace
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public String getNamespace() {
    return namespace;
  }


  public void setNamespace(String namespace) {
    this.namespace = namespace;
  }


  public ClusterJobs category(String category) {
    
    this.category = category;
    return this;
  }

   /**
   * Get category
   * @return category
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public String getCategory() {
    return category;
  }


  public void setCategory(String category) {
    this.category = category;
  }


  public ClusterJobs description(String description) {
    
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


  public ClusterJobs enabled(Boolean enabled) {
    
    this.enabled = enabled;
    return this;
  }

   /**
   * Get enabled
   * @return enabled
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public Boolean getEnabled() {
    return enabled;
  }


  public void setEnabled(Boolean enabled) {
    this.enabled = enabled;
  }


  public ClusterJobs dateCreated(OffsetDateTime dateCreated) {
    
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


  public ClusterJobs lastUpdated(OffsetDateTime lastUpdated) {
    
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


  public ClusterJobs lastRun(OffsetDateTime lastRun) {
    
    this.lastRun = lastRun;
    return this;
  }

   /**
   * Get lastRun
   * @return lastRun
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public OffsetDateTime getLastRun() {
    return lastRun;
  }


  public void setLastRun(OffsetDateTime lastRun) {
    this.lastRun = lastRun;
  }


  public ClusterJobs lastResult(String lastResult) {
    
    this.lastResult = lastResult;
    return this;
  }

   /**
   * Get lastResult
   * @return lastResult
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public String getLastResult() {
    return lastResult;
  }


  public void setLastResult(String lastResult) {
    this.lastResult = lastResult;
  }


  public ClusterJobs createdBy(InlineResponse20083LoadBalancerNodeCreatedBy createdBy) {
    
    this.createdBy = createdBy;
    return this;
  }

   /**
   * Get createdBy
   * @return createdBy
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public InlineResponse20083LoadBalancerNodeCreatedBy getCreatedBy() {
    return createdBy;
  }


  public void setCreatedBy(InlineResponse20083LoadBalancerNodeCreatedBy createdBy) {
    this.createdBy = createdBy;
  }


  public ClusterJobs targetType(String targetType) {
    
    this.targetType = targetType;
    return this;
  }

   /**
   * Get targetType
   * @return targetType
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public String getTargetType() {
    return targetType;
  }


  public void setTargetType(String targetType) {
    this.targetType = targetType;
  }


  public ClusterJobs targets(List<Object> targets) {
    
    this.targets = targets;
    return this;
  }

  public ClusterJobs addTargetsItem(Object targetsItem) {
    if (this.targets == null) {
      this.targets = new ArrayList<Object>();
    }
    this.targets.add(targetsItem);
    return this;
  }

   /**
   * Get targets
   * @return targets
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public List<Object> getTargets() {
    return targets;
  }


  public void setTargets(List<Object> targets) {
    this.targets = targets;
  }


  public ClusterJobs customConfig(Object customConfig) {
    
    this.customConfig = customConfig;
    return this;
  }

   /**
   * Get customConfig
   * @return customConfig
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public Object getCustomConfig() {
    return customConfig;
  }


  public void setCustomConfig(Object customConfig) {
    this.customConfig = customConfig;
  }


  public ClusterJobs customOptions(Object customOptions) {
    
    this.customOptions = customOptions;
    return this;
  }

   /**
   * Get customOptions
   * @return customOptions
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public Object getCustomOptions() {
    return customOptions;
  }


  public void setCustomOptions(Object customOptions) {
    this.customOptions = customOptions;
  }


  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (o == null || getClass() != o.getClass()) {
      return false;
    }
    ClusterJobs clusterJobs = (ClusterJobs) o;
    return Objects.equals(this.id, clusterJobs.id) &&
        Objects.equals(this.name, clusterJobs.name) &&
        Objects.equals(this.labels, clusterJobs.labels) &&
        Objects.equals(this.type, clusterJobs.type) &&
        Objects.equals(this.jobSummary, clusterJobs.jobSummary) &&
        Objects.equals(this.scheduleMode, clusterJobs.scheduleMode) &&
        Objects.equals(this.dateTime, clusterJobs.dateTime) &&
        Objects.equals(this.status, clusterJobs.status) &&
        Objects.equals(this.namespace, clusterJobs.namespace) &&
        Objects.equals(this.category, clusterJobs.category) &&
        Objects.equals(this.description, clusterJobs.description) &&
        Objects.equals(this.enabled, clusterJobs.enabled) &&
        Objects.equals(this.dateCreated, clusterJobs.dateCreated) &&
        Objects.equals(this.lastUpdated, clusterJobs.lastUpdated) &&
        Objects.equals(this.lastRun, clusterJobs.lastRun) &&
        Objects.equals(this.lastResult, clusterJobs.lastResult) &&
        Objects.equals(this.createdBy, clusterJobs.createdBy) &&
        Objects.equals(this.targetType, clusterJobs.targetType) &&
        Objects.equals(this.targets, clusterJobs.targets) &&
        Objects.equals(this.customConfig, clusterJobs.customConfig) &&
        Objects.equals(this.customOptions, clusterJobs.customOptions);
  }

  @Override
  public int hashCode() {
    return Objects.hash(id, name, labels, type, jobSummary, scheduleMode, dateTime, status, namespace, category, description, enabled, dateCreated, lastUpdated, lastRun, lastResult, createdBy, targetType, targets, customConfig, customOptions);
  }


  @Override
  public String toString() {
    StringBuilder sb = new StringBuilder();
    sb.append("class ClusterJobs {\n");
    sb.append("    id: ").append(toIndentedString(id)).append("\n");
    sb.append("    name: ").append(toIndentedString(name)).append("\n");
    sb.append("    labels: ").append(toIndentedString(labels)).append("\n");
    sb.append("    type: ").append(toIndentedString(type)).append("\n");
    sb.append("    jobSummary: ").append(toIndentedString(jobSummary)).append("\n");
    sb.append("    scheduleMode: ").append(toIndentedString(scheduleMode)).append("\n");
    sb.append("    dateTime: ").append(toIndentedString(dateTime)).append("\n");
    sb.append("    status: ").append(toIndentedString(status)).append("\n");
    sb.append("    namespace: ").append(toIndentedString(namespace)).append("\n");
    sb.append("    category: ").append(toIndentedString(category)).append("\n");
    sb.append("    description: ").append(toIndentedString(description)).append("\n");
    sb.append("    enabled: ").append(toIndentedString(enabled)).append("\n");
    sb.append("    dateCreated: ").append(toIndentedString(dateCreated)).append("\n");
    sb.append("    lastUpdated: ").append(toIndentedString(lastUpdated)).append("\n");
    sb.append("    lastRun: ").append(toIndentedString(lastRun)).append("\n");
    sb.append("    lastResult: ").append(toIndentedString(lastResult)).append("\n");
    sb.append("    createdBy: ").append(toIndentedString(createdBy)).append("\n");
    sb.append("    targetType: ").append(toIndentedString(targetType)).append("\n");
    sb.append("    targets: ").append(toIndentedString(targets)).append("\n");
    sb.append("    customConfig: ").append(toIndentedString(customConfig)).append("\n");
    sb.append("    customOptions: ").append(toIndentedString(customOptions)).append("\n");
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

