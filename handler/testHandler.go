package handler

import (
	"net/http"

	"github.com/ZafirProjects/QuodOrbisChallenge/view/test"
)

type TestHandler struct{}

func (h *TestHandler) HandleTestShow(w http.ResponseWriter, r *http.Request) {
	render(w, r, test.RenderTest())
}
