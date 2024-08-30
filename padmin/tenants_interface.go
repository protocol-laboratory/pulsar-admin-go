package padmin

type Tenants interface {
	Create(tenantName string, info TenantInfo) error
	Delete(tenantName string) error
	List() ([]string, error)
}

type TenantInfo struct {
	AdminRoles      []string `json:"adminRoles,omitempty"`
	AllowedClusters []string `json:"allowedClusters,omitempty"`
}
