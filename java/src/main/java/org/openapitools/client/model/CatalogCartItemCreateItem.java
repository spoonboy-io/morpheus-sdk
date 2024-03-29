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
import org.openapitools.client.model.CatalogCartItemCreateItemType;

/**
 * CatalogCartItemCreateItem
 */
@javax.annotation.Generated(value = "org.openapitools.codegen.languages.JavaClientCodegen", date = "2023-08-21T11:26:47.028Z[GMT]")
public class CatalogCartItemCreateItem {
  public static final String SERIALIZED_NAME_TYPE = "type";
  @SerializedName(SERIALIZED_NAME_TYPE)
  private CatalogCartItemCreateItemType type;

  public static final String SERIALIZED_NAME_QUANTITY = "quantity";
  @SerializedName(SERIALIZED_NAME_QUANTITY)
  private Long quantity;

  public static final String SERIALIZED_NAME_CONFIG = "config";
  @SerializedName(SERIALIZED_NAME_CONFIG)
  private Object config;

  public static final String SERIALIZED_NAME_CONTEXT = "context";
  @SerializedName(SERIALIZED_NAME_CONTEXT)
  private String context;

  public static final String SERIALIZED_NAME_TARGET = "target";
  @SerializedName(SERIALIZED_NAME_TARGET)
  private Long target;


  public CatalogCartItemCreateItem type(CatalogCartItemCreateItemType type) {
    
    this.type = type;
    return this;
  }

   /**
   * Get type
   * @return type
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public CatalogCartItemCreateItemType getType() {
    return type;
  }


  public void setType(CatalogCartItemCreateItemType type) {
    this.type = type;
  }


  public CatalogCartItemCreateItem quantity(Long quantity) {
    
    this.quantity = quantity;
    return this;
  }

   /**
   * Quantity for this catalog item. Will be overridden to 1 if quantity not allowed by the item type. 
   * @return quantity
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "Quantity for this catalog item. Will be overridden to 1 if quantity not allowed by the item type. ")

  public Long getQuantity() {
    return quantity;
  }


  public void setQuantity(Long quantity) {
    this.quantity = quantity;
  }


  public CatalogCartItemCreateItem config(Object config) {
    
    this.config = config;
    return this;
  }

   /**
   * Config Object, required options depend on the catalog item type&#39;s associated option types. The values passed in here are injected into the instance config or app spec or workflow script(s) defined by the type. 
   * @return config
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "Config Object, required options depend on the catalog item type's associated option types. The values passed in here are injected into the instance config or app spec or workflow script(s) defined by the type. ")

  public Object getConfig() {
    return config;
  }


  public void setConfig(Object config) {
    this.config = config;
  }


  public CatalogCartItemCreateItem context(String context) {
    
    this.context = context;
    return this;
  }

   /**
   * Context Type for running the workflow, determines if a target resource must be selected. &#x60;instance&#x60;, &#x60;server&#x60;, or &#x60;appliance&#x60;. This may only be passed if the type allows it, usually the type determines the context for the user. Only applies to type &#x60;workflow&#x60;. 
   * @return context
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "Context Type for running the workflow, determines if a target resource must be selected. `instance`, `server`, or `appliance`. This may only be passed if the type allows it, usually the type determines the context for the user. Only applies to type `workflow`. ")

  public String getContext() {
    return context;
  }


  public void setContext(String context) {
    this.context = context;
  }


  public CatalogCartItemCreateItem target(Long target) {
    
    this.target = target;
    return this;
  }

   /**
   * Resource (Instance or Server) ID for context when running the &#x60;workflow&#x60;. Only applies to type &#x60;workflow&#x60; and only required when context is &#x60;instance&#x60; or &#x60;server&#x60;. 
   * @return target
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "Resource (Instance or Server) ID for context when running the `workflow`. Only applies to type `workflow` and only required when context is `instance` or `server`. ")

  public Long getTarget() {
    return target;
  }


  public void setTarget(Long target) {
    this.target = target;
  }


  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (o == null || getClass() != o.getClass()) {
      return false;
    }
    CatalogCartItemCreateItem catalogCartItemCreateItem = (CatalogCartItemCreateItem) o;
    return Objects.equals(this.type, catalogCartItemCreateItem.type) &&
        Objects.equals(this.quantity, catalogCartItemCreateItem.quantity) &&
        Objects.equals(this.config, catalogCartItemCreateItem.config) &&
        Objects.equals(this.context, catalogCartItemCreateItem.context) &&
        Objects.equals(this.target, catalogCartItemCreateItem.target);
  }

  @Override
  public int hashCode() {
    return Objects.hash(type, quantity, config, context, target);
  }


  @Override
  public String toString() {
    StringBuilder sb = new StringBuilder();
    sb.append("class CatalogCartItemCreateItem {\n");
    sb.append("    type: ").append(toIndentedString(type)).append("\n");
    sb.append("    quantity: ").append(toIndentedString(quantity)).append("\n");
    sb.append("    config: ").append(toIndentedString(config)).append("\n");
    sb.append("    context: ").append(toIndentedString(context)).append("\n");
    sb.append("    target: ").append(toIndentedString(target)).append("\n");
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

