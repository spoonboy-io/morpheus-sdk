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

/**
 * The PolicyCreatePolicyType model module.
 * @module model/PolicyCreatePolicyType
 * @version 6.2.1
 */
class PolicyCreatePolicyType {
    /**
     * Constructs a new <code>PolicyCreatePolicyType</code>.
     * @alias module:model/PolicyCreatePolicyType
     */
    constructor() { 
        
        PolicyCreatePolicyType.initialize(this);
    }

    /**
     * Initializes the fields of this object.
     * This method is used by the constructors of any subclasses, in order to implement multiple inheritance (mix-ins).
     * Only for internal use.
     */
    static initialize(obj) { 
    }

    /**
     * Constructs a <code>PolicyCreatePolicyType</code> from a plain JavaScript object, optionally creating a new instance.
     * Copies all relevant properties from <code>data</code> to <code>obj</code> if supplied or a new instance if not.
     * @param {Object} data The plain JavaScript object bearing properties of interest.
     * @param {module:model/PolicyCreatePolicyType} obj Optional instance to populate.
     * @return {module:model/PolicyCreatePolicyType} The populated <code>PolicyCreatePolicyType</code> instance.
     */
    static constructFromObject(data, obj) {
        if (data) {
            obj = obj || new PolicyCreatePolicyType();

            if (data.hasOwnProperty('code')) {
                obj['code'] = ApiClient.convertToType(data['code'], 'String');
            }
        }
        return obj;
    }


}

/**
 * The policy type code. See `Retrieves all Policy Types` endpoint for listing.
 * @member {module:model/PolicyCreatePolicyType.CodeEnum} code
 */
PolicyCreatePolicyType.prototype['code'] = undefined;





/**
 * Allowed values for the <code>code</code> property.
 * @enum {String}
 * @readonly
 */
PolicyCreatePolicyType['CodeEnum'] = {

    /**
     * value: "deleteApproval"
     * @const
     */
    "deleteApproval": "deleteApproval",

    /**
     * value: "provisionApproval"
     * @const
     */
    "provisionApproval": "provisionApproval",

    /**
     * value: "reconfigureApproval"
     * @const
     */
    "reconfigureApproval": "reconfigureApproval",

    /**
     * value: "workflowApproval"
     * @const
     */
    "workflowApproval": "workflowApproval",

    /**
     * value: "createBackup"
     * @const
     */
    "createBackup": "createBackup",

    /**
     * value: "backupStorage"
     * @const
     */
    "backupStorage": "backupStorage",

    /**
     * value: "maxPrice"
     * @const
     */
    "maxPrice": "maxPrice",

    /**
     * value: "serverNaming"
     * @const
     */
    "serverNaming": "serverNaming",

    /**
     * value: "cypher"
     * @const
     */
    "cypher": "cypher",

    /**
     * value: "delayedRemoval"
     * @const
     */
    "delayedRemoval": "delayedRemoval",

    /**
     * value: "lifecycle"
     * @const
     */
    "lifecycle": "lifecycle",

    /**
     * value: "storageShareQuota"
     * @const
     */
    "storageShareQuota": "storageShareQuota",

    /**
     * value: "hostNaming"
     * @const
     */
    "hostNaming": "hostNaming",

    /**
     * value: "naming"
     * @const
     */
    "naming": "naming",

    /**
     * value: "maxContainers"
     * @const
     */
    "maxContainers": "maxContainers",

    /**
     * value: "maxCores"
     * @const
     */
    "maxCores": "maxCores",

    /**
     * value: "maxHosts"
     * @const
     */
    "maxHosts": "maxHosts",

    /**
     * value: "maxPools"
     * @const
     */
    "maxPools": "maxPools",

    /**
     * value: "maxMemory"
     * @const
     */
    "maxMemory": "maxMemory",

    /**
     * value: "maxPoolMembers"
     * @const
     */
    "maxPoolMembers": "maxPoolMembers",

    /**
     * value: "maxSnapshots"
     * @const
     */
    "maxSnapshots": "maxSnapshots",

    /**
     * value: "maxStorage"
     * @const
     */
    "maxStorage": "maxStorage",

    /**
     * value: "maxVirtualServers"
     * @const
     */
    "maxVirtualServers": "maxVirtualServers",

    /**
     * value: "maxVms"
     * @const
     */
    "maxVms": "maxVms",

    /**
     * value: "motd"
     * @const
     */
    "motd": "motd",

    /**
     * value: "maxNetworks"
     * @const
     */
    "maxNetworks": "maxNetworks",

    /**
     * value: "storageBucketQuota"
     * @const
     */
    "storageBucketQuota": "storageBucketQuota",

    /**
     * value: "powerSchedule"
     * @const
     */
    "powerSchedule": "powerSchedule",

    /**
     * value: "maxRouters"
     * @const
     */
    "maxRouters": "maxRouters",

    /**
     * value: "shutdown"
     * @const
     */
    "shutdown": "shutdown",

    /**
     * value: "storageServerQuota"
     * @const
     */
    "storageServerQuota": "storageServerQuota",

    /**
     * value: "tags"
     * @const
     */
    "tags": "tags",

    /**
     * value: "createUser"
     * @const
     */
    "createUser": "createUser",

    /**
     * value: "createUserGroup"
     * @const
     */
    "createUserGroup": "createUserGroup",

    /**
     * value: "workflow"
     * @const
     */
    "workflow": "workflow"
};



export default PolicyCreatePolicyType;

