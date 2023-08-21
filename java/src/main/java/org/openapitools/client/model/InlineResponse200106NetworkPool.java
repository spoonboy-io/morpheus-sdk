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
import org.openapitools.client.model.InlineResponse200106NetworkPoolIpRanges;
import org.openapitools.client.model.InlineResponse20040AppDeployInstance;
import org.openapitools.client.model.InlineResponse20094Network;

/**
 * InlineResponse200106NetworkPool
 */
@javax.annotation.Generated(value = "org.openapitools.codegen.languages.JavaClientCodegen", date = "2023-08-21T11:26:47.028Z[GMT]")
public class InlineResponse200106NetworkPool {
  public static final String SERIALIZED_NAME_ID = "id";
  @SerializedName(SERIALIZED_NAME_ID)
  private Long id;

  public static final String SERIALIZED_NAME_TYPE = "type";
  @SerializedName(SERIALIZED_NAME_TYPE)
  private InlineResponse20094Network type;

  public static final String SERIALIZED_NAME_ACCOUNT = "account";
  @SerializedName(SERIALIZED_NAME_ACCOUNT)
  private InlineResponse20040AppDeployInstance account;

  public static final String SERIALIZED_NAME_CATEGORY = "category";
  @SerializedName(SERIALIZED_NAME_CATEGORY)
  private String category;

  public static final String SERIALIZED_NAME_CODE = "code";
  @SerializedName(SERIALIZED_NAME_CODE)
  private String code;

  public static final String SERIALIZED_NAME_NAME = "name";
  @SerializedName(SERIALIZED_NAME_NAME)
  private String name;

  public static final String SERIALIZED_NAME_DISPLAY_NAME = "displayName";
  @SerializedName(SERIALIZED_NAME_DISPLAY_NAME)
  private String displayName;

  public static final String SERIALIZED_NAME_INTERNAL_ID = "internalId";
  @SerializedName(SERIALIZED_NAME_INTERNAL_ID)
  private String internalId;

  public static final String SERIALIZED_NAME_EXTERNAL_ID = "externalId";
  @SerializedName(SERIALIZED_NAME_EXTERNAL_ID)
  private String externalId;

  public static final String SERIALIZED_NAME_DNS_DOMAIN = "dnsDomain";
  @SerializedName(SERIALIZED_NAME_DNS_DOMAIN)
  private String dnsDomain;

  public static final String SERIALIZED_NAME_DNS_SEARCH_PATH = "dnsSearchPath";
  @SerializedName(SERIALIZED_NAME_DNS_SEARCH_PATH)
  private String dnsSearchPath;

  public static final String SERIALIZED_NAME_HOST_PREFIX = "hostPrefix";
  @SerializedName(SERIALIZED_NAME_HOST_PREFIX)
  private String hostPrefix;

  public static final String SERIALIZED_NAME_HTTP_PROXY = "httpProxy";
  @SerializedName(SERIALIZED_NAME_HTTP_PROXY)
  private String httpProxy;

  public static final String SERIALIZED_NAME_DNS_SERVERS = "dnsServers";
  @SerializedName(SERIALIZED_NAME_DNS_SERVERS)
  private List<String> dnsServers = null;

  public static final String SERIALIZED_NAME_DNS_SUFFIX_LIST = "dnsSuffixList";
  @SerializedName(SERIALIZED_NAME_DNS_SUFFIX_LIST)
  private List<String> dnsSuffixList = null;

  public static final String SERIALIZED_NAME_DHCP_SERVER = "dhcpServer";
  @SerializedName(SERIALIZED_NAME_DHCP_SERVER)
  private Boolean dhcpServer;

  public static final String SERIALIZED_NAME_DHCP_IP = "dhcpIp";
  @SerializedName(SERIALIZED_NAME_DHCP_IP)
  private String dhcpIp;

  public static final String SERIALIZED_NAME_GATEWAY = "gateway";
  @SerializedName(SERIALIZED_NAME_GATEWAY)
  private String gateway;

