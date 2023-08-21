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
import IdentitySourcesCustomSSOConfigConfig from './IdentitySourcesCustomSSOConfigConfig';
import IdentitySourcesLDAPConfigDefaultAccountRole from './IdentitySourcesLDAPConfigDefaultAccountRole';
import InlineResponse20040AppDeployInstance from './InlineResponse20040AppDeployInstance';

/**
 * The IdentitySourcesCustomSSOConfig model module.
 * @module model/IdentitySourcesCustomSSOConfig
 * @version 6.2.1
 */
class IdentitySourcesCustomSSOConfig {
    /**
     * Constructs a new <code>IdentitySourcesCustomSSOConfig</code>.
     * @alias module:model/IdentitySourcesCustomSSOConfig
     */
    constructor() { 
        
        IdentitySourcesCustomSSOConfig.initialize(this);
    }

    /**
     * Initializes the fields of this object.
     * This method is used by the constructors of any subclasses, in order to implement multiple inheritance (mix-ins).
     * Only for internal use.
     */
    static initialize(obj) { 
    }

    /**
     * Constructs a <code>IdentitySourcesCustomSSOConfig</code> from a plain JavaScript object, optionally creating a new instance.
     * Copies all relevant properties from <code>data</code> to <code>obj</code> if supplied or a new instance if not.
     * @param {Object} data The plain JavaScript object bearing properties of interest.
     * @param {module:model/IdentitySourcesCustomSSOConfig} obj Optional instance to populate.
     * @return {module:model/IdentitySourcesCustomSSOConfig} The populated <code>IdentitySourcesCustomSSOConfig</code> instance.
     */
    static constructFromObject(data, obj) {
        if (data) {
            obj = obj || new IdentitySourcesCustomSSOConfig();

            if (data.hasOwnProperty('id')) {
                obj['id'] = ApiClient.convertToType(data['id'], 'Number');
            }
            if (data.hasOwnProperty('name')) {
                obj['name'] = ApiClient.convertToType(data['name'], 'String');
            }
            if (data.hasOwnProperty('description')) {
                obj['description'] = ApiClient.convertToType(data['description'], 'String');
            }
            if (data.hasOwnProperty('code')) {
                obj['code'] = ApiClient.convertToType(data['code'], 'String');
            }
            if (data.hasOwnProperty('type')) {
                obj['type'] = ApiClient.convertToType(data['type'], 'String');
            }
            if (data.hasOwnProperty('active')) {
                obj['active'] = ApiClient.convertToType(data['active'], 'Boolean');
            }
            if (data.hasOwnProperty('deleted')) {
                obj['deleted'] = ApiClient.convertToType(data['deleted'], 'Boolean');
            }
            if (data.hasOwnProperty('autoSyncOnLogin')) {
                obj['autoSyncOnLogin'] = ApiClient.convertToType(data['autoSyncOnLogin'], 'Boolean');
            }
            if (data.hasOwnProperty('externalLogin')) {
                obj['externalLogin'] = ApiClient.convertToType(data['externalLogin'], 'Boolean');
            }
            if (data.hasOwnProperty('allowCustomMappings')) {
                obj['allowCustomMappings'] = ApiClient.convertToType(data['allowCustomMappings'], 'Boolean');
            }
            if (data.hasOwnProperty('manualRoleAssignment')) {
                obj['manualRoleAssignment'] = ApiClient.convertToType(data['manualRoleAssignment'], 'Boolean');
            }
            if (data.hasOwnProperty('account')) {
                obj['account'] = InlineResponse20040AppDeployInstance.constructFromObject(data['account']);
            }
            if (data.hasOwnProperty('defaultAccountRole')) {
                obj['defaultAccountRole'] = IdentitySourcesLDAPConfigDefaultAccountRole.constructFromObject(data['defaultAccountRole']);
            }
            if (data.hasOwnProperty('config')) {
                obj['config'] = IdentitySourcesCustomSSOConfigConfig.constructFromObject(data['config']);
            }
            if (data.hasOwnProperty('roleMappings')) {
                obj['roleMappings'] = ApiClient.convertToType(data['roleMappings'], [Object]);
            }
            if (data.hasOwnProperty('subdomain')) {
                obj['subdomain'] = ApiClient.convertToType(data['subdomain'], 'String');
            }
            if (data.hasOwnProperty('loginURL')) {
                obj['loginURL'] = ApiClient.convertToType(data['loginURL'], 'String');
            }
            if (data.hasOwnProperty('providerSettings')) {
                obj['providerSettings'] = ApiClient.convertToType(data['providerSettings'], Object);
            }
            if (data.hasOwnProperty('dateCreated')) {
                obj['dateCreated'] = ApiClient.convertToType(data['dateCreated'], 'Date');
            }
            if (data.hasOwnProperty('lastUpdated')) {
                obj['lastUpdated'] = ApiClient.convertToType(data['lastUpdated'], 'Date');
            }
        }
        return obj;
    }


}

