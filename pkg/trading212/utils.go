package trading212

// SecureString is a string type that doesn't print.
// Used to avoid leaking credentials when logging.
type SecureString string

func (s SecureString) String() string {
	return "[REDACTED]"
}
