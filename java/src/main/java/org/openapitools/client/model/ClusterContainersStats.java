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
import org.threeten.bp.OffsetDateTime;

/**
 * ClusterContainersStats
 */
@javax.annotation.Generated(value = "org.openapitools.codegen.languages.JavaClientCodegen", date = "2023-08-21T11:26:47.028Z[GMT]")
public class ClusterContainersStats {
  public static final String SERIALIZED_NAME_TS = "ts";
  @SerializedName(SERIALIZED_NAME_TS)
  private OffsetDateTime ts;

  public static final String SERIALIZED_NAME_RUNNING = "running";
  @SerializedName(SERIALIZED_NAME_RUNNING)
  private Boolean running;

  public static final String SERIALIZED_NAME_USER_CPU_USAGE = "userCpuUsage";
  @SerializedName(SERIALIZED_NAME_USER_CPU_USAGE)
  private BigDecimal userCpuUsage;

  public static final String SERIALIZED_NAME_SYSTEM_CPU_USAGE = "systemCpuUsage";
  @SerializedName(SERIALIZED_NAME_SYSTEM_CPU_USAGE)
  private BigDecimal systemCpuUsage;

  public static final String SERIALIZED_NAME_USED_MEMORY = "usedMemory";
  @SerializedName(SERIALIZED_NAME_USED_MEMORY)
  private Long usedMemory;

  public static final String SERIALIZED_NAME_MAX_MEMORY = "maxMemory";
  @SerializedName(SERIALIZED_NAME_MAX_MEMORY)
  private Long maxMemory;

  public static final String SERIALIZED_NAME_CACHE_MEMORY = "cacheMemory";
  @SerializedName(SERIALIZED_NAME_CACHE_MEMORY)
  private Long cacheMemory;

  public static final String SERIALIZED_NAME_MAX_STORAGE = "maxStorage";
  @SerializedName(SERIALIZED_NAME_MAX_STORAGE)
  private Long maxStorage;

  public static final String SERIALIZED_NAME_USED_STORAGE = "usedStorage";
  @SerializedName(SERIALIZED_NAME_USED_STORAGE)
  private Long usedStorage;

  public static final String SERIALIZED_NAME_READ_I_O_P_S = "readIOPS";
  @SerializedName(SERIALIZED_NAME_READ_I_O_P_S)
  private Long readIOPS;

  public static final String SERIALIZED_NAME_WRITE_I_O_P_S = "writeIOPS";
  @SerializedName(SERIALIZED_NAME_WRITE_I_O_P_S)
  private Long writeIOPS;

  public static final String SERIALIZED_NAME_TOTAL_I_O_P_S = "totalIOPS";
  @SerializedName(SERIALIZED_NAME_TOTAL_I_O_P_S)
  private Long totalIOPS;

  public static final String SERIALIZED_NAME_NET_TX_USAGE = "netTxUsage";
  @SerializedName(SERIALIZED_NAME_NET_TX_USAGE)
  private Long netTxUsage;

  public static final String SERIALIZED_NAME_NET_RX_USAGE = "netRxUsage";
  @SerializedName(SERIALIZED_NAME_NET_RX_USAGE)
  private Long netRxUsage;


  public ClusterContainersStats ts(OffsetDateTime ts) {
    
    this.ts = ts;
    return this;
  }