/**
 * @member {Number} id
 */
IdentitySourcesCustomSSOConfig.prototype['id'] = undefined;

/**
 * @member {String} name
 */
IdentitySourcesCustomSSOConfig.prototype['name'] = undefined;

/**
 * @member {String} description
 */
IdentitySourcesCustomSSOConfig.prototype['description'] = undefined;

/**
 * @member {String} code
 */
IdentitySourcesCustomSSOConfig.prototype['code'] = undefined;

/**
 * @member {String} type
 */
IdentitySourcesCustomSSOConfig.prototype['type'] = undefined;

/**
 * @member {Boolean} active
 */
IdentitySourcesCustomSSOConfig.prototype['active'] = undefined;

/**
 * @member {Boolean} deleted
 */
IdentitySourcesCustomSSOConfig.prototype['deleted'] = undefined;

/**
 * @member {Boolean} autoSyncOnLogin
 */
IdentitySourcesCustomSSOConfig.prototype['autoSyncOnLogin'] = undefined;

/**
 * @member {Boolean} externalLogin
 */
IdentitySourcesCustomSSOConfig.prototype['externalLogin'] = undefined;

/**
 * @member {Boolean} allowCustomMappings
 */
IdentitySourcesCustomSSOConfig.prototype['allowCustomMappings'] = undefined;

/**
 * @member {Boolean} manualRoleAssignment
 */
IdentitySourcesCustomSSOConfig.prototype['manualRoleAssignment'] = undefined;

/**
 * @member {module:model/InlineResponse20040AppDeployInstance} account
 */
IdentitySourcesCustomSSOConfig.prototype['account'] = undefined;

/**
 * @member {module:model/IdentitySourcesLDAPConfigDefaultAccountRole} defaultAccountRole
 */
IdentitySourcesCustomSSOConfig.prototype['defaultAccountRole'] = undefined;

/**
 * @member {module:model/IdentitySourcesCustomSSOConfigConfig} config
 */
IdentitySourcesCustomSSOConfig.prototype['config'] = undefined;

/**
 * @member {Array.<Object>} roleMappings
 */
IdentitySourcesCustomSSOConfig.prototype['roleMappings'] = undefined;

/**
 * @member {String} subdomain
 */
IdentitySourcesCustomSSOConfig.prototype['subdomain'] = undefined;

/**
 * @member {String} loginURL
 */
IdentitySourcesCustomSSOConfig.prototype['loginURL'] = undefined;

/**
 * @member {Object} providerSettings
 */
IdentitySourcesCustomSSOConfig.prototype['providerSettings'] = undefined;

/**
 * @member {Date} dateCreated
 */
IdentitySourcesCustomSSOConfig.prototype['dateCreated'] = undefined;

/**
 * @member {Date} lastUpdated
 */
IdentitySourcesCustomSSOConfig.prototype['lastUpdated'] = undefined;






export default IdentitySourcesCustomSSOConfig;

