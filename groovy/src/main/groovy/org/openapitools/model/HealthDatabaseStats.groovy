package org.openapitools.model;

import groovy.transform.Canonical
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;

@Canonical
class HealthDatabaseStats {
    
    String abortedClients
    
    String abortedConnects
    
    String binlogSnapshotFile
    
    String binlogSnapshotPosition
    
    String binlogCacheDiskUse
    
    String binlogCacheUse
    
    String binlogSnapshotGtidExecuted
    
    String binlogStmtCacheDiskUse
    
    String binlogStmtCacheUse
    
    String bytesReceived
    
    String bytesSent
    
    String comAdminCommands
    
    String comAssignToKeycache
    
    String comAlterDb
    
    String comAlterDbUpgrade
    
    String comAlterEvent
    
    String comAlterFunction
    
    String comAlterInstance
    
    String comAlterProcedure
    
    String comAlterServer
    
    String comAlterTable
    
    String comAlterTablespace
    
    String comAlterUser
    
    String comAnalyze
    
    String comBegin
    
    String comBinlog
    
    String comCallProcedure
    
    String comChangeDb
    
    String comChangeMaster
    
    String comChangeReplFilter
    
    String comCheck
    
    String comChecksum
    
    String comCommit
    
    String comCreateCompressionDictionary
    
    String comCreateDb
    
    String comCreateEvent
    
    String comCreateFunction
    
    String comCreateIndex
    
    String comCreateProcedure
    
    String comCreateServer
    
    String comCreateTable
    
    String comCreateTrigger
    
    String comCreateUdf
    
    String comCreateUser
    
    String comCreateView
    
    String comDeallocSql
    
    String comDelete
    
    String comDeleteMulti
    
    String comDo
    
    String comDropCompressionDictionary
    
    String comDropDb
    
    String comDropEvent
    
    String comDropFunction
    
    String comDropIndex
    
    String comDropProcedure
    
    String comDropServer
    
    String comDropTable
    
    String comDropTrigger
    
    String comDropUser
    
    String comDropView
    
    String comEmptyQuery
    
    String comExecuteSql
    
    String comExplainOther
    
    String comFlush
    
    String comGetDiagnostics
    
    String comGrant
    
    String comHaClose
    
    String comHaOpen
    
    String comHaRead
    
    String comHelp
    
    String comInsert
    
    String comInsertSelect
    
    String comInstallPlugin
    
    String comKill
    
    String comLoad
    
    String comLockTables
    
    String comLockTablesForBackup
    
    String comLockBinlogForBackup
    
    String comOptimize
    
    String comPreloadKeys
    
    String comPrepareSql
    
    String comPurge
    
    String comPurgeBeforeDate
    
    String comReleaseSavepoint
    
    String comRenameTable
    
    String comRenameUser
    
    String comRepair
    
    String comReplace
    
    String comReplaceSelect
    
    String comReset
    
    String comResignal
    
    String comRevoke
    
    String comRevokeAll
    
    String comRollback
    
    String comRollbackToSavepoint
    
    String comSavepoint
    
    String comSelect
    
    String comSetOption
    
    String comSignal
    
    String comShowBinlogEvents
    
    String comShowBinlogs
    
    String comShowCharsets
    
    String comShowClientStatistics
    
    String comShowCollations
    
    String comShowCreateDb
    
    String comShowCreateEvent
    
    String comShowCreateFunc
    
    String comShowCreateProc
    
    String comShowCreateTable
    
    String comShowCreateTrigger
    
    String comShowDatabases
    
    String comShowEngineLogs
    
    String comShowEngineMutex
    
    String comShowEngineStatus
    
    String comShowEvents
    
    String comShowErrors
    
    String comShowFields
    
    String comShowFunctionCode
    
    String comShowFunctionStatus
    
    String comShowGrants
    
    String comShowIndexStatistics
    
    String comShowKeys
    
    String comShowMasterStatus
    
    String comShowOpenTables
    
    String comShowPlugins
    
    String comShowPrivileges
    
    String comShowProcedureCode
    
    String comShowProcedureStatus
    
    String comShowProcesslist
    
    String comShowProfile
    
    String comShowProfiles
    
    String comShowRelaylogEvents
    
    String comShowSlaveHosts
    
    String comShowSlaveStatus
    
    String comShowStatus
    
    String comShowStorageEngines
    
    String comShowTableStatistics
    
    String comShowTableStatus
    
    String comShowTables
    
    String comShowThreadStatistics
    
    String comShowTriggers
    
    String comShowUserStatistics
    
    String comShowVariables
    
    String comShowWarnings
    
    String comShowCreateUser
    
    String comShutdown
    
    String comSlaveStart
    
