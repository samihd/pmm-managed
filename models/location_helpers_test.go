// pmm-managed
// Copyright (C) 2017 Percona LLC
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU Affero General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Affero General Public License for more details.
//
// You should have received a copy of the GNU Affero General Public License
// along with this program. If not, see <https://www.gnu.org/licenses/>.

package models_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"gopkg.in/reform.v1"
	"gopkg.in/reform.v1/dialects/postgresql"

	"github.com/percona/pmm-managed/models"
	"github.com/percona/pmm-managed/utils/testdb"
	"github.com/percona/pmm-managed/utils/tests"
)

func TestBackupLocations(t *testing.T) {
	sqlDB := testdb.Open(t, models.SkipFixtures, nil)
	defer func() {
		require.NoError(t, sqlDB.Close())
	}()
	db := reform.NewDB(sqlDB, postgresql.Dialect, reform.NewPrintfLogger(t.Logf))

	t.Run("create - pmm client", func(t *testing.T) {
		tx, err := db.Begin()
		require.NoError(t, err)
		defer func() {
			require.NoError(t, tx.Rollback())
		}()

		q := tx.Querier

		params := models.CreateBackupLocationParams{
			Name:        "some name",
			Description: "some desc",
			BackupLocationConfig: models.BackupLocationConfig{
				PMMClientConfig: &models.PMMClientLocationConfig{
					Path: "/tmp",
				},
			},
		}

		location, err := models.CreateBackupLocation(q, params)
		require.NoError(t, err)
		assert.Equal(t, models.PMMClientBackupLocationType, location.Type)
		assert.Equal(t, params.Name, location.Name)
		assert.Equal(t, params.Description, location.Description)
		assert.Equal(t, params.PMMClientConfig.Path, location.PMMClientConfig.Path)
		assert.NotEmpty(t, location.ID)
	})

	t.Run("create - s3", func(t *testing.T) {
		tx, err := db.Begin()
		require.NoError(t, err)
		defer func() {
			require.NoError(t, tx.Rollback())
		}()

		q := tx.Querier

		params := models.CreateBackupLocationParams{
			Name:        "some name",
			Description: "some desc",
			BackupLocationConfig: models.BackupLocationConfig{
				S3Config: &models.S3LocationConfig{
					Endpoint:   "https://example.com/",
					AccessKey:  "access_key",
					SecretKey:  "secret_key",
					BucketName: "example_bucket",
				},
			},
		}

		location, err := models.CreateBackupLocation(q, params)
		require.NoError(t, err)
		assert.Equal(t, models.S3BackupLocationType, location.Type)
		assert.Equal(t, params.Name, location.Name)
		assert.Equal(t, params.Description, location.Description)
		assert.Equal(t, params.S3Config.Endpoint, location.S3Config.Endpoint)
		assert.Equal(t, params.S3Config.AccessKey, location.S3Config.AccessKey)
		assert.Equal(t, params.S3Config.SecretKey, location.S3Config.SecretKey)
		assert.Equal(t, params.S3Config.BucketName, location.S3Config.BucketName)

		assert.NotEmpty(t, location.ID)
	})

	t.Run("create - two configs", func(t *testing.T) {
		tx, err := db.Begin()
		require.NoError(t, err)
		defer func() {
			require.NoError(t, tx.Rollback())
		}()

		q := tx.Querier

		params := models.CreateBackupLocationParams{
			Name:        "some name",
			Description: "some desc",
			BackupLocationConfig: models.BackupLocationConfig{
				PMMClientConfig: &models.PMMClientLocationConfig{
					Path: "/tmp",
				},
				S3Config: &models.S3LocationConfig{
					Endpoint:   "https://example.com/",
					AccessKey:  "access_key",
					SecretKey:  "secret_key",
					BucketName: "example_bucket",
				},
			},
		}

		_, err = models.CreateBackupLocation(q, params)
		tests.AssertGRPCError(t, status.New(codes.InvalidArgument, "Only one config is allowed."), err)
	})

	t.Run("list", func(t *testing.T) {
		tx, err := db.Begin()
		require.NoError(t, err)
		defer func() {
			require.NoError(t, tx.Rollback())
		}()

		q := tx.Querier

		params1 := models.CreateBackupLocationParams{
			Name:        "some name",
			Description: "some desc",
			BackupLocationConfig: models.BackupLocationConfig{
				PMMClientConfig: &models.PMMClientLocationConfig{
					Path: "/tmp",
				},
			},
		}
		params2 := models.CreateBackupLocationParams{
			Name:        "some name2",
			Description: "some desc2",
			BackupLocationConfig: models.BackupLocationConfig{
				S3Config: &models.S3LocationConfig{
					Endpoint:   "https://example.com/",
					AccessKey:  "access_key",
					SecretKey:  "secret_key",
					BucketName: "example_bucket",
				},
			},
		}

		loc1, err := models.CreateBackupLocation(q, params1)
		require.NoError(t, err)
		loc2, err := models.CreateBackupLocation(q, params2)
		require.NoError(t, err)

		actual, err := models.FindBackupLocations(q)
		require.NoError(t, err)

		findLocID := func(id string) func() bool {
			return func() bool {
				for _, location := range actual {
					if location.ID == id {
						return true
					}
				}
				return false
			}
		}

		assert.Condition(t, findLocID(loc1.ID), "First location not found")
		assert.Condition(t, findLocID(loc2.ID), "Second location not found")
	})

	t.Run("update", func(t *testing.T) {
		tx, err := db.Begin()
		require.NoError(t, err)
		defer func() {
			require.NoError(t, tx.Rollback())
		}()

		q := tx.Querier

		createParams := models.CreateBackupLocationParams{
			Name:        "some name",
			Description: "some desc",
			BackupLocationConfig: models.BackupLocationConfig{
				PMMClientConfig: &models.PMMClientLocationConfig{
					Path: "/tmp",
				},
			},
		}

		location, err := models.CreateBackupLocation(q, createParams)
		require.NoError(t, err)

		changeParams := models.ChangeBackupLocationParams{
			Name:        "some name modified",
			Description: "",
			BackupLocationConfig: models.BackupLocationConfig{
				PMMServerConfig: &models.PMMServerLocationConfig{
					Path: "/tmp/nested",
				},
			},
		}

		updatedLoc, err := models.ChangeBackupLocation(q, location.ID, changeParams)
		require.NoError(t, err)
		assert.Equal(t, changeParams.Name, updatedLoc.Name)
		// empty description in request, we expect no change
		assert.Equal(t, createParams.Description, updatedLoc.Description)
		assert.Nil(t, updatedLoc.PMMClientConfig)
		assert.Equal(t, changeParams.PMMServerConfig.Path, updatedLoc.PMMServerConfig.Path)
		assert.Equal(t, updatedLoc.Type, models.PMMServerBackupLocationType)

		findLoc, err := models.FindBackupLocationByID(q, location.ID)
		require.NoError(t, err)

		assert.Equal(t, updatedLoc, findLoc)
	})

	t.Run("remove", func(t *testing.T) {
		tx, err := db.Begin()
		require.NoError(t, err)
		defer func() {
			require.NoError(t, tx.Rollback())
		}()

		q := tx.Querier

		params := models.CreateBackupLocationParams{
			Name:        "some name",
			Description: "some desc",
			BackupLocationConfig: models.BackupLocationConfig{
				PMMClientConfig: &models.PMMClientLocationConfig{
					Path: "/tmp",
				},
			},
		}

		loc, err := models.CreateBackupLocation(q, params)
		require.NoError(t, err)

		err = models.RemoveBackupLocation(q, loc.ID, models.RemoveRestrict)
		require.NoError(t, err)

		locations, err := models.FindBackupLocations(q)
		require.NoError(t, err)
		assert.Empty(t, locations)
	})
}

