package main

import (
	"log"
	"os"
	"os/exec"
	"text/template"
)

type WPConfig struct {
	DBName         string
	DBUser         string
	DBPassword     string
	DBHost         string
	DBCharset      string
	DBCollate      string
	AuthKey        string
	SecureAuthKey  string
	LoggedInKey    string
	NonceKey       string
	AuthSalt       string
	SecureAuthSalt string
	LoggedInSalt   string
	NonceSalt      string
	TablePrefix    string
	WPDebug        string
	RedisClient    string
	RedisHost      string
	RedisPort      string
	RedisScheme    string
	RedisPath      string
	RedisServers   string
}

const wpConfigTemplate = `<?php
define( 'DB_NAME', '{{ .DBName }}' );
define( 'DB_USER', '{{ .DBUser }}' );
define( 'DB_PASSWORD', '{{ .DBPassword }}' );
define( 'DB_HOST', '{{ .DBHost }}' );

{{- if ne .DBCharset "" }}
define( 'DB_CHARSET', '{{ .DBCharset }}' );
{{- else }}
define( 'DB_CHARSET', 'utf8' );
{{- end }}

define( 'DB_COLLATE', '{{ .DBCollate }}' );

define( 'AUTH_KEY',         '{{ .AuthKey }}' );
define( 'SECURE_AUTH_KEY',  '{{ .SecureAuthKey }}' );
define( 'LOGGED_IN_KEY',    '{{ .LoggedInKey }}' );
define( 'NONCE_KEY',        '{{ .NonceKey }}' );
define( 'AUTH_SALT',        '{{ .AuthSalt }}' );
define( 'SECURE_AUTH_SALT', '{{ .SecureAuthSalt }}' );
define( 'LOGGED_IN_SALT',   '{{ .LoggedInSalt }}' );
define( 'NONCE_SALT',       '{{ .NonceSalt }}' );

$table_prefix = '{{ .TablePrefix }}';

{{- if eq .WPDebug "true" }}
define( 'WP_DEBUG', true );
{{- end }}

{{- if ne .RedisClient "" }}
define( 'WP_REDIS_CLIENT', '{{ .RedisClient }}' );

{{- if eq .RedisClient "relay" }}
define( 'WP_REDIS_DATABASE', 0 );
define( 'WP_REDIS_PREFIX', 'db3:' );
define( 'WP_REDIS_IGBINARY', true );
{{- end }}

{{- end }}

{{- if ne .RedisHost "" }}
define( 'WP_REDIS_HOST', '{{ .RedisHost }}' );
{{- end }}

{{- if ne .RedisPort "" }}
define( 'WP_REDIS_PORT', '{{ .RedisPort }}' );
{{- end }}

{{- if eq .RedisScheme "unix" }}
define( 'WP_REDIS_SCHEME', 'unix' );
define( 'WP_REDIS_PATH', '{{ .RedisPath }}' );
{{- end }}

{{- if ne .RedisServers "" }}
define( 'WP_REDIS_SERVERS', {{ .RedisServers }} );
{{- end }}


// If we're behind a proxy server and using HTTPS, we need to alert WordPress of that fact
if (isset($_SERVER['HTTP_X_FORWARDED_PROTO']) && strpos($_SERVER['HTTP_X_FORWARDED_PROTO'], 'https') !== false) {
	$_SERVER['HTTPS'] = 'on';
}

if ( ! defined( 'ABSPATH' ) ) {
	define( 'ABSPATH', __DIR__ . '/' );
}
require_once ABSPATH . 'wp-settings.php';
`

func main() {
	// TODO: Validate that all required environment variables are set
	config := WPConfig{
		DBName:         os.Getenv("WORDPRESS_DB_NAME"),
		DBUser:         os.Getenv("WORDPRESS_DB_USER"),
		DBPassword:     os.Getenv("WORDPRESS_DB_PASSWORD"),
		DBHost:         os.Getenv("WORDPRESS_DB_HOST"),
		DBCharset:      os.Getenv("WORDPRESS_DB_CHARSET"),
		DBCollate:      os.Getenv("WORDPRESS_DB_COLLATE"),
		AuthKey:        os.Getenv("WORDPRESS_AUTH_KEY"),
		SecureAuthKey:  os.Getenv("WORDPRESS_SECURE_AUTH_KEY"),
		LoggedInKey:    os.Getenv("WORDPRESS_LOGGED_IN_KEY"),
		NonceKey:       os.Getenv("WORDPRESS_NONCE_KEY"),
		AuthSalt:       os.Getenv("WORDPRESS_AUTH_SALT"),
		SecureAuthSalt: os.Getenv("WORDPRESS_SECURE_AUTH_SALT"),
		LoggedInSalt:   os.Getenv("WORDPRESS_LOGGED_IN_SALT"),
		NonceSalt:      os.Getenv("WORDPRESS_NONCE_SALT"),
		TablePrefix:    os.Getenv("WORDPRESS_TABLE_PREFIX"),
		WPDebug:        os.Getenv("WORDPRESS_DEBUG"),
		RedisClient:    os.Getenv("WP_REDIS_CLIENT"),
		RedisHost:      os.Getenv("WP_REDIS_HOST"),
		RedisPort:      os.Getenv("WP_REDIS_PORT"),
		RedisScheme:    os.Getenv("WP_REDIS_SCHEME"),
		RedisPath:      os.Getenv("WP_REDIS_PATH"),
		RedisServers:   os.Getenv("WP_REDIS_SERVERS"),
	}
	tmpl, err := template.New("wp-config").Parse(wpConfigTemplate)
	if err != nil {
		log.Fatalf("Error parsing template: %v", err)
	}

	f, err := os.Create("/var/www/html/wp-config.php")
	if err != nil {
		log.Fatalf("Error creating wp-config.php: %v", err)
	}
	defer f.Close()

	// Write the configuration to wp-config.php
	if err := tmpl.Execute(f, config); err != nil {
		log.Fatalf("Error writing wp-config.php: %v", err)
	}

	// If a command is provided, execute it
	if len(os.Args) > 1 {
		cmd := os.Args[1]
		args := os.Args[2:]
		if err := exec.Command(cmd, args...).Run(); err != nil {
			log.Fatalf("Error executing command: %v", err)
		}
	}
}
