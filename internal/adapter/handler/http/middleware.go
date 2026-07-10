package http

import (
	"net/http"
	"net/url"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/jenish-brainztechs/go-backend/internal/core/domain"
	"github.com/jenish-brainztechs/go-backend/internal/core/port"
)

const (
	authorizationHeaderKey  = "Authorization"
	authorizationType       = "Bearer"
	authorizationPayloadKey = "authorization_payload"
	sessionContextKey       = "session_state"
)

func CORSMiddleware(allowedOrigins string) gin.HandlerFunc {
	originMap := map[string]bool{}
	allowWildcard := false
	wildcardSuffixes := []string{}

	normalizeOrigin := func(value string) string {
		return strings.TrimRight(strings.TrimSpace(value), "/")
	}

	parseOriginHost := func(value string) string {
		if value == "" {
			return ""
		}
		if strings.Contains(value, "://") {
			if parsed, err := url.Parse(value); err == nil && parsed.Host != "" {
				host := parsed.Hostname()
				return strings.TrimSpace(host)
			}
		}
		return strings.TrimSpace(value)
	}

	addWildcardSuffix := func(raw string) {
		host := parseOriginHost(raw)
		if strings.HasPrefix(host, "*.") {
			suffix := strings.TrimPrefix(host, "*.")
			if suffix != "" {
				wildcardSuffixes = append(wildcardSuffixes, strings.ToLower(suffix))
			}
		}
	}

	for _, o := range strings.Split(allowedOrigins, ",") {
		trimmed := normalizeOrigin(o)
		if trimmed == "" {
			continue
		}
		if trimmed == "*" {
			allowWildcard = true
			continue
		}
		if strings.Contains(trimmed, "*.") {
			addWildcardSuffix(trimmed)
			continue
		}
		originMap[trimmed] = true
	}

	return func(c *gin.Context) {
		origin := normalizeOrigin(c.GetHeader("Origin"))
		allowed := false

		if allowWildcard {
			allowed = true
		} else if origin != "" && originMap[origin] {
			allowed = true
		} else if origin != "" && len(wildcardSuffixes) > 0 {
			if parsed, err := url.Parse(origin); err == nil {
				host := strings.ToLower(parsed.Hostname())
				for _, suffix := range wildcardSuffixes {
					if host != suffix && strings.HasSuffix(host, "."+suffix) {
						allowed = true
						break
					}
				}
			}
		} else if origin == "" {
			// no Origin header (curl/server-to-server). Not a browser CORS request.
			allowed = true
		}

		// If this is a browser request (Origin present) and it's not allowed, reject preflight
		if origin != "" && !allowed && c.Request.Method == http.MethodOptions {
			c.AbortWithStatus(http.StatusForbidden)
			return
		}

		if allowed {
			if allowWildcard {
				// IMPORTANT: if you ever use credentials (cookies), DO NOT use "*"
				c.Header("Access-Control-Allow-Origin", "*")
			} else if origin != "" {
				c.Header("Access-Control-Allow-Origin", origin)
				c.Header("Access-Control-Allow-Credentials", "true")
				c.Header("Vary", "Origin")
			}

			// Allow headers: echo what browser requested to avoid missing-header failures
			if reqHeaders := c.GetHeader("Access-Control-Request-Headers"); reqHeaders != "" {
				c.Header("Access-Control-Allow-Headers", reqHeaders)
				c.Header("Vary", "Origin, Access-Control-Request-Headers")
			} else {
				c.Header("Access-Control-Allow-Headers", "Content-Type, Authorization, X-Requested-With")
			}

			c.Header("Access-Control-Allow-Methods", "GET, POST, PUT, PATCH, DELETE, OPTIONS")
			c.Header("Access-Control-Max-Age", "3600")
		}

		if c.Request.Method == http.MethodOptions {
			c.AbortWithStatus(http.StatusNoContent)
			return
		}

		c.Next()
	}
}
func authMiddleware(token port.TokenService) gin.HandlerFunc {
	return func(ctx *gin.Context) {

		authorizationHeader := ctx.GetHeader(authorizationHeaderKey)
		isEmpty := len(authorizationHeader) == 0

		if isEmpty {
			handleAbort(ctx, domain.ErrEmptyAuthorizationHeader)
			return

		}

		fields := strings.Fields(authorizationHeader)

		isValid := len(fields) == 2

		if !isValid {
			handleAbort(ctx, domain.ErrInvalidAuthorizationHeader)
			return
		}
		currentAuthorizationtype := fields[0]

		if currentAuthorizationtype != authorizationType {
			handleAbort(ctx, domain.ErrInvalidAuthorizationType)
			return

		}
		accessToken := fields[1]

		payload, err := token.VerifyAccessToken(accessToken)

		if err != nil {
			handleAbort(ctx, err)
			return

		}
		ctx.Set(authorizationPayloadKey, payload)
		ctx.Next()

	}
}

func adminMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		payload, exists := ctx.Get(authorizationPayloadKey)
		if !exists {
			validationError(ctx, domain.ErrEmptyAuthorizationHeader)
			return
		}

		userPayload, ok := payload.(*domain.TokenPayload)
		if !ok {
			validationError(ctx, domain.ErrInvalidAuthorizationHeader)
			return
		}

		if userPayload.RoleName != "ROLE_ADMIN" {
			handleAbort(ctx, domain.ErrUnauthorized)
			return

		}
		ctx.Next()

	}
}
