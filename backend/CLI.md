# Clinker docs

Self-built CLI tool for migrating, seeding, and generating keys for Linker.  
Clinker is built upon [Cobra](https://github.com/spf13/cobra).

## Setting up Clinker for use

Because this isn't meant for public use, the setup is a bit cumbersome:
- First, make sure you've followed every step in installing Go from [this tutorial](https://go.dev/doc/install)
- Then, from the project's root directory, run:
```bash
go build backend/cmd/cli/clinker.go
mv clinker YOUR/PERSONAL/go/bin     # Replace with your actual go/bin path
```

## Using Clinker

The CLI doesn't process any flags except for `--help` (or `-h`), just works with commands like so:
```bash
clinker [COMMAND] [COMMAND]
# e.g. clinker key generate
```
Using `--help` gives you a list of commands you can use to follow up the last command.

## Developing Clinker

For convenience, here's a Bash script that's made for quicker Clinker developing.  
The script is meant to run only from the `linker/backend` directory but feel free to modify it:
```bash
devClinker()
{
	if [[ $PWD != *"linker/backend" ]]; then
		echo "Wrong directory"
		return 2
	fi
	rm -rf ~/go/bin/clinker
	go build ./cmd/cli/clinker.go
	mv clinker ~/go/bin/
}
```
After adding this to your shell's source file, you'll need to source it and then you can use it by running `devClinker` on the command line.
