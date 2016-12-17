package gh

type Giter interface {
	Clone(repo, ref string) (string, error)
}
