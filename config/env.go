package config

type Config struct {
	PublicHost             string
	Port                   string
	DBUser                 string
	DBPassword             string
	DBAddress              string
	DBName                 string
	JWTExpirationInSeconds int64
	JWTSecret              string
	ImageURL               string
}

var Envs = initConfig()

// Implement initConfig function that returns Config struct with default values
// Example:
// func initConfig() Config {
// 	return Config{
// 		PublicHost:             "http://localhost",
// 		Port:                   ":8080",
// 		DBUser:                 "username",
// 		DBPassword:             "paswd",
// 		DBAddress:              "127.0.0.1:3306",
// 		DBName:                 "db name",
// 		JWTSecret:              "your-secret-here",
// 		JWTExpirationInSeconds: 3600 * 24 * 7, // 1 week
// 		ImageURL:               "https://images.pexels.com/photos/27567342/pexels-photo-27567342/free-photo-of-a-barn-sits-on-a-hillside-in-the-middle-of-a-field.jpeg?auto=compress&cs=tinysrgb&w=1260&h=750&dpr=1",
// 	}
// }
