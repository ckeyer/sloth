package gh

// Giter
type Giter interface {
	Clone(repo, ref string) (string, error)
}
