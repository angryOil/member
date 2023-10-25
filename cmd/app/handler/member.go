package handler

import (
	"github.com/gorilla/mux"
	"net/http"
)

type Handler struct {
}

// 카페에 요청을 받을것이므로 카페의 요청을 믿음(카페에서 권한 체크를 한후에 요청한다고 예상 )

// 가입한 카페 id list
// 카페 가입 요청
// 카페 가입 요청 관리
// 밴
// 밴리스트 확인

func NewHandler() Handler {
	m := mux.NewRouter()
	h := Handler{}

	// 가장 많이 사용될 기능 cafe 의 모든액션하기전 거쳐갈 기능임(auth 기능)
	m.HandleFunc("/member/info", h.getMyInfo).Methods(http.MethodGet)
	// 카페 가입 요청 ?cafeId=1111 이런식으로 요청받을 예정(url query)
	m.HandleFunc("/member/join/request", h.requestJoin).Methods(http.MethodPost)
	// 카페 가입 요청 관리페이지
	m.HandleFunc("/member/join", h.getJoinReqList).Methods(http.MethodGet)
	// 카페 가입 처리
	m.HandleFunc("/member/join", h.processJoinReq).Methods(http.MethodPatch)
	// 밴목록 확인
	m.HandleFunc("/member/ban", h.getBanedList).Methods(http.MethodGet)
	// 밴상태 수정
	m.HandleFunc("/member/ban", h.updateBan).Methods(http.MethodPatch)
	return h
}

func (h Handler) getMyInfo(w http.ResponseWriter, r *http.Request) {

}

func (h Handler) requestJoin(w http.ResponseWriter, r *http.Request) {

}

func (h Handler) processJoinReq(w http.ResponseWriter, r *http.Request) {

}

func (h Handler) getJoinReqList(w http.ResponseWriter, r *http.Request) {

}

func (h Handler) getBanedList(w http.ResponseWriter, r *http.Request) {

}

func (h Handler) updateBan(w http.ResponseWriter, r *http.Request) {

}
