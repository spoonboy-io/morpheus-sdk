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
import org.openapitools.client.model.LicenseCurrentUsage;
import org.openapitools.client.model.LicenseLicense;

/**
 * Object for logging in to the [Morpheus Hub](https://morpheushub.com) with existing credentials. This is only required for &#x60;hubmode&#x3D;login&#x60;.
 */
@ApiModel(description = "Object for logging in to the [Morpheus Hub](https://morpheushub.com) with existing credentials. This is only required for `hubmode=login`.")
@javax.annotation.Generated(value = "org.openapitools.codegen.languages.JavaClientCodegen", date = "2023-08-21T11:26:47.028Z[GMT]")
public class License {
  public static final String SERIALIZED_NAME_LICENSE = "license";
  @SerializedName(SERIALIZED_NAME_LICENSE)
  private LicenseLicense license;

  public static final String SERIALIZED_NAME_CURRENT_USAGE = "currentUsage";
  @SerializedName(SERIALIZED_NAME_CURRENT_USAGE)
  private LicenseCurrentUsage currentUsage;


  public License license(LicenseLicense license) {
    
    this.license = license;
    return this;
  }

   /**
   * Get license
   * @return license
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public LicenseLicense getLicense() {
    return license;
  }


  public void setLicense(LicenseLicense license) {
    this.license = license;
  }


  public License currentUsage(LicenseCurrentUsage currentUsage) {
    
    this.currentUsage = currentUsage;
    return this;
  }

   /**
   * Get currentUsage
   * @return currentUsage
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public LicenseCurrentUsage getCurrentUsage() {
    return currentUsage;
  }


  public void setCurrentUsage(LicenseCurrentUsage currentUsage) {
    this.currentUsage = currentUsage;
  }


  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (o == null || getClass() != o.getClass()) {
      return false;
    }
    License license = (License) o;
    return Objects.equals(this.license, license.license) &&
        Objects.equals(this.currentUsage, license.currentUsage);
  }

  @Override
  public int hashCode() {
    return Objects.hash(license, currentUsage);
  }


  @Override
  public String toString() {
    StringBuilder sb = new StringBuilder();
    sb.append("class License {\n");
    sb.append("    license: ").append(toIndentedString(license)).append("\n");
    sb.append("    currentUsage: ").append(toIndentedString(currentUsage)).append("\n");
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

