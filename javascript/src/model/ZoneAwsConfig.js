/**
 * Morpheus API
 * Morpheus is a powerful cloud management tool that provides provisioning, monitoring, logging, backups, and application deployment strategies.  This document describes the Morpheus API protocol and the available endpoints. Sections are organized in the same manner as they appear in the Morpheus UI.
 *
 * The version of the OpenAPI document: 6.2.1
 * Contact: dev@morpheusdata.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 *
 */

import ApiClient from '../ApiClient';
import ZoneVcenterConfigNetworkServer from './ZoneVcenterConfigNetworkServer';

/**
 * The ZoneAwsConfig model module.
 * @module model/ZoneAwsConfig
 * @version 6.2.1
 */
class ZoneAwsConfig {
    /**
     * Constructs a new <code>ZoneAwsConfig</code>.
     * @alias module:model/ZoneAwsConfig
     */
    constructor() { 
        
        ZoneAwsConfig.initialize(this);
    }

    /**
     * Initializes the fields of this object.
     * This method is used by the constructors of any subclasses, in order to implement multiple inheritance (mix-ins).
     * Only for internal use.
     */
    static initialize(obj) { 
    }

    /**
     * Constructs a <code>ZoneAwsConfig</code> from a plain JavaScript object, optionally creating a new instance.
     * Copies all relevant properties from <code>data</code> to <code>obj</code> if supplied or a new instance if not.
     * @param {Object} data The plain JavaScript object bearing properties of interest.
     * @param {module:model/ZoneAwsConfig} obj Optional instance to populate.
     * @return {module:model/ZoneAwsConfig} The populated <code>ZoneAwsConfig</code> instance.
     */
    static constructFromObject(data, obj) {
        if (data) {
            obj = obj || new ZoneAwsConfig();

            if (data.hasOwnProperty('endpoint')) {
                obj['endpoint'] = ApiClient.convertToType(data['endpoint'], 'String');
            }
            if (data.hasOwnProperty('accessKey')) {
                obj['accessKey'] = ApiClient.convertToType(data['accessKey'], 'String');
            }
            if (data.hasOwnProperty('secretKey')) {
                obj['secretKey'] = ApiClient.convertToType(data['secretKey'], 'String');
            }
            if (data.hasOwnProperty('_useHostCredentials')) {
                obj['_useHostCredentials'] = ApiClient.convertToType(data['_useHostCredentials'], 'String');
            }
            if (data.hasOwnProperty('stsAssumeRole')) {
                obj['stsAssumeRole'] = ApiClient.convertToType(data['stsAssumeRole'], 'String');
            }
            if (data.hasOwnProperty('isVpc')) {
                obj['isVpc'] = ApiClient.convertToType(data['isVpc'], 'String');
            }
            if (data.hasOwnProperty('vpc')) {
                obj['vpc'] = ApiClient.convertToType(data['vpc'], 'String');
            }
            if (data.hasOwnProperty('imageStoreId')) {
                obj['imageStoreId'] = ApiClient.convertToType(data['imageStoreId'], 'String');
            }
            if (data.hasOwnProperty('ebsEncryption')) {
                obj['ebsEncryption'] = ApiClient.convertToType(data['ebsEncryption'], 'String');
            }
            if (data.hasOwnProperty('costingReport')) {
                obj['costingReport'] = ApiClient.convertToType(data['costingReport'], 'String');
            }
            if (data.hasOwnProperty('costingFolder')) {
                obj['costingFolder'] = ApiClient.convertToType(data['costingFolder'], 'String');
            }
            if (data.hasOwnProperty('costingBucket')) {
                obj['costingBucket'] = ApiClient.convertToType(data['costingBucket'], 'String');
            }
            if (data.hasOwnProperty('costingBucketName')) {
                obj['costingBucketName'] = ApiClient.convertToType(data['costingBucketName'], 'String');
            }
            if (data.hasOwnProperty('costingRegion')) {
                obj['costingRegion'] = ApiClient.convertToType(data['costingRegion'], 'String');
            }
            if (data.hasOwnProperty('costingAccessKey')) {
                obj['costingAccessKey'] = ApiClient.convertToType(data['costingAccessKey'], 'String');
            }
            if (data.hasOwnProperty('costingSecretKey')) {
                obj['costingSecretKey'] = ApiClient.convertToType(data['costingSecretKey'], 'String');
            }
            if (data.hasOwnProperty('costingReportName')) {
                obj['costingReportName'] = ApiClient.convertToType(data['costingReportName'], 'String');
            }
            if (data.hasOwnProperty('applianceUrl')) {
                obj['applianceUrl'] = ApiClient.convertToType(data['applianceUrl'], 'String');
            }
            if (data.hasOwnProperty('datacenterName')) {
                obj['datacenterName'] = ApiClient.convertToType(data['datacenterName'], 'String');
            }
            if (data.hasOwnProperty('networkServer.id')) {
                obj['networkServer.id'] = ApiClient.convertToType(data['networkServer.id'], 'String');
            }
            if (data.hasOwnProperty('networkServer')) {
                obj['networkServer'] = ZoneVcenterConfigNetworkServer.constructFromObject(data['networkServer']);
            }
            if (data.hasOwnProperty('securityServer')) {
                obj['securityServer'] = ApiClient.convertToType(data['securityServer'], 'String');
            }
            if (data.hasOwnProperty('certificateProvider')) {
                obj['certificateProvider'] = ApiClient.convertToType(data['certificateProvider'], 'String');
            }
            if (data.hasOwnProperty('backupMode')) {
                obj['backupMode'] = ApiClient.convertToType(data['backupMode'], 'String');
            }
            if (data.hasOwnProperty('replicationMode')) {
                obj['replicationMode'] = ApiClient.convertToType(data['replicationMode'], 'String');
            }
            if (data.hasOwnProperty('dnsIntegrationId')) {
                obj['dnsIntegrationId'] = ApiClient.convertToType(data['dnsIntegrationId'], 'String');
            }
            if (data.hasOwnProperty('serviceRegistryId')) {
                obj['serviceRegistryId'] = ApiClient.convertToType(data['serviceRegistryId'], 'String');
            }
            if (data.hasOwnProperty('configManagementId')) {
                obj['configManagementId'] = ApiClient.convertToType(data['configManagementId'], 'String');
            }
            if (data.hasOwnProperty('configCmdbDiscovery')) {
                obj['configCmdbDiscovery'] = ApiClient.convertToType(data['configCmdbDiscovery'], 'Boolean');
            }
            if (data.hasOwnProperty('secretKeyHash')) {
                obj['secretKeyHash'] = ApiClient.convertToType(data['secretKeyHash'], 'String');
            }
            if (data.hasOwnProperty('costingSecretKeyHash')) {
                obj['costingSecretKeyHash'] = ApiClient.convertToType(data['costingSecretKeyHash'], 'String');
            }
        }
        return obj;
    }


}

