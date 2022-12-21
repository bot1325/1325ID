package commands

import (
	"errors"
)

const (
	fmtCmdAutheliaShort = "authelia %s"

	fmtCmdAutheliaLong = `authelia %s

An open-source authentication and authorization server providing
two-factor authentication and single sign-on (SSO) for your
applications via a web portal.

Documentation is available at: https://www.authelia.com/`

	cmdAutheliaExample = `authelia --config /etc/authelia/config.yml --config /etc/authelia/access-control.yml
authelia --config /etc/authelia/config.yml,/etc/authelia/access-control.yml
authelia --config /etc/authelia/config/`

	fmtAutheliaBuild = `Last Tag: %s
State: %s
Branch: %s
Commit: %s
Build Number: %s
Build OS: %s
Build Arch: %s
Build Date: %s
Extra: %s
`

	cmdAutheliaBuildInfoShort = "Show the build information of Authelia"

	cmdAutheliaBuildInfoLong = `Show the build information of Authelia.

This outputs detailed version information about the specific version
of the Authelia binary. This information is embedded into Authelia
by the continuous integration.

This could be vital in debugging if you're not using a particular
tagged build of Authelia. It's suggested to provide it along with
your issue.
`
	cmdAutheliaBuildInfoExample = `authelia build-info`

	cmdAutheliaAccessControlShort = "Helpers for the access control system"

	cmdAutheliaAccessControlLong = `Helpers for the access control system.`

	cmdAutheliaAccessControlExample = `authelia access-control --help`

	cmdAutheliaAccessControlCheckPolicyShort = "Checks a request against the access control rules to determine what policy would be applied"

	cmdAutheliaAccessControlCheckPolicyLong = `
Checks a request against the access control rules to determine what policy would be applied.

Legend:

	#		The rule position in the configuration.
	*		The first fully matched rule.
	~		Potential match i.e. if the user was authenticated they may match this rule.
	hit     The criteria in this column is a match to the request.
	miss    The criteria in this column is not match to the request.
	may     The criteria in this column is potentially a match to the request.

Notes:

	A rule that potentially matches a request will cause a redirection to occur in order to perform one-factor
	authentication. This is so Authelia can adequately determine if the rule actually matches.
`
	cmdAutheliaAccessControlCheckPolicyExample = `authelia access-control check-policy --config config.yml --url https://example.com
authelia access-control check-policy --config config.yml --url https://example.com --username john
authelia access-control check-policy --config config.yml --url https://example.com --groups admin,public
authelia access-control check-policy --config config.yml --url https://example.com --username john --method GET
authelia access-control check-policy --config config.yml --url https://example.com --username john --method GET --verbose`

	cmdAutheliaStorageShort = "Manage the Authelia storage"

	cmdAutheliaStorageLong = `Manage the Authelia storage.

This subcommand has several methods to interact with the Authelia SQL Database. This allows doing several advanced
operations which would be much harder to do manually.
`

	cmdAutheliaStorageExample = `authelia storage --help`

	cmdAutheliaStorageEncryptionShort = "Manage storage encryption"

	cmdAutheliaStorageEncryptionLong = `Manage storage encryption.

This subcommand allows management of the storage encryption.`

	cmdAutheliaStorageEncryptionExample = `authelia storage encryption --help`

	cmdAutheliaStorageEncryptionCheckShort = "Checks the encryption key against the database data"

	cmdAutheliaStorageEncryptionCheckLong = `Checks the encryption key against the database data.

This is useful for validating all data that can be encrypted is intact.`

	cmdAutheliaStorageEncryptionCheckExample = `authelia storage encryption check
authelia storage encryption check --verbose
authelia storage encryption check --verbose --config config.yml
authelia storage encryption check --verbose --encryption-key b3453fde-ecc2-4a1f-9422-2707ddbed495 --postgres.host postgres --postgres.password autheliapw`

	cmdAutheliaStorageEncryptionChangeKeyShort = "Changes the encryption key"

	cmdAutheliaStorageEncryptionChangeKeyLong = `Changes the encryption key.

This subcommand allows you to change the encryption key of an Authelia SQL database.`

	cmdAutheliaStorageEncryptionChangeKeyExample = `authelia storage encryption change-key --config config.yml --new-encryption-key 0e95cb49-5804-4ad9-be82-bb04a9ddecd8
authelia storage encryption change-key --encryption-key b3453fde-ecc2-4a1f-9422-2707ddbed495 --new-encryption-key 0e95cb49-5804-4ad9-be82-bb04a9ddecd8 --postgres.host postgres --postgres.password autheliapw`

	cmdAutheliaStorageUserShort = "Manages user settings"

	cmdAutheliaStorageUserLong = `Manages user settings.

This subcommand allows modifying and exporting user settings.`

	cmdAutheliaStorageUserExample = `authelia storage user --help`

	cmdAutheliaStorageUserIdentifiersShort = "Manage user opaque identifiers"

	cmdAutheliaStorageUserIdentifiersLong = `Manage user opaque identifiers.

This subcommand allows performing various tasks related to the opaque identifiers for users.`

	cmdAutheliaStorageUserIdentifiersExample = `authelia storage user identifiers --help`

	cmdAutheliaStorageUserIdentifiersExportShort = "Export the identifiers to a YAML file"

	cmdAutheliaStorageUserIdentifiersExportLong = `Export the identifiers to a YAML file.

This subcommand allows exporting the opaque identifiers for users in order to back them up.`

	cmdAutheliaStorageUserIdentifiersExportExample = `authelia storage user identifiers export
authelia storage user identifiers export --file export.yaml
authelia storage user identifiers export --file export.yaml --config config.yml
authelia storage user identifiers export --file export.yaml --encryption-key b3453fde-ecc2-4a1f-9422-2707ddbed495 --postgres.host postgres --postgres.password autheliapw`

	cmdAutheliaStorageUserIdentifiersImportShort = "Import the identifiers from a YAML file"

	cmdAutheliaStorageUserIdentifiersImportLong = `Import the identifiers from a YAML file.

This subcommand allows you to import the opaque identifiers for users from a YAML file.

The YAML file can either be automatically generated using the authelia storage user identifiers export command, or
manually provided the file is in the same format.`

	cmdAutheliaStorageUserIdentifiersImportExample = `authelia storage user identifiers import
authelia storage user identifiers import --file export.yaml
authelia storage user identifiers import --file export.yaml --config config.yml
authelia storage user identifiers import --file export.yaml --encryption-key b3453fde-ecc2-4a1f-9422-2707ddbed495 --postgres.host postgres --postgres.password autheliapw`

	cmdAutheliaStorageUserIdentifiersGenerateShort = "Generate opaque identifiers in bulk"

	cmdAutheliaStorageUserIdentifiersGenerateLong = `Generate opaque identifiers in bulk.

This subcommand allows various options for generating the opaque identifies for users in bulk.`

	cmdAutheliaStorageUserIdentifiersGenerateExample = `authelia storage user identifiers generate --users john,mary
authelia storage user identifiers generate --users john,mary --services openid
authelia storage user identifiers generate --users john,mary --services openid --sectors=",example.com,test.com"
authelia storage user identifiers generate --users john,mary --services openid --sectors=",example.com,test.com" --config config.yml
authelia storage user identifiers generate --users john,mary --services openid --sectors=",example.com,test.com" --encryption-key b3453fde-ecc2-4a1f-9422-2707ddbed495 --postgres.host postgres --postgres.password autheliapw`

	cmdAutheliaStorageUserIdentifiersAddShort = "Add an opaque identifier for a user to the database"

	cmdAutheliaStorageUserIdentifiersAddLong = `Add an opaque identifier for a user to the database.

This subcommand allows manually adding an opaque identifier for a user to the database provided it's in the correct format.`

	cmdAutheliaStorageUserIdentifiersAddExample = `authelia storage user identifiers add john --identifier f0919359-9d15-4e15-bcba-83b41620a073
authelia storage user identifiers add john --identifier f0919359-9d15-4e15-bcba-83b41620a073 --config config.yml
authelia storage user identifiers add john --identifier f0919359-9d15-4e15-bcba-83b41620a073 --encryption-key b3453fde-ecc2-4a1f-9422-2707ddbed495 --postgres.host postgres --postgres.password autheliapw`

	cmdAutheliaStorageUserWebAuthnShort = "Manage Webauthn devices"

	cmdAutheliaStorageUserWebAuthnLong = `Manage Webauthn devices.

This subcommand allows interacting with Webauthn devices.`

	cmdAutheliaStorageUserWebAuthnExample = `authelia storage user webauthn --help`

	cmdAutheliaStorageUserWebAuthnListShort = "List WebAuthn devices"

	cmdAutheliaStorageUserWebAuthnListLong = `List WebAuthn devices.

This subcommand allows listing WebAuthn devices.`

	cmdAutheliaStorageUserWebAuthnListExample = `authelia storage user webauthn list
authelia storage user webauthn list john
authelia storage user webauthn list --config config.yml
authelia storage user webauthn list john --config config.yml
authelia storage user webauthn list --encryption-key b3453fde-ecc2-4a1f-9422-2707ddbed495 --postgres.host postgres --postgres.password autheliapw
authelia storage user webauthn list john --encryption-key b3453fde-ecc2-4a1f-9422-2707ddbed495 --postgres.host postgres --postgres.password autheliapw`

	cmdAutheliaStorageUserWebAuthnDeleteShort = "Delete a WebAuthn device"

	cmdAutheliaStorageUserWebAuthnDeleteLong = `Delete a WebAuthn device.

This subcommand allows deleting a WebAuthn device directly from the database.`

	cmdAutheliaStorageUserWebAuthnDeleteExample = `authelia storage user webauthn delete john --all
authelia storage user webauthn delete john --all --config config.yml
authelia storage user webauthn delete john --all --encryption-key b3453fde-ecc2-4a1f-9422-2707ddbed495 --postgres.host postgres --postgres.password autheliapw
authelia storage user webauthn delete john --description Primary
authelia storage user webauthn delete john --description Primary --config config.yml
authelia storage user webauthn delete john --description Primary --encryption-key b3453fde-ecc2-4a1f-9422-2707ddbed495 --postgres.host postgres --postgres.password autheliapw
authelia storage user webauthn delete --kid abc123
authelia storage user webauthn delete --kid abc123 --config config.yml
authelia storage user webauthn delete --kid abc123 --encryption-key b3453fde-ecc2-4a1f-9422-2707ddbed495 --postgres.host postgres --postgres.password autheliapw`

	cmdAutheliaStorageUserTOTPShort = "Manage TOTP configurations"

	cmdAutheliaStorageUserTOTPLong = `Manage TOTP configurations.

This subcommand allows deleting, exporting, and creating user TOTP configurations.`

	cmdAutheliaStorageUserTOTPExample = `authelia storage user totp --help`

	cmdAutheliaStorageUserTOTPGenerateShort = "Generate a TOTP configuration for a user"

	cmdAutheliaStorageUserTOTPGenerateLong = `Generate a TOTP configuration for a user.

This subcommand allows generating a new TOTP configuration for a user,
and overwriting the existing configuration if applicable.`

	cmdAutheliaStorageUserTOTPGenerateExample = `authelia storage user totp generate john
authelia storage user totp generate john --period 90
authelia storage user totp generate john --digits 8
authelia storage user totp generate john --algorithm SHA512
authelia storage user totp generate john --algorithm SHA512 --config config.yml
authelia storage user totp generate john --algorithm SHA512 --config config.yml --path john.png`

	cmdAutheliaStorageUserTOTPDeleteShort = "Delete a TOTP configuration for a user"

	cmdAutheliaStorageUserTOTPDeleteLong = `Delete a TOTP configuration for a user.

This subcommand allows deleting a TOTP configuration directly from the database for a given user.`

	cmdAutheliaStorageUserTOTPDeleteExample = `authelia storage user totp delete john
authelia storage user totp delete john --config config.yml
authelia storage user totp delete john --encryption-key b3453fde-ecc2-4a1f-9422-2707ddbed495 --postgres.host postgres --postgres.password autheliapw`

	cmdAutheliaStorageUserTOTPExportShort = "Perform exports of the TOTP configurations"

	cmdAutheliaStorageUserTOTPExportLong = `Perform exports of the TOTP configurations.

This subcommand allows exporting TOTP configurations to various formats.`

	cmdAutheliaStorageUserTOTPExportExample = `authelia storage user totp export --format csv
authelia storage user totp export --format png --dir ./totp-qr
authelia storage user totp export --format png --dir ./totp-qr --config config.yml
authelia storage user totp export --format png --dir ./totp-qr --encryption-key b3453fde-ecc2-4a1f-9422-2707ddbed495 --postgres.host postgres --postgres.password autheliapw`

	cmdAutheliaStorageSchemaInfoShort = "Show the storage information"

	cmdAutheliaStorageSchemaInfoLong = `Show the storage information.

This subcommand shows advanced information about the storage schema useful in some diagnostic tasks.`

	cmdAutheliaStorageSchemaInfoExample = `authelia storage schema-info
authelia storage schema-info --config config.yml
authelia storage schema-info --encryption-key b3453fde-ecc2-4a1f-9422-2707ddbed495 --postgres.host postgres --postgres.password autheliapw`

	cmdAutheliaStorageMigrateShort = "Perform or list migrations"

	cmdAutheliaStorageMigrateLong = `Perform or list migrations.

This subcommand handles schema migration tasks.`

	cmdAutheliaStorageMigrateExample = `authelia storage migrate --help`

	cmdAutheliaStorageMigrateHistoryShort = "Show migration history"

	cmdAutheliaStorageMigrateHistoryLong = `Show migration history.

This subcommand allows users to list previous migrations.`

	cmdAutheliaStorageMigrateHistoryExample = `authelia storage migrate history
authelia storage migrate history --config config.yml
authelia storage migrate history --encryption-key b3453fde-ecc2-4a1f-9422-2707ddbed495 --postgres.host postgres --postgres.password autheliapw`

	cmdAutheliaStorageMigrateListUpShort = "List the up migrations available"

	cmdAutheliaStorageMigrateListUpLong = `List the up migrations available.

This subcommand lists the schema migrations available in this version of Authelia which are greater than the current
schema version of the database.`

	cmdAutheliaStorageMigrateListUpExample = `authelia storage migrate list-up
authelia storage migrate list-up --config config.yml
authelia storage migrate list-up --encryption-key b3453fde-ecc2-4a1f-9422-2707ddbed495 --postgres.host postgres --postgres.password autheliapw`

	cmdAutheliaStorageMigrateListDownShort = "List the down migrations available"

	cmdAutheliaStorageMigrateListDownLong = `List the down migrations available.

This subcommand lists the schema migrations available in this version of Authelia which are less than the current
schema version of the database.`

	cmdAutheliaStorageMigrateListDownExample = `authelia storage migrate list-down
authelia storage migrate list-down --config config.yml
authelia storage migrate list-down --encryption-key b3453fde-ecc2-4a1f-9422-2707ddbed495 --postgres.host postgres --postgres.password autheliapw`

	cmdAutheliaStorageMigrateUpShort = "Perform a migration up"

	cmdAutheliaStorageMigrateUpLong = `Perform a migration up.

This subcommand performs the schema migrations available in this version of Authelia which are greater than the current
schema version of the database. By default this will migrate up to the latest available, but you can customize this.`

	cmdAutheliaStorageMigrateUpExample = `authelia storage migrate up
authelia storage migrate up --config config.yml
authelia storage migrate up --target 20 --config config.yml
authelia storage migrate up --encryption-key b3453fde-ecc2-4a1f-9422-2707ddbed495 --postgres.host postgres --postgres.password autheliapw`

	cmdAutheliaStorageMigrateDownShort = "Perform a migration down"

	cmdAutheliaStorageMigrateDownLong = `Perform a migration down.

This subcommand performs the schema migrations available in this version of Authelia which are less than the current
schema version of the database.`

	cmdAutheliaStorageMigrateDownExample = `authelia storage migrate down --target 20
authelia storage migrate down --target 20 --config config.yml
authelia storage migrate down --target 20 --encryption-key b3453fde-ecc2-4a1f-9422-2707ddbed495 --postgres.host postgres --postgres.password autheliapw`

	cmdAutheliaValidateConfigShort = "Check a configuration against the internal configuration validation mechanisms"

	cmdAutheliaValidateConfigLong = `Check a configuration against the internal configuration validation mechanisms.

This subcommand allows validation of the YAML and Environment configurations so that a configuration can be checked
prior to deploying it.`

	cmdAutheliaValidateConfigExample = `authelia validate-config
authelia validate-config --config config.yml`

	cmdAutheliaCryptoShort = "Perform cryptographic operations"

	cmdAutheliaCryptoLong = `Perform cryptographic operations.

This subcommand allows preforming cryptographic certificate, key pair, etc tasks.`

	cmdAutheliaCryptoExample = `authelia crypto --help`

	cmdAutheliaCryptoRandShort = "Generate a cryptographically secure random string"

	cmdAutheliaCryptoRandLong = `Generate a cryptographically secure random string.

This subcommand allows generating cryptographically secure random strings for use for encryption keys, HMAC keys, etc.`

	cmdAutheliaCryptoRandExample = `authelia crypto rand --help
authelia crypto rand --length 80
authelia crypto rand -n 80
authelia crypto rand --charset alphanumeric
authelia crypto rand --charset alphabetic
authelia crypto rand --charset ascii
authelia crypto rand --charset numeric
authelia crypto rand --charset numeric-hex
authelia crypto rand --characters 0123456789ABCDEF`

	cmdAutheliaCryptoHashShort = "Perform cryptographic hash operations"

	cmdAutheliaCryptoHashLong = `Perform cryptographic hash operations.

This subcommand allows preforming hashing cryptographic tasks.`

	cmdAutheliaCryptoHashExample = `authelia crypto hash --help`

	cmdAutheliaCryptoHashValidateShort = "Perform cryptographic hash validations"

	cmdAutheliaCryptoHashValidateLong = `Perform cryptographic hash validations.

This subcommand allows preforming cryptographic hash validations. i.e. checking hash digests against a password.`

	cmdAutheliaCryptoHashValidateExample = `authelia crypto hash validate --help
authelia crypto hash validate '$5$rounds=500000$WFjMpdCQxIkbNl0k$M0qZaZoK8Gwdh8Cw5diHgGfe5pE0iJvxcVG3.CVnQe.' -- 'p@ssw0rd'`

	cmdAutheliaCryptoHashGenerateShort = "Generate cryptographic hash digests"

	cmdAutheliaCryptoHashGenerateLong = `Generate cryptographic hash digests.

This subcommand allows generating cryptographic hash digests.

See the help for the subcommands if you want to override the configuration or defaults.`

	cmdAutheliaCryptoHashGenerateExample = `authelia crypto hash generate --help`

	fmtCmdAutheliaCryptoHashGenerateSubShort = "Generate cryptographic %s hash digests"

	fmtCmdAutheliaCryptoHashGenerateSubLong = `Generate cryptographic %s hash digests.

This subcommand allows generating cryptographic %s hash digests.`

	fmtCmdAutheliaCryptoHashGenerateSubExample = `authelia crypto hash generate %s --help`

	cmdAutheliaCryptoCertificateShort = "Perform certificate cryptographic operations"

	cmdAutheliaCryptoCertificateLong = `Perform certificate cryptographic operations.

This subcommand allows preforming certificate cryptographic tasks.`

	cmdAutheliaCryptoCertificateExample = `authelia crypto certificate --help`

	fmtCmdAutheliaCryptoCertificateSubShort = "Perform %s certificate cryptographic operations"

	fmtCmdAutheliaCryptoCertificateSubLong = `Perform %s certificate cryptographic operations.

This subcommand allows preforming %s certificate cryptographic tasks.`

	fmtCmdAutheliaCryptoCertificateSubExample = `authelia crypto certificate %s --help`

	fmtCmdAutheliaCryptoCertificateGenerateRequestShort = "Generate an %s private key and %s"

	fmtCmdAutheliaCryptoCertificateGenerateRequestLong = `Generate an %s private key and %s.

This subcommand allows generating an %s private key and %s.`

	cmdAutheliaCryptoCertificateRSAGenerateExample = `authelia crypto certificate rsa generate --help`

	cmdAutheliaCryptoCertificateECDSAGenerateExample = `authelia crypto certificate ecdsa generate --help`

	cmdAutheliaCryptoCertificateEd25519GenerateExample = `authelia crypto certificate ed25519 request --help`

	cmdAutheliaCryptoCertificateRSARequestExample = `authelia crypto certificate rsa request --help`

	cmdAutheliaCryptoCertificateECDSARequestExample = `authelia crypto certificate ecdsa request --help`

	cmdAutheliaCryptoCertificateEd25519RequestExample = `authelia crypto certificate ed25519 request --help`

	cmdAutheliaCryptoPairShort = "Perform key pair cryptographic operations"

	cmdAutheliaCryptoPairLong = `Perform key pair cryptographic operations.

This subcommand allows preforming key pair cryptographic tasks.`

	cmdAutheliaCryptoPairExample = `authelia crypto pair --help`

	cmdAutheliaCryptoPairSubShort = "Perform %s key pair cryptographic operations"

	cmdAutheliaCryptoPairSubLong = `Perform %s key pair cryptographic operations.

This subcommand allows preforming %s key pair cryptographic tasks.`

	cmdAutheliaCryptoPairRSAExample = `authelia crypto pair rsa --help`

	cmdAutheliaCryptoPairECDSAExample = `authelia crypto pair ecdsa --help`

	cmdAutheliaCryptoPairEd25519Example = `authelia crypto pair ed25519 --help`

	fmtCmdAutheliaCryptoPairGenerateShort = "Generate a cryptographic %s key pair"

	fmtCmdAutheliaCryptoPairGenerateLong = `Generate a cryptographic %s key pair.

This subcommand allows generating an %s key pair.`

	cmdAutheliaCryptoPairRSAGenerateExample = `authelia crypto pair rsa generate --help`

	cmdAutheliaCryptoPairECDSAGenerateExample = `authelia crypto pair ecdsa generate --help`

	cmdAutheliaCryptoPairEd25519GenerateExample = `authelia crypto pair ed25519 generate --help`

	cmdAutheliaHashPasswordShort = "Hash a password to be used in file-based users database"

	cmdAutheliaHashPasswordLong = `Hash a password to be used in file-based users database.`

	//nolint:gosec // This is an example.
	cmdAutheliaHashPasswordExample = `authelia hash-password -- 'mypass'
authelia hash-password --sha512 -- 'mypass'
authelia hash-password --iterations=4 -- 'mypass'
authelia hash-password --memory=128 -- 'mypass'
authelia hash-password --parallelism=1 -- 'mypass'
authelia hash-password --key-length=64 -- 'mypass'`
)

