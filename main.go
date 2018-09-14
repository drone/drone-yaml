package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"

	"github.com/drone/drone-yaml-v1/config/compiler"
	"github.com/drone/drone-yaml/yaml"
	"github.com/drone/drone-yaml/yaml/converter"
	"github.com/drone/drone-yaml/yaml/linter"
	"github.com/drone/drone-yaml/yaml/pretty"
	"github.com/drone/drone-yaml/yaml/signer"

	"gopkg.in/alecthomas/kingpin.v2"
)

var (
	format     = kingpin.Command("fmt", "format the yaml file")
	formatPriv = format.Flag("privileged", "privileged mode").Short('p').Bool()
	formatSave = format.Flag("save", "save result to source").Short('s').Bool()
	formatFile = format.Arg("source", "source file location").Default(".drone.yml").File()

	lint     = kingpin.Command("lint", "lint the yaml file")
	lintPriv = lint.Flag("privileged", "privileged mode").Short('p').Bool()
	lintFile = lint.Arg("source", "source file location").Default(".drone.yml").File()

	sign     = kingpin.Command("sign", "sign the yaml file")
	signKey  = sign.Arg("key", "secret key").Required().String()
	signFile = sign.Arg("source", "source file location").Default(".drone.yml").File()
	signSave = sign.Flag("save", "save result to source").Short('s').Bool()

	verify     = kingpin.Command("verify", "verify the yaml signature")
	verifyKey  = verify.Arg("key", "secret key").Required().String()
	verifyFile = verify.Arg("source", "source file location").Default(".drone.yml").File()

	compile     = kingpin.Command("compile", "compile the yaml file")
	compileIn   = compile.Arg("source", "source file location").Default(".drone.yml").File()
	compileName = compile.Flag("name", "pipeline name").String()
)

func main() {
	switch kingpin.Parse() {
	case format.FullCommand():
		kingpin.FatalIfError(runFormat(), "")
	case lint.FullCommand():
		kingpin.FatalIfError(runLint(), "")
	case sign.FullCommand():
		kingpin.FatalIfError(runSign(), "")
	case verify.FullCommand():
		kingpin.FatalIfError(runVerify(), "")
	case compile.FullCommand():
		kingpin.FatalIfError(runCompile(), "")
	}
}

func runFormat() error {
	f := *formatFile
	m, err := yaml.Parse(f)
	if err != nil {
		return err
	}

	b := new(bytes.Buffer)
	pretty.Print(b, m)

	if *formatSave {
		return ioutil.WriteFile(f.Name(), b.Bytes(), 0644)
	}
	_, err = io.Copy(os.Stderr, b)
	return err
}

func runLint() error {
	f := *lintFile
	m, err := yaml.Parse(f)
	if err != nil {
		return err
	}
	for _, r := range m.Resources {
		err := linter.Lint(r, *lintPriv)
		if err != nil {
			return err
		}
	}
	return nil
}

func runSign() error {
	f := *signFile
	d, err := ioutil.ReadAll(f)
	if err != nil {
		return err
	}

	k := signer.KeyString(*signKey)

	if *signSave {
		out, err := signer.SignUpdate(d, k)
		if err != nil {
			return err
		}
		return ioutil.WriteFile(f.Name(), out, 0644)
	}

	hmac, err := signer.Sign(d, k)
	if err != nil {
		return err
	}
	fmt.Println(hmac)
	return nil
}

func runVerify() error {
	f := *verifyFile
	d, err := ioutil.ReadAll(f)
	if err != nil {
		return err
	}

	k := signer.KeyString(*verifyKey)
	ok, err := signer.Verify(d, k)
	if err != nil {
		return err
	} else if !ok {
		return errors.New("cannot verify yaml signature")
	}

	fmt.Println("success: yaml signature verified")
	return nil
}

