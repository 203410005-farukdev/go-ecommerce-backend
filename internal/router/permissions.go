package router

// PermissionDefinition describes a single route-level access rule seeded into the database.
type PermissionDefinition struct {
	Name        string
	Module      string
	Method      string
	Path        string
	Description string
}

// PermissionDefinitions returns the full list of route permissions used by the RBAC seeder.
func PermissionDefinitions() []PermissionDefinition {
	return []PermissionDefinition{
		{Name: "auth:login", Module: "AUTH", Method: "POST", Path: "/api/v1/auth/login", Description: "Login user"},
		{Name: "auth:register", Module: "AUTH", Method: "POST", Path: "/api/v1/auth/register", Description: "Register user"},
		{Name: "auth:refresh", Module: "AUTH", Method: "POST", Path: "/api/v1/auth/refresh", Description: "Refresh access token"},
		{Name: "auth:logout", Module: "AUTH", Method: "POST", Path: "/api/v1/auth/logout", Description: "Logout user"},
		{Name: "user:me", Module: "USERS", Method: "GET", Path: "/api/v1/me", Description: "Get current user"},
		{Name: "user:list", Module: "USERS", Method: "GET", Path: "/api/v1/users", Description: "List all users"},
		{Name: "log:list", Module: "SYSTEM", Method: "GET", Path: "/api/v1/logs", Description: "List request logs"},
		{Name: "role:list", Module: "RBAC", Method: "GET", Path: "/api/v1/roles", Description: "List roles"},
		{Name: "role:create", Module: "RBAC", Method: "POST", Path: "/api/v1/roles", Description: "Create role"},
		{Name: "role:update", Module: "RBAC", Method: "PUT", Path: "/api/v1/roles/:id", Description: "Update role"},
		{Name: "role:delete", Module: "RBAC", Method: "DELETE", Path: "/api/v1/roles/:id", Description: "Delete role"},
		{Name: "role:permissions:list", Module: "RBAC", Method: "GET", Path: "/api/v1/roles/:id/permissions", Description: "Get role permissions"},
		{Name: "permission:list", Module: "RBAC", Method: "GET", Path: "/api/v1/permissions", Description: "List permissions"},
		{Name: "permission:create", Module: "RBAC", Method: "POST", Path: "/api/v1/permissions", Description: "Create permission"},
		{Name: "permission:update", Module: "RBAC", Method: "PUT", Path: "/api/v1/permissions/:id", Description: "Update permission"},
		{Name: "permission:delete", Module: "RBAC", Method: "DELETE", Path: "/api/v1/permissions/:id", Description: "Delete permission"},
		{Name: "role_permission:assign", Module: "RBAC", Method: "POST", Path: "/api/v1/roles/:id/permissions", Description: "Assign permission to role"},
		{Name: "role_permission:revoke", Module: "RBAC", Method: "DELETE", Path: "/api/v1/roles/:id/permissions/:permission_id", Description: "Revoke permission from role"},
		{Name: "user:assign_role", Module: "USERS", Method: "PATCH", Path: "/api/v1/users/:id/role", Description: "Assign role to user"},
		{Name: "user:permissions:list", Module: "USERS", Method: "GET", Path: "/api/v1/me/permissions", Description: "Get user permissions"},
		// Category permissions
		{Name: "category:list", Module: "PRODUCTS", Method: "GET", Path: "/api/v1/categories", Description: "List categories"},
		{Name: "category:create", Module: "PRODUCTS", Method: "POST", Path: "/api/v1/categories", Description: "Create category"},
		{Name: "category:update", Module: "PRODUCTS", Method: "PUT", Path: "/api/v1/categories/:id", Description: "Update category"},
		{Name: "category:delete", Module: "PRODUCTS", Method: "DELETE", Path: "/api/v1/categories/:id", Description: "Delete category"},
		// Subcategory permissions
		{Name: "subcategory:list", Module: "PRODUCTS", Method: "GET", Path: "/api/v1/subcategories", Description: "List subcategories"},
		{Name: "subcategory:create", Module: "PRODUCTS", Method: "POST", Path: "/api/v1/subcategories", Description: "Create subcategory"},
		{Name: "subcategory:update", Module: "PRODUCTS", Method: "PUT", Path: "/api/v1/subcategories/:id", Description: "Update subcategory"},
		{Name: "subcategory:delete", Module: "PRODUCTS", Method: "DELETE", Path: "/api/v1/subcategories/:id", Description: "Delete subcategory"},
		// Product permissions
		{Name: "product:list", Module: "PRODUCTS", Method: "GET", Path: "/api/v1/products", Description: "List products"},
		{Name: "product:create", Module: "PRODUCTS", Method: "POST", Path: "/api/v1/products", Description: "Create product"},
		{Name: "product:update", Module: "PRODUCTS", Method: "PUT", Path: "/api/v1/products/:id", Description: "Update product"},
		{Name: "product:delete", Module: "PRODUCTS", Method: "DELETE", Path: "/api/v1/products/:id", Description: "Delete product"},
		// Product variant permissions
		{Name: "product_variant:list", Module: "PRODUCTS", Method: "GET", Path: "/api/v1/products/:id/variants", Description: "List product variants"},
		{Name: "product_variant:create", Module: "PRODUCTS", Method: "POST", Path: "/api/v1/products/:id/variants", Description: "Create product variant"},
		{Name: "product_variant:update", Module: "PRODUCTS", Method: "PUT", Path: "/api/v1/products/:id/variants/:variant_id", Description: "Update product variant"},
		{Name: "product_variant:delete", Module: "PRODUCTS", Method: "DELETE", Path: "/api/v1/products/:id/variants/:variant_id", Description: "Delete product variant"},
		// Stock permissions
		{Name: "stock:list", Module: "INVENTORY", Method: "GET", Path: "/api/v1/stock", Description: "List stock"},
		{Name: "stock:update", Module: "INVENTORY", Method: "PUT", Path: "/api/v1/stock/:id", Description: "Update stock"},
		{Name: "stock_movement:list", Module: "INVENTORY", Method: "GET", Path: "/api/v1/stock-movements", Description: "List stock movements"},
		{Name: "stock_movement:create", Module: "INVENTORY", Method: "POST", Path: "/api/v1/stock-movements", Description: "Create stock movement"},
		// Setting permissions
		{Name: "setting:list", Module: "SETTINGS", Method: "GET", Path: "/api/v1/settings", Description: "Get settings"},
		{Name: "setting:update", Module: "SETTINGS", Method: "PUT", Path: "/api/v1/settings", Description: "Update settings"},
		// Dashboard permissions
		{Name: "dashboard:stats", Module: "DASHBOARD", Method: "GET", Path: "/api/v1/dashboard/stats", Description: "Get dashboard stats"},
	}
}
