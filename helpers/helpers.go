package helpers

// Helper function to convert string to *string
func ToPtr(s string) *string {
	if s == "" {
		return nil // Ensures nil values instead of empty strings
	}
	return &s
}
