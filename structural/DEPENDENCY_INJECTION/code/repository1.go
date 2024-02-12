package code

type Repository1 struct {
}

func (r *Repository1) GetUser(id int) string {
	return "user1"
}
