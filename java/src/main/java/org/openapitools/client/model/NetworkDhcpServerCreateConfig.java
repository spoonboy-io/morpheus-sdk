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
 * Configuration object with parameters that vary by pool type.
 */
@ApiModel(description = "Configuration object with parameters that vary by pool type.")
@javax.annotation.Generated(value = "org.openapitools.codegen.languages.JavaClientCodegen", date = "2023-08-21T11:26:47.028Z[GMT]")
public class NetworkDhcpServerCreateConfig {
  public static final String SERIALIZED_NAME_EDGE_CLUSTER = "edgeCluster";
  @SerializedName(SERIALIZED_NAME_EDGE_CLUSTER)
  private String edgeCluster;

  public static final String SERIALIZED_NAME_PREFERRED_EDGE_NODE1 = "preferredEdgeNode1";
  @SerializedName(SERIALIZED_NAME_PREFERRED_EDGE_NODE1)
  private String preferredEdgeNode1;

  public static final String SERIALIZED_NAME_PREFERRED_EDGE_NODE2 = "preferredEdgeNode2";
  @SerializedName(SERIALIZED_NAME_PREFERRED_EDGE_NODE2)
  private String preferredEdgeNode2;


  public NetworkDhcpServerCreateConfig edgeCluster(String edgeCluster) {
    
    this.edgeCluster = edgeCluster;
    return this;
  }

   /**
   * Edge Cluster
   * @return edgeCluster
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "Edge Cluster")

  public String getEdgeCluster() {
    return edgeCluster;
  }


  public void setEdgeCluster(String edgeCluster) {
    this.edgeCluster = edgeCluster;
  }


  public NetworkDhcpServerCreateConfig preferredEdgeNode1(String preferredEdgeNode1) {
    
    this.preferredEdgeNode1 = preferredEdgeNode1;
    return this;
  }

   /**
   * Active Edge Node Options obtained by calling option source with :optionSource &#x3D; nsxtEdgeNodes and networkServerId param
   * @return preferredEdgeNode1
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "Active Edge Node Options obtained by calling option source with :optionSource = nsxtEdgeNodes and networkServerId param")

  public String getPreferredEdgeNode1() {
    return preferredEdgeNode1;
  }


  public void setPreferredEdgeNode1(String preferredEdgeNode1) {
    this.preferredEdgeNode1 = preferredEdgeNode1;
  }


  public NetworkDhcpServerCreateConfig preferredEdgeNode2(String preferredEdgeNode2) {
    
    this.preferredEdgeNode2 = preferredEdgeNode2;
    return this;
  }

   /**
   * Standby Edge Node Options obtained by calling option source with optionSource &#x3D; nsxtEdgeNodes and networkServerId param
   * @return preferredEdgeNode2
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "Standby Edge Node Options obtained by calling option source with optionSource = nsxtEdgeNodes and networkServerId param")

  public String getPreferredEdgeNode2() {
    return preferredEdgeNode2;
  }


  public void setPreferredEdgeNode2(String preferredEdgeNode2) {
    this.preferredEdgeNode2 = preferredEdgeNode2;
  }


  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (o == null || getClass() != o.getClass()) {
      return false;
    }
    NetworkDhcpServerCreateConfig networkDhcpServerCreateConfig = (NetworkDhcpServerCreateConfig) o;
    return Objects.equals(this.edgeCluster, networkDhcpServerCreateConfig.edgeCluster) &&
        Objects.equals(this.preferredEdgeNode1, networkDhcpServerCreateConfig.preferredEdgeNode1) &&
        Objects.equals(this.preferredEdgeNode2, networkDhcpServerCreateConfig.preferredEdgeNode2);
  }

  @Override
  public int hashCode() {
    return Objects.hash(edgeCluster, preferredEdgeNode1, preferredEdgeNode2);
  }


  @Override
  public String toString() {
    StringBuilder sb = new StringBuilder();
    sb.append("class NetworkDhcpServerCreateConfig {\n");
    sb.append("    edgeCluster: ").append(toIndentedString(edgeCluster)).append("\n");
    sb.append("    preferredEdgeNode1: ").append(toIndentedString(preferredEdgeNode1)).append("\n");
    sb.append("    preferredEdgeNode2: ").append(toIndentedString(preferredEdgeNode2)).append("\n");
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

