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
import org.openapitools.client.model.GuidanceStatsSeverity;
import org.openapitools.client.model.GuidanceStatsType;
import org.openapitools.client.model.GuidanceVmwareSizingSavings;

/**
 * GuidanceStats
 */
@javax.annotation.Generated(value = "org.openapitools.codegen.languages.JavaClientCodegen", date = "2023-08-21T11:26:47.028Z[GMT]")
public class GuidanceStats {
  public static final String SERIALIZED_NAME_TOTAL = "total";
  @SerializedName(SERIALIZED_NAME_TOTAL)
  private Long total;

  public static final String SERIALIZED_NAME_SAVINGS = "savings";
  @SerializedName(SERIALIZED_NAME_SAVINGS)
  private GuidanceVmwareSizingSavings savings;

  public static final String SERIALIZED_NAME_SEVERITY = "severity";
  @SerializedName(SERIALIZED_NAME_SEVERITY)
  private GuidanceStatsSeverity severity;

  public static final String SERIALIZED_NAME_TYPE = "type";
  @SerializedName(SERIALIZED_NAME_TYPE)
  private GuidanceStatsType type;


  public GuidanceStats total(Long total) {
    
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


  public GuidanceStats savings(GuidanceVmwareSizingSavings savings) {
    
    this.savings = savings;
    return this;
  }

   /**
   * Get savings
   * @return savings
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public GuidanceVmwareSizingSavings getSavings() {
    return savings;
  }


  public void setSavings(GuidanceVmwareSizingSavings savings) {
    this.savings = savings;
  }


  public GuidanceStats severity(GuidanceStatsSeverity severity) {
    
    this.severity = severity;
    return this;
  }

   /**
   * Get severity
   * @return severity
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public GuidanceStatsSeverity getSeverity() {
    return severity;
  }


  public void setSeverity(GuidanceStatsSeverity severity) {
    this.severity = severity;
  }


  public GuidanceStats type(GuidanceStatsType type) {
    
    this.type = type;
    return this;
  }

   /**
   * Get type
   * @return type
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public GuidanceStatsType getType() {
    return type;
  }


  public void setType(GuidanceStatsType type) {
    this.type = type;
  }


  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (o == null || getClass() != o.getClass()) {
      return false;
    }
    GuidanceStats guidanceStats = (GuidanceStats) o;
    return Objects.equals(this.total, guidanceStats.total) &&
        Objects.equals(this.savings, guidanceStats.savings) &&
        Objects.equals(this.severity, guidanceStats.severity) &&
        Objects.equals(this.type, guidanceStats.type);
  }

  @Override
  public int hashCode() {
    return Objects.hash(total, savings, severity, type);
  }


  @Override
  public String toString() {
    StringBuilder sb = new StringBuilder();
    sb.append("class GuidanceStats {\n");
    sb.append("    total: ").append(toIndentedString(total)).append("\n");
    sb.append("    savings: ").append(toIndentedString(savings)).append("\n");
    sb.append("    severity: ").append(toIndentedString(severity)).append("\n");
    sb.append("    type: ").append(toIndentedString(type)).append("\n");
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
