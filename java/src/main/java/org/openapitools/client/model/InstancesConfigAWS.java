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
 * InstancesConfigAWS
 */
@javax.annotation.Generated(value = "org.openapitools.codegen.languages.JavaClientCodegen", date = "2023-08-21T11:26:47.028Z[GMT]")
public class InstancesConfigAWS {
  public static final String SERIALIZED_NAME_NO_AGENT = "noAgent";
  @SerializedName(SERIALIZED_NAME_NO_AGENT)
  private Boolean noAgent = false;

  public static final String SERIALIZED_NAME_IS_E_C2 = "isEC2";
  @SerializedName(SERIALIZED_NAME_IS_E_C2)
  private String isEC2 = "false";

  public static final String SERIALIZED_NAME_AVAILABILITY_ID = "availabilityId";
  @SerializedName(SERIALIZED_NAME_AVAILABILITY_ID)
  private String availabilityId;

  public static final String SERIALIZED_NAME_SECURITY_ID = "securityId";
  @SerializedName(SERIALIZED_NAME_SECURITY_ID)
  private String securityId;

  public static final String SERIALIZED_NAME_PUBLIC_IP_TYPE = "publicIpType";
  @SerializedName(SERIALIZED_NAME_PUBLIC_IP_TYPE)
  private String publicIpType;

  public static final String SERIALIZED_NAME_INSTANCE_PROFILE = "instanceProfile";
  @SerializedName(SERIALIZED_NAME_INSTANCE_PROFILE)
  private String instanceProfile;

  public static final String SERIALIZED_NAME_KMS_KEY_ID = "kmsKeyId";
  @SerializedName(SERIALIZED_NAME_KMS_KEY_ID)
  private String kmsKeyId;


  public InstancesConfigAWS noAgent(Boolean noAgent) {
    
    this.noAgent = noAgent;
    return this;
  }

   /**
   * Skipping Agent installation will result in a lack of logging and guest operating system statistics. Automation scripts may also be adversely affected.
   * @return noAgent
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "Skipping Agent installation will result in a lack of logging and guest operating system statistics. Automation scripts may also be adversely affected.")

  public Boolean getNoAgent() {
    return noAgent;
  }


  public void setNoAgent(Boolean noAgent) {
    this.noAgent = noAgent;
  }


  public InstancesConfigAWS isEC2(String isEC2) {
    
    this.isEC2 = isEC2;
    return this;
  }

   /**
   * Amazon Cloud Type
   * @return isEC2
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "Amazon Cloud Type")

  public String getIsEC2() {
    return isEC2;
  }


  public void setIsEC2(String isEC2) {
    this.isEC2 = isEC2;
  }


  public InstancesConfigAWS availabilityId(String availabilityId) {
    
    this.availabilityId = availabilityId;
    return this;
  }

   /**
   * Amazon Zone
   * @return availabilityId
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "Amazon Zone")

  public String getAvailabilityId() {
    return availabilityId;
  }


  public void setAvailabilityId(String availabilityId) {
    this.availabilityId = availabilityId;
  }


  public InstancesConfigAWS securityId(String securityId) {
    
    this.securityId = securityId;
    return this;
  }

   /**
   * Security Group
   * @return securityId
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "Security Group")

  public String getSecurityId() {
    return securityId;
  }


  public void setSecurityId(String securityId) {
    this.securityId = securityId;
  }


  public InstancesConfigAWS publicIpType(String publicIpType) {
    
    this.publicIpType = publicIpType;
    return this;
  }

   /**
   * Public IP
   * @return publicIpType
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(example = "elasticIp", value = "Public IP")

  public String getPublicIpType() {
    return publicIpType;
  }


  public void setPublicIpType(String publicIpType) {
    this.publicIpType = publicIpType;
  }


  public InstancesConfigAWS instanceProfile(String instanceProfile) {
    
    this.instanceProfile = instanceProfile;
    return this;
  }

   /**
   * IAM Profile
   * @return instanceProfile
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "IAM Profile")

  public String getInstanceProfile() {
    return instanceProfile;
  }


  public void setInstanceProfile(String instanceProfile) {
    this.instanceProfile = instanceProfile;
  }


  public InstancesConfigAWS kmsKeyId(String kmsKeyId) {
    
    this.kmsKeyId = kmsKeyId;
    return this;
  }

   /**
   * KMS Key ID
   * @return kmsKeyId
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "KMS Key ID")

  public String getKmsKeyId() {
    return kmsKeyId;
  }


  public void setKmsKeyId(String kmsKeyId) {
    this.kmsKeyId = kmsKeyId;
  }


  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (o == null || getClass() != o.getClass()) {
      return false;
    }
    InstancesConfigAWS instancesConfigAWS = (InstancesConfigAWS) o;
    return Objects.equals(this.noAgent, instancesConfigAWS.noAgent) &&
        Objects.equals(this.isEC2, instancesConfigAWS.isEC2) &&
        Objects.equals(this.availabilityId, instancesConfigAWS.availabilityId) &&
        Objects.equals(this.securityId, instancesConfigAWS.securityId) &&
        Objects.equals(this.publicIpType, instancesConfigAWS.publicIpType) &&
        Objects.equals(this.instanceProfile, instancesConfigAWS.instanceProfile) &&
        Objects.equals(this.kmsKeyId, instancesConfigAWS.kmsKeyId);
  }

  @Override
  public int hashCode() {
    return Objects.hash(noAgent, isEC2, availabilityId, securityId, publicIpType, instanceProfile, kmsKeyId);
  }


  @Override
  public String toString() {
    StringBuilder sb = new StringBuilder();
    sb.append("class InstancesConfigAWS {\n");
    sb.append("    noAgent: ").append(toIndentedString(noAgent)).append("\n");
    sb.append("    isEC2: ").append(toIndentedString(isEC2)).append("\n");
    sb.append("    availabilityId: ").append(toIndentedString(availabilityId)).append("\n");
    sb.append("    securityId: ").append(toIndentedString(securityId)).append("\n");
    sb.append("    publicIpType: ").append(toIndentedString(publicIpType)).append("\n");
    sb.append("    instanceProfile: ").append(toIndentedString(instanceProfile)).append("\n");
    sb.append("    kmsKeyId: ").append(toIndentedString(kmsKeyId)).append("\n");
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

