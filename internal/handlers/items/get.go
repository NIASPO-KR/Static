package items

import (
	"net/http"

	"static/internal/usecase"
	httpErr "static/pkg/http/error"
	"static/pkg/http/writer"
)

func GetItems(uc usecase.ItemsUseCase) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		items, err := uc.GetItems(r.Context())
		if err != nil {
			httpErr.InternalError(w, err)
			return
		}

		writer.WriteJson(w, items)
	}
}
