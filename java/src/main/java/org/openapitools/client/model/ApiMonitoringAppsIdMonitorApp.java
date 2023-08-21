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
 * Payload for creating a new monitoring check app
 */
@ApiModel(description = "Payload for creating a new monitoring check app")
@javax.annotation.Generated(value = "org.openapitools.codegen.languages.JavaClientCodegen", date = "2023-08-21T11:26:47.028Z[GMT]")
public class ApiMonitoringAppsIdMonitorApp {
  public static final String SERIALIZED_NAME_NAME = "name";
  @SerializedName(SERIALIZED_NAME_NAME)
  private String name;

  public static final String SERIALIZED_NAME_DESCRIPTION = "description";
  @SerializedName(SERIALIZED_NAME_DESCRIPTION)
  private String description;

  public static final String SERIALIZED_NAME_IN_UPTIME = "inUptime";
  @SerializedName(SERIALIZED_NAME_IN_UPTIME)
  private Boolean inUptime = true;

  /**
   * Severity level of incidents that are created when this check fails
   */
  @JsonAdapter(SeverityEnum.Adapter.class)
  public enum SeverityEnum {
    INFO("info"),
    
    WARNING("warning"),
    
    CRITICAL("critical");

    private String value;

    SeverityEnum(String value) {
      this.value = value;
    }

    public String getValue() {
      return value;
    }

    @Override
    public String toString() {
      return String.valueOf(value);
    }

    public static SeverityEnum fromValue(String value) {
      for (SeverityEnum b : SeverityEnum.values()) {
        if (b.value.equals(value)) {
          return b;
        }
      }
      throw new IllegalArgumentException("Unexpected value '" + value + "'");
    }

    public static class Adapter extends TypeAdapter<SeverityEnum> {
      @Override
      public void write(final JsonWriter jsonWriter, final SeverityEnum enumeration) throws IOException {
        jsonWriter.value(enumeration.getValue());
      }

      @Override
      public SeverityEnum read(final JsonReader jsonReader) throws IOException {
        String value =  jsonReader.nextString();
        return SeverityEnum.fromValue(value);
      }
    }
  }

  public static final String SERIALIZED_NAME_SEVERITY = "severity";
  @SerializedName(SERIALIZED_NAME_SEVERITY)
  private SeverityEnum severity = SeverityEnum.CRITICAL;

  public static final String SERIALIZED_NAME_ACTIVE = "active";
  @SerializedName(SERIALIZED_NAME_ACTIVE)
  private Boolean active = true;

  public static final String SERIALIZED_NAME_CHECKS = "checks";
  @SerializedName(SERIALIZED_NAME_CHECKS)
  private List<Integer> checks = null;

  public static final String SERIALIZED_NAME_CHECK_GROUPS = "checkGroups";
  @SerializedName(SERIALIZED_NAME_CHECK_GROUPS)
  private List<Integer> checkGroups = null;


  public ApiMonitoringAppsIdMonitorApp name(String name) {
    
    this.name = name;
    return this;
  }

