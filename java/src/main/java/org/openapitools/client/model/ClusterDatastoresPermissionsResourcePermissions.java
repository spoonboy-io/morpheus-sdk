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
import org.openapitools.client.model.ApiBlueprintsIdUpdatePermissionsResourcePermissionSites;
import org.openapitools.client.model.ResourcePermissionsSites;

/**
 * ClusterDatastoresPermissionsResourcePermissions
 */
@javax.annotation.Generated(value = "org.openapitools.codegen.languages.JavaClientCodegen", date = "2023-08-21T11:26:47.028Z[GMT]")
public class ClusterDatastoresPermissionsResourcePermissions {
  public static final String SERIALIZED_NAME_DEFAULT_STORE = "defaultStore";
  @SerializedName(SERIALIZED_NAME_DEFAULT_STORE)
  private Boolean defaultStore;

  public static final String SERIALIZED_NAME_ALL_PLANS = "allPlans";
  @SerializedName(SERIALIZED_NAME_ALL_PLANS)
  private Boolean allPlans;

  public static final String SERIALIZED_NAME_DEFAULT_TARGET = "defaultTarget";
  @SerializedName(SERIALIZED_NAME_DEFAULT_TARGET)
  private Boolean defaultTarget;

  public static final String SERIALIZED_NAME_CAN_MANAGE = "canManage";
  @SerializedName(SERIALIZED_NAME_CAN_MANAGE)
  private Boolean canManage;

  public static final String SERIALIZED_NAME_ALL = "all";
  @SerializedName(SERIALIZED_NAME_ALL)
  private Boolean all;

  public static final String SERIALIZED_NAME_ACCOUNT = "account";
  @SerializedName(SERIALIZED_NAME_ACCOUNT)
  private ApiBlueprintsIdUpdatePermissionsResourcePermissionSites account;

  public static final String SERIALIZED_NAME_SITES = "sites";
  @SerializedName(SERIALIZED_NAME_SITES)
  private List<ResourcePermissionsSites> sites = null;

  public static final String SERIALIZED_NAME_PLANS = "plans";
  @SerializedName(SERIALIZED_NAME_PLANS)
  private List<ResourcePermissionsSites> plans = null;


  public ClusterDatastoresPermissionsResourcePermissions defaultStore(Boolean defaultStore) {
    
    this.defaultStore = defaultStore;
    return this;
  }

