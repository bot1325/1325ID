package commands

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"

	"github.com/authelia/authelia/v4/internal/configuration/schema"
)

func newStorageCmd() (cmd *cobra.Command) {
	cmd = &cobra.Command{
		Use:               "storage",
		Short:             cmdAutheliaStorageShort,
		Long:              cmdAutheliaStorageLong,
		Example:           cmdAutheliaStorageExample,
		Args:              cobra.NoArgs,
		PersistentPreRunE: storagePersistentPreRunE,

		DisableAutoGenTag: true,
	}

	cmdWithConfigFlags(cmd, true, []string{"configuration.yml"})

	cmd.PersistentFlags().String(cmdFlagNameEncryptionKey, "", "the storage encryption key to use")

	cmd.PersistentFlags().String(cmdFlagNameSQLite3Path, "", "the SQLite database path")

	cmd.PersistentFlags().String(cmdFlagNameMySQLHost, "", "the MySQL hostname")
	cmd.PersistentFlags().Int(cmdFlagNameMySQLPort, 3306, "the MySQL port")
	cmd.PersistentFlags().String(cmdFlagNameMySQLDatabase, "authelia", "the MySQL database name")
	cmd.PersistentFlags().String(cmdFlagNameMySQLUsername, "authelia", "the MySQL username")
	cmd.PersistentFlags().String(cmdFlagNameMySQLPassword, "", "the MySQL password")

	cmd.PersistentFlags().String(cmdFlagNamePostgreSQLHost, "", "the PostgreSQL hostname")
	cmd.PersistentFlags().Int(cmdFlagNamePostgreSQLPort, 5432, "the PostgreSQL port")
	cmd.PersistentFlags().String(cmdFlagNamePostgreSQLDatabase, "authelia", "the PostgreSQL database name")
	cmd.PersistentFlags().String(cmdFlagNamePostgreSQLSchema, "public", "the PostgreSQL schema name")
	cmd.PersistentFlags().String(cmdFlagNamePostgreSQLUsername, "authelia", "the PostgreSQL username")
	cmd.PersistentFlags().String(cmdFlagNamePostgreSQLPassword, "", "the PostgreSQL password")
	cmd.PersistentFlags().String("postgres.ssl.mode", "disable", "the PostgreSQL ssl mode")
	cmd.PersistentFlags().String("postgres.ssl.root_certificate", "", "the PostgreSQL ssl root certificate file location")
	cmd.PersistentFlags().String("postgres.ssl.certificate", "", "the PostgreSQL ssl certificate file location")
	cmd.PersistentFlags().String("postgres.ssl.key", "", "the PostgreSQL ssl key file location")

	cmd.AddCommand(
		newStorageMigrateCmd(),
		newStorageSchemaInfoCmd(),
		newStorageEncryptionCmd(),
		newStorageUserCmd(),
	)

	return cmd
}

func newStorageEncryptionCmd() (cmd *cobra.Command) {
	cmd = &cobra.Command{
		Use:     "encryption",
		Short:   cmdAutheliaStorageEncryptionShort,
		Long:    cmdAutheliaStorageEncryptionLong,
		Example: cmdAutheliaStorageEncryptionExample,

		DisableAutoGenTag: true,
	}

	cmd.AddCommand(
		newStorageEncryptionChangeKeyCmd(),
		newStorageEncryptionCheckCmd(),
	)

	return cmd
}

func newStorageEncryptionCheckCmd() (cmd *cobra.Command) {
	cmd = &cobra.Command{
		Use:     "check",
		Short:   cmdAutheliaStorageEncryptionCheckShort,
		Long:    cmdAutheliaStorageEncryptionCheckLong,
		Example: cmdAutheliaStorageEncryptionCheckExample,
		RunE:    storageSchemaEncryptionCheckRunE,

		DisableAutoGenTag: true,
	}

	cmd.Flags().Bool(cmdFlagNameVerbose, false, "enables verbose checking of every row of encrypted data")

	return cmd
}

func newStorageEncryptionChangeKeyCmd() (cmd *cobra.Command) {
	cmd = &cobra.Command{
		Use:     "change-key",
		Short:   cmdAutheliaStorageEncryptionChangeKeyShort,
		Long:    cmdAutheliaStorageEncryptionChangeKeyLong,
		Example: cmdAutheliaStorageEncryptionChangeKeyExample,
		RunE:    storageSchemaEncryptionChangeKeyRunE,

		DisableAutoGenTag: true,
	}

	cmd.Flags().String(cmdFlagNameNewEncryptionKey, "", "the new key to encrypt the data with")

	return cmd
}