   /**
   * Unique name scoped to your account for the check app
   * @return name
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(example = "My Check App", value = "Unique name scoped to your account for the check app")

  public String getName() {
    return name;
  }


  public void setName(String name) {
    this.name = name;
  }


  public ApiMonitoringAppsIdMonitorApp description(String description) {
    
    this.description = description;
    return this;
  }

   /**
   * Optional description field
   * @return description
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(example = "My cool description", value = "Optional description field")

  public String getDescription() {
    return description;
  }


  public void setDescription(String description) {
    this.description = description;
  }


  public ApiMonitoringAppsIdMonitorApp inUptime(Boolean inUptime) {
    
    this.inUptime = inUptime;
    return this;
  }

   /**
   * Used to determine if check should affect account wide availability calculations
   * @return inUptime
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "Used to determine if check should affect account wide availability calculations")

  public Boolean getInUptime() {
    return inUptime;
  }


  public void setInUptime(Boolean inUptime) {
    this.inUptime = inUptime;
  }


  public ApiMonitoringAppsIdMonitorApp severity(SeverityEnum severity) {
    
    this.severity = severity;
    return this;
  }

   /**
   * Severity level of incidents that are created when this check fails
   * @return severity
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "Severity level of incidents that are created when this check fails")

  public SeverityEnum getSeverity() {
    return severity;
  }


  public void setSeverity(SeverityEnum severity) {
    this.severity = severity;
  }


  public ApiMonitoringAppsIdMonitorApp active(Boolean active) {
    
    this.active = active;
    return this;
  }

   /**
   * Used to determine if check app is active
   * @return active
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "Used to determine if check app is active")

  public Boolean getActive() {
    return active;
  }


  public void setActive(Boolean active) {
    this.active = active;
  }


  public ApiMonitoringAppsIdMonitorApp checks(List<Integer> checks) {
    
    this.checks = checks;
    return this;
  }

  public ApiMonitoringAppsIdMonitorApp addChecksItem(Integer checksItem) {
    if (this.checks == null) {
      this.checks = new ArrayList<Integer>();
    }
    this.checks.add(checksItem);
    return this;
  }

   /**
   * Get checks
   * @return checks
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public List<Integer> getChecks() {
    return checks;
  }


  public void setChecks(List<Integer> checks) {
    this.checks = checks;
  }


  public ApiMonitoringAppsIdMonitorApp checkGroups(List<Integer> checkGroups) {
    
    this.checkGroups = checkGroups;
    return this;
  }

  public ApiMonitoringAppsIdMonitorApp addCheckGroupsItem(Integer checkGroupsItem) {
    if (this.checkGroups == null) {
      this.checkGroups = new ArrayList<Integer>();
    }
    this.checkGroups.add(checkGroupsItem);
    return this;
  }

   /**
   * Get checkGroups
   * @return checkGroups
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public List<Integer> getCheckGroups() {
    return checkGroups;
  }


  public void setCheckGroups(List<Integer> checkGroups) {
    this.checkGroups = checkGroups;
  }


  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (o == null || getClass() != o.getClass()) {
      return false;
    }
    ApiMonitoringAppsIdMonitorApp apiMonitoringAppsIdMonitorApp = (ApiMonitoringAppsIdMonitorApp) o;
    return Objects.equals(this.name, apiMonitoringAppsIdMonitorApp.name) &&
        Objects.equals(this.description, apiMonitoringAppsIdMonitorApp.description) &&
        Objects.equals(this.inUptime, apiMonitoringAppsIdMonitorApp.inUptime) &&
        Objects.equals(this.severity, apiMonitoringAppsIdMonitorApp.severity) &&
        Objects.equals(this.active, apiMonitoringAppsIdMonitorApp.active) &&
        Objects.equals(this.checks, apiMonitoringAppsIdMonitorApp.checks) &&
        Objects.equals(this.checkGroups, apiMonitoringAppsIdMonitorApp.checkGroups);
  }

  @Override
  public int hashCode() {
    return Objects.hash(name, description, inUptime, severity, active, checks, checkGroups);
  }


  @Override
  public String toString() {
    StringBuilder sb = new StringBuilder();
    sb.append("class ApiMonitoringAppsIdMonitorApp {\n");
    sb.append("    name: ").append(toIndentedString(name)).append("\n");
    sb.append("    description: ").append(toIndentedString(description)).append("\n");
    sb.append("    inUptime: ").append(toIndentedString(inUptime)).append("\n");
    sb.append("    severity: ").append(toIndentedString(severity)).append("\n");
    sb.append("    active: ").append(toIndentedString(active)).append("\n");
    sb.append("    checks: ").append(toIndentedString(checks)).append("\n");
    sb.append("    checkGroups: ").append(toIndentedString(checkGroups)).append("\n");
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

