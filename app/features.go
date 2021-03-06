package app

//go:generate mockgen -destination=../mocks/app/mock_features.go -package=mock_app github.com/rudderlabs/rudder-server/app MigratorFeature,SuppressUserFeature

import (
	backendconfig "github.com/rudderlabs/rudder-server/config/backend-config"
	"github.com/rudderlabs/rudder-server/jobsdb"
	"github.com/rudderlabs/rudder-server/utils/types"
)

// MigratorFeature handles migration of nodes during cluster's scale up/down.
type MigratorFeature interface {
	Setup(*jobsdb.HandleT, *jobsdb.HandleT, *jobsdb.HandleT, func(), func())
}

// MigratorFeatureSetup is a function that initializes a Migrator feature, based on application instance
type MigratorFeatureSetup func(Interface) MigratorFeature

var migratorFeatureSetup MigratorFeatureSetup

// RegisterMigratorFeature registers a Migration implementation
func RegisterMigratorFeature(f MigratorFeatureSetup) {
	migratorFeatureSetup = f
}

// SuppressUserFeature handles webhook event requests
type SuppressUserFeature interface {
	Setup(backendConfig backendconfig.BackendConfig) types.SuppressUserI
}

// SuppressUserFeatureSetup is a function that initializes a Webhook feature, based on application instance
type SuppressUserFeatureSetup func(Interface) SuppressUserFeature

var suppressUserFeatureSetup SuppressUserFeatureSetup

// RegisterSuppressUserFeature registers a suppress user feature implementation
func RegisterSuppressUserFeature(f SuppressUserFeatureSetup) {
	suppressUserFeatureSetup = f
}

// ProtocolsFeature handles event schemas
type ProtocolsFeature interface {
	Setup() types.ProtocolsI
}

// ProtocolsFeatureSetup is a function that initializes a Protocols feature, based on application instance
type ProtocolsFeatureSetup func(Interface) ProtocolsFeature

var protocolsFeatureSetup ProtocolsFeatureSetup

// RegisterProtocolsFeature registers a protocols feature implementation
func RegisterProtocolsFeature(f ProtocolsFeatureSetup) {
	protocolsFeatureSetup = f
}

/*********************************
DestinationConfig Env Support
*********************************/

// ProtocolsFeature handles event schemas
type ConfigEnvFeature interface {
	Setup() types.ConfigEnvI
}

// ProtocolsFeatureSetup is a function that initializes a Protocols feature, based on application instance
type ConfigEnvFeatureSetup func(Interface) ConfigEnvFeature

var configEnvFeatureSetup ConfigEnvFeatureSetup

// RegisterProtocolsFeature registers a protocols feature implementation
func RegisterConfigEnvFeature(f ConfigEnvFeatureSetup) {
	configEnvFeatureSetup = f
}

// Features contains optional implementations of Enterprise only features.
type Features struct {
	Migrator     MigratorFeature
	SuppressUser SuppressUserFeature
	Protocols    ProtocolsFeature
	ConfigEnv    ConfigEnvFeature
}
