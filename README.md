# Git for Humans

An Intuitive [Git](https://git-scm.com/) Commands for Humans

## Introduction

[Git](https://git-scm.com/) is cool. It makes teamwork and collaborations easy.

But have you ever have to type unnecessarily long commands (only to start using Git in your project):

```shell
git init 
git add .
git commit -m "Initial commit"
```

Or when you face another commands to upload your changes to a remote repository:

```shell
git remote add origin https://github.com/git4humans/git4humans

git add .
git commit -m "Add publish functionality" 
git push origin main
```

Or simply want to save all recent changes locally:

```shell
git add .
git commit -m "fix save message"
```

Using **Git for Humans**, the above scenarios can be done in a more succinct way (and let the tool guide you in a more friendly & interactive way):

```shell 
g start 
```

```shell 
g up
```

```shell
g save
```
## Installation 

Installing **Git for Humans** is easy, the only prerequisite is you already have [Git](https://git-scm.com/book/en/v2/Getting-Started-Installing-Git) installed in your system.

### Linux & macOS 

You can install on Linux & macOS by downloading the appropriate zip file from the [release page](https://github.com/git4humans/git4humans/releases), then unzip it into any location of your choice.

Then add the path of **git4humans** directory from the zip file to the **PATH** environment variable:

```shell
PATH=/path/to/git4humans:$PATH
```
You may verify by typing the following shell command:

```shell
g
```

### Windows

The easiest way to install on Windows is download the installer from the [release page](https://github.com/git4humans/git4humans/releases) and run it. The file is small and installation process is fast. 

You may verify by running the following cmd command:

```shell
g
```

Another is manual installation by downloading the appropriate zip file from the [release page](https://github.com/git4humans/git4humans/releases). 

Unzip the file into any location of your choice. Then add the path of **git4humans** directory from the zip file to the **PATH** environment variable. For example, if you unzip at the root drive C, then the path will be **C:\git4humans**.

For a guide to add path to environment variable, [read the following instructions](https://www.architectryan.com/2018/03/17/add-to-the-path-on-windows-10/).

### From Source

To install from source, you need to have [Go installed](https://go.dev/dl/) in your system.

Copy the source project:

```shell
git clone https://github.com/git4humans/git4humans
cd git4humans 
```
Build it:

Linux / macOS

```shell
go build -o g cmd/g/main.go
```

Windows 

```shell
go build -o g.exe cmd/g/main.go
```

Then add the **git4humans** directory to your **PATH** environment variable.

## Principles

The basic principle of **Git for Humans** is **conventions over configurations**. It means the tool will make an educated guess (based on conventions) for our everyday use. 

As an example, when dealing with remote repositories, **Git for Humans** will use `origin` as the default name for our remote repository. Unless we state it spesifically, the tool will assume we want to use `origin`. 

As a result, to add a remote repository, we can use a simple `g repo` command:

```shell 
g repo https://github.com/pytorch/pytorch 
```

instead of the long ones:

```
git remote add origin https://github.com/pytorch/pytorch

g repo origin https://github.com/pytorch/pytorch
g repo + origin https://github.com/pytorch/pytorch
g repo add origin https://github.com/pytorch/pytorch
```

Since we all know, unless we are working on a big, complex project with lots of branches and collaborators, we will likely use a default remote repository with the default name `origin`. 

No need to state it explicitly.
The tool will understand.

Besides, for more heavy and complex use, we can always fallback to the original Git commands. [All Git commands](https://git-scm.com/docs) is applicable in **Git for Humans**.

### Interactivity & Guidance

Another principle is **interactivity and guidance**, which means the tool will guide us in an interactive way for some operations that need a confirmation (or an additional input, such as file name or commit message).

As an example, when uploading changes to a remote repository using `g up` command, **Git for Humans** will check the status of our project. If we have some new or modified files, the tool will ask for confirmation whether we want to save and upload them. It will also show the list of changes so we can easily decide.

```shell
On branch main

You have some unstaged changes in your project:
 M README.md
 M command.go
 M git.go

Save the changes (Y/N)? _
```
We may type `y` to confirm that yes, we want to save and upload the changes. Then the tool will ask to type a commit message.

```shell
Save the changes (Y/N)? y

Saving with message: _
```

This way, **Git for Humans** tries its best to guide us, and make our experience with Git more fun and enjoyable.

## Usage 

The infamous [Pareto Principle](https://en.wikipedia.org/wiki/Pareto_principle) can be found to apply to [Git commands](https://git-scm.com/docs): roughly 80% of our experience with Git will be around 20% of all the commands. 

**Git for Humans** is focusing on these 20% commands that we most likely will touch in everyday use. For more advanced and complex use, we can fallback to the original Git commands. [All Git commands](https://git-scm.com/docs) is applicable in **Git for Humans** (e.g., we can use `git fetch` and `g fetch` interchangeably).

**new** 

**start**

**copy**

**refresh**

**add**

**delete**

**move**

**rename**

**restore**

**save** 

**repo** 

**sync**

**up**

**branch** 

**switch**

**merge**

**reset**

**tag** 

**status**

**log**

**last**

## Tutorial for Everyday Use

## Contribute 

You may contribute to improve **Git for Humans** in any way you choose. That includes [reporting issues](https://github.com/git4humans/git4humans/issues), [pull requests](https://github.com/git4humans/git4humans/pulls), [suggesting ideas](https://github.com/git4humans/git4humans/issues), and everything else.

To play with the source, you need to [install Go](https://go.dev/dl/) and build the project as follow.

Linux / macOS

```shell
go build -o g cmd/g/main.go
```

Windows 

```shell
go build -o g.exe cmd/g/main.go
```

We also include [default configuration](.goreleaser.yaml) to build the project with [Goreleaser](https://goreleaser.com/).

```shell
goreleaser release --snapshot --skip-publish --rm-dist
```

## License 

**Git for Humans** is open-sourced software licensed under the [MIT license](https://opensource.org/licenses/MIT).