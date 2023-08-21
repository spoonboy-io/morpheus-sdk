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
import org.openapitools.client.model.GuidanceAzureReservationsConfigDetailList;
import org.openapitools.client.model.GuidanceAzureReservationsConfigServices;
import org.openapitools.client.model.GuidanceAzureReservationsConfigServicesAzureVmsSummary;

/**
 * GuidanceAzureReservationsConfig
 */
@javax.annotation.Generated(value = "org.openapitools.codegen.languages.JavaClientCodegen", date = "2023-08-21T11:26:47.028Z[GMT]")
public class GuidanceAzureReservationsConfig {
  public static final String SERIALIZED_NAME_SUCCESS = "success";
  @SerializedName(SERIALIZED_NAME_SUCCESS)
  private Boolean success;

  public static final String SERIALIZED_NAME_DETAIL_LIST = "detailList";
  @SerializedName(SERIALIZED_NAME_DETAIL_LIST)
  private List<GuidanceAzureReservationsConfigDetailList> detailList = null;

  public static final String SERIALIZED_NAME_SERVICES = "services";
  @SerializedName(SERIALIZED_NAME_SERVICES)
  private GuidanceAzureReservationsConfigServices services;

  public static final String SERIALIZED_NAME_SUMMARY = "summary";
  @SerializedName(SERIALIZED_NAME_SUMMARY)
  private GuidanceAzureReservationsConfigServicesAzureVmsSummary summary;


  public GuidanceAzureReservationsConfig success(Boolean success) {
    
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


  public GuidanceAzureReservationsConfig detailList(List<GuidanceAzureReservationsConfigDetailList> detailList) {
    
    this.detailList = detailList;
    return this;
  }

  public GuidanceAzureReservationsConfig addDetailListItem(GuidanceAzureReservationsConfigDetailList detailListItem) {
    if (this.detailList == null) {
      this.detailList = new ArrayList<GuidanceAzureReservationsConfigDetailList>();
    }
    this.detailList.add(detailListItem);
    return this;
  }

   /**
   * Get detailList
   * @return detailList
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public List<GuidanceAzureReservationsConfigDetailList> getDetailList() {
    return detailList;
  }


  public void setDetailList(List<GuidanceAzureReservationsConfigDetailList> detailList) {
    this.detailList = detailList;
  }


  public GuidanceAzureReservationsConfig services(GuidanceAzureReservationsConfigServices services) {
    
    this.services = services;
    return this;
  }

   /**
   * Get services
   * @return services
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public GuidanceAzureReservationsConfigServices getServices() {
    return services;
  }


  public void setServices(GuidanceAzureReservationsConfigServices services) {
    this.services = services;
  }


  public GuidanceAzureReservationsConfig summary(GuidanceAzureReservationsConfigServicesAzureVmsSummary summary) {
    
    this.summary = summary;
    return this;
  }

   /**
   * Get summary
   * @return summary
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public GuidanceAzureReservationsConfigServicesAzureVmsSummary getSummary() {
    return summary;
  }


  public void setSummary(GuidanceAzureReservationsConfigServicesAzureVmsSummary summary) {
    this.summary = summary;
  }


  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (o == null || getClass() != o.getClass()) {
      return false;
    }
    GuidanceAzureReservationsConfig guidanceAzureReservationsConfig = (GuidanceAzureReservationsConfig) o;
    return Objects.equals(this.success, guidanceAzureReservationsConfig.success) &&
        Objects.equals(this.detailList, guidanceAzureReservationsConfig.detailList) &&
        Objects.equals(this.services, guidanceAzureReservationsConfig.services) &&
        Objects.equals(this.summary, guidanceAzureReservationsConfig.summary);
  }

  @Override
  public int hashCode() {
    return Objects.hash(success, detailList, services, summary);
  }


  @Override
  public String toString() {
    StringBuilder sb = new StringBuilder();
    sb.append("class GuidanceAzureReservationsConfig {\n");
    sb.append("    success: ").append(toIndentedString(success)).append("\n");
    sb.append("    detailList: ").append(toIndentedString(detailList)).append("\n");
    sb.append("    services: ").append(toIndentedString(services)).append("\n");
    sb.append("    summary: ").append(toIndentedString(summary)).append("\n");
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

