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
import java.util.ArrayList;
import java.util.List;

/**
 * AppStats
 */
@javax.annotation.Generated(value = "org.openapitools.codegen.languages.JavaClientCodegen", date = "2023-08-21T11:26:47.028Z[GMT]")
public class AppStats {
  public static final String SERIALIZED_NAME_USED_MEMORY = "usedMemory";
  @SerializedName(SERIALIZED_NAME_USED_MEMORY)
  private Long usedMemory;

  public static final String SERIALIZED_NAME_MAX_MEMORY = "maxMemory";
  @SerializedName(SERIALIZED_NAME_MAX_MEMORY)
  private Long maxMemory;

  public static final String SERIALIZED_NAME_USED_STORAGE = "usedStorage";
  @SerializedName(SERIALIZED_NAME_USED_STORAGE)
  private Long usedStorage;

  public static final String SERIALIZED_NAME_MAX_STORAGE = "maxStorage";
  @SerializedName(SERIALIZED_NAME_MAX_STORAGE)
  private Long maxStorage;

  public static final String SERIALIZED_NAME_RUNNING = "running";
  @SerializedName(SERIALIZED_NAME_RUNNING)
  private Long running;

  public static final String SERIALIZED_NAME_TOTAL = "total";
  @SerializedName(SERIALIZED_NAME_TOTAL)
  private Long total;

  public static final String SERIALIZED_NAME_CPU_USAGE = "cpuUsage";
  @SerializedName(SERIALIZED_NAME_CPU_USAGE)
  private BigDecimal cpuUsage;

  public static final String SERIALIZED_NAME_INSTANCE_COUNT = "instanceCount";
  @SerializedName(SERIALIZED_NAME_INSTANCE_COUNT)
  private Long instanceCount;

  public static final String SERIALIZED_NAME_INSTANCE_DAY_COUNT = "instanceDayCount";
  @SerializedName(SERIALIZED_NAME_INSTANCE_DAY_COUNT)
  private List<Long> instanceDayCount = null;

  public static final String SERIALIZED_NAME_INSTANCE_DAY_COUNT_TOTAL = "instanceDayCountTotal";
  @SerializedName(SERIALIZED_NAME_INSTANCE_DAY_COUNT_TOTAL)
  private Long instanceDayCountTotal;


  public AppStats usedMemory(Long usedMemory) {
    
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


  public AppStats maxMemory(Long maxMemory) {
    
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


  public AppStats usedStorage(Long usedStorage) {
    
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


  public AppStats maxStorage(Long maxStorage) {
    
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


  public AppStats running(Long running) {
    
    this.running = running;
    return this;
  }

   /**
   * Get running
   * @return running
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public Long getRunning() {
    return running;
  }


  public void setRunning(Long running) {
    this.running = running;
  }


  public AppStats total(Long total) {
    
    this.total = total;
    return this;
  }

   /**
   * Get total
   * @return total
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public Long getTotal() {
    return total;
  }


  public void setTotal(Long total) {
    this.total = total;
  }


  public AppStats cpuUsage(BigDecimal cpuUsage) {
    
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


  public AppStats instanceCount(Long instanceCount) {
    
    this.instanceCount = instanceCount;
    return this;
  }

   /**
   * Get instanceCount
   * @return instanceCount
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public Long getInstanceCount() {
    return instanceCount;
  }


  public void setInstanceCount(Long instanceCount) {
    this.instanceCount = instanceCount;
  }


  public AppStats instanceDayCount(List<Long> instanceDayCount) {
    
    this.instanceDayCount = instanceDayCount;
    return this;
  }

  public AppStats addInstanceDayCountItem(Long instanceDayCountItem) {
    if (this.instanceDayCount == null) {
      this.instanceDayCount = new ArrayList<Long>();
    }
    this.instanceDayCount.add(instanceDayCountItem);
    return this;
  }

   /**
   * Get instanceDayCount
   * @return instanceDayCount
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public List<Long> getInstanceDayCount() {
    return instanceDayCount;
  }


  public void setInstanceDayCount(List<Long> instanceDayCount) {
    this.instanceDayCount = instanceDayCount;
  }


  public AppStats instanceDayCountTotal(Long instanceDayCountTotal) {
    
    this.instanceDayCountTotal = instanceDayCountTotal;
    return this;
  }

   /**
   * Get instanceDayCountTotal
   * @return instanceDayCountTotal
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public Long getInstanceDayCountTotal() {
    return instanceDayCountTotal;
  }


  public void setInstanceDayCountTotal(Long instanceDayCountTotal) {
    this.instanceDayCountTotal = instanceDayCountTotal;
  }


  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (o == null || getClass() != o.getClass()) {
      return false;
    }
    AppStats appStats = (AppStats) o;
    return Objects.equals(this.usedMemory, appStats.usedMemory) &&
        Objects.equals(this.maxMemory, appStats.maxMemory) &&
        Objects.equals(this.usedStorage, appStats.usedStorage) &&
        Objects.equals(this.maxStorage, appStats.maxStorage) &&
        Objects.equals(this.running, appStats.running) &&
        Objects.equals(this.total, appStats.total) &&
        Objects.equals(this.cpuUsage, appStats.cpuUsage) &&
        Objects.equals(this.instanceCount, appStats.instanceCount) &&
        Objects.equals(this.instanceDayCount, appStats.instanceDayCount) &&
        Objects.equals(this.instanceDayCountTotal, appStats.instanceDayCountTotal);
  }

  @Override
  public int hashCode() {
    return Objects.hash(usedMemory, maxMemory, usedStorage, maxStorage, running, total, cpuUsage, instanceCount, instanceDayCount, instanceDayCountTotal);
  }


  @Override
  public String toString() {
    StringBuilder sb = new StringBuilder();
    sb.append("class AppStats {\n");
    sb.append("    usedMemory: ").append(toIndentedString(usedMemory)).append("\n");
    sb.append("    maxMemory: ").append(toIndentedString(maxMemory)).append("\n");
    sb.append("    usedStorage: ").append(toIndentedString(usedStorage)).append("\n");
    sb.append("    maxStorage: ").append(toIndentedString(maxStorage)).append("\n");
    sb.append("    running: ").append(toIndentedString(running)).append("\n");
    sb.append("    total: ").append(toIndentedString(total)).append("\n");
    sb.append("    cpuUsage: ").append(toIndentedString(cpuUsage)).append("\n");
    sb.append("    instanceCount: ").append(toIndentedString(instanceCount)).append("\n");
    sb.append("    instanceDayCount: ").append(toIndentedString(instanceDayCount)).append("\n");
    sb.append("    instanceDayCountTotal: ").append(toIndentedString(instanceDayCountTotal)).append("\n");
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

