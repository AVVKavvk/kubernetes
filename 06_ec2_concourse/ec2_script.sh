#!/bin/bash

echo "This is on ec2 server and i am running this script"
echo "Hello World"


# Define some colors for fun
GREEN='\033[0;32m'
NC='\033[0m' # No Color

echo -e "${GREEN}--- EC2 Status Report ---${NC}"

# 1. Print a greeting with the current user and date
echo "Hello, $USER! Today is $(date)."

# 2. Show the system uptime
echo -e "\n[System Uptime]"
uptime -p

# 3. Show a quick loop 
echo -e "\n[Running a quick diagnostic loop...]"
for i in {1..3}
do
   echo "Checking status of service $i... All good!"
   sleep 0.5
done

# 4. Check Disk Usage
echo -e "\n[Disk Usage]"
df -h | grep '^/dev/'

echo -e "\n${GREEN}Script execution finished.${NC}"