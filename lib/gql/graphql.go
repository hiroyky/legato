package gql

import (
	"encoding/base64"
	"fmt"
	"github.com/hiroyky/legato/domain/errors"
	"strconv"
	"strings"
)

func EncodeID(filedName string, rawID int) string {
	if strings.Contains(filedName, ":") {
		panic(fmt.Sprintf("Invalid field name: %s", filedName))
	}
	format := fmt.Sprintf("%s:%d", filedName, rawID)
	return base64.StdEncoding.EncodeToString([]byte(format))
}

func DecodeID(graphID string) (int, error) {
	decoded, err := base64.StdEncoding.DecodeString(graphID)
	if err != nil {
		return 0, errors.Wrap(err, fmt.Sprintf("Failed decode GraphQL ID: %s", graphID))
	}
	parts := strings.Split(string(decoded), ":")
	if len(parts) != 2 {
		return 0, errors.Wrap(err, fmt.Sprintf("Failed decode GraphQL ID: %s", graphID))
	}

	id, err := strconv.Atoi(parts[1])
	if err != nil {
		return 0, errors.Wrap(err, fmt.Sprintf("Failed decode GraphQL ID: %s", graphID))
	}
	return id, nil
}

func DecodeIDIfNotNil(graphID *string) (*int, error) {
	if graphID == nil {
		return nil, nil
	}

	id, err := DecodeID(*graphID)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Failed decode GraphQL ID: %s", *graphID))
	}
	return &id, nil
}

func DecodedIDIntPtr(graphID string) (*int, error) {
	decoded, err := DecodeID(graphID)
	if err != nil {
		return nil, err
	}
	intID := int(decoded)
	return &intID, nil
}
