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
import OneOfobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobject from './OneOfobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobject';

/**
 * The PolicyCloudUpdatePolicyType model module.
 * @module model/PolicyCloudUpdatePolicyType
 * @version 6.2.1
 */
class PolicyCloudUpdatePolicyType {
    /**
     * Constructs a new <code>PolicyCloudUpdatePolicyType</code>.
     * @alias module:model/PolicyCloudUpdatePolicyType
     */
    constructor() { 
        
        PolicyCloudUpdatePolicyType.initialize(this);
    }

    /**
     * Initializes the fields of this object.
     * This method is used by the constructors of any subclasses, in order to implement multiple inheritance (mix-ins).
     * Only for internal use.
     */
    static initialize(obj) { 
    }

    /**
     * Constructs a <code>PolicyCloudUpdatePolicyType</code> from a plain JavaScript object, optionally creating a new instance.
     * Copies all relevant properties from <code>data</code> to <code>obj</code> if supplied or a new instance if not.
     * @param {Object} data The plain JavaScript object bearing properties of interest.
     * @param {module:model/PolicyCloudUpdatePolicyType} obj Optional instance to populate.
     * @return {module:model/PolicyCloudUpdatePolicyType} The populated <code>PolicyCloudUpdatePolicyType</code> instance.
     */
    static constructFromObject(data, obj) {
        if (data) {
            obj = obj || new PolicyCloudUpdatePolicyType();

            if (data.hasOwnProperty('code')) {
                obj['code'] = ApiClient.convertToType(data['code'], 'String');
            }
            if (data.hasOwnProperty('config')) {
                obj['config'] = ApiClient.convertToType(data['config'], OneOfobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobject);
            }
            if (data.hasOwnProperty('enabled')) {
                obj['enabled'] = ApiClient.convertToType(data['enabled'], 'Boolean');
            }
            if (data.hasOwnProperty('refType')) {
                obj['refType'] = ApiClient.convertToType(data['refType'], 'String');
            }
            if (data.hasOwnProperty('refId')) {
                obj['refId'] = ApiClient.convertToType(data['refId'], 'Number');
            }
            if (data.hasOwnProperty('accounts')) {
                obj['accounts'] = ApiClient.convertToType(data['accounts'], ['Number']);
            }
            if (data.hasOwnProperty('eachUser')) {
                obj['eachUser'] = ApiClient.convertToType(data['eachUser'], 'Boolean');
            }
        }
        return obj;
    }


}

/**
 * The policy type
 * @member {module:model/PolicyCloudUpdatePolicyType.CodeEnum} code
 */
PolicyCloudUpdatePolicyType.prototype['code'] = undefined;

/**
 * A map of config values. The expected values vary by policyType.
 * @member {module:model/OneOfobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobject} config
 */
PolicyCloudUpdatePolicyType.prototype['config'] = undefined;

/**
 * Set to false to disable
 * @member {Boolean} enabled
 * @default true
 */
PolicyCloudUpdatePolicyType.prototype['enabled'] = true;

/**
 * Scope object type
 * @member {module:model/PolicyCloudUpdatePolicyType.RefTypeEnum} refType
 */
PolicyCloudUpdatePolicyType.prototype['refType'] = undefined;

/**
 * Scope object ID (`cloud`)
 * @member {Number} refId
 */
PolicyCloudUpdatePolicyType.prototype['refId'] = undefined;

/**
 * Array of tenants to scope the policy to
 * @member {Array.<Number>} accounts
 */
PolicyCloudUpdatePolicyType.prototype['accounts'] = undefined;

/**
 * Apply individually to each user in role.  Only when `refType` equals `Role`
 * @member {Boolean} eachUser
 */
PolicyCloudUpdatePolicyType.prototype['eachUser'] = undefined;





/**
 * Allowed values for the <code>code</code> property.
 * @enum {String}
 * @readonly
 */
