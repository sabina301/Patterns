package code

type Connection struct {
	Id string
}

func (connection *Connection) GetID() string {
	return connection.Id
}
