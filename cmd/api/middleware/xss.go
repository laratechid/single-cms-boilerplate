package middleware

import (
	"net/http"
	"regexp"

	"github.com/gin-gonic/gin"
)

// Precompiled regular expressions for common security issues
var (
	// Regular expressions to detect potentially malicious content
	sqlInjectionPattern = regexp.MustCompile(`(?i)\b(union\b.*\bselect|select\b.*\bfrom|insert\b.*\binto|delete\b.*\bfrom|update\b.*\bset|drop\b.*\btable|alter\b.*\btable|create\b.*\btable|truncate\b.*\btable|;|--|/\*|\*/|\bOR\b.*\b1=1|\bAND\b.*\b1=1|=\s*NULL|\bWAITFOR\b.*\bDELAY)\b`)
	xssPattern          = regexp.MustCompile(`(?i)<\s*script|<\s*img|<\s*a\s*href|javascript:|vbscript:`)
	phpInjectionPattern = regexp.MustCompile(`(?i)(\.php|<\?php|\.\./)`)
	filePattern         = regexp.MustCompile(`(?i)\.(js|json|xml|html|php|asp|jsp|pl|py)$`)
)

func SecurityMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Iterate over all query parameters
		for _, values := range c.Request.URL.Query() {
			for _, value := range values {
				// Perform security checks on each query parameter value
				if isMalicious(value) {
					// Log the issue and return a 403 Forbidden response
					c.JSON(http.StatusForbidden, gin.H{
						"status":  "error",
						"message": "Forbidden: Malicious query parameter detected.",
					})
					c.Abort() // Stop further processing
					return
				}
			}
		}

		// If all checks passed, proceed with the request
		c.Next()
	}
}

// isMalicious checks if a query parameter value contains malicious patterns
func isMalicious(value string) bool {
	// Check against common injection patterns
	if sqlInjectionPattern.MatchString(value) || xssPattern.MatchString(value) || phpInjectionPattern.MatchString(value) || filePattern.MatchString(value) {
		return true
	}
	return false
}
