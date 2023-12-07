package usecase

import (
	"context"
	"cosmart-library/domain"
	"cosmart-library/pkg/logger"
)

func (u *usecase) GetBookList(ctx context.Context, request domain.GetBookListRequest) (response domain.GetBookListResponse, err error) {
	if request.Subject == "" {
		logger.Warn("[GetBookList] subject is empty, referring to default subject")
		request.Subject = "fiction"
	}

	books, err := u.repository.FetchRawBooksBySubject(ctx, request.Subject)
	if err != nil {
		logger.Error("[GetBookList][FetchRawBooksBySubject] error calling repository", err.Error())
		return
	}

	response = domain.GetBookListResponse{
		Books: books,
		Page:  1,
	}

	return
}
