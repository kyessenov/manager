// Copyright 2016 Google Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package model

import "github.com/golang/protobuf/proto"

// ConfigKey is the identity of the configuration object
type ConfigKey struct {
	// Config object kind, e.g.  "MyKind"
	Kind string
	// Config object name, e.g. "my-name"
	Name string
	// Config object namespace, e.g. "default"
	Namespace string
}

// Config object holds the normalized config objects defined by Kind schema
type Config struct {
	ConfigKey
	// Spec holds the configuration struct
	Spec interface{}
	// Status holds the status information (may be null)
	Status interface{}
}

// Registry of the configuration objects
type Registry interface {
	Get(key ConfigKey) (*Config, bool)
	Delete(key ConfigKey)
	Put(obj Config) error
	List(kind string, namespace string) []*Config
}

// KindMap defines bijection between Kind name and proto message name
type KindMap map[string]ProtoSchema

// ProtoSchema provides custom validation checks
type ProtoSchema struct {
	// MessageName refers to the protobuf message type name
	MessageName string
	// StatusMessageName refers to the protubuf message type name for the StatusMessageName
	StatusMessageName string
	// Description of the configuration type
	Description string
	// Validate configuration as a protobuf message
	Validate func(o proto.Message) error
}
