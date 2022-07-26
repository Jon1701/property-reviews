package controllers

import (
	"github.com/gin-gonic/gin"
)

type PaginationParameters struct {
	AfterID  *string
	BeforeID *string
	Limit    *uint64
}

// Serializes pagination parameters from Context into a struct.
func SerializePaginationParameters(c *gin.Context, params *PaginationParameters) {
	values := c.Keys

	if values["afterID"] == nil || values["afterID"] == "" {
		params.AfterID = nil
	} else {
		v := string(values["afterID"].(string))
		params.AfterID = &v
	}

	if values["beforeID"] == nil || values["beforeID"] == "" {
		params.BeforeID = nil
	} else {
		v := string(values["beforeID"].(string))
		params.BeforeID = &v
	}

	if values["limit"] == nil || values["limit"] == "" {
		params.Limit = nil
	} else {
		v := uint64(values["limit"].(uint64))
		params.Limit = &v
	}
}
