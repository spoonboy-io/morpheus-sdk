<?php
/**
 * BackupTypeInstance
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
 * BackupTypeInstance Class Doc Comment
 *
 * @category Class
 * @package  OpenAPI\Client
 * @author   OpenAPI Generator team
 * @link     https://openapi-generator.tech
 * @implements \ArrayAccess<TKey, TValue>
 * @template TKey int|null
 * @template TValue mixed|null  
 */
class BackupTypeInstance implements ModelInterface, ArrayAccess, \JsonSerializable
{
    public const DISCRIMINATOR = null;

    /**
      * The original name of the model.
      *
      * @var string
      */
    protected static $openAPIModelName = 'backupTypeInstance';

    /**
      * Array of property to type mappings. Used for (de)serialization
      *
      * @var string[]
      */
    protected static $openAPITypes = [
        'location_type' => 'string',
        'name' => 'string',
        'instance_id' => 'int',
        'container_id' => 'int',
        'backup_type' => 'string',
        'job_action' => 'string',
        'job_id' => 'int',
        'job_name' => 'string',
        'job_schedule' => 'int',
        'retention_count' => 'int'
    ];

    /**
      * Array of property to format mappings. Used for (de)serialization
      *
      * @var string[]
      * @phpstan-var array<string, string|null>
      * @psalm-var array<string, string|null>
      */
    protected static $openAPIFormats = [
        'location_type' => null,
        'name' => null,
        'instance_id' => 'int64',
        'container_id' => 'int64',
        'backup_type' => null,
        'job_action' => null,
        'job_id' => 'int64',
        'job_name' => null,
        'job_schedule' => 'int64',
        'retention_count' => 'int64'
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
        'location_type' => 'locationType',
        'name' => 'name',
        'instance_id' => 'instanceId',
        'container_id' => 'containerId',
        'backup_type' => 'backupType',
        'job_action' => 'jobAction',
        'job_id' => 'jobId',
        'job_name' => 'jobName',
        'job_schedule' => 'jobSchedule',
        'retention_count' => 'retentionCount'
    ];

    /**
     * Array of attributes to setter functions (for deserialization of responses)
     *
     * @var string[]
     */
    protected static $setters = [
        'location_type' => 'setLocationType',
        'name' => 'setName',
        'instance_id' => 'setInstanceId',
        'container_id' => 'setContainerId',
        'backup_type' => 'setBackupType',
        'job_action' => 'setJobAction',
        'job_id' => 'setJobId',
        'job_name' => 'setJobName',
        'job_schedule' => 'setJobSchedule',
        'retention_count' => 'setRetentionCount'
    ];

