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
 * GuidanceSettings
 */
@javax.annotation.Generated(value = "org.openapitools.codegen.languages.JavaClientCodegen", date = "2023-08-21T11:26:47.028Z[GMT]")
public class GuidanceSettings {
  public static final String SERIALIZED_NAME_CPU_AVG_CUTOFF_POWER = "cpuAvgCutoffPower";
  @SerializedName(SERIALIZED_NAME_CPU_AVG_CUTOFF_POWER)
  private Integer cpuAvgCutoffPower;

  public static final String SERIALIZED_NAME_CPU_MAX_CUTOFF_POWER = "cpuMaxCutoffPower";
  @SerializedName(SERIALIZED_NAME_CPU_MAX_CUTOFF_POWER)
  private Integer cpuMaxCutoffPower;

  public static final String SERIALIZED_NAME_NETWORK_CUTOFF_POWER = "networkCutoffPower";
  @SerializedName(SERIALIZED_NAME_NETWORK_CUTOFF_POWER)
  private Integer networkCutoffPower;

  public static final String SERIALIZED_NAME_CPU_UP_AVG_STANDARD_CUTOFF_RIGHT_SIZE = "cpuUpAvgStandardCutoffRightSize";
  @SerializedName(SERIALIZED_NAME_CPU_UP_AVG_STANDARD_CUTOFF_RIGHT_SIZE)
  private Integer cpuUpAvgStandardCutoffRightSize;

  public static final String SERIALIZED_NAME_CPU_UP_MAX_STANDARD_CUTOFF_RIGHT_SIZE = "cpuUpMaxStandardCutoffRightSize";
  @SerializedName(SERIALIZED_NAME_CPU_UP_MAX_STANDARD_CUTOFF_RIGHT_SIZE)
  private Integer cpuUpMaxStandardCutoffRightSize;

  public static final String SERIALIZED_NAME_MEMORY_UP_AVG_STANDARD_CUTOFF_RIGHT_SIZE = "memoryUpAvgStandardCutoffRightSize";
  @SerializedName(SERIALIZED_NAME_MEMORY_UP_AVG_STANDARD_CUTOFF_RIGHT_SIZE)
  private Integer memoryUpAvgStandardCutoffRightSize;

  public static final String SERIALIZED_NAME_MEMORY_DOWN_AVG_STANDARD_CUTOFF_RIGHT_SIZE = "memoryDownAvgStandardCutoffRightSize";
  @SerializedName(SERIALIZED_NAME_MEMORY_DOWN_AVG_STANDARD_CUTOFF_RIGHT_SIZE)
  private Integer memoryDownAvgStandardCutoffRightSize;

  public static final String SERIALIZED_NAME_MEMORY_DOWN_MAX_STANDARD_CUTOFF_RIGHT_SIZE = "memoryDownMaxStandardCutoffRightSize";
  @SerializedName(SERIALIZED_NAME_MEMORY_DOWN_MAX_STANDARD_CUTOFF_RIGHT_SIZE)
  private Integer memoryDownMaxStandardCutoffRightSize;


  public GuidanceSettings cpuAvgCutoffPower(Integer cpuAvgCutoffPower) {
    
    this.cpuAvgCutoffPower = cpuAvgCutoffPower;
    return this;
  }

