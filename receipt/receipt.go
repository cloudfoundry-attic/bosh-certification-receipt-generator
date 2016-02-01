package receipt

import (
	"errors"

	"github.com/cloudfoundry-incubator/bosh-certification-receipt-generator/artifact"
)

type Receipt struct {
	Releases []artifact.Artifact `json:"releases"`
	Stemcell artifact.Artifact   `json:"stemcell"`
}

func New(releases []artifact.Artifact, stemcell artifact.Artifact) (Receipt, error) {
	allArtifacts := releases
	allArtifacts = append(allArtifacts, stemcell)

	if len(releases) == 0 {
		return Receipt{}, errors.New("must include at least one release in receipt")
	}

	err := artifact.ValidateAll(allArtifacts...)
	if err != nil {
		return Receipt{}, err
	}

	return Receipt{Releases: releases, Stemcell: stemcell}, nil
}
