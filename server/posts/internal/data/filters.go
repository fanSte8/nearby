package data

import "nearby/common/validator"

type Pagination struct {
	Page     int
	PageSize int
}

func ValidatePagination(v *validator.Validator, p Pagination) {
	v.Check(p.Page > 0, "page", "must be greater than zero")
	v.Check(p.PageSize > 0, "pageSize", "must be greater than zero")
}

func (p Pagination) limit() int {
	return p.PageSize
}

func (p Pagination) offset() int {
	return (p.Page - 1) * p.PageSize
}
