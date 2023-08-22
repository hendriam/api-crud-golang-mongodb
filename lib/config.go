package lib

type Config struct {
	Server struct {
		Host string
		Port int
	}

	Database struct {
		MongoDB struct {
			Dsn string
		}
	}

	Log struct {
		Level string
	}
}

func LoadConfig() Config {
	return Config{
		Server: struct {
			Host string
			Port int
		}{
			Host: "localhost",
			Port: 8080,
		},

		Database: struct{ MongoDB struct{ Dsn string } }{
			MongoDB: struct{ Dsn string }{
				Dsn: "mongodb://127.0.0.1:27017",
			},
		},

		Log: struct{ Level string }{
			Level: "info",
		},
	}
}
