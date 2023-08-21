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
import org.openapitools.client.model.OptionType;

/**
 * CatalogType
 */
@javax.annotation.Generated(value = "org.openapitools.codegen.languages.JavaClientCodegen", date = "2023-08-21T11:26:47.028Z[GMT]")
public class CatalogType {
  public static final String SERIALIZED_NAME_ID = "id";
  @SerializedName(SERIALIZED_NAME_ID)
  private Long id;

  public static final String SERIALIZED_NAME_NAME = "name";
  @SerializedName(SERIALIZED_NAME_NAME)
  private String name;

  public static final String SERIALIZED_NAME_DESCRIPTION = "description";
  @SerializedName(SERIALIZED_NAME_DESCRIPTION)
  private String description;

  public static final String SERIALIZED_NAME_TYPE = "type";
  @SerializedName(SERIALIZED_NAME_TYPE)
  private String type;

  public static final String SERIALIZED_NAME_CONTEXT = "context";
  @SerializedName(SERIALIZED_NAME_CONTEXT)
  private String context;

  public static final String SERIALIZED_NAME_FEATURED = "featured";
  @SerializedName(SERIALIZED_NAME_FEATURED)
  private Boolean featured;

  public static final String SERIALIZED_NAME_ALLOW_QUANTITY = "allowQuantity";
  @SerializedName(SERIALIZED_NAME_ALLOW_QUANTITY)
  private Boolean allowQuantity;

  public static final String SERIALIZED_NAME_IMAGE_PATH = "imagePath";
  @SerializedName(SERIALIZED_NAME_IMAGE_PATH)
  private String imagePath;

  public static final String SERIALIZED_NAME_DARK_IMAGE_PATH = "darkImagePath";
  @SerializedName(SERIALIZED_NAME_DARK_IMAGE_PATH)
  private String darkImagePath;

  public static final String SERIALIZED_NAME_OPTION_TYPES = "optionTypes";
  @SerializedName(SERIALIZED_NAME_OPTION_TYPES)
  private List<OptionType> optionTypes = null;


  public CatalogType id(Long id) {
    
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


  public CatalogType name(String name) {
    
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


  public CatalogType description(String description) {
    
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


  public CatalogType type(String type) {
    
    this.type = type;
    return this;
  }

   /**
   * Get type
   * @return type
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public String getType() {
    return type;
  }


  public void setType(String type) {
    this.type = type;
  }


  public CatalogType context(String context) {
    
    this.context = context;
    return this;
  }

   /**
   * Get context
   * @return context
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public String getContext() {
    return context;
  }


  public void setContext(String context) {
    this.context = context;
  }


  public CatalogType featured(Boolean featured) {
    
    this.featured = featured;
    return this;
  }

   /**
   * Get featured
   * @return featured
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public Boolean getFeatured() {
    return featured;
  }


  public void setFeatured(Boolean featured) {
    this.featured = featured;
  }


  public CatalogType allowQuantity(Boolean allowQuantity) {
    
    this.allowQuantity = allowQuantity;
    return this;
  }

   /**
   * Can users order more than one of this item at a time.
   * @return allowQuantity
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "Can users order more than one of this item at a time.")

  public Boolean getAllowQuantity() {
    return allowQuantity;
  }


  public void setAllowQuantity(Boolean allowQuantity) {
    this.allowQuantity = allowQuantity;
  }


  public CatalogType imagePath(String imagePath) {
    
    this.imagePath = imagePath;
    return this;
  }

   /**
   * Get imagePath
   * @return imagePath
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public String getImagePath() {
    return imagePath;
  }


  public void setImagePath(String imagePath) {
    this.imagePath = imagePath;
  }


  public CatalogType darkImagePath(String darkImagePath) {
    
    this.darkImagePath = darkImagePath;
    return this;
  }

   /**
   * Get darkImagePath
   * @return darkImagePath
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public String getDarkImagePath() {
    return darkImagePath;
  }


  public void setDarkImagePath(String darkImagePath) {
    this.darkImagePath = darkImagePath;
  }


  public CatalogType optionTypes(List<OptionType> optionTypes) {
    
    this.optionTypes = optionTypes;
    return this;
  }

  public CatalogType addOptionTypesItem(OptionType optionTypesItem) {
    if (this.optionTypes == null) {
      this.optionTypes = new ArrayList<OptionType>();
    }
    this.optionTypes.add(optionTypesItem);
    return this;
  }

   /**
   * Get optionTypes
   * @return optionTypes
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public List<OptionType> getOptionTypes() {
    return optionTypes;
  }


  public void setOptionTypes(List<OptionType> optionTypes) {
    this.optionTypes = optionTypes;
  }


  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (o == null || getClass() != o.getClass()) {
      return false;
    }
    CatalogType catalogType = (CatalogType) o;
    return Objects.equals(this.id, catalogType.id) &&
        Objects.equals(this.name, catalogType.name) &&
        Objects.equals(this.description, catalogType.description) &&
        Objects.equals(this.type, catalogType.type) &&
        Objects.equals(this.context, catalogType.context) &&
        Objects.equals(this.featured, catalogType.featured) &&
        Objects.equals(this.allowQuantity, catalogType.allowQuantity) &&
        Objects.equals(this.imagePath, catalogType.imagePath) &&
        Objects.equals(this.darkImagePath, catalogType.darkImagePath) &&
        Objects.equals(this.optionTypes, catalogType.optionTypes);
  }

  @Override
  public int hashCode() {
    return Objects.hash(id, name, description, type, context, featured, allowQuantity, imagePath, darkImagePath, optionTypes);
  }


  @Override
  public String toString() {
    StringBuilder sb = new StringBuilder();
    sb.append("class CatalogType {\n");
    sb.append("    id: ").append(toIndentedString(id)).append("\n");
    sb.append("    name: ").append(toIndentedString(name)).append("\n");
    sb.append("    description: ").append(toIndentedString(description)).append("\n");
    sb.append("    type: ").append(toIndentedString(type)).append("\n");
    sb.append("    context: ").append(toIndentedString(context)).append("\n");
    sb.append("    featured: ").append(toIndentedString(featured)).append("\n");
    sb.append("    allowQuantity: ").append(toIndentedString(allowQuantity)).append("\n");
    sb.append("    imagePath: ").append(toIndentedString(imagePath)).append("\n");
    sb.append("    darkImagePath: ").append(toIndentedString(darkImagePath)).append("\n");
    sb.append("    optionTypes: ").append(toIndentedString(optionTypes)).append("\n");
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
