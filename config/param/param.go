// Copyright 2016 Netflix, Inc.
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

package param

// properties
const (
	Enabled          = "chaosbum.enabled"
	Leashed          = "chaosbum.leashed"
	ScheduleEnabled  = "chaosbum.schedule_enabled"
	Accounts         = "chaosbum.accounts"
	StartHour        = "chaosbum.start_hour"
	EndHour          = "chaosbum.end_hour"
	TimeZone         = "chaosbum.time_zone"
	CronPath         = "chaosbum.cron_path"
	TermPath         = "chaosbum.term_path"
	TermAccount      = "chaosbum.term_account"
	MaxApps          = "chaosbum.max_apps"
	Trackers         = "chaosbum.trackers"
	ErrorCounter     = "chaosbum.error_counter"
	Decryptor        = "chaosbum.decryptor"
	OutageChecker    = "chaosbum.outage_checker"
	CronExpression   = "chaosbum.cron_expression"
	ScheduleCronPath = "chaosbum.schedule_cron_path"
	SchedulePath     = "chaosbum.schedule_path"
	LogPath          = "chaosbum.log_path"

	// spinnaker
	SpinnakerEndpoint          = "spinnaker.endpoint"
	SpinnakerCertificate       = "spinnaker.certificate"
	SpinnakerEncryptedPassword = "spinnaker.encrypted_password"
	SpinnakerUser              = "spinnaker.user"
	SpinnakerX509Cert          = "spinnaker.x509_cert"
	SpinnakerX509Key           = "spinnaker.x509_key"
	// database
	DatabaseHost              = "database.host"
	DatabasePort              = "database.port"
	DatabaseUser              = "database.user"
	DatabaseEncryptedPassword = "database.encrypted_password"
	DatabaseName              = "database.name"

	// dynamic property provider
	DynamicProvider = "dynamic.provider"
	DynamicEndpoint = "dynamic.endpoint"
	DynamicPath     = "dynamic.path"
)
