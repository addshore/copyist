// Copyright 2020 The Cockroach Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or
// implied. See the License for the specific language governing
// permissions and limitations under the License.

// pqtestold tests an ancient version of PQ, before support was added for
// QueryContext, ExecContext, and BeginTx, in order to test that copyist works
// even when drivers do not support those functions.

package pqtestold

import (
	"testing"

	"github.com/cockroachdb/copyist/drivertest/commontest"

	_ "github.com/lib/pq/old"
)

const (
	driverName     = "postgres"
	dataSourceName = "postgresql://root@localhost:26888?sslmode=disable"

	// Don't use default CRDB port in case another instance is already running.
	dockerArgs = "-p 26888:26257 cockroachdb/cockroach:v20.2.4 start-single-node --insecure"
)

// TestMain runs all PQ driver-specific tests. To use:
//
//   1. Run the tests with the "-record" command-line flag. This will run the
//      tests against the real PQ driver and create recording files in the
//      testdata directory. This tests generation of recordings.
//   2. Run the test without the "-record" flag. This will run the tests against
//      the copyist driver that plays back the recordings created by step #1.
//      This tests playback of recording.
//
func TestMain(m *testing.M) {
	commontest.RunAllTests(m, driverName, dataSourceName, dockerArgs)
}

// TestQuery fetches a single customer.
func TestQuery(t *testing.T) {
	commontest.RunTestQuery(t, driverName, dataSourceName)
}

// TestInsert inserts a row and ensures that it's been committed.
func TestInsert(t *testing.T) {
	commontest.RunTestInsert(t, driverName, dataSourceName)
}

// TestDataTypes queries data types that are interesting for the SQL driver.
func TestDataTypes(t *testing.T) {
	commontest.RunTestDataTypes(t, driverName, dataSourceName)
}

// TestFloatLiterals tests the generation of float literal values, with and
// without fractions and exponents.
func TestFloatLiterals(t *testing.T) {
	// Run twice in order to regress problem with float round-tripping.
	t.Run("run 1", func(t *testing.T) {
		commontest.RunTestFloatLiterals(t, driverName, dataSourceName)
	})
	t.Run("run 2", func(t *testing.T) {
		commontest.RunTestFloatLiterals(t, driverName, dataSourceName)
	})
}

// TestTxns commits and aborts transactions.
func TestTxns(t *testing.T) {
	commontest.RunTestTxns(t, driverName, dataSourceName)
}

// TestSqlx tests usage of the `sqlx` package with copyist.
func TestSqlx(t *testing.T) {
	commontest.RunTestSqlx(t, driverName, dataSourceName)
}
