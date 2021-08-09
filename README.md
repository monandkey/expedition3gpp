
# expedition3gpp

## Purpose

The purpose is to download the standardization documents developed by 3GPP.

It also caches HTML information to speed up document retrieval the second and subsequent times.

## Download or Build

In order to use this tool, you need to download the binary or compile the source code.

### Download

```bash
$ echo "linux"
$ wget https://github.com/monandkey/expedition3gpp/files/6824166/expedition3gpp-linux-amd64.gz
$
$ echo "windows"
$ wget https://github.com/monandkey/expedition3gpp/files/6824169/expedition3gpp-windows.zip
```

### Build

```bash
$ cd expedition3gpp
$ cd make
$ cd docker compose up -d
```

## Usage

### Main

```
$ expedition3gpp
Download the 3GPP document

Usage:
  expedition3gpp [flags]
  expedition3gpp [command]

Available Commands:
  completion  generate the autocompletion script for the specified shell
  download    Download the 3GPP documentation.
  help        Help about any command
  init        Create the config file
  search      Search for 3GPP documentation.

Flags:
  -h, --help      help for expedition3gpp
  -v, --version   display version

Use "expedition3gpp [command] --help" for more information about a command.
```

### init

This command does not need to be explicitly executed.

The default configuration will be created when the normal command is executed.

Execute this command if you want to change the parameters of the config.

```bash
$ expedition3gpp init -h
Create the config file

Usage:
  expedition3gpp init [flags]

Flags:
      --cache-enable               Enable or disable the cache
                                   true  -> enable
                                   false -> disable
                                    (default true)
      --cache-location string      Specify the location to save the cache.
                                   windows -> C:\Users\testuser
                                   linux   -> /home/testuser
                                    (default "HOMEDIR")
      --cache-retention-time int   Specify the validity period for saving the cache.
                                   [0...4294967295]
                                    (default 14400)
  -h, --help                       help for init
      --strage-location string     Specify the location to save the config.
                                   windows -> C:\Users\testuser
                                   linux   -> /home/testuser
                                    (default "HOMEDIR")
```

### search

Use this command to get the version information of the document you want to download.

However, the download command will allow you to `download` even if you do not specify the version.

```bash
$ expedition3gpp search -h
Usage:
  expedition3gpp search [flags]

Flags:
      --document-number string    3GPP Document Number
      --document-version string   3GPP Document Version
  -h, --help                      help for search
      --no-cache                  Not using cache
```

### download

Download the document by specifying the document number and version.

```bash
$ expedition3gpp download -h
Usage:
  expedition3gpp download [flags]

Flags:
      --document-number string    3GPP Document Number
      --document-version string   3GPP Document Version
  -h, --help                      help for download
      --no-cache                  Not using cache
```

## NOTE

The functions to determine the cache expiration date and whether or not to use the cache are currently not supported.

We plan to add these features in the future.
