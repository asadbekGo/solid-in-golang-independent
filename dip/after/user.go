package after

type IUser interface {
	GetSender() *User
	GetReceiver() *User
}

type User struct {
	Id    int
	Name  string
	Phone int
}

func (u *User) GetSender() *User {
	return &User{
		Id:    1,
		Name:  "Bob",
		Phone: 1234,
	}
}

func (u *User) GetReceiver() *User {
	return &User{
		Id:    2,
		Name:  "John",
		Phone: 4321,
	}
}
