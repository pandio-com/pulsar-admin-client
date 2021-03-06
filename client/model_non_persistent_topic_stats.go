/*
 * Pulsar Admin REST API
 *
 * This provides the REST API for admin operations
 *
 * API version: v2
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package client

type NonPersistentTopicStats struct {
	MsgRateIn           float64                                   `json:"msgRateIn,omitempty"`
	MsgThroughputIn     float64                                   `json:"msgThroughputIn,omitempty"`
	MsgRateOut          float64                                   `json:"msgRateOut,omitempty"`
	MsgThroughputOut    float64                                   `json:"msgThroughputOut,omitempty"`
	BytesInCounter      int64                                     `json:"bytesInCounter,omitempty"`
	MsgInCounter        int64                                     `json:"msgInCounter,omitempty"`
	BytesOutCounter     int64                                     `json:"bytesOutCounter,omitempty"`
	MsgOutCounter       int64                                     `json:"msgOutCounter,omitempty"`
	AverageMsgSize      float64                                   `json:"averageMsgSize,omitempty"`
	StorageSize         int64                                     `json:"storageSize,omitempty"`
	BacklogSize         int64                                     `json:"backlogSize,omitempty"`
	Publishers          []NonPersistentPublisherStats             `json:"publishers,omitempty"`
	Subscriptions       map[string]NonPersistentSubscriptionStats `json:"subscriptions,omitempty"`
	Replication         map[string]NonPersistentReplicatorStats   `json:"replication,omitempty"`
	DeduplicationStatus string                                    `json:"deduplicationStatus,omitempty"`
	MsgDropRate         float64                                   `json:"msgDropRate,omitempty"`
}
