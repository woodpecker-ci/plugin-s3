package main

import (
	"os"

	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"github.com/urfave/cli/v3"
)

var version = "0.0.0+0"

func main() {
	app := cli.NewApp()
	app.Name = "s3 plugin"
	app.Usage = "s3 plugin"
	app.Version = version
	app.Action = run
	app.Flags = []cli.Flag{
		&cli.StringFlag{
			Name:    "endpoint",
			Usage:   "endpoint for the s3 connection",
			EnvVars: []string{"PLUGIN_ENDPOINT", "S3_ENDPOINT"},
		},
		&cli.StringFlag{
			Name:    "access-key",
			Usage:   "aws access key",
			EnvVars: []string{"PLUGIN_ACCESS_KEY", "AWS_ACCESS_KEY_ID"},
		},
		&cli.StringFlag{
			Name:    "secret-key",
			Usage:   "aws secret key",
			EnvVars: []string{"PLUGIN_SECRET_KEY", "AWS_SECRET_ACCESS_KEY"},
		},
		&cli.StringFlag{
			Name:    "assume-role",
			Usage:   "aws iam role to assume",
			EnvVars: []string{"PLUGIN_ASSUME_ROLE", "ASSUME_ROLE"},
		},
		&cli.StringFlag{
			Name:    "assume-role-session-name",
			Usage:   "aws iam role session name to assume",
			Value:   "woodpecker-s3",
			EnvVars: []string{"PLUGIN_ASSUME_ROLE_SESSION_NAME", "ASSUME_ROLE_SESSION_NAME"},
		},
		&cli.StringFlag{
			Name:    "bucket",
			Usage:   "aws bucket",
			Value:   "us-east-1",
			EnvVars: []string{"PLUGIN_BUCKET", "S3_BUCKET"},
		},
		&cli.StringFlag{
			Name:    "region",
			Usage:   "aws region",
			Value:   "us-east-1",
			EnvVars: []string{"PLUGIN_REGION", "S3_REGION"},
		},
		&cli.StringFlag{
			Name:    "acl",
			Usage:   "upload files with acl",
			Value:   "private",
			EnvVars: []string{"PLUGIN_ACL"},
		},
		&cli.StringFlag{
			Name:    "source",
			Usage:   "upload files from source folder",
			EnvVars: []string{"PLUGIN_SOURCE"},
		},
		&cli.StringFlag{
			Name:    "target",
			Usage:   "upload files to target folder",
			EnvVars: []string{"PLUGIN_TARGET"},
		},
		&cli.StringFlag{
			Name:    "strip-prefix",
			Usage:   "strip the prefix from the target",
			EnvVars: []string{"PLUGIN_STRIP_PREFIX"},
		},
		&cli.StringSliceFlag{
			Name:    "exclude",
			Usage:   "ignore files matching exclude pattern",
			EnvVars: []string{"PLUGIN_EXCLUDE"},
		},
		&cli.StringFlag{
			Name:    "encryption",
			Usage:   "server-side encryption algorithm, defaults to none",
			EnvVars: []string{"PLUGIN_ENCRYPTION"},
		},
		&cli.BoolFlag{
			Name:    "dry-run",
			Usage:   "dry run for debug purposes",
			EnvVars: []string{"PLUGIN_DRY_RUN"},
		},
		&cli.BoolFlag{
			Name:    "path-style",
			Usage:   "use path style for bucket paths",
			EnvVars: []string{"PLUGIN_PATH_STYLE"},
		},
		&cli.GenericFlag{
			Name:    "content-type",
			Usage:   "set content type header for uploaded objects",
			EnvVars: []string{"PLUGIN_CONTENT_TYPE"},
			Value:   &StringMapFlag{},
		},
		&cli.GenericFlag{
			Name:    "content-encoding",
			Usage:   "set content encoding header for uploaded objects",
			EnvVars: []string{"PLUGIN_CONTENT_ENCODING"},
			Value:   &StringMapFlag{},
		},
		&cli.GenericFlag{
			Name:    "cache-control",
			Usage:   "set cache-control header for uploaded objects",
			EnvVars: []string{"PLUGIN_CACHE_CONTROL"},
			Value:   &StringMapFlag{},
		},
		&cli.StringFlag{
			Name:    "storage-class",
			Usage:   "set storage class to choose the best backend",
			EnvVars: []string{"PLUGIN_STORAGE_CLASS"},
		},
		&cli.StringFlag{
			Name:    "env-file",
			Usage:   "source env file",
			EnvVars: []string{"PLUGIN_ENV_FILE"},
		},
		&cli.BoolFlag{
			Name:    "compress",
			Usage:   "prior to upload, compress files and use gzip content-encoding",
			EnvVars: []string{"PLUGIN_COMPRESS"},
		},
		&cli.BoolFlag{
			Name:    "overwrite",
			Usage:   "overwrite existing files",
			EnvVars: []string{"PLUGIN_OVERWRITE"},
		},
	}

	if err := app.Run(os.Args); err != nil {
		logrus.Fatal(err)
	}
}

func run(c *cli.Context) error {
	if c.String("env-file") != "" {
		_ = godotenv.Load(c.String("env-file"))
	}

	plugin := Plugin{
		Endpoint:              c.String("endpoint"),
		Key:                   c.String("access-key"),
		Secret:                c.String("secret-key"),
		AssumeRole:            c.String("assume-role"),
		AssumeRoleSessionName: c.String("assume-role-session-name"),
		Bucket:                c.String("bucket"),
		Region:                c.String("region"),
		Access:                c.String("acl"),
		Source:                c.String("source"),
		Target:                c.String("target"),
		StripPrefix:           c.String("strip-prefix"),
		Exclude:               c.StringSlice("exclude"),
		Encryption:            c.String("encryption"),
		ContentType:           c.Generic("content-type").(*StringMapFlag).Get(),
		ContentEncoding:       c.Generic("content-encoding").(*StringMapFlag).Get(),
		CacheControl:          c.Generic("cache-control").(*StringMapFlag).Get(),
		StorageClass:          c.String("storage-class"),
		PathStyle:             c.Bool("path-style"),
		DryRun:                c.Bool("dry-run"),
		Compress:              c.Bool("compress"),
		Overwrite:             c.Bool("overwrite"),
	}

	return plugin.Exec()
}
