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
import org.openapitools.client.model.InlineResponse20094Network;

/**
 * IntegrationObjectLayout
 */
@javax.annotation.Generated(value = "org.openapitools.codegen.languages.JavaClientCodegen", date = "2023-08-21T11:26:47.028Z[GMT]")
public class IntegrationObjectLayout {
  public static final String SERIALIZED_NAME_ID = "id";
  @SerializedName(SERIALIZED_NAME_ID)
  private Long id;

  public static final String SERIALIZED_NAME_NAME = "name";
  @SerializedName(SERIALIZED_NAME_NAME)
  private String name;

  public static final String SERIALIZED_NAME_CODE = "code";
  @SerializedName(SERIALIZED_NAME_CODE)
  private String code;

  public static final String SERIALIZED_NAME_PROVISION_TYPE = "provisionType";
  @SerializedName(SERIALIZED_NAME_PROVISION_TYPE)
  private InlineResponse20094Network provisionType;

  public static final String SERIALIZED_NAME_INSTANCE_TYPE = "instanceType";
  @SerializedName(SERIALIZED_NAME_INSTANCE_TYPE)
  private InlineResponse20094Network instanceType;

  public static final String SERIALIZED_NAME_INSTANCE_VERSION = "instanceVersion";
  @SerializedName(SERIALIZED_NAME_INSTANCE_VERSION)
  private String instanceVersion;


  public IntegrationObjectLayout id(Long id) {
    
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


  public IntegrationObjectLayout name(String name) {
    
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


  public IntegrationObjectLayout code(String code) {
    
    this.code = code;
    return this;
  }

   /**
   * Get code
   * @return code
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public String getCode() {
    return code;
  }


  public void setCode(String code) {
    this.code = code;
  }


  public IntegrationObjectLayout provisionType(InlineResponse20094Network provisionType) {
    
    this.provisionType = provisionType;
    return this;
  }

   /**
   * Get provisionType
   * @return provisionType
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public InlineResponse20094Network getProvisionType() {
    return provisionType;
  }


  public void setProvisionType(InlineResponse20094Network provisionType) {
    this.provisionType = provisionType;
  }


  public IntegrationObjectLayout instanceType(InlineResponse20094Network instanceType) {
    
    this.instanceType = instanceType;
    return this;
  }

   /**
   * Get instanceType
   * @return instanceType
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public InlineResponse20094Network getInstanceType() {
    return instanceType;
  }


  public void setInstanceType(InlineResponse20094Network instanceType) {
    this.instanceType = instanceType;
  }


  public IntegrationObjectLayout instanceVersion(String instanceVersion) {
    
    this.instanceVersion = instanceVersion;
    return this;
  }

   /**
   * Get instanceVersion
   * @return instanceVersion
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public String getInstanceVersion() {
    return instanceVersion;
  }


  public void setInstanceVersion(String instanceVersion) {
    this.instanceVersion = instanceVersion;
  }


  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (o == null || getClass() != o.getClass()) {
      return false;
    }
    IntegrationObjectLayout integrationObjectLayout = (IntegrationObjectLayout) o;
    return Objects.equals(this.id, integrationObjectLayout.id) &&
        Objects.equals(this.name, integrationObjectLayout.name) &&
        Objects.equals(this.code, integrationObjectLayout.code) &&
        Objects.equals(this.provisionType, integrationObjectLayout.provisionType) &&
        Objects.equals(this.instanceType, integrationObjectLayout.instanceType) &&
        Objects.equals(this.instanceVersion, integrationObjectLayout.instanceVersion);
  }

  @Override
  public int hashCode() {
    return Objects.hash(id, name, code, provisionType, instanceType, instanceVersion);
  }


  @Override
  public String toString() {
    StringBuilder sb = new StringBuilder();
    sb.append("class IntegrationObjectLayout {\n");
    sb.append("    id: ").append(toIndentedString(id)).append("\n");
    sb.append("    name: ").append(toIndentedString(name)).append("\n");
    sb.append("    code: ").append(toIndentedString(code)).append("\n");
    sb.append("    provisionType: ").append(toIndentedString(provisionType)).append("\n");
    sb.append("    instanceType: ").append(toIndentedString(instanceType)).append("\n");
    sb.append("    instanceVersion: ").append(toIndentedString(instanceVersion)).append("\n");
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

