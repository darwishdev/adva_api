package auth

import "time"
 
type Maker interface {
	CreateToken(username string , userId int32 , permissions map[string]map[string]bool , duration time.Duration) (string, *Payload, error) 
	VerifyToken(token string) (*Payload, error)
}
