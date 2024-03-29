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
import org.openapitools.client.model.InstanceScheduleCreateThreshold;
import org.threeten.bp.OffsetDateTime;

/**
 * InstanceScheduleCreate
 */
@javax.annotation.Generated(value = "org.openapitools.codegen.languages.JavaClientCodegen", date = "2023-08-21T11:26:47.028Z[GMT]")
public class InstanceScheduleCreate {
  /**
   * Gets or Sets scheduleType
   */
  @JsonAdapter(ScheduleTypeEnum.Adapter.class)
  public enum ScheduleTypeEnum {
    DAYOFWEEK("dayOfWeek"),
    
    EXACT("exact");

    private String value;

    ScheduleTypeEnum(String value) {
      this.value = value;
    }

    public String getValue() {
      return value;
    }

    @Override
    public String toString() {
      return String.valueOf(value);
    }

    public static ScheduleTypeEnum fromValue(String value) {
      for (ScheduleTypeEnum b : ScheduleTypeEnum.values()) {
        if (b.value.equals(value)) {
          return b;
        }
      }
      throw new IllegalArgumentException("Unexpected value '" + value + "'");
    }

    public static class Adapter extends TypeAdapter<ScheduleTypeEnum> {
      @Override
      public void write(final JsonWriter jsonWriter, final ScheduleTypeEnum enumeration) throws IOException {
        jsonWriter.value(enumeration.getValue());
      }

      @Override
      public ScheduleTypeEnum read(final JsonReader jsonReader) throws IOException {
        String value =  jsonReader.nextString();
        return ScheduleTypeEnum.fromValue(value);
      }
    }
  }

  public static final String SERIALIZED_NAME_SCHEDULE_TYPE = "scheduleType";
  @SerializedName(SERIALIZED_NAME_SCHEDULE_TYPE)
  private ScheduleTypeEnum scheduleType = ScheduleTypeEnum.DAYOFWEEK;

  public static final String SERIALIZED_NAME_SCHEDULE_TIMEZONE = "scheduleTimezone";
  @SerializedName(SERIALIZED_NAME_SCHEDULE_TIMEZONE)
  private String scheduleTimezone = "UTC";

  public static final String SERIALIZED_NAME_START_DAY_OF_WEEK = "startDayOfWeek";
  @SerializedName(SERIALIZED_NAME_START_DAY_OF_WEEK)
  private Integer startDayOfWeek;

  public static final String SERIALIZED_NAME_START_TIME = "startTime";
  @SerializedName(SERIALIZED_NAME_START_TIME)
  private String startTime;

  public static final String SERIALIZED_NAME_END_DAY_OF_WEEK = "endDayOfWeek";
  @SerializedName(SERIALIZED_NAME_END_DAY_OF_WEEK)
  private Integer endDayOfWeek;

  public static final String SERIALIZED_NAME_END_TIME = "endTime";
  @SerializedName(SERIALIZED_NAME_END_TIME)
  private String endTime;

  public static final String SERIALIZED_NAME_START_DATE = "startDate";
  @SerializedName(SERIALIZED_NAME_START_DATE)
  private OffsetDateTime startDate;

  public static final String SERIALIZED_NAME_END_DATE = "endDate";
  @SerializedName(SERIALIZED_NAME_END_DATE)
  private OffsetDateTime endDate;

  public static final String SERIALIZED_NAME_THRESHOLD = "threshold";
  @SerializedName(SERIALIZED_NAME_THRESHOLD)
  private InstanceScheduleCreateThreshold threshold;


  public InstanceScheduleCreate scheduleType(ScheduleTypeEnum scheduleType) {
    
    this.scheduleType = scheduleType;
    return this;
  }

