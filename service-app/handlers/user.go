package handlers

//
//import (
//	"context"
//	"errors"
//	"fmt"
//	"log"
//	"net/http"
//	"service-app/auth"
//	"service-app/data/user"
//	"service-app/web"
//)
//
//type userHandlers struct {
//	*user.DbService
//	*auth.Auth
//}
//
//func (h *userHandlers) SignUp(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
//	v, ok := ctx.Value(web.KeyValue).(*web.Values)
//	if !ok {
//		return errors.New("value not found in the context")
//	}
//	var nu user.NewUser
//	err := web.Decode(r, &nu)
//	if err != nil {
//		return err
//	}
//	//impl it
//	//if nu.Name == "" || nu.Email == "" {
//	//
//	//}
//	usr, err := h.Create(ctx, nu, v.Now)
//
//	if err != nil {
//		return fmt.Errorf("user: %+v %w", &usr, err)
//	}
//
//	return web.Respond(ctx, w, usr, http.StatusCreated)
//
//}
//
//func (h *userHandlers) Login(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
//	v, ok := ctx.Value(web.KeyValue).(*web.Values)
//	if !ok {
//		return errors.New("value not found in the context")
//	}
//
//	var login struct {
//		Email    string `json:"email"`
//		Password string `json:"password"`
//	}
//
//	err := web.Decode(r, &login)
//	if err != nil {
//		return err
//	}
//
//	claims, err := h.Authenticate(ctx, v.Now, login.Email, login.Password)
//	log.Println(err)
//	if err != nil {
//		return web.NewRequestError(errors.New("please provide a valid email and password"), http.StatusUnauthorized)
//
//	}
//
//	var tkn struct {
//		Token string `json:"token"`
//	}
//
//	tkn.Token, err = h.GenerateToken(claims)
//
//	if err != nil {
//		return fmt.Errorf("generating token %w", err)
//	}
//	h.SetCookie(w, tkn.Token)
//
//	return web.Respond(ctx, w, tkn, http.StatusOK)
//
//}
//
//func (h *userHandlers) SetCookie(w http.ResponseWriter, token string) {
//
//	cookie := http.Cookie{
//		Name:     "token",
//		Value:    token,
//		HttpOnly: true,
//	}
//	http.SetCookie(w, &cookie)
//
//}
