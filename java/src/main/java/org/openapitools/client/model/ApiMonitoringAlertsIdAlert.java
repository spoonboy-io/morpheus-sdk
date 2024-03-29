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
import org.openapitools.client.model.ApiMonitoringAlertsAlertContacts;

/**
 * Payload for creating a new monitoring alert
 */
@ApiModel(description = "Payload for creating a new monitoring alert")
@javax.annotation.Generated(value = "org.openapitools.codegen.languages.JavaClientCodegen", date = "2023-08-21T11:26:47.028Z[GMT]")
public class ApiMonitoringAlertsIdAlert {
  public static final String SERIALIZED_NAME_NAME = "name";
  @SerializedName(SERIALIZED_NAME_NAME)
  private String name;

  public static final String SERIALIZED_NAME_MIN_DURATION = "minDuration";
  @SerializedName(SERIALIZED_NAME_MIN_DURATION)
  private Integer minDuration = 0;

  /**
   * Severity level threshold for sending notifications.
   */
  @JsonAdapter(MinSeverityEnum.Adapter.class)
  public enum MinSeverityEnum {
    INFO("info"),
    
    WARNING("warning"),
    
    CRITICAL("critical");

    private String value;

    MinSeverityEnum(String value) {
      this.value = value;
    }

    public String getValue() {
      return value;
    }

    @Override
    public String toString() {
      return String.valueOf(value);
    }

    public static MinSeverityEnum fromValue(String value) {
      for (MinSeverityEnum b : MinSeverityEnum.values()) {
        if (b.value.equals(value)) {
          return b;
        }
      }
      throw new IllegalArgumentException("Unexpected value '" + value + "'");
    }

    public static class Adapter extends TypeAdapter<MinSeverityEnum> {
      @Override
      public void write(final JsonWriter jsonWriter, final MinSeverityEnum enumeration) throws IOException {
        jsonWriter.value(enumeration.getValue());
      }

      @Override
      public MinSeverityEnum read(final JsonReader jsonReader) throws IOException {
        String value =  jsonReader.nextString();
        return MinSeverityEnum.fromValue(value);
      }
    }
  }

  public static final String SERIALIZED_NAME_MIN_SEVERITY = "minSeverity";
  @SerializedName(SERIALIZED_NAME_MIN_SEVERITY)
  private MinSeverityEnum minSeverity = MinSeverityEnum.CRITICAL;

  public static final String SERIALIZED_NAME_ACTIVE = "active";
  @SerializedName(SERIALIZED_NAME_ACTIVE)
  private Boolean active = true;

  public static final String SERIALIZED_NAME_ALL_CHECKS = "allChecks";
  @SerializedName(SERIALIZED_NAME_ALL_CHECKS)
  private Boolean allChecks = false;

  public static final String SERIALIZED_NAME_ALL_GROUPS = "allGroups";
  @SerializedName(SERIALIZED_NAME_ALL_GROUPS)
  private Boolean allGroups = false;

  public static final String SERIALIZED_NAME_ALL_APPS = "allApps";
  @SerializedName(SERIALIZED_NAME_ALL_APPS)
  private Boolean allApps = false;

  public static final String SERIALIZED_NAME_CHECKS = "checks";
  @SerializedName(SERIALIZED_NAME_CHECKS)
  private List<Integer> checks = null;

  public static final String SERIALIZED_NAME_GROUPS = "groups";
  @SerializedName(SERIALIZED_NAME_GROUPS)
  private List<Integer> groups = null;

  public static final String SERIALIZED_NAME_APPS = "apps";
  @SerializedName(SERIALIZED_NAME_APPS)
  private List<Integer> apps = null;

  public static final String SERIALIZED_NAME_CONTACTS = "contacts";
  @SerializedName(SERIALIZED_NAME_CONTACTS)
  private List<ApiMonitoringAlertsAlertContacts> contacts = null;


  public ApiMonitoringAlertsIdAlert name(String name) {
    
    this.name = name;
    return this;
  }

