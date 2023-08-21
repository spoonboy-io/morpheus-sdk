<?php
/**
 * Job
 *
 * PHP version 7.2
 *
 * @category Class
 * @package  OpenAPI\Client
 * @author   OpenAPI Generator team
 * @link     https://openapi-generator.tech
 */

/**
 * Morpheus API
 *
 * Morpheus is a powerful cloud management tool that provides provisioning, monitoring, logging, backups, and application deployment strategies.  This document describes the Morpheus API protocol and the available endpoints. Sections are organized in the same manner as they appear in the Morpheus UI.
 *
 * The version of the OpenAPI document: 6.2.1
 * Contact: dev@morpheusdata.com
 * Generated by: https://openapi-generator.tech
 * OpenAPI Generator version: 5.0.0
 */

/**
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */

namespace OpenAPI\Client\Model;

use \ArrayAccess;
use \OpenAPI\Client\ObjectSerializer;

/**
 * Job Class Doc Comment
 *
 * @category Class
 * @package  OpenAPI\Client
 * @author   OpenAPI Generator team
 * @link     https://openapi-generator.tech
 * @implements \ArrayAccess<TKey, TValue>
 * @template TKey int|null
 * @template TValue mixed|null  
 */
class Job implements ModelInterface, ArrayAccess, \JsonSerializable
{
    public const DISCRIMINATOR = null;

    /**
      * The original name of the model.
      *
      * @var string
      */
    protected static $openAPIModelName = 'job';

    /**
      * Array of property to type mappings. Used for (de)serialization
      *
      * @var string[]
      */
    protected static $openAPITypes = [
        'id' => 'int',
        'name' => 'string',
        'labels' => 'string[]',
        'type' => '\OpenAPI\Client\Model\InlineResponse20094Network',
        'workflow' => '\OpenAPI\Client\Model\JobWorkflow',
        'task' => '\OpenAPI\Client\Model\JobTask',
        'security_package' => '\OpenAPI\Client\Model\JobSecurityPackage',
        'job_summary' => 'string',
        'schedule_mode' => 'OneOfStringLong',
        'date_time' => 'string',
        'status' => 'string',
        'namespace' => 'string',
        'category' => 'string',
        'description' => 'string',
        'enabled' => 'bool',
        'date_created' => '\DateTime',
        'last_updated' => '\DateTime',
        'last_run' => '\DateTime',
        'last_result' => 'string',
        'created_by' => '\OpenAPI\Client\Model\JobCreatedBy',
        'target_type' => 'string',
        'targets' => '\OpenAPI\Client\Model\JobTargets[]',
        'scan_path' => 'string',
        'security_profile' => 'string',
        'custom_config' => 'string',
        'custom_options' => '\OpenAPI\Client\Model\JobCustomOptions'
    ];

    /**
      * Array of property to format mappings. Used for (de)serialization
      *
      * @var string[]
      * @phpstan-var array<string, string|null>
      * @psalm-var array<string, string|null>
      */
    protected static $openAPIFormats = [
        'id' => 'int64',
        'name' => null,
        'labels' => null,
        'type' => null,
        'workflow' => null,
        'task' => null,
        'security_package' => null,
        'job_summary' => null,
        'schedule_mode' => null,
        'date_time' => null,
        'status' => null,
        'namespace' => null,
        'category' => null,
        'description' => null,
        'enabled' => null,
        'date_created' => 'date-time',
        'last_updated' => 'date-time',
        'last_run' => 'date-time',
        'last_result' => null,
        'created_by' => null,
        'target_type' => null,
        'targets' => null,
        'scan_path' => null,
        'security_profile' => null,
        'custom_config' => null,
        'custom_options' => null
    ];

    /**
     * Array of property to type mappings. Used for (de)serialization
     *
     * @return array
     */
    public static function openAPITypes()
    {
        return self::$openAPITypes;
    }

    /**
     * Array of property to format mappings. Used for (de)serialization
     *
     * @return array
     */
    public static function openAPIFormats()
    {
        return self::$openAPIFormats;
    }

