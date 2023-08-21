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
 * RolePermissionCatalogItemType
 */
@javax.annotation.Generated(value = "org.openapitools.codegen.languages.JavaClientCodegen", date = "2023-08-21T11:26:47.028Z[GMT]")
public class RolePermissionCatalogItemType {
  public static final String SERIALIZED_NAME_CATALOG_ITEM_TYPE_ID = "catalogItemTypeId";
  @SerializedName(SERIALIZED_NAME_CATALOG_ITEM_TYPE_ID)
  private Integer catalogItemTypeId;

  /**
   * The new access level.
   */
  @JsonAdapter(AccessEnum.Adapter.class)
  public enum AccessEnum {
    DEFAULT("default"),
    
    FULL("full"),
    
    NONE("none");

    private String value;

    AccessEnum(String value) {
      this.value = value;
    }

    public String getValue() {
      return value;
    }

    @Override
    public String toString() {
      return String.valueOf(value);
    }

    public static AccessEnum fromValue(String value) {
      for (AccessEnum b : AccessEnum.values()) {
        if (b.value.equals(value)) {
          return b;
        }
      }
      throw new IllegalArgumentException("Unexpected value '" + value + "'");
    }

    public static class Adapter extends TypeAdapter<AccessEnum> {
      @Override
      public void write(final JsonWriter jsonWriter, final AccessEnum enumeration) throws IOException {
        jsonWriter.value(enumeration.getValue());
      }

      @Override
      public AccessEnum read(final JsonReader jsonReader) throws IOException {
        String value =  jsonReader.nextString();
        return AccessEnum.fromValue(value);
      }
    }
  }

  public static final String SERIALIZED_NAME_ACCESS = "access";
  @SerializedName(SERIALIZED_NAME_ACCESS)
  private AccessEnum access;


  public RolePermissionCatalogItemType catalogItemTypeId(Integer catalogItemTypeId) {
    
    this.catalogItemTypeId = catalogItemTypeId;
    return this;
  }

   /**
   * &#x60;id&#x60; of the catalog item type
   * @return catalogItemTypeId
  **/
  @ApiModelProperty(required = true, value = "`id` of the catalog item type")

  public Integer getCatalogItemTypeId() {
    return catalogItemTypeId;
  }


  public void setCatalogItemTypeId(Integer catalogItemTypeId) {
    this.catalogItemTypeId = catalogItemTypeId;
  }


  public RolePermissionCatalogItemType access(AccessEnum access) {
    
    this.access = access;
    return this;
  }

   /**
   * The new access level.
   * @return access
  **/
  @ApiModelProperty(required = true, value = "The new access level.")

  public AccessEnum getAccess() {
    return access;
  }


  public void setAccess(AccessEnum access) {
    this.access = access;
  }


  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (o == null || getClass() != o.getClass()) {
      return false;
    }
    RolePermissionCatalogItemType rolePermissionCatalogItemType = (RolePermissionCatalogItemType) o;
    return Objects.equals(this.catalogItemTypeId, rolePermissionCatalogItemType.catalogItemTypeId) &&
        Objects.equals(this.access, rolePermissionCatalogItemType.access);
  }

  @Override
  public int hashCode() {
    return Objects.hash(catalogItemTypeId, access);
  }


  @Override
  public String toString() {
    StringBuilder sb = new StringBuilder();
    sb.append("class RolePermissionCatalogItemType {\n");
    sb.append("    catalogItemTypeId: ").append(toIndentedString(catalogItemTypeId)).append("\n");
    sb.append("    access: ").append(toIndentedString(access)).append("\n");
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

