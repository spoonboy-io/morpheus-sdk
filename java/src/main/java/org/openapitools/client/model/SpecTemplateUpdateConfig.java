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
import org.openapitools.client.model.SpecTemplateUpdateConfigCloudformation;

/**
 * The Cloud Formation type supports some additional configuration parameters
 */
@ApiModel(description = "The Cloud Formation type supports some additional configuration parameters")
@javax.annotation.Generated(value = "org.openapitools.codegen.languages.JavaClientCodegen", date = "2023-08-21T11:26:47.028Z[GMT]")
public class SpecTemplateUpdateConfig {
  public static final String SERIALIZED_NAME_CLOUDFORMATION = "cloudformation";
  @SerializedName(SERIALIZED_NAME_CLOUDFORMATION)
  private SpecTemplateUpdateConfigCloudformation cloudformation;


  public SpecTemplateUpdateConfig cloudformation(SpecTemplateUpdateConfigCloudformation cloudformation) {
    
    this.cloudformation = cloudformation;
    return this;
  }

   /**
   * Get cloudformation
   * @return cloudformation
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public SpecTemplateUpdateConfigCloudformation getCloudformation() {
    return cloudformation;
  }


  public void setCloudformation(SpecTemplateUpdateConfigCloudformation cloudformation) {
    this.cloudformation = cloudformation;
  }


  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (o == null || getClass() != o.getClass()) {
      return false;
    }
    SpecTemplateUpdateConfig specTemplateUpdateConfig = (SpecTemplateUpdateConfig) o;
    return Objects.equals(this.cloudformation, specTemplateUpdateConfig.cloudformation);
  }

  @Override
  public int hashCode() {
    return Objects.hash(cloudformation);
  }


  @Override
  public String toString() {
    StringBuilder sb = new StringBuilder();
    sb.append("class SpecTemplateUpdateConfig {\n");
    sb.append("    cloudformation: ").append(toIndentedString(cloudformation)).append("\n");
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

