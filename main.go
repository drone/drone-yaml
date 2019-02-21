// Copyright 2019 Drone.IO Inc. All rights reserved.
// Use of this source code is governed by the Drone Non-Commercial License
// that can be found in the LICENSE file.

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
	"strings"

	"github.com/drone/drone-runtime/engine"
	"github.com/drone/drone-yaml/yaml"
	"github.com/drone/drone-yaml/yaml/compiler"
	"github.com/drone/drone-yaml/yaml/compiler/transform"
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

	convert     = kingpin.Command("convert", "convert the yaml file")
	convertSave = convert.Flag("save", "save result to source").Short('s').Bool()
	convertFile = convert.Arg("source", "source file location").Default(".drone.yml").File()

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
	case convert.FullCommand():
		kingpin.FatalIfError(runConvert(), "")
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

func runConvert() error {
	f := *convertFile
	d, err := ioutil.ReadAll(f)
	if err != nil {
		return err
	}
	m := converter.Metadata{
		Filename: f.Name(),
	}
	b, err := converter.Convert(d, m)
	if err != nil {
		return err
	}
	if *formatSave {
		return ioutil.WriteFile(f.Name(), b, 0644)
	}
	_, err = io.Copy(os.Stderr, bytes.NewReader(b))
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
	trusted    = compile.Flag("trusted", "trusted mode").Bool()
	labels     = compile.Flag("label", "container labels").StringMap()
	clone      = compile.Flag("clone", "clone step").Bool()
	volume     = compile.Flag("volume", "attached volumes").StringMap()
	network    = compile.Flag("network", "attached networks").Strings()
	environ    = compile.Flag("env", "environment variable").StringMap()
	dind       = compile.Flag("dind", "dind images").Default("plugins/docker").Strings()
	event      = compile.Flag("event", "event type").PlaceHolder("<event>").Enum("push", "pull_request", "tag", "deployment")
	repo       = compile.Flag("repo", "repository name").PlaceHolder("octocat/hello-world").String()
	remote     = compile.Flag("git-remote", "git remote url").PlaceHolder("https://github.com/octocat/hello-world.git").String()
	branch     = compile.Flag("git-branch", "git commit branch").PlaceHolder("master").String()
	ref        = compile.Flag("git-ref", "git commit ref").PlaceHolder("refs/heads/master").String()
	sha        = compile.Flag("git-sha", "git commit sha").String()
	creds      = compile.Flag("git-creds", "git credentials").URLList()
	instance   = compile.Flag("instance", "drone instance hostname").PlaceHolder("drone.company.com").String()
	deploy     = compile.Flag("deploy-to", "target deployment").PlaceHolder("production").String()
	secrets    = compile.Flag("secret", "secret variable").StringMap()
	registries = compile.Flag("registry", "registry credentials").URLList()
	username   = compile.Flag("netrc-username", "netrc username").PlaceHolder("<token>").String()
	password   = compile.Flag("netrc-password", "netrc password").PlaceHolder("x-oauth-basic").String()
	machine    = compile.Flag("netrc-machine", "netrc machine").PlaceHolder("github.com").String()
	memlimit   = compile.Flag("mem-limit", "memory limit").PlaceHolder("1GB").Bytes()
	cpulimit   = compile.Flag("cpu-limit", "cpu limit").PlaceHolder("2").Int64()
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

	// the user has the option to disable the git clone
	// if the pipeline is being executed on the local
	// codebase.
	if *clone == false {
		p.Clone.Disable = true
	}

	var auths []*engine.DockerAuth
	for _, uri := range *registries {
		if uri.User == nil {
			log.Fatalln("Expect registry format [user]:[password]@hostname")
		}
		password, ok := uri.User.Password()
		if !ok {
			log.Fatalln("Invalid or missing registry password")
		}
		auths = append(auths, &engine.DockerAuth{
			Address:  uri.Host,
			Username: uri.User.Username(),
			Password: password,
		})
	}

	comp := new(compiler.Compiler)
	comp.GitCredentialsFunc = defaultCreds // TODO create compiler.GitCredentialsFunc and compiler.GlobalGitCredentialsFunc
	comp.NetrcFunc = nil                   // TODO create compiler.NetrcFunc and compiler.GlobalNetrcFunc
	comp.PrivilegedFunc = compiler.DindFunc(*dind)
	comp.SkipFunc = compiler.SkipFunc(
		compiler.SkipData{
			Branch:   *branch,
			Event:    *event,
			Instance: *instance,
			Ref:      *ref,
			Repo:     *repo,
			Target:   *deploy,
		},
	)
	comp.TransformFunc = transform.Combine(
		transform.WithAuths(auths),
		transform.WithEnviron(*environ),
		transform.WithEnviron(defaultEnvs()),
		transform.WithLables(*labels),
		transform.WithLimits(int64(*memlimit), int64(*cpulimit)),
		transform.WithNetrc(*machine, *username, *password),
		transform.WithNetworks(*network),
		transform.WithProxy(),
		transform.WithSecrets(*secrets),
		transform.WithVolumes(*volume),
	)
	compiled := comp.Compile(p)

	// // for drone-exec we will need to change the workspace
	// // to a host volume mount, to the current working dir.
	// for _, volume := range compiled.Docker.Volumes {
	// 	if volume.Metadata.Name == "workspace" {
	// 		volume.EmptyDir = nil
	// 		volume.HostPath = &engine.VolumeHostPath{
	// 			Path: "", // pwd
	// 		}
	// 		break
	// 	}
	// }
	// // then we need to change the base mount for every container
	// // to use the workspace base + path.
	// for _, container := range compiled.Steps {
	// 	for _, volume := range container.Volumes {
	// 		if volume.Name == "workspace" {
	// 			volume.Path = container.Envs["DRONE_WORKSPACE"]
	// 		}
	// 	}
	// }

	enc := json.NewEncoder(os.Stdout)
	enc.SetIndent("", "  ")
	return enc.Encode(compiled)
}

// helper function returns the git credential function,
// used to return a git credentials file.
func defaultCreds() []byte {
	urls := *creds
	if len(urls) == 0 {
		return nil
	}
	var buf bytes.Buffer
	for _, url := range urls {
		buf.WriteString(url.String())
		buf.WriteByte('\n')
	}
	return buf.Bytes()
}

// helper function returns the minimum required environment
// variables to clone a repository. All other environment
// variables should be passed via the --env flag.
func defaultEnvs() map[string]string {
	envs := map[string]string{}
	envs["DRONE_COMMIT_BRANCH"] = *branch
	envs["DRONE_COMMIT_SHA"] = *sha
	envs["DRONE_COMMIT_REF"] = *ref
	envs["DRONE_REMOTE_URL"] = *remote
	envs["DRONE_BUILD_EVENT"] = *event
	if strings.HasPrefix(*ref, "refs/tags/") {
		envs["DRONE_TAG"] = strings.TrimPrefix(*ref, "refs/tags/")
	}
	return envs
}