func newStorageUserCmd() (cmd *cobra.Command) {
	cmd = &cobra.Command{
		Use:     "user",
		Short:   cmdAutheliaStorageUserShort,
		Long:    cmdAutheliaStorageUserLong,
		Example: cmdAutheliaStorageUserExample,

		DisableAutoGenTag: true,
	}

	cmd.AddCommand(
		newStorageUserIdentifiersCmd(),
		newStorageUserTOTPCmd(),
		newStorageUserWebAuthnCmd(),
	)

	return cmd
}

func newStorageUserIdentifiersCmd() (cmd *cobra.Command) {
	cmd = &cobra.Command{
		Use:     "identifiers",
		Short:   cmdAutheliaStorageUserIdentifiersShort,
		Long:    cmdAutheliaStorageUserIdentifiersLong,
		Example: cmdAutheliaStorageUserIdentifiersExample,

		DisableAutoGenTag: true,
	}

	cmd.AddCommand(
		newStorageUserIdentifiersExportCmd(),
		newStorageUserIdentifiersImportCmd(),
		newStorageUserIdentifiersGenerateCmd(),
		newStorageUserIdentifiersAddCmd(),
	)

	return cmd
}

func newStorageUserIdentifiersExportCmd() (cmd *cobra.Command) {
	cmd = &cobra.Command{
		Use:     "export",
		Short:   cmdAutheliaStorageUserIdentifiersExportShort,
		Long:    cmdAutheliaStorageUserIdentifiersExportLong,
		Example: cmdAutheliaStorageUserIdentifiersExportExample,
		RunE:    storageUserIdentifiersExport,

		DisableAutoGenTag: true,
	}

	cmd.Flags().StringP(cmdFlagNameFile, "f", "user-opaque-identifiers.yml", "The file name for the YAML export")

	return cmd
}

func newStorageUserIdentifiersImportCmd() (cmd *cobra.Command) {
	cmd = &cobra.Command{
		Use:     "import",
		Short:   cmdAutheliaStorageUserIdentifiersImportShort,
		Long:    cmdAutheliaStorageUserIdentifiersImportLong,
		Example: cmdAutheliaStorageUserIdentifiersImportExample,
		RunE:    storageUserIdentifiersImport,

		DisableAutoGenTag: true,
	}

	cmd.Flags().StringP(cmdFlagNameFile, "f", "user-opaque-identifiers.yml", "The file name for the YAML import")

	return cmd
}

func newStorageUserIdentifiersGenerateCmd() (cmd *cobra.Command) {
	cmd = &cobra.Command{
		Use:     "generate",
		Short:   cmdAutheliaStorageUserIdentifiersGenerateShort,
		Long:    cmdAutheliaStorageUserIdentifiersGenerateLong,
		Example: cmdAutheliaStorageUserIdentifiersGenerateExample,
		RunE:    storageUserIdentifiersGenerate,

		DisableAutoGenTag: true,
	}

	cmd.Flags().StringSlice(cmdFlagNameUsers, nil, "The list of users to generate the opaque identifiers for")
	cmd.Flags().StringSlice(cmdFlagNameServices, []string{identifierServiceOpenIDConnect}, fmt.Sprintf("The list of services to generate the opaque identifiers for, valid values are: %s", strings.Join(validIdentifierServices, ", ")))
	cmd.Flags().StringSlice(cmdFlagNameSectors, []string{""}, "The list of sectors to generate identifiers for")

	return cmd
}

func newStorageUserIdentifiersAddCmd() (cmd *cobra.Command) {
	cmd = &cobra.Command{
		Use:     "add <username>",
		Short:   cmdAutheliaStorageUserIdentifiersAddShort,
		Long:    cmdAutheliaStorageUserIdentifiersAddLong,
		Example: cmdAutheliaStorageUserIdentifiersAddExample,
		Args:    cobra.ExactArgs(1),
		RunE:    storageUserIdentifiersAdd,

		DisableAutoGenTag: true,
	}

	cmd.Flags().String(cmdFlagNameIdentifier, "", "The optional version 4 UUID to use, if not set a random one will be used")
	cmd.Flags().String(cmdFlagNameService, identifierServiceOpenIDConnect, fmt.Sprintf("The service to add the identifier for, valid values are: %s", strings.Join(validIdentifierServices, ", ")))
	cmd.Flags().String(cmdFlagNameSector, "", "The sector identifier to use (should usually be blank)")

	return cmd
}

