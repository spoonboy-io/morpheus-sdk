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
import org.openapitools.client.model.OptionTypeListCredential;
import org.openapitools.client.model.TaskAnsiblePlaybookConfigFile;
import org.openapitools.client.model.TaskPowershellConfigTaskOptions;
import org.openapitools.client.model.TaskPowershellConfigTaskType;
import org.threeten.bp.OffsetDateTime;

/**
 * TaskPowershellConfig
 */
@javax.annotation.Generated(value = "org.openapitools.codegen.languages.JavaClientCodegen", date = "2023-08-21T11:26:47.028Z[GMT]")
public class TaskPowershellConfig {
  public static final String SERIALIZED_NAME_ID = "id";
  @SerializedName(SERIALIZED_NAME_ID)
  private Long id;

  public static final String SERIALIZED_NAME_ACCOUNT_ID = "accountId";
  @SerializedName(SERIALIZED_NAME_ACCOUNT_ID)
  private Long accountId;

  public static final String SERIALIZED_NAME_NAME = "name";
  @SerializedName(SERIALIZED_NAME_NAME)
  private String name;

  public static final String SERIALIZED_NAME_CODE = "code";
  @SerializedName(SERIALIZED_NAME_CODE)
  private String code;

  public static final String SERIALIZED_NAME_TASK_TYPE = "taskType";
  @SerializedName(SERIALIZED_NAME_TASK_TYPE)
  private TaskPowershellConfigTaskType taskType;

  public static final String SERIALIZED_NAME_LABELS = "labels";
  @SerializedName(SERIALIZED_NAME_LABELS)
  private List<String> labels = null;

  public static final String SERIALIZED_NAME_VISIBILITY = "visibility";
  @SerializedName(SERIALIZED_NAME_VISIBILITY)
  private String visibility;

  public static final String SERIALIZED_NAME_TASK_OPTIONS = "taskOptions";
  @SerializedName(SERIALIZED_NAME_TASK_OPTIONS)
  private TaskPowershellConfigTaskOptions taskOptions;

  public static final String SERIALIZED_NAME_FILE = "file";
  @SerializedName(SERIALIZED_NAME_FILE)
  private TaskAnsiblePlaybookConfigFile file;

  public static final String SERIALIZED_NAME_RESULT_TYPE = "resultType";
  @SerializedName(SERIALIZED_NAME_RESULT_TYPE)
  private String resultType;

  public static final String SERIALIZED_NAME_EXECUTE_TARGET = "executeTarget";
  @SerializedName(SERIALIZED_NAME_EXECUTE_TARGET)
  private String executeTarget;

  public static final String SERIALIZED_NAME_RETRYABLE = "retryable";
  @SerializedName(SERIALIZED_NAME_RETRYABLE)
  private Boolean retryable;

  public static final String SERIALIZED_NAME_RETRY_COUNT = "retryCount";
  @SerializedName(SERIALIZED_NAME_RETRY_COUNT)
  private Long retryCount;

  public static final String SERIALIZED_NAME_RETRY_DELAY_SECONDS = "retryDelaySeconds";
  @SerializedName(SERIALIZED_NAME_RETRY_DELAY_SECONDS)
  private Long retryDelaySeconds;

  public static final String SERIALIZED_NAME_ALLOW_CUSTOM_CONFIG = "allowCustomConfig";
  @SerializedName(SERIALIZED_NAME_ALLOW_CUSTOM_CONFIG)
  private Boolean allowCustomConfig;

  public static final String SERIALIZED_NAME_CREDENTIAL = "credential";
  @SerializedName(SERIALIZED_NAME_CREDENTIAL)
  private OptionTypeListCredential credential;

  public static final String SERIALIZED_NAME_DATE_CREATED = "dateCreated";
  @SerializedName(SERIALIZED_NAME_DATE_CREATED)
  private OffsetDateTime dateCreated;

  public static final String SERIALIZED_NAME_LAST_UPDATED = "lastUpdated";
  @SerializedName(SERIALIZED_NAME_LAST_UPDATED)
  private OffsetDateTime lastUpdated;


