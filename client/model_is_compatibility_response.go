/*
 * Pulsar Admin REST API
 *
 * This provides the REST API for admin operations
 *
 * API version: v2
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package client

type IsCompatibilityResponse struct {
	SchemaCompatibilityStrategy string `json:"schemaCompatibilityStrategy,omitempty"`
	Compatibility               bool   `json:"compatibility,omitempty"`
}
