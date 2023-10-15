package model

type Product struct {
	ID        int
	Name      string
	Variants  []Variant
	CreatedAt string
	UpdatedAt string
	// CreatedAt *time.Time
	// UpdatedAt *time.Time
}