  public TaskPowershellConfig id(Long id) {
    
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


  public TaskPowershellConfig accountId(Long accountId) {
    
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


  public TaskPowershellConfig name(String name) {
    
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


  public TaskPowershellConfig code(String code) {
    
    this.code = code;
    return this;
  }

   /**
   * Get code
   * @return code
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public String getCode() {
    return code;
  }


  public void setCode(String code) {
    this.code = code;
  }


  public TaskPowershellConfig taskType(TaskPowershellConfigTaskType taskType) {
    
    this.taskType = taskType;
    return this;
  }

   /**
   * Get taskType
   * @return taskType
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public TaskPowershellConfigTaskType getTaskType() {
    return taskType;
  }


  public void setTaskType(TaskPowershellConfigTaskType taskType) {
    this.taskType = taskType;
  }


  public TaskPowershellConfig labels(List<String> labels) {
    
    this.labels = labels;
    return this;
  }

  public TaskPowershellConfig addLabelsItem(String labelsItem) {
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


  public TaskPowershellConfig visibility(String visibility) {
    
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


  public TaskPowershellConfig taskOptions(TaskPowershellConfigTaskOptions taskOptions) {
    
    this.taskOptions = taskOptions;
    return this;
  }

   /**
   * Get taskOptions
   * @return taskOptions
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public TaskPowershellConfigTaskOptions getTaskOptions() {
    return taskOptions;
  }


  public void setTaskOptions(TaskPowershellConfigTaskOptions taskOptions) {
    this.taskOptions = taskOptions;
  }


  public TaskPowershellConfig file(TaskAnsiblePlaybookConfigFile file) {
    
    this.file = file;
    return this;
  }

   /**
   * Get file
   * @return file
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public TaskAnsiblePlaybookConfigFile getFile() {
    return file;
  }


  public void setFile(TaskAnsiblePlaybookConfigFile file) {
    this.file = file;
  }


  public TaskPowershellConfig resultType(String resultType) {
    
    this.resultType = resultType;
    return this;
  }

   /**
   * Get resultType
   * @return resultType
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public String getResultType() {
    return resultType;
  }


  public void setResultType(String resultType) {
    this.resultType = resultType;
  }


  public TaskPowershellConfig executeTarget(String executeTarget) {
    
    this.executeTarget = executeTarget;
    return this;
  }

   /**
   * Get executeTarget
   * @return executeTarget
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public String getExecuteTarget() {
    return executeTarget;
  }


  public void setExecuteTarget(String executeTarget) {
    this.executeTarget = executeTarget;
  }


  public TaskPowershellConfig retryable(Boolean retryable) {
    
    this.retryable = retryable;
    return this;
  }

   /**
   * Get retryable
   * @return retryable
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public Boolean getRetryable() {
    return retryable;
  }


  public void setRetryable(Boolean retryable) {
    this.retryable = retryable;
  }


  public TaskPowershellConfig retryCount(Long retryCount) {
    
    this.retryCount = retryCount;
    return this;
  }

   /**
   * Get retryCount
   * @return retryCount
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public Long getRetryCount() {
    return retryCount;
  }


  public void setRetryCount(Long retryCount) {
    this.retryCount = retryCount;
  }


  public TaskPowershellConfig retryDelaySeconds(Long retryDelaySeconds) {
    
    this.retryDelaySeconds = retryDelaySeconds;
    return this;
  }

   /**
   * Get retryDelaySeconds
   * @return retryDelaySeconds
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public Long getRetryDelaySeconds() {
    return retryDelaySeconds;
  }


  public void setRetryDelaySeconds(Long retryDelaySeconds) {
    this.retryDelaySeconds = retryDelaySeconds;
  }


  public TaskPowershellConfig allowCustomConfig(Boolean allowCustomConfig) {
    
    this.allowCustomConfig = allowCustomConfig;
    return this;
  }

   /**
   * Get allowCustomConfig
   * @return allowCustomConfig
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public Boolean getAllowCustomConfig() {
    return allowCustomConfig;
  }


  public void setAllowCustomConfig(Boolean allowCustomConfig) {
    this.allowCustomConfig = allowCustomConfig;
  }


  public TaskPowershellConfig credential(OptionTypeListCredential credential) {
    
    this.credential = credential;
    return this;
  }

   /**
   * Get credential
   * @return credential
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public OptionTypeListCredential getCredential() {
    return credential;
  }


  public void setCredential(OptionTypeListCredential credential) {
    this.credential = credential;
  }


  public TaskPowershellConfig dateCreated(OffsetDateTime dateCreated) {
    
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


  public TaskPowershellConfig lastUpdated(OffsetDateTime lastUpdated) {
    
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


  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (o == null || getClass() != o.getClass()) {
      return false;
    }
    TaskPowershellConfig taskPowershellConfig = (TaskPowershellConfig) o;
    return Objects.equals(this.id, taskPowershellConfig.id) &&
        Objects.equals(this.accountId, taskPowershellConfig.accountId) &&
        Objects.equals(this.name, taskPowershellConfig.name) &&
        Objects.equals(this.code, taskPowershellConfig.code) &&
        Objects.equals(this.taskType, taskPowershellConfig.taskType) &&
        Objects.equals(this.labels, taskPowershellConfig.labels) &&
        Objects.equals(this.visibility, taskPowershellConfig.visibility) &&
        Objects.equals(this.taskOptions, taskPowershellConfig.taskOptions) &&
        Objects.equals(this.file, taskPowershellConfig.file) &&
        Objects.equals(this.resultType, taskPowershellConfig.resultType) &&
        Objects.equals(this.executeTarget, taskPowershellConfig.executeTarget) &&
        Objects.equals(this.retryable, taskPowershellConfig.retryable) &&
        Objects.equals(this.retryCount, taskPowershellConfig.retryCount) &&
        Objects.equals(this.retryDelaySeconds, taskPowershellConfig.retryDelaySeconds) &&
        Objects.equals(this.allowCustomConfig, taskPowershellConfig.allowCustomConfig) &&
        Objects.equals(this.credential, taskPowershellConfig.credential) &&
        Objects.equals(this.dateCreated, taskPowershellConfig.dateCreated) &&
        Objects.equals(this.lastUpdated, taskPowershellConfig.lastUpdated);
  }

  @Override
  public int hashCode() {
    return Objects.hash(id, accountId, name, code, taskType, labels, visibility, taskOptions, file, resultType, executeTarget, retryable, retryCount, retryDelaySeconds, allowCustomConfig, credential, dateCreated, lastUpdated);
  }


  @Override
  public String toString() {
    StringBuilder sb = new StringBuilder();
    sb.append("class TaskPowershellConfig {\n");
    sb.append("    id: ").append(toIndentedString(id)).append("\n");
    sb.append("    accountId: ").append(toIndentedString(accountId)).append("\n");
    sb.append("    name: ").append(toIndentedString(name)).append("\n");
    sb.append("    code: ").append(toIndentedString(code)).append("\n");
    sb.append("    taskType: ").append(toIndentedString(taskType)).append("\n");
    sb.append("    labels: ").append(toIndentedString(labels)).append("\n");
    sb.append("    visibility: ").append(toIndentedString(visibility)).append("\n");
    sb.append("    taskOptions: ").append(toIndentedString(taskOptions)).append("\n");
    sb.append("    file: ").append(toIndentedString(file)).append("\n");
    sb.append("    resultType: ").append(toIndentedString(resultType)).append("\n");
    sb.append("    executeTarget: ").append(toIndentedString(executeTarget)).append("\n");
    sb.append("    retryable: ").append(toIndentedString(retryable)).append("\n");
    sb.append("    retryCount: ").append(toIndentedString(retryCount)).append("\n");
    sb.append("    retryDelaySeconds: ").append(toIndentedString(retryDelaySeconds)).append("\n");
    sb.append("    allowCustomConfig: ").append(toIndentedString(allowCustomConfig)).append("\n");
    sb.append("    credential: ").append(toIndentedString(credential)).append("\n");
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

