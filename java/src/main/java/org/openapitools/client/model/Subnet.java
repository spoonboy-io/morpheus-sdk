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
import org.openapitools.client.model.AppStateInputProviders;
import org.openapitools.client.model.InlineResponse20040AppDeployInstance;
import org.openapitools.client.model.InlineResponse20079LoadBalancerMonitorLoadBalancerType;
import org.openapitools.client.model.SubnetResourcePermission;

/**
 * Subnet
 */
@javax.annotation.Generated(value = "org.openapitools.codegen.languages.JavaClientCodegen", date = "2023-08-21T11:26:47.028Z[GMT]")
public class Subnet {
  public static final String SERIALIZED_NAME_ID = "id";
  @SerializedName(SERIALIZED_NAME_ID)
  private Long id;

  public static final String SERIALIZED_NAME_CODE = "code";
  @SerializedName(SERIALIZED_NAME_CODE)
  private String code;

  public static final String SERIALIZED_NAME_NAME = "name";
  @SerializedName(SERIALIZED_NAME_NAME)
  private String name;

  public static final String SERIALIZED_NAME_LABELS = "labels";
  @SerializedName(SERIALIZED_NAME_LABELS)
  private List<String> labels = null;

  public static final String SERIALIZED_NAME_ACTIVE = "active";
  @SerializedName(SERIALIZED_NAME_ACTIVE)
  private Boolean active;

  public static final String SERIALIZED_NAME_DESCRIPTION = "description";
  @SerializedName(SERIALIZED_NAME_DESCRIPTION)
  private String description;

  public static final String SERIALIZED_NAME_EXTERNAL_ID = "externalId";
  @SerializedName(SERIALIZED_NAME_EXTERNAL_ID)
  private String externalId;

  public static final String SERIALIZED_NAME_UNIQUE_ID = "uniqueId";
  @SerializedName(SERIALIZED_NAME_UNIQUE_ID)
  private String uniqueId;

  public static final String SERIALIZED_NAME_ADDRESS_PREFIX = "addressPrefix";
  @SerializedName(SERIALIZED_NAME_ADDRESS_PREFIX)
  private String addressPrefix;

  public static final String SERIALIZED_NAME_CIDR = "cidr";
  @SerializedName(SERIALIZED_NAME_CIDR)
  private String cidr;

  public static final String SERIALIZED_NAME_GATEWAY = "gateway";
  @SerializedName(SERIALIZED_NAME_GATEWAY)
  private String gateway;

  public static final String SERIALIZED_NAME_NETMASK = "netmask";
  @SerializedName(SERIALIZED_NAME_NETMASK)
  private String netmask;

  public static final String SERIALIZED_NAME_SUBNET_ADDRESS = "subnetAddress";
  @SerializedName(SERIALIZED_NAME_SUBNET_ADDRESS)
  private String subnetAddress;

  public static final String SERIALIZED_NAME_TFTP_SERVER = "tftpServer";
  @SerializedName(SERIALIZED_NAME_TFTP_SERVER)
  private String tftpServer;

  public static final String SERIALIZED_NAME_BOOT_FILE = "bootFile";
  @SerializedName(SERIALIZED_NAME_BOOT_FILE)
  private String bootFile;

  public static final String SERIALIZED_NAME_POOL = "pool";
  @SerializedName(SERIALIZED_NAME_POOL)
  private String pool;

  public static final String SERIALIZED_NAME_DHCP_SERVER = "dhcpServer";
  @SerializedName(SERIALIZED_NAME_DHCP_SERVER)
  private Boolean dhcpServer;

  public static final String SERIALIZED_NAME_HAS_FLOATING_IPS = "hasFloatingIps";
  @SerializedName(SERIALIZED_NAME_HAS_FLOATING_IPS)
  private Boolean hasFloatingIps;

  public static final String SERIALIZED_NAME_DHCP_IP = "dhcpIp";
  @SerializedName(SERIALIZED_NAME_DHCP_IP)
  private String dhcpIp;

