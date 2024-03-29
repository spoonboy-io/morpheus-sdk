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
import InlineResponse20094Network from './InlineResponse20094Network';
import JobCreatedBy from './JobCreatedBy';
import JobCustomOptions from './JobCustomOptions';
import JobSecurityPackage from './JobSecurityPackage';
import JobTargets from './JobTargets';
import JobTask from './JobTask';
import JobWorkflow from './JobWorkflow';
import OneOfstringlong from './OneOfstringlong';

/**
 * The Job model module.
 * @module model/Job
 * @version 6.2.1
 */
class Job {
    /**
     * Constructs a new <code>Job</code>.
     * @alias module:model/Job
     */
    constructor() { 
        
        Job.initialize(this);
    }

    /**
     * Initializes the fields of this object.
     * This method is used by the constructors of any subclasses, in order to implement multiple inheritance (mix-ins).
     * Only for internal use.
     */
    static initialize(obj) { 
    }

    /**
     * Constructs a <code>Job</code> from a plain JavaScript object, optionally creating a new instance.
     * Copies all relevant properties from <code>data</code> to <code>obj</code> if supplied or a new instance if not.
     * @param {Object} data The plain JavaScript object bearing properties of interest.
     * @param {module:model/Job} obj Optional instance to populate.
     * @return {module:model/Job} The populated <code>Job</code> instance.
     */
    static constructFromObject(data, obj) {
        if (data) {
            obj = obj || new Job();

            if (data.hasOwnProperty('id')) {
                obj['id'] = ApiClient.convertToType(data['id'], 'Number');
            }
            if (data.hasOwnProperty('name')) {
                obj['name'] = ApiClient.convertToType(data['name'], 'String');
            }
            if (data.hasOwnProperty('labels')) {
                obj['labels'] = ApiClient.convertToType(data['labels'], ['String']);
            }
            if (data.hasOwnProperty('type')) {
                obj['type'] = InlineResponse20094Network.constructFromObject(data['type']);
            }
            if (data.hasOwnProperty('workflow')) {
                obj['workflow'] = JobWorkflow.constructFromObject(data['workflow']);
            }
            if (data.hasOwnProperty('task')) {
                obj['task'] = JobTask.constructFromObject(data['task']);
            }
            if (data.hasOwnProperty('securityPackage')) {
                obj['securityPackage'] = JobSecurityPackage.constructFromObject(data['securityPackage']);
            }
            if (data.hasOwnProperty('jobSummary')) {
                obj['jobSummary'] = ApiClient.convertToType(data['jobSummary'], 'String');
            }
            if (data.hasOwnProperty('scheduleMode')) {
                obj['scheduleMode'] = ApiClient.convertToType(data['scheduleMode'], OneOfstringlong);
            }
            if (data.hasOwnProperty('dateTime')) {
                obj['dateTime'] = ApiClient.convertToType(data['dateTime'], 'String');
            }
            if (data.hasOwnProperty('status')) {
                obj['status'] = ApiClient.convertToType(data['status'], 'String');
            }
            if (data.hasOwnProperty('namespace')) {
                obj['namespace'] = ApiClient.convertToType(data['namespace'], 'String');
            }
            if (data.hasOwnProperty('category')) {
                obj['category'] = ApiClient.convertToType(data['category'], 'String');
            }
            if (data.hasOwnProperty('description')) {
                obj['description'] = ApiClient.convertToType(data['description'], 'String');
            }
            if (data.hasOwnProperty('enabled')) {
                obj['enabled'] = ApiClient.convertToType(data['enabled'], 'Boolean');
            }
            if (data.hasOwnProperty('dateCreated')) {
                obj['dateCreated'] = ApiClient.convertToType(data['dateCreated'], 'Date');
            }
            if (data.hasOwnProperty('lastUpdated')) {
                obj['lastUpdated'] = ApiClient.convertToType(data['lastUpdated'], 'Date');
            }
            if (data.hasOwnProperty('lastRun')) {
                obj['lastRun'] = ApiClient.convertToType(data['lastRun'], 'Date');
            }
            if (data.hasOwnProperty('lastResult')) {
                obj['lastResult'] = ApiClient.convertToType(data['lastResult'], 'String');
            }
            if (data.hasOwnProperty('createdBy')) {
                obj['createdBy'] = JobCreatedBy.constructFromObject(data['createdBy']);
            }
            if (data.hasOwnProperty('targetType')) {
                obj['targetType'] = ApiClient.convertToType(data['targetType'], 'String');
            }
            if (data.hasOwnProperty('targets')) {
                obj['targets'] = ApiClient.convertToType(data['targets'], [JobTargets]);
            }
            if (data.hasOwnProperty('scanPath')) {
                obj['scanPath'] = ApiClient.convertToType(data['scanPath'], 'String');
            }
            if (data.hasOwnProperty('securityProfile')) {
                obj['securityProfile'] = ApiClient.convertToType(data['securityProfile'], 'String');
            }
            if (data.hasOwnProperty('customConfig')) {
                obj['customConfig'] = ApiClient.convertToType(data['customConfig'], 'String');
            }
            if (data.hasOwnProperty('customOptions')) {
                obj['customOptions'] = JobCustomOptions.constructFromObject(data['customOptions']);
            }
        }
        return obj;
    }


}

