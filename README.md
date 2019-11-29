# extensions/brew

This extension helps with creating a brew formula for your tap repository

## Parameters

| Parameter   | Type   | Values                                                                  |
| ----------- | ------ | ----------------------------------------------------------------------- |
| binaryURL   | string | The url to the binary file you want to install with this formula        |
| description | string | The description of what the application in the formula does             |
| formula     | string | The name of the formula                                                 |
| formulaDir  | string | The directory where the formulas are stored; defaults to Formula        |
| homepage    | string | The homepage of the application                                         |
| tapRepoDir  | string | The directory to the `homebrew-<repo>` repository                       |
| version     | string | The version of the application; defaults to the Estafette build version |

## Usage

In order to use this extension in your `.estafette.yaml` manifest use the following snippets:

### Linting

```yaml
  create-brew-formula:
    image: extensions/brew:stable
    formula: myformula
    description: This is my awesome own formula that does abc
    homepage: https://estafette.io
    binaryURL: https://github.com/estafette/estafette/releases/download/v${ESTAFETTE_BUILD_VERSION}/estafette-v${ESTAFETTE_BUILD_VERSION}-darwin-amd64.zip
    tapRepoDir: homebrew-tap
    formulaDir: Formula    
```
