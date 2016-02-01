package receipt_test

import (
	"github.com/cloudfoundry-incubator/bosh-certification-receipt-generator/artifact"
	"github.com/cloudfoundry-incubator/bosh-certification-receipt-generator/receipt"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Receipt", func() {
	Describe("New", func() {
		It("creates a new Receipt from valid Artifacts", func() {
			release, err := artifact.New("release/version")
			Expect(err).ToNot(HaveOccurred())

			stemcell, err := artifact.New("stemcell/version")
			Expect(err).ToNot(HaveOccurred())

			subject, err := receipt.New([]artifact.Artifact{release}, stemcell)
			Expect(err).ToNot(HaveOccurred())

			Expect(len(subject.Releases)).To(Equal(1))
			Expect(subject.Releases[0]).To(Equal(release))
			Expect(subject.Stemcell).To(Equal(stemcell))
		})

		It("returns an error when no releases are specified", func() {
			stemcell, err := artifact.New("stemcell/version")
			Expect(err).ToNot(HaveOccurred())

			subject, err := receipt.New([]artifact.Artifact{}, stemcell)
			Expect(subject).To(Equal(receipt.Receipt{}))
			Expect(err).To(MatchError("must include at least one release in receipt"))
		})

		It("returns an error when any artifact is invalid", func() {
			validRelease, err := artifact.New("release/version")
			Expect(err).ToNot(HaveOccurred())

			invalidRelease := artifact.Artifact{Name: "invalid"}
			Expect(err).ToNot(HaveOccurred())

			stemcell, err := artifact.New("stemcell/version")
			Expect(err).ToNot(HaveOccurred())

			subject, err := receipt.New([]artifact.Artifact{validRelease, invalidRelease}, stemcell)
			Expect(subject).To(Equal(receipt.Receipt{}))
			Expect(err).To(HaveOccurred())
			Expect(err.Error()).To(ContainSubstring("artifact info must specify a version"))
		})
	})
})
