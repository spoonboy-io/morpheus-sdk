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
import org.openapitools.client.model.Vdi;

/**
 * InlineResponse200164
 */
@javax.annotation.Generated(value = "org.openapitools.codegen.languages.JavaClientCodegen", date = "2023-08-21T11:26:47.028Z[GMT]")
public class InlineResponse200164 {
  public static final String SERIALIZED_NAME_DESKTOP = "desktop";
  @SerializedName(SERIALIZED_NAME_DESKTOP)
  private Vdi desktop;


  public InlineResponse200164 desktop(Vdi desktop) {
    
    this.desktop = desktop;
    return this;
  }

   /**
   * Get desktop
   * @return desktop
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public Vdi getDesktop() {
    return desktop;
  }


  public void setDesktop(Vdi desktop) {
    this.desktop = desktop;
  }


  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (o == null || getClass() != o.getClass()) {
      return false;
    }
    InlineResponse200164 inlineResponse200164 = (InlineResponse200164) o;
    return Objects.equals(this.desktop, inlineResponse200164.desktop);
  }

  @Override
  public int hashCode() {
    return Objects.hash(desktop);
  }


  @Override
  public String toString() {
    StringBuilder sb = new StringBuilder();
    sb.append("class InlineResponse200164 {\n");
    sb.append("    desktop: ").append(toIndentedString(desktop)).append("\n");
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

