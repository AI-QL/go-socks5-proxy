package main

import (
	"log"            // Importing the log package for logging
	"net"            // Importing the net package for network operations
	"os"             // Importing the os package for environment variables and standard I/O

	"github.com/AI-QL/go-socks5"  // Importing the go-socks5 package for SOCKS5 proxy functionality
	"github.com/caarlos0/env/v11" // Importing the env package for environment variable parsing
)

// Define a struct to hold the configuration parameters
type params struct {
	User            string   `env:"PROXY_USER" envDefault:""`       // Username for proxy authentication
	Password        string   `env:"PROXY_PASSWORD" envDefault:""`   // Password for proxy authentication
	Port            string   `env:"PROXY_PORT" envDefault:"1080"`   // Port number for the proxy server
	AllowedDestFqdn string   `env:"ALLOWED_DEST_FQDN" envDefault:""` // Allowed destination FQDN (Fully Qualified Domain Name)
	AllowedIPs      []string `env:"ALLOWED_IPS" envSeparator:"," envDefault:""` // Allowed IP addresses, comma-separated
}

func main() {
	// Initialize the configuration struct
	cfg := params{}

	// Parse environment variables into the configuration struct
	err := env.Parse(&cfg)
	if err != nil {
		// Log any errors that occur during parsing
		log.Printf("%+v\n", err)
	}

	// Initialize the SOCKS5 configuration
	socks5conf := &socks5.Config{
		Logger: log.New(os.Stdout, "", log.LstdFlags), // Set the logger to standard output
	}

	// If both User and Password are provided, configure authentication
	if cfg.User+cfg.Password != "" {
		creds := socks5.StaticCredentials{
			cfg.User: cfg.Password, // Use the provided username and password
		}
		cator := socks5.UserPassAuthenticator{Credentials: creds} // Create an authenticator
		socks5conf.AuthMethods = []socks5.Authenticator{cator} // Set the authenticator in the configuration
	}

	// If AllowedDestFqdn is provided, set the rule to allow only the specified destination
	if cfg.AllowedDestFqdn != "" {
		socks5conf.Rules = PermitDestAddrPattern(cfg.AllowedDestFqdn)
	}

	// Create a new SOCKS5 server with the configured settings
	server, err := socks5.New(socks5conf)
	if err != nil {
		// Log and exit if there is an error creating the server
		log.Fatal(err)
	}

	// Set the IP whitelist if allowed IPs are provided
	if len(cfg.AllowedIPs) > 0 {
		whitelist := make([]net.IP, len(cfg.AllowedIPs)) // Create a slice to hold the allowed IPs
		for i, ip := range cfg.AllowedIPs {
			whitelist[i] = net.ParseIP(ip) // Parse each IP and add it to the whitelist
		}
		server.SetIPAllowlist(whitelist) // Set the whitelist on the server
	}

	// Log the port on which the proxy service is listening
	log.Printf("Start listening proxy service on port %s\n", cfg.Port)

	// Start the server to listen and serve on the specified port
	if err := server.ListenAndServe("tcp", ":"+cfg.Port); err != nil {
		// Log and exit if there is an error starting the server
		log.Fatal(err)
	}
}