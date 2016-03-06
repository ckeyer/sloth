package types

type User struct {
	ID       int64 `gorm:"primary_key,auto_increment"`
	Name     string
	Password string
	*GithubAccount
}

type GithubAccount struct {
	ID       int64 `gorm:"primary_key,auto_increment"`
	UserName string
	Password string
}
