package main

import (
	"bytes"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"io"
	"net/http"
	"os"
	"runtime"

	"text/template"

	"github.com/alecthomas/kingpin"
	foundation "github.com/estafette/estafette-foundation"
	"github.com/iancoleman/strcase"
	"github.com/rs/zerolog/log"
	"gopkg.in/yaml.v2"
)

var (
	appgroup  string
	app       string
	version   string
	branch    string
	revision  string
	buildDate string
	goVersion = runtime.Version()
)

var (
	paramsYAML   = kingpin.Flag("params-yaml", "Extension parameters, created from custom properties.").Envar("ESTAFETTE_EXTENSION_CUSTOM_PROPERTIES_YAML").Required().String()
	buildVersion = kingpin.Flag("build-version", "Version number, used if not passed explicitly.").Envar("ESTAFETTE_BUILD_VERSION").String()
)

func main() {

	// parse command line parameters
	kingpin.Parse()

	// init log format from envvar ESTAFETTE_LOG_FORMAT
	foundation.InitLoggingFromEnv(appgroup, app, version, branch, revision, buildDate)

	log.Info().Msg("Unmarshalling parameters / custom properties...")
	var params Params
	err := yaml.Unmarshal([]byte(*paramsYAML), &params)
	if err != nil {
		log.Fatal().Err(err).Msg("Failed unmarshalling parameters")
	}

	// set defaults
	params.SetDefaults(*buildVersion)

	// validate parameters
	valid, warnings := params.Validate()
	if !valid {
		log.Fatal().Msgf("Some parameters are not valid: %v", warnings)
	}

	// create target file to render template to
	targetFilePath := fmt.Sprintf("%v/%v/%v.rb", params.TapReposityDirectory, params.FormulaDirectory, params.Formula)
	targetFile, err := os.Create(targetFilePath)
	defer targetFile.Close()

	// read and parse template
	formulaTemplate, err := template.ParseFiles("/templates/formula.rb")
	if err != nil {
		log.Fatal().Err(err).Msg("Failed parsing template")
	}

	// download binary
	resp, err := http.Get(params.BinaryURL)
	if err != nil {
		log.Fatal().Err(err).Msgf("Failed retrieving binary from url %v", params.BinaryURL)
	}
	defer resp.Body.Close()

	// calculate sh256 checksum
	hasher := sha256.New()
	if _, err := io.Copy(hasher, resp.Body); err != nil {
		log.Fatal().Err(err).Msgf("Failed calculating sha256 checksum for binary from url %v", params.BinaryURL)
	}

	data := struct {
		Formula          string
		FormulaClassName string
		Description      string
		Homepage         string
		BinaryURL        string
		Version          string
		Sha256           string
	}{
		params.Formula,
		strcase.ToCamel(params.Formula),
		params.Description,
		params.Homepage,
		params.BinaryURL,
		params.Version,
		hex.EncodeToString(hasher.Sum(nil)),
	}

	// write rendered template to file
	err = formulaTemplate.Execute(targetFile, data)
	if err != nil {
		log.Fatal().Err(err).Msg("Failed writing template to file")
	}

	// log template to stdout
	var renderedTemplate bytes.Buffer
	err = formulaTemplate.Execute(&renderedTemplate, data)
	if err != nil {
		log.Fatal().Err(err).Msg("Failed rendering template for stdout")
	}

	log.Info().Msg(renderedTemplate.String())
}
