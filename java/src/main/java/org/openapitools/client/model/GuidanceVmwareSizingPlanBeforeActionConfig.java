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
import org.openapitools.client.model.GuidanceVmwareSizingPlanBeforeActionConfigRanges;

/**
 * GuidanceVmwareSizingPlanBeforeActionConfig
 */
@javax.annotation.Generated(value = "org.openapitools.codegen.languages.JavaClientCodegen", date = "2023-08-21T11:26:47.028Z[GMT]")
public class GuidanceVmwareSizingPlanBeforeActionConfig {
  public static final String SERIALIZED_NAME_STORAGE_SIZE_TYPE = "storageSizeType";
  @SerializedName(SERIALIZED_NAME_STORAGE_SIZE_TYPE)
  private String storageSizeType;

  public static final String SERIALIZED_NAME_MEMORY_SIZE_TYPE = "memorySizeType";
  @SerializedName(SERIALIZED_NAME_MEMORY_SIZE_TYPE)
  private String memorySizeType;

  public static final String SERIALIZED_NAME_RANGES = "ranges";
  @SerializedName(SERIALIZED_NAME_RANGES)
  private GuidanceVmwareSizingPlanBeforeActionConfigRanges ranges;


  public GuidanceVmwareSizingPlanBeforeActionConfig storageSizeType(String storageSizeType) {
    
    this.storageSizeType = storageSizeType;
    return this;
  }

   /**
   * Get storageSizeType
   * @return storageSizeType
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public String getStorageSizeType() {
    return storageSizeType;
  }


  public void setStorageSizeType(String storageSizeType) {
    this.storageSizeType = storageSizeType;
  }


  public GuidanceVmwareSizingPlanBeforeActionConfig memorySizeType(String memorySizeType) {
    
    this.memorySizeType = memorySizeType;
    return this;
  }

   /**
   * Get memorySizeType
   * @return memorySizeType
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public String getMemorySizeType() {
    return memorySizeType;
  }


  public void setMemorySizeType(String memorySizeType) {
    this.memorySizeType = memorySizeType;
  }


  public GuidanceVmwareSizingPlanBeforeActionConfig ranges(GuidanceVmwareSizingPlanBeforeActionConfigRanges ranges) {
    
    this.ranges = ranges;
    return this;
  }

   /**
   * Get ranges
   * @return ranges
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public GuidanceVmwareSizingPlanBeforeActionConfigRanges getRanges() {
    return ranges;
  }


  public void setRanges(GuidanceVmwareSizingPlanBeforeActionConfigRanges ranges) {
    this.ranges = ranges;
  }


  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (o == null || getClass() != o.getClass()) {
      return false;
    }
    GuidanceVmwareSizingPlanBeforeActionConfig guidanceVmwareSizingPlanBeforeActionConfig = (GuidanceVmwareSizingPlanBeforeActionConfig) o;
    return Objects.equals(this.storageSizeType, guidanceVmwareSizingPlanBeforeActionConfig.storageSizeType) &&
        Objects.equals(this.memorySizeType, guidanceVmwareSizingPlanBeforeActionConfig.memorySizeType) &&
        Objects.equals(this.ranges, guidanceVmwareSizingPlanBeforeActionConfig.ranges);
  }

  @Override
  public int hashCode() {
    return Objects.hash(storageSizeType, memorySizeType, ranges);
  }


  @Override
  public String toString() {
    StringBuilder sb = new StringBuilder();
    sb.append("class GuidanceVmwareSizingPlanBeforeActionConfig {\n");
    sb.append("    storageSizeType: ").append(toIndentedString(storageSizeType)).append("\n");
    sb.append("    memorySizeType: ").append(toIndentedString(memorySizeType)).append("\n");
    sb.append("    ranges: ").append(toIndentedString(ranges)).append("\n");
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
