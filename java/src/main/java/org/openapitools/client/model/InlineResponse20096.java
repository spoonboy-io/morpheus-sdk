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
import org.openapitools.client.model.InlineResponse20096NetworkRouterBgpNeighbors;

/**
 * InlineResponse20096
 */
@javax.annotation.Generated(value = "org.openapitools.codegen.languages.JavaClientCodegen", date = "2023-08-21T11:26:47.028Z[GMT]")
public class InlineResponse20096 {
  public static final String SERIALIZED_NAME_NETWORK_ROUTER_BGP_NEIGHBORS = "networkRouterBgpNeighbors";
  @SerializedName(SERIALIZED_NAME_NETWORK_ROUTER_BGP_NEIGHBORS)
  private List<InlineResponse20096NetworkRouterBgpNeighbors> networkRouterBgpNeighbors = null;


  public InlineResponse20096 networkRouterBgpNeighbors(List<InlineResponse20096NetworkRouterBgpNeighbors> networkRouterBgpNeighbors) {
    
    this.networkRouterBgpNeighbors = networkRouterBgpNeighbors;
    return this;
  }

  public InlineResponse20096 addNetworkRouterBgpNeighborsItem(InlineResponse20096NetworkRouterBgpNeighbors networkRouterBgpNeighborsItem) {
    if (this.networkRouterBgpNeighbors == null) {
      this.networkRouterBgpNeighbors = new ArrayList<InlineResponse20096NetworkRouterBgpNeighbors>();
    }
    this.networkRouterBgpNeighbors.add(networkRouterBgpNeighborsItem);
    return this;
  }

   /**
   * Get networkRouterBgpNeighbors
   * @return networkRouterBgpNeighbors
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public List<InlineResponse20096NetworkRouterBgpNeighbors> getNetworkRouterBgpNeighbors() {
    return networkRouterBgpNeighbors;
  }


  public void setNetworkRouterBgpNeighbors(List<InlineResponse20096NetworkRouterBgpNeighbors> networkRouterBgpNeighbors) {
    this.networkRouterBgpNeighbors = networkRouterBgpNeighbors;
  }


  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (o == null || getClass() != o.getClass()) {
      return false;
    }
    InlineResponse20096 inlineResponse20096 = (InlineResponse20096) o;
    return Objects.equals(this.networkRouterBgpNeighbors, inlineResponse20096.networkRouterBgpNeighbors);
  }

  @Override
  public int hashCode() {
    return Objects.hash(networkRouterBgpNeighbors);
  }


  @Override
  public String toString() {
    StringBuilder sb = new StringBuilder();
    sb.append("class InlineResponse20096 {\n");
    sb.append("    networkRouterBgpNeighbors: ").append(toIndentedString(networkRouterBgpNeighbors)).append("\n");
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
