package handler

import (
	"github.com/gorilla/mux"
	"member/internal/controller"
	"net/http"
)

type Handler struct {
	c controller.MemberController
}

// 카페에 요청을 받을것이므로 카페의 요청을 믿음(카페에서 권한 체크를 한후에 요청한다고 예상 )

// 가입한 카페 id list
// 카페 가입 요청
// 카페 가입 요청 관리
// 밴
// 밴리스트 확인

func NewHandler(c controller.MemberController) http.Handler {
	m := mux.NewRouter()
	h := Handler{c: c}

	// 가장 많이 사용될 기능 cafe 의 모든액션하기전 거쳐갈 기능임(auth 기능)
	m.HandleFunc("/member/{cafeId:[0-9]+}/info", h.getMyInfo).Methods(http.MethodGet)
	// 가입한 멤버 리스트 조회
	m.HandleFunc("/member/{cafeId:[0-9]+}", h.getMemberList).Methods(http.MethodGet)
	// 카페가입 요청
	m.HandleFunc("/member/{cafeId:[0-9]+}/join/request", h.requestJoin).Methods(http.MethodPost)
	// 밴목록 확인
	m.HandleFunc("/member/{cafeId:[0-9]+}/ban", h.getBanedList).Methods(http.MethodGet)
	// 밴상태 수정
	m.HandleFunc("/member/{cafeId:[0-9]+}/ban", h.updateBan).Methods(http.MethodPatch)

	// 카페 가입 요청 관리페이지
	//m.HandleFunc("/member/join", h.getJoinReqList).Methods(http.MethodGet)
	// 카페 가입 처리
	//m.HandleFunc("/member/join", h.processJoinReq).Methods(http.MethodPatch)
	return m
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

func (h Handler) getMemberList(w http.ResponseWriter, r *http.Request) {

}