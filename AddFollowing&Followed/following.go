package main

import (
	"encoding/json"
	"net/http"
	"github.com/julienschmidt/httprouter"
	"google.golang.org/appengine"
	"google.golang.org/appengine/datastore"
	"google.golang.org/appengine/log"
)

func following(res http.ResponseWriter, req *http.Request, ps httprouter.Params){
	ctx := appengine.NewContext(req)
	memItem, err := getSession(req)
	if err != nil{
		http.Redirect(res,req,"/form/login",http.StatusFound)
		return
	}
	var sd SessionData
	json.Unmarshal(memItem.Value, &sd)
	sd.LoggedIn = true
	Following, FollowingMe, err := getRelations(req,sd.UserName)
	if err == nil {
		sd.Following, sd.FollowingMe = Following, FollowingMe
	}else{
		log.Errorf(ctx, "Error getting Relations for %v: %v ---", sd.UserName,err)
		sd.Following, sd.FollowingMe = []Relation{}, []Relation{}
	}
	tpl.ExecuteTemplate(res, "following.html", &sd)
}

func getRelations(req *http.Request, username string) ([]Relation, []Relation, error) {
	ctx := appengine.NewContext(req)
	var Following, FollowingMe []Relation
	q := datastore.NewQuery("Relation")
	_, err := q.Filter("Follower =",username).Order("Following").GetAll(ctx,&Following)
	if err != nil {
		return []Relation{},[]Relation{},err
	}
	q = datastore.NewQuery("Relation")
	_, err = q.Filter("Following =",username).Order("Follower").GetAll(ctx,&FollowingMe)
	if err != nil {
		return []Relation{},[]Relation{},err
	}
	return Following, FollowingMe, nil
}

func follow(res http.ResponseWriter, req *http.Request, ps httprouter.Params) {
	ctx := appengine.NewContext(req)
	memItem, err := getSession(req)
	if err != nil{
		http.Redirect(res,req,"/form/login",http.StatusFound)
		return
	}
	var user User
	json.Unmarshal(memItem.Value, &user)
	Following := ps.ByName("user")
	if Following == user.UserName {
		http.Redirect(res,req,"/",http.StatusFound)
		return
	}
	key := datastore.NewKey(ctx,"Relation",user.UserName+Following,0,nil)
	_, err = datastore.Put(ctx,key,&Relation{user.UserName,Following})
	if err != nil {
		log.Errorf(ctx, "error adding todo: %v", err)
		http.Error(res, err.Error(), 500)
		return
	}
	http.Redirect(res,req,"/",http.StatusFound)
}

func unfollow(res http.ResponseWriter, req *http.Request, ps httprouter.Params) {
	ctx := appengine.NewContext(req)
	memItem, err := getSession(req)
	if err != nil{
		http.Redirect(res,req,"/form/login",http.StatusFound)
		return
	}
	var user User
	json.Unmarshal(memItem.Value, &user)
	Following := ps.ByName("user")
	if Following == user.UserName {
		http.Redirect(res,req,"/",http.StatusFound)
		return
	}
	key := datastore.NewKey(ctx,"Relation",user.UserName+Following,0,nil)
	err = datastore.Delete(ctx,key)
	if err != nil {
		log.Errorf(ctx, "error deleting todo: %v", err)
		http.Error(res, err.Error(), 500)
		return
	}
	http.Redirect(res,req,"/",http.StatusFound)
}