PolicyCloudUpdatePolicyType['CodeEnum'] = {

    /**
     * value: "Approve Delete"
     * @const
     */
    "Approve Delete": "Approve Delete",

    /**
     * value: "Approve Provision"
     * @const
     */
    "Approve Provision": "Approve Provision",

    /**
     * value: "Approve Reconfigure"
     * @const
     */
    "Approve Reconfigure": "Approve Reconfigure",

    /**
     * value: "Backup Creation"
     * @const
     */
    "Backup Creation": "Backup Creation",

    /**
     * value: "Backup Targets"
     * @const
     */
    "Backup Targets": "Backup Targets",

    /**
     * value: "Budget"
     * @const
     */
    "Budget": "Budget",

    /**
     * value: "Cluster Resource Name"
     * @const
     */
    "Cluster Resource Name": "Cluster Resource Name",

    /**
     * value: "Cypher Access"
     * @const
     */
    "Cypher Access": "Cypher Access",

    /**
     * value: "Delayed Delete"
     * @const
     */
    "Delayed Delete": "Delayed Delete",

    /**
     * value: "Expiration"
     * @const
     */
    "Expiration": "Expiration",

    /**
     * value: "File Share Storage Quota"
     * @const
     */
    "File Share Storage Quota": "File Share Storage Quota",

    /**
     * value: "Hostname"
     * @const
     */
    "Hostname": "Hostname",

    /**
     * value: "Instance Name"
     * @const
     */
    "Instance Name": "Instance Name",

    /**
     * value: "Max Containers"
     * @const
     */
    "Max Containers": "Max Containers",

    /**
     * value: "Max Cores"
     * @const
     */
    "Max Cores": "Max Cores",

    /**
     * value: "Max Hosts"
     * @const
     */
    "Max Hosts": "Max Hosts",

    /**
     * value: "Max Load Balancer Pools"
     * @const
     */
    "Max Load Balancer Pools": "Max Load Balancer Pools",

    /**
     * value: "Max Memory"
     * @const
     */
    "Max Memory": "Max Memory",

    /**
     * value: "Max Pool Members"
     * @const
     */
    "Max Pool Members": "Max Pool Members",

    /**
     * value: "Max Storage"
     * @const
     */
    "Max Storage": "Max Storage",

    /**
     * value: "Max Virtual Servers"
     * @const
     */
    "Max Virtual Servers": "Max Virtual Servers",

    /**
     * value: "Max VMs"
     * @const
     */
    "Max VMs": "Max VMs",

    /**
     * value: "Message of the Day"
     * @const
     */
    "Message of the Day": "Message of the Day",

    /**
     * value: "Network Quota"
     * @const
     */
    "Network Quota": "Network Quota",

    /**
     * value: "Object Storage Quota"
     * @const
     */
    "Object Storage Quota": "Object Storage Quota",

    /**
     * value: "Power Schedule"
     * @const
     */
    "Power Schedule": "Power Schedule",

    /**
     * value: "Router Quota"
     * @const
     */
    "Router Quota": "Router Quota",

    /**
     * value: "Shutdown"
     * @const
     */
    "Shutdown": "Shutdown",

    /**
     * value: "Storage Server Storage Quota"
     * @const
     */
    "Storage Server Storage Quota": "Storage Server Storage Quota",

    /**
     * value: "Tags"
     * @const
     */
    "Tags": "Tags",

    /**
     * value: "User Creation"
     * @const
     */
    "User Creation": "User Creation",

    /**
     * value: "User Group Creation"
     * @const
     */
    "User Group Creation": "User Group Creation",

    /**
     * value: "Workflow"
     * @const
     */
    "Workflow": "Workflow"
};


/**
 * Allowed values for the <code>refType</code> property.
 * @enum {String}
 * @readonly
 */
PolicyCloudUpdatePolicyType['RefTypeEnum'] = {

    /**
     * value: "ComputeZone"
     * @const
     */
    "ComputeZone": "ComputeZone"
};



export default PolicyCloudUpdatePolicyType;

