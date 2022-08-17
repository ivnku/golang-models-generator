# Golang Models Generator

Golang Models Generator was created for generating structs which represent tables of an already existent database ("database first" approach).

## Usage

If you don't have golang installed on your machine then there is an already  built exe file in the `bin` folder (for linux you need to build it by yourself).

1. First of all, rename `config/config_test.yml` to `config/config.yml` and fill all the variables with your data. After you did it you can place this file in one folder with the binary otherwise you will need to provide a path to the config with the `-c` flag
2. You can edit the template file for a model as you wish, there is no requirements for a content of a template (but it must be a valid golang template).
3. Now you can run the binary: `./modgen.exe -c "<path_to_config>"` or without `-c` flag if config is in one folder with the binary. There are following available flags for execution:
    - `-c` - the path to the config file, './config.yml' by default
    - `-p` - the path where models will be generated to
    - `-t` - the name of the table you need to create a model for. If no name specified - generate models for all tables