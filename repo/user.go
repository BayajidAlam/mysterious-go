package repo

type User struct {
	ID          int    `json:"id"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	Email       string `json:"email"`
	Password    string `json:"password"`
	IsShopOwner bool   `json:"is_shop_owner"`
}

type UserRepo interface {
	Create(u User) (*User, error)
	Get(email string, password string) (*User, error)
}

type userRepo struct {
	userList []*User
}

func NewUserRepo() *userRepo {
	repo := &userRepo{}
	return repo
}

func (r *userRepo) Create(u User) (*User, error) {
	u.ID = len(r.userList) + 1
	r.userList = append(r.userList, &u)
	return &u, nil
}

func (r *userRepo) Get(email string, password string) (*User, error) {
	for _, user := range r.userList {
		if user.Email == email && user.Password == password {
			return user, nil
		}
	}
	return nil, nil
}
