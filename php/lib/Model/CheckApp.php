<?php
/**
 * CheckApp
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
 * CheckApp Class Doc Comment
 *
 * @category Class
 * @package  OpenAPI\Client
 * @author   OpenAPI Generator team
 * @link     https://openapi-generator.tech
 * @implements \ArrayAccess<TKey, TValue>
 * @template TKey int|null
 * @template TValue mixed|null  
 */
class CheckApp implements ModelInterface, ArrayAccess, \JsonSerializable
{
    public const DISCRIMINATOR = null;

    /**
      * The original name of the model.
      *
      * @var string
      */
    protected static $openAPIModelName = 'checkApp';

    /**
      * Array of property to type mappings. Used for (de)serialization
      *
      * @var string[]
      */
    protected static $openAPITypes = [
        'id' => 'int',
        'account' => '\OpenAPI\Client\Model\ApiBlueprintsIdUpdatePermissionsResourcePermissionSites',
        'active' => 'bool',
        'app' => '\OpenAPI\Client\Model\InlineResponse20082LoadBalancerInstanceSslCert',
        'name' => 'string',
        'description' => 'string',
        'in_uptime' => 'bool',
        'last_check_status' => 'string',
        'last_warning_date' => '\DateTime',
        'last_error_date' => '\DateTime',
        'last_success_date' => '\DateTime',
        'last_run_date' => '\DateTime',
        'last_error' => 'string',
        'last_timer' => 'int',
        'health' => 'int',
        'history' => 'string',
        'severity' => 'string',
        'create_incident' => 'bool',
        'muted' => 'bool',
        'created_by' => '\OpenAPI\Client\Model\InlineResponse200107NetworkPoolCreatedBy',
        'date_created' => '\DateTime',
        'last_updated' => '\DateTime',
        'availability' => 'string',
        'checks' => 'int[]',
        'check_groups' => 'int[]'
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
        'active' => null,
        'app' => null,
        'name' => null,
        'description' => null,
        'in_uptime' => null,
        'last_check_status' => null,
        'last_warning_date' => 'date-time',
        'last_error_date' => 'date-time',
        'last_success_date' => 'date-time',
        'last_run_date' => 'date-time',
        'last_error' => null,
        'last_timer' => 'int64',
        'health' => 'int64',
        'history' => null,
        'severity' => null,
        'create_incident' => null,
        'muted' => null,
        'created_by' => null,
        'date_created' => 'date-time',
        'last_updated' => 'date-time',
        'availability' => null,
        'checks' => 'int64',
        'check_groups' => 'int64'
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
        'active' => 'active',
        'app' => 'app',
        'name' => 'name',
        'description' => 'description',
        'in_uptime' => 'inUptime',
        'last_check_status' => 'lastCheckStatus',
        'last_warning_date' => 'lastWarningDate',
        'last_error_date' => 'lastErrorDate',
        'last_success_date' => 'lastSuccessDate',
        'last_run_date' => 'lastRunDate',
        'last_error' => 'lastError',
        'last_timer' => 'lastTimer',
        'health' => 'health',
        'history' => 'history',
        'severity' => 'severity',
        'create_incident' => 'createIncident',
        'muted' => 'muted',
        'created_by' => 'createdBy',
        'date_created' => 'dateCreated',
        'last_updated' => 'lastUpdated',
        'availability' => 'availability',
        'checks' => 'checks',
        'check_groups' => 'checkGroups'
    ];

    /**
     * Array of attributes to setter functions (for deserialization of responses)
     *
     * @var string[]
     */
    protected static $setters = [
        'id' => 'setId',
        'account' => 'setAccount',
        'active' => 'setActive',
        'app' => 'setApp',
        'name' => 'setName',
        'description' => 'setDescription',
        'in_uptime' => 'setInUptime',
        'last_check_status' => 'setLastCheckStatus',
        'last_warning_date' => 'setLastWarningDate',
        'last_error_date' => 'setLastErrorDate',
        'last_success_date' => 'setLastSuccessDate',
        'last_run_date' => 'setLastRunDate',
        'last_error' => 'setLastError',
        'last_timer' => 'setLastTimer',
        'health' => 'setHealth',
        'history' => 'setHistory',
        'severity' => 'setSeverity',
        'create_incident' => 'setCreateIncident',
        'muted' => 'setMuted',
        'created_by' => 'setCreatedBy',
        'date_created' => 'setDateCreated',
        'last_updated' => 'setLastUpdated',
        'availability' => 'setAvailability',
        'checks' => 'setChecks',
        'check_groups' => 'setCheckGroups'
    ];