/**
 * @member {String} endpoint
 */
ZoneAwsConfig.prototype['endpoint'] = undefined;

/**
 * @member {String} accessKey
 */
ZoneAwsConfig.prototype['accessKey'] = undefined;

/**
 * @member {String} secretKey
 */
ZoneAwsConfig.prototype['secretKey'] = undefined;

/**
 * @member {String} _useHostCredentials
 */
ZoneAwsConfig.prototype['_useHostCredentials'] = undefined;

/**
 * @member {String} stsAssumeRole
 */
ZoneAwsConfig.prototype['stsAssumeRole'] = undefined;

/**
 * @member {String} isVpc
 */
ZoneAwsConfig.prototype['isVpc'] = undefined;

/**
 * @member {String} vpc
 */
ZoneAwsConfig.prototype['vpc'] = undefined;

/**
 * @member {String} imageStoreId
 */
ZoneAwsConfig.prototype['imageStoreId'] = undefined;

/**
 * @member {String} ebsEncryption
 */
ZoneAwsConfig.prototype['ebsEncryption'] = undefined;

/**
 * @member {String} costingReport
 */
ZoneAwsConfig.prototype['costingReport'] = undefined;

/**
 * @member {String} costingFolder
 */
ZoneAwsConfig.prototype['costingFolder'] = undefined;

/**
 * @member {String} costingBucket
 */
ZoneAwsConfig.prototype['costingBucket'] = undefined;

/**
 * @member {String} costingBucketName
 */
ZoneAwsConfig.prototype['costingBucketName'] = undefined;

/**
 * @member {String} costingRegion
 */
ZoneAwsConfig.prototype['costingRegion'] = undefined;

/**
 * @member {String} costingAccessKey
 */
ZoneAwsConfig.prototype['costingAccessKey'] = undefined;

/**
 * @member {String} costingSecretKey
 */
ZoneAwsConfig.prototype['costingSecretKey'] = undefined;

/**
 * @member {String} costingReportName
 */
ZoneAwsConfig.prototype['costingReportName'] = undefined;

/**
 * @member {String} applianceUrl
 */
ZoneAwsConfig.prototype['applianceUrl'] = undefined;

/**
 * @member {String} datacenterName
 */
ZoneAwsConfig.prototype['datacenterName'] = undefined;

/**
 * @member {String} networkServer.id
 */
ZoneAwsConfig.prototype['networkServer.id'] = undefined;

/**
 * @member {module:model/ZoneVcenterConfigNetworkServer} networkServer
 */
ZoneAwsConfig.prototype['networkServer'] = undefined;

/**
 * @member {String} securityServer
 */
ZoneAwsConfig.prototype['securityServer'] = undefined;

/**
 * @member {String} certificateProvider
 */
ZoneAwsConfig.prototype['certificateProvider'] = undefined;

/**
 * @member {String} backupMode
 */
ZoneAwsConfig.prototype['backupMode'] = undefined;

/**
 * @member {String} replicationMode
 */
ZoneAwsConfig.prototype['replicationMode'] = undefined;

/**
 * @member {String} dnsIntegrationId
 */
ZoneAwsConfig.prototype['dnsIntegrationId'] = undefined;

/**
 * @member {String} serviceRegistryId
 */
ZoneAwsConfig.prototype['serviceRegistryId'] = undefined;

/**
 * @member {String} configManagementId
 */
ZoneAwsConfig.prototype['configManagementId'] = undefined;

/**
 * @member {Boolean} configCmdbDiscovery
 */
ZoneAwsConfig.prototype['configCmdbDiscovery'] = undefined;

/**
 * @member {String} secretKeyHash
 */
ZoneAwsConfig.prototype['secretKeyHash'] = undefined;

/**
 * @member {String} costingSecretKeyHash
 */
ZoneAwsConfig.prototype['costingSecretKeyHash'] = undefined;






export default ZoneAwsConfig;

