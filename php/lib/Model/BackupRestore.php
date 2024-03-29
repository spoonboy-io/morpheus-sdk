<?php
/**
 * BackupRestore
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
 * BackupRestore Class Doc Comment
 *
 * @category Class
 * @package  OpenAPI\Client
 * @author   OpenAPI Generator team
 * @link     https://openapi-generator.tech
 * @implements \ArrayAccess<TKey, TValue>
 * @template TKey int|null
 * @template TValue mixed|null  
 */
class BackupRestore implements ModelInterface, ArrayAccess, \JsonSerializable
{
    public const DISCRIMINATOR = null;

    /**
      * The original name of the model.
      *
      * @var string
      */
    protected static $openAPIModelName = 'backupRestore';

    /**
      * Array of property to type mappings. Used for (de)serialization
      *
      * @var string[]
      */
    protected static $openAPITypes = [
        'id' => 'int',
        'backup_result_id' => 'int',
        'backup_id' => 'int',
        'backup' => '\OpenAPI\Client\Model\BackupRestoreBackup',
        'container_id' => 'int',
        'container' => '\OpenAPI\Client\Model\BackupRestoreContainer',
        'instance' => '\OpenAPI\Client\Model\BackupInstance',
        'restore_to_new' => 'bool',
        'status' => 'string',
        'error_message' => 'string',
        'start_date' => '\DateTime',
        'end_date' => '\DateTime',
        'duration_millis' => 'int',
        'external_id' => 'string',
        'external_status_ref' => 'string',
        'date_created' => '\DateTime',
        'last_updated' => '\DateTime'
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
        'backup_result_id' => 'int64',
        'backup_id' => 'int64',
        'backup' => null,
        'container_id' => 'int64',
        'container' => null,
        'instance' => null,
        'restore_to_new' => null,
        'status' => null,
        'error_message' => null,
        'start_date' => 'date-time',
        'end_date' => 'date-time',
        'duration_millis' => 'int64',
        'external_id' => null,
        'external_status_ref' => null,
        'date_created' => 'date-time',
        'last_updated' => 'date-time'
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
        'backup_result_id' => 'backupResultId',
        'backup_id' => 'backupId',
        'backup' => 'backup',
        'container_id' => 'containerId',
        'container' => 'container',
        'instance' => 'instance',
        'restore_to_new' => 'restoreToNew',
        'status' => 'status',
        'error_message' => 'errorMessage',
        'start_date' => 'startDate',
        'end_date' => 'endDate',
        'duration_millis' => 'durationMillis',
        'external_id' => 'externalId',
        'external_status_ref' => 'externalStatusRef',
        'date_created' => 'dateCreated',
        'last_updated' => 'lastUpdated'
    ];

    /**
     * Array of attributes to setter functions (for deserialization of responses)
     *
     * @var string[]
     */
    protected static $setters = [
        'id' => 'setId',
        'backup_result_id' => 'setBackupResultId',
        'backup_id' => 'setBackupId',
        'backup' => 'setBackup',
        'container_id' => 'setContainerId',
        'container' => 'setContainer',
        'instance' => 'setInstance',
        'restore_to_new' => 'setRestoreToNew',
        'status' => 'setStatus',
        'error_message' => 'setErrorMessage',
        'start_date' => 'setStartDate',
        'end_date' => 'setEndDate',
        'duration_millis' => 'setDurationMillis',
        'external_id' => 'setExternalId',
        'external_status_ref' => 'setExternalStatusRef',
        'date_created' => 'setDateCreated',
        'last_updated' => 'setLastUpdated'
    ];

