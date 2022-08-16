package errors

import (
	"regexp"
)

// https://www.postgresql.org/docs/9.4/errcodes-appendix.html

func IsViolatingUniqueConstraintError(err error) bool {
	unique := regexp.MustCompile(`\(SQLSTATE 23505\)$`)
	return unique.MatchString(err.Error())
}

func IsViolatingNotNullConstraintError(err error) bool {
	notNull := regexp.MustCompile(`\(SQLSTATE 23502\)$`)
	return notNull.MatchString(err.Error())
}

func IsViolatingForeignKeyConstraintError(err error) bool {
	foreignKey := regexp.MustCompile(`\(SQLSTATE 23503\)$`)
	return foreignKey.MatchString(err.Error())
}
