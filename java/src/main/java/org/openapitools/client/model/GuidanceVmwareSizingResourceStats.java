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
 * GuidanceVmwareSizingResourceStats
 */
@javax.annotation.Generated(value = "org.openapitools.codegen.languages.JavaClientCodegen", date = "2023-08-21T11:26:47.028Z[GMT]")
public class GuidanceVmwareSizingResourceStats {
  public static final String SERIALIZED_NAME_TS = "ts";
  @SerializedName(SERIALIZED_NAME_TS)
  private String ts;

  public static final String SERIALIZED_NAME_FREE_MEMORY = "freeMemory";
  @SerializedName(SERIALIZED_NAME_FREE_MEMORY)
  private Long freeMemory;

  public static final String SERIALIZED_NAME_USED_MEMORY = "usedMemory";
  @SerializedName(SERIALIZED_NAME_USED_MEMORY)
  private Long usedMemory;

  public static final String SERIALIZED_NAME_FREE_SWAP = "freeSwap";
  @SerializedName(SERIALIZED_NAME_FREE_SWAP)
  private Long freeSwap;

  public static final String SERIALIZED_NAME_USED_SWAP = "usedSwap";
  @SerializedName(SERIALIZED_NAME_USED_SWAP)
  private Long usedSwap;

  public static final String SERIALIZED_NAME_CPU_IDLE_TIME = "cpuIdleTime";
  @SerializedName(SERIALIZED_NAME_CPU_IDLE_TIME)
  private Long cpuIdleTime;

  public static final String SERIALIZED_NAME_CPU_SYSTEM_TIME = "cpuSystemTime";
  @SerializedName(SERIALIZED_NAME_CPU_SYSTEM_TIME)
  private Long cpuSystemTime;

  public static final String SERIALIZED_NAME_CPU_USER_TIME = "cpuUserTime";
  @SerializedName(SERIALIZED_NAME_CPU_USER_TIME)
  private Long cpuUserTime;

  public static final String SERIALIZED_NAME_CPU_TOTAL_TIME = "cpuTotalTime";
  @SerializedName(SERIALIZED_NAME_CPU_TOTAL_TIME)
  private Long cpuTotalTime;

  public static final String SERIALIZED_NAME_CPU_USAGE = "cpuUsage";
  @SerializedName(SERIALIZED_NAME_CPU_USAGE)
  private BigDecimal cpuUsage;

  public static final String SERIALIZED_NAME_USED_STORAGE = "usedStorage";
  @SerializedName(SERIALIZED_NAME_USED_STORAGE)
  private Long usedStorage;

  public static final String SERIALIZED_NAME_RESERVED_STORAGE = "reservedStorage";
  @SerializedName(SERIALIZED_NAME_RESERVED_STORAGE)
  private Long reservedStorage;

  public static final String SERIALIZED_NAME_MAX_STORAGE = "maxStorage";
  @SerializedName(SERIALIZED_NAME_MAX_STORAGE)
  private Long maxStorage;

  public static final String SERIALIZED_NAME_NET_TX_USAGE = "netTxUsage";
  @SerializedName(SERIALIZED_NAME_NET_TX_USAGE)
  private Long netTxUsage;

  public static final String SERIALIZED_NAME_NET_RX_USAGE = "netRxUsage";
  @SerializedName(SERIALIZED_NAME_NET_RX_USAGE)
  private Long netRxUsage;

  public static final String SERIALIZED_NAME_NETWORK_BANDWIDTH = "networkBandwidth";
  @SerializedName(SERIALIZED_NAME_NETWORK_BANDWIDTH)
  private Long networkBandwidth;


  public GuidanceVmwareSizingResourceStats ts(String ts) {
    
    this.ts = ts;
    return this;
  }

