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

/**
 * InstanceScheduleCreateThreshold
 */
@javax.annotation.Generated(value = "org.openapitools.codegen.languages.JavaClientCodegen", date = "2023-08-21T11:26:47.028Z[GMT]")
public class InstanceScheduleCreateThreshold {
  public static final String SERIALIZED_NAME_SOURCE_THRESHOLD_ID = "sourceThresholdId";
  @SerializedName(SERIALIZED_NAME_SOURCE_THRESHOLD_ID)
  private Long sourceThresholdId;

  public static final String SERIALIZED_NAME_AUTO_UP = "autoUp";
  @SerializedName(SERIALIZED_NAME_AUTO_UP)
  private Boolean autoUp = false;

  public static final String SERIALIZED_NAME_AUTO_DOWN = "autoDown";
  @SerializedName(SERIALIZED_NAME_AUTO_DOWN)
  private Boolean autoDown = false;

  public static final String SERIALIZED_NAME_MIN_COUNT = "minCount";
  @SerializedName(SERIALIZED_NAME_MIN_COUNT)
  private Integer minCount;

  public static final String SERIALIZED_NAME_MAX_COUNT = "maxCount";
  @SerializedName(SERIALIZED_NAME_MAX_COUNT)
  private Integer maxCount;

  public static final String SERIALIZED_NAME_CPU_ENABLED = "cpuEnabled";
  @SerializedName(SERIALIZED_NAME_CPU_ENABLED)
  private Boolean cpuEnabled = false;

  public static final String SERIALIZED_NAME_MIN_CPU = "minCpu";
  @SerializedName(SERIALIZED_NAME_MIN_CPU)
  private Double minCpu = 0d;

  public static final String SERIALIZED_NAME_MAX_CPU = "maxCpu";
  @SerializedName(SERIALIZED_NAME_MAX_CPU)
  private Double maxCpu = 0d;

  public static final String SERIALIZED_NAME_MEMORY_ENABLED = "memoryEnabled";
  @SerializedName(SERIALIZED_NAME_MEMORY_ENABLED)
  private Boolean memoryEnabled = false;

  public static final String SERIALIZED_NAME_MIN_MEMORY = "minMemory";
  @SerializedName(SERIALIZED_NAME_MIN_MEMORY)
  private Double minMemory = 0d;

  public static final String SERIALIZED_NAME_MAX_MEMORY = "maxMemory";
  @SerializedName(SERIALIZED_NAME_MAX_MEMORY)
  private Double maxMemory = 0d;

  public static final String SERIALIZED_NAME_DISK_ENABLED = "diskEnabled";
  @SerializedName(SERIALIZED_NAME_DISK_ENABLED)
  private Boolean diskEnabled = false;

  public static final String SERIALIZED_NAME_MIN_DISK = "minDisk";
  @SerializedName(SERIALIZED_NAME_MIN_DISK)
  private Double minDisk = 0d;

  public static final String SERIALIZED_NAME_MAX_DISK = "maxDisk";
  @SerializedName(SERIALIZED_NAME_MAX_DISK)
  private Double maxDisk = 0d;


  public InstanceScheduleCreateThreshold sourceThresholdId(Long sourceThresholdId) {
    
    this.sourceThresholdId = sourceThresholdId;
    return this;
  }

