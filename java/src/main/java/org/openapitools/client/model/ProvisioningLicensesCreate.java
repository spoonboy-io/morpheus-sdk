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

/**
 * ProvisioningLicensesCreate
 */
@javax.annotation.Generated(value = "org.openapitools.codegen.languages.JavaClientCodegen", date = "2023-08-21T11:26:47.028Z[GMT]")
public class ProvisioningLicensesCreate {
  public static final String SERIALIZED_NAME_NAME = "name";
  @SerializedName(SERIALIZED_NAME_NAME)
  private String name;

  public static final String SERIALIZED_NAME_LICENSE_TYPE = "licenseType";
  @SerializedName(SERIALIZED_NAME_LICENSE_TYPE)
  private String licenseType;

  public static final String SERIALIZED_NAME_LICENSE_KEY = "licenseKey";
  @SerializedName(SERIALIZED_NAME_LICENSE_KEY)
  private String licenseKey;

  public static final String SERIALIZED_NAME_ORG_NAME = "orgName";
  @SerializedName(SERIALIZED_NAME_ORG_NAME)
  private String orgName;

  public static final String SERIALIZED_NAME_FULL_NAME = "fullName";
  @SerializedName(SERIALIZED_NAME_FULL_NAME)
  private String fullName;

  public static final String SERIALIZED_NAME_LICENSE_VERSION = "licenseVersion";
  @SerializedName(SERIALIZED_NAME_LICENSE_VERSION)
  private String licenseVersion;

  public static final String SERIALIZED_NAME_COPIES = "copies";
  @SerializedName(SERIALIZED_NAME_COPIES)
  private Long copies = 1l;

  public static final String SERIALIZED_NAME_DESCRIPTION = "description";
  @SerializedName(SERIALIZED_NAME_DESCRIPTION)
  private String description;

  public static final String SERIALIZED_NAME_VIRTUAL_IMAGES = "virtualImages";
  @SerializedName(SERIALIZED_NAME_VIRTUAL_IMAGES)
  private List<Long> virtualImages = null;

  public static final String SERIALIZED_NAME_TENANTS = "tenants";
  @SerializedName(SERIALIZED_NAME_TENANTS)
  private List<Long> tenants = null;


  public ProvisioningLicensesCreate name(String name) {
    
    this.name = name;
    return this;
  }

   /**
   * Name
   * @return name
  **/
  @ApiModelProperty(required = true, value = "Name")

  public String getName() {
    return name;
  }


  public void setName(String name) {
    this.name = name;
  }


  public ProvisioningLicensesCreate licenseType(String licenseType) {
    
    this.licenseType = licenseType;
    return this;
  }

   /**
   * License Type - The license type code.
   * @return licenseType
  **/
  @ApiModelProperty(example = "win", required = true, value = "License Type - The license type code.")

  public String getLicenseType() {
    return licenseType;
  }


  public void setLicenseType(String licenseType) {
    this.licenseType = licenseType;
  }


  public ProvisioningLicensesCreate licenseKey(String licenseKey) {
    
    this.licenseKey = licenseKey;
    return this;
  }

   /**
   * License Key - The license key, to be kept a secret.
   * @return licenseKey
  **/
  @ApiModelProperty(required = true, value = "License Key - The license key, to be kept a secret.")

  public String getLicenseKey() {
    return licenseKey;
  }


  public void setLicenseKey(String licenseKey) {
    this.licenseKey = licenseKey;
  }


  public ProvisioningLicensesCreate orgName(String orgName) {
    
    this.orgName = orgName;
    return this;
  }

