package media

import (
	"github.com/pkg/errors"
	"github.com/spf13/viper"
	"golang.org/x/exp/slices"
	"os"
	"regexp"
)

type Media struct {
	Name             string
	Match            string
	OutputExpression string
	OutputLocation   string
}

func New(name, match, outputExpression, outputLocation string) (*Media, error) {
	media := &Media{
		Name:             name,
		Match:            match,
		OutputExpression: outputExpression,
		OutputLocation:   outputLocation,
	}

	err := validate(media)
	return media, err
}

func validate(media *Media) error {
	if media == nil {
		return errors.New("media is nil")
	}

	_, err := regexp.Compile(media.Match)
	if err != nil {
		return errors.Wrap(err, "the match isn't a valid regex")
	}

	_, err = regexp.Compile(media.OutputExpression)
	if err != nil {
		return errors.Wrap(err, "the output expression isn't a valid regex")
	}

	err = os.MkdirAll(media.OutputLocation, 0755)
	if err != nil {
		return errors.Wrap(err, "could not make sure the output location exists")
	}

	return nil
}

func Add(media *Media) error {
	err := validate(media)
	if err != nil {
		return errors.Wrap(err, "invalid media")
	}

	exists, err := Exists(media.Name)
	if err != nil {
		return errors.Wrapf(err, "error when checking if media with name %q already exists", media.Name)
	}
	if exists {
		return errors.Errorf("media with name %q already exists", media.Name)
	}

	allMedia, err := List()
	if err != nil {
		return errors.Wrap(err, "error getting all media")
	}

	allMedia = append(allMedia, *media)
	viper.Set("media", allMedia)
	err = viper.WriteConfig()
	if err != nil {
		return errors.Wrap(err, "error writing changes to disk")
	}

	return nil
}

func List() ([]Media, error) {
	var media []Media
	err := viper.UnmarshalKey("media", &media)
	if err != nil {
		return nil, errors.Wrap(err, "error reading from config")
	}

	return media, nil
}

func Exists(name string) (bool, error) {
	media, err := List()
	if err != nil {
		return false, errors.Wrap(err, "error getting all media")
	}
	return slices.ContainsFunc(media, func(m Media) bool {
		return m.Name == name
	}), nil
}