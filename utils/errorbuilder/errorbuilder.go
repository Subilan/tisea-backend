package errorbuilder

import "tisea-backend/utils/common"

type CommonError struct {
	Description string `json:"desc"`
	OccurredAt  int64  `json:"at"`
	Internal    string  `json:"internal"`
}

func New(desc string, internal error) CommonError {
	return CommonError{Description: desc, OccurredAt: common.GetTimestampMilli(), Internal: internal.Error()}
}

func NewWithoutInternal(desc string) CommonError {
	return CommonError{Description: desc, OccurredAt: common.GetTimestampMilli(), Internal: ""}
}

func NewWithoutDescription(internal error) CommonError {
	return CommonError{Description: "N/A", OccurredAt: common.GetTimestampMilli(), Internal: internal.Error()}
}
