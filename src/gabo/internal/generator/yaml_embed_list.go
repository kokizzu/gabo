package generator

import _ "embed"

var (
	//go:embed data/format-go.yaml
	_formatGoYaml string
	//go:embed data/lint-android.yaml
	_lintAndroidYaml string
	//go:embed data/lint-docker.yaml
	_lintDockerYaml string
	//go:embed data/lint-go-incomplete.yaml
	_lintGoTemplate string
	//go:embed data/lint-markdown.yaml
	_lintMarkdownYaml string
	//go:embed data/lint-python.yaml
	_lintPythonYaml string
	//go:embed data/lint-shell-script.yaml
	_lintShellScriptYaml string
	//go:embed data/lint-solidity.yaml
	_lintSolidityYaml string
	//go:embed data/lint-yaml.yaml
	_lintYamlYaml string

	//go:embed data/build-android-incomplete.yaml
	_buildAndroidTemplate string
	//go:embed data/build-docker-incomplete.yaml
	_buildDockerTemplate string

	//go:embed data/translate-android.yaml
	_translateAndroidYaml string
	//go:embed data/validate-openapi-schema.incomplete.yaml
	_generateOpenAPISchemaValidatorTemplate string
)
