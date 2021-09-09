package main

type Repository struct {
	Id   int
	Name string
}

var reposDb = []Repository{
	{Id: 0, Name: "gmail"}, {Id: 1, Name: "react"}, {Id: 2, Name: "oso"},
}

func GetRepositoryById(id int) Repository {
	return reposDb[id]
}

type RepositoryRole struct {
	Role   string
	RepoId int
}

type User struct {
	Roles []RepositoryRole
}

var usersDb = map[string]User{
	"larry": {
		Roles: []RepositoryRole{
			{Role: "admin", RepoId: 0},
		},
	},
	"anne": {
		Roles: []RepositoryRole{
			{Role: "maintainer", RepoId: 1},
		},
	},
	"graham": {
		Roles: []RepositoryRole{
			{Role: "contributor", RepoId: 2},
		},
	},
}

func GetCurrentUser() User {
	return usersDb["larry"]
}
