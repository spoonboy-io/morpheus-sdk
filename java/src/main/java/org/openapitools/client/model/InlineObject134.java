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
import org.openapitools.client.model.ApiLoadBalancersLoadBalancerIdProfilesLoadBalancerProfile;

/**
 * InlineObject134
 */
@javax.annotation.Generated(value = "org.openapitools.codegen.languages.JavaClientCodegen", date = "2023-08-21T11:26:47.028Z[GMT]")
public class InlineObject134 {
  public static final String SERIALIZED_NAME_LOAD_BALANCER_PROFILE = "loadBalancerProfile";
  @SerializedName(SERIALIZED_NAME_LOAD_BALANCER_PROFILE)
  private ApiLoadBalancersLoadBalancerIdProfilesLoadBalancerProfile loadBalancerProfile;


  public InlineObject134 loadBalancerProfile(ApiLoadBalancersLoadBalancerIdProfilesLoadBalancerProfile loadBalancerProfile) {
    
    this.loadBalancerProfile = loadBalancerProfile;
    return this;
  }

   /**
   * Get loadBalancerProfile
   * @return loadBalancerProfile
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public ApiLoadBalancersLoadBalancerIdProfilesLoadBalancerProfile getLoadBalancerProfile() {
    return loadBalancerProfile;
  }


  public void setLoadBalancerProfile(ApiLoadBalancersLoadBalancerIdProfilesLoadBalancerProfile loadBalancerProfile) {
    this.loadBalancerProfile = loadBalancerProfile;
  }


  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (o == null || getClass() != o.getClass()) {
      return false;
    }
    InlineObject134 inlineObject134 = (InlineObject134) o;
    return Objects.equals(this.loadBalancerProfile, inlineObject134.loadBalancerProfile);
  }

  @Override
  public int hashCode() {
    return Objects.hash(loadBalancerProfile);
  }


  @Override
  public String toString() {
    StringBuilder sb = new StringBuilder();
    sb.append("class InlineObject134 {\n");
    sb.append("    loadBalancerProfile: ").append(toIndentedString(loadBalancerProfile)).append("\n");
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