    String comSlaveStop
    
    String comGroupReplicationStart
    
    String comGroupReplicationStop
    
    String comStmtExecute
    
    String comStmtClose
    
    String comStmtFetch
    
    String comStmtPrepare
    
    String comStmtReset
    
    String comStmtSendLongData
    
    String comTruncate
    
    String comUninstallPlugin
    
    String comUnlockBinlog
    
    String comUnlockTables
    
    String comUpdate
    
    String comUpdateMulti
    
    String comXaCommit
    
    String comXaEnd
    
    String comXaPrepare
    
    String comXaRecover
    
    String comXaRollback
    
    String comXaStart
    
    String comStmtReprepare
    
    String compression
    
    String connectionErrorsAccept
    
    String connectionErrorsInternal
    
    String connectionErrorsMaxConnections
    
    String connectionErrorsPeerAddress
    
    String connectionErrorsSelect
    
    String connectionErrorsTcpwrap
    
    String connections
    
    String createdTmpDiskTables
    
    String createdTmpFiles
    
    String createdTmpTables
    
    String delayedErrors
    
    String delayedInsertThreads
    
    String delayedWrites
    
    String flushCommands
    
    String handlerCommit
    
    String handlerDelete
    
    String handlerDiscover
    
    String handlerExternalLock
    
    String handlerMrrInit
    
    String handlerPrepare
    
    String handlerReadFirst
    
    String handlerReadKey
    
    String handlerReadLast
    
    String handlerReadNext
    
    String handlerReadPrev
    
    String handlerReadRnd
    
    String handlerReadRndNext
    
    String handlerRollback
    
    String handlerSavepoint
    
    String handlerSavepointRollback
    
    String handlerUpdate
    
    String handlerWrite
    
    String innodbBackgroundLogSync
    
    String innodbBufferPoolDumpStatus
    
    String innodbBufferPoolLoadStatus
    
    String innodbBufferPoolResizeStatus
    
    String innodbBufferPoolPagesData
    
    String innodbBufferPoolBytesData
    
    String innodbBufferPoolPagesDirty
    
    String innodbBufferPoolBytesDirty
    
    String innodbBufferPoolPagesFlushed
    
    String innodbBufferPoolPagesFree
    
    String innodbBufferPoolPagesLRUFlushed
    
    String innodbBufferPoolPagesMadeNotYoung
    
    String innodbBufferPoolPagesMadeYoung
    
    String innodbBufferPoolPagesMisc
    
    String innodbBufferPoolPagesOld
    
    String innodbBufferPoolPagesTotal
    
    String innodbBufferPoolReadAheadRnd
    
    String innodbBufferPoolReadAhead
    
    String innodbBufferPoolReadAheadEvicted
    
    String innodbBufferPoolReadRequests
    
    String innodbBufferPoolReads
    
    String innodbBufferPoolWaitFree
    
    String innodbBufferPoolWriteRequests
    
    String innodbCheckpointAge
    
    String innodbCheckpointMaxAge
    
    String innodbDataFsyncs
    
    String innodbDataPendingFsyncs
    
    String innodbDataPendingReads
    
    String innodbDataPendingWrites
    
    String innodbDataRead
    
    String innodbDataReads
    
    String innodbDataWrites
    
    String innodbDataWritten
    
    String innodbDblwrPagesWritten
    
    String innodbDblwrWrites
    
    String innodbIbufFreeList
    
    String innodbIbufSegmentSize
    
    String innodbLogWaits
    
    String innodbLogWriteRequests
    
    String innodbLogWrites
    
    String innodbLsnCurrent
    
    String innodbLsnFlushed
    
    String innodbLsnLastCheckpoint
    
    String innodbMasterThreadActiveLoops
    
    String innodbMasterThreadIdleLoops
    
    String innodbMaxTrxId
    
    String innodbMemAdaptiveHash
    
    String innodbMemDictionary
    
    String innodbOldestViewLowLimitTrxId
    
    String innodbOsLogFsyncs
    
    String innodbOsLogPendingFsyncs
    
    String innodbOsLogPendingWrites
    
    String innodbOsLogWritten
    
    String innodbPageSize
    
    String innodbPagesCreated
    
    String innodbPagesRead
    
    String innodbPages0Read
    
    String innodbPagesWritten
    
    String innodbPurgeTrxId
    
    String innodbPurgeUndoNo
    
    String innodbRowLockCurrentWaits
    
    String innodbRowLockTime
    
    String innodbRowLockTimeAvg
    
    String innodbRowLockTimeMax
    
    String innodbRowLockWaits
    
    String innodbRowsDeleted
    
    String innodbRowsInserted
    
    String innodbRowsRead
    
