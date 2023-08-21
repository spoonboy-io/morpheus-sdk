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
 * HealthMemory
 */
@javax.annotation.Generated(value = "org.openapitools.codegen.languages.JavaClientCodegen", date = "2023-08-21T11:26:47.028Z[GMT]")
public class HealthMemory {
  public static final String SERIALIZED_NAME_SUCCESS = "success";
  @SerializedName(SERIALIZED_NAME_SUCCESS)
  private Boolean success;

  public static final String SERIALIZED_NAME_MAX_MEMORY = "maxMemory";
  @SerializedName(SERIALIZED_NAME_MAX_MEMORY)
  private Long maxMemory;

  public static final String SERIALIZED_NAME_TOTAL_MEMORY = "totalMemory";
  @SerializedName(SERIALIZED_NAME_TOTAL_MEMORY)
  private Long totalMemory;

  public static final String SERIALIZED_NAME_FREE_MEMORY = "freeMemory";
  @SerializedName(SERIALIZED_NAME_FREE_MEMORY)
  private Long freeMemory;

  public static final String SERIALIZED_NAME_USED_MEMORY = "usedMemory";
  @SerializedName(SERIALIZED_NAME_USED_MEMORY)
  private Long usedMemory;

  public static final String SERIALIZED_NAME_SYSTEM_MEMORY = "systemMemory";
  @SerializedName(SERIALIZED_NAME_SYSTEM_MEMORY)
  private Long systemMemory;

  public static final String SERIALIZED_NAME_COMMITTED_MEMORY = "committedMemory";
  @SerializedName(SERIALIZED_NAME_COMMITTED_MEMORY)
  private Long committedMemory;

  public static final String SERIALIZED_NAME_SYSTEM_FREE_MEMORY = "systemFreeMemory";
  @SerializedName(SERIALIZED_NAME_SYSTEM_FREE_MEMORY)
  private Long systemFreeMemory;

  public static final String SERIALIZED_NAME_SYSTEM_SWAP = "systemSwap";
  @SerializedName(SERIALIZED_NAME_SYSTEM_SWAP)
  private Long systemSwap;

  public static final String SERIALIZED_NAME_SYSTEM_FREE_SWAP = "systemFreeSwap";
  @SerializedName(SERIALIZED_NAME_SYSTEM_FREE_SWAP)
  private Long systemFreeSwap;

  public static final String SERIALIZED_NAME_SWAP_PERCENT = "swapPercent";
  @SerializedName(SERIALIZED_NAME_SWAP_PERCENT)
  private Long swapPercent;

  public static final String SERIALIZED_NAME_MEMORY_PERCENT = "memoryPercent";
  @SerializedName(SERIALIZED_NAME_MEMORY_PERCENT)
  private BigDecimal memoryPercent;

  public static final String SERIALIZED_NAME_SYSTEM_MEMORY_PERCENT = "systemMemoryPercent";
  @SerializedName(SERIALIZED_NAME_SYSTEM_MEMORY_PERCENT)
  private BigDecimal systemMemoryPercent;

  public static final String SERIALIZED_NAME_STATUS = "status";
  @SerializedName(SERIALIZED_NAME_STATUS)
  private String status;


  public HealthMemory success(Boolean success) {
    
    this.success = success;
    return this;
  }