   /**
   * Get defaultStore
   * @return defaultStore
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public Boolean getDefaultStore() {
    return defaultStore;
  }


  public void setDefaultStore(Boolean defaultStore) {
    this.defaultStore = defaultStore;
  }


  public ClusterDatastoresPermissionsResourcePermissions allPlans(Boolean allPlans) {
    
    this.allPlans = allPlans;
    return this;
  }

   /**
   * Get allPlans
   * @return allPlans
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public Boolean getAllPlans() {
    return allPlans;
  }


  public void setAllPlans(Boolean allPlans) {
    this.allPlans = allPlans;
  }


  public ClusterDatastoresPermissionsResourcePermissions defaultTarget(Boolean defaultTarget) {
    
    this.defaultTarget = defaultTarget;
    return this;
  }

   /**
   * Get defaultTarget
   * @return defaultTarget
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public Boolean getDefaultTarget() {
    return defaultTarget;
  }


  public void setDefaultTarget(Boolean defaultTarget) {
    this.defaultTarget = defaultTarget;
  }


  public ClusterDatastoresPermissionsResourcePermissions canManage(Boolean canManage) {
    
    this.canManage = canManage;
    return this;
  }

   /**
   * Get canManage
   * @return canManage
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public Boolean getCanManage() {
    return canManage;
  }


  public void setCanManage(Boolean canManage) {
    this.canManage = canManage;
  }


  public ClusterDatastoresPermissionsResourcePermissions all(Boolean all) {
    
    this.all = all;
    return this;
  }

   /**
   * Get all
   * @return all
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public Boolean getAll() {
    return all;
  }


  public void setAll(Boolean all) {
    this.all = all;
  }


  public ClusterDatastoresPermissionsResourcePermissions account(ApiBlueprintsIdUpdatePermissionsResourcePermissionSites account) {
    
    this.account = account;
    return this;
  }

   /**
   * Get account
   * @return account
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public ApiBlueprintsIdUpdatePermissionsResourcePermissionSites getAccount() {
    return account;
  }


  public void setAccount(ApiBlueprintsIdUpdatePermissionsResourcePermissionSites account) {
    this.account = account;
  }


  public ClusterDatastoresPermissionsResourcePermissions sites(List<ResourcePermissionsSites> sites) {
    
    this.sites = sites;
    return this;
  }

  public ClusterDatastoresPermissionsResourcePermissions addSitesItem(ResourcePermissionsSites sitesItem) {
    if (this.sites == null) {
      this.sites = new ArrayList<ResourcePermissionsSites>();
    }
    this.sites.add(sitesItem);
    return this;
  }

   /**
   * Get sites
   * @return sites
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public List<ResourcePermissionsSites> getSites() {
    return sites;
  }


  public void setSites(List<ResourcePermissionsSites> sites) {
    this.sites = sites;
  }


  public ClusterDatastoresPermissionsResourcePermissions plans(List<ResourcePermissionsSites> plans) {
    
    this.plans = plans;
    return this;
  }

  public ClusterDatastoresPermissionsResourcePermissions addPlansItem(ResourcePermissionsSites plansItem) {
    if (this.plans == null) {
      this.plans = new ArrayList<ResourcePermissionsSites>();
    }
    this.plans.add(plansItem);
    return this;
  }

   /**
   * Get plans
   * @return plans
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public List<ResourcePermissionsSites> getPlans() {
    return plans;
  }


  public void setPlans(List<ResourcePermissionsSites> plans) {
    this.plans = plans;
  }


  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (o == null || getClass() != o.getClass()) {
      return false;
    }
    ClusterDatastoresPermissionsResourcePermissions clusterDatastoresPermissionsResourcePermissions = (ClusterDatastoresPermissionsResourcePermissions) o;
    return Objects.equals(this.defaultStore, clusterDatastoresPermissionsResourcePermissions.defaultStore) &&
        Objects.equals(this.allPlans, clusterDatastoresPermissionsResourcePermissions.allPlans) &&
        Objects.equals(this.defaultTarget, clusterDatastoresPermissionsResourcePermissions.defaultTarget) &&
        Objects.equals(this.canManage, clusterDatastoresPermissionsResourcePermissions.canManage) &&
        Objects.equals(this.all, clusterDatastoresPermissionsResourcePermissions.all) &&
        Objects.equals(this.account, clusterDatastoresPermissionsResourcePermissions.account) &&
        Objects.equals(this.sites, clusterDatastoresPermissionsResourcePermissions.sites) &&
        Objects.equals(this.plans, clusterDatastoresPermissionsResourcePermissions.plans);
  }

  @Override
  public int hashCode() {
    return Objects.hash(defaultStore, allPlans, defaultTarget, canManage, all, account, sites, plans);
  }


  @Override
  public String toString() {
    StringBuilder sb = new StringBuilder();
    sb.append("class ClusterDatastoresPermissionsResourcePermissions {\n");
    sb.append("    defaultStore: ").append(toIndentedString(defaultStore)).append("\n");
    sb.append("    allPlans: ").append(toIndentedString(allPlans)).append("\n");
    sb.append("    defaultTarget: ").append(toIndentedString(defaultTarget)).append("\n");
    sb.append("    canManage: ").append(toIndentedString(canManage)).append("\n");
    sb.append("    all: ").append(toIndentedString(all)).append("\n");
    sb.append("    account: ").append(toIndentedString(account)).append("\n");
    sb.append("    sites: ").append(toIndentedString(sites)).append("\n");
    sb.append("    plans: ").append(toIndentedString(plans)).append("\n");
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

