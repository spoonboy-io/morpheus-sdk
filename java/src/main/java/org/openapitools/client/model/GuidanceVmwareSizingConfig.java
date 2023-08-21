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
 * GuidanceVmwareSizingConfig
 */
@javax.annotation.Generated(value = "org.openapitools.codegen.languages.JavaClientCodegen", date = "2023-08-21T11:26:47.028Z[GMT]")
public class GuidanceVmwareSizingConfig {
  public static final String SERIALIZED_NAME_EXISTS = "exists";
  @SerializedName(SERIALIZED_NAME_EXISTS)
  private Boolean exists;

  public static final String SERIALIZED_NAME_OBJECT_ID = "objectId";
  @SerializedName(SERIALIZED_NAME_OBJECT_ID)
  private Long objectId;

  public static final String SERIALIZED_NAME_CPU_TOTAL_TIME_COUNT = "cpuTotalTimeCount";
  @SerializedName(SERIALIZED_NAME_CPU_TOTAL_TIME_COUNT)
  private BigDecimal cpuTotalTimeCount;

  public static final String SERIALIZED_NAME_CPU_TOTAL_TIME_MIN = "cpuTotalTimeMin";
  @SerializedName(SERIALIZED_NAME_CPU_TOTAL_TIME_MIN)
  private BigDecimal cpuTotalTimeMin;

  public static final String SERIALIZED_NAME_CPU_TOTAL_TIME_MAX = "cpuTotalTimeMax";
  @SerializedName(SERIALIZED_NAME_CPU_TOTAL_TIME_MAX)
  private BigDecimal cpuTotalTimeMax;

  public static final String SERIALIZED_NAME_CPU_TOTAL_TIME_AVG = "cpuTotalTimeAvg";
  @SerializedName(SERIALIZED_NAME_CPU_TOTAL_TIME_AVG)
  private BigDecimal cpuTotalTimeAvg;

  public static final String SERIALIZED_NAME_CPU_TOTAL_TIME_SUM = "cpuTotalTimeSum";
  @SerializedName(SERIALIZED_NAME_CPU_TOTAL_TIME_SUM)
  private BigDecimal cpuTotalTimeSum;

  public static final String SERIALIZED_NAME_CPU_IDLE_TIME_COUNT = "cpuIdleTimeCount";
  @SerializedName(SERIALIZED_NAME_CPU_IDLE_TIME_COUNT)
  private BigDecimal cpuIdleTimeCount;

  public static final String SERIALIZED_NAME_CPU_IDLE_TIME_MIN = "cpuIdleTimeMin";
  @SerializedName(SERIALIZED_NAME_CPU_IDLE_TIME_MIN)
  private BigDecimal cpuIdleTimeMin;

  public static final String SERIALIZED_NAME_CPU_IDLE_TIME_MAX = "cpuIdleTimeMax";
  @SerializedName(SERIALIZED_NAME_CPU_IDLE_TIME_MAX)
  private BigDecimal cpuIdleTimeMax;

  public static final String SERIALIZED_NAME_CPU_IDLE_TIME_AVG = "cpuIdleTimeAvg";
  @SerializedName(SERIALIZED_NAME_CPU_IDLE_TIME_AVG)
  private BigDecimal cpuIdleTimeAvg;

  public static final String SERIALIZED_NAME_CPU_IDLE_TIME_SUM = "cpuIdleTimeSum";
  @SerializedName(SERIALIZED_NAME_CPU_IDLE_TIME_SUM)
  private BigDecimal cpuIdleTimeSum;

  public static final String SERIALIZED_NAME_CPU_USAGE_COUNT = "cpuUsageCount";
  @SerializedName(SERIALIZED_NAME_CPU_USAGE_COUNT)
  private BigDecimal cpuUsageCount;

  public static final String SERIALIZED_NAME_CPU_USAGE_MIN = "cpuUsageMin";
  @SerializedName(SERIALIZED_NAME_CPU_USAGE_MIN)
  private BigDecimal cpuUsageMin;

  public static final String SERIALIZED_NAME_CPU_USAGE_MAX = "cpuUsageMax";
  @SerializedName(SERIALIZED_NAME_CPU_USAGE_MAX)
  private BigDecimal cpuUsageMax;

  public static final String SERIALIZED_NAME_CPU_USAGE_AVG = "cpuUsageAvg";
  @SerializedName(SERIALIZED_NAME_CPU_USAGE_AVG)
  private BigDecimal cpuUsageAvg;

  public static final String SERIALIZED_NAME_CPU_USAGE_SUM = "cpuUsageSum";
  @SerializedName(SERIALIZED_NAME_CPU_USAGE_SUM)
  private BigDecimal cpuUsageSum;

  public static final String SERIALIZED_NAME_MAX_MEMORY_COUNT = "maxMemoryCount";
  @SerializedName(SERIALIZED_NAME_MAX_MEMORY_COUNT)
  private BigDecimal maxMemoryCount;

  public static final String SERIALIZED_NAME_MAX_MEMORY_MIN = "maxMemoryMin";
  @SerializedName(SERIALIZED_NAME_MAX_MEMORY_MIN)
  private BigDecimal maxMemoryMin;

  public static final String SERIALIZED_NAME_MAX_MEMORY_MAX = "maxMemoryMax";
  @SerializedName(SERIALIZED_NAME_MAX_MEMORY_MAX)
  private BigDecimal maxMemoryMax;