  public static final String SERIALIZED_NAME_DNS_PRIMARY = "dnsPrimary";
  @SerializedName(SERIALIZED_NAME_DNS_PRIMARY)
  private String dnsPrimary;

  public static final String SERIALIZED_NAME_DNS_SECONDARY = "dnsSecondary";
  @SerializedName(SERIALIZED_NAME_DNS_SECONDARY)
  private String dnsSecondary;

  public static final String SERIALIZED_NAME_DHCP_START = "dhcpStart";
  @SerializedName(SERIALIZED_NAME_DHCP_START)
  private String dhcpStart;

  public static final String SERIALIZED_NAME_DHCP_END = "dhcpEnd";
  @SerializedName(SERIALIZED_NAME_DHCP_END)
  private String dhcpEnd;

  public static final String SERIALIZED_NAME_DHCP_RANGE = "dhcpRange";
  @SerializedName(SERIALIZED_NAME_DHCP_RANGE)
  private String dhcpRange;

  public static final String SERIALIZED_NAME_NETWORK_PROXY = "networkProxy";
  @SerializedName(SERIALIZED_NAME_NETWORK_PROXY)
  private String networkProxy;

  public static final String SERIALIZED_NAME_NETWORK_DOMAIN = "networkDomain";
  @SerializedName(SERIALIZED_NAME_NETWORK_DOMAIN)
  private String networkDomain;

  public static final String SERIALIZED_NAME_SEARCH_DOMAINS = "searchDomains";
  @SerializedName(SERIALIZED_NAME_SEARCH_DOMAINS)
  private String searchDomains;

  public static final String SERIALIZED_NAME_DEFAULT_NETWORK = "defaultNetwork";
  @SerializedName(SERIALIZED_NAME_DEFAULT_NETWORK)
  private Boolean defaultNetwork;

  public static final String SERIALIZED_NAME_ASSIGN_PUBLIC_IP = "assignPublicIp";
  @SerializedName(SERIALIZED_NAME_ASSIGN_PUBLIC_IP)
  private Boolean assignPublicIp;

  public static final String SERIALIZED_NAME_VISIBILITY = "visibility";
  @SerializedName(SERIALIZED_NAME_VISIBILITY)
  private String visibility;

  public static final String SERIALIZED_NAME_STATUS = "status";
  @SerializedName(SERIALIZED_NAME_STATUS)
  private AppStateInputProviders status;

  public static final String SERIALIZED_NAME_NETWORK = "network";
  @SerializedName(SERIALIZED_NAME_NETWORK)
  private InlineResponse20040AppDeployInstance network;

  public static final String SERIALIZED_NAME_TYPE = "type";
  @SerializedName(SERIALIZED_NAME_TYPE)
  private InlineResponse20079LoadBalancerMonitorLoadBalancerType type;

  public static final String SERIALIZED_NAME_ACCOUNT = "account";
  @SerializedName(SERIALIZED_NAME_ACCOUNT)
  private InlineResponse20040AppDeployInstance account;

  public static final String SERIALIZED_NAME_SECURITY_GROUPS = "securityGroups";
  @SerializedName(SERIALIZED_NAME_SECURITY_GROUPS)
  private List<Object> securityGroups = null;

  public static final String SERIALIZED_NAME_TENANTS = "tenants";
  @SerializedName(SERIALIZED_NAME_TENANTS)
  private List<InlineResponse20040AppDeployInstance> tenants = null;

  public static final String SERIALIZED_NAME_RESOURCE_PERMISSION = "resourcePermission";
  @SerializedName(SERIALIZED_NAME_RESOURCE_PERMISSION)
  private SubnetResourcePermission resourcePermission;


