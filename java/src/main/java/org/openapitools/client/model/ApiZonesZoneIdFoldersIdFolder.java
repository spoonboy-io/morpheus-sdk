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
import org.openapitools.client.model.ApiZonesZoneIdDataStoresIdDatastoreResourcePermissions;
import org.openapitools.client.model.ApiZonesZoneIdFoldersIdFolderTenantPermissions;

/**
 * ApiZonesZoneIdFoldersIdFolder
 */
@javax.annotation.Generated(value = "org.openapitools.codegen.languages.JavaClientCodegen", date = "2023-08-21T11:26:47.028Z[GMT]")
public class ApiZonesZoneIdFoldersIdFolder {
  public static final String SERIALIZED_NAME_DEFAULT_FOLDER = "defaultFolder";
  @SerializedName(SERIALIZED_NAME_DEFAULT_FOLDER)
  private Boolean defaultFolder = false;

  public static final String SERIALIZED_NAME_DEFAULT_IMAGE = "defaultImage";
  @SerializedName(SERIALIZED_NAME_DEFAULT_IMAGE)
  private Boolean defaultImage = false;

  public static final String SERIALIZED_NAME_ACTIVE = "active";
  @SerializedName(SERIALIZED_NAME_ACTIVE)
  private Boolean active;

  /**
   * Setting &#x60;private&#x60; or &#x60;public&#x60;
   */
  @JsonAdapter(VisibilityEnum.Adapter.class)
  public enum VisibilityEnum {
    PUBLIC("public"),
    
    PRIVATE("private");

    private String value;

    VisibilityEnum(String value) {
      this.value = value;
    }

    public String getValue() {
      return value;
    }

    @Override
    public String toString() {
      return String.valueOf(value);
    }

    public static VisibilityEnum fromValue(String value) {
      for (VisibilityEnum b : VisibilityEnum.values()) {
        if (b.value.equals(value)) {
          return b;
        }
      }
      throw new IllegalArgumentException("Unexpected value '" + value + "'");
    }

    public static class Adapter extends TypeAdapter<VisibilityEnum> {
      @Override
      public void write(final JsonWriter jsonWriter, final VisibilityEnum enumeration) throws IOException {
        jsonWriter.value(enumeration.getValue());
      }

      @Override
      public VisibilityEnum read(final JsonReader jsonReader) throws IOException {
        String value =  jsonReader.nextString();
        return VisibilityEnum.fromValue(value);
      }
    }
  }

  public static final String SERIALIZED_NAME_VISIBILITY = "visibility";
  @SerializedName(SERIALIZED_NAME_VISIBILITY)
  private VisibilityEnum visibility = VisibilityEnum.PRIVATE;

  public static final String SERIALIZED_NAME_TENANT_PERMISSIONS = "tenantPermissions";
  @SerializedName(SERIALIZED_NAME_TENANT_PERMISSIONS)
  private List<ApiZonesZoneIdFoldersIdFolderTenantPermissions> tenantPermissions = null;

  public static final String SERIALIZED_NAME_RESOURCE_PERMISSIONS = "resourcePermissions";
  @SerializedName(SERIALIZED_NAME_RESOURCE_PERMISSIONS)
  private ApiZonesZoneIdDataStoresIdDatastoreResourcePermissions resourcePermissions;


  public ApiZonesZoneIdFoldersIdFolder defaultFolder(Boolean defaultFolder) {
    
    this.defaultFolder = defaultFolder;
    return this;
  }

