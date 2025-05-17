package repository

// UnitOfWork wraps all your aggregate repositories inside a transaction.
type UnitOfWork interface {
	// Tenant Repositories
	Tenant() TenantRepository
	Organization() OrganizationRepository
	User() UserRepository

	// Commit Transaction management
	// Commit the transaction (or return an error to rollback)
	Commit() error
	Rollback() error
}
