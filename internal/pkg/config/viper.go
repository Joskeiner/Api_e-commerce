package config

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/viper"
)

type Viper struct {
	*viper.Viper
}

func newViper() (Config, error) {
	v := viper.New()

	path, err := os.Executable()
	if err != nil {
		return nil, err
	}
	dir := filepath.Dir(path)

	v.AddConfigPath(dir)
	fmt.Println(dir)
	v.SetConfigFile(".env")
	v.AutomaticEnv()

	err = v.ReadInConfig()

	if err != nil {
		return nil, err
	}
	return &Viper{
		v,
	}, nil
}

func (v *Viper) Load() (*Container, error) {
	app, err := v.newAppConfig()
	if err != nil {
		return nil, err
	}

	http, err := v.newHttpConfig()
	if err != nil {
		return nil, err
	}

	/*db, err := v.newDBConfig()
	if err != nil {
		return nil, err
	}

	token, err := v.newTokenConfig()
	if err != nil {
		return nil, err
	}*/

	return &Container{
		App:  app,
		Http: http,
		// Databese: db,
		// Token:    token,
	}, nil
}

func (v *Viper) newAppConfig() (*App, error) {
	var app App
	err := v.Viper.Unmarshal(&app)
	if err != nil {
		return nil, err
	}
	return &app, nil
}

func (v *Viper) newHttpConfig() (*Http, error) {
	var http Http

	err := v.Viper.Unmarshal(&http)
	if err != nil {
		return nil, err
	}

	return &http, nil
}

/*func (v *Viper) newDBConfig() (*Databese, error) {
	var db Databese

	err := v.Viper.Unmarshal(&db)
	if err != nil {
		return nil, err
	}
	return &db, nil
}*/

/*func (v *Viper) newTokenConfig() (*Token, error) {
	var token Token
	err := v.Viper.Unmarshal(&token)
	if err != nil {
		return nil, err
	}
	return &token, nil
}*/
