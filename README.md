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

Using **Git for Humans**, the above two scenarios can be done in a more succinct way (and let the tool guide you in a friendly interactive way):

```shell 
gt start 
```

```shell 
gt publish
```
## Installation 

Installation is easy. The only requirement is you have already installed [Git](https://git-scm.com/book/en/v2/Getting-Started-Installing-Git) in your system.

### Linux & macOS 

You can install on Linux & macOS by downloading appropriate zip file from the [release page](https://github.com/git4humans/git4humans/releases). Then unzip it into any location of your choice.

Then add the path of your **git4humans** directory to the **PATH** environment variable:

```shell
PATH=/path/to/your/git4humans:$PATH
```
You can check your installation by typing the following shell command:

```shell
gt
```

### Windows

The easiest way to install for Windows is download the installer from the [release page](https://github.com/git4humans/git4humans/releases) and run it. The file size is small and installation process is fast. 

You can check your installation by running the following command in cmd:

```shell
gt
```

Another is a manual install by downloading appropriate zip file from the [release page](https://github.com/git4humans/git4humans/releases). 

Then unzip the file into any location of your choice. Then add the path of your **git4humans** directory to the **PATH** environment variable. For example, you may unzip at the root of drive C, then the path will be **C:\git4humans**. 

For a guide to add to the **PATH** environment variable, [read the following instructions](https://www.architectryan.com/2018/03/17/add-to-the-path-on-windows-10/).

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

Then copy the output (gt or gt.exe) into any directory of your choice. Then add that directory to your **PATH** environment variable.

## Principles

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

Since we all know, unless you are working on a big, complex project, you mostly need to add a default remote repository under the default name `origin`. 

No need to state it explicitly. The tool will understand.

For an advanced and more complex use, you can always fallback to the [original Git commands](https://git-scm.com/docs). Every [standard Git commands](https://git-scm.com/docs) is applicable in **Git for Humans**.

Another principle is **interactivity and guidance**. It means the tool will guide you in an interactive way for some process that needs confirmation (or an additional input, such as filename or commit message).

As an example, when you publish changes to a remote repository using `gt publish`, **Git for Humans** will check the status of your project. 

If you have some untracked or modified files that has not been committed, it will ask for confirmation if you want to save and include it with your current publish command. The tool will also show details of the changes so you can easily decide.

```shell
On branch main
Your branch is up to date with 'origin/main'.

Changes not staged for commit:
  (use "gt + <file>..." to update what will be committed)
  (use "gt restore <file>..." to discard changes in working directory)
        modified:   README.md
        modified:   start.go

no changes added to commit (use "gt +" and/or "gt save")

You have some recent changes in this branch.

Do you want to save all the changes? (y/n) _
```
You may type `y` to confirm that yes, you want to include all the changes with your current publish. Then the tool will ask you to input the commit message.

```shell
You have some recent changes in this branch.

Do you want to save all the changes? (y/n) y

Save with message: _
```

In this way, **Git for Humans** will try its best to guide you to make your Git experience more fun and enjoyable.

## Usage 

**gt new** 

**gt start**

**gt copy**

## Tutorial for Everyday Use

## Contribute 

## License 

**Git for Humans** is open-sourced software licensed under the [MIT license](https://opensource.org/licenses/MIT).