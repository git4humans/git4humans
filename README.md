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
git add .
git commit -m "Add publish functionality"
git remote add origin https://github.com/git4humans/git4humans 
git push origin main
```

Using **Git for Humans**, the above two scenarios can be done in a more succinct way (and let the tool guide you in a friendly way):

```shell 
gt start 
```

```shell 
gt publish
```
## Installation 

Installation is easy in any operating system of your choice. The only requirement is you have already installed [Git](https://git-scm.com/book/en/v2/Getting-Started-Installing-Git) in your system.

### Linux & macOS 

You can install on Linux & macOS by downloading appropriate zip file from the [release page](https://github.com/git4humans/git4humans/releases). Then unzip it at any location of your choice.

Then add the path of your **git4humans** directory to the **PATH** environment variable:

```shell
PATH=/path/to/your/git4humans:$PATH
```
You can check your installation by typing the following shell command:

```shell
gt
```

### Windows

The easiest way to install for Windows is download the installer from the [release page](https://github.com/git4humans/git4humans/releases) and run it. The file is small and installation process is fast. 

You can check your installation by running the following command in cmd:

```shell
gt
```

Another is a manual install by downloading appropriate zip file from the [release page](https://github.com/git4humans/git4humans/releases). Then unzip it at any location of your choice.

For example, you may unzip it at the root of drive C.

Then add the path of your **git4humans** directory to the **PATH** environment variable. If you unzip it at the root drive C, the path will be **C:\git4humans**. For a guide to add to the **PATH** environment variable, you can [read the following instructions](https://www.architectryan.com/2018/03/17/add-to-the-path-on-windows-10/).

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

The basic principle of **Git for Humans** is **conventions over configurations**. It means the tool will make an educated guess based on conventions for your simple everyday use. 

As an example, when dealing with remote repositories, **Git for Humans** will use `origin` as the default name for your remote repository. Unless you state it spesifically, the tool will assume you want to use `origin`. 

As a result, to add a remote repository, for example, you can use simple `gt repo` command:

```shell 
gt repo https://github.com/pytorch/pytorch 
```
or 

```shell 
gt repo + https://github.com/pytorch/pytorch 
```

instead of the long command:

```
git remote add origin https://github.com/pytorch/pytorch

gt repo origin https://github.com/pytorch/pytorch
gt repo + origin https://github.com/pytorch/pytorch
```

Since we all know, unless you are working in a big, complex project, you mostly only need to add a default remote repository under the default name `origin`. 

No need to explicitly state it. The tool will understand.

For an advanced and more complex use, you can always fallback to the [original Git commands](https://git-scm.com/docs). Every [standard Git commands](https://git-scm.com/docs) is applicable in **Git for Humans**.

Another principle is **interactivity and guidance**.

## Usage 

**gt new** 

**gt start**

**gt copy**

## Tutorial for Everyday Use

## Contribute 

## License 

**Git for Humans** is open-sourced software licensed under the [MIT license](https://opensource.org/licenses/MIT).