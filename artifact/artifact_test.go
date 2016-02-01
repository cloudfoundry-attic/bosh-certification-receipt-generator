package artifact_test

import (
	"fmt"

	"github.com/cloudfoundry-incubator/bosh-certification-receipt-generator/artifact"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Artifact", func() {
	Describe("New", func() {
		It("creates a new Artifact from a string in `name/version` format", func() {
			exampleInfo := "name/version"
			a, err := artifact.New(exampleInfo)

			Expect(err).ToNot(HaveOccurred())
			Expect(a.Name).To(Equal("name"))
			Expect(a.Version).To(Equal("version"))
		})

		It("returns an error if the raw info is invalid", func() {
			invalidInfo := "name"
			a, err := artifact.New(invalidInfo)
			Expect(a).To(Equal(artifact.Artifact{}))
			Expect(err).To(MatchError("artifact info: `name` must have format: name/version"))

			invalidInfo = "name/version/something"
			a, err = artifact.New(invalidInfo)
			Expect(a).To(Equal(artifact.Artifact{}))
			Expect(err).To(MatchError("name and version for artifact info: `name/version/something` cannot contain `/`"))

			invalidInfo = "name/"
			a, err = artifact.New(invalidInfo)
			Expect(a).To(Equal(artifact.Artifact{}))
			Expect(err).To(MatchError("artifact info must specify a version"))

			invalidInfo = "/version"
			a, err = artifact.New(invalidInfo)
			Expect(a).To(Equal(artifact.Artifact{}))
			Expect(err).To(MatchError("artifact info must specify a name"))
		})
	})

	Describe("ValidateAll", func() {
		It("returns nil if all Artifacts are valid", func() {
			artifacts := []artifact.Artifact{}
			for i := 0; i < 2; i++ {
				artifact, err := artifact.New(fmt.Sprintf("name/%d", i+1))
				Expect(err).ToNot(HaveOccurred())
				artifacts = append(artifacts, artifact)
			}

			err := artifact.ValidateAll(artifacts...)
			Expect(err).ToNot(HaveOccurred())
		})

		It("returns an error message describing any/all invalid Artifacts", func() {
			artifacts := []artifact.Artifact{}
			for i := 0; i < 2; i++ {
				artifact, err := artifact.New(fmt.Sprintf("name/%d", i+1))
				Expect(err).ToNot(HaveOccurred())
				artifacts = append(artifacts, artifact)
			}

			artifacts = append(artifacts, artifact.Artifact{Name: "Foo"})

			err := artifact.ValidateAll(artifacts...)
			Expect(err).To(MatchError("artifact 3 of 3 failed validation: artifact info must specify a version"))
		})
	})
})
