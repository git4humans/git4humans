# Git for Humans

An Intuitive [Git](https://git-scm.com/) Commands for Humans

## Introduction

[Git](https://git-scm.com/) is cool. It makes teamwork and collaborations easy.

But have you ever have to type unnecessary long commands like the following (only to start using Git in your project):

```shell
git init 
git add .
git commit -m "Initial commit"
```

Or when you need to face another long commands for publishing changes to a remote repository:

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

Using **Git for Humans**, the above scenarios can be done in a more succinct way (and the tool will guide you in an interactive friendly way):

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

Installation is easy. The only requirement is you have already installed [Git](https://git-scm.com/book/en/v2/Getting-Started-Installing-Git) in your system.

### Linux & macOS 

You can install on Linux & macOS by downloading appropriate zip file from the [release page](https://github.com/git4humans/git4humans/releases). Then unzip it into any location of your choice.

Then add the path of your **git4humans** directory to the **PATH** environment variable:

```shell
PATH=/path/to/your/git4humans:$PATH
```
You can check by typing the following shell command:

```shell
gt
```

### Windows

The easiest way to install for Windows is download the installer from the [release page](https://github.com/git4humans/git4humans/releases) and run it. The file is small and installation process is fast. 

You can check by running the following command in cmd:

```shell
gt
```

Another is manual installation by downloading appropriate zip file from the [release page](https://github.com/git4humans/git4humans/releases). 

Unzip the file into any location of your choice. Then add the path of your **git4humans** directory to the **PATH** environment variable. For example, if you unzip at the root drive C, then the path will be **C:\git4humans**.

For a guide to add path to the **PATH** environment variable, [read the following instructions](https://www.architectryan.com/2018/03/17/add-to-the-path-on-windows-10/).

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

### Conventions
The basic principle of **Git for Humans** is **conventions over configurations**. It means the tool will make an educated guess (based on conventions) for your simple everyday use. 

As an example, when dealing with remote repositories, **Git for Humans** will use `origin` as the default name for your remote repository. Unless you state it spesifically, the tool will assume you want to use `origin`. 

As a result, to add a remote repository, you can use simple `gt repo` command:

```shell 
gt repo https://github.com/pytorch/pytorch 
```

instead of the long command:

```
git remote add origin https://github.com/pytorch/pytorch

gt repo origin https://github.com/pytorch/pytorch
gt repo + origin https://github.com/pytorch/pytorch
gt repo add origin https://github.com/pytorch/pytorch
```

Since we all know, unless you are working on a big, complex project, you mostly will add your remote repository under the default name `origin`. 

You don't need to state it explicitly. The tool will understand.

Besides, for more complex and advanced use, you can always fallback to the [original Git commands](https://git-scm.com/docs). Every [standard Git commands](https://git-scm.com/docs) is applicable in **Git for Humans**.

### Interactivity & Guidance

Another principle is **interactivity and guidance**. It means the tool will guide you in an interactive way for some process that needs confirmation (or additional input, such as file name or commit message).

As an example, when you publish changes to a remote repository using `gt publish`, **Git for Humans** will check the status of your project. 

If you have some new or modified files that has not been committed, the tool will ask for confirmation whether you want to save and include the changes in this publish command. It will also show details of the changes so you can easily decide.

```shell
On branch main
Your branch is up to date with 'origin/main'.

Changes not staged for commit:
  (use "gt + <file>..." to update what will be committed)
  (use "gt restore <file>..." to discard changes in working directory)
        modified:   README.md
        modified:   start.go

no changes added to commit (use "gt +" and/or "gt save")

You have unstaged changes in this branch.

Do you want to save all the changes? (y/n) _
```
You may type `y` to confirm that yes, you want to include all the changes with your publish. Then the tool will ask you to input the commit message.

```shell
Do you want to save all the changes? (y/n) y

Save with message: _
```

This way, **Git for Humans** will try its best to guide, to make your experience with Git more fun and enjoyable.

## Usage 

The infamous [80/20 Pareto Principle](https://en.wikipedia.org/wiki/Pareto_principle) can be found to apply with Git commands. Roughly 80% of your experience with Git will be circling around 20% of the [complete commands](https://git-scm.com/docs).

**Git for Humans** is focusing into these 20% commands that you will likely touch in everyday use. For more complex and advanced use, you can always fallback to the [original Git commands](https://git-scm.com/docs), since every [standard Git commands](https://git-scm.com/docs) is applicable in **Git for Humans** (e.g., you can use `git bisect` and `gt bisect` interchangeably).

**new** 

**start**

**copy**

## Tutorial for Everyday Use

## Contribute 

## License 

**Git for Humans** is open-sourced software licensed under the [MIT license](https://opensource.org/licenses/MIT).