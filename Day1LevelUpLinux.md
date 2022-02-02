<!--linux skill up challenge-->

#### Day 1 - Linux Skill Up Challenge

**Aim** : Get to know your server

PS: This was posted a month ago on Reddit. Not sure how this works but my goal is just to learn so I took it up .

**Tools Used** :
- Oracle Virtual Machine - CentOS Distribution
- Windows Subsystem for Linux - Ubuntu Distribution

**Procedure :**
1) Configure and run server on CentOS VM</br>
   Login to CentOS.</br>
   Commands Used :</br>
   i)sudo yum install openssh</br>
    used to install openssh on your VM (here we need to install the openssh server)</br>
   ii)sudo systemctl status sshd</br>
     used to check is ssh is running</br>
   iii)sudo systemctl start sshd</br>
     used to start the ssh service</br>
   iv)sudo systemctl stop sshd</br>
     used to stop the ssh service</br>
   v)ip a</br>
     used to get IP address of the VM</br>

2) Connect to VM server via Ubuntu</br>
   Open Ubuntu terminal.</br>
   Commands Used :</br>
   i)sudo apt-get install openssh-client</br>
     used to install openssh client</br>
   ii)ssh user_Id_of_VM@ipAddress_of_VM</br>
   
   Operations performed on Remote Server</br>
   i)free -h</br>
   ii)uname -a</br>
   iii)ls</br>
   iv)uptime</br>
   v)df -h</br>

  **Optional**:-</br>
   Additional commands incase we want to configure a server on Ubuntu terminal</br>
   i)sudo apt-get install openssh-server</br>
     used to install openssh server</br>
   ii)sudo service ssh status</br>
     used to check is ssh if running</br>
   iii)sudo service ssh start</br>
     used to start the ssh service</br>
   iv)sudo service ssh stop</br>
     used to stop the ssh service</br>

**Issues faced**:</br>
1) ssh command returned Resource temporarily unavailable</br>
Analysis :-</br>
CentOs VM was running on Host Only Network mode</br>
Solution :-</br>
Change network settings to NAT mode</br>

**Resources**:-</br>
1)https://www.reddit.com/r/linuxupskillchallenge/comments/rugjxr/day_1_get_to_know_your_server/</br>
2)https://www.cyberciti.biz/faq/how-to-install-ssh-on-ubuntu-linux-using-apt-get/</br>
3)https://linuxhint.com/enable_ssh_centos8/</br>
