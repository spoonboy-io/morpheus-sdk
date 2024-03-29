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
 * IntegrationDockerRepoConfigIntegration
 */
@javax.annotation.Generated(value = "org.openapitools.codegen.languages.JavaClientCodegen", date = "2023-08-21T11:26:47.028Z[GMT]")
public class IntegrationDockerRepoConfigIntegration {
  public static final String SERIALIZED_NAME_NAME = "name";
  @SerializedName(SERIALIZED_NAME_NAME)
  private String name;

  /**
   * Integration Type Code
   */
  @JsonAdapter(TypeEnum.Adapter.class)
  public enum TypeEnum {
    DOCKER_REGISTRY("docker.registry");

    private String value;

    TypeEnum(String value) {
      this.value = value;
    }

    public String getValue() {
      return value;
    }

    @Override
    public String toString() {
      return String.valueOf(value);
    }

    public static TypeEnum fromValue(String value) {
      for (TypeEnum b : TypeEnum.values()) {
        if (b.value.equals(value)) {
          return b;
        }
      }
      throw new IllegalArgumentException("Unexpected value '" + value + "'");
    }

    public static class Adapter extends TypeAdapter<TypeEnum> {
      @Override
      public void write(final JsonWriter jsonWriter, final TypeEnum enumeration) throws IOException {
        jsonWriter.value(enumeration.getValue());
      }

      @Override
      public TypeEnum read(final JsonReader jsonReader) throws IOException {
        String value =  jsonReader.nextString();
        return TypeEnum.fromValue(value);
      }
    }
  }

  public static final String SERIALIZED_NAME_TYPE = "type";
  @SerializedName(SERIALIZED_NAME_TYPE)
  private TypeEnum type;

  public static final String SERIALIZED_NAME_SERVICE_URL = "serviceUrl";
  @SerializedName(SERIALIZED_NAME_SERVICE_URL)
  private String serviceUrl;

  public static final String SERIALIZED_NAME_SERVICE_USERNAME = "serviceUsername";
  @SerializedName(SERIALIZED_NAME_SERVICE_USERNAME)
  private String serviceUsername;

  public static final String SERIALIZED_NAME_SERVICE_PASSWORD = "servicePassword";
  @SerializedName(SERIALIZED_NAME_SERVICE_PASSWORD)
  private String servicePassword;


  public IntegrationDockerRepoConfigIntegration name(String name) {
    
    this.name = name;
    return this;
  }

   /**
   * Name, a unique identifier for the integration
   * @return name
  **/
  @ApiModelProperty(example = "Sample Integration", required = true, value = "Name, a unique identifier for the integration")

  public String getName() {
    return name;
  }


  public void setName(String name) {
    this.name = name;
  }


  public IntegrationDockerRepoConfigIntegration type(TypeEnum type) {
    
    this.type = type;
    return this;
  }

   /**
   * Integration Type Code
   * @return type
  **/
  @ApiModelProperty(required = true, value = "Integration Type Code")

  public TypeEnum getType() {
    return type;
  }


  public void setType(TypeEnum type) {
    this.type = type;
  }


  public IntegrationDockerRepoConfigIntegration serviceUrl(String serviceUrl) {
    
    this.serviceUrl = serviceUrl;
    return this;
  }

   /**
   * Repository URL
   * @return serviceUrl
  **/
  @ApiModelProperty(required = true, value = "Repository URL")

  public String getServiceUrl() {
    return serviceUrl;
  }


  public void setServiceUrl(String serviceUrl) {
    this.serviceUrl = serviceUrl;
  }


  public IntegrationDockerRepoConfigIntegration serviceUsername(String serviceUsername) {
    
    this.serviceUsername = serviceUsername;
    return this;
  }

   /**
   * Username
   * @return serviceUsername
  **/
  @ApiModelProperty(required = true, value = "Username")

  public String getServiceUsername() {
    return serviceUsername;
  }


  public void setServiceUsername(String serviceUsername) {
    this.serviceUsername = serviceUsername;
  }


  public IntegrationDockerRepoConfigIntegration servicePassword(String servicePassword) {
    
    this.servicePassword = servicePassword;
    return this;
  }

   /**
   * Password
   * @return servicePassword
  **/
  @ApiModelProperty(required = true, value = "Password")

  public String getServicePassword() {
    return servicePassword;
  }


  public void setServicePassword(String servicePassword) {
    this.servicePassword = servicePassword;
  }


  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (o == null || getClass() != o.getClass()) {
      return false;
    }
    IntegrationDockerRepoConfigIntegration integrationDockerRepoConfigIntegration = (IntegrationDockerRepoConfigIntegration) o;
    return Objects.equals(this.name, integrationDockerRepoConfigIntegration.name) &&
        Objects.equals(this.type, integrationDockerRepoConfigIntegration.type) &&
        Objects.equals(this.serviceUrl, integrationDockerRepoConfigIntegration.serviceUrl) &&
        Objects.equals(this.serviceUsername, integrationDockerRepoConfigIntegration.serviceUsername) &&
        Objects.equals(this.servicePassword, integrationDockerRepoConfigIntegration.servicePassword);
  }

  @Override
  public int hashCode() {
    return Objects.hash(name, type, serviceUrl, serviceUsername, servicePassword);
  }


  @Override
  public String toString() {
    StringBuilder sb = new StringBuilder();
    sb.append("class IntegrationDockerRepoConfigIntegration {\n");
    sb.append("    name: ").append(toIndentedString(name)).append("\n");
    sb.append("    type: ").append(toIndentedString(type)).append("\n");
    sb.append("    serviceUrl: ").append(toIndentedString(serviceUrl)).append("\n");
    sb.append("    serviceUsername: ").append(toIndentedString(serviceUsername)).append("\n");
    sb.append("    servicePassword: ").append(toIndentedString(servicePassword)).append("\n");
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

