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
 * RolePermissionDefaultReportType
 */
@javax.annotation.Generated(value = "org.openapitools.codegen.languages.JavaClientCodegen", date = "2023-08-21T11:26:47.028Z[GMT]")
public class RolePermissionDefaultReportType {
  /**
   * &#x60;ReportTypes&#x60; is the code for Default Report Type Access
   */
  @JsonAdapter(PermissionCodeEnum.Adapter.class)
  public enum PermissionCodeEnum {
    REPORTTYPES("ReportTypes");

    private String value;

    PermissionCodeEnum(String value) {
      this.value = value;
    }

    public String getValue() {
      return value;
    }

    @Override
    public String toString() {
      return String.valueOf(value);
    }

    public static PermissionCodeEnum fromValue(String value) {
      for (PermissionCodeEnum b : PermissionCodeEnum.values()) {
        if (b.value.equals(value)) {
          return b;
        }
      }
      throw new IllegalArgumentException("Unexpected value '" + value + "'");
    }

    public static class Adapter extends TypeAdapter<PermissionCodeEnum> {
      @Override
      public void write(final JsonWriter jsonWriter, final PermissionCodeEnum enumeration) throws IOException {
        jsonWriter.value(enumeration.getValue());
      }

      @Override
      public PermissionCodeEnum read(final JsonReader jsonReader) throws IOException {
        String value =  jsonReader.nextString();
        return PermissionCodeEnum.fromValue(value);
      }
    }
  }

  public static final String SERIALIZED_NAME_PERMISSION_CODE = "permissionCode";
  @SerializedName(SERIALIZED_NAME_PERMISSION_CODE)
  private PermissionCodeEnum permissionCode;

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


  public RolePermissionDefaultReportType permissionCode(PermissionCodeEnum permissionCode) {
    
    this.permissionCode = permissionCode;
    return this;
  }

   /**
   * &#x60;ReportTypes&#x60; is the code for Default Report Type Access
   * @return permissionCode
  **/
  @ApiModelProperty(required = true, value = "`ReportTypes` is the code for Default Report Type Access")

  public PermissionCodeEnum getPermissionCode() {
    return permissionCode;
  }


  public void setPermissionCode(PermissionCodeEnum permissionCode) {
    this.permissionCode = permissionCode;
  }


  public RolePermissionDefaultReportType access(AccessEnum access) {
    
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
    RolePermissionDefaultReportType rolePermissionDefaultReportType = (RolePermissionDefaultReportType) o;
    return Objects.equals(this.permissionCode, rolePermissionDefaultReportType.permissionCode) &&
        Objects.equals(this.access, rolePermissionDefaultReportType.access);
  }

  @Override
  public int hashCode() {
    return Objects.hash(permissionCode, access);
  }


  @Override
  public String toString() {
    StringBuilder sb = new StringBuilder();
    sb.append("class RolePermissionDefaultReportType {\n");
    sb.append("    permissionCode: ").append(toIndentedString(permissionCode)).append("\n");
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