func TestCreateBackupLocationValidation(t *testing.T) {
	sqlDB := testdb.Open(t, models.SkipFixtures, nil)
	defer func() {
		require.NoError(t, sqlDB.Close())
	}()
	db := reform.NewDB(sqlDB, postgresql.Dialect, reform.NewPrintfLogger(t.Logf))

	tableTests := []struct {
		name     string
		params   models.CreateBackupLocationParams
		errorMsg string
	}{
		{
			name: "normal client config",
			params: models.CreateBackupLocationParams{
				Name: "client-1",
				BackupLocationConfig: models.BackupLocationConfig{
					PMMClientConfig: &models.PMMClientLocationConfig{
						Path: "/tmp",
					},
				},
			},
			errorMsg: "",
		},
		{
			name: "client config - missing path",
			params: models.CreateBackupLocationParams{
				Name: "client-2",
				BackupLocationConfig: models.BackupLocationConfig{
					PMMClientConfig: &models.PMMClientLocationConfig{
						Path: "",
					},
				},
			},
			errorMsg: "rpc error: code = InvalidArgument desc = PMM client config path field is empty.",
		},
		{
			name: "normal s3 config",
			params: models.CreateBackupLocationParams{
				Name: "s3-1",
				BackupLocationConfig: models.BackupLocationConfig{
					S3Config: &models.S3LocationConfig{
						Endpoint:   "https://s3.us-west-2.amazonaws.com/",
						AccessKey:  "access_key",
						SecretKey:  "secret_key",
						BucketName: "example_bucket",
					},
				},
			},
			errorMsg: "",
		},
		{
			name: "s3 config - missing endpoint",
			params: models.CreateBackupLocationParams{
				Name: "s3-2",
				BackupLocationConfig: models.BackupLocationConfig{
					S3Config: &models.S3LocationConfig{
						Endpoint:   "",
						AccessKey:  "access_key",
						SecretKey:  "secret_key",
						BucketName: "example_bucket",
					},
				},
			},
			errorMsg: "rpc error: code = InvalidArgument desc = S3 endpoint field is empty.",
		},
		{
			name: "s3 config - missing access key",
			params: models.CreateBackupLocationParams{
				Name: "s3-3",
				BackupLocationConfig: models.BackupLocationConfig{
					S3Config: &models.S3LocationConfig{
						Endpoint:   "https://s3.us-west-2.amazonaws.com/",
						AccessKey:  "",
						SecretKey:  "secret_key",
						BucketName: "example_bucket",
					},
				},
			},
			errorMsg: "rpc error: code = InvalidArgument desc = S3 accessKey field is empty.",
		},
		{
			name: "s3 config - missing secret key",
			params: models.CreateBackupLocationParams{
				Name: "s3-4",
				BackupLocationConfig: models.BackupLocationConfig{
					S3Config: &models.S3LocationConfig{
						Endpoint:   "https://s3.us-west-2.amazonaws.com/",
						AccessKey:  "secret_key",
						SecretKey:  "",
						BucketName: "example_bucket",
					},
				},
			},
			errorMsg: "rpc error: code = InvalidArgument desc = S3 secretKey field is empty.",
		},
		{
			name: "s3 config - missing bucket name",
			params: models.CreateBackupLocationParams{
				Name: "s3-5",
				BackupLocationConfig: models.BackupLocationConfig{
					S3Config: &models.S3LocationConfig{
						Endpoint:   "https://s3.us-west-2.amazonaws.com/",
						AccessKey:  "secret_key",
						SecretKey:  "example_key",
						BucketName: "",
					},
				},
			},
			errorMsg: "rpc error: code = InvalidArgument desc = S3 bucketName field is empty.",
		},
		{
			name: "s3 config - invalid endpoint",
			params: models.CreateBackupLocationParams{
				Name: "s3-6",
				BackupLocationConfig: models.BackupLocationConfig{
					S3Config: &models.S3LocationConfig{
						Endpoint:   "#invalidendpoint",
						AccessKey:  "secret_key",
						SecretKey:  "example_key",
						BucketName: "example_bucket",
					},
				},
			},
			errorMsg: "rpc error: code = InvalidArgument desc = No host found in the Endpoint.",
		},
		{
			name: "s3 config - invalid endpoint, path is not allowed",
			params: models.CreateBackupLocationParams{
				Name: "s3-7",
				BackupLocationConfig: models.BackupLocationConfig{
					S3Config: &models.S3LocationConfig{
						Endpoint:   "https://s3.us-west-2.amazonaws.com/path",
						AccessKey:  "secret_key",
						SecretKey:  "example_key",
						BucketName: "example_bucket",
					},
				},
			},
			errorMsg: "rpc error: code = InvalidArgument desc = Path is not allowed for Endpoint.",
		},
		{
			name: "s3 config - invalid scheme",
			params: models.CreateBackupLocationParams{
				Name: "s3-8",
				BackupLocationConfig: models.BackupLocationConfig{
					S3Config: &models.S3LocationConfig{
						Endpoint:   "tcp://s3.us-west-2.amazonaws.com",
						AccessKey:  "secret_key",
						SecretKey:  "example_key",
						BucketName: "example_bucket",
					},
				},
			},
			errorMsg: "rpc error: code = InvalidArgument desc = Invalid scheme 'tcp'",
		},
	}

	for _, test := range tableTests {
		test := test

		t.Run(test.name, func(t *testing.T) {
			tx, err := db.Begin()
			require.NoError(t, err)
			defer func() {
				require.NoError(t, tx.Rollback())
			}()

			q := tx.Querier

			c, err := models.CreateBackupLocation(q, test.params)
			if test.errorMsg != "" {
				assert.EqualError(t, err, test.errorMsg)
				return
			}
			assert.NoError(t, err)
			assert.NotNil(t, c)
		})
	}
}

