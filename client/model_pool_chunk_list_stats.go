/*
 * Pulsar Admin REST API
 *
 * This provides the REST API for admin operations
 *
 * API version: v2
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

type PoolChunkListStats struct {
	MinUsage int32            `json:"minUsage,omitempty"`
	MaxUsage int32            `json:"maxUsage,omitempty"`
	Chunks   []PoolChunkStats `json:"chunks,omitempty"`
}