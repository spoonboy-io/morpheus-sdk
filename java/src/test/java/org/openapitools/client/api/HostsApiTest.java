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


package org.openapitools.client.api;

import org.openapitools.client.ApiException;
import java.math.BigDecimal;
import org.openapitools.client.model.DefaultError;
import org.openapitools.client.model.InlineObject220;
import org.openapitools.client.model.InlineObject221;
import org.openapitools.client.model.InlineObject222;
import org.openapitools.client.model.InlineObject223;
import org.openapitools.client.model.InlineObject224;
import org.openapitools.client.model.InlineObject225;
import org.openapitools.client.model.InlineObject226;
import org.openapitools.client.model.InlineObject271;
import org.openapitools.client.model.InlineResponse200137;
import org.openapitools.client.model.InlineResponse200138;
import org.openapitools.client.model.InlineResponse200141;
import org.openapitools.client.model.InlineResponse200168;
import org.openapitools.client.model.InlineResponse20050;
import org.openapitools.client.model.Model200Success;
import org.openapitools.client.model.NetworkInterfaceUpdate;
import org.threeten.bp.OffsetDateTime;
import org.junit.Test;
import org.junit.Ignore;

import java.util.ArrayList;
import java.util.HashMap;
import java.util.List;
import java.util.Map;

/**
 * API tests for HostsApi
 */
@Ignore
public class HostsApiTest {

    private final HostsApi api = new HostsApi();

    
    /**
     * Get a Specific Host
     *
     * This endpoint retrieves a specific host.
     *
     * @throws ApiException
     *          if the Api call fails
     */
    @Test
    public void getHostTest() throws ApiException {
        Long id = null;
        InlineResponse200137 response = api.getHost(id);

        // TODO: test validations
    }
    
    /**
     * Get list of snapshots for a Host
     *
     * Get list of snapshots for a Host
     *
     * @throws ApiException
     *          if the Api call fails
     */
    @Test
    public void getHostSnpshotsTest() throws ApiException {
        Long id = null;
        InlineResponse200138 response = api.getHostSnpshots(id);

        // TODO: test validations
    }
    
    /**
     * Get a Specific Host Type
     *
     * This endpoint will retrieve a specific host type by id
     *
     * @throws ApiException
     *          if the Api call fails
     */
    @Test
    public void getHostTypeTest() throws ApiException {
        Long id = null;
        InlineResponse20050 response = api.getHostType(id);

        // TODO: test validations
    }
    
    /**
     * Retrieves a Server Wiki Page
     *
     * This endpoint retrieves a server Wiki page. 
     *
     * @throws ApiException
     *          if the Api call fails
     */
    @Test
    public void getWikiServerTest() throws ApiException {
        Long id = null;
        InlineResponse200168 response = api.getWikiServer(id);

        // TODO: test validations
    }
    
    /**
     * Host Types
     *
     * Fetch a paginated list of available host types. This returns the configuration options for each type.
     *
     * @throws ApiException
     *          if the Api call fails
     */
    @Test
    public void listHostTypesTest() throws ApiException {
        Long max = null;
        Long offset = null;
        String sort = null;
        String direction = null;
        String name = null;
        String code = null;
        String phrase = null;
        String provisionType = null;
        String zoneType = null;
        Boolean creatable = null;
        Object response = api.listHostTypes(max, offset, sort, direction, name, code, phrase, provisionType, zoneType, creatable);

        // TODO: test validations
    }
    
    /**
     * Get All Hosts
     *
     * This endpoint retrieves a paginated list of hosts.
     *
     * @throws ApiException
     *          if the Api call fails
     */
    @Test
    public void listHostsTest() throws ApiException {
        String name = null;
        String phrase = null;
        Long zoneId = null;
        Long siteId = null;
        Long clusterId = null;
        Boolean managed = null;
        String serverType = null;
        String powerState = null;
        String ip = null;
        Boolean vm = null;
        Boolean vmHypervisor = null;
        Boolean bareMetalHost = null;
        String status = null;
        Boolean agentInstalled = null;
        Long max = null;
        Long offset = null;
        OffsetDateTime lastUpdated = null;
        Long createdBy = null;
        String labels = null;
        String allLabels = null;
        String tags = null;
        String metadata = null;
        String uuid = null;
        String externalId = null;
        String internalId = null;
        String externalUniquelId = null;
        Object response = api.listHosts(name, phrase, zoneId, siteId, clusterId, managed, serverType, powerState, ip, vm, vmHypervisor, bareMetalHost, status, agentInstalled, max, offset, lastUpdated, createdBy, labels, allLabels, tags, metadata, uuid, externalId, internalId, externalUniquelId);

        // TODO: test validations
    }
    
    /**
     * Get Available Service Plans for a Host
     *
     * This endpoint retrieves all the Service Plans available for the specified cloud and host type. It may be used to get the list of available plans when creating a new host or resizing an existing host.
     *
     * @throws ApiException
     *          if the Api call fails
     */
    @Test
    public void listServerServicePlansTest() throws ApiException {
        Long zoneId = null;
        Long serverTypeId = null;
        Long siteId = null;
        InlineResponse200141 response = api.listServerServicePlans(zoneId, serverTypeId, siteId);

        // TODO: test validations
    }
    
    /**
     * Delete a Host
     *
     * Will delete a host asynchronously.
     *
     * @throws ApiException
     *          if the Api call fails
     */
    @Test
    public void removeHostTest() throws ApiException {
        Long id = null;
        String removeResources = null;
        String removeInstances = null;
        String preserveVolumes = null;
        String releaseFloatingIps = null;
        String releaseEIPs = null;
        String force = null;
        Model200Success response = api.removeHost(id, removeResources, removeInstances, preserveVolumes, releaseFloatingIps, releaseEIPs, force);

        // TODO: test validations
    }
    
