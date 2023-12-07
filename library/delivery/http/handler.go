package http

import (
	"cosmart-library/domain"
	"cosmart-library/pkg/logger"
	"net/http"
	"strings"

	"github.com/microcosm-cc/bluemonday"
)

// ArticleHandler  represent the httphandler for article
type libraryHandler struct {
	libraryUsecase domain.LibraryUsecaseInterface
	sanitizer      *bluemonday.Policy
}

func InitHTTPHandler(mux *http.ServeMux, uc domain.LibraryUsecaseInterface) {
	handler := &libraryHandler{
		libraryUsecase: uc,
		sanitizer:      bluemonday.UGCPolicy(),
	}

	mux.Handle("/book/list", domain.GET(handler.GetBookListHandler))
}

func (l *libraryHandler) GetBookListHandler(w http.ResponseWriter, r *http.Request) (resp domain.GeneralResponse, err error) {
	subject := strings.ToLower(l.sanitizer.Sanitize(r.URL.Query().Get("subject")))

	bookList, err := l.libraryUsecase.GetBookList(r.Context(), domain.GetBookListRequest{
		Subject: subject,
	})
	if err != nil {
		logger.Error("[GetBookListHandler][GetBookList] error calling usecase", err.Error())
		return
	}

	resp.Data = bookList

	return
}
