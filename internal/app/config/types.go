package config

type Config struct {
	Server TServerConfig
	Middleware TMiddlewareConfig
	Kafka TKafkaConfig
}

type TServerConfig struct {
	Environment string
	Port 	 int
}

type TMiddlewareConfig struct {
	ClientID           bool // Whether Client-ID middleware is enabled.
	DeviceID           bool // Whether Device-ID middleware is enabled.
	RequestID          bool // Whether Request-ID middleware is enabled.
	Idempotency        bool // Whether Idempotency middleware is enabled.
	GzipCompression    bool // Whether Gzip compression middleware is enabled.
	UserAgent          bool // Whether User-Agent middleware is enabled.
	SecurityHeaders    bool // Whether security headers middleware is enabled.
	CORS               bool // Whether CORS middleware is enabled.
	Cache              bool // Whether default cache behavior is enabled.
	ContentNegotiation bool // Whether content negotiation middleware is enabled.
	Referer            bool // Whether Referer header middleware is enabled.
	Cookies            bool // Whether Cookies middleware is enabled.
	RateLimit          bool // Whether rate-limiting middleware is enabled.
	Logger             bool // Whether logging middleware is enabled.
	Metrics            bool // Whether metrics middleware is enabled.
	Auth               bool // Whether authentication middleware is enabled.
	Profiler           bool // Whether profiling middleware is enabled.
}

type TKafkaConfig struct {
	Brokers       []string          `mapstructure:"brokers"`
	ClientID      string            `mapstructure:"client_id"`
	Acks          string            `mapstructure:"acks"`
	Async         bool              `mapstructure:"async"`
	RetryAttempts int               `mapstructure:"retry_attempts"`
	WriteTimeout  int               `mapstructure:"write_timeout"`
	ReadTimeout   int               `mapstructure:"read_timeout"`

	Topics map[string]string `mapstructure:"topics"` // "logs", "events", "context"

	// Auth (optional)
	EnableTLS     bool   `mapstructure:"enable_tls"`
	EnableSASL    bool   `mapstructure:"enable_sasl"`
	SASLUser      string `mapstructure:"sasl_user"`
	SASLPassword  string `mapstructure:"sasl_password"`
	SASLMechanism string `mapstructure:"sasl_mechanism"`
}

// Global variables holding configuration instances.
var (
	ServerConfig     TServerConfig     // Stores server configuration.
	KafkaConfig      TKafkaConfig      // Stores Kafka configuration.
	MiddlewareConfig TMiddlewareConfig // Stores middleware configuration.
)