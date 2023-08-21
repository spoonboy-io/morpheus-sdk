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
import java.math.BigDecimal;

/**
 * BillingInstancesPricesUsed
 */
@javax.annotation.Generated(value = "org.openapitools.codegen.languages.JavaClientCodegen", date = "2023-08-21T11:26:47.028Z[GMT]")
public class BillingInstancesPricesUsed {
  public static final String SERIALIZED_NAME_TYPE = "type";
  @SerializedName(SERIALIZED_NAME_TYPE)
  private String type;

  public static final String SERIALIZED_NAME_PRICE_PER_UNIT = "pricePerUnit";
  @SerializedName(SERIALIZED_NAME_PRICE_PER_UNIT)
  private BigDecimal pricePerUnit;

  public static final String SERIALIZED_NAME_COST_PER_UNIT = "costPerUnit";
  @SerializedName(SERIALIZED_NAME_COST_PER_UNIT)
  private BigDecimal costPerUnit;

  public static final String SERIALIZED_NAME_QUANTITY = "quantity";
  @SerializedName(SERIALIZED_NAME_QUANTITY)
  private Long quantity;


  public BillingInstancesPricesUsed type(String type) {
    
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


  public BillingInstancesPricesUsed pricePerUnit(BigDecimal pricePerUnit) {
    
    this.pricePerUnit = pricePerUnit;
    return this;
  }

   /**
   * Get pricePerUnit
   * @return pricePerUnit
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public BigDecimal getPricePerUnit() {
    return pricePerUnit;
  }


  public void setPricePerUnit(BigDecimal pricePerUnit) {
    this.pricePerUnit = pricePerUnit;
  }


  public BillingInstancesPricesUsed costPerUnit(BigDecimal costPerUnit) {
    
    this.costPerUnit = costPerUnit;
    return this;
  }

   /**
   * Get costPerUnit
   * @return costPerUnit
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public BigDecimal getCostPerUnit() {
    return costPerUnit;
  }


  public void setCostPerUnit(BigDecimal costPerUnit) {
    this.costPerUnit = costPerUnit;
  }


  public BillingInstancesPricesUsed quantity(Long quantity) {
    
    this.quantity = quantity;
    return this;
  }

   /**
   * Get quantity
   * @return quantity
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public Long getQuantity() {
    return quantity;
  }


  public void setQuantity(Long quantity) {
    this.quantity = quantity;
  }


  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (o == null || getClass() != o.getClass()) {
      return false;
    }
    BillingInstancesPricesUsed billingInstancesPricesUsed = (BillingInstancesPricesUsed) o;
    return Objects.equals(this.type, billingInstancesPricesUsed.type) &&
        Objects.equals(this.pricePerUnit, billingInstancesPricesUsed.pricePerUnit) &&
        Objects.equals(this.costPerUnit, billingInstancesPricesUsed.costPerUnit) &&
        Objects.equals(this.quantity, billingInstancesPricesUsed.quantity);
  }

  @Override
  public int hashCode() {
    return Objects.hash(type, pricePerUnit, costPerUnit, quantity);
  }


  @Override
  public String toString() {
    StringBuilder sb = new StringBuilder();
    sb.append("class BillingInstancesPricesUsed {\n");
    sb.append("    type: ").append(toIndentedString(type)).append("\n");
    sb.append("    pricePerUnit: ").append(toIndentedString(pricePerUnit)).append("\n");
    sb.append("    costPerUnit: ").append(toIndentedString(costPerUnit)).append("\n");
    sb.append("    quantity: ").append(toIndentedString(quantity)).append("\n");
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

