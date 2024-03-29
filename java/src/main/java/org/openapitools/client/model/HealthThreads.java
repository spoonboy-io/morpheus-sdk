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
import org.openapitools.client.model.HealthThreadsBusyThreads;

/**
 * HealthThreads
 */
@javax.annotation.Generated(value = "org.openapitools.codegen.languages.JavaClientCodegen", date = "2023-08-21T11:26:47.028Z[GMT]")
public class HealthThreads {
  public static final String SERIALIZED_NAME_SUCCESS = "success";
  @SerializedName(SERIALIZED_NAME_SUCCESS)
  private Boolean success;

  public static final String SERIALIZED_NAME_THREAD_LIST = "threadList";
  @SerializedName(SERIALIZED_NAME_THREAD_LIST)
  private List<Object> threadList = null;

  public static final String SERIALIZED_NAME_BUSY_THREADS = "busyThreads";
  @SerializedName(SERIALIZED_NAME_BUSY_THREADS)
  private List<HealthThreadsBusyThreads> busyThreads = null;

  public static final String SERIALIZED_NAME_BLOCKED_THREADS = "blockedThreads";
  @SerializedName(SERIALIZED_NAME_BLOCKED_THREADS)
  private List<Object> blockedThreads = null;

  public static final String SERIALIZED_NAME_RUNNING_THREADS = "runningThreads";
  @SerializedName(SERIALIZED_NAME_RUNNING_THREADS)
  private List<Object> runningThreads = null;

  public static final String SERIALIZED_NAME_TOTAL_CPU_TIME = "totalCpuTime";
  @SerializedName(SERIALIZED_NAME_TOTAL_CPU_TIME)
  private Long totalCpuTime;

  public static final String SERIALIZED_NAME_TOTAL_THREADS = "totalThreads";
  @SerializedName(SERIALIZED_NAME_TOTAL_THREADS)
  private Long totalThreads;

  public static final String SERIALIZED_NAME_RUNNING_WEB_THREADS = "runningWebThreads";
  @SerializedName(SERIALIZED_NAME_RUNNING_WEB_THREADS)
  private Long runningWebThreads;

  public static final String SERIALIZED_NAME_STATUS = "status";
  @SerializedName(SERIALIZED_NAME_STATUS)
  private String status;


