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
 * InstanceScheduleThreshold
 */
@javax.annotation.Generated(value = "org.openapitools.codegen.languages.JavaClientCodegen", date = "2023-08-21T11:26:47.028Z[GMT]")
public class InstanceScheduleThreshold {
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

  public static final String SERIALIZED_NAME_SCALE_INCREMENT = "scaleIncrement";
  @SerializedName(SERIALIZED_NAME_SCALE_INCREMENT)
  private Long scaleIncrement = 1l;

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


  public InstanceScheduleThreshold autoUp(Boolean autoUp) {
    
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


  public InstanceScheduleThreshold autoDown(Boolean autoDown) {
    
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


  public InstanceScheduleThreshold minCount(Integer minCount) {
    
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


  public InstanceScheduleThreshold maxCount(Integer maxCount) {
    
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


  public InstanceScheduleThreshold scaleIncrement(Long scaleIncrement) {
    
    this.scaleIncrement = scaleIncrement;
    return this;
  }

   /**
   * The number of nodes that are added or removed at one time when scaling up or down
   * @return scaleIncrement
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "The number of nodes that are added or removed at one time when scaling up or down")

  public Long getScaleIncrement() {
    return scaleIncrement;
  }


  public void setScaleIncrement(Long scaleIncrement) {
    this.scaleIncrement = scaleIncrement;
  }


  public InstanceScheduleThreshold cpuEnabled(Boolean cpuEnabled) {
    
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


  public InstanceScheduleThreshold minCpu(Double minCpu) {
    
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


  public InstanceScheduleThreshold maxCpu(Double maxCpu) {
    
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


  public InstanceScheduleThreshold memoryEnabled(Boolean memoryEnabled) {
    
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


  public InstanceScheduleThreshold minMemory(Double minMemory) {
    
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


  public InstanceScheduleThreshold maxMemory(Double maxMemory) {
    
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


  public InstanceScheduleThreshold diskEnabled(Boolean diskEnabled) {
    
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


  public InstanceScheduleThreshold minDisk(Double minDisk) {
    
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


  public InstanceScheduleThreshold maxDisk(Double maxDisk) {
    
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
    InstanceScheduleThreshold instanceScheduleThreshold = (InstanceScheduleThreshold) o;
    return Objects.equals(this.autoUp, instanceScheduleThreshold.autoUp) &&
        Objects.equals(this.autoDown, instanceScheduleThreshold.autoDown) &&
        Objects.equals(this.minCount, instanceScheduleThreshold.minCount) &&
        Objects.equals(this.maxCount, instanceScheduleThreshold.maxCount) &&
        Objects.equals(this.scaleIncrement, instanceScheduleThreshold.scaleIncrement) &&
        Objects.equals(this.cpuEnabled, instanceScheduleThreshold.cpuEnabled) &&
        Objects.equals(this.minCpu, instanceScheduleThreshold.minCpu) &&
        Objects.equals(this.maxCpu, instanceScheduleThreshold.maxCpu) &&
        Objects.equals(this.memoryEnabled, instanceScheduleThreshold.memoryEnabled) &&
        Objects.equals(this.minMemory, instanceScheduleThreshold.minMemory) &&
        Objects.equals(this.maxMemory, instanceScheduleThreshold.maxMemory) &&
        Objects.equals(this.diskEnabled, instanceScheduleThreshold.diskEnabled) &&
        Objects.equals(this.minDisk, instanceScheduleThreshold.minDisk) &&
        Objects.equals(this.maxDisk, instanceScheduleThreshold.maxDisk);
  }

  @Override
  public int hashCode() {
    return Objects.hash(autoUp, autoDown, minCount, maxCount, scaleIncrement, cpuEnabled, minCpu, maxCpu, memoryEnabled, minMemory, maxMemory, diskEnabled, minDisk, maxDisk);
  }


  @Override
  public String toString() {
    StringBuilder sb = new StringBuilder();
    sb.append("class InstanceScheduleThreshold {\n");
    sb.append("    autoUp: ").append(toIndentedString(autoUp)).append("\n");
    sb.append("    autoDown: ").append(toIndentedString(autoDown)).append("\n");
    sb.append("    minCount: ").append(toIndentedString(minCount)).append("\n");
    sb.append("    maxCount: ").append(toIndentedString(maxCount)).append("\n");
    sb.append("    scaleIncrement: ").append(toIndentedString(scaleIncrement)).append("\n");
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

