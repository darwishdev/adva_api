package auth

import (
	"errors"
	"time"

	"github.com/google/uuid" 
)

// Different types of error returned by the VerifyToken function
var (
	ErrInvalidToken = errors.New("token is invalid")
	ErrExpiredToken = errors.New("token has expired")
)

type Payload struct {
	ID          uuid.UUID                  `json:"id"`
	UserId      int32                      `json:"user_id"`
	Username    string                     `json:"username"`
	Permissions map[string]map[string]bool `json:"Authorities"`
	IssuedAt    time.Time                  `json:"issued_at"`
	ExpiredAt   time.Time                  `json:"expired_at"`
}

func NewPayload(username string, userID int32, permissions map[string]map[string]bool, duration time.Duration) (*Payload, error) {
	tokenID, err := uuid.NewRandom()
	if err != nil {
		return nil, err
	}

	payload := &Payload{
		ID:     tokenID,
		UserId: userID,

		Permissions: permissions,
		Username:    username,
		IssuedAt:    time.Now(),
		ExpiredAt:   time.Now().Add(duration),
	}

	return payload, nil
}

// Valid checks if the token payload is valid or not
func (payload *Payload) Valid() error {
	if time.Now().After(payload.ExpiredAt) {
		return ErrExpiredToken
	}
	return nil
}

// // Can checks if the user has certain Authorities
// func (payload *Payload) Can(group string, action string) bool {
// 	if payload.Permissions == nil {
// 		return false
// 	}
// 	return payload.Permissions[group][action]
// }

// func (payload *Payload) GetAccessableActionsForGroup(group string) *rmsv1.ListDataOptions {
// 	resp := rmsv1.ListDataOptions{
// 		Title:       fmt.Sprintf("%s_list", group),
// 		Description: fmt.Sprintf("%s_description", group),
// 	}
// 	var (
// 		singularizedGroup string = inflector.Singularize(group)
// 		redirectRoute     string = fmt.Sprintf("%s_list", group)
// 		requestProperty   string = fmt.Sprintf("%sId", singularizedGroup)
// 		create            string = fmt.Sprintf("%s_create", singularizedGroup)
// 		update            string = fmt.Sprintf("%s_update", singularizedGroup)
// 		deleteRestore     string = fmt.Sprintf("%s_delete_restore", singularizedGroup)
// 	)
// 	if payload.Permissions == nil {
// 		return &resp
// 	}

// 	authorities := payload.Permissions[group]
// 	log.Debug().Interface("perms", payload.Permissions).Interface("autho", group).Interface("creaet", create).Msg("jjhi")
// 	if len(authorities) == 0 {
// 		return &resp
// 	}
// 	log.Debug().Interface("create2", strcase.ToCamel(create)).Str("create", create).Msg("gello")

// 	if authorities[strcase.ToCamel(create)] {
// 		resp.CreateHandler = &rmsv1.CreateHandler{
// 			RedirectRoute: redirectRoute,
// 			Title:         create,
// 			Endpoint:      strcase.ToLowerCamel(create),
// 			RouteName:     create,
// 		}
// 		// resp.ImportHandler = &rmsv1.ImportHandler{
// 		// 	Endpoint:           importEndpoint,
// 		// 	ImportTemplateLink: importTemplateLink,
// 		// }
// 	}
// 	if authorities[strcase.ToCamel(update)] {
// 		resp.UpdateHandler = &rmsv1.UpdateHandler{
// 			RedirectRoute: redirectRoute,
// 			Title:         update,
// 			Endpoint:      strcase.ToLowerCamel(update),
// 			RouteName:     update,
// 		}
// 	}
// 	if authorities[strcase.ToCamel(deleteRestore)] {
// 		resp.DeleteRestoreHandler = &rmsv1.DeleteRestoreHandler{
// 			Endpoint:        strcase.ToLowerCamel(deleteRestore),
// 			RequestProperty: requestProperty,
// 		}
// 	}
// 	return &resp
// }
