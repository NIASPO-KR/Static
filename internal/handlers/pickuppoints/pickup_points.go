package pickuppoints

import (
	"net/http"

	"static/internal/usecase"
	httpErr "static/pkg/http/error"
	"static/pkg/http/writer"
)

func GetPickupPoints(uc usecase.PickupPointsUseCase) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		pickupPoints, err := uc.GetPickupPoints(r.Context())
		if err != nil {
			httpErr.InternalError(w, err)
			return
		}

		writer.WriteJson(w, pickupPoints)
	}
}