    /**
     * Array of attributes where the key is the local name,
     * and the value is the original name
     *
     * @var string[]
     */
    protected static $attributeMap = [
        'id' => 'id',
        'name' => 'name',
        'labels' => 'labels',
        'type' => 'type',
        'workflow' => 'workflow',
        'task' => 'task',
        'security_package' => 'securityPackage',
        'job_summary' => 'jobSummary',
        'schedule_mode' => 'scheduleMode',
        'date_time' => 'dateTime',
        'status' => 'status',
        'namespace' => 'namespace',
        'category' => 'category',
        'description' => 'description',
        'enabled' => 'enabled',
        'date_created' => 'dateCreated',
        'last_updated' => 'lastUpdated',
        'last_run' => 'lastRun',
        'last_result' => 'lastResult',
        'created_by' => 'createdBy',
        'target_type' => 'targetType',
        'targets' => 'targets',
        'scan_path' => 'scanPath',
        'security_profile' => 'securityProfile',
        'custom_config' => 'customConfig',
        'custom_options' => 'customOptions'
    ];

    /**
     * Array of attributes to setter functions (for deserialization of responses)
     *
     * @var string[]
     */
    protected static $setters = [
        'id' => 'setId',
        'name' => 'setName',
        'labels' => 'setLabels',
        'type' => 'setType',
        'workflow' => 'setWorkflow',
        'task' => 'setTask',
        'security_package' => 'setSecurityPackage',
        'job_summary' => 'setJobSummary',
        'schedule_mode' => 'setScheduleMode',
        'date_time' => 'setDateTime',
        'status' => 'setStatus',
        'namespace' => 'setNamespace',
        'category' => 'setCategory',
        'description' => 'setDescription',
        'enabled' => 'setEnabled',
        'date_created' => 'setDateCreated',
        'last_updated' => 'setLastUpdated',
        'last_run' => 'setLastRun',
        'last_result' => 'setLastResult',
        'created_by' => 'setCreatedBy',
        'target_type' => 'setTargetType',
        'targets' => 'setTargets',
        'scan_path' => 'setScanPath',
        'security_profile' => 'setSecurityProfile',
        'custom_config' => 'setCustomConfig',
        'custom_options' => 'setCustomOptions'
    ];

    /**
     * Array of attributes to getter functions (for serialization of requests)
     *
     * @var string[]
     */
    protected static $getters = [
        'id' => 'getId',
        'name' => 'getName',
        'labels' => 'getLabels',
        'type' => 'getType',
        'workflow' => 'getWorkflow',
        'task' => 'getTask',
        'security_package' => 'getSecurityPackage',
        'job_summary' => 'getJobSummary',
        'schedule_mode' => 'getScheduleMode',
        'date_time' => 'getDateTime',
        'status' => 'getStatus',
        'namespace' => 'getNamespace',
        'category' => 'getCategory',
        'description' => 'getDescription',
        'enabled' => 'getEnabled',
        'date_created' => 'getDateCreated',
        'last_updated' => 'getLastUpdated',
        'last_run' => 'getLastRun',
        'last_result' => 'getLastResult',
        'created_by' => 'getCreatedBy',
        'target_type' => 'getTargetType',
        'targets' => 'getTargets',
        'scan_path' => 'getScanPath',
        'security_profile' => 'getSecurityProfile',
        'custom_config' => 'getCustomConfig',
        'custom_options' => 'getCustomOptions'
    ];

    /**
     * Array of attributes where the key is the local name,
     * and the value is the original name
     *
     * @return array
     */
    public static function attributeMap()
    {
        return self::$attributeMap;
    }

    /**
     * Array of attributes to setter functions (for deserialization of responses)
     *
     * @return array
     */
    public static function setters()
    {
        return self::$setters;
    }

    /**
     * Array of attributes to getter functions (for serialization of requests)
     *
     * @return array
     */
    public static function getters()
    {
        return self::$getters;
    }

    /**
     * The original name of the model.
     *
     * @return string
     */
    public function getModelName()
    {
        return self::$openAPIModelName;
    }

    

    

    /**
     * Associative array for storing property values
     *
     * @var mixed[]
     */
    protected $container = [];

