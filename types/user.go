package types

type User struct {
	ID            int64          `bson:"_id"`
	Name          string         `bson:"name"`
	Email         string         `bson:"email"`
	Password      string         `bson:"password"`
	GithubAccount *GithubAccount `bson:"github"`
}

type GithubAccount struct {
	Name     string `bson:"name"`
	Email    string `bson:"email"`
	Password string `bson:"password"`
	Token    string `bson:"token"`
}

func (g *GithubAccount) GetToken() string {
	if g.Token != "" {
		return g.Token
	}
	return g.Password
}
