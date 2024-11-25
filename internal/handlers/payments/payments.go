package payments

import (
	"net/http"

	"static/internal/usecase/static"
	httpErr "static/pkg/http/error"
	"static/pkg/http/writer"
)

func GetPayments(uc static.PaymentsUseCase) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		payments, err := uc.GetPayments(r.Context())
		if err != nil {
			httpErr.InternalError(w, err)
			return
		}

		writer.WriteJson(w, payments)
	}
}
