package mfa

import (
	"fmt"
	"strings"

	"github.com/aws/aws-sdk-go/aws/arn"
	"github.com/pkg/errors"
)

const (
	errParseARN = "failed to parse %q as ARN"
)

// callerIdentityToSerial a caller identity ARN to a MFA serial
func callerIdentityToSerial(i *string) (string, error) {

	a, err := arn.Parse(*i)

	if err != nil {
		return "", errors.Wrapf(err, errParseARN, *i)
	}

	return strings.Replace(a.String(), ":user/", ":mfa/", 1), nil

}

// serialToName converts a MFA serial to a source name
func serialToName(i *string) (string, error) {

	a, err := arn.Parse(*i)

	if err != nil {
		return "", errors.Wrapf(err, errParseARN, *i)
	}

	var (
		issuer = fmt.Sprintf("aws/iam/%s", a.AccountID)
		name   = strings.TrimPrefix(a.Resource, "mfa/")
	)

	return strings.Join([]string{
		issuer,
		name,
	}, ":"), nil

}