    /**
     * Constructor
     *
     * @param mixed[] $data Associated array of property values
     *                      initializing the model
     */
    public function __construct(array $data = null)
    {
        $this->container['id'] = $data['id'] ?? null;
        $this->container['name'] = $data['name'] ?? null;
        $this->container['labels'] = $data['labels'] ?? null;
        $this->container['type'] = $data['type'] ?? null;
        $this->container['workflow'] = $data['workflow'] ?? null;
        $this->container['task'] = $data['task'] ?? null;
        $this->container['security_package'] = $data['security_package'] ?? null;
        $this->container['job_summary'] = $data['job_summary'] ?? null;
        $this->container['schedule_mode'] = $data['schedule_mode'] ?? null;
        $this->container['date_time'] = $data['date_time'] ?? null;
        $this->container['status'] = $data['status'] ?? null;
        $this->container['namespace'] = $data['namespace'] ?? null;
        $this->container['category'] = $data['category'] ?? null;
        $this->container['description'] = $data['description'] ?? null;
        $this->container['enabled'] = $data['enabled'] ?? null;
        $this->container['date_created'] = $data['date_created'] ?? null;
        $this->container['last_updated'] = $data['last_updated'] ?? null;
        $this->container['last_run'] = $data['last_run'] ?? null;
        $this->container['last_result'] = $data['last_result'] ?? null;
        $this->container['created_by'] = $data['created_by'] ?? null;
        $this->container['target_type'] = $data['target_type'] ?? null;
        $this->container['targets'] = $data['targets'] ?? null;
        $this->container['scan_path'] = $data['scan_path'] ?? null;
        $this->container['security_profile'] = $data['security_profile'] ?? null;
        $this->container['custom_config'] = $data['custom_config'] ?? null;
        $this->container['custom_options'] = $data['custom_options'] ?? null;
    }

    /**
     * Show all the invalid properties with reasons.
     *
     * @return array invalid properties with reasons
     */
    public function listInvalidProperties()
    {
        $invalidProperties = [];

        return $invalidProperties;
    }

    /**
     * Validate all the properties in the model
     * return true if all passed
     *
     * @return bool True if all properties are valid
     */
    public function valid()
    {
        return count($this->listInvalidProperties()) === 0;
    }


    /**
     * Gets id
     *
     * @return int|null
     */
    public function getId()
    {
        return $this->container['id'];
    }

    /**
     * Sets id
     *
     * @param int|null $id id
     *
     * @return self
     */
    public function setId($id)
    {
        $this->container['id'] = $id;

        return $this;
    }

    /**
     * Gets name
     *
     * @return string|null
     */
    public function getName()
    {
        return $this->container['name'];
    }

    /**
     * Sets name
     *
     * @param string|null $name name
     *
     * @return self
     */
    public function setName($name)
    {
        $this->container['name'] = $name;

        return $this;
    }

    /**
     * Gets labels
     *
     * @return string[]|null
     */
    public function getLabels()
    {
        return $this->container['labels'];
    }

    /**
     * Sets labels
     *
     * @param string[]|null $labels labels
     *
     * @return self
     */
    public function setLabels($labels)
    {
        $this->container['labels'] = $labels;

        return $this;
    }

    /**
     * Gets type
     *
     * @return \OpenAPI\Client\Model\InlineResponse20094Network|null
     */
    public function getType()
    {
        return $this->container['type'];
    }

    /**
     * Sets type
     *
     * @param \OpenAPI\Client\Model\InlineResponse20094Network|null $type type
     *
     * @return self
     */
    public function setType($type)
    {
        $this->container['type'] = $type;

        return $this;
    }

    /**
     * Gets workflow
     *
     * @return \OpenAPI\Client\Model\JobWorkflow|null
     */
    public function getWorkflow()
    {
        return $this->container['workflow'];
    }

    /**
     * Sets workflow
     *
     * @param \OpenAPI\Client\Model\JobWorkflow|null $workflow workflow
     *
     * @return self
     */
    public function setWorkflow($workflow)
    {
        $this->container['workflow'] = $workflow;

        return $this;
    }

    /**
     * Gets task
     *
     * @return \OpenAPI\Client\Model\JobTask|null
     */
    public function getTask()
    {
        return $this->container['task'];
    }

    /**
     * Sets task
     *
     * @param \OpenAPI\Client\Model\JobTask|null $task task
     *
     * @return self
     */
    public function setTask($task)
    {
        $this->container['task'] = $task;

        return $this;
    }

    /**
     * Gets security_package
     *
     * @return \OpenAPI\Client\Model\JobSecurityPackage|null
     */
    public function getSecurityPackage()
    {
        return $this->container['security_package'];
    }

    /**
     * Sets security_package
     *
     * @param \OpenAPI\Client\Model\JobSecurityPackage|null $security_package security_package
     *
     * @return self
     */
    public function setSecurityPackage($security_package)
    {
        $this->container['security_package'] = $security_package;

        return $this;
    }

    /**
     * Gets job_summary
     *
     * @return string|null
     */
    public function getJobSummary()
    {
        return $this->container['job_summary'];
    }