func newStorageUserWebAuthnCmd() (cmd *cobra.Command) {
	cmd = &cobra.Command{
		Use:     "webauthn",
		Short:   cmdAutheliaStorageUserWebAuthnShort,
		Long:    cmdAutheliaStorageUserWebAuthnLong,
		Example: cmdAutheliaStorageUserWebAuthnExample,

		DisableAutoGenTag: true,
	}

	cmd.AddCommand(
		newStorageUserWebAuthnListCmd(),
		newStorageUserWebAuthnDeleteCmd(),
	)

	return cmd
}

func newStorageUserWebAuthnListCmd() (cmd *cobra.Command) {
	cmd = &cobra.Command{
		Use:     "list [username]",
		Short:   cmdAutheliaStorageUserWebAuthnListShort,
		Long:    cmdAutheliaStorageUserWebAuthnListLong,
		Example: cmdAutheliaStorageUserWebAuthnListExample,
		RunE:    storageWebAuthnListRunE,
		Args:    cobra.MaximumNArgs(1),

		DisableAutoGenTag: true,
	}

	return cmd
}

func newStorageUserWebAuthnDeleteCmd() (cmd *cobra.Command) {
	cmd = &cobra.Command{
		Use:     "delete [username]",
		Short:   cmdAutheliaStorageUserWebAuthnDeleteShort,
		Long:    cmdAutheliaStorageUserWebAuthnDeleteLong,
		Example: cmdAutheliaStorageUserWebAuthnDeleteExample,
		RunE:    storageWebAuthnDeleteRunE,
		Args:    cobra.MaximumNArgs(1),

		DisableAutoGenTag: true,
	}

	cmd.Flags().Bool(cmdFlagNameAll, false, "delete all of the users webauthn devices")
	cmd.Flags().String(cmdFlagNameDescription, "", "delete a users webauthn device by description")
	cmd.Flags().String(cmdFlagNameKeyID, "", "delete a users webauthn device by key id")

	return cmd
}

func newStorageUserTOTPCmd() (cmd *cobra.Command) {
	cmd = &cobra.Command{
		Use:     "totp",
		Short:   cmdAutheliaStorageUserTOTPShort,
		Long:    cmdAutheliaStorageUserTOTPLong,
		Example: cmdAutheliaStorageUserTOTPExample,

		DisableAutoGenTag: true,
	}

	cmd.AddCommand(
		newStorageUserTOTPGenerateCmd(),
		newStorageUserTOTPDeleteCmd(),
		newStorageUserTOTPExportCmd(),
	)

	return cmd
}

func newStorageUserTOTPGenerateCmd() (cmd *cobra.Command) {
	cmd = &cobra.Command{
		Use:     "generate <username>",
		Short:   cmdAutheliaStorageUserTOTPGenerateShort,
		Long:    cmdAutheliaStorageUserTOTPGenerateLong,
		Example: cmdAutheliaStorageUserTOTPGenerateExample,
		RunE:    storageTOTPGenerateRunE,
		Args:    cobra.ExactArgs(1),

		DisableAutoGenTag: true,
	}

	cmd.Flags().String(cmdFlagNameSecret, "", "set the shared secret as base32 encoded bytes (no padding), it's recommended that you do not use this option unless you're restoring a configuration")
	cmd.Flags().Uint(cmdFlagNameSecretSize, schema.TOTPSecretSizeDefault, "set the secret size")
	cmd.Flags().Uint(cmdFlagNamePeriod, 30, "set the period between rotations")
	cmd.Flags().Uint(cmdFlagNameDigits, 6, "set the number of digits")
	cmd.Flags().String(cmdFlagNameAlgorithm, "SHA1", "set the algorithm to either SHA1 (supported by most applications), SHA256, or SHA512")
	cmd.Flags().String(cmdFlagNameIssuer, "Authelia", "set the issuer description")
	cmd.Flags().BoolP(cmdFlagNameForce, "f", false, "forces the configuration to be generated regardless if it exists or not")
	cmd.Flags().StringP(cmdFlagNamePath, "p", "", "path to a file to create a PNG file with the QR code (optional)")

	return cmd
}

func newStorageUserTOTPDeleteCmd() (cmd *cobra.Command) {
	cmd = &cobra.Command{
		Use:     "delete <username>",
		Short:   cmdAutheliaStorageUserTOTPDeleteShort,
		Long:    cmdAutheliaStorageUserTOTPDeleteLong,
		Example: cmdAutheliaStorageUserTOTPDeleteExample,
		RunE:    storageTOTPDeleteRunE,
		Args:    cobra.ExactArgs(1),

		DisableAutoGenTag: true,
	}

	return cmd
}

