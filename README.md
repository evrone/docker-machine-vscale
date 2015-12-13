# Docker Machine Vscale Driver

[![Vexor status](https://ci.vexor.io/projects/2c3724f0-4de6-4e28-9bdd-19fb71795812/status.svg)](https://ci.vexor.io/ui/projects/2c3724f0-4de6-4e28-9bdd-19fb71795812/builds)

This is a plugin for [Docker Machine](https://docs.docker.com/machine/) allowing
to create Doker hosts on [Vscale]( http://vscale.io ) cloud services.

## Installation

Compile driver for your platform

```console
$ make install
```

## Usage

After compile you can use driver for creating docker hosts.
Get Vscale access token from [your profile](https://vscale.io/panel/settings/tokens/) then run:

```console
$ docker-machine create -d vscale --vscale-access-token YOUR_VSCALE_ACCESS_TOKEN machine_name
```

You should see simple log of operations:

```
Running pre-create checks...
Creating machine...
(vscale) Creating SSH key...
(vscale) Creating Vscale scalet...
(vscale) Waiting for IP address to be assigned to the Scalet...
Waiting for machine to be running, this may take a few minutes...
Machine is running, waiting for SSH to be available...
Detecting operating system of created instance...
Detecting the provisioner...
Provisioning with ubuntu(upstart)...
Installing Docker...
Copying certs to the local machine directory...
Copying certs to the remote machine...
Setting Docker configuration on the remote daemon...
Checking connection to Docker...
Docker is up and running!
To see how to connect Docker to this machine, run: docker-machine env vscale
```

Just insert in command line this:
```console
$ docker-machine env vscale
```
and follow instructions.

If you did everything correctly, run the command, you will see information about the host:

```console
$ docker info
```

## Contribution Guidelines

01. Fork
02. Change
03. PR