   /**
   * Get success
   * @return success
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public Boolean getSuccess() {
    return success;
  }


  public void setSuccess(Boolean success) {
    this.success = success;
  }


  public HealthMemory maxMemory(Long maxMemory) {
    
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


  public HealthMemory totalMemory(Long totalMemory) {
    
    this.totalMemory = totalMemory;
    return this;
  }

   /**
   * Get totalMemory
   * @return totalMemory
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public Long getTotalMemory() {
    return totalMemory;
  }


  public void setTotalMemory(Long totalMemory) {
    this.totalMemory = totalMemory;
  }


  public HealthMemory freeMemory(Long freeMemory) {
    
    this.freeMemory = freeMemory;
    return this;
  }

   /**
   * Get freeMemory
   * @return freeMemory
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public Long getFreeMemory() {
    return freeMemory;
  }


  public void setFreeMemory(Long freeMemory) {
    this.freeMemory = freeMemory;
  }


  public HealthMemory usedMemory(Long usedMemory) {
    
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


  public HealthMemory systemMemory(Long systemMemory) {
    
    this.systemMemory = systemMemory;
    return this;
  }

   /**
   * Get systemMemory
   * @return systemMemory
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public Long getSystemMemory() {
    return systemMemory;
  }


  public void setSystemMemory(Long systemMemory) {
    this.systemMemory = systemMemory;
  }


  public HealthMemory committedMemory(Long committedMemory) {
    
    this.committedMemory = committedMemory;
    return this;
  }

   /**
   * Get committedMemory
   * @return committedMemory
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public Long getCommittedMemory() {
    return committedMemory;
  }


  public void setCommittedMemory(Long committedMemory) {
    this.committedMemory = committedMemory;
  }


  public HealthMemory systemFreeMemory(Long systemFreeMemory) {
    
    this.systemFreeMemory = systemFreeMemory;
    return this;
  }

   /**
   * Get systemFreeMemory
   * @return systemFreeMemory
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public Long getSystemFreeMemory() {
    return systemFreeMemory;
  }


  public void setSystemFreeMemory(Long systemFreeMemory) {
    this.systemFreeMemory = systemFreeMemory;
  }


  public HealthMemory systemSwap(Long systemSwap) {
    
    this.systemSwap = systemSwap;
    return this;
  }

   /**
   * Get systemSwap
   * @return systemSwap
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public Long getSystemSwap() {
    return systemSwap;
  }


  public void setSystemSwap(Long systemSwap) {
    this.systemSwap = systemSwap;
  }


  public HealthMemory systemFreeSwap(Long systemFreeSwap) {
    
    this.systemFreeSwap = systemFreeSwap;
    return this;
  }

   /**
   * Get systemFreeSwap
   * @return systemFreeSwap
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public Long getSystemFreeSwap() {
    return systemFreeSwap;
  }


  public void setSystemFreeSwap(Long systemFreeSwap) {
    this.systemFreeSwap = systemFreeSwap;
  }


  public HealthMemory swapPercent(Long swapPercent) {
    
    this.swapPercent = swapPercent;
    return this;
  }

   /**
   * Get swapPercent
   * @return swapPercent
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public Long getSwapPercent() {
    return swapPercent;
  }


  public void setSwapPercent(Long swapPercent) {
    this.swapPercent = swapPercent;
  }


  public HealthMemory memoryPercent(BigDecimal memoryPercent) {
    
    this.memoryPercent = memoryPercent;
    return this;
  }

   /**
   * Get memoryPercent
   * @return memoryPercent
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public BigDecimal getMemoryPercent() {
    return memoryPercent;
  }


  public void setMemoryPercent(BigDecimal memoryPercent) {
    this.memoryPercent = memoryPercent;
  }


  public HealthMemory systemMemoryPercent(BigDecimal systemMemoryPercent) {
    
    this.systemMemoryPercent = systemMemoryPercent;
    return this;
  }

   /**
   * Get systemMemoryPercent
   * @return systemMemoryPercent
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public BigDecimal getSystemMemoryPercent() {
    return systemMemoryPercent;
  }


  public void setSystemMemoryPercent(BigDecimal systemMemoryPercent) {
    this.systemMemoryPercent = systemMemoryPercent;
  }


  public HealthMemory status(String status) {
    
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


  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (o == null || getClass() != o.getClass()) {
      return false;
    }
    HealthMemory healthMemory = (HealthMemory) o;
    return Objects.equals(this.success, healthMemory.success) &&
        Objects.equals(this.maxMemory, healthMemory.maxMemory) &&
        Objects.equals(this.totalMemory, healthMemory.totalMemory) &&
        Objects.equals(this.freeMemory, healthMemory.freeMemory) &&
        Objects.equals(this.usedMemory, healthMemory.usedMemory) &&
        Objects.equals(this.systemMemory, healthMemory.systemMemory) &&
        Objects.equals(this.committedMemory, healthMemory.committedMemory) &&
        Objects.equals(this.systemFreeMemory, healthMemory.systemFreeMemory) &&
        Objects.equals(this.systemSwap, healthMemory.systemSwap) &&
        Objects.equals(this.systemFreeSwap, healthMemory.systemFreeSwap) &&
        Objects.equals(this.swapPercent, healthMemory.swapPercent) &&
        Objects.equals(this.memoryPercent, healthMemory.memoryPercent) &&
        Objects.equals(this.systemMemoryPercent, healthMemory.systemMemoryPercent) &&
        Objects.equals(this.status, healthMemory.status);
  }

  @Override
  public int hashCode() {
    return Objects.hash(success, maxMemory, totalMemory, freeMemory, usedMemory, systemMemory, committedMemory, systemFreeMemory, systemSwap, systemFreeSwap, swapPercent, memoryPercent, systemMemoryPercent, status);
  }


  @Override
  public String toString() {
    StringBuilder sb = new StringBuilder();
    sb.append("class HealthMemory {\n");
    sb.append("    success: ").append(toIndentedString(success)).append("\n");
    sb.append("    maxMemory: ").append(toIndentedString(maxMemory)).append("\n");
    sb.append("    totalMemory: ").append(toIndentedString(totalMemory)).append("\n");
    sb.append("    freeMemory: ").append(toIndentedString(freeMemory)).append("\n");
    sb.append("    usedMemory: ").append(toIndentedString(usedMemory)).append("\n");
    sb.append("    systemMemory: ").append(toIndentedString(systemMemory)).append("\n");
    sb.append("    committedMemory: ").append(toIndentedString(committedMemory)).append("\n");
    sb.append("    systemFreeMemory: ").append(toIndentedString(systemFreeMemory)).append("\n");
    sb.append("    systemSwap: ").append(toIndentedString(systemSwap)).append("\n");
    sb.append("    systemFreeSwap: ").append(toIndentedString(systemFreeSwap)).append("\n");
    sb.append("    swapPercent: ").append(toIndentedString(swapPercent)).append("\n");
    sb.append("    memoryPercent: ").append(toIndentedString(memoryPercent)).append("\n");
    sb.append("    systemMemoryPercent: ").append(toIndentedString(systemMemoryPercent)).append("\n");
    sb.append("    status: ").append(toIndentedString(status)).append("\n");
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
