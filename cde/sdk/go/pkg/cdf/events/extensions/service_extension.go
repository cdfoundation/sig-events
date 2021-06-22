package extensions

import (
	"github.com/cloudevents/sdk-go/v2/binding"
	"github.com/cloudevents/sdk-go/v2/types"
)

const (
	ServiceEnvIdExtension   = "serviceenvid"
	ServiceNameExtension    = "servicename"
	ServiceVersionExtension = "serviceversion"
)

// ServiceExtension represents the extension for extension context
type ServiceExtension struct {
	ServiceEnvId   string `json:"serviceenvid"`
	ServiceName    string `json:"servicename"`
	ServiceVersion string `json:"serviceversion"`
}

func (ae *ServiceExtension) ReadTransformer() binding.TransformerFunc {
	return func(reader binding.MessageMetadataReader, writer binding.MessageMetadataWriter) error {
		ServiceEnvIdExtension := reader.GetExtension(ServiceEnvIdExtension)
		if ServiceEnvIdExtension != nil {
			aeiFormatted, err := types.Format(ServiceEnvIdExtension)
			if err != nil {
				return err
			}
			ae.ServiceEnvId = aeiFormatted
		}
		ServiceNameExtension := reader.GetExtension(ServiceNameExtension)
		if ServiceNameExtension != nil {
			aenFormatted, err := types.Format(ServiceNameExtension)
			if err != nil {
				return err
			}
			ae.ServiceName = aenFormatted
		}
		ServiceVersionExtension := reader.GetExtension(ServiceVersionExtension)
		if ServiceVersionExtension != nil {
			aevFormatted, err := types.Format(ServiceVersionExtension)
			if err != nil {
				return err
			}
			ae.ServiceVersion = aevFormatted
		}
		return nil
	}
}

func (ae *ServiceExtension) WriteTransformer() binding.TransformerFunc {
	return func(reader binding.MessageMetadataReader, writer binding.MessageMetadataWriter) error {
		err := writer.SetExtension(ServiceEnvIdExtension, ae.ServiceEnvId)
		if err != nil {
			return nil
		}
		err = writer.SetExtension(ServiceNameExtension, ae.ServiceName)
		if err != nil {
			return nil
		}
		err = writer.SetExtension(ServiceVersionExtension, ae.ServiceVersion)
		if err != nil {
			return nil
		}
		return nil
	}
}
