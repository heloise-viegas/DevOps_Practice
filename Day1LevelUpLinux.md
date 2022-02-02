######linux skill up challenge

#### Day 1 - Linux Skill Up Challenge

**Aim** : Get to know your server

PS: This was posted a month ago on Reddit. Not sure how this works but my goal is just to learn so I took it up .

**Tools Used : **
- Oracle Virtual Machine - CentOS Distribution
- Windows Subsystem for Linux - Ubuntu Distribution

**Procedure :**
1) Configure and run server on CentOS VM
   Login to CentOS.
   Commands Used :
   i)sudo yum install openssh
   used to install openssh on your VM (here we need to install the openssh server)
   ii)sudo systemctl status sshd
     used to check is ssh is running
   iii)sudo systemctl start sshd
     used to start the ssh service
   iv)sudo systemctl stop sshd
     used to stop the ssh service
   v)ip a
     used to get IP address of the VM

2) Connect to VM server via Ubuntu
   Open Ubuntu terminal.
   Commands Used :
   i)sudo apt-get install openssh-client
     used to install openssh client
   ii)ssh user_Id_of_VM@ipAddress_of_VM
   
   Operations performed on Remote Server
   i)free -h
   ii)uname -a
   iii)ls
   iv)uptime
   v)df -h

  **Optional :-** 
   Additional commands incase we want to configure a server on Ubuntu terminal  
   i)sudo apt-get install openssh-server
     used to install openssh server
   ii)sudo service ssh status  
     used to check is ssh if running
   iii)sudo service ssh start
     used to start the ssh service
   iv)sudo service ssh stop
     used to stop the ssh service

**Issues faced :**
1) ssh command returned Resource temporarily unavailable
Analysis :-
CentOs VM was running on Host Only Network mode
Solution :-
Change network settings to NAT mode

**Resources :-** 
1)https://www.reddit.com/r/linuxupskillchallenge/comments/rugjxr/day_1_get_to_know_your_server/
2)https://www.cyberciti.biz/faq/how-to-install-ssh-on-ubuntu-linux-using-apt-get/
3)https://linuxhint.com/enable_ssh_centos8/