    /**
     * Restart a Host
     *
     * This will restart a host.
     *
     * @throws ApiException
     *          if the Api call fails
     */
    @Test
    public void restartHostTest() throws ApiException {
        Long id = null;
        Object response = api.restartHost(id);

        // TODO: test validations
    }
    
    /**
     * Start a Host
     *
     * This will start a host.
     *
     * @throws ApiException
     *          if the Api call fails
     */
    @Test
    public void startHostTest() throws ApiException {
        Long id = null;
        Model200Success response = api.startHost(id);

        // TODO: test validations
    }
    
    /**
     * Stop a Host
     *
     * This will stop a host.
     *
     * @throws ApiException
     *          if the Api call fails
     */
    @Test
    public void stopHostTest() throws ApiException {
        Long id = null;
        Model200Success response = api.stopHost(id);

        // TODO: test validations
    }
    
    /**
     * Updating a Host
     *
     * Updating a Host
     *
     * @throws ApiException
     *          if the Api call fails
     */
    @Test
    public void updateHostTest() throws ApiException {
        Long id = null;
        InlineObject220 inlineObject220 = null;
        InlineResponse200137 response = api.updateHost(id, inlineObject220);

        // TODO: test validations
    }
    
    /**
     * Assign To Tenant
     *
     * This will change the ownership of the host to the specified Tenant account. This is only available to Master Tenant users.
     *
     * @throws ApiException
     *          if the Api call fails
     */
    @Test
    public void updateHostAssignTenantTest() throws ApiException {
        Long id = null;
        Long accountId = null;
        InlineObject221 inlineObject221 = null;
        Object response = api.updateHostAssignTenant(id, accountId, inlineObject221);

        // TODO: test validations
    }
    
    /**
     * Change Server Cloud
     *
     * This api call is reserved for migrating servers from one cloud to another. This could be due to moving clusters or resource pool scoping of a server without losing the data.
     *
     * @throws ApiException
     *          if the Api call fails
     */
    @Test
    public void updateHostCloudTest() throws ApiException {
        InlineObject226 inlineObject226 = null;
        Object response = api.updateHostCloud(inlineObject226);

        // TODO: test validations
    }
    
    /**
     * Run Workflow on a Host
     *
     * This will run a provisioning workflow on a host.  For operational workflows, see Execute a Workflow. 
     *
     * @throws ApiException
     *          if the Api call fails
     */
    @Test
    public void updateHostExecuteWorkflowTest() throws ApiException {
        Long id = null;
        Long workflowId = null;
        String workflowName = null;
        InlineObject225 inlineObject225 = null;
        Model200Success response = api.updateHostExecuteWorkflow(id, workflowId, workflowName, inlineObject225);

        // TODO: test validations
    }
    
    /**
     * Install Agent
     *
     * This will make the host a managed server, and install the agent.
     *
     * @throws ApiException
     *          if the Api call fails
     */
    @Test
    public void updateHostInstallAgentTest() throws ApiException {
        Long id = null;
        InlineObject222 inlineObject222 = null;
        Object response = api.updateHostInstallAgent(id, inlineObject222);

        // TODO: test validations
    }
    
    /**
     * Convert To Managed
     *
     * This will make the host a managed server, and install the agent.
     *
     * @throws ApiException
     *          if the Api call fails
     */
    @Test
    public void updateHostManagedTest() throws ApiException {
        Long id = null;
        InlineObject223 inlineObject223 = null;
        Object response = api.updateHostManaged(id, inlineObject223);

        // TODO: test validations
    }
    
    /**
     * Resize a Host
     *
     * Will resize a host asynchronously. This endpoint also allows for NIC reconfiguration by passing a new array of &#x60;networkInterfaces&#x60;.
     *
     * @throws ApiException
     *          if the Api call fails
     */
    @Test
    public void updateHostResizeTest() throws ApiException {
        Long id = null;
        InlineObject224 inlineObject224 = null;
        Object response = api.updateHostResize(id, inlineObject224);

        // TODO: test validations
    }
    
    /**
     * Upgrade Agent
     *
     * This will upgrade the version of the agent installed on the host.
     *
     * @throws ApiException
     *          if the Api call fails
     */
    @Test
    public void updateHostUpgradeAgentTest() throws ApiException {
        Long id = null;
        Model200Success response = api.updateHostUpgradeAgent(id);

        // TODO: test validations
    }
    
    /**
     * Updating a label for a Server&#39;s Network
     *
     * Updating a Server&#39;s Network&#39;s Label
     *
     * @throws ApiException
     *          if the Api call fails
     */
    @Test
    public void updateServerNetworkInterfaceTest() throws ApiException {
        Long id = null;
        BigDecimal networkInterfaceId = null;
        NetworkInterfaceUpdate networkInterfaceUpdate = null;
        Object response = api.updateServerNetworkInterface(id, networkInterfaceId, networkInterfaceUpdate);

        // TODO: test validations
    }
    
    /**
     * Update a Server Wiki Page
     *
     * Updates a server Wiki page. 
     *
     * @throws ApiException
     *          if the Api call fails
     */
    @Test
    public void updateWikiServerTest() throws ApiException {
        Long id = null;
        InlineObject271 inlineObject271 = null;
        Object response = api.updateWikiServer(id, inlineObject271);

        // TODO: test validations
    }
    
}
