// util/validation.go

package util

import "regexp"

// ValidateEmail 함수는 전달된 이메일이 유효한지 검증합니다.
func ValidateEmail(email string) bool {
    emailRegex := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
    return emailRegex.MatchString(email)
}

// SanitizeString 함수는 전달된 문자열에서 특수문자를 제거합니다.
func SanitizeString(str string) string {
    specialCharsRegex := regexp.MustCompile(`[^\w\d\s]+`)
    return specialCharsRegex.ReplaceAllString(str, "")
}
