/*
 * Pulsar Admin REST API
 *
 * This provides the REST API for admin operations
 *
 * API version: v2
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

type PoolChunkStats struct {
	Usage     int32 `json:"usage,omitempty"`
	ChunkSize int32 `json:"chunkSize,omitempty"`
	FreeBytes int32 `json:"freeBytes,omitempty"`
}
