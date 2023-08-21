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
 * BudgetStatsCurrent
 */
@javax.annotation.Generated(value = "org.openapitools.codegen.languages.JavaClientCodegen", date = "2023-08-21T11:26:47.028Z[GMT]")
public class BudgetStatsCurrent {
  public static final String SERIALIZED_NAME_ESTIMATED_COST = "estimatedCost";
  @SerializedName(SERIALIZED_NAME_ESTIMATED_COST)
  private String estimatedCost;

  public static final String SERIALIZED_NAME_LAST_COST = "lastCost";
  @SerializedName(SERIALIZED_NAME_LAST_COST)
  private String lastCost;


  public BudgetStatsCurrent estimatedCost(String estimatedCost) {
    
    this.estimatedCost = estimatedCost;
    return this;
  }

   /**
   * Get estimatedCost
   * @return estimatedCost
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public String getEstimatedCost() {
    return estimatedCost;
  }


  public void setEstimatedCost(String estimatedCost) {
    this.estimatedCost = estimatedCost;
  }


  public BudgetStatsCurrent lastCost(String lastCost) {
    
    this.lastCost = lastCost;
    return this;
  }

   /**
   * Get lastCost
   * @return lastCost
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public String getLastCost() {
    return lastCost;
  }


  public void setLastCost(String lastCost) {
    this.lastCost = lastCost;
  }


  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (o == null || getClass() != o.getClass()) {
      return false;
    }
    BudgetStatsCurrent budgetStatsCurrent = (BudgetStatsCurrent) o;
    return Objects.equals(this.estimatedCost, budgetStatsCurrent.estimatedCost) &&
        Objects.equals(this.lastCost, budgetStatsCurrent.lastCost);
  }

  @Override
  public int hashCode() {
    return Objects.hash(estimatedCost, lastCost);
  }


  @Override
  public String toString() {
    StringBuilder sb = new StringBuilder();
    sb.append("class BudgetStatsCurrent {\n");
    sb.append("    estimatedCost: ").append(toIndentedString(estimatedCost)).append("\n");
    sb.append("    lastCost: ").append(toIndentedString(lastCost)).append("\n");
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
