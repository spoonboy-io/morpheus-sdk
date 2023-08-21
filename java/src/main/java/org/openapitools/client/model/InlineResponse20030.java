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
import org.openapitools.client.model.ClusterMasters;

/**
 * InlineResponse20030
 */
@javax.annotation.Generated(value = "org.openapitools.codegen.languages.JavaClientCodegen", date = "2023-08-21T11:26:47.028Z[GMT]")
public class InlineResponse20030 {
  public static final String SERIALIZED_NAME_MASTERS = "masters";
  @SerializedName(SERIALIZED_NAME_MASTERS)
  private List<ClusterMasters> masters = null;


  public InlineResponse20030 masters(List<ClusterMasters> masters) {
    
    this.masters = masters;
    return this;
  }

  public InlineResponse20030 addMastersItem(ClusterMasters mastersItem) {
    if (this.masters == null) {
      this.masters = new ArrayList<ClusterMasters>();
    }
    this.masters.add(mastersItem);
    return this;
  }

   /**
   * Get masters
   * @return masters
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public List<ClusterMasters> getMasters() {
    return masters;
  }


  public void setMasters(List<ClusterMasters> masters) {
    this.masters = masters;
  }


  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (o == null || getClass() != o.getClass()) {
      return false;
    }
    InlineResponse20030 inlineResponse20030 = (InlineResponse20030) o;
    return Objects.equals(this.masters, inlineResponse20030.masters);
  }

  @Override
  public int hashCode() {
    return Objects.hash(masters);
  }


  @Override
  public String toString() {
    StringBuilder sb = new StringBuilder();
    sb.append("class InlineResponse20030 {\n");
    sb.append("    masters: ").append(toIndentedString(masters)).append("\n");
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