   /**
   * Unique name scoped to your account for the alert
   * @return name
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(example = "My Alert", value = "Unique name scoped to your account for the alert")

  public String getName() {
    return name;
  }


  public void setName(String name) {
    this.name = name;
  }


  public ApiMonitoringAlertsIdAlert minDuration(Integer minDuration) {
    
    this.minDuration = minDuration;
    return this;
  }

   /**
   * Duration in minutes of the delay before sending notification(s)
   * @return minDuration
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "Duration in minutes of the delay before sending notification(s)")

  public Integer getMinDuration() {
    return minDuration;
  }


  public void setMinDuration(Integer minDuration) {
    this.minDuration = minDuration;
  }


  public ApiMonitoringAlertsIdAlert minSeverity(MinSeverityEnum minSeverity) {
    
    this.minSeverity = minSeverity;
    return this;
  }

   /**
   * Severity level threshold for sending notifications.
   * @return minSeverity
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "Severity level threshold for sending notifications.")

  public MinSeverityEnum getMinSeverity() {
    return minSeverity;
  }


  public void setMinSeverity(MinSeverityEnum minSeverity) {
    this.minSeverity = minSeverity;
  }


  public ApiMonitoringAlertsIdAlert active(Boolean active) {
    
    this.active = active;
    return this;
  }

   /**
   * Set to false to disable notifications
   * @return active
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "Set to false to disable notifications")

  public Boolean getActive() {
    return active;
  }


  public void setActive(Boolean active) {
    this.active = active;
  }


  public ApiMonitoringAlertsIdAlert allChecks(Boolean allChecks) {
    
    this.allChecks = allChecks;
    return this;
  }

   /**
   * Trigger for all checks
   * @return allChecks
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "Trigger for all checks")

  public Boolean getAllChecks() {
    return allChecks;
  }


  public void setAllChecks(Boolean allChecks) {
    this.allChecks = allChecks;
  }


  public ApiMonitoringAlertsIdAlert allGroups(Boolean allGroups) {
    
    this.allGroups = allGroups;
    return this;
  }

   /**
   * Trigger for all check groups
   * @return allGroups
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "Trigger for all check groups")

  public Boolean getAllGroups() {
    return allGroups;
  }


  public void setAllGroups(Boolean allGroups) {
    this.allGroups = allGroups;
  }


  public ApiMonitoringAlertsIdAlert allApps(Boolean allApps) {
    
    this.allApps = allApps;
    return this;
  }

   /**
   * Trigger for all monitor apps
   * @return allApps
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "Trigger for all monitor apps")

  public Boolean getAllApps() {
    return allApps;
  }


  public void setAllApps(Boolean allApps) {
    this.allApps = allApps;
  }


  public ApiMonitoringAlertsIdAlert checks(List<Integer> checks) {
    
    this.checks = checks;
    return this;
  }

  public ApiMonitoringAlertsIdAlert addChecksItem(Integer checksItem) {
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


  public ApiMonitoringAlertsIdAlert groups(List<Integer> groups) {
    
    this.groups = groups;
    return this;
  }

  public ApiMonitoringAlertsIdAlert addGroupsItem(Integer groupsItem) {
    if (this.groups == null) {
      this.groups = new ArrayList<Integer>();
    }
    this.groups.add(groupsItem);
    return this;
  }

   /**
   * Get groups
   * @return groups
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public List<Integer> getGroups() {
    return groups;
  }


  public void setGroups(List<Integer> groups) {
    this.groups = groups;
  }


  public ApiMonitoringAlertsIdAlert apps(List<Integer> apps) {
    
    this.apps = apps;
    return this;
  }

  public ApiMonitoringAlertsIdAlert addAppsItem(Integer appsItem) {
    if (this.apps == null) {
      this.apps = new ArrayList<Integer>();
    }
    this.apps.add(appsItem);
    return this;
  }

   /**
   * Get apps
   * @return apps
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public List<Integer> getApps() {
    return apps;
  }


  public void setApps(List<Integer> apps) {
    this.apps = apps;
  }


  public ApiMonitoringAlertsIdAlert contacts(List<ApiMonitoringAlertsAlertContacts> contacts) {
    
    this.contacts = contacts;
    return this;
  }

  public ApiMonitoringAlertsIdAlert addContactsItem(ApiMonitoringAlertsAlertContacts contactsItem) {
    if (this.contacts == null) {
      this.contacts = new ArrayList<ApiMonitoringAlertsAlertContacts>();
    }
    this.contacts.add(contactsItem);
    return this;
  }

   /**
   * Get contacts
   * @return contacts
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public List<ApiMonitoringAlertsAlertContacts> getContacts() {
    return contacts;
  }


  public void setContacts(List<ApiMonitoringAlertsAlertContacts> contacts) {
    this.contacts = contacts;
  }


  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (o == null || getClass() != o.getClass()) {
      return false;
    }
    ApiMonitoringAlertsIdAlert apiMonitoringAlertsIdAlert = (ApiMonitoringAlertsIdAlert) o;
    return Objects.equals(this.name, apiMonitoringAlertsIdAlert.name) &&
        Objects.equals(this.minDuration, apiMonitoringAlertsIdAlert.minDuration) &&
        Objects.equals(this.minSeverity, apiMonitoringAlertsIdAlert.minSeverity) &&
        Objects.equals(this.active, apiMonitoringAlertsIdAlert.active) &&
        Objects.equals(this.allChecks, apiMonitoringAlertsIdAlert.allChecks) &&
        Objects.equals(this.allGroups, apiMonitoringAlertsIdAlert.allGroups) &&
        Objects.equals(this.allApps, apiMonitoringAlertsIdAlert.allApps) &&
        Objects.equals(this.checks, apiMonitoringAlertsIdAlert.checks) &&
        Objects.equals(this.groups, apiMonitoringAlertsIdAlert.groups) &&
        Objects.equals(this.apps, apiMonitoringAlertsIdAlert.apps) &&
        Objects.equals(this.contacts, apiMonitoringAlertsIdAlert.contacts);
  }

  @Override
  public int hashCode() {
    return Objects.hash(name, minDuration, minSeverity, active, allChecks, allGroups, allApps, checks, groups, apps, contacts);
  }


  @Override
  public String toString() {
    StringBuilder sb = new StringBuilder();
    sb.append("class ApiMonitoringAlertsIdAlert {\n");
    sb.append("    name: ").append(toIndentedString(name)).append("\n");
    sb.append("    minDuration: ").append(toIndentedString(minDuration)).append("\n");
    sb.append("    minSeverity: ").append(toIndentedString(minSeverity)).append("\n");
    sb.append("    active: ").append(toIndentedString(active)).append("\n");
    sb.append("    allChecks: ").append(toIndentedString(allChecks)).append("\n");
    sb.append("    allGroups: ").append(toIndentedString(allGroups)).append("\n");
    sb.append("    allApps: ").append(toIndentedString(allApps)).append("\n");
    sb.append("    checks: ").append(toIndentedString(checks)).append("\n");
    sb.append("    groups: ").append(toIndentedString(groups)).append("\n");
    sb.append("    apps: ").append(toIndentedString(apps)).append("\n");
    sb.append("    contacts: ").append(toIndentedString(contacts)).append("\n");
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

