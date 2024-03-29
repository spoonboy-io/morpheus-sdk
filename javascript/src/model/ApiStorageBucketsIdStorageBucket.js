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
import OneOfobjectobjectobjectobjectobjectobjectobject from './OneOfobjectobjectobjectobjectobjectobjectobject';

/**
 * The ApiStorageBucketsIdStorageBucket model module.
 * @module model/ApiStorageBucketsIdStorageBucket
 * @version 6.2.1
 */
class ApiStorageBucketsIdStorageBucket {
    /**
     * Constructs a new <code>ApiStorageBucketsIdStorageBucket</code>.
     * @alias module:model/ApiStorageBucketsIdStorageBucket
     */
    constructor() { 
        
        ApiStorageBucketsIdStorageBucket.initialize(this);
    }

    /**
     * Initializes the fields of this object.
     * This method is used by the constructors of any subclasses, in order to implement multiple inheritance (mix-ins).
     * Only for internal use.
     */
    static initialize(obj) { 
    }

    /**
     * Constructs a <code>ApiStorageBucketsIdStorageBucket</code> from a plain JavaScript object, optionally creating a new instance.
     * Copies all relevant properties from <code>data</code> to <code>obj</code> if supplied or a new instance if not.
     * @param {Object} data The plain JavaScript object bearing properties of interest.
     * @param {module:model/ApiStorageBucketsIdStorageBucket} obj Optional instance to populate.
     * @return {module:model/ApiStorageBucketsIdStorageBucket} The populated <code>ApiStorageBucketsIdStorageBucket</code> instance.
     */
    static constructFromObject(data, obj) {
        if (data) {
            obj = obj || new ApiStorageBucketsIdStorageBucket();

            if (data.hasOwnProperty('name')) {
                obj['name'] = ApiClient.convertToType(data['name'], 'String');
            }
            if (data.hasOwnProperty('providerType')) {
                obj['providerType'] = ApiClient.convertToType(data['providerType'], 'String');
            }
            if (data.hasOwnProperty('defaultBackupTarget')) {
                obj['defaultBackupTarget'] = ApiClient.convertToType(data['defaultBackupTarget'], 'Boolean');
            }
            if (data.hasOwnProperty('copyToStore')) {
                obj['copyToStore'] = ApiClient.convertToType(data['copyToStore'], 'Boolean');
            }
            if (data.hasOwnProperty('defaultDeploymentTarget')) {
                obj['defaultDeploymentTarget'] = ApiClient.convertToType(data['defaultDeploymentTarget'], 'Boolean');
            }
            if (data.hasOwnProperty('defaultVirtualImageTarget')) {
                obj['defaultVirtualImageTarget'] = ApiClient.convertToType(data['defaultVirtualImageTarget'], 'Boolean');
            }
            if (data.hasOwnProperty('retentionPolicyType')) {
                obj['retentionPolicyType'] = ApiClient.convertToType(data['retentionPolicyType'], 'String');
            }
            if (data.hasOwnProperty('retentionPolicyDays')) {
                obj['retentionPolicyDays'] = ApiClient.convertToType(data['retentionPolicyDays'], 'Number');
            }
            if (data.hasOwnProperty('retentionProvider')) {
                obj['retentionProvider'] = ApiClient.convertToType(data['retentionProvider'], 'String');
            }
            if (data.hasOwnProperty('bucketName')) {
                obj['bucketName'] = ApiClient.convertToType(data['bucketName'], 'String');
            }
            if (data.hasOwnProperty('createBucket')) {
                obj['createBucket'] = ApiClient.convertToType(data['createBucket'], 'Boolean');
            }
            if (data.hasOwnProperty('config')) {
                obj['config'] = ApiClient.convertToType(data['config'], OneOfobjectobjectobjectobjectobjectobjectobject);
            }
        }
        return obj;
    }


}

/**
 * A unique name scoped to your account for the storage bucket
 * @member {String} name
 */
