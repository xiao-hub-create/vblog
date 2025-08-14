package token

import "time"

// 用户身份令牌
type Token struct {
	//主键
	Id uint `json:"id" gorm:"primaryKey;column:id"`
	//用户id
	RefUserId uint `json:"user_id" gorm:"column:ref_user_id"`
	//关联查询
	RefUserName string `json:"ref_user_name" gorm:"-"`

	//颁发时间
	IssueAt time.Time `json:"issue_at" gorm:"column:issue_at"` //默认值为当前时间

	//访问Token
	AccessToken string `json:"access_token" gorm:"column:access_token;unique;index"`
	//访问Token过期时间
	AccessTokenExpireAt *time.Time `json:"access_token_expire_at" gorm:"column:access_token_expire_at"`

	//刷新Token
	RefreshToken string `json:"refresh_token" gorm:"column:refresh_token;unique;index"`
	//刷新Token过期时间
	RefreshTokenExpireAt *time.Time `json:"refresh_token_expire_at" gorm:"column:refresh_token_expire_at"`
}

func (t *Token) TableName() string {
	return "tokens"
}
