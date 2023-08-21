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
 * Metadata about the returned array of results, provides pagination information.
 */
@ApiModel(description = "Metadata about the returned array of results, provides pagination information.")
@javax.annotation.Generated(value = "org.openapitools.codegen.languages.JavaClientCodegen", date = "2023-08-21T11:26:47.028Z[GMT]")
public class MetaObject {
  public static final String SERIALIZED_NAME_OFFSET = "offset";
  @SerializedName(SERIALIZED_NAME_OFFSET)
  private Long offset = 0l;

  public static final String SERIALIZED_NAME_MAX = "max";
  @SerializedName(SERIALIZED_NAME_MAX)
  private Long max = 25l;

  public static final String SERIALIZED_NAME_SIZE = "size";
  @SerializedName(SERIALIZED_NAME_SIZE)
  private Long size = 0l;

  public static final String SERIALIZED_NAME_TOTAL = "total";
  @SerializedName(SERIALIZED_NAME_TOTAL)
  private Long total = 0l;


  public MetaObject offset(Long offset) {
    
    this.offset = offset;
    return this;
  }

   /**
   * Offset records, the number of records to skip, can be used to paginate over results.
   * @return offset
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "Offset records, the number of records to skip, can be used to paginate over results.")

  public Long getOffset() {
    return offset;
  }


  public void setOffset(Long offset) {
    this.offset = offset;
  }


  public MetaObject max(Long max) {
    
    this.max = max;
    return this;
  }

   /**
   * Max size, the maximum number of records to include in the response.
   * @return max
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "Max size, the maximum number of records to include in the response.")

  public Long getMax() {
    return max;
  }


  public void setMax(Long max) {
    this.max = max;
  }


  public MetaObject size(Long size) {
    
    this.size = size;
    return this;
  }

   /**
   * Number of records returned in the response
   * @return size
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "Number of records returned in the response")

  public Long getSize() {
    return size;
  }


  public void setSize(Long size) {
    this.size = size;
  }


  public MetaObject total(Long total) {
    
    this.total = total;
    return this;
  }

   /**
   * Total number of records found
   * @return total
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "Total number of records found")

  public Long getTotal() {
    return total;
  }


  public void setTotal(Long total) {
    this.total = total;
  }


  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (o == null || getClass() != o.getClass()) {
      return false;
    }
    MetaObject metaObject = (MetaObject) o;
    return Objects.equals(this.offset, metaObject.offset) &&
        Objects.equals(this.max, metaObject.max) &&
        Objects.equals(this.size, metaObject.size) &&
        Objects.equals(this.total, metaObject.total);
  }

  @Override
  public int hashCode() {
    return Objects.hash(offset, max, size, total);
  }


  @Override
  public String toString() {
    StringBuilder sb = new StringBuilder();
    sb.append("class MetaObject {\n");
    sb.append("    offset: ").append(toIndentedString(offset)).append("\n");
    sb.append("    max: ").append(toIndentedString(max)).append("\n");
    sb.append("    size: ").append(toIndentedString(size)).append("\n");
    sb.append("    total: ").append(toIndentedString(total)).append("\n");
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

