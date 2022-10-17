package middleware

import (
	"crypto/sha256"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)


/*// AuthLogin checks if the token avaliable
func AuthLogin() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		_, err := auth.ValidateJwt(ctx)
		if err != nil {
			ctx.AbortWithError(http.StatusUnauthorized, err)
			return
		}

		ctx.Next()
	}
}
*/

// AuthUser checks if token is avaliable & issuer is self
// func AuthUser() gin.HandlerFunc {
// 	return func(ctx *gin.Context) {

// 		_, err := auth.ValidateJwt(ctx)
// 		if err != nil {
// 			ctx.AbortWithError(http.StatusUnauthorized, err)
// 			return
// 		}
// 		ctx.Next()
// 	}
// }

// // AuthAdmin checks if token is avaliable & issuer is Admin
// func AuthAdmin() gin.HandlerFunc {
// 	return func(ctx *gin.Context) {
// 		me, err := auth.ValidateJwt(ctx)
// 		if err != nil {
// 			ctx.AbortWithError(http.StatusUnauthorized, err)
// 			return
// 		}

// 		logging.Info("%v", me)
// 		if me.IsAdmin != true {
// 			ctx.AbortWithError(http.StatusUnauthorized, errors.New("user not admin"))
// 			return
// 		}
// 		ctx.Next()
// 	}
// }

// AuthTimeDigest authenticates based on a time-based hash.
func AuthTimeDigest() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// Obtain secret digest string.
		secret := ctx.Query("secret")
		if len(secret) <= 0 {
			ctx.AbortWithError(http.StatusUnauthorized,
				errors.New("missing query parameter"))
			return
		}

		// Generate digest from current Unix timestamp.
		unit := int64(8 * 60 * 60) // number of seconds in 8 hours.
		mealsSinceEpoch := time.Now().Unix() / unit
		plain := fmt.Sprintf("once upon a time %d meals ago", mealsSinceEpoch)
		digest := fmt.Sprintf("%x", sha256.Sum256([]byte(plain)))

		// Check if digest strings match.
		if secret != digest {
			ctx.AbortWithError(http.StatusUnauthorized,
				errors.New("invalid secret"))
			return
		}
		ctx.Next()
	}
}