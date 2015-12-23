package main

import (
	"github.com/julienschmidt/httprouter"
	"github.com/nu7hatch/gouuid"
	"golang.org/x/crypto/bcrypt"
	"google.golang.org/appengine"
	"google.golang.org/appengine/datastore"
	"google.golang.org/appengine/log"
	"google.golang.org/appengine/memcache"
	"html/template"
	"net/http"
	"time"
	"encoding/json"
)

type User struct {
	Email    string
	UserName string `datastore:"-"`
	Password string `json:"-"`
}

type sessionData struct {
	User
	LoggedIn  bool
	LoginFail bool
}

var myTmpl *template.Template

func init() {
	r := httprouter.New()
	http.Handle("/", r)
	r.GET("/", index)
	r.POST("/login", login)
	r.GET("/logout", logout)
	r.GET("/signup",signup)
	r.POST("/register",registeruser)
	r.GET("/upload",uploadtemplate)
	r.POST("/uploadfile",uploadfile)
	myTmpl = template.Must(myTmpl.ParseGlob("assets/*.html"))
}

func index(res http.ResponseWriter, req *http.Request, ps httprouter.Params) {
	//if logged in, redirect to upload template
	memItem := getSession(req)
	if len(memItem.Value) > 0 {
		http.Redirect(res,req,"/upload",http.StatusFound)
		return
	} else {
		//serve template for home page
		myTmpl.ExecuteTemplate(res, "index", nil)
	}
}

func login(res http.ResponseWriter, req *http.Request, ps httprouter.Params) {
	//if logged in, redirect to upload template
	memItem := getSession(req)
	if len(memItem.Value) > 0 {
		http.Redirect(res,req,"/upload",http.StatusFound)
		return
	}
	//check user login info
	ctx := appengine.NewContext(req)
	key := datastore.NewKey(ctx, "Users", req.FormValue("username"), 0, nil)
	var user User
	err := datastore.Get(ctx, key, &user)
	if err != nil || bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.FormValue("password"))) != nil {
		// failure logging in
		myTmpl.ExecuteTemplate(res, "index", true)
		return
	} else {
		user.UserName = req.FormValue("username")
		// success logging in
		createSession(res, req, user)
		// redirect
		http.Redirect(res, req, "/upload", http.StatusFound)
	}
}

func signup(res http.ResponseWriter, req *http.Request, ps httprouter.Params) {
	//if logged in, redirect to upload template
	memItem := getSession(req)
	if len(memItem.Value) > 0 {
		http.Redirect(res,req,"/upload",http.StatusFound)
		return
	} else {
		//serve template for home page
		myTmpl.ExecuteTemplate(res, "signup", nil)
	}
}

func registeruser(res http.ResponseWriter, req *http.Request, ps httprouter.Params) {
	//if logged in, redirect to upload template
	memItem := getSession(req)
	if len(memItem.Value) > 0 {
		http.Redirect(res,req,"/upload",http.StatusFound)
		return
	} else {
		//register the user from submitted signup form (check if user exists)
		if checkusername(res,req,req.FormValue("username")) {
			myTmpl.ExecuteTemplate(res, "signup", true)
			return
		}
		ctx := appengine.NewContext(req)
		hashedPass, err := bcrypt.GenerateFromPassword([]byte(req.FormValue("password")), bcrypt.DefaultCost)
		if err != nil {
			log.Errorf(ctx, "error creating password: %v", err)
			http.Error(res, err.Error(), 500)
			return
		}
		user := User{
			Email:    req.FormValue("email"),
			UserName: req.FormValue("username"),
			Password: string(hashedPass),
		}
		key := datastore.NewKey(ctx, "Users", user.UserName, 0, nil)
		key, err = datastore.Put(ctx, key, &user)
		if err != nil {
			log.Errorf(ctx, "error adding todo: %v", err)
			http.Error(res, err.Error(), 500)
			return
		}

		createSession(res, req, user)
		// redirect
		http.Redirect(res, req, "/upload", 302)
	}
}

func checkusername(res http.ResponseWriter, req *http.Request, username string) bool {
	ctx := appengine.NewContext(req)
	var user User
	key := datastore.NewKey(ctx, "Users", username, 0, nil)
	err := datastore.Get(ctx, key, &user)
	// if there is an err, there is NO user
	log.Infof(ctx, "ERR: %v", err)
	if err != datastore.ErrNoSuchEntity {
		return true
	} else {
		return false
	}
}

func uploadtemplate(res http.ResponseWriter, req *http.Request, ps httprouter.Params) {
	//if logged in, serve upload template
	memItem := getSession(req)
	if len(memItem.Value) > 0 {
		//serve template for uploading file
		myTmpl.ExecuteTemplate(res, "upload", nil)
		return
	} else {
		//redirect to index
		http.Redirect(res,req,"/",http.StatusFound)
	}
}

func uploadfile(res http.ResponseWriter, req *http.Request, ps httprouter.Params) {
	//upload file
}

func createSession(res http.ResponseWriter, req *http.Request, user User) {
	ctx := appengine.NewContext(req)
	// SET COOKIE
	id, _ := uuid.NewV4()
	cookie := &http.Cookie{
		Name:  "session",
		Value: id.String(),
		Path:  "/",
		//		UNCOMMENT WHEN DEPLOYED:
		//		Secure: true,
		//		HttpOnly: true,
	}
	http.SetCookie(res, cookie)

	// SET MEMCACHE session data (sd)
	json, err := json.Marshal(user)
	if err != nil {
		log.Errorf(ctx, "error marshalling during user creation: %v", err)
		http.Error(res, err.Error(), 500)
		return
	}
	sd := memcache.Item{
		Key:   id.String(),
		Value: json,
		//		Expiration: time.Duration(20*time.Minute),
		Expiration: time.Duration(20 * time.Second),
	}
	memcache.Set(ctx, &sd)
}

func getSession(req *http.Request) *memcache.Item {
	cookie, err := req.Cookie("session")
	if err != nil {
	return &memcache.Item{}
	}

	ctx := appengine.NewContext(req)
	item, err := memcache.Get(ctx, cookie.Value)
	if err != nil {
	return &memcache.Item{}
	}
	return item
}

func logout(res http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	ctx := appengine.NewContext(req)

	cookie, err := req.Cookie("session")
	// cookie is not set
	if err != nil {
		http.Redirect(res, req, "/", 302)
		return
	}

	// clear memcache
	sd := memcache.Item{
		Key:        cookie.Value,
		Value:      []byte(""),
		Expiration: time.Duration(1 * time.Microsecond),
	}
	memcache.Set(ctx, &sd)

	// clear the cookie
	cookie.MaxAge = -1
	http.SetCookie(res, cookie)

	// redirect
	http.Redirect(res, req, "/", 302)
}
