Package yaml provides a parser, linter, formatter and compiler for the [drone](https://github.com/drone/drone) configuration file format.

Lint the yaml file:

```text
$ drone-yaml lint samples/simple.yml
```

Format the yaml file:

```text
$ drone-yaml fmt samples/simple.yml
$ drone-yaml fmt samples/simple.yml --save
```

Sign the yaml file using a 32-bit secret key:

```text
$ drone-yaml sign 642909eb4c3d47e33999235c0598353c samples/simple.yml
$ drone-yaml sign 642909eb4c3d47e33999235c0598353c samples/simple.yml --save
```

Verify the yaml file signature:

```text
$ drone-yaml verify 642909eb4c3d47e33999235c0598353c samples/simple.yml
```

Compile the yaml file:

```text
$ drone-yaml compile samples/simple.yml > samples/simple.json
```
