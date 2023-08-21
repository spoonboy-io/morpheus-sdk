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
import org.openapitools.client.model.CheckGroup;

/**
 * InlineResponse20019
 */
@javax.annotation.Generated(value = "org.openapitools.codegen.languages.JavaClientCodegen", date = "2023-08-21T11:26:47.028Z[GMT]")
public class InlineResponse20019 {
  public static final String SERIALIZED_NAME_CHECK_GROUP = "checkGroup";
  @SerializedName(SERIALIZED_NAME_CHECK_GROUP)
  private CheckGroup checkGroup;

  public static final String SERIALIZED_NAME_CHECKS = "checks";
  @SerializedName(SERIALIZED_NAME_CHECKS)
  private List<Check> checks = null;


  public InlineResponse20019 checkGroup(CheckGroup checkGroup) {
    
    this.checkGroup = checkGroup;
    return this;
  }

   /**
   * Get checkGroup
   * @return checkGroup
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public CheckGroup getCheckGroup() {
    return checkGroup;
  }


  public void setCheckGroup(CheckGroup checkGroup) {
    this.checkGroup = checkGroup;
  }


  public InlineResponse20019 checks(List<Check> checks) {
    
    this.checks = checks;
    return this;
  }

  public InlineResponse20019 addChecksItem(Check checksItem) {
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


  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (o == null || getClass() != o.getClass()) {
      return false;
    }
    InlineResponse20019 inlineResponse20019 = (InlineResponse20019) o;
    return Objects.equals(this.checkGroup, inlineResponse20019.checkGroup) &&
        Objects.equals(this.checks, inlineResponse20019.checks);
  }

  @Override
  public int hashCode() {
    return Objects.hash(checkGroup, checks);
  }


  @Override
  public String toString() {
    StringBuilder sb = new StringBuilder();
    sb.append("class InlineResponse20019 {\n");
    sb.append("    checkGroup: ").append(toIndentedString(checkGroup)).append("\n");
    sb.append("    checks: ").append(toIndentedString(checks)).append("\n");
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
