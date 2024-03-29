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
 * StorageType
 */
@javax.annotation.Generated(value = "org.openapitools.codegen.languages.JavaClientCodegen", date = "2023-08-21T11:26:47.028Z[GMT]")
public class StorageType {
  public static final String SERIALIZED_NAME_ID = "id";
  @SerializedName(SERIALIZED_NAME_ID)
  private Long id;

  public static final String SERIALIZED_NAME_CODE = "code";
  @SerializedName(SERIALIZED_NAME_CODE)
  private String code;

  public static final String SERIALIZED_NAME_NAME = "name";
  @SerializedName(SERIALIZED_NAME_NAME)
  private String name;

  public static final String SERIALIZED_NAME_DISPLAY_ORDER = "displayOrder";
  @SerializedName(SERIALIZED_NAME_DISPLAY_ORDER)
  private Long displayOrder;

  public static final String SERIALIZED_NAME_DEFAULT_TYPE = "defaultType";
  @SerializedName(SERIALIZED_NAME_DEFAULT_TYPE)
  private Boolean defaultType;

  public static final String SERIALIZED_NAME_CUSTOM_LABEL = "customLabel";
  @SerializedName(SERIALIZED_NAME_CUSTOM_LABEL)
  private Boolean customLabel;

  public static final String SERIALIZED_NAME_CUSTOM_SIZE = "customSize";
  @SerializedName(SERIALIZED_NAME_CUSTOM_SIZE)
  private Boolean customSize;

  public static final String SERIALIZED_NAME_CUSTOM_SIZE_OPTIONS = "customSizeOptions";
  @SerializedName(SERIALIZED_NAME_CUSTOM_SIZE_OPTIONS)
  private String customSizeOptions;


  public StorageType id(Long id) {
    
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


  public StorageType code(String code) {
    
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


  public StorageType name(String name) {
    
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


  public StorageType displayOrder(Long displayOrder) {
    
    this.displayOrder = displayOrder;
    return this;
  }

   /**
   * Get displayOrder
   * @return displayOrder
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public Long getDisplayOrder() {
    return displayOrder;
  }


  public void setDisplayOrder(Long displayOrder) {
    this.displayOrder = displayOrder;
  }


  public StorageType defaultType(Boolean defaultType) {
    
    this.defaultType = defaultType;
    return this;
  }

   /**
   * Get defaultType
   * @return defaultType
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public Boolean getDefaultType() {
    return defaultType;
  }


  public void setDefaultType(Boolean defaultType) {
    this.defaultType = defaultType;
  }


  public StorageType customLabel(Boolean customLabel) {
    
    this.customLabel = customLabel;
    return this;
  }

   /**
   * Get customLabel
   * @return customLabel
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public Boolean getCustomLabel() {
    return customLabel;
  }


  public void setCustomLabel(Boolean customLabel) {
    this.customLabel = customLabel;
  }


  public StorageType customSize(Boolean customSize) {
    
    this.customSize = customSize;
    return this;
  }

   /**
   * Get customSize
   * @return customSize
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public Boolean getCustomSize() {
    return customSize;
  }


  public void setCustomSize(Boolean customSize) {
    this.customSize = customSize;
  }


  public StorageType customSizeOptions(String customSizeOptions) {
    
    this.customSizeOptions = customSizeOptions;
    return this;
  }

   /**
   * Get customSizeOptions
   * @return customSizeOptions
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public String getCustomSizeOptions() {
    return customSizeOptions;
  }


  public void setCustomSizeOptions(String customSizeOptions) {
    this.customSizeOptions = customSizeOptions;
  }


  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (o == null || getClass() != o.getClass()) {
      return false;
    }
    StorageType storageType = (StorageType) o;
    return Objects.equals(this.id, storageType.id) &&
        Objects.equals(this.code, storageType.code) &&
        Objects.equals(this.name, storageType.name) &&
        Objects.equals(this.displayOrder, storageType.displayOrder) &&
        Objects.equals(this.defaultType, storageType.defaultType) &&
        Objects.equals(this.customLabel, storageType.customLabel) &&
        Objects.equals(this.customSize, storageType.customSize) &&
        Objects.equals(this.customSizeOptions, storageType.customSizeOptions);
  }

  @Override
  public int hashCode() {
    return Objects.hash(id, code, name, displayOrder, defaultType, customLabel, customSize, customSizeOptions);
  }


  @Override
  public String toString() {
    StringBuilder sb = new StringBuilder();
    sb.append("class StorageType {\n");
    sb.append("    id: ").append(toIndentedString(id)).append("\n");
    sb.append("    code: ").append(toIndentedString(code)).append("\n");
    sb.append("    name: ").append(toIndentedString(name)).append("\n");
    sb.append("    displayOrder: ").append(toIndentedString(displayOrder)).append("\n");
    sb.append("    defaultType: ").append(toIndentedString(defaultType)).append("\n");
    sb.append("    customLabel: ").append(toIndentedString(customLabel)).append("\n");
    sb.append("    customSize: ").append(toIndentedString(customSize)).append("\n");
    sb.append("    customSizeOptions: ").append(toIndentedString(customSizeOptions)).append("\n");
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

