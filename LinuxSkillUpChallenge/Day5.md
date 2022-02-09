
**[osboxes@osboxes ~]$ sudo dnf install apache2**
[sudo] password for osboxes:
Last metadata expiration check: 7:47:42 ago on Tue Feb  8 01:51:14 2022.
No match for argument: apache2
Error: Unable to find a match: apache2
**[osboxes@osboxes ~]$ sudo dnf search apache**
Last metadata expiration check: 7:49:28 ago on Tue Feb  8 01:51:14 2022.
========================= Name & Summary Matched: apache ==========================
apache-commons-logging.noarch : Apache Commons Logging
pcp-pmda-apache.x86_64 : Performance Co-Pilot (PCP) metrics for the Apache
                       : webserver
============================== Name Matched: apache ===============================
apache-commons-cli.noarch : Command Line Interface Library for Java
apache-commons-codec.noarch : Implementations of common encoders and decoders
apache-commons-io.noarch : Utilities to assist with developing IO functionality
apache-commons-lang3.noarch : Provides a host of helper utilities for the java.lang
                            : API
============================= Summary Matched: apache =============================
apr.i686 : Apache Portable Runtime library
apr.x86_64 : Apache Portable Runtime library
apr-util.i686 : Apache Portable Runtime Utility library
apr-util.x86_64 : Apache Portable Runtime Utility library
httpd.x86_64 : Apache HTTP Server
**[osboxes@osboxes ~]$ sudo dnf install apache http server**
Last metadata expiration check: 7:50:30 ago on Tue Feb  8 01:51:14 2022.
No match for argument: apache
No match for argument: http
No match for argument: server
Error: Unable to find a match: apache http server
**[osboxes@osboxes ~]$ sudo dnf install httpd**
Last metadata expiration check: 7:50:50 ago on Tue Feb  8 01:51:14 2022.
Dependencies resolved.
Installed:
  apr-1.6.3-12.el8.x86_64
  apr-util-1.6.1-6.el8.x86_64
  apr-util-bdb-1.6.1-6.el8.x86_64
  apr-util-openssl-1.6.1-6.el8.x86_64
  centos-logos-httpd-85.8-2.el8.noarch
  httpd-2.4.37-43.module_el8.5.0+1022+b541f3b1.x86_64
  httpd-filesystem-2.4.37-43.module_el8.5.0+1022+b541f3b1.noarch
  httpd-tools-2.4.37-43.module_el8.5.0+1022+b541f3b1.x86_64
  mod_http2-1.15.7-3.module_el8.4.0+778+c970deab.x86_64

