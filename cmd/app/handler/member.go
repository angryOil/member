package handler

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"member/internal/controller"
	"member/internal/controller/req"
	page2 "member/page"
	"net/http"
	"strconv"
	"strings"
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
	m.HandleFunc("/members/{cafeId:[0-9]+}/info/{userId:[0-9]+}", h.getMemberInfo).Methods(http.MethodGet)

	// 내 카페 id 리스트 조회
	m.HandleFunc("/members/list/{userId:[0-9]+}", h.getMyCafeIdList).Methods(http.MethodGet)

	// 카페가입 요청
	m.HandleFunc("/members/{cafeId:[0-9]+}/join/{userId:[0-9]+}", h.requestJoin).Methods(http.MethodPost)

	// 관리자 기능

	// 가입한 멤버 리스트 조회
	m.HandleFunc("/members/admin/{cafeId:[0-9]+}", h.getMemberList).Methods(http.MethodGet)
	// 밴목록 확인
	m.HandleFunc("/members/admin/{cafeId:[0-9]+}/ban", h.getBanedList).Methods(http.MethodGet)
	// 밴상태 수정
	m.HandleFunc("/members/admin/{cafeId:[0-9]+}/ban", h.updateBan).Methods(http.MethodPatch)

	// 카페 가입 요청 관리페이지
	//m.HandleFunc("/member/join", h.getJoinReqList).Methods(http.MethodGet)
	// 카페 가입 처리
	//m.HandleFunc("/member/join", h.processJoinReq).Methods(http.MethodPatch)
	return m
}

func (h Handler) getMemberInfo(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userId, err := strconv.Atoi(vars["userId"])
	if err != nil {
		http.Error(w, "invalid user id", http.StatusBadRequest)
		return
	}
	cafeId, err := strconv.Atoi(vars["cafeId"])
	if err != nil {
		http.Error(w, "invalid cafe id", http.StatusBadRequest)
	}
	dto, err := h.c.GetMemberInfo(r.Context(), cafeId, userId)
	if err != nil {
		if strings.Contains(err.Error(), "invalid") {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}

	data, err := json.Marshal(dto)
	if err != nil {
		log.Println("getMemberInfo marshal err: ", err)
	}
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(data)
}

func (h Handler) getMyCafeIdList(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userId, err := strconv.Atoi(vars["userId"])
	if err != nil {
		http.Error(w, "invalid user id ", http.StatusBadRequest)
		return
	}
	page, err := strconv.Atoi(r.URL.Query().Get("page"))
	if err != nil {
		page = 0
	}
	size, err := strconv.Atoi(r.URL.Query().Get("size"))
	if err != nil {
		size = 0
	}
	idsDto, err := h.c.GetJoinCafeIds(r.Context(), userId, page2.NewReqPage(page, size))
	data, err := json.Marshal(idsDto)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(data)
}

func (h Handler) requestJoin(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userId, err := strconv.Atoi(vars["userId"])
	if err != nil {
		http.Error(w, "invalid userId", http.StatusUnauthorized)
		return
	}
	cafeId, err := strconv.Atoi(vars["cafeId"])
	if err != nil {
		http.Error(w, "invalid cafeId", http.StatusBadRequest)
		return
	}

	var dto req.JoinMemberDto
	err = json.NewDecoder(r.Body).Decode(&dto)
	if err != nil {
		http.Error(w, "json decode error", http.StatusBadRequest)
		return
	}
	err = h.c.RequestJoin(r.Context(), dto, cafeId, userId)
	if err != nil {
		if strings.Contains(err.Error(), "invalid") || strings.Contains(err.Error(), "empty") {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		if strings.Contains(err.Error(), "duplicate") {
			http.Error(w, err.Error(), http.StatusConflict)
			return
		}
		log.Println("requestJoin err :", err)
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
}

// 관리자  cafe API 에서 자체적으로 올바른 요청인지(권한)을 확인해야함
func (h Handler) getMemberList(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	cafeId, err := strconv.Atoi(vars["cafeId"])
	if err != nil {
		http.Error(w, "invalid cafe id ", http.StatusBadRequest)
		return
	}

	page, err := strconv.Atoi(r.URL.Query().Get("page"))
	if err != nil {
		page = 0
	}
	size, err := strconv.Atoi(r.URL.Query().Get("size"))
	if err != nil {
		size = 0
	}
	reqPage := page2.NewReqPage(page, size)
	memberInfoListCountDto, err := h.c.GetMemberList(r.Context(), cafeId, reqPage)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	data, err := json.Marshal(&memberInfoListCountDto)
	if err != nil {
		log.Println("getMemberList marshal err: ", err)
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(data)
}

func (h Handler) getBanedList(w http.ResponseWriter, r *http.Request) {

}

func (h Handler) updateBan(w http.ResponseWriter, r *http.Request) {

}