    /**
     * Sets job_summary
     *
     * @param string|null $job_summary job_summary
     *
     * @return self
     */
    public function setJobSummary($job_summary)
    {
        $this->container['job_summary'] = $job_summary;

        return $this;
    }

    /**
     * Gets schedule_mode
     *
     * @return OneOfStringLong|null
     */
    public function getScheduleMode()
    {
        return $this->container['schedule_mode'];
    }

    /**
     * Sets schedule_mode
     *
     * @param OneOfStringLong|null $schedule_mode schedule_mode
     *
     * @return self
     */
    public function setScheduleMode($schedule_mode)
    {
        $this->container['schedule_mode'] = $schedule_mode;

        return $this;
    }

    /**
     * Gets date_time
     *
     * @return string|null
     */
    public function getDateTime()
    {
        return $this->container['date_time'];
    }

    /**
     * Sets date_time
     *
     * @param string|null $date_time date_time
     *
     * @return self
     */
    public function setDateTime($date_time)
    {
        $this->container['date_time'] = $date_time;

        return $this;
    }

    /**
     * Gets status
     *
     * @return string|null
     */
    public function getStatus()
    {
        return $this->container['status'];
    }

    /**
     * Sets status
     *
     * @param string|null $status status
     *
     * @return self
     */
    public function setStatus($status)
    {
        $this->container['status'] = $status;

        return $this;
    }

    /**
     * Gets namespace
     *
     * @return string|null
     */
    public function getNamespace()
    {
        return $this->container['namespace'];
    }

    /**
     * Sets namespace
     *
     * @param string|null $namespace namespace
     *
     * @return self
     */
    public function setNamespace($namespace)
    {
        $this->container['namespace'] = $namespace;

        return $this;
    }

    /**
     * Gets category
     *
     * @return string|null
     */
    public function getCategory()
    {
        return $this->container['category'];
    }

    /**
     * Sets category
     *
     * @param string|null $category category
     *
     * @return self
     */
    public function setCategory($category)
    {
        $this->container['category'] = $category;

        return $this;
    }

    /**
     * Gets description
     *
     * @return string|null
     */
    public function getDescription()
    {
        return $this->container['description'];
    }

    /**
     * Sets description
     *
     * @param string|null $description description
     *
     * @return self
     */
    public function setDescription($description)
    {
        $this->container['description'] = $description;

        return $this;
    }

    /**
     * Gets enabled
     *
     * @return bool|null
     */
    public function getEnabled()
    {
        return $this->container['enabled'];
    }

    /**
     * Sets enabled
     *
     * @param bool|null $enabled enabled
     *
     * @return self
     */
    public function setEnabled($enabled)
    {
        $this->container['enabled'] = $enabled;

        return $this;
    }

    /**
     * Gets date_created
     *
     * @return \DateTime|null
     */
    public function getDateCreated()
    {
        return $this->container['date_created'];
    }

    /**
     * Sets date_created
     *
     * @param \DateTime|null $date_created date_created
     *
     * @return self
     */
    public function setDateCreated($date_created)
    {
        $this->container['date_created'] = $date_created;

        return $this;
    }

    /**
     * Gets last_updated
     *
     * @return \DateTime|null
     */
    public function getLastUpdated()
    {
        return $this->container['last_updated'];
    }

    /**
     * Sets last_updated
     *
     * @param \DateTime|null $last_updated last_updated
     *
     * @return self
     */
    public function setLastUpdated($last_updated)
    {
        $this->container['last_updated'] = $last_updated;

        return $this;
    }

    /**
     * Gets last_run
     *
     * @return \DateTime|null
     */
    public function getLastRun()
    {
        return $this->container['last_run'];
    }

    /**
     * Sets last_run
     *
     * @param \DateTime|null $last_run last_run
     *
     * @return self
     */
    public function setLastRun($last_run)
    {
        $this->container['last_run'] = $last_run;

        return $this;
    }

    /**
     * Gets last_result
     *
     * @return string|null
     */
    public function getLastResult()
    {
        return $this->container['last_result'];
    }

    /**
     * Sets last_result
     *
     * @param string|null $last_result last_result
     *
     * @return self
     */
    public function setLastResult($last_result)
    {
        $this->container['last_result'] = $last_result;

        return $this;
    }

    /**
     * Gets created_by
     *
     * @return \OpenAPI\Client\Model\JobCreatedBy|null
     */
    public function getCreatedBy()
    {
        return $this->container['created_by'];
    }

