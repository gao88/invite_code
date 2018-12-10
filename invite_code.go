package invite_code

import (
	"errors"

	"github.com/zxfonline/bases"
)

const (
	BASEID = uint64(1000000000)
)

//UserId 转 唯一邀请码
func GetInviteCode(userId uint64) (invitecode string) {
	invitecode = bases.ToBase(int(userId+BASEID), 58)
	return
}

//唯一邀请码 转 UserId
func GetUserId(inviteCode string) (uint64, error) {
	userid := uint64(bases.FromBase(inviteCode, 58)) - BASEID
	if userid < 0 {
		return 0, errors.New("invalid invite code")
	}
	return userid, nil
}
