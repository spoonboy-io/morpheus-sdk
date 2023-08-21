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
import org.openapitools.client.model.ApiJobsIdJobTargets;
import org.openapitools.client.model.ApiJobsIdJobTask;
import org.openapitools.client.model.OneOfstringlong;
import org.threeten.bp.OffsetDateTime;

/**
 * ApiJobsIdJob
 */
@javax.annotation.Generated(value = "org.openapitools.codegen.languages.JavaClientCodegen", date = "2023-08-21T11:26:47.028Z[GMT]")
public class ApiJobsIdJob {
  public static final String SERIALIZED_NAME_NAME = "name";
  @SerializedName(SERIALIZED_NAME_NAME)
  private String name;

  public static final String SERIALIZED_NAME_LABELS = "labels";
  @SerializedName(SERIALIZED_NAME_LABELS)
  private List<String> labels = null;

  public static final String SERIALIZED_NAME_ENABLED = "enabled";
  @SerializedName(SERIALIZED_NAME_ENABLED)
  private Boolean enabled = true;

  public static final String SERIALIZED_NAME_TASK = "task";
  @SerializedName(SERIALIZED_NAME_TASK)
  private ApiJobsIdJobTask task;

  public static final String SERIALIZED_NAME_WORKFLOW = "workflow";
  @SerializedName(SERIALIZED_NAME_WORKFLOW)
  private ApiJobsIdJobTask workflow;

  public static final String SERIALIZED_NAME_SCAN_PATH = "scanPath";
  @SerializedName(SERIALIZED_NAME_SCAN_PATH)
  private String scanPath;

  public static final String SERIALIZED_NAME_SECURITY_PROFILE = "securityProfile";
  @SerializedName(SERIALIZED_NAME_SECURITY_PROFILE)
  private String securityProfile;

  /**
   * Target type where job will execute
   */
  @JsonAdapter(TargetTypeEnum.Adapter.class)
  public enum TargetTypeEnum {
    APPLIANCE("appliance"),
    
    INSTANCE("instance"),
    
    INSTANCE_LABEL("instance-label"),
    
    SERVER("server"),
    
    SERVER_LABEL("server-label");

    private String value;

    TargetTypeEnum(String value) {
      this.value = value;
    }

    public String getValue() {
      return value;
    }

    @Override
    public String toString() {
      return String.valueOf(value);
    }

    public static TargetTypeEnum fromValue(String value) {
      for (TargetTypeEnum b : TargetTypeEnum.values()) {
        if (b.value.equals(value)) {
          return b;
        }
      }
      throw new IllegalArgumentException("Unexpected value '" + value + "'");
    }

    public static class Adapter extends TypeAdapter<TargetTypeEnum> {
      @Override
      public void write(final JsonWriter jsonWriter, final TargetTypeEnum enumeration) throws IOException {
        jsonWriter.value(enumeration.getValue());
      }

      @Override
      public TargetTypeEnum read(final JsonReader jsonReader) throws IOException {
        String value =  jsonReader.nextString();
        return TargetTypeEnum.fromValue(value);
      }
    }
  }

  public static final String SERIALIZED_NAME_TARGET_TYPE = "targetType";
  @SerializedName(SERIALIZED_NAME_TARGET_TYPE)
  private TargetTypeEnum targetType;

  public static final String SERIALIZED_NAME_TARGETS = "targets";
  @SerializedName(SERIALIZED_NAME_TARGETS)
  private List<ApiJobsIdJobTargets> targets = null;

  public static final String SERIALIZED_NAME_INSTANCE_LABEL = "instanceLabel";
  @SerializedName(SERIALIZED_NAME_INSTANCE_LABEL)
  private String instanceLabel;

  public static final String SERIALIZED_NAME_SERVER_LABEL = "serverLabel";
  @SerializedName(SERIALIZED_NAME_SERVER_LABEL)
  private String serverLabel;

  public static final String SERIALIZED_NAME_SCHEDULE_MODE = "scheduleMode";
  @SerializedName(SERIALIZED_NAME_SCHEDULE_MODE)
  private OneOfstringlong scheduleMode = null;

  public static final String SERIALIZED_NAME_CUSTOM_OPTIONS = "customOptions";
  @SerializedName(SERIALIZED_NAME_CUSTOM_OPTIONS)
  private Object customOptions;

  public static final String SERIALIZED_NAME_CUSTOM_CONFIG = "customConfig";
  @SerializedName(SERIALIZED_NAME_CUSTOM_CONFIG)
  private String customConfig;

  public static final String SERIALIZED_NAME_DATE_TIME = "dateTime";
  @SerializedName(SERIALIZED_NAME_DATE_TIME)
  private OffsetDateTime dateTime;

  public static final String SERIALIZED_NAME_RUN = "run";
  @SerializedName(SERIALIZED_NAME_RUN)
  private Boolean run;


  public ApiJobsIdJob name(String name) {
    
    this.name = name;
    return this;
  }

