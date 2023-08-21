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
import org.openapitools.client.model.PreseedScriptsCreate;

/**
 * InlineObject196
 */
@javax.annotation.Generated(value = "org.openapitools.codegen.languages.JavaClientCodegen", date = "2023-08-21T11:26:47.028Z[GMT]")
public class InlineObject196 {
  public static final String SERIALIZED_NAME_PRESEED_SCRIPT = "preseedScript";
  @SerializedName(SERIALIZED_NAME_PRESEED_SCRIPT)
  private PreseedScriptsCreate preseedScript;


  public InlineObject196 preseedScript(PreseedScriptsCreate preseedScript) {
    
    this.preseedScript = preseedScript;
    return this;
  }

   /**
   * Get preseedScript
   * @return preseedScript
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public PreseedScriptsCreate getPreseedScript() {
    return preseedScript;
  }


  public void setPreseedScript(PreseedScriptsCreate preseedScript) {
    this.preseedScript = preseedScript;
  }


  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (o == null || getClass() != o.getClass()) {
      return false;
    }
    InlineObject196 inlineObject196 = (InlineObject196) o;
    return Objects.equals(this.preseedScript, inlineObject196.preseedScript);
  }

  @Override
  public int hashCode() {
    return Objects.hash(preseedScript);
  }


  @Override
  public String toString() {
    StringBuilder sb = new StringBuilder();
    sb.append("class InlineObject196 {\n");
    sb.append("    preseedScript: ").append(toIndentedString(preseedScript)).append("\n");
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

