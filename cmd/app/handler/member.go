package handler

import (
	"github.com/gorilla/mux"
	"net/http"
)

type Handler struct {
}

// 가입한 카페 id list
// 카페 가입 요청
// 카페 가입 요청 관리
// 밴
//

func NewHandler() Handler {
	m := mux.NewRouter()
	h := Handler{}
	// 카페 가입 요청 ?cafeId=1111 이런식으로 요청받을 예정(url query)
	m.HandleFunc("/member/join", h.applyRequest).Methods(http.MethodPost)
	
	m.HandleFunc("/member", h.applyRequest).Methods(http.MethodPost)
	return h
}

func (h Handler) applyRequest(w http.ResponseWriter, r *http.Request) {

}