const (
	storageMigrateDirectionUp   = "up"
	storageMigrateDirectionDown = "down"
)

const (
	storageTOTPExportFormatCSV = "csv"
	storageTOTPExportFormatURI = "uri"
	storageTOTPExportFormatPNG = "png"
)

var (
	validStorageTOTPExportFormats = []string{storageTOTPExportFormatCSV, storageTOTPExportFormatURI, storageTOTPExportFormatPNG}
)

const (
	timeLayoutCertificateNotBefore = "Jan 2 15:04:05 2006"
)

const (
	cmdFlagNameDirectory = "directory"

	cmdFlagNamePathCA = "path.ca"

	cmdFlagNameFilePrivateKey    = "file.private-key"
	cmdFlagNameFilePublicKey     = "file.public-key"
	cmdFlagNameFileCertificate   = "file.certificate"
	cmdFlagNameFileCAPrivateKey  = "file.ca-private-key"
	cmdFlagNameFileCACertificate = "file.ca-certificate"
	cmdFlagNameFileCSR           = "file.csr"

	cmdFlagNameExtendedUsage = "extended-usage"
	cmdFlagNameSignature     = "signature"
	cmdFlagNameCA            = "ca"
	cmdFlagNameSANs          = "sans"

	cmdFlagNameCommonName         = "common-name"
	cmdFlagNameOrganization       = "organization"
	cmdFlagNameOrganizationalUnit = "organizational-unit"
	cmdFlagNameCountry            = "country"
	cmdFlagNameProvince           = "province"
	cmdFlagNameLocality           = "locality"
	cmdFlagNameStreetAddress      = "street-address"
	cmdFlagNamePostcode           = "postcode"

	cmdFlagNameNotBefore = "not-before"
	cmdFlagNameDuration  = "duration"

	cmdFlagNamePKCS8 = "pkcs8"
	cmdFlagNameBits  = "bits"
	cmdFlagNameCurve = "curve"

	cmdFlagNamePassword         = "password"
	cmdFlagNameRandom           = "random"
	cmdFlagNameRandomLength     = "random.length"
	cmdFlagNameRandomCharSet    = "random.charset"
	cmdFlagNameRandomCharacters = "random.characters"
	cmdFlagNameNoConfirm        = "no-confirm"
	cmdFlagNameVariant          = "variant"
	cmdFlagNameCost             = "cost"
	cmdFlagNameIterations       = "iterations"
	cmdFlagNameParallelism      = "parallelism"
	cmdFlagNameBlockSize        = "block-size"
	cmdFlagNameMemory           = "memory"
	cmdFlagNameKeySize          = "key-size"
	cmdFlagNameSaltSize         = "salt-size"
	cmdFlagNameProfile          = "profile"
	cmdFlagNameSHA512           = "sha512"
	cmdFlagNameConfig           = "config"

	cmdFlagNameConfigExpFilters = "config.experimental.filters"

	cmdFlagNameCharSet     = "charset"
	cmdFlagValueCharSet    = "alphanumeric"
	cmdFlagUsageCharset    = "sets the charset for the random password, options are 'ascii', 'alphanumeric', 'alphabetic', 'numeric', 'numeric-hex', and 'rfc3986'"
	cmdFlagNameCharacters  = "characters"
	cmdFlagUsageCharacters = "sets the explicit characters for the random string"
	cmdFlagNameLength      = "length"
	cmdFlagUsageLength     = "sets the character length for the random string"

	cmdFlagNameNewEncryptionKey = "new-encryption-key"

	cmdFlagNameFile        = "file"
	cmdFlagNameUsers       = "users"
	cmdFlagNameServices    = "services"
	cmdFlagNameSectors     = "sectors"
	cmdFlagNameIdentifier  = "identifier"
	cmdFlagNameService     = "service"
	cmdFlagNameSector      = "sector"
	cmdFlagNameDescription = "description"
	cmdFlagNameAll         = "all"
	cmdFlagNameKeyID       = "kid"
	cmdFlagNameVerbose     = "verbose"
	cmdFlagNameSecret      = "secret"
	cmdFlagNameSecretSize  = "secret-size"
	cmdFlagNamePeriod      = "period"
	cmdFlagNameDigits      = "digits"
	cmdFlagNameAlgorithm   = "algorithm"
	cmdFlagNameIssuer      = "issuer"
	cmdFlagNameForce       = "force"
	cmdFlagNameFormat      = "format"
	cmdFlagNamePath        = "path"
	cmdFlagNameTarget      = "target"
	cmdFlagNameDestroyData = "destroy-data"

	cmdFlagNameEncryptionKey      = "encryption-key"
	cmdFlagNameSQLite3Path        = "sqlite.path"
	cmdFlagNameMySQLHost          = "mysql.host"
	cmdFlagNameMySQLPort          = "mysql.port"
	cmdFlagNameMySQLDatabase      = "mysql.database"
	cmdFlagNameMySQLUsername      = "mysql.username"
	cmdFlagNameMySQLPassword      = "mysql.password"
	cmdFlagNamePostgreSQLHost     = "postgres.host"
	cmdFlagNamePostgreSQLPort     = "postgres.port"
	cmdFlagNamePostgreSQLDatabase = "postgres.database"
	cmdFlagNamePostgreSQLSchema   = "postgres.schema"
	cmdFlagNamePostgreSQLUsername = "postgres.username"
	cmdFlagNamePostgreSQLPassword = "postgres.password"
)

