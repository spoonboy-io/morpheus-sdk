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
 * GuidanceVmwareSizingResourceServerOs
 */
@javax.annotation.Generated(value = "org.openapitools.codegen.languages.JavaClientCodegen", date = "2023-08-21T11:26:47.028Z[GMT]")
public class GuidanceVmwareSizingResourceServerOs {
  public static final String SERIALIZED_NAME_ID = "id";
  @SerializedName(SERIALIZED_NAME_ID)
  private Long id;

  public static final String SERIALIZED_NAME_CODE = "code";
  @SerializedName(SERIALIZED_NAME_CODE)
  private String code;

  public static final String SERIALIZED_NAME_NAME = "name";
  @SerializedName(SERIALIZED_NAME_NAME)
  private String name;

  public static final String SERIALIZED_NAME_DESCRIPTION = "description";
  @SerializedName(SERIALIZED_NAME_DESCRIPTION)
  private String description;

  public static final String SERIALIZED_NAME_VENDOR = "vendor";
  @SerializedName(SERIALIZED_NAME_VENDOR)
  private String vendor;

  public static final String SERIALIZED_NAME_CATEGORY = "category";
  @SerializedName(SERIALIZED_NAME_CATEGORY)
  private String category;

  public static final String SERIALIZED_NAME_OS_FAMILY = "osFamily";
  @SerializedName(SERIALIZED_NAME_OS_FAMILY)
  private String osFamily;

  public static final String SERIALIZED_NAME_OS_VERSION = "osVersion";
  @SerializedName(SERIALIZED_NAME_OS_VERSION)
  private String osVersion;

  public static final String SERIALIZED_NAME_BIT_COUNT = "bitCount";
  @SerializedName(SERIALIZED_NAME_BIT_COUNT)
  private Long bitCount;

  public static final String SERIALIZED_NAME_PLATFORM = "platform";
  @SerializedName(SERIALIZED_NAME_PLATFORM)
  private String platform;


  public GuidanceVmwareSizingResourceServerOs id(Long id) {
    
    this.id = id;
    return this;
  }

   /**
   * Get id
   * @return id
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public Long getId() {
    return id;
  }


  public void setId(Long id) {
    this.id = id;
  }


  public GuidanceVmwareSizingResourceServerOs code(String code) {
    
    this.code = code;
    return this;
  }

   /**
   * Get code
   * @return code
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public String getCode() {
    return code;
  }


  public void setCode(String code) {
    this.code = code;
  }


  public GuidanceVmwareSizingResourceServerOs name(String name) {
    
    this.name = name;
    return this;
  }

   /**
   * Get name
   * @return name
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public String getName() {
    return name;
  }


  public void setName(String name) {
    this.name = name;
  }


  public GuidanceVmwareSizingResourceServerOs description(String description) {
    
    this.description = description;
    return this;
  }

   /**
   * Get description
   * @return description
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public String getDescription() {
    return description;
  }


  public void setDescription(String description) {
    this.description = description;
  }


  public GuidanceVmwareSizingResourceServerOs vendor(String vendor) {
    
    this.vendor = vendor;
    return this;
  }

   /**
   * Get vendor
   * @return vendor
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public String getVendor() {
    return vendor;
  }


  public void setVendor(String vendor) {
    this.vendor = vendor;
  }


  public GuidanceVmwareSizingResourceServerOs category(String category) {
    
    this.category = category;
    return this;
  }

   /**
   * Get category
   * @return category
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public String getCategory() {
    return category;
  }


  public void setCategory(String category) {
    this.category = category;
  }


  public GuidanceVmwareSizingResourceServerOs osFamily(String osFamily) {
    
    this.osFamily = osFamily;
    return this;
  }

   /**
   * Get osFamily
   * @return osFamily
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public String getOsFamily() {
    return osFamily;
  }


  public void setOsFamily(String osFamily) {
    this.osFamily = osFamily;
  }


  public GuidanceVmwareSizingResourceServerOs osVersion(String osVersion) {
    
    this.osVersion = osVersion;
    return this;
  }

   /**
   * Get osVersion
   * @return osVersion
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public String getOsVersion() {
    return osVersion;
  }


  public void setOsVersion(String osVersion) {
    this.osVersion = osVersion;
  }


  public GuidanceVmwareSizingResourceServerOs bitCount(Long bitCount) {
    
    this.bitCount = bitCount;
    return this;
  }

   /**
   * Get bitCount
   * @return bitCount
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public Long getBitCount() {
    return bitCount;
  }


  public void setBitCount(Long bitCount) {
    this.bitCount = bitCount;
  }


  public GuidanceVmwareSizingResourceServerOs platform(String platform) {
    
    this.platform = platform;
    return this;
  }

   /**
   * Get platform
   * @return platform
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public String getPlatform() {
    return platform;
  }


  public void setPlatform(String platform) {
    this.platform = platform;
  }


  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (o == null || getClass() != o.getClass()) {
      return false;
    }
    GuidanceVmwareSizingResourceServerOs guidanceVmwareSizingResourceServerOs = (GuidanceVmwareSizingResourceServerOs) o;
    return Objects.equals(this.id, guidanceVmwareSizingResourceServerOs.id) &&
        Objects.equals(this.code, guidanceVmwareSizingResourceServerOs.code) &&
        Objects.equals(this.name, guidanceVmwareSizingResourceServerOs.name) &&
        Objects.equals(this.description, guidanceVmwareSizingResourceServerOs.description) &&
        Objects.equals(this.vendor, guidanceVmwareSizingResourceServerOs.vendor) &&
        Objects.equals(this.category, guidanceVmwareSizingResourceServerOs.category) &&
        Objects.equals(this.osFamily, guidanceVmwareSizingResourceServerOs.osFamily) &&
        Objects.equals(this.osVersion, guidanceVmwareSizingResourceServerOs.osVersion) &&
        Objects.equals(this.bitCount, guidanceVmwareSizingResourceServerOs.bitCount) &&
        Objects.equals(this.platform, guidanceVmwareSizingResourceServerOs.platform);
  }

  @Override
  public int hashCode() {
    return Objects.hash(id, code, name, description, vendor, category, osFamily, osVersion, bitCount, platform);
  }


  @Override
  public String toString() {
    StringBuilder sb = new StringBuilder();
    sb.append("class GuidanceVmwareSizingResourceServerOs {\n");
    sb.append("    id: ").append(toIndentedString(id)).append("\n");
    sb.append("    code: ").append(toIndentedString(code)).append("\n");
    sb.append("    name: ").append(toIndentedString(name)).append("\n");
    sb.append("    description: ").append(toIndentedString(description)).append("\n");
    sb.append("    vendor: ").append(toIndentedString(vendor)).append("\n");
    sb.append("    category: ").append(toIndentedString(category)).append("\n");
    sb.append("    osFamily: ").append(toIndentedString(osFamily)).append("\n");
    sb.append("    osVersion: ").append(toIndentedString(osVersion)).append("\n");
    sb.append("    bitCount: ").append(toIndentedString(bitCount)).append("\n");
    sb.append("    platform: ").append(toIndentedString(platform)).append("\n");
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

