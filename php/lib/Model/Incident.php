<?php
/**
 * Incident
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
 * Incident Class Doc Comment
 *
 * @category Class
 * @package  OpenAPI\Client
 * @author   OpenAPI Generator team
 * @link     https://openapi-generator.tech
 * @implements \ArrayAccess<TKey, TValue>
 * @template TKey int|null
 * @template TValue mixed|null  
 */
class Incident implements ModelInterface, ArrayAccess, \JsonSerializable
{
    public const DISCRIMINATOR = null;

    /**
      * The original name of the model.
      *
      * @var string
      */
    protected static $openAPIModelName = 'incident';

    /**
      * Array of property to type mappings. Used for (de)serialization
      *
      * @var string[]
      */
    protected static $openAPITypes = [
        'id' => 'int',
        'account' => '\OpenAPI\Client\Model\ApiBlueprintsIdUpdatePermissionsResourcePermissionSites',
        'app' => 'string',
        'auto_close' => 'bool',
        'channel_id' => 'string',
        'check_groups' => '\OpenAPI\Client\Model\InlineResponse20040AppDeployInstance[]',
        'checks' => '\OpenAPI\Client\Model\Check[]',
        'comment' => 'string',
        'display_name' => 'string',
        'duration' => 'string',
        'end_date' => '\DateTime',
        'in_uptime' => 'bool',
        'muted' => 'bool',
        'last_check_time' => '\DateTime',
        'last_error' => 'string',
        'last_message' => 'string',
        'name' => 'string',
        'resolution' => 'string',
        'severity' => 'string',
        'severity_id' => 'int',
        'start_date' => '\DateTime',
        'status' => 'string',
        'visibility' => 'string'
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
        'account' => null,
        'app' => null,
        'auto_close' => null,
        'channel_id' => null,
        'check_groups' => null,
        'checks' => null,
        'comment' => null,
        'display_name' => null,
        'duration' => null,
        'end_date' => 'date-time',
        'in_uptime' => null,
        'muted' => null,
        'last_check_time' => 'date-time',
        'last_error' => null,
        'last_message' => null,
        'name' => null,
        'resolution' => null,
        'severity' => null,
        'severity_id' => 'int64',
        'start_date' => 'date-time',
        'status' => null,
        'visibility' => null
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
        'account' => 'account',
        'app' => 'app',
        'auto_close' => 'autoClose',
        'channel_id' => 'channelId',
        'check_groups' => 'checkGroups',
        'checks' => 'checks',
        'comment' => 'comment',
        'display_name' => 'displayName',
        'duration' => 'duration',
        'end_date' => 'endDate',
        'in_uptime' => 'inUptime',
        'muted' => 'muted',
        'last_check_time' => 'lastCheckTime',
        'last_error' => 'lastError',
        'last_message' => 'lastMessage',
        'name' => 'name',
        'resolution' => 'resolution',
        'severity' => 'severity',
        'severity_id' => 'severityId',
        'start_date' => 'startDate',
        'status' => 'status',
        'visibility' => 'visibility'
    ];

    /**
     * Array of attributes to setter functions (for deserialization of responses)
     *
     * @var string[]
     */
    protected static $setters = [
        'id' => 'setId',
        'account' => 'setAccount',
        'app' => 'setApp',
        'auto_close' => 'setAutoClose',
        'channel_id' => 'setChannelId',
        'check_groups' => 'setCheckGroups',
        'checks' => 'setChecks',
        'comment' => 'setComment',
        'display_name' => 'setDisplayName',
        'duration' => 'setDuration',
        'end_date' => 'setEndDate',
        'in_uptime' => 'setInUptime',
        'muted' => 'setMuted',
        'last_check_time' => 'setLastCheckTime',
        'last_error' => 'setLastError',
        'last_message' => 'setLastMessage',
        'name' => 'setName',
        'resolution' => 'setResolution',
        'severity' => 'setSeverity',
        'severity_id' => 'setSeverityId',
        'start_date' => 'setStartDate',
        'status' => 'setStatus',
        'visibility' => 'setVisibility'
    ];

