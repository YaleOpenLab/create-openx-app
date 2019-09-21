# Template

Template is a series of basic handlers and a model smart contract built on top of Stellar modelled based on [opensolar](https://github.com/YaleOpenLab/opensolar/)

## Editing template files

The structure of the template is centred around the `core` subrepo and the `rpc` subrepo contains handlers for others to query from. Inside core is `openx` which contains handlers for interacting with openx. Template files are designed to be modular in nature and apps that build on top of the template should leverage this as much as possible.
