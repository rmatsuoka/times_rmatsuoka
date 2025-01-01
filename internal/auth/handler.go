package auth

import (
	"context"
	"net/http"
	"net/url"
	"os"

	"github.com/gorilla/sessions"
	oauth1 "github.com/rmatsuoka/dghubble-oauth1"
	"github.com/rmatsuoka/times_rmatsuoka/internal/currnet"
	"github.com/rmatsuoka/times_rmatsuoka/internal/users"
	"github.com/rmatsuoka/times_rmatsuoka/internal/x/xhttp"
)

var (
	sessionKey = os.Getenv("SESSION_KEY")

	consumerKey      = os.Getenv("CONSUMER_KEY")
	consumerSecret   = os.Getenv("CONSUMER_SECRET")
	callbackURL      = os.Getenv("CALLBACK_URL")
	requestTokenPath = "https://www.hatena.com/oauth/initiate"
	authorizePath    = "https://www.hatena.ne.jp/oauth/authorize"
	accessTokenPath  = "https://www.hatena.com/oauth/token"
)

var store = sessions.NewCookieStore([]byte(sessionKey))

type Auth struct {
	config       *oauth1.Config
	callbackFunc func(context.Context, HatenaMy) (users.ID, error)
}

func New() *Auth {
	return &Auth{
		config: &oauth1.Config{
			ConsumerKey:    consumerKey,
			ConsumerSecret: consumerSecret,
			CallbackURL:    callbackURL,
			Endpoint: oauth1.Endpoint{
				RequestTokenURL: requestTokenPath,
				AuthorizeURL:    authorizePath,
				AccessTokenURL:  accessTokenPath,
			},
			Realm:      "",
			Signer:     nil,
			Noncer:     nil,
			HTTPClient: http.DefaultClient,
		},
	}
}

func (a *Auth) Install(handle func(string, http.Handler)) {
	handleFunc := func(pattern string, h http.HandlerFunc) {
		handle(pattern, h)
	}

	handleFunc("GET /auth/signin", a.signin)
	handleFunc("GET /auth/callback", a.callback)
}

func (a *Auth) signin(w http.ResponseWriter, req *http.Request) {
	token, secret, err := a.config.RequestToken(url.Values{"scope": []string{"read_public"}})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	session, err := store.Get(req, "user")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	session.Values["request_token"] = token
	session.Values["request_secret"] = secret
	if err := session.Save(req, w); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	u, err := a.config.AuthorizationURL(token)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	http.Redirect(w, req, u.String(), http.StatusFound)
}

type SessionKey string

const (
	sessionKeyRequestToken    = "request_token"
	sessionKeyRequestSecret   = "request_secret"
	sessionKeyURLName         = "url_name"
	sessionKeyDisplayName     = "display_name"
	sessionKeyProfileImageURL = "profile_image_url"
)

func (a *Auth) callback(w http.ResponseWriter, req *http.Request) {
	_, verifier, err := oauth1.ParseAuthorizationCallback(req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	session, err := store.Get(req, "user")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	requestToken, ok := session.Values["request_token"]
	if !ok {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}
	delete(session.Values, "request_token")

	requestSecret, ok := session.Values["request_secret"]
	if !ok {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}
	delete(session.Values, "request_secret")

	accessToken, accessSecret, err := a.config.AccessToken(requestToken.(string), requestSecret.(string), verifier)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	client := &xhttp.Client{
		Client: a.config.Client(req.Context(), oauth1.NewToken(accessToken, accessSecret)),
	}

	var my HatenaMy
	err = client.GetJSON(req.Context(), "https://n.hatena.com/applications/my.json", &my)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	id, err := a.callbackFunc(req.Context(), my)
	session.Values["user_id"] = id

	if err := session.Save(req, w); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	http.Redirect(w, req, "/", http.StatusFound)
}

type HatenaMy struct {
	URLName         string `json:"url_name"`
	DisplayName     string `json:"display_name"`
	ProfileImageURL string `json:"profile_image_url"`
}

func AuthHandler(handler http.Handler, fallback func(w http.ResponseWriter, req *http.Request)) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		session, err := store.Get(req, "user")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		userIDStr, ok := session.Values["url_name"]
		if !ok {
			fallback(w, req)
			return
		}

		userID, ok := userIDStr.(string)
		if !ok {
			http.Error(w, "bad request", http.StatusBadRequest)
			return
		}

		ctx := currnet.ContextWithUserID(req.Context(), users.ID(userID))
		req = req.WithContext(ctx)
		handler.ServeHTTP(w, req)
	})
}
