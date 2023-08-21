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
import org.openapitools.client.model.InlineResponse20040AppDeployInstance;
import org.openapitools.client.model.InlineResponse20079LoadBalancerMonitorLoadBalancerType;

/**
 * ProvisioningLicense
 */
@javax.annotation.Generated(value = "org.openapitools.codegen.languages.JavaClientCodegen", date = "2023-08-21T11:26:47.028Z[GMT]")
public class ProvisioningLicense {
  public static final String SERIALIZED_NAME_ID = "id";
  @SerializedName(SERIALIZED_NAME_ID)
  private Long id;

  public static final String SERIALIZED_NAME_NAME = "name";
  @SerializedName(SERIALIZED_NAME_NAME)
  private String name;

  public static final String SERIALIZED_NAME_DESCRIPTION = "description";
  @SerializedName(SERIALIZED_NAME_DESCRIPTION)
  private String description;

  public static final String SERIALIZED_NAME_LICENSE_TYPE = "licenseType";
  @SerializedName(SERIALIZED_NAME_LICENSE_TYPE)
  private InlineResponse20079LoadBalancerMonitorLoadBalancerType licenseType;

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
  private Long copies;

  public static final String SERIALIZED_NAME_RESERVATION_COUNT = "reservationCount";
  @SerializedName(SERIALIZED_NAME_RESERVATION_COUNT)
  private Long reservationCount;

  public static final String SERIALIZED_NAME_TENANTS = "tenants";
  @SerializedName(SERIALIZED_NAME_TENANTS)
  private List<Object> tenants = null;

  public static final String SERIALIZED_NAME_VIRTUAL_IMAGES = "virtualImages";
  @SerializedName(SERIALIZED_NAME_VIRTUAL_IMAGES)
  private List<InlineResponse20040AppDeployInstance> virtualImages = null;

  public static final String SERIALIZED_NAME_ACCOUNT = "account";
  @SerializedName(SERIALIZED_NAME_ACCOUNT)
  private InlineResponse20040AppDeployInstance account;


