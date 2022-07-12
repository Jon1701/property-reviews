package validation

import (
	"net/url"

	"github.com/Jon1701/property-reviews/app/errormessages"
)

// Checks if a given URL is valid.
func IsValidURL(s string) bool {
	_, err := url.ParseRequestURI(s)

	return err == nil
}

// Check if a given URL has the correct length.
func IsURLOfCorrectLength(url string) bool {
	return len(url) >= errormessages.WebsiteMinLength && len(url) <= errormessages.WebsiteMaxLength
}