   /**
   * A name for the Job
   * @return name
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(example = "Sample Job Name", value = "A name for the Job")

  public String getName() {
    return name;
  }


  public void setName(String name) {
    this.name = name;
  }


  public ApiJobsIdJob labels(List<String> labels) {
    
    this.labels = labels;
    return this;
  }

  public ApiJobsIdJob addLabelsItem(String labelsItem) {
    if (this.labels == null) {
      this.labels = new ArrayList<String>();
    }
    this.labels.add(labelsItem);
    return this;
  }

   /**
   * Array of label strings, can be used for filtering.
   * @return labels
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "Array of label strings, can be used for filtering.")

  public List<String> getLabels() {
    return labels;
  }


  public void setLabels(List<String> labels) {
    this.labels = labels;
  }


  public ApiJobsIdJob enabled(Boolean enabled) {
    
    this.enabled = enabled;
    return this;
  }

   /**
   * Use this to set enabled state
   * @return enabled
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "Use this to set enabled state")

  public Boolean getEnabled() {
    return enabled;
  }


  public void setEnabled(Boolean enabled) {
    this.enabled = enabled;
  }


  public ApiJobsIdJob task(ApiJobsIdJobTask task) {
    
    this.task = task;
    return this;
  }

   /**
   * Get task
   * @return task
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public ApiJobsIdJobTask getTask() {
    return task;
  }


  public void setTask(ApiJobsIdJobTask task) {
    this.task = task;
  }


  public ApiJobsIdJob workflow(ApiJobsIdJobTask workflow) {
    
    this.workflow = workflow;
    return this;
  }

   /**
   * Get workflow
   * @return workflow
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public ApiJobsIdJobTask getWorkflow() {
    return workflow;
  }


  public void setWorkflow(ApiJobsIdJobTask workflow) {
    this.workflow = workflow;
  }


  public ApiJobsIdJob scanPath(String scanPath) {
    
    this.scanPath = scanPath;
    return this;
  }

   /**
   * Scan Checklist. Only applies to type scap-package.
   * @return scanPath
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(example = "/test_CentOS_Linux_7_Benchmark_v3/test_CentOS_Linux_7_Benchmark_v3.1.1-xccdf.xml", value = "Scan Checklist. Only applies to type scap-package.")

  public String getScanPath() {
    return scanPath;
  }


  public void setScanPath(String scanPath) {
    this.scanPath = scanPath;
  }


  public ApiJobsIdJob securityProfile(String securityProfile) {
    
    this.securityProfile = securityProfile;
    return this;
  }

   /**
   * Security Profile. Only applies to type scap-package.
   * @return securityProfile
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(example = "xccdf_org.cisecurity.benchmarks_profile_Level_2_-_Server", value = "Security Profile. Only applies to type scap-package.")

  public String getSecurityProfile() {
    return securityProfile;
  }


  public void setSecurityProfile(String securityProfile) {
    this.securityProfile = securityProfile;
  }


  public ApiJobsIdJob targetType(TargetTypeEnum targetType) {
    
    this.targetType = targetType;
    return this;
  }

   /**
   * Target type where job will execute
   * @return targetType
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "Target type where job will execute")

  public TargetTypeEnum getTargetType() {
    return targetType;
  }


  public void setTargetType(TargetTypeEnum targetType) {
    this.targetType = targetType;
  }


  public ApiJobsIdJob targets(List<ApiJobsIdJobTargets> targets) {
    
    this.targets = targets;
    return this;
  }

  public ApiJobsIdJob addTargetsItem(ApiJobsIdJobTargets targetsItem) {
    if (this.targets == null) {
      this.targets = new ArrayList<ApiJobsIdJobTargets>();
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

  public List<ApiJobsIdJobTargets> getTargets() {
    return targets;
  }


  public void setTargets(List<ApiJobsIdJobTargets> targets) {
    this.targets = targets;
  }


  public ApiJobsIdJob instanceLabel(String instanceLabel) {
    
    this.instanceLabel = instanceLabel;
    return this;
  }

   /**
   * Instance Label. Only applicable if &#x60;targetType&#x60; is &#x60;instance-label&#x60;.
   * @return instanceLabel
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "Instance Label. Only applicable if `targetType` is `instance-label`.")

  public String getInstanceLabel() {
    return instanceLabel;
  }


  public void setInstanceLabel(String instanceLabel) {
    this.instanceLabel = instanceLabel;
  }


  public ApiJobsIdJob serverLabel(String serverLabel) {
    
    this.serverLabel = serverLabel;
    return this;
  }

   /**
   * Server Label. Only applicable if &#x60;targetType&#x60; is &#x60;server-label&#x60;.
   * @return serverLabel
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "Server Label. Only applicable if `targetType` is `server-label`.")

  public String getServerLabel() {
    return serverLabel;
  }


  public void setServerLabel(String serverLabel) {
    this.serverLabel = serverLabel;
  }


  public ApiJobsIdJob scheduleMode(OneOfstringlong scheduleMode) {
    
    this.scheduleMode = scheduleMode;
    return this;
  }

   /**
   * Get scheduleMode
   * @return scheduleMode
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public OneOfstringlong getScheduleMode() {
    return scheduleMode;
  }


  public void setScheduleMode(OneOfstringlong scheduleMode) {
    this.scheduleMode = scheduleMode;
  }


  public ApiJobsIdJob customOptions(Object customOptions) {
    
    this.customOptions = customOptions;
    return this;
  }

   /**
   * Map of options to be used as values in the workflow tasks. These correspond to option types.
   * @return customOptions
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "Map of options to be used as values in the workflow tasks. These correspond to option types.")

  public Object getCustomOptions() {
    return customOptions;
  }


  public void setCustomOptions(Object customOptions) {
    this.customOptions = customOptions;
  }


  public ApiJobsIdJob customConfig(String customConfig) {
    
    this.customConfig = customConfig;
    return this;
  }

   /**
   * Job custom configuration (String in JSON format)
   * @return customConfig
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "Job custom configuration (String in JSON format)")

  public String getCustomConfig() {
    return customConfig;
  }


  public void setCustomConfig(String customConfig) {
    this.customConfig = customConfig;
  }


  public ApiJobsIdJob dateTime(OffsetDateTime dateTime) {
    
    this.dateTime = dateTime;
    return this;
  }

   /**
   * Date and Time to execute the job. Use UTC time in the format 2020-02-15T05:00:00Z. Required when scheduleMode is &#39;dateTime&#39;.
   * @return dateTime
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(example = "2020-02-15T05:00Z", value = "Date and Time to execute the job. Use UTC time in the format 2020-02-15T05:00:00Z. Required when scheduleMode is 'dateTime'.")

  public OffsetDateTime getDateTime() {
    return dateTime;
  }


  public void setDateTime(OffsetDateTime dateTime) {
    this.dateTime = dateTime;
  }


  public ApiJobsIdJob run(Boolean run) {
    
    this.run = run;
    return this;
  }

   /**
   * If true, executes job
   * @return run
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "If true, executes job")

  public Boolean getRun() {
    return run;
  }


  public void setRun(Boolean run) {
    this.run = run;
  }


  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (o == null || getClass() != o.getClass()) {
      return false;
    }
    ApiJobsIdJob apiJobsIdJob = (ApiJobsIdJob) o;
    return Objects.equals(this.name, apiJobsIdJob.name) &&
        Objects.equals(this.labels, apiJobsIdJob.labels) &&
        Objects.equals(this.enabled, apiJobsIdJob.enabled) &&
        Objects.equals(this.task, apiJobsIdJob.task) &&
        Objects.equals(this.workflow, apiJobsIdJob.workflow) &&
        Objects.equals(this.scanPath, apiJobsIdJob.scanPath) &&
        Objects.equals(this.securityProfile, apiJobsIdJob.securityProfile) &&
        Objects.equals(this.targetType, apiJobsIdJob.targetType) &&
        Objects.equals(this.targets, apiJobsIdJob.targets) &&
        Objects.equals(this.instanceLabel, apiJobsIdJob.instanceLabel) &&
        Objects.equals(this.serverLabel, apiJobsIdJob.serverLabel) &&
        Objects.equals(this.scheduleMode, apiJobsIdJob.scheduleMode) &&
        Objects.equals(this.customOptions, apiJobsIdJob.customOptions) &&
        Objects.equals(this.customConfig, apiJobsIdJob.customConfig) &&
        Objects.equals(this.dateTime, apiJobsIdJob.dateTime) &&
        Objects.equals(this.run, apiJobsIdJob.run);
  }

  @Override
  public int hashCode() {
    return Objects.hash(name, labels, enabled, task, workflow, scanPath, securityProfile, targetType, targets, instanceLabel, serverLabel, scheduleMode, customOptions, customConfig, dateTime, run);
  }


  @Override
  public String toString() {
    StringBuilder sb = new StringBuilder();
    sb.append("class ApiJobsIdJob {\n");
    sb.append("    name: ").append(toIndentedString(name)).append("\n");
    sb.append("    labels: ").append(toIndentedString(labels)).append("\n");
    sb.append("    enabled: ").append(toIndentedString(enabled)).append("\n");
    sb.append("    task: ").append(toIndentedString(task)).append("\n");
    sb.append("    workflow: ").append(toIndentedString(workflow)).append("\n");
    sb.append("    scanPath: ").append(toIndentedString(scanPath)).append("\n");
    sb.append("    securityProfile: ").append(toIndentedString(securityProfile)).append("\n");
    sb.append("    targetType: ").append(toIndentedString(targetType)).append("\n");
    sb.append("    targets: ").append(toIndentedString(targets)).append("\n");
    sb.append("    instanceLabel: ").append(toIndentedString(instanceLabel)).append("\n");
    sb.append("    serverLabel: ").append(toIndentedString(serverLabel)).append("\n");
    sb.append("    scheduleMode: ").append(toIndentedString(scheduleMode)).append("\n");
    sb.append("    customOptions: ").append(toIndentedString(customOptions)).append("\n");
    sb.append("    customConfig: ").append(toIndentedString(customConfig)).append("\n");
    sb.append("    dateTime: ").append(toIndentedString(dateTime)).append("\n");
    sb.append("    run: ").append(toIndentedString(run)).append("\n");
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