  public Subnet id(Long id) {
    
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


  public Subnet code(String code) {
    
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


  public Subnet name(String name) {
    
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


  public Subnet labels(List<String> labels) {
    
    this.labels = labels;
    return this;
  }

  public Subnet addLabelsItem(String labelsItem) {
    if (this.labels == null) {
      this.labels = new ArrayList<String>();
    }
    this.labels.add(labelsItem);
    return this;
  }

   /**
   * Get labels
   * @return labels
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public List<String> getLabels() {
    return labels;
  }


  public void setLabels(List<String> labels) {
    this.labels = labels;
  }


  public Subnet active(Boolean active) {
    
    this.active = active;
    return this;
  }

   /**
   * Get active
   * @return active
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public Boolean getActive() {
    return active;
  }


  public void setActive(Boolean active) {
    this.active = active;
  }


  public Subnet description(String description) {
    
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


  public Subnet externalId(String externalId) {
    
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


  public Subnet uniqueId(String uniqueId) {
    
    this.uniqueId = uniqueId;
    return this;
  }

   /**
   * Get uniqueId
   * @return uniqueId
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public String getUniqueId() {
    return uniqueId;
  }


  public void setUniqueId(String uniqueId) {
    this.uniqueId = uniqueId;
  }


  public Subnet addressPrefix(String addressPrefix) {
    
    this.addressPrefix = addressPrefix;
    return this;
  }

   /**
   * Get addressPrefix
   * @return addressPrefix
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public String getAddressPrefix() {
    return addressPrefix;
  }


  public void setAddressPrefix(String addressPrefix) {
    this.addressPrefix = addressPrefix;
  }


  public Subnet cidr(String cidr) {
    
    this.cidr = cidr;
    return this;
  }

   /**
   * Get cidr
   * @return cidr
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public String getCidr() {
    return cidr;
  }


  public void setCidr(String cidr) {
    this.cidr = cidr;
  }


  public Subnet gateway(String gateway) {
    
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


  public Subnet netmask(String netmask) {
    
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


  public Subnet subnetAddress(String subnetAddress) {
    
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


  public Subnet tftpServer(String tftpServer) {
    
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


  public Subnet bootFile(String bootFile) {
    
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


  public Subnet pool(String pool) {
    
    this.pool = pool;
    return this;
  }

   /**
   * Get pool
   * @return pool
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public String getPool() {
    return pool;
  }


  public void setPool(String pool) {
    this.pool = pool;
  }


  public Subnet dhcpServer(Boolean dhcpServer) {
    
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


  public Subnet hasFloatingIps(Boolean hasFloatingIps) {
    
    this.hasFloatingIps = hasFloatingIps;
    return this;
  }

   /**
   * Get hasFloatingIps
   * @return hasFloatingIps
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public Boolean getHasFloatingIps() {
    return hasFloatingIps;
  }


  public void setHasFloatingIps(Boolean hasFloatingIps) {
    this.hasFloatingIps = hasFloatingIps;
  }


  public Subnet dhcpIp(String dhcpIp) {
    
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


  public Subnet dnsPrimary(String dnsPrimary) {
    
    this.dnsPrimary = dnsPrimary;
    return this;
  }

   /**
   * Get dnsPrimary
   * @return dnsPrimary
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public String getDnsPrimary() {
    return dnsPrimary;
  }


  public void setDnsPrimary(String dnsPrimary) {
    this.dnsPrimary = dnsPrimary;
  }


  public Subnet dnsSecondary(String dnsSecondary) {
    
    this.dnsSecondary = dnsSecondary;
    return this;
  }

   /**
   * Get dnsSecondary
   * @return dnsSecondary
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public String getDnsSecondary() {
    return dnsSecondary;
  }


  public void setDnsSecondary(String dnsSecondary) {
    this.dnsSecondary = dnsSecondary;
  }


  public Subnet dhcpStart(String dhcpStart) {
    
    this.dhcpStart = dhcpStart;
    return this;
  }

   /**
   * Get dhcpStart
   * @return dhcpStart
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public String getDhcpStart() {
    return dhcpStart;
  }


  public void setDhcpStart(String dhcpStart) {
    this.dhcpStart = dhcpStart;
  }


  public Subnet dhcpEnd(String dhcpEnd) {
    
    this.dhcpEnd = dhcpEnd;
    return this;
  }

   /**
   * Get dhcpEnd
   * @return dhcpEnd
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public String getDhcpEnd() {
    return dhcpEnd;
  }


  public void setDhcpEnd(String dhcpEnd) {
    this.dhcpEnd = dhcpEnd;
  }


  public Subnet dhcpRange(String dhcpRange) {
    
    this.dhcpRange = dhcpRange;
    return this;
  }

   /**
   * Get dhcpRange
   * @return dhcpRange
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public String getDhcpRange() {
    return dhcpRange;
  }


  public void setDhcpRange(String dhcpRange) {
    this.dhcpRange = dhcpRange;
  }


  public Subnet networkProxy(String networkProxy) {
    
    this.networkProxy = networkProxy;
    return this;
  }

   /**
   * Get networkProxy
   * @return networkProxy
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public String getNetworkProxy() {
    return networkProxy;
  }


  public void setNetworkProxy(String networkProxy) {
    this.networkProxy = networkProxy;
  }


  public Subnet networkDomain(String networkDomain) {
    
    this.networkDomain = networkDomain;
    return this;
  }

   /**
   * Get networkDomain
   * @return networkDomain
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public String getNetworkDomain() {
    return networkDomain;
  }


  public void setNetworkDomain(String networkDomain) {
    this.networkDomain = networkDomain;
  }


  public Subnet searchDomains(String searchDomains) {
    
    this.searchDomains = searchDomains;
    return this;
  }

   /**
   * Get searchDomains
   * @return searchDomains
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public String getSearchDomains() {
    return searchDomains;
  }


  public void setSearchDomains(String searchDomains) {
    this.searchDomains = searchDomains;
  }


  public Subnet defaultNetwork(Boolean defaultNetwork) {
    
    this.defaultNetwork = defaultNetwork;
    return this;
  }

   /**
   * Get defaultNetwork
   * @return defaultNetwork
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public Boolean getDefaultNetwork() {
    return defaultNetwork;
  }


  public void setDefaultNetwork(Boolean defaultNetwork) {
    this.defaultNetwork = defaultNetwork;
  }


  public Subnet assignPublicIp(Boolean assignPublicIp) {
    
    this.assignPublicIp = assignPublicIp;
    return this;
  }

   /**
   * Get assignPublicIp
   * @return assignPublicIp
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public Boolean getAssignPublicIp() {
    return assignPublicIp;
  }


  public void setAssignPublicIp(Boolean assignPublicIp) {
    this.assignPublicIp = assignPublicIp;
  }


  public Subnet visibility(String visibility) {
    
    this.visibility = visibility;
    return this;
  }

   /**
   * Get visibility
   * @return visibility
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public String getVisibility() {
    return visibility;
  }


  public void setVisibility(String visibility) {
    this.visibility = visibility;
  }


  public Subnet status(AppStateInputProviders status) {
    
    this.status = status;
    return this;
  }

   /**
   * Get status
   * @return status
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public AppStateInputProviders getStatus() {
    return status;
  }


  public void setStatus(AppStateInputProviders status) {
    this.status = status;
  }


  public Subnet network(InlineResponse20040AppDeployInstance network) {
    
    this.network = network;
    return this;
  }

   /**
   * Get network
   * @return network
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public InlineResponse20040AppDeployInstance getNetwork() {
    return network;
  }


  public void setNetwork(InlineResponse20040AppDeployInstance network) {
    this.network = network;
  }


  public Subnet type(InlineResponse20079LoadBalancerMonitorLoadBalancerType type) {
    
    this.type = type;
    return this;
  }

   /**
   * Get type
   * @return type
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public InlineResponse20079LoadBalancerMonitorLoadBalancerType getType() {
    return type;
  }


  public void setType(InlineResponse20079LoadBalancerMonitorLoadBalancerType type) {
    this.type = type;
  }


  public Subnet account(InlineResponse20040AppDeployInstance account) {
    
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


  public Subnet securityGroups(List<Object> securityGroups) {
    
    this.securityGroups = securityGroups;
    return this;
  }

  public Subnet addSecurityGroupsItem(Object securityGroupsItem) {
    if (this.securityGroups == null) {
      this.securityGroups = new ArrayList<Object>();
    }
    this.securityGroups.add(securityGroupsItem);
    return this;
  }

   /**
   * Get securityGroups
   * @return securityGroups
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public List<Object> getSecurityGroups() {
    return securityGroups;
  }


  public void setSecurityGroups(List<Object> securityGroups) {
    this.securityGroups = securityGroups;
  }


  public Subnet tenants(List<InlineResponse20040AppDeployInstance> tenants) {
    
    this.tenants = tenants;
    return this;
  }

  public Subnet addTenantsItem(InlineResponse20040AppDeployInstance tenantsItem) {
    if (this.tenants == null) {
      this.tenants = new ArrayList<InlineResponse20040AppDeployInstance>();
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

  public List<InlineResponse20040AppDeployInstance> getTenants() {
    return tenants;
  }


  public void setTenants(List<InlineResponse20040AppDeployInstance> tenants) {
    this.tenants = tenants;
  }


  public Subnet resourcePermission(SubnetResourcePermission resourcePermission) {
    
    this.resourcePermission = resourcePermission;
    return this;
  }

   /**
   * Get resourcePermission
   * @return resourcePermission
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public SubnetResourcePermission getResourcePermission() {
    return resourcePermission;
  }


  public void setResourcePermission(SubnetResourcePermission resourcePermission) {
    this.resourcePermission = resourcePermission;
  }


  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (o == null || getClass() != o.getClass()) {
      return false;
    }
    Subnet subnet = (Subnet) o;
    return Objects.equals(this.id, subnet.id) &&
        Objects.equals(this.code, subnet.code) &&
        Objects.equals(this.name, subnet.name) &&
        Objects.equals(this.labels, subnet.labels) &&
        Objects.equals(this.active, subnet.active) &&
        Objects.equals(this.description, subnet.description) &&
        Objects.equals(this.externalId, subnet.externalId) &&
        Objects.equals(this.uniqueId, subnet.uniqueId) &&
        Objects.equals(this.addressPrefix, subnet.addressPrefix) &&
        Objects.equals(this.cidr, subnet.cidr) &&
        Objects.equals(this.gateway, subnet.gateway) &&
        Objects.equals(this.netmask, subnet.netmask) &&
        Objects.equals(this.subnetAddress, subnet.subnetAddress) &&
        Objects.equals(this.tftpServer, subnet.tftpServer) &&
        Objects.equals(this.bootFile, subnet.bootFile) &&
        Objects.equals(this.pool, subnet.pool) &&
        Objects.equals(this.dhcpServer, subnet.dhcpServer) &&
        Objects.equals(this.hasFloatingIps, subnet.hasFloatingIps) &&
        Objects.equals(this.dhcpIp, subnet.dhcpIp) &&
        Objects.equals(this.dnsPrimary, subnet.dnsPrimary) &&
        Objects.equals(this.dnsSecondary, subnet.dnsSecondary) &&
        Objects.equals(this.dhcpStart, subnet.dhcpStart) &&
        Objects.equals(this.dhcpEnd, subnet.dhcpEnd) &&
        Objects.equals(this.dhcpRange, subnet.dhcpRange) &&
        Objects.equals(this.networkProxy, subnet.networkProxy) &&
        Objects.equals(this.networkDomain, subnet.networkDomain) &&
        Objects.equals(this.searchDomains, subnet.searchDomains) &&
        Objects.equals(this.defaultNetwork, subnet.defaultNetwork) &&
        Objects.equals(this.assignPublicIp, subnet.assignPublicIp) &&
        Objects.equals(this.visibility, subnet.visibility) &&
        Objects.equals(this.status, subnet.status) &&
        Objects.equals(this.network, subnet.network) &&
        Objects.equals(this.type, subnet.type) &&
        Objects.equals(this.account, subnet.account) &&
        Objects.equals(this.securityGroups, subnet.securityGroups) &&
        Objects.equals(this.tenants, subnet.tenants) &&
        Objects.equals(this.resourcePermission, subnet.resourcePermission);
  }

  @Override
  public int hashCode() {
    return Objects.hash(id, code, name, labels, active, description, externalId, uniqueId, addressPrefix, cidr, gateway, netmask, subnetAddress, tftpServer, bootFile, pool, dhcpServer, hasFloatingIps, dhcpIp, dnsPrimary, dnsSecondary, dhcpStart, dhcpEnd, dhcpRange, networkProxy, networkDomain, searchDomains, defaultNetwork, assignPublicIp, visibility, status, network, type, account, securityGroups, tenants, resourcePermission);
  }


  @Override
  public String toString() {
    StringBuilder sb = new StringBuilder();
    sb.append("class Subnet {\n");
    sb.append("    id: ").append(toIndentedString(id)).append("\n");
    sb.append("    code: ").append(toIndentedString(code)).append("\n");
    sb.append("    name: ").append(toIndentedString(name)).append("\n");
    sb.append("    labels: ").append(toIndentedString(labels)).append("\n");
    sb.append("    active: ").append(toIndentedString(active)).append("\n");
    sb.append("    description: ").append(toIndentedString(description)).append("\n");
    sb.append("    externalId: ").append(toIndentedString(externalId)).append("\n");
    sb.append("    uniqueId: ").append(toIndentedString(uniqueId)).append("\n");
    sb.append("    addressPrefix: ").append(toIndentedString(addressPrefix)).append("\n");
    sb.append("    cidr: ").append(toIndentedString(cidr)).append("\n");
    sb.append("    gateway: ").append(toIndentedString(gateway)).append("\n");
    sb.append("    netmask: ").append(toIndentedString(netmask)).append("\n");
    sb.append("    subnetAddress: ").append(toIndentedString(subnetAddress)).append("\n");
    sb.append("    tftpServer: ").append(toIndentedString(tftpServer)).append("\n");
    sb.append("    bootFile: ").append(toIndentedString(bootFile)).append("\n");
    sb.append("    pool: ").append(toIndentedString(pool)).append("\n");
    sb.append("    dhcpServer: ").append(toIndentedString(dhcpServer)).append("\n");
    sb.append("    hasFloatingIps: ").append(toIndentedString(hasFloatingIps)).append("\n");
    sb.append("    dhcpIp: ").append(toIndentedString(dhcpIp)).append("\n");
    sb.append("    dnsPrimary: ").append(toIndentedString(dnsPrimary)).append("\n");
    sb.append("    dnsSecondary: ").append(toIndentedString(dnsSecondary)).append("\n");
    sb.append("    dhcpStart: ").append(toIndentedString(dhcpStart)).append("\n");
    sb.append("    dhcpEnd: ").append(toIndentedString(dhcpEnd)).append("\n");
    sb.append("    dhcpRange: ").append(toIndentedString(dhcpRange)).append("\n");
    sb.append("    networkProxy: ").append(toIndentedString(networkProxy)).append("\n");
    sb.append("    networkDomain: ").append(toIndentedString(networkDomain)).append("\n");
    sb.append("    searchDomains: ").append(toIndentedString(searchDomains)).append("\n");
    sb.append("    defaultNetwork: ").append(toIndentedString(defaultNetwork)).append("\n");
    sb.append("    assignPublicIp: ").append(toIndentedString(assignPublicIp)).append("\n");
    sb.append("    visibility: ").append(toIndentedString(visibility)).append("\n");
    sb.append("    status: ").append(toIndentedString(status)).append("\n");
    sb.append("    network: ").append(toIndentedString(network)).append("\n");
    sb.append("    type: ").append(toIndentedString(type)).append("\n");
    sb.append("    account: ").append(toIndentedString(account)).append("\n");
    sb.append("    securityGroups: ").append(toIndentedString(securityGroups)).append("\n");
    sb.append("    tenants: ").append(toIndentedString(tenants)).append("\n");
    sb.append("    resourcePermission: ").append(toIndentedString(resourcePermission)).append("\n");
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

