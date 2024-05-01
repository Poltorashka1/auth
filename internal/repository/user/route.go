package repoUser

import (
	serviceUserModel "auth/internal/service/user/model"
	"context"
)

func (r *UserRepos) GetRouteRole(ctx context.Context, data serviceUserModel.CheckUserRoleData) ([]string, error) {
	//q := db.NewQuery("GetRouteRole",
	//	`select role from "Roles" where route = $1`,
	//	[]interface{}{data.Route})
	//r.cache.Get(ctx, "")

	return nil, nil
}