func newStorageUserTOTPExportCmd() (cmd *cobra.Command) {
	cmd = &cobra.Command{
		Use:     "export",
		Short:   cmdAutheliaStorageUserTOTPExportShort,
		Long:    cmdAutheliaStorageUserTOTPExportLong,
		Example: cmdAutheliaStorageUserTOTPExportExample,
		RunE:    storageTOTPExportRunE,

		DisableAutoGenTag: true,
	}

	cmd.Flags().String(cmdFlagNameFormat, storageTOTPExportFormatURI, fmt.Sprintf("sets the output format, valid values are: %s", strings.Join(validStorageTOTPExportFormats, ", ")))
	cmd.Flags().String("dir", "", "used with the png output format to specify which new directory to save the files in")

	return cmd
}

func newStorageSchemaInfoCmd() (cmd *cobra.Command) {
	cmd = &cobra.Command{
		Use:     "schema-info",
		Short:   cmdAutheliaStorageSchemaInfoShort,
		Long:    cmdAutheliaStorageSchemaInfoLong,
		Example: cmdAutheliaStorageSchemaInfoExample,
		RunE:    storageSchemaInfoRunE,

		DisableAutoGenTag: true,
	}

	return cmd
}

// NewMigrationCmd returns a new Migration Cmd.
func newStorageMigrateCmd() (cmd *cobra.Command) {
	cmd = &cobra.Command{
		Use:     "migrate",
		Short:   cmdAutheliaStorageMigrateShort,
		Long:    cmdAutheliaStorageMigrateLong,
		Example: cmdAutheliaStorageMigrateExample,
		Args:    cobra.NoArgs,

		DisableAutoGenTag: true,
	}

	cmd.AddCommand(
		newStorageMigrateUpCmd(), newStorageMigrateDownCmd(),
		newStorageMigrateListUpCmd(), newStorageMigrateListDownCmd(),
		newStorageMigrateHistoryCmd(),
	)

	return cmd
}

func newStorageMigrateHistoryCmd() (cmd *cobra.Command) {
	cmd = &cobra.Command{
		Use:     "history",
		Short:   cmdAutheliaStorageMigrateHistoryShort,
		Long:    cmdAutheliaStorageMigrateHistoryLong,
		Example: cmdAutheliaStorageMigrateHistoryExample,
		Args:    cobra.NoArgs,
		RunE:    storageMigrateHistoryRunE,

		DisableAutoGenTag: true,
	}

	return cmd
}

func newStorageMigrateListUpCmd() (cmd *cobra.Command) {
	cmd = &cobra.Command{
		Use:     "list-up",
		Short:   cmdAutheliaStorageMigrateListUpShort,
		Long:    cmdAutheliaStorageMigrateListUpLong,
		Example: cmdAutheliaStorageMigrateListUpExample,
		Args:    cobra.NoArgs,
		RunE:    newStorageMigrateListRunE(true),

		DisableAutoGenTag: true,
	}

	return cmd
}

func newStorageMigrateListDownCmd() (cmd *cobra.Command) {
	cmd = &cobra.Command{
		Use:     "list-down",
		Short:   cmdAutheliaStorageMigrateListDownShort,
		Long:    cmdAutheliaStorageMigrateListDownLong,
		Example: cmdAutheliaStorageMigrateListDownExample,
		Args:    cobra.NoArgs,
		RunE:    newStorageMigrateListRunE(false),

		DisableAutoGenTag: true,
	}

	return cmd
}

func newStorageMigrateUpCmd() (cmd *cobra.Command) {
	cmd = &cobra.Command{
		Use:     storageMigrateDirectionUp,
		Short:   cmdAutheliaStorageMigrateUpShort,
		Long:    cmdAutheliaStorageMigrateUpLong,
		Example: cmdAutheliaStorageMigrateUpExample,
		Args:    cobra.NoArgs,
		RunE:    newStorageMigrationRunE(true),

		DisableAutoGenTag: true,
	}

	cmd.Flags().IntP(cmdFlagNameTarget, "t", 0, "sets the version to migrate to, by default this is the latest version")

	return cmd
}

func newStorageMigrateDownCmd() (cmd *cobra.Command) {
	cmd = &cobra.Command{
		Use:     storageMigrateDirectionDown,
		Short:   cmdAutheliaStorageMigrateDownShort,
		Long:    cmdAutheliaStorageMigrateDownLong,
		Example: cmdAutheliaStorageMigrateDownExample,
		Args:    cobra.NoArgs,
		RunE:    newStorageMigrationRunE(false),

		DisableAutoGenTag: true,
	}

	cmd.Flags().IntP(cmdFlagNameTarget, "t", 0, "sets the version to migrate to")
	cmd.Flags().Bool(cmdFlagNameDestroyData, false, "confirms you want to destroy data with this migration")

	return cmd
}
