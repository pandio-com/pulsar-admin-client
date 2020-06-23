/*
 * Pulsar Admin REST API
 *
 * This provides the REST API for admin operations
 *
 * API version: v2
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

type NamespaceBundleStats struct {
	MsgRateIn        float64 `json:"msgRateIn,omitempty"`
	MsgThroughputIn  float64 `json:"msgThroughputIn,omitempty"`
	MsgRateOut       float64 `json:"msgRateOut,omitempty"`
	MsgThroughputOut float64 `json:"msgThroughputOut,omitempty"`
	ConsumerCount    int32   `json:"consumerCount,omitempty"`
	ProducerCount    int32   `json:"producerCount,omitempty"`
	Topics           int64   `json:"topics,omitempty"`
	CacheSize        int64   `json:"cacheSize,omitempty"`
}