   /**
   * Get ts
   * @return ts
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public String getTs() {
    return ts;
  }


  public void setTs(String ts) {
    this.ts = ts;
  }


  public GuidanceVmwareSizingResourceStats freeMemory(Long freeMemory) {
    
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


  public GuidanceVmwareSizingResourceStats usedMemory(Long usedMemory) {
    
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


  public GuidanceVmwareSizingResourceStats freeSwap(Long freeSwap) {
    
    this.freeSwap = freeSwap;
    return this;
  }

   /**
   * Get freeSwap
   * @return freeSwap
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public Long getFreeSwap() {
    return freeSwap;
  }


  public void setFreeSwap(Long freeSwap) {
    this.freeSwap = freeSwap;
  }


  public GuidanceVmwareSizingResourceStats usedSwap(Long usedSwap) {
    
    this.usedSwap = usedSwap;
    return this;
  }

   /**
   * Get usedSwap
   * @return usedSwap
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public Long getUsedSwap() {
    return usedSwap;
  }


  public void setUsedSwap(Long usedSwap) {
    this.usedSwap = usedSwap;
  }


  public GuidanceVmwareSizingResourceStats cpuIdleTime(Long cpuIdleTime) {
    
    this.cpuIdleTime = cpuIdleTime;
    return this;
  }

   /**
   * Get cpuIdleTime
   * @return cpuIdleTime
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public Long getCpuIdleTime() {
    return cpuIdleTime;
  }


  public void setCpuIdleTime(Long cpuIdleTime) {
    this.cpuIdleTime = cpuIdleTime;
  }


  public GuidanceVmwareSizingResourceStats cpuSystemTime(Long cpuSystemTime) {
    
    this.cpuSystemTime = cpuSystemTime;
    return this;
  }

   /**
   * Get cpuSystemTime
   * @return cpuSystemTime
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public Long getCpuSystemTime() {
    return cpuSystemTime;
  }


  public void setCpuSystemTime(Long cpuSystemTime) {
    this.cpuSystemTime = cpuSystemTime;
  }


  public GuidanceVmwareSizingResourceStats cpuUserTime(Long cpuUserTime) {
    
    this.cpuUserTime = cpuUserTime;
    return this;
  }

   /**
   * Get cpuUserTime
   * @return cpuUserTime
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public Long getCpuUserTime() {
    return cpuUserTime;
  }


  public void setCpuUserTime(Long cpuUserTime) {
    this.cpuUserTime = cpuUserTime;
  }


  public GuidanceVmwareSizingResourceStats cpuTotalTime(Long cpuTotalTime) {
    
    this.cpuTotalTime = cpuTotalTime;
    return this;
  }

   /**
   * Get cpuTotalTime
   * @return cpuTotalTime
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public Long getCpuTotalTime() {
    return cpuTotalTime;
  }


  public void setCpuTotalTime(Long cpuTotalTime) {
    this.cpuTotalTime = cpuTotalTime;
  }


  public GuidanceVmwareSizingResourceStats cpuUsage(BigDecimal cpuUsage) {
    
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


  public GuidanceVmwareSizingResourceStats usedStorage(Long usedStorage) {
    
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


  public GuidanceVmwareSizingResourceStats reservedStorage(Long reservedStorage) {
    
    this.reservedStorage = reservedStorage;
    return this;
  }

   /**
   * Get reservedStorage
   * @return reservedStorage
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public Long getReservedStorage() {
    return reservedStorage;
  }


  public void setReservedStorage(Long reservedStorage) {
    this.reservedStorage = reservedStorage;
  }


  public GuidanceVmwareSizingResourceStats maxStorage(Long maxStorage) {
    
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


  public GuidanceVmwareSizingResourceStats netTxUsage(Long netTxUsage) {
    
    this.netTxUsage = netTxUsage;
    return this;
  }

   /**
   * Get netTxUsage
   * @return netTxUsage
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public Long getNetTxUsage() {
    return netTxUsage;
  }


  public void setNetTxUsage(Long netTxUsage) {
    this.netTxUsage = netTxUsage;
  }


  public GuidanceVmwareSizingResourceStats netRxUsage(Long netRxUsage) {
    
    this.netRxUsage = netRxUsage;
    return this;
  }

   /**
   * Get netRxUsage
   * @return netRxUsage
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public Long getNetRxUsage() {
    return netRxUsage;
  }


  public void setNetRxUsage(Long netRxUsage) {
    this.netRxUsage = netRxUsage;
  }


  public GuidanceVmwareSizingResourceStats networkBandwidth(Long networkBandwidth) {
    
    this.networkBandwidth = networkBandwidth;
    return this;
  }

   /**
   * Get networkBandwidth
   * @return networkBandwidth
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public Long getNetworkBandwidth() {
    return networkBandwidth;
  }


  public void setNetworkBandwidth(Long networkBandwidth) {
    this.networkBandwidth = networkBandwidth;
  }


  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (o == null || getClass() != o.getClass()) {
      return false;
    }
    GuidanceVmwareSizingResourceStats guidanceVmwareSizingResourceStats = (GuidanceVmwareSizingResourceStats) o;
    return Objects.equals(this.ts, guidanceVmwareSizingResourceStats.ts) &&
        Objects.equals(this.freeMemory, guidanceVmwareSizingResourceStats.freeMemory) &&
        Objects.equals(this.usedMemory, guidanceVmwareSizingResourceStats.usedMemory) &&
        Objects.equals(this.freeSwap, guidanceVmwareSizingResourceStats.freeSwap) &&
        Objects.equals(this.usedSwap, guidanceVmwareSizingResourceStats.usedSwap) &&
        Objects.equals(this.cpuIdleTime, guidanceVmwareSizingResourceStats.cpuIdleTime) &&
        Objects.equals(this.cpuSystemTime, guidanceVmwareSizingResourceStats.cpuSystemTime) &&
        Objects.equals(this.cpuUserTime, guidanceVmwareSizingResourceStats.cpuUserTime) &&
        Objects.equals(this.cpuTotalTime, guidanceVmwareSizingResourceStats.cpuTotalTime) &&
        Objects.equals(this.cpuUsage, guidanceVmwareSizingResourceStats.cpuUsage) &&
        Objects.equals(this.usedStorage, guidanceVmwareSizingResourceStats.usedStorage) &&
        Objects.equals(this.reservedStorage, guidanceVmwareSizingResourceStats.reservedStorage) &&
        Objects.equals(this.maxStorage, guidanceVmwareSizingResourceStats.maxStorage) &&
        Objects.equals(this.netTxUsage, guidanceVmwareSizingResourceStats.netTxUsage) &&
        Objects.equals(this.netRxUsage, guidanceVmwareSizingResourceStats.netRxUsage) &&
        Objects.equals(this.networkBandwidth, guidanceVmwareSizingResourceStats.networkBandwidth);
  }

  @Override
  public int hashCode() {
    return Objects.hash(ts, freeMemory, usedMemory, freeSwap, usedSwap, cpuIdleTime, cpuSystemTime, cpuUserTime, cpuTotalTime, cpuUsage, usedStorage, reservedStorage, maxStorage, netTxUsage, netRxUsage, networkBandwidth);
  }


  @Override
  public String toString() {
    StringBuilder sb = new StringBuilder();
    sb.append("class GuidanceVmwareSizingResourceStats {\n");
    sb.append("    ts: ").append(toIndentedString(ts)).append("\n");
    sb.append("    freeMemory: ").append(toIndentedString(freeMemory)).append("\n");
    sb.append("    usedMemory: ").append(toIndentedString(usedMemory)).append("\n");
    sb.append("    freeSwap: ").append(toIndentedString(freeSwap)).append("\n");
    sb.append("    usedSwap: ").append(toIndentedString(usedSwap)).append("\n");
    sb.append("    cpuIdleTime: ").append(toIndentedString(cpuIdleTime)).append("\n");
    sb.append("    cpuSystemTime: ").append(toIndentedString(cpuSystemTime)).append("\n");
    sb.append("    cpuUserTime: ").append(toIndentedString(cpuUserTime)).append("\n");
    sb.append("    cpuTotalTime: ").append(toIndentedString(cpuTotalTime)).append("\n");
    sb.append("    cpuUsage: ").append(toIndentedString(cpuUsage)).append("\n");
    sb.append("    usedStorage: ").append(toIndentedString(usedStorage)).append("\n");
    sb.append("    reservedStorage: ").append(toIndentedString(reservedStorage)).append("\n");
    sb.append("    maxStorage: ").append(toIndentedString(maxStorage)).append("\n");
    sb.append("    netTxUsage: ").append(toIndentedString(netTxUsage)).append("\n");
    sb.append("    netRxUsage: ").append(toIndentedString(netRxUsage)).append("\n");
    sb.append("    networkBandwidth: ").append(toIndentedString(networkBandwidth)).append("\n");
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