    /**
     * Array of attributes to getter functions (for serialization of requests)
     *
     * @var string[]
     */
    protected static $getters = [
        'id' => 'getId',
        'backup_result_id' => 'getBackupResultId',
        'backup_id' => 'getBackupId',
        'backup' => 'getBackup',
        'container_id' => 'getContainerId',
        'container' => 'getContainer',
        'instance' => 'getInstance',
        'restore_to_new' => 'getRestoreToNew',
        'status' => 'getStatus',
        'error_message' => 'getErrorMessage',
        'start_date' => 'getStartDate',
        'end_date' => 'getEndDate',
        'duration_millis' => 'getDurationMillis',
        'external_id' => 'getExternalId',
        'external_status_ref' => 'getExternalStatusRef',
        'date_created' => 'getDateCreated',
        'last_updated' => 'getLastUpdated'
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
        $this->container['backup_result_id'] = $data['backup_result_id'] ?? null;
        $this->container['backup_id'] = $data['backup_id'] ?? null;
        $this->container['backup'] = $data['backup'] ?? null;
        $this->container['container_id'] = $data['container_id'] ?? null;
        $this->container['container'] = $data['container'] ?? null;
        $this->container['instance'] = $data['instance'] ?? null;
        $this->container['restore_to_new'] = $data['restore_to_new'] ?? null;
        $this->container['status'] = $data['status'] ?? null;
        $this->container['error_message'] = $data['error_message'] ?? null;
        $this->container['start_date'] = $data['start_date'] ?? null;
        $this->container['end_date'] = $data['end_date'] ?? null;
        $this->container['duration_millis'] = $data['duration_millis'] ?? null;
        $this->container['external_id'] = $data['external_id'] ?? null;
        $this->container['external_status_ref'] = $data['external_status_ref'] ?? null;
        $this->container['date_created'] = $data['date_created'] ?? null;
        $this->container['last_updated'] = $data['last_updated'] ?? null;
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
     * @param int|null $id Backup Result ID
     *
     * @return self
     */
    public function setId($id)
    {
        $this->container['id'] = $id;

        return $this;
    }

    /**
     * Gets backup_result_id
     *
     * @return int|null
     */
    public function getBackupResultId()
    {
        return $this->container['backup_result_id'];
    }

    /**
     * Sets backup_result_id
     *
     * @param int|null $backup_result_id backup_result_id
     *
     * @return self
     */
    public function setBackupResultId($backup_result_id)
    {
        $this->container['backup_result_id'] = $backup_result_id;

        return $this;
    }

    /**
     * Gets backup_id
     *
     * @return int|null
     */
    public function getBackupId()
    {
        return $this->container['backup_id'];
    }

    /**
     * Sets backup_id
     *
     * @param int|null $backup_id backup_id
     *
     * @return self
     */
    public function setBackupId($backup_id)
    {
        $this->container['backup_id'] = $backup_id;

        return $this;
    }

    /**
     * Gets backup
     *
     * @return \OpenAPI\Client\Model\BackupRestoreBackup|null
     */
    public function getBackup()
    {
        return $this->container['backup'];
    }

    /**
     * Sets backup
     *
     * @param \OpenAPI\Client\Model\BackupRestoreBackup|null $backup backup
     *
     * @return self
     */
    public function setBackup($backup)
    {
        $this->container['backup'] = $backup;

        return $this;
    }

    /**
     * Gets container_id
     *
     * @return int|null
     */
    public function getContainerId()
    {
        return $this->container['container_id'];
    }

    /**
     * Sets container_id
     *
     * @param int|null $container_id container_id
     *
     * @return self
     */
    public function setContainerId($container_id)
    {
        $this->container['container_id'] = $container_id;

        return $this;
    }

    /**
     * Gets container
     *
     * @return \OpenAPI\Client\Model\BackupRestoreContainer|null
     */
    public function getContainer()
    {
        return $this->container['container'];
    }

    /**
     * Sets container
     *
     * @param \OpenAPI\Client\Model\BackupRestoreContainer|null $container container
     *
     * @return self
     */
    public function setContainer($container)
    {
        $this->container['container'] = $container;

        return $this;
    }

    /**
     * Gets instance
     *
     * @return \OpenAPI\Client\Model\BackupInstance|null
     */
    public function getInstance()
    {
        return $this->container['instance'];
    }

    /**
     * Sets instance
     *
     * @param \OpenAPI\Client\Model\BackupInstance|null $instance instance
     *
     * @return self
     */
    public function setInstance($instance)
    {
        $this->container['instance'] = $instance;

        return $this;
    }

    /**
     * Gets restore_to_new
     *
     * @return bool|null
     */
    public function getRestoreToNew()
    {
        return $this->container['restore_to_new'];
    }

    /**
     * Sets restore_to_new
     *
     * @param bool|null $restore_to_new restore_to_new
     *
     * @return self
     */
    public function setRestoreToNew($restore_to_new)
    {
        $this->container['restore_to_new'] = $restore_to_new;

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
     * Gets error_message
     *
     * @return string|null
     */
    public function getErrorMessage()
    {
        return $this->container['error_message'];
    }

    /**
     * Sets error_message
     *
     * @param string|null $error_message error_message
     *
     * @return self
     */
    public function setErrorMessage($error_message)
    {
        $this->container['error_message'] = $error_message;

        return $this;
    }

    /**
     * Gets start_date
     *
     * @return \DateTime|null
     */
    public function getStartDate()
    {
        return $this->container['start_date'];
    }

    /**
     * Sets start_date
     *
     * @param \DateTime|null $start_date start_date
     *
     * @return self
     */
    public function setStartDate($start_date)
    {
        $this->container['start_date'] = $start_date;

        return $this;
    }

    /**
     * Gets end_date
     *
     * @return \DateTime|null
     */
    public function getEndDate()
    {
        return $this->container['end_date'];
    }

    /**
     * Sets end_date
     *
     * @param \DateTime|null $end_date end_date
     *
     * @return self
     */
    public function setEndDate($end_date)
    {
        $this->container['end_date'] = $end_date;

        return $this;
    }

    /**
     * Gets duration_millis
     *
     * @return int|null
     */
    public function getDurationMillis()
    {
        return $this->container['duration_millis'];
    }

    /**
     * Sets duration_millis
     *
     * @param int|null $duration_millis duration_millis
     *
     * @return self
     */
    public function setDurationMillis($duration_millis)
    {
        $this->container['duration_millis'] = $duration_millis;

        return $this;
    }

    /**
     * Gets external_id
     *
     * @return string|null
     */
    public function getExternalId()
    {
        return $this->container['external_id'];
    }

    /**
     * Sets external_id
     *
     * @param string|null $external_id external_id
     *
     * @return self
     */
    public function setExternalId($external_id)
    {
        $this->container['external_id'] = $external_id;

        return $this;
    }

    /**
     * Gets external_status_ref
     *
     * @return string|null
     */
    public function getExternalStatusRef()
    {
        return $this->container['external_status_ref'];
    }

    /**
     * Sets external_status_ref
     *
     * @param string|null $external_status_ref external_status_ref
     *
     * @return self
     */
    public function setExternalStatusRef($external_status_ref)
    {
        $this->container['external_status_ref'] = $external_status_ref;

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
     * @param \DateTime|null $date_created Date Created
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
     * @param \DateTime|null $last_updated Last Updated
     *
     * @return self
     */
    public function setLastUpdated($last_updated)
    {
        $this->container['last_updated'] = $last_updated;

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


