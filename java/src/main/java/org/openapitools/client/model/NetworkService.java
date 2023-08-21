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
 * NetworkService
 */
@javax.annotation.Generated(value = "org.openapitools.codegen.languages.JavaClientCodegen", date = "2023-08-21T11:26:47.028Z[GMT]")
public class NetworkService {
  public static final String SERIALIZED_NAME_SERVICE_TYPE = "serviceType";
  @SerializedName(SERIALIZED_NAME_SERVICE_TYPE)
  private String serviceType;

  public static final String SERIALIZED_NAME_SERVICE_TYPE_NAME = "serviceTypeName";
  @SerializedName(SERIALIZED_NAME_SERVICE_TYPE_NAME)
  private String serviceTypeName;

  public static final String SERIALIZED_NAME_TYPE = "type";
  @SerializedName(SERIALIZED_NAME_TYPE)
  private String type;

  public static final String SERIALIZED_NAME_TYPE_NAME = "typeName";
  @SerializedName(SERIALIZED_NAME_TYPE_NAME)
  private String typeName;

  public static final String SERIALIZED_NAME_NAME = "name";
  @SerializedName(SERIALIZED_NAME_NAME)
  private String name;

  public static final String SERIALIZED_NAME_ID = "id";
  @SerializedName(SERIALIZED_NAME_ID)
  private Long id;

  public static final String SERIALIZED_NAME_INTEGRATION_ID = "integrationId";
  @SerializedName(SERIALIZED_NAME_INTEGRATION_ID)
  private Long integrationId;


  public NetworkService serviceType(String serviceType) {
    
    this.serviceType = serviceType;
    return this;
  }

   /**
   * Get serviceType
   * @return serviceType
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public String getServiceType() {
    return serviceType;
  }


  public void setServiceType(String serviceType) {
    this.serviceType = serviceType;
  }


  public NetworkService serviceTypeName(String serviceTypeName) {
    
    this.serviceTypeName = serviceTypeName;
    return this;
  }

   /**
   * Get serviceTypeName
   * @return serviceTypeName
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public String getServiceTypeName() {
    return serviceTypeName;
  }


  public void setServiceTypeName(String serviceTypeName) {
    this.serviceTypeName = serviceTypeName;
  }


  public NetworkService type(String type) {
    
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


  public NetworkService typeName(String typeName) {
    
    this.typeName = typeName;
    return this;
  }

   /**
   * Get typeName
   * @return typeName
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public String getTypeName() {
    return typeName;
  }


  public void setTypeName(String typeName) {
    this.typeName = typeName;
  }


  public NetworkService name(String name) {
    
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


  public NetworkService id(Long id) {
    
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


  public NetworkService integrationId(Long integrationId) {
    
    this.integrationId = integrationId;
    return this;
  }

   /**
   * Get integrationId
   * @return integrationId
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public Long getIntegrationId() {
    return integrationId;
  }


  public void setIntegrationId(Long integrationId) {
    this.integrationId = integrationId;
  }


  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (o == null || getClass() != o.getClass()) {
      return false;
    }
    NetworkService networkService = (NetworkService) o;
    return Objects.equals(this.serviceType, networkService.serviceType) &&
        Objects.equals(this.serviceTypeName, networkService.serviceTypeName) &&
        Objects.equals(this.type, networkService.type) &&
        Objects.equals(this.typeName, networkService.typeName) &&
        Objects.equals(this.name, networkService.name) &&
        Objects.equals(this.id, networkService.id) &&
        Objects.equals(this.integrationId, networkService.integrationId);
  }

  @Override
  public int hashCode() {
    return Objects.hash(serviceType, serviceTypeName, type, typeName, name, id, integrationId);
  }


  @Override
  public String toString() {
    StringBuilder sb = new StringBuilder();
    sb.append("class NetworkService {\n");
    sb.append("    serviceType: ").append(toIndentedString(serviceType)).append("\n");
    sb.append("    serviceTypeName: ").append(toIndentedString(serviceTypeName)).append("\n");
    sb.append("    type: ").append(toIndentedString(type)).append("\n");
    sb.append("    typeName: ").append(toIndentedString(typeName)).append("\n");
    sb.append("    name: ").append(toIndentedString(name)).append("\n");
    sb.append("    id: ").append(toIndentedString(id)).append("\n");
    sb.append("    integrationId: ").append(toIndentedString(integrationId)).append("\n");
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
