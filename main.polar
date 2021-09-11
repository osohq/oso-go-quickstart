# TODO: update for main

actor User {}

resource Repository {
	permissions = ["read", "push", "delete"];
	roles = ["contributor", "maintainer", "admin"];

	"read" if "contributor";
	"push" if "maintainer";
	"delete" if "admin";

	"maintainer" if "admin";
	"contributor" if "maintainer";
}

has_role(user: User, role_name: String, repository: Repository) if
  role in user.Roles and
  role matches { Role: role_name, RepoId: repository.Id };

allow(actor, action, resource) if
  has_permission(actor, action, resource);