    /**
     * Sets created_by
     *
     * @param \OpenAPI\Client\Model\JobCreatedBy|null $created_by created_by
     *
     * @return self
     */
    public function setCreatedBy($created_by)
    {
        $this->container['created_by'] = $created_by;

        return $this;
    }

    /**
     * Gets target_type
     *
     * @return string|null
     */
    public function getTargetType()
    {
        return $this->container['target_type'];
    }

    /**
     * Sets target_type
     *
     * @param string|null $target_type target_type
     *
     * @return self
     */
    public function setTargetType($target_type)
    {
        $this->container['target_type'] = $target_type;

        return $this;
    }

    /**
     * Gets targets
     *
     * @return \OpenAPI\Client\Model\JobTargets[]|null
     */
    public function getTargets()
    {
        return $this->container['targets'];
    }

    /**
     * Sets targets
     *
     * @param \OpenAPI\Client\Model\JobTargets[]|null $targets targets
     *
     * @return self
     */
    public function setTargets($targets)
    {
        $this->container['targets'] = $targets;

        return $this;
    }

    /**
     * Gets scan_path
     *
     * @return string|null
     */
    public function getScanPath()
    {
        return $this->container['scan_path'];
    }

    /**
     * Sets scan_path
     *
     * @param string|null $scan_path Scan Checklist. Only applies to type scap-package.
     *
     * @return self
     */
    public function setScanPath($scan_path)
    {
        $this->container['scan_path'] = $scan_path;

        return $this;
    }

    /**
     * Gets security_profile
     *
     * @return string|null
     */
    public function getSecurityProfile()
    {
        return $this->container['security_profile'];
    }

    /**
     * Sets security_profile
     *
     * @param string|null $security_profile Security Profile. Only applies to type scap-package.
     *
     * @return self
     */
    public function setSecurityProfile($security_profile)
    {
        $this->container['security_profile'] = $security_profile;

        return $this;
    }

    /**
     * Gets custom_config
     *
     * @return string|null
     */
    public function getCustomConfig()
    {
        return $this->container['custom_config'];
    }

    /**
     * Sets custom_config
     *
     * @param string|null $custom_config custom_config
     *
     * @return self
     */
    public function setCustomConfig($custom_config)
    {
        $this->container['custom_config'] = $custom_config;

        return $this;
    }

    /**
     * Gets custom_options
     *
     * @return \OpenAPI\Client\Model\JobCustomOptions|null
     */
    public function getCustomOptions()
    {
        return $this->container['custom_options'];
    }

    /**
     * Sets custom_options
     *
     * @param \OpenAPI\Client\Model\JobCustomOptions|null $custom_options custom_options
     *
     * @return self
     */
    public function setCustomOptions($custom_options)
    {
        $this->container['custom_options'] = $custom_options;

        return $this;
    }
    /**
     * Returns true if offset exists. False otherwise.
     *
     * @param integer $offset Offset
     *
     * @return boolean
     */
    public function offsetExists($offset)
    {
        return isset($this->container[$offset]);
    }

    /**
     * Gets offset.
     *
     * @param integer $offset Offset
     *
     * @return mixed|null
     */
    public function offsetGet($offset)
    {
        return $this->container[$offset] ?? null;
    }

    /**
     * Sets value based on offset.
     *
     * @param int|null $offset Offset
     * @param mixed    $value  Value to be set
     *
     * @return void
     */
    public function offsetSet($offset, $value)
    {
        if (is_null($offset)) {
            $this->container[] = $value;
        } else {
            $this->container[$offset] = $value;
        }
    }

    /**
     * Unsets offset.
     *
     * @param integer $offset Offset
     *
     * @return void
     */
    public function offsetUnset($offset)
    {
        unset($this->container[$offset]);
    }

    /**
     * Serializes the object to a value that can be serialized natively by json_encode().
     * @link https://www.php.net/manual/en/jsonserializable.jsonserialize.php
     *
     * @return mixed Returns data which can be serialized by json_encode(), which is a value
     * of any type other than a resource.
     */
    public function jsonSerialize()
    {
       return ObjectSerializer::sanitizeForSerialization($this);
    }

    /**
     * Gets the string presentation of the object
     *
     * @return string
     */
    public function __toString()
    {
        return json_encode(
            ObjectSerializer::sanitizeForSerialization($this),
            JSON_PRETTY_PRINT
        );
    }

    /**
     * Gets a header-safe presentation of the object
     *
     * @return string
     */
    public function toHeaderValue()
    {
        return json_encode(ObjectSerializer::sanitizeForSerialization($this));
    }
}