   /**
   * Get ts
   * @return ts
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public OffsetDateTime getTs() {
    return ts;
  }


  public void setTs(OffsetDateTime ts) {
    this.ts = ts;
  }


  public ClusterContainersStats running(Boolean running) {
    
    this.running = running;
    return this;
  }

   /**
   * Get running
   * @return running
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public Boolean getRunning() {
    return running;
  }


  public void setRunning(Boolean running) {
    this.running = running;
  }


  public ClusterContainersStats userCpuUsage(BigDecimal userCpuUsage) {
    
    this.userCpuUsage = userCpuUsage;
    return this;
  }

   /**
   * Get userCpuUsage
   * @return userCpuUsage
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public BigDecimal getUserCpuUsage() {
    return userCpuUsage;
  }


  public void setUserCpuUsage(BigDecimal userCpuUsage) {
    this.userCpuUsage = userCpuUsage;
  }


  public ClusterContainersStats systemCpuUsage(BigDecimal systemCpuUsage) {
    
    this.systemCpuUsage = systemCpuUsage;
    return this;
  }

   /**
   * Get systemCpuUsage
   * @return systemCpuUsage
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public BigDecimal getSystemCpuUsage() {
    return systemCpuUsage;
  }


  public void setSystemCpuUsage(BigDecimal systemCpuUsage) {
    this.systemCpuUsage = systemCpuUsage;
  }


  public ClusterContainersStats usedMemory(Long usedMemory) {
    
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


  public ClusterContainersStats maxMemory(Long maxMemory) {
    
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


  public ClusterContainersStats cacheMemory(Long cacheMemory) {
    
    this.cacheMemory = cacheMemory;
    return this;
  }

   /**
   * Get cacheMemory
   * @return cacheMemory
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public Long getCacheMemory() {
    return cacheMemory;
  }


  public void setCacheMemory(Long cacheMemory) {
    this.cacheMemory = cacheMemory;
  }


  public ClusterContainersStats maxStorage(Long maxStorage) {
    
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


  public ClusterContainersStats usedStorage(Long usedStorage) {
    
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


  public ClusterContainersStats readIOPS(Long readIOPS) {
    
    this.readIOPS = readIOPS;
    return this;
  }

   /**
   * Get readIOPS
   * @return readIOPS
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public Long getReadIOPS() {
    return readIOPS;
  }


  public void setReadIOPS(Long readIOPS) {
    this.readIOPS = readIOPS;
  }


  public ClusterContainersStats writeIOPS(Long writeIOPS) {
    
    this.writeIOPS = writeIOPS;
    return this;
  }

   /**
   * Get writeIOPS
   * @return writeIOPS
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public Long getWriteIOPS() {
    return writeIOPS;
  }


  public void setWriteIOPS(Long writeIOPS) {
    this.writeIOPS = writeIOPS;
  }


  public ClusterContainersStats totalIOPS(Long totalIOPS) {
    
    this.totalIOPS = totalIOPS;
    return this;
  }

   /**
   * Get totalIOPS
   * @return totalIOPS
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public Long getTotalIOPS() {
    return totalIOPS;
  }


  public void setTotalIOPS(Long totalIOPS) {
    this.totalIOPS = totalIOPS;
  }


  public ClusterContainersStats netTxUsage(Long netTxUsage) {
    
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


  public ClusterContainersStats netRxUsage(Long netRxUsage) {
    
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


  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (o == null || getClass() != o.getClass()) {
      return false;
    }
    ClusterContainersStats clusterContainersStats = (ClusterContainersStats) o;
    return Objects.equals(this.ts, clusterContainersStats.ts) &&
        Objects.equals(this.running, clusterContainersStats.running) &&
        Objects.equals(this.userCpuUsage, clusterContainersStats.userCpuUsage) &&
        Objects.equals(this.systemCpuUsage, clusterContainersStats.systemCpuUsage) &&
        Objects.equals(this.usedMemory, clusterContainersStats.usedMemory) &&
        Objects.equals(this.maxMemory, clusterContainersStats.maxMemory) &&
        Objects.equals(this.cacheMemory, clusterContainersStats.cacheMemory) &&
        Objects.equals(this.maxStorage, clusterContainersStats.maxStorage) &&
        Objects.equals(this.usedStorage, clusterContainersStats.usedStorage) &&
        Objects.equals(this.readIOPS, clusterContainersStats.readIOPS) &&
        Objects.equals(this.writeIOPS, clusterContainersStats.writeIOPS) &&
        Objects.equals(this.totalIOPS, clusterContainersStats.totalIOPS) &&
        Objects.equals(this.netTxUsage, clusterContainersStats.netTxUsage) &&
        Objects.equals(this.netRxUsage, clusterContainersStats.netRxUsage);
  }

  @Override
  public int hashCode() {
    return Objects.hash(ts, running, userCpuUsage, systemCpuUsage, usedMemory, maxMemory, cacheMemory, maxStorage, usedStorage, readIOPS, writeIOPS, totalIOPS, netTxUsage, netRxUsage);
  }


  @Override
  public String toString() {
    StringBuilder sb = new StringBuilder();
    sb.append("class ClusterContainersStats {\n");
    sb.append("    ts: ").append(toIndentedString(ts)).append("\n");
    sb.append("    running: ").append(toIndentedString(running)).append("\n");
    sb.append("    userCpuUsage: ").append(toIndentedString(userCpuUsage)).append("\n");
    sb.append("    systemCpuUsage: ").append(toIndentedString(systemCpuUsage)).append("\n");
    sb.append("    usedMemory: ").append(toIndentedString(usedMemory)).append("\n");
    sb.append("    maxMemory: ").append(toIndentedString(maxMemory)).append("\n");
    sb.append("    cacheMemory: ").append(toIndentedString(cacheMemory)).append("\n");
    sb.append("    maxStorage: ").append(toIndentedString(maxStorage)).append("\n");
    sb.append("    usedStorage: ").append(toIndentedString(usedStorage)).append("\n");
    sb.append("    readIOPS: ").append(toIndentedString(readIOPS)).append("\n");
    sb.append("    writeIOPS: ").append(toIndentedString(writeIOPS)).append("\n");
    sb.append("    totalIOPS: ").append(toIndentedString(totalIOPS)).append("\n");
    sb.append("    netTxUsage: ").append(toIndentedString(netTxUsage)).append("\n");
    sb.append("    netRxUsage: ").append(toIndentedString(netRxUsage)).append("\n");
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
