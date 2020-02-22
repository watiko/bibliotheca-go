package auth

var (
	ExportEmailFromTokenClaims                    = emailFromTokenClaims
	ExportAuthWithJWTMiddleware                   = authWithJWTMiddleware
	ExportNewFirebaseKeyGetter                    = newFirebaseKeyGetter
	ExportNewValidationKeyGetter                  = newValidationKeyGetter
	ExportNewJWTMiddlewareWithValidationKeyGetter = newJWTMiddlewareWithValidationKeyGetter
)
