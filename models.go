package main

type Repository struct {
	Id       int
	Name     string
	IsPublic bool
}

var reposDb = map[string]Repository{
	"gmail": {Id: 1, Name: "gmail"},
	"react": {Id: 2, Name: "react", IsPublic: true},
	"oso":   {Id: 3, Name: "oso"},
}

func GetRepositoryByName(name string) Repository {
	return reposDb[name]
}

type RepositoryRole struct {
	Role   string
	RepoId int
}

type User struct {
	Roles []RepositoryRole
}

var usersDb = map[string]User{
	"larry":  {Roles: []RepositoryRole{{Role: "admin", RepoId: 1}}},
	"anne":   {Roles: []RepositoryRole{{Role: "maintainer", RepoId: 2}}},
	"graham": {Roles: []RepositoryRole{{Role: "contributor", RepoId: 3}}},
}

func GetCurrentUser() User {
	return usersDb["larry"]
}
