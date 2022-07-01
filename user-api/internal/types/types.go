// Code generated by goctl. DO NOT EDIT.
package types

type User struct {
	Id       int64  `json:"id"`
	Mobile   string `json:"mobile"`
	Nickname string `json:"nickname"`
	Sex      int64  `json:"sex"`
	Avatar   string `json:"avatar"`
}

type RegisterReq struct {
	Mobile         string `json:"mobile"`
	Password       string `json:"password" validate:"required"`
	RepeatPassword string `json:"repeatPassword" validate:"required"`
	Email          string `json:"email"`
	AuthCode       string `json:"authCode" validate:"required"`
	Area           string `json:"area"`
}

type RegisterResp struct {
	AccessToken  string `json:"accessToken"`
	AccessExpire int64  `json:"accessExpire"`
	RefreshAfter int64  `json:"refreshAfter"`
}

type UserInfoReq struct {
}

type UserInfoResp struct {
	UserInfo User `json:"userInfo"`
}
