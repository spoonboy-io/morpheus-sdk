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
 * TaskNestedWorkflowConfigTaskOptions
 */
@javax.annotation.Generated(value = "org.openapitools.codegen.languages.JavaClientCodegen", date = "2023-08-21T11:26:47.028Z[GMT]")
public class TaskNestedWorkflowConfigTaskOptions {
  public static final String SERIALIZED_NAME_OPERATIONAL_WORKFLOW_ID = "operationalWorkflowId";
  @SerializedName(SERIALIZED_NAME_OPERATIONAL_WORKFLOW_ID)
  private String operationalWorkflowId;

  public static final String SERIALIZED_NAME_OPERATIONAL_WORKFLOW_NAME = "operationalWorkflowName";
  @SerializedName(SERIALIZED_NAME_OPERATIONAL_WORKFLOW_NAME)
  private String operationalWorkflowName;


  public TaskNestedWorkflowConfigTaskOptions operationalWorkflowId(String operationalWorkflowId) {
    
    this.operationalWorkflowId = operationalWorkflowId;
    return this;
  }

   /**
   * Operational Workflow ID
   * @return operationalWorkflowId
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(required = true, value = "Operational Workflow ID")

  public String getOperationalWorkflowId() {
    return operationalWorkflowId;
  }


  public void setOperationalWorkflowId(String operationalWorkflowId) {
    this.operationalWorkflowId = operationalWorkflowId;
  }


  public TaskNestedWorkflowConfigTaskOptions operationalWorkflowName(String operationalWorkflowName) {
    
    this.operationalWorkflowName = operationalWorkflowName;
    return this;
  }

   /**
   * Operational Workflow Name
   * @return operationalWorkflowName
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "Operational Workflow Name")

  public String getOperationalWorkflowName() {
    return operationalWorkflowName;
  }


  public void setOperationalWorkflowName(String operationalWorkflowName) {
    this.operationalWorkflowName = operationalWorkflowName;
  }


  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (o == null || getClass() != o.getClass()) {
      return false;
    }
    TaskNestedWorkflowConfigTaskOptions taskNestedWorkflowConfigTaskOptions = (TaskNestedWorkflowConfigTaskOptions) o;
    return Objects.equals(this.operationalWorkflowId, taskNestedWorkflowConfigTaskOptions.operationalWorkflowId) &&
        Objects.equals(this.operationalWorkflowName, taskNestedWorkflowConfigTaskOptions.operationalWorkflowName);
  }

  @Override
  public int hashCode() {
    return Objects.hash(operationalWorkflowId, operationalWorkflowName);
  }


  @Override
  public String toString() {
    StringBuilder sb = new StringBuilder();
    sb.append("class TaskNestedWorkflowConfigTaskOptions {\n");
    sb.append("    operationalWorkflowId: ").append(toIndentedString(operationalWorkflowId)).append("\n");
    sb.append("    operationalWorkflowName: ").append(toIndentedString(operationalWorkflowName)).append("\n");
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

