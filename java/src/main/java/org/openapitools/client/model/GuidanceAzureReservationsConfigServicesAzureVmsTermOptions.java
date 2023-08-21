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
import org.openapitools.client.model.GuidanceAzureReservationsConfigServicesAzureVmsSummary;

/**
 * GuidanceAzureReservationsConfigServicesAzureVmsTermOptions
 */
@javax.annotation.Generated(value = "org.openapitools.codegen.languages.JavaClientCodegen", date = "2023-08-21T11:26:47.028Z[GMT]")
public class GuidanceAzureReservationsConfigServicesAzureVmsTermOptions {
  public static final String SERIALIZED_NAME_CODE = "code";
  @SerializedName(SERIALIZED_NAME_CODE)
  private String code;

  public static final String SERIALIZED_NAME_NAME = "name";
  @SerializedName(SERIALIZED_NAME_NAME)
  private String name;

  public static final String SERIALIZED_NAME_DETAIL_LIST = "detailList";
  @SerializedName(SERIALIZED_NAME_DETAIL_LIST)
  private List<GuidanceAzureReservationsConfigDetailList> detailList = null;

  public static final String SERIALIZED_NAME_SUMMARY = "summary";
  @SerializedName(SERIALIZED_NAME_SUMMARY)
  private GuidanceAzureReservationsConfigServicesAzureVmsSummary summary;


  public GuidanceAzureReservationsConfigServicesAzureVmsTermOptions code(String code) {
    
    this.code = code;
    return this;
  }

   /**
   * Get code
   * @return code
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public String getCode() {
    return code;
  }


  public void setCode(String code) {
    this.code = code;
  }


  public GuidanceAzureReservationsConfigServicesAzureVmsTermOptions name(String name) {
    
    this.name = name;
    return this;
  }

   /**
   * Get name
   * @return name
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public String getName() {
    return name;
  }


  public void setName(String name) {
    this.name = name;
  }


  public GuidanceAzureReservationsConfigServicesAzureVmsTermOptions detailList(List<GuidanceAzureReservationsConfigDetailList> detailList) {
    
    this.detailList = detailList;
    return this;
  }

  public GuidanceAzureReservationsConfigServicesAzureVmsTermOptions addDetailListItem(GuidanceAzureReservationsConfigDetailList detailListItem) {
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


  public GuidanceAzureReservationsConfigServicesAzureVmsTermOptions summary(GuidanceAzureReservationsConfigServicesAzureVmsSummary summary) {
    
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
    GuidanceAzureReservationsConfigServicesAzureVmsTermOptions guidanceAzureReservationsConfigServicesAzureVmsTermOptions = (GuidanceAzureReservationsConfigServicesAzureVmsTermOptions) o;
    return Objects.equals(this.code, guidanceAzureReservationsConfigServicesAzureVmsTermOptions.code) &&
        Objects.equals(this.name, guidanceAzureReservationsConfigServicesAzureVmsTermOptions.name) &&
        Objects.equals(this.detailList, guidanceAzureReservationsConfigServicesAzureVmsTermOptions.detailList) &&
        Objects.equals(this.summary, guidanceAzureReservationsConfigServicesAzureVmsTermOptions.summary);
  }

  @Override
  public int hashCode() {
    return Objects.hash(code, name, detailList, summary);
  }


  @Override
  public String toString() {
    StringBuilder sb = new StringBuilder();
    sb.append("class GuidanceAzureReservationsConfigServicesAzureVmsTermOptions {\n");
    sb.append("    code: ").append(toIndentedString(code)).append("\n");
    sb.append("    name: ").append(toIndentedString(name)).append("\n");
    sb.append("    detailList: ").append(toIndentedString(detailList)).append("\n");
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

