package showvoting

import (
	"html/template"
	"net/http"

	router "github.com/julienschmidt/httprouter"
	"google.golang.org/appengine"
	"google.golang.org/appengine/datastore"
	"google.golang.org/appengine/log"
	"google.golang.org/appengine/user"
)

type Show struct {
	Title string
	Votes int
}

type Vote struct {
	Key  string
	User string
}

var myTmpl *template.Template

func init() {
	r := router.New()
	http.Handle("/", r)
	r.GET("/vote/:key", voteShow)
	r.GET("/add", addTemplate)
	r.GET("/process", addShow)
	r.GET("/process/:title", addShowDirect)
	r.GET("/", index)
	myTmpl = template.Must(template.ParseFiles("assets/index.html", "assets/add.html"))
}

func index(w http.ResponseWriter, r *http.Request, ps router.Params) {
	c := appengine.NewContext(r)
	var voteResults []Show

	//Retrieve Shows from datastore
	showKeys, err := datastore.NewQuery("Show").Order("-Votes").Order("Title").GetAll(c, &voteResults)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Errorf(c, "GetAll: %v", err)
		return
	}

	//Create input to pass to html template
	tempInput := make([]struct {
		Show Show
		Key  string
	}, len(showKeys))

	//Fill input object with data
	for key, val := range voteResults {
		tempInput[key].Key = showKeys[key].Encode()
		tempInput[key].Show = val
	}

	myTmpl.ExecuteTemplate(w, "index", tempInput)
}

//Cast a vote for the show
func voteShow(w http.ResponseWriter, r *http.Request, ps router.Params) {
	c := appengine.NewContext(r)

	//Decode the Key for the show being voted for
	keyStr := ps.ByName("key")
	k, err := datastore.DecodeKey(keyStr)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Errorf(c, "DecodeKey: %v", err)
		return
	}

	//Get Show from datastore
	var s Show
	err = datastore.Get(c, k, &s)
	if err == datastore.ErrNoSuchEntity {
		//Show not found in the datastore
		http.NotFound(w, r)
		return
	} else if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Errorf(c, "DecodeKey: %v", err)
		return
	}

	//Check if user has already cast a vote for this show
	u := user.Current(c)
	i, err := datastore.NewQuery("Vote").Filter("Key =", keyStr).Filter("User =", u.Email).Count(c)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Errorf(c, "Vote Validation: %v", err)
		return
	}
	if i == 0 {
		//Put new vote in datastore
		newVote := Vote{keyStr, u.Email}
		_, err = datastore.Put(c, datastore.NewIncompleteKey(c, "Vote", nil), &newVote)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			log.Errorf(c, "Put Vote: %v", err)
			return
		}
	}

	//Count number of Votes and update s.Votes
	i, err = datastore.NewQuery("Vote").Filter("Key =", keyStr).Count(c)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Errorf(c, "Vote Count: %v", err)
		return
	}
	s.Votes = i

	//Put updated show back in to datastore
	_, err = datastore.Put(c, k, &s)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Errorf(c, "Put Updated Show: %v", err)
		return
	}

	http.Redirect(w, r, "/", http.StatusFound)
}

//Respond with the template for adding a new show
func addTemplate(w http.ResponseWriter, r *http.Request, ps router.Params) {
	myTmpl.ExecuteTemplate(w, "add", nil)
}

//Redirect addShow request to other URL to addshow
func addShow(w http.ResponseWriter, r *http.Request, ps router.Params) {
	title := r.FormValue("title")
	http.Redirect(w, r, "process/"+title, http.StatusFound)
}

//Processing a submitted show to be added to the list
func addShowDirect(w http.ResponseWriter, r *http.Request, ps router.Params) {
	c := appengine.NewContext(r)

	title := ps.ByName("title")

	q := datastore.NewQuery("Show").Filter("Title =", title)

	i, err := q.Count(c)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Errorf(c, "Query Filter Title: %v", err)
		return
	}

	//Check if Show is already in datastore
	var kStr string
	if i == 0 {
		newShow := Show{title, 0}
		k, err := datastore.Put(c, datastore.NewIncompleteKey(c, "Show", nil), &newShow)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			log.Errorf(c, "Put New Show: %v", err)
			return
		}
		kStr = k.Encode()
	} else {
		//Get key for show that already exists
		k, err := q.GetAll(c, &[]Show{})
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			log.Errorf(c, "Query Show: %v", err)
			return
		}
		kStr = k[0].Encode()
	}

	http.Redirect(w, r, "/vote/"+kStr, http.StatusFound)
}

//TODO: Solve race condition on Show.Votes
//TODO: Add logout feature
//TODO: Add unvote feature (remove Vote entity from the datastore)
