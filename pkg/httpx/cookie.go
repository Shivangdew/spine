package httpx

import "time"

type Cookie struct {
	Name  string
	Value string

	Path    string
	Domain  string
	MaxAge  int
	Expires *time.Time

	HttpOnly bool
	Secure   bool
	SameSite SameSite

	Priority string // "Low" | "Medium" | "High"
}

type SameSite string

const (
	SameSiteLax    SameSite = "Lax"
	SameSiteStrict SameSite = "Strict"
	SameSiteNone   SameSite = "None"
)

const (
	AccessTokenCookieName  = "accessToken"
	RefreshTokenCookieName = "refreshToken"
)

func AccessTokenCookie(token string, ttl time.Duration) Cookie {
	return Cookie{
		Name:     AccessTokenCookieName,
		Value:    token,
		Path:     "/",
		HttpOnly: true,

		Secure:   true,
		SameSite: SameSiteNone,

		MaxAge: int(ttl.Seconds()),
	}
}

func RefreshTokenCookie(token string, ttl time.Duration) Cookie {
	return Cookie{
		Name:     RefreshTokenCookieName,
		Value:    token,
		Path:     "/",
		HttpOnly: true,

		Secure:   true,
		SameSite: SameSiteNone,

		MaxAge: int(ttl.Seconds()),
	}
}

func DefaultRefreshTokenCookie(token string) Cookie {
	return RefreshTokenCookie(token, 7*24*time.Hour)
}

func ClearAccessTokenCookie() Cookie {
	return Cookie{
		Name:   AccessTokenCookieName,
		Value:  "",
		Path:   "/",
		MaxAge: -1,
	}
}

func ClearRefreshTokenCookie() Cookie {
	return Cookie{
		Name:   RefreshTokenCookieName,
		Value:  "",
		Path:   "/",
		MaxAge: -1,
	}
}
