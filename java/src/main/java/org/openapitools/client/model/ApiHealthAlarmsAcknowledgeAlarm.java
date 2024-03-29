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

/**
 * ApiHealthAlarmsAcknowledgeAlarm
 */
@javax.annotation.Generated(value = "org.openapitools.codegen.languages.JavaClientCodegen", date = "2023-08-21T11:26:47.028Z[GMT]")
public class ApiHealthAlarmsAcknowledgeAlarm {
  public static final String SERIALIZED_NAME_ACKNOWLEDGED = "acknowledged";
  @SerializedName(SERIALIZED_NAME_ACKNOWLEDGED)
  private Boolean acknowledged;

  public static final String SERIALIZED_NAME_IDS = "ids";
  @SerializedName(SERIALIZED_NAME_IDS)
  private List<Long> ids = null;

  public static final String SERIALIZED_NAME_ALL = "all";
  @SerializedName(SERIALIZED_NAME_ALL)
  private Boolean all = false;


  public ApiHealthAlarmsAcknowledgeAlarm acknowledged(Boolean acknowledged) {
    
    this.acknowledged = acknowledged;
    return this;
  }

   /**
   * Pass &#x60;true&#x60; to ackowledge an alarm, or pass &#x60;false&#x60; to unacknowledge it.
   * @return acknowledged
  **/
  @ApiModelProperty(required = true, value = "Pass `true` to ackowledge an alarm, or pass `false` to unacknowledge it.")

  public Boolean getAcknowledged() {
    return acknowledged;
  }


  public void setAcknowledged(Boolean acknowledged) {
    this.acknowledged = acknowledged;
  }


  public ApiHealthAlarmsAcknowledgeAlarm ids(List<Long> ids) {
    
    this.ids = ids;
    return this;
  }

  public ApiHealthAlarmsAcknowledgeAlarm addIdsItem(Long idsItem) {
    if (this.ids == null) {
      this.ids = new ArrayList<Long>();
    }
    this.ids.add(idsItem);
    return this;
  }

   /**
   * Array of Alarm ID(s)to be updated.
   * @return ids
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "Array of Alarm ID(s)to be updated.")

  public List<Long> getIds() {
    return ids;
  }


  public void setIds(List<Long> ids) {
    this.ids = ids;
  }


  public ApiHealthAlarmsAcknowledgeAlarm all(Boolean all) {
    
    this.all = all;
    return this;
  }

   /**
   * Pass &#x60;true&#x60; to update all alarms instead of passing ids. This will update any active alarm that is not already acknowledged. 
   * @return all
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "Pass `true` to update all alarms instead of passing ids. This will update any active alarm that is not already acknowledged. ")

  public Boolean getAll() {
    return all;
  }


  public void setAll(Boolean all) {
    this.all = all;
  }


  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (o == null || getClass() != o.getClass()) {
      return false;
    }
    ApiHealthAlarmsAcknowledgeAlarm apiHealthAlarmsAcknowledgeAlarm = (ApiHealthAlarmsAcknowledgeAlarm) o;
    return Objects.equals(this.acknowledged, apiHealthAlarmsAcknowledgeAlarm.acknowledged) &&
        Objects.equals(this.ids, apiHealthAlarmsAcknowledgeAlarm.ids) &&
        Objects.equals(this.all, apiHealthAlarmsAcknowledgeAlarm.all);
  }

  @Override
  public int hashCode() {
    return Objects.hash(acknowledged, ids, all);
  }


  @Override
  public String toString() {
    StringBuilder sb = new StringBuilder();
    sb.append("class ApiHealthAlarmsAcknowledgeAlarm {\n");
    sb.append("    acknowledged: ").append(toIndentedString(acknowledged)).append("\n");
    sb.append("    ids: ").append(toIndentedString(ids)).append("\n");
    sb.append("    all: ").append(toIndentedString(all)).append("\n");
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