   /**
   * Source Scale Threshold to apply as a template. All threshold settings with be copied from this threshold, and can be overridden by also passing each setting explicitly.
   * @return sourceThresholdId
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "Source Scale Threshold to apply as a template. All threshold settings with be copied from this threshold, and can be overridden by also passing each setting explicitly.")

  public Long getSourceThresholdId() {
    return sourceThresholdId;
  }


  public void setSourceThresholdId(Long sourceThresholdId) {
    this.sourceThresholdId = sourceThresholdId;
  }


  public InstanceScheduleCreateThreshold autoUp(Boolean autoUp) {
    
    this.autoUp = autoUp;
    return this;
  }

   /**
   * Auto Upscale
   * @return autoUp
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "Auto Upscale")

  public Boolean getAutoUp() {
    return autoUp;
  }


  public void setAutoUp(Boolean autoUp) {
    this.autoUp = autoUp;
  }


  public InstanceScheduleCreateThreshold autoDown(Boolean autoDown) {
    
    this.autoDown = autoDown;
    return this;
  }

   /**
   * Auto Downscale
   * @return autoDown
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "Auto Downscale")

  public Boolean getAutoDown() {
    return autoDown;
  }


  public void setAutoDown(Boolean autoDown) {
    this.autoDown = autoDown;
  }


  public InstanceScheduleCreateThreshold minCount(Integer minCount) {
    
    this.minCount = minCount;
    return this;
  }

   /**
   * The minimum number of nodes to scale down to
   * @return minCount
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(example = "1", value = "The minimum number of nodes to scale down to")

  public Integer getMinCount() {
    return minCount;
  }


  public void setMinCount(Integer minCount) {
    this.minCount = minCount;
  }


  public InstanceScheduleCreateThreshold maxCount(Integer maxCount) {
    
    this.maxCount = maxCount;
    return this;
  }

   /**
   * The maximum number of nodes to scale up to
   * @return maxCount
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(example = "3", value = "The maximum number of nodes to scale up to")

  public Integer getMaxCount() {
    return maxCount;
  }


  public void setMaxCount(Integer maxCount) {
    this.maxCount = maxCount;
  }


  public InstanceScheduleCreateThreshold cpuEnabled(Boolean cpuEnabled) {
    
    this.cpuEnabled = cpuEnabled;
    return this;
  }

   /**
   * Enable CPU Threshold
   * @return cpuEnabled
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "Enable CPU Threshold")

  public Boolean getCpuEnabled() {
    return cpuEnabled;
  }


  public void setCpuEnabled(Boolean cpuEnabled) {
    this.cpuEnabled = cpuEnabled;
  }


  public InstanceScheduleCreateThreshold minCpu(Double minCpu) {
    
    this.minCpu = minCpu;
    return this;
  }

   /**
   * Min CPU (%)
   * @return minCpu
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "Min CPU (%)")

  public Double getMinCpu() {
    return minCpu;
  }


  public void setMinCpu(Double minCpu) {
    this.minCpu = minCpu;
  }


  public InstanceScheduleCreateThreshold maxCpu(Double maxCpu) {
    
    this.maxCpu = maxCpu;
    return this;
  }

   /**
   * Max CPU (%)
   * @return maxCpu
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "Max CPU (%)")

  public Double getMaxCpu() {
    return maxCpu;
  }


  public void setMaxCpu(Double maxCpu) {
    this.maxCpu = maxCpu;
  }


  public InstanceScheduleCreateThreshold memoryEnabled(Boolean memoryEnabled) {
    
    this.memoryEnabled = memoryEnabled;
    return this;
  }

   /**
   * Enable Memory Threshold
   * @return memoryEnabled
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "Enable Memory Threshold")

  public Boolean getMemoryEnabled() {
    return memoryEnabled;
  }


  public void setMemoryEnabled(Boolean memoryEnabled) {
    this.memoryEnabled = memoryEnabled;
  }


  public InstanceScheduleCreateThreshold minMemory(Double minMemory) {
    
    this.minMemory = minMemory;
    return this;
  }

   /**
   * Min Memory (%)
   * @return minMemory
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "Min Memory (%)")

  public Double getMinMemory() {
    return minMemory;
  }


  public void setMinMemory(Double minMemory) {
    this.minMemory = minMemory;
  }


  public InstanceScheduleCreateThreshold maxMemory(Double maxMemory) {
    
    this.maxMemory = maxMemory;
    return this;
  }

   /**
   * Max Memory (%)
   * @return maxMemory
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "Max Memory (%)")

  public Double getMaxMemory() {
    return maxMemory;
  }


  public void setMaxMemory(Double maxMemory) {
    this.maxMemory = maxMemory;
  }


  public InstanceScheduleCreateThreshold diskEnabled(Boolean diskEnabled) {
    
    this.diskEnabled = diskEnabled;
    return this;
  }

   /**
   * Enable Disk Threshold
   * @return diskEnabled
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "Enable Disk Threshold")

  public Boolean getDiskEnabled() {
    return diskEnabled;
  }


  public void setDiskEnabled(Boolean diskEnabled) {
    this.diskEnabled = diskEnabled;
  }


  public InstanceScheduleCreateThreshold minDisk(Double minDisk) {
    
    this.minDisk = minDisk;
    return this;
  }

   /**
   * Min Disk (%)
   * @return minDisk
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "Min Disk (%)")

  public Double getMinDisk() {
    return minDisk;
  }


  public void setMinDisk(Double minDisk) {
    this.minDisk = minDisk;
  }


  public InstanceScheduleCreateThreshold maxDisk(Double maxDisk) {
    
    this.maxDisk = maxDisk;
    return this;
  }

   /**
   * Max Disk (%)
   * @return maxDisk
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "Max Disk (%)")

  public Double getMaxDisk() {
    return maxDisk;
  }


  public void setMaxDisk(Double maxDisk) {
    this.maxDisk = maxDisk;
  }


  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (o == null || getClass() != o.getClass()) {
      return false;
    }
    InstanceScheduleCreateThreshold instanceScheduleCreateThreshold = (InstanceScheduleCreateThreshold) o;
    return Objects.equals(this.sourceThresholdId, instanceScheduleCreateThreshold.sourceThresholdId) &&
        Objects.equals(this.autoUp, instanceScheduleCreateThreshold.autoUp) &&
        Objects.equals(this.autoDown, instanceScheduleCreateThreshold.autoDown) &&
        Objects.equals(this.minCount, instanceScheduleCreateThreshold.minCount) &&
        Objects.equals(this.maxCount, instanceScheduleCreateThreshold.maxCount) &&
        Objects.equals(this.cpuEnabled, instanceScheduleCreateThreshold.cpuEnabled) &&
        Objects.equals(this.minCpu, instanceScheduleCreateThreshold.minCpu) &&
        Objects.equals(this.maxCpu, instanceScheduleCreateThreshold.maxCpu) &&
        Objects.equals(this.memoryEnabled, instanceScheduleCreateThreshold.memoryEnabled) &&
        Objects.equals(this.minMemory, instanceScheduleCreateThreshold.minMemory) &&
        Objects.equals(this.maxMemory, instanceScheduleCreateThreshold.maxMemory) &&
        Objects.equals(this.diskEnabled, instanceScheduleCreateThreshold.diskEnabled) &&
        Objects.equals(this.minDisk, instanceScheduleCreateThreshold.minDisk) &&
        Objects.equals(this.maxDisk, instanceScheduleCreateThreshold.maxDisk);
  }

  @Override
  public int hashCode() {
    return Objects.hash(sourceThresholdId, autoUp, autoDown, minCount, maxCount, cpuEnabled, minCpu, maxCpu, memoryEnabled, minMemory, maxMemory, diskEnabled, minDisk, maxDisk);
  }


  @Override
  public String toString() {
    StringBuilder sb = new StringBuilder();
    sb.append("class InstanceScheduleCreateThreshold {\n");
    sb.append("    sourceThresholdId: ").append(toIndentedString(sourceThresholdId)).append("\n");
    sb.append("    autoUp: ").append(toIndentedString(autoUp)).append("\n");
    sb.append("    autoDown: ").append(toIndentedString(autoDown)).append("\n");
    sb.append("    minCount: ").append(toIndentedString(minCount)).append("\n");
    sb.append("    maxCount: ").append(toIndentedString(maxCount)).append("\n");
    sb.append("    cpuEnabled: ").append(toIndentedString(cpuEnabled)).append("\n");
    sb.append("    minCpu: ").append(toIndentedString(minCpu)).append("\n");
    sb.append("    maxCpu: ").append(toIndentedString(maxCpu)).append("\n");
    sb.append("    memoryEnabled: ").append(toIndentedString(memoryEnabled)).append("\n");
    sb.append("    minMemory: ").append(toIndentedString(minMemory)).append("\n");
    sb.append("    maxMemory: ").append(toIndentedString(maxMemory)).append("\n");
    sb.append("    diskEnabled: ").append(toIndentedString(diskEnabled)).append("\n");
    sb.append("    minDisk: ").append(toIndentedString(minDisk)).append("\n");
    sb.append("    maxDisk: ").append(toIndentedString(maxDisk)).append("\n");
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

