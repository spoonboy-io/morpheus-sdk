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
 * For a full list of available NAT options, see natOptionTypes in the specific Network Router Type
 */
@ApiModel(description = "For a full list of available NAT options, see natOptionTypes in the specific Network Router Type")
@javax.annotation.Generated(value = "org.openapitools.codegen.languages.JavaClientCodegen", date = "2023-08-21T11:26:47.028Z[GMT]")
public class ApiNetworksRoutersRouterIdNatsNetworkRouterNAT {
  public static final String SERIALIZED_NAME_NAME = "name";
  @SerializedName(SERIALIZED_NAME_NAME)
  private Object name = null;


  public ApiNetworksRoutersRouterIdNatsNetworkRouterNAT name(Object name) {
    
    this.name = name;
    return this;
  }

   /**
   * NAT name
   * @return name
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(required = true, value = "NAT name")

  public Object getName() {
    return name;
  }


  public void setName(Object name) {
    this.name = name;
  }


  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (o == null || getClass() != o.getClass()) {
      return false;
    }
    ApiNetworksRoutersRouterIdNatsNetworkRouterNAT apiNetworksRoutersRouterIdNatsNetworkRouterNAT = (ApiNetworksRoutersRouterIdNatsNetworkRouterNAT) o;
    return Objects.equals(this.name, apiNetworksRoutersRouterIdNatsNetworkRouterNAT.name);
  }

  @Override
  public int hashCode() {
    return Objects.hash(name);
  }


  @Override
  public String toString() {
    StringBuilder sb = new StringBuilder();
    sb.append("class ApiNetworksRoutersRouterIdNatsNetworkRouterNAT {\n");
    sb.append("    name: ").append(toIndentedString(name)).append("\n");
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

