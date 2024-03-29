<?php
/**
 * BackupResult
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
 * BackupResult Class Doc Comment
 *
 * @category Class
 * @package  OpenAPI\Client
 * @author   OpenAPI Generator team
 * @link     https://openapi-generator.tech
 * @implements \ArrayAccess<TKey, TValue>
 * @template TKey int|null
 * @template TValue mixed|null  
 */
class BackupResult implements ModelInterface, ArrayAccess, \JsonSerializable
{
    public const DISCRIMINATOR = null;

    /**
      * The original name of the model.
      *
      * @var string
      */
    protected static $openAPIModelName = 'backupResult';

    /**
      * Array of property to type mappings. Used for (de)serialization
      *
      * @var string[]
      */
    protected static $openAPITypes = [
        'id' => 'int',
        'backup' => '\OpenAPI\Client\Model\BackupJobBackups',
        'backup_set_id' => 'string',
        'instance_id' => 'int',
        'container_id' => 'int',
        'server_id' => 'int',
        'status' => 'string',
        'error_message' => 'string',
        'start_date' => '\DateTime',
        'end_date' => '\DateTime',
        'duration_millis' => 'int',
        'size_in_bytes' => 'int',
        'size_in_mb' => 'int',
        'volume_path' => 'string',
        'result_archive' => 'string',
        'result_path' => 'string',
        'external_id' => 'string',
        'snapshot_id' => 'string',
        'snapshot_external_id' => 'string',
        'created_by' => '\OpenAPI\Client\Model\InlineResponse200108NetworkFloatingIpCreatedBy',
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
        'backup' => null,
        'backup_set_id' => null,
        'instance_id' => 'int64',
        'container_id' => 'int64',
        'server_id' => 'int64',
        'status' => null,
        'error_message' => null,
        'start_date' => 'date-time',
        'end_date' => 'date-time',
        'duration_millis' => 'int64',
        'size_in_bytes' => 'int64',
        'size_in_mb' => 'int64',
        'volume_path' => null,
        'result_archive' => null,
        'result_path' => null,
        'external_id' => null,
        'snapshot_id' => null,
        'snapshot_external_id' => null,
        'created_by' => null,
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
        'backup' => 'backup',
        'backup_set_id' => 'backupSetId',
        'instance_id' => 'instanceId',
        'container_id' => 'containerId',
        'server_id' => 'serverId',
        'status' => 'status',
        'error_message' => 'errorMessage',
        'start_date' => 'startDate',
        'end_date' => 'endDate',
        'duration_millis' => 'durationMillis',
        'size_in_bytes' => 'sizeInBytes',
        'size_in_mb' => 'sizeInMb',
        'volume_path' => 'volumePath',
        'result_archive' => 'resultArchive',
        'result_path' => 'resultPath',
        'external_id' => 'externalId',
        'snapshot_id' => 'snapshotId',
        'snapshot_external_id' => 'snapshotExternalId',
        'created_by' => 'createdBy',
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
        'backup' => 'setBackup',
        'backup_set_id' => 'setBackupSetId',
        'instance_id' => 'setInstanceId',
        'container_id' => 'setContainerId',
        'server_id' => 'setServerId',
        'status' => 'setStatus',
        'error_message' => 'setErrorMessage',
        'start_date' => 'setStartDate',
        'end_date' => 'setEndDate',
        'duration_millis' => 'setDurationMillis',
        'size_in_bytes' => 'setSizeInBytes',
        'size_in_mb' => 'setSizeInMb',
        'volume_path' => 'setVolumePath',
        'result_archive' => 'setResultArchive',
        'result_path' => 'setResultPath',
        'external_id' => 'setExternalId',
        'snapshot_id' => 'setSnapshotId',
        'snapshot_external_id' => 'setSnapshotExternalId',
        'created_by' => 'setCreatedBy',
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
        'backup' => 'getBackup',
        'backup_set_id' => 'getBackupSetId',
        'instance_id' => 'getInstanceId',
        'container_id' => 'getContainerId',
        'server_id' => 'getServerId',
        'status' => 'getStatus',
        'error_message' => 'getErrorMessage',
        'start_date' => 'getStartDate',
        'end_date' => 'getEndDate',
        'duration_millis' => 'getDurationMillis',
        'size_in_bytes' => 'getSizeInBytes',
        'size_in_mb' => 'getSizeInMb',
        'volume_path' => 'getVolumePath',
        'result_archive' => 'getResultArchive',
        'result_path' => 'getResultPath',
        'external_id' => 'getExternalId',
        'snapshot_id' => 'getSnapshotId',
        'snapshot_external_id' => 'getSnapshotExternalId',
        'created_by' => 'getCreatedBy',
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
        $this->container['backup'] = $data['backup'] ?? null;
        $this->container['backup_set_id'] = $data['backup_set_id'] ?? null;
        $this->container['instance_id'] = $data['instance_id'] ?? null;
        $this->container['container_id'] = $data['container_id'] ?? null;
        $this->container['server_id'] = $data['server_id'] ?? null;
        $this->container['status'] = $data['status'] ?? null;
        $this->container['error_message'] = $data['error_message'] ?? null;
        $this->container['start_date'] = $data['start_date'] ?? null;
        $this->container['end_date'] = $data['end_date'] ?? null;
        $this->container['duration_millis'] = $data['duration_millis'] ?? null;
        $this->container['size_in_bytes'] = $data['size_in_bytes'] ?? null;
        $this->container['size_in_mb'] = $data['size_in_mb'] ?? null;
        $this->container['volume_path'] = $data['volume_path'] ?? null;
        $this->container['result_archive'] = $data['result_archive'] ?? null;
        $this->container['result_path'] = $data['result_path'] ?? null;
        $this->container['external_id'] = $data['external_id'] ?? null;
        $this->container['snapshot_id'] = $data['snapshot_id'] ?? null;
        $this->container['snapshot_external_id'] = $data['snapshot_external_id'] ?? null;
        $this->container['created_by'] = $data['created_by'] ?? null;
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
     * Gets backup
     *
     * @return \OpenAPI\Client\Model\BackupJobBackups|null
     */
    public function getBackup()
    {
        return $this->container['backup'];
    }

    /**
     * Sets backup
     *
     * @param \OpenAPI\Client\Model\BackupJobBackups|null $backup backup
     *
     * @return self
     */
    public function setBackup($backup)
    {
        $this->container['backup'] = $backup;

        return $this;
    }

    /**
     * Gets backup_set_id
     *
     * @return string|null
     */
    public function getBackupSetId()
    {
        return $this->container['backup_set_id'];
    }

    /**
     * Sets backup_set_id
     *
     * @param string|null $backup_set_id backup_set_id
     *
     * @return self
     */
    public function setBackupSetId($backup_set_id)
    {
        $this->container['backup_set_id'] = $backup_set_id;

        return $this;
    }

    /**
     * Gets instance_id
     *
     * @return int|null
     */
    public function getInstanceId()
    {
        return $this->container['instance_id'];
    }

    /**
     * Sets instance_id
     *
     * @param int|null $instance_id instance_id
     *
     * @return self
     */
    public function setInstanceId($instance_id)
    {
        $this->container['instance_id'] = $instance_id;

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
     * Gets server_id
     *
     * @return int|null
     */
    public function getServerId()
    {
        return $this->container['server_id'];
    }

    /**
     * Sets server_id
     *
     * @param int|null $server_id server_id
     *
     * @return self
     */
    public function setServerId($server_id)
    {
        $this->container['server_id'] = $server_id;

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
     * Gets size_in_bytes
     *
     * @return int|null
     */
    public function getSizeInBytes()
    {
        return $this->container['size_in_bytes'];
    }

    /**
     * Sets size_in_bytes
     *
     * @param int|null $size_in_bytes size_in_bytes
     *
     * @return self
     */
    public function setSizeInBytes($size_in_bytes)
    {
        $this->container['size_in_bytes'] = $size_in_bytes;

        return $this;
    }

    /**
     * Gets size_in_mb
     *
     * @return int|null
     */
    public function getSizeInMb()
    {
        return $this->container['size_in_mb'];
    }

    /**
     * Sets size_in_mb
     *
     * @param int|null $size_in_mb size_in_mb
     *
     * @return self
     */
    public function setSizeInMb($size_in_mb)
    {
        $this->container['size_in_mb'] = $size_in_mb;

        return $this;
    }

    /**
     * Gets volume_path
     *
     * @return string|null
     */
    public function getVolumePath()
    {
        return $this->container['volume_path'];
    }

    /**
     * Sets volume_path
     *
     * @param string|null $volume_path volume_path
     *
     * @return self
     */
    public function setVolumePath($volume_path)
    {
        $this->container['volume_path'] = $volume_path;

        return $this;
    }

    /**
     * Gets result_archive
     *
     * @return string|null
     */
    public function getResultArchive()
    {
        return $this->container['result_archive'];
    }

    /**
     * Sets result_archive
     *
     * @param string|null $result_archive result_archive
     *
     * @return self
     */
    public function setResultArchive($result_archive)
    {
        $this->container['result_archive'] = $result_archive;

        return $this;
    }

    /**
     * Gets result_path
     *
     * @return string|null
     */
    public function getResultPath()
    {
        return $this->container['result_path'];
    }

    /**
     * Sets result_path
     *
     * @param string|null $result_path result_path
     *
     * @return self
     */
    public function setResultPath($result_path)
    {
        $this->container['result_path'] = $result_path;

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
     * Gets snapshot_id
     *
     * @return string|null
     */
    public function getSnapshotId()
    {
        return $this->container['snapshot_id'];
    }

    /**
     * Sets snapshot_id
     *
     * @param string|null $snapshot_id snapshot_id
     *
     * @return self
     */
    public function setSnapshotId($snapshot_id)
    {
        $this->container['snapshot_id'] = $snapshot_id;

        return $this;
    }

    /**
     * Gets snapshot_external_id
     *
     * @return string|null
     */
    public function getSnapshotExternalId()
    {
        return $this->container['snapshot_external_id'];
    }

    /**
     * Sets snapshot_external_id
     *
     * @param string|null $snapshot_external_id snapshot_external_id
     *
     * @return self
     */
    public function setSnapshotExternalId($snapshot_external_id)
    {
        $this->container['snapshot_external_id'] = $snapshot_external_id;

        return $this;
    }

    /**
     * Gets created_by
     *
     * @return \OpenAPI\Client\Model\InlineResponse200108NetworkFloatingIpCreatedBy|null
     */
    public function getCreatedBy()
    {
        return $this->container['created_by'];
    }

    /**
     * Sets created_by
     *
     * @param \OpenAPI\Client\Model\InlineResponse200108NetworkFloatingIpCreatedBy|null $created_by created_by
     *
     * @return self
     */
    public function setCreatedBy($created_by)
    {
        $this->container['created_by'] = $created_by;

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