var (
	trusted      = compile.Flag("trusted", "trusted mode").Bool()
	clone        = compile.Flag("clone", "clone step").Bool()
	volume       = compile.Flag("volume", "attached volumes").Strings()
	network      = compile.Flag("network", "attached networks").Strings()
	environ      = compile.Flag("env", "environment variable").StringMap()
	images       = compile.Flag("privileged", "privileged images").Default("plugins/docker").Strings()
	base         = compile.Flag("base", "workspace base path").Default("/workspace").String()
	path         = compile.Flag("path", "wrokspace path").String()
	event        = compile.Flag("event", "event type").PlaceHolder("<event>").Enum("push", "pull_request", "tag", "deployment")
	repo         = compile.Flag("repo", "repository name").PlaceHolder("octocat/hello-world").String()
	branch       = compile.Flag("git-branch", "git commit branch").PlaceHolder("master").String()
	ref          = compile.Flag("git-ref", "git commit ref").PlaceHolder("refs/heads/master").String()
	deploy       = compile.Flag("deploy-to", "target deployment").PlaceHolder("production").String()
	platform     = compile.Flag("platform", "target platform").PlaceHolder("linux/amd64").String()
	secrets      = compile.Flag("secret", "secret variable").StringMap()
	registries   = compile.Flag("registry", "registry credentials").URLList()
	username     = compile.Flag("netrc-login", "netrc username").PlaceHolder("<token>").String()
	password     = compile.Flag("netrc-password", "netrc password").PlaceHolder("x-oauth-basic").String()
	machine      = compile.Flag("netrc-machine", "netrc machine").PlaceHolder("github.com").String()
	cpuset       = compile.Flag("cpu-set", "cpu set").PlaceHolder("0,1").String()
	cpushares    = compile.Flag("cpu-shares", "cpu shares").PlaceHolder("75").Int64()
	cpuquota     = compile.Flag("cpu-quota", "cpu quota").PlaceHolder("7500").Int64()
	memlimit     = compile.Flag("mem-limit", "memory limit").PlaceHolder("1GB").Bytes()
	memswaplimit = compile.Flag("mem-swap-limit", "memory swap limit").PlaceHolder("1GB").Bytes()
	shmsize      = compile.Flag("shmsize", "shmsize").PlaceHolder("1GB").Bytes()
)

func runCompile() error {
	m, err := yaml.Parse(*compileIn)
	if err != nil {
		return err
	}

	var p *yaml.Pipeline
	for _, r := range m.Resources {
		v, ok := r.(*yaml.Pipeline)
		if !ok {
			continue
		}
		if *compileName == "" ||
			*compileName == v.Name {
			p = v
			break
		}
	}

	if p == nil {
		return errors.New("cannot find pipeline resource")
	}

	var secretList []compiler.Secret
	for k, v := range *secrets {
		secretList = append(secretList, compiler.Secret{
			Name:  k,
			Value: v,
		})
	}

	var registryList []compiler.Registry
	for _, uri := range *registries {
		if uri.User == nil {
			log.Fatalln("Expect registry format [user]:[password]@hostname")
		}
		password, ok := uri.User.Password()
		if !ok {
			log.Fatalln("Invalid or missing registry password")
		}
		registryList = append(registryList, compiler.Registry{
			Hostname: uri.Host,
			Username: uri.User.Username(),
			Password: password,
		})
	}

	var opts = []compiler.Option{
		compiler.WithClone(*clone),
		compiler.WithEnviron(*environ),
		compiler.WithLimits(
			compiler.Resources{
				CPUQuota:     *cpuquota,
				CPUShares:    *cpushares,
				CPUSet:       *cpuset,
				ShmSize:      int64(*shmsize),
				MemLimit:     int64(*memlimit),
				MemSwapLimit: int64(*memswaplimit),
			},
		),
		compiler.WithMetadata(
			compiler.Metadata{
				Branch:      *branch,
				Event:       *event,
				Ref:         *ref,
				Repo:        *repo,
				Platform:    *platform,
				Environment: *deploy,
			},
		),
		compiler.WithNetrc(*username, *password, *machine),
		compiler.WithNetworks(*network...),
		compiler.WithPrivileged(*images...),
		compiler.WithRegistry(registryList...),
		compiler.WithSecret(secretList...),
		compiler.WithVolumes(*volume...),
		compiler.WithWorkspace(*base, *path),
	}

	config := converter.Convert(p)
	out, err := compiler.New(opts...).Compile(config)
	if err != nil {
		return err
	}

	enc := json.NewEncoder(os.Stdout)
	enc.SetIndent("", "  ")
	return enc.Encode(out)
}
