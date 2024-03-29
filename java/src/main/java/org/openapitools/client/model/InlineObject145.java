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
import org.openapitools.client.model.NetworkStaticRouteCreate;

/**
 * InlineObject145
 */
@javax.annotation.Generated(value = "org.openapitools.codegen.languages.JavaClientCodegen", date = "2023-08-21T11:26:47.028Z[GMT]")
public class InlineObject145 {
  public static final String SERIALIZED_NAME_NETWORK_ROUTE = "networkRoute";
  @SerializedName(SERIALIZED_NAME_NETWORK_ROUTE)
  private NetworkStaticRouteCreate networkRoute;


  public InlineObject145 networkRoute(NetworkStaticRouteCreate networkRoute) {
    
    this.networkRoute = networkRoute;
    return this;
  }

   /**
   * Get networkRoute
   * @return networkRoute
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public NetworkStaticRouteCreate getNetworkRoute() {
    return networkRoute;
  }


  public void setNetworkRoute(NetworkStaticRouteCreate networkRoute) {
    this.networkRoute = networkRoute;
  }


  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (o == null || getClass() != o.getClass()) {
      return false;
    }
    InlineObject145 inlineObject145 = (InlineObject145) o;
    return Objects.equals(this.networkRoute, inlineObject145.networkRoute);
  }

  @Override
  public int hashCode() {
    return Objects.hash(networkRoute);
  }


  @Override
  public String toString() {
    StringBuilder sb = new StringBuilder();
    sb.append("class InlineObject145 {\n");
    sb.append("    networkRoute: ").append(toIndentedString(networkRoute)).append("\n");
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

