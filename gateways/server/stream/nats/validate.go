/*
Copyright 2018 BlackRock, Inc.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

	http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package nats

import (
	"context"

	"github.com/argoproj/argo-events/gateways"
	"github.com/argoproj/argo-events/pkg/apis/eventsources/v1alpha1"
	"github.com/ghodss/yaml"
	"github.com/pkg/errors"
)

// ValidateEventSource validates nats event source
func (listener *EventListener) ValidateEventSource(ctx context.Context, eventSource *gateways.EventSource) (*gateways.ValidEventSource, error) {
	var natsGridEventSource *v1alpha1.NATSEventsSource
	if err := yaml.Unmarshal(eventSource.Value, &natsGridEventSource); err != nil {
		listener.Logger.WithError(err).Error("failed to parse the event source")
		return &gateways.ValidEventSource{
			IsValid: false,
			Reason:  err.Error(),
		}, nil
	}

	if err := validate(natsGridEventSource); err != nil {
		listener.Logger.WithError(err).Error("failed to validate nats event source")
		return &gateways.ValidEventSource{
			IsValid: false,
			Reason:  err.Error(),
		}, nil
	}

	return &gateways.ValidEventSource{
		IsValid: true,
	}, nil
}

func validate(eventSource *v1alpha1.NATSEventsSource) error {
	if eventSource == nil {
		return errors.New("configuration must be non empty")
	}
	if eventSource.URL == "" {
		return errors.New("url must be specified")
	}
	if eventSource.Subject == "" {
		return errors.New("subject must be specified")
	}
	return nil
}