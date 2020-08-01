package discord

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"

	b64 "encoding/base64"

	"github.com/arturoguerra/destinyarena-api/internal/utils"
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
	"gopkg.in/go-playground/validator.v9"
)

type (
	ReqPayload struct {
		Code string `json:"code" validate:"required"`
	}

	RespOAuthPayload struct {
		AccessToken  string `json:"access_token"`
		TokenType    string `json:"token_type"`
		ExpiresIn    string `json:"expires_in"`
		RefreshToken string `json:"refresh_token"`
		Scope        string `json:"scope"`
	}

	User struct {
		ID            string `json:"id"`
		Username      string `json:"username"`
		Discriminator string `json:"discriminator"`
	}

	Claims struct {
		User
		jwt.StandardClaims
	}
)

func GetUser(token string) (*User, error) {
	client := new(http.Client)
	authtoken := fmt.Sprintf("%s %s", "Bearer", token)
	log.Debugln(authtoken)

	req, err := http.NewRequest("GET", fmt.Sprintf("%s/v6/users/@me", cfg.BaseURL), nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Authorization", authtoken)
	req.Header.Set("Content-Type", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		log.Error(err)
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Error(err)
		return nil, err
	}

	var user User
	json.Unmarshal(body, &user)

	return &user, nil
}

func getToken(p *ReqPayload) (*RespOAuthPayload, error) {
	client := new(http.Client)
	apiurl := fmt.Sprintf("%s/oauth2/token", cfg.BaseURL)

	data := url.Values{}
	data.Set("grant_type", "authorization_code")
	data.Set("code", p.Code)
	data.Set("redirect_uri", cfg.RedirectURI)
	data.Set("scope", "identify guilds")

	fmt.Println(apiurl)
	req, err := http.NewRequest("POST", apiurl, strings.NewReader(data.Encode()))
	if err != nil {
		log.Error(err)
		return nil, err
	}

	creds := b64.StdEncoding.EncodeToString([]byte(cfg.ClientID + ":" + secrets.Discord))
	req.Header.Set("Authorization", fmt.Sprintf("Basic %s", creds))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("Content-Length", strconv.Itoa(len(data.Encode())))

	resp, err := client.Do(req)
	if err != nil {
		log.Error(err)
		return nil, err
	}

	defer resp.Body.Close()

	authBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Error(err)
		return nil, err
	}

	fmt.Println(resp.StatusCode)
	fmt.Println(string(authBody))
	if resp.StatusCode != 200 {
		return nil, errors.New("Invalid error code")
	}

	var authPayload RespOAuthPayload
	json.Unmarshal(authBody, &authPayload)

	return &authPayload, nil
}

func Callback(c echo.Context) (err error) {
	payload := new(ReqPayload)
	if err = c.Bind(payload); err != nil {
		return c.String(http.StatusBadRequest, "Invalid code payload")
	}

	v := validator.New()
	if errs := v.Struct(payload); errs != nil {
		return c.String(http.StatusBadRequest, "Invalid Payload")
	}

	authPayload, err := getToken(payload)
	if err != nil {
		fmt.Println(err)
		return c.String(401, "Error while getting Auth token from discord")
	}

	accessToken := authPayload.AccessToken

	user, err := GetUser(accessToken)
	if err != nil {
		return c.String(http.StatusInternalServerError, "Discord API is probably down again")
	}

	err, ok := checkGuilds(accessToken)
	if !ok {
		return c.String(http.StatusUnauthorized, "Please join our server before attempting to register")
	} else if err != nil {
		log.Error(err)
		return c.String(http.StatusInternalServerError, "Well rip again")
	}

	claims := &Claims{
		User: *user,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * time.Duration(1)).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
	}

	token, err := utils.SignJWT(claims)
	if err != nil {
		return c.String(http.StatusInternalServerError, "Something went wrong while generation Token")
	}

	r := map[string]interface{}{
		"token": token,
	}

	return c.JSON(http.StatusOK, r)
}
