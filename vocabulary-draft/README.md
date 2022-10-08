***Note! This draft vocabulary/specification is no longer maintained. Please go to the official [CDEvents website](https://cdevents.dev/) to follow the progress there.***

# Continuous Delivery Events

Continuous Delivery Events (CDEvents) is a common event format for the Continuous Delivery (CD) ecosystem.

CDEvents are an enabler for interoperability between different CI/CD tools as they provide a tool agnostic communication protocol.

## Events Format

CDEvents are [CloudEvents](https://cloudevents.io/) with specific extensions and payload structure, which is based on the [CDEvents vocabulary](introduction.md).

Events producer may use the payload to provide extra context to the events consumers; the payload however is not meant to
transport large amounts of data. Data such as logs or software artifacts should be linked to from the event and not
embedded into the events. We follow the [CloudEvents recommendation](https://github.com/cloudevents/spec/blob/v1.0.1/spec.md#size-limits)
on events size and size limits.

### Specification

The aim of the Events SIG is to produce a version specification for CDEvents. The specification was initially defined in the [draft CDE vocabulary](introduction.md).

## Tools

The draft specification of the CDEvents is combined with an [SDK and CLI](https://github.com/cdfoundation/sig-events/blob/main/cde/sdk/go/README.md).

## Proof of Concept

There is a [Proof of Concept](../poc/README.md) demonstrating how to use draft CDEvents specification and its SDK.