    /**
     * Array of attributes to getter functions (for serialization of requests)
     *
     * @var string[]
     */
    protected static $getters = [
        'id' => 'getId',
        'account' => 'getAccount',
        'app' => 'getApp',
        'auto_close' => 'getAutoClose',
        'channel_id' => 'getChannelId',
        'check_groups' => 'getCheckGroups',
        'checks' => 'getChecks',
        'comment' => 'getComment',
        'display_name' => 'getDisplayName',
        'duration' => 'getDuration',
        'end_date' => 'getEndDate',
        'in_uptime' => 'getInUptime',
        'muted' => 'getMuted',
        'last_check_time' => 'getLastCheckTime',
        'last_error' => 'getLastError',
        'last_message' => 'getLastMessage',
        'name' => 'getName',
        'resolution' => 'getResolution',
        'severity' => 'getSeverity',
        'severity_id' => 'getSeverityId',
        'start_date' => 'getStartDate',
        'status' => 'getStatus',
        'visibility' => 'getVisibility'
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
        $this->container['account'] = $data['account'] ?? null;
        $this->container['app'] = $data['app'] ?? null;
        $this->container['auto_close'] = $data['auto_close'] ?? null;
        $this->container['channel_id'] = $data['channel_id'] ?? null;
        $this->container['check_groups'] = $data['check_groups'] ?? null;
        $this->container['checks'] = $data['checks'] ?? null;
        $this->container['comment'] = $data['comment'] ?? null;
        $this->container['display_name'] = $data['display_name'] ?? null;
        $this->container['duration'] = $data['duration'] ?? null;
        $this->container['end_date'] = $data['end_date'] ?? null;
        $this->container['in_uptime'] = $data['in_uptime'] ?? null;
        $this->container['muted'] = $data['muted'] ?? null;
        $this->container['last_check_time'] = $data['last_check_time'] ?? null;
        $this->container['last_error'] = $data['last_error'] ?? null;
        $this->container['last_message'] = $data['last_message'] ?? null;
        $this->container['name'] = $data['name'] ?? null;
        $this->container['resolution'] = $data['resolution'] ?? null;
        $this->container['severity'] = $data['severity'] ?? null;
        $this->container['severity_id'] = $data['severity_id'] ?? null;
        $this->container['start_date'] = $data['start_date'] ?? null;
        $this->container['status'] = $data['status'] ?? null;
        $this->container['visibility'] = $data['visibility'] ?? null;
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
     * Gets account
     *
     * @return \OpenAPI\Client\Model\ApiBlueprintsIdUpdatePermissionsResourcePermissionSites|null
     */
    public function getAccount()
    {
        return $this->container['account'];
    }

    /**
     * Sets account
     *
     * @param \OpenAPI\Client\Model\ApiBlueprintsIdUpdatePermissionsResourcePermissionSites|null $account account
     *
     * @return self
     */
    public function setAccount($account)
    {
        $this->container['account'] = $account;

        return $this;
    }

    /**
     * Gets app
     *
     * @return string|null
     */
    public function getApp()
    {
        return $this->container['app'];
    }

    /**
     * Sets app
     *
     * @param string|null $app app
     *
     * @return self
     */
    public function setApp($app)
    {
        $this->container['app'] = $app;

        return $this;
    }

    /**
     * Gets auto_close
     *
     * @return bool|null
     */
    public function getAutoClose()
    {
        return $this->container['auto_close'];
    }

    /**
     * Sets auto_close
     *
     * @param bool|null $auto_close auto_close
     *
     * @return self
     */
    public function setAutoClose($auto_close)
    {
        $this->container['auto_close'] = $auto_close;

        return $this;
    }

    /**
     * Gets channel_id
     *
     * @return string|null
     */
    public function getChannelId()
    {
        return $this->container['channel_id'];
    }

    /**
     * Sets channel_id
     *
     * @param string|null $channel_id channel_id
     *
     * @return self
     */
    public function setChannelId($channel_id)
    {
        $this->container['channel_id'] = $channel_id;

        return $this;
    }

    /**
     * Gets check_groups
     *
     * @return \OpenAPI\Client\Model\InlineResponse20040AppDeployInstance[]|null
     */
    public function getCheckGroups()
    {
        return $this->container['check_groups'];
    }

    /**
     * Sets check_groups
     *
     * @param \OpenAPI\Client\Model\InlineResponse20040AppDeployInstance[]|null $check_groups check_groups
     *
     * @return self
     */
    public function setCheckGroups($check_groups)
    {
        $this->container['check_groups'] = $check_groups;

        return $this;
    }

    /**
     * Gets checks
     *
     * @return \OpenAPI\Client\Model\Check[]|null
     */
    public function getChecks()
    {
        return $this->container['checks'];
    }

    /**
     * Sets checks
     *
     * @param \OpenAPI\Client\Model\Check[]|null $checks checks
     *
     * @return self
     */
    public function setChecks($checks)
    {
        $this->container['checks'] = $checks;

        return $this;
    }

    /**
     * Gets comment
     *
     * @return string|null
     */
    public function getComment()
    {
        return $this->container['comment'];
    }

    /**
     * Sets comment
     *
     * @param string|null $comment comment
     *
     * @return self
     */
    public function setComment($comment)
    {
        $this->container['comment'] = $comment;

        return $this;
    }

    /**
     * Gets display_name
     *
     * @return string|null
     */
    public function getDisplayName()
    {
        return $this->container['display_name'];
    }

    /**
     * Sets display_name
     *
     * @param string|null $display_name display_name
     *
     * @return self
     */
    public function setDisplayName($display_name)
    {
        $this->container['display_name'] = $display_name;

        return $this;
    }

    /**
     * Gets duration
     *
     * @return string|null
     */
    public function getDuration()
    {
        return $this->container['duration'];
    }

    /**
     * Sets duration
     *
     * @param string|null $duration duration
     *
     * @return self
     */
    public function setDuration($duration)
    {
        $this->container['duration'] = $duration;

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
     * Gets in_uptime
     *
     * @return bool|null
     */
    public function getInUptime()
    {
        return $this->container['in_uptime'];
    }

    /**
     * Sets in_uptime
     *
     * @param bool|null $in_uptime in_uptime
     *
     * @return self
     */
    public function setInUptime($in_uptime)
    {
        $this->container['in_uptime'] = $in_uptime;

        return $this;
    }

    /**
     * Gets muted
     *
     * @return bool|null
     */
    public function getMuted()
    {
        return $this->container['muted'];
    }

    /**
     * Sets muted
     *
     * @param bool|null $muted muted
     *
     * @return self
     */
    public function setMuted($muted)
    {
        $this->container['muted'] = $muted;

        return $this;
    }

    /**
     * Gets last_check_time
     *
     * @return \DateTime|null
     */
    public function getLastCheckTime()
    {
        return $this->container['last_check_time'];
    }

    /**
     * Sets last_check_time
     *
     * @param \DateTime|null $last_check_time last_check_time
     *
     * @return self
     */
    public function setLastCheckTime($last_check_time)
    {
        $this->container['last_check_time'] = $last_check_time;

        return $this;
    }

    /**
     * Gets last_error
     *
     * @return string|null
     */
    public function getLastError()
    {
        return $this->container['last_error'];
    }

    /**
     * Sets last_error
     *
     * @param string|null $last_error last_error
     *
     * @return self
     */
    public function setLastError($last_error)
    {
        $this->container['last_error'] = $last_error;

        return $this;
    }

    /**
     * Gets last_message
     *
     * @return string|null
     */
    public function getLastMessage()
    {
        return $this->container['last_message'];
    }

    /**
     * Sets last_message
     *
     * @param string|null $last_message last_message
     *
     * @return self
     */
    public function setLastMessage($last_message)
    {
        $this->container['last_message'] = $last_message;

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
     * Gets resolution
     *
     * @return string|null
     */
    public function getResolution()
    {
        return $this->container['resolution'];
    }

    /**
     * Sets resolution
     *
     * @param string|null $resolution resolution
     *
     * @return self
     */
    public function setResolution($resolution)
    {
        $this->container['resolution'] = $resolution;

        return $this;
    }

    /**
     * Gets severity
     *
     * @return string|null
     */
    public function getSeverity()
    {
        return $this->container['severity'];
    }

    /**
     * Sets severity
     *
     * @param string|null $severity severity
     *
     * @return self
     */
    public function setSeverity($severity)
    {
        $this->container['severity'] = $severity;

        return $this;
    }

    /**
     * Gets severity_id
     *
     * @return int|null
     */
    public function getSeverityId()
    {
        return $this->container['severity_id'];
    }

    /**
     * Sets severity_id
     *
     * @param int|null $severity_id severity_id
     *
     * @return self
     */
    public function setSeverityId($severity_id)
    {
        $this->container['severity_id'] = $severity_id;

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
     * Gets visibility
     *
     * @return string|null
     */
    public function getVisibility()
    {
        return $this->container['visibility'];
    }

    /**
     * Sets visibility
     *
     * @param string|null $visibility visibility
     *
     * @return self
     */
    public function setVisibility($visibility)
    {
        $this->container['visibility'] = $visibility;

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