  public HealthThreads success(Boolean success) {
    
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


  public HealthThreads threadList(List<Object> threadList) {
    
    this.threadList = threadList;
    return this;
  }

  public HealthThreads addThreadListItem(Object threadListItem) {
    if (this.threadList == null) {
      this.threadList = new ArrayList<Object>();
    }
    this.threadList.add(threadListItem);
    return this;
  }

   /**
   * Get threadList
   * @return threadList
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public List<Object> getThreadList() {
    return threadList;
  }


  public void setThreadList(List<Object> threadList) {
    this.threadList = threadList;
  }


  public HealthThreads busyThreads(List<HealthThreadsBusyThreads> busyThreads) {
    
    this.busyThreads = busyThreads;
    return this;
  }

  public HealthThreads addBusyThreadsItem(HealthThreadsBusyThreads busyThreadsItem) {
    if (this.busyThreads == null) {
      this.busyThreads = new ArrayList<HealthThreadsBusyThreads>();
    }
    this.busyThreads.add(busyThreadsItem);
    return this;
  }

   /**
   * Get busyThreads
   * @return busyThreads
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public List<HealthThreadsBusyThreads> getBusyThreads() {
    return busyThreads;
  }


  public void setBusyThreads(List<HealthThreadsBusyThreads> busyThreads) {
    this.busyThreads = busyThreads;
  }


  public HealthThreads blockedThreads(List<Object> blockedThreads) {
    
    this.blockedThreads = blockedThreads;
    return this;
  }

  public HealthThreads addBlockedThreadsItem(Object blockedThreadsItem) {
    if (this.blockedThreads == null) {
      this.blockedThreads = new ArrayList<Object>();
    }
    this.blockedThreads.add(blockedThreadsItem);
    return this;
  }

   /**
   * Get blockedThreads
   * @return blockedThreads
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public List<Object> getBlockedThreads() {
    return blockedThreads;
  }


  public void setBlockedThreads(List<Object> blockedThreads) {
    this.blockedThreads = blockedThreads;
  }


  public HealthThreads runningThreads(List<Object> runningThreads) {
    
    this.runningThreads = runningThreads;
    return this;
  }

  public HealthThreads addRunningThreadsItem(Object runningThreadsItem) {
    if (this.runningThreads == null) {
      this.runningThreads = new ArrayList<Object>();
    }
    this.runningThreads.add(runningThreadsItem);
    return this;
  }

   /**
   * Get runningThreads
   * @return runningThreads
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public List<Object> getRunningThreads() {
    return runningThreads;
  }


  public void setRunningThreads(List<Object> runningThreads) {
    this.runningThreads = runningThreads;
  }


  public HealthThreads totalCpuTime(Long totalCpuTime) {
    
    this.totalCpuTime = totalCpuTime;
    return this;
  }

   /**
   * Get totalCpuTime
   * @return totalCpuTime
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public Long getTotalCpuTime() {
    return totalCpuTime;
  }


  public void setTotalCpuTime(Long totalCpuTime) {
    this.totalCpuTime = totalCpuTime;
  }


  public HealthThreads totalThreads(Long totalThreads) {
    
    this.totalThreads = totalThreads;
    return this;
  }

   /**
   * Get totalThreads
   * @return totalThreads
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public Long getTotalThreads() {
    return totalThreads;
  }


  public void setTotalThreads(Long totalThreads) {
    this.totalThreads = totalThreads;
  }


  public HealthThreads runningWebThreads(Long runningWebThreads) {
    
    this.runningWebThreads = runningWebThreads;
    return this;
  }

   /**
   * Get runningWebThreads
   * @return runningWebThreads
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public Long getRunningWebThreads() {
    return runningWebThreads;
  }


  public void setRunningWebThreads(Long runningWebThreads) {
    this.runningWebThreads = runningWebThreads;
  }


  public HealthThreads status(String status) {
    
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
    HealthThreads healthThreads = (HealthThreads) o;
    return Objects.equals(this.success, healthThreads.success) &&
        Objects.equals(this.threadList, healthThreads.threadList) &&
        Objects.equals(this.busyThreads, healthThreads.busyThreads) &&
        Objects.equals(this.blockedThreads, healthThreads.blockedThreads) &&
        Objects.equals(this.runningThreads, healthThreads.runningThreads) &&
        Objects.equals(this.totalCpuTime, healthThreads.totalCpuTime) &&
        Objects.equals(this.totalThreads, healthThreads.totalThreads) &&
        Objects.equals(this.runningWebThreads, healthThreads.runningWebThreads) &&
        Objects.equals(this.status, healthThreads.status);
  }

  @Override
  public int hashCode() {
    return Objects.hash(success, threadList, busyThreads, blockedThreads, runningThreads, totalCpuTime, totalThreads, runningWebThreads, status);
  }


  @Override
  public String toString() {
    StringBuilder sb = new StringBuilder();
    sb.append("class HealthThreads {\n");
    sb.append("    success: ").append(toIndentedString(success)).append("\n");
    sb.append("    threadList: ").append(toIndentedString(threadList)).append("\n");
    sb.append("    busyThreads: ").append(toIndentedString(busyThreads)).append("\n");
    sb.append("    blockedThreads: ").append(toIndentedString(blockedThreads)).append("\n");
    sb.append("    runningThreads: ").append(toIndentedString(runningThreads)).append("\n");
    sb.append("    totalCpuTime: ").append(toIndentedString(totalCpuTime)).append("\n");
    sb.append("    totalThreads: ").append(toIndentedString(totalThreads)).append("\n");
    sb.append("    runningWebThreads: ").append(toIndentedString(runningWebThreads)).append("\n");
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

