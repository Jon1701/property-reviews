package middleware

import (
	"net/http"
	"strconv"
	"strings"

	"github.com/Jon1701/property-reviews/app/errormessages"
	"github.com/gin-gonic/gin"
)

// Sanitizes the pagination parameters from the query string.
func SanitizePaginationParameters() gin.HandlerFunc {
	return func(c *gin.Context) {
		afterID := strings.Trim(c.Query("afterID"), " ")
		beforeID := strings.Trim(c.Query("beforeID"), " ")
		limitStr := strings.Trim(c.Query("limit"), " ")
		limit, err := strconv.ParseUint(limitStr, 10, 32)

		if limitStr != "" {
			// Unable to parse.
			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{
					"message": errormessages.InvalidPaginationLimit,
				})
				c.Abort()
				return
			}

			// Check value range.
			isLimitValid := limit >= errormessages.PaginationLimitMinValue && limit <= errormessages.PaginationLimitMaxValue
			if !isLimitValid {
				c.JSON(http.StatusBadRequest, gin.H{
					"message": errormessages.InvalidPaginationLimit,
				})
				c.Abort()
				return
			}
		}

		c.Set("afterID", afterID)
		c.Set("beforeID", beforeID)
		c.Set("limit", limit)

		c.Next()
	}
}
