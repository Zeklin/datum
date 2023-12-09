package route

import (
	"net/http"

	echo "github.com/datumforge/echox"
)

// VerifyEmail verifies a user's email address by validating the token in the request.
// This endpoint is intended to be called by frontend applications after the user has
// followed the link in the verification email. If the user is not verified and the
// token is valid then the user is logged in. If the user is already verified then a
// 204 response is returned.

// TODO: implement verify handler
func registerVerifyHandler(router *echo.Echo) (err error) { //nolint:unused
	_, err = router.AddRoute(echo.Route{
		Method: http.MethodGet,
		Path:   "/verify",
		Handler: func(c echo.Context) error {
			return c.JSON(http.StatusNotImplemented, echo.Map{
				"error": "Not implemented",
			})
		},
	})

	return
}
