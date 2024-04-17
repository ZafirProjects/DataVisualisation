package handler

import (
	"net/http"

	"github.com/ZafirProjects/QuodOrbisChallenge/view/welcome"
)

type WelcomeHandler struct{}

func (h *WelcomeHandler) HandleWelcomeRequest(w http.ResponseWriter, r *http.Request) {
	render(w, r, welcome.RenderWelcomePage())
}
