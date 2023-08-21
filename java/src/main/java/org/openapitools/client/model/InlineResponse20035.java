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
import org.openapitools.client.model.InlineResponse20035Actions;

/**
 * InlineResponse20035
 */
@javax.annotation.Generated(value = "org.openapitools.codegen.languages.JavaClientCodegen", date = "2023-08-21T11:26:47.028Z[GMT]")
public class InlineResponse20035 {
  public static final String SERIALIZED_NAME_CONTAINER_IDS = "containerIds";
  @SerializedName(SERIALIZED_NAME_CONTAINER_IDS)
  private List<Long> containerIds = null;

  public static final String SERIALIZED_NAME_ACTIONS = "actions";
  @SerializedName(SERIALIZED_NAME_ACTIONS)
  private List<InlineResponse20035Actions> actions = null;


  public InlineResponse20035 containerIds(List<Long> containerIds) {
    
    this.containerIds = containerIds;
    return this;
  }

  public InlineResponse20035 addContainerIdsItem(Long containerIdsItem) {
    if (this.containerIds == null) {
      this.containerIds = new ArrayList<Long>();
    }
    this.containerIds.add(containerIdsItem);
    return this;
  }

   /**
   * Get containerIds
   * @return containerIds
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public List<Long> getContainerIds() {
    return containerIds;
  }


  public void setContainerIds(List<Long> containerIds) {
    this.containerIds = containerIds;
  }


  public InlineResponse20035 actions(List<InlineResponse20035Actions> actions) {
    
    this.actions = actions;
    return this;
  }

  public InlineResponse20035 addActionsItem(InlineResponse20035Actions actionsItem) {
    if (this.actions == null) {
      this.actions = new ArrayList<InlineResponse20035Actions>();
    }
    this.actions.add(actionsItem);
    return this;
  }

   /**
   * Get actions
   * @return actions
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public List<InlineResponse20035Actions> getActions() {
    return actions;
  }


  public void setActions(List<InlineResponse20035Actions> actions) {
    this.actions = actions;
  }


  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (o == null || getClass() != o.getClass()) {
      return false;
    }
    InlineResponse20035 inlineResponse20035 = (InlineResponse20035) o;
    return Objects.equals(this.containerIds, inlineResponse20035.containerIds) &&
        Objects.equals(this.actions, inlineResponse20035.actions);
  }

  @Override
  public int hashCode() {
    return Objects.hash(containerIds, actions);
  }


  @Override
  public String toString() {
    StringBuilder sb = new StringBuilder();
    sb.append("class InlineResponse20035 {\n");
    sb.append("    containerIds: ").append(toIndentedString(containerIds)).append("\n");
    sb.append("    actions: ").append(toIndentedString(actions)).append("\n");
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
