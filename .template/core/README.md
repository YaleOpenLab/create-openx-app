# Core

Core contains the core functionality of the template platform. There are different types of files in core, intended to give people the ability to initialize their platforms with the maximum amount of customising options:

1. Gotemp files - gotemp files are files that replace their namesake .go files (eg. contract_stellar.gotemp replaces contract.go if a user wishes to have their platform support only Stellar). Gotemp files can be deleted without any consequence to the template itself but cannot exist as .go files in parallel to their parent .go files.

2. Option files - Option files are intended to provide extra options to entitites within the platform (eg. investor_options gives a user the option to add the voting options for investors). Option files can exist in parallel to their parent .go files and can be deleted without any consequence to their parent .go files.
