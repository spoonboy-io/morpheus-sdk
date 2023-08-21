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
 * GuidanceVmwareSizingPlanBeforeActionConfigRanges
 */
@javax.annotation.Generated(value = "org.openapitools.codegen.languages.JavaClientCodegen", date = "2023-08-21T11:26:47.028Z[GMT]")
public class GuidanceVmwareSizingPlanBeforeActionConfigRanges {
  public static final String SERIALIZED_NAME_MIN_STORAGE = "minStorage";
  @SerializedName(SERIALIZED_NAME_MIN_STORAGE)
  private String minStorage;

  public static final String SERIALIZED_NAME_MAX_STORAGE = "maxStorage";
  @SerializedName(SERIALIZED_NAME_MAX_STORAGE)
  private String maxStorage;

  public static final String SERIALIZED_NAME_MIN_MEMORY = "minMemory";
  @SerializedName(SERIALIZED_NAME_MIN_MEMORY)
  private String minMemory;

  public static final String SERIALIZED_NAME_MAX_MEMORY = "maxMemory";
  @SerializedName(SERIALIZED_NAME_MAX_MEMORY)
  private String maxMemory;

  public static final String SERIALIZED_NAME_MIN_CORES = "minCores";
  @SerializedName(SERIALIZED_NAME_MIN_CORES)
  private String minCores;

  public static final String SERIALIZED_NAME_MAX_CORES = "maxCores";
  @SerializedName(SERIALIZED_NAME_MAX_CORES)
  private String maxCores;


  public GuidanceVmwareSizingPlanBeforeActionConfigRanges minStorage(String minStorage) {
    
    this.minStorage = minStorage;
    return this;
  }

   /**
   * Get minStorage
   * @return minStorage
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public String getMinStorage() {
    return minStorage;
  }


  public void setMinStorage(String minStorage) {
    this.minStorage = minStorage;
  }


  public GuidanceVmwareSizingPlanBeforeActionConfigRanges maxStorage(String maxStorage) {
    
    this.maxStorage = maxStorage;
    return this;
  }

   /**
   * Get maxStorage
   * @return maxStorage
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public String getMaxStorage() {
    return maxStorage;
  }


  public void setMaxStorage(String maxStorage) {
    this.maxStorage = maxStorage;
  }


  public GuidanceVmwareSizingPlanBeforeActionConfigRanges minMemory(String minMemory) {
    
    this.minMemory = minMemory;
    return this;
  }

   /**
   * Get minMemory
   * @return minMemory
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public String getMinMemory() {
    return minMemory;
  }


  public void setMinMemory(String minMemory) {
    this.minMemory = minMemory;
  }


  public GuidanceVmwareSizingPlanBeforeActionConfigRanges maxMemory(String maxMemory) {
    
    this.maxMemory = maxMemory;
    return this;
  }

   /**
   * Get maxMemory
   * @return maxMemory
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public String getMaxMemory() {
    return maxMemory;
  }


  public void setMaxMemory(String maxMemory) {
    this.maxMemory = maxMemory;
  }


  public GuidanceVmwareSizingPlanBeforeActionConfigRanges minCores(String minCores) {
    
    this.minCores = minCores;
    return this;
  }

   /**
   * Get minCores
   * @return minCores
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public String getMinCores() {
    return minCores;
  }


  public void setMinCores(String minCores) {
    this.minCores = minCores;
  }


  public GuidanceVmwareSizingPlanBeforeActionConfigRanges maxCores(String maxCores) {
    
    this.maxCores = maxCores;
    return this;
  }

   /**
   * Get maxCores
   * @return maxCores
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public String getMaxCores() {
    return maxCores;
  }


  public void setMaxCores(String maxCores) {
    this.maxCores = maxCores;
  }


  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (o == null || getClass() != o.getClass()) {
      return false;
    }
    GuidanceVmwareSizingPlanBeforeActionConfigRanges guidanceVmwareSizingPlanBeforeActionConfigRanges = (GuidanceVmwareSizingPlanBeforeActionConfigRanges) o;
    return Objects.equals(this.minStorage, guidanceVmwareSizingPlanBeforeActionConfigRanges.minStorage) &&
        Objects.equals(this.maxStorage, guidanceVmwareSizingPlanBeforeActionConfigRanges.maxStorage) &&
        Objects.equals(this.minMemory, guidanceVmwareSizingPlanBeforeActionConfigRanges.minMemory) &&
        Objects.equals(this.maxMemory, guidanceVmwareSizingPlanBeforeActionConfigRanges.maxMemory) &&
        Objects.equals(this.minCores, guidanceVmwareSizingPlanBeforeActionConfigRanges.minCores) &&
        Objects.equals(this.maxCores, guidanceVmwareSizingPlanBeforeActionConfigRanges.maxCores);
  }

  @Override
  public int hashCode() {
    return Objects.hash(minStorage, maxStorage, minMemory, maxMemory, minCores, maxCores);
  }


  @Override
  public String toString() {
    StringBuilder sb = new StringBuilder();
    sb.append("class GuidanceVmwareSizingPlanBeforeActionConfigRanges {\n");
    sb.append("    minStorage: ").append(toIndentedString(minStorage)).append("\n");
    sb.append("    maxStorage: ").append(toIndentedString(maxStorage)).append("\n");
    sb.append("    minMemory: ").append(toIndentedString(minMemory)).append("\n");
    sb.append("    maxMemory: ").append(toIndentedString(maxMemory)).append("\n");
    sb.append("    minCores: ").append(toIndentedString(minCores)).append("\n");
    sb.append("    maxCores: ").append(toIndentedString(maxCores)).append("\n");
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
