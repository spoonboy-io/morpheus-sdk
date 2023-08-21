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
import org.openapitools.client.model.Check;
import org.openapitools.client.model.CheckApp;
import org.openapitools.client.model.CheckGroup;
import org.openapitools.client.model.Incident;

/**
 * InlineResponse20016
 */
@javax.annotation.Generated(value = "org.openapitools.codegen.languages.JavaClientCodegen", date = "2023-08-21T11:26:47.028Z[GMT]")
public class InlineResponse20016 {
  public static final String SERIALIZED_NAME_MONITOR_APP = "monitorApp";
  @SerializedName(SERIALIZED_NAME_MONITOR_APP)
  private CheckApp monitorApp;

  public static final String SERIALIZED_NAME_CHECK_GROUPS = "checkGroups";
  @SerializedName(SERIALIZED_NAME_CHECK_GROUPS)
  private List<CheckGroup> checkGroups = null;

  public static final String SERIALIZED_NAME_CHECKS = "checks";
  @SerializedName(SERIALIZED_NAME_CHECKS)
  private List<Check> checks = null;

  public static final String SERIALIZED_NAME_OPEN_INCIDENTS = "openIncidents";
  @SerializedName(SERIALIZED_NAME_OPEN_INCIDENTS)
  private List<Incident> openIncidents = null;


  public InlineResponse20016 monitorApp(CheckApp monitorApp) {
    
    this.monitorApp = monitorApp;
    return this;
  }

   /**
   * Get monitorApp
   * @return monitorApp
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public CheckApp getMonitorApp() {
    return monitorApp;
  }


  public void setMonitorApp(CheckApp monitorApp) {
    this.monitorApp = monitorApp;
  }


  public InlineResponse20016 checkGroups(List<CheckGroup> checkGroups) {
    
    this.checkGroups = checkGroups;
    return this;
  }

  public InlineResponse20016 addCheckGroupsItem(CheckGroup checkGroupsItem) {
    if (this.checkGroups == null) {
      this.checkGroups = new ArrayList<CheckGroup>();
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

  public List<CheckGroup> getCheckGroups() {
    return checkGroups;
  }


  public void setCheckGroups(List<CheckGroup> checkGroups) {
    this.checkGroups = checkGroups;
  }


  public InlineResponse20016 checks(List<Check> checks) {
    
    this.checks = checks;
    return this;
  }

  public InlineResponse20016 addChecksItem(Check checksItem) {
    if (this.checks == null) {
      this.checks = new ArrayList<Check>();
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

  public List<Check> getChecks() {
    return checks;
  }


  public void setChecks(List<Check> checks) {
    this.checks = checks;
  }


  public InlineResponse20016 openIncidents(List<Incident> openIncidents) {
    
    this.openIncidents = openIncidents;
    return this;
  }

  public InlineResponse20016 addOpenIncidentsItem(Incident openIncidentsItem) {
    if (this.openIncidents == null) {
      this.openIncidents = new ArrayList<Incident>();
    }
    this.openIncidents.add(openIncidentsItem);
    return this;
  }

   /**
   * Get openIncidents
   * @return openIncidents
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public List<Incident> getOpenIncidents() {
    return openIncidents;
  }


  public void setOpenIncidents(List<Incident> openIncidents) {
    this.openIncidents = openIncidents;
  }


  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (o == null || getClass() != o.getClass()) {
      return false;
    }
    InlineResponse20016 inlineResponse20016 = (InlineResponse20016) o;
    return Objects.equals(this.monitorApp, inlineResponse20016.monitorApp) &&
        Objects.equals(this.checkGroups, inlineResponse20016.checkGroups) &&
        Objects.equals(this.checks, inlineResponse20016.checks) &&
        Objects.equals(this.openIncidents, inlineResponse20016.openIncidents);
  }

  @Override
  public int hashCode() {
    return Objects.hash(monitorApp, checkGroups, checks, openIncidents);
  }


  @Override
  public String toString() {
    StringBuilder sb = new StringBuilder();
    sb.append("class InlineResponse20016 {\n");
    sb.append("    monitorApp: ").append(toIndentedString(monitorApp)).append("\n");
    sb.append("    checkGroups: ").append(toIndentedString(checkGroups)).append("\n");
    sb.append("    checks: ").append(toIndentedString(checks)).append("\n");
    sb.append("    openIncidents: ").append(toIndentedString(openIncidents)).append("\n");
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