func TestVerifyBackupLocationValidation(t *testing.T) {
	tableTests := []struct {
		name     string
		params   *models.VerifyBackupLocationParams
		errorMsg string
	}{
		{
			name: "client config - missing path",
			params: &models.VerifyBackupLocationParams{
				BackupLocationConfig: models.BackupLocationConfig{
					PMMClientConfig: &models.PMMClientLocationConfig{
						Path: "",
					},
				},
			},
			errorMsg: "rpc error: code = InvalidArgument desc = PMM client config path field is empty.",
		},
		{
			name:     "s3 config - missing config",
			params:   &models.VerifyBackupLocationParams{},
			errorMsg: "rpc error: code = InvalidArgument desc = Missing location config.",
		},
		{
			name: "s3 config - missing endpoint",
			params: &models.VerifyBackupLocationParams{
				BackupLocationConfig: models.BackupLocationConfig{
					S3Config: &models.S3LocationConfig{
						Endpoint:   "",
						AccessKey:  "access_key",
						SecretKey:  "secret_key",
						BucketName: "example_bucket",
					},
				},
			},
			errorMsg: "rpc error: code = InvalidArgument desc = S3 endpoint field is empty.",
		},
		{
			name: "s3 config - missing access key",
			params: &models.VerifyBackupLocationParams{
				BackupLocationConfig: models.BackupLocationConfig{
					S3Config: &models.S3LocationConfig{
						Endpoint:   "https://s3.us-west-2.amazonaws.com/",
						AccessKey:  "",
						SecretKey:  "secret_key",
						BucketName: "example_bucket",
					},
				},
			},
			errorMsg: "rpc error: code = InvalidArgument desc = S3 accessKey field is empty.",
		},
		{
			name: "s3 config - missing secret key",
			params: &models.VerifyBackupLocationParams{
				BackupLocationConfig: models.BackupLocationConfig{
					S3Config: &models.S3LocationConfig{
						Endpoint:   "https://s3.us-west-2.amazonaws.com/",
						AccessKey:  "secret_key",
						SecretKey:  "",
						BucketName: "example_bucket",
					},
				},
			},
			errorMsg: "rpc error: code = InvalidArgument desc = S3 secretKey field is empty.",
		},
		{
			name: "s3 config - missing bucket name",
			params: &models.VerifyBackupLocationParams{
				BackupLocationConfig: models.BackupLocationConfig{
					S3Config: &models.S3LocationConfig{
						Endpoint:   "https://s3.us-west-2.amazonaws.com/",
						AccessKey:  "secret_key",
						SecretKey:  "example_key",
						BucketName: "",
					},
				},
			},
			errorMsg: "rpc error: code = InvalidArgument desc = S3 bucketName field is empty.",
		},
		{
			name: "s3 config - invalid endpoint",
			params: &models.VerifyBackupLocationParams{
				BackupLocationConfig: models.BackupLocationConfig{
					S3Config: &models.S3LocationConfig{
						Endpoint:   "#invalidendpoint",
						AccessKey:  "secret_key",
						SecretKey:  "example_key",
						BucketName: "example_bucket",
					},
				},
			},
			errorMsg: "rpc error: code = InvalidArgument desc = No host found in the Endpoint.",
		},
		{
			name: "s3 config - invalid endpoint, path is not allowed",
			params: &models.VerifyBackupLocationParams{
				BackupLocationConfig: models.BackupLocationConfig{
					S3Config: &models.S3LocationConfig{
						Endpoint:   "https://s3.us-west-2.amazonaws.com/path",
						AccessKey:  "secret_key",
						SecretKey:  "example_key",
						BucketName: "example_bucket",
					},
				},
			},
			errorMsg: "rpc error: code = InvalidArgument desc = Path is not allowed for Endpoint.",
		},
		{
			name: "s3 config - invalid scheme",
			params: &models.VerifyBackupLocationParams{
				BackupLocationConfig: models.BackupLocationConfig{
					S3Config: &models.S3LocationConfig{
						Endpoint:   "tcp://s3.us-west-2.amazonaws.com",
						AccessKey:  "secret_key",
						SecretKey:  "example_key",
						BucketName: "example_bucket",
					},
				},
			},
			errorMsg: "rpc error: code = InvalidArgument desc = Invalid scheme 'tcp'",
		},
		{
			name: "s3 config - invalid bucket name",
			params: &models.VerifyBackupLocationParams{
				BackupLocationConfig: models.BackupLocationConfig{
					S3Config: &models.S3LocationConfig{
						Endpoint:   "s3.us-west-2.amazonaws.com",
						AccessKey:  "secret_key",
						SecretKey:  "example_key",
						BucketName: "invalid@bucket",
					},
				},
			},
			errorMsg: "rpc error: code = Internal desc = Bucket name contains invalid characters",
		},
	}

	for _, test := range tableTests {
		test := test

		t.Run(test.name, func(t *testing.T) {
			err := models.VerifyBackupLocationConfig(test.params)
			if test.errorMsg != "" {
				assert.EqualError(t, err, test.errorMsg)
				return
			}
			assert.NoError(t, err)
		})
	}
}