  public ProvisioningLicense id(Long id) {
    
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


  public ProvisioningLicense name(String name) {
    
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


  public ProvisioningLicense description(String description) {
    
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


  public ProvisioningLicense licenseType(InlineResponse20079LoadBalancerMonitorLoadBalancerType licenseType) {
    
    this.licenseType = licenseType;
    return this;
  }

   /**
   * Get licenseType
   * @return licenseType
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public InlineResponse20079LoadBalancerMonitorLoadBalancerType getLicenseType() {
    return licenseType;
  }


  public void setLicenseType(InlineResponse20079LoadBalancerMonitorLoadBalancerType licenseType) {
    this.licenseType = licenseType;
  }


  public ProvisioningLicense licenseKey(String licenseKey) {
    
    this.licenseKey = licenseKey;
    return this;
  }

   /**
   * Get licenseKey
   * @return licenseKey
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public String getLicenseKey() {
    return licenseKey;
  }


  public void setLicenseKey(String licenseKey) {
    this.licenseKey = licenseKey;
  }


  public ProvisioningLicense orgName(String orgName) {
    
    this.orgName = orgName;
    return this;
  }

   /**
   * Get orgName
   * @return orgName
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public String getOrgName() {
    return orgName;
  }


  public void setOrgName(String orgName) {
    this.orgName = orgName;
  }


  public ProvisioningLicense fullName(String fullName) {
    
    this.fullName = fullName;
    return this;
  }

   /**
   * Get fullName
   * @return fullName
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public String getFullName() {
    return fullName;
  }


  public void setFullName(String fullName) {
    this.fullName = fullName;
  }


  public ProvisioningLicense licenseVersion(String licenseVersion) {
    
    this.licenseVersion = licenseVersion;
    return this;
  }

   /**
   * Get licenseVersion
   * @return licenseVersion
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public String getLicenseVersion() {
    return licenseVersion;
  }


  public void setLicenseVersion(String licenseVersion) {
    this.licenseVersion = licenseVersion;
  }


  public ProvisioningLicense copies(Long copies) {
    
    this.copies = copies;
    return this;
  }

   /**
   * Get copies
   * @return copies
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public Long getCopies() {
    return copies;
  }


  public void setCopies(Long copies) {
    this.copies = copies;
  }


  public ProvisioningLicense reservationCount(Long reservationCount) {
    
    this.reservationCount = reservationCount;
    return this;
  }

   /**
   * Get reservationCount
   * @return reservationCount
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public Long getReservationCount() {
    return reservationCount;
  }


  public void setReservationCount(Long reservationCount) {
    this.reservationCount = reservationCount;
  }


  public ProvisioningLicense tenants(List<Object> tenants) {
    
    this.tenants = tenants;
    return this;
  }

  public ProvisioningLicense addTenantsItem(Object tenantsItem) {
    if (this.tenants == null) {
      this.tenants = new ArrayList<Object>();
    }
    this.tenants.add(tenantsItem);
    return this;
  }

   /**
   * Get tenants
   * @return tenants
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public List<Object> getTenants() {
    return tenants;
  }


  public void setTenants(List<Object> tenants) {
    this.tenants = tenants;
  }


  public ProvisioningLicense virtualImages(List<InlineResponse20040AppDeployInstance> virtualImages) {
    
    this.virtualImages = virtualImages;
    return this;
  }

  public ProvisioningLicense addVirtualImagesItem(InlineResponse20040AppDeployInstance virtualImagesItem) {
    if (this.virtualImages == null) {
      this.virtualImages = new ArrayList<InlineResponse20040AppDeployInstance>();
    }
    this.virtualImages.add(virtualImagesItem);
    return this;
  }

   /**
   * Get virtualImages
   * @return virtualImages
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public List<InlineResponse20040AppDeployInstance> getVirtualImages() {
    return virtualImages;
  }


  public void setVirtualImages(List<InlineResponse20040AppDeployInstance> virtualImages) {
    this.virtualImages = virtualImages;
  }


  public ProvisioningLicense account(InlineResponse20040AppDeployInstance account) {
    
    this.account = account;
    return this;
  }

   /**
   * Get account
   * @return account
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public InlineResponse20040AppDeployInstance getAccount() {
    return account;
  }


  public void setAccount(InlineResponse20040AppDeployInstance account) {
    this.account = account;
  }


  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (o == null || getClass() != o.getClass()) {
      return false;
    }
    ProvisioningLicense provisioningLicense = (ProvisioningLicense) o;
    return Objects.equals(this.id, provisioningLicense.id) &&
        Objects.equals(this.name, provisioningLicense.name) &&
        Objects.equals(this.description, provisioningLicense.description) &&
        Objects.equals(this.licenseType, provisioningLicense.licenseType) &&
        Objects.equals(this.licenseKey, provisioningLicense.licenseKey) &&
        Objects.equals(this.orgName, provisioningLicense.orgName) &&
        Objects.equals(this.fullName, provisioningLicense.fullName) &&
        Objects.equals(this.licenseVersion, provisioningLicense.licenseVersion) &&
        Objects.equals(this.copies, provisioningLicense.copies) &&
        Objects.equals(this.reservationCount, provisioningLicense.reservationCount) &&
        Objects.equals(this.tenants, provisioningLicense.tenants) &&
        Objects.equals(this.virtualImages, provisioningLicense.virtualImages) &&
        Objects.equals(this.account, provisioningLicense.account);
  }

  @Override
  public int hashCode() {
    return Objects.hash(id, name, description, licenseType, licenseKey, orgName, fullName, licenseVersion, copies, reservationCount, tenants, virtualImages, account);
  }


  @Override
  public String toString() {
    StringBuilder sb = new StringBuilder();
    sb.append("class ProvisioningLicense {\n");
    sb.append("    id: ").append(toIndentedString(id)).append("\n");
    sb.append("    name: ").append(toIndentedString(name)).append("\n");
    sb.append("    description: ").append(toIndentedString(description)).append("\n");
    sb.append("    licenseType: ").append(toIndentedString(licenseType)).append("\n");
    sb.append("    licenseKey: ").append(toIndentedString(licenseKey)).append("\n");
    sb.append("    orgName: ").append(toIndentedString(orgName)).append("\n");
    sb.append("    fullName: ").append(toIndentedString(fullName)).append("\n");
    sb.append("    licenseVersion: ").append(toIndentedString(licenseVersion)).append("\n");
    sb.append("    copies: ").append(toIndentedString(copies)).append("\n");
    sb.append("    reservationCount: ").append(toIndentedString(reservationCount)).append("\n");
    sb.append("    tenants: ").append(toIndentedString(tenants)).append("\n");
    sb.append("    virtualImages: ").append(toIndentedString(virtualImages)).append("\n");
    sb.append("    account: ").append(toIndentedString(account)).append("\n");
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
