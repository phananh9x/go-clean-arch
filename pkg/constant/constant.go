package constant

const (
	CodeSuccess             = 200
	CodeCreated             = 201
	CodeNoContent           = 204
	CodeBadRequest          = 400
	CodeUnAuthorize         = 401
	CodeForbidden           = 403
	CodeNotFound            = 404
	CodeMethodNotAllowed    = 405
	CodeUnprocessableEntity = 422
	CodeTooManyRequest      = 429
	CodeInternalServerError = 500
	CodeServerUnavailable   = 503
)

const (
	_      = iota //blank identifier
	KB int = 1 << (10 * iota)
	MB
	GB
	TB
	PB
)

const (
	//PartnerStatusPending ...
	PartnerStatusPending = iota
	//PartnerStatusSuccessful ...
	PartnerStatusSuccessful
	//PartnerStatusProcessing ...
	PartnerStatusProcessing
	//PartnerStatusNotExistedTrans ...
	PartnerStatusNotExistedTrans
	//PartnerStatusFailure ...
	PartnerStatusFailure
	//PartnerStatusBlockPrefix ...
	PartnerStatusBlockPrefix
)

const (
	MessageSuccess        = "Success"
	MessageSomethingError = "Something error"
)
