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
 * LoadBalancerCreateResourcePermission
 */
@javax.annotation.Generated(value = "org.openapitools.codegen.languages.JavaClientCodegen", date = "2023-08-21T11:26:47.028Z[GMT]")
public class LoadBalancerCreateResourcePermission {
  public static final String SERIALIZED_NAME_ALL = "all";
  @SerializedName(SERIALIZED_NAME_ALL)
  private Boolean all;

  public static final String SERIALIZED_NAME_SITES = "sites";
  @SerializedName(SERIALIZED_NAME_SITES)
  private List<Long> sites = null;


  public LoadBalancerCreateResourcePermission all(Boolean all) {
    
    this.all = all;
    return this;
  }

   /**
   * Pass true to allow access to all groups
   * @return all
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "Pass true to allow access to all groups")

  public Boolean getAll() {
    return all;
  }


  public void setAll(Boolean all) {
    this.all = all;
  }


  public LoadBalancerCreateResourcePermission sites(List<Long> sites) {
    
    this.sites = sites;
    return this;
  }

  public LoadBalancerCreateResourcePermission addSitesItem(Long sitesItem) {
    if (this.sites == null) {
      this.sites = new ArrayList<Long>();
    }
    this.sites.add(sitesItem);
    return this;
  }

   /**
   * Array of groups that are allowed access
   * @return sites
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "Array of groups that are allowed access")

  public List<Long> getSites() {
    return sites;
  }


  public void setSites(List<Long> sites) {
    this.sites = sites;
  }


  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (o == null || getClass() != o.getClass()) {
      return false;
    }
    LoadBalancerCreateResourcePermission loadBalancerCreateResourcePermission = (LoadBalancerCreateResourcePermission) o;
    return Objects.equals(this.all, loadBalancerCreateResourcePermission.all) &&
        Objects.equals(this.sites, loadBalancerCreateResourcePermission.sites);
  }

  @Override
  public int hashCode() {
    return Objects.hash(all, sites);
  }


  @Override
  public String toString() {
    StringBuilder sb = new StringBuilder();
    sb.append("class LoadBalancerCreateResourcePermission {\n");
    sb.append("    all: ").append(toIndentedString(all)).append("\n");
    sb.append("    sites: ").append(toIndentedString(sites)).append("\n");
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

