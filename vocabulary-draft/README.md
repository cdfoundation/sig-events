# Continuous Delivery Events

Continuous Delivery Events (CD Events or CDE) is a common event format for the Continuous Delivery (CD) ecosystem.

CD Events are an enabler for interoperability between different CI/CD tools as they provide a tool agnostic communication protocol.

## Name

Continuous Delivery Events is a temporary name that we adopted until we have a final name.
If you have an idea for a name, feel free to share it via [this form](https://docs.google.com/forms/d/1CVPooDG16B6JaBqQysH7V6UehYcZcSw_PM1HrtxwNns/edit).

## Events Format

CD Events are [CloudEvents](https://cloudevents.io/) with CDE specific extensions and payload structure, which is based
on the [CDE vocabulary](introduction.md).

Events producer may use the payload to provide extra context to the events consumers; the payload however is not meant to
transport large amounts of data. Data such as logs or software artifacts should be linked to from the event and not
embedded into the events. We follow the [CloudEvents recommendation](https://github.com/cloudevents/spec/blob/v1.0.1/spec.md#size-limits)
on events size and size limits.

### Specification

The aim of the Events SIG is to produce a version specification for CDE. The draft specification is defined in the
form a this [draft CDE vocabulary](introduction.md).

## Tools

Developers may create CD Events via the [CDE SDK and CLI](../cde/sdk/go/README.md).

## Proof of Concept

The Events SIG is working on a [Proof of Concept](../poc/README.md) to demonstrate how to use CD Events and the SDK.
