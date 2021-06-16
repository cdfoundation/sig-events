package extensions

import (
	"github.com/cloudevents/sdk-go/v2/binding"
	"github.com/cloudevents/sdk-go/v2/types"
)

const (
	ArtifactIdExtension      = "artifactid"
	ArtifactNameExtension    = "artifactname"
	ArtifactVersionExtension = "artifactversion"
)

// ArtifactExtension represents the extension for extension context
type ArtifactExtension struct {
	ArtifactId      string `json:"artifactid"`
	ArtifactName    string `json:"artifactname"`
	ArtifactVersion string `json:"artifactversion"`
}

func (ae *ArtifactExtension) ReadTransformer() binding.TransformerFunc {
	return func(reader binding.MessageMetadataReader, writer binding.MessageMetadataWriter) error {
		artifactIdExtension := reader.GetExtension(ArtifactIdExtension)
		if artifactIdExtension != nil {
			aeiFormatted, err := types.Format(artifactIdExtension)
			if err != nil {
				return err
			}
			ae.ArtifactId = aeiFormatted
		}
		artifactNameExtension := reader.GetExtension(ArtifactNameExtension)
		if artifactNameExtension != nil {
			aenFormatted, err := types.Format(artifactNameExtension)
			if err != nil {
				return err
			}
			ae.ArtifactName = aenFormatted
		}
		artifactVersionExtension := reader.GetExtension(ArtifactVersionExtension)
		if artifactVersionExtension != nil {
			aevFormatted, err := types.Format(artifactVersionExtension)
			if err != nil {
				return err
			}
			ae.ArtifactVersion = aevFormatted
		}
		return nil
	}
}

func (ae *ArtifactExtension) WriteTransformer() binding.TransformerFunc {
	return func(reader binding.MessageMetadataReader, writer binding.MessageMetadataWriter) error {
		err := writer.SetExtension(ArtifactIdExtension, ae.ArtifactId)
		if err != nil {
			return nil
		}
		err = writer.SetExtension(ArtifactNameExtension, ae.ArtifactName)
		if err != nil {
			return nil
		}
		err = writer.SetExtension(ArtifactVersionExtension, ae.ArtifactVersion)
		if err != nil {
			return nil
		}
		return nil
	}
}
