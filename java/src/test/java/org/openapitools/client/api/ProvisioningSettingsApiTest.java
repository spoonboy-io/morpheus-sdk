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
import org.openapitools.client.model.DefaultError;
import org.openapitools.client.model.InlineObject204;
import org.openapitools.client.model.InlineResponse200128;
import org.openapitools.client.model.Model200Success;
import org.junit.Test;
import org.junit.Ignore;

import java.util.ArrayList;
import java.util.HashMap;
import java.util.List;
import java.util.Map;

/**
 * API tests for ProvisioningSettingsApi
 */
@Ignore
public class ProvisioningSettingsApiTest {

    private final ProvisioningSettingsApi api = new ProvisioningSettingsApi();

    
    /**
     * Get Provisioning Settings
     *
     * This endpoint retrieves provisioning settings.
     *
     * @throws ApiException
     *          if the Api call fails
     */
    @Test
    public void listProvisioningSettingsTest() throws ApiException {
        InlineResponse200128 response = api.listProvisioningSettings();

        // TODO: test validations
    }
    
    /**
     * Update Provisioning Settings
     *
     * Update Provisioning Settings
     *
     * @throws ApiException
     *          if the Api call fails
     */
    @Test
    public void updateProvisioningSettingsTest() throws ApiException {
        InlineObject204 inlineObject204 = null;
        Model200Success response = api.updateProvisioningSettings(inlineObject204);

        // TODO: test validations
    }
    
}