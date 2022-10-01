# Git for Humans

An Intuitive [Git](https://git-scm.com/) Commands for Humans

## Introduction

[Git](https://git-scm.com/) is cool. It makes teamwork and collaborations easy.

But have you ever have to type unnecessarily long commands like so (only to start using Git in your project):

```shell
git init 
git add .
git commit -m "Initial commit"
```

Or when you need to face another commands for publishing changes to a remote repository:

```shell
git remote add origin https://github.com/git4humans/git4humans

git add .
git commit -m "Add publish functionality" 
git push origin main
```

Or simply want to save all your recent changes locally:

```shell
git add .
git commit -m "fix save message"
```

Using **Git for Humans**, the above scenarios can be done in a more succinct way (and let the tool guide you in a more friendly & interactive way):

```shell 
gt start 
```

```shell 
gt publish
```

```shell
gt save
```
## Installation 

Installing **Git for Humans** is easy, the only prerequisite is you already have [Git](https://git-scm.com/book/en/v2/Getting-Started-Installing-Git) installed in your system.

### Linux & macOS 

You can install on Linux & macOS by downloading the appropriate zip file from the [release page](https://github.com/git4humans/git4humans/releases), then unzip it into any location of your choice.

Then add the path of your **git4humans** directory to the **PATH** environment variable:

```shell
PATH=/path/to/git4humans:$PATH
```
You may verify by typing the following shell command:

```shell
gt
```

### Windows

The easiest way to install for Windows is download the installer from the [release page](https://github.com/git4humans/git4humans/releases) and run it. The file is small and installation process is fast. 

You may verify by running the following cmd command:

```shell
gt
```

Another is manual installation by downloading the appropriate zip file from the [release page](https://github.com/git4humans/git4humans/releases). 

Unzip the file into any location of your choice. Then add the path of your **git4humans** directory to the **PATH** environment variable. For example, if you unzip at the root drive C, then the path will be **C:\git4humans**.

For a guide to add path to environment variable, [read the following instructions](https://www.architectryan.com/2018/03/17/add-to-the-path-on-windows-10/).

### From Source

To install from source, you need to have [Go](https://go.dev/dl/) installed in your system.

Copy the source project:

```shell
git clone https://github.com/git4humans/git4humans
cd git4humans 
```
Build it:

Linux / macOS

```shell
go build -o gt cmd/gt/main.go
```

Windows 

```shell
go build -o gt.exe cmd/gt/main.go
```

Then add the **git4humans** directory to your **PATH** environment variable.

## Principles

The basic principle of **Git for Humans** is **conventions over configurations**. It means the tool will make an educated guess (based on conventions) for your simple everyday use. 

As an example, when dealing with remote repositories, **Git for Humans** will use `origin` as the default name for your remote repository. Unless you state it spesifically, the tool will assume you want to use `origin`. 

As a result, to add a remote repository, you can use simple `gt repo` command:

```shell 
gt repo https://github.com/pytorch/pytorch 
```

instead of the long ones:

```
git remote add origin https://github.com/pytorch/pytorch

gt repo origin https://github.com/pytorch/pytorch
gt repo + origin https://github.com/pytorch/pytorch
gt repo add origin https://github.com/pytorch/pytorch
```

Since we all know, unless you are working on a big, complex project, you most likely will use a single shared remote repository with the default name `origin`. 

No need to state it explicitly. The tool will understand.

Besides, for more complex use, you can always fallback to the [original Git commands](https://git-scm.com/docs). Every [standard Git commands](https://git-scm.com/docs) is applicable in **Git for Humans**.

### Interactivity & Guidance

Another principle is **interactivity and guidance**, which means the tool will guide you in an interactive way for some process that needs your confirmation (or an additional input, such as file name or commit message).

As an example, when you publish changes to a remote repository using `gt publish`, **Git for Humans** will check the status of your project. If you have some new or modified files that is not committed, the tool will ask for confirmation whether you want to save and include them with this publish command. It will also show the changes so you can easily decide.

```shell
On branch main
Your branch is up to date with 'origin/main'.

Changes not staged for commit:
  (use "gt + <file>..." to update what will be committed)
  (use "gt restore <file>..." to discard changes in working directory)
        modified:   README.md
        modified:   start.go

no changes added to commit (use "gt +" and/or "gt save")


Warn: you have some unstaged changes in this branch.

Want to save this changes? (y/n) _
```
You may type `y` to confirm that yes, you want to include all the changes with your publish. Then the tool will ask you to type a commit message.

```shell
Want to save this changes? (y/n) y

Save with message: _
```

This way, **Git for Humans** will try its best to guide you, and make your experience with Git more fun and enjoyable.

## Usage 

The infamous [Pareto Principle](https://en.wikipedia.org/wiki/Pareto_principle) can be found to apply to [Git commands](https://git-scm.com/docs): roughly 80% of your experience with Git will be circling around 20% of the commands. 

**Git for Humans** is focusing on these 20% commands that you will most likely touch in everyday use. For more complex use, you can always fallback to the [original Git commands](https://git-scm.com/docs). Every [standard Git commands](https://git-scm.com/docs) is applicable in **Git for Humans** (e.g., you can use `git bisect` and `gt bisect` interchangeably).

**new** 

**start**

**copy**

## Tutorial for Everyday Use

## Contribute 

You may contribute to improve **Git for Humans** in any way you choose. That includes [reporting issues](https://github.com/git4humans/git4humans/issues), [pull requests](https://github.com/git4humans/git4humans/pulls), [suggesting ideas](https://github.com/git4humans/git4humans/issues), and everything else.

To play with the source, you need to install [Go](https://go.dev/dl/) and build the project as follow.

Linux / macOS

```shell
go build -o gt cmd/gt/main.go
```

Windows 

```shell
go build -o gt.exe cmd/gt/main.go
```

We also include [default configuration](.goreleaser.yaml) to build the project with [Goreleaser](https://goreleaser.com/).

```shell
goreleaser release --snapshot --skip-publish --rm-dist
```

## License 

**Git for Humans** is open-sourced software licensed under the [MIT license](https://opensource.org/licenses/MIT).