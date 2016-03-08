package types

type User struct {
	ID            int64          `bson:"_id"`
	Name          string         `bson:"name"`
	Email         string         `bson:"email"`
	Password      string         `bson:"password"`
	Role          string         `bson:"role"` // admin member
	GithubAccount *GithubAccount `bson:"github"`
}

type GithubAccount struct {
	Name     string `bson:"name"`
	Email    string `bson:"email"`
	Password string `bson:"password"`
	Token    string `bson:"token"`
}

type WebhookConfig struct {
	RepoName string `bson:""`
	Secret   string `bson:"secret"`
}

func (g *GithubAccount) GetToken() string {
	if g.Token != "" {
		return g.Token
	}
	return g.Password
}
