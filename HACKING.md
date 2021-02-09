# Hacking on `terraform-provider-oneview` 

We welcome your contributions to the OneView provider for Terraform! See CONTRIBUTING.md for more details.

## Go get

This project is `go get`able! In order to use the latest code in the master branch and also build the
binary to use with Hashicorp's Terraform, just issue the following command:

```bash
go get -u github.com/HewlettPackard/terraform-provider-oneview
```

Where `-u` will also pull down the latest version of the code rather than use the version (if any) you
have stored already on your `$GOPATH`.

The command above will download the dependency to your `$GOPATH` and also build the binary and place it into `$GOPATH/bin`
so you can later copy it to a different location, along with the terraform binary
[you can get from here](https://www.terraform.io/downloads.html).

## Go dependency management

We use [govendor](https://github.com/kardianos/govendor) as a dependency management tool, because i
of how clean the result is. When adding new external dependencies to `terraform-provider-oneview`, you
must also add those dependencies into the repository so they can be version-controlled. To do so,
first pull the dependency into your own `$GOPATH` -- since `terraform-provider-oneview` expects a
proper Go installation with a `$GOPATH` set, which it's set by default since Go 1.8 to be at
`$HOME/go` -- and then you can run:

```bash
govendor add +external
```

Which will find for extra differences in the dependencies and will add the new ones that weren't before
inside the `vendor/` folder.

#### Pinning an specific version of a dependency or update existent dependencies

You can update existing dependencies or add new ones and also pin them to an specific semver or
git SHA by running:

```bash
# Add new dependencies by fetching them with an specific git SHA
govendor fetch github.com/HewlettPackard/dependency@a4bbce9fcae005b22ae5443f6af064d80a6f5a55

# Add new dependencies by using the latest code in the git tag or branch "v1.*.*"
govendor fetch github.com/HewlettPackard/dependency@v1

# Add new dependencies by using an exact git branch tag
govendor fetch github.com/HewlettPackard/dependency@=v1
```

For more information on how to use `govendor`, please refer to
[their documentation](https://github.com/kardianos/govendor).