   /**
   * Get defaultFolder
   * @return defaultFolder
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public Boolean getDefaultFolder() {
    return defaultFolder;
  }


  public void setDefaultFolder(Boolean defaultFolder) {
    this.defaultFolder = defaultFolder;
  }


  public ApiZonesZoneIdFoldersIdFolder defaultImage(Boolean defaultImage) {
    
    this.defaultImage = defaultImage;
    return this;
  }

   /**
   * Get defaultImage
   * @return defaultImage
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public Boolean getDefaultImage() {
    return defaultImage;
  }


  public void setDefaultImage(Boolean defaultImage) {
    this.defaultImage = defaultImage;
  }


  public ApiZonesZoneIdFoldersIdFolder active(Boolean active) {
    
    this.active = active;
    return this;
  }

   /**
   * Activate &#x60;true&#x60; or disable &#x60;false&#x60; the folder
   * @return active
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "Activate `true` or disable `false` the folder")

  public Boolean getActive() {
    return active;
  }


  public void setActive(Boolean active) {
    this.active = active;
  }


  public ApiZonesZoneIdFoldersIdFolder visibility(VisibilityEnum visibility) {
    
    this.visibility = visibility;
    return this;
  }

   /**
   * Setting &#x60;private&#x60; or &#x60;public&#x60;
   * @return visibility
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "Setting `private` or `public`")

  public VisibilityEnum getVisibility() {
    return visibility;
  }


  public void setVisibility(VisibilityEnum visibility) {
    this.visibility = visibility;
  }


  public ApiZonesZoneIdFoldersIdFolder tenantPermissions(List<ApiZonesZoneIdFoldersIdFolderTenantPermissions> tenantPermissions) {
    
    this.tenantPermissions = tenantPermissions;
    return this;
  }

  public ApiZonesZoneIdFoldersIdFolder addTenantPermissionsItem(ApiZonesZoneIdFoldersIdFolderTenantPermissions tenantPermissionsItem) {
    if (this.tenantPermissions == null) {
      this.tenantPermissions = new ArrayList<ApiZonesZoneIdFoldersIdFolderTenantPermissions>();
    }
    this.tenantPermissions.add(tenantPermissionsItem);
    return this;
  }

   /**
   * Get tenantPermissions
   * @return tenantPermissions
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public List<ApiZonesZoneIdFoldersIdFolderTenantPermissions> getTenantPermissions() {
    return tenantPermissions;
  }


  public void setTenantPermissions(List<ApiZonesZoneIdFoldersIdFolderTenantPermissions> tenantPermissions) {
    this.tenantPermissions = tenantPermissions;
  }


  public ApiZonesZoneIdFoldersIdFolder resourcePermissions(ApiZonesZoneIdDataStoresIdDatastoreResourcePermissions resourcePermissions) {
    
    this.resourcePermissions = resourcePermissions;
    return this;
  }

   /**
   * Get resourcePermissions
   * @return resourcePermissions
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public ApiZonesZoneIdDataStoresIdDatastoreResourcePermissions getResourcePermissions() {
    return resourcePermissions;
  }


  public void setResourcePermissions(ApiZonesZoneIdDataStoresIdDatastoreResourcePermissions resourcePermissions) {
    this.resourcePermissions = resourcePermissions;
  }


  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (o == null || getClass() != o.getClass()) {
      return false;
    }
    ApiZonesZoneIdFoldersIdFolder apiZonesZoneIdFoldersIdFolder = (ApiZonesZoneIdFoldersIdFolder) o;
    return Objects.equals(this.defaultFolder, apiZonesZoneIdFoldersIdFolder.defaultFolder) &&
        Objects.equals(this.defaultImage, apiZonesZoneIdFoldersIdFolder.defaultImage) &&
        Objects.equals(this.active, apiZonesZoneIdFoldersIdFolder.active) &&
        Objects.equals(this.visibility, apiZonesZoneIdFoldersIdFolder.visibility) &&
        Objects.equals(this.tenantPermissions, apiZonesZoneIdFoldersIdFolder.tenantPermissions) &&
        Objects.equals(this.resourcePermissions, apiZonesZoneIdFoldersIdFolder.resourcePermissions);
  }

  @Override
  public int hashCode() {
    return Objects.hash(defaultFolder, defaultImage, active, visibility, tenantPermissions, resourcePermissions);
  }


  @Override
  public String toString() {
    StringBuilder sb = new StringBuilder();
    sb.append("class ApiZonesZoneIdFoldersIdFolder {\n");
    sb.append("    defaultFolder: ").append(toIndentedString(defaultFolder)).append("\n");
    sb.append("    defaultImage: ").append(toIndentedString(defaultImage)).append("\n");
    sb.append("    active: ").append(toIndentedString(active)).append("\n");
    sb.append("    visibility: ").append(toIndentedString(visibility)).append("\n");
    sb.append("    tenantPermissions: ").append(toIndentedString(tenantPermissions)).append("\n");
    sb.append("    resourcePermissions: ").append(toIndentedString(resourcePermissions)).append("\n");
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

