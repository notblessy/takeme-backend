package utils

import (
	"errors"
	"strconv"

	"github.com/sirupsen/logrus"
)

// DefaultSize :nodoc:
const DefaultSize int = 10

// DefaultPage :nodoc:
const DefaultPage int = 1

// ResponseWithPagination :nodoc:
type ResponseWithPagination struct {
	Records     interface{}            `json:"records"`
	PageSummary map[string]interface{} `json:"pageSummary"`
}

// WithPaging builds response with pagination
func WithPaging(result interface{}, total int64, pageRequest, sizeRequest string) ResponseWithPagination {
	var (
		page int
		size int
	)

	page, err := strconv.Atoi(pageRequest)
	if err != nil {
		page = DefaultPage
	}

	size, err = strconv.Atoi(sizeRequest)
	if err != nil {
		size = DefaultSize
	}

	offset := (page - 1) * size

	var hasNext bool
	if offset+size < int(total) {
		hasNext = true
	}

	return ResponseWithPagination{
		Records: result,
		PageSummary: map[string]interface{}{
			"size":    size,
			"page":    page,
			"hasNext": hasNext,
			"total":   total,
		},
	}
}

// GetPageAndSize :nodoc:
func GetPageAndSize(q map[string]string) (offset, size int) {
	logger := logrus.WithField("queries", Dump(q))
	if q == nil {
		logger.Error(errors.New("invalid queries").Error())
		return 0, 10
	}

	offset, err := strconv.Atoi(q["page"])
	if err != nil {
		logger.Error(err)
		offset = 0
	}

	size, err = strconv.Atoi(q["size"])
	if err != nil {
		logger.Error(err)
		size = 10
	}

	return (offset - 1) * size, size
}
