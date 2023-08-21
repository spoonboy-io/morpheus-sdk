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
 * ClusterServerCreatePlan
 */
@javax.annotation.Generated(value = "org.openapitools.codegen.languages.JavaClientCodegen", date = "2023-08-21T11:26:47.028Z[GMT]")
public class ClusterServerCreatePlan {
  public static final String SERIALIZED_NAME_ID = "id";
  @SerializedName(SERIALIZED_NAME_ID)
  private Long id;

  public static final String SERIALIZED_NAME_CODE = "code";
  @SerializedName(SERIALIZED_NAME_CODE)
  private String code;

  public static final String SERIALIZED_NAME_OPTIONS = "options";
  @SerializedName(SERIALIZED_NAME_OPTIONS)
  private Object options;


  public ClusterServerCreatePlan id(Long id) {
    
    this.id = id;
    return this;
  }

   /**
   * The id for the memory and storage option pre-configured within Morpheus.
   * @return id
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "The id for the memory and storage option pre-configured within Morpheus.")

  public Long getId() {
    return id;
  }


  public void setId(Long id) {
    this.id = id;
  }


  public ClusterServerCreatePlan code(String code) {
    
    this.code = code;
    return this;
  }

   /**
   * The code for the memory and storage option pre-configured within Morpheus.
   * @return code
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "The code for the memory and storage option pre-configured within Morpheus.")

  public String getCode() {
    return code;
  }


  public void setCode(String code) {
    this.code = code;
  }


  public ClusterServerCreatePlan options(Object options) {
    
    this.options = options;
    return this;
  }

   /**
   * Map of custom options depending on selected service plan . An example would be &#x60;maxMemory&#x60;, or &#x60;maxCores&#x60;.
   * @return options
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "Map of custom options depending on selected service plan . An example would be `maxMemory`, or `maxCores`.")

  public Object getOptions() {
    return options;
  }


  public void setOptions(Object options) {
    this.options = options;
  }


  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (o == null || getClass() != o.getClass()) {
      return false;
    }
    ClusterServerCreatePlan clusterServerCreatePlan = (ClusterServerCreatePlan) o;
    return Objects.equals(this.id, clusterServerCreatePlan.id) &&
        Objects.equals(this.code, clusterServerCreatePlan.code) &&
        Objects.equals(this.options, clusterServerCreatePlan.options);
  }

  @Override
  public int hashCode() {
    return Objects.hash(id, code, options);
  }


  @Override
  public String toString() {
    StringBuilder sb = new StringBuilder();
    sb.append("class ClusterServerCreatePlan {\n");
    sb.append("    id: ").append(toIndentedString(id)).append("\n");
    sb.append("    code: ").append(toIndentedString(code)).append("\n");
    sb.append("    options: ").append(toIndentedString(options)).append("\n");
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
