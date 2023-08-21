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
 * InlineObject47
 */
@javax.annotation.Generated(value = "org.openapitools.codegen.languages.JavaClientCodegen", date = "2023-08-21T11:26:47.028Z[GMT]")
public class InlineObject47 {
  public static final String SERIALIZED_NAME_MODE = "mode";
  @SerializedName(SERIALIZED_NAME_MODE)
  private String mode;

  public static final String SERIALIZED_NAME_REBUILD = "rebuild";
  @SerializedName(SERIALIZED_NAME_REBUILD)
  private String rebuild;

  public static final String SERIALIZED_NAME_PERIOD = "period";
  @SerializedName(SERIALIZED_NAME_PERIOD)
  private String period;


  public InlineObject47 mode(String mode) {
    
    this.mode = mode;
    return this;
  }

   /**
   * Refresh Mode. Run the &#x60;daily&#x60; or &#x60;costing&#x60; job instead of the default &#x60;hourly&#x60; refresh job.
   * @return mode
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "Refresh Mode. Run the `daily` or `costing` job instead of the default `hourly` refresh job.")

  public String getMode() {
    return mode;
  }


  public void setMode(String mode) {
    this.mode = mode;
  }


  public InlineObject47 rebuild(String rebuild) {
    
    this.rebuild = rebuild;
    return this;
  }

   /**
   * Rebuild. Pass &#x60;true&#x60; to purge existing invoices for the period before refreshing. Only applies to mode &#x60;costing&#x60;.
   * @return rebuild
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "Rebuild. Pass `true` to purge existing invoices for the period before refreshing. Only applies to mode `costing`.")

  public String getRebuild() {
    return rebuild;
  }


  public void setRebuild(String rebuild) {
    this.rebuild = rebuild;
  }


  public InlineObject47 period(String period) {
    
    this.period = period;
    return this;
  }

   /**
   * Period. Invoice billing period to refresh in the format &#x60;YYYYMM&#x60;. The default period is the current month. Only applies to mode &#x60;costing&#x60;.
   * @return period
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "Period. Invoice billing period to refresh in the format `YYYYMM`. The default period is the current month. Only applies to mode `costing`.")

  public String getPeriod() {
    return period;
  }


  public void setPeriod(String period) {
    this.period = period;
  }


  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (o == null || getClass() != o.getClass()) {
      return false;
    }
    InlineObject47 inlineObject47 = (InlineObject47) o;
    return Objects.equals(this.mode, inlineObject47.mode) &&
        Objects.equals(this.rebuild, inlineObject47.rebuild) &&
        Objects.equals(this.period, inlineObject47.period);
  }

  @Override
  public int hashCode() {
    return Objects.hash(mode, rebuild, period);
  }


  @Override
  public String toString() {
    StringBuilder sb = new StringBuilder();
    sb.append("class InlineObject47 {\n");
    sb.append("    mode: ").append(toIndentedString(mode)).append("\n");
    sb.append("    rebuild: ").append(toIndentedString(rebuild)).append("\n");
    sb.append("    period: ").append(toIndentedString(period)).append("\n");
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