   /**
   * Get scheduleType
   * @return scheduleType
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public ScheduleTypeEnum getScheduleType() {
    return scheduleType;
  }


  public void setScheduleType(ScheduleTypeEnum scheduleType) {
    this.scheduleType = scheduleType;
  }


  public InstanceScheduleCreate scheduleTimezone(String scheduleTimezone) {
    
    this.scheduleTimezone = scheduleTimezone;
    return this;
  }

   /**
   * Time Zone eg. America/New_York, Europe/Amsterdam, etc. Only used and required for scheduleType &#x60;dayOfWeek&#x60;
   * @return scheduleTimezone
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "Time Zone eg. America/New_York, Europe/Amsterdam, etc. Only used and required for scheduleType `dayOfWeek`")

  public String getScheduleTimezone() {
    return scheduleTimezone;
  }


  public void setScheduleTimezone(String scheduleTimezone) {
    this.scheduleTimezone = scheduleTimezone;
  }


  public InstanceScheduleCreate startDayOfWeek(Integer startDayOfWeek) {
    
    this.startDayOfWeek = startDayOfWeek;
    return this;
  }

   /**
   * Start day of the week 1-7 (Sun-Sat). Only used and required for scheduleType &#x60;dayOfWeek&#x60;
   * @return startDayOfWeek
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(example = "1", value = "Start day of the week 1-7 (Sun-Sat). Only used and required for scheduleType `dayOfWeek`")

  public Integer getStartDayOfWeek() {
    return startDayOfWeek;
  }


  public void setStartDayOfWeek(Integer startDayOfWeek) {
    this.startDayOfWeek = startDayOfWeek;
  }


  public InstanceScheduleCreate startTime(String startTime) {
    
    this.startTime = startTime;
    return this;
  }

   /**
   * Start time of the day in 24-hour format. Only used and required for scheduleType &#x60;dayOfWeek&#x60;
   * @return startTime
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(example = "01:00", value = "Start time of the day in 24-hour format. Only used and required for scheduleType `dayOfWeek`")

  public String getStartTime() {
    return startTime;
  }


  public void setStartTime(String startTime) {
    this.startTime = startTime;
  }


  public InstanceScheduleCreate endDayOfWeek(Integer endDayOfWeek) {
    
    this.endDayOfWeek = endDayOfWeek;
    return this;
  }

   /**
   * End day of the week 1-7 (Sun-Sat). Only used and required for scheduleType &#x60;dayOfWeek&#x60;
   * @return endDayOfWeek
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(example = "7", value = "End day of the week 1-7 (Sun-Sat). Only used and required for scheduleType `dayOfWeek`")

  public Integer getEndDayOfWeek() {
    return endDayOfWeek;
  }


  public void setEndDayOfWeek(Integer endDayOfWeek) {
    this.endDayOfWeek = endDayOfWeek;
  }


  public InstanceScheduleCreate endTime(String endTime) {
    
    this.endTime = endTime;
    return this;
  }

   /**
   * End time of the day in 24-hour format. Only used and required for scheduleType &#x60;dayOfWeek&#x60;
   * @return endTime
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(example = "20:15", value = "End time of the day in 24-hour format. Only used and required for scheduleType `dayOfWeek`")

  public String getEndTime() {
    return endTime;
  }


  public void setEndTime(String endTime) {
    this.endTime = endTime;
  }


  public InstanceScheduleCreate startDate(OffsetDateTime startDate) {
    
    this.startDate = startDate;
    return this;
  }

   /**
   * Start Date. Only used and required for scheduleType &#x60;exact&#x60;
   * @return startDate
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(example = "2022-12-25T00:00Z", value = "Start Date. Only used and required for scheduleType `exact`")

  public OffsetDateTime getStartDate() {
    return startDate;
  }


  public void setStartDate(OffsetDateTime startDate) {
    this.startDate = startDate;
  }


  public InstanceScheduleCreate endDate(OffsetDateTime endDate) {
    
    this.endDate = endDate;
    return this;
  }

   /**
   * End Date. Only used and required for scheduleType &#x60;exact&#x60;
   * @return endDate
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(example = "2023-01-01T00:00Z", value = "End Date. Only used and required for scheduleType `exact`")

  public OffsetDateTime getEndDate() {
    return endDate;
  }


  public void setEndDate(OffsetDateTime endDate) {
    this.endDate = endDate;
  }


  public InstanceScheduleCreate threshold(InstanceScheduleCreateThreshold threshold) {
    
    this.threshold = threshold;
    return this;
  }

   /**
   * Get threshold
   * @return threshold
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public InstanceScheduleCreateThreshold getThreshold() {
    return threshold;
  }


  public void setThreshold(InstanceScheduleCreateThreshold threshold) {
    this.threshold = threshold;
  }


  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (o == null || getClass() != o.getClass()) {
      return false;
    }
    InstanceScheduleCreate instanceScheduleCreate = (InstanceScheduleCreate) o;
    return Objects.equals(this.scheduleType, instanceScheduleCreate.scheduleType) &&
        Objects.equals(this.scheduleTimezone, instanceScheduleCreate.scheduleTimezone) &&
        Objects.equals(this.startDayOfWeek, instanceScheduleCreate.startDayOfWeek) &&
        Objects.equals(this.startTime, instanceScheduleCreate.startTime) &&
        Objects.equals(this.endDayOfWeek, instanceScheduleCreate.endDayOfWeek) &&
        Objects.equals(this.endTime, instanceScheduleCreate.endTime) &&
        Objects.equals(this.startDate, instanceScheduleCreate.startDate) &&
        Objects.equals(this.endDate, instanceScheduleCreate.endDate) &&
        Objects.equals(this.threshold, instanceScheduleCreate.threshold);
  }

  @Override
  public int hashCode() {
    return Objects.hash(scheduleType, scheduleTimezone, startDayOfWeek, startTime, endDayOfWeek, endTime, startDate, endDate, threshold);
  }


  @Override
  public String toString() {
    StringBuilder sb = new StringBuilder();
    sb.append("class InstanceScheduleCreate {\n");
    sb.append("    scheduleType: ").append(toIndentedString(scheduleType)).append("\n");
    sb.append("    scheduleTimezone: ").append(toIndentedString(scheduleTimezone)).append("\n");
    sb.append("    startDayOfWeek: ").append(toIndentedString(startDayOfWeek)).append("\n");
    sb.append("    startTime: ").append(toIndentedString(startTime)).append("\n");
    sb.append("    endDayOfWeek: ").append(toIndentedString(endDayOfWeek)).append("\n");
    sb.append("    endTime: ").append(toIndentedString(endTime)).append("\n");
    sb.append("    startDate: ").append(toIndentedString(startDate)).append("\n");
    sb.append("    endDate: ").append(toIndentedString(endDate)).append("\n");
    sb.append("    threshold: ").append(toIndentedString(threshold)).append("\n");
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

