package commands

import (
	"bytes"
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"io"
	"io/ioutil"
	"os"

	"github.com/pkg/errors"
	"github.com/spf13/cobra"

	"github.com/chef/automate/lib/io/fileutils"
)

func init() {
	RootCmd.AddCommand(initCmd)
}

const (
	// Location of secrets_key generated by compliance service in
	// previous A2 versions. This key is migrated to
	// secrets-service if it exists.
	complianceKeyPath = "/hab/svc/compliance-service/data/secrets_key"
	// Number of bytes of random data to use to generate the key
	defaultKeyLength = 16
	// Permissions on key
	defaultKeyMode = 0600
)

var initCmd = &cobra.Command{
	Use:   "init KEYPATH",
	Short: "Initializes the shared secret used for symmetric encryption. Migrated from compliance if needed",
	Args:  cobra.ExactArgs(1),
	Run: func(_ *cobra.Command, args []string) {
		keyPath := args[0]
		keyContent, err := ioutil.ReadFile(keyPath)
		if err != nil && !os.IsNotExist(err) {
			fatalf("failed to check existing key: %s", err.Error())
		}

		if os.IsNotExist(err) || len(keyContent) < 2 {
			err = generateOrMigrateKey(keyPath)
			if err != nil {
				fatalf("%s", err.Error())
			}
		}

		err = os.Chmod(keyPath, defaultKeyMode)
		if err != nil {
			fatalf("failed to change key permissions: %s", err.Error())
		}
	},
}

func generateOrMigrateKey(newKeyPath string) error {
	complianceKeyExists, err := fileutils.PathExists(complianceKeyPath)
	if err != nil {
		return errors.Wrap(err, "failed to check for compliance key")
	}

	var keyContentReader io.Reader
	if complianceKeyExists {
		fmt.Println("Migrating secret key from compliance service")
		complianceKeyFile, err := os.Open(complianceKeyPath)
		if err != nil {
			return errors.Wrap(err, "failed to open compliance key")
		}
		keyContentReader = complianceKeyFile
	} else {
		fmt.Println("Generating new secret key")
		key, err := generateKey()
		if err != nil {
			return errors.Wrap(err, "could not generate key")
		}
		keyContentReader = bytes.NewBuffer(key)
	}

	err = fileutils.AtomicWrite(newKeyPath, keyContentReader, fileutils.WithAtomicWriteFileMode(defaultKeyMode))
	if err != nil {
		return errors.Wrap(err, "could not write key file")
	}

	return nil
}

func generateKey() ([]byte, error) {
	randomBytes := make([]byte, defaultKeyLength)
	encodedBytes := make([]byte, defaultKeyLength*2)
	if _, err := rand.Read(randomBytes); err != nil {
		return nil, errors.Wrap(err, "failed to read random data while generating random bytes")
	}

	hex.Encode(encodedBytes, randomBytes)
	return encodedBytes, nil
}

func fatalf(msg string, args ...interface{}) {
	fmt.Fprintf(os.Stderr, msg, args...)
	os.Exit(1)
}
