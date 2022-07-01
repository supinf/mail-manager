package model

type User struct {
	APIKey      *string `json:"apiKey" dynamo:"api_key,hash"`
	APIKeyID    string  `json:"apiKeyID" dynamo:"api_key_id"`
	Name        string  `json:"name" dynamo:"name" index:"name-index,hash"`
	UsagePlanID string  `json:"usagePlanID" dynamo:"usage_plan_id"`
	Mail        string  `json:"mail" dynamo:"mail"`
	Role        Role    `json:"role" dynamo:"role"`
}

// Role ロール
type Role int64

const (
	// RoleNone なし
	RoleNone Role = 0
	// RoleAdmin 管理者
	RoleAdmin Role = 1
)

func FindUserByHash(hashKey string) (*User, error) {
	var resp User

	query := newQuery()
	query = query.Table(User{}).Get("api_key", hashKey)

	if err := query.One(&resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func ListUserByGSI(hashKey string) ([]User, error) {
	var list []User

	query := newQuery()
	query = query.Table(User{}).IndexGet("name-index", "name", hashKey)

	if err := query.All(&list); err != nil {
		return nil, err
	}
	return list, nil
}

func (u *User) Create() error {
	query := newQuery()
	return query.Table(User{}).Put(u).Run()
}
