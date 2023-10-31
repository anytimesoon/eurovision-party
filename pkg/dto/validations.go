package dto

import (
	"fmt"
	"github.com/anytimesoon/eurovision-party/pkg/enum"
	"github.com/anytimesoon/eurovision-party/pkg/errs"
	"strings"
)

var topScore uint8 = 10

func isPresent(attr string, attrName string) string {
	if attr == "" {
		return attrName + " must not be blank"
	} else {
		return ""
	}
}

func isWithinRange(vote uint8, attrName string) string {
	if vote > topScore {
		return fmt.Sprintf("Your vote for %s is too high. The maximum is 10", attrName)
	} else if vote < 1 {
		return fmt.Sprintf("Your must vote for %s is too low. The minimum is 1", attrName)
	} else {
		return ""
	}
}

func isValidCat(cat string) string {
	switch cat {
	case enum.Song, enum.Costume, enum.Performance, enum.Props:
		return ""
	default:
		return cat + " is not a valid category"
	}
}

func messagesToError(messages []string) *errs.AppError {
	if len(messages) > 0 {
		return errs.NewInvalidError(strings.Join(messages, ","))
	} else {
		return nil
	}
}
