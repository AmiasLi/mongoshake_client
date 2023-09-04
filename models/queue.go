package models

type Queue struct {
	SyncerReplicaSetName string             `json:"syncer_replica_set_name"`
	LogsQueueSize        int                `json:"logs_queue_size"`
	PendingQueueSize     int                `json:"pending_queue_size"`
	SyncerInnerQueue     []SyncerInnerQueue `json:"syncer_inner_queue"`
	PersisterBufferUsed  int                `json:"persister_buffer_used"`
}

type SyncerInnerQueue struct {
	QueueId          int `json:"queue_id"`
	PendingQueueUsed int `json:"pending_queue_used"`
	LogsQueueUsed    int `json:"logs_queue_used"`
}
