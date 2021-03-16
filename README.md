# ADBATCH

The adbatch is a simple way to execute adb commands to all connected devices

## How to use

Execute the same command used in the adb but replacing by adbatch

Example
```
adbatch shell ls
```

It will execute the `shell ls` command to all connected devices, print the output and save the output for each device in a file called `[DEVICE_ID]-out.txt` in the current directory

## Compiling from source

Requirements

**Go Lang** with **go compiler**: [https://golang.org/](https://golang.org/)

First clone this repo and enter the directory

```bash
git clone https://github.com/brunohgv/adbatch
cd adbatch
```
Then use the GO CLI to install the program

```bash
go install
```

Done! the adbatch is installed

To test you can open the terminal and type
```bash
adbatch

# It shall display the following message:
# Adbatch is working!
# Run an adb command using adbatch to broadcast to all connected devices.
#
```

## Installing from Release

Go to [The releases page](https://github.com/brunohgv/adbatch/releases)

Download the compressed file containing your OS
* windows
* linux
* darwin (macOS)

Extract the content of the compressed file

Add the path to the compressed file to your **path variables**

To test you can open the terminal and type
```bash
adbatch

# It shall display the following message:
# Adbatch is working!
# Run an adb command using adbatch to broadcast to all connected devices.
#
```
