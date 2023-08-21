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
import java.math.BigDecimal;

/**
 * InstanceStats
 */
@javax.annotation.Generated(value = "org.openapitools.codegen.languages.JavaClientCodegen", date = "2023-08-21T11:26:47.028Z[GMT]")
public class InstanceStats {
  public static final String SERIALIZED_NAME_USED_STORAGE = "usedStorage";
  @SerializedName(SERIALIZED_NAME_USED_STORAGE)
  private Long usedStorage;

  public static final String SERIALIZED_NAME_MAX_STORAGE = "maxStorage";
  @SerializedName(SERIALIZED_NAME_MAX_STORAGE)
  private Long maxStorage;

  public static final String SERIALIZED_NAME_USED_MEMORY = "usedMemory";
  @SerializedName(SERIALIZED_NAME_USED_MEMORY)
  private Long usedMemory;

  public static final String SERIALIZED_NAME_MAX_MEMORY = "maxMemory";
  @SerializedName(SERIALIZED_NAME_MAX_MEMORY)
  private Long maxMemory;

  public static final String SERIALIZED_NAME_USED_CPU = "usedCpu";
  @SerializedName(SERIALIZED_NAME_USED_CPU)
  private BigDecimal usedCpu;

  public static final String SERIALIZED_NAME_CPU_USAGE = "cpuUsage";
  @SerializedName(SERIALIZED_NAME_CPU_USAGE)
  private BigDecimal cpuUsage;

  public static final String SERIALIZED_NAME_CPU_USAGE_PEAK = "cpuUsagePeak";
  @SerializedName(SERIALIZED_NAME_CPU_USAGE_PEAK)
  private BigDecimal cpuUsagePeak;

  public static final String SERIALIZED_NAME_CPU_USAGE_AVG = "cpuUsageAvg";
  @SerializedName(SERIALIZED_NAME_CPU_USAGE_AVG)
  private BigDecimal cpuUsageAvg;


  public InstanceStats usedStorage(Long usedStorage) {
    
    this.usedStorage = usedStorage;
    return this;
  }

   /**
   * Get usedStorage
   * @return usedStorage
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public Long getUsedStorage() {
    return usedStorage;
  }


  public void setUsedStorage(Long usedStorage) {
    this.usedStorage = usedStorage;
  }


  public InstanceStats maxStorage(Long maxStorage) {
    
    this.maxStorage = maxStorage;
    return this;
  }

   /**
   * Get maxStorage
   * @return maxStorage
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public Long getMaxStorage() {
    return maxStorage;
  }


  public void setMaxStorage(Long maxStorage) {
    this.maxStorage = maxStorage;
  }


  public InstanceStats usedMemory(Long usedMemory) {
    
    this.usedMemory = usedMemory;
    return this;
  }

   /**
   * Get usedMemory
   * @return usedMemory
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public Long getUsedMemory() {
    return usedMemory;
  }


  public void setUsedMemory(Long usedMemory) {
    this.usedMemory = usedMemory;
  }


  public InstanceStats maxMemory(Long maxMemory) {
    
    this.maxMemory = maxMemory;
    return this;
  }

   /**
   * Get maxMemory
   * @return maxMemory
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public Long getMaxMemory() {
    return maxMemory;
  }


  public void setMaxMemory(Long maxMemory) {
    this.maxMemory = maxMemory;
  }


  public InstanceStats usedCpu(BigDecimal usedCpu) {
    
    this.usedCpu = usedCpu;
    return this;
  }

   /**
   * Get usedCpu
   * @return usedCpu
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public BigDecimal getUsedCpu() {
    return usedCpu;
  }


  public void setUsedCpu(BigDecimal usedCpu) {
    this.usedCpu = usedCpu;
  }


  public InstanceStats cpuUsage(BigDecimal cpuUsage) {
    
    this.cpuUsage = cpuUsage;
    return this;
  }

   /**
   * Get cpuUsage
   * @return cpuUsage
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public BigDecimal getCpuUsage() {
    return cpuUsage;
  }


  public void setCpuUsage(BigDecimal cpuUsage) {
    this.cpuUsage = cpuUsage;
  }


  public InstanceStats cpuUsagePeak(BigDecimal cpuUsagePeak) {
    
    this.cpuUsagePeak = cpuUsagePeak;
    return this;
  }

   /**
   * Get cpuUsagePeak
   * @return cpuUsagePeak
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public BigDecimal getCpuUsagePeak() {
    return cpuUsagePeak;
  }


  public void setCpuUsagePeak(BigDecimal cpuUsagePeak) {
    this.cpuUsagePeak = cpuUsagePeak;
  }


  public InstanceStats cpuUsageAvg(BigDecimal cpuUsageAvg) {
    
    this.cpuUsageAvg = cpuUsageAvg;
    return this;
  }

   /**
   * Get cpuUsageAvg
   * @return cpuUsageAvg
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public BigDecimal getCpuUsageAvg() {
    return cpuUsageAvg;
  }


  public void setCpuUsageAvg(BigDecimal cpuUsageAvg) {
    this.cpuUsageAvg = cpuUsageAvg;
  }


  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (o == null || getClass() != o.getClass()) {
      return false;
    }
    InstanceStats instanceStats = (InstanceStats) o;
    return Objects.equals(this.usedStorage, instanceStats.usedStorage) &&
        Objects.equals(this.maxStorage, instanceStats.maxStorage) &&
        Objects.equals(this.usedMemory, instanceStats.usedMemory) &&
        Objects.equals(this.maxMemory, instanceStats.maxMemory) &&
        Objects.equals(this.usedCpu, instanceStats.usedCpu) &&
        Objects.equals(this.cpuUsage, instanceStats.cpuUsage) &&
        Objects.equals(this.cpuUsagePeak, instanceStats.cpuUsagePeak) &&
        Objects.equals(this.cpuUsageAvg, instanceStats.cpuUsageAvg);
  }

  @Override
  public int hashCode() {
    return Objects.hash(usedStorage, maxStorage, usedMemory, maxMemory, usedCpu, cpuUsage, cpuUsagePeak, cpuUsageAvg);
  }


  @Override
  public String toString() {
    StringBuilder sb = new StringBuilder();
    sb.append("class InstanceStats {\n");
    sb.append("    usedStorage: ").append(toIndentedString(usedStorage)).append("\n");
    sb.append("    maxStorage: ").append(toIndentedString(maxStorage)).append("\n");
    sb.append("    usedMemory: ").append(toIndentedString(usedMemory)).append("\n");
    sb.append("    maxMemory: ").append(toIndentedString(maxMemory)).append("\n");
    sb.append("    usedCpu: ").append(toIndentedString(usedCpu)).append("\n");
    sb.append("    cpuUsage: ").append(toIndentedString(cpuUsage)).append("\n");
    sb.append("    cpuUsagePeak: ").append(toIndentedString(cpuUsagePeak)).append("\n");
    sb.append("    cpuUsageAvg: ").append(toIndentedString(cpuUsageAvg)).append("\n");
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

