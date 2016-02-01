package artifact

import (
	"errors"
	"fmt"
	"strings"
)

type Artifact struct {
	Name    string `json:"name"`
	Version string `json:"version"`
}

func New(raw string) (Artifact, error) {
	values := strings.Split(raw, "/")
	l := len(values)

	switch {
	case (l < 2):
		return Artifact{}, fmt.Errorf("artifact info: `%s` must have format: name/version", raw)
	case (l > 2):
		return Artifact{}, fmt.Errorf("name and version for artifact info: `%s` cannot contain `/`", raw)
	}

	a := Artifact{Name: values[0], Version: values[1]}
	err := a.validate()
	if err != nil {
		return Artifact{}, err
	}

	return a, nil
}

func (a *Artifact) validate() error {
	if a.Name == "" {
		return errors.New("artifact info must specify a name")
	}

	if a.Version == "" {
		return errors.New("artifact info must specify a version")
	}

	return nil
}

func ValidateAll(artifacts ...Artifact) error {
	var errMsgs []string

	for i := range artifacts {
		validationErr := artifacts[i].validate()
		if validationErr != nil {
			errMsgs = append(errMsgs, fmt.Sprintf("artifact %d of %d failed validation: %s", i+1, len(artifacts), validationErr))
		}
	}

	if len(errMsgs) == 0 {
		return nil
	}

	joinedMsg := strings.Join(errMsgs, "\n")

	return errors.New(joinedMsg)
}
