// Copyright 2021 ZUP IT SERVICOS EM TECNOLOGIA E INOVACAO SA
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

package enums

const (
	EnvAnalyticDatabaseURI          = "HORUSEC_DATABASE_SQL_URI"
	DefaultAnalyticDatabaseURIValue = "postgresql://root:root@142.93.254.133:5432/horusec_analytic_db?sslmode=disable"
	EnvHorusecDatabaseURI           = "HORUSEC_DATABASE_HORUSEC_SQL_URI"
	DefaultHorusecDatabaseURIValue  = "postgresql://root:root@142.93.254.133:5432/horusec_db?sslmode=disable"
	SummarySuccess                  = "success"
	SummaryFailed                   = "failed"
	TotalOfTables                   = 4
	MigrationV1ToV2Name             = "20210609_v1_to_v2"
	MigrationTable                  = "horusec_migrations"
)