/**
 * @member {Number} id
 */
Job.prototype['id'] = undefined;

/**
 * @member {String} name
 */
Job.prototype['name'] = undefined;

/**
 * @member {Array.<String>} labels
 */
Job.prototype['labels'] = undefined;

/**
 * @member {module:model/InlineResponse20094Network} type
 */
Job.prototype['type'] = undefined;

/**
 * @member {module:model/JobWorkflow} workflow
 */
Job.prototype['workflow'] = undefined;

/**
 * @member {module:model/JobTask} task
 */
Job.prototype['task'] = undefined;

/**
 * @member {module:model/JobSecurityPackage} securityPackage
 */
Job.prototype['securityPackage'] = undefined;

/**
 * @member {String} jobSummary
 */
Job.prototype['jobSummary'] = undefined;

/**
 * @member {module:model/OneOfstringlong} scheduleMode
 */
Job.prototype['scheduleMode'] = undefined;

/**
 * @member {String} dateTime
 */
Job.prototype['dateTime'] = undefined;

/**
 * @member {String} status
 */
Job.prototype['status'] = undefined;

/**
 * @member {String} namespace
 */
Job.prototype['namespace'] = undefined;

/**
 * @member {String} category
 */
Job.prototype['category'] = undefined;

/**
 * @member {String} description
 */
Job.prototype['description'] = undefined;

/**
 * @member {Boolean} enabled
 */
Job.prototype['enabled'] = undefined;

/**
 * @member {Date} dateCreated
 */
Job.prototype['dateCreated'] = undefined;

/**
 * @member {Date} lastUpdated
 */
Job.prototype['lastUpdated'] = undefined;

/**
 * @member {Date} lastRun
 */
Job.prototype['lastRun'] = undefined;

/**
 * @member {String} lastResult
 */
Job.prototype['lastResult'] = undefined;

/**
 * @member {module:model/JobCreatedBy} createdBy
 */
Job.prototype['createdBy'] = undefined;

/**
 * @member {String} targetType
 */
Job.prototype['targetType'] = undefined;

/**
 * @member {Array.<module:model/JobTargets>} targets
 */
Job.prototype['targets'] = undefined;

/**
 * Scan Checklist. Only applies to type scap-package.
 * @member {String} scanPath
 */
Job.prototype['scanPath'] = undefined;

/**
 * Security Profile. Only applies to type scap-package.
 * @member {String} securityProfile
 */
Job.prototype['securityProfile'] = undefined;

/**
 * @member {String} customConfig
 */
Job.prototype['customConfig'] = undefined;

/**
 * @member {module:model/JobCustomOptions} customOptions
 */
Job.prototype['customOptions'] = undefined;






export default Job;

