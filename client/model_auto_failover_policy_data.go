/*
 * Pulsar Admin REST API
 *
 * This provides the REST API for admin operations
 *
 * API version: v2
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

// The auto failover policy configuration data
type AutoFailoverPolicyData struct {
	// The auto failover policy type
	PolicyType string `json:"policy_type,omitempty"`
	// The parameters applied to the auto failover policy specified by `policy_type`. The parameters for 'min_available' are :   - 'min_limit': the limit of minimal number of available brokers in primary group before auto failover   - 'usage_threshold': the resource usage threshold. If the usage of a broker is beyond this value, it would be marked as unavailable
	Parameters map[string]string `json:"parameters,omitempty"`
}