    String innodbRowsUpdated
    
    String innodbNumOpenFiles
    
    String innodbTruncatedStatusWrites
    
    String innodbAvailableUndoLogs
    
    String innodbSecondaryIndexTriggeredClusterReads
    
    String innodbSecondaryIndexTriggeredClusterReadsAvoided
    
    String innodbBufferedAioSubmitted
    
    String innodbScanPagesContiguous
    
    String innodbScanPagesDisjointed
    
    String innodbScanPagesTotalSeekDistance
    
    String innodbScanDataSize
    
    String innodbScanDeletedRecsSize
    
    String innodbScrubBackgroundPageReorganizations
    
    String innodbScrubBackgroundPageSplits
    
    String innodbScrubBackgroundPageSplitFailuresUnderflow
    
    String innodbScrubBackgroundPageSplitFailuresOutOfFilespace
    
    String innodbScrubBackgroundPageSplitFailuresMissingIndex
    
    String innodbScrubBackgroundPageSplitFailuresUnknown
    
    String innodbScrubLog
    
    String innodbEncryptionRotationPagesReadFromCache
    
    String innodbEncryptionRotationPagesReadFromDisk
    
    String innodbEncryptionRotationPagesModified
    
    String innodbEncryptionRotationPagesFlushed
    
    String innodbEncryptionRotationEstimatedIops
    
    String innodbEncryptionKeyRotationListLength
    
    String innodbEncryptionNMergeBlocksEncrypted
    
    String innodbEncryptionNMergeBlocksDecrypted
    
    String innodbEncryptionNRowlogBlocksEncrypted
    
    String innodbEncryptionNRowlogBlocksDecrypted
    
    String innodbNumPagesEncrypted
    
    String innodbNumPagesDecrypted
    
    String innodbEncryptionRedoKeyVersion
    
    String keyBlocksNotFlushed
    
    String keyBlocksUnused
    
    String keyBlocksUsed
    
    String keyReadRequests
    
    String keyReads
    
    String keyWriteRequests
    
    String keyWrites
    
    String lastQueryCost
    
    String lastQueryPartialPlans
    
    String lockedConnects
    
    String maxExecutionTimeExceeded
    
    String maxExecutionTimeSet
    
    String maxExecutionTimeSetFailed
    
    String maxUsedConnections
    
    String maxUsedConnectionsTime
    
    String notFlushedDelayedRows
    
    String ongoingAnonymousTransactionCount
    
    String openFiles
    
    String openStreams
    
    String openTableDefinitions
    
    String openTables
    
    String openedFiles
    
    String openedTableDefinitions
    
    String openedTables
    
    String performanceSchemaAccountsLost
    
    String performanceSchemaCondClassesLost
    
    String performanceSchemaCondInstancesLost
    
    String performanceSchemaDigestLost
    
    String performanceSchemaFileClassesLost
    
    String performanceSchemaFileHandlesLost
    
    String performanceSchemaFileInstancesLost
    
    String performanceSchemaHostsLost
    
    String performanceSchemaIndexStatLost
    
    String performanceSchemaLockerLost
    
    String performanceSchemaMemoryClassesLost
    
    String performanceSchemaMetadataLockLost
    
    String performanceSchemaMutexClassesLost
    
    String performanceSchemaMutexInstancesLost
    
    String performanceSchemaNestedStatementLost
    
    String performanceSchemaPreparedStatementsLost
    
    String performanceSchemaProgramLost
    
    String performanceSchemaRwlockClassesLost
    
    String performanceSchemaRwlockInstancesLost
    
    String performanceSchemaSessionConnectAttrsLost
    
    String performanceSchemaSocketClassesLost
    
    String performanceSchemaSocketInstancesLost
    
    String performanceSchemaStageClassesLost
    
    String performanceSchemaStatementClassesLost
    
    String performanceSchemaTableHandlesLost
    
    String performanceSchemaTableInstancesLost
    
    String performanceSchemaTableLockStatLost
    
    String performanceSchemaThreadClassesLost
    
    String performanceSchemaThreadInstancesLost
    
    String performanceSchemaUsersLost
    
    String preparedStmtCount
    
    String qcacheFreeBlocks
    
    String qcacheFreeMemory
    
    String qcacheHits
    
    String qcacheInserts
    
    String qcacheLowmemPrunes
    
    String qcacheNotCached
    
    String qcacheQueriesInCache
    
    String qcacheTotalBlocks
    
    String queries
    
    String questions
    
    String rsaPublicKey
    
    String selectFullJoin
    
    String selectFullRangeJoin
    
    String selectRange
    
    String selectRangeCheck
    
    String selectScan
    
    String slaveOpenTempTables
    