const (
	cmdUseHashPassword  = "hash-password [flags] -- [password]"
	cmdUseHash          = "hash"
	cmdUseHashArgon2    = "argon2"
	cmdUseHashSHA2Crypt = "sha2crypt"
	cmdUseHashPBKDF2    = "pbkdf2"
	cmdUseHashBCrypt    = "bcrypt"
	cmdUseHashSCrypt    = "scrypt"

	cmdUseCrypto      = "crypto"
	cmdUseRand        = "rand"
	cmdUseCertificate = "certificate"
	cmdUseGenerate    = "generate"
	cmdUseValidate    = "validate"
	cmdUseFmtValidate = "%s [flags] -- <digest>"
	cmdUseRequest     = "request"
	cmdUsePair        = "pair"
	cmdUseRSA         = "rsa"
	cmdUseECDSA       = "ecdsa"
	cmdUseEd25519     = "ed25519"
)

const (
	cryptoCertPubCertOut = "certificate"
	cryptoCertCSROut     = "certificate signing request"

	prefixFilePassword = "authentication_backend.file.password"
)

var (
	errNoStorageProvider = errors.New("no storage provider configured")
)

const (
	identifierServiceOpenIDConnect = "openid"
	invalid                        = "invalid"
)

var (
	validIdentifierServices = []string{identifierServiceOpenIDConnect}
)