  public static final String SERIALIZED_NAME_NETMASK = "netmask";
  @SerializedName(SERIALIZED_NAME_NETMASK)
  private String netmask;

  public static final String SERIALIZED_NAME_SUBNET_ADDRESS = "subnetAddress";
  @SerializedName(SERIALIZED_NAME_SUBNET_ADDRESS)
  private String subnetAddress;

  public static final String SERIALIZED_NAME_IP_COUNT = "ipCount";
  @SerializedName(SERIALIZED_NAME_IP_COUNT)
  private Long ipCount;

  public static final String SERIALIZED_NAME_FREE_COUNT = "freeCount";
  @SerializedName(SERIALIZED_NAME_FREE_COUNT)
  private Long freeCount;

  public static final String SERIALIZED_NAME_POOL_ENABLED = "poolEnabled";
  @SerializedName(SERIALIZED_NAME_POOL_ENABLED)
  private Boolean poolEnabled;

  public static final String SERIALIZED_NAME_TFTP_SERVER = "tftpServer";
  @SerializedName(SERIALIZED_NAME_TFTP_SERVER)
  private String tftpServer;

  public static final String SERIALIZED_NAME_BOOT_FILE = "bootFile";
  @SerializedName(SERIALIZED_NAME_BOOT_FILE)
  private String bootFile;

  public static final String SERIALIZED_NAME_REF_TYPE = "refType";
  @SerializedName(SERIALIZED_NAME_REF_TYPE)
  private String refType;

  public static final String SERIALIZED_NAME_REF_ID = "refId";
  @SerializedName(SERIALIZED_NAME_REF_ID)
  private String refId;

  public static final String SERIALIZED_NAME_PARENT_TYPE = "parentType";
  @SerializedName(SERIALIZED_NAME_PARENT_TYPE)
  private String parentType;

  public static final String SERIALIZED_NAME_PARENT_ID = "parentId";
  @SerializedName(SERIALIZED_NAME_PARENT_ID)
  private String parentId;

  public static final String SERIALIZED_NAME_POOL_GROUP = "poolGroup";
  @SerializedName(SERIALIZED_NAME_POOL_GROUP)
  private String poolGroup;

  public static final String SERIALIZED_NAME_IP_RANGES = "ipRanges";
  @SerializedName(SERIALIZED_NAME_IP_RANGES)
  private List<InlineResponse200106NetworkPoolIpRanges> ipRanges = null;