  public static final String SERIALIZED_NAME_MAX_MEMORY_AVG = "maxMemoryAvg";
  @SerializedName(SERIALIZED_NAME_MAX_MEMORY_AVG)
  private BigDecimal maxMemoryAvg;

  public static final String SERIALIZED_NAME_MAX_MEMORY_SUM = "maxMemorySum";
  @SerializedName(SERIALIZED_NAME_MAX_MEMORY_SUM)
  private BigDecimal maxMemorySum;

  public static final String SERIALIZED_NAME_CPU_USER_TIME_COUNT = "cpuUserTimeCount";
  @SerializedName(SERIALIZED_NAME_CPU_USER_TIME_COUNT)
  private BigDecimal cpuUserTimeCount;

  public static final String SERIALIZED_NAME_CPU_USER_TIME_MIN = "cpuUserTimeMin";
  @SerializedName(SERIALIZED_NAME_CPU_USER_TIME_MIN)
  private BigDecimal cpuUserTimeMin;

  public static final String SERIALIZED_NAME_CPU_USER_TIME_MAX = "cpuUserTimeMax";
  @SerializedName(SERIALIZED_NAME_CPU_USER_TIME_MAX)
  private BigDecimal cpuUserTimeMax;

  public static final String SERIALIZED_NAME_CPU_USER_TIME_AVG = "cpuUserTimeAvg";
  @SerializedName(SERIALIZED_NAME_CPU_USER_TIME_AVG)
  private BigDecimal cpuUserTimeAvg;

  public static final String SERIALIZED_NAME_CPU_USER_TIME_SUM = "cpuUserTimeSum";
  @SerializedName(SERIALIZED_NAME_CPU_USER_TIME_SUM)
  private BigDecimal cpuUserTimeSum;

  public static final String SERIALIZED_NAME_CPU_SYSTEM_TIME_COUNT = "cpuSystemTimeCount";
  @SerializedName(SERIALIZED_NAME_CPU_SYSTEM_TIME_COUNT)
  private BigDecimal cpuSystemTimeCount;

  public static final String SERIALIZED_NAME_CPU_SYSTEM_TIME_MIN = "cpuSystemTimeMin";
  @SerializedName(SERIALIZED_NAME_CPU_SYSTEM_TIME_MIN)
  private BigDecimal cpuSystemTimeMin;

  public static final String SERIALIZED_NAME_CPU_SYSTEM_TIME_MAX = "cpuSystemTimeMax";
  @SerializedName(SERIALIZED_NAME_CPU_SYSTEM_TIME_MAX)
  private BigDecimal cpuSystemTimeMax;

  public static final String SERIALIZED_NAME_CPU_SYSTEM_TIME_AVG = "cpuSystemTimeAvg";
  @SerializedName(SERIALIZED_NAME_CPU_SYSTEM_TIME_AVG)
  private BigDecimal cpuSystemTimeAvg;

  public static final String SERIALIZED_NAME_CPU_SYSTEM_TIME_SUM = "cpuSystemTimeSum";
  @SerializedName(SERIALIZED_NAME_CPU_SYSTEM_TIME_SUM)
  private BigDecimal cpuSystemTimeSum;

  public static final String SERIALIZED_NAME_USED_MEMORY_COUNT = "usedMemoryCount";
  @SerializedName(SERIALIZED_NAME_USED_MEMORY_COUNT)
  private BigDecimal usedMemoryCount;

  public static final String SERIALIZED_NAME_USED_MEMORY_MIN = "usedMemoryMin";
  @SerializedName(SERIALIZED_NAME_USED_MEMORY_MIN)
  private BigDecimal usedMemoryMin;

  public static final String SERIALIZED_NAME_USED_MEMORY_MAX = "usedMemoryMax";
  @SerializedName(SERIALIZED_NAME_USED_MEMORY_MAX)
  private BigDecimal usedMemoryMax;

  public static final String SERIALIZED_NAME_USED_MEMORY_AVG = "usedMemoryAvg";
  @SerializedName(SERIALIZED_NAME_USED_MEMORY_AVG)
  private BigDecimal usedMemoryAvg;

  public static final String SERIALIZED_NAME_USED_MEMORY_SUM = "usedMemorySum";
  @SerializedName(SERIALIZED_NAME_USED_MEMORY_SUM)
  private BigDecimal usedMemorySum;

  public static final String SERIALIZED_NAME_FREE_MEMORY_COUNT = "freeMemoryCount";
  @SerializedName(SERIALIZED_NAME_FREE_MEMORY_COUNT)
  private BigDecimal freeMemoryCount;

  public static final String SERIALIZED_NAME_FREE_MEMORY_MIN = "freeMemoryMin";
  @SerializedName(SERIALIZED_NAME_FREE_MEMORY_MIN)
  private BigDecimal freeMemoryMin;

  public static final String SERIALIZED_NAME_FREE_MEMORY_MAX = "freeMemoryMax";
  @SerializedName(SERIALIZED_NAME_FREE_MEMORY_MAX)
  private BigDecimal freeMemoryMax;

  public static final String SERIALIZED_NAME_FREE_MEMORY_AVG = "freeMemoryAvg";
  @SerializedName(SERIALIZED_NAME_FREE_MEMORY_AVG)
  private BigDecimal freeMemoryAvg;

  public static final String SERIALIZED_NAME_FREE_MEMORY_SUM = "freeMemorySum";
  @SerializedName(SERIALIZED_NAME_FREE_MEMORY_SUM)
  private BigDecimal freeMemorySum;


