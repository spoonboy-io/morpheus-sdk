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
import org.openapitools.client.model.InstanceUpdateInstance;
import org.openapitools.client.model.InstancesConfigCustomOptions;

/**
 * InstanceUpdate
 */
@javax.annotation.Generated(value = "org.openapitools.codegen.languages.JavaClientCodegen", date = "2023-08-21T11:26:47.028Z[GMT]")
public class InstanceUpdate {
  public static final String SERIALIZED_NAME_INSTANCE = "instance";
  @SerializedName(SERIALIZED_NAME_INSTANCE)
  private InstanceUpdateInstance instance;

  public static final String SERIALIZED_NAME_CONFIG = "config";
  @SerializedName(SERIALIZED_NAME_CONFIG)
  private InstancesConfigCustomOptions config;


  public InstanceUpdate instance(InstanceUpdateInstance instance) {
    
    this.instance = instance;
    return this;
  }

   /**
   * Get instance
   * @return instance
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public InstanceUpdateInstance getInstance() {
    return instance;
  }


  public void setInstance(InstanceUpdateInstance instance) {
    this.instance = instance;
  }


  public InstanceUpdate config(InstancesConfigCustomOptions config) {
    
    this.config = config;
    return this;
  }

   /**
   * Get config
   * @return config
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public InstancesConfigCustomOptions getConfig() {
    return config;
  }


  public void setConfig(InstancesConfigCustomOptions config) {
    this.config = config;
  }


  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (o == null || getClass() != o.getClass()) {
      return false;
    }
    InstanceUpdate instanceUpdate = (InstanceUpdate) o;
    return Objects.equals(this.instance, instanceUpdate.instance) &&
        Objects.equals(this.config, instanceUpdate.config);
  }

  @Override
  public int hashCode() {
    return Objects.hash(instance, config);
  }


  @Override
  public String toString() {
    StringBuilder sb = new StringBuilder();
    sb.append("class InstanceUpdate {\n");
    sb.append("    instance: ").append(toIndentedString(instance)).append("\n");
    sb.append("    config: ").append(toIndentedString(config)).append("\n");
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