   /**
   * Power Shutdown Average CPU (%). Lower limit for average CPU usage
   * @return cpuAvgCutoffPower
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(example = "75", value = "Power Shutdown Average CPU (%). Lower limit for average CPU usage")

  public Integer getCpuAvgCutoffPower() {
    return cpuAvgCutoffPower;
  }


  public void setCpuAvgCutoffPower(Integer cpuAvgCutoffPower) {
    this.cpuAvgCutoffPower = cpuAvgCutoffPower;
  }


  public GuidanceSettings cpuMaxCutoffPower(Integer cpuMaxCutoffPower) {
    
    this.cpuMaxCutoffPower = cpuMaxCutoffPower;
    return this;
  }

   /**
   * Power Shutdown Maximum CPU (%). Lower limit for peak CPU usage
   * @return cpuMaxCutoffPower
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(example = "500", value = "Power Shutdown Maximum CPU (%). Lower limit for peak CPU usage")

  public Integer getCpuMaxCutoffPower() {
    return cpuMaxCutoffPower;
  }


  public void setCpuMaxCutoffPower(Integer cpuMaxCutoffPower) {
    this.cpuMaxCutoffPower = cpuMaxCutoffPower;
  }


  public GuidanceSettings networkCutoffPower(Integer networkCutoffPower) {
    
    this.networkCutoffPower = networkCutoffPower;
    return this;
  }

   /**
   * Power Shutdown Network threshold (bytes). Lower limit for average network bandwidth
   * @return networkCutoffPower
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(example = "2000", value = "Power Shutdown Network threshold (bytes). Lower limit for average network bandwidth")

  public Integer getNetworkCutoffPower() {
    return networkCutoffPower;
  }


  public void setNetworkCutoffPower(Integer networkCutoffPower) {
    this.networkCutoffPower = networkCutoffPower;
  }


  public GuidanceSettings cpuUpAvgStandardCutoffRightSize(Integer cpuUpAvgStandardCutoffRightSize) {
    
    this.cpuUpAvgStandardCutoffRightSize = cpuUpAvgStandardCutoffRightSize;
    return this;
  }

   /**
   * CPU Up-size Average CPU (%). Upper limit for CPU usage
   * @return cpuUpAvgStandardCutoffRightSize
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(example = "50", value = "CPU Up-size Average CPU (%). Upper limit for CPU usage")

  public Integer getCpuUpAvgStandardCutoffRightSize() {
    return cpuUpAvgStandardCutoffRightSize;
  }


  public void setCpuUpAvgStandardCutoffRightSize(Integer cpuUpAvgStandardCutoffRightSize) {
    this.cpuUpAvgStandardCutoffRightSize = cpuUpAvgStandardCutoffRightSize;
  }


  public GuidanceSettings cpuUpMaxStandardCutoffRightSize(Integer cpuUpMaxStandardCutoffRightSize) {
    
    this.cpuUpMaxStandardCutoffRightSize = cpuUpMaxStandardCutoffRightSize;
    return this;
  }

   /**
   * CPU Up-size Maximum CPU (%). Upper limit for peak CPU usage
   * @return cpuUpMaxStandardCutoffRightSize
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(example = "99", value = "CPU Up-size Maximum CPU (%). Upper limit for peak CPU usage")

  public Integer getCpuUpMaxStandardCutoffRightSize() {
    return cpuUpMaxStandardCutoffRightSize;
  }


  public void setCpuUpMaxStandardCutoffRightSize(Integer cpuUpMaxStandardCutoffRightSize) {
    this.cpuUpMaxStandardCutoffRightSize = cpuUpMaxStandardCutoffRightSize;
  }


  public GuidanceSettings memoryUpAvgStandardCutoffRightSize(Integer memoryUpAvgStandardCutoffRightSize) {
    
    this.memoryUpAvgStandardCutoffRightSize = memoryUpAvgStandardCutoffRightSize;
    return this;
  }

   /**
   * Memory Up-size Minimum Free Memory (%). Lower limit for average free memory usage
   * @return memoryUpAvgStandardCutoffRightSize
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(example = "10", value = "Memory Up-size Minimum Free Memory (%). Lower limit for average free memory usage")

  public Integer getMemoryUpAvgStandardCutoffRightSize() {
    return memoryUpAvgStandardCutoffRightSize;
  }


  public void setMemoryUpAvgStandardCutoffRightSize(Integer memoryUpAvgStandardCutoffRightSize) {
    this.memoryUpAvgStandardCutoffRightSize = memoryUpAvgStandardCutoffRightSize;
  }


  public GuidanceSettings memoryDownAvgStandardCutoffRightSize(Integer memoryDownAvgStandardCutoffRightSize) {
    
    this.memoryDownAvgStandardCutoffRightSize = memoryDownAvgStandardCutoffRightSize;
    return this;
  }

   /**
   * Memory Down-size Maximum Free Memory (%). Upper limit for average free memory
   * @return memoryDownAvgStandardCutoffRightSize
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(example = "60", value = "Memory Down-size Maximum Free Memory (%). Upper limit for average free memory")

  public Integer getMemoryDownAvgStandardCutoffRightSize() {
    return memoryDownAvgStandardCutoffRightSize;
  }


  public void setMemoryDownAvgStandardCutoffRightSize(Integer memoryDownAvgStandardCutoffRightSize) {
    this.memoryDownAvgStandardCutoffRightSize = memoryDownAvgStandardCutoffRightSize;
  }


  public GuidanceSettings memoryDownMaxStandardCutoffRightSize(Integer memoryDownMaxStandardCutoffRightSize) {
    
    this.memoryDownMaxStandardCutoffRightSize = memoryDownMaxStandardCutoffRightSize;
    return this;
  }

   /**
   * Memory Down-size Maximum Free Memory (%). Upper limit for peak memory usage
   * @return memoryDownMaxStandardCutoffRightSize
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(example = "30", value = "Memory Down-size Maximum Free Memory (%). Upper limit for peak memory usage")

  public Integer getMemoryDownMaxStandardCutoffRightSize() {
    return memoryDownMaxStandardCutoffRightSize;
  }


  public void setMemoryDownMaxStandardCutoffRightSize(Integer memoryDownMaxStandardCutoffRightSize) {
    this.memoryDownMaxStandardCutoffRightSize = memoryDownMaxStandardCutoffRightSize;
  }


  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (o == null || getClass() != o.getClass()) {
      return false;
    }
    GuidanceSettings guidanceSettings = (GuidanceSettings) o;
    return Objects.equals(this.cpuAvgCutoffPower, guidanceSettings.cpuAvgCutoffPower) &&
        Objects.equals(this.cpuMaxCutoffPower, guidanceSettings.cpuMaxCutoffPower) &&
        Objects.equals(this.networkCutoffPower, guidanceSettings.networkCutoffPower) &&
        Objects.equals(this.cpuUpAvgStandardCutoffRightSize, guidanceSettings.cpuUpAvgStandardCutoffRightSize) &&
        Objects.equals(this.cpuUpMaxStandardCutoffRightSize, guidanceSettings.cpuUpMaxStandardCutoffRightSize) &&
        Objects.equals(this.memoryUpAvgStandardCutoffRightSize, guidanceSettings.memoryUpAvgStandardCutoffRightSize) &&
        Objects.equals(this.memoryDownAvgStandardCutoffRightSize, guidanceSettings.memoryDownAvgStandardCutoffRightSize) &&
        Objects.equals(this.memoryDownMaxStandardCutoffRightSize, guidanceSettings.memoryDownMaxStandardCutoffRightSize);
  }

  @Override
  public int hashCode() {
    return Objects.hash(cpuAvgCutoffPower, cpuMaxCutoffPower, networkCutoffPower, cpuUpAvgStandardCutoffRightSize, cpuUpMaxStandardCutoffRightSize, memoryUpAvgStandardCutoffRightSize, memoryDownAvgStandardCutoffRightSize, memoryDownMaxStandardCutoffRightSize);
  }


  @Override
  public String toString() {
    StringBuilder sb = new StringBuilder();
    sb.append("class GuidanceSettings {\n");
    sb.append("    cpuAvgCutoffPower: ").append(toIndentedString(cpuAvgCutoffPower)).append("\n");
    sb.append("    cpuMaxCutoffPower: ").append(toIndentedString(cpuMaxCutoffPower)).append("\n");
    sb.append("    networkCutoffPower: ").append(toIndentedString(networkCutoffPower)).append("\n");
    sb.append("    cpuUpAvgStandardCutoffRightSize: ").append(toIndentedString(cpuUpAvgStandardCutoffRightSize)).append("\n");
    sb.append("    cpuUpMaxStandardCutoffRightSize: ").append(toIndentedString(cpuUpMaxStandardCutoffRightSize)).append("\n");
    sb.append("    memoryUpAvgStandardCutoffRightSize: ").append(toIndentedString(memoryUpAvgStandardCutoffRightSize)).append("\n");
    sb.append("    memoryDownAvgStandardCutoffRightSize: ").append(toIndentedString(memoryDownAvgStandardCutoffRightSize)).append("\n");
    sb.append("    memoryDownMaxStandardCutoffRightSize: ").append(toIndentedString(memoryDownMaxStandardCutoffRightSize)).append("\n");
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

