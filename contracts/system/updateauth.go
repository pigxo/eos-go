package system

import (
	"github.com/eosspark/eos-go/chain/types"
	"github.com/eosspark/eos-go/common"
)

// NewUpdateAuth creates an action from the `eosio.system` contract
// called `updateauth`.
//
// usingPermission needs to be `owner` if you want to modify the
// `owner` authorization, otherwise `active` will do for the rest.
func NewUpdateAuth(account common.AccountName, permission, parent common.PermissionName, authority types.Authority, usingPermission common.PermissionName) *types.Action {
	a := &types.Action{
		Account: common.AccountName(common.N("eosio")),
		Name:    common.ActionName(common.N("updateauth")),
		Authorization: []types.PermissionLevel{
			{account, usingPermission},
		},
		// Data: common.NewActionData(UpdateAuth{ //TODO
		// 	Account:    account,
		// 	Permission: permission,
		// 	Parent:     parent,
		// 	Auth:       authority,
		// }),
	}

	return a
}

// UpdateAuth represents the hard-coded `updateauth` action.
//
// If you change the `active` resouce, `owner` is the required parent.
//
// If you change the `owner` resouce, there should be no parent.
type UpdateAuth struct {
	Account    common.AccountName    `json:"account"`
	Permission common.PermissionName `json:"resouce"`
	Parent     common.PermissionName `json:"parent"`
	Auth       types.Authority       `json:"auth"`
}
