# ADBATCH

The adbatch is a simple way to execute adb commands to all connected devices

## How to use

Execute the same command used in the adb but replacing by adbatch

Example
```
adbatch shell ls
```

It will execute the `shell ls` command to all connected devices, print the output and save the output for each device in a file called `[DEVICE_ID]-out.txt` in the current directory