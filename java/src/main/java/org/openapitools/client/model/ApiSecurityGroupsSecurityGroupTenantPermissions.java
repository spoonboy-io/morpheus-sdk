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
 * ApiSecurityGroupsSecurityGroupTenantPermissions
 */
@javax.annotation.Generated(value = "org.openapitools.codegen.languages.JavaClientCodegen", date = "2023-08-21T11:26:47.028Z[GMT]")
public class ApiSecurityGroupsSecurityGroupTenantPermissions {
  public static final String SERIALIZED_NAME_ACCOUNTS = "accounts";
  @SerializedName(SERIALIZED_NAME_ACCOUNTS)
  private List<Long> accounts = null;

  public static final String SERIALIZED_NAME_CAN_MANAGE_ACCOUNTS = "canManageAccounts";
  @SerializedName(SERIALIZED_NAME_CAN_MANAGE_ACCOUNTS)
  private List<Long> canManageAccounts = null;


  public ApiSecurityGroupsSecurityGroupTenantPermissions accounts(List<Long> accounts) {
    
    this.accounts = accounts;
    return this;
  }

  public ApiSecurityGroupsSecurityGroupTenantPermissions addAccountsItem(Long accountsItem) {
    if (this.accounts == null) {
      this.accounts = new ArrayList<Long>();
    }
    this.accounts.add(accountsItem);
    return this;
  }

   /**
   * Array of tenant account ids that are allowed access
   * @return accounts
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(example = "[1,3]", value = "Array of tenant account ids that are allowed access")

  public List<Long> getAccounts() {
    return accounts;
  }


  public void setAccounts(List<Long> accounts) {
    this.accounts = accounts;
  }


  public ApiSecurityGroupsSecurityGroupTenantPermissions canManageAccounts(List<Long> canManageAccounts) {
    
    this.canManageAccounts = canManageAccounts;
    return this;
  }

  public ApiSecurityGroupsSecurityGroupTenantPermissions addCanManageAccountsItem(Long canManageAccountsItem) {
    if (this.canManageAccounts == null) {
      this.canManageAccounts = new ArrayList<Long>();
    }
    this.canManageAccounts.add(canManageAccountsItem);
    return this;
  }

   /**
   * Array of tenant account ids that can manage
   * @return canManageAccounts
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(example = "[1,3]", value = "Array of tenant account ids that can manage")

  public List<Long> getCanManageAccounts() {
    return canManageAccounts;
  }


  public void setCanManageAccounts(List<Long> canManageAccounts) {
    this.canManageAccounts = canManageAccounts;
  }


  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (o == null || getClass() != o.getClass()) {
      return false;
    }
    ApiSecurityGroupsSecurityGroupTenantPermissions apiSecurityGroupsSecurityGroupTenantPermissions = (ApiSecurityGroupsSecurityGroupTenantPermissions) o;
    return Objects.equals(this.accounts, apiSecurityGroupsSecurityGroupTenantPermissions.accounts) &&
        Objects.equals(this.canManageAccounts, apiSecurityGroupsSecurityGroupTenantPermissions.canManageAccounts);
  }

  @Override
  public int hashCode() {
    return Objects.hash(accounts, canManageAccounts);
  }


  @Override
  public String toString() {
    StringBuilder sb = new StringBuilder();
    sb.append("class ApiSecurityGroupsSecurityGroupTenantPermissions {\n");
    sb.append("    accounts: ").append(toIndentedString(accounts)).append("\n");
    sb.append("    canManageAccounts: ").append(toIndentedString(canManageAccounts)).append("\n");
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
