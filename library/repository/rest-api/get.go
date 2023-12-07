package restapi

import (
	"context"
	"cosmart-library/domain"
	"cosmart-library/library/types"
	"cosmart-library/pkg/logger"
	"encoding/json"
	"fmt"
	"net/http"
)

func (a *api) FetchBookListBySubject(ctx context.Context, subject string) (books []domain.Book, err error) {
	url := fmt.Sprintf("%s/subjects/%s.json", a.baseURL, subject)
	fmt.Println(url)
	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	response, err := a.http.Do(request)
	if err != nil {
		logger.Error("[FetchBookListBySubject] error http request", err.Error())
		return nil, err
	}
	defer response.Body.Close()

	data := types.BookAPIResponse{}

	err = json.NewDecoder(response.Body).Decode(&data)
	if err != nil {
		return nil, err
	}

	for _, book := range data.Books {
		books = append(books, domain.Book{
			ID:      book.Key,
			Subject: data.Name,
			Title:   book.Title,
			Author: func() []string {
				var authors []string
				for _, a := range book.Authors {
					authors = append(authors, a.Name)
				}
				return authors
			}(),
			EditionNumber: book.Edition,
		})
	}

	return
}