ApiStorageBucketsIdStorageBucket.prototype['name'] = undefined;

/**
 * The type of storage bucket
 * @member {module:model/ApiStorageBucketsIdStorageBucket.ProviderTypeEnum} providerType
 */
ApiStorageBucketsIdStorageBucket.prototype['providerType'] = undefined;

/**
 * Default Backup Target
 * @member {Boolean} defaultBackupTarget
 * @default false
 */
ApiStorageBucketsIdStorageBucket.prototype['defaultBackupTarget'] = false;

/**
 * Archive Snapshots
 * @member {Boolean} copyToStore
 */
ApiStorageBucketsIdStorageBucket.prototype['copyToStore'] = undefined;

/**
 * Default Deployment Target
 * @member {Boolean} defaultDeploymentTarget
 * @default false
 */
ApiStorageBucketsIdStorageBucket.prototype['defaultDeploymentTarget'] = false;

/**
 * Default Virtual Image Store
 * @member {Boolean} defaultVirtualImageTarget
 * @default false
 */
ApiStorageBucketsIdStorageBucket.prototype['defaultVirtualImageTarget'] = false;

/**
 * Cleanup mode. `backup` - Move old files to a backup provider. `delete` - Delete old files. `none` - Keep all files.
 * @member {module:model/ApiStorageBucketsIdStorageBucket.RetentionPolicyTypeEnum} retentionPolicyType
 * @default 'none'
 */
ApiStorageBucketsIdStorageBucket.prototype['retentionPolicyType'] = 'none';

/**
 * The number of days old a file must be before it is deleted.
 * @member {Number} retentionPolicyDays
 */
ApiStorageBucketsIdStorageBucket.prototype['retentionPolicyDays'] = undefined;

/**
 * The backup Storage Bucket where old files are moved to.
 * @member {String} retentionProvider
 */
ApiStorageBucketsIdStorageBucket.prototype['retentionProvider'] = undefined;

/**
 * The name of the bucket. Only applies to `Amazon`, `Azure`, `CIFS`, `NFSv3`, `Openstack Swift`, and `Rackspace CDN`.
 * @member {String} bucketName
 */
ApiStorageBucketsIdStorageBucket.prototype['bucketName'] = undefined;

/**
 * Create the bucket if it does not exist. Only applies to `Amazon`, `Azure`, `Openstack Swift`, and `Rackspace CDN`.
 * @member {Boolean} createBucket
 * @default false
 */
ApiStorageBucketsIdStorageBucket.prototype['createBucket'] = false;

/**
 * @member {module:model/OneOfobjectobjectobjectobjectobjectobjectobject} config
 */
ApiStorageBucketsIdStorageBucket.prototype['config'] = undefined;





/**
 * Allowed values for the <code>providerType</code> property.
 * @enum {String}
 * @readonly
 */
ApiStorageBucketsIdStorageBucket['ProviderTypeEnum'] = {

    /**
     * value: "s3"
     * @const
     */
    "s3": "s3",

    /**
     * value: "azure"
     * @const
     */
    "azure": "azure",

    /**
     * value: "cifs"
     * @const
     */
    "cifs": "cifs",

    /**
     * value: "local"
     * @const
     */
    "local": "local",

    /**
     * value: "nfs"
     * @const
     */
    "nfs": "nfs",

    /**
     * value: "openstack"
     * @const
     */
    "openstack": "openstack",

    /**
     * value: "rackspace"
     * @const
     */
    "rackspace": "rackspace"
};


/**
 * Allowed values for the <code>retentionPolicyType</code> property.
 * @enum {String}
 * @readonly
 */
ApiStorageBucketsIdStorageBucket['RetentionPolicyTypeEnum'] = {

    /**
     * value: "backup"
     * @const
     */
    "backup": "backup",

    /**
     * value: "delete"
     * @const
     */
    "delete": "delete",

    /**
     * value: "none"
     * @const
     */
    "none": "none"
};



export default ApiStorageBucketsIdStorageBucket;

