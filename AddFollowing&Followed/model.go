package main

import "time"

type User struct {
	Email    string
	UserName string `datastore:"-"`
	Password string `json:"-"`
}

type SessionData struct {
	User
	LoggedIn  bool
	LoginFail bool
	Tweets    []Tweet
	Following []Relation
	FollowingMe []Relation
}

type Tweet struct {
	Msg      string
	Time     time.Time
	UserName string
}

type Relation struct {
	Follower string
	Following string
}
