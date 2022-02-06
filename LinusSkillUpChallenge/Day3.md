<!--linux skill up challenge-->

#### Day 3 - Linux Skill Up Challenge

**Aim** : Power Trip

PS: This was posted a month ago on Reddit. Not sure how this works but my goal is just to learn so I took it up .

**Tools Used** :
- Oracle Virtual Machine - CentOS Distribution
- Windows Subsystem for Linux - Ubuntu Distribution

**Procedure :**
1) Connect to VM server via Ubuntu</br>
   Open Ubuntu terminal.</br>
   Commands Used :</br>
   i)ssh james@192.145.87.21</br>
   
Operations performed on Remote Server</br>
**-shadow file stores hashed passwords**<br/>
[osboxes@osboxes ~]$ sudo cat /etc/shadow<br/>

**reboot server**<br/>
[osboxes@osboxes ~]$ sudo reboot<br/>
Connection to 192.168.85.56 closed by remote host.<br/>
Connection to 192.168.85.56 closed.<br/>
heloise@DESKTOP-P5D1DV6:~$ ssh osboxes@192.168.85.56<br/>
ssh: connect to host 192.168.85.56 port 22: Resource temporarily unavailable<br/>
heloise@DESKTOP-P5D1DV6:~$ ssh osboxes@192.168.85.56<br/>
osboxes@192.168.85.56's password:<br/>
Activate the web console with: systemctl enable --now cockpit.socket<br/>
Last login: Sat Feb  5 01:24:00 2022<br/>
[osboxes@osboxes ~]$ uptime<br/>
 01:25:12 up 1 min,  2 users,  load average: 0.14, 0.10, 0.04<br/>
 
 **Change server host name**<br/>
[osboxes@osboxes ~]$ sudo hostnamectl set-hostname osbox<br/>
[sudo] password for osboxes:<br/>
[osboxes@osboxes ~]$ hostname<br/>
osbox<br/>
[osboxes@osboxes ~]$ sudo hostnamectl set-hostname osboxes<br/>
[sudo] password for osboxes:<br/>
[osboxes@osboxes ~]$ hostname<br/>
osboxes<br/>

**Change server date and time zone**<br/>
[osboxes@osboxes ~]$ timedatectl<br/>
               Local time: Sat 2022-02-05 02:26:36 EST<br/>
           Universal time: Sat 2022-02-05 07:26:36 UTC<br/>
                 RTC time: Sat 2022-02-05 07:26:34<br/>
                Time zone: America/New_York (EST, -0500)<br/>
System clock synchronized: yes<br/>
              NTP service: active<br/>
          RTC in local TZ: no<br/>
[osboxes@osboxes ~]$ timedatectl list-timezones<br/>
Africa/Abidjan<br/>
Africa/Accra<br/>
Africa/Addis_Ababa<br/>
Africa/Algiers<br/>
Africa/Asmara<br/>
Africa/Bamako<br/>
Africa/Bangui<br/>
Africa/Banjul<br/>
Africa/Bissau<br/>

[osboxes@osboxes ~]$ sudo timedatectl set-timezone Australia/Sydney<br/>
[sudo] password for osboxes:<br/>
[osboxes@osboxes ~]$ timedatectl<br/>
               Local time: Sat 2022-02-05 18:29:23 AEDT<br/>
           Universal time: Sat 2022-02-05 07:29:23 UTC<br/>
                 RTC time: Sat 2022-02-05 07:29:22<br/>
                Time zone: Australia/Sydney (AEDT, +1100)<br/>
System clock synchronized: yes<br/>
              NTP service: active<br/>
          RTC in local TZ: no<br/>
[osboxes@osboxes ~]$ sudo timedatectl set-timezone America/New_York<br/>
[osboxes@osboxes ~]$ timedatectl<br/>
               Local time: Sat 2022-02-05 02:30:04 EST<br/>
           Universal time: Sat 2022-02-05 07:30:04 UTC<br/>
                 RTC time: Sat 2022-02-05 07:30:02<br/>
                Time zone: America/New_York (EST, -0500)<br/>
System clock synchronized: yes<br/>
              NTP service: active<br/>
          RTC in local TZ: no<br/>
[osboxes@osboxes ~]$<br/>


**Resources**:-</br>
1)https://www.reddit.com/r/linuxupskillchallenge/comments/rw1uw0/day_3_power_trip/</br>
