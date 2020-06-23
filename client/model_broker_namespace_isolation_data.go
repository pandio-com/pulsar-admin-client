/*
 * Pulsar Admin REST API
 *
 * This provides the REST API for admin operations
 *
 * API version: v2
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package client

// The namespace isolation data for a given broker
type BrokerNamespaceIsolationData struct {
	// The broker name
	BrokerName string `json:"brokerName,omitempty"`
	// Policy name
	PolicyName string `json:"policyName,omitempty"`
	// Is Primary broker
	IsPrimary bool `json:"isPrimary,omitempty"`
	// The namespace-isolation policies attached to this broker
	NamespaceRegex []string `json:"namespaceRegex,omitempty"`
}
