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
 * InlineResponse200157
 */
@javax.annotation.Generated(value = "org.openapitools.codegen.languages.JavaClientCodegen", date = "2023-08-21T11:26:47.028Z[GMT]")
public class InlineResponse200157 {
  public static final String SERIALIZED_NAME_CLIENTS = "clients";
  @SerializedName(SERIALIZED_NAME_CLIENTS)
  private List<Object> clients = null;


  public InlineResponse200157 clients(List<Object> clients) {
    
    this.clients = clients;
    return this;
  }

  public InlineResponse200157 addClientsItem(Object clientsItem) {
    if (this.clients == null) {
      this.clients = new ArrayList<Object>();
    }
    this.clients.add(clientsItem);
    return this;
  }

   /**
   * Get clients
   * @return clients
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public List<Object> getClients() {
    return clients;
  }


  public void setClients(List<Object> clients) {
    this.clients = clients;
  }


  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (o == null || getClass() != o.getClass()) {
      return false;
    }
    InlineResponse200157 inlineResponse200157 = (InlineResponse200157) o;
    return Objects.equals(this.clients, inlineResponse200157.clients);
  }

  @Override
  public int hashCode() {
    return Objects.hash(clients);
  }


  @Override
  public String toString() {
    StringBuilder sb = new StringBuilder();
    sb.append("class InlineResponse200157 {\n");
    sb.append("    clients: ").append(toIndentedString(clients)).append("\n");
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