  public GuidanceVmwareSizingConfig exists(Boolean exists) {
    
    this.exists = exists;
    return this;
  }

   /**
   * Get exists
   * @return exists
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public Boolean getExists() {
    return exists;
  }


  public void setExists(Boolean exists) {
    this.exists = exists;
  }


  public GuidanceVmwareSizingConfig objectId(Long objectId) {
    
    this.objectId = objectId;
    return this;
  }

   /**
   * Get objectId
   * @return objectId
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public Long getObjectId() {
    return objectId;
  }


  public void setObjectId(Long objectId) {
    this.objectId = objectId;
  }


  public GuidanceVmwareSizingConfig cpuTotalTimeCount(BigDecimal cpuTotalTimeCount) {
    
    this.cpuTotalTimeCount = cpuTotalTimeCount;
    return this;
  }

   /**
   * Get cpuTotalTimeCount
   * @return cpuTotalTimeCount
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public BigDecimal getCpuTotalTimeCount() {
    return cpuTotalTimeCount;
  }


  public void setCpuTotalTimeCount(BigDecimal cpuTotalTimeCount) {
    this.cpuTotalTimeCount = cpuTotalTimeCount;
  }


  public GuidanceVmwareSizingConfig cpuTotalTimeMin(BigDecimal cpuTotalTimeMin) {
    
    this.cpuTotalTimeMin = cpuTotalTimeMin;
    return this;
  }

   /**
   * Get cpuTotalTimeMin
   * @return cpuTotalTimeMin
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public BigDecimal getCpuTotalTimeMin() {
    return cpuTotalTimeMin;
  }


  public void setCpuTotalTimeMin(BigDecimal cpuTotalTimeMin) {
    this.cpuTotalTimeMin = cpuTotalTimeMin;
  }


  public GuidanceVmwareSizingConfig cpuTotalTimeMax(BigDecimal cpuTotalTimeMax) {
    
    this.cpuTotalTimeMax = cpuTotalTimeMax;
    return this;
  }

   /**
   * Get cpuTotalTimeMax
   * @return cpuTotalTimeMax
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public BigDecimal getCpuTotalTimeMax() {
    return cpuTotalTimeMax;
  }


  public void setCpuTotalTimeMax(BigDecimal cpuTotalTimeMax) {
    this.cpuTotalTimeMax = cpuTotalTimeMax;
  }


  public GuidanceVmwareSizingConfig cpuTotalTimeAvg(BigDecimal cpuTotalTimeAvg) {
    
    this.cpuTotalTimeAvg = cpuTotalTimeAvg;
    return this;
  }

   /**
   * Get cpuTotalTimeAvg
   * @return cpuTotalTimeAvg
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public BigDecimal getCpuTotalTimeAvg() {
    return cpuTotalTimeAvg;
  }


  public void setCpuTotalTimeAvg(BigDecimal cpuTotalTimeAvg) {
    this.cpuTotalTimeAvg = cpuTotalTimeAvg;
  }


  public GuidanceVmwareSizingConfig cpuTotalTimeSum(BigDecimal cpuTotalTimeSum) {
    
    this.cpuTotalTimeSum = cpuTotalTimeSum;
    return this;
  }

   /**
   * Get cpuTotalTimeSum
   * @return cpuTotalTimeSum
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public BigDecimal getCpuTotalTimeSum() {
    return cpuTotalTimeSum;
  }


  public void setCpuTotalTimeSum(BigDecimal cpuTotalTimeSum) {
    this.cpuTotalTimeSum = cpuTotalTimeSum;
  }


  public GuidanceVmwareSizingConfig cpuIdleTimeCount(BigDecimal cpuIdleTimeCount) {
    
    this.cpuIdleTimeCount = cpuIdleTimeCount;
    return this;
  }

   /**
   * Get cpuIdleTimeCount
   * @return cpuIdleTimeCount
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public BigDecimal getCpuIdleTimeCount() {
    return cpuIdleTimeCount;
  }


  public void setCpuIdleTimeCount(BigDecimal cpuIdleTimeCount) {
    this.cpuIdleTimeCount = cpuIdleTimeCount;
  }


  public GuidanceVmwareSizingConfig cpuIdleTimeMin(BigDecimal cpuIdleTimeMin) {
    
    this.cpuIdleTimeMin = cpuIdleTimeMin;
    return this;
  }

   /**
   * Get cpuIdleTimeMin
   * @return cpuIdleTimeMin
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public BigDecimal getCpuIdleTimeMin() {
    return cpuIdleTimeMin;
  }


  public void setCpuIdleTimeMin(BigDecimal cpuIdleTimeMin) {
    this.cpuIdleTimeMin = cpuIdleTimeMin;
  }


  public GuidanceVmwareSizingConfig cpuIdleTimeMax(BigDecimal cpuIdleTimeMax) {
    
    this.cpuIdleTimeMax = cpuIdleTimeMax;
    return this;
  }

   /**
   * Get cpuIdleTimeMax
   * @return cpuIdleTimeMax
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public BigDecimal getCpuIdleTimeMax() {
    return cpuIdleTimeMax;
  }


  public void setCpuIdleTimeMax(BigDecimal cpuIdleTimeMax) {
    this.cpuIdleTimeMax = cpuIdleTimeMax;
  }


  public GuidanceVmwareSizingConfig cpuIdleTimeAvg(BigDecimal cpuIdleTimeAvg) {
    
    this.cpuIdleTimeAvg = cpuIdleTimeAvg;
    return this;
  }

   /**
   * Get cpuIdleTimeAvg
   * @return cpuIdleTimeAvg
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public BigDecimal getCpuIdleTimeAvg() {
    return cpuIdleTimeAvg;
  }


  public void setCpuIdleTimeAvg(BigDecimal cpuIdleTimeAvg) {
    this.cpuIdleTimeAvg = cpuIdleTimeAvg;
  }


  public GuidanceVmwareSizingConfig cpuIdleTimeSum(BigDecimal cpuIdleTimeSum) {
    
    this.cpuIdleTimeSum = cpuIdleTimeSum;
    return this;
  }

   /**
   * Get cpuIdleTimeSum
   * @return cpuIdleTimeSum
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public BigDecimal getCpuIdleTimeSum() {
    return cpuIdleTimeSum;
  }


  public void setCpuIdleTimeSum(BigDecimal cpuIdleTimeSum) {
    this.cpuIdleTimeSum = cpuIdleTimeSum;
  }


  public GuidanceVmwareSizingConfig cpuUsageCount(BigDecimal cpuUsageCount) {
    
    this.cpuUsageCount = cpuUsageCount;
    return this;
  }

   /**
   * Get cpuUsageCount
   * @return cpuUsageCount
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public BigDecimal getCpuUsageCount() {
    return cpuUsageCount;
  }


  public void setCpuUsageCount(BigDecimal cpuUsageCount) {
    this.cpuUsageCount = cpuUsageCount;
  }


  public GuidanceVmwareSizingConfig cpuUsageMin(BigDecimal cpuUsageMin) {
    
    this.cpuUsageMin = cpuUsageMin;
    return this;
  }

   /**
   * Get cpuUsageMin
   * @return cpuUsageMin
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public BigDecimal getCpuUsageMin() {
    return cpuUsageMin;
  }


  public void setCpuUsageMin(BigDecimal cpuUsageMin) {
    this.cpuUsageMin = cpuUsageMin;
  }


  public GuidanceVmwareSizingConfig cpuUsageMax(BigDecimal cpuUsageMax) {
    
    this.cpuUsageMax = cpuUsageMax;
    return this;
  }

   /**
   * Get cpuUsageMax
   * @return cpuUsageMax
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public BigDecimal getCpuUsageMax() {
    return cpuUsageMax;
  }


  public void setCpuUsageMax(BigDecimal cpuUsageMax) {
    this.cpuUsageMax = cpuUsageMax;
  }


  public GuidanceVmwareSizingConfig cpuUsageAvg(BigDecimal cpuUsageAvg) {
    
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


  public GuidanceVmwareSizingConfig cpuUsageSum(BigDecimal cpuUsageSum) {
    
    this.cpuUsageSum = cpuUsageSum;
    return this;
  }

   /**
   * Get cpuUsageSum
   * @return cpuUsageSum
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public BigDecimal getCpuUsageSum() {
    return cpuUsageSum;
  }


  public void setCpuUsageSum(BigDecimal cpuUsageSum) {
    this.cpuUsageSum = cpuUsageSum;
  }


  public GuidanceVmwareSizingConfig maxMemoryCount(BigDecimal maxMemoryCount) {
    
    this.maxMemoryCount = maxMemoryCount;
    return this;
  }

   /**
   * Get maxMemoryCount
   * @return maxMemoryCount
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public BigDecimal getMaxMemoryCount() {
    return maxMemoryCount;
  }


  public void setMaxMemoryCount(BigDecimal maxMemoryCount) {
    this.maxMemoryCount = maxMemoryCount;
  }


  public GuidanceVmwareSizingConfig maxMemoryMin(BigDecimal maxMemoryMin) {
    
    this.maxMemoryMin = maxMemoryMin;
    return this;
  }

   /**
   * Get maxMemoryMin
   * @return maxMemoryMin
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public BigDecimal getMaxMemoryMin() {
    return maxMemoryMin;
  }


  public void setMaxMemoryMin(BigDecimal maxMemoryMin) {
    this.maxMemoryMin = maxMemoryMin;
  }


  public GuidanceVmwareSizingConfig maxMemoryMax(BigDecimal maxMemoryMax) {
    
    this.maxMemoryMax = maxMemoryMax;
    return this;
  }

   /**
   * Get maxMemoryMax
   * @return maxMemoryMax
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public BigDecimal getMaxMemoryMax() {
    return maxMemoryMax;
  }


  public void setMaxMemoryMax(BigDecimal maxMemoryMax) {
    this.maxMemoryMax = maxMemoryMax;
  }


  public GuidanceVmwareSizingConfig maxMemoryAvg(BigDecimal maxMemoryAvg) {
    
    this.maxMemoryAvg = maxMemoryAvg;
    return this;
  }

   /**
   * Get maxMemoryAvg
   * @return maxMemoryAvg
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public BigDecimal getMaxMemoryAvg() {
    return maxMemoryAvg;
  }


  public void setMaxMemoryAvg(BigDecimal maxMemoryAvg) {
    this.maxMemoryAvg = maxMemoryAvg;
  }


  public GuidanceVmwareSizingConfig maxMemorySum(BigDecimal maxMemorySum) {
    
    this.maxMemorySum = maxMemorySum;
    return this;
  }

   /**
   * Get maxMemorySum
   * @return maxMemorySum
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public BigDecimal getMaxMemorySum() {
    return maxMemorySum;
  }


  public void setMaxMemorySum(BigDecimal maxMemorySum) {
    this.maxMemorySum = maxMemorySum;
  }


  public GuidanceVmwareSizingConfig cpuUserTimeCount(BigDecimal cpuUserTimeCount) {
    
    this.cpuUserTimeCount = cpuUserTimeCount;
    return this;
  }

   /**
   * Get cpuUserTimeCount
   * @return cpuUserTimeCount
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public BigDecimal getCpuUserTimeCount() {
    return cpuUserTimeCount;
  }


  public void setCpuUserTimeCount(BigDecimal cpuUserTimeCount) {
    this.cpuUserTimeCount = cpuUserTimeCount;
  }


  public GuidanceVmwareSizingConfig cpuUserTimeMin(BigDecimal cpuUserTimeMin) {
    
    this.cpuUserTimeMin = cpuUserTimeMin;
    return this;
  }

   /**
   * Get cpuUserTimeMin
   * @return cpuUserTimeMin
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public BigDecimal getCpuUserTimeMin() {
    return cpuUserTimeMin;
  }


  public void setCpuUserTimeMin(BigDecimal cpuUserTimeMin) {
    this.cpuUserTimeMin = cpuUserTimeMin;
  }


  public GuidanceVmwareSizingConfig cpuUserTimeMax(BigDecimal cpuUserTimeMax) {
    
    this.cpuUserTimeMax = cpuUserTimeMax;
    return this;
  }

   /**
   * Get cpuUserTimeMax
   * @return cpuUserTimeMax
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public BigDecimal getCpuUserTimeMax() {
    return cpuUserTimeMax;
  }


  public void setCpuUserTimeMax(BigDecimal cpuUserTimeMax) {
    this.cpuUserTimeMax = cpuUserTimeMax;
  }


  public GuidanceVmwareSizingConfig cpuUserTimeAvg(BigDecimal cpuUserTimeAvg) {
    
    this.cpuUserTimeAvg = cpuUserTimeAvg;
    return this;
  }

   /**
   * Get cpuUserTimeAvg
   * @return cpuUserTimeAvg
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public BigDecimal getCpuUserTimeAvg() {
    return cpuUserTimeAvg;
  }


  public void setCpuUserTimeAvg(BigDecimal cpuUserTimeAvg) {
    this.cpuUserTimeAvg = cpuUserTimeAvg;
  }


  public GuidanceVmwareSizingConfig cpuUserTimeSum(BigDecimal cpuUserTimeSum) {
    
    this.cpuUserTimeSum = cpuUserTimeSum;
    return this;
  }

   /**
   * Get cpuUserTimeSum
   * @return cpuUserTimeSum
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public BigDecimal getCpuUserTimeSum() {
    return cpuUserTimeSum;
  }


  public void setCpuUserTimeSum(BigDecimal cpuUserTimeSum) {
    this.cpuUserTimeSum = cpuUserTimeSum;
  }


  public GuidanceVmwareSizingConfig cpuSystemTimeCount(BigDecimal cpuSystemTimeCount) {
    
    this.cpuSystemTimeCount = cpuSystemTimeCount;
    return this;
  }

   /**
   * Get cpuSystemTimeCount
   * @return cpuSystemTimeCount
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public BigDecimal getCpuSystemTimeCount() {
    return cpuSystemTimeCount;
  }


  public void setCpuSystemTimeCount(BigDecimal cpuSystemTimeCount) {
    this.cpuSystemTimeCount = cpuSystemTimeCount;
  }


  public GuidanceVmwareSizingConfig cpuSystemTimeMin(BigDecimal cpuSystemTimeMin) {
    
    this.cpuSystemTimeMin = cpuSystemTimeMin;
    return this;
  }

   /**
   * Get cpuSystemTimeMin
   * @return cpuSystemTimeMin
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public BigDecimal getCpuSystemTimeMin() {
    return cpuSystemTimeMin;
  }


  public void setCpuSystemTimeMin(BigDecimal cpuSystemTimeMin) {
    this.cpuSystemTimeMin = cpuSystemTimeMin;
  }


  public GuidanceVmwareSizingConfig cpuSystemTimeMax(BigDecimal cpuSystemTimeMax) {
    
    this.cpuSystemTimeMax = cpuSystemTimeMax;
    return this;
  }

   /**
   * Get cpuSystemTimeMax
   * @return cpuSystemTimeMax
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public BigDecimal getCpuSystemTimeMax() {
    return cpuSystemTimeMax;
  }


  public void setCpuSystemTimeMax(BigDecimal cpuSystemTimeMax) {
    this.cpuSystemTimeMax = cpuSystemTimeMax;
  }


  public GuidanceVmwareSizingConfig cpuSystemTimeAvg(BigDecimal cpuSystemTimeAvg) {
    
    this.cpuSystemTimeAvg = cpuSystemTimeAvg;
    return this;
  }

   /**
   * Get cpuSystemTimeAvg
   * @return cpuSystemTimeAvg
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public BigDecimal getCpuSystemTimeAvg() {
    return cpuSystemTimeAvg;
  }


  public void setCpuSystemTimeAvg(BigDecimal cpuSystemTimeAvg) {
    this.cpuSystemTimeAvg = cpuSystemTimeAvg;
  }


  public GuidanceVmwareSizingConfig cpuSystemTimeSum(BigDecimal cpuSystemTimeSum) {
    
    this.cpuSystemTimeSum = cpuSystemTimeSum;
    return this;
  }

   /**
   * Get cpuSystemTimeSum
   * @return cpuSystemTimeSum
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public BigDecimal getCpuSystemTimeSum() {
    return cpuSystemTimeSum;
  }


  public void setCpuSystemTimeSum(BigDecimal cpuSystemTimeSum) {
    this.cpuSystemTimeSum = cpuSystemTimeSum;
  }


  public GuidanceVmwareSizingConfig usedMemoryCount(BigDecimal usedMemoryCount) {
    
    this.usedMemoryCount = usedMemoryCount;
    return this;
  }

   /**
   * Get usedMemoryCount
   * @return usedMemoryCount
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public BigDecimal getUsedMemoryCount() {
    return usedMemoryCount;
  }


  public void setUsedMemoryCount(BigDecimal usedMemoryCount) {
    this.usedMemoryCount = usedMemoryCount;
  }


  public GuidanceVmwareSizingConfig usedMemoryMin(BigDecimal usedMemoryMin) {
    
    this.usedMemoryMin = usedMemoryMin;
    return this;
  }

   /**
   * Get usedMemoryMin
   * @return usedMemoryMin
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public BigDecimal getUsedMemoryMin() {
    return usedMemoryMin;
  }


  public void setUsedMemoryMin(BigDecimal usedMemoryMin) {
    this.usedMemoryMin = usedMemoryMin;
  }


  public GuidanceVmwareSizingConfig usedMemoryMax(BigDecimal usedMemoryMax) {
    
    this.usedMemoryMax = usedMemoryMax;
    return this;
  }

   /**
   * Get usedMemoryMax
   * @return usedMemoryMax
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public BigDecimal getUsedMemoryMax() {
    return usedMemoryMax;
  }


  public void setUsedMemoryMax(BigDecimal usedMemoryMax) {
    this.usedMemoryMax = usedMemoryMax;
  }


  public GuidanceVmwareSizingConfig usedMemoryAvg(BigDecimal usedMemoryAvg) {
    
    this.usedMemoryAvg = usedMemoryAvg;
    return this;
  }

   /**
   * Get usedMemoryAvg
   * @return usedMemoryAvg
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public BigDecimal getUsedMemoryAvg() {
    return usedMemoryAvg;
  }


  public void setUsedMemoryAvg(BigDecimal usedMemoryAvg) {
    this.usedMemoryAvg = usedMemoryAvg;
  }


  public GuidanceVmwareSizingConfig usedMemorySum(BigDecimal usedMemorySum) {
    
    this.usedMemorySum = usedMemorySum;
    return this;
  }

   /**
   * Get usedMemorySum
   * @return usedMemorySum
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public BigDecimal getUsedMemorySum() {
    return usedMemorySum;
  }


  public void setUsedMemorySum(BigDecimal usedMemorySum) {
    this.usedMemorySum = usedMemorySum;
  }


  public GuidanceVmwareSizingConfig freeMemoryCount(BigDecimal freeMemoryCount) {
    
    this.freeMemoryCount = freeMemoryCount;
    return this;
  }

   /**
   * Get freeMemoryCount
   * @return freeMemoryCount
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public BigDecimal getFreeMemoryCount() {
    return freeMemoryCount;
  }


  public void setFreeMemoryCount(BigDecimal freeMemoryCount) {
    this.freeMemoryCount = freeMemoryCount;
  }


  public GuidanceVmwareSizingConfig freeMemoryMin(BigDecimal freeMemoryMin) {
    
    this.freeMemoryMin = freeMemoryMin;
    return this;
  }

   /**
   * Get freeMemoryMin
   * @return freeMemoryMin
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public BigDecimal getFreeMemoryMin() {
    return freeMemoryMin;
  }


  public void setFreeMemoryMin(BigDecimal freeMemoryMin) {
    this.freeMemoryMin = freeMemoryMin;
  }


  public GuidanceVmwareSizingConfig freeMemoryMax(BigDecimal freeMemoryMax) {
    
    this.freeMemoryMax = freeMemoryMax;
    return this;
  }

   /**
   * Get freeMemoryMax
   * @return freeMemoryMax
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public BigDecimal getFreeMemoryMax() {
    return freeMemoryMax;
  }


  public void setFreeMemoryMax(BigDecimal freeMemoryMax) {
    this.freeMemoryMax = freeMemoryMax;
  }


  public GuidanceVmwareSizingConfig freeMemoryAvg(BigDecimal freeMemoryAvg) {
    
    this.freeMemoryAvg = freeMemoryAvg;
    return this;
  }

   /**
   * Get freeMemoryAvg
   * @return freeMemoryAvg
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public BigDecimal getFreeMemoryAvg() {
    return freeMemoryAvg;
  }


  public void setFreeMemoryAvg(BigDecimal freeMemoryAvg) {
    this.freeMemoryAvg = freeMemoryAvg;
  }


  public GuidanceVmwareSizingConfig freeMemorySum(BigDecimal freeMemorySum) {
    
    this.freeMemorySum = freeMemorySum;
    return this;
  }

   /**
   * Get freeMemorySum
   * @return freeMemorySum
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public BigDecimal getFreeMemorySum() {
    return freeMemorySum;
  }


  public void setFreeMemorySum(BigDecimal freeMemorySum) {
    this.freeMemorySum = freeMemorySum;
  }


  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (o == null || getClass() != o.getClass()) {
      return false;
    }
    GuidanceVmwareSizingConfig guidanceVmwareSizingConfig = (GuidanceVmwareSizingConfig) o;
    return Objects.equals(this.exists, guidanceVmwareSizingConfig.exists) &&
        Objects.equals(this.objectId, guidanceVmwareSizingConfig.objectId) &&
        Objects.equals(this.cpuTotalTimeCount, guidanceVmwareSizingConfig.cpuTotalTimeCount) &&
        Objects.equals(this.cpuTotalTimeMin, guidanceVmwareSizingConfig.cpuTotalTimeMin) &&
        Objects.equals(this.cpuTotalTimeMax, guidanceVmwareSizingConfig.cpuTotalTimeMax) &&
        Objects.equals(this.cpuTotalTimeAvg, guidanceVmwareSizingConfig.cpuTotalTimeAvg) &&
        Objects.equals(this.cpuTotalTimeSum, guidanceVmwareSizingConfig.cpuTotalTimeSum) &&
        Objects.equals(this.cpuIdleTimeCount, guidanceVmwareSizingConfig.cpuIdleTimeCount) &&
        Objects.equals(this.cpuIdleTimeMin, guidanceVmwareSizingConfig.cpuIdleTimeMin) &&
        Objects.equals(this.cpuIdleTimeMax, guidanceVmwareSizingConfig.cpuIdleTimeMax) &&
        Objects.equals(this.cpuIdleTimeAvg, guidanceVmwareSizingConfig.cpuIdleTimeAvg) &&
        Objects.equals(this.cpuIdleTimeSum, guidanceVmwareSizingConfig.cpuIdleTimeSum) &&
        Objects.equals(this.cpuUsageCount, guidanceVmwareSizingConfig.cpuUsageCount) &&
        Objects.equals(this.cpuUsageMin, guidanceVmwareSizingConfig.cpuUsageMin) &&
        Objects.equals(this.cpuUsageMax, guidanceVmwareSizingConfig.cpuUsageMax) &&
        Objects.equals(this.cpuUsageAvg, guidanceVmwareSizingConfig.cpuUsageAvg) &&
        Objects.equals(this.cpuUsageSum, guidanceVmwareSizingConfig.cpuUsageSum) &&
        Objects.equals(this.maxMemoryCount, guidanceVmwareSizingConfig.maxMemoryCount) &&
        Objects.equals(this.maxMemoryMin, guidanceVmwareSizingConfig.maxMemoryMin) &&
        Objects.equals(this.maxMemoryMax, guidanceVmwareSizingConfig.maxMemoryMax) &&
        Objects.equals(this.maxMemoryAvg, guidanceVmwareSizingConfig.maxMemoryAvg) &&
        Objects.equals(this.maxMemorySum, guidanceVmwareSizingConfig.maxMemorySum) &&
        Objects.equals(this.cpuUserTimeCount, guidanceVmwareSizingConfig.cpuUserTimeCount) &&
        Objects.equals(this.cpuUserTimeMin, guidanceVmwareSizingConfig.cpuUserTimeMin) &&
        Objects.equals(this.cpuUserTimeMax, guidanceVmwareSizingConfig.cpuUserTimeMax) &&
        Objects.equals(this.cpuUserTimeAvg, guidanceVmwareSizingConfig.cpuUserTimeAvg) &&
        Objects.equals(this.cpuUserTimeSum, guidanceVmwareSizingConfig.cpuUserTimeSum) &&
        Objects.equals(this.cpuSystemTimeCount, guidanceVmwareSizingConfig.cpuSystemTimeCount) &&
        Objects.equals(this.cpuSystemTimeMin, guidanceVmwareSizingConfig.cpuSystemTimeMin) &&
        Objects.equals(this.cpuSystemTimeMax, guidanceVmwareSizingConfig.cpuSystemTimeMax) &&
        Objects.equals(this.cpuSystemTimeAvg, guidanceVmwareSizingConfig.cpuSystemTimeAvg) &&
        Objects.equals(this.cpuSystemTimeSum, guidanceVmwareSizingConfig.cpuSystemTimeSum) &&
        Objects.equals(this.usedMemoryCount, guidanceVmwareSizingConfig.usedMemoryCount) &&
        Objects.equals(this.usedMemoryMin, guidanceVmwareSizingConfig.usedMemoryMin) &&
        Objects.equals(this.usedMemoryMax, guidanceVmwareSizingConfig.usedMemoryMax) &&
        Objects.equals(this.usedMemoryAvg, guidanceVmwareSizingConfig.usedMemoryAvg) &&
        Objects.equals(this.usedMemorySum, guidanceVmwareSizingConfig.usedMemorySum) &&
        Objects.equals(this.freeMemoryCount, guidanceVmwareSizingConfig.freeMemoryCount) &&
        Objects.equals(this.freeMemoryMin, guidanceVmwareSizingConfig.freeMemoryMin) &&
        Objects.equals(this.freeMemoryMax, guidanceVmwareSizingConfig.freeMemoryMax) &&
        Objects.equals(this.freeMemoryAvg, guidanceVmwareSizingConfig.freeMemoryAvg) &&
        Objects.equals(this.freeMemorySum, guidanceVmwareSizingConfig.freeMemorySum);
  }

  @Override
  public int hashCode() {
    return Objects.hash(exists, objectId, cpuTotalTimeCount, cpuTotalTimeMin, cpuTotalTimeMax, cpuTotalTimeAvg, cpuTotalTimeSum, cpuIdleTimeCount, cpuIdleTimeMin, cpuIdleTimeMax, cpuIdleTimeAvg, cpuIdleTimeSum, cpuUsageCount, cpuUsageMin, cpuUsageMax, cpuUsageAvg, cpuUsageSum, maxMemoryCount, maxMemoryMin, maxMemoryMax, maxMemoryAvg, maxMemorySum, cpuUserTimeCount, cpuUserTimeMin, cpuUserTimeMax, cpuUserTimeAvg, cpuUserTimeSum, cpuSystemTimeCount, cpuSystemTimeMin, cpuSystemTimeMax, cpuSystemTimeAvg, cpuSystemTimeSum, usedMemoryCount, usedMemoryMin, usedMemoryMax, usedMemoryAvg, usedMemorySum, freeMemoryCount, freeMemoryMin, freeMemoryMax, freeMemoryAvg, freeMemorySum);
  }


  @Override
  public String toString() {
    StringBuilder sb = new StringBuilder();
    sb.append("class GuidanceVmwareSizingConfig {\n");
    sb.append("    exists: ").append(toIndentedString(exists)).append("\n");
    sb.append("    objectId: ").append(toIndentedString(objectId)).append("\n");
    sb.append("    cpuTotalTimeCount: ").append(toIndentedString(cpuTotalTimeCount)).append("\n");
    sb.append("    cpuTotalTimeMin: ").append(toIndentedString(cpuTotalTimeMin)).append("\n");
    sb.append("    cpuTotalTimeMax: ").append(toIndentedString(cpuTotalTimeMax)).append("\n");
    sb.append("    cpuTotalTimeAvg: ").append(toIndentedString(cpuTotalTimeAvg)).append("\n");
    sb.append("    cpuTotalTimeSum: ").append(toIndentedString(cpuTotalTimeSum)).append("\n");
    sb.append("    cpuIdleTimeCount: ").append(toIndentedString(cpuIdleTimeCount)).append("\n");
    sb.append("    cpuIdleTimeMin: ").append(toIndentedString(cpuIdleTimeMin)).append("\n");
    sb.append("    cpuIdleTimeMax: ").append(toIndentedString(cpuIdleTimeMax)).append("\n");
    sb.append("    cpuIdleTimeAvg: ").append(toIndentedString(cpuIdleTimeAvg)).append("\n");
    sb.append("    cpuIdleTimeSum: ").append(toIndentedString(cpuIdleTimeSum)).append("\n");
    sb.append("    cpuUsageCount: ").append(toIndentedString(cpuUsageCount)).append("\n");
    sb.append("    cpuUsageMin: ").append(toIndentedString(cpuUsageMin)).append("\n");
    sb.append("    cpuUsageMax: ").append(toIndentedString(cpuUsageMax)).append("\n");
    sb.append("    cpuUsageAvg: ").append(toIndentedString(cpuUsageAvg)).append("\n");
    sb.append("    cpuUsageSum: ").append(toIndentedString(cpuUsageSum)).append("\n");
    sb.append("    maxMemoryCount: ").append(toIndentedString(maxMemoryCount)).append("\n");
    sb.append("    maxMemoryMin: ").append(toIndentedString(maxMemoryMin)).append("\n");
    sb.append("    maxMemoryMax: ").append(toIndentedString(maxMemoryMax)).append("\n");
    sb.append("    maxMemoryAvg: ").append(toIndentedString(maxMemoryAvg)).append("\n");
    sb.append("    maxMemorySum: ").append(toIndentedString(maxMemorySum)).append("\n");
    sb.append("    cpuUserTimeCount: ").append(toIndentedString(cpuUserTimeCount)).append("\n");
    sb.append("    cpuUserTimeMin: ").append(toIndentedString(cpuUserTimeMin)).append("\n");
    sb.append("    cpuUserTimeMax: ").append(toIndentedString(cpuUserTimeMax)).append("\n");
    sb.append("    cpuUserTimeAvg: ").append(toIndentedString(cpuUserTimeAvg)).append("\n");
    sb.append("    cpuUserTimeSum: ").append(toIndentedString(cpuUserTimeSum)).append("\n");
    sb.append("    cpuSystemTimeCount: ").append(toIndentedString(cpuSystemTimeCount)).append("\n");
    sb.append("    cpuSystemTimeMin: ").append(toIndentedString(cpuSystemTimeMin)).append("\n");
    sb.append("    cpuSystemTimeMax: ").append(toIndentedString(cpuSystemTimeMax)).append("\n");
    sb.append("    cpuSystemTimeAvg: ").append(toIndentedString(cpuSystemTimeAvg)).append("\n");
    sb.append("    cpuSystemTimeSum: ").append(toIndentedString(cpuSystemTimeSum)).append("\n");
    sb.append("    usedMemoryCount: ").append(toIndentedString(usedMemoryCount)).append("\n");
    sb.append("    usedMemoryMin: ").append(toIndentedString(usedMemoryMin)).append("\n");
    sb.append("    usedMemoryMax: ").append(toIndentedString(usedMemoryMax)).append("\n");
    sb.append("    usedMemoryAvg: ").append(toIndentedString(usedMemoryAvg)).append("\n");
    sb.append("    usedMemorySum: ").append(toIndentedString(usedMemorySum)).append("\n");
    sb.append("    freeMemoryCount: ").append(toIndentedString(freeMemoryCount)).append("\n");
    sb.append("    freeMemoryMin: ").append(toIndentedString(freeMemoryMin)).append("\n");
    sb.append("    freeMemoryMax: ").append(toIndentedString(freeMemoryMax)).append("\n");
    sb.append("    freeMemoryAvg: ").append(toIndentedString(freeMemoryAvg)).append("\n");
    sb.append("    freeMemorySum: ").append(toIndentedString(freeMemorySum)).append("\n");
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
