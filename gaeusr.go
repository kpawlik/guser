package guser

import (
	"appengine"
	"appengine/datastore"
	"appengine/user"
)

const (
	UsrTableName = "Usr"
)

/*************************************
User
*************************************/
type User struct {
	Name   string
	Active bool
}

func NewUser(name string) *User {
	return &User{name, true}
}

func (u User) Key(c appengine.Context) *datastore.Key {
	return datastore.NewKey(c, UsrTableName, u.Name, 0, nil)
}

func (u *User) Save(c appengine.Context) error {
	_, err := datastore.Put(c, u.Key(c), u)
	return err
}

func GetUser(c appengine.Context, name string) (u *User) {
	u = &User{}
	key := datastore.NewKey(c, UsrTableName, name, 0, nil)
	if datastore.Get(c, key, u) == datastore.ErrNoSuchEntity {
		u = nil
	}
	return
}

func GetUserFromContext(c appengine.Context) (u *User) {
	cusr := user.Current(c)
	return GetUser(c, cusr.String())
}

func Users(c appengine.Context) (usrs []*User, err error) {
	q := datastore.NewQuery(UsrTableName)
	for t := q.Run(c); ; {
		result := &User{}
		_, err = t.Next(result)
		if err == datastore.Done {
			err = nil
			break
		}
		if err != nil {
			return
		}
		usrs = append(usrs, result)
	}
	return
}
