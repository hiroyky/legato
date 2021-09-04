package errors

import (
	"fmt"
	native "github.com/pkg/errors"
)

type ErrorCode string

const (
	TrackNotFoundError       ErrorCode = "TrackNotFoundError"
	AlbumNotFoundError       ErrorCode = "AlbumNotFoundError"
	AlbumArtistNotFoundError ErrorCode = "AlbumArtistNotFoundError"
	GenreNotFoundError       ErrorCode = "GenreNotFoundError"

	GetTrackFatal   ErrorCode = "GetTrackFatal"
	CountTrackFatal ErrorCode = "CountTrackFatal"
	GetAlbumFatal   ErrorCode = "GetAlbumFatal"
	CountAlbumFatal ErrorCode = "CountAlbumFatal"
)

func (code ErrorCode) ToString() string {
	return string(code)
}

type LegatoError interface {
	ErrorCode() ErrorCode
	Error() string
}

type legatoError struct {
	errorCode ErrorCode
	baseError error
}

func (e *legatoError) Error() string {
	if e.baseError == nil {
		return e.ErrorCode().ToString()
	}
	return fmt.Sprintf("%s %s", e.ErrorCode(), e.baseError.Error())
}

func (e *legatoError) ErrorCode() ErrorCode {
	return e.errorCode
}

func New(errorCode ErrorCode, baseError error) LegatoError {
	return &legatoError{
		baseError: baseError,
		errorCode: errorCode,
	}
}

func Wrap(err error, message string) error {
	return native.Wrap(err, message)
}

func Cause(err error) error {
	return native.Cause(err)
}

// ExtractLegatoError 与えられたerrorからLegatoError型を抽出します。抽出できなかった場合、nilが返ります。
func ExtractLegatoError(err error) LegatoError {
	e, ok := Cause(err).(LegatoError)
	if !ok {
		return nil
	}
	return e
}

// IsErrorCode 与えられたerrorのエラーコードが一致するかどうかを判定します。
func IsErrorCode(err error, code ErrorCode) bool {
	are := ExtractLegatoError(err)
	if are == nil {
		return false
	}
	return are.ErrorCode() == code
}