   /**
   * Org Name - The Organization Name (if applicable) related to the license key
   * @return orgName
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "Org Name - The Organization Name (if applicable) related to the license key")

  public String getOrgName() {
    return orgName;
  }


  public void setOrgName(String orgName) {
    this.orgName = orgName;
  }


  public ProvisioningLicensesCreate fullName(String fullName) {
    
    this.fullName = fullName;
    return this;
  }

   /**
   * Full Name - The Full Name (if applicable) related to the license key
   * @return fullName
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "Full Name - The Full Name (if applicable) related to the license key")

  public String getFullName() {
    return fullName;
  }


  public void setFullName(String fullName) {
    this.fullName = fullName;
  }


  public ProvisioningLicensesCreate licenseVersion(String licenseVersion) {
    
    this.licenseVersion = licenseVersion;
    return this;
  }

   /**
   * License Version
   * @return licenseVersion
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "License Version")

  public String getLicenseVersion() {
    return licenseVersion;
  }


  public void setLicenseVersion(String licenseVersion) {
    this.licenseVersion = licenseVersion;
  }


  public ProvisioningLicensesCreate copies(Long copies) {
    
    this.copies = copies;
    return this;
  }

   /**
   * Copies - The number of times the key can be used.
   * @return copies
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "Copies - The number of times the key can be used.")

  public Long getCopies() {
    return copies;
  }


  public void setCopies(Long copies) {
    this.copies = copies;
  }


  public ProvisioningLicensesCreate description(String description) {
    
    this.description = description;
    return this;
  }

   /**
   * Description
   * @return description
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "Description")

  public String getDescription() {
    return description;
  }


  public void setDescription(String description) {
    this.description = description;
  }


  public ProvisioningLicensesCreate virtualImages(List<Long> virtualImages) {
    
    this.virtualImages = virtualImages;
    return this;
  }

  public ProvisioningLicensesCreate addVirtualImagesItem(Long virtualImagesItem) {
    if (this.virtualImages == null) {
      this.virtualImages = new ArrayList<Long>();
    }
    this.virtualImages.add(virtualImagesItem);
    return this;
  }

   /**
   * Virtual Images - Array of Virtual Image IDs to associate with license.
   * @return virtualImages
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "Virtual Images - Array of Virtual Image IDs to associate with license.")

  public List<Long> getVirtualImages() {
    return virtualImages;
  }


  public void setVirtualImages(List<Long> virtualImages) {
    this.virtualImages = virtualImages;
  }


  public ProvisioningLicensesCreate tenants(List<Long> tenants) {
    
    this.tenants = tenants;
    return this;
  }

  public ProvisioningLicensesCreate addTenantsItem(Long tenantsItem) {
    if (this.tenants == null) {
      this.tenants = new ArrayList<Long>();
    }
    this.tenants.add(tenantsItem);
    return this;
  }

   /**
   * Tenants - Array of tenants that are allowed to use the key.
   * @return tenants
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "Tenants - Array of tenants that are allowed to use the key.")

  public List<Long> getTenants() {
    return tenants;
  }


  public void setTenants(List<Long> tenants) {
    this.tenants = tenants;
  }


  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (o == null || getClass() != o.getClass()) {
      return false;
    }
    ProvisioningLicensesCreate provisioningLicensesCreate = (ProvisioningLicensesCreate) o;
    return Objects.equals(this.name, provisioningLicensesCreate.name) &&
        Objects.equals(this.licenseType, provisioningLicensesCreate.licenseType) &&
        Objects.equals(this.licenseKey, provisioningLicensesCreate.licenseKey) &&
        Objects.equals(this.orgName, provisioningLicensesCreate.orgName) &&
        Objects.equals(this.fullName, provisioningLicensesCreate.fullName) &&
        Objects.equals(this.licenseVersion, provisioningLicensesCreate.licenseVersion) &&
        Objects.equals(this.copies, provisioningLicensesCreate.copies) &&
        Objects.equals(this.description, provisioningLicensesCreate.description) &&
        Objects.equals(this.virtualImages, provisioningLicensesCreate.virtualImages) &&
        Objects.equals(this.tenants, provisioningLicensesCreate.tenants);
  }

  @Override
  public int hashCode() {
    return Objects.hash(name, licenseType, licenseKey, orgName, fullName, licenseVersion, copies, description, virtualImages, tenants);
  }


  @Override
  public String toString() {
    StringBuilder sb = new StringBuilder();
    sb.append("class ProvisioningLicensesCreate {\n");
    sb.append("    name: ").append(toIndentedString(name)).append("\n");
    sb.append("    licenseType: ").append(toIndentedString(licenseType)).append("\n");
    sb.append("    licenseKey: ").append(toIndentedString(licenseKey)).append("\n");
    sb.append("    orgName: ").append(toIndentedString(orgName)).append("\n");
    sb.append("    fullName: ").append(toIndentedString(fullName)).append("\n");
    sb.append("    licenseVersion: ").append(toIndentedString(licenseVersion)).append("\n");
    sb.append("    copies: ").append(toIndentedString(copies)).append("\n");
    sb.append("    description: ").append(toIndentedString(description)).append("\n");
    sb.append("    virtualImages: ").append(toIndentedString(virtualImages)).append("\n");
    sb.append("    tenants: ").append(toIndentedString(tenants)).append("\n");
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

