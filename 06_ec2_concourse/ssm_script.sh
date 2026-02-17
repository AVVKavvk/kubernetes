#!/bin/bash

echo "This is on ec2 server and i am running this script"
echo "Hello World"

echo "--- EC2 Status Report ---"

# 1. Print a greeting with the current user and date
echo "Hello, $USER! Today is $(date)."

# 2. Show the system uptime
echo ""
echo "[System Uptime]"
uptime -p

# 3. Show a quick loop
echo ""
echo "[Running a quick diagnostic loop...]"
for i in {1..3}
do
   echo "Checking status of service $i... All good!"
   sleep 0.5
done

# 4. Check Disk Usage
echo ""
echo "[Disk Usage]"
df -h | grep '^/dev/'

echo ""
echo "Script execution finished."