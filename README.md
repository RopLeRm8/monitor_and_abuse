## Monitoring activity hours per process system

In order to make this app work in background, make sure to schedule it accordingly.

Instructions (windows only):

You would have to enter Task Scheduler in order to schedule the script to run once user has logged in

ALT+R</br>
Enter inside:

```
taskschd.msc
```

Then, create basic task, make sure to make it trigger once user is logon. Make sure it triggers the monitor.exe as an action.

Thats it! You're all set.
