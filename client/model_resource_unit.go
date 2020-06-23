/*
 * Pulsar Admin REST API
 *
 * This provides the REST API for admin operations
 *
 * API version: v2
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

type ResourceUnit struct {
	AvailableResource *ResourceDescription `json:"availableResource,omitempty"`
	ResourceId        string               `json:"resourceId,omitempty"`
}