  public InlineResponse200106NetworkPool id(Long id) {
    
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


  public InlineResponse200106NetworkPool type(InlineResponse20094Network type) {
    
    this.type = type;
    return this;
  }

   /**
   * Get type
   * @return type
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public InlineResponse20094Network getType() {
    return type;
  }


  public void setType(InlineResponse20094Network type) {
    this.type = type;
  }


  public InlineResponse200106NetworkPool account(InlineResponse20040AppDeployInstance account) {
    
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


  public InlineResponse200106NetworkPool category(String category) {
    
    this.category = category;
    return this;
  }

   /**
   * Get category
   * @return category
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public String getCategory() {
    return category;
  }


  public void setCategory(String category) {
    this.category = category;
  }


  public InlineResponse200106NetworkPool code(String code) {
    
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


  public InlineResponse200106NetworkPool name(String name) {
    
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


  public InlineResponse200106NetworkPool displayName(String displayName) {
    
    this.displayName = displayName;
    return this;
  }

   /**
   * Get displayName
   * @return displayName
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public String getDisplayName() {
    return displayName;
  }


  public void setDisplayName(String displayName) {
    this.displayName = displayName;
  }


  public InlineResponse200106NetworkPool internalId(String internalId) {
    
    this.internalId = internalId;
    return this;
  }

   /**
   * Get internalId
   * @return internalId
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public String getInternalId() {
    return internalId;
  }


  public void setInternalId(String internalId) {
    this.internalId = internalId;
  }


  public InlineResponse200106NetworkPool externalId(String externalId) {
    
    this.externalId = externalId;
    return this;
  }

   /**
   * Get externalId
   * @return externalId
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public String getExternalId() {
    return externalId;
  }


  public void setExternalId(String externalId) {
    this.externalId = externalId;
  }


  public InlineResponse200106NetworkPool dnsDomain(String dnsDomain) {
    
    this.dnsDomain = dnsDomain;
    return this;
  }

   /**
   * Get dnsDomain
   * @return dnsDomain
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public String getDnsDomain() {
    return dnsDomain;
  }


  public void setDnsDomain(String dnsDomain) {
    this.dnsDomain = dnsDomain;
  }


  public InlineResponse200106NetworkPool dnsSearchPath(String dnsSearchPath) {
    
    this.dnsSearchPath = dnsSearchPath;
    return this;
  }

   /**
   * Get dnsSearchPath
   * @return dnsSearchPath
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public String getDnsSearchPath() {
    return dnsSearchPath;
  }


  public void setDnsSearchPath(String dnsSearchPath) {
    this.dnsSearchPath = dnsSearchPath;
  }


  public InlineResponse200106NetworkPool hostPrefix(String hostPrefix) {
    
    this.hostPrefix = hostPrefix;
    return this;
  }

   /**
   * Get hostPrefix
   * @return hostPrefix
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public String getHostPrefix() {
    return hostPrefix;
  }


  public void setHostPrefix(String hostPrefix) {
    this.hostPrefix = hostPrefix;
  }


  public InlineResponse200106NetworkPool httpProxy(String httpProxy) {
    
    this.httpProxy = httpProxy;
    return this;
  }

   /**
   * Get httpProxy
   * @return httpProxy
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public String getHttpProxy() {
    return httpProxy;
  }


  public void setHttpProxy(String httpProxy) {
    this.httpProxy = httpProxy;
  }


  public InlineResponse200106NetworkPool dnsServers(List<String> dnsServers) {
    
    this.dnsServers = dnsServers;
    return this;
  }

  public InlineResponse200106NetworkPool addDnsServersItem(String dnsServersItem) {
    if (this.dnsServers == null) {
      this.dnsServers = new ArrayList<String>();
    }
    this.dnsServers.add(dnsServersItem);
    return this;
  }

   /**
   * Get dnsServers
   * @return dnsServers
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public List<String> getDnsServers() {
    return dnsServers;
  }


  public void setDnsServers(List<String> dnsServers) {
    this.dnsServers = dnsServers;
  }


  public InlineResponse200106NetworkPool dnsSuffixList(List<String> dnsSuffixList) {
    
    this.dnsSuffixList = dnsSuffixList;
    return this;
  }

  public InlineResponse200106NetworkPool addDnsSuffixListItem(String dnsSuffixListItem) {
    if (this.dnsSuffixList == null) {
      this.dnsSuffixList = new ArrayList<String>();
    }
    this.dnsSuffixList.add(dnsSuffixListItem);
    return this;
  }

   /**
   * Get dnsSuffixList
   * @return dnsSuffixList
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public List<String> getDnsSuffixList() {
    return dnsSuffixList;
  }


  public void setDnsSuffixList(List<String> dnsSuffixList) {
    this.dnsSuffixList = dnsSuffixList;
  }


  public InlineResponse200106NetworkPool dhcpServer(Boolean dhcpServer) {
    
    this.dhcpServer = dhcpServer;
    return this;
  }

   /**
   * Get dhcpServer
   * @return dhcpServer
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public Boolean getDhcpServer() {
    return dhcpServer;
  }


  public void setDhcpServer(Boolean dhcpServer) {
    this.dhcpServer = dhcpServer;
  }


  public InlineResponse200106NetworkPool dhcpIp(String dhcpIp) {
    
    this.dhcpIp = dhcpIp;
    return this;
  }

   /**
   * Get dhcpIp
   * @return dhcpIp
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public String getDhcpIp() {
    return dhcpIp;
  }


  public void setDhcpIp(String dhcpIp) {
    this.dhcpIp = dhcpIp;
  }


  public InlineResponse200106NetworkPool gateway(String gateway) {
    
    this.gateway = gateway;
    return this;
  }

   /**
   * Get gateway
   * @return gateway
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public String getGateway() {
    return gateway;
  }


  public void setGateway(String gateway) {
    this.gateway = gateway;
  }


  public InlineResponse200106NetworkPool netmask(String netmask) {
    
    this.netmask = netmask;
    return this;
  }

   /**
   * Get netmask
   * @return netmask
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public String getNetmask() {
    return netmask;
  }


  public void setNetmask(String netmask) {
    this.netmask = netmask;
  }


  public InlineResponse200106NetworkPool subnetAddress(String subnetAddress) {
    
    this.subnetAddress = subnetAddress;
    return this;
  }

   /**
   * Get subnetAddress
   * @return subnetAddress
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public String getSubnetAddress() {
    return subnetAddress;
  }


  public void setSubnetAddress(String subnetAddress) {
    this.subnetAddress = subnetAddress;
  }


  public InlineResponse200106NetworkPool ipCount(Long ipCount) {
    
    this.ipCount = ipCount;
    return this;
  }

   /**
   * Get ipCount
   * @return ipCount
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public Long getIpCount() {
    return ipCount;
  }


  public void setIpCount(Long ipCount) {
    this.ipCount = ipCount;
  }


  public InlineResponse200106NetworkPool freeCount(Long freeCount) {
    
    this.freeCount = freeCount;
    return this;
  }

   /**
   * Get freeCount
   * @return freeCount
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public Long getFreeCount() {
    return freeCount;
  }


  public void setFreeCount(Long freeCount) {
    this.freeCount = freeCount;
  }


  public InlineResponse200106NetworkPool poolEnabled(Boolean poolEnabled) {
    
    this.poolEnabled = poolEnabled;
    return this;
  }

   /**
   * Get poolEnabled
   * @return poolEnabled
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public Boolean getPoolEnabled() {
    return poolEnabled;
  }


  public void setPoolEnabled(Boolean poolEnabled) {
    this.poolEnabled = poolEnabled;
  }


  public InlineResponse200106NetworkPool tftpServer(String tftpServer) {
    
    this.tftpServer = tftpServer;
    return this;
  }

   /**
   * Get tftpServer
   * @return tftpServer
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public String getTftpServer() {
    return tftpServer;
  }


  public void setTftpServer(String tftpServer) {
    this.tftpServer = tftpServer;
  }


  public InlineResponse200106NetworkPool bootFile(String bootFile) {
    
    this.bootFile = bootFile;
    return this;
  }

   /**
   * Get bootFile
   * @return bootFile
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public String getBootFile() {
    return bootFile;
  }


  public void setBootFile(String bootFile) {
    this.bootFile = bootFile;
  }


  public InlineResponse200106NetworkPool refType(String refType) {
    
    this.refType = refType;
    return this;
  }

   /**
   * Get refType
   * @return refType
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public String getRefType() {
    return refType;
  }


  public void setRefType(String refType) {
    this.refType = refType;
  }


  public InlineResponse200106NetworkPool refId(String refId) {
    
    this.refId = refId;
    return this;
  }

   /**
   * Get refId
   * @return refId
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public String getRefId() {
    return refId;
  }


  public void setRefId(String refId) {
    this.refId = refId;
  }


  public InlineResponse200106NetworkPool parentType(String parentType) {
    
    this.parentType = parentType;
    return this;
  }

   /**
   * Get parentType
   * @return parentType
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public String getParentType() {
    return parentType;
  }


  public void setParentType(String parentType) {
    this.parentType = parentType;
  }


  public InlineResponse200106NetworkPool parentId(String parentId) {
    
    this.parentId = parentId;
    return this;
  }

   /**
   * Get parentId
   * @return parentId
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public String getParentId() {
    return parentId;
  }


  public void setParentId(String parentId) {
    this.parentId = parentId;
  }


  public InlineResponse200106NetworkPool poolGroup(String poolGroup) {
    
    this.poolGroup = poolGroup;
    return this;
  }

   /**
   * Get poolGroup
   * @return poolGroup
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public String getPoolGroup() {
    return poolGroup;
  }


  public void setPoolGroup(String poolGroup) {
    this.poolGroup = poolGroup;
  }


  public InlineResponse200106NetworkPool ipRanges(List<InlineResponse200106NetworkPoolIpRanges> ipRanges) {
    
    this.ipRanges = ipRanges;
    return this;
  }

  public InlineResponse200106NetworkPool addIpRangesItem(InlineResponse200106NetworkPoolIpRanges ipRangesItem) {
    if (this.ipRanges == null) {
      this.ipRanges = new ArrayList<InlineResponse200106NetworkPoolIpRanges>();
    }
    this.ipRanges.add(ipRangesItem);
    return this;
  }

   /**
   * Get ipRanges
   * @return ipRanges
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public List<InlineResponse200106NetworkPoolIpRanges> getIpRanges() {
    return ipRanges;
  }


  public void setIpRanges(List<InlineResponse200106NetworkPoolIpRanges> ipRanges) {
    this.ipRanges = ipRanges;
  }


  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (o == null || getClass() != o.getClass()) {
      return false;
    }
    InlineResponse200106NetworkPool inlineResponse200106NetworkPool = (InlineResponse200106NetworkPool) o;
    return Objects.equals(this.id, inlineResponse200106NetworkPool.id) &&
        Objects.equals(this.type, inlineResponse200106NetworkPool.type) &&
        Objects.equals(this.account, inlineResponse200106NetworkPool.account) &&
        Objects.equals(this.category, inlineResponse200106NetworkPool.category) &&
        Objects.equals(this.code, inlineResponse200106NetworkPool.code) &&
        Objects.equals(this.name, inlineResponse200106NetworkPool.name) &&
        Objects.equals(this.displayName, inlineResponse200106NetworkPool.displayName) &&
        Objects.equals(this.internalId, inlineResponse200106NetworkPool.internalId) &&
        Objects.equals(this.externalId, inlineResponse200106NetworkPool.externalId) &&
        Objects.equals(this.dnsDomain, inlineResponse200106NetworkPool.dnsDomain) &&
        Objects.equals(this.dnsSearchPath, inlineResponse200106NetworkPool.dnsSearchPath) &&
        Objects.equals(this.hostPrefix, inlineResponse200106NetworkPool.hostPrefix) &&
        Objects.equals(this.httpProxy, inlineResponse200106NetworkPool.httpProxy) &&
        Objects.equals(this.dnsServers, inlineResponse200106NetworkPool.dnsServers) &&
        Objects.equals(this.dnsSuffixList, inlineResponse200106NetworkPool.dnsSuffixList) &&
        Objects.equals(this.dhcpServer, inlineResponse200106NetworkPool.dhcpServer) &&
        Objects.equals(this.dhcpIp, inlineResponse200106NetworkPool.dhcpIp) &&
        Objects.equals(this.gateway, inlineResponse200106NetworkPool.gateway) &&
        Objects.equals(this.netmask, inlineResponse200106NetworkPool.netmask) &&
        Objects.equals(this.subnetAddress, inlineResponse200106NetworkPool.subnetAddress) &&
        Objects.equals(this.ipCount, inlineResponse200106NetworkPool.ipCount) &&
        Objects.equals(this.freeCount, inlineResponse200106NetworkPool.freeCount) &&
        Objects.equals(this.poolEnabled, inlineResponse200106NetworkPool.poolEnabled) &&
        Objects.equals(this.tftpServer, inlineResponse200106NetworkPool.tftpServer) &&
        Objects.equals(this.bootFile, inlineResponse200106NetworkPool.bootFile) &&
        Objects.equals(this.refType, inlineResponse200106NetworkPool.refType) &&
        Objects.equals(this.refId, inlineResponse200106NetworkPool.refId) &&
        Objects.equals(this.parentType, inlineResponse200106NetworkPool.parentType) &&
        Objects.equals(this.parentId, inlineResponse200106NetworkPool.parentId) &&
        Objects.equals(this.poolGroup, inlineResponse200106NetworkPool.poolGroup) &&
        Objects.equals(this.ipRanges, inlineResponse200106NetworkPool.ipRanges);
  }

  @Override
  public int hashCode() {
    return Objects.hash(id, type, account, category, code, name, displayName, internalId, externalId, dnsDomain, dnsSearchPath, hostPrefix, httpProxy, dnsServers, dnsSuffixList, dhcpServer, dhcpIp, gateway, netmask, subnetAddress, ipCount, freeCount, poolEnabled, tftpServer, bootFile, refType, refId, parentType, parentId, poolGroup, ipRanges);
  }


  @Override
  public String toString() {
    StringBuilder sb = new StringBuilder();
    sb.append("class InlineResponse200106NetworkPool {\n");
    sb.append("    id: ").append(toIndentedString(id)).append("\n");
    sb.append("    type: ").append(toIndentedString(type)).append("\n");
    sb.append("    account: ").append(toIndentedString(account)).append("\n");
    sb.append("    category: ").append(toIndentedString(category)).append("\n");
    sb.append("    code: ").append(toIndentedString(code)).append("\n");
    sb.append("    name: ").append(toIndentedString(name)).append("\n");
    sb.append("    displayName: ").append(toIndentedString(displayName)).append("\n");
    sb.append("    internalId: ").append(toIndentedString(internalId)).append("\n");
    sb.append("    externalId: ").append(toIndentedString(externalId)).append("\n");
    sb.append("    dnsDomain: ").append(toIndentedString(dnsDomain)).append("\n");
    sb.append("    dnsSearchPath: ").append(toIndentedString(dnsSearchPath)).append("\n");
    sb.append("    hostPrefix: ").append(toIndentedString(hostPrefix)).append("\n");
    sb.append("    httpProxy: ").append(toIndentedString(httpProxy)).append("\n");
    sb.append("    dnsServers: ").append(toIndentedString(dnsServers)).append("\n");
    sb.append("    dnsSuffixList: ").append(toIndentedString(dnsSuffixList)).append("\n");
    sb.append("    dhcpServer: ").append(toIndentedString(dhcpServer)).append("\n");
    sb.append("    dhcpIp: ").append(toIndentedString(dhcpIp)).append("\n");
    sb.append("    gateway: ").append(toIndentedString(gateway)).append("\n");
    sb.append("    netmask: ").append(toIndentedString(netmask)).append("\n");
    sb.append("    subnetAddress: ").append(toIndentedString(subnetAddress)).append("\n");
    sb.append("    ipCount: ").append(toIndentedString(ipCount)).append("\n");
    sb.append("    freeCount: ").append(toIndentedString(freeCount)).append("\n");
    sb.append("    poolEnabled: ").append(toIndentedString(poolEnabled)).append("\n");
    sb.append("    tftpServer: ").append(toIndentedString(tftpServer)).append("\n");
    sb.append("    bootFile: ").append(toIndentedString(bootFile)).append("\n");
    sb.append("    refType: ").append(toIndentedString(refType)).append("\n");
    sb.append("    refId: ").append(toIndentedString(refId)).append("\n");
    sb.append("    parentType: ").append(toIndentedString(parentType)).append("\n");
    sb.append("    parentId: ").append(toIndentedString(parentId)).append("\n");
    sb.append("    poolGroup: ").append(toIndentedString(poolGroup)).append("\n");
    sb.append("    ipRanges: ").append(toIndentedString(ipRanges)).append("\n");
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

