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
import org.junit.Assert;
import org.junit.Ignore;
import org.junit.Test;


/**
 * Model tests for Subnet
 */
public class SubnetTest {
    private final Subnet model = new Subnet();

    /**
     * Model tests for Subnet
     */
    @Test
    public void testSubnet() {
        // TODO: test Subnet
    }

    /**
     * Test the property 'id'
     */
    @Test
    public void idTest() {
        // TODO: test id
    }

    /**
     * Test the property 'code'
     */
    @Test
    public void codeTest() {
        // TODO: test code
    }

    /**
     * Test the property 'name'
     */
    @Test
    public void nameTest() {
        // TODO: test name
    }

    /**
     * Test the property 'labels'
     */
    @Test
    public void labelsTest() {
        // TODO: test labels
    }

    /**
     * Test the property 'active'
     */
    @Test
    public void activeTest() {
        // TODO: test active
    }

    /**
     * Test the property 'description'
     */
    @Test
    public void descriptionTest() {
        // TODO: test description
    }

    /**
     * Test the property 'externalId'
     */
    @Test
    public void externalIdTest() {
        // TODO: test externalId
    }

    /**
     * Test the property 'uniqueId'
     */
    @Test
    public void uniqueIdTest() {
        // TODO: test uniqueId
    }

    /**
     * Test the property 'addressPrefix'
     */
    @Test
    public void addressPrefixTest() {
        // TODO: test addressPrefix
    }

    /**
     * Test the property 'cidr'
     */
    @Test
    public void cidrTest() {
        // TODO: test cidr
    }

    /**
     * Test the property 'gateway'
     */
    @Test
    public void gatewayTest() {
        // TODO: test gateway
    }

    /**
     * Test the property 'netmask'
     */
    @Test
    public void netmaskTest() {
        // TODO: test netmask
    }

    /**
     * Test the property 'subnetAddress'
     */
    @Test
    public void subnetAddressTest() {
        // TODO: test subnetAddress
    }

    /**
     * Test the property 'tftpServer'
     */
    @Test
    public void tftpServerTest() {
        // TODO: test tftpServer
    }

    /**
     * Test the property 'bootFile'
     */
    @Test
    public void bootFileTest() {
        // TODO: test bootFile
    }

    /**
     * Test the property 'pool'
     */
    @Test
    public void poolTest() {
        // TODO: test pool
    }

    /**
     * Test the property 'dhcpServer'
     */
    @Test
    public void dhcpServerTest() {
        // TODO: test dhcpServer
    }

    /**
     * Test the property 'hasFloatingIps'
     */
    @Test
    public void hasFloatingIpsTest() {
        // TODO: test hasFloatingIps
    }

    /**
     * Test the property 'dhcpIp'
     */
    @Test
    public void dhcpIpTest() {
        // TODO: test dhcpIp
    }

    /**
     * Test the property 'dnsPrimary'
     */
    @Test
    public void dnsPrimaryTest() {
        // TODO: test dnsPrimary
    }

    /**
     * Test the property 'dnsSecondary'
     */
    @Test
    public void dnsSecondaryTest() {
        // TODO: test dnsSecondary
    }

    /**
     * Test the property 'dhcpStart'
     */
    @Test
    public void dhcpStartTest() {
        // TODO: test dhcpStart
    }

    /**
     * Test the property 'dhcpEnd'
     */
    @Test
    public void dhcpEndTest() {
        // TODO: test dhcpEnd
    }

    /**
     * Test the property 'dhcpRange'
     */
    @Test
    public void dhcpRangeTest() {
        // TODO: test dhcpRange
    }

    /**
     * Test the property 'networkProxy'
     */
    @Test
    public void networkProxyTest() {
        // TODO: test networkProxy
    }

    /**
     * Test the property 'networkDomain'
     */
    @Test
    public void networkDomainTest() {
        // TODO: test networkDomain
    }

    /**
     * Test the property 'searchDomains'
     */
    @Test
    public void searchDomainsTest() {
        // TODO: test searchDomains
    }

    /**
     * Test the property 'defaultNetwork'
     */
    @Test
    public void defaultNetworkTest() {
        // TODO: test defaultNetwork
    }

    /**
     * Test the property 'assignPublicIp'
     */
    @Test
    public void assignPublicIpTest() {
        // TODO: test assignPublicIp
    }

    /**
     * Test the property 'visibility'
     */
    @Test
    public void visibilityTest() {
        // TODO: test visibility
    }

    /**
     * Test the property 'status'
     */
    @Test
    public void statusTest() {
        // TODO: test status
    }

    /**
     * Test the property 'network'
     */
    @Test
    public void networkTest() {
        // TODO: test network
    }

    /**
     * Test the property 'type'
     */
    @Test
    public void typeTest() {
        // TODO: test type
    }

    /**
     * Test the property 'account'
     */
    @Test
    public void accountTest() {
        // TODO: test account
    }

    /**
     * Test the property 'securityGroups'
     */
    @Test
    public void securityGroupsTest() {
        // TODO: test securityGroups
    }

    /**
     * Test the property 'tenants'
     */
    @Test
    public void tenantsTest() {
        // TODO: test tenants
    }

    /**
     * Test the property 'resourcePermission'
     */
    @Test
    public void resourcePermissionTest() {
        // TODO: test resourcePermission
    }

}