    /**
     * Array of attributes to getter functions (for serialization of requests)
     *
     * @var string[]
     */
    protected static $getters = [
        'id' => 'getId',
        'account' => 'getAccount',
        'active' => 'getActive',
        'app' => 'getApp',
        'name' => 'getName',
        'description' => 'getDescription',
        'in_uptime' => 'getInUptime',
        'last_check_status' => 'getLastCheckStatus',
        'last_warning_date' => 'getLastWarningDate',
        'last_error_date' => 'getLastErrorDate',
        'last_success_date' => 'getLastSuccessDate',
        'last_run_date' => 'getLastRunDate',
        'last_error' => 'getLastError',
        'last_timer' => 'getLastTimer',
        'health' => 'getHealth',
        'history' => 'getHistory',
        'severity' => 'getSeverity',
        'create_incident' => 'getCreateIncident',
        'muted' => 'getMuted',
        'created_by' => 'getCreatedBy',
        'date_created' => 'getDateCreated',
        'last_updated' => 'getLastUpdated',
        'availability' => 'getAvailability',
        'checks' => 'getChecks',
        'check_groups' => 'getCheckGroups'
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
        $this->container['active'] = $data['active'] ?? null;
        $this->container['app'] = $data['app'] ?? null;
        $this->container['name'] = $data['name'] ?? null;
        $this->container['description'] = $data['description'] ?? null;
        $this->container['in_uptime'] = $data['in_uptime'] ?? null;
        $this->container['last_check_status'] = $data['last_check_status'] ?? null;
        $this->container['last_warning_date'] = $data['last_warning_date'] ?? null;
        $this->container['last_error_date'] = $data['last_error_date'] ?? null;
        $this->container['last_success_date'] = $data['last_success_date'] ?? null;
        $this->container['last_run_date'] = $data['last_run_date'] ?? null;
        $this->container['last_error'] = $data['last_error'] ?? null;
        $this->container['last_timer'] = $data['last_timer'] ?? null;
        $this->container['health'] = $data['health'] ?? null;
        $this->container['history'] = $data['history'] ?? null;
        $this->container['severity'] = $data['severity'] ?? null;
        $this->container['create_incident'] = $data['create_incident'] ?? null;
        $this->container['muted'] = $data['muted'] ?? null;
        $this->container['created_by'] = $data['created_by'] ?? null;
        $this->container['date_created'] = $data['date_created'] ?? null;
        $this->container['last_updated'] = $data['last_updated'] ?? null;
        $this->container['availability'] = $data['availability'] ?? null;
        $this->container['checks'] = $data['checks'] ?? null;
        $this->container['check_groups'] = $data['check_groups'] ?? null;
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
     * Gets active
     *
     * @return bool|null
     */
    public function getActive()
    {
        return $this->container['active'];
    }

    /**
     * Sets active
     *
     * @param bool|null $active active
     *
     * @return self
     */
    public function setActive($active)
    {
        $this->container['active'] = $active;

        return $this;
    }

    /**
     * Gets app
     *
     * @return \OpenAPI\Client\Model\InlineResponse20082LoadBalancerInstanceSslCert|null
     */
    public function getApp()
    {
        return $this->container['app'];
    }

    /**
     * Sets app
     *
     * @param \OpenAPI\Client\Model\InlineResponse20082LoadBalancerInstanceSslCert|null $app app
     *
     * @return self
     */
    public function setApp($app)
    {
        $this->container['app'] = $app;

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
     * Gets last_check_status
     *
     * @return string|null
     */
    public function getLastCheckStatus()
    {
        return $this->container['last_check_status'];
    }

    /**
     * Sets last_check_status
     *
     * @param string|null $last_check_status last_check_status
     *
     * @return self
     */
    public function setLastCheckStatus($last_check_status)
    {
        $this->container['last_check_status'] = $last_check_status;

        return $this;
    }

    /**
     * Gets last_warning_date
     *
     * @return \DateTime|null
     */
    public function getLastWarningDate()
    {
        return $this->container['last_warning_date'];
    }

    /**
     * Sets last_warning_date
     *
     * @param \DateTime|null $last_warning_date last_warning_date
     *
     * @return self
     */
    public function setLastWarningDate($last_warning_date)
    {
        $this->container['last_warning_date'] = $last_warning_date;

        return $this;
    }

    /**
     * Gets last_error_date
     *
     * @return \DateTime|null
     */
    public function getLastErrorDate()
    {
        return $this->container['last_error_date'];
    }

    /**
     * Sets last_error_date
     *
     * @param \DateTime|null $last_error_date last_error_date
     *
     * @return self
     */
    public function setLastErrorDate($last_error_date)
    {
        $this->container['last_error_date'] = $last_error_date;

        return $this;
    }

    /**
     * Gets last_success_date
     *
     * @return \DateTime|null
     */
    public function getLastSuccessDate()
    {
        return $this->container['last_success_date'];
    }

    /**
     * Sets last_success_date
     *
     * @param \DateTime|null $last_success_date last_success_date
     *
     * @return self
     */
    public function setLastSuccessDate($last_success_date)
    {
        $this->container['last_success_date'] = $last_success_date;

        return $this;
    }

    /**
     * Gets last_run_date
     *
     * @return \DateTime|null
     */
    public function getLastRunDate()
    {
        return $this->container['last_run_date'];
    }

    /**
     * Sets last_run_date
     *
     * @param \DateTime|null $last_run_date last_run_date
     *
     * @return self
     */
    public function setLastRunDate($last_run_date)
    {
        $this->container['last_run_date'] = $last_run_date;

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
     * Gets last_timer
     *
     * @return int|null
     */
    public function getLastTimer()
    {
        return $this->container['last_timer'];
    }

    /**
     * Sets last_timer
     *
     * @param int|null $last_timer last_timer
     *
     * @return self
     */
    public function setLastTimer($last_timer)
    {
        $this->container['last_timer'] = $last_timer;

        return $this;
    }

    /**
     * Gets health
     *
     * @return int|null
     */
    public function getHealth()
    {
        return $this->container['health'];
    }

    /**
     * Sets health
     *
     * @param int|null $health health
     *
     * @return self
     */
    public function setHealth($health)
    {
        $this->container['health'] = $health;

        return $this;
    }

    /**
     * Gets history
     *
     * @return string|null
     */
    public function getHistory()
    {
        return $this->container['history'];
    }

    /**
     * Sets history
     *
     * @param string|null $history history
     *
     * @return self
     */
    public function setHistory($history)
    {
        $this->container['history'] = $history;

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
     * Gets create_incident
     *
     * @return bool|null
     */
    public function getCreateIncident()
    {
        return $this->container['create_incident'];
    }

    /**
     * Sets create_incident
     *
     * @param bool|null $create_incident create_incident
     *
     * @return self
     */
    public function setCreateIncident($create_incident)
    {
        $this->container['create_incident'] = $create_incident;

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
     * Gets created_by
     *
     * @return \OpenAPI\Client\Model\InlineResponse200107NetworkPoolCreatedBy|null
     */
    public function getCreatedBy()
    {
        return $this->container['created_by'];
    }

    /**
     * Sets created_by
     *
     * @param \OpenAPI\Client\Model\InlineResponse200107NetworkPoolCreatedBy|null $created_by created_by
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
     * Gets availability
     *
     * @return string|null
     */
    public function getAvailability()
    {
        return $this->container['availability'];
    }

    /**
     * Sets availability
     *
     * @param string|null $availability availability
     *
     * @return self
     */
    public function setAvailability($availability)
    {
        $this->container['availability'] = $availability;

        return $this;
    }

    /**
     * Gets checks
     *
     * @return int[]|null
     */
    public function getChecks()
    {
        return $this->container['checks'];
    }

    /**
     * Sets checks
     *
     * @param int[]|null $checks checks
     *
     * @return self
     */
    public function setChecks($checks)
    {
        $this->container['checks'] = $checks;

        return $this;
    }

    /**
     * Gets check_groups
     *
     * @return int[]|null
     */
    public function getCheckGroups()
    {
        return $this->container['check_groups'];
    }

    /**
     * Sets check_groups
     *
     * @param int[]|null $check_groups check_groups
     *
     * @return self
     */
    public function setCheckGroups($check_groups)
    {
        $this->container['check_groups'] = $check_groups;

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