Complete!
**[osboxes@osboxes ~]$ curl https://ipinfo.io/ip**
**43.249.184.35[osboxes@osboxes ~]$ http://43.249.184.35**
-bash: http://43.249.184.35: No such file or directory
[osboxes@osboxes ~]$ systemctl status httpd
● httpd.service - The Apache HTTP Server
   Loaded: loaded (/usr/lib/systemd/system/httpd.service; disabled; vendor preset:>
   Active: inactive (dead)
     Docs: man:httpd.service(8)
lines 1-4/4 (END)
^C
**[osboxes@osboxes ~]$ sudo systemctl start httpd**
[sudo] password for osboxes:
**[osboxes@osboxes ~]$ systemctl status httpd**
● httpd.service - The Apache HTTP Server
   Loaded: loaded (/usr/lib/systemd/system/httpd.service; disabled; vendor preset:>
   Active: active (running) since Tue 2022-02-08 09:53:04 EST; 3s ago
     Docs: man:httpd.service(8)
 Main PID: 82971 (httpd)
   Status: "Started, listening on: port 80"
    Tasks: 213 (limit: 4944)
   Memory: 25.4M
   CGroup: /system.slice/httpd.service
           ├─82971 /usr/sbin/httpd -DFOREGROUND
           ├─82972 /usr/sbin/httpd -DFOREGROUND
           ├─82973 /usr/sbin/httpd -DFOREGROUND
           ├─82974 /usr/sbin/httpd -DFOREGROUND
           └─82975 /usr/sbin/httpd -DFOREGROUND

Feb 08 09:53:04 osboxes systemd[1]: Starting The Apache HTTP Server...
Feb 08 09:53:04 osboxes httpd[82971]: AH00558: httpd: Could not reliably determine>
Feb 08 09:53:04 osboxes systemd[1]: Started The Apache HTTP Server.
Feb 08 09:53:04 osboxes httpd[82971]: Server configured, listening on: port 80
lines 1-19/19 (END)
^C
[osboxes@osboxes ~]$ http://43.249.184.35
-bash: http://43.249.184.35: No such file or directory
[osboxes@osboxes ~]$ http://192.168.31.255
-bash: http://192.168.31.255: No such file or directory
[osboxes@osboxes ~]$ http://192.168.31.27
-bash: http://192.168.31.27: No such file or directory
[osboxes@osboxes ~]$ sudo cat /etc/httpd.conf
[sudo] password for osboxes:
cat: /etc/httpd.conf: No such file or directory
[osboxes@osboxes ~]$ sudo -i
[root@osboxes ~]# cd /etc
[root@osboxes etc]# ls -al
total 1192
drwxr-xr-x. 105 root root     8192 Feb  8 09:42 .
dr-xr-xr-x.  17 root root      224 Feb  8 03:40 ..
etc...

[root@osboxes etc]# cat httpd
cat: httpd: Is a directory
[root@osboxes etc]# cd httpd
[root@osboxes httpd]# ls -al
total 12
drwxr-xr-x.   5 root root  105 Feb  8 09:42 .
drwxr-xr-x. 105 root root 8192 Feb  8 09:42 ..

[root@osboxes httpd]# cat conf.d
cat: conf.d: Is a directory
[root@osboxes httpd]# cd conf.d
[root@osboxes conf.d]# ls -al
total 16
drwxr-xr-x. 2 root root   82 Feb  8 09:42 .
etc..
[root@osboxes conf.d]# pwd
/etc/httpd/conf.d
[root@osboxes conf.d]# cat welcome.conf
#
# This configuration file enables the default "Welcome" page if there
# is no default index page present for the root URL.  To disable the
# Welcome page, comment out all the lines below.
#
# NOTE: if this file is removed, it will be restored on upgrades.
#
<LocationMatch "^/+$">
    Options -Indexes
    ErrorDocument 403 /.noindex.html
</LocationMatch>

<Directory /usr/share/httpd/noindex>
    AllowOverride None
    Require all granted
</Directory>

Alias /.noindex.html /usr/share/httpd/noindex/index.html
Alias /poweredby.png /usr/share/httpd/icons/apache_pb3.png[root@osboxes conf.d]#
[root@osboxes conf.d]# cat userdir.conf
#
# UserDir: The name of the directory that is appended onto a user's home
# directory if a ~user request is received.
#
# The path to the end user account 'public_html' directory must be
# accessible to the webserver userid.  This usually means that ~userid
# must have permissions of 711, ~userid/public_html must have permissions
# of 755, and documents contained therein must be world-readable.
# Otherwise, the client will only receive a "403 Forbidden" message.
#
<IfModule mod_userdir.c>
    #
    # UserDir is disabled by default since it can confirm the presence
    # of a username on the system (depending on home directory
    # permissions).
    #
    UserDir disabled

    #
    # To enable requests to /~user/ to serve the user's public_html
    # directory, remove the "UserDir disabled" line above, and uncomment
    # the following line instead:
    #
    #UserDir public_html
</IfModule>

#
# Control access to UserDir directories.  The following is an example
# for a site where these directories are restricted to read-only.
#
<Directory "/home/*/public_html">
    AllowOverride FileInfo AuthConfig Limit Indexes
    Options MultiViews Indexes SymLinksIfOwnerMatch IncludesNoExec
    Require method GET POST OPTIONS
</Directory>

**[root@osboxes conf.d]# cd /var/log/httpd**
**[root@osboxes httpd]# ls -al**
total 8
drwx------.  2 root root   41 Feb  8 09:53 .
drwxr-xr-x. 10 root root 4096 Feb  8 09:42 ..
-rw-r--r--.  1 root root    0 Feb  8 09:53 access_log
-rw-r--r--.  1 root root 1039 Feb  8 09:53 error_log
[root@osboxes httpd]# cat access_log
[root@osboxes httpd]# vim access_log
[root@osboxes httpd]# file access_log
access_log: empty


> Complete!
> [osboxes@osboxes ~]$
> [osboxes@osboxes ~]$ sudo dnf install apache2
> [sudo] password for osboxes:
> Last metadata expiration check: 7:47:42 ago on Tue Feb  8 01:51:14 2022.
> No match for argument: apache2
> Error: Unable to find a match: apache2
> [osboxes@osboxes ~]$ sudo dnf search apache
> Last metadata expiration check: 7:49:28 ago on Tue Feb  8 01:51:14 2022.
> ========================= Name & Summary Matched: apache ==========================
> apache-commons-logging.noarch : Apache Commons Logging
> pcp-pmda-apache.x86_64 : Performance Co-Pilot (PCP) metrics for the Apache
>                        : webserver
>
> Installed:
>   apr-1.6.3-12.el8.x86_64
>   apr-util-1.6.1-6.el8.x86_64
>   apr-util-bdb-1.6.1-6.el8.x86_64
>   apr-util-openssl-1.6.1-6.el8.x86_64
>   centos-logos-httpd-85.8-2.el8.noarch
>   httpd-2.4.37-43.module_el8.5.0+1022+b541f3b1.x86_64
>   httpd-filesystem-2.4.37-43.module_el8.5.0+1022+b541f3b1.noarch
>   httpd-tools-2.4.37-43.module_el8.5.0+1022+b541f3b1.x86_64
>   mod_http2-1.15.7-3.module_el8.4.0+778+c970deab.x86_64
>
> Complete!
> [osboxes@osboxes ~]$ curl https://ipinfo.io/ip
> 43.249.184.35[osboxes@osboxes ~]$ http://43.249.184.35
w> -bash: http://43.249.184.35: No such file or directory
> [osboxes@osboxes ~]$ systemctl status httpd
> ● httpd.service - The Apache HTTP Server
>    Loaded: loaded (/usr/lib/systemd/system/httpd.service; disabled; vendor preset:>
>    Active: inactive (dead)
>      Docs: man:httpd.service(8)
> lines 1-4/4 (END)
> ^C
> [osboxes@osboxes ~]$ sudo systemctl start httpd
> [sudo] password for osboxes:
> [osboxes@osboxes ~]$ systemctl status httpd
> ● httpd.service - The Apache HTTP Server
>    Loaded: loaded (/usr/lib/systemd/system/httpd.service; disabled; vendor preset:>
>    Active: active (running) since Tue 2022-02-08 09:53:04 EST; 3s ago
>      Docs: man:httpd.service(8)
>  Main PID: 82971 (httpd)
>    Status: "Started, listening on: port 80"
>     Tasks: 213 (limit: 4944)
>    Memory: 25.4M
>    CGroup: /system.slice/httpd.service
>            ├─82971 /usr/sbin/httpd -DFOREGROUND
>            ├─82972 /usr/sbin/httpd -DFOREGROUND
>            ├─82973 /usr/sbin/httpd -DFOREGROUND
>            ├─82974 /usr/sbin/httpd -DFOREGROUND
>            └─82975 /usr/sbin/httpd -DFOREGROUND
>
> Feb 08 09:53:04 osboxes systemd[1]: Starting The Apache HTTP Server...
> Feb 08 09:53:04 osboxes httpd[82971]: AH00558: httpd: Could not reliably determine>
> Feb 08 09:53:04 osboxes systemd[1]: Started The Apache HTTP Server.
> Feb 08 09:53:04 osboxes httpd[82971]: Server configured, listening on: port 80
> lines 1-19/19 (END)
> ^C
> [osboxes@osboxes ~]$ http://43.249.184.35
> -bash: http://43.249.184.35: No such file or directory
> [osboxes@osboxes ~]$ http://192.168.31.255
> -bash: http://192.168.31.255: No such file or directory
> [osboxes@osboxes ~]$ http://192.168.31.27
> -bash: http://192.168.31.27: No such file or directory
> [osboxes@osboxes ~]$ sudo cat /etc/httpd.conf
> [sudo] password for osboxes:
> cat: /etc/httpd.conf: No such file or directory
> [osboxes@osboxes ~]$ sudo -i
> [root@osboxes ~]# cd /etc
> [root@osboxes etc]# ls -al
> total 1192
>
[root@osboxes httpd]# cd conf.d
[root@osboxes conf.d]# ls -al
total 16
drwxr-xr-x. 2 root root   82 Feb  8 09:42 .
drwxr-xr-x. 5 root root  105 Feb  8 09:42 ..
-rw-r--r--. 1 root root  400 Nov 11 23:58 README
-rw-r--r--. 1 root root 2926 Nov 11 23:58 autoindex.conf
-rw-r--r--. 1 root root 1252 Nov 11 23:54 userdir.conf
-rw-r--r--. 1 root root  574 Nov 11 23:54 welcome.conf
[root@osboxes conf.d]# pwd
/etc/httpd/conf.d
[root@osboxes conf.d]# cat welcome.conf
#
# This configuration file enables the default "Welcome" page if there
# is no default index page present for the root URL.  To disable the
# Welcome page, comment out all the lines below.
#
# NOTE: if this file is removed, it will be restored on upgrades.
#
<LocationMatch "^/+$">
    Options -Indexes
    ErrorDocument 403 /.noindex.html
</LocationMatch>

<Directory /usr/share/httpd/noindex>
    AllowOverride None
    Require all granted
</Directory>

Alias /.noindex.html /usr/share/httpd/noindex/index.html
Alias /poweredby.png /usr/share/httpd/icons/apache_pb3.png[root@osboxes conf.d]#
[root@osboxes conf.d]# cat userdir.conf
#
# UserDir: The name of the directory that is appended onto a user's home
# directory if a ~user request is received.
#
# The path to the end user account 'public_html' directory must be
# accessible to the webserver userid.  This usually means that ~userid
# must have permissions of 711, ~userid/public_html must have permissions
# of 755, and documents contained therein must be world-readable.
# Otherwise, the client will only receive a "403 Forbidden" message.
#
<IfModule mod_userdir.c>
    #
    # UserDir is disabled by default since it can confirm the presence
    # of a username on the system (depending on home directory
    # permissions).
    #
    UserDir disabled

    #
    # To enable requests to /~user/ to serve the user's public_html
    # directory, remove the "UserDir disabled" line above, and uncomment
    # the following line instead:
    #
    #UserDir public_html
</IfModule>

#
# Control access to UserDir directories.  The following is an example
# for a site where these directories are restricted to read-only.
#
<Directory "/home/*/public_html">
    AllowOverride FileInfo AuthConfig Limit Indexes
    Options MultiViews Indexes SymLinksIfOwnerMatch IncludesNoExec
    Require method GET POST OPTIONS
</Directory>

[root@osboxes conf.d]# cd /var/log/httpd
[root@osboxes httpd]# ls -al
total 8
drwx------.  2 root root   41 Feb  8 09:53 .
drwxr-xr-x. 10 r> drwxr-xr-x.   2 root root       60 Feb  8 03:47 udisks2
> drwxr-xr-x.   2 root root       45 Feb  8 03:47 unbound
> -rw-r--r--.   1 root root      587 May 11  2019 updatedb.conf
> -rw-r--r--.   1 root root     1523 May 11  2019 usb_modeswitch.conf
> -rw-r--r--.   1 root root       28 Jun 19  2021 vconsole.conf
> -rw-r--r--.   1 root root     1982 Sep 22 07:10 vimrc
> -rw-r--r--.   1 root root     1204 Sep 22 07:10 virc
> drwxr-xr-x.   4 root root      208 Feb  8 03:48 vmware-tools
> -rw-r--r--.   1 root root     4925 Apr 26  2020 wgetrc
> -rw-r--r--.   1 root root      642 Dec  9  2016 xattr.conf
> drwxr-xr-x.   4 root root       38 Jun 22  2021 xdg
> drwxr-xr-x.   2 root root        6 Jun 22  2021 xinetd.d
> drwxr-xr-x.   2 root root       57 Feb  8 03:48 yum
> lrwxrwxrwx.   1 root root       12 Sep 17 15:05 yum.conf -> dnf/dnf.conf
> drwxr-xr-x.   2 root root     4096 Feb  8 03:48 yum.repos.d
> [root@osboxes etc]# cat httpd
> cat: httpd: Is a directory
> [root@osboxes etc]# cd httpd
> [root@osboxes httpd]# ls -al
> total 12
> drwxr-xr-x.   5 root root  105 Feb  8 09:42 .
> drwxr-xr-x. 105 root root 8192 Feb  8 09:42 ..
> drwxr-xr-x.   2 root root   37 Feb  8 09:42 conf
> drwxr-xr-x.   2 root root   82 Feb  8 09:42 conf.d
> drwxr-xr-x.   2 root root  226 Feb  8 09:42 conf.modules.d
> lrwxrwxrwx.   1 root root   19 Nov 11 23:58 logs -> ../../var/log/httpd
> lrwxrwxrwx.   1 root root   29 Nov 11 23:58 modules -> ../../usr/lib64/httpd/modules
> lrwxrwxrwx.   1 root root   10 Nov 11 23:58 run -> /run/httpd
> lrwxrwxrwx.   1 root root   19 Nov 11 23:58 state -> ../../var/lib/httpd
> [root@osboxes httpd]# cat conf.d
> cat: conf.d: Is a directory
> [root@osboxes httpd]# cd conf.d
> [root@osboxes conf.d]# ls -al
> total 16
> drwxr-xr-x. 2 root root   82 Feb  8 09:42 .
> drwxr-xr-x. 5 root root  105 Feb  8 09:42 ..
> -rw-r--r--. 1 root root  400 Nov 11 23:58 README
> -rw-r--r--. 1 root root 2926 Nov 11 23:58 autoindex.conf
> -rw-r--r--. 1 root root 1252 Nov 11 23:54 userdir.conf
> -rw-r--r--. 1 root root  574 Nov 11 23:54 welcome.conf
> [root@osboxes conf.d]# pwd
> /etc/httpd/conf.d
> [root@osboxes conf.d]# cat welcome.conf
> #
> # This configuration file enables the default "Welcome" page if there
> # is no default index page present for the root URL.  To disable the
> # Welcome page, comment out all the lines below.
> #
> # NOTE: if this file is removed, it will be restored on upgrades.
> #
> <LocationMatch "^/+$">
>     Options -Indexes
>     ErrorDocument 403 /.noindex.html
> </LocationMatch>
>
> <Directory /usr/share/httpd/noindex>
>     AllowOverride None
>     Require all granted
> </Directory>
>
> Alias /.noindex.html /usr/share/httpd/noindex/index.html
> Alias /poweredby.png /usr/share/httpd/icons/apache_pb3.png[root@osboxes conf.d]#
> [root@osboxes conf.d]# cat userdir.conf
> #
> # UserDir: The name of the directory that is appended onto a user's home
-bash: osboxes@192.168.31.27s password:
Activate the web console with: systemctl enable --now cockpit.socket

Last login: Mon Feb  7 22:08:57 2022
[osboxes@osboxes ~]$ sudo apt update
[sudo] password for osboxes:
sudo: apt: command not found
[osboxes@osboxes ~]$ sudo dnf update
Last metadata expiration check: 1:43:52 ago on Tue Feb  8 01:51:14 2022.
Dependencies resolved.
===================================================================================
 Package                       Arch   Version                      Repo       Size


Complete!
[osboxes@osboxes ~]$
[osboxes@osboxes ~]$ sudo dnf install apache2
[sudo] password for osboxes:
Last metadata expiration check: 7:47:42 ago on Tue Feb  8 01:51:14 2022.
No match for argument: apache2
Error: Unable to find a match: apache2
[osboxes@osboxes ~]$ sudo dnf search apache
Last metadata expiration check: 7:49:28 ago on Tue Feb  8 01:51:14 2022.
========================= Name & Summary Matched: apache ==========================
apache-commons-logging.noarch : Apache Commons Logging

[osboxes@osboxes ~]$ sudo dnf install httpd


Installed:
  apr-1.6.3-12.el8.x86_64
  apr-util-1.6.1-6.el8.x86_64
  apr-util-bdb-1.6.1-6.el8.x86_64
  apr-util-openssl-1.6.1-6.el8.x86_64
  centos-logos-httpd-85.8-2.el8.noarch
  httpd-2.4.37-43.module_el8.5.0+1022+b541f3b1.x86_64
  httpd-filesystem-2.4.37-43.module_el8.5.0+1022+b541f3b1.noarch
  httpd-tools-2.4.37-43.module_el8.5.0+1022+b541f3b1.x86_64
  mod_http2-1.15.7-3.module_el8.4.0+778+c970deab.x86_64

Complete!
[osboxes@osboxes ~]$ curl https://ipinfo.io/ip
43.249.184.35[osboxes@osboxes ~]$ http://43.249.184.35
-bash: http://43.249.184.35: No such file or directory
[osboxes@osboxes ~]$ systemctl status httpd
● httpd.service - The Apache HTTP Server
   Loaded: loaded (/usr/lib/systemd/system/httpd.service; disabled; vendor preset:>
   Active: inactive (dead)
     Docs: man:httpd.service(8)
lines 1-4/4 (END)
^C
[osboxes@osboxes ~]$ sudo systemctl start httpd
[sudo] password for osboxes:
[osboxes@osboxes ~]$ systemctl status httpd
● httpd.service - The Apache HTTP Server
   Loaded: loaded (/usr/lib/systemd/system/httpd.service; disabled; vendor preset:>
   Active: active (running) since Tue 2022-02-08 09:53:04 EST; 3s ago
     Docs: man:httpd.service(8)
 Main PID: 82971 (httpd)
   Status: "Started, listening on: port 80"
    Tasks: 213 (limit: 4944)
   Memory: 25.4M
   CGroup: /system.slice/httpd.service
           ├─82971 /usr/sbin/httpd -DFOREGROUND
           ├─82972 /usr/sbin/httpd -DFOREGROUND
           ├─82973 /usr/sbin/httpd -DFOREGROUND
           ├─82974 /usr/sbin/httpd -DFOREGROUND
           └─82975 /usr/sbin/httpd -DFOREGROUND

Feb 08 09:53:04 osboxes systemd[1]: Starting The Apache HTTP Server...
Feb 08 09:53:04 osboxes httpd[82971]: AH00558: httpd: Could not reliably determine>
Feb 08 09:53:04 osboxes systemd[1]: Started The Apache HTTP Server.
Feb 08 09:53:04 osboxes httpd[82971]: Server configured, listening on: port 80
lines 1-19/19 (END)
^C
[osboxes@osboxes ~]$ http://43.249.184.35
-bash: http://43.249.184.35: No such file or directory
[osboxes@osboxes ~]$ http://192.168.31.255
-bash: http://192.168.31.255: No such file or directory
[osboxes@osboxes ~]$ http://192.168.31.27
-bash: http://192.168.31.27: No such file or directory
[osboxes@osboxes ~]$ sudo cat /etc/httpd.conf
[sudo] password for osboxes:
cat: /etc/httpd.conf: No such file or directory
[osboxes@osboxes ~]$ sudo -i
[root@osboxes ~]# cd /etc
[root@osboxes etc]# ls -al
total 1192

**[root@osboxes httpd]# [root@osboxes conf.d]# cd /var/log/httpd**
-bash: [root@osboxes: command not found
**[root@osboxes httpd]# [root@osboxes httpd]# ls -al**