    /**
     * Array of attributes to getter functions (for serialization of requests)
     *
     * @var string[]
     */
    protected static $getters = [
        'location_type' => 'getLocationType',
        'name' => 'getName',
        'instance_id' => 'getInstanceId',
        'container_id' => 'getContainerId',
        'backup_type' => 'getBackupType',
        'job_action' => 'getJobAction',
        'job_id' => 'getJobId',
        'job_name' => 'getJobName',
        'job_schedule' => 'getJobSchedule',
        'retention_count' => 'getRetentionCount'
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

    const LOCATION_TYPE_INSTANCE = 'instance';
    const BACKUP_TYPE_ALIBABA_SNAPSHOT = 'alibabaSnapshot';
    const BACKUP_TYPE_AMAZON_SNAPSHOT = 'amazonSnapshot';
    const BACKUP_TYPE_AVAMAR_VM_WARE_BACKUP = 'avamarVMWareBackup';
    const BACKUP_TYPE_AZURE_SNAPSHOT = 'azureSnapshot';
    const BACKUP_TYPE_BLUEMIX_SNAPSHOT = 'bluemixSnapshot';
    const BACKUP_TYPE_COMMVAULT_FILE_BACKUP = 'commvaultFileBackup';
    const BACKUP_TYPE_COMMVAULT_OPENSTACK_BACKUP = 'commvaultOpenstackBackup';
    const BACKUP_TYPE_COMMVAULT_VM_WARE_BACKUP = 'commvaultVMWareBackup';
    const BACKUP_TYPE_DIGITALOCEAN_SNAPSHOT = 'digitaloceanSnapshot';
    const BACKUP_TYPE_DIRECTORY_BACKUP = 'directoryBackup';
    const BACKUP_TYPE_ESXI_SNAPSHOT = 'esxiSnapshot';
    const BACKUP_TYPE_FILE_BACKUP = 'fileBackup';
    const BACKUP_TYPE_FUSION_SNAPSHOT = 'fusionSnapshot';
    const BACKUP_TYPE_GOOGLE_SNAPSHOT = 'googleSnapshot';
    const BACKUP_TYPE_HUAWEI_SNAPSHOT = 'huaweiSnapshot';
    const BACKUP_TYPE_HYPERV_SNAPSHOT = 'hypervSnapshot';
    const BACKUP_TYPE_KVM = 'kvm';
    const BACKUP_TYPE_LVM_IMAGE = 'lvmImage';
    const BACKUP_TYPE_LVM_MIGRATION = 'lvmMigration';
    const BACKUP_TYPE_LVM_SNAPSHOT = 'lvmSnapshot';
    const BACKUP_TYPE_MONGO_DB = 'MongoDB';
    const BACKUP_TYPE_MORPHEUS_APPLIANCE = 'morpheusAppliance';
    const BACKUP_TYPE_MORPHEUS_CONTAINER_BACKUP = 'morpheusContainerBackup';
    const BACKUP_TYPE_MORPHEUS_STORAGE_BACKUP = 'morpheusStorageBackup';
    const BACKUP_TYPE_MORPHEUS_VM_BACKUP = 'morpheusVmBackup';
    const BACKUP_TYPE_MY_SQL = 'MySQL';
    const BACKUP_TYPE_NUTANIX_SNAPSHOT = 'nutanixSnapshot';
    const BACKUP_TYPE_OPENSTACK_SNAPSHOT = 'openstackSnapshot';
    const BACKUP_TYPE_OPENTELEKOM_SNAPSHOT = 'opentelekomSnapshot';
    const BACKUP_TYPE_ORACLE_CLOUD_SNAPSHOT = 'oracleCloudSnapshot';
    const BACKUP_TYPE_POSTGRES = 'Postgres';
    const BACKUP_TYPE_POWERVC_SNAPSHOT = 'powervcSnapshot';
    const BACKUP_TYPE_RUBRIK_VM_WARE_BACKUP = 'rubrikVMWareBackup';
    const BACKUP_TYPE_SCVMM_SNAPSHOT = 'scvmmSnapshot';
    const BACKUP_TYPE_SOFTLAYER_SNAPSHOT = 'softlayerSnapshot';
    const BACKUP_TYPE_SQL_SERVER = 'SqlServer';
    const BACKUP_TYPE_TAR_DIRECTORY_BACKUP = 'tarDirectoryBackup';
    const BACKUP_TYPE_UP_CLOUD_SNAPSHOT = 'upCloudSnapshot';
    const BACKUP_TYPE_VCD_SNAPSHOT = 'vcdSnapshot';
    const BACKUP_TYPE_VEEAM_HYPERV_BACKUP = 'veeamHypervBackup';
    const BACKUP_TYPE_VEEAM_SCVMM_BACKUP = 'veeamScvmmBackup';
    const BACKUP_TYPE_VEEAM_VCD_BACKUP = 'veeamVcdBackup';
    const BACKUP_TYPE_VEEAM_VM_WARE_BACKUP = 'veeamVMWareBackup';
    const BACKUP_TYPE_VIRTUSTREAM_SNAPSHOT = 'virtustreamSnapshot';
    const BACKUP_TYPE_VMWARE_SNAPSHOT = 'vmwareSnapshot';
    const BACKUP_TYPE_WIN_MIGRATION = 'winMigration';
    const BACKUP_TYPE_XEN_SNAPSHOT = 'xenSnapshot';
    const JOB_ACTION__NEW = 'new';
    const JOB_ACTION__CLONE = 'clone';
    const JOB_ACTION_ADD_TO = 'addTo';
    

    
    /**
     * Gets allowable values of the enum
     *
     * @return string[]
     */
    public function getLocationTypeAllowableValues()
    {
        return [
            self::LOCATION_TYPE_INSTANCE,
        ];
    }
    
    /**
     * Gets allowable values of the enum
     *
     * @return string[]
     */
    public function getBackupTypeAllowableValues()
    {
        return [
            self::BACKUP_TYPE_ALIBABA_SNAPSHOT,
            self::BACKUP_TYPE_AMAZON_SNAPSHOT,
            self::BACKUP_TYPE_AVAMAR_VM_WARE_BACKUP,
            self::BACKUP_TYPE_AZURE_SNAPSHOT,
            self::BACKUP_TYPE_BLUEMIX_SNAPSHOT,
            self::BACKUP_TYPE_COMMVAULT_FILE_BACKUP,
            self::BACKUP_TYPE_COMMVAULT_OPENSTACK_BACKUP,
            self::BACKUP_TYPE_COMMVAULT_VM_WARE_BACKUP,
            self::BACKUP_TYPE_DIGITALOCEAN_SNAPSHOT,
            self::BACKUP_TYPE_DIRECTORY_BACKUP,
            self::BACKUP_TYPE_ESXI_SNAPSHOT,
            self::BACKUP_TYPE_FILE_BACKUP,
            self::BACKUP_TYPE_FUSION_SNAPSHOT,
            self::BACKUP_TYPE_GOOGLE_SNAPSHOT,
            self::BACKUP_TYPE_HUAWEI_SNAPSHOT,
            self::BACKUP_TYPE_HYPERV_SNAPSHOT,
            self::BACKUP_TYPE_KVM,
            self::BACKUP_TYPE_LVM_IMAGE,
            self::BACKUP_TYPE_LVM_MIGRATION,
            self::BACKUP_TYPE_LVM_SNAPSHOT,
            self::BACKUP_TYPE_MONGO_DB,
            self::BACKUP_TYPE_MORPHEUS_APPLIANCE,
            self::BACKUP_TYPE_MORPHEUS_CONTAINER_BACKUP,
            self::BACKUP_TYPE_MORPHEUS_STORAGE_BACKUP,
            self::BACKUP_TYPE_MORPHEUS_VM_BACKUP,
            self::BACKUP_TYPE_MY_SQL,
            self::BACKUP_TYPE_NUTANIX_SNAPSHOT,
            self::BACKUP_TYPE_OPENSTACK_SNAPSHOT,
            self::BACKUP_TYPE_OPENTELEKOM_SNAPSHOT,
            self::BACKUP_TYPE_ORACLE_CLOUD_SNAPSHOT,
            self::BACKUP_TYPE_POSTGRES,
            self::BACKUP_TYPE_POWERVC_SNAPSHOT,
            self::BACKUP_TYPE_RUBRIK_VM_WARE_BACKUP,
            self::BACKUP_TYPE_SCVMM_SNAPSHOT,
            self::BACKUP_TYPE_SOFTLAYER_SNAPSHOT,
            self::BACKUP_TYPE_SQL_SERVER,
            self::BACKUP_TYPE_TAR_DIRECTORY_BACKUP,
            self::BACKUP_TYPE_UP_CLOUD_SNAPSHOT,
            self::BACKUP_TYPE_VCD_SNAPSHOT,
            self::BACKUP_TYPE_VEEAM_HYPERV_BACKUP,
            self::BACKUP_TYPE_VEEAM_SCVMM_BACKUP,
            self::BACKUP_TYPE_VEEAM_VCD_BACKUP,
            self::BACKUP_TYPE_VEEAM_VM_WARE_BACKUP,
            self::BACKUP_TYPE_VIRTUSTREAM_SNAPSHOT,
            self::BACKUP_TYPE_VMWARE_SNAPSHOT,
            self::BACKUP_TYPE_WIN_MIGRATION,
            self::BACKUP_TYPE_XEN_SNAPSHOT,
        ];
    }
    
    /**
     * Gets allowable values of the enum
     *
     * @return string[]
     */
    public function getJobActionAllowableValues()
    {
        return [
            self::JOB_ACTION__NEW,
            self::JOB_ACTION__CLONE,
            self::JOB_ACTION_ADD_TO,
        ];
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
        $this->container['location_type'] = $data['location_type'] ?? null;
        $this->container['name'] = $data['name'] ?? null;
        $this->container['instance_id'] = $data['instance_id'] ?? null;
        $this->container['container_id'] = $data['container_id'] ?? null;
        $this->container['backup_type'] = $data['backup_type'] ?? null;
        $this->container['job_action'] = $data['job_action'] ?? null;
        $this->container['job_id'] = $data['job_id'] ?? null;
        $this->container['job_name'] = $data['job_name'] ?? null;
        $this->container['job_schedule'] = $data['job_schedule'] ?? null;
        $this->container['retention_count'] = $data['retention_count'] ?? null;
    }

    /**
     * Show all the invalid properties with reasons.
     *
     * @return array invalid properties with reasons
     */
    public function listInvalidProperties()
    {
        $invalidProperties = [];

        if ($this->container['location_type'] === null) {
            $invalidProperties[] = "'location_type' can't be null";
        }
        $allowedValues = $this->getLocationTypeAllowableValues();
        if (!is_null($this->container['location_type']) && !in_array($this->container['location_type'], $allowedValues, true)) {
            $invalidProperties[] = sprintf(
                "invalid value '%s' for 'location_type', must be one of '%s'",
                $this->container['location_type'],
                implode("', '", $allowedValues)
            );
        }

        if ($this->container['name'] === null) {
            $invalidProperties[] = "'name' can't be null";
        }
        if ($this->container['instance_id'] === null) {
            $invalidProperties[] = "'instance_id' can't be null";
        }
        if ($this->container['container_id'] === null) {
            $invalidProperties[] = "'container_id' can't be null";
        }
        if ($this->container['backup_type'] === null) {
            $invalidProperties[] = "'backup_type' can't be null";
        }
        $allowedValues = $this->getBackupTypeAllowableValues();
        if (!is_null($this->container['backup_type']) && !in_array($this->container['backup_type'], $allowedValues, true)) {
            $invalidProperties[] = sprintf(
                "invalid value '%s' for 'backup_type', must be one of '%s'",
                $this->container['backup_type'],
                implode("', '", $allowedValues)
            );
        }

        if ($this->container['job_action'] === null) {
            $invalidProperties[] = "'job_action' can't be null";
        }
        $allowedValues = $this->getJobActionAllowableValues();
        if (!is_null($this->container['job_action']) && !in_array($this->container['job_action'], $allowedValues, true)) {
            $invalidProperties[] = sprintf(
                "invalid value '%s' for 'job_action', must be one of '%s'",
                $this->container['job_action'],
                implode("', '", $allowedValues)
            );
        }

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
     * Gets location_type
     *
     * @return string
     */
    public function getLocationType()
    {
        return $this->container['location_type'];
    }

    /**
     * Sets location_type
     *
     * @param string $location_type location_type
     *
     * @return self
     */
    public function setLocationType($location_type)
    {
        $allowedValues = $this->getLocationTypeAllowableValues();
        if (!in_array($location_type, $allowedValues, true)) {
            throw new \InvalidArgumentException(
                sprintf(
                    "Invalid value '%s' for 'location_type', must be one of '%s'",
                    $location_type,
                    implode("', '", $allowedValues)
                )
            );
        }
        $this->container['location_type'] = $location_type;

        return $this;
    }

    /**
     * Gets name
     *
     * @return string
     */
    public function getName()
    {
        return $this->container['name'];
    }

    /**
     * Sets name
     *
     * @param string $name A name for the backup
     *
     * @return self
     */
    public function setName($name)
    {
        $this->container['name'] = $name;

        return $this;
    }

    /**
     * Gets instance_id
     *
     * @return int
     */
    public function getInstanceId()
    {
        return $this->container['instance_id'];
    }

    /**
     * Sets instance_id
     *
     * @param int $instance_id The ID of the instance to backup
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
     * @return int
     */
    public function getContainerId()
    {
        return $this->container['container_id'];
    }

    /**
     * Sets container_id
     *
     * @param int $container_id The ID of the container to backup
     *
     * @return self
     */
    public function setContainerId($container_id)
    {
        $this->container['container_id'] = $container_id;

        return $this;
    }

    /**
     * Gets backup_type
     *
     * @return string
     */
    public function getBackupType()
    {
        return $this->container['backup_type'];
    }

    /**
     * Sets backup_type
     *
     * @param string $backup_type The backup type code, options vary by the type of cloud and source
     *
     * @return self
     */
    public function setBackupType($backup_type)
    {
        $allowedValues = $this->getBackupTypeAllowableValues();
        if (!in_array($backup_type, $allowedValues, true)) {
            throw new \InvalidArgumentException(
                sprintf(
                    "Invalid value '%s' for 'backup_type', must be one of '%s'",
                    $backup_type,
                    implode("', '", $allowedValues)
                )
            );
        }
        $this->container['backup_type'] = $backup_type;

        return $this;
    }

    /**
     * Gets job_action
     *
     * @return string
     */
    public function getJobAction()
    {
        return $this->container['job_action'];
    }

    /**
     * Sets job_action
     *
     * @param string $job_action Create a new backup job, clone an existing job or add the new backup to an existing job
     *
     * @return self
     */
    public function setJobAction($job_action)
    {
        $allowedValues = $this->getJobActionAllowableValues();
        if (!in_array($job_action, $allowedValues, true)) {
            throw new \InvalidArgumentException(
                sprintf(
                    "Invalid value '%s' for 'job_action', must be one of '%s'",
                    $job_action,
                    implode("', '", $allowedValues)
                )
            );
        }
        $this->container['job_action'] = $job_action;

        return $this;
    }

    /**
     * Gets job_id
     *
     * @return int|null
     */
    public function getJobId()
    {
        return $this->container['job_id'];
    }

    /**
     * Sets job_id
     *
     * @param int|null $job_id The ID of the job to clone or add to. Only applies to jobAction `clone` and `addTo`.
     *
     * @return self
     */
    public function setJobId($job_id)
    {
        $this->container['job_id'] = $job_id;

        return $this;
    }

    /**
     * Gets job_name
     *
     * @return string|null
     */
    public function getJobName()
    {
        return $this->container['job_name'];
    }

    /**
     * Sets job_name
     *
     * @param string|null $job_name Name for new job. Only applies to jobAction `new` and `clone`.
     *
     * @return self
     */
    public function setJobName($job_name)
    {
        $this->container['job_name'] = $job_name;

        return $this;
    }

    /**
     * Gets job_schedule
     *
     * @return int|null
     */
    public function getJobSchedule()
    {
        return $this->container['job_schedule'];
    }

    /**
     * Sets job_schedule
     *
     * @param int|null $job_schedule The ID of the execute schedule for new job. See Execute Schedules. Only applies to jobAction `new` and `clone`.
     *
     * @return self
     */
    public function setJobSchedule($job_schedule)
    {
        $this->container['job_schedule'] = $job_schedule;

        return $this;
    }

    /**
     * Gets retention_count
     *
     * @return int|null
     */
    public function getRetentionCount()
    {
        return $this->container['retention_count'];
    }

    /**
     * Sets retention_count
     *
     * @param int|null $retention_count Retention Count for new job. By default the backup settings value will be used. Only applies to jobAction `new` and `clone`.
     *
     * @return self
     */
    public function setRetentionCount($retention_count)
    {
        $this->container['retention_count'] = $retention_count;

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