    String slowLaunchThreads
    
    String slowQueries
    
    String sortMergePasses
    
    String sortRange
    
    String sortRows
    
    String sortScan
    
    String sslAcceptRenegotiates
    
    String sslAccepts
    
    String sslCallbackCacheHits
    
    String sslCipher
    
    String sslCipherList
    
    String sslClientConnects
    
    String sslConnectRenegotiates
    
    String sslCtxVerifyDepth
    
    String sslCtxVerifyMode
    
    String sslDefaultTimeout
    
    String sslFinishedAccepts
    
    String sslFinishedConnects
    
    String sslServerNotAfter
    
    String sslServerNotBefore
    
    String sslSessionCacheHits
    
    String sslSessionCacheMisses
    
    String sslSessionCacheMode
    
    String sslSessionCacheOverflows
    
    String sslSessionCacheSize
    
    String sslSessionCacheTimeouts
    
    String sslSessionsReused
    
    String sslUsedSessionCacheEntries
    
    String sslVerifyDepth
    
    String sslVerifyMode
    
    String sslVersion
    
    String tableLocksImmediate
    
    String tableLocksWaited
    
    String tableOpenCacheHits
    
    String tableOpenCacheMisses
    
    String tableOpenCacheOverflows
    
    String tcLogMaxPagesUsed
    
    String tcLogPageSize
    
    String tcLogPageWaits
    
    String threadpoolIdleThreads
    
    String threadpoolThreads
    
    String threadsCached
    
    String threadsConnected
    
    String threadsCreated
    
    String threadsRunning
    
    String uptime
    
    String uptimeSinceFlushStatus
    
    String wsrepLocalStateUuid
    
    String wsrepProtocolVersion
    
    String wsrepLastApplied
    
    String wsrepLastCommitted
    
    String wsrepReplicated
    
    String wsrepReplicatedBytes
    
    String wsrepReplKeys
    
    String wsrepReplKeysBytes
    
    String wsrepReplDataBytes
    
    String wsrepReplOtherBytes
    
    String wsrepReceived
    
    String wsrepReceivedBytes
    
    String wsrepLocalCommits
    
    String wsrepLocalCertFailures
    
    String wsrepLocalReplays
    
    String wsrepLocalSendQueue
    
    String wsrepLocalSendQueueMax
    
    String wsrepLocalSendQueueMin
    
    String wsrepLocalSendQueueAvg
    
    String wsrepLocalRecvQueue
    
    String wsrepLocalRecvQueueMax
    
    String wsrepLocalRecvQueueMin
    
    String wsrepLocalRecvQueueAvg
    
    String wsrepLocalCachedDownto
    
    String wsrepFlowControlPausedNs
    
    String wsrepFlowControlPaused
    
    String wsrepFlowControlSent
    
    String wsrepFlowControlRecv
    
    String wsrepFlowControlInterval
    
    String wsrepFlowControlIntervalLow
    
    String wsrepFlowControlIntervalHigh
    
    String wsrepFlowControlStatus
    
    String wsrepFlowControlActive
    
    String wsrepFlowControlRequested
    
    String wsrepCertDepsDistance
    
    String wsrepApplyOooe
    
    String wsrepApplyOool
    
    String wsrepApplyWindow
    
    String wsrepCommitOooe
    
    String wsrepCommitOool
    
    String wsrepCommitWindow
    
    String wsrepLocalState
    
    String wsrepLocalStateComment
    
    String wsrepCertIndexSize
    
    String wsrepCertBucketCount
    
    String wsrepGcachePoolSize
    
    String wsrepCausalReads
    
    String wsrepCertInterval
    
    String wsrepOpenTransactions
    
    String wsrepOpenConnections
    
    String wsrepIstReceiveStatus
    
    String wsrepIstReceiveSeqnoStart
    
    String wsrepIstReceiveSeqnoCurrent
    
    String wsrepIstReceiveSeqnoEnd
    
    String wsrepIncomingAddresses
    
    String wsrepClusterWeight
    
    String wsrepDesyncCount
    
    String wsrepEvsDelayed
    
    String wsrepEvsEvictList
    
    String wsrepEvsReplLatency
    
    String wsrepEvsState
    
    String wsrepGcommUuid
    
    String wsrepGmcastSegment
    
    String wsrepClusterConfId
    
    String wsrepClusterSize
    
    String wsrepClusterStateUuid
    
    String wsrepClusterStatus
    
    String wsrepConnected
    
    String wsrepLocalBfAborts
    
    String wsrepLocalIndex
    
    String wsrepProviderName
    
    String wsrepProviderVendor
    
    String wsrepProviderVersion
    
    String wsrepReady
}
