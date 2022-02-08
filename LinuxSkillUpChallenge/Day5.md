  grub2-tools-efi-1:2.02-106.el8.x86_64
  kernel-4.18.0-348.7.1.el8_5.x86_64
  kernel-core-4.18.0-348.7.1.el8_5.x86_64
  kernel-modules-4.18.0-348.7.1.el8_5.x86_64
  libbpf-0.4.0-1.el8.x86_64

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
httpd-devel.x86_64 : Development interfaces for the Apache HTTP server
httpd-filesystem.noarch : The basic directory layout for the Apache HTTP server
httpd-manual.noarch : Documentation for the Apache HTTP server
httpd-tools.x86_64 : Tools for use with the Apache HTTP Server
keycloak-httpd-client-install.noarch : Tools to configure Apache HTTPD as Keycloak
                                     : client
librdkafka.i686 : The Apache Kafka C library
librdkafka.x86_64 : The Apache Kafka C library
mod_auth_gssapi.x86_64 : A GSSAPI Authentication module for Apache
mod_auth_mellon.x86_64 : A SAML 2.0 authentication module for the Apache Httpd
                       : Server
mod_dav_svn.x86_64 : Apache httpd module for Subversion server
mod_fcgid.x86_64 : FastCGI interface module for Apache 2
mod_http2.x86_64 : module implementing HTTP/2 for Apache 2
mod_intercept_form_submit.x86_64 : Apache module to intercept login form submission
                                 : and run PAM authentication
mod_ldap.x86_64 : LDAP authentication modules for the Apache HTTP Server
mod_lookup_identity.x86_64 : Apache module to retrieve additional information about
                           : the authenticated user
mod_md.x86_64 : Certificate provisioning using ACME for the Apache HTTP Server
mod_proxy_html.x86_64 : HTML and XML content filters for the Apache HTTP Server
mod_security.x86_64 : Security module for the Apache HTTP Server
mod_session.x86_64 : Session interface for the Apache HTTP Server
mod_ssl.x86_64 : SSL/TLS module for the Apache HTTP Server
pcp-export-pcp2spark.x86_64 : Performance Co-Pilot tools for exporting PCP metrics
                            : to Apache Spark
python3-keycloak-httpd-client-install.noarch : Tools to configure Apache HTTPD as
                                             : Keycloak client
python3-mod_wsgi.x86_64 : A WSGI interface for Python web applications in Apache
python38-mod_wsgi.x86_64 : A WSGI interface for Python web applications in Apache
python39-mod_wsgi.x86_64 : A WSGI interface for Python web applications in Apache
[osboxes@osboxes ~]$ sudo dnf install apache http server
Last metadata expiration check: 7:50:30 ago on Tue Feb  8 01:51:14 2022.
No match for argument: apache
No match for argument: http
No match for argument: server
Error: Unable to find a match: apache http server
[osboxes@osboxes ~]$ sudo dnf install httpd
Last metadata expiration check: 7:50:50 ago on Tue Feb  8 01:51:14 2022.
Dependencies resolved.
===================================================================================
 Package            Arch   Version                                 Repo       Size
===================================================================================
Installing:
 httpd              x86_64 2.4.37-43.module_el8.5.0+1022+b541f3b1  appstream 1.4 M
Installing dependencies:
 apr                x86_64 1.6.3-12.el8                            appstream 129 k
 apr-util           x86_64 1.6.1-6.el8                             appstream 105 k
 centos-logos-httpd noarch 85.8-2.el8                              baseos     75 k
 httpd-filesystem   noarch 2.4.37-43.module_el8.5.0+1022+b541f3b1  appstream  39 k
 httpd-tools        x86_64 2.4.37-43.module_el8.5.0+1022+b541f3b1  appstream 107 k
 mod_http2          x86_64 1.15.7-3.module_el8.4.0+778+c970deab    appstream 154 k
Installing weak dependencies:
 apr-util-bdb       x86_64 1.6.1-6.el8                             appstream  25 k
 apr-util-openssl   x86_64 1.6.1-6.el8                             appstream  27 k
Enabling module streams:
 httpd                     2.4

Transaction Summary
===================================================================================
Install  9 Packages

Total download size: 2.1 M
Installed size: 5.6 M
Is this ok [y/N]: y
Downloading Packages:
(1/9): apr-util-bdb-1.6.1-6.el8.x86_64.rpm          11 kB/s |  25 kB     00:02
(2/9): apr-util-1.6.1-6.el8.x86_64.rpm              37 kB/s | 105 kB     00:02
(3/9): apr-util-openssl-1.6.1-6.el8.x86_64.rpm      43 kB/s |  27 kB     00:00
(4/9): apr-1.6.3-12.el8.x86_64.rpm                  46 kB/s | 129 kB     00:02
(5/9): httpd-filesystem-2.4.37-43.module_el8.5.0+1  21 kB/s |  39 kB     00:01
(6/9): httpd-2.4.37-43.module_el8.5.0+1022+b541f3b 568 kB/s | 1.4 MB     00:02
(7/9): mod_http2-1.15.7-3.module_el8.4.0+778+c970d 185 kB/s | 154 kB     00:00
(8/9): centos-logos-httpd-85.8-2.el8.noarch.rpm    135 kB/s |  75 kB     00:00
(9/9): httpd-tools-2.4.37-43.module_el8.5.0+1022+b  34 kB/s | 107 kB     00:03
-----------------------------------------------------------------------------------
Total                                              349 kB/s | 2.1 MB     00:06
Running transaction check
Transaction check succeeded.
Running transaction test
Transaction test succeeded.
Running transaction
  Preparing        :                                                           1/1
  Installing       : apr-1.6.3-12.el8.x86_64                                   1/9
  Running scriptlet: apr-1.6.3-12.el8.x86_64                                   1/9
  Installing       : apr-util-bdb-1.6.1-6.el8.x86_64                           2/9
  Installing       : apr-util-openssl-1.6.1-6.el8.x86_64                       3/9
  Installing       : apr-util-1.6.1-6.el8.x86_64                               4/9
  Running scriptlet: apr-util-1.6.1-6.el8.x86_64                               4/9
  Installing       : httpd-tools-2.4.37-43.module_el8.5.0+1022+b541f3b1.x86_   5/9
  Installing       : centos-logos-httpd-85.8-2.el8.noarch                      6/9
  Running scriptlet: httpd-filesystem-2.4.37-43.module_el8.5.0+1022+b541f3b1   7/9
  Installing       : httpd-filesystem-2.4.37-43.module_el8.5.0+1022+b541f3b1   7/9
  Installing       : mod_http2-1.15.7-3.module_el8.4.0+778+c970deab.x86_64     8/9
  Installing       : httpd-2.4.37-43.module_el8.5.0+1022+b541f3b1.x86_64       9/9
  Running scriptlet: httpd-2.4.37-43.module_el8.5.0+1022+b541f3b1.x86_64       9/9
  Verifying        : apr-1.6.3-12.el8.x86_64                                   1/9
  Verifying        : apr-util-1.6.1-6.el8.x86_64                               2/9
  Verifying        : apr-util-bdb-1.6.1-6.el8.x86_64                           3/9
  Verifying        : apr-util-openssl-1.6.1-6.el8.x86_64                       4/9
  Verifying        : httpd-2.4.37-43.module_el8.5.0+1022+b541f3b1.x86_64       5/9
  Verifying        : httpd-filesystem-2.4.37-43.module_el8.5.0+1022+b541f3b1   6/9
  Verifying        : httpd-tools-2.4.37-43.module_el8.5.0+1022+b541f3b1.x86_   7/9
  Verifying        : mod_http2-1.15.7-3.module_el8.4.0+778+c970deab.x86_64     8/9
  Verifying        : centos-logos-httpd-85.8-2.el8.noarch                      9/9

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
drwxr-xr-x. 105 root root     8192 Feb  8 09:42 .
dr-xr-xr-x.  17 root root      224 Feb  8 03:40 ..
-rw-------.   1 root root        0 Jun 19  2021 .pwd.lock
-rw-r--r--.   1 root root      208 Jun 19  2021 .updated
-rw-r--r--.   1 root root     4536 Jul 14  2021 DIR_COLORS
-rw-r--r--.   1 root root     5214 Jul 14  2021 DIR_COLORS.256color
-rw-r--r--.   1 root root     4618 Jul 14  2021 DIR_COLORS.lightbgcolor
-rw-r--r--.   1 root root       94 May 10  2019 GREP_COLORS
drwxr-xr-x.   7 root root      134 Feb  8 03:46 NetworkManager
drwxr-xr-x.   2 root root       48 Jun 19  2021 PackageKit
drwxr-xr-x.   6 root root       70 Jun 22  2021 X11
-rw-r--r--.   1 root root       16 Jun 19  2021 adjtime
-rw-r--r--.   1 root root     1529 May 15  2020 aliases
drwxr-xr-x.   2 root root      194 Jul 28  2021 alternatives
-rw-r--r--.   1 root root      541 Nov  8  2019 anacrontab
-rw-r--r--.   1 root root        1 May 11  2019 at.deny
drwxr-x---.   4 root root      100 Jun 19  2021 audit
drwxr-xr-x.   3 root root      228 Feb  8 09:35 authselect
drwxr-xr-x.   2 root root      135 Feb  8 03:49 bash_completion.d
-rw-r--r--.   1 root root     3019 May 15  2020 bashrc
-rw-r--r--.   1 root root      535 Apr 21  2021 bindresvport.blacklist
drwxr-xr-x.   2 root root        6 Dec 21 15:08 binfmt.d
-rw-r--r--.   1 root root       30 Nov  9 19:17 centos-release
-rw-r--r--.   1 root root       42 Nov  9 19:17 centos-release-upstream
drwxr-xr-x.   2 root root        6 Jul 28  2021 chkconfig.d
-rw-r--r--.   1 root root     1085 Jun 24  2021 chrony.conf
-rw-r-----.   1 root chrony    540 May 12  2021 chrony.keys
drwxr-xr-x.   2 root root       26 Dec 21 15:14 cifs-utils
drwxr-xr-x.   3 root root       19 Jun 19  2021 cni
drwxr-xr-x.   4 root root       42 Sep 10 10:51 cockpit
drwxr-xr-x.   6 root root      139 Feb  8 03:48 containers
drwxr-xr-x.   2 root root       39 Jun 19  2021 cron.d
drwxr-xr-x.   2 root root       23 Jun 19  2021 cron.daily
-rw-r--r--.   1 root root        0 Nov  8  2019 cron.deny
drwxr-xr-x.   2 root root       22 Jun 19  2021 cron.hourly
drwxr-xr-x.   2 root root        6 Jan 12  2021 cron.monthly
drwxr-xr-x.   2 root root        6 Jan 12  2021 cron.weekly
-rw-r--r--.   1 root root      451 Jan 12  2021 crontab
drwxr-xr-x.   6 root root       81 Jun 19  2021 crypto-policies
-rw-------.   1 root root        0 Jun 19  2021 crypttab
-rw-r--r--.   1 root root     1629 May 15  2020 csh.cshrc
-rw-r--r--.   1 root root     1078 May 15  2020 csh.login
drwxr-xr-x.   4 root root       78 May  8  2021 dbus-1
drwxr-xr-x.   3 root root       16 Jun 19  2021 dconf
drwxr-xr-x.   2 root root       33 Feb  8 03:48 default
drwxr-xr-x.   2 root root       40 Feb  8 03:46 depmod.d
drwxr-x---.   3 root root       45 Jul 22  2021 dhcp
drwxr-xr-x.   8 root root      128 Sep 17 15:05 dnf
-rw-r--r--.   1 root root      117 Sep 29 17:07 dracut.conf
drwxr-xr-x.   2 root root        6 Sep 29 17:07 dracut.conf.d
-rw-r--r--.   1 root root        0 May 15  2020 environment
-rw-r--r--.   1 root root     1362 Aug 24 19:13 ethertypes
-rw-r--r--.   1 root root        0 Sep 10  2018 exports
lrwxrwxrwx.   1 root root       56 Nov 12 13:19 favicon.png -> /usr/share/icons/hicolor/16x16/apps/fedora-logo-icon.png
-rw-r--r--.   1 root root       66 Sep 10  2018 filesystems
drwxr-x---.   8 root root      149 Feb  8 03:48 firewalld
drwxr-xr-x.   3 root root       38 Jun 19  2021 fonts
-rw-r--r--.   1 root root       20 Jan 13  2021 fprintd.conf
-rw-r--r--.   1 root root      709 Jun 19  2021 fstab
-rw-r--r--.   1 root root       38 May 11  2019 fuse.conf
drwxr-xr-x.   2 root root       25 Jun 29  2021 gcrypt
drwxr-xr-x.   2 root root        6 May 15  2020 gnupg
drwxr-xr-x.   4 root root       40 Jun 19  2021 groff
-rw-r--r--.   1 root root      691 Feb  8 09:42 group
-rw-r--r--.   1 root root      678 Jun 19  2021 group-
drwx------.   2 root root      288 Feb  8 03:44 grub.d
lrwxrwxrwx.   1 root root       22 Nov  8 01:39 grub2.cfg -> ../boot/grub2/grub.cfg
----------.   1 root root      555 Feb  8 09:42 gshadow
----------.   1 root root      544 Jun 19  2021 gshadow-
drwxr-xr-x.   3 root root       20 Aug 26 01:12 gss
-rw-r--r--.   1 root root        9 Sep 10  2018 host.conf
-rw-r--r--.   1 root root        8 Feb  5 02:10 hostname
-rw-r--r--.   1 root root      158 Sep 10  2018 hosts
drwxr-xr-x.   5 root root      105 Feb  8 09:42 httpd
-rw-r--r--.   1 root root     4849 Jul 30  2021 idmapd.conf
lrwxrwxrwx.   1 root root       11 Jul 28  2021 init.d -> rc.d/init.d
-rw-r--r--.   1 root root      490 Dec 21 15:08 inittab
-rw-r--r--.   1 root root      942 Sep 10  2018 inputrc
drwxr-xr-x.   2 root root      159 Oct 15 13:08 iproute2
drwxr-xr-x.   2 root root       52 Aug 10 01:10 iscsi
-rw-r--r--.   1 root root       23 Nov  9 19:17 issue
drwxr-xr-x.   2 root root       27 Jun 19  2021 issue.d
-rw-r--r--.   1 root root       22 Nov  9 19:17 issue.net
drwxr-xr-x.   4 root root       33 Dec 21 15:15 kdump
-rw-r--r--.   1 root root     8550 Feb  8 03:47 kdump.conf
drwxr-xr-x.   4 root root       41 Dec 21 15:08 kernel
-rw-r--r--.   1 root root      812 Aug 26 01:05 krb5.conf
drwxr-xr-x.   2 root root       55 Feb  8 03:48 krb5.conf.d
-rw-r--r--.   1 root root    21275 Feb  8 09:42 ld.so.cache
-rw-r--r--.   1 root root       28 Aug 24 19:10 ld.so.conf
drwxr-xr-x.   2 root root      129 Feb  8 03:44 ld.so.conf.d
-rw-r-----.   1 root root      191 Nov  4  2019 libaudit.conf
drwxr-xr-x.   3 root root       20 Jul  2  2021 libblockdev
drwxr-xr-x.   2 root root      246 May 17  2021 libibverbs.d
drwxr-xr-x.   2 root root       35 Jun 19  2021 libnl
drwxr-xr-x.   6 root root       70 Jun 19  2021 libreport
drwxr-xr-x.   2 root root       62 May  5  2021 libssh
-rw-r--r--.   1 root root     2391 Jul 23  2015 libuser.conf
-rw-r--r--.   1 root root       19 Jun 19  2021 locale.conf
lrwxrwxrwx.   1 root root       38 Feb  5 02:29 localtime -> ../usr/share/zoneinfo/America/New_York
-rw-r--r--.   1 root root     2512 Aug 18 15:04 login.defs
-rw-r--r--.   1 root root      438 Feb 19  2018 logrotate.conf
drwxr-xr-x.   2 root root      188 Feb  8 09:42 logrotate.d
drwxr-xr-x.   3 root root       43 Apr 22  2021 lsm
drwxr-xr-x.   6 root root      100 Feb  8 03:46 lvm
-r--r--r--.   1 root root       33 Jun 19  2021 machine-id
-rw-r--r--.   1 root root      111 May  8  2021 magic
-rw-r--r--.   1 root root      272 May 11  2017 mailcap
-rw-r--r--.   1 root root     5122 Dec 21 15:15 makedumpfile.conf.sample
-rw-r--r--.   1 root root     5165 Jun 29  2021 man_db.conf
drwxr-xr-x.   2 root root      181 Feb  5 05:47 mc
drwxr-xr-x.   3 root root       41 Aug  9  2021 mcelog
drwxr-xr-x.   3 root root       32 Jul 24  2021 microcode_ctl
-rw-r--r--.   1 root root    60352 May 11  2017 mime.types
-rw-r--r--.   1 root root     1108 Jun 24  2021 mke2fs.conf
drwxr-xr-x.   2 root root       93 Feb  8 03:48 modprobe.d
drwxr-xr-x.   2 root root        6 Dec 21 15:08 modules-load.d
-rw-r--r--.   1 root root        0 Sep 10  2018 motd
drwxr-xr-x.   2 root root       21 Jun 19  2021 motd.d
lrwxrwxrwx.   1 root root       19 Jun 19  2021 mtab -> ../proc/self/mounts
drwxr-xr-x.   2 root root        6 Jul 28  2021 multipath
-rw-r--r--.   1 root root     9450 May 11  2019 nanorc
-rw-r--r--.   1 root root      767 Apr 21  2021 netconfig
-rw-r--r--.   1 root root       58 Sep 10  2018 networks
drwx------.   3 root root       66 Aug 24 19:25 nftables
lrwxrwxrwx.   1 root root       29 Feb  8 09:35 nsswitch.conf -> /etc/authselect/nsswitch.conf
-rw-r--r--.   1 root root     2197 Mar 11  2021 nsswitch.conf.bak
drwxr-xr-x.   2 root root       35 Sep 17 15:06 nvme
drwxr-xr-x.   2 root root        6 Dec 16  2020 oddjob
-rw-r--r--.   1 root root     4922 Dec 16  2020 oddjobd.conf
drwxr-xr-x.   2 root root       70 Jun 19  2021 oddjobd.conf.d
drwxr-xr-x.   3 root root       36 Aug 10 01:12 openldap
drwxr-xr-x.   2 root root        6 Jun 22  2021 opt
lrwxrwxrwx.   1 root root       21 Nov  9 19:17 os-release -> ../usr/lib/os-release
drwxr-xr-x.   2 root root     4096 Feb  8 09:35 pam.d
-rw-r--r--.   1 root root     1594 Feb  8 09:42 passwd
-rw-r--r--.   1 root root     1541 Jun 19  2021 passwd-
-rw-r--r--.   1 root root     2872 May 14  2019 pinforc
drwxr-xr-x.   3 root root       21 Jun 19  2021 pkcs11
drwxr-xr-x.   8 root root       88 Jun 22  2021 pki
drwxr-xr-x.   2 root root       28 Aug 24 19:28 plymouth
drwxr-xr-x.   5 root root       52 Jun 22  2021 pm
drwxr-xr-x.   5 root root       72 Jun  2  2021 polkit-1
drwxr-xr-x.   2 root root        6 Jan 19  2021 popt.d
drwxr-xr-x.   2 root root       24 Feb  8 03:43 prelink.conf.d
-rw-r--r--.   1 root root      233 Sep 10  2018 printcap
-rw-r--r--.   1 root root     2123 May 15  2020 profile
drwxr-xr-x.   2 root root     4096 Feb  8 03:41 profile.d
-rw-r--r--.   1 root root     6568 Sep 10  2018 protocols
drwxr-xr-x.  10 root root      127 Jul 28  2021 rc.d
lrwxrwxrwx.   1 root root       13 Dec 21 15:08 rc.local -> rc.d/rc.local
lrwxrwxrwx.   1 root root       10 Jul 28  2021 rc0.d -> rc.d/rc0.d
lrwxrwxrwx.   1 root root       10 Jul 28  2021 rc1.d -> rc.d/rc1.d
lrwxrwxrwx.   1 root root       10 Jul 28  2021 rc2.d -> rc.d/rc2.d
lrwxrwxrwx.   1 root root       10 Jul 28  2021 rc3.d -> rc.d/rc3.d
lrwxrwxrwx.   1 root root       10 Jul 28  2021 rc4.d -> rc.d/rc4.d
lrwxrwxrwx.   1 root root       10 Jul 28  2021 rc5.d -> rc.d/rc5.d
lrwxrwxrwx.   1 root root       10 Jul 28  2021 rc6.d -> rc.d/rc6.d
drwxr-xr-x.   3 root root       38 May 17  2021 rdma
lrwxrwxrwx.   1 root root       14 Nov  9 19:17 redhat-release -> centos-release
-rw-r--r--.   1 root root       54 Feb  7 22:08 resolv.conf
drwxr-xr-x.   3 root root       24 Jun 19  2021 rhsm
-rw-r--r--.   1 root root     1634 Aug  1  2018 rpc
drwxr-xr-x.   2 root root       25 Feb  8 03:48 rpm
-rw-r--r--.   1 root root     3186 Aug 10 12:43 rsyslog.conf
drwxr-xr-x.   2 root root        6 Aug 10 12:46 rsyslog.d
drwxr-xr-x.   2 root root       35 Dec 21 15:14 rwtab.d
drwxr-xr-x.   2 root root       61 Feb  8 03:44 samba
drwxr-xr-x.   2 root root        6 May 15  2020 sasl2
drwxr-xr-x.   7 root root     4096 May  7  2021 security
drwxr-xr-x.   3 root root       57 Dec 21 15:17 selinux
-rw-r--r--.   1 root root   692252 May 15  2020 services
-rw-r--r--.   1 root root      216 Sep 30 01:08 sestatus.conf
drwxr-xr-x.   2 root root       33 Sep 30 01:08 setroubleshoot
----------.   1 root root      972 Feb  8 09:42 shadow
----------.   1 root root      950 Jun 19  2021 shadow-
-rw-r--r--.   1 root root       44 Sep 10  2018 shells
drwxr-xr-x.   2 root root       62 Jun 22  2021 skel
drwxr-xr-x.   3 root root       74 Jun 19  2021 smartmontools
drwxr-xr-x.   6 root root       86 Feb  8 03:46 sos
drwxr-xr-x.   3 root root      245 Jul 13  2021 ssh
drwxr-xr-x.   2 root root       19 Feb  8 03:43 ssl
drwx------.   4 sssd sssd       31 Dec 21 15:14 sssd
-rw-r--r--.   1 root root       21 Jun 19  2021 subgid
-rw-r--r--.   1 root root        0 Sep 10  2018 subgid-
-rw-r--r--.   1 root root       21 Jun 19  2021 subuid
-rw-r--r--.   1 root root        0 Sep 10  2018 subuid-
-rw-r-----.   1 root root     3181 Oct 25 10:27 sudo-ldap.conf
-rw-r-----.   1 root root     1786 Oct 25 10:27 sudo.conf
-r--r-----.   1 root root     4328 Oct 25 10:27 sudoers
drwxr-x---.   2 root root        6 Oct 25 10:29 sudoers.d
drwxr-xr-x.   5 root root     4096 Feb  8 09:42 sysconfig
-rw-r--r--.   1 root root      449 Dec 21 15:08 sysctl.conf
drwxr-xr-x.   2 root root       28 Dec 21 15:08 sysctl.d
lrwxrwxrwx.   1 root root       14 Nov  9 19:17 system-release -> centos-release
-rw-r--r--.   1 root root       23 Nov  9 19:17 system-release-cpe
drwxr-xr-x.   4 root root      150 Feb  8 03:44 systemd
-rw-r-----.   1 root tss      7046 Nov 16  2020 tcsd.conf
drwxr-xr-x.   2 root root        6 Jun  1  2021 terminfo
drwxr-xr-x.   2 root root        6 Dec 21 15:08 tmpfiles.d
-rw-r--r--.   1 root root      375 Aug 24 19:20 trusted-key.key
drwxr-xr-x.   3 root root      136 Feb  8 03:48 tuned
drwxr-xr-x.   4 root root       68 Feb  8 09:37 udev
drwxr-xr-x.   2 root root       60 Feb  8 03:47 udisks2
drwxr-xr-x.   2 root root       45 Feb  8 03:47 unbound
-rw-r--r--.   1 root root      587 May 11  2019 updatedb.conf
-rw-r--r--.   1 root root     1523 May 11  2019 usb_modeswitch.conf
-rw-r--r--.   1 root root       28 Jun 19  2021 vconsole.conf
-rw-r--r--.   1 root root     1982 Sep 22 07:10 vimrc
-rw-r--r--.   1 root root     1204 Sep 22 07:10 virc
drwxr-xr-x.   4 root root      208 Feb  8 03:48 vmware-tools
-rw-r--r--.   1 root root     4925 Apr 26  2020 wgetrc
-rw-r--r--.   1 root root      642 Dec  9  2016 xattr.conf
drwxr-xr-x.   4 root root       38 Jun 22  2021 xdg
drwxr-xr-x.   2 root root        6 Jun 22  2021 xinetd.d
drwxr-xr-x.   2 root root       57 Feb  8 03:48 yum
lrwxrwxrwx.   1 root root       12 Sep 17 15:05 yum.conf -> dnf/dnf.conf
drwxr-xr-x.   2 root root     4096 Feb  8 03:48 yum.repos.d
[root@osboxes etc]# cat httpd
cat: httpd: Is a directory
[root@osboxes etc]# cd httpd
[root@osboxes httpd]# ls -al
total 12
drwxr-xr-x.   5 root root  105 Feb  8 09:42 .
drwxr-xr-x. 105 root root 8192 Feb  8 09:42 ..
drwxr-xr-x.   2 root root   37 Feb  8 09:42 conf
drwxr-xr-x.   2 root root   82 Feb  8 09:42 conf.d
drwxr-xr-x.   2 root root  226 Feb  8 09:42 conf.modules.d
lrwxrwxrwx.   1 root root   19 Nov 11 23:58 logs -> ../../var/log/httpd
lrwxrwxrwx.   1 root root   29 Nov 11 23:58 modules -> ../../usr/lib64/httpd/modules
lrwxrwxrwx.   1 root root   10 Nov 11 23:58 run -> /run/httpd
lrwxrwxrwx.   1 root root   19 Nov 11 23:58 state -> ../../var/lib/httpd
[root@osboxes httpd]# cat conf.d
cat: conf.d: Is a directory
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
drwxr-xr-x. 10 root root 4096 Feb  8 09:42 ..
-rw-r--r--.  1 root root    0 Feb  8 09:53 access_log
-rw-r--r--.  1 root root 1039 Feb  8 09:53 error_log
[root@osboxes httpd]# cat access_log
[root@osboxes httpd]# vim access_log
[root@osboxes httpd]# file access_log
access_log: empty
[root@osboxes httpd]# Welcome to Ubuntu 20.04 LTS (GNU/Linux 4.4.0-19041-Microsoft x86_64)
-bash: syntax error near unexpected token `('
[root@osboxes httpd]#
[root@osboxes httpd]#  * Documentation:  https://help.ubuntu.com
o-bash: access_log: command not found
[root@osboxes httpd]#  * Management:     https://landscape.canonical.com
-bash: access_log: command not found
[root@osboxes httpd]#  * Support:        https://ubuntu.com/advantage
-bash: access_log: command not found
[root@osboxes httpd]#
[root@osboxes httpd]#   System information as of Tue Feb  8 19:37:59 IST 2022
-bash: System: command not found
[root@osboxes httpd]#
[root@osboxes httpd]#   System load:    0.52      Users logged in:        0
-bash: System: command not found
[root@osboxes httpd]#   Usage of /home: unknown   IPv4 address for eth1:  192.168.114.1
-bash: Usage: command not found
[root@osboxes httpd]#   Memory usage:   60%       IPv4 address for eth2:  192.168.192.1
-bash: Memory: command not found
[root@osboxes httpd]#   Swap usage:     0%        IPv4 address for wifi0: 192.168.31.140
-bash: Swap: command not found
[root@osboxes httpd]#   Processes:      7
-bash: Processes:: command not found
[root@osboxes httpd]#
[root@osboxes httpd]# 0 updates can be installed immediately.
-bash: 0: command not found
[root@osboxes httpd]# 0 of these updates are security updates.
-bash: 0: command not found
[root@osboxes httpd]#
[root@osboxes httpd]#
[root@osboxes httpd]# The list of available updates is more than a week old.
-bash: The: command not found
[root@osboxes httpd]# To check for new updates run: sudo apt update
-bash: To: command not found
[root@osboxes httpd]#
[root@osboxes httpd]#
[root@osboxes httpd]# This message is shown once once a day. To disable it please create the
-bash: This: command not found
[root@osboxes httpd]# /home/heloise/.hushlogin file.
-bash: /home/heloise/.hushlogin: No such file or directory
[root@osboxes httpd]# heloise@DESKTOP-P5D1DV6:~$ ssh osboxes@192.168.31.27
-bash: heloise@DESKTOP-P5D1DV6:~$: command not found
[root@osboxes httpd]# osboxes@192.168.31.27's password:
> Activate the web console with: systemctl enable --now cockpit.socket
>
> Last login: Mon Feb  7 22:08:57 2022
> [osboxes@osboxes ~]$ sudo apt update
> [sudo] password for osboxes:
> sudo: apt: command not found
> [osboxes@osboxes ~]$ sudo dnf update
> Last metadata expiration check: 1:43:52 ago on Tue Feb  8 01:51:14 2022.
> Dependencies resolved.
> ===================================================================================
>  Package                       Arch   Version                      Repo       Size
> ===================================================================================
> Installing:
>  kernel                        x86_64 4.18.0-348.7.1.el8_5         baseos    7.0 M
>  kernel-core                   x86_64 4.18.0-348.7.1.el8_5         baseos     38 M
>  kernel-modules                x86_64 4.18.0-348.7.1.el8_5         baseos     30 M
> Upgrading:
>  NetworkManager                x86_64 1:1.32.10-4.el8              baseos    2.6 M
>  NetworkManager-config-server  noarch 1:1.32.10-4.el8              baseos    131 k
>  NetworkManager-libnm          x86_64 1:1.32.10-4.el8              baseos    1.8 M
>  NetworkManager-team           x86_64 1:1.32.10-4.el8              baseos    148 k
>  NetworkManager-tui            x86_64 1:1.32.10-4.el8              baseos    336 k
>  adcli                         x86_64 0.8.2-12.el8                 baseos    118 k
>  alsa-sof-firmware             noarch 1.8-1.el8                    baseos    676 k
>  authselect                    x86_64 1.2.2-3.el8                  baseos    133 k
>  authselect-compat             x86_64 1.2.2-3.el8                  appstream  38 k
>  authselect-libs               x86_64 1.2.2-3.el8                  baseos    222 k
>  bash                          x86_64 4.4.20-2.el8                 baseos    1.5 M
>  bind-export-libs              x86_64 32:9.11.26-6.el8             baseos    1.1 M
>  bind-libs                     x86_64 32:9.11.26-6.el8             appstream 174 k
>  bind-libs-lite                x86_64 32:9.11.26-6.el8             appstream 1.2 M
>  bind-license                  noarch 32:9.11.26-6.el8             appstream 102 k
>  bind-utils                    x86_64 32:9.11.26-6.el8             appstream 451 k
>  binutils                      x86_64 2.30-108.el8_5.1             baseos    5.8 M
>  bpftool                       x86_64 4.18.0-348.7.1.el8_5         baseos    7.7 M
>  buildah                       x86_64 1.22.3-2.module_el8.5.0+911+f19012f9
>                                                                    appstream 7.7 M
>  ca-certificates               noarch 2021.2.50-80.0.el8_4         baseos    390 k
>  centos-gpg-keys               noarch 1:8-3.el8                    baseos     12 k
>  centos-linux-release          noarch 8.5-1.2111.el8               baseos     22 k
>  centos-linux-repos            noarch 8-3.el8                      baseos     20 k
>  centos-logos                  x86_64 85.8-2.el8                   baseos    700 k
>  chkconfig                     x86_64 1.19.1-1.el8                 baseos    198 k
>  chrony                        x86_64 4.1-1.el8                    baseos    327 k
>  cockpit                       x86_64 251.1-1.el8                  baseos     78 k
>  cockpit-bridge                x86_64 251.1-1.el8                  baseos    539 k
>  cockpit-packagekit            noarch 251.1-1.el8                  appstream 703 k
>  cockpit-podman                noarch 33-1.module_el8.5.0+890+6b136101
>                                                                    appstream 437 k
>  cockpit-storaged              noarch 251.1-1.el8                  appstream 655 k
>  cockpit-system                noarch 251.1-1.el8                  baseos    3.2 M
>  cockpit-ws                    x86_64 251.1-1.el8                  baseos    1.3 M
>  conmon                        x86_64 2:2.0.29-1.module_el8.5.0+890+6b136101
>                                                                    appstream  52 k
>  container-selinux             noarch 2:2.167.0-1.module_el8.5.0+911+f19012f9
>                                                                    appstream  54 k
>  containernetworking-plugins   x86_64 1.0.0-1.module_el8.5.0+890+6b136101
>                                                                    appstream  19 M
>  containers-common             noarch 2:1-2.module_el8.5.0+890+6b136101
>                                                                    appstream  79 k
>  coreutils                     x86_64 8.30-12.el8                  baseos    1.2 M
>  coreutils-common              x86_64 8.30-12.el8                  baseos    2.0 M
>  criu                          x86_64 3.15-3.module_el8.5.0+890+6b136101
>                                                                    appstream 518 k
>  crypto-policies               noarch 20210617-1.gitc776d3e.el8    baseos     63 k
>  crypto-policies-scripts       noarch 20210617-1.gitc776d3e.el8    baseos     83 k
>  cups-libs                     x86_64 1:2.2.6-40.el8               baseos    433 k
>  curl                          x86_64 7.61.1-22.el8                baseos    351 k
>  dbus                          x86_64 1:1.12.8-14.el8              baseos     41 k
>  dbus-common                   noarch 1:1.12.8-14.el8              baseos     46 k
>  dbus-daemon                   x86_64 1:1.12.8-14.el8              baseos    240 k
>  dbus-libs                     x86_64 1:1.12.8-14.el8              baseos    184 k
>  dbus-tools                    x86_64 1:1.12.8-14.el8              baseos     85 k
>  device-mapper                 x86_64 8:1.02.177-10.el8            baseos    377 k
>  device-mapper-event           x86_64 8:1.02.177-10.el8            baseos    271 k
>  device-mapper-event-libs      x86_64 8:1.02.177-10.el8            baseos    270 k
>  device-mapper-libs            x86_64 8:1.02.177-10.el8            baseos    409 k
>  device-mapper-multipath       x86_64 0.8.4-17.el8                 baseos    197 k
>  device-mapper-multipath-libs  x86_64 0.8.4-17.el8                 baseos    322 k
>  device-mapper-persistent-data x86_64 0.9.0-4.el8                  baseos    925 k
>  dhcp-client                   x86_64 12:4.3.6-45.el8              baseos    318 k
>  dhcp-common                   noarch 12:4.3.6-45.el8              baseos    207 k
>  dhcp-libs                     x86_64 12:4.3.6-45.el8              baseos    148 k
>  dmidecode                     x86_64 1:3.2-10.el8                 baseos     91 k
>  dnf                           noarch 4.7.0-4.el8                  baseos    544 k
>  dnf-data                      noarch 4.7.0-4.el8                  baseos    154 k
>  dnf-plugins-core              noarch 4.0.21-3.el8                 baseos     70 k
>  dracut                        x86_64 049-191.git20210920.el8      baseos    374 k
>  dracut-config-rescue          x86_64 049-191.git20210920.el8      baseos     61 k
>  dracut-network                x86_64 049-191.git20210920.el8      baseos    108 k
>  dracut-squash                 x86_64 049-191.git20210920.el8      baseos     61 k
>  e2fsprogs                     x86_64 1.45.6-2.el8                 baseos    1.0 M
>  e2fsprogs-libs                x86_64 1.45.6-2.el8                 baseos    233 k
>  elfutils-debuginfod-client    x86_64 0.185-1.el8                  baseos     66 k
>  elfutils-default-yama-scope   noarch 0.185-1.el8                  baseos     49 k
>  elfutils-libelf               x86_64 0.185-1.el8                  baseos    221 k
>  elfutils-libs                 x86_64 0.185-1.el8                  baseos    292 k
>  emacs-filesystem              noarch 1:26.1-7.el8                 baseos     70 k
>  ethtool                       x86_64 2:5.8-7.el8                  baseos    209 k
>  file                          x86_64 5.33-20.el8                  baseos     77 k
>  file-libs                     x86_64 5.33-20.el8                  baseos    543 k
>  filesystem                    x86_64 3.8-6.el8                    baseos    1.1 M
>  firewalld                     noarch 0.9.3-7.el8                  baseos    502 k
>  firewalld-filesystem          noarch 0.9.3-7.el8                  baseos     77 k
>  fontconfig                    x86_64 2.13.1-4.el8                 baseos    274 k
>  fstrm                         x86_64 0.6.1-2.el8                  appstream  29 k
>  fuse-overlayfs                x86_64 1.7.1-1.module_el8.5.0+890+6b136101
>                                                                    appstream  73 k
>  glib2                         x86_64 2.56.4-156.el8               baseos    2.5 M
>  glibc                         x86_64 2.28-164.el8                 baseos    3.6 M
>  glibc-common                  x86_64 2.28-164.el8                 baseos    1.3 M
>  glibc-langpack-en             x86_64 2.28-164.el8                 baseos    828 k
>  gnutls                        x86_64 3.6.16-4.el8                 baseos    1.0 M
>  gpgme                         x86_64 1.13.1-9.el8                 baseos    336 k
>  grub2-common                  noarch 1:2.02-106.el8               baseos    891 k
>  grub2-pc                      x86_64 1:2.02-106.el8               baseos     42 k
>  grub2-pc-modules              noarch 1:2.02-106.el8               baseos    916 k
>  grub2-tools                   x86_64 1:2.02-106.el8               baseos    2.0 M
>  grub2-tools-extra             x86_64 1:2.02-106.el8               baseos    1.1 M
>  grub2-tools-minimal           x86_64 1:2.02-106.el8               baseos    210 k
>  grubby                        x86_64 8.40-42.el8                  baseos     49 k
>  gsettings-desktop-schemas     x86_64 3.32.0-6.el8                 baseos    633 k
>  hdparm                        x86_64 9.54-4.el8                   baseos    100 k
>  hwdata                        noarch 0.314-8.10.el8               baseos    1.7 M
>  iproute                       x86_64 5.12.0-4.el8                 baseos    775 k
>  iptables                      x86_64 1.8.4-20.el8                 baseos    585 k
>  iptables-ebtables             x86_64 1.8.4-20.el8                 baseos     72 k
>  iptables-libs                 x86_64 1.8.4-20.el8                 baseos    107 k
>  iscsi-initiator-utils         x86_64 6.2.1.4-4.git095f59c.el8     baseos    378 k
>  iscsi-initiator-utils-iscsiuio
>                                x86_64 6.2.1.4-4.git095f59c.el8     baseos    100 k
>  iwl100-firmware               noarch 39.31.5.1-103.el8.1          baseos    172 k
>  iwl1000-firmware              noarch 1:39.31.5.1-103.el8.1        baseos    235 k
>  iwl105-firmware               noarch 18.168.6.1-103.el8.1         baseos    256 k
>  iwl135-firmware               noarch 18.168.6.1-103.el8.1         baseos    265 k
>  iwl2000-firmware              noarch 18.168.6.1-103.el8.1         baseos    259 k
>  iwl2030-firmware              noarch 18.168.6.1-103.el8.1         baseos    268 k
>  iwl3160-firmware              noarch 1:25.30.13.0-103.el8.1       baseos    1.7 M
>  iwl5000-firmware              noarch 8.83.5.1_1-103.el8.1         baseos    316 k
>  iwl5150-firmware              noarch 8.24.2.2-103.el8.1           baseos    169 k
>  iwl6000-firmware              noarch 9.221.4.1-103.el8.1          baseos    189 k
>  iwl6000g2a-firmware           noarch 18.168.6.1-103.el8.1         baseos    332 k
>  iwl6000g2b-firmware           noarch 18.168.6.1-103.el8.1         baseos    332 k
>  iwl6050-firmware              noarch 41.28.5.1-103.el8.1          baseos    265 k
>  iwl7260-firmware              noarch 1:25.30.13.0-103.el8.1       baseos     18 M
>  json-c                        x86_64 0.13.1-2.el8                 baseos     40 k
>  kernel-tools                  x86_64 4.18.0-348.7.1.el8_5         baseos    7.2 M
>  kernel-tools-libs             x86_64 4.18.0-348.7.1.el8_5         baseos    7.0 M
>  kexec-tools                   x86_64 2.0.20-57.el8_5.1            baseos    514 k
>  keyutils-libs                 x86_64 1.5.10-9.el8                 baseos     34 k
>  kmod                          x86_64 25-18.el8                    baseos    126 k
>  kmod-kvdo                     x86_64 6.2.5.72-81.el8              baseos    341 k
>  kmod-libs                     x86_64 25-18.el8                    baseos     68 k
>  kpartx                        x86_64 0.8.4-17.el8                 baseos    113 k
>  kpatch                        noarch 0.9.2-5.el8                  baseos     17 k
>  kpatch-dnf                    noarch 0.2-5.el8                    baseos     17 k
>  krb5-libs                     x86_64 1.18.2-14.el8                baseos    840 k
>  libX11                        x86_64 1.6.8-5.el8                  appstream 611 k
>  libX11-common                 noarch 1.6.8-5.el8                  appstream 158 k
>  libblkid                      x86_64 2.32.1-28.el8                baseos    217 k
>  libblockdev                   x86_64 2.24-7.el8                   appstream 131 k
>  libblockdev-crypto            x86_64 2.24-7.el8                   appstream  81 k
>  libblockdev-fs                x86_64 2.24-7.el8                   appstream  86 k
>  libblockdev-loop              x86_64 2.24-7.el8                   appstream  70 k
>  libblockdev-lvm               x86_64 2.24-7.el8                   appstream  86 k
>  libblockdev-mdraid            x86_64 2.24-7.el8                   appstream  76 k
>  libblockdev-part              x86_64 2.24-7.el8                   appstream  80 k
>  libblockdev-swap              x86_64 2.24-7.el8                   appstream  72 k
>  libblockdev-utils             x86_64 2.24-7.el8                   appstream  80 k
>  libcap                        x86_64 2.26-5.el8                   baseos     60 k
>  libcap-ng                     x86_64 0.7.11-1.el8                 baseos     33 k
>  libcom_err                    x86_64 1.45.6-2.el8                 baseos     49 k
>  libcomps                      x86_64 0.1.16-2.el8                 baseos     82 k
>  libcurl                       x86_64 7.61.1-22.el8                baseos    301 k
>  libdb                         x86_64 5.3.28-42.el8_4              baseos    751 k
>  libdb-utils                   x86_64 5.3.28-42.el8_4              baseos    150 k
>  libdnf                        x86_64 0.63.0-3.el8                 baseos    700 k
>  libdrm                        x86_64 2.4.106-2.el8                appstream 167 k
>  libertas-usb8388-firmware     noarch 2:20210702-103.gitd79c2677.el8
>                                                                    baseos    135 k
>  libfastjson                   x86_64 0.99.9-1.el8                 appstream  38 k
>  libfdisk                      x86_64 2.32.1-28.el8                baseos    251 k
>  libgcc                        x86_64 8.5.0-4.el8_5                baseos     79 k
>  libgcrypt                     x86_64 1.8.5-6.el8                  baseos    463 k
>  libgomp                       x86_64 8.5.0-4.el8_5                baseos    206 k
>  libibverbs                    x86_64 35.0-1.el8                   baseos    335 k
>  libipa_hbac                   x86_64 2.5.2-2.el8_5.3              baseos    116 k
>  libldb                        x86_64 2.3.0-2.el8                  baseos    189 k
>  libmodulemd                   x86_64 2.13.0-1.el8                 baseos    233 k
>  libmount                      x86_64 2.32.1-28.el8                baseos    234 k
>  libndp                        x86_64 1.7-6.el8                    baseos     40 k
>  libnfsidmap                   x86_64 1:2.3.3-46.el8               baseos    121 k
>  librelp                       x86_64 1.9.0-1.el8                  appstream  76 k
>  librepo                       x86_64 1.14.0-2.el8                 baseos     93 k
>  libsepol                      x86_64 2.9-3.el8                    baseos    340 k
>  libslirp                      x86_64 4.4.0-1.module_el8.5.0+890+6b136101
>                                                                    appstream  70 k
>  libsmartcols                  x86_64 2.32.1-28.el8                baseos    177 k
>  libsmbclient                  x86_64 4.14.5-7.el8_5               baseos    148 k
>  libsolv                       x86_64 0.7.19-1.el8                 baseos    374 k
>  libss                         x86_64 1.45.6-2.el8                 baseos     54 k
>  libssh                        x86_64 0.9.4-3.el8                  baseos    215 k
>  libssh-config                 noarch 0.9.4-3.el8                  baseos     19 k
>  libsss_autofs                 x86_64 2.5.2-2.el8_5.3              baseos    118 k
>  libsss_certmap                x86_64 2.5.2-2.el8_5.3              baseos    155 k
>  libsss_idmap                  x86_64 2.5.2-2.el8_5.3              baseos    120 k
>  libsss_nss_idmap              x86_64 2.5.2-2.el8_5.3              baseos    127 k
>  libsss_sudo                   x86_64 2.5.2-2.el8_5.3              baseos    116 k
>  libstdc++                     x86_64 8.5.0-4.el8_5                baseos    453 k
>  libstoragemgmt                x86_64 1.9.1-1.el8                  baseos    246 k
>  libtalloc                     x86_64 2.3.2-1.el8                  baseos     49 k
>  libtevent                     x86_64 0.11.0-0.el8                 baseos     50 k
>  libtirpc                      x86_64 1.1.4-5.el8                  baseos    112 k
>  libudisks2                    x86_64 2.9.0-7.el8                  appstream 184 k
>  libuuid                       x86_64 2.32.1-28.el8                baseos     96 k
>  libwbclient                   x86_64 4.14.5-7.el8_5               baseos    121 k
>  libxcrypt                     x86_64 4.1.1-6.el8                  baseos     73 k
>  libxml2                       x86_64 2.9.7-9.el8_4.2              baseos    696 k
>  linux-firmware                noarch 20210702-103.gitd79c2677.el8 baseos    161 M
>  lshw                          x86_64 B.02.19.2-6.el8              baseos    341 k
>  lsscsi                        x86_64 0.32-3.el8                   baseos     71 k
>  lua-libs                      x86_64 5.3.4-12.el8                 baseos    118 k
>  lvm2                          x86_64 8:2.03.12-10.el8             baseos    1.6 M
>  lvm2-libs                     x86_64 8:2.03.12-10.el8             baseos    1.2 M
>  lz4-libs                      x86_64 1.8.3-3.el8_4                baseos     66 k
>  man-db                        x86_64 2.7.6.1-18.el8               baseos    887 k
>  man-pages-overrides           noarch 8.5.0.1-1.el8                appstream  98 k
>  mcelog                        x86_64 3:175-1.el8                  baseos     83 k
>  mdadm                         x86_64 4.2-rc2.el8                  baseos    460 k
>  microcode_ctl                 x86_64 4:20210608-1.el8             baseos    5.5 M
>  ncurses                       x86_64 6.1-9.20180224.el8           baseos    387 k
>  ncurses-base                  noarch 6.1-9.20180224.el8           baseos     81 k
>  ncurses-libs                  x86_64 6.1-9.20180224.el8           baseos    334 k
>  nettle                        x86_64 3.4.1-7.el8                  baseos    301 k
>  nftables                      x86_64 1:0.9.3-21.el8               baseos    321 k
>  nmap-ncat                     x86_64 2:7.70-6.el8                 appstream 237 k
>  nspr                          x86_64 4.32.0-1.el8_4               appstream 142 k
>  nss                           x86_64 3.67.0-7.el8_5               appstream 741 k
>  nss-softokn                   x86_64 3.67.0-7.el8_5               appstream 487 k
>  nss-softokn-freebl            x86_64 3.67.0-7.el8_5               appstream 395 k
>  nss-sysinit                   x86_64 3.67.0-7.el8_5               appstream  73 k
>  nss-util                      x86_64 3.67.0-7.el8_5               appstream 137 k
>  numactl-libs                  x86_64 2.0.12-13.el8                baseos     36 k
>  nvme-cli                      x86_64 1.14-3.el8                   baseos    463 k
>  open-vm-tools                 x86_64 11.2.5-2.el8                 appstream 771 k
>  openldap                      x86_64 2.4.46-18.el8                baseos    352 k
>  openssh                       x86_64 8.0p1-10.el8                 baseos    522 k
>  openssh-clients               x86_64 8.0p1-10.el8                 baseos    668 k
>  openssh-server                x86_64 8.0p1-10.el8                 baseos    485 k
>  openssl                       x86_64 1:1.1.1k-5.el8_5             baseos    709 k
>  openssl-libs                  x86_64 1:1.1.1k-5.el8_5             baseos    1.5 M
>  os-prober                     x86_64 1.74-9.el8                   baseos     51 k
>  pam                           x86_64 1.3.1-15.el8                 baseos    739 k
>  parted                        x86_64 3.2-39.el8                   baseos    555 k
>  pcre                          x86_64 8.42-6.el8                   baseos    211 k
>  platform-python               x86_64 3.6.8-41.el8                 baseos     85 k
>  platform-python-pip           noarch 9.0.3-20.el8                 baseos    1.7 M
>  plymouth                      x86_64 0.9.4-10.20200615git1e36e30.el8
>                                                                    appstream 127 k
>  plymouth-core-libs            x86_64 0.9.4-10.20200615git1e36e30.el8
>                                                                    appstream 122 k
>  plymouth-scripts              x86_64 0.9.4-10.20200615git1e36e30.el8
>                                                                    appstream  44 k
>  podman                        x86_64 3.3.1-9.module_el8.5.0+988+b1f0b741
>                                                                    appstream  12 M
>  podman-catatonit              x86_64 3.3.1-9.module_el8.5.0+988+b1f0b741
>                                                                    appstream 340 k
>  policycoreutils               x86_64 2.9-16.el8                   baseos    373 k
>  policycoreutils-python-utils  noarch 2.9-16.el8                   baseos    252 k
>  polkit                        x86_64 0.115-12.el8                 baseos    154 k
>  polkit-libs                   x86_64 0.115-12.el8                 baseos     76 k
>  python3-bind                  noarch 32:9.11.26-6.el8             appstream 150 k
>  python3-dnf                   noarch 4.7.0-4.el8                  baseos    545 k
>  python3-dnf-plugins-core      noarch 4.0.21-3.el8                 baseos    234 k
>  python3-firewall              noarch 0.9.3-7.el8                  baseos    432 k
>  python3-gpg                   x86_64 1.13.1-9.el8                 baseos    245 k
>  python3-hawkey                x86_64 0.63.0-3.el8                 baseos    116 k
>  python3-libcomps              x86_64 0.1.16-2.el8                 baseos     51 k
>  python3-libdnf                x86_64 0.63.0-3.el8                 baseos    777 k
>  python3-libs                  x86_64 3.6.8-41.el8                 baseos    7.8 M
>  python3-libstoragemgmt        x86_64 1.9.1-1.el8                  baseos    175 k
>      replacing  python3-libstoragemgmt-clibs.x86_64 1.8.7-1.el8
>  python3-libxml2               x86_64 2.9.7-9.el8_4.2              baseos    237 k
>  python3-lxml                  x86_64 4.2.3-3.el8                  appstream 1.5 M
>  python3-nftables              x86_64 1:0.9.3-21.el8               baseos     29 k
>  python3-perf                  x86_64 4.18.0-348.7.1.el8_5         baseos    7.1 M
>  python3-pip-wheel             noarch 9.0.3-20.el8                 baseos    1.0 M
>  python3-policycoreutils       noarch 2.9-16.el8                   baseos    2.2 M
>  python3-psutil                x86_64 5.4.3-11.el8                 appstream 373 k
>  python3-rpm                   x86_64 4.14.3-19.el8                baseos    154 k
>  python3-sssdconfig            noarch 2.5.2-2.el8_5.3              baseos    142 k
>  python3-syspurpose            x86_64 1.28.21-3.el8                baseos    312 k
>  python3-unbound               x86_64 1.7.3-17.el8                 appstream 119 k
>  quota                         x86_64 1:4.04-14.el8                baseos    214 k
>  quota-nls                     noarch 1:4.04-14.el8                baseos     95 k
>  rdma-core                     x86_64 35.0-1.el8                   baseos     59 k
>  realmd                        x86_64 0.16.3-23.el8                baseos    238 k
>  rpm                           x86_64 4.14.3-19.el8                baseos    543 k
>  rpm-build-libs                x86_64 4.14.3-19.el8                baseos    156 k
>  rpm-libs                      x86_64 4.14.3-19.el8                baseos    344 k
>  rpm-plugin-selinux            x86_64 4.14.3-19.el8                baseos     77 k
>  rpm-plugin-systemd-inhibit    x86_64 4.14.3-19.el8                baseos     78 k
>  rsyslog                       x86_64 8.2102.0-5.el8               appstream 752 k
>  rsyslog-gnutls                x86_64 8.2102.0-5.el8               appstream  32 k
>  rsyslog-gssapi                x86_64 8.2102.0-5.el8               appstream  34 k
>  rsyslog-relp                  x86_64 8.2102.0-5.el8               appstream  33 k
>  runc                          x86_64 1.0.2-1.module_el8.5.0+911+f19012f9
>                                                                    appstream 3.1 M
>  samba-client-libs             x86_64 4.14.5-7.el8_5               baseos    5.4 M
>  samba-common                  noarch 4.14.5-7.el8_5               baseos    221 k
>  samba-common-libs             x86_64 4.14.5-7.el8_5               baseos    173 k
>  selinux-policy                noarch 3.14.3-80.el8_5.2            baseos    636 k
>  selinux-policy-targeted       noarch 3.14.3-80.el8_5.2            baseos     15 M
>  setroubleshoot-plugins        noarch 3.3.14-1.el8                 appstream 358 k
>  setroubleshoot-server         x86_64 3.3.24-4.el8                 appstream 388 k
>  shadow-utils                  x86_64 2:4.6-14.el8                 baseos    1.2 M
>  slirp4netns                   x86_64 1.1.8-1.module_el8.5.0+890+6b136101
>                                                                    appstream  51 k
>  sos                           noarch 4.1-5.el8                    baseos    706 k
>  sqlite                        x86_64 3.26.0-15.el8                baseos    668 k
>  sqlite-libs                   x86_64 3.26.0-15.el8                baseos    581 k
>  sssd                          x86_64 2.5.2-2.el8_5.3              baseos    107 k
>  sssd-ad                       x86_64 2.5.2-2.el8_5.3              baseos    271 k
>  sssd-client                   x86_64 2.5.2-2.el8_5.3              baseos    205 k
>  sssd-common                   x86_64 2.5.2-2.el8_5.3              baseos    1.6 M
>  sssd-common-pac               x86_64 2.5.2-2.el8_5.3              baseos    179 k
>  sssd-ipa                      x86_64 2.5.2-2.el8_5.3              baseos    348 k
>  sssd-kcm                      x86_64 2.5.2-2.el8_5.3              baseos    254 k
>  sssd-krb5                     x86_64 2.5.2-2.el8_5.3              baseos    150 k
>  sssd-krb5-common              x86_64 2.5.2-2.el8_5.3              baseos    186 k
>  sssd-ldap                     x86_64 2.5.2-2.el8_5.3              baseos    209 k
>  sssd-nfs-idmap                x86_64 2.5.2-2.el8_5.3              baseos    116 k
>  sssd-proxy                    x86_64 2.5.2-2.el8_5.3              baseos    148 k
>  strace                        x86_64 5.7-3.el8                    baseos    1.1 M
>  sudo                          x86_64 1.8.29-7.el8_4.1             baseos    925 k
>  systemd                       x86_64 239-51.el8_5.2               baseos    3.6 M
>  systemd-libs                  x86_64 239-51.el8_5.2               baseos    1.1 M
>  systemd-pam                   x86_64 239-51.el8_5.2               baseos    477 k
>  systemd-udev                  x86_64 239-51.el8_5.2               baseos    1.6 M
>  tcpdump                       x86_64 14:4.9.3-2.el8               appstream 452 k
>  tpm2-tools                    x86_64 4.1.1-5.el8                  baseos    1.0 M
>  tpm2-tss                      x86_64 2.3.2-4.el8                  baseos    275 k
>  tuned                         noarch 2.16.0-1.el8                 baseos    312 k
>  tzdata                        noarch 2021e-1.el8                  baseos    474 k
>  udisks2                       x86_64 2.9.0-7.el8                  appstream 474 k
>  udisks2-iscsi                 x86_64 2.9.0-7.el8                  appstream  30 k
>  udisks2-lvm2                  x86_64 2.9.0-7.el8                  appstream  45 k
>  unbound-libs                  x86_64 1.7.3-17.el8                 appstream 503 k
>  unzip                         x86_64 6.0-45.el8_4                 baseos    195 k
>  util-linux                    x86_64 2.32.1-28.el8                baseos    2.5 M
>  util-linux-user               x86_64 2.32.1-28.el8                baseos    100 k
>  vdo                           x86_64 6.2.5.74-14.el8              baseos    663 k
>  vim-common                    x86_64 2:8.0.1763-16.el8            appstream 6.3 M
>  vim-enhanced                  x86_64 2:8.0.1763-16.el8            appstream 1.4 M
>  vim-filesystem                noarch 2:8.0.1763-16.el8            appstream  49 k
>  vim-minimal                   x86_64 2:8.0.1763-16.el8            baseos    573 k
>  virt-what                     x86_64 1.18-12.el8                  baseos     36 k
>  which                         x86_64 2.21-16.el8                  baseos     49 k
>  xfsprogs                      x86_64 5.0.0-9.el8                  baseos    1.1 M
>  yum                           noarch 4.7.0-4.el8                  baseos    205 k
> Installing dependencies:
>  grub2-tools-efi               x86_64 1:2.02-106.el8               baseos    474 k
>  libbpf                        x86_64 0.4.0-1.el8                  baseos    110 k
>
> Transaction Summary
> ===================================================================================
> Install    5 Packages
> Upgrade  324 Packages
>
> Total download size: 500 M
> Is this ok [y/N]: y
> Downloading Packages:
> (1/329): grub2-tools-efi-2.02-106.el8.x86_64.rpm    96 kB/s | 474 kB     00:04
> (2/329): kernel-4.18.0-348.7.1.el8_5.x86_64.rpm    392 kB/s | 7.0 MB     00:18
> (3/329): libbpf-0.4.0-1.el8.x86_64.rpm             122 kB/s | 110 kB     00:00
> (4/329): authselect-compat-1.2.2-3.el8.x86_64.rpm   60 kB/s |  38 kB     00:00
> (5/329): bind-libs-9.11.26-6.el8.x86_64.rpm        144 kB/s | 174 kB     00:01
> (6/329): bind-libs-lite-9.11.26-6.el8.x86_64.rpm   336 kB/s | 1.2 MB     00:03
> (7/329): bind-license-9.11.26-6.el8.noarch.rpm     115 kB/s | 102 kB     00:00
> (8/329): bind-utils-9.11.26-6.el8.x86_64.rpm       239 kB/s | 451 kB     00:01
> (9/329): kernel-core-4.18.0-348.7.1.el8_5.x86_64.r 1.2 MB/s |  38 MB     00:31
> (10/329): kernel-modules-4.18.0-348.7.1.el8_5.x86_ 1.1 MB/s |  30 MB     00:26
> (11/329): cockpit-packagekit-251.1-1.el8.noarch.rp 654 kB/s | 703 kB     00:01
> (12/329): cockpit-podman-33-1.module_el8.5.0+890+6 348 kB/s | 437 kB     00:01
> (13/329): conmon-2.0.29-1.module_el8.5.0+890+6b136  80 kB/s |  52 kB     00:00
> (14/329): cockpit-storaged-251.1-1.el8.noarch.rpm  624 kB/s | 655 kB     00:01
> (15/329): container-selinux-2.167.0-1.module_el8.5  88 kB/s |  54 kB     00:00
> (16/329): containers-common-1-2.module_el8.5.0+890 126 kB/s |  79 kB     00:00
> (17/329): criu-3.15-3.module_el8.5.0+890+6b136101. 407 kB/s | 518 kB     00:01
> (18/329): fstrm-0.6.1-2.el8.x86_64.rpm              49 kB/s |  29 kB     00:00
> (19/329): fuse-overlayfs-1.7.1-1.module_el8.5.0+89 120 kB/s |  73 kB     00:00
> (20/329): buildah-1.22.3-2.module_el8.5.0+911+f190 721 kB/s | 7.7 MB     00:10
> (21/329): libX11-common-1.6.8-5.el8.noarch.rpm     259 kB/s | 158 kB     00:00
> (22/329): libX11-1.6.8-5.el8.x86_64.rpm            390 kB/s | 611 kB     00:01
> (23/329): libblockdev-2.24-7.el8.x86_64.rpm        217 kB/s | 131 kB     00:00
> (24/329): libblockdev-crypto-2.24-7.el8.x86_64.rpm 135 kB/s |  81 kB     00:00
> (25/329): libblockdev-fs-2.24-7.el8.x86_64.rpm     144 kB/s |  86 kB     00:00
> (26/329): libblockdev-loop-2.24-7.el8.x86_64.rpm   117 kB/s |  70 kB     00:00
> (27/329): libblockdev-lvm-2.24-7.el8.x86_64.rpm    142 kB/s |  86 kB     00:00
> (28/329): libblockdev-mdraid-2.24-7.el8.x86_64.rpm  85 kB/s |  76 kB     00:00
> (29/329): libblockdev-part-2.24-7.el8.x86_64.rpm   134 kB/s |  80 kB     00:00
> (30/329): libblockdev-swap-2.24-7.el8.x86_64.rpm   122 kB/s |  72 kB     00:00
> (31/329): libblockdev-utils-2.24-7.el8.x86_64.rpm   89 kB/s |  80 kB     00:00
> (32/329): libdrm-2.4.106-2.el8.x86_64.rpm          186 kB/s | 167 kB     00:00
> (33/329): libfastjson-0.99.9-1.el8.x86_64.rpm       63 kB/s |  38 kB     00:00
> (34/329): librelp-1.9.0-1.el8.x86_64.rpm           127 kB/s |  76 kB     00:00
> (35/329): libslirp-4.4.0-1.module_el8.5.0+890+6b13  78 kB/s |  70 kB     00:00
> (36/329): libudisks2-2.9.0-7.el8.x86_64.rpm        200 kB/s | 184 kB     00:00
> (37/329): man-pages-overrides-8.5.0.1-1.el8.noarch 109 kB/s |  98 kB     00:00
> (38/329): nmap-ncat-7.70-6.el8.x86_64.rpm          265 kB/s | 237 kB     00:00
> (39/329): nspr-4.32.0-1.el8_4.x86_64.rpm           138 kB/s | 142 kB     00:01
> (40/329): nss-3.67.0-7.el8_5.x86_64.rpm            531 kB/s | 741 kB     00:01
> (41/329): nss-softokn-3.67.0-7.el8_5.x86_64.rpm    384 kB/s | 487 kB     00:01
> (42/329): nss-softokn-freebl-3.67.0-7.el8_5.x86_64 430 kB/s | 395 kB     00:00
> (43/329): nss-sysinit-3.67.0-7.el8_5.x86_64.rpm    121 kB/s |  73 kB     00:00
> (44/329): nss-util-3.67.0-7.el8_5.x86_64.rpm       227 kB/s | 137 kB     00:00
> (45/329): plymouth-0.9.4-10.20200615git1e36e30.el8 209 kB/s | 127 kB     00:00
> (46/329): containernetworking-plugins-1.0.0-1.modu 1.2 MB/s |  19 MB     00:15
> (47/329): plymouth-core-libs-0.9.4-10.20200615git1 177 kB/s | 122 kB     00:00
> (48/329): open-vm-tools-11.2.5-2.el8.x86_64.rpm    421 kB/s | 771 kB     00:01
> (49/329): plymouth-scripts-0.9.4-10.20200615git1e3  74 kB/s |  44 kB     00:00
> (50/329): podman-catatonit-3.3.1-9.module_el8.5.0+ 272 kB/s | 340 kB     00:01
> (51/329): python3-bind-9.11.26-6.el8.noarch.rpm    162 kB/s | 150 kB     00:00
> (52/329): python3-psutil-5.4.3-11.el8.x86_64.rpm   304 kB/s | 373 kB     00:01
> (53/329): python3-lxml-4.2.3-3.el8.x86_64.rpm      903 kB/s | 1.5 MB     00:01
> (54/329): python3-unbound-1.7.3-17.el8.x86_64.rpm  192 kB/s | 119 kB     00:00
> (55/329): rsyslog-gnutls-8.2102.0-5.el8.x86_64.rpm  53 kB/s |  32 kB     00:00
> (56/329): rsyslog-8.2102.0-5.el8.x86_64.rpm        625 kB/s | 752 kB     00:01
> (57/329): rsyslog-gssapi-8.2102.0-5.el8.x86_64.rpm  56 kB/s |  34 kB     00:00
> (58/329): rsyslog-relp-8.2102.0-5.el8.x86_64.rpm    56 kB/s |  33 kB     00:00
> (59/329): setroubleshoot-plugins-3.3.14-1.el8.noar 306 kB/s | 358 kB     00:01
> (60/329): setroubleshoot-server-3.3.24-4.el8.x86_6 424 kB/s | 388 kB     00:00
> (61/329): slirp4netns-1.1.8-1.module_el8.5.0+890+6  84 kB/s |  51 kB     00:00
> (62/329): tcpdump-4.9.3-2.el8.x86_64.rpm           375 kB/s | 452 kB     00:01
> (63/329): runc-1.0.2-1.module_el8.5.0+911+f19012f9 667 kB/s | 3.1 MB     00:04
> (64/329): podman-3.3.1-9.module_el8.5.0+988+b1f0b7 1.3 MB/s |  12 MB     00:09
> (65/329): udisks2-2.9.0-7.el8.x86_64.rpm           492 kB/s | 474 kB     00:00
> (66/329): udisks2-iscsi-2.9.0-7.el8.x86_64.rpm      52 kB/s |  30 kB     00:00
> (67/329): udisks2-lvm2-2.9.0-7.el8.x86_64.rpm       77 kB/s |  45 kB     00:00
> (68/329): unbound-libs-1.7.3-17.el8.x86_64.rpm     540 kB/s | 503 kB     00:00
> (69/329): vim-filesystem-8.0.1763-16.el8.noarch.rp  82 kB/s |  49 kB     00:00
> (70/329): vim-enhanced-8.0.1763-16.el8.x86_64.rpm  485 kB/s | 1.4 MB     00:02
> (71/329): NetworkManager-config-server-1.32.10-4.e 214 kB/s | 131 kB     00:00
> (72/329): NetworkManager-1.32.10-4.el8.x86_64.rpm  1.0 MB/s | 2.6 MB     00:02
> (73/329): NetworkManager-team-1.32.10-4.el8.x86_64 239 kB/s | 148 kB     00:00
> (74/329): vim-common-8.0.1763-16.el8.x86_64.rpm    1.2 MB/s | 6.3 MB     00:05
> (75/329): NetworkManager-tui-1.32.10-4.el8.x86_64. 350 kB/s | 336 kB     00:00
> (76/329): adcli-0.8.2-12.el8.x86_64.rpm            193 kB/s | 118 kB     00:00
> (77/329): NetworkManager-libnm-1.32.10-4.el8.x86_6 638 kB/s | 1.8 MB     00:02
> (78/329): alsa-sof-firmware-1.8-1.el8.noarch.rpm   653 kB/s | 676 kB     00:01
> (79/329): authselect-1.2.2-3.el8.x86_64.rpm        116 kB/s | 133 kB     00:01
> (80/329): authselect-libs-1.2.2-3.el8.x86_64.rpm   363 kB/s | 222 kB     00:00
> (81/329): bind-export-libs-9.11.26-6.el8.x86_64.rp 704 kB/s | 1.1 MB     00:01
> (82/329): bash-4.4.20-2.el8.x86_64.rpm             764 kB/s | 1.5 MB     00:02
> (83/329): ca-certificates-2021.2.50-80.0.el8_4.noa 209 kB/s | 390 kB     00:01
> (84/329): centos-gpg-keys-8-3.el8.noarch.rpm        21 kB/s |  12 kB     00:00
> (85/329): centos-linux-release-8.5-1.2111.el8.noar  36 kB/s |  22 kB     00:00
> (86/329): centos-linux-repos-8-3.el8.noarch.rpm     34 kB/s |  20 kB     00:00
> (87/329): centos-logos-85.8-2.el8.x86_64.rpm       357 kB/s | 700 kB     00:01
> (88/329): bpftool-4.18.0-348.7.1.el8_5.x86_64.rpm  1.2 MB/s | 7.7 MB     00:06
> (89/329): chkconfig-1.19.1-1.el8.x86_64.rpm        323 kB/s | 198 kB     00:00
> (90/329): chrony-4.1-1.el8.x86_64.rpm              367 kB/s | 327 kB     00:00
> (91/329): cockpit-251.1-1.el8.x86_64.rpm            86 kB/s |  78 kB     00:00
> (92/329): cockpit-bridge-251.1-1.el8.x86_64.rpm    579 kB/s | 539 kB     00:00
> (93/329): binutils-2.30-108.el8_5.1.x86_64.rpm     569 kB/s | 5.8 MB     00:10
> (94/329): cockpit-ws-251.1-1.el8.x86_64.rpm        889 kB/s | 1.3 MB     00:01
> (95/329): cockpit-system-251.1-1.el8.noarch.rpm    1.1 MB/s | 3.2 MB     00:02
> (96/329): crypto-policies-20210617-1.gitc776d3e.el 107 kB/s |  63 kB     00:00
> (97/329): crypto-policies-scripts-20210617-1.gitc7 130 kB/s |  83 kB     00:00
> (98/329): coreutils-8.30-12.el8.x86_64.rpm         498 kB/s | 1.2 MB     00:02
> (99/329): coreutils-common-8.30-12.el8.x86_64.rpm  1.0 MB/s | 2.0 MB     00:01
> (100/329): cups-libs-2.2.6-40.el8.x86_64.rpm       523 kB/s | 433 kB     00:00
> (101/329): dbus-1.12.8-14.el8.x86_64.rpm            46 kB/s |  41 kB     00:00
> (102/329): curl-7.61.1-22.el8.x86_64.rpm           285 kB/s | 351 kB     00:01
> (103/329): dbus-common-1.12.8-14.el8.noarch.rpm     78 kB/s |  46 kB     00:00
> (104/329): dbus-daemon-1.12.8-14.el8.x86_64.rpm    380 kB/s | 240 kB     00:00
> (105/329): dbus-libs-1.12.8-14.el8.x86_64.rpm      294 kB/s | 184 kB     00:00
> (106/329): dbus-tools-1.12.8-14.el8.x86_64.rpm     139 kB/s |  85 kB     00:00
> (107/329): device-mapper-event-1.02.177-10.el8.x86 415 kB/s | 271 kB     00:00
> (108/329): device-mapper-1.02.177-10.el8.x86_64.rp 452 kB/s | 377 kB     00:00
> (109/329): device-mapper-multipath-0.8.4-17.el8.x8 309 kB/s | 197 kB     00:00
> (110/329): device-mapper-libs-1.02.177-10.el8.x86_ 446 kB/s | 409 kB     00:00
> (111/329): device-mapper-event-libs-1.02.177-10.el 147 kB/s | 270 kB     00:01
> (112/329): device-mapper-multipath-libs-0.8.4-17.e 401 kB/s | 322 kB     00:00
> (113/329): dhcp-client-4.3.6-45.el8.x86_64.rpm     390 kB/s | 318 kB     00:00
> (114/329): device-mapper-persistent-data-0.9.0-4.e 732 kB/s | 925 kB     00:01
> (115/329): dhcp-common-4.3.6-45.el8.noarch.rpm     223 kB/s | 207 kB     00:00
> (116/329): dhcp-libs-4.3.6-45.el8.x86_64.rpm       241 kB/s | 148 kB     00:00
> (117/329): dmidecode-3.2-10.el8.x86_64.rpm         150 kB/s |  91 kB     00:00
> (118/329): dnf-plugins-core-4.0.21-3.el8.noarch.rp 116 kB/s |  70 kB     00:00
> (119/329): dnf-data-4.7.0-4.el8.noarch.rpm         173 kB/s | 154 kB     00:00
> (120/329): dnf-4.7.0-4.el8.noarch.rpm              354 kB/s | 544 kB     00:01
> (121/329): dracut-config-rescue-049-191.git2021092 101 kB/s |  61 kB     00:00
> (122/329): dracut-049-191.git20210920.el8.x86_64.r 415 kB/s | 374 kB     00:00
> (123/329): dracut-network-049-191.git20210920.el8. 117 kB/s | 108 kB     00:00
> (124/329): dracut-squash-049-191.git20210920.el8.x  87 kB/s |  61 kB     00:00
> (125/329): e2fsprogs-libs-1.45.6-2.el8.x86_64.rpm  379 kB/s | 233 kB     00:00
> (126/329): elfutils-debuginfod-client-0.185-1.el8. 108 kB/s |  66 kB     00:00
> (127/329): elfutils-default-yama-scope-0.185-1.el8  83 kB/s |  49 kB     00:00
> (128/329): e2fsprogs-1.45.6-2.el8.x86_64.rpm       565 kB/s | 1.0 MB     00:01
> (129/329): emacs-filesystem-26.1-7.el8.noarch.rpm  115 kB/s |  70 kB     00:00
> (130/329): elfutils-libelf-0.185-1.el8.x86_64.rpm  181 kB/s | 221 kB     00:01
> (131/329): elfutils-libs-0.185-1.el8.x86_64.rpm    316 kB/s | 292 kB     00:00
> (132/329): ethtool-5.8-7.el8.x86_64.rpm            335 kB/s | 209 kB     00:00
> (133/329): file-5.33-20.el8.x86_64.rpm             126 kB/s |  77 kB     00:00
> (134/329): file-libs-5.33-20.el8.x86_64.rpm        351 kB/s | 543 kB     00:01
> (135/329): filesystem-3.8-6.el8.x86_64.rpm         836 kB/s | 1.1 MB     00:01
> (136/329): firewalld-0.9.3-7.el8.noarch.rpm        378 kB/s | 502 kB     00:01
> (137/329): firewalld-filesystem-0.9.3-7.el8.noarch 129 kB/s |  77 kB     00:00
> (138/329): fontconfig-2.13.1-4.el8.x86_64.rpm      297 kB/s | 274 kB     00:00
> (139/329): glibc-common-2.28-164.el8.x86_64.rpm    853 kB/s | 1.3 MB     00:01
> (140/329): glibc-2.28-164.el8.x86_64.rpm           1.1 MB/s | 3.6 MB     00:03
> (141/329): glib2-2.56.4-156.el8.x86_64.rpm         671 kB/s | 2.5 MB     00:03
> (142/329): glibc-langpack-en-2.28-164.el8.x86_64.r 537 kB/s | 828 kB     00:01
> (143/329): gpgme-1.13.1-9.el8.x86_64.rpm           403 kB/s | 336 kB     00:00
> (144/329): gnutls-3.6.16-4.el8.x86_64.rpm          787 kB/s | 1.0 MB     00:01
> (145/329): grub2-pc-2.02-106.el8.x86_64.rpm         69 kB/s |  42 kB     00:00
> (146/329): grub2-common-2.02-106.el8.noarch.rpm    566 kB/s | 891 kB     00:01
> (147/329): grub2-pc-modules-2.02-106.el8.noarch.rp 723 kB/s | 916 kB     00:01
> (148/329): grub2-tools-minimal-2.02-106.el8.x86_64 320 kB/s | 210 kB     00:00
> (149/329): grub2-tools-extra-2.02-106.el8.x86_64.r 694 kB/s | 1.1 MB     00:01
> (150/329): grubby-8.40-42.el8.x86_64.rpm            82 kB/s |  49 kB     00:00
> (151/329): grub2-tools-2.02-106.el8.x86_64.rpm     813 kB/s | 2.0 MB     00:02
> (152/329): hdparm-9.54-4.el8.x86_64.rpm            163 kB/s | 100 kB     00:00
> (153/329): gsettings-desktop-schemas-3.32.0-6.el8. 508 kB/s | 633 kB     00:01
> (154/329): iproute-5.12.0-4.el8.x86_64.rpm         617 kB/s | 775 kB     00:01
> (155/329): iptables-1.8.4-20.el8.x86_64.rpm        405 kB/s | 585 kB     00:01
> (156/329): iptables-ebtables-1.8.4-20.el8.x86_64.r 107 kB/s |  72 kB     00:00
> (157/329): hwdata-0.314-8.10.el8.noarch.rpm        740 kB/s | 1.7 MB     00:02
> (158/329): iptables-libs-1.8.4-20.el8.x86_64.rpm   176 kB/s | 107 kB     00:00
> (159/329): iscsi-initiator-utils-iscsiuio-6.2.1.4- 111 kB/s | 100 kB     00:00
> (160/329): iscsi-initiator-utils-6.2.1.4-4.git095f 314 kB/s | 378 kB     00:01
> (161/329): iwl1000-firmware-39.31.5.1-103.el8.1.no 370 kB/s | 235 kB     00:00
> (162/329): iwl100-firmware-39.31.5.1-103.el8.1.noa 143 kB/s | 172 kB     00:01
> (163/329): iwl105-firmware-18.168.6.1-103.el8.1.no 279 kB/s | 256 kB     00:00
> (164/329): iwl135-firmware-18.168.6.1-103.el8.1.no 294 kB/s | 265 kB     00:00
> (165/329): iwl2000-firmware-18.168.6.1-103.el8.1.n 281 kB/s | 259 kB     00:00
> (166/329): iwl2030-firmware-18.168.6.1-103.el8.1.n 295 kB/s | 268 kB     00:00
> (167/329): iwl5000-firmware-8.83.5.1_1-103.el8.1.n 335 kB/s | 316 kB     00:00
> (168/329): iwl5150-firmware-8.24.2.2-103.el8.1.noa 272 kB/s | 169 kB     00:00
> (169/329): iwl6000-firmware-9.221.4.1-103.el8.1.no 303 kB/s | 189 kB     00:00
> (170/329): iwl3160-firmware-25.30.13.0-103.el8.1.n 985 kB/s | 1.7 MB     00:01
> (171/329): iwl6000g2a-firmware-18.168.6.1-103.el8. 276 kB/s | 332 kB     00:01
> (172/329): iwl6000g2b-firmware-18.168.6.1-103.el8. 366 kB/s | 332 kB     00:00
> (173/329): iwl6050-firmware-41.28.5.1-103.el8.1.no 293 kB/s | 265 kB     00:00
> (174/329): json-c-0.13.1-2.el8.x86_64.rpm           67 kB/s |  40 kB     00:00
> (175/329): kernel-tools-libs-4.18.0-348.7.1.el8_5. 1.0 MB/s | 7.0 MB     00:07
> (176/329): kernel-tools-4.18.0-348.7.1.el8_5.x86_6 841 kB/s | 7.2 MB     00:08
> (177/329): kexec-tools-2.0.20-57.el8_5.1.x86_64.rp 336 kB/s | 514 kB     00:01
> (178/329): keyutils-libs-1.5.10-9.el8.x86_64.rpm    56 kB/s |  34 kB     00:00
> (179/329): kmod-25-18.el8.x86_64.rpm               206 kB/s | 126 kB     00:00
> (180/329): kmod-libs-25-18.el8.x86_64.rpm          109 kB/s |  68 kB     00:00
> (181/329): kmod-kvdo-6.2.5.72-81.el8.x86_64.rpm    268 kB/s | 341 kB     00:01
> (182/329): kpartx-0.8.4-17.el8.x86_64.rpm          183 kB/s | 113 kB     00:00
> (183/329): kpatch-0.9.2-5.el8.noarch.rpm            28 kB/s |  17 kB     00:00
> (184/329): kpatch-dnf-0.2-5.el8.noarch.rpm          29 kB/s |  17 kB     00:00
> (185/329): libblkid-2.32.1-28.el8.x86_64.rpm       237 kB/s | 217 kB     00:00
> (186/329): libcap-2.26-5.el8.x86_64.rpm             99 kB/s |  60 kB     00:00
> (187/329): krb5-libs-1.18.2-14.el8.x86_64.rpm      392 kB/s | 840 kB     00:02
> (188/329): iwl7260-firmware-25.30.13.0-103.el8.1.n 1.3 MB/s |  18 MB     00:14
> (189/329): libcap-ng-0.7.11-1.el8.x86_64.rpm        55 kB/s |  33 kB     00:00
> (190/329): libcom_err-1.45.6-2.el8.x86_64.rpm       61 kB/s |  49 kB     00:00
> (191/329): libcomps-0.1.16-2.el8.x86_64.rpm         77 kB/s |  82 kB     00:01
> (192/329): libdb-5.3.28-42.el8_4.x86_64.rpm        470 kB/s | 751 kB     00:01
> (193/329): libdb-utils-5.3.28-42.el8_4.x86_64.rpm  160 kB/s | 150 kB     00:00
> (194/329): libertas-usb8388-firmware-20210702-103. 127 kB/s | 135 kB     00:01
> (195/329): libcurl-7.61.1-22.el8.x86_64.rpm         77 kB/s | 301 kB     00:03
> (196/329): libfdisk-2.32.1-28.el8.x86_64.rpm       261 kB/s | 251 kB     00:00
> (197/329): libdnf-0.63.0-3.el8.x86_64.rpm          305 kB/s | 700 kB     00:02
> (198/329): libgcc-8.5.0-4.el8_5.x86_64.rpm         131 kB/s |  79 kB     00:00
> (199/329): libgomp-8.5.0-4.el8_5.x86_64.rpm        331 kB/s | 206 kB     00:00
> (200/329): libgcrypt-1.8.5-6.el8.x86_64.rpm        376 kB/s | 463 kB     00:01
> (201/329): libipa_hbac-2.5.2-2.el8_5.3.x86_64.rpm  192 kB/s | 116 kB     00:00
> (202/329): libldb-2.3.0-2.el8.x86_64.rpm           208 kB/s | 189 kB     00:00
> (203/329): libibverbs-35.0-1.el8.x86_64.rpm        185 kB/s | 335 kB     00:01
> (204/329): libmodulemd-2.13.0-1.el8.x86_64.rpm     256 kB/s | 233 kB     00:00
> (205/329): libndp-1.7-6.el8.x86_64.rpm              67 kB/s |  40 kB     00:00
> (206/329): libmount-2.32.1-28.el8.x86_64.rpm       196 kB/s | 234 kB     00:01
> (207/329): librepo-1.14.0-2.el8.x86_64.rpm         153 kB/s |  93 kB     00:00
> (208/329): libnfsidmap-2.3.3-46.el8.x86_64.rpm      69 kB/s | 121 kB     00:01
> (209/329): libsepol-2.9-3.el8.x86_64.rpm           267 kB/s | 340 kB     00:01
> (210/329): libsmartcols-2.32.1-28.el8.x86_64.rpm   139 kB/s | 177 kB     00:01
> (211/329): libsmbclient-4.14.5-7.el8_5.x86_64.rpm  152 kB/s | 148 kB     00:00
> (212/329): libss-1.45.6-2.el8.x86_64.rpm            85 kB/s |  54 kB     00:00
> (213/329): libsolv-0.7.19-1.el8.x86_64.rpm         303 kB/s | 374 kB     00:01
> (214/329): libssh-config-0.9.4-3.el8.noarch.rpm     31 kB/s |  19 kB     00:00
> (215/329): libssh-0.9.4-3.el8.x86_64.rpm           234 kB/s | 215 kB     00:00
> (216/329): libsss_autofs-2.5.2-2.el8_5.3.x86_64.rp 130 kB/s | 118 kB     00:00
> (217/329): libsss_certmap-2.5.2-2.el8_5.3.x86_64.r 172 kB/s | 155 kB     00:00
> (218/329): libsss_idmap-2.5.2-2.el8_5.3.x86_64.rpm 100 kB/s | 120 kB     00:01
> (219/329): libsss_nss_idmap-2.5.2-2.el8_5.3.x86_64 140 kB/s | 127 kB     00:00
> (220/329): libsss_sudo-2.5.2-2.el8_5.3.x86_64.rpm  129 kB/s | 116 kB     00:00
> (221/329): libtalloc-2.3.2-1.el8.x86_64.rpm         81 kB/s |  49 kB     00:00
> (222/329): libstoragemgmt-1.9.1-1.el8.x86_64.rpm   205 kB/s | 246 kB     00:01
> (223/329): libstdc++-8.5.0-4.el8_5.x86_64.rpm      247 kB/s | 453 kB     00:01
> (224/329): libtevent-0.11.0-0.el8.x86_64.rpm        57 kB/s |  50 kB     00:00
> (225/329): libuuid-2.32.1-28.el8.x86_64.rpm        151 kB/s |  96 kB     00:00
> (226/329): libtirpc-1.1.4-5.el8.x86_64.rpm         123 kB/s | 112 kB     00:00
> (227/329): libwbclient-4.14.5-7.el8_5.x86_64.rpm   133 kB/s | 121 kB     00:00
> (228/329): libxcrypt-4.1.1-6.el8.x86_64.rpm        120 kB/s |  73 kB     00:00
> (229/329): libxml2-2.9.7-9.el8_4.2.x86_64.rpm      382 kB/s | 696 kB     00:01
> (230/329): lshw-B.02.19.2-6.el8.x86_64.rpm         222 kB/s | 341 kB     00:01
> (231/329): lsscsi-0.32-3.el8.x86_64.rpm            115 kB/s |  71 kB     00:00
> (232/329): lua-libs-5.3.4-12.el8.x86_64.rpm        130 kB/s | 118 kB     00:00
> (233/329): lvm2-2.03.12-10.el8.x86_64.rpm          508 kB/s | 1.6 MB     00:03
> (234/329): lvm2-libs-2.03.12-10.el8.x86_64.rpm     377 kB/s | 1.2 MB     00:03
> (235/329): lz4-libs-1.8.3-3.el8_4.x86_64.rpm       106 kB/s |  66 kB     00:00
> (236/329): mcelog-175-1.el8.x86_64.rpm             135 kB/s |  83 kB     00:00
> (237/329): man-db-2.7.6.1-18.el8.x86_64.rpm        482 kB/s | 887 kB     00:01
> (238/329): mdadm-4.2-rc2.el8.x86_64.rpm            301 kB/s | 460 kB     00:01
> (239/329): ncurses-6.1-9.20180224.el8.x86_64.rpm   314 kB/s | 387 kB     00:01
> (240/329): ncurses-base-6.1-9.20180224.el8.noarch. 133 kB/s |  81 kB     00:00
> (241/329): ncurses-libs-6.1-9.20180224.el8.x86_64. 268 kB/s | 334 kB     00:01
> (242/329): nettle-3.4.1-7.el8.x86_64.rpm           244 kB/s | 301 kB     00:01
> (243/329): nftables-0.9.3-21.el8.x86_64.rpm        257 kB/s | 321 kB     00:01
> (244/329): numactl-libs-2.0.12-13.el8.x86_64.rpm    60 kB/s |  36 kB     00:00
> (245/329): microcode_ctl-20210608-1.el8.x86_64.rpm 827 kB/s | 5.5 MB     00:06
> (246/329): nvme-cli-1.14-3.el8.x86_64.rpm          490 kB/s | 463 kB     00:00
> (247/329): openldap-2.4.46-18.el8.x86_64.rpm       244 kB/s | 352 kB     00:01
> (248/329): openssh-8.0p1-10.el8.x86_64.rpm         233 kB/s | 522 kB     00:02
> (249/329): openssh-clients-8.0p1-10.el8.x86_64.rpm 409 kB/s | 668 kB     00:01
> (250/329): openssh-server-8.0p1-10.el8.x86_64.rpm  395 kB/s | 485 kB     00:01
> (251/329): openssl-1.1.1k-5.el8_5.x86_64.rpm       320 kB/s | 709 kB     00:02
> (252/329): os-prober-1.74-9.el8.x86_64.rpm          85 kB/s |  51 kB     00:00
> (253/329): openssl-libs-1.1.1k-5.el8_5.x86_64.rpm  435 kB/s | 1.5 MB     00:03
> (254/329): pam-1.3.1-15.el8.x86_64.rpm             303 kB/s | 739 kB     00:02
> (255/329): parted-3.2-39.el8.x86_64.rpm            658 kB/s | 555 kB     00:00
> (256/329): platform-python-3.6.8-41.el8.x86_64.rpm 105 kB/s |  85 kB     00:00
> (257/329): pcre-8.42-6.el8.x86_64.rpm              175 kB/s | 211 kB     00:01
> (258/329): policycoreutils-2.9-16.el8.x86_64.rpm   322 kB/s | 373 kB     00:01
> (259/329): policycoreutils-python-utils-2.9-16.el8 413 kB/s | 252 kB     00:00
> (260/329): polkit-0.115-12.el8.x86_64.rpm          255 kB/s | 154 kB     00:00
> (261/329): polkit-libs-0.115-12.el8.x86_64.rpm     129 kB/s |  76 kB     00:00
> (262/329): platform-python-pip-9.0.3-20.el8.noarch 471 kB/s | 1.7 MB     00:03
> (263/329): python3-dnf-4.7.0-4.el8.noarch.rpm      448 kB/s | 545 kB     00:01
> (264/329): python3-dnf-plugins-core-4.0.21-3.el8.n 179 kB/s | 234 kB     00:01
> (265/329): python3-firewall-0.9.3-7.el8.noarch.rpm 472 kB/s | 432 kB     00:00
> (266/329): python3-hawkey-0.63.0-3.el8.x86_64.rpm  197 kB/s | 116 kB     00:00
> (267/329): python3-gpg-1.13.1-9.el8.x86_64.rpm     201 kB/s | 245 kB     00:01
> (268/329): python3-libcomps-0.1.16-2.el8.x86_64.rp  87 kB/s |  51 kB     00:00
> (269/329): python3-libdnf-0.63.0-3.el8.x86_64.rpm  357 kB/s | 777 kB     00:02
> (270/329): python3-libstoragemgmt-1.9.1-1.el8.x86_ 184 kB/s | 175 kB     00:00
> (271/329): python3-libxml2-2.9.7-9.el8_4.2.x86_64. 183 kB/s | 237 kB     00:01
> (272/329): python3-nftables-0.9.3-21.el8.x86_64.rp  44 kB/s |  29 kB     00:00
> (273/329): python3-libs-3.6.8-41.el8.x86_64.rpm    1.2 MB/s | 7.8 MB     00:06
> (274/329): python3-pip-wheel-9.0.3-20.el8.noarch.r 836 kB/s | 1.0 MB     00:01
> (275/329): python3-policycoreutils-2.9-16.el8.noar 1.0 MB/s | 2.2 MB     00:02
> (276/329): python3-rpm-4.14.3-19.el8.x86_64.rpm    258 kB/s | 154 kB     00:00
> (277/329): python3-sssdconfig-2.5.2-2.el8_5.3.noar 230 kB/s | 142 kB     00:00
> (278/329): python3-syspurpose-1.28.21-3.el8.x86_64 128 kB/s | 312 kB     00:02
> (279/329): quota-4.04-14.el8.x86_64.rpm            180 kB/s | 214 kB     00:01
> (280/329): quota-nls-4.04-14.el8.noarch.rpm        156 kB/s |  95 kB     00:00
> (281/329): python3-perf-4.18.0-348.7.1.el8_5.x86_6 697 kB/s | 7.1 MB     00:10
> (282/329): rdma-core-35.0-1.el8.x86_64.rpm          99 kB/s |  59 kB     00:00
> (283/329): realmd-0.16.3-23.el8.x86_64.rpm         193 kB/s | 238 kB     00:01
> (284/329): rpm-build-libs-4.14.3-19.el8.x86_64.rpm 165 kB/s | 156 kB     00:00
> (285/329): rpm-4.14.3-19.el8.x86_64.rpm            311 kB/s | 543 kB     00:01
> (286/329): rpm-plugin-selinux-4.14.3-19.el8.x86_64 131 kB/s |  77 kB     00:00
> (287/329): rpm-libs-4.14.3-19.el8.x86_64.rpm       281 kB/s | 344 kB     00:01
> (288/329): rpm-plugin-systemd-inhibit-4.14.3-19.el  91 kB/s |  78 kB     00:00
> (289/329): samba-common-4.14.5-7.el8_5.noarch.rpm  244 kB/s | 221 kB     00:00
> (290/329): samba-common-libs-4.14.5-7.el8_5.x86_64 190 kB/s | 173 kB     00:00
> (291/329): selinux-policy-3.14.3-80.el8_5.2.noarch 418 kB/s | 636 kB     00:01
> (292/329): samba-client-libs-4.14.5-7.el8_5.x86_64 843 kB/s | 5.4 MB     00:06
> (293/329): shadow-utils-4.6-14.el8.x86_64.rpm      647 kB/s | 1.2 MB     00:01
> (294/329): sos-4.1-5.el8.noarch.rpm                441 kB/s | 706 kB     00:01
> (295/329): sqlite-3.26.0-15.el8.x86_64.rpm         426 kB/s | 668 kB     00:01
> (296/329): sqlite-libs-3.26.0-15.el8.x86_64.rpm    434 kB/s | 581 kB     00:01
> (297/329): sssd-2.5.2-2.el8_5.3.x86_64.rpm         174 kB/s | 107 kB     00:00
> (298/329): sssd-ad-2.5.2-2.el8_5.3.x86_64.rpm      278 kB/s | 271 kB     00:00
> (299/329): sssd-client-2.5.2-2.el8_5.3.x86_64.rpm  221 kB/s | 205 kB     00:00
> (300/329): selinux-policy-targeted-3.14.3-80.el8_5 1.1 MB/s |  15 MB     00:13
> (301/329): sssd-common-pac-2.5.2-2.el8_5.3.x86_64. 277 kB/s | 179 kB     00:00
> (302/329): sssd-ipa-2.5.2-2.el8_5.3.x86_64.rpm     387 kB/s | 348 kB     00:00
> (303/329): sssd-common-2.5.2-2.el8_5.3.x86_64.rpm  491 kB/s | 1.6 MB     00:03
> (304/329): sssd-kcm-2.5.2-2.el8_5.3.x86_64.rpm     397 kB/s | 254 kB     00:00
> (305/329): sssd-krb5-2.5.2-2.el8_5.3.x86_64.rpm    167 kB/s | 150 kB     00:00
> (306/329): sssd-krb5-common-2.5.2-2.el8_5.3.x86_64 151 kB/s | 186 kB     00:01
> (307/329): sssd-ldap-2.5.2-2.el8_5.3.x86_64.rpm    240 kB/s | 209 kB     00:00
> (308/329): sssd-nfs-idmap-2.5.2-2.el8_5.3.x86_64.r 123 kB/s | 116 kB     00:00
> (309/329): sssd-proxy-2.5.2-2.el8_5.3.x86_64.rpm   246 kB/s | 148 kB     00:00
> (310/329): sudo-1.8.29-7.el8_4.1.x86_64.rpm        564 kB/s | 925 kB     00:01
> (311/329): strace-5.7-3.el8.x86_64.rpm             535 kB/s | 1.1 MB     00:02
> (312/329): systemd-libs-239-51.el8_5.2.x86_64.rpm  499 kB/s | 1.1 MB     00:02
> (313/329): systemd-239-51.el8_5.2.x86_64.rpm       1.2 MB/s | 3.6 MB     00:03
> (314/329): systemd-pam-239-51.el8_5.2.x86_64.rpm   377 kB/s | 477 kB     00:01
> (315/329): systemd-udev-239-51.el8_5.2.x86_64.rpm  931 kB/s | 1.6 MB     00:01
> (316/329): tpm2-tss-2.3.2-4.el8.x86_64.rpm         446 kB/s | 275 kB     00:00
> (317/329): tpm2-tools-4.1.1-5.el8.x86_64.rpm       455 kB/s | 1.0 MB     00:02
> (318/329): tuned-2.16.0-1.el8.noarch.rpm           341 kB/s | 312 kB     00:00
> (319/329): unzip-6.0-45.el8_4.x86_64.rpm           189 kB/s | 195 kB     00:01
> (320/329): tzdata-2021e-1.el8.noarch.rpm           382 kB/s | 474 kB     00:01
> (321/329): util-linux-user-2.32.1-28.el8.x86_64.rp 169 kB/s | 100 kB     00:00
> (322/329): vdo-6.2.5.74-14.el8.x86_64.rpm          550 kB/s | 663 kB     00:01
> (323/329): vim-minimal-8.0.1763-16.el8.x86_64.rpm  594 kB/s | 573 kB     00:00
> (324/329): util-linux-2.32.1-28.el8.x86_64.rpm     749 kB/s | 2.5 MB     00:03
> (325/329): virt-what-1.18-12.el8.x86_64.rpm         63 kB/s |  36 kB     00:00
> (326/329): which-2.21-16.el8.x86_64.rpm             82 kB/s |  49 kB     00:00
> (327/329): yum-4.7.0-4.el8.noarch.rpm              225 kB/s | 205 kB     00:00
> (328/329): xfsprogs-5.0.0-9.el8.x86_64.rpm         735 kB/s | 1.1 MB     00:01
> (329/329): linux-firmware-20210702-103.gitd79c2677 1.2 MB/s | 161 MB     02:12
> -----------------------------------------------------------------------------------
> Total                                              1.9 MB/s | 500 MB     04:26
> Running transaction check
> Transaction check succeeded.
> Running transaction test
> Transaction test succeeded.
> Running transaction
>   Running scriptlet: filesystem-3.8-6.el8.x86_64                               1/1
>   Running scriptlet: kmod-kvdo-6.2.5.72-81.el8.x86_64                          1/1
>   Preparing        :                                                           1/1
>   Running scriptlet: libgcc-8.5.0-4.el8_5.x86_64                               1/1
>   Upgrading        : libgcc-8.5.0-4.el8_5.x86_64                             1/654
>   Running scriptlet: libgcc-8.5.0-4.el8_5.x86_64                             1/654
>   Upgrading        : bind-license-32:9.11.26-6.el8.noarch                    2/654
>   Upgrading        : python3-pip-wheel-9.0.3-20.el8.noarch                   3/654
>   Upgrading        : filesystem-3.8-6.el8.x86_64                             4/654
>   Upgrading        : tzdata-2021e-1.el8.noarch                               5/654
>   Upgrading        : quota-nls-1:4.04-14.el8.noarch                          6/654
>   Upgrading        : ncurses-base-6.1-9.20180224.el8.noarch                  7/654
>   Upgrading        : ncurses-libs-6.1-9.20180224.el8.x86_64                  8/654
>   Running scriptlet: glibc-2.28-164.el8.x86_64                               9/654
>   Upgrading        : glibc-2.28-164.el8.x86_64                               9/654
>   Running scriptlet: glibc-2.28-164.el8.x86_64                               9/654
>   Upgrading        : bash-4.4.20-2.el8.x86_64                               10/654
>   Running scriptlet: bash-4.4.20-2.el8.x86_64                               10/654
>   Upgrading        : glibc-common-2.28-164.el8.x86_64                       11/654
>   Upgrading        : glibc-langpack-en-2.28-164.el8.x86_64                  12/654
>   Upgrading        : libcom_err-1.45.6-2.el8.x86_64                         13/654
>   Running scriptlet: libcom_err-1.45.6-2.el8.x86_64                         13/654
>   Upgrading        : libcap-2.26-5.el8.x86_64                               14/654
>   Upgrading        : libtalloc-2.3.2-1.el8.x86_64                           15/654
>   Upgrading        : sqlite-libs-3.26.0-15.el8.x86_64                       16/654
>   Upgrading        : elfutils-libelf-0.185-1.el8.x86_64                     17/654
>   Upgrading        : libuuid-2.32.1-28.el8.x86_64                           18/654
>   Running scriptlet: libuuid-2.32.1-28.el8.x86_64                           18/654
>   Upgrading        : libxml2-2.9.7-9.el8_4.2.x86_64                         19/654
>   Upgrading        : libtevent-0.11.0-0.el8.x86_64                          20/654
>   Upgrading        : libxcrypt-4.1.1-6.el8.x86_64                           21/654
>   Upgrading        : libsepol-2.9-3.el8.x86_64                              22/654
>   Running scriptlet: libsepol-2.9-3.el8.x86_64                              22/654
>   Upgrading        : libstdc++-8.5.0-4.el8_5.x86_64                         23/654
>   Running scriptlet: libstdc++-8.5.0-4.el8_5.x86_64                         23/654
>   Upgrading        : gpgme-1.13.1-9.el8.x86_64                              24/654
>   Upgrading        : lua-libs-5.3.4-12.el8.x86_64                           25/654
>   Upgrading        : which-2.21-16.el8.x86_64                               26/654
>   Upgrading        : chkconfig-1.19.1-1.el8.x86_64                          27/654
>   Upgrading        : grub2-common-1:2.02-106.el8.noarch                     28/654
>   Upgrading        : nspr-4.32.0-1.el8_4.x86_64                             29/654
>   Running scriptlet: nspr-4.32.0-1.el8_4.x86_64                             29/654
>   Upgrading        : json-c-0.13.1-2.el8.x86_64                             30/654
>   Upgrading        : keyutils-libs-1.5.10-9.el8.x86_64                      31/654
>   Upgrading        : libsss_idmap-2.5.2-2.el8_5.3.x86_64                    32/654
>   Running scriptlet: libsss_idmap-2.5.2-2.el8_5.3.x86_64                    32/654
>   Upgrading        : nss-util-3.67.0-7.el8_5.x86_64                         33/654
>   Upgrading        : file-libs-5.33-20.el8.x86_64                           34/654
>   Upgrading        : file-5.33-20.el8.x86_64                                35/654
>   Upgrading        : libcap-ng-0.7.11-1.el8.x86_64                          36/654
>   Upgrading        : libsmartcols-2.32.1-28.el8.x86_64                      37/654
>   Running scriptlet: libsmartcols-2.32.1-28.el8.x86_64                      37/654
>   Upgrading        : fstrm-0.6.1-2.el8.x86_64                               38/654
>   Upgrading        : device-mapper-persistent-data-0.9.0-4.el8.x86_64       39/654
>   Upgrading        : e2fsprogs-libs-1.45.6-2.el8.x86_64                     40/654
>   Running scriptlet: e2fsprogs-libs-1.45.6-2.el8.x86_64                     40/654
>   Upgrading        : nettle-3.4.1-7.el8.x86_64                              41/654
>   Running scriptlet: nettle-3.4.1-7.el8.x86_64                              41/654
>   Upgrading        : dmidecode-1:3.2-10.el8.x86_64                          42/654
>   Upgrading        : ethtool-2:5.8-7.el8.x86_64                             43/654
>   Upgrading        : iptables-libs-1.8.4-20.el8.x86_64                      44/654
>   Running scriptlet: iptables-1.8.4-20.el8.x86_64                           45/654
>   Upgrading        : iptables-1.8.4-20.el8.x86_64                           45/654
>   Running scriptlet: iptables-1.8.4-20.el8.x86_64                           45/654
>   Upgrading        : nftables-1:0.9.3-21.el8.x86_64                         46/654
>   Running scriptlet: nftables-1:0.9.3-21.el8.x86_64                         46/654
>   Upgrading        : libgcrypt-1.8.5-6.el8.x86_64                           47/654
>   Running scriptlet: libgcrypt-1.8.5-6.el8.x86_64                           47/654
>   Upgrading        : lz4-libs-1.8.3-3.el8_4.x86_64                          48/654
>   Upgrading        : iptables-ebtables-1.8.4-20.el8.x86_64                  49/654
>   Running scriptlet: iptables-ebtables-1.8.4-20.el8.x86_64                  49/654
>   Upgrading        : nss-softokn-freebl-3.67.0-7.el8_5.x86_64               50/654
>   Upgrading        : nss-softokn-3.67.0-7.el8_5.x86_64                      51/654
>   Upgrading        : grub2-pc-modules-1:2.02-106.el8.noarch                 52/654
>   Upgrading        : libcomps-0.1.16-2.el8.x86_64                           53/654
>   Installing       : libbpf-0.4.0-1.el8.x86_64                              54/654
>   Upgrading        : sqlite-3.26.0-15.el8.x86_64                            55/654
>   Upgrading        : libss-1.45.6-2.el8.x86_64                              56/654
>   Running scriptlet: libss-1.45.6-2.el8.x86_64                              56/654
>   Upgrading        : coreutils-common-8.30-12.el8.x86_64                    57/654
>   Running scriptlet: coreutils-common-8.30-12.el8.x86_64                    57/654
>   Upgrading        : kernel-tools-libs-4.18.0-348.7.1.el8_5.x86_64          58/654
>   Running scriptlet: kernel-tools-libs-4.18.0-348.7.1.el8_5.x86_64          58/654
>   Upgrading        : containernetworking-plugins-1.0.0-1.module_el8.5.0+    59/654
>   Upgrading        : libdrm-2.4.106-2.el8.x86_64                            60/654
>   Upgrading        : libfastjson-0.99.9-1.el8.x86_64                        61/654
>   Running scriptlet: libfastjson-0.99.9-1.el8.x86_64                        61/654
>   Upgrading        : hdparm-9.54-4.el8.x86_64                               62/654
>   Upgrading        : libndp-1.7-6.el8.x86_64                                63/654
>   Running scriptlet: libndp-1.7-6.el8.x86_64                                63/654
>   Upgrading        : libsss_autofs-2.5.2-2.el8_5.3.x86_64                   64/654
>   Upgrading        : libsss_nss_idmap-2.5.2-2.el8_5.3.x86_64                65/654
>   Running scriptlet: libsss_nss_idmap-2.5.2-2.el8_5.3.x86_64                65/654
>   Upgrading        : libsss_sudo-2.5.2-2.el8_5.3.x86_64                     66/654
>   Running scriptlet: libsss_sudo-2.5.2-2.el8_5.3.x86_64                     66/654
>   Upgrading        : ncurses-6.1-9.20180224.el8.x86_64                      67/654
>   Upgrading        : numactl-libs-2.0.12-13.el8.x86_64                      68/654
>   Running scriptlet: numactl-libs-2.0.12-13.el8.x86_64                      68/654
>   Upgrading        : pcre-8.42-6.el8.x86_64                                 69/654
>   Running scriptlet: pcre-8.42-6.el8.x86_64                                 69/654
>   Upgrading        : vim-minimal-2:8.0.1763-16.el8.x86_64                   70/654
>   Upgrading        : linux-firmware-20210702-103.gitd79c2677.el8.noarch     71/654
>   Upgrading        : libssh-config-0.9.4-3.el8.noarch                       72/654
>   Upgrading        : hwdata-0.314-8.10.el8.noarch                           73/654
>   Upgrading        : firewalld-filesystem-0.9.3-7.el8.noarch                74/654
>   Upgrading        : dnf-data-4.7.0-4.el8.noarch                            75/654
>   Upgrading        : dhcp-common-12:4.3.6-45.el8.noarch                     76/654
>   Upgrading        : dbus-common-1:1.12.8-14.el8.noarch                     77/654
>   Upgrading        : grub2-tools-minimal-1:2.02-106.el8.x86_64              78/654
>   Upgrading        : platform-python-pip-9.0.3-20.el8.noarch                79/654
>   Upgrading        : python3-libs-3.6.8-41.el8.x86_64                       80/654
>   Upgrading        : libssh-0.9.4-3.el8.x86_64                              81/654
>   Upgrading        : openldap-2.4.46-18.el8.x86_64                          82/654
>   Upgrading        : platform-python-3.6.8-41.el8.x86_64                    83/654
>   Running scriptlet: platform-python-3.6.8-41.el8.x86_64                    83/654
>   Upgrading        : grubby-8.40-42.el8.x86_64                              84/654
>   Upgrading        : libdb-utils-5.3.28-42.el8_4.x86_64                     85/654
>   Upgrading        : curl-7.61.1-22.el8.x86_64                              86/654
>   Upgrading        : libcurl-7.61.1-22.el8.x86_64                           87/654
>   Upgrading        : crypto-policies-scripts-20210617-1.gitc776d3e.el8.n    88/654
>   Upgrading        : crypto-policies-20210617-1.gitc776d3e.el8.noarch       89/654
>   Running scriptlet: crypto-policies-20210617-1.gitc776d3e.el8.noarch       89/654
>   Upgrading        : gnutls-3.6.16-4.el8.x86_64                             90/654
>   Upgrading        : elfutils-default-yama-scope-0.185-1.el8.noarch         91/654
>   Running scriptlet: elfutils-default-yama-scope-0.185-1.el8.noarch         91/654
>   Upgrading        : device-mapper-8:1.02.177-10.el8.x86_64                 92/654
>   Upgrading        : kpartx-0.8.4-17.el8.x86_64                             93/654
>   Upgrading        : krb5-libs-1.18.2-14.el8.x86_64                         94/654
>   Upgrading        : libtirpc-1.1.4-5.el8.x86_64                            95/654
>   Running scriptlet: libtirpc-1.1.4-5.el8.x86_64                            95/654
>   Upgrading        : elfutils-debuginfod-client-0.185-1.el8.x86_64          96/654
>   Upgrading        : elfutils-libs-0.185-1.el8.x86_64                       97/654
>   Upgrading        : rpm-4.14.3-19.el8.x86_64                               98/654
>   Upgrading        : libfdisk-2.32.1-28.el8.x86_64                          99/654
>   Running scriptlet: libfdisk-2.32.1-28.el8.x86_64                          99/654
>   Upgrading        : libmount-2.32.1-28.el8.x86_64                         100/654
>   Running scriptlet: libmount-2.32.1-28.el8.x86_64                         100/654
>   Upgrading        : dbus-libs-1:1.12.8-14.el8.x86_64                      101/654
>   Running scriptlet: dbus-libs-1:1.12.8-14.el8.x86_64                      101/654
>   Upgrading        : dbus-tools-1:1.12.8-14.el8.x86_64                     102/654
>   Upgrading        : device-mapper-libs-8:1.02.177-10.el8.x86_64           103/654
>   Upgrading        : coreutils-8.30-12.el8.x86_64                          104/654
>   Upgrading        : systemd-libs-239-51.el8_5.2.x86_64                    105/654
>   Running scriptlet: systemd-libs-239-51.el8_5.2.x86_64                    105/654
>   Upgrading        : libblkid-2.32.1-28.el8.x86_64                         106/654
>   Running scriptlet: libblkid-2.32.1-28.el8.x86_64                         106/654
>   Upgrading        : shadow-utils-2:4.6-14.el8.x86_64                      107/654
>   Running scriptlet: ca-certificates-2021.2.50-80.0.el8_4.noarch           108/654
>   Upgrading        : ca-certificates-2021.2.50-80.0.el8_4.noarch           108/654
>   Running scriptlet: ca-certificates-2021.2.50-80.0.el8_4.noarch           108/654
>   Upgrading        : openssl-libs-1:1.1.1k-5.el8_5.x86_64                  109/654
>   Running scriptlet: openssl-libs-1:1.1.1k-5.el8_5.x86_64                  109/654
>   Upgrading        : kmod-libs-25-18.el8.x86_64                            110/654
>   Running scriptlet: kmod-libs-25-18.el8.x86_64                            110/654
>   Upgrading        : libdb-5.3.28-42.el8_4.x86_64                          111/654
>   Running scriptlet: libdb-5.3.28-42.el8_4.x86_64                          111/654
>   Upgrading        : pam-1.3.1-15.el8.x86_64                               112/654
>   Running scriptlet: pam-1.3.1-15.el8.x86_64                               112/654
>   Upgrading        : util-linux-2.32.1-28.el8.x86_64                       113/654
>   Running scriptlet: util-linux-2.32.1-28.el8.x86_64                       113/654
>   Upgrading        : rpm-libs-4.14.3-19.el8.x86_64                         114/654
>   Running scriptlet: rpm-libs-4.14.3-19.el8.x86_64                         114/654
>   Upgrading        : kmod-25-18.el8.x86_64                                 115/654
>   Running scriptlet: dbus-daemon-1:1.12.8-14.el8.x86_64                    116/654
>   Upgrading        : dbus-daemon-1:1.12.8-14.el8.x86_64                    116/654
>   Running scriptlet: dbus-daemon-1:1.12.8-14.el8.x86_64                    116/654
>   Upgrading        : dracut-049-191.git20210920.el8.x86_64                 117/654
>   Upgrading        : systemd-pam-239-51.el8_5.2.x86_64                     118/654
>   Upgrading        : os-prober-1.74-9.el8.x86_64                           119/654
>   Running scriptlet: grub2-tools-1:2.02-106.el8.x86_64                     120/654
>   Upgrading        : grub2-tools-1:2.02-106.el8.x86_64                     120/654
>   Running scriptlet: grub2-tools-1:2.02-106.el8.x86_64                     120/654
>   Upgrading        : dbus-1:1.12.8-14.el8.x86_64                           121/654
>   Running scriptlet: systemd-239-51.el8_5.2.x86_64                         122/654
>   Upgrading        : systemd-239-51.el8_5.2.x86_64                         122/654
>   Running scriptlet: systemd-239-51.el8_5.2.x86_64                         122/654
>   Upgrading        : systemd-udev-239-51.el8_5.2.x86_64                    123/654
>   Running scriptlet: systemd-udev-239-51.el8_5.2.x86_64                    123/654
>   Upgrading        : glib2-2.56.4-156.el8.x86_64                           124/654
>   Upgrading        : libldb-2.3.0-2.el8.x86_64                             125/654
>   Upgrading        : polkit-libs-0.115-12.el8.x86_64                       126/654
>   Running scriptlet: polkit-libs-0.115-12.el8.x86_64                       126/654
>   Running scriptlet: polkit-0.115-12.el8.x86_64                            127/654
>   Upgrading        : polkit-0.115-12.el8.x86_64                            127/654
>   Running scriptlet: polkit-0.115-12.el8.x86_64                            127/654
>   Upgrading        : cockpit-bridge-251.1-1.el8.x86_64                     128/654
>   Upgrading        : libmodulemd-2.13.0-1.el8.x86_64                       129/654
>   Upgrading        : policycoreutils-2.9-16.el8.x86_64                     130/654
>   Running scriptlet: policycoreutils-2.9-16.el8.x86_64                     130/654
>   Upgrading        : libsss_certmap-2.5.2-2.el8_5.3.x86_64                 131/654
>   Running scriptlet: libsss_certmap-2.5.2-2.el8_5.3.x86_64                 131/654
>   Upgrading        : device-mapper-event-libs-8:1.02.177-10.el8.x86_64     132/654
>   Upgrading        : librepo-1.14.0-2.el8.x86_64                           133/654
>   Installing       : kernel-core-4.18.0-348.7.1.el8_5.x86_64               134/654
>   Running scriptlet: kernel-core-4.18.0-348.7.1.el8_5.x86_64               134/654
>   Upgrading        : rsyslog-8.2102.0-5.el8.x86_64                         135/654
>   Running scriptlet: rsyslog-8.2102.0-5.el8.x86_64                         135/654
>   Running scriptlet: samba-common-4.14.5-7.el8_5.noarch                    136/654
>   Upgrading        : samba-common-4.14.5-7.el8_5.noarch                    136/654
>   Running scriptlet: samba-common-4.14.5-7.el8_5.noarch                    136/654
>   Upgrading        : libsolv-0.7.19-1.el8.x86_64                           137/654
>   Upgrading        : iproute-5.12.0-4.el8.x86_64                           138/654
>   Upgrading        : parted-3.2-39.el8.x86_64                              139/654
>   Running scriptlet: parted-3.2-39.el8.x86_64                              139/654
>   Upgrading        : libblockdev-utils-2.24-7.el8.x86_64                   140/654
>   Upgrading        : libdnf-0.63.0-3.el8.x86_64                            141/654
>   Upgrading        : python3-libdnf-0.63.0-3.el8.x86_64                    142/654
>   Upgrading        : python3-hawkey-0.63.0-3.el8.x86_64                    143/654
>   Upgrading        : rpm-plugin-selinux-4.14.3-19.el8.x86_64               144/654
>   Upgrading        : selinux-policy-3.14.3-80.el8_5.2.noarch               145/654
>   Running scriptlet: selinux-policy-3.14.3-80.el8_5.2.noarch               145/654
>   Running scriptlet: selinux-policy-targeted-3.14.3-80.el8_5.2.noarch      146/654
>   Upgrading        : selinux-policy-targeted-3.14.3-80.el8_5.2.noarch      146/654
>   Running scriptlet: selinux-policy-targeted-3.14.3-80.el8_5.2.noarch      146/654
>   Upgrading        : NetworkManager-libnm-1:1.32.10-4.el8.x86_64           147/654
>   Running scriptlet: NetworkManager-libnm-1:1.32.10-4.el8.x86_64           147/654
>   Running scriptlet: NetworkManager-1:1.32.10-4.el8.x86_64                 148/654
>   Upgrading        : NetworkManager-1:1.32.10-4.el8.x86_64                 148/654
>   Running scriptlet: NetworkManager-1:1.32.10-4.el8.x86_64                 148/654
>   Upgrading        : iscsi-initiator-utils-iscsiuio-6.2.1.4-4.git095f59c   149/654
>   Running scriptlet: iscsi-initiator-utils-iscsiuio-6.2.1.4-4.git095f59c   149/654
>   Upgrading        : iscsi-initiator-utils-6.2.1.4-4.git095f59c.el8.x86_   150/654
>   Running scriptlet: iscsi-initiator-utils-6.2.1.4-4.git095f59c.el8.x86_   150/654
>   Upgrading        : fuse-overlayfs-1.7.1-1.module_el8.5.0+890+6b136101.   151/654
>   Running scriptlet: fuse-overlayfs-1.7.1-1.module_el8.5.0+890+6b136101.   151/654
>   Running scriptlet: openssh-8.0p1-10.el8.x86_64                           152/654
>   Upgrading        : openssh-8.0p1-10.el8.x86_64                           152/654
>   Upgrading        : bind-libs-lite-32:9.11.26-6.el8.x86_64                153/654
>   Upgrading        : python3-libxml2-2.9.7-9.el8_4.2.x86_64                154/654
>   Upgrading        : sos-4.1-5.el8.noarch                                  155/654
>   Upgrading        : bind-libs-32:9.11.26-6.el8.x86_64                     156/654
>   Upgrading        : NetworkManager-team-1:1.32.10-4.el8.x86_64            157/654
>   Upgrading        : libblockdev-2.24-7.el8.x86_64                         158/654
>   Upgrading        : libblockdev-fs-2.24-7.el8.x86_64                      159/654
>   Upgrading        : libblockdev-loop-2.24-7.el8.x86_64                    160/654
>   Upgrading        : libblockdev-part-2.24-7.el8.x86_64                    161/654
>   Upgrading        : libblockdev-swap-2.24-7.el8.x86_64                    162/654
>   Installing       : kernel-modules-4.18.0-348.7.1.el8_5.x86_64            163/654
>   Running scriptlet: kernel-modules-4.18.0-348.7.1.el8_5.x86_64            163/654
>   Running scriptlet: kmod-kvdo-6.2.5.72-81.el8.x86_64                      164/654
>   Upgrading        : kmod-kvdo-6.2.5.72-81.el8.x86_64                      164/654
>   Running scriptlet: kmod-kvdo-6.2.5.72-81.el8.x86_64                      164/654
>   Upgrading        : device-mapper-event-8:1.02.177-10.el8.x86_64          165/654
>   Running scriptlet: device-mapper-event-8:1.02.177-10.el8.x86_64          165/654
>   Upgrading        : lvm2-libs-8:2.03.12-10.el8.x86_64                     166/654
>   Upgrading        : lvm2-8:2.03.12-10.el8.x86_64                          167/654
>   Running scriptlet: lvm2-8:2.03.12-10.el8.x86_64                          167/654
>   Upgrading        : libblockdev-lvm-2.24-7.el8.x86_64                     168/654
>   Upgrading        : python3-policycoreutils-2.9-16.el8.noarch             169/654
>   Upgrading        : policycoreutils-python-utils-2.9-16.el8.noarch        170/654
>   Running scriptlet: container-selinux-2:2.167.0-1.module_el8.5.0+911+f1   171/654
>   Upgrading        : container-selinux-2:2.167.0-1.module_el8.5.0+911+f1   171/654
>   Running scriptlet: container-selinux-2:2.167.0-1.module_el8.5.0+911+f1   171/654
>   Upgrading        : cockpit-packagekit-251.1-1.el8.noarch                 172/654
>   Upgrading        : conmon-2:2.0.29-1.module_el8.5.0+890+6b136101.x86_6   173/654
>   Upgrading        : libslirp-4.4.0-1.module_el8.5.0+890+6b136101.x86_64   174/654
>   Upgrading        : slirp4netns-1.1.8-1.module_el8.5.0+890+6b136101.x86   175/654
>   Upgrading        : libudisks2-2.9.0-7.el8.x86_64                         176/654
>   Upgrading        : libipa_hbac-2.5.2-2.el8_5.3.x86_64                    177/654
>   Running scriptlet: libipa_hbac-2.5.2-2.el8_5.3.x86_64                    177/654
>   Running scriptlet: libstoragemgmt-1.9.1-1.el8.x86_64                     178/654
>   Upgrading        : libstoragemgmt-1.9.1-1.el8.x86_64                     178/654
>   Running scriptlet: libstoragemgmt-1.9.1-1.el8.x86_64                     178/654
>   Upgrading        : python3-libstoragemgmt-1.9.1-1.el8.x86_64             179/654
>   Running scriptlet: unbound-libs-1.7.3-17.el8.x86_64                      180/654
>   Upgrading        : unbound-libs-1.7.3-17.el8.x86_64                      180/654
>   Running scriptlet: unbound-libs-1.7.3-17.el8.x86_64                      180/654
>   Upgrading        : python3-unbound-1.7.3-17.el8.x86_64                   181/654
>   Running scriptlet: authselect-libs-1.2.2-3.el8.x86_64                    182/654
>   Upgrading        : authselect-libs-1.2.2-3.el8.x86_64                    182/654
 >   Upgrading        : mdadm-4.2-rc2.el8.x86_64                              183/644
>   Running scriptlet: mdadm-4.2-rc2.el8.x86_64                              183/654
>   Upgrading        : libblockdev-mdraid-2.24-7.el8.x86_64                  184/654
>   Upgrading        : grub2-tools-extra-1:2.02-106.el8.x86_64               185/654
>   Upgrading        : dracut-squash-049-191.git20210920.el8.x86_64          186/654
>   Upgrading        : rpm-build-libs-4.14.3-19.el8.x86_64                   187/654
>   Running scriptlet: rpm-build-libs-4.14.3-19.el8.x86_64                   187/654
>   Upgrading        : python3-rpm-4.14.3-19.el8.x86_64                      188/654
>   Running scriptlet: setroubleshoot-server-3.3.24-4.el8.x86_64             189/654
>   Upgrading        : setroubleshoot-server-3.3.24-4.el8.x86_64             189/654
>   Running scriptlet: setroubleshoot-server-3.3.24-4.el8.x86_64             189/654
>   Upgrading        : setroubleshoot-plugins-3.3.14-1.el8.noarch            190/654
>   Upgrading        : rpm-plugin-systemd-inhibit-4.14.3-19.el8.x86_64       191/654
>   Upgrading        : virt-what-1.18-12.el8.x86_64                          192/654
>   Upgrading        : sssd-client-2.5.2-2.el8_5.3.x86_64                    193/654
>   Running scriptlet: sssd-client-2.5.2-2.el8_5.3.x86_64                    193/654
>   Upgrading        : sudo-1.8.29-7.el8_4.1.x86_64                          194/654
>   Running scriptlet: sudo-1.8.29-7.el8_4.1.x86_64                          194/654
>   Upgrading        : librelp-1.9.0-1.el8.x86_64                            195/654
>   Running scriptlet: librelp-1.9.0-1.el8.x86_64                            195/654
>   Upgrading        : bind-export-libs-32:9.11.26-6.el8.x86_64              196/654
>   Running scriptlet: bind-export-libs-32:9.11.26-6.el8.x86_64              196/654
>   Upgrading        : openssl-1:1.1.1k-5.el8_5.x86_64                       197/654
>   Running scriptlet: tpm2-tss-2.3.2-4.el8.x86_64                           198/654
>   Upgrading        : tpm2-tss-2.3.2-4.el8.x86_64                           198/654
>   Running scriptlet: tpm2-tss-2.3.2-4.el8.x86_64                           198/654
>   Upgrading        : e2fsprogs-1.45.6-2.el8.x86_64                         199/654
>   Upgrading        : xfsprogs-5.0.0-9.el8.x86_64                           200/654
>   Running scriptlet: xfsprogs-5.0.0-9.el8.x86_64                           200/654
>   Upgrading        : plymouth-core-libs-0.9.4-10.20200615git1e36e30.el8.   201/654
>   Running scriptlet: plymouth-core-libs-0.9.4-10.20200615git1e36e30.el8.   201/654
>   Upgrading        : plymouth-scripts-0.9.4-10.20200615git1e36e30.el8.x8   202/654
>   Upgrading        : plymouth-0.9.4-10.20200615git1e36e30.el8.x86_64       203/654
>   Upgrading        : device-mapper-multipath-libs-0.8.4-17.el8.x86_64      204/654
>   Running scriptlet: device-mapper-multipath-libs-0.8.4-17.el8.x86_64      204/654
>   Upgrading        : device-mapper-multipath-0.8.4-17.el8.x86_64           205/654
>   Running scriptlet: device-mapper-multipath-0.8.4-17.el8.x86_64           205/654
>   Upgrading        : dhcp-libs-12:4.3.6-45.el8.x86_64                      206/654
>   Upgrading        : dhcp-client-12:4.3.6-45.el8.x86_64                    207/654
>   Upgrading        : dracut-network-049-191.git20210920.el8.x86_64         208/654
>   Upgrading        : kexec-tools-2.0.20-57.el8_5.1.x86_64                  209/654
>   Running scriptlet: kexec-tools-2.0.20-57.el8_5.1.x86_64                  209/654
>   Upgrading        : cockpit-system-251.1-1.el8.noarch                     210/654
>   Upgrading        : rdma-core-35.0-1.el8.x86_64                           211/654
>   Running scriptlet: rdma-core-35.0-1.el8.x86_64                           211/654
>   Upgrading        : nss-3.67.0-7.el8_5.x86_64                             212/654
>   Upgrading        : nss-sysinit-3.67.0-7.el8_5.x86_64                     213/654
>   Upgrading        : libblockdev-crypto-2.24-7.el8.x86_64                  214/654
>   Upgrading        : udisks2-2.9.0-7.el8.x86_64                            215/654
>   Running scriptlet: udisks2-2.9.0-7.el8.x86_64                            215/654
>   Upgrading        : udisks2-iscsi-2.9.0-7.el8.x86_64                      216/654
>   Upgrading        : udisks2-lvm2-2.9.0-7.el8.x86_64                       217/654
>   Upgrading        : cockpit-storaged-251.1-1.el8.noarch                   218/654
>   Upgrading        : binutils-2.30-108.el8_5.1.x86_64                      219/654
>   Running scriptlet: binutils-2.30-108.el8_5.1.x86_64                      219/654
>   Upgrading        : centos-logos-85.8-2.el8.x86_64                        220/654
>   Running scriptlet: centos-logos-85.8-2.el8.x86_64                        220/654
>   Running scriptlet: cockpit-ws-251.1-1.el8.x86_64                         221/654
>   Upgrading        : cockpit-ws-251.1-1.el8.x86_64                         221/654
>   Running scriptlet: cockpit-ws-251.1-1.el8.x86_64                         221/654
>   Upgrading        : libnfsidmap-1:2.3.3-46.el8.x86_64                     222/654
>   Upgrading        : sssd-nfs-idmap-2.5.2-2.el8_5.3.x86_64                 223/654
>   Running scriptlet: sssd-common-2.5.2-2.el8_5.3.x86_64                    224/654
>   Upgrading        : sssd-common-2.5.2-2.el8_5.3.x86_64                    224/654
>   Running scriptlet: sssd-common-2.5.2-2.el8_5.3.x86_64                    224/654
>   Running scriptlet: sssd-krb5-common-2.5.2-2.el8_5.3.x86_64               225/654
>   Upgrading        : sssd-krb5-common-2.5.2-2.el8_5.3.x86_64               225/654
>   Upgrading        : sssd-krb5-2.5.2-2.el8_5.3.x86_64                      226/654
>   Upgrading        : sssd-ldap-2.5.2-2.el8_5.3.x86_64                      227/654
>   Running scriptlet: sssd-proxy-2.5.2-2.el8_5.3.x86_64                     228/654
>   Upgrading        : sssd-proxy-2.5.2-2.el8_5.3.x86_64                     228/654
>   Upgrading        : adcli-0.8.2-12.el8.x86_64                             229/654
>   Running scriptlet: adcli-0.8.2-12.el8.x86_64                             229/654
>   Upgrading        : cups-libs-1:2.2.6-40.el8.x86_64                       230/654
>   Upgrading        : libwbclient-4.14.5-7.el8_5.x86_64                     231/654
>   Upgrading        : samba-client-libs-4.14.5-7.el8_5.x86_64               232/654
>   Upgrading        : samba-common-libs-4.14.5-7.el8_5.x86_64               233/654
>   Upgrading        : sssd-common-pac-2.5.2-2.el8_5.3.x86_64                234/654
>   Upgrading        : libsmbclient-4.14.5-7.el8_5.x86_64                    235/654
>   Upgrading        : criu-3.15-3.module_el8.5.0+890+6b136101.x86_64        236/654
>   Upgrading        : runc-1.0.2-1.module_el8.5.0+911+f19012f9.x86_64       237/654
>   Upgrading        : python3-bind-32:9.11.26-6.el8.noarch                  238/654
>   Upgrading        : bind-utils-32:9.11.26-6.el8.x86_64                    239/654
>   Upgrading        : sssd-ad-2.5.2-2.el8_5.3.x86_64                        240/654
>   Running scriptlet: sssd-ipa-2.5.2-2.el8_5.3.x86_64                       241/654
>   Upgrading        : sssd-ipa-2.5.2-2.el8_5.3.x86_64                       241/654
>   Upgrading        : kernel-tools-4.18.0-348.7.1.el8_5.x86_64              242/654
>   Upgrading        : python3-gpg-1.13.1-9.el8.x86_64                       243/654
>   Upgrading        : python3-libcomps-0.1.16-2.el8.x86_64                  244/654
>   Upgrading        : python3-dnf-4.7.0-4.el8.noarch                        245/654
>   Upgrading        : dnf-4.7.0-4.el8.noarch                                246/654
>   Running scriptlet: dnf-4.7.0-4.el8.noarch                                246/654
>   Upgrading        : kpatch-dnf-0.2-5.el8.noarch                           247/654
>   Running scriptlet: kpatch-dnf-0.2-5.el8.noarch                           247/654
> To enable automatic kpatch-patch subscription, run:
>         $ dnf kpatch auto
>
>   Upgrading        : python3-dnf-plugins-core-4.0.21-3.el8.noarch          248/654
>   Upgrading        : python3-nftables-1:0.9.3-21.el8.x86_64                249/654
>   Upgrading        : python3-firewall-0.9.3-7.el8.noarch                   250/654
>   Upgrading        : python3-perf-4.18.0-348.7.1.el8_5.x86_64              251/654
>   Upgrading        : python3-sssdconfig-2.5.2-2.el8_5.3.noarch             252/654
>   Upgrading        : sssd-2.5.2-2.el8_5.3.x86_64                           253/654
>   Upgrading        : authselect-1.2.2-3.el8.x86_64                         254/654
>   Upgrading        : realmd-0.16.3-23.el8.x86_64                           255/654
>   Running scriptlet: realmd-0.16.3-23.el8.x86_64                           255/654
>   Upgrading        : python3-syspurpose-1.28.21-3.el8.x86_64               256/654
>   Upgrading        : centos-gpg-keys-1:8-3.el8.noarch                      257/654
>   Upgrading        : centos-linux-release-8.5-1.2111.el8.noarch            258/654
>   Upgrading        : centos-linux-repos-8-3.el8.noarch                     259/654
>   Upgrading        : containers-common-2:1-2.module_el8.5.0+890+6b136101   260/654
>   Upgrading        : podman-catatonit-3.3.1-9.module_el8.5.0+988+b1f0b74   261/654
>   Upgrading        : podman-3.3.1-9.module_el8.5.0+988+b1f0b741.x86_64     262/654
>   Upgrading        : vim-filesystem-2:8.0.1763-16.el8.noarch               263/654
>   Upgrading        : vim-common-2:8.0.1763-16.el8.x86_64                   264/654
>   Upgrading        : libX11-common-1.6.8-5.el8.noarch                      265/654
>   Upgrading        : libX11-1.6.8-5.el8.x86_64                             266/654
>   Upgrading        : vim-enhanced-2:8.0.1763-16.el8.x86_64                 267/654
>   Upgrading        : cockpit-podman-33-1.module_el8.5.0+890+6b136101.noa   268/654
>   Upgrading        : buildah-1.22.3-2.module_el8.5.0+911+f19012f9.x86_64   269/654
>   Upgrading        : tuned-2.16.0-1.el8.noarch                             270/654
>   Running scriptlet: tuned-2.16.0-1.el8.noarch                             270/654
>   Upgrading        : authselect-compat-1.2.2-3.el8.x86_64                  271/654
>   Upgrading        : firewalld-0.9.3-7.el8.noarch                          272/654
>   Running scriptlet: firewalld-0.9.3-7.el8.noarch                          272/654
>   Upgrading        : dnf-plugins-core-4.0.21-3.el8.noarch                  273/654
>   Upgrading        : kpatch-0.9.2-5.el8.noarch                             274/654
>   Upgrading        : yum-4.7.0-4.el8.noarch                                275/654
>   Upgrading        : sssd-kcm-2.5.2-2.el8_5.3.x86_64                       276/654
>   Running scriptlet: sssd-kcm-2.5.2-2.el8_5.3.x86_64                       276/654
>   Upgrading        : cockpit-251.1-1.el8.x86_64                            277/654
>   Upgrading        : libibverbs-35.0-1.el8.x86_64                          278/654
>   Running scriptlet: libibverbs-35.0-1.el8.x86_64                          278/654
>   Upgrading        : tpm2-tools-4.1.1-5.el8.x86_64                         279/654
>   Upgrading        : rsyslog-relp-8.2102.0-5.el8.x86_64                    280/654
>   Upgrading        : grub2-pc-1:2.02-106.el8.x86_64                        281/654
>   Upgrading        : vdo-6.2.5.74-14.el8.x86_64                            282/654
>   Running scriptlet: vdo-6.2.5.74-14.el8.x86_64                            282/654
>   Installing       : kernel-4.18.0-348.7.1.el8_5.x86_64                    283/654
>   Upgrading        : openssh-clients-8.0p1-10.el8.x86_64                   284/654
>   Running scriptlet: openssh-server-8.0p1-10.el8.x86_64                    285/654
>   Upgrading        : openssh-server-8.0p1-10.el8.x86_64                    285/654
>   Running scriptlet: openssh-server-8.0p1-10.el8.x86_64                    285/654
>   Upgrading        : NetworkManager-tui-1:1.32.10-4.el8.x86_64             286/654
>   Upgrading        : open-vm-tools-11.2.5-2.el8.x86_64                     287/654
>   Running scriptlet: open-vm-tools-11.2.5-2.el8.x86_64                     287/654
>   Upgrading        : rsyslog-gnutls-8.2102.0-5.el8.x86_64                  288/654
>   Upgrading        : rsyslog-gssapi-8.2102.0-5.el8.x86_64                  289/654
>   Upgrading        : gsettings-desktop-schemas-3.32.0-6.el8.x86_64         290/654
>   Running scriptlet: chrony-4.1-1.el8.x86_64                               291/654
>   Upgrading        : chrony-4.1-1.el8.x86_64                               291/654
>   Running scriptlet: chrony-4.1-1.el8.x86_64                               291/654
>   Upgrading        : mcelog-3:175-1.el8.x86_64                             292/654
>   Running scriptlet: mcelog-3:175-1.el8.x86_64                             292/654
>   Upgrading        : microcode_ctl-4:20210608-1.el8.x86_64                 293/654
>   Running scriptlet: microcode_ctl-4:20210608-1.el8.x86_64                 293/654
>   Installing       : grub2-tools-efi-1:2.02-106.el8.x86_64                 294/654
>   Upgrading        : dracut-config-rescue-049-191.git20210920.el8.x86_64   295/654
>   Upgrading        : nvme-cli-1.14-3.el8.x86_64                            296/654
>   Running scriptlet: nvme-cli-1.14-3.el8.x86_64                            296/654
>   Upgrading        : util-linux-user-2.32.1-28.el8.x86_64                  297/654
>   Upgrading        : nmap-ncat-2:7.70-6.el8.x86_64                         298/654
>   Running scriptlet: nmap-ncat-2:7.70-6.el8.x86_64                         298/654
>   Running scriptlet: tcpdump-14:4.9.3-2.el8.x86_64                         299/654
>   Upgrading        : tcpdump-14:4.9.3-2.el8.x86_64                         299/654
>   Upgrading        : fontconfig-2.13.1-4.el8.x86_64                        300/654
>   Running scriptlet: fontconfig-2.13.1-4.el8.x86_64                        300/654
>   Running scriptlet: man-db-2.7.6.1-18.el8.x86_64                          301/654
>   Upgrading        : man-db-2.7.6.1-18.el8.x86_64                          301/654
>   Running scriptlet: man-db-2.7.6.1-18.el8.x86_64                          301/654
>   Upgrading        : strace-5.7-3.el8.x86_64                               302/654
>   Upgrading        : quota-1:4.04-14.el8.x86_64                            303/654
>   Upgrading        : python3-lxml-4.2.3-3.el8.x86_64                       304/654
>   Upgrading        : python3-psutil-5.4.3-11.el8.x86_64                    305/654
>   Upgrading        : lshw-B.02.19.2-6.el8.x86_64                           306/654
>   Upgrading        : bpftool-4.18.0-348.7.1.el8_5.x86_64                   307/654
>   Upgrading        : libgomp-8.5.0-4.el8_5.x86_64                          308/654
>   Running scriptlet: libgomp-8.5.0-4.el8_5.x86_64                          308/654
>   Upgrading        : unzip-6.0-45.el8_4.x86_64                             309/654
>   Upgrading        : lsscsi-0.32-3.el8.x86_64                              310/654
>   Upgrading        : libertas-usb8388-firmware-2:20210702-103.gitd79c267   311/654
>   Upgrading        : iwl7260-firmware-1:25.30.13.0-103.el8.1.noarch        312/654
>   Upgrading        : iwl6050-firmware-41.28.5.1-103.el8.1.noarch           313/654
>   Upgrading        : iwl6000g2b-firmware-18.168.6.1-103.el8.1.noarch       314/654
>   Upgrading        : iwl6000g2a-firmware-18.168.6.1-103.el8.1.noarch       315/654
>   Upgrading        : iwl6000-firmware-9.221.4.1-103.el8.1.noarch           316/654
>   Upgrading        : iwl5150-firmware-8.24.2.2-103.el8.1.noarch            317/654
>   Upgrading        : iwl5000-firmware-8.83.5.1_1-103.el8.1.noarch          318/654
>   Upgrading        : iwl3160-firmware-1:25.30.13.0-103.el8.1.noarch        319/654
>   Upgrading        : iwl2030-firmware-18.168.6.1-103.el8.1.noarch          320/654
>   Upgrading        : iwl2000-firmware-18.168.6.1-103.el8.1.noarch          321/654
>   Upgrading        : iwl135-firmware-18.168.6.1-103.el8.1.noarch           322/654
>   Upgrading        : iwl105-firmware-18.168.6.1-103.el8.1.noarch           323/654
>   Upgrading        : iwl1000-firmware-1:39.31.5.1-103.el8.1.noarch         324/654
>   Upgrading        : iwl100-firmware-39.31.5.1-103.el8.1.noarch            325/654
>   Upgrading        : emacs-filesystem-1:26.1-7.el8.noarch                  326/654
>   Upgrading        : alsa-sof-firmware-1.8-1.el8.noarch                    327/654
>   Upgrading        : NetworkManager-config-server-1:1.32.10-4.el8.noarch   328/654
>   Upgrading        : man-pages-overrides-8.5.0.1-1.el8.noarch              329/654
>   Running scriptlet: tuned-2.15.0-2.el8.noarch                             330/654
>   Cleanup          : tuned-2.15.0-2.el8.noarch                             330/654
>   Running scriptlet: tuned-2.15.0-2.el8.noarch                             330/654
>   Running scriptlet: microcode_ctl-4:20210216-1.el8.x86_64                 331/654
>   Cleanup          : microcode_ctl-4:20210216-1.el8.x86_64                 331/654
>   Running scriptlet: microcode_ctl-4:20210216-1.el8.x86_64                 331/654
>   Running scriptlet: firewalld-0.8.2-6.el8.noarch                          332/654
>   Cleanup          : firewalld-0.8.2-6.el8.noarch                          332/654
>   Running scriptlet: firewalld-0.8.2-6.el8.noarch                          332/654
>   Cleanup          : grub2-pc-1:2.02-99.el8.x86_64                         333/654
>   Running scriptlet: iptables-ebtables-1.8.4-17.el8.x86_64                 334/654
>   Cleanup          : iptables-ebtables-1.8.4-17.el8.x86_64                 334/654
>   Running scriptlet: iptables-ebtables-1.8.4-17.el8.x86_64                 334/654
>   Cleanup          : kpatch-0.9.2-3.el8.noarch                             335/654
>   Cleanup          : cockpit-238.2-1.el8.x86_64                            336/654
>   Cleanup          : cockpit-storaged-238.2-1.el8.noarch                   337/654
>   Cleanup          : cockpit-system-238.2-1.el8.noarch                     338/654
>   Cleanup          : authselect-compat-1.2.2-2.el8.x86_64                  339/654
>   Running scriptlet: open-vm-tools-11.2.0-2.el8.x86_64                     340/654
>   Cleanup          : open-vm-tools-11.2.0-2.el8.x86_64                     340/654
>   Running scriptlet: open-vm-tools-11.2.0-2.el8.x86_64                     340/654
>   Running scriptlet: cockpit-ws-238.2-1.el8.x86_64                         341/654
>   Cleanup          : cockpit-ws-238.2-1.el8.x86_64                         341/654
>   Running scriptlet: cockpit-ws-238.2-1.el8.x86_64                         341/654
>   Running scriptlet: openssh-server-8.0p1-5.el8.x86_64                     342/654
>   Cleanup          : openssh-server-8.0p1-5.el8.x86_64                     342/654
>   Running scriptlet: openssh-server-8.0p1-5.el8.x86_64                     342/654
>   Running scriptlet: sssd-kcm-2.4.0-9.el8.x86_64                           343/654
>   Cleanup          : sssd-kcm-2.4.0-9.el8.x86_64                           343/654
>   Running scriptlet: sssd-kcm-2.4.0-9.el8.x86_64                           343/654
>   Cleanup          : NetworkManager-tui-1:1.30.0-7.el8.x86_64              344/654
>   Running scriptlet: kexec-tools-2.0.20-46.el8.x86_64                      345/654
>   Cleanup          : kexec-tools-2.0.20-46.el8.x86_64                      345/654
>   Running scriptlet: kexec-tools-2.0.20-46.el8.x86_64                      345/654
>   Obsoleting       : python3-libstoragemgmt-clibs-1.8.7-1.el8.x86_64       346/654
>   Running scriptlet: libstoragemgmt-1.8.7-1.el8.x86_64                     347/654
>   Cleanup          : libstoragemgmt-1.8.7-1.el8.x86_64                     347/654
>   Running scriptlet: libstoragemgmt-1.8.7-1.el8.x86_64                     347/654
>   Cleanup          : python3-libstoragemgmt-1.8.7-1.el8.noarch             348/654
>   Running scriptlet: device-mapper-multipath-0.8.4-10.el8.x86_64           349/654
>   Cleanup          : device-mapper-multipath-0.8.4-10.el8.x86_64           349/654
>   Running scriptlet: device-mapper-multipath-0.8.4-10.el8.x86_64           349/654
>   Running scriptlet: binutils-2.30-93.el8.x86_64                           350/654
>   Cleanup          : binutils-2.30-93.el8.x86_64                           350/654
>   Running scriptlet: binutils-2.30-93.el8.x86_64                           350/654
>   Running scriptlet: vdo-6.2.4.14-14.el8.x86_64                            351/654
>   Cleanup          : vdo-6.2.4.14-14.el8.x86_64                            351/654
>   Running scriptlet: vdo-6.2.4.14-14.el8.x86_64                            351/654
>   Cleanup          : python3-lxml-4.2.3-2.el8.x86_64                       352/654
>   Cleanup          : device-mapper-multipath-libs-0.8.4-10.el8.x86_64      353/654
>   Running scriptlet: device-mapper-multipath-libs-0.8.4-10.el8.x86_64      353/654
>   Cleanup          : sudo-1.8.29-7.el8.x86_64                              354/654
>   Running scriptlet: chrony-3.5-2.el8.x86_64                               355/654
>   Cleanup          : chrony-3.5-2.el8.x86_64                               355/654
>   Running scriptlet: chrony-3.5-2.el8.x86_64                               355/654
>   Cleanup          : openssh-clients-8.0p1-5.el8.x86_64                    356/654
>   Cleanup          : realmd-0.16.3-22.el8.x86_64                           357/654
>   Cleanup          : kernel-tools-4.18.0-305.3.1.el8.x86_64                358/654
>   Cleanup          : buildah-1.19.7-1.module_el8.4.0+781+acf4c33b.x86_64   359/654
>   Cleanup          : dracut-network-049-135.git20210121.el8.x86_64         360/654
>   Cleanup          : dhcp-client-12:4.3.6-44.0.1.el8.x86_64                361/654
>   Cleanup          : dhcp-libs-12:4.3.6-44.0.1.el8.x86_64                  362/654
>   Running scriptlet: kmod-kvdo-6.2.4.26-77.el8.x86_64                      363/654
>   Cleanup          : kmod-kvdo-6.2.4.26-77.el8.x86_64                      363/654
>   Running scriptlet: kmod-kvdo-6.2.4.26-77.el8.x86_64                      363/654
>   Cleanup          : kpatch-dnf-0.2-3.el8.noarch                           364/654
>   Cleanup          : centos-logos-85.5-1.el8.x86_64                        365/654
>   Running scriptlet: centos-logos-85.5-1.el8.x86_64                        365/654
>   Cleanup          : sos-4.0-11.el8.noarch                                 366/654
>   Cleanup          : dracut-squash-049-135.git20210121.el8.x86_64          367/654
>   Cleanup          : python3-firewall-0.8.2-6.el8.noarch                   368/654
>   Cleanup          : python3-nftables-1:0.9.3-18.el8.x86_64                369/654
>   Cleanup          : python3-syspurpose-1.28.13-2.el8.x86_64               370/654
>   Cleanup          : dracut-config-rescue-049-135.git20210121.el8.x86_64   371/654
>   Cleanup          : cockpit-podman-29-2.module_el8.4.0+781+acf4c33b.noa   372/654
>   Cleanup          : podman-3.0.1-6.module_el8.4.0+781+acf4c33b.x86_64     373/654
>   Cleanup          : bind-export-libs-32:9.11.26-3.el8.x86_64              374/654
>   Running scriptlet: bind-export-libs-32:9.11.26-3.el8.x86_64              374/654
>   Cleanup          : openssh-8.0p1-5.el8.x86_64                            375/654
>   Cleanup          : openssl-1:1.1.1g-15.el8_3.x86_64                      376/654
>   Cleanup          : lshw-B.02.19.2-5.el8.x86_64                           377/654
>   Cleanup          : vim-enhanced-2:8.0.1763-15.el8.x86_64                 378/654
>   Cleanup          : nmap-ncat-2:7.70-5.el8.x86_64                         379/654
>   Running scriptlet: nmap-ncat-2:7.70-5.el8.x86_64                         379/654
>   Cleanup          : iproute-5.9.0-4.el8.x86_64                            380/654
>   Cleanup          : bpftool-4.18.0-305.3.1.el8.x86_64                     381/654
>   Cleanup          : iptables-1.8.4-17.el8.x86_64                          382/654
>   Running scriptlet: iptables-1.8.4-17.el8.x86_64                          382/654
>   Cleanup          : libibverbs-32.0-4.el8.x86_64                          383/654
>   Running scriptlet: libibverbs-32.0-4.el8.x86_64                          383/654
>   Cleanup          : fontconfig-2.13.1-3.el8.x86_64                        384/654
>   Cleanup          : conmon-2:2.0.26-1.module_el8.4.0+781+acf4c33b.x86_6   385/654
>   Cleanup          : runc-1.0.0-70.rc92.module_el8.4.0+673+eabfc99d.x86_   386/654
>   Cleanup          : criu-3.15-1.module_el8.4.0+641+6116a774.x86_64        387/654
>   Cleanup          : python3-perf-4.18.0-305.3.1.el8.x86_64                388/654
>   Cleanup          : strace-5.7-2.el8.x86_64                               389/654
>   Running scriptlet: libgomp-8.4.1-1.el8.x86_64                            390/654
>   Cleanup          : libgomp-8.4.1-1.el8.x86_64                            390/654
>   Running scriptlet: libgomp-8.4.1-1.el8.x86_64                            390/654
>   Cleanup          : plymouth-0.9.4-9.20200615git1e36e30.el8.x86_64        391/654
>   Running scriptlet: plymouth-0.9.4-9.20200615git1e36e30.el8.x86_64        391/654
>   Cleanup          : plymouth-core-libs-0.9.4-9.20200615git1e36e30.el8.x   392/654
>   Running scriptlet: plymouth-core-libs-0.9.4-9.20200615git1e36e30.el8.x   392/654
>   Running scriptlet: nftables-1:0.9.3-18.el8.x86_64                        393/654
>   Cleanup          : nftables-1:0.9.3-18.el8.x86_64                        393/654
>   Running scriptlet: nftables-1:0.9.3-18.el8.x86_64                        393/654
>   Cleanup          : util-linux-user-2.32.1-27.el8.x86_64                  394/654
>   Running scriptlet: mcelog-3:173-0.el8.x86_64                             395/654
>   Cleanup          : mcelog-3:173-0.el8.x86_64                             395/654
>   Running scriptlet: mcelog-3:173-0.el8.x86_64                             395/654
>   Running scriptlet: authselect-1.2.2-2.el8.x86_64                         396/654
>   Cleanup          : authselect-1.2.2-2.el8.x86_64                         396/654
>   Cleanup          : tpm2-tools-4.1.1-2.el8.x86_64                         397/654
>   Cleanup          : tpm2-tss-2.3.2-3.el8.x86_64                           398/654
>   Running scriptlet: tpm2-tss-2.3.2-3.el8.x86_64                           398/654
>   Cleanup          : man-db-2.7.6.1-17.el8.x86_64                          399/654
>   Cleanup          : NetworkManager-team-1:1.30.0-7.el8.x86_64             400/654
>   Running scriptlet: NetworkManager-1:1.30.0-7.el8.x86_64                  401/654
>   Cleanup          : NetworkManager-1:1.30.0-7.el8.x86_64                  401/654
>   Running scriptlet: NetworkManager-1:1.30.0-7.el8.x86_64                  401/654
>   Cleanup          : NetworkManager-libnm-1:1.30.0-7.el8.x86_64            402/654
>   Running scriptlet: NetworkManager-libnm-1:1.30.0-7.el8.x86_64            402/654
>   Cleanup          : grub2-tools-extra-1:2.02-99.el8.x86_64                403/654
>   Cleanup          : authselect-libs-1.2.2-2.el8.x86_64                    404/654
>   Cleanup          : rdma-core-32.0-4.el8.x86_64                           405/654
>   Cleanup          : udisks2-lvm2-2.9.0-6.el8.x86_64                       406/654
>   Cleanup          : libblockdev-lvm-2.24-5.el8.x86_64                     407/654
>   Running scriptlet: lvm2-8:2.03.11-5.el8.x86_64                           408/654
>   Cleanup          : lvm2-8:2.03.11-5.el8.x86_64                           408/654
>   Running scriptlet: lvm2-8:2.03.11-5.el8.x86_64                           408/654
>   Cleanup          : lvm2-libs-8:2.03.11-5.el8.x86_64                      409/654
>   Running scriptlet: device-mapper-event-8:1.02.175-5.el8.x86_64           410/654
>   Cleanup          : device-mapper-event-8:1.02.175-5.el8.x86_64           410/654
>   Cleanup          : device-mapper-persistent-data-0.8.5-4.el8.x86_64      411/654
>   Cleanup          : device-mapper-event-libs-8:1.02.175-5.el8.x86_64      412/654
>   Cleanup          : elfutils-debuginfod-client-0.182-3.el8.x86_64         413/654
>   Cleanup          : python3-psutil-5.4.3-10.el8.x86_64                    414/654
>   Cleanup          : libX11-1.6.8-4.el8.x86_64                             415/654
>   Cleanup          : container-selinux-2:2.158.0-1.module_el8.4.0+781+ac   416/654
>   Running scriptlet: container-selinux-2:2.158.0-1.module_el8.4.0+781+ac   416/654
>   Cleanup          : sssd-2.4.0-9.el8.x86_64                               417/654
>   Cleanup          : rpm-plugin-selinux-4.14.3-13.el8.x86_64               418/654
>   Cleanup          : selinux-policy-targeted-3.14.3-67.el8.noarch          419/654
>   Running scriptlet: selinux-policy-targeted-3.14.3-67.el8.noarch          419/654
>   Cleanup          : selinux-policy-3.14.3-67.el8.noarch                   420/654
>   Running scriptlet: selinux-policy-3.14.3-67.el8.noarch                   420/654
>   Cleanup          : plymouth-scripts-0.9.4-9.20200615git1e36e30.el8.x86   421/654
>   Cleanup          : containers-common-1:1.2.2-8.module_el8.4.0+781+acf4   422/654
>   Cleanup          : python3-sssdconfig-2.4.0-9.el8.noarch                 423/654
>   Cleanup          : cockpit-packagekit-238.2-1.el8.noarch                 424/654
>   Cleanup          : grub2-pc-modules-1:2.02-99.el8.noarch                 425/654
>   Cleanup          : yum-4.4.2-11.el8.noarch                               426/654
>   Running scriptlet: dnf-4.4.2-11.el8.noarch                               427/654
>   Cleanup          : dnf-4.4.2-11.el8.noarch                               427/654
>   Running scriptlet: dnf-4.4.2-11.el8.noarch                               427/654
>   Cleanup          : gsettings-desktop-schemas-3.32.0-5.el8.x86_64         428/654
>   Cleanup          : dnf-plugins-core-4.0.18-4.el8.noarch                  429/654
>   Cleanup          : python3-dnf-plugins-core-4.0.18-4.el8.noarch          430/654
>   Cleanup          : python3-dnf-4.4.2-11.el8.noarch                       431/654
>   Cleanup          : centos-linux-release-8.4-1.2105.el8.noarch            432/654
>   Cleanup          : centos-linux-repos-8-2.el8.noarch                     433/654
>   Cleanup          : setroubleshoot-plugins-3.3.13-1.el8.noarch            434/654
>   Cleanup          : centos-gpg-keys-1:8-2.el8.noarch                      435/654
>   Cleanup          : dnf-data-4.4.2-11.el8.noarch                          436/654
>   Cleanup          : libX11-common-1.6.8-4.el8.noarch                      437/654
>   Cleanup          : hwdata-0.314-8.8.el8.noarch                           438/654
>   Cleanup          : podman-catatonit-3.0.1-6.module_el8.4.0+781+acf4c33   439/654
>   Cleanup          : dhcp-common-12:4.3.6-44.0.1.el8.noarch                440/654
>   Cleanup          : firewalld-filesystem-0.8.2-6.el8.noarch               441/654
>   Cleanup          : linux-firmware-20201218-102.git05789708.el8.noarch    442/654
>   Cleanup          : libertas-usb8388-firmware-2:20201218-102.git0578970   443/654
>   Cleanup          : iwl7260-firmware-1:25.30.13.0-102.el8.1.noarch        444/654
>   Cleanup          : iwl6050-firmware-41.28.5.1-102.el8.1.noarch           445/654
>   Cleanup          : iwl6000g2b-firmware-18.168.6.1-102.el8.1.noarch       446/654
>   Cleanup          : iwl6000g2a-firmware-18.168.6.1-102.el8.1.noarch       447/654
>   Cleanup          : iwl6000-firmware-9.221.4.1-102.el8.1.noarch           448/654
>   Cleanup          : iwl5150-firmware-8.24.2.2-102.el8.1.noarch            449/654
>   Cleanup          : iwl5000-firmware-8.83.5.1_1-102.el8.1.noarch          450/654
>   Cleanup          : iwl3160-firmware-1:25.30.13.0-102.el8.1.noarch        451/654
>   Cleanup          : iwl2030-firmware-18.168.6.1-102.el8.1.noarch          452/654
>   Cleanup          : iwl2000-firmware-18.168.6.1-102.el8.1.noarch          453/654
>   Cleanup          : iwl135-firmware-18.168.6.1-102.el8.1.noarch           454/654
>   Cleanup          : iwl105-firmware-18.168.6.1-102.el8.1.noarch           455/654
>   Cleanup          : iwl1000-firmware-1:39.31.5.1-102.el8.1.noarch         456/654
>   Cleanup          : iwl100-firmware-39.31.5.1-102.el8.1.noarch            457/654
>   Cleanup          : emacs-filesystem-1:26.1-5.el8.noarch                  458/654
>   Cleanup          : alsa-sof-firmware-1.6.1-2.el8.noarch                  459/654
>   Cleanup          : NetworkManager-config-server-1:1.30.0-7.el8.noarch    460/654
>   Cleanup          : man-pages-overrides-8.3.0.2-2.el8.noarch              461/654
>   Cleanup          : sssd-ipa-2.4.0-9.el8.x86_64                           462/654
>   Cleanup          : sssd-ad-2.4.0-9.el8.x86_64                            463/654
>   Cleanup          : libsmbclient-4.13.3-3.el8.x86_64                      464/654
>   Cleanup          : sssd-common-pac-2.4.0-9.el8.x86_64                    465/654
>   Running scriptlet: libwbclient-4.13.3-3.el8.x86_64                       466/654
>   Cleanup          : libwbclient-4.13.3-3.el8.x86_64                       466/654
>   Cleanup          : samba-client-libs-4.13.3-3.el8.x86_64                 467/654
>   Cleanup          : samba-common-libs-4.13.3-3.el8.x86_64                 468/654
 >   Cleanup          : python3-hawkey-0.55.0-7.el8.x86_64                    469/644
>   Cleanup          : python3-libdnf-0.55.0-7.el8.x86_64                    470/654
>   Cleanup          : libdnf-0.55.0-7.el8.x86_64                            471/654
>   Cleanup          : sssd-ldap-2.4.0-9.el8.x86_64                          472/654
>   Cleanup          : sssd-proxy-2.4.0-9.el8.x86_64                         473/654
>   Cleanup          : bind-utils-32:9.11.26-3.el8.x86_64                    474/654
>   Cleanup          : cockpit-bridge-238.2-1.el8.x86_64                     475/654
>   Cleanup          : cups-libs-1:2.2.6-38.el8.x86_64                       476/654
>   Cleanup          : sssd-krb5-2.4.0-9.el8.x86_64                          477/654
>   Cleanup          : bind-libs-32:9.11.26-3.el8.x86_64                     478/654
>   Cleanup          : bind-libs-lite-32:9.11.26-3.el8.x86_64                479/654
>   Cleanup          : adcli-0.8.2-9.el8.x86_64                              480/654
>   Running scriptlet: adcli-0.8.2-9.el8.x86_64                              480/654
>   Cleanup          : sssd-krb5-common-2.4.0-9.el8.x86_64                   481/654
>   Running scriptlet: sssd-common-2.4.0-9.el8.x86_64                        482/654
>   Cleanup          : sssd-common-2.4.0-9.el8.x86_64                        482/654
>   Running scriptlet: sssd-common-2.4.0-9.el8.x86_64                        482/654
>   Running scriptlet: sssd-client-2.4.0-9.el8.x86_64                        483/654
>   Cleanup          : sssd-client-2.4.0-9.el8.x86_64                        483/654
>   Running scriptlet: sssd-client-2.4.0-9.el8.x86_64                        483/654
>   Cleanup          : librepo-1.12.0-3.el8.x86_64                           484/654
>   Cleanup          : setroubleshoot-server-3.3.24-3.el8.x86_64             485/654
>   Running scriptlet: setroubleshoot-server-3.3.24-3.el8.x86_64             485/654
>   Cleanup          : python3-libxml2-2.9.7-9.el8.x86_64                    486/654
>   Running scriptlet: polkit-0.115-11.el8.x86_64                            487/654
>   Cleanup          : polkit-0.115-11.el8.x86_64                            487/654
>   Running scriptlet: polkit-0.115-11.el8.x86_64                            487/654
>   Cleanup          : python3-rpm-4.14.3-13.el8.x86_64                      488/654
>   Cleanup          : rpm-build-libs-4.14.3-13.el8.x86_64                   489/654
>   Running scriptlet: rpm-build-libs-4.14.3-13.el8.x86_64                   489/654
>   Cleanup          : libstdc++-8.4.1-1.el8.x86_64                          490/654
>   Running scriptlet: libstdc++-8.4.1-1.el8.x86_64                          490/654
>   Cleanup          : fuse-overlayfs-1.4.0-2.module_el8.4.0+673+eabfc99d.   491/654
>   Cleanup          : rpm-plugin-systemd-inhibit-4.14.3-13.el8.x86_64       492/654
>   Cleanup          : libldb-2.2.0-2.el8.x86_64                             493/654
>   Cleanup          : sqlite-3.26.0-13.el8.x86_64                           494/654
>   Cleanup          : libsss_nss_idmap-2.4.0-9.el8.x86_64                   495/654
>   Running scriptlet: libsss_nss_idmap-2.4.0-9.el8.x86_64                   495/654
>   Cleanup          : libsolv-0.7.16-2.el8.x86_64                           496/654
>   Cleanup          : libsss_certmap-2.4.0-9.el8.x86_64                     497/654
>   Running scriptlet: libsss_certmap-2.4.0-9.el8.x86_64                     497/654
>   Cleanup          : slirp4netns-1.1.8-1.module_el8.4.0+641+6116a774.x86   498/654
>   Cleanup          : libtevent-0.10.2-2.el8.x86_64                         499/654
>   Cleanup          : python3-gpg-1.13.1-7.el8.x86_64                       500/654
>   Cleanup          : python3-unbound-1.7.3-15.el8.x86_64                   501/654
>   Running scriptlet: unbound-libs-1.7.3-15.el8.x86_64                      502/654
>   Cleanup          : unbound-libs-1.7.3-15.el8.x86_64                      502/654
>   Running scriptlet: unbound-libs-1.7.3-15.el8.x86_64                      502/654
>   Cleanup          : libmodulemd-2.9.4-2.el8.x86_64                        503/654
>   Cleanup          : gpgme-1.13.1-7.el8.x86_64                             504/654
>   Cleanup          : python3-libcomps-0.1.11-5.el8.x86_64                  505/654
>   Cleanup          : libcomps-0.1.11-5.el8.x86_64                          506/654
>   Cleanup          : libxml2-2.9.7-9.el8.x86_64                            507/654
>   Cleanup          : numactl-libs-2.0.12-11.el8.x86_64                     508/654
>   Running scriptlet: numactl-libs-2.0.12-11.el8.x86_64                     508/654
>   Cleanup          : libdrm-2.4.103-1.el8.x86_64                           509/654
>   Cleanup          : udisks2-iscsi-2.9.0-6.el8.x86_64                      510/654
>   Running scriptlet: udisks2-2.9.0-6.el8.x86_64                            511/654
>   Cleanup          : udisks2-2.9.0-6.el8.x86_64                            511/654
>   Running scriptlet: udisks2-2.9.0-6.el8.x86_64                            511/654
>   Running scriptlet: iscsi-initiator-utils-6.2.1.2-1.gita8fcb37.el8.x86_   512/654
>   Cleanup          : iscsi-initiator-utils-6.2.1.2-1.gita8fcb37.el8.x86_   512/654
>   Running scriptlet: iscsi-initiator-utils-6.2.1.2-1.gita8fcb37.el8.x86_   512/654
>   Cleanup          : libblockdev-crypto-2.24-5.el8.x86_64                  513/654
>   Cleanup          : nss-3.53.1-17.el8_3.x86_64                            514/654
>   Cleanup          : libblockdev-fs-2.24-5.el8.x86_64                      515/654
>   Cleanup          : nss-softokn-3.53.1-17.el8_3.x86_64                    516/654
>   Cleanup          : e2fsprogs-1.45.6-1.el8.x86_64                         517/654
>   Cleanup          : xfsprogs-5.0.0-8.el8.x86_64                           518/654
>   Running scriptlet: xfsprogs-5.0.0-8.el8.x86_64                           518/654
>   Running scriptlet: iscsi-initiator-utils-iscsiuio-6.2.1.2-1.gita8fcb37   519/654
>   Cleanup          : iscsi-initiator-utils-iscsiuio-6.2.1.2-1.gita8fcb37   519/654
>   Running scriptlet: iscsi-initiator-utils-iscsiuio-6.2.1.2-1.gita8fcb37   519/654
>   Cleanup          : libblockdev-part-2.24-5.el8.x86_64                    520/654
>   Cleanup          : libblockdev-2.24-5.el8.x86_64                         521/654
>   Cleanup          : nss-sysinit-3.53.1-17.el8_3.x86_64                    522/654
>   Cleanup          : libblockdev-swap-2.24-5.el8.x86_64                    523/654
>   Cleanup          : polkit-libs-0.115-11.el8.x86_64                       524/654
>   Running scriptlet: polkit-libs-0.115-11.el8.x86_64                       524/654
>   Cleanup          : libss-1.45.6-1.el8.x86_64                             525/654
>   Running scriptlet: libss-1.45.6-1.el8.x86_64                             525/654
>   Cleanup          : libblockdev-loop-2.24-5.el8.x86_64                    526/654
>   Cleanup          : libblockdev-mdraid-2.24-5.el8.x86_64                  527/654
>   Running scriptlet: mdadm-4.1-15.el8.x86_64                               528/654
>   Cleanup          : mdadm-4.1-15.el8.x86_64                               528/654
>   Running scriptlet: mdadm-4.1-15.el8.x86_64                               528/654
>   Cleanup          : libblockdev-utils-2.24-5.el8.x86_64                   529/654
>   Running scriptlet: parted-3.2-38.el8.x86_64                              530/654
>   Cleanup          : parted-3.2-38.el8.x86_64                              530/654
>   Running scriptlet: parted-3.2-38.el8.x86_64                              530/654
>   Cleanup          : nss-softokn-freebl-3.53.1-17.el8_3.x86_64             531/654
>   Cleanup          : nss-util-3.53.1-17.el8_3.x86_64                       532/654
>   Cleanup          : nspr-4.25.0-2.el8_2.x86_64                            533/654
>   Running scriptlet: nspr-4.25.0-2.el8_2.x86_64                            533/654
>   Cleanup          : quota-1:4.04-12.el8.x86_64                            534/654
>   Cleanup          : e2fsprogs-libs-1.45.6-1.el8.x86_64                    535/654
>   Running scriptlet: e2fsprogs-libs-1.45.6-1.el8.x86_64                    535/654
>   Cleanup          : nvme-cli-1.12-3.el8.0.1.x86_64                        536/654
>   Cleanup          : tcpdump-14:4.9.3-1.el8.x86_64                         537/654
>   Cleanup          : libsss_idmap-2.4.0-9.el8.x86_64                       538/654
>   Running scriptlet: libsss_idmap-2.4.0-9.el8.x86_64                       538/654
>   Cleanup          : libsss_sudo-2.4.0-9.el8.x86_64                        539/654
>   Running scriptlet: libsss_sudo-2.4.0-9.el8.x86_64                        539/654
>   Cleanup          : vim-minimal-2:8.0.1763-15.el8.x86_64                  540/654
>   Cleanup          : rsyslog-gssapi-8.1911.0-7.el8.x86_64                  541/654
>   Cleanup          : sssd-nfs-idmap-2.4.0-9.el8.x86_64                     542/654
>   Cleanup          : libnfsidmap-1:2.3.3-41.el8.x86_64                     543/654
>   Cleanup          : fstrm-0.6.0-3.el8.1.x86_64                            544/654
>   Cleanup          : json-c-0.13.1-0.4.el8.x86_64                          545/654
>   Cleanup          : containernetworking-plugins-0.9.1-1.module_el8.4.0+   546/654
>   Cleanup          : ethtool-2:5.8-5.el8.x86_64                            547/654
>   Cleanup          : rsyslog-relp-8.1911.0-7.el8.x86_64                    548/654
>   Cleanup          : librelp-1.2.16-1.el8.x86_64                           549/654
>   Running scriptlet: librelp-1.2.16-1.el8.x86_64                           549/654
>   Cleanup          : libslirp-4.3.1-1.module_el8.4.0+575+63b40ad7.x86_64   550/654
>   Cleanup          : libipa_hbac-2.4.0-9.el8.x86_64                        551/654
>   Running scriptlet: libipa_hbac-2.4.0-9.el8.x86_64                        551/654
>   Cleanup          : rsyslog-gnutls-8.1911.0-7.el8.x86_64                  552/654
>   Running scriptlet: rsyslog-8.1911.0-7.el8.x86_64                         553/654
>   Cleanup          : rsyslog-8.1911.0-7.el8.x86_64                         553/654
>   Running scriptlet: rsyslog-8.1911.0-7.el8.x86_64                         553/654
>   Cleanup          : libfastjson-0.99.8-2.el8.x86_64                       554/654
>   Running scriptlet: libfastjson-0.99.8-2.el8.x86_64                       554/654
>   Cleanup          : libudisks2-2.9.0-6.el8.x86_64                         555/654
>   Cleanup          : glib2-2.56.4-9.el8.x86_64                             556/654
>   Cleanup          : pcre-8.42-4.el8.x86_64                                557/654
>   Cleanup          : libtalloc-2.3.1-2.el8.x86_64                          558/654
>   Cleanup          : libndp-1.7-5.el8.x86_64                               559/654
>   Running scriptlet: libndp-1.7-5.el8.x86_64                               559/654
>   Cleanup          : vim-common-2:8.0.1763-15.el8.x86_64                   560/654
>   Cleanup          : virt-what-1.18-6.el8.x86_64                           561/654
>   Cleanup          : hdparm-9.54-3.el8.x86_64                              562/654
>   Cleanup          : unzip-6.0-44.el8.x86_64                               563/654
>   Cleanup          : libsss_autofs-2.4.0-9.el8.x86_64                      564/654
>   Cleanup          : kernel-tools-libs-4.18.0-305.3.1.el8.x86_64           565/654
>   Running scriptlet: kernel-tools-libs-4.18.0-305.3.1.el8.x86_64           565/654
>   Cleanup          : lsscsi-0.32-2.el8.x86_64                              566/654
>   Cleanup          : python3-bind-32:9.11.26-3.el8.noarch                  567/654
>   Cleanup          : samba-common-4.13.3-3.el8.noarch                      568/654
>   Cleanup          : policycoreutils-python-utils-2.9-14.el8.noarch        569/654
>   Cleanup          : python3-policycoreutils-2.9-14.el8.noarch             570/654
>   Cleanup          : bind-license-32:9.11.26-3.el8.noarch                  571/654
>   Cleanup          : vim-filesystem-2:8.0.1763-15.el8.noarch               572/654
>   Cleanup          : quota-nls-1:4.04-12.el8.noarch                        573/654
>   Running scriptlet: policycoreutils-2.9-14.el8.x86_64                     574/654
>   Cleanup          : policycoreutils-2.9-14.el8.x86_64                     574/654
>   Cleanup          : platform-python-pip-9.0.3-19.el8.noarch               575/654
>   Cleanup          : dbus-tools-1:1.12.8-12.el8.x86_64                     576/654
>   Cleanup          : dbus-libs-1:1.12.8-12.el8.x86_64                      577/654
>   Running scriptlet: dbus-libs-1:1.12.8-12.el8.x86_64                      577/654
>   Running scriptlet: dbus-daemon-1:1.12.8-12.el8.x86_64                    578/654
>   Cleanup          : dbus-daemon-1:1.12.8-12.el8.x86_64                    578/654
>   Running scriptlet: dbus-daemon-1:1.12.8-12.el8.x86_64                    578/654
>   Cleanup          : systemd-pam-239-45.el8.x86_64                         579/654
>   Cleanup          : libfdisk-2.32.1-27.el8.x86_64                         580/654
>   Running scriptlet: libfdisk-2.32.1-27.el8.x86_64                         580/654
>   Cleanup          : libmount-2.32.1-27.el8.x86_64                         581/654
>   Running scriptlet: libmount-2.32.1-27.el8.x86_64                         581/654
>   Cleanup          : ca-certificates-2020.2.41-80.0.el8_2.noarch           582/654
>   Cleanup          : shadow-utils-2:4.6-12.el8.x86_64                      583/654
>   Cleanup          : libblkid-2.32.1-27.el8.x86_64                         584/654
>   Running scriptlet: libblkid-2.32.1-27.el8.x86_64                         584/654
>   Cleanup          : systemd-libs-239-45.el8.x86_64                        585/654
>   Cleanup          : pam-1.3.1-14.el8.x86_64                               586/654
>   Running scriptlet: pam-1.3.1-14.el8.x86_64                               586/654
>   Cleanup          : util-linux-2.32.1-27.el8.x86_64                       587/654
>   Cleanup          : python3-libs-3.6.8-37.el8.x86_64                      588/654
>   Cleanup          : libtirpc-1.1.4-4.el8.x86_64                           589/654
>   Running scriptlet: libtirpc-1.1.4-4.el8.x86_64                           589/654
>   Cleanup          : platform-python-3.6.8-37.el8.x86_64                   590/654
>   Running scriptlet: platform-python-3.6.8-37.el8.x86_64                   590/654
>   Cleanup          : kmod-25-17.el8.x86_64                                 591/654
>   Cleanup          : kmod-libs-25-17.el8.x86_64                            592/654
>   Running scriptlet: kmod-libs-25-17.el8.x86_64                            592/654
>   Cleanup          : libdb-utils-5.3.28-40.el8.x86_64                      593/654
>   Cleanup          : libcurl-7.61.1-18.el8.x86_64                          594/654
>   Cleanup          : libssh-0.9.4-2.el8.x86_64                             595/654
>   Cleanup          : krb5-libs-1.18.2-8.el8.x86_64                         596/654
>   Cleanup          : openldap-2.4.46-16.el8.x86_64                         597/654
>   Cleanup          : curl-7.61.1-18.el8.x86_64                             598/654
>   Cleanup          : rpm-4.14.3-13.el8.x86_64                              599/654
>   Cleanup          : rpm-libs-4.14.3-13.el8.x86_64                         600/654
>   Running scriptlet: rpm-libs-4.14.3-13.el8.x86_64                         600/654
>   Cleanup          : libdb-5.3.28-40.el8.x86_64                            601/654
>   Running scriptlet: libdb-5.3.28-40.el8.x86_64                            601/654
>   Cleanup          : coreutils-8.30-8.el8.x86_64                           602/654
>   Cleanup          : openssl-libs-1:1.1.1g-15.el8_3.x86_64                 603/654
>   Running scriptlet: openssl-libs-1:1.1.1g-15.el8_3.x86_64                 603/654
>   Cleanup          : gnutls-3.6.14-7.el8_3.x86_64                          604/654
>   Cleanup          : crypto-policies-20210209-1.gitbfb6bed.el8_3.noarch    605/654
>   Cleanup          : crypto-policies-scripts-20210209-1.gitbfb6bed.el8_3   606/654
>   Cleanup          : grubby-8.40-41.el8.x86_64                             607/654
>   Running scriptlet: grub2-tools-1:2.02-99.el8.x86_64                      608/654
>   Cleanup          : grub2-tools-1:2.02-99.el8.x86_64                      608/654
>   Cleanup          : dracut-049-135.git20210121.el8.x86_64                 609/654
>   Cleanup          : kpartx-0.8.4-10.el8.x86_64                            610/654
>   Cleanup          : os-prober-1.74-6.el8.x86_64                           611/654
>   Cleanup          : systemd-udev-239-45.el8.x86_64                        612/654
>   Running scriptlet: systemd-udev-239-45.el8.x86_64                        612/654
>   Cleanup          : grub2-tools-minimal-1:2.02-99.el8.x86_64              613/654
>   Cleanup          : elfutils-libs-0.182-3.el8.x86_64                      614/654
>   Cleanup          : elfutils-default-yama-scope-0.182-3.el8.noarch        615/654
>   Cleanup          : device-mapper-libs-8:1.02.175-5.el8.x86_64            616/654
>   Cleanup          : device-mapper-8:1.02.175-5.el8.x86_64                 617/654
>   Cleanup          : dbus-1:1.12.8-12.el8.x86_64                           618/654
>   Running scriptlet: systemd-239-45.el8.x86_64                             619/654
>   Cleanup          : systemd-239-45.el8.x86_64                             619/654
>   Cleanup          : libsmartcols-2.32.1-27.el8.x86_64                     620/654
>   Running scriptlet: libsmartcols-2.32.1-27.el8.x86_64                     620/654
>   Cleanup          : sqlite-libs-3.26.0-13.el8.x86_64                      621/654
>   Cleanup          : libuuid-2.32.1-27.el8.x86_64                          622/654
>   Running scriptlet: libuuid-2.32.1-27.el8.x86_64                          622/654
>   Cleanup          : iptables-libs-1.8.4-17.el8.x86_64                     623/654
>   Cleanup          : lua-libs-5.3.4-11.el8.x86_64                          624/654
>   Cleanup          : libcom_err-1.45.6-1.el8.x86_64                        625/654
>   Running scriptlet: libcom_err-1.45.6-1.el8.x86_64                        625/654
>   Cleanup          : libgcrypt-1.8.5-4.el8.x86_64                          626/654
>   Running scriptlet: libgcrypt-1.8.5-4.el8.x86_64                          626/654
>   Running scriptlet: nettle-3.4.1-2.el8.x86_64                             627/654
>   Cleanup          : nettle-3.4.1-2.el8.x86_64                             627/654
>   Running scriptlet: nettle-3.4.1-2.el8.x86_64                             627/654
>   Cleanup          : ncurses-6.1-7.20180224.el8.x86_64                     628/654
>   Cleanup          : chkconfig-1.13-2.el8.x86_64                           629/654
>   Cleanup          : libsepol-2.9-2.el8.x86_64                             630/654
>   Running scriptlet: libsepol-2.9-2.el8.x86_64                             630/654
>   Cleanup          : libcap-ng-0.7.9-5.el8.x86_64                          631/654
>   Cleanup          : libcap-2.26-4.el8.x86_64                              632/654
>   Cleanup          : libxcrypt-4.1.1-4.el8.x86_64                          633/654
>   Cleanup          : elfutils-libelf-0.182-3.el8.x86_64                    634/654
>   Cleanup          : keyutils-libs-1.5.10-6.el8.x86_64                     635/654
>   Cleanup          : file-5.33-16.el8_3.1.x86_64                           636/654
>   Cleanup          : file-libs-5.33-16.el8_3.1.x86_64                      637/654
>   Cleanup          : which-2.21-12.el8.x86_64                              638/654
>   Cleanup          : dmidecode-1:3.2-8.el8.x86_64                          639/654
>   Running scriptlet: coreutils-common-8.30-8.el8.x86_64                    640/654
>   Cleanup          : coreutils-common-8.30-8.el8.x86_64                    640/654
>   Cleanup          : grub2-common-1:2.02-99.el8.noarch                     641/654
>   Cleanup          : libssh-config-0.9.4-2.el8.noarch                      642/654
>   Cleanup          : python3-pip-wheel-9.0.3-19.el8.noarch                 643/654
>   Cleanup          : dbus-common-1:1.12.8-12.el8.noarch                    644/654
>   Cleanup          : lz4-libs-1.8.3-2.el8.x86_64                           645/654
>   Cleanup          : ncurses-libs-6.1-7.20180224.el8.x86_64                646/654
>   Cleanup          : bash-4.4.19-14.el8.x86_64                             647/654
>   Running scriptlet: bash-4.4.19-14.el8.x86_64                             647/654
>   Cleanup          : glibc-2.28-151.el8.x86_64                             648/654
>   Cleanup          : glibc-langpack-en-2.28-151.el8.x86_64                 649/654
>   Cleanup          : glibc-common-2.28-151.el8.x86_64                      650/654
>   Cleanup          : tzdata-2021a-1.el8.noarch                             651/654
>   Cleanup          : filesystem-3.8-3.el8.x86_64                           652/654
>   Cleanup          : ncurses-base-6.1-7.20180224.el8.noarch                653/654
>   Cleanup          : libgcc-8.4.1-1.el8.x86_64                             654/654
>   Running scriptlet: libgcc-8.4.1-1.el8.x86_64                             654/654
>   Running scriptlet: filesystem-3.8-6.el8.x86_64                           654/654
>   Running scriptlet: crypto-policies-scripts-20210617-1.gitc776d3e.el8.n   654/654
>   Running scriptlet: ca-certificates-2021.2.50-80.0.el8_4.noarch           654/654
>   Running scriptlet: kernel-core-4.18.0-348.7.1.el8_5.x86_64               654/654
>   Running scriptlet: kmod-kvdo-6.2.5.72-81.el8.x86_64                      654/654
>   Running scriptlet: container-selinux-2:2.167.0-1.module_el8.5.0+911+f1   654/654
>
>   Running scriptlet: authselect-libs-1.2.2-3.el8.x86_64                    654/654
>   Running scriptlet: nss-3.67.0-7.el8_5.x86_64                             654/654
>   Running scriptlet: centos-logos-85.8-2.el8.x86_64                        654/654
>   Running scriptlet: sssd-common-2.5.2-2.el8_5.3.x86_64                    654/654
>   Running scriptlet: libwbclient-4.14.5-7.el8_5.x86_64                     654/654
>   Running scriptlet: tuned-2.16.0-1.el8.noarch                             654/654
>   Running scriptlet: authselect-compat-1.2.2-3.el8.x86_64                  654/654
>   Running scriptlet: microcode_ctl-4:20210608-1.el8.x86_64                 654/654
>   Running scriptlet: libgcc-8.4.1-1.el8.x86_64                             654/654
>   Running scriptlet: glibc-common-2.28-164.el8.x86_64                      654/654
>   Running scriptlet: systemd-239-51.el8_5.2.x86_64                         654/654
>   Running scriptlet: systemd-udev-239-51.el8_5.2.x86_64                    654/654
>   Running scriptlet: glib2-2.56.4-156.el8.x86_64                           654/654
>   Running scriptlet: vim-common-2:8.0.1763-16.el8.x86_64                   654/654
>   Running scriptlet: fontconfig-2.13.1-4.el8.x86_64                        654/654
>   Running scriptlet: man-db-2.7.6.1-18.el8.x86_64                          654/654
>   Verifying        : grub2-tools-efi-1:2.02-106.el8.x86_64                   1/654
>   Verifying        : kernel-4.18.0-348.7.1.el8_5.x86_64                      2/654
>   Verifying        : kernel-core-4.18.0-348.7.1.el8_5.x86_64                 3/654
>   Verifying        : kernel-modules-4.18.0-348.7.1.el8_5.x86_64              4/654
>   Verifying        : libbpf-0.4.0-1.el8.x86_64                               5/654
>   Verifying        : authselect-compat-1.2.2-3.el8.x86_64                    6/654
>   Verifying        : authselect-compat-1.2.2-2.el8.x86_64                    7/654
>   Verifying        : bind-libs-32:9.11.26-6.el8.x86_64                       8/654
>   Verifying        : bind-libs-32:9.11.26-3.el8.x86_64                       9/654
>   Verifying        : bind-libs-lite-32:9.11.26-6.el8.x86_64                 10/654
>   Verifying        : bind-libs-lite-32:9.11.26-3.el8.x86_64                 11/654
>   Verifying        : bind-license-32:9.11.26-6.el8.noarch                   12/654
>   Verifying        : bind-license-32:9.11.26-3.el8.noarch                   13/654
>   Verifying        : bind-utils-32:9.11.26-6.el8.x86_64                     14/654
>   Verifying        : bind-utils-32:9.11.26-3.el8.x86_64                     15/654
>   Verifying        : buildah-1.22.3-2.module_el8.5.0+911+f19012f9.x86_64    16/654
>   Verifying        : buildah-1.19.7-1.module_el8.4.0+781+acf4c33b.x86_64    17/654
>   Verifying        : cockpit-packagekit-251.1-1.el8.noarch                  18/654
>   Verifying        : cockpit-packagekit-238.2-1.el8.noarch                  19/654
>   Verifying        : cockpit-podman-33-1.module_el8.5.0+890+6b136101.noa    20/654
>   Verifying        : cockpit-podman-29-2.module_el8.4.0+781+acf4c33b.noa    21/654
>   Verifying        : cockpit-storaged-251.1-1.el8.noarch                    22/654
>   Verifying        : cockpit-storaged-238.2-1.el8.noarch                    23/654
>   Verifying        : conmon-2:2.0.29-1.module_el8.5.0+890+6b136101.x86_6    24/654
>   Verifying        : conmon-2:2.0.26-1.module_el8.4.0+781+acf4c33b.x86_6    25/654
>   Verifying        : container-selinux-2:2.167.0-1.module_el8.5.0+911+f1    26/654
>   Verifying        : container-selinux-2:2.158.0-1.module_el8.4.0+781+ac    27/654
>   Verifying        : containernetworking-plugins-1.0.0-1.module_el8.5.0+    28/654
>   Verifying        : containernetworking-plugins-0.9.1-1.module_el8.4.0+    29/654
>   Verifying        : containers-common-2:1-2.module_el8.5.0+890+6b136101    30/654
>   Verifying        : containers-common-1:1.2.2-8.module_el8.4.0+781+acf4    31/654
>   Verifying        : criu-3.15-3.module_el8.5.0+890+6b136101.x86_64         32/654
>   Verifying        : criu-3.15-1.module_el8.4.0+641+6116a774.x86_64         33/654
>   Verifying        : fstrm-0.6.1-2.el8.x86_64                               34/654
>   Verifying        : fstrm-0.6.0-3.el8.1.x86_64                             35/654
>   Verifying        : fuse-overlayfs-1.7.1-1.module_el8.5.0+890+6b136101.    36/654
>   Verifying        : fuse-overlayfs-1.4.0-2.module_el8.4.0+673+eabfc99d.    37/654
>   Verifying        : libX11-1.6.8-5.el8.x86_64                              38/654
>   Verifying        : libX11-1.6.8-4.el8.x86_64                              39/654
>   Verifying        : libX11-common-1.6.8-5.el8.noarch                       40/654
>   Verifying        : libX11-common-1.6.8-4.el8.noarch                       41/654
>   Verifying        : libblockdev-2.24-7.el8.x86_64                          42/654
>   Verifying        : libblockdev-2.24-5.el8.x86_64                          43/654
>   Verifying        : libblockdev-crypto-2.24-7.el8.x86_64                   44/654
>   Verifying        : libblockdev-crypto-2.24-5.el8.x86_64                   45/654
>   Verifying        : libblockdev-fs-2.24-7.el8.x86_64                       46/654
>   Verifying        : libblockdev-fs-2.24-5.el8.x86_64                       47/654
>   Verifying        : libblockdev-loop-2.24-7.el8.x86_64                     48/654
>   Verifying        : libblockdev-loop-2.24-5.el8.x86_64                     49/654
>   Verifying        : libblockdev-lvm-2.24-7.el8.x86_64                      50/654
>   Verifying        : libblockdev-lvm-2.24-5.el8.x86_64                      51/654
>   Verifying        : libblockdev-mdraid-2.24-7.el8.x86_64                   52/654
>   Verifying        : libblockdev-mdraid-2.24-5.el8.x86_64                   53/654
>   Verifying        : libblockdev-part-2.24-7.el8.x86_64                     54/654
>   Verifying        : libblockdev-part-2.24-5.el8.x86_64                     55/654
>   Verifying        : libblockdev-swap-2.24-7.el8.x86_64                     56/654
>   Verifying        : libblockdev-swap-2.24-5.el8.x86_64                     57/654
>   Verifying        : libblockdev-utils-2.24-7.el8.x86_64                    58/654
>   Verifying        : libblockdev-utils-2.24-5.el8.x86_64                    59/654
>   Verifying        : libdrm-2.4.106-2.el8.x86_64                            60/654
>   Verifying        : libdrm-2.4.103-1.el8.x86_64                            61/654
>   Verifying        : libfastjson-0.99.9-1.el8.x86_64                        62/654
>   Verifying        : libfastjson-0.99.8-2.el8.x86_64                        63/654
>   Verifying        : librelp-1.9.0-1.el8.x86_64                             64/654
>   Verifying        : librelp-1.2.16-1.el8.x86_64                            65/654
>   Verifying        : libslirp-4.4.0-1.module_el8.5.0+890+6b136101.x86_64    66/654
>   Verifying        : libslirp-4.3.1-1.module_el8.4.0+575+63b40ad7.x86_64    67/654
>   Verifying        : libudisks2-2.9.0-7.el8.x86_64                          68/654
>   Verifying        : libudisks2-2.9.0-6.el8.x86_64                          69/654
>   Verifying        : man-pages-overrides-8.5.0.1-1.el8.noarch               70/654
>   Verifying        : man-pages-overrides-8.3.0.2-2.el8.noarch               71/654
>   Verifying        : nmap-ncat-2:7.70-6.el8.x86_64                          72/654
>   Verifying        : nmap-ncat-2:7.70-5.el8.x86_64                          73/654
>   Verifying        : nspr-4.32.0-1.el8_4.x86_64                             74/654
>   Verifying        : nspr-4.25.0-2.el8_2.x86_64                             75/654
>   Verifying        : nss-3.67.0-7.el8_5.x86_64                              76/654
>   Verifying        : nss-3.53.1-17.el8_3.x86_64                             77/654
>   Verifying        : nss-softokn-3.67.0-7.el8_5.x86_64                      78/654
>   Verifying        : nss-softokn-3.53.1-17.el8_3.x86_64                     79/654
>   Verifying        : nss-softokn-freebl-3.67.0-7.el8_5.x86_64               80/654
>   Verifying        : nss-softokn-freebl-3.53.1-17.el8_3.x86_64              81/654
>   Verifying        : nss-sysinit-3.67.0-7.el8_5.x86_64                      82/654
>   Verifying        : nss-sysinit-3.53.1-17.el8_3.x86_64                     83/654
>   Verifying        : nss-util-3.67.0-7.el8_5.x86_64                         84/654
>   Verifying        : nss-util-3.53.1-17.el8_3.x86_64                        85/654
>   Verifying        : open-vm-tools-11.2.5-2.el8.x86_64                      86/654
>   Verifying        : open-vm-tools-11.2.0-2.el8.x86_64                      87/654
>   Verifying        : plymouth-0.9.4-10.20200615git1e36e30.el8.x86_64        88/654
>   Verifying        : plymouth-0.9.4-9.20200615git1e36e30.el8.x86_64         89/654
>   Verifying        : plymouth-core-libs-0.9.4-10.20200615git1e36e30.el8.    90/654
>   Verifying        : plymouth-core-libs-0.9.4-9.20200615git1e36e30.el8.x    91/654
>   Verifying        : plymouth-scripts-0.9.4-10.20200615git1e36e30.el8.x8    92/654
>   Verifying        : plymouth-scripts-0.9.4-9.20200615git1e36e30.el8.x86    93/654
>   Verifying        : podman-3.3.1-9.module_el8.5.0+988+b1f0b741.x86_64      94/654
>   Verifying        : podman-3.0.1-6.module_el8.4.0+781+acf4c33b.x86_64      95/654
>   Verifying        : podman-catatonit-3.3.1-9.module_el8.5.0+988+b1f0b74    96/654
>   Verifying        : podman-catatonit-3.0.1-6.module_el8.4.0+781+acf4c33    97/654
>   Verifying        : python3-bind-32:9.11.26-6.el8.noarch                   98/654
>   Verifying        : python3-bind-32:9.11.26-3.el8.noarch                   99/654
>   Verifying        : python3-lxml-4.2.3-3.el8.x86_64                       100/654
>   Verifying        : python3-lxml-4.2.3-2.el8.x86_64                       101/654
>   Verifying        : python3-psutil-5.4.3-11.el8.x86_64                    102/654
>   Verifying        : python3-psutil-5.4.3-10.el8.x86_64                    103/654
>   Verifying        : python3-unbound-1.7.3-17.el8.x86_64                   104/654
>   Verifying        : python3-unbound-1.7.3-15.el8.x86_64                   105/654
>   Verifying        : rsyslog-8.2102.0-5.el8.x86_64                         106/654
>   Verifying        : rsyslog-8.1911.0-7.el8.x86_64                         107/654
>   Verifying        : rsyslog-gnutls-8.2102.0-5.el8.x86_64                  108/654
>   Verifying        : rsyslog-gnutls-8.1911.0-7.el8.x86_64                  109/654
>   Verifying        : rsyslog-gssapi-8.2102.0-5.el8.x86_64                  110/654
>   Verifying        : rsyslog-gssapi-8.1911.0-7.el8.x86_64                  111/654
>   Verifying        : rsyslog-relp-8.2102.0-5.el8.x86_64                    112/654
>   Verifying        : rsyslog-relp-8.1911.0-7.el8.x86_64                    113/654
>   Verifying        : runc-1.0.2-1.module_el8.5.0+911+f19012f9.x86_64       114/654
>   Verifying        : runc-1.0.0-70.rc92.module_el8.4.0+673+eabfc99d.x86_   115/654
>   Verifying        : setroubleshoot-plugins-3.3.14-1.el8.noarch            116/654
>   Verifying        : setroubleshoot-plugins-3.3.13-1.el8.noarch            117/654
>   Verifying        : setroubleshoot-server-3.3.24-4.el8.x86_64             118/654
>   Verifying        : setroubleshoot-server-3.3.24-3.el8.x86_64             119/654
>   Verifying        : slirp4netns-1.1.8-1.module_el8.5.0+890+6b136101.x86   120/654
>   Verifying        : slirp4netns-1.1.8-1.module_el8.4.0+641+6116a774.x86   121/654
>   Verifying        : tcpdump-14:4.9.3-2.el8.x86_64                         122/654
>   Verifying        : tcpdump-14:4.9.3-1.el8.x86_64                         123/654
>   Verifying        : udisks2-2.9.0-7.el8.x86_64                            124/654
>   Verifying        : udisks2-2.9.0-6.el8.x86_64                            125/654
>   Verifying        : udisks2-iscsi-2.9.0-7.el8.x86_64                      126/654
>   Verifying        : udisks2-iscsi-2.9.0-6.el8.x86_64                      127/654
>   Verifying        : udisks2-lvm2-2.9.0-7.el8.x86_64                       128/654
>   Verifying        : udisks2-lvm2-2.9.0-6.el8.x86_64                       129/654
>   Verifying        : unbound-libs-1.7.3-17.el8.x86_64                      130/654
>   Verifying        : unbound-libs-1.7.3-15.el8.x86_64                      131/654
>   Verifying        : vim-common-2:8.0.1763-16.el8.x86_64                   132/654
>   Verifying        : vim-common-2:8.0.1763-15.el8.x86_64                   133/654
>   Verifying        : vim-enhanced-2:8.0.1763-16.el8.x86_64                 134/654
>   Verifying        : vim-enhanced-2:8.0.1763-15.el8.x86_64                 135/654
>   Verifying        : vim-filesystem-2:8.0.1763-16.el8.noarch               136/654
>   Verifying        : vim-filesystem-2:8.0.1763-15.el8.noarch               137/654
>   Verifying        : NetworkManager-1:1.32.10-4.el8.x86_64                 138/654
>   Verifying        : NetworkManager-1:1.30.0-7.el8.x86_64                  139/654
>   Verifying        : NetworkManager-config-server-1:1.32.10-4.el8.noarch   140/654
>   Verifying        : NetworkManager-config-server-1:1.30.0-7.el8.noarch    141/654
>   Verifying        : NetworkManager-libnm-1:1.32.10-4.el8.x86_64           142/654
>   Verifying        : NetworkManager-libnm-1:1.30.0-7.el8.x86_64            143/654
>   Verifying        : NetworkManager-team-1:1.32.10-4.el8.x86_64            144/654
>   Verifying        : NetworkManager-team-1:1.30.0-7.el8.x86_64             145/654
>   Verifying        : NetworkManager-tui-1:1.32.10-4.el8.x86_64             146/654
>   Verifying        : NetworkManager-tui-1:1.30.0-7.el8.x86_64              147/654
>   Verifying        : adcli-0.8.2-12.el8.x86_64                             148/654
>   Verifying        : adcli-0.8.2-9.el8.x86_64                              149/654
>   Verifying        : alsa-sof-firmware-1.8-1.el8.noarch                    150/654
>   Verifying        : alsa-sof-firmware-1.6.1-2.el8.noarch                  151/654
>   Verifying        : authselect-1.2.2-3.el8.x86_64                         152/654
>   Verifying        : authselect-1.2.2-2.el8.x86_64                         153/654
>   Verifying        : authselect-libs-1.2.2-3.el8.x86_64                    154/654
>   Verifying        : authselect-libs-1.2.2-2.el8.x86_64                    155/654
>   Verifying        : bash-4.4.20-2.el8.x86_64                              156/654
>   Verifying        : bash-4.4.19-14.el8.x86_64                             157/654
>   Verifying        : bind-export-libs-32:9.11.26-6.el8.x86_64              158/654
>   Verifying        : bind-export-libs-32:9.11.26-3.el8.x86_64              159/654
>   Verifying        : binutils-2.30-108.el8_5.1.x86_64                      160/654
>   Verifying        : binutils-2.30-93.el8.x86_64                           161/654
>   Verifying        : bpftool-4.18.0-348.7.1.el8_5.x86_64                   162/654
>   Verifying        : bpftool-4.18.0-305.3.1.el8.x86_64                     163/654
>   Verifying        : ca-certificates-2021.2.50-80.0.el8_4.noarch           164/654
>   Verifying        : ca-certificates-2020.2.41-80.0.el8_2.noarch           165/654
>   Verifying        : centos-gpg-keys-1:8-3.el8.noarch                      166/654
>   Verifying        : centos-gpg-keys-1:8-2.el8.noarch                      167/654
>   Verifying        : centos-linux-release-8.5-1.2111.el8.noarch            168/654
>   Verifying        : centos-linux-release-8.4-1.2105.el8.noarch            169/654
>   Verifying        : centos-linux-repos-8-3.el8.noarch                     170/654
>   Verifying        : centos-linux-repos-8-2.el8.noarch                     171/654
>   Verifying        : centos-logos-85.8-2.el8.x86_64                        172/654
>   Verifying        : centos-logos-85.5-1.el8.x86_64                        173/654
>   Verifying        : chkconfig-1.19.1-1.el8.x86_64                         174/654
>   Verifying        : chkconfig-1.13-2.el8.x86_64                           175/654
>   Verifying        : chrony-4.1-1.el8.x86_64                               176/654
>   Verifying        : chrony-3.5-2.el8.x86_64                               177/654
>   Verifying        : cockpit-251.1-1.el8.x86_64                            178/654
>   Verifying        : cockpit-238.2-1.el8.x86_64                            179/654
>   Verifying        : cockpit-bridge-251.1-1.el8.x86_64                     180/654
>   Verifying        : cockpit-bridge-238.2-1.el8.x86_64                     181/654
>   Verifying        : cockpit-system-251.1-1.el8.noarch                     182/654
>   Verifying        : cockpit-system-238.2-1.el8.noarch                     183/654
>   Verifying        : cockpit-ws-251.1-1.el8.x86_64                         184/654
>   Verifying        : cockpit-ws-238.2-1.el8.x86_64                         185/654
>   Verifying        : coreutils-8.30-12.el8.x86_64                          186/654
>   Verifying        : coreutils-8.30-8.el8.x86_64                           187/654
>   Verifying        : coreutils-common-8.30-12.el8.x86_64                   188/654
>   Verifying        : coreutils-common-8.30-8.el8.x86_64                    189/654
>   Verifying        : crypto-policies-20210617-1.gitc776d3e.el8.noarch      190/654
>   Verifying        : crypto-policies-20210209-1.gitbfb6bed.el8_3.noarch    191/654
>   Verifying        : crypto-policies-scripts-20210617-1.gitc776d3e.el8.n   192/654
>   Verifying        : crypto-policies-scripts-20210209-1.gitbfb6bed.el8_3   193/654
>   Verifying        : cups-libs-1:2.2.6-40.el8.x86_64                       194/654
>   Verifying        : cups-libs-1:2.2.6-38.el8.x86_64                       195/654
>   Verifying        : curl-7.61.1-22.el8.x86_64                             196/654
>   Verifying        : curl-7.61.1-18.el8.x86_64                             197/654
>   Verifying        : dbus-1:1.12.8-14.el8.x86_64                           198/654
>   Verifying        : dbus-1:1.12.8-12.el8.x86_64                           199/654
>   Verifying        : dbus-common-1:1.12.8-14.el8.noarch                    200/654
>   Verifying        : dbus-common-1:1.12.8-12.el8.noarch                    201/654
>   Verifying        : dbus-daemon-1:1.12.8-14.el8.x86_64                    202/654
>   Verifying        : dbus-daemon-1:1.12.8-12.el8.x86_64                    203/654
>   Verifying        : dbus-libs-1:1.12.8-14.el8.x86_64                      204/654
>   Verifying        : dbus-libs-1:1.12.8-12.el8.x86_64                      205/654
>   Verifying        : dbus-tools-1:1.12.8-14.el8.x86_64                     206/654
>   Verifying        : dbus-tools-1:1.12.8-12.el8.x86_64                     207/654
>   Verifying        : device-mapper-8:1.02.177-10.el8.x86_64                208/654
>   Verifying        : device-mapper-8:1.02.175-5.el8.x86_64                 209/654
>   Verifying        : device-mapper-event-8:1.02.177-10.el8.x86_64          210/654
>   Verifying        : device-mapper-event-8:1.02.175-5.el8.x86_64           211/654
>   Verifying        : device-mapper-event-libs-8:1.02.177-10.el8.x86_64     212/654
>   Verifying        : device-mapper-event-libs-8:1.02.175-5.el8.x86_64      213/654
>   Verifying        : device-mapper-libs-8:1.02.177-10.el8.x86_64           214/654
>   Verifying        : device-mapper-libs-8:1.02.175-5.el8.x86_64            215/654
>   Verifying        : device-mapper-multipath-0.8.4-17.el8.x86_64           216/654
>   Verifying        : device-mapper-multipath-0.8.4-10.el8.x86_64           217/654
>   Verifying        : device-mapper-multipath-libs-0.8.4-17.el8.x86_64      218/654
 >   Verifying        : device-mapper-multipath-libs-0.8.4-10.el8.x86_64      219/644
>   Verifying        : device-mapper-persistent-data-0.9.0-4.el8.x86_64      220/654
>   Verifying        : device-mapper-persistent-data-0.8.5-4.el8.x86_64      221/654
>   Verifying        : dhcp-client-12:4.3.6-45.el8.x86_64                    222/654
>   Verifying        : dhcp-client-12:4.3.6-44.0.1.el8.x86_64                223/654
>   Verifying        : dhcp-common-12:4.3.6-45.el8.noarch                    224/654
>   Verifying        : dhcp-common-12:4.3.6-44.0.1.el8.noarch                225/654
>   Verifying        : dhcp-libs-12:4.3.6-45.el8.x86_64                      226/654
>   Verifying        : dhcp-libs-12:4.3.6-44.0.1.el8.x86_64                  227/654
>   Verifying        : dmidecode-1:3.2-10.el8.x86_64                         228/654
>   Verifying        : dmidecode-1:3.2-8.el8.x86_64                          229/654
>   Verifying        : dnf-4.7.0-4.el8.noarch                                230/654
>   Verifying        : dnf-4.4.2-11.el8.noarch                               231/654
>   Verifying        : dnf-data-4.7.0-4.el8.noarch                           232/654
>   Verifying        : dnf-data-4.4.2-11.el8.noarch                          233/654
>   Verifying        : dnf-plugins-core-4.0.21-3.el8.noarch                  234/654
>   Verifying        : dnf-plugins-core-4.0.18-4.el8.noarch                  235/654
>   Verifying        : dracut-049-191.git20210920.el8.x86_64                 236/654
>   Verifying        : dracut-049-135.git20210121.el8.x86_64                 237/654
>   Verifying        : dracut-config-rescue-049-191.git20210920.el8.x86_64   238/654
>   Verifying        : dracut-config-rescue-049-135.git20210121.el8.x86_64   239/654
>   Verifying        : dracut-network-049-191.git20210920.el8.x86_64         240/654
>   Verifying        : dracut-network-049-135.git20210121.el8.x86_64         241/654
>   Verifying        : dracut-squash-049-191.git20210920.el8.x86_64          242/654
>   Verifying        : dracut-squash-049-135.git20210121.el8.x86_64          243/654
>   Verifying        : e2fsprogs-1.45.6-2.el8.x86_64                         244/654
>   Verifying        : e2fsprogs-1.45.6-1.el8.x86_64                         245/654
>   Verifying        : e2fsprogs-libs-1.45.6-2.el8.x86_64                    246/654
>   Verifying        : e2fsprogs-libs-1.45.6-1.el8.x86_64                    247/654
>   Verifying        : elfutils-debuginfod-client-0.185-1.el8.x86_64         248/654
>   Verifying        : elfutils-debuginfod-client-0.182-3.el8.x86_64         249/654
>   Verifying        : elfutils-default-yama-scope-0.185-1.el8.noarch        250/654
>   Verifying        : elfutils-default-yama-scope-0.182-3.el8.noarch        251/654
>   Verifying        : elfutils-libelf-0.185-1.el8.x86_64                    252/654
>   Verifying        : elfutils-libelf-0.182-3.el8.x86_64                    253/654
>   Verifying        : elfutils-libs-0.185-1.el8.x86_64                      254/654
>   Verifying        : elfutils-libs-0.182-3.el8.x86_64                      255/654
>   Verifying        : emacs-filesystem-1:26.1-7.el8.noarch                  256/654
>   Verifying        : emacs-filesystem-1:26.1-5.el8.noarch                  257/654
>   Verifying        : ethtool-2:5.8-7.el8.x86_64                            258/654
>   Verifying        : ethtool-2:5.8-5.el8.x86_64                            259/654
>   Verifying        : file-5.33-20.el8.x86_64                               260/654
>   Verifying        : file-5.33-16.el8_3.1.x86_64                           261/654
>   Verifying        : file-libs-5.33-20.el8.x86_64                          262/654
>   Verifying        : file-libs-5.33-16.el8_3.1.x86_64                      263/654
>   Verifying        : filesystem-3.8-6.el8.x86_64                           264/654
>   Verifying        : filesystem-3.8-3.el8.x86_64                           265/654
>   Verifying        : firewalld-0.9.3-7.el8.noarch                          266/654
>   Verifying        : firewalld-0.8.2-6.el8.noarch                          267/654
8>   Verifying        : firewalld-filesystem-0.9.3-7.el8.noarch               268/644
>   Verifying        : firewalld-filesystem-0.8.2-6.el8.noarch               269/654
>   Verifying        : fontconfig-2.13.1-4.el8.x86_64                        270/654
>   Verifying        : fontconfig-2.13.1-3.el8.x86_64                        271/654
>   Verifying        : glib2-2.56.4-156.el8.x86_64                           272/654
>   Verifying        : glib2-2.56.4-9.el8.x86_64                             273/654
>   Verifying        : glibc-2.28-164.el8.x86_64                             274/654
>   Verifying        : glibc-2.28-151.el8.x86_64                             275/654
>   Verifying        : glibc-common-2.28-164.el8.x86_64                      276/654
>   Verifying        : glibc-common-2.28-151.el8.x86_64                      277/654
>   Verifying        : glibc-langpack-en-2.28-164.el8.x86_64                 278/654
>   Verifying        : glibc-langpack-en-2.28-151.el8.x86_64                 279/654
>   Verifying        : gnutls-3.6.16-4.el8.x86_64                            280/654
>   Verifying        : gnutls-3.6.14-7.el8_3.x86_64                          281/654
>   Verifying        : gpgme-1.13.1-9.el8.x86_64                             282/654
>   Verifying        : gpgme-1.13.1-7.el8.x86_64                             283/654
>   Verifying        : grub2-common-1:2.02-106.el8.noarch                    284/654
>   Verifying        : grub2-common-1:2.02-99.el8.noarch                     285/654
>   Verifying        : grub2-pc-1:2.02-106.el8.x86_64                        286/654
>   Verifying        : grub2-pc-1:2.02-99.el8.x86_64                         287/654
>   Verifying        : grub2-pc-modules-1:2.02-106.el8.noarch                288/654
>   Verifying        : grub2-pc-modules-1:2.02-99.el8.noarch                 289/654
>   Verifying        : grub2-tools-1:2.02-106.el8.x86_64                     290/654
>   Verifying        : grub2-tools-1:2.02-99.el8.x86_64                      291/654
>   Verifying        : grub2-tools-extra-1:2.02-106.el8.x86_64               292/654
>   Verifying        : grub2-tools-extra-1:2.02-99.el8.x86_64                293/654
>   Verifying        : grub2-tools-minimal-1:2.02-106.el8.x86_64             294/654
>   Verifying        : grub2-tools-minimal-1:2.02-99.el8.x86_64              295/654
>   Verifying        : grubby-8.40-42.el8.x86_64                             296/654
>   Verifying        : grubby-8.40-41.el8.x86_64                             297/654
>   Verifying        : gsettings-desktop-schemas-3.32.0-6.el8.x86_64         298/654
>   Verifying        : gsettings-desktop-schemas-3.32.0-5.el8.x86_64         299/654
>   Verifying        : hdparm-9.54-4.el8.x86_64                              300/654
>   Verifying        : hdparm-9.54-3.el8.x86_64                              301/654
>   Verifying        : hwdata-0.314-8.10.el8.noarch                          302/654
>   Verifying        : hwdata-0.314-8.8.el8.noarch                           303/654
>   Verifying        : iproute-5.12.0-4.el8.x86_64                           304/654
>   Verifying        : iproute-5.9.0-4.el8.x86_64                            305/654
>   Verifying        : iptables-1.8.4-20.el8.x86_64                          306/654
>   Verifying        : iptables-1.8.4-17.el8.x86_64                          307/654
>   Verifying        : iptables-ebtables-1.8.4-20.el8.x86_64                 308/654
>   Verifying        : iptables-ebtables-1.8.4-17.el8.x86_64                 309/654
>   Verifying        : iptables-libs-1.8.4-20.el8.x86_64                     310/654
>   Verifying        : iptables-libs-1.8.4-17.el8.x86_64                     311/654
>   Verifying        : iscsi-initiator-utils-6.2.1.4-4.git095f59c.el8.x86_   312/654
>   Verifying        : iscsi-initiator-utils-6.2.1.2-1.gita8fcb37.el8.x86_   313/654
>   Verifying        : iscsi-initiator-utils-iscsiuio-6.2.1.4-4.git095f59c   314/654
>   Verifying        : iscsi-initiator-utils-iscsiuio-6.2.1.2-1.gita8fcb37   315/654
>   Verifying        : iwl100-firmware-39.31.5.1-103.el8.1.noarch            316/654
>   Verifying        : iwl100-firmware-39.31.5.1-102.el8.1.noarch            317/654
>   Verifying        : iwl1000-firmware-1:39.31.5.1-103.el8.1.noarch         318/654
>   Verifying        : iwl1000-firmware-1:39.31.5.1-102.el8.1.noarch         319/654
>   Verifying        : iwl105-firmware-18.168.6.1-103.el8.1.noarch           320/654
>   Verifying        : iwl105-firmware-18.168.6.1-102.el8.1.noarch           321/654
>   Verifying        : iwl135-firmware-18.168.6.1-103.el8.1.noarch           322/654
>   Verifying        : iwl135-firmware-18.168.6.1-102.el8.1.noarch           323/654
>   Verifying        : iwl2000-firmware-18.168.6.1-103.el8.1.noarch          324/654
>   Verifying        : iwl2000-firmware-18.168.6.1-102.el8.1.noarch          325/654
>   Verifying        : iwl2030-firmware-18.168.6.1-103.el8.1.noarch          326/654
>   Verifying        : iwl2030-firmware-18.168.6.1-102.el8.1.noarch          327/654
>   Verifying        : iwl3160-firmware-1:25.30.13.0-103.el8.1.noarch        328/654
>   Verifying        : iwl3160-firmware-1:25.30.13.0-102.el8.1.noarch        329/654
>   Verifying        : iwl5000-firmware-8.83.5.1_1-103.el8.1.noarch          330/654
>   Verifying        : iwl5000-firmware-8.83.5.1_1-102.el8.1.noarch          331/654
>   Verifying        : iwl5150-firmware-8.24.2.2-103.el8.1.noarch            332/654
>   Verifying        : iwl5150-firmware-8.24.2.2-102.el8.1.noarch            333/654
>   Verifying        : iwl6000-firmware-9.221.4.1-103.el8.1.noarch           334/654
>   Verifying        : iwl6000-firmware-9.221.4.1-102.el8.1.noarch           335/654
>   Verifying        : iwl6000g2a-firmware-18.168.6.1-103.el8.1.noarch       336/654
>   Verifying        : iwl6000g2a-firmware-18.168.6.1-102.el8.1.noarch       337/654
>   Verifying        : iwl6000g2b-firmware-18.168.6.1-103.el8.1.noarch       338/654
>   Verifying        : iwl6000g2b-firmware-18.168.6.1-102.el8.1.noarch       339/654
>   Verifying        : iwl6050-firmware-41.28.5.1-103.el8.1.noarch           340/654
>   Verifying        : iwl6050-firmware-41.28.5.1-102.el8.1.noarch           341/654
>   Verifying        : iwl7260-firmware-1:25.30.13.0-103.el8.1.noarch        342/654
>   Verifying        : iwl7260-firmware-1:25.30.13.0-102.el8.1.noarch        343/654
>   Verifying        : json-c-0.13.1-2.el8.x86_64                            344/654
>   Verifying        : json-c-0.13.1-0.4.el8.x86_64                          345/654
>   Verifying        : kernel-tools-4.18.0-348.7.1.el8_5.x86_64              346/654
>   Verifying        : kernel-tools-4.18.0-305.3.1.el8.x86_64                347/654
>   Verifying        : kernel-tools-libs-4.18.0-348.7.1.el8_5.x86_64         348/654
>   Verifying        : kernel-tools-libs-4.18.0-305.3.1.el8.x86_64           349/654
>   Verifying        : kexec-tools-2.0.20-57.el8_5.1.x86_64                  350/654
>   Verifying        : kexec-tools-2.0.20-46.el8.x86_64                      351/654
>   Verifying        : keyutils-libs-1.5.10-9.el8.x86_64                     352/654
>   Verifying        : keyutils-libs-1.5.10-6.el8.x86_64                     353/654
>   Verifying        : kmod-25-18.el8.x86_64                                 354/654
>   Verifying        : kmod-25-17.el8.x86_64                                 355/654
>   Verifying        : kmod-kvdo-6.2.5.72-81.el8.x86_64                      356/654
>   Verifying        : kmod-kvdo-6.2.4.26-77.el8.x86_64                      357/654
>   Verifying        : kmod-libs-25-18.el8.x86_64                            358/654
>   Verifying        : kmod-libs-25-17.el8.x86_64                            359/654
>   Verifying        : kpartx-0.8.4-17.el8.x86_64                            360/654
>   Verifying        : kpartx-0.8.4-10.el8.x86_64                            361/654
>   Verifying        : kpatch-0.9.2-5.el8.noarch                             362/654
>   Verifying        : kpatch-0.9.2-3.el8.noarch                             363/654
>   Verifying        : kpatch-dnf-0.2-5.el8.noarch                           364/654
>   Verifying        : kpatch-dnf-0.2-3.el8.noarch                           365/654
>   Verifying        : krb5-libs-1.18.2-14.el8.x86_64                        366/654
>   Verifying        : krb5-libs-1.18.2-8.el8.x86_64                         367/654
>   Verifying        : libblkid-2.32.1-28.el8.x86_64                         368/654
>   Verifying        : libblkid-2.32.1-27.el8.x86_64                         369/654
>   Verifying        : libcap-2.26-5.el8.x86_64                              370/654
>   Verifying        : libcap-2.26-4.el8.x86_64                              371/654
>   Verifying        : libcap-ng-0.7.11-1.el8.x86_64                         372/654
>   Verifying        : libcap-ng-0.7.9-5.el8.x86_64                          373/654
>   Verifying        : libcom_err-1.45.6-2.el8.x86_64                        374/654
>   Verifying        : libcom_err-1.45.6-1.el8.x86_64                        375/654
>   Verifying        : libcomps-0.1.16-2.el8.x86_64                          376/654
>   Verifying        : libcomps-0.1.11-5.el8.x86_64                          377/654
>   Verifying        : libcurl-7.61.1-22.el8.x86_64                          378/654
>   Verifying        : libcurl-7.61.1-18.el8.x86_64                          379/654
>   Verifying        : libdb-5.3.28-42.el8_4.x86_64                          380/654
>   Verifying        : libdb-5.3.28-40.el8.x86_64                            381/654
>   Verifying        : libdb-utils-5.3.28-42.el8_4.x86_64                    382/654
>   Verifying        : libdb-utils-5.3.28-40.el8.x86_64                      383/654
>   Verifying        : libdnf-0.63.0-3.el8.x86_64                            384/654
>   Verifying        : libdnf-0.55.0-7.el8.x86_64                            385/654
>   Verifying        : libertas-usb8388-firmware-2:20210702-103.gitd79c267   386/654
>   Verifying        : libertas-usb8388-firmware-2:20201218-102.git0578970   387/654
>   Verifying        : libfdisk-2.32.1-28.el8.x86_64                         388/654
>   Verifying        : libfdisk-2.32.1-27.el8.x86_64                         389/654
>   Verifying        : libgcc-8.5.0-4.el8_5.x86_64                           390/654
>   Verifying        : libgcc-8.4.1-1.el8.x86_64                             391/654
>   Verifying        : libgcrypt-1.8.5-6.el8.x86_64                          392/654
>   Verifying        : libgcrypt-1.8.5-4.el8.x86_64                          393/654
>   Verifying        : libgomp-8.5.0-4.el8_5.x86_64                          394/654
>   Verifying        : libgomp-8.4.1-1.el8.x86_64                            395/654
>   Verifying        : libibverbs-35.0-1.el8.x86_64                          396/654
>   Verifying        : libibverbs-32.0-4.el8.x86_64                          397/654
>   Verifying        : libipa_hbac-2.5.2-2.el8_5.3.x86_64                    398/654
>   Verifying        : libipa_hbac-2.4.0-9.el8.x86_64                        399/654
>   Verifying        : libldb-2.3.0-2.el8.x86_64                             400/654
>   Verifying        : libldb-2.2.0-2.el8.x86_64                             401/654
>   Verifying        : libmodulemd-2.13.0-1.el8.x86_64                       402/654
>   Verifying        : libmodulemd-2.9.4-2.el8.x86_64                        403/654
>   Verifying        : libmount-2.32.1-28.el8.x86_64                         404/654
>   Verifying        : libmount-2.32.1-27.el8.x86_64                         405/654
>   Verifying        : libndp-1.7-6.el8.x86_64                               406/654
>   Verifying        : libndp-1.7-5.el8.x86_64                               407/654
>   Verifying        : libnfsidmap-1:2.3.3-46.el8.x86_64                     408/654
>   Verifying        : libnfsidmap-1:2.3.3-41.el8.x86_64                     409/654
>   Verifying        : librepo-1.14.0-2.el8.x86_64                           410/654
>   Verifying        : librepo-1.12.0-3.el8.x86_64                           411/654
>   Verifying        : libsepol-2.9-3.el8.x86_64                             412/654
>   Verifying        : libsepol-2.9-2.el8.x86_64                             413/654
>   Verifying        : libsmartcols-2.32.1-28.el8.x86_64                     414/654
>   Verifying        : libsmartcols-2.32.1-27.el8.x86_64                     415/654
>   Verifying        : libsmbclient-4.14.5-7.el8_5.x86_64                    416/654
>   Verifying        : libsmbclient-4.13.3-3.el8.x86_64                      417/654
>   Verifying        : libsolv-0.7.19-1.el8.x86_64                           418/654
>   Verifying        : libsolv-0.7.16-2.el8.x86_64                           419/654
>   Verifying        : libss-1.45.6-2.el8.x86_64                             420/654
>   Verifying        : libss-1.45.6-1.el8.x86_64                             421/654
>   Verifying        : libssh-0.9.4-3.el8.x86_64                             422/654
>   Verifying        : libssh-0.9.4-2.el8.x86_64                             423/654
>   Verifying        : libssh-config-0.9.4-3.el8.noarch                      424/654
>   Verifying        : libssh-config-0.9.4-2.el8.noarch                      425/654
>   Verifying        : libsss_autofs-2.5.2-2.el8_5.3.x86_64                  426/654
>   Verifying        : libsss_autofs-2.4.0-9.el8.x86_64                      427/654
>   Verifying        : libsss_certmap-2.5.2-2.el8_5.3.x86_64                 428/654
>   Verifying        : libsss_certmap-2.4.0-9.el8.x86_64                     429/654
>   Verifying        : libsss_idmap-2.5.2-2.el8_5.3.x86_64                   430/654
>   Verifying        : libsss_idmap-2.4.0-9.el8.x86_64                       431/654
>   Verifying        : libsss_nss_idmap-2.5.2-2.el8_5.3.x86_64               432/654
>   Verifying        : libsss_nss_idmap-2.4.0-9.el8.x86_64                   433/654
>   Verifying        : libsss_sudo-2.5.2-2.el8_5.3.x86_64                    434/654
>   Verifying        : libsss_sudo-2.4.0-9.el8.x86_64                        435/654
>   Verifying        : libstdc++-8.5.0-4.el8_5.x86_64                        436/654
>   Verifying        : libstdc++-8.4.1-1.el8.x86_64                          437/654
>   Verifying        : libstoragemgmt-1.9.1-1.el8.x86_64                     438/654
>   Verifying        : libstoragemgmt-1.8.7-1.el8.x86_64                     439/654
>   Verifying        : libtalloc-2.3.2-1.el8.x86_64                          440/654
>   Verifying        : libtalloc-2.3.1-2.el8.x86_64                          441/654
>   Verifying        : libtevent-0.11.0-0.el8.x86_64                         442/654
>   Verifying        : libtevent-0.10.2-2.el8.x86_64                         443/654
>   Verifying        : libtirpc-1.1.4-5.el8.x86_64                           444/654
>   Verifying        : libtirpc-1.1.4-4.el8.x86_64                           445/654
>   Verifying        : libuuid-2.32.1-28.el8.x86_64                          446/654
>   Verifying        : libuuid-2.32.1-27.el8.x86_64                          447/654
>   Verifying        : libwbclient-4.14.5-7.el8_5.x86_64                     448/654
>   Verifying        : libwbclient-4.13.3-3.el8.x86_64                       449/654
>   Verifying        : libxcrypt-4.1.1-6.el8.x86_64                          450/654
>   Verifying        : libxcrypt-4.1.1-4.el8.x86_64                          451/654
>   Verifying        : libxml2-2.9.7-9.el8_4.2.x86_64                        452/654
>   Verifying        : libxml2-2.9.7-9.el8.x86_64                            453/654
>   Verifying        : linux-firmware-20210702-103.gitd79c2677.el8.noarch    454/654
>   Verifying        : linux-firmware-20201218-102.git05789708.el8.noarch    455/654
>   Verifying        : lshw-B.02.19.2-6.el8.x86_64                           456/654
>   Verifying        : lshw-B.02.19.2-5.el8.x86_64                           457/654
>   Verifying        : lsscsi-0.32-3.el8.x86_64                              458/654
>   Verifying        : lsscsi-0.32-2.el8.x86_64                              459/654
>   Verifying        : lua-libs-5.3.4-12.el8.x86_64                          460/654
>   Verifying        : lua-libs-5.3.4-11.el8.x86_64                          461/654
>   Verifying        : lvm2-8:2.03.12-10.el8.x86_64                          462/654
>   Verifying        : lvm2-8:2.03.11-5.el8.x86_64                           463/654
>   Verifying        : lvm2-libs-8:2.03.12-10.el8.x86_64                     464/654
>   Verifying        : lvm2-libs-8:2.03.11-5.el8.x86_64                      465/654
0>   Verifying        : lz4-libs-1.8.3-3.el8_4.x86_64                         466/644
>   Verifying        : lz4-libs-1.8.3-2.el8.x86_64                           467/654
>   Verifying        : man-db-2.7.6.1-18.el8.x86_64                          468/654
>   Verifying        : man-db-2.7.6.1-17.el8.x86_64                          469/654
>   Verifying        : mcelog-3:175-1.el8.x86_64                             470/654
>   Verifying        : mcelog-3:173-0.el8.x86_64                             471/654
>   Verifying        : mdadm-4.2-rc2.el8.x86_64                              472/654
>   Verifying        : mdadm-4.1-15.el8.x86_64                               473/654
>   Verifying        : microcode_ctl-4:20210608-1.el8.x86_64                 474/654
>   Verifying        : microcode_ctl-4:20210216-1.el8.x86_64                 475/654
>   Verifying        : ncurses-6.1-9.20180224.el8.x86_64                     476/654
>   Verifying        : ncurses-6.1-7.20180224.el8.x86_64                     477/654
>   Verifying        : ncurses-base-6.1-9.20180224.el8.noarch                478/654
>   Verifying        : ncurses-base-6.1-7.20180224.el8.noarch                479/654
>   Verifying        : ncurses-libs-6.1-9.20180224.el8.x86_64                480/654
>   Verifying        : ncurses-libs-6.1-7.20180224.el8.x86_64                481/654
>   Verifying        : nettle-3.4.1-7.el8.x86_64                             482/654
>   Verifying        : nettle-3.4.1-2.el8.x86_64                             483/654
>   Verifying        : nftables-1:0.9.3-21.el8.x86_64                        484/654
>   Verifying        : nftables-1:0.9.3-18.el8.x86_64                        485/654
>   Verifying        : numactl-libs-2.0.12-13.el8.x86_64                     486/654
>   Verifying        : numactl-libs-2.0.12-11.el8.x86_64                     487/654
>   Verifying        : nvme-cli-1.14-3.el8.x86_64                            488/654
>   Verifying        : nvme-cli-1.12-3.el8.0.1.x86_64                        489/654
>   Verifying        : openldap-2.4.46-18.el8.x86_64                         490/654
>   Verifying        : openldap-2.4.46-16.el8.x86_64                         491/654
>   Verifying        : openssh-8.0p1-10.el8.x86_64                           492/654
>   Verifying        : openssh-8.0p1-5.el8.x86_64                            493/654
>   Verifying        : openssh-clients-8.0p1-10.el8.x86_64                   494/654
>   Verifying        : openssh-clients-8.0p1-5.el8.x86_64                    495/654
>   Verifying        : openssh-server-8.0p1-10.el8.x86_64                    496/654
>   Verifying        : openssh-server-8.0p1-5.el8.x86_64                     497/654
>   Verifying        : openssl-1:1.1.1k-5.el8_5.x86_64                       498/654
>   Verifying        : openssl-1:1.1.1g-15.el8_3.x86_64                      499/654
>   Verifying        : openssl-libs-1:1.1.1k-5.el8_5.x86_64                  500/654
>   Verifying        : openssl-libs-1:1.1.1g-15.el8_3.x86_64                 501/654
>   Verifying        : os-prober-1.74-9.el8.x86_64                           502/654
>   Verifying        : os-prober-1.74-6.el8.x86_64                           503/654
>   Verifying        : pam-1.3.1-15.el8.x86_64                               504/654
>   Verifying        : pam-1.3.1-14.el8.x86_64                               505/654
>   Verifying        : parted-3.2-39.el8.x86_64                              506/654
>   Verifying        : parted-3.2-38.el8.x86_64                              507/654
>   Verifying        : pcre-8.42-6.el8.x86_64                                508/654
>   Verifying        : pcre-8.42-4.el8.x86_64                                509/654
>   Verifying        : platform-python-3.6.8-41.el8.x86_64                   510/654
>   Verifying        : platform-python-3.6.8-37.el8.x86_64                   511/654
>   Verifying        : platform-python-pip-9.0.3-20.el8.noarch               512/654
>   Verifying        : platform-python-pip-9.0.3-19.el8.noarch               513/654
>   Verifying        : policycoreutils-2.9-16.el8.x86_64                     514/654
y>   Verifying        : policycoreutils-2.9-14.el8.x86_64                     515/644
>   Verifying        : policycoreutils-python-utils-2.9-16.el8.noarch        516/654
>   Verifying        : policycoreutils-python-utils-2.9-14.el8.noarch        517/654
>   Verifying        : polkit-0.115-12.el8.x86_64                            518/654
>   Verifying        : polkit-0.115-11.el8.x86_64                            519/654
>   Verifying        : polkit-libs-0.115-12.el8.x86_64                       520/654
>   Verifying        : polkit-libs-0.115-11.el8.x86_64                       521/654
>   Verifying        : python3-dnf-4.7.0-4.el8.noarch                        522/654
>   Verifying        : python3-dnf-4.4.2-11.el8.noarch                       523/654
>   Verifying        : python3-dnf-plugins-core-4.0.21-3.el8.noarch          524/654
>   Verifying        : python3-dnf-plugins-core-4.0.18-4.el8.noarch          525/654
>   Verifying        : python3-firewall-0.9.3-7.el8.noarch                   526/654
>   Verifying        : python3-firewall-0.8.2-6.el8.noarch                   527/654
>   Verifying        : python3-gpg-1.13.1-9.el8.x86_64                       528/654
>   Verifying        : python3-gpg-1.13.1-7.el8.x86_64                       529/654
>   Verifying        : python3-hawkey-0.63.0-3.el8.x86_64                    530/654
>   Verifying        : python3-hawkey-0.55.0-7.el8.x86_64                    531/654
>   Verifying        : python3-libcomps-0.1.16-2.el8.x86_64                  532/654
>   Verifying        : python3-libcomps-0.1.11-5.el8.x86_64                  533/654
>   Verifying        : python3-libdnf-0.63.0-3.el8.x86_64                    534/654
>   Verifying        : python3-libdnf-0.55.0-7.el8.x86_64                    535/654
>   Verifying        : python3-libs-3.6.8-41.el8.x86_64                      536/654
>   Verifying        : python3-libs-3.6.8-37.el8.x86_64                      537/654
>   Verifying        : python3-libstoragemgmt-1.9.1-1.el8.x86_64             538/654
>   Verifying        : python3-libstoragemgmt-1.8.7-1.el8.noarch             539/654
>   Verifying        : python3-libstoragemgmt-clibs-1.8.7-1.el8.x86_64       540/654
>   Verifying        : python3-libxml2-2.9.7-9.el8_4.2.x86_64                541/654
>   Verifying        : python3-libxml2-2.9.7-9.el8.x86_64                    542/654
>   Verifying        : python3-nftables-1:0.9.3-21.el8.x86_64                543/654
>   Verifying        : python3-nftables-1:0.9.3-18.el8.x86_64                544/654
>   Verifying        : python3-perf-4.18.0-348.7.1.el8_5.x86_64              545/654
>   Verifying        : python3-perf-4.18.0-305.3.1.el8.x86_64                546/654
>   Verifying        : python3-pip-wheel-9.0.3-20.el8.noarch                 547/654
>   Verifying        : python3-pip-wheel-9.0.3-19.el8.noarch                 548/654
>   Verifying        : python3-policycoreutils-2.9-16.el8.noarch             549/654
>   Verifying        : python3-policycoreutils-2.9-14.el8.noarch             550/654
>   Verifying        : python3-rpm-4.14.3-19.el8.x86_64                      551/654
>   Verifying        : python3-rpm-4.14.3-13.el8.x86_64                      552/654
>   Verifying        : python3-sssdconfig-2.5.2-2.el8_5.3.noarch             553/654
>   Verifying        : python3-sssdconfig-2.4.0-9.el8.noarch                 554/654
>   Verifying        : python3-syspurpose-1.28.21-3.el8.x86_64               555/654
>   Verifying        : python3-syspurpose-1.28.13-2.el8.x86_64               556/654
>   Verifying        : quota-1:4.04-14.el8.x86_64                            557/654
>   Verifying        : quota-1:4.04-12.el8.x86_64                            558/654
>   Verifying        : quota-nls-1:4.04-14.el8.noarch                        559/654
>   Verifying        : quota-nls-1:4.04-12.el8.noarch                        560/654
>   Verifying        : rdma-core-35.0-1.el8.x86_64                           561/654
>   Verifying        : rdma-core-32.0-4.el8.x86_64                           562/654
>   Verifying        : realmd-0.16.3-23.el8.x86_64                           563/654
>   Verifying        : realmd-0.16.3-22.el8.x86_64                           564/654
>   Verifying        : rpm-4.14.3-19.el8.x86_64                              565/654
>   Verifying        : rpm-4.14.3-13.el8.x86_64                              566/654
>   Verifying        : rpm-build-libs-4.14.3-19.el8.x86_64                   567/654
>   Verifying        : rpm-build-libs-4.14.3-13.el8.x86_64                   568/654
>   Verifying        : rpm-libs-4.14.3-19.el8.x86_64                         569/654
>   Verifying        : rpm-libs-4.14.3-13.el8.x86_64                         570/654
>   Verifying        : rpm-plugin-selinux-4.14.3-19.el8.x86_64               571/654
>   Verifying        : rpm-plugin-selinux-4.14.3-13.el8.x86_64               572/654
>   Verifying        : rpm-plugin-systemd-inhibit-4.14.3-19.el8.x86_64       573/654
>   Verifying        : rpm-plugin-systemd-inhibit-4.14.3-13.el8.x86_64       574/654
>   Verifying        : samba-client-libs-4.14.5-7.el8_5.x86_64               575/654
>   Verifying        : samba-client-libs-4.13.3-3.el8.x86_64                 576/654
>   Verifying        : samba-common-4.14.5-7.el8_5.noarch                    577/654
>   Verifying        : samba-common-4.13.3-3.el8.noarch                      578/654
>   Verifying        : samba-common-libs-4.14.5-7.el8_5.x86_64               579/654
>   Verifying        : samba-common-libs-4.13.3-3.el8.x86_64                 580/654
>   Verifying        : selinux-policy-3.14.3-80.el8_5.2.noarch               581/654
>   Verifying        : selinux-policy-3.14.3-67.el8.noarch                   582/654
>   Verifying        : selinux-policy-targeted-3.14.3-80.el8_5.2.noarch      583/654
>   Verifying        : selinux-policy-targeted-3.14.3-67.el8.noarch          584/654
>   Verifying        : shadow-utils-2:4.6-14.el8.x86_64                      585/654
>   Verifying        : shadow-utils-2:4.6-12.el8.x86_64                      586/654
>   Verifying        : sos-4.1-5.el8.noarch                                  587/654
>   Verifying        : sos-4.0-11.el8.noarch                                 588/654
>   Verifying        : sqlite-3.26.0-15.el8.x86_64                           589/654
>   Verifying        : sqlite-3.26.0-13.el8.x86_64                           590/654
>   Verifying        : sqlite-libs-3.26.0-15.el8.x86_64                      591/654
>   Verifying        : sqlite-libs-3.26.0-13.el8.x86_64                      592/654
>   Verifying        : sssd-2.5.2-2.el8_5.3.x86_64                           593/654
>   Verifying        : sssd-2.4.0-9.el8.x86_64                               594/654
>   Verifying        : sssd-ad-2.5.2-2.el8_5.3.x86_64                        595/654
>   Verifying        : sssd-ad-2.4.0-9.el8.x86_64                            596/654
>   Verifying        : sssd-client-2.5.2-2.el8_5.3.x86_64                    597/654
>   Verifying        : sssd-client-2.4.0-9.el8.x86_64                        598/654
>   Verifying        : sssd-common-2.5.2-2.el8_5.3.x86_64                    599/654
>   Verifying        : sssd-common-2.4.0-9.el8.x86_64                        600/654
>   Verifying        : sssd-common-pac-2.5.2-2.el8_5.3.x86_64                601/654
>   Verifying        : sssd-common-pac-2.4.0-9.el8.x86_64                    602/654
>   Verifying        : sssd-ipa-2.5.2-2.el8_5.3.x86_64                       603/654
>   Verifying        : sssd-ipa-2.4.0-9.el8.x86_64                           604/654
>   Verifying        : sssd-kcm-2.5.2-2.el8_5.3.x86_64                       605/654
>   Verifying        : sssd-kcm-2.4.0-9.el8.x86_64                           606/654
>   Verifying        : sssd-krb5-2.5.2-2.el8_5.3.x86_64                      607/654
>   Verifying        : sssd-krb5-2.4.0-9.el8.x86_64                          608/654
>   Verifying        : sssd-krb5-common-2.5.2-2.el8_5.3.x86_64               609/654
>   Verifying        : sssd-krb5-common-2.4.0-9.el8.x86_64                   610/654
>   Verifying        : sssd-ldap-2.5.2-2.el8_5.3.x86_64                      611/654
>   Verifying        : sssd-ldap-2.4.0-9.el8.x86_64                          612/654
>   Verifying        : sssd-nfs-idmap-2.5.2-2.el8_5.3.x86_64                 613/654
>   Verifying        : sssd-nfs-idmap-2.4.0-9.el8.x86_64                     614/654
>   Verifying        : sssd-proxy-2.5.2-2.el8_5.3.x86_64                     615/654
>   Verifying        : sssd-proxy-2.4.0-9.el8.x86_64                         616/654
>   Verifying        : strace-5.7-3.el8.x86_64                               617/654
>   Verifying        : strace-5.7-2.el8.x86_64                               618/654
>   Verifying        : sudo-1.8.29-7.el8_4.1.x86_64                          619/654
>   Verifying        : sudo-1.8.29-7.el8.x86_64                              620/654
>   Verifying        : systemd-239-51.el8_5.2.x86_64                         621/654
>   Verifying        : systemd-239-45.el8.x86_64                             622/654
>   Verifying        : systemd-libs-239-51.el8_5.2.x86_64                    623/654
>   Verifying        : systemd-libs-239-45.el8.x86_64                        624/654
>   Verifying        : systemd-pam-239-51.el8_5.2.x86_64                     625/654
>   Verifying        : systemd-pam-239-45.el8.x86_64                         626/654
>   Verifying        : systemd-udev-239-51.el8_5.2.x86_64                    627/654
>   Verifying        : systemd-udev-239-45.el8.x86_64                        628/654
>   Verifying        : tpm2-tools-4.1.1-5.el8.x86_64                         629/654
>   Verifying        : tpm2-tools-4.1.1-2.el8.x86_64                         630/654
>   Verifying        : tpm2-tss-2.3.2-4.el8.x86_64                           631/654
>   Verifying        : tpm2-tss-2.3.2-3.el8.x86_64                           632/654
>   Verifying        : tuned-2.16.0-1.el8.noarch                             633/654
>   Verifying        : tuned-2.15.0-2.el8.noarch                             634/654
>   Verifying        : tzdata-2021e-1.el8.noarch                             635/654
>   Verifying        : tzdata-2021a-1.el8.noarch                             636/654
>   Verifying        : unzip-6.0-45.el8_4.x86_64                             637/654
>   Verifying        : unzip-6.0-44.el8.x86_64                               638/654
>   Verifying        : util-linux-2.32.1-28.el8.x86_64                       639/654
>   Verifying        : util-linux-2.32.1-27.el8.x86_64                       640/654
>   Verifying        : util-linux-user-2.32.1-28.el8.x86_64                  641/654
>   Verifying        : util-linux-user-2.32.1-27.el8.x86_64                  642/654
>   Verifying        : vdo-6.2.5.74-14.el8.x86_64                            643/654
>   Verifying        : vdo-6.2.4.14-14.el8.x86_64                            644/654
>   Verifying        : vim-minimal-2:8.0.1763-16.el8.x86_64                  645/654
>   Verifying        : vim-minimal-2:8.0.1763-15.el8.x86_64                  646/654
>   Verifying        : virt-what-1.18-12.el8.x86_64                          647/654
>   Verifying        : virt-what-1.18-6.el8.x86_64                           648/654
>   Verifying        : which-2.21-16.el8.x86_64                              649/654
>   Verifying        : which-2.21-12.el8.x86_64                              650/654
>   Verifying        : xfsprogs-5.0.0-9.el8.x86_64                           651/654
>   Verifying        : xfsprogs-5.0.0-8.el8.x86_64                           652/654
>   Verifying        : yum-4.7.0-4.el8.noarch                                653/654
>   Verifying        : yum-4.4.2-11.el8.noarch                               654/654
>
> Upgraded:
>   NetworkManager-1:1.32.10-4.el8.x86_64
>   NetworkManager-config-server-1:1.32.10-4.el8.noarch
>   NetworkManager-libnm-1:1.32.10-4.el8.x86_64
>   NetworkManager-team-1:1.32.10-4.el8.x86_64
>   NetworkManager-tui-1:1.32.10-4.el8.x86_64
>   adcli-0.8.2-12.el8.x86_64
>   alsa-sof-firmware-1.8-1.el8.noarch
>   authselect-1.2.2-3.el8.x86_64
>   authselect-compat-1.2.2-3.el8.x86_64
>   authselect-libs-1.2.2-3.el8.x86_64
>   bash-4.4.20-2.el8.x86_64
>   bind-export-libs-32:9.11.26-6.el8.x86_64
>   bind-libs-32:9.11.26-6.el8.x86_64
>   bind-libs-lite-32:9.11.26-6.el8.x86_64
>   bind-license-32:9.11.26-6.el8.noarch
>   bind-utils-32:9.11.26-6.el8.x86_64
>   binutils-2.30-108.el8_5.1.x86_64
>   bpftool-4.18.0-348.7.1.el8_5.x86_64
>   buildah-1.22.3-2.module_el8.5.0+911+f19012f9.x86_64
>   ca-certificates-2021.2.50-80.0.el8_4.noarch
>   centos-gpg-keys-1:8-3.el8.noarch
>   centos-linux-release-8.5-1.2111.el8.noarch
>   centos-linux-repos-8-3.el8.noarch
>   centos-logos-85.8-2.el8.x86_64
>   chkconfig-1.19.1-1.el8.x86_64
>   chrony-4.1-1.el8.x86_64
>   cockpit-251.1-1.el8.x86_64
>   cockpit-bridge-251.1-1.el8.x86_64
>   cockpit-packagekit-251.1-1.el8.noarch
>   cockpit-podman-33-1.module_el8.5.0+890+6b136101.noarch
>   cockpit-storaged-251.1-1.el8.noarch
>   cockpit-system-251.1-1.el8.noarch
>   cockpit-ws-251.1-1.el8.x86_64
>   conmon-2:2.0.29-1.module_el8.5.0+890+6b136101.x86_64
>   container-selinux-2:2.167.0-1.module_el8.5.0+911+f19012f9.noarch
>   containernetworking-plugins-1.0.0-1.module_el8.5.0+890+6b136101.x86_64
>   containers-common-2:1-2.module_el8.5.0+890+6b136101.noarch
>   coreutils-8.30-12.el8.x86_64
>   coreutils-common-8.30-12.el8.x86_64
>   criu-3.15-3.module_el8.5.0+890+6b136101.x86_64
>   crypto-policies-20210617-1.gitc776d3e.el8.noarch
>   crypto-policies-scripts-20210617-1.gitc776d3e.el8.noarch
>   cups-libs-1:2.2.6-40.el8.x86_64
>   curl-7.61.1-22.el8.x86_64
>   dbus-1:1.12.8-14.el8.x86_64
>   dbus-common-1:1.12.8-14.el8.noarch
>   dbus-daemon-1:1.12.8-14.el8.x86_64
>   dbus-libs-1:1.12.8-14.el8.x86_64
>   dbus-tools-1:1.12.8-14.el8.x86_64
>   device-mapper-8:1.02.177-10.el8.x86_64
>   device-mapper-event-8:1.02.177-10.el8.x86_64
>   device-mapper-event-libs-8:1.02.177-10.el8.x86_64
>   device-mapper-libs-8:1.02.177-10.el8.x86_64
>   device-mapper-multipath-0.8.4-17.el8.x86_64
>   device-mapper-multipath-libs-0.8.4-17.el8.x86_64
>   device-mapper-persistent-data-0.9.0-4.el8.x86_64
>   dhcp-client-12:4.3.6-45.el8.x86_64
>   dhcp-common-12:4.3.6-45.el8.noarch
>   dhcp-libs-12:4.3.6-45.el8.x86_64
>   dmidecode-1:3.2-10.el8.x86_64
>   dnf-4.7.0-4.el8.noarch
>   dnf-data-4.7.0-4.el8.noarch
>   dnf-plugins-core-4.0.21-3.el8.noarch
>   dracut-049-191.git20210920.el8.x86_64
>   dracut-config-rescue-049-191.git20210920.el8.x86_64
>   dracut-network-049-191.git20210920.el8.x86_64
>   dracut-squash-049-191.git20210920.el8.x86_64
>   e2fsprogs-1.45.6-2.el8.x86_64
>   e2fsprogs-libs-1.45.6-2.el8.x86_64
>   elfutils-debuginfod-client-0.185-1.el8.x86_64
>   elfutils-default-yama-scope-0.185-1.el8.noarch
>   elfutils-libelf-0.185-1.el8.x86_64
>   elfutils-libs-0.185-1.el8.x86_64
>   emacs-filesystem-1:26.1-7.el8.noarch
>   ethtool-2:5.8-7.el8.x86_64
>   file-5.33-20.el8.x86_64
>   file-libs-5.33-20.el8.x86_64
>   filesystem-3.8-6.el8.x86_64
>   firewalld-0.9.3-7.el8.noarch
>   firewalld-filesystem-0.9.3-7.el8.noarch
>   fontconfig-2.13.1-4.el8.x86_64
>   fstrm-0.6.1-2.el8.x86_64
>   fuse-overlayfs-1.7.1-1.module_el8.5.0+890+6b136101.x86_64
>   glib2-2.56.4-156.el8.x86_64
>   glibc-2.28-164.el8.x86_64
>   glibc-common-2.28-164.el8.x86_64
>   glibc-langpack-en-2.28-164.el8.x86_64
>   gnutls-3.6.16-4.el8.x86_64
>   gpgme-1.13.1-9.el8.x86_64
>   grub2-common-1:2.02-106.el8.noarch
>   grub2-pc-1:2.02-106.el8.x86_64
>   grub2-pc-modules-1:2.02-106.el8.noarch
>   grub2-tools-1:2.02-106.el8.x86_64
>   grub2-tools-extra-1:2.02-106.el8.x86_64
>   grub2-tools-minimal-1:2.02-106.el8.x86_64
>   grubby-8.40-42.el8.x86_64
>   gsettings-desktop-schemas-3.32.0-6.el8.x86_64
>   hdparm-9.54-4.el8.x86_64
>   hwdata-0.314-8.10.el8.noarch
>   iproute-5.12.0-4.el8.x86_64
>   iptables-1.8.4-20.el8.x86_64
>   iptables-ebtables-1.8.4-20.el8.x86_64
>   iptables-libs-1.8.4-20.el8.x86_64
>   iscsi-initiator-utils-6.2.1.4-4.git095f59c.el8.x86_64
>   iscsi-initiator-utils-iscsiuio-6.2.1.4-4.git095f59c.el8.x86_64
>   iwl100-firmware-39.31.5.1-103.el8.1.noarch
>   iwl1000-firmware-1:39.31.5.1-103.el8.1.noarch
>   iwl105-firmware-18.168.6.1-103.el8.1.noarch
>   iwl135-firmware-18.168.6.1-103.el8.1.noarch
>   iwl2000-firmware-18.168.6.1-103.el8.1.noarch
>   iwl2030-firmware-18.168.6.1-103.el8.1.noarch
>   iwl3160-firmware-1:25.30.13.0-103.el8.1.noarch
>   iwl5000-firmware-8.83.5.1_1-103.el8.1.noarch
>   iwl5150-firmware-8.24.2.2-103.el8.1.noarch
>   iwl6000-firmware-9.221.4.1-103.el8.1.noarch
>   iwl6000g2a-firmware-18.168.6.1-103.el8.1.noarch
>   iwl6000g2b-firmware-18.168.6.1-103.el8.1.noarch
>   iwl6050-firmware-41.28.5.1-103.el8.1.noarch
>   iwl7260-firmware-1:25.30.13.0-103.el8.1.noarch
>   json-c-0.13.1-2.el8.x86_64
>   kernel-tools-4.18.0-348.7.1.el8_5.x86_64
>   kernel-tools-libs-4.18.0-348.7.1.el8_5.x86_64
>   kexec-tools-2.0.20-57.el8_5.1.x86_64
>   keyutils-libs-1.5.10-9.el8.x86_64
>   kmod-25-18.el8.x86_64
>   kmod-kvdo-6.2.5.72-81.el8.x86_64
>   kmod-libs-25-18.el8.x86_64
>   kpartx-0.8.4-17.el8.x86_64
>   kpatch-0.9.2-5.el8.noarch
>   kpatch-dnf-0.2-5.el8.noarch
>   krb5-libs-1.18.2-14.el8.x86_64
>   libX11-1.6.8-5.el8.x86_64
>   libX11-common-1.6.8-5.el8.noarch
>   libblkid-2.32.1-28.el8.x86_64
>   libblockdev-2.24-7.el8.x86_64
>   libblockdev-crypto-2.24-7.el8.x86_64
>   libblockdev-fs-2.24-7.el8.x86_64
>   libblockdev-loop-2.24-7.el8.x86_64
>   libblockdev-lvm-2.24-7.el8.x86_64
>   libblockdev-mdraid-2.24-7.el8.x86_64
>   libblockdev-part-2.24-7.el8.x86_64
>   libblockdev-swap-2.24-7.el8.x86_64
>   libblockdev-utils-2.24-7.el8.x86_64
>   libcap-2.26-5.el8.x86_64
>   libcap-ng-0.7.11-1.el8.x86_64
>   libcom_err-1.45.6-2.el8.x86_64
>   libcomps-0.1.16-2.el8.x86_64
>   libcurl-7.61.1-22.el8.x86_64
>   libdb-5.3.28-42.el8_4.x86_64
>   libdb-utils-5.3.28-42.el8_4.x86_64
>   libdnf-0.63.0-3.el8.x86_64
>   libdrm-2.4.106-2.el8.x86_64
>   libertas-usb8388-firmware-2:20210702-103.gitd79c2677.el8.noarch
>   libfastjson-0.99.9-1.el8.x86_64
>   libfdisk-2.32.1-28.el8.x86_64
>   libgcc-8.5.0-4.el8_5.x86_64
>   libgcrypt-1.8.5-6.el8.x86_64
>   libgomp-8.5.0-4.el8_5.x86_64
>   libibverbs-35.0-1.el8.x86_64
>   libipa_hbac-2.5.2-2.el8_5.3.x86_64
>   libldb-2.3.0-2.el8.x86_64
>   libmodulemd-2.13.0-1.el8.x86_64
>   libmount-2.32.1-28.el8.x86_64
>   libndp-1.7-6.el8.x86_64
>   libnfsidmap-1:2.3.3-46.el8.x86_64
>   librelp-1.9.0-1.el8.x86_64
>   librepo-1.14.0-2.el8.x86_64
>   libsepol-2.9-3.el8.x86_64
>   libslirp-4.4.0-1.module_el8.5.0+890+6b136101.x86_64
>   libsmartcols-2.32.1-28.el8.x86_64
>   libsmbclient-4.14.5-7.el8_5.x86_64
>   libsolv-0.7.19-1.el8.x86_64
>   libss-1.45.6-2.el8.x86_64
>   libssh-0.9.4-3.el8.x86_64
>   libssh-config-0.9.4-3.el8.noarch
>   libsss_autofs-2.5.2-2.el8_5.3.x86_64
>   libsss_certmap-2.5.2-2.el8_5.3.x86_64
>   libsss_idmap-2.5.2-2.el8_5.3.x86_64
>   libsss_nss_idmap-2.5.2-2.el8_5.3.x86_64
>   libsss_sudo-2.5.2-2.el8_5.3.x86_64
>   libstdc++-8.5.0-4.el8_5.x86_64
>   libstoragemgmt-1.9.1-1.el8.x86_64
>   libtalloc-2.3.2-1.el8.x86_64
>   libtevent-0.11.0-0.el8.x86_64
>   libtirpc-1.1.4-5.el8.x86_64
>   libudisks2-2.9.0-7.el8.x86_64
>   libuuid-2.32.1-28.el8.x86_64
>   libwbclient-4.14.5-7.el8_5.x86_64
>   libxcrypt-4.1.1-6.el8.x86_64
>   libxml2-2.9.7-9.el8_4.2.x86_64
>   linux-firmware-20210702-103.gitd79c2677.el8.noarch
>   lshw-B.02.19.2-6.el8.x86_64
>   lsscsi-0.32-3.el8.x86_64
>   lua-libs-5.3.4-12.el8.x86_64
>   lvm2-8:2.03.12-10.el8.x86_64
>   lvm2-libs-8:2.03.12-10.el8.x86_64
>   lz4-libs-1.8.3-3.el8_4.x86_64
>   man-db-2.7.6.1-18.el8.x86_64
>   man-pages-overrides-8.5.0.1-1.el8.noarch
>   mcelog-3:175-1.el8.x86_64
>   mdadm-4.2-rc2.el8.x86_64
>   microcode_ctl-4:20210608-1.el8.x86_64
>   ncurses-6.1-9.20180224.el8.x86_64
>   ncurses-base-6.1-9.20180224.el8.noarch
>   ncurses-libs-6.1-9.20180224.el8.x86_64
.4.1-7.el8.x86_64
  nftables-1:0.9.3-21.el8.x86_64
  nmap-ncat-2:7.70-6.el8.x86_64
  nspr-4.32.0-1.el8_4.x86_64
  nss-3.67.0-7.el8_5.x86_64
  nss-softokn-3.67.0-7.el8_5.x86_64
  nss-softokn-freebl-3.67.0-7.el8_5.x86_64
  nss-sysinit-3.67.0-7.el8_5.x86_64
  nss-util-3.67.0-7.el8_5.x86_64
  numactl-libs-2.0.12-13.el8.x86_64
  nvme-cli-1.14-3.el8.x86_64
  open-vm-tools-11.2.5-2.el8.x86_64
  openldap-2.4.46-18.el8.x86_64
  openssh-8.0p1-10.el8.x86_64
  openssh-clients-8.0p1-10.el8.x86_64
  openssh-server-8.0p1-10.el8.x86_64
  openssl-1:1.1.1k-5.el8_5.x86_64
  openssl-libs-1:1.1.1k-5.el8_5.x86_64
  os-prober-1.74-9.el8.x86_64
  pam-1.3.1-15.el8.x86_64
  parted-3.2-39.el8.x86_64
  pcre-8.42-6.el8.x86_64
  platform-python-3.6.8-41.el8.x86_64
  platform-python-pip-9.0.3-20.el8.noarch
  plymouth-0.9.4-10.20200615git1e36e30.el8.x86_64
  plymouth-core-libs-0.9.4-10.20200615git1e36e30.el8.x86_64
  plymouth-scripts-0.9.4-10.20200615git1e36e30.el8.x86_64
  podman-3.3.1-9.module_el8.5.0+988+b1f0b741.x86_64
  podman-catatonit-3.3.1-9.module_el8.5.0+988+b1f0b741.x86_64
  policycoreutils-2.9-16.el8.x86_64
  policycoreutils-python-utils-2.9-16.el8.noarch
  polkit-0.115-12.el8.x86_64
  polkit-libs-0.115-12.el8.x86_64
  python3-bind-32:9.11.26-6.el8.noarch
  python3-dnf-4.7.0-4.el8.noarch
  python3-dnf-plugins-core-4.0.21-3.el8.noarch
  python3-firewall-0.9.3-7.el8.noarch
  python3-gpg-1.13.1-9.el8.x86_64
  python3-hawkey-0.63.0-3.el8.x86_64
  python3-libcomps-0.1.16-2.el8.x86_64
  python3-libdnf-0.63.0-3.el8.x86_64
  python3-libs-3.6.8-41.el8.x86_64
  python3-libstoragemgmt-1.9.1-1.el8.x86_64
  python3-libxml2-2.9.7-9.el8_4.2.x86_64
  python3-lxml-4.2.3-3.el8.x86_64
  python3-nftables-1:0.9.3-21.el8.x86_64
  python3-perf-4.18.0-348.7.1.el8_5.x86_64
  python3-pip-wheel-9.0.3-20.el8.noarch
  python3-policycoreutils-2.9-16.el8.noarch
  python3-psutil-5.4.3-11.el8.x86_64
  python3-rpm-4.14.3-19.el8.x86_64
  python3-sssdconfig-2.5.2-2.el8_5.3.noarch
  python3-syspurpose-1.28.21-3.el8.x86_64
  python3-unbound-1.7.3-17.el8.x86_64
  quota-1:4.04-14.el8.x86_64
  quota-nls-1:4.04-14.el8.noarch
  rdma-core-35.0-1.el8.x86_64
  realmd-0.16.3-23.el8.x86_64
  rpm-4.14.3-19.el8.x86_64
  rpm-build-libs-4.14.3-19.el8.x86_64
  rpm-libs-4.14.3-19.el8.x86_64
  rpm-plugin-selinux-4.14.3-19.el8.x86_64
  rpm-plugin-systemd-inhibit-4.14.3-19.el8.x86_64
  rsyslog-8.2102.0-5.el8.x86_64
  rsyslog-gnutls-8.2102.0-5.el8.x86_64
  rsyslog-gssapi-8.2102.0-5.el8.x86_64
  rsyslog-relp-8.2102.0-5.el8.x86_64
  runc-1.0.2-1.module_el8.5.0+911+f19012f9.x86_64
  samba-client-libs-4.14.5-7.el8_5.x86_64
  samba-common-4.14.5-7.el8_5.noarch
  samba-common-libs-4.14.5-7.el8_5.x86_64
  selinux-policy-3.14.3-80.el8_5.2.noarch
  selinux-policy-targeted-3.14.3-80.el8_5.2.noarch
  setroubleshoot-plugins-3.3.14-1.el8.noarch
  setroubleshoot-server-3.3.24-4.el8.x86_64
  shadow-utils-2:4.6-14.el8.x86_64
  slirp4netns-1.1.8-1.module_el8.5.0+890+6b136101.x86_64
  sos-4.1-5.el8.noarch
  sqlite-3.26.0-15.el8.x86_64
  sqlite-libs-3.26.0-15.el8.x86_64
  sssd-2.5.2-2.el8_5.3.x86_64
  sssd-ad-2.5.2-2.el8_5.3.x86_64
  sssd-client-2.5.2-2.el8_5.3.x86_64
  sssd-common-2.5.2-2.el8_5.3.x86_64
  sssd-common-pac-2.5.2-2.el8_5.3.x86_64
  sssd-ipa-2.5.2-2.el8_5.3.x86_64
  sssd-kcm-2.5.2-2.el8_5.3.x86_64
  sssd-krb5-2.5.2-2.el8_5.3.x86_64
  sssd-krb5-common-2.5.2-2.el8_5.3.x86_64
  sssd-ldap-2.5.2-2.el8_5.3.x86_64
  sssd-nfs-idmap-2.5.2-2.el8_5.3.x86_64
  sssd-proxy-2.5.2-2.el8_5.3.x86_64
  strace-5.7-3.el8.x86_64
  sudo-1.8.29-7.el8_4.1.x86_64
  systemd-239-51.el8_5.2.x86_64
  systemd-libs-239-51.el8_5.2.x86_64
  systemd-pam-239-51.el8_5.2.x86_64
  systemd-udev-239-51.el8_5.2.x86_64
  tcpdump-14:4.9.3-2.el8.x86_64
  tpm2-tools-4.1.1-5.el8.x86_64
  tpm2-tss-2.3.2-4.el8.x86_64
  tuned-2.16.0-1.el8.noarch
  tzdata-2021e-1.el8.noarch
  udisks2-2.9.0-7.el8.x86_64
  udisks2-iscsi-2.9.0-7.el8.x86_64
  udisks2-lvm2-2.9.0-7.el8.x86_64
  unbound-libs-1.7.3-17.el8.x86_64
  unzip-6.0-45.el8_4.x86_64
  util-linux-2.32.1-28.el8.x86_64
  util-linux-user-2.32.1-28.el8.x86_64
  vdo-6.2.5.74-14.el8.x86_64
  vim-com>   nettle-3.4.1-7.el8.x86_64
>   nftables-1:0.9.3-21.el8.x86_64
>   nmap-ncat-2:7.70-6.el8.x86_64
>   nspr-4.32.0-1.el8_4.x86_64
>   nss-3.67.0-7.el8_5.x86_64
>   nss-softokn-3.67.0-7.el8_5.x86_64
>   nss-softokn-freebl-3.67.0-7.el8_5.x86_64
>   nss-sysinit-3.67.0-7.el8_5.x86_64
>   nss-util-3.67.0-7.el8_5.x86_64
>   numactl-libs-2.0.12-13.el8.x86_64
>   nvme-cli-1.14-3.el8.x86_64
>   open-vm-tools-11.2.5-2.el8.x86_64
>   openldap-2.4.46-18.el8.x86_64
>   openssh-8.0p1-10.el8.x86_64
>   openssh-clients-8.0p1-10.el8.x86_64
>   openssh-server-8.0p1-10.el8.x86_64
>   openssl-1:1.1.1k-5.el8_5.x86_64
>   openssl-libs-1:1.1.1k-5.el8_5.x86_64
>   os-prober-1.74-9.el8.x86_64
>   pam-1.3.1-15.el8.x86_64
>   parted-3.2-39.el8.x86_64
>   pcre-8.42-6.el8.x86_64
>   platform-python-3.6.8-41.el8.x86_64
>   platform-python-pip-9.0.3-20.el8.noarch
mon-2:8.0.1>   plymouth-0.9.4-10.20200615git1e36e30.el8.x86_64
>   plymouth-core-libs-0.9.4-10.20200615git1e36e30.el8.x86_64
>   plymouth-scripts-0.9.4-10.20200615git1e36e30.el8.x86_64
>   podman-3.3.1-9.module_el8.5.0+988+b1f0b741.x86_64
>   podman-catatonit-3.3.1-9.module_el8.5.0+988+b1f0b741.x86_64
>   policycoreutils-2.9-16.el8.x86_64
>   policycoreutils-python-utils-2.9-16.el8.noarch
>   polkit-0.115-12.el8.x86_64
>   polkit-libs-0.115-12.el8.x86_64
>   python3-bind-32:9.11.26-6.el8.noarch
>   python3-dnf-4.7.0-4.el8.noarch
>   python3-dnf-plugins-core-4.0.21-3.el8.noarch
>   python3-firewall-0.9.3-7.el8.noarch
>   python3-gpg-1.13.1-9.el8.x86_64
>   python3-hawkey-0.63.0-3.el8.x86_64
>   python3-libcomps-0.1.16-2.el8.x86_64
>   python3-libdnf-0.63.0-3.el8.x86_64
>   python3-libs-3.6.8-41.el8.x86_64
>   python3-libstoragemgmt-1.9.1-1.el8.x86_64
>   python3-libxml2-2.9.7-9.el8_4.2.x86_64
>   python3-lxml-4.2.3-3.el8.x86_64
>   python3-nftables-1:0.9.3-21.el8.x86_64
>   python3-perf-4.18.0-348.7.1.el8_5.x86_64
>   python3-pip-wheel-9.0.3-20.el8.noarch
>   python3-policycoreutils-2.9-16.el8.noarch
>   python3-psutil-5.4.3-11.el8.x86_64
>   python3-rpm-4.14.3-19.el8.x86_64
>   python3-sssdconfig-2.5.2-2.el8_5.3.noarch
>   python3-syspurpose-1.28.21-3.el8.x86_64
>   python3-unbound-1.7.3-17.el8.x86_64
>   quota-1:4.04-14.el8.x86_64
>   quota-nls-1:4.04-14.el8.noarch
>   rdma-core-35.0-1.el8.x86_64
>   realmd-0.16.3-23.el8.x86_64
>   rpm-4.14.3-19.el8.x86_64
>   rpm-build-libs-4.14.3-19.el8.x86_64
>   rpm-libs-4.14.3-19.el8.x86_64
>   rpm-plugin-selinux-4.14.3-19.el8.x86_64
>   rpm-plugin-systemd-inhibit-4.14.3-19.el8.x86_64
>   rsyslog-8.2102.0-5.el8.x86_64
>   rsyslog-gnutls-8.2102.0-5.el8.x86_64
>   rsyslog-gssapi-8.2102.0-5.el8.x86_64
>   rsyslog-relp-8.2102.0-5.el8.x86_64
>   runc-1.0.2-1.module_el8.5.0+911+f19012f9.x86_64
>   samba-client-libs-4.14.5-7.el8_5.x86_64
>   samba-common-4.14.5-7.el8_5.noarch
>   samba-common-libs-4.14.5-7.el8_5.x86_64
>   selinux-policy-3.14.3-80.el8_5.2.noarch
>   selinux-policy-targeted-3.14.3-80.el8_5.2.noarch
>   setroubleshoot-plugins-3.3.14-1.el8.noarch
>   setroubleshoot-server-3.3.24-4.el8.x86_64
>   shadow-utils-2:4.6-14.el8.x86_64
>   slirp4netns-1.1.8-1.module_el8.5.0+890+6b136101.x86_64
>   sos-4.1-5.el8.noarch
>   sqlite-3.26.0-15.el8.x86_64
>   sqlite-libs-3.26.0-15.el8.x86_64
>   sssd-2.5.2-2.el8_5.3.x86_64
>   sssd-ad-2.5.2-2.el8_5.3.x86_64
>   sssd-client-2.5.2-2.el8_5.3.x86_64
>   sssd-common-2.5.2-2.el8_5.3.x86_64
>   sssd-common-pac-2.5.2-2.el8_5.3.x86_64
>   sssd-ipa-2.5.2-2.el8_5.3.x86_64
>   sssd-kcm-2.5.2-2.el8_5.3.x86_64
>   sssd-krb5-2.5.2-2.el8_5.3.x86_64
>   sssd-krb5-common-2.5.2-2.el8_5.3.x86_64
>   sssd-ldap-2.5.2-2.el8_5.3.x86_64
>   sssd-nfs-idmap-2.5.2-2.el8_5.3.x86_64
>   sssd-proxy-2.5.2-2.el8_5.3.x86_64
>   strace-5.7-3.el8.x86_64
>   sudo-1.8.29-7.el8_4.1.x86_64
>   systemd-239-51.el8_5.2.x86_64
>   systemd-libs-239-51.el8_5.2.x86_64
>   systemd-pam-239-51.el8_5.2.x86_64
>   systemd-udev-239-51.el8_5.2.x86_64
>   tcpdump-14:4.9.3-2.el8.x86_64
>   tpm2-tools-4.1.1-5.el8.x86_64
>   tpm2-tss-2.3.2-4.el8.x86_64
>   tuned-2.16.0-1.el8.noarch
>   tzdata-2021e-1.el8.noarch
>   udisks2-2.9.0-7.el8.x86_64
>   udisks2-iscsi-2.9.0-7.el8.x86_64
>   udisks2-lvm2-2.9.0-7.el8.x86_64
>   unbound-libs-1.7.3-17.el8.x86_64
>   unzip-6.0-45.el8_4.x86_64
>   util-linux-2.32.1-28.el8.x86_64
>   util-linux-user-2.32.1-28.el8.x86_64
>   vdo-6.2.5.74-14.el8.x86_64
>   vim-common-2:8.0.1763-16.el8.x86_64
>   vim-enhanced-2:8.0.1763-16.el8.x86_64
>   vim-filesystem-2:8.0.1763-16.el8.noarch
>   vim-minimal-2:8.0.1763-16.el8.x86_64
>   virt-what-1.18-12.el8.x86_64
>   which-2.21-16.el8.x86_64
>   xfsprogs-5.0.0-9.el8.x86_64
>   yum-4.7.0-4.el8.noarch
> Installed:
>   grub2-tools-efi-1:2.02-106.el8.x86_64
>   kernel-4.18.0-348.7.1.el8_5.x86_64
>   kernel-core-4.18.0-348.7.1.el8_5.x86_64
>   kernel-modules-4.18.0-348.7.1.el8_5.x86_64
>   libbpf-0.4.0-1.el8.x86_64
>
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
> ============================== Name Matched: apache ===============================
> apache-commons-cli.noarch : Command Line Interface Library for Java
> apache-commons-codec.noarch : Implementations of common encoders and decoders
> apache-commons-io.noarch : Utilities to assist with developing IO functionality
> apache-commons-lang3.noarch : Provides a host of helper utilities for the java.lang
>                             : API
> ============================= Summary Matched: apache =============================
> apr.i686 : Apache Portable Runtime library
> apr.x86_64 : Apache Portable Runtime library
> apr-util.i686 : Apache Portable Runtime Utility library
> apr-util.x86_64 : Apache Portable Runtime Utility library
> httpd.x86_64 : Apache HTTP Server
> httpd-devel.x86_64 : Development interfaces for the Apache HTTP server
> httpd-filesystem.noarch : The basic directory layout for the Apache HTTP server
> httpd-manual.noarch : Documentation for the Apache HTTP server
> httpd-tools.x86_64 : Tools for use with the Apache HTTP Server
> keycloak-httpd-client-install.noarch : Tools to configure Apache HTTPD as Keycloak
>                                      : client
> librdkafka.i686 : The Apache Kafka C library
> librdkafka.x86_64 : The Apache Kafka C library
> mod_auth_gssapi.x86_64 : A GSSAPI Authentication module for Apache
> mod_auth_mellon.x86_64 : A SAML 2.0 authentication module for the Apache Httpd
>                        : Server
> mod_dav_svn.x86_64 : Apache httpd module for Subversion server
> mod_fcgid.x86_64 : FastCGI interface module for Apache 2
> mod_http2.x86_64 : module implementing HTTP/2 for Apache 2
> mod_intercept_form_submit.x86_64 : Apache module to intercept login form submission
>                                  : and run PAM authentication
> mod_ldap.x86_64 : LDAP authentication modules for the Apache HTTP Server
> mod_lookup_identity.x86_64 : Apache module to retrieve additional information about
>                            : the authenticated user
> mod_md.x86_64 : Certificate provisioning using ACME for the Apache HTTP Server
> mod_proxy_html.x86_64 : HTML and XML content filters for the Apache HTTP Server
> mod_security.x86_64 : Security module for the Apache HTTP Server
> mod_session.x86_64 : Session interface for the Apache HTTP Server
> mod_ssl.x86_64 : SSL/TLS module for the Apache HTTP Server
> pcp-export-pcp2spark.x86_64 : Performance Co-Pilot tools for exporting PCP metrics
>                             : to Apache Spark
> python3-keycloak-httpd-client-install.noarch : Tools to configure Apache HTTPD as
>                                              : Keycloak client
> python3-mod_wsgi.x86_64 : A WSGI interface for Python web applications in Apache
> python38-mod_wsgi.x86_64 : A WSGI interface for Python web applications in Apache
> python39-mod_wsgi.x86_64 : A WSGI interface for Python web applications in Apache
> [osboxes@osboxes ~]$ sudo dnf install apache http server
> Last metadata expiration check: 7:50:30 ago on Tue Feb  8 01:51:14 2022.
> No match for argument: apache
> No match for argument: http
> No match for argument: server
> Error: Unable to find a match: apache http server
> [osboxes@osboxes ~]$ sudo dnf install httpd
> Last metadata expiration check: 7:50:50 ago on Tue Feb  8 01:51:14 2022.
> Dependencies resolved.
> ===================================================================================
>  Package            Arch   Version                                 Repo       Size
> ===================================================================================
> Installing:
>  httpd              x86_64 2.4.37-43.module_el8.5.0+1022+b541f3b1  appstream 1.4 M
> Installing dependencies:
>  apr                x86_64 1.6.3-12.el8                            appstream 129 k
>  apr-util           x86_64 1.6.1-6.el8                             appstream 105 k
>  centos-logos-httpd noarch 85.8-2.el8                              baseos     75 k
>  httpd-filesystem   noarch 2.4.37-43.module_el8.5.0+1022+b541f3b1  appstream  39 k
>  httpd-tools        x86_64 2.4.37-43.module_el8.5.0+1022+b541f3b1  appstream 107 k
>  mod_http2          x86_64 1.15.7-3.module_el8.4.0+778+c970deab    appstream 154 k
> Installing weak dependencies:
>  apr-util-bdb       x86_64 1.6.1-6.el8                             appstream  25 k
>  apr-util-openssl   x86_64 1.6.1-6.el8                             appstream  27 k
> Enabling module streams:
>  httpd                     2.4
>
> Transaction Summary
> ===================================================================================
> Install  9 Packages
>
> Total download size: 2.1 M
> Installed size: 5.6 M
> Is this ok [y/N]: y
> Downloading Packages:
> (1/9): apr-util-bdb-1.6.1-6.el8.x86_64.rpm          11 kB/s |  25 kB     00:02
> (2/9): apr-util-1.6.1-6.el8.x86_64.rpm              37 kB/s | 105 kB     00:02
> (3/9): apr-util-openssl-1.6.1-6.el8.x86_64.rpm      43 kB/s |  27 kB     00:00
> (4/9): apr-1.6.3-12.el8.x86_64.rpm                  46 kB/s | 129 kB     00:02
> (5/9): httpd-filesystem-2.4.37-43.module_el8.5.0+1  21 kB/s |  39 kB     00:01
> (6/9): httpd-2.4.37-43.module_el8.5.0+1022+b541f3b 568 kB/s | 1.4 MB     00:02
> (7/9): mod_http2-1.15.7-3.module_el8.4.0+778+c970d 185 kB/s | 154 kB     00:00
> (8/9): centos-logos-httpd-85.8-2.el8.noarch.rpm    135 kB/s |  75 kB     00:00
> (9/9): httpd-tools-2.4.37-43.module_el8.5.0+1022+b  34 kB/s | 107 kB     00:03
> -----------------------------------------------------------------------------------
> Total                                              349 kB/s | 2.1 MB     00:06
> Running transaction check
> Transaction check succeeded.
> Running transaction test
> Transaction test succeeded.
> Running transaction
>   Preparing        :                                                           1/1
>   Installing       : apr-1.6.3-12.el8.x86_64                                   1/9
>   Running scriptlet: apr-1.6.3-12.el8.x86_64                                   1/9
>   Installing       : apr-util-bdb-1.6.1-6.el8.x86_64                           2/9
>   Installing       : apr-util-openssl-1.6.1-6.el8.x86_64                       3/9
>   Installing       : apr-util-1.6.1-6.el8.x86_64                               4/9
>   Running scriptlet: apr-util-1.6.1-6.el8.x86_64                               4/9
>   Installing       : httpd-tools-2.4.37-43.module_el8.5.0+1022+b541f3b1.x86_   5/9
>   Installing       : centos-logos-httpd-85.8-2.el8.noarch                      6/9
>   Running scriptlet: httpd-filesystem-2.4.37-43.module_el8.5.0+1022+b541f3b1   7/9
>   Installing       : httpd-filesystem-2.4.37-43.module_el8.5.0+1022+b541f3b1   7/9
>   Installing       : mod_http2-1.15.7-3.module_el8.4.0+778+c970deab.x86_64     8/9
>   Installing       : httpd-2.4.37-43.module_el8.5.0+1022+b541f3b1.x86_64       9/9
>   Running scriptlet: httpd-2.4.37-43.module_el8.5.0+1022+b541f3b1.x86_64       9/9
>   Verifying        : apr-1.6.3-12.el8.x86_64                                   1/9
>   Verifying        : apr-util-1.6.1-6.el8.x86_64                               2/9
>   Verifying        : apr-util-bdb-1.6.1-6.el8.x86_64                           3/9
>   Verifying        : apr-util-openssl-1.6.1-6.el8.x86_64                       4/9
>   Verifying        : httpd-2.4.37-43.module_el8.5.0+1022+b541f3b1.x86_64       5/9
>   Verifying        : httpd-filesystem-2.4.37-43.module_el8.5.0+1022+b541f3b1   6/9
>   Verifying        : httpd-tools-2.4.37-43.module_el8.5.0+1022+b541f3b1.x86_   7/9
>   Verifying        : mod_http2-1.15.7-3.module_el8.4.0+778+c970deab.x86_64     8/9
>   Verifying        : centos-logos-httpd-85.8-2.el8.noarch                      9/9
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
> drwxr-xr-x. 105 root root     8192 Feb  8 09:42 .
> dr-xr-xr-x.  17 root root      224 Feb  8 03:40 ..
> -rw-------.   1 root root        0 Jun 19  2021 .pwd.lock
> -rw-r--r--.   1 root root      208 Jun 19  2021 .updated
> -rw-r--r--.   1 root root     4536 Jul 14  2021 DIR_COLORS
> -rw-r--r--.   1 root root     5214 Jul 14  2021 DIR_COLORS.256color
> -rw-r--r--.   1 root root     4618 Jul 14  2021 DIR_COLORS.lightbgcolor
> -rw-r--r--.   1 root root       94 May 10  2019 GREP_COLORS
> drwxr-xr-x.   7 root root      134 Feb  8 03:46 NetworkManager
> drwxr-xr-x.   2 root root       48 Jun 19  2021 PackageKit
> drwxr-xr-x.   6 root root       70 Jun 22  2021 X11
> -rw-r--r--.   1 root root       16 Jun 19  2021 adjtime
> -rw-r--r--.   1 root root     1529 May 15  2020 aliases
> drwxr-xr-x.   2 root root      194 Jul 28  2021 alternatives
> -rw-r--r--.   1 root root      541 Nov  8  2019 anacrontab
> -rw-r--r--.   1 root root        1 May 11  2019 at.deny
> drwxr-x---.   4 root root      100 Jun 19  2021 audit
> drwxr-xr-x.   3 root root      228 Feb  8 09:35 authselect
> drwxr-xr-x.   2 root root      135 Feb  8 03:49 bash_completion.d
> -rw-r--r--.   1 root root     3019 May 15  2020 bashrc
> -rw-r--r--.   1 root root      535 Apr 21  2021 bindresvport.blacklist
> drwxr-xr-x.   2 root root        6 Dec 21 15:08 binfmt.d
> -rw-r--r--.   1 root root       30 Nov  9 19:17 centos-release
> -rw-r--r--.   1 root root       42 Nov  9 19:17 centos-release-upstream
> drwxr-xr-x.   2 root root        6 Jul 28  2021 chkconfig.d
> -rw-r--r--.   1 root root     1085 Jun 24  2021 chrony.conf
> -rw-r-----.   1 root chrony    540 May 12  2021 chrony.keys
> drwxr-xr-x.   2 root root       26 Dec 21 15:14 cifs-utils
> drwxr-xr-x.   3 root root       19 Jun 19  2021 cni
> drwxr-xr-x.   4 root root       42 Sep 10 10:51 cockpit
> drwxr-xr-x.   6 root root      139 Feb  8 03:48 containers
> drwxr-xr-x.   2 root root       39 Jun 19  2021 cron.d
> drwxr-xr-x.   2 root root       23 Jun 19  2021 cron.daily
> -rw-r--r--.   1 root root        0 Nov  8  2019 cron.deny
> drwxr-xr-x.   2 root root       22 Jun 19  2021 cron.hourly
> drwxr-xr-x.   2 root root        6 Jan 12  2021 cron.monthly
> drwxr-xr-x.   2 root root        6 Jan 12  2021 cron.weekly
> -rw-r--r--.   1 root root      451 Jan 12  2021 crontab
> drwxr-xr-x.   6 root root       81 Jun 19  2021 crypto-policies
> -rw-------.   1 root root        0 Jun 19  2021 crypttab
> -rw-r--r--.   1 root root     1629 May 15  2020 csh.cshrc
> -rw-r--r--.   1 root root     1078 May 15  2020 csh.login
> drwxr-xr-x.   4 root root       78 May  8  2021 dbus-1
> drwxr-xr-x.   3 root root       16 Jun 19  2021 dconf
> drwxr-xr-x.   2 root root       33 Feb  8 03:48 default
> drwxr-xr-x.   2 root root       40 Feb  8 03:46 depmod.d
> drwxr-x---.   3 root root       45 Jul 22  2021 dhcp
> drwxr-xr-x.   8 root root      128 Sep 17 15:05 dnf
> -rw-r--r--.   1 root root      117 Sep 29 17:07 dracut.conf
> drwxr-xr-x.   2 root root        6 Sep 29 17:07 dracut.conf.d
> -rw-r--r--.   1 root root        0 May 15  2020 environment
> -rw-r--r--.   1 root root     1362 Aug 24 19:13 ethertypes
> -rw-r--r--.   1 root root        0 Sep 10  2018 exports
> lrwxrwxrwx.   1 root root       56 Nov 12 13:19 favicon.png -> /usr/share/icons/hicolor/16x16/apps/fedora-logo-icon.png
> -rw-r--r--.   1 root root       66 Sep 10  2018 filesystems
> drwxr-x---.   8 root root      149 Feb  8 03:48 firewalld
> drwxr-xr-x.   3 root root       38 Jun 19  2021 fonts
> -rw-r--r--.   1 root root       20 Jan 13  2021 fprintd.conf
> -rw-r--r--.   1 root root      709 Jun 19  2021 fstab
> -rw-r--r--.   1 root root       38 May 11  2019 fuse.conf
> drwxr-xr-x.   2 root root       25 Jun 29  2021 gcrypt
> drwxr-xr-x.   2 root root        6 May 15  2020 gnupg
> drwxr-xr-x.   4 root root       40 Jun 19  2021 groff
> -rw-r--r--.   1 root root      691 Feb  8 09:42 group
> -rw-r--r--.   1 root root      678 Jun 19  2021 group-
> drwx------.   2 root root      288 Feb  8 03:44 grub.d
> lrwxrwxrwx.   1 root root       22 Nov  8 01:39 grub2.cfg -> ../boot/grub2/grub.cfg
> ----------.   1 root root      555 Feb  8 09:42 gshadow
> ----------.   1 root root      544 Jun 19  2021 gshadow-
> drwxr-xr-x.   3 root root       20 Aug 26 01:12 gss
> -rw-r--r--.   1 root root        9 Sep 10  2018 host.conf
> -rw-r--r--.   1 root root        8 Feb  5 02:10 hostname
> -rw-r--r--.   1 root root      158 Sep 10  2018 hosts
> drwxr-xr-x.   5 root root      105 Feb  8 09:42 httpd
> -rw-r--r--.   1 root root     4849 Jul 30  2021 idmapd.conf
> lrwxrwxrwx.   1 root root       11 Jul 28  2021 init.d -> rc.d/init.d
> -rw-r--r--.   1 root root      490 Dec 21 15:08 inittab
> -rw-r--r--.   1 root root      942 Sep 10  2018 inputrc
> drwxr-xr-x.   2 root root      159 Oct 15 13:08 iproute2
> drwxr-xr-x.   2 root root       52 Aug 10 01:10 iscsi
> -rw-r--r--.   1 root root       23 Nov  9 19:17 issue
> drwxr-xr-x.   2 root root       27 Jun 19  2021 issue.d
> -rw-r--r--.   1 root root       22 Nov  9 19:17 issue.net
> drwxr-xr-x.   4 root root       33 Dec 21 15:15 kdump
> -rw-r--r--.   1 root root     8550 Feb  8 03:47 kdump.conf
> drwxr-xr-x.   4 root root       41 Dec 21 15:08 kernel
> -rw-r--r--.   1 root root      812 Aug 26 01:05 krb5.conf
> drwxr-xr-x.   2 root root       55 Feb  8 03:48 krb5.conf.d
> -rw-r--r--.   1 root root    21275 Feb  8 09:42 ld.so.cache
> -rw-r--r--.   1 root root       28 Aug 24 19:10 ld.so.conf
> drwxr-xr-x.   2 root root      129 Feb  8 03:44 ld.so.conf.d
> -rw-r-----.   1 root root      191 Nov  4  2019 libaudit.conf
> drwxr-xr-x.   3 root root       20 Jul  2  2021 libblockdev
> drwxr-xr-x.   2 root root      246 May 17  2021 libibverbs.d
> drwxr-xr-x.   2 root root       35 Jun 19  2021 libnl
> drwxr-xr-x.   6 root root       70 Jun 19  2021 libreport
> drwxr-xr-x.   2 root root       62 May  5  2021 libssh
> -rw-r--r--.   1 root root     2391 Jul 23  2015 libuser.conf
> -rw-r--r--.   1 root root       19 Jun 19  2021 locale.conf
> lrwxrwxrwx.   1 root root       38 Feb  5 02:29 localtime -> ../usr/share/zoneinfo/America/New_York
> -rw-r--r--.   1 root root     2512 Aug 18 15:04 login.defs
> -rw-r--r--.   1 root root      438 Feb 19  2018 logrotate.conf
> drwxr-xr-x.   2 root root      188 Feb  8 09:42 logrotate.d
> drwxr-xr-x.   3 root root       43 Apr 22  2021 lsm
> drwxr-xr-x.   6 root root      100 Feb  8 03:46 lvm
> -r--r--r--.   1 root root       33 Jun 19  2021 machine-id
> -rw-r--r--.   1 root root      111 May  8  2021 magic
> -rw-r--r--.   1 root root      272 May 11  2017 mailcap
> -rw-r--r--.   1 root root     5122 Dec 21 15:15 makedumpfile.conf.sample
> -rw-r--r--.   1 root root     5165 Jun 29  2021 man_db.conf
> drwxr-xr-x.   2 root root      181 Feb  5 05:47 mc
> drwxr-xr-x.   3 root root       41 Aug  9  2021 mcelog
> drwxr-xr-x.   3 root root       32 Jul 24  2021 microcode_ctl
> -rw-r--r--.   1 root root    60352 May 11  2017 mime.types
> -rw-r--r--.   1 root root     1108 Jun 24  2021 mke2fs.conf
> drwxr-xr-x.   2 root root       93 Feb  8 03:48 modprobe.d
> drwxr-xr-x.   2 root root        6 Dec 21 15:08 modules-load.d
> -rw-r--r--.   1 root root        0 Sep 10  2018 motd
> drwxr-xr-x.   2 root root       21 Jun 19  2021 motd.d
> lrwxrwxrwx.   1 root root       19 Jun 19  2021 mtab -> ../proc/self/mounts
> drwxr-xr-x.   2 root root        6 Jul 28  2021 multipath
> -rw-r--r--.   1 root root     9450 May 11  2019 nanorc
> -rw-r--r--.   1 root root      767 Apr 21  2021 netconfig
> -rw-r--r--.   1 root root       58 Sep 10  2018 networks
> drwx------.   3 root root       66 Aug 24 19:25 nftables
> lrwxrwxrwx.   1 root root       29 Feb  8 09:35 nsswitch.conf -> /etc/authselect/nsswitch.conf
> -rw-r--r--.   1 root root     2197 Mar 11  2021 nsswitch.conf.bak
> drwxr-xr-x.   2 root root       35 Sep 17 15:06 nvme
> drwxr-xr-x.   2 root root        6 Dec 16  2020 oddjob
> -rw-r--r--.   1 root root     4922 Dec 16  2020 oddjobd.conf
> drwxr-xr-x.   2 root root       70 Jun 19  2021 oddjobd.conf.d
> drwxr-xr-x.   3 root root       36 Aug 10 01:12 openldap
> drwxr-xr-x.   2 root root        6 Jun 22  2021 opt
> lrwxrwxrwx.   1 root root       21 Nov  9 19:17 os-release -> ../usr/lib/os-release
> drwxr-xr-x.   2 root root     4096 Feb  8 09:35 pam.d
> -rw-r--r--.   1 root root     1594 Feb  8 09:42 passwd
> -rw-r--r--.   1 root root     1541 Jun 19  2021 passwd-
> -rw-r--r--.   1 root root     2872 May 14  2019 pinforc
> drwxr-xr-x.   3 root root       21 Jun 19  2021 pkcs11
> drwxr-xr-x.   8 root root       88 Jun 22  2021 pki
> drwxr-xr-x.   2 root root       28 Aug 24 19:28 plymouth
> drwxr-xr-x.   5 root root       52 Jun 22  2021 pm
> drwxr-xr-x.   5 root root       72 Jun  2  2021 polkit-1
> drwxr-xr-x.   2 root root        6 Jan 19  2021 popt.d
> drwxr-xr-x.   2 root root       24 Feb  8 03:43 prelink.conf.d
> -rw-r--r--.   1 root root      233 Sep 10  2018 printcap
> -rw-r--r--.   1 root root     2123 May 15  2020 profile
> drwxr-xr-x.   2 root root     4096 Feb  8 03:41 profile.d
> -rw-r--r--.   1 root root     6568 Sep 10  2018 protocols
> drwxr-xr-x.  10 root root      127 Jul 28  2021 rc.d
> lrwxrwxrwx.   1 root root       13 Dec 21 15:08 rc.local -> rc.d/rc.local
> lrwxrwxrwx.   1 root root       10 Jul 28  2021 rc0.d -> rc.d/rc0.d
> lrwxrwxrwx.   1 root root       10 Jul 28  2021 rc1.d -> rc.d/rc1.d
> lrwxrwxrwx.   1 root root       10 Jul 28  2021 rc2.d -> rc.d/rc2.d
> lrwxrwxrwx.   1 root root       10 Jul 28  2021 rc3.d -> rc.d/rc3.d
> lrwxrwxrwx.   1 root root       10 Jul 28  2021 rc4.d -> rc.d/rc4.d
> lrwxrwxrwx.   1 root root       10 Jul 28  2021 rc5.d -> rc.d/rc5.d
> lrwxrwxrwx.   1 root root       10 Jul 28  2021 rc6.d -> rc.d/rc6.d
> drwxr-xr-x.   3 root root       38 May 17  2021 rdma
> lrwxrwxrwx.   1 root root       14 Nov  9 19:17 redhat-release -> centos-release
> -rw-r--r--.   1 root root       54 Feb  7 22:08 resolv.conf
> drwxr-xr-x.   3 root root       24 Jun 19  2021 rhsm
> -rw-r--r--.   1 root root     1634 Aug  1  2018 rpc
> drwxr-xr-x.   2 root root       25 Feb  8 03:48 rpm
> -rw-r--r--.   1 root root     3186 Aug 10 12:43 rsyslog.conf
> drwxr-xr-x.   2 root root        6 Aug 10 12:46 rsyslog.d
> drwxr-xr-x.   2 root root       35 Dec 21 15:14 rwtab.d
> drwxr-xr-x.   2 root root       61 Feb  8 03:44 samba
> drwxr-xr-x.   2 root root        6 May 15  2020 sasl2
> drwxr-xr-x.   7 root root     4096 May  7  2021 security
lrwxrwxrwx.   1 root root   19 Nov 11 23:58 logs -> ../../var/log/httpd
lrwxrwxrwx.   1 root root   29 Nov 11 23:58 modules -> ../../usr/lib64/httpd/modules
lrwxrwxrwx.   1 root root   10 Nov 11 23:58 run -> /run/httpd
lrwxrwxrwx.   1 root root   19 Nov 11 23:58 state -> ../../var/lib/httpd
[root@osboxes h> drwxr-xr-x.   3 root root       57 Dec 21 15:17 selinux
> -rw-r--r--.   1 root root   692252 May 15  2020 services
> -rw-r--r--.   1 root root      216 Sep 30 01:08 sestatus.conf
> drwxr-xr-x.   2 root root       33 Sep 30 01:08 setroubleshoot
> ----------.   1 root root      972 Feb  8 09:42 shadow
> ----------.   1 root root      950 Jun 19  2021 shadow-
> -rw-r--r--.   1 root root       44 Sep 10  2018 shells
> drwxr-xr-x.   2 root root       62 Jun 22  2021 skel
> drwxr-xr-x.   3 root root       74 Jun 19  2021 smartmontools
> drwxr-xr-x.   6 root root       86 Feb  8 03:46 sos
> drwxr-xr-x.   3 root root      245 Jul 13  2021 ssh
> drwxr-xr-x.   2 root root       19 Feb  8 03:43 ssl
> drwx------.   4 sssd sssd       31 Dec 21 15:14 sssd
> -rw-r--r--.   1 root root       21 Jun 19  2021 subgid
> -rw-r--r--.   1 root root        0 Sep 10  2018 subgid-
> -rw-r--r--.   1 root root       21 Jun 19  2021 subuid
> -rw-r--r--.   1 root root        0 Sep 10  2018 subuid-
> -rw-r-----.   1 root root     3181 Oct 25 10:27 sudo-ldap.conf
> -rw-r-----.   1 root root     1786 Oct 25 10:27 sudo.conf
> -r--r-----.   1 root root     4328 Oct 25 10:27 sudoers
> drwxr-x---.   2 root root        6 Oct 25 10:29 sudoers.d
> drwxr-xr-x.   5 root root     4096 Feb  8 09:42 sysconfig
> -rw-r--r--.   1 root root      449 Dec 21 15:08 sysctl.conf
> drwxr-xr-x.   2 root root       28 Dec 21 15:08 sysctl.d
> lrwxrwxrwx.   1 root root       14 Nov  9 19:17 system-release -> centos-release
> -rw-r--r--.   1 root root       23 Nov  9 19:17 system-release-cpe
> drwxr-xr-x.   4 root root      150 Feb  8 03:44 systemd
> -rw-r-----.   1 root tss      7046 Nov 16  2020 tcsd.conf
> drwxr-xr-x.   2 root root        6 Jun  1  2021 terminfo
> drwxr-xr-x.   2 root root        6 Dec 21 15:08 tmpfiles.d
> -rw-r--r--.   1 root root      375 Aug 24 19:20 trusted-key.key
> drwxr-xr-x.   3 root root      136 Feb  8 03:48 tuned
> drwxr-xr-x.   4 root root       68 Feb  8 09:37 udev
ttpd]# cat conf.d
cat: conf.d: Is a directory
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
===================================================================================
Installing:
 kernel                        x86_64 4.18.0-348.7.1.el8_5         baseos    7.0 M
 kernel-core                   x86_64 4.18.0-348.7.1.el8_5         baseos     38 M
 kernel-modules                x86_64 4.18.0-348.7.1.el8_5         baseos     30 M
Upgrading:
 NetworkManager                x86_64 1:1.32.10-4.el8              baseos    2.6 M
 NetworkManager-config-server  noarch 1:1.32.10-4.el8              baseos    131 k
 NetworkManager-libnm          x86_64 1:1.32.10-4.el8              baseos    1.8 M
 NetworkManager-team           x86_64 1:1.32.10-4.el8              baseos    148 k
 NetworkManager-tui            x86_64 1:1.32.10-4.el8              baseos    336 k
 adcli                         x86_64 0.8.2-12.el8                 baseos    118 k
 alsa-sof-firmware             noarch 1.8-1.el8                    baseos    676 k
 authselect                    x86_64 1.2.2-3.el8                  baseos    133 k
 authselect-compat             x86_64 1.2.2-3.el8                  appstream  38 k
 authselect-libs               x86_64 1.2.2-3.el8                  baseos    222 k
 bash                          x86_64 4.4.20-2.el8                 baseos    1.5 M
 bind-export-libs              x86_64 32:9.11.26-6.el8             baseos    1.1 M
 bind-libs                     x86_64 32:9.11.26-6.el8             appstream 174 k
 bind-libs-lite                x86_64 32:9.11.26-6.el8             appstream 1.2 M
 bind-license                  noarch 32:9.11.26-6.el8             appstream 102 k
 bind-utils                    x86_64 32:9.11.26-6.el8             appstream 451 k
 binutils                      x86_64 2.30-108.el8_5.1             baseos    5.8 M
 bpftool                       x86_64 4.18.0-348.7.1.el8_5         baseos    7.7 M
 buildah                       x86_64 1.22.3-2.module_el8.5.0+911+f19012f9
                                                                   appstream 7.7 M
 ca-certificates               noarch 2021.2.50-80.0.el8_4         baseos    390 k
 centos-gpg-keys               noarch 1:8-3.el8                    baseos     12 k
 centos-linux-release          noarch 8.5-1.2111.el8               baseos     22 k
 centos-linux-repos            noarch 8-3.el8                      baseos     20 k
 centos-logos                  x86_64 85.8-2.el8                   baseos    700 k
 chkconfig                     x86_64 1.19.1-1.el8                 baseos    198 k
 chrony                        x86_64 4.1-1.el8                    baseos    327 k
 cockpit                       x86_64 251.1-1.el8                  baseos     78 k
 cockpit-bridge                x86_64 251.1-1.el8                  baseos    539 k
 cockpit-packagekit            noarch 251.1-1.el8                  appstream 703 k
 cockpit-podman                noarch 33-1.module_el8.5.0+890+6b136101
                                                                   appstream 437 k
 cockpit-storaged              noarch 251.1-1.el8                  appstream 655 k
 cockpit-system                noarch 251.1-1.el8                  baseos    3.2 M
 cockpit-ws                    x86_64 251.1-1.el8                  baseos    1.3 M
 conmon                        x86_64 2:2.0.29-1.module_el8.5.0+890+6b136101
                                                                   appstream  52 k
 container-selinux             noarch 2:2.167.0-1.module_el8.5.0+911+f19012f9
                                                                   appstream  54 k
 containernetworking-plugins   x86_64 1.0.0-1.module_el8.5.0+890+6b136101
                                                                   appstream  19 M
 containers-common             noarch 2:1-2.module_el8.5.0+890+6b136101
                                                                   appstream  79 k
 coreutils                     x86_64 8.30-12.el8                  baseos    1.2 M
 coreutils-common              x86_64 8.30-12.el8                  baseos    2.0 M
 criu                          x86_64 3.15-3.module_el8.5.0+890+6b136101
                                                                   appstream 518 k
 crypto-policies               noarch 20210617-1.gitc776d3e.el8    baseos     63 k
 crypto-policies-scripts       noarch 20210617-1.gitc776d3e.el8    baseos     83 k
 cups-libs                     x86_64 1:2.2.6-40.el8               baseos    433 k
 curl                          x86_64 7.61.1-22.el8                baseos    351 k
 dbus                          x86_64 1:1.12.8-14.el8              baseos     41 k
 dbus-common                   noarch 1:1.12.8-14.el8              baseos     46 k
 dbus-daemon                   x86_64 1:1.12.8-14.el8              baseos    240 k
 dbus-libs                     x86_64 1:1.12.8-14.el8              baseos    184 k
 dbus-tools                    x86_64 1:1.12.8-14.el8              baseos     85 k
 device-mapper                 x86_64 8:1.02.177-10.el8            baseos    377 k
 device-mapper-event           x86_64 8:1.02.177-10.el8            baseos    271 k
 device-mapper-event-libs      x86_64 8:1.02.177-10.el8            baseos    270 k
 device-mapper-libs            x86_64 8:1.02.177-10.el8            baseos    409 k
 device-mapper-multipath       x86_64 0.8.4-17.el8                 baseos    197 k
 device-mapper-multipath-libs  x86_64 0.8.4-17.el8                 baseos    322 k
 device-mapper-persistent-data x86_64 0.9.0-4.el8                  baseos    925 k
 dhcp-client                   x86_64 12:4.3.6-45.el8              baseos    318 k
 dhcp-common                   noarch 12:4.3.6-45.el8              baseos    207 k
 dhcp-libs                     x86_64 12:4.3.6-45.el8              baseos    148 k
 dmidecode                     x86_64 1:3.2-10.el8                 baseos     91 k
 dnf                           noarch 4.7.0-4.el8                  baseos    544 k
 dnf-data                      noarch 4.7.0-4.el8                  baseos    154 k
 dnf-plugins-core              noarch 4.0.21-3.el8                 baseos     70 k
 dracut                        x86_64 049-191.git20210920.el8      baseos    374 k
 dracut-config-rescue          x86_64 049-191.git20210920.el8      baseos     61 k
 dracut-network                x86_64 049-191.git20210920.el8      baseos    108 k
 dracut-squash                 x86_64 049-191.git20210920.el8      baseos     61 k
 e2fsprogs                     x86_64 1.45.6-2.el8                 baseos    1.0 M
 e2fsprogs-libs                x86_64 1.45.6-2.el8                 baseos    233 k
 elfutils-debuginfod-client    x86_64 0.185-1.el8                  baseos     66 k
 elfutils-default-yama-scope   noarch 0.185-1.el8                  baseos     49 k
 elfutils-libelf               x86_64 0.185-1.el8                  baseos    221 k
 elfutils-libs                 x86_64 0.185-1.el8                  baseos    292 k
 emacs-filesystem              noarch 1:26.1-7.el8                 baseos     70 k
 ethtool                       x86_64 2:5.8-7.el8                  baseos    209 k
 file                          x86_64 5.33-20.el8                  baseos     77 k
 file-libs                     x86_64 5.33-20.el8                  baseos    543 k
 filesystem                    x86_64 3.8-6.el8                    baseos    1.1 M
 firewalld                     noarch 0.9.3-7.el8                  baseos    502 k
 firewalld-filesystem          noarch 0.9.3-7.el8                  baseos     77 k
 fontconfig                    x86_64 2.13.1-4.el8                 baseos    274 k
 fstrm                         x86_64 0.6.1-2.el8                  appstream  29 k
 fuse-overlayfs                x86_64 1.7.1-1.module_el8.5.0+890+6b136101
                                                                   appstream  73 k
 glib2                         x86_64 2.56.4-156.el8               baseos    2.5 M
 glibc                         x86_64 2.28-164.el8                 baseos    3.6 M
 glibc-common                  x86_64 2.28-164.el8                 baseos    1.3 M
 glibc-langpack-en             x86_64 2.28-164.el8                 baseos    828 k
 gnutls                        x86_64 3.6.16-4.el8                 baseos    1.0 M
 gpgme                         x86_64 1.13.1-9.el8                 baseos    336 k
 grub2-common                  noarch 1:2.02-106.el8               baseos    891 k
 grub2-pc                      x86_64 1:2.02-106.el8               baseos     42 k
 grub2-pc-modules              noarch 1:2.02-106.el8               baseos    916 k
 grub2-tools                   x86_64 1:2.02-106.el8               baseos    2.0 M
 grub2-tools-extra             x86_64 1:2.02-106.el8               baseos    1.1 M
 grub2-tools-minimal           x86_64 1:2.02-106.el8               baseos    210 k
 grubby                        x86_64 8.40-42.el8                  baseos     49 k
 gsettings-desktop-schemas     x86_64 3.32.0-6.el8                 baseos    633 k
 hdparm                        x86_64 9.54-4.el8                   baseos    100 k
 hwdata                        noarch 0.314-8.10.el8               baseos    1.7 M
 iproute                       x86_64 5.12.0-4.el8                 baseos    775 k
 iptables                      x86_64 1.8.4-20.el8                 baseos    585 k
 iptables-ebtables             x86_64 1.8.4-20.el8                 baseos     72 k
 iptables-libs                 x86_64 1.8.4-20.el8                 baseos    107 k
 iscsi-initiator-utils         x86_64 6.2.1.4-4.git095f59c.el8     baseos    378 k
 iscsi-initiator-utils-iscsiuio
                               x86_64 6.2.1.4-4.git095f59c.el8     baseos    100 k
 iwl100-firmware               noarch 39.31.5.1-103.el8.1          baseos    172 k
 iwl1000-firmware              noarch 1:39.31.5.1-103.el8.1        baseos    235 k
 iwl105-firmware               noarch 18.168.6.1-103.el8.1         baseos    256 k
 iwl135-firmware               noarch 18.168.6.1-103.el8.1         baseos    265 k
 iwl2000-firmware              noarch 18.168.6.1-103.el8.1         baseos    259 k
 iwl2030-firmware              noarch 18.168.6.1-103.el8.1         baseos    268 k
 iwl3160-firmware              noarch 1:25.30.13.0-103.el8.1       baseos    1.7 M
 iwl5000-firmware              noarch 8.83.5.1_1-103.el8.1         baseos    316 k
 iwl5150-firmware              noarch 8.24.2.2-103.el8.1           baseos    169 k
 iwl6000-firmware              noarch 9.221.4.1-103.el8.1          baseos    189 k
 iwl6000g2a-firmware           noarch 18.168.6.1-103.el8.1         baseos    332 k
 iwl6000g2b-firmware           noarch 18.168.6.1-103.el8.1         baseos    332 k
 iwl6050-firmware              noarch 41.28.5.1-103.el8.1          baseos    265 k
 iwl7260-firmware              noarch 1:25.30.13.0-103.el8.1       baseos     18 M
 json-c                        x86_64 0.13.1-2.el8                 baseos     40 k
 kernel-tools                  x86_64 4.18.0-348.7.1.el8_5         baseos    7.2 M
 kernel-tools-libs             x86_64 4.18.0-348.7.1.el8_5         baseos    7.0 M
 kexec-tools                   x86_64 2.0.20-57.el8_5.1            baseos    514 k
 keyutils-libs                 x86_64 1.5.10-9.el8                 baseos     34 k
 kmod                          x86_64 25-18.el8                    baseos    126 k
 kmod-kvdo                     x86_64 6.2.5.72-81.el8              baseos    341 k
 kmod-libs                     x86_64 25-18.el8                    baseos     68 k
 kpartx                        x86_64 0.8.4-17.el8                 baseos    113 k
 kpatch                        noarch 0.9.2-5.el8                  baseos     17 k
 kpatch-dnf                    noarch 0.2-5.el8                    baseos     17 k
 krb5-libs                     x86_64 1.18.2-14.el8                baseos    840 k
 libX11                        x86_64 1.6.8-5.el8                  appstream 611 k
 libX11-common                 noarch 1.6.8-5.el8                  appstream 158 k
 libblkid                      x86_64 2.32.1-28.el8                baseos    217 k
 libblockdev                   x86_64 2.24-7.el8                   appstream 131 k
 libblockdev-crypto            x86_64 2.24-7.el8                   appstream  81 k
 libblockdev-fs                x86_64 2.24-7.el8                   appstream  86 k
 libblockdev-loop              x86_64 2.24-7.el8                   appstream  70 k
 libblockdev-lvm               x86_64 2.24-7.el8                   appstream  86 k
 libblockdev-mdraid            x86_64 2.24-7.el8                   appstream  76 k
 libblockdev-part              x86_64 2.24-7.el8                   appstream  80 k
 libblockdev-swap              x86_64 2.24-7.el8                   appstream  72 k
 libblockdev-utils             x86_64 2.24-7.el8                   appstream  80 k
 libcap                        x86_64 2.26-5.el8                   baseos     60 k
 libcap-ng                     x86_64 0.7.11-1.el8                 baseos     33 k
 libcom_err                    x86_64 1.45.6-2.el8                 baseos     49 k
 libcomps                      x86_64 0.1.16-2.el8                 baseos     82 k
 libcurl                       x86_64 7.61.1-22.el8                baseos    301 k
 libdb                         x86_64 5.3.28-42.el8_4              baseos    751 k
 libdb-utils                   x86_64 5.3.28-42.el8_4              baseos    150 k
 libdnf                        x86_64 0.63.0-3.el8                 baseos    700 k
 libdrm                        x86_64 2.4.106-2.el8                appstream 167 k
 libertas-usb8388-firmware     noarch 2:20210702-103.gitd79c2677.el8
                                                                   baseos    135 k
 libfastjson                   x86_64 0.99.9-1.el8                 appstream  38 k
 libfdisk                      x86_64 2.32.1-28.el8                baseos    251 k
 libgcc                        x86_64 8.5.0-4.el8_5                baseos     79 k
 libgcrypt                     x86_64 1.8.5-6.el8                  baseos    463 k
 libgomp                       x86_64 8.5.0-4.el8_5                baseos    206 k
 libibverbs                    x86_64 35.0-1.el8                   baseos    335 k
 libipa_hbac                   x86_64 2.5.2-2.el8_5.3              baseos    116 k
 libldb                        x86_64 2.3.0-2.el8                  baseos    189 k
 libmodulemd                   x86_64 2.13.0-1.el8                 baseos    233 k
 libmount                      x86_64 2.32.1-28.el8                baseos    234 k
 libndp                        x86_64 1.7-6.el8                    baseos     40 k
 libnfsidmap                   x86_64 1:2.3.3-46.el8               baseos    121 k
 librelp                       x86_64 1.9.0-1.el8                  appstream  76 k
 librepo                       x86_64 1.14.0-2.el8                 baseos     93 k
 libsepol                      x86_64 2.9-3.el8                    baseos    340 k
 libslirp                      x86_64 4.4.0-1.module_el8.5.0+890+6b136101
                                                                   appstream  70 k
 libsmartcols                  x86_64 2.32.1-28.el8                baseos    177 k
 libsmbclient                  x86_64 4.14.5-7.el8_5               baseos    148 k
 libsolv                       x86_64 0.7.19-1.el8                 baseos    374 k
 libss                         x86_64 1.45.6-2.el8                 baseos     54 k
 libssh                        x86_64 0.9.4-3.el8                  baseos    215 k
 libssh-config                 noarch 0.9.4-3.el8                  baseos     19 k
 libsss_autofs                 x86_64 2.5.2-2.el8_5.3              baseos    118 k
 libsss_certmap                x86_64 2.5.2-2.el8_5.3              baseos    155 k
 libsss_idmap                  x86_64 2.5.2-2.el8_5.3              baseos    120 k
 libsss_nss_idmap              x86_64 2.5.2-2.el8_5.3              baseos    127 k
 libsss_sudo                   x86_64 2.5.2-2.el8_5.3              baseos    116 k
 libstdc++                     x86_64 8.5.0-4.el8_5                baseos    453 k
 libstoragemgmt                x86_64 1.9.1-1.el8                  baseos    246 k
 libtalloc                     x86_64 2.3.2-1.el8                  baseos     49 k
 libtevent                     x86_64 0.11.0-0.el8                 baseos     50 k
 libtirpc                      x86_64 1.1.4-5.el8                  baseos    112 k
 libudisks2                    x86_64 2.9.0-7.el8                  appstream 184 k
 libuuid                       x86_64 2.32.1-28.el8                baseos     96 k
 libwbclient                   x86_64 4.14.5-7.el8_5               baseos    121 k
 libxcrypt                     x86_64 4.1.1-6.el8                  baseos     73 k
 libxml2                       x86_64 2.9.7-9.el8_4.2              baseos    696 k
 linux-firmware                noarch 20210702-103.gitd79c2677.el8 baseos    161 M
 lshw                          x86_64 B.02.19.2-6.el8              baseos    341 k
 lsscsi                        x86_64 0.32-3.el8                   baseos     71 k
 lua-libs                      x86_64 5.3.4-12.el8                 baseos    118 k
 lvm2                          x86_64 8:2.03.12-10.el8             baseos    1.6 M
 lvm2-libs                     x86_64 8:2.03.12-10.el8             baseos    1.2 M
 lz4-libs                      x86_64 1.8.3-3.el8_4                baseos     66 k
 man-db                        x86_64 2.7.6.1-18.el8               baseos    887 k
 man-pages-overrides           noarch 8.5.0.1-1.el8                appstream  98 k
 mcelog                        x86_64 3:175-1.el8                  baseos     83 k
 mdadm                         x86_64 4.2-rc2.el8                  baseos    460 k
 microcode_ctl                 x86_64 4:20210608-1.el8             baseos    5.5 M
 ncurses                       x86_64 6.1-9.20180224.el8           baseos    387 k
 ncurses-base                  noarch 6.1-9.20180224.el8           baseos     81 k
 ncurses-libs                  x86_64 6.1-9.20180224.el8           baseos    334 k
 nettle                        x86_64 3.4.1-7.el8                  baseos    301 k
 nftables                      x86_64 1:0.9.3-21.el8               baseos    321 k
 nmap-ncat                     x86_64 2:7.70-6.el8                 appstream 237 k
 nspr                          x86_64 4.32.0-1.el8_4               appstream 142 k
 nss                           x86_64 3.67.0-7.el8_5               appstream 741 k
 nss-softokn                   x86_64 3.67.0-7.el8_5               appstream 487 k
 nss-softokn-freebl            x86_64 3.67.0-7.el8_5               appstream 395 k
 nss-sysinit                   x86_64 3.67.0-7.el8_5               appstream  73 k
 nss-util                      x86_64 3.67.0-7.el8_5               appstream 137 k
 numactl-libs                  x86_64 2.0.12-13.el8                baseos     36 k
 nvme-cli                      x86_64 1.14-3.el8                   baseos    463 k
 open-vm-tools                 x86_64 11.2.5-2.el8                 appstream 771 k
 openldap                      x86_64 2.4.46-18.el8                baseos    352 k
 openssh                       x86_64 8.0p1-10.el8                 baseos    522 k
 openssh-clients               x86_64 8.0p1-10.el8                 baseos    668 k
 openssh-server                x86_64 8.0p1-10.el8                 baseos    485 k
 openssl                       x86_64 1:1.1.1k-5.el8_5             baseos    709 k
 openssl-libs                  x86_64 1:1.1.1k-5.el8_5             baseos    1.5 M
 os-prober                     x86_64 1.74-9.el8                   baseos     51 k
 pam                           x86_64 1.3.1-15.el8                 baseos    739 k
 parted                        x86_64 3.2-39.el8                   baseos    555 k
 pcre                          x86_64 8.42-6.el8                   baseos    211 k
 platform-python               x86_64 3.6.8-41.el8                 baseos     85 k
 platform-python-pip           noarch 9.0.3-20.el8                 baseos    1.7 M
 plymouth                      x86_64 0.9.4-10.20200615git1e36e30.el8
                                                                   appstream 127 k
 plymouth-core-libs            x86_64 0.9.4-10.20200615git1e36e30.el8
                                                                   appstream 122 k
 plymouth-scripts              x86_64 0.9.4-10.20200615git1e36e30.el8
                                                                   appstream  44 k
 podman                        x86_64 3.3.1-9.module_el8.5.0+988+b1f0b741
                                                                   appstream  12 M
 podman-catatonit              x86_64 3.3.1-9.module_el8.5.0+988+b1f0b741
                                                                   appstream 340 k
 policycoreutils               x86_64 2.9-16.el8                   baseos    373 k
 policycoreutils-python-utils  noarch 2.9-16.el8                   baseos    252 k
 polkit                        x86_64 0.115-12.el8                 baseos    154 k
 polkit-libs                   x86_64 0.115-12.el8                 baseos     76 k
 python3-bind                  noarch 32:9.11.26-6.el8             appstream 150 k
 python3-dnf                   noarch 4.7.0-4.el8                  baseos    545 k
 python3-dnf-plugins-core      noarch 4.0.21-3.el8                 baseos    234 k
 python3-firewall              noarch 0.9.3-7.el8                  baseos    432 k
 python3-gpg                   x86_64 1.13.1-9.el8                 baseos    245 k
 python3-hawkey                x86_64 0.63.0-3.el8                 baseos    116 k
 python3-libcomps              x86_64 0.1.16-2.el8                 baseos     51 k
 python3-libdnf                x86_64 0.63.0-3.el8                 baseos    777 k
 python3-libs                  x86_64 3.6.8-41.el8                 baseos    7.8 M
 python3-libstoragemgmt        x86_64 1.9.1-1.el8                  baseos    175 k
     replacing  python3-libstoragemgmt-clibs.x86_64 1.8.7-1.el8
 python3-libxml2               x86_64 2.9.7-9.el8_4.2              baseos    237 k
 python3-lxml                  x86_64 4.2.3-3.el8                  appstream 1.5 M
 python3-nftables              x86_64 1:0.9.3-21.el8               baseos     29 k
 python3-perf                  x86_64 4.18.0-348.7.1.el8_5         baseos    7.1 M
 python3-pip-wheel             noarch 9.0.3-20.el8                 baseos    1.0 M
 python3-policycoreutils       noarch 2.9-16.el8                   baseos    2.2 M
 python3-psutil                x86_64 5.4.3-11.el8                 appstream 373 k
 python3-rpm                   x86_64 4.14.3-19.el8                baseos    154 k
 python3-sssdconfig            noarch 2.5.2-2.el8_5.3              baseos    142 k
 python3-syspurpose            x86_64 1.28.21-3.el8                baseos    312 k
 python3-unbound               x86_64 1.7.3-17.el8                 appstream 119 k
 quota                         x86_64 1:4.04-14.el8                baseos    214 k
 quota-nls                     noarch 1:4.04-14.el8                baseos     95 k
 rdma-core                     x86_64 35.0-1.el8                   baseos     59 k
 realmd                        x86_64 0.16.3-23.el8                baseos    238 k
 rpm                           x86_64 4.14.3-19.el8                baseos    543 k
 rpm-build-libs                x86_64 4.14.3-19.el8                baseos    156 k
 rpm-libs                      x86_64 4.14.3-19.el8                baseos    344 k
 rpm-plugin-selinux            x86_64 4.14.3-19.el8                baseos     77 k
 rpm-plugin-systemd-inhibit    x86_64 4.14.3-19.el8                baseos     78 k
 rsyslog                       x86_64 8.2102.0-5.el8               appstream 752 k
 rsyslog-gnutls                x86_64 8.2102.0-5.el8               appstream  32 k
 rsyslog-gssapi                x86_64 8.2102.0-5.el8               appstream  34 k
 rsyslog-relp                  x86_64 8.2102.0-5.el8               appstream  33 k
 runc                          x86_64 1.0.2-1.module_el8.5.0+911+f19012f9
                                                                   appstream 3.1 M
 samba-client-libs             x86_64 4.14.5-7.el8_5               baseos    5.4 M
 samba-common                  noarch 4.14.5-7.el8_5               baseos    221 k
 samba-common-libs             x86_64 4.14.5-7.el8_5               baseos    173 k
 selinux-policy                noarch 3.14.3-80.el8_5.2            baseos    636 k
 selinux-policy-targeted       noarch 3.14.3-80.el8_5.2            baseos     15 M
 setroubleshoot-plugins        noarch 3.3.14-1.el8                 appstream 358 k
 setroubleshoot-server         x86_64 3.3.24-4.el8                 appstream 388 k
 shadow-utils                  x86_64 2:4.6-14.el8                 baseos    1.2 M
 slirp4netns                   x86_64 1.1.8-1.module_el8.5.0+890+6b136101
                                                                   appstream  51 k
 sos                           noarch 4.1-5.el8                    baseos    706 k
 sqlite                        x86_64 3.26.0-15.el8                baseos    668 k
 sqlite-libs                   x86_64 3.26.0-15.el8                baseos    581 k
 sssd                          x86_64 2.5.2-2.el8_5.3              baseos    107 k
 sssd-ad                       x86_64 2.5.2-2.el8_5.3              baseos    271 k
 sssd-client                   x86_64 2.5.2-2.el8_5.3              baseos    205 k
 sssd-common                   x86_64 2.5.2-2.el8_5.3              baseos    1.6 M
 sssd-common-pac               x86_64 2.5.2-2.el8_5.3              baseos    179 k
 sssd-ipa                      x86_64 2.5.2-2.el8_5.3              baseos    348 k
 sssd-kcm                      x86_64 2.5.2-2.el8_5.3              baseos    254 k
 sssd-krb5                     x86_64 2.5.2-2.el8_5.3              baseos    150 k
 sssd-krb5-common              x86_64 2.5.2-2.el8_5.3              baseos    186 k
 sssd-ldap                     x86_64 2.5.2-2.el8_5.3              baseos    209 k
 sssd-nfs-idmap                x86_64 2.5.2-2.el8_5.3              baseos    116 k
 sssd-proxy                    x86_64 2.5.2-2.el8_5.3              baseos    148 k
 strace                        x86_64 5.7-3.el8                    baseos    1.1 M
 sudo                          x86_64 1.8.29-7.el8_4.1             baseos    925 k
 systemd                       x86_64 239-51.el8_5.2               baseos    3.6 M
 systemd-libs                  x86_64 239-51.el8_5.2               baseos    1.1 M
 systemd-pam                   x86_64 239-51.el8_5.2               baseos    477 k
 systemd-udev                  x86_64 239-51.el8_5.2               baseos    1.6 M
 tcpdump                       x86_64 14:4.9.3-2.el8               appstream 452 k
 tpm2-tools                    x86_64 4.1.1-5.el8                  baseos    1.0 M
 tpm2-tss                      x86_64 2.3.2-4.el8                  baseos    275 k
 tuned                         noarch 2.16.0-1.el8                 baseos    312 k
 tzdata                        noarch 2021e-1.el8                  baseos    474 k
 udisks2                       x86_64 2.9.0-7.el8                  appstream 474 k
 udisks2-iscsi                 x86_64 2.9.0-7.el8                  appstream  30 k
 udisks2-lvm2                  x86_64 2.9.0-7.el8                  appstream  45 k
 unbound-libs                  x86_64 1.7.3-17.el8                 appstream 503 k
 unzip                         x86_64 6.0-45.el8_4                 baseos    195 k
 util-linux                    x86_64 2.32.1-28.el8                baseos    2.5 M
 util-linux-user               x86_64 2.32.1-28.el8                baseos    100 k
 vdo                           x86_64 6.2.5.74-14.el8              baseos    663 k
 vim-common                    x86_64 2:8.0.1763-16.el8            appstream 6.3 M
 vim-enhanced                  x86_64 2:8.0.1763-16.el8            appstream 1.4 M
 vim-filesystem                noarch 2:8.0.1763-16.el8            appstream  49 k
 vim-minimal                   x86_64 2:8.0.1763-16.el8            baseos    573 k
 virt-what                     x86_64 1.18-12.el8                  baseos     36 k
 which                         x86_64 2.21-16.el8                  baseos     49 k
 xfsprogs                      x86_64 5.0.0-9.el8                  baseos    1.1 M
 yum                           noarch 4.7.0-4.el8                  baseos    205 k
Installing dependencies:
 grub2-tools-efi               x86_64 1:2.02-106.el8               baseos    474 k
 libbpf                        x86_64 0.4.0-1.el8                  baseos    110 k

Transaction Summary
===================================================================================
Install    5 Packages
Upgrade  324 Packages

Total download size: 500 M
Is this ok [y/N]: y
Downloading Packages:
(1/329): grub2-tools-efi-2.02-106.el8.x86_64.rpm    96 kB/s | 474 kB     00:04
(2/329): kernel-4.18.0-348.7.1.el8_5.x86_64.rpm    392 kB/s | 7.0 MB     00:18
(3/329): libbpf-0.4.0-1.el8.x86_64.rpm             122 kB/s | 110 kB     00:00
(4/329): authselect-compat-1.2.2-3.el8.x86_64.rpm   60 kB/s |  38 kB     00:00
(5/329): bind-libs-9.11.26-6.el8.x86_64.rpm        144 kB/s | 174 kB     00:01
(6/329): bind-libs-lite-9.11.26-6.el8.x86_64.rpm   336 kB/s | 1.2 MB     00:03
(7/329): bind-license-9.11.26-6.el8.noarch.rpm     115 kB/s | 102 kB     00:00
(8/329): bind-utils-9.11.26-6.el8.x86_64.rpm       239 kB/s | 451 kB     00:01
(9/329): kernel-core-4.18.0-348.7.1.el8_5.x86_64.r 1.2 MB/s |  38 MB     00:31
(10/329): kernel-modules-4.18.0-348.7.1.el8_5.x86_ 1.1 MB/s |  30 MB     00:26
(11/329): cockpit-packagekit-251.1-1.el8.noarch.rp 654 kB/s | 703 kB     00:01
(12/329): cockpit-podman-33-1.module_el8.5.0+890+6 348 kB/s | 437 kB     00:01
(13/329): conmon-2.0.29-1.module_el8.5.0+890+6b136  80 kB/s |  52 kB     00:00
(14/329): cockpit-storaged-251.1-1.el8.noarch.rpm  624 kB/s | 655 kB     00:01
(15/329): container-selinux-2.167.0-1.module_el8.5  88 kB/s |  54 kB     00:00
(16/329): containers-common-1-2.module_el8.5.0+890 126 kB/s |  79 kB     00:00
(17/329): criu-3.15-3.module_el8.5.0+890+6b136101. 407 kB/s | 518 kB     00:01
(18/329): fstrm-0.6.1-2.el8.x86_64.rpm              49 kB/s |  29 kB     00:00
(19/329): fuse-overlayfs-1.7.1-1.module_el8.5.0+89 120 kB/s |  73 kB     00:00
(20/329): buildah-1.22.3-2.module_el8.5.0+911+f190 721 kB/s | 7.7 MB     00:10
(21/329): libX11-common-1.6.8-5.el8.noarch.rpm     259 kB/s | 158 kB     00:00
(22/329): libX11-1.6.8-5.el8.x86_64.rpm            390 kB/s | 611 kB     00:01
(23/329): libblockdev-2.24-7.el8.x86_64.rpm        217 kB/s | 131 kB     00:00
(24/329): libblockdev-crypto-2.24-7.el8.x86_64.rpm 135 kB/s |  81 kB     00:00
(25/329): libblockdev-fs-2.24-7.el8.x86_64.rpm     144 kB/s |  86 kB     00:00
(26/329): libblockdev-loop-2.24-7.el8.x86_64.rpm   117 kB/s |  70 kB     00:00
(27/329): libblockdev-lvm-2.24-7.el8.x86_64.rpm    142 kB/s |  86 kB     00:00
(28/329): libblockdev-mdraid-2.24-7.el8.x86_64.rpm  85 kB/s |  76 kB     00:00
(29/329): libblockdev-part-2.24-7.el8.x86_64.rpm   134 kB/s |  80 kB     00:00
(30/329): libblockdev-swap-2.24-7.el8.x86_64.rpm   122 kB/s |  72 kB     00:00
(31/329): libblockdev-utils-2.24-7.el8.x86_64.rpm   89 kB/s |  80 kB     00:00
(32/329): libdrm-2.4.106-2.el8.x86_64.rpm          186 kB/s | 167 kB     00:00
(33/329): libfastjson-0.99.9-1.el8.x86_64.rpm       63 kB/s |  38 kB     00:00
(34/329): librelp-1.9.0-1.el8.x86_64.rpm           127 kB/s |  76 kB     00:00
(35/329): libslirp-4.4.0-1.module_el8.5.0+890+6b13  78 kB/s |  70 kB     00:00
(36/329): libudisks2-2.9.0-7.el8.x86_64.rpm        200 kB/s | 184 kB     00:00
(37/329): man-pages-overrides-8.5.0.1-1.el8.noarch 109 kB/s |  98 kB     00:00
(38/329): nmap-ncat-7.70-6.el8.x86_64.rpm          265 kB/s | 237 kB     00:00
(39/329): nspr-4.32.0-1.el8_4.x86_64.rpm           138 kB/s | 142 kB     00:01
(40/329): nss-3.67.0-7.el8_5.x86_64.rpm            531 kB/s | 741 kB     00:01
(41/329): nss-softokn-3.67.0-7.el8_5.x86_64.rpm    384 kB/s | 487 kB     00:01
(42/329): nss-softokn-freebl-3.67.0-7.el8_5.x86_64 430 kB/s | 395 kB     00:00
(43/329): nss-sysinit-3.67.0-7.el8_5.x86_64.rpm    121 kB/s |  73 kB     00:00
(44/329): nss-util-3.67.0-7.el8_5.x86_64.rpm       227 kB/s | 137 kB     00:00
(45/329): plymouth-0.9.4-10.20200615git1e36e30.el8 209 kB/s | 127 kB     00:00
(46/329): containernetworking-plugins-1.0.0-1.modu 1.2 MB/s |  19 MB     00:15
(47/329): plymouth-core-libs-0.9.4-10.20200615git1 177 kB/s | 122 kB     00:00
(48/329): open-vm-tools-11.2.5-2.el8.x86_64.rpm    421 kB/s | 771 kB     00:01
(49/329): plymouth-scripts-0.9.4-10.20200615git1e3  74 kB/s |  44 kB     00:00
(50/329): podman-catatonit-3.3.1-9.module_el8.5.0+ 272 kB/s | 340 kB     00:01
(51/329): python3-bind-9.11.26-6.el8.noarch.rpm    162 kB/s | 150 kB     00:00
(52/329): python3-psutil-5.4.3-11.el8.x86_64.rpm   304 kB/s | 373 kB     00:01
(53/329): python3-lxml-4.2.3-3.el8.x86_64.rpm      903 kB/s | 1.5 MB     00:01
(54/329): python3-unbound-1.7.3-17.el8.x86_64.rpm  192 kB/s | 119 kB     00:00
(55/329): rsyslog-gnutls-8.2102.0-5.el8.x86_64.rpm  53 kB/s |  32 kB     00:00
(56/329): rsyslog-8.2102.0-5.el8.x86_64.rpm        625 kB/s | 752 kB     00:01
(57/329): rsyslog-gssapi-8.2102.0-5.el8.x86_64.rpm  56 kB/s |  34 kB     00:00
(58/329): rsyslog-relp-8.2102.0-5.el8.x86_64.rpm    56 kB/s |  33 kB     00:00
(59/329): setroubleshoot-plugins-3.3.14-1.el8.noar 306 kB/s | 358 kB     00:01
(60/329): setroubleshoot-server-3.3.24-4.el8.x86_6 424 kB/s | 388 kB     00:00
(61/329): slirp4netns-1.1.8-1.module_el8.5.0+890+6  84 kB/s |  51 kB     00:00
(62/329): tcpdump-4.9.3-2.el8.x86_64.rpm           375 kB/s | 452 kB     00:01
(63/329): runc-1.0.2-1.module_el8.5.0+911+f19012f9 667 kB/s | 3.1 MB     00:04
(64/329): podman-3.3.1-9.module_el8.5.0+988+b1f0b7 1.3 MB/s |  12 MB     00:09
(65/329): udisks2-2.9.0-7.el8.x86_64.rpm           492 kB/s | 474 kB     00:00
(66/329): udisks2-iscsi-2.9.0-7.el8.x86_64.rpm      52 kB/s |  30 kB     00:00
(67/329): udisks2-lvm2-2.9.0-7.el8.x86_64.rpm       77 kB/s |  45 kB     00:00
(68/329): unbound-libs-1.7.3-17.el8.x86_64.rpm     540 kB/s | 503 kB     00:00
(69/329): vim-filesystem-8.0.1763-16.el8.noarch.rp  82 kB/s |  49 kB     00:00
(70/329): vim-enhanced-8.0.1763-16.el8.x86_64.rpm  485 kB/s | 1.4 MB     00:02
(71/329): NetworkManager-config-server-1.32.10-4.e 214 kB/s | 131 kB     00:00
(72/329): NetworkManager-1.32.10-4.el8.x86_64.rpm  1.0 MB/s | 2.6 MB     00:02
(73/329): NetworkManager-team-1.32.10-4.el8.x86_64 239 kB/s | 148 kB     00:00
(74/329): vim-common-8.0.1763-16.el8.x86_64.rpm    1.2 MB/s | 6.3 MB     00:05
(75/329): NetworkManager-tui-1.32.10-4.el8.x86_64. 350 kB/s | 336 kB     00:00
(76/329): adcli-0.8.2-12.el8.x86_64.rpm            193 kB/s | 118 kB     00:00
(77/329): NetworkManager-libnm-1.32.10-4.el8.x86_6 638 kB/s | 1.8 MB     00:02
(78/329): alsa-sof-firmware-1.8-1.el8.noarch.rpm   653 kB/s | 676 kB     00:01
(79/329): authselect-1.2.2-3.el8.x86_64.rpm        116 kB/s | 133 kB     00:01
(80/329): authselect-libs-1.2.2-3.el8.x86_64.rpm   363 kB/s | 222 kB     00:00
(81/329): bind-export-libs-9.11.26-6.el8.x86_64.rp 704 kB/s | 1.1 MB     00:01
(82/329): bash-4.4.20-2.el8.x86_64.rpm             764 kB/s | 1.5 MB     00:02
(83/329): ca-certificates-2021.2.50-80.0.el8_4.noa 209 kB/s | 390 kB     00:01
(84/329): centos-gpg-keys-8-3.el8.noarch.rpm        21 kB/s |  12 kB     00:00
(85/329): centos-linux-release-8.5-1.2111.el8.noar  36 kB/s |  22 kB     00:00
(86/329): centos-linux-repos-8-3.el8.noarch.rpm     34 kB/s |  20 kB     00:00
(87/329): centos-logos-85.8-2.el8.x86_64.rpm       357 kB/s | 700 kB     00:01
(88/329): bpftool-4.18.0-348.7.1.el8_5.x86_64.rpm  1.2 MB/s | 7.7 MB     00:06
(89/329): chkconfig-1.19.1-1.el8.x86_64.rpm        323 kB/s | 198 kB     00:00
(90/329): chrony-4.1-1.el8.x86_64.rpm              367 kB/s | 327 kB     00:00
(91/329): cockpit-251.1-1.el8.x86_64.rpm            86 kB/s |  78 kB     00:00
(92/329): cockpit-bridge-251.1-1.el8.x86_64.rpm    579 kB/s | 539 kB     00:00
(93/329): binutils-2.30-108.el8_5.1.x86_64.rpm     569 kB/s | 5.8 MB     00:10
(94/329): cockpit-ws-251.1-1.el8.x86_64.rpm        889 kB/s | 1.3 MB     00:01
(95/329): cockpit-system-251.1-1.el8.noarch.rpm    1.1 MB/s | 3.2 MB     00:02
(96/329): crypto-policies-20210617-1.gitc776d3e.el 107 kB/s |  63 kB     00:00
(97/329): crypto-policies-scripts-20210617-1.gitc7 130 kB/s |  83 kB     00:00
(98/329): coreutils-8.30-12.el8.x86_64.rpm         498 kB/s | 1.2 MB     00:02
(99/329): coreutils-common-8.30-12.el8.x86_64.rpm  1.0 MB/s | 2.0 MB     00:01
(100/329): cups-libs-2.2.6-40.el8.x86_64.rpm       523 kB/s | 433 kB     00:00
(101/329): dbus-1.12.8-14.el8.x86_64.rpm            46 kB/s |  41 kB     00:00
(102/329): curl-7.61.1-22.el8.x86_64.rpm           285 kB/s | 351 kB     00:01
(103/329): dbus-common-1.12.8-14.el8.noarch.rpm     78 kB/s |  46 kB     00:00
(104/329): dbus-daemon-1.12.8-14.el8.x86_64.rpm    380 kB/s | 240 kB     00:00
(105/329): dbus-libs-1.12.8-14.el8.x86_64.rpm      294 kB/s | 184 kB     00:00
(106/329): dbus-tools-1.12.8-14.el8.x86_64.rpm     139 kB/s |  85 kB     00:00
(107/329): device-mapper-event-1.02.177-10.el8.x86 415 kB/s | 271 kB     00:00
(108/329): device-mapper-1.02.177-10.el8.x86_64.rp 452 kB/s | 377 kB     00:00
(109/329): device-mapper-multipath-0.8.4-17.el8.x8 309 kB/s | 197 kB     00:00
(110/329): device-mapper-libs-1.02.177-10.el8.x86_ 446 kB/s | 409 kB     00:00
(111/329): device-mapper-event-libs-1.02.177-10.el 147 kB/s | 270 kB     00:01
(112/329): device-mapper-multipath-libs-0.8.4-17.e 401 kB/s | 322 kB     00:00
(113/329): dhcp-client-4.3.6-45.el8.x86_64.rpm     390 kB/s | 318 kB     00:00
(114/329): device-mapper-persistent-data-0.9.0-4.e 732 kB/s | 925 kB     00:01
(115/329): dhcp-common-4.3.6-45.el8.noarch.rpm     223 kB/s | 207 kB     00:00
(116/329): dhcp-libs-4.3.6-45.el8.x86_64.rpm       241 kB/s | 148 kB     00:00
(117/329): dmidecode-3.2-10.el8.x86_64.rpm         150 kB/s |  91 kB     00:00
(118/329): dnf-plugins-core-4.0.21-3.el8.noarch.rp 116 kB/s |  70 kB     00:00
(119/329): dnf-data-4.7.0-4.el8.noarch.rpm         173 kB/s | 154 kB     00:00
(120/329): dnf-4.7.0-4.el8.noarch.rpm              354 kB/s | 544 kB     00:01
(121/329): dracut-config-rescue-049-191.git2021092 101 kB/s |  61 kB     00:00
(122/329): dracut-049-191.git20210920.el8.x86_64.r 415 kB/s | 374 kB     00:00
(123/329): dracut-network-049-191.git20210920.el8. 117 kB/s | 108 kB     00:00
(124/329): dracut-squash-049-191.git20210920.el8.x  87 kB/s |  61 kB     00:00
(125/329): e2fsprogs-libs-1.45.6-2.el8.x86_64.rpm  379 kB/s | 233 kB     00:00
(126/329): elfutils-debuginfod-client-0.185-1.el8. 108 kB/s |  66 kB     00:00
(127/329): elfutils-default-yama-scope-0.185-1.el8  83 kB/s |  49 kB     00:00
(128/329): e2fsprogs-1.45.6-2.el8.x86_64.rpm       565 kB/s | 1.0 MB     00:01
(129/329): emacs-filesystem-26.1-7.el8.noarch.rpm  115 kB/s |  70 kB     00:00
(130/329): elfutils-libelf-0.185-1.el8.x86_64.rpm  181 kB/s | 221 kB     00:01
(131/329): elfutils-libs-0.185-1.el8.x86_64.rpm    316 kB/s | 292 kB     00:00
(132/329): ethtool-5.8-7.el8.x86_64.rpm            335 kB/s | 209 kB     00:00
(133/329): file-5.33-20.el8.x86_64.rpm             126 kB/s |  77 kB     00:00
(134/329): file-libs-5.33-20.el8.x86_64.rpm        351 kB/s | 543 kB     00:01
(135/329): filesystem-3.8-6.el8.x86_64.rpm         836 kB/s | 1.1 MB     00:01
(136/329): firewalld-0.9.3-7.el8.noarch.rpm        378 kB/s | 502 kB     00:01
(137/329): firewalld-filesystem-0.9.3-7.el8.noarch 129 kB/s |  77 kB     00:00
(138/329): fontconfig-2.13.1-4.el8.x86_64.rpm      297 kB/s | 274 kB     00:00
(139/329): glibc-common-2.28-164.el8.x86_64.rpm    853 kB/s | 1.3 MB     00:01
(140/329): glibc-2.28-164.el8.x86_64.rpm           1.1 MB/s | 3.6 MB     00:03
(141/329): glib2-2.56.4-156.el8.x86_64.rpm         671 kB/s | 2.5 MB     00:03
(142/329): glibc-langpack-en-2.28-164.el8.x86_64.r 537 kB/s | 828 kB     00:01
(143/329): gpgme-1.13.1-9.el8.x86_64.rpm           403 kB/s | 336 kB     00:00
(144/329): gnutls-3.6.16-4.el8.x86_64.rpm          787 kB/s | 1.0 MB     00:01
(145/329): grub2-pc-2.02-106.el8.x86_64.rpm         69 kB/s |  42 kB     00:00
(146/329): grub2-common-2.02-106.el8.noarch.rpm    566 kB/s | 891 kB     00:01
(147/329): grub2-pc-modules-2.02-106.el8.noarch.rp 723 kB/s | 916 kB     00:01
(148/329): grub2-tools-minimal-2.02-106.el8.x86_64 320 kB/s | 210 kB     00:00
(149/329): grub2-tools-extra-2.02-106.el8.x86_64.r 694 kB/s | 1.1 MB     00:01
(150/329): grubby-8.40-42.el8.x86_64.rpm            82 kB/s |  49 kB     00:00
(151/329): grub2-tools-2.02-106.el8.x86_64.rpm     813 kB/s | 2.0 MB     00:02
(152/329): hdparm-9.54-4.el8.x86_64.rpm            163 kB/s | 100 kB     00:00
(153/329): gsettings-desktop-schemas-3.32.0-6.el8. 508 kB/s | 633 kB     00:01
(154/329): iproute-5.12.0-4.el8.x86_64.rpm         617 kB/s | 775 kB     00:01
(155/329): iptables-1.8.4-20.el8.x86_64.rpm        405 kB/s | 585 kB     00:01
(156/329): iptables-ebtables-1.8.4-20.el8.x86_64.r 107 kB/s |  72 kB     00:00
(157/329): hwdata-0.314-8.10.el8.noarch.rpm        740 kB/s | 1.7 MB     00:02
(158/329): iptables-libs-1.8.4-20.el8.x86_64.rpm   176 kB/s | 107 kB     00:00
(159/329): iscsi-initiator-utils-iscsiuio-6.2.1.4- 111 kB/s | 100 kB     00:00
(160/329): iscsi-initiator-utils-6.2.1.4-4.git095f 314 kB/s | 378 kB     00:01
(161/329): iwl1000-firmware-39.31.5.1-103.el8.1.no 370 kB/s | 235 kB     00:00
(162/329): iwl100-firmware-39.31.5.1-103.el8.1.noa 143 kB/s | 172 kB     00:01
(163/329): iwl105-firmware-18.168.6.1-103.el8.1.no 279 kB/s | 256 kB     00:00
(164/329): iwl135-firmware-18.168.6.1-103.el8.1.no 294 kB/s | 265 kB     00:00
(165/329): iwl2000-firmware-18.168.6.1-103.el8.1.n 281 kB/s | 259 kB     00:00
(166/329): iwl2030-firmware-18.168.6.1-103.el8.1.n 295 kB/s | 268 kB     00:00
(167/329): iwl5000-firmware-8.83.5.1_1-103.el8.1.n 335 kB/s | 316 kB     00:00
(168/329): iwl5150-firmware-8.24.2.2-103.el8.1.noa 272 kB/s | 169 kB     00:00
(169/329): iwl6000-firmware-9.221.4.1-103.el8.1.no 303 kB/s | 189 kB     00:00
(170/329): iwl3160-firmware-25.30.13.0-103.el8.1.n 985 kB/s | 1.7 MB     00:01
(171/329): iwl6000g2a-firmware-18.168.6.1-103.el8. 276 kB/s | 332 kB     00:01
(172/329): iwl6000g2b-firmware-18.168.6.1-103.el8. 366 kB/s | 332 kB     00:00
(173/329): iwl6050-firmware-41.28.5.1-103.el8.1.no 293 kB/s | 265 kB     00:00
(174/329): json-c-0.13.1-2.el8.x86_64.rpm           67 kB/s |  40 kB     00:00
(175/329): kernel-tools-libs-4.18.0-348.7.1.el8_5. 1.0 MB/s | 7.0 MB     00:07
(176/329): kernel-tools-4.18.0-348.7.1.el8_5.x86_6 841 kB/s | 7.2 MB     00:08
(177/329): kexec-tools-2.0.20-57.el8_5.1.x86_64.rp 336 kB/s | 514 kB     00:01
(178/329): keyutils-libs-1.5.10-9.el8.x86_64.rpm    56 kB/s |  34 kB     00:00
(179/329): kmod-25-18.el8.x86_64.rpm               206 kB/s | 126 kB     00:00
(180/329): kmod-libs-25-18.el8.x86_64.rpm          109 kB/s |  68 kB     00:00
(181/329): kmod-kvdo-6.2.5.72-81.el8.x86_64.rpm    268 kB/s | 341 kB     00:01
(182/329): kpartx-0.8.4-17.el8.x86_64.rpm          183 kB/s | 113 kB     00:00
(183/329): kpatch-0.9.2-5.el8.noarch.rpm            28 kB/s |  17 kB     00:00
(184/329): kpatch-dnf-0.2-5.el8.noarch.rpm          29 kB/s |  17 kB     00:00
(185/329): libblkid-2.32.1-28.el8.x86_64.rpm       237 kB/s | 217 kB     00:00
(186/329): libcap-2.26-5.el8.x86_64.rpm             99 kB/s |  60 kB     00:00
(187/329): krb5-libs-1.18.2-14.el8.x86_64.rpm      392 kB/s | 840 kB     00:02
(188/329): iwl7260-firmware-25.30.13.0-103.el8.1.n 1.3 MB/s |  18 MB     00:14
(189/329): libcap-ng-0.7.11-1.el8.x86_64.rpm        55 kB/s |  33 kB     00:00
(190/329): libcom_err-1.45.6-2.el8.x86_64.rpm       61 kB/s |  49 kB     00:00
(191/329): libcomps-0.1.16-2.el8.x86_64.rpm         77 kB/s |  82 kB     00:01
(192/329): libdb-5.3.28-42.el8_4.x86_64.rpm        470 kB/s | 751 kB     00:01
(193/329): libdb-utils-5.3.28-42.el8_4.x86_64.rpm  160 kB/s | 150 kB     00:00
(194/329): libertas-usb8388-firmware-20210702-103. 127 kB/s | 135 kB     00:01
(195/329): libcurl-7.61.1-22.el8.x86_64.rpm         77 kB/s | 301 kB     00:03
(196/329): libfdisk-2.32.1-28.el8.x86_64.rpm       261 kB/s | 251 kB     00:00
(197/329): libdnf-0.63.0-3.el8.x86_64.rpm          305 kB/s | 700 kB     00:02
(198/329): libgcc-8.5.0-4.el8_5.x86_64.rpm         131 kB/s |  79 kB     00:00
(199/329): libgomp-8.5.0-4.el8_5.x86_64.rpm        331 kB/s | 206 kB     00:00
(200/329): libgcrypt-1.8.5-6.el8.x86_64.rpm        376 kB/s | 463 kB     00:01
(201/329): libipa_hbac-2.5.2-2.el8_5.3.x86_64.rpm  192 kB/s | 116 kB     00:00
(202/329): libldb-2.3.0-2.el8.x86_64.rpm           208 kB/s | 189 kB     00:00
(203/329): libibverbs-35.0-1.el8.x86_64.rpm        185 kB/s | 335 kB     00:01
(204/329): libmodulemd-2.13.0-1.el8.x86_64.rpm     256 kB/s | 233 kB     00:00
(205/329): libndp-1.7-6.el8.x86_64.rpm              67 kB/s |  40 kB     00:00
(206/329): libmount-2.32.1-28.el8.x86_64.rpm       196 kB/s | 234 kB     00:01
(207/329): librepo-1.14.0-2.el8.x86_64.rpm         153 kB/s |  93 kB     00:00
(208/329): libnfsidmap-2.3.3-46.el8.x86_64.rpm      69 kB/s | 121 kB     00:01
(209/329): libsepol-2.9-3.el8.x86_64.rpm           267 kB/s | 340 kB     00:01
(210/329): libsmartcols-2.32.1-28.el8.x86_64.rpm   139 kB/s | 177 kB     00:01
(211/329): libsmbclient-4.14.5-7.el8_5.x86_64.rpm  152 kB/s | 148 kB     00:00
(212/329): libss-1.45.6-2.el8.x86_64.rpm            85 kB/s |  54 kB     00:00
(213/329): libsolv-0.7.19-1.el8.x86_64.rpm         303 kB/s | 374 kB     00:01
(214/329): libssh-config-0.9.4-3.el8.noarch.rpm     31 kB/s |  19 kB     00:00
(215/329): libssh-0.9.4-3.el8.x86_64.rpm           234 kB/s | 215 kB     00:00
(216/329): libsss_autofs-2.5.2-2.el8_5.3.x86_64.rp 130 kB/s | 118 kB     00:00
(217/329): libsss_certmap-2.5.2-2.el8_5.3.x86_64.r 172 kB/s | 155 kB     00:00
(218/329): libsss_idmap-2.5.2-2.el8_5.3.x86_64.rpm 100 kB/s | 120 kB     00:01
(219/329): libsss_nss_idmap-2.5.2-2.el8_5.3.x86_64 140 kB/s | 127 kB     00:00
(220/329): libsss_sudo-2.5.2-2.el8_5.3.x86_64.rpm  129 kB/s | 116 kB     00:00
(221/329): libtalloc-2.3.2-1.el8.x86_64.rpm         81 kB/s |  49 kB     00:00
(222/329): libstoragemgmt-1.9.1-1.el8.x86_64.rpm   205 kB/s | 246 kB     00:01
(223/329): libstdc++-8.5.0-4.el8_5.x86_64.rpm      247 kB/s | 453 kB     00:01
(224/329): libtevent-0.11.0-0.el8.x86_64.rpm        57 kB/s |  50 kB     00:00
(225/329): libuuid-2.32.1-28.el8.x86_64.rpm        151 kB/s |  96 kB     00:00
(226/329): libtirpc-1.1.4-5.el8.x86_64.rpm         123 kB/s | 112 kB     00:00
(227/329): libwbclient-4.14.5-7.el8_5.x86_64.rpm   133 kB/s | 121 kB     00:00
(228/329): libxcrypt-4.1.1-6.el8.x86_64.rpm        120 kB/s |  73 kB     00:00
(229/329): libxml2-2.9.7-9.el8_4.2.x86_64.rpm      382 kB/s | 696 kB     00:01
(230/329): lshw-B.02.19.2-6.el8.x86_64.rpm         222 kB/s | 341 kB     00:01
(231/329): lsscsi-0.32-3.el8.x86_64.rpm            115 kB/s |  71 kB     00:00
(232/329): lua-libs-5.3.4-12.el8.x86_64.rpm        130 kB/s | 118 kB     00:00
(233/329): lvm2-2.03.12-10.el8.x86_64.rpm          508 kB/s | 1.6 MB     00:03
(234/329): lvm2-libs-2.03.12-10.el8.x86_64.rpm     377 kB/s | 1.2 MB     00:03
(235/329): lz4-libs-1.8.3-3.el8_4.x86_64.rpm       106 kB/s |  66 kB     00:00
(236/329): mcelog-175-1.el8.x86_64.rpm             135 kB/s |  83 kB     00:00
(237/329): man-db-2.7.6.1-18.el8.x86_64.rpm        482 kB/s | 887 kB     00:01
(238/329): mdadm-4.2-rc2.el8.x86_64.rpm            301 kB/s | 460 kB     00:01
(239/329): ncurses-6.1-9.20180224.el8.x86_64.rpm   314 kB/s | 387 kB     00:01
(240/329): ncurses-base-6.1-9.20180224.el8.noarch. 133 kB/s |  81 kB     00:00
(241/329): ncurses-libs-6.1-9.20180224.el8.x86_64. 268 kB/s | 334 kB     00:01
(242/329): nettle-3.4.1-7.el8.x86_64.rpm           244 kB/s | 301 kB     00:01
(243/329): nftables-0.9.3-21.el8.x86_64.rpm        257 kB/s | 321 kB     00:01
(244/329): numactl-libs-2.0.12-13.el8.x86_64.rpm    60 kB/s |  36 kB     00:00
(245/329): microcode_ctl-20210608-1.el8.x86_64.rpm 827 kB/s | 5.5 MB     00:06
(246/329): nvme-cli-1.14-3.el8.x86_64.rpm          490 kB/s | 463 kB     00:00
(247/329): openldap-2.4.46-18.el8.x86_64.rpm       244 kB/s | 352 kB     00:01
(248/329): openssh-8.0p1-10.el8.x86_64.rpm         233 kB/s | 522 kB     00:02
(249/329): openssh-clients-8.0p1-10.el8.x86_64.rpm 409 kB/s | 668 kB     00:01
(250/329): openssh-server-8.0p1-10.el8.x86_64.rpm  395 kB/s | 485 kB     00:01
(251/329): openssl-1.1.1k-5.el8_5.x86_64.rpm       320 kB/s | 709 kB     00:02
(252/329): os-prober-1.74-9.el8.x86_64.rpm          85 kB/s |  51 kB     00:00
(253/329): openssl-libs-1.1.1k-5.el8_5.x86_64.rpm  435 kB/s | 1.5 MB     00:03
(254/329): pam-1.3.1-15.el8.x86_64.rpm             303 kB/s | 739 kB     00:02
(255/329): parted-3.2-39.el8.x86_64.rpm            658 kB/s | 555 kB     00:00
(256/329): platform-python-3.6.8-41.el8.x86_64.rpm 105 kB/s |  85 kB     00:00
(257/329): pcre-8.42-6.el8.x86_64.rpm              175 kB/s | 211 kB     00:01
(258/329): policycoreutils-2.9-16.el8.x86_64.rpm   322 kB/s | 373 kB     00:01
(259/329): policycoreutils-python-utils-2.9-16.el8 413 kB/s | 252 kB     00:00
(260/329): polkit-0.115-12.el8.x86_64.rpm          255 kB/s | 154 kB     00:00
(261/329): polkit-libs-0.115-12.el8.x86_64.rpm     129 kB/s |  76 kB     00:00
(262/329): platform-python-pip-9.0.3-20.el8.noarch 471 kB/s | 1.7 MB     00:03
(263/329): python3-dnf-4.7.0-4.el8.noarch.rpm      448 kB/s | 545 kB     00:01
(264/329): python3-dnf-plugins-core-4.0.21-3.el8.n 179 kB/s | 234 kB     00:01
(265/329): python3-firewall-0.9.3-7.el8.noarch.rpm 472 kB/s | 432 kB     00:00
(266/329): python3-hawkey-0.63.0-3.el8.x86_64.rpm  197 kB/s | 116 kB     00:00
(267/329): python3-gpg-1.13.1-9.el8.x86_64.rpm     201 kB/s | 245 kB     00:01
(268/329): python3-libcomps-0.1.16-2.el8.x86_64.rp  87 kB/s |  51 kB     00:00
(269/329): python3-libdnf-0.63.0-3.el8.x86_64.rpm  357 kB/s | 777 kB     00:02
(270/329): python3-libstoragemgmt-1.9.1-1.el8.x86_ 184 kB/s | 175 kB     00:00
(271/329): python3-libxml2-2.9.7-9.el8_4.2.x86_64. 183 kB/s | 237 kB     00:01
(272/329): python3-nftables-0.9.3-21.el8.x86_64.rp  44 kB/s |  29 kB     00:00
(273/329): python3-libs-3.6.8-41.el8.x86_64.rpm    1.2 MB/s | 7.8 MB     00:06
(274/329): python3-pip-wheel-9.0.3-20.el8.noarch.r 836 kB/s | 1.0 MB     00:01
(275/329): python3-policycoreutils-2.9-16.el8.noar 1.0 MB/s | 2.2 MB     00:02
(276/329): python3-rpm-4.14.3-19.el8.x86_64.rpm    258 kB/s | 154 kB     00:00
(277/329): python3-sssdconfig-2.5.2-2.el8_5.3.noar 230 kB/s | 142 kB     00:00
(278/329): python3-syspurpose-1.28.21-3.el8.x86_64 128 kB/s | 312 kB     00:02
(279/329): quota-4.04-14.el8.x86_64.rpm            180 kB/s | 214 kB     00:01
(280/329): quota-nls-4.04-14.el8.noarch.rpm        156 kB/s |  95 kB     00:00
(281/329): python3-perf-4.18.0-348.7.1.el8_5.x86_6 697 kB/s | 7.1 MB     00:10
(282/329): rdma-core-35.0-1.el8.x86_64.rpm          99 kB/s |  59 kB     00:00
(283/329): realmd-0.16.3-23.el8.x86_64.rpm         193 kB/s | 238 kB     00:01
(284/329): rpm-build-libs-4.14.3-19.el8.x86_64.rpm 165 kB/s | 156 kB     00:00
(285/329): rpm-4.14.3-19.el8.x86_64.rpm            311 kB/s | 543 kB     00:01
(286/329): rpm-plugin-selinux-4.14.3-19.el8.x86_64 131 kB/s |  77 kB     00:00
(287/329): rpm-libs-4.14.3-19.el8.x86_64.rpm       281 kB/s | 344 kB     00:01
(288/329): rpm-plugin-systemd-inhibit-4.14.3-19.el  91 kB/s |  78 kB     00:00
(289/329): samba-common-4.14.5-7.el8_5.noarch.rpm  244 kB/s | 221 kB     00:00
(290/329): samba-common-libs-4.14.5-7.el8_5.x86_64 190 kB/s | 173 kB     00:00
(291/329): selinux-policy-3.14.3-80.el8_5.2.noarch 418 kB/s | 636 kB     00:01
(292/329): samba-client-libs-4.14.5-7.el8_5.x86_64 843 kB/s | 5.4 MB     00:06
(293/329): shadow-utils-4.6-14.el8.x86_64.rpm      647 kB/s | 1.2 MB     00:01
(294/329): sos-4.1-5.el8.noarch.rpm                441 kB/s | 706 kB     00:01
(295/329): sqlite-3.26.0-15.el8.x86_64.rpm         426 kB/s | 668 kB     00:01
(296/329): sqlite-libs-3.26.0-15.el8.x86_64.rpm    434 kB/s | 581 kB     00:01
(297/329): sssd-2.5.2-2.el8_5.3.x86_64.rpm         174 kB/s | 107 kB     00:00
(298/329): sssd-ad-2.5.2-2.el8_5.3.x86_64.rpm      278 kB/s | 271 kB     00:00
(299/329): sssd-client-2.5.2-2.el8_5.3.x86_64.rpm  221 kB/s | 205 kB     00:00
(300/329): selinux-policy-targeted-3.14.3-80.el8_5 1.1 MB/s |  15 MB     00:13
(301/329): sssd-common-pac-2.5.2-2.el8_5.3.x86_64. 277 kB/s | 179 kB     00:00
(302/329): sssd-ipa-2.5.2-2.el8_5.3.x86_64.rpm     387 kB/s | 348 kB     00:00
(303/329): sssd-common-2.5.2-2.el8_5.3.x86_64.rpm  491 kB/s | 1.6 MB     00:03
(304/329): sssd-kcm-2.5.2-2.el8_5.3.x86_64.rpm     397 kB/s | 254 kB     00:00
(305/329): sssd-krb5-2.5.2-2.el8_5.3.x86_64.rpm    167 kB/s | 150 kB     00:00
(306/329): sssd-krb5-common-2.5.2-2.el8_5.3.x86_64 151 kB/s | 186 kB     00:01
(307/329): sssd-ldap-2.5.2-2.el8_5.3.x86_64.rpm    240 kB/s | 209 kB     00:00
(308/329): sssd-nfs-idmap-2.5.2-2.el8_5.3.x86_64.r 123 kB/s | 116 kB     00:00
(309/329): sssd-proxy-2.5.2-2.el8_5.3.x86_64.rpm   246 kB/s | 148 kB     00:00
(310/329): sudo-1.8.29-7.el8_4.1.x86_64.rpm        564 kB/s | 925 kB     00:01
(311/329): strace-5.7-3.el8.x86_64.rpm             535 kB/s | 1.1 MB     00:02
(312/329): systemd-libs-239-51.el8_5.2.x86_64.rpm  499 kB/s | 1.1 MB     00:02
(313/329): systemd-239-51.el8_5.2.x86_64.rpm       1.2 MB/s | 3.6 MB     00:03
(314/329): systemd-pam-239-51.el8_5.2.x86_64.rpm   377 kB/s | 477 kB     00:01
(315/329): systemd-udev-239-51.el8_5.2.x86_64.rpm  931 kB/s | 1.6 MB     00:01
(316/329): tpm2-tss-2.3.2-4.el8.x86_64.rpm         446 kB/s | 275 kB     00:00
(317/329): tpm2-tools-4.1.1-5.el8.x86_64.rpm       455 kB/s | 1.0 MB     00:02
(318/329): tuned-2.16.0-1.el8.noarch.rpm           341 kB/s | 312 kB     00:00
(319/329): unzip-6.0-45.el8_4.x86_64.rpm           189 kB/s | 195 kB     00:01
(320/329): tzdata-2021e-1.el8.noarch.rpm           382 kB/s | 474 kB     00:01
(321/329): util-linux-user-2.32.1-28.el8.x86_64.rp 169 kB/s | 100 kB     00:00
(322/329): vdo-6.2.5.74-14.el8.x86_64.rpm          550 kB/s | 663 kB     00:01
(323/329): vim-minimal-8.0.1763-16.el8.x86_64.rpm  594 kB/s | 573 kB     00:00
(324/329): util-linux-2.32.1-28.el8.x86_64.rpm     749 kB/s | 2.5 MB     00:03
(325/329): virt-what-1.18-12.el8.x86_64.rpm         63 kB/s |  36 kB     00:00
(326/329): which-2.21-16.el8.x86_64.rpm             82 kB/s |  49 kB     00:00
(327/329): yum-4.7.0-4.el8.noarch.rpm              225 kB/s | 205 kB     00:00
(328/329): xfsprogs-5.0.0-9.el8.x86_64.rpm         735 kB/s | 1.1 MB     00:01
(329/329): linux-firmware-20210702-103.gitd79c2677 1.2 MB/s | 161 MB     02:12
-----------------------------------------------------------------------------------
Total                                              1.9 MB/s | 500 MB     04:26
Running transaction check
Transaction check succeeded.
Running transaction test
Transaction test succeeded.
Running transaction
  Running scriptlet: filesystem-3.8-6.el8.x86_64                               1/1
  Running scriptlet: kmod-kvdo-6.2.5.72-81.el8.x86_64                          1/1
  Preparing        :                                                           1/1
  Running scriptlet: libgcc-8.5.0-4.el8_5.x86_64                               1/1
  Upgrading        : libgcc-8.5.0-4.el8_5.x86_64                             1/654
  Running scriptlet: libgcc-8.5.0-4.el8_5.x86_64                             1/654
  Upgrading        : bind-license-32:9.11.26-6.el8.noarch                    2/654
  Upgrading        : python3-pip-wheel-9.0.3-20.el8.noarch                   3/654
  Upgrading        : filesystem-3.8-6.el8.x86_64                             4/654
  Upgrading        : tzdata-2021e-1.el8.noarch                               5/654
  Upgrading        : quota-nls-1:4.04-14.el8.noarch                          6/654
  Upgrading        : ncurses-base-6.1-9.20180224.el8.noarch                  7/654
  Upgrading        : ncurses-libs-6.1-9.20180224.el8.x86_64                  8/654
  Running scriptlet: glibc-2.28-164.el8.x86_64                               9/654
  Upgrading        : glibc-2.28-164.el8.x86_64                               9/654
  Running scriptlet: glibc-2.28-164.el8.x86_64                               9/654
  Upgrading        : bash-4.4.20-2.el8.x86_64                               10/654
  Running scriptlet: bash-4.4.20-2.el8.x86_64                               10/654
  Upgrading        : glibc-common-2.28-164.el8.x86_64                       11/654
  Upgrading        : glibc-langpack-en-2.28-164.el8.x86_64                  12/654
  Upgrading        : libcom_err-1.45.6-2.el8.x86_64                         13/654
  Running scriptlet: libcom_err-1.45.6-2.el8.x86_64                         13/654
  Upgrading        : libcap-2.26-5.el8.x86_64                               14/654
  Upgrading        : libtalloc-2.3.2-1.el8.x86_64                           15/654
  Upgrading        : sqlite-libs-3.26.0-15.el8.x86_64                       16/654
  Upgrading        : elfutils-libelf-0.185-1.el8.x86_64                     17/654
  Upgrading        : libuuid-2.32.1-28.el8.x86_64                           18/654
  Running scriptlet: libuuid-2.32.1-28.el8.x86_64                           18/654
  Upgrading        : libxml2-2.9.7-9.el8_4.2.x86_64                         19/654
  Upgrading        : libtevent-0.11.0-0.el8.x86_64                          20/654
  Upgrading        : libxcrypt-4.1.1-6.el8.x86_64                           21/654
  Upgrading        : libsepol-2.9-3.el8.x86_64                              22/654
  Running scriptlet: libsepol-2.9-3.el8.x86_64                              22/654
  Upgrading        : libstdc++-8.5.0-4.el8_5.x86_64                         23/654
  Running scriptlet: libstdc++-8.5.0-4.el8_5.x86_64                         23/654
  Upgrading        : gpgme-1.13.1-9.el8.x86_64                              24/654
  Upgrading        : lua-libs-5.3.4-12.el8.x86_64                           25/654
  Upgrading        : which-2.21-16.el8.x86_64                               26/654
  Upgrading        : chkconfig-1.19.1-1.el8.x86_64                          27/654
  Upgrading        : grub2-common-1:2.02-106.el8.noarch                     28/654
  Upgrading        : nspr-4.32.0-1.el8_4.x86_64                             29/654
  Running scriptlet: nspr-4.32.0-1.el8_4.x86_64                             29/654
  Upgrading        : json-c-0.13.1-2.el8.x86_64                             30/654
  Upgrading        : keyutils-libs-1.5.10-9.el8.x86_64                      31/654
  Upgrading        : libsss_idmap-2.5.2-2.el8_5.3.x86_64                    32/654
  Running scriptlet: libsss_idmap-2.5.2-2.el8_5.3.x86_64                    32/654
  Upgrading        : nss-util-3.67.0-7.el8_5.x86_64                         33/654
  Upgrading        : file-libs-5.33-20.el8.x86_64                           34/654
  Upgrading        : file-5.33-20.el8.x86_64                                35/654
  Upgrading        : libcap-ng-0.7.11-1.el8.x86_64                          36/654
  Upgrading        : libsmartcols-2.32.1-28.el8.x86_64                      37/654
  Running scriptlet: libsmartcols-2.32.1-28.el8.x86_64                      37/654
  Upgrading        : fstrm-0.6.1-2.el8.x86_64                               38/654
  Upgrading        : device-mapper-persistent-data-0.9.0-4.el8.x86_64       39/654
  Upgrading        : e2fsprogs-libs-1.45.6-2.el8.x86_64                     40/654
  Running scriptlet: e2fsprogs-libs-1.45.6-2.el8.x86_64                     40/654
  Upgrading        : nettle-3.4.1-7.el8.x86_64                              41/654
  Running scriptlet: nettle-3.4.1-7.el8.x86_64                              41/654
  Upgrading        : dmidecode-1:3.2-10.el8.x86_64                          42/654
  Upgrading        : ethtool-2:5.8-7.el8.x86_64                             43/654
  Upgrading        : iptables-libs-1.8.4-20.el8.x86_64                      44/654
  Running scriptlet: iptables-1.8.4-20.el8.x86_64                           45/654
  Upgrading        : iptables-1.8.4-20.el8.x86_64                           45/654
  Running scriptlet: iptables-1.8.4-20.el8.x86_64                           45/654
  Upgrading        : nftables-1:0.9.3-21.el8.x86_64                         46/654
  Running scriptlet: nftables-1:0.9.3-21.el8.x86_64                         46/654
  Upgrading        : libgcrypt-1.8.5-6.el8.x86_64                           47/654
  Running scriptlet: libgcrypt-1.8.5-6.el8.x86_64                           47/654
  Upgrading        : lz4-libs-1.8.3-3.el8_4.x86_64                          48/654
  Upgrading        : iptables-ebtables-1.8.4-20.el8.x86_64                  49/654
  Running scriptlet: iptables-ebtables-1.8.4-20.el8.x86_64                  49/654
  Upgrading        : nss-softokn-freebl-3.67.0-7.el8_5.x86_64               50/654
  Upgrading        : nss-softokn-3.67.0-7.el8_5.x86_64                      51/654
  Upgrading        : grub2-pc-modules-1:2.02-106.el8.noarch                 52/654
  Upgrading        : libcomps-0.1.16-2.el8.x86_64                           53/654
  Installing       : libbpf-0.4.0-1.el8.x86_64                              54/654
  Upgrading        : sqlite-3.26.0-15.el8.x86_64                            55/654
  Upgrading        : libss-1.45.6-2.el8.x86_64                              56/654
  Running scriptlet: libss-1.45.6-2.el8.x86_64                              56/654
  Upgrading        : coreutils-common-8.30-12.el8.x86_64                    57/654
  Running scriptlet: coreutils-common-8.30-12.el8.x86_64                    57/654
  Upgrading        : kernel-tools-libs-4.18.0-348.7.1.el8_5.x86_64          58/654
  Running scriptlet: kernel-tools-libs-4.18.0-348.7.1.el8_5.x86_64          58/654
  Upgrading        : containernetworking-plugins-1.0.0-1.module_el8.5.0+    59/654
  Upgrading        : libdrm-2.4.106-2.el8.x86_64                            60/654
  Upgrading        : libfastjson-0.99.9-1.el8.x86_64                        61/654
  Running scriptlet: libfastjson-0.99.9-1.el8.x86_64                        61/654
  Upgrading        : hdparm-9.54-4.el8.x86_64                               62/654
  Upgrading        : libndp-1.7-6.el8.x86_64                                63/654
  Running scriptlet: libndp-1.7-6.el8.x86_64                                63/654
  Upgrading        : libsss_autofs-2.5.2-2.el8_5.3.x86_64                   64/654
  Upgrading        : libsss_nss_idmap-2.5.2-2.el8_5.3.x86_64                65/654
  Running scriptlet: libsss_nss_idmap-2.5.2-2.el8_5.3.x86_64                65/654
  Upgrading        : libsss_sudo-2.5.2-2.el8_5.3.x86_64                     66/654
  Running scriptlet: libsss_sudo-2.5.2-2.el8_5.3.x86_64                     66/654
  Upgrading        : ncurses-6.1-9.20180224.el8.x86_64                      67/654
  Upgrading        : numactl-libs-2.0.12-13.el8.x86_64                      68/654
  Running scriptlet: numactl-libs-2.0.12-13.el8.x86_64                      68/654
  Upgrading        : pcre-8.42-6.el8.x86_64                                 69/654
  Running scriptlet: pcre-8.42-6.el8.x86_64                                 69/654
  Upgrading        : vim-minimal-2:8.0.1763-16.el8.x86_64                   70/654
  Upgrading        : linux-firmware-20210702-103.gitd79c2677.el8.noarch     71/654
  Upgrading        : libssh-config-0.9.4-3.el8.noarch                       72/654
  Upgrading        : hwdata-0.314-8.10.el8.noarch                           73/654
  Upgrading        : firewalld-filesystem-0.9.3-7.el8.noarch                74/654
  Upgrading        : dnf-data-4.7.0-4.el8.noarch                            75/654
  Upgrading        : dhcp-common-12:4.3.6-45.el8.noarch                     76/654
  Upgrading        : dbus-common-1:1.12.8-14.el8.noarch                     77/654
  Upgrading        : grub2-tools-minimal-1:2.02-106.el8.x86_64              78/654
  Upgrading        : platform-python-pip-9.0.3-20.el8.noarch                79/654
  Upgrading        : python3-libs-3.6.8-41.el8.x86_64                       80/654
  Upgrading        : libssh-0.9.4-3.el8.x86_64                              81/654
  Upgrading        : openldap-2.4.46-18.el8.x86_64                          82/654
  Upgrading        : platform-python-3.6.8-41.el8.x86_64                    83/654
  Running scriptlet: platform-python-3.6.8-41.el8.x86_64                    83/654
  Upgrading        : grubby-8.40-42.el8.x86_64                              84/654
  Upgrading        : libdb-utils-5.3.28-42.el8_4.x86_64                     85/654
  Upgrading        : curl-7.61.1-22.el8.x86_64                              86/654
  Upgrading        : libcurl-7.61.1-22.el8.x86_64                           87/654
  Upgrading        : crypto-policies-scripts-20210617-1.gitc776d3e.el8.n    88/654
  Upgrading        : crypto-policies-20210617-1.gitc776d3e.el8.noarch       89/654
  Running scriptlet: crypto-policies-20210617-1.gitc776d3e.el8.noarch       89/654
  Upgrading        : gnutls-3.6.16-4.el8.x86_64                             90/654
  Upgrading        : elfutils-default-yama-scope-0.185-1.el8.noarch         91/654
  Running scriptlet: elfutils-default-yama-scope-0.185-1.el8.noarch         91/654
  Upgrading        : device-mapper-8:1.02.177-10.el8.x86_64                 92/654
  Upgrading        : kpartx-0.8.4-17.el8.x86_64                             93/654
  Upgrading        : krb5-libs-1.18.2-14.el8.x86_64                         94/654
  Upgrading        : libtirpc-1.1.4-5.el8.x86_64                            95/654
  Running scriptlet: libtirpc-1.1.4-5.el8.x86_64                            95/654
  Upgrading        : elfutils-debuginfod-client-0.185-1.el8.x86_64          96/654
  Upgrading        : elfutils-libs-0.185-1.el8.x86_64                       97/654
  Upgrading        : rpm-4.14.3-19.el8.x86_64                               98/654
  Upgrading        : libfdisk-2.32.1-28.el8.x86_64                          99/654
  Running scriptlet: libfdisk-2.32.1-28.el8.x86_64                          99/654
  Upgrading        : libmount-2.32.1-28.el8.x86_64                         100/654
  Running scriptlet: libmount-2.32.1-28.el8.x86_64                         100/654
  Upgrading        : dbus-libs-1:1.12.8-14.el8.x86_64                      101/654
  Running scriptlet: dbus-libs-1:1.12.8-14.el8.x86_64                      101/654
  Upgrading        : dbus-tools-1:1.12.8-14.el8.x86_64                     102/654
  Upgrading        : device-mapper-libs-8:1.02.177-10.el8.x86_64           103/654
  Upgrading        : coreutils-8.30-12.el8.x86_64                          104/654
  Upgrading        : systemd-libs-239-51.el8_5.2.x86_64                    105/654
  Running scriptlet: systemd-libs-239-51.el8_5.2.x86_64                    105/654
  Upgrading        : libblkid-2.32.1-28.el8.x86_64                         106/654
  Running scriptlet: libblkid-2.32.1-28.el8.x86_64                         106/654
  Upgrading        : shadow-utils-2:4.6-14.el8.x86_64                      107/654
  Running scriptlet: ca-certificates-2021.2.50-80.0.el8_4.noarch           108/654
  Upgrading        : ca-certificates-2021.2.50-80.0.el8_4.noarch           108/654
  Running scriptlet: ca-certificates-2021.2.50-80.0.el8_4.noarch           108/654
  Upgrading        : openssl-libs-1:1.1.1k-5.el8_5.x86_64                  109/654
  Running scriptlet: openssl-libs-1:1.1.1k-5.el8_5.x86_64                  109/654
  Upgrading        : kmod-libs-25-18.el8.x86_64                            110/654
  Running scriptlet: kmod-libs-25-18.el8.x86_64                            110/654
  Upgrading        : libdb-5.3.28-42.el8_4.x86_64                          111/654
  Running scriptlet: libdb-5.3.28-42.el8_4.x86_64                          111/654
  Upgrading        : pam-1.3.1-15.el8.x86_64                               112/654
  Running scriptlet: pam-1.3.1-15.el8.x86_64                               112/654
  Upgrading        : util-linux-2.32.1-28.el8.x86_64                       113/654
  Running scriptlet: util-linux-2.32.1-28.el8.x86_64                       113/654
  Upgrading        : rpm-libs-4.14.3-19.el8.x86_64                         114/654
  Running scriptlet: rpm-libs-4.14.3-19.el8.x86_64                         114/654
  Upgrading        : kmod-25-18.el8.x86_64                                 115/654
  Running scriptlet: dbus-daemon-1:1.12.8-14.el8.x86_64                    116/654
  Upgrading        : dbus-daemon-1:1.12.8-14.el8.x86_64                    116/654
  Running scriptlet: dbus-daemon-1:1.12.8-14.el8.x86_64                    116/654
  Upgrading        : dracut-049-191.git20210920.el8.x86_64                 117/654
  Upgrading        : systemd-pam-239-51.el8_5.2.x86_64                     118/654
  Upgrading        : os-prober-1.74-9.el8.x86_64                           119/654
  Running scriptlet: grub2-tools-1:2.02-106.el8.x86_64                     120/654
  Upgrading        : grub2-tools-1:2.02-106.el8.x86_64                     120/654
  Running scriptlet: grub2-tools-1:2.02-106.el8.x86_64                     120/654
  Upgrading        : dbus-1:1.12.8-14.el8.x86_64                           121/654
  Running scriptlet: systemd-239-51.el8_5.2.x86_64                         122/654
  Upgrading        : systemd-239-51.el8_5.2.x86_64                         122/654
  Running scriptlet: systemd-239-51.el8_5.2.x86_64                         122/654
  Upgrading        : systemd-udev-239-51.el8_5.2.x86_64                    123/654
  Running scriptlet: systemd-udev-239-51.el8_5.2.x86_64                    123/654
  Upgrading        : glib2-2.56.4-156.el8.x86_64                           124/654
  Upgrading        : libldb-2.3.0-2.el8.x86_64                             125/654
  Upgrading        : polkit-libs-0.115-12.el8.x86_64                       126/654
  Running scriptlet: polkit-libs-0.115-12.el8.x86_64                       126/654
  Running scriptlet: polkit-0.115-12.el8.x86_64                            127/654
  Upgrading        : polkit-0.115-12.el8.x86_64                            127/654
  Running scriptlet: polkit-0.115-12.el8.x86_64                            127/654
  Upgrading        : cockpit-bridge-251.1-1.el8.x86_64                     128/654
  Upgrading        : libmodulemd-2.13.0-1.el8.x86_64                       129/654
  Upgrading        : policycoreutils-2.9-16.el8.x86_64                     130/654
  Running scriptlet: policycoreutils-2.9-16.el8.x86_64                     130/654
  Upgrading        : libsss_certmap-2.5.2-2.el8_5.3.x86_64                 131/654
  Running scriptlet: libsss_certmap-2.5.2-2.el8_5.3.x86_64                 131/654
  Upgrading        : device-mapper-event-libs-8:1.02.177-10.el8.x86_64     132/654
  Upgrading        : librepo-1.14.0-2.el8.x86_64                           133/654
  Installing       : kernel-core-4.18.0-348.7.1.el8_5.x86_64               134/654
  Running scriptlet: kernel-core-4.18.0-348.7.1.el8_5.x86_64               134/654
  Upgrading        : rsyslog-8.2102.0-5.el8.x86_64                         135/654
  Running scriptlet: rsyslog-8.2102.0-5.el8.x86_64                         135/654
  Running scriptlet: samba-common-4.14.5-7.el8_5.noarch                    136/654
  Upgrading        : samba-common-4.14.5-7.el8_5.noarch                    136/654
  Running scriptlet: samba-common-4.14.5-7.el8_5.noarch                    136/654
  Upgrading        : libsolv-0.7.19-1.el8.x86_64                           137/654
  Upgrading        : iproute-5.12.0-4.el8.x86_64                           138/654
  Upgrading        : parted-3.2-39.el8.x86_64                              139/654
  Running scriptlet: parted-3.2-39.el8.x86_64                              139/654
  Upgrading        : libblockdev-utils-2.24-7.el8.x86_64                   140/654
  Upgrading        : libdnf-0.63.0-3.el8.x86_64                            141/654
  Upgrading        : python3-libdnf-0.63.0-3.el8.x86_64                    142/654
  Upgrading        : python3-hawkey-0.63.0-3.el8.x86_64                    143/654
  Upgrading        : rpm-plugin-selinux-4.14.3-19.el8.x86_64               144/654
  Upgrading        : selinux-policy-3.14.3-80.el8_5.2.noarch               145/654
  Running scriptlet: selinux-policy-3.14.3-80.el8_5.2.noarch               145/654
  Running scriptlet: selinux-policy-targeted-3.14.3-80.el8_5.2.noarch      146/654
  Upgrading        : selinux-policy-targeted-3.14.3-80.el8_5.2.noarch      146/654
  Running scriptlet: selinux-policy-targeted-3.14.3-80.el8_5.2.noarch      146/654
  Upgrading        : NetworkManager-libnm-1:1.32.10-4.el8.x86_64           147/654
  Running scriptlet: NetworkManager-libnm-1:1.32.10-4.el8.x86_64           147/654
  Running scriptlet: NetworkManager-1:1.32.10-4.el8.x86_64                 148/654
  Upgrading        : NetworkManager-1:1.32.10-4.el8.x86_64                 148/654
  Running scriptlet: NetworkManager-1:1.32.10-4.el8.x86_64                 148/654
  Upgrading        : iscsi-initiator-utils-iscsiuio-6.2.1.4-4.git095f59c   149/654
  Running scriptlet: iscsi-initiator-utils-iscsiuio-6.2.1.4-4.git095f59c   149/654
  Upgrading        : iscsi-initiator-utils-6.2.1.4-4.git095f59c.el8.x86_   150/654
  Running scriptlet: iscsi-initiator-utils-6.2.1.4-4.git095f59c.el8.x86_   150/654
  Upgrading        : fuse-overlayfs-1.7.1-1.module_el8.5.0+890+6b136101.   151/654
  Running scriptlet: fuse-overlayfs-1.7.1-1.module_el8.5.0+890+6b136101.   151/654
  Running scriptlet: openssh-8.0p1-10.el8.x86_64                           152/654
  Upgrading        : openssh-8.0p1-10.el8.x86_64                           152/654
  Upgrading        : bind-libs-lite-32:9.11.26-6.el8.x86_64                153/654
  Upgrading        : python3-libxml2-2.9.7-9.el8_4.2.x86_64                154/654
  Upgrading        : sos-4.1-5.el8.noarch                                  155/654
  Upgrading        : bind-libs-32:9.11.26-6.el8.x86_64                     156/654
  Upgrading        : NetworkManager-team-1:1.32.10-4.el8.x86_64            157/654
  Upgrading        : libblockdev-2.24-7.el8.x86_64                         158/654
  Upgrading        : libblockdev-fs-2.24-7.el8.x86_64                      159/654
  Upgrading        : libblockdev-loop-2.24-7.el8.x86_64                    160/654
  Upgrading        : libblockdev-part-2.24-7.el8.x86_64                    161/654
  Upgrading        : libblockdev-swap-2.24-7.el8.x86_64                    162/654
  Installing       : kernel-modules-4.18.0-348.7.1.el8_5.x86_64            163/654
  Running scriptlet: kernel-modules-4.18.0-348.7.1.el8_5.x86_64            163/654
  Running scriptlet: kmod-kvdo-6.2.5.72-81.el8.x86_64                      164/654
  Upgrading        : kmod-kvdo-6.2.5.72-81.el8.x86_64                      164/654
  Running scriptlet: kmod-kvdo-6.2.5.72-81.el8.x86_64                      164/654
  Upgrading        : device-mapper-event-8:1.02.177-10.el8.x86_64          165/654
  Running scriptlet: device-mapper-event-8:1.02.177-10.el8.x86_64          165/654
  Upgrading        : lvm2-libs-8:2.03.12-10.el8.x86_64                     166/654
  Upgrading        : lvm2-8:2.03.12-10.el8.x86_64                          167/654
  Running scriptlet: lvm2-8:2.03.12-10.el8.x86_64                          167/654
  Upgrading        : libblockdev-lvm-2.24-7.el8.x86_64                     168/654
  Upgrading        : python3-policycoreutils-2.9-16.el8.noarch             169/654
  Upgrading        : policycoreutils-python-utils-2.9-16.el8.noarch        170/654
  Running scriptlet: container-selinux-2:2.167.0-1.module_el8.5.0+911+f1   171/654
  Upgrading        : container-selinux-2:2.167.0-1.module_el8.5.0+911+f1   171/654
  Running scriptlet: container-selinux-2:2.167.0-1.module_el8.5.0+911+f1   171/654
  Upgrading        : cockpit-packagekit-251.1-1.el8.noarch                 172/654
  Upgrading        : conmon-2:2.0.29-1.module_el8.5.0+890+6b136101.x86_6   173/654
  Upgrading        : libslirp-4.4.0-1.module_el8.5.0+890+6b136101.x86_64   174/654
  Upgrading        : slirp4netns-1.1.8-1.module_el8.5.0+890+6b136101.x86   175/654
  Upgrading        : libudisks2-2.9.0-7.el8.x86_64                         176/654
  Upgrading        : libipa_hbac-2.5.2-2.el8_5.3.x86_64                    177/654
  Running scriptlet: libipa_hbac-2.5.2-2.el8_5.3.x86_64                    177/654
  Running scriptlet: libstoragemgmt-1.9.1-1.el8.x86_64                     178/654
  Upgrading        : libstoragemgmt-1.9.1-1.el8.x86_64                     178/654
  Running scriptlet: libstoragemgmt-1.9.1-1.el8.x86_64                     178/654
  Upgrading        : python3-libstoragemgmt-1.9.1-1.el8.x86_64             179/654
  Running scriptlet: unbound-libs-1.7.3-17.el8.x86_64                      180/654
  Upgrading        : unbound-libs-1.7.3-17.el8.x86_64                      180/654
  Running scriptlet: unbound-libs-1.7.3-17.el8.x86_64                      180/654
  Upgrading        : python3-unbound-1.7.3-17.el8.x86_64                   181/654
  Running scriptlet: authselect-libs-1.2.2-3.el8.x86_64                    182/654
  Upgrading        : authselect-libs-1.2.2-3.el8.x86_64                    182/654
  Upgrading        : mdadm-4.2-rc2.el8.x86_64                              183/654
  Running scriptlet: mdadm-4.2-rc2.el8.x86_64                              183/654
  Upgrading        : libblockdev-mdraid-2.24-7.el8.x86_64                  184/654
  Upgrading        : grub2-tools-extra-1:2.02-106.el8.x86_64               185/654
  Upgrading        : dracut-squash-049-191.git20210920.el8.x86_64          186/654
  Upgrading        : rpm-build-libs-4.14.3-19.el8.x86_64                   187/654
  Running scriptlet: rpm-build-libs-4.14.3-19.el8.x86_64                   187/654
  Upgrading        : python3-rpm-4.14.3-19.el8.x86_64                      188/654
  Running scriptlet: setroubleshoot-server-3.3.24-4.el8.x86_64             189/654
  Upgrading        : setroubleshoot-server-3.3.24-4.el8.x86_64             189/654
  Running scriptlet: setroubleshoot-server-3.3.24-4.el8.x86_64             189/654
  Upgrading        : setroubleshoot-plugins-3.3.14-1.el8.noarch            190/654
  Upgrading        : rpm-plugin-systemd-inhibit-4.14.3-19.el8.x86_64       191/654
  Upgrading        : virt-what-1.18-12.el8.x86_64                          192/654
  Upgrading        : sssd-client-2.5.2-2.el8_5.3.x86_64                    193/654
  Running scriptlet: sssd-client-2.5.2-2.el8_5.3.x86_64                    193/654
  Upgrading        : sudo-1.8.29-7.el8_4.1.x86_64                          194/654
  Running scriptlet: sudo-1.8.29-7.el8_4.1.x86_64                          194/654
  Upgrading        : librelp-1.9.0-1.el8.x86_64                            195/654
  Running scriptlet: librelp-1.9.0-1.el8.x86_64                            195/654
  Upgrading        : bind-export-libs-32:9.11.26-6.el8.x86_64              196/654
  Running scriptlet: bind-export-libs-32:9.11.26-6.el8.x86_64              196/654
  Upgrading        : openssl-1:1.1.1k-5.el8_5.x86_64                       197/654
  Running scriptlet: tpm2-tss-2.3.2-4.el8.x86_64                           198/654
  Upgrading        : tpm2-tss-2.3.2-4.el8.x86_64                           198/654
  Running scriptlet: tpm2-tss-2.3.2-4.el8.x86_64                           198/654
  Upgrading        : e2fsprogs-1.45.6-2.el8.x86_64                         199/654
  Upgrading        : xfsprogs-5.0.0-9.el8.x86_64                           200/654
  Running scriptlet: xfsprogs-5.0.0-9.el8.x86_64                           200/654
  Upgrading        : plymouth-core-libs-0.9.4-10.20200615git1e36e30.el8.   201/654
  Running scriptlet: plymouth-core-libs-0.9.4-10.20200615git1e36e30.el8.   201/654
  Upgrading        : plymouth-scripts-0.9.4-10.20200615git1e36e30.el8.x8   202/654
  Upgrading        : plymouth-0.9.4-10.20200615git1e36e30.el8.x86_64       203/654
  Upgrading        : device-mapper-multipath-libs-0.8.4-17.el8.x86_64      204/654
  Running scriptlet: device-mapper-multipath-libs-0.8.4-17.el8.x86_64      204/654
  Upgrading        : device-mapper-multipath-0.8.4-17.el8.x86_64           205/654
  Running scriptlet: device-mapper-multipath-0.8.4-17.el8.x86_64           205/654
  Upgrading        : dhcp-libs-12:4.3.6-45.el8.x86_64                      206/654
  Upgrading        : dhcp-client-12:4.3.6-45.el8.x86_64                    207/654
  Upgrading        : dracut-network-049-191.git20210920.el8.x86_64         208/654
  Upgrading        : kexec-tools-2.0.20-57.el8_5.1.x86_64                  209/654
  Running scriptlet: kexec-tools-2.0.20-57.el8_5.1.x86_64                  209/654
  Upgrading        : cockpit-system-251.1-1.el8.noarch                     210/654
  Upgrading        : rdma-core-35.0-1.el8.x86_64                           211/654
  Running scriptlet: rdma-core-35.0-1.el8.x86_64                           211/654
  Upgrading        : nss-3.67.0-7.el8_5.x86_64                             212/654
  Upgrading        : nss-sysinit-3.67.0-7.el8_5.x86_64                     213/654
  Upgrading        : libblockdev-crypto-2.24-7.el8.x86_64                  214/654
  Upgrading        : udisks2-2.9.0-7.el8.x86_64                            215/654
  Running scriptlet: udisks2-2.9.0-7.el8.x86_64                            215/654
  Upgrading        : udisks2-iscsi-2.9.0-7.el8.x86_64                      216/654
  Upgrading        : udisks2-lvm2-2.9.0-7.el8.x86_64                       217/654
  Upgrading        : cockpit-storaged-251.1-1.el8.noarch                   218/654
  Upgrading        : binutils-2.30-108.el8_5.1.x86_64                      219/654
  Running scriptlet: binutils-2.30-108.el8_5.1.x86_64                      219/654
  Upgrading        : centos-logos-85.8-2.el8.x86_64                        220/654
  Running scriptlet: centos-logos-85.8-2.el8.x86_64                        220/654
  Running scriptlet: cockpit-ws-251.1-1.el8.x86_64                         221/654
  Upgrading        : cockpit-ws-251.1-1.el8.x86_64                         221/654
  Running scriptlet: cockpit-ws-251.1-1.el8.x86_64                         221/654
  Upgrading        : libnfsidmap-1:2.3.3-46.el8.x86_64                     222/654
  Upgrading        : sssd-nfs-idmap-2.5.2-2.el8_5.3.x86_64                 223/654
  Running scriptlet: sssd-common-2.5.2-2.el8_5.3.x86_64                    224/654
  Upgrading        : sssd-common-2.5.2-2.el8_5.3.x86_64                    224/654
  Running scriptlet: sssd-common-2.5.2-2.el8_5.3.x86_64                    224/654
  Running scriptlet: sssd-krb5-common-2.5.2-2.el8_5.3.x86_64               225/654
  Upgrading        : sssd-krb5-common-2.5.2-2.el8_5.3.x86_64               225/654
  Upgrading        : sssd-krb5-2.5.2-2.el8_5.3.x86_64                      226/654
  Upgrading        : sssd-ldap-2.5.2-2.el8_5.3.x86_64                      227/654
  Running scriptlet: sssd-proxy-2.5.2-2.el8_5.3.x86_64                     228/654
  Upgrading        : sssd-proxy-2.5.2-2.el8_5.3.x86_64                     228/654
  Upgrading        : adcli-0.8.2-12.el8.x86_64                             229/654
  Running scriptlet: adcli-0.8.2-12.el8.x86_64                             229/654
  Upgrading        : cups-libs-1:2.2.6-40.el8.x86_64                       230/654
  Upgrading        : libwbclient-4.14.5-7.el8_5.x86_64                     231/654
  Upgrading        : samba-client-libs-4.14.5-7.el8_5.x86_64               232/654
  Upgrading        : samba-common-libs-4.14.5-7.el8_5.x86_64               233/654
  Upgrading        : sssd-common-pac-2.5.2-2.el8_5.3.x86_64                234/654
  Upgrading        : libsmbclient-4.14.5-7.el8_5.x86_64                    235/654
  Upgrading        : criu-3.15-3.module_el8.5.0+890+6b136101.x86_64        236/654
  Upgrading        : runc-1.0.2-1.module_el8.5.0+911+f19012f9.x86_64       237/654
  Upgrading        : python3-bind-32:9.11.26-6.el8.noarch                  238/654
  Upgrading        : bind-utils-32:9.11.26-6.el8.x86_64                    239/654
  Upgrading        : sssd-ad-2.5.2-2.el8_5.3.x86_64                        240/654
  Running scriptlet: sssd-ipa-2.5.2-2.el8_5.3.x86_64                       241/654
  Upgrading        : sssd-ipa-2.5.2-2.el8_5.3.x86_64                       241/654
  Upgrading        : kernel-tools-4.18.0-348.7.1.el8_5.x86_64              242/654
  Upgrading        : python3-gpg-1.13.1-9.el8.x86_64                       243/654
  Upgrading        : python3-libcomps-0.1.16-2.el8.x86_64                  244/654
  Upgrading        : python3-dnf-4.7.0-4.el8.noarch                        245/654
  Upgrading        : dnf-4.7.0-4.el8.noarch                                246/654
  Running scriptlet: dnf-4.7.0-4.el8.noarch                                246/654
  Upgrading        : kpatch-dnf-0.2-5.el8.noarch                           247/654
  Running scriptlet: kpatch-dnf-0.2-5.el8.noarch                           247/654
To enable automatic kpatch-patch subscription, run:
        $ dnf kpatch auto

  Upgrading        : python3-dnf-plugins-core-4.0.21-3.el8.noarch          248/654
  Upgrading        : python3-nftables-1:0.9.3-21.el8.x86_64                249/654
  Upgrading        : python3-firewall-0.9.3-7.el8.noarch                   250/654
  Upgrading        : python3-perf-4.18.0-348.7.1.el8_5.x86_64              251/654
  Upgrading        : python3-sssdconfig-2.5.2-2.el8_5.3.noarch             252/654
  Upgrading        : sssd-2.5.2-2.el8_5.3.x86_64                           253/654
  Upgrading        : authselect-1.2.2-3.el8.x86_64                         254/654
  Upgrading        : realmd-0.16.3-23.el8.x86_64                           255/654
  Running scriptlet: realmd-0.16.3-23.el8.x86_64                           255/654
  Upgrading        : python3-syspurpose-1.28.21-3.el8.x86_64               256/654
  Upgrading        : centos-gpg-keys-1:8-3.el8.noarch                      257/654
  Upgrading        : centos-linux-release-8.5-1.2111.el8.noarch            258/654
  Upgrading        : centos-linux-repos-8-3.el8.noarch                     259/654
  Upgrading        : containers-common-2:1-2.module_el8.5.0+890+6b136101   260/654
  Upgrading        : podman-catatonit-3.3.1-9.module_el8.5.0+988+b1f0b74   261/654
  Upgrading        : podman-3.3.1-9.module_el8.5.0+988+b1f0b741.x86_64     262/654
  Upgrading        : vim-filesystem-2:8.0.1763-16.el8.noarch               263/654
  Upgrading        : vim-common-2:8.0.1763-16.el8.x86_64                   264/654
  Upgrading        : libX11-common-1.6.8-5.el8.noarch                      265/654
  Upgrading        : libX11-1.6.8-5.el8.x86_64                             266/654
  Upgrading        : vim-enhanced-2:8.0.1763-16.el8.x86_64                 267/654
  Upgrading        : cockpit-podman-33-1.module_el8.5.0+890+6b136101.noa   268/654
  Upgrading        : buildah-1.22.3-2.module_el8.5.0+911+f19012f9.x86_64   269/654
  Upgrading        : tuned-2.16.0-1.el8.noarch                             270/654
  Running scriptlet: tuned-2.16.0-1.el8.noarch                             270/654
  Upgrading        : authselect-compat-1.2.2-3.el8.x86_64                  271/654
  Upgrading        : firewalld-0.9.3-7.el8.noarch                          272/654
  Running scriptlet: firewalld-0.9.3-7.el8.noarch                          272/654
  Upgrading        : dnf-plugins-core-4.0.21-3.el8.noarch                  273/654
  Upgrading        : kpatch-0.9.2-5.el8.noarch                             274/654
  Upgrading        : yum-4.7.0-4.el8.noarch                                275/654
  Upgrading        : sssd-kcm-2.5.2-2.el8_5.3.x86_64                       276/654
  Running scriptlet: sssd-kcm-2.5.2-2.el8_5.3.x86_64                       276/654
  Upgrading        : cockpit-251.1-1.el8.x86_64                            277/654
  Upgrading        : libibverbs-35.0-1.el8.x86_64                          278/654
  Running scriptlet: libibverbs-35.0-1.el8.x86_64                          278/654
  Upgrading        : tpm2-tools-4.1.1-5.el8.x86_64                         279/654
  Upgrading        : rsyslog-relp-8.2102.0-5.el8.x86_64                    280/654
  Upgrading        : grub2-pc-1:2.02-106.el8.x86_64                        281/654
  Upgrading        : vdo-6.2.5.74-14.el8.x86_64                            282/654
  Running scriptlet: vdo-6.2.5.74-14.el8.x86_64                            282/654
  Installing       : kernel-4.18.0-348.7.1.el8_5.x86_64                    283/654
  Upgrading        : openssh-clients-8.0p1-10.el8.x86_64                   284/654
  Running scriptlet: openssh-server-8.0p1-10.el8.x86_64                    285/654
  Upgrading        : openssh-server-8.0p1-10.el8.x86_64                    285/654
  Running scriptlet: openssh-server-8.0p1-10.el8.x86_64                    285/654
  Upgrading        : NetworkManager-tui-1:1.32.10-4.el8.x86_64             286/654
  Upgrading        : open-vm-tools-11.2.5-2.el8.x86_64                     287/654
  Running scriptlet: open-vm-tools-11.2.5-2.el8.x86_64                     287/654
  Upgrading        : rsyslog-gnutls-8.2102.0-5.el8.x86_64                  288/654
  Upgrading        : rsyslog-gssapi-8.2102.0-5.el8.x86_64                  289/654
  Upgrading        : gsettings-desktop-schemas-3.32.0-6.el8.x86_64         290/654
  Running scriptlet: chrony-4.1-1.el8.x86_64                               291/654
  Upgrading        : chrony-4.1-1.el8.x86_64                               291/654
  Running scriptlet: chrony-4.1-1.el8.x86_64                               291/654
  Upgrading        : mcelog-3:175-1.el8.x86_64                             292/654
  Running scriptlet: mcelog-3:175-1.el8.x86_64                             292/654
  Upgrading        : microcode_ctl-4:20210608-1.el8.x86_64                 293/654
  Running scriptlet: microcode_ctl-4:20210608-1.el8.x86_64                 293/654
  Installing       : grub2-tools-efi-1:2.02-106.el8.x86_64                 294/654
  Upgrading        : dracut-config-rescue-049-191.git20210920.el8.x86_64   295/654
  Upgrading        : nvme-cli-1.14-3.el8.x86_64                            296/654
  Running scriptlet: nvme-cli-1.14-3.el8.x86_64                            296/654
  Upgrading        : util-linux-user-2.32.1-28.el8.x86_64                  297/654
  Upgrading        : nmap-ncat-2:7.70-6.el8.x86_64                         298/654
  Running scriptlet: nmap-ncat-2:7.70-6.el8.x86_64                         298/654
  Running scriptlet: tcpdump-14:4.9.3-2.el8.x86_64                         299/654
  Upgrading        : tcpdump-14:4.9.3-2.el8.x86_64                         299/654
  Upgrading        : fontconfig-2.13.1-4.el8.x86_64                        300/654
  Running scriptlet: fontconfig-2.13.1-4.el8.x86_64                        300/654
  Running scriptlet: man-db-2.7.6.1-18.el8.x86_64                          301/654
  Upgrading        : man-db-2.7.6.1-18.el8.x86_64                          301/654
  Running scriptlet: man-db-2.7.6.1-18.el8.x86_64                          301/654
  Upgrading        : strace-5.7-3.el8.x86_64                               302/654
  Upgrading        : quota-1:4.04-14.el8.x86_64                            303/654
  Upgrading        : python3-lxml-4.2.3-3.el8.x86_64                       304/654
  Upgrading        : python3-psutil-5.4.3-11.el8.x86_64                    305/654
  Upgrading        : lshw-B.02.19.2-6.el8.x86_64                           306/654
  Upgrading        : bpftool-4.18.0-348.7.1.el8_5.x86_64                   307/654
  Upgrading        : libgomp-8.5.0-4.el8_5.x86_64                          308/654
  Running scriptlet: libgomp-8.5.0-4.el8_5.x86_64                          308/654
  Upgrading        : unzip-6.0-45.el8_4.x86_64                             309/654
  Upgrading        : lsscsi-0.32-3.el8.x86_64                              310/654
  Upgrading        : libertas-usb8388-firmware-2:20210702-103.gitd79c267   311/654
  Upgrading        : iwl7260-firmware-1:25.30.13.0-103.el8.1.noarch        312/654
  Upgrading        : iwl6050-firmware-41.28.5.1-103.el8.1.noarch           313/654
  Upgrading        : iwl6000g2b-firmware-18.168.6.1-103.el8.1.noarch       314/654
  Upgrading        : iwl6000g2a-firmware-18.168.6.1-103.el8.1.noarch       315/654
  Upgrading        : iwl6000-firmware-9.221.4.1-103.el8.1.noarch           316/654
  Upgrading        : iwl5150-firmware-8.24.2.2-103.el8.1.noarch            317/654
  Upgrading        : iwl5000-firmware-8.83.5.1_1-103.el8.1.noarch          318/654
  Upgrading        : iwl3160-firmware-1:25.30.13.0-103.el8.1.noarch        319/654
  Upgrading        : iwl2030-firmware-18.168.6.1-103.el8.1.noarch          320/654
  Upgrading        : iwl2000-firmware-18.168.6.1-103.el8.1.noarch          321/654
  Upgrading        : iwl135-firmware-18.168.6.1-103.el8.1.noarch           322/654
  Upgrading        : iwl105-firmware-18.168.6.1-103.el8.1.noarch           323/654
  Upgrading        : iwl1000-firmware-1:39.31.5.1-103.el8.1.noarch         324/654
  Upgrading        : iwl100-firmware-39.31.5.1-103.el8.1.noarch            325/654
  Upgrading        : emacs-filesystem-1:26.1-7.el8.noarch                  326/654
  Upgrading        : alsa-sof-firmware-1.8-1.el8.noarch                    327/654
  Upgrading        : NetworkManager-config-server-1:1.32.10-4.el8.noarch   328/654
  Upgrading        : man-pages-overrides-8.5.0.1-1.el8.noarch              329/654
  Running scriptlet: tuned-2.15.0-2.el8.noarch                             330/654
  Cleanup          : tuned-2.15.0-2.el8.noarch                             330/654
  Running scriptlet: tuned-2.15.0-2.el8.noarch                             330/654
  Running scriptlet: microcode_ctl-4:20210216-1.el8.x86_64                 331/654
  Cleanup          : microcode_ctl-4:20210216-1.el8.x86_64                 331/654
  Running scriptlet: microcode_ctl-4:20210216-1.el8.x86_64                 331/654
  Running scriptlet: firewalld-0.8.2-6.el8.noarch                          332/654
  Cleanup          : firewalld-0.8.2-6.el8.noarch                          332/654
  Running scriptlet: firewalld-0.8.2-6.el8.noarch                          332/654
  Cleanup          : grub2-pc-1:2.02-99.el8.x86_64                         333/654
  Running scriptlet: iptables-ebtables-1.8.4-17.el8.x86_64                 334/654
  Cleanup          : iptables-ebtables-1.8.4-17.el8.x86_64                 334/654
  Running scriptlet: iptables-ebtables-1.8.4-17.el8.x86_64                 334/654
  Cleanup          : kpatch-0.9.2-3.el8.noarch                             335/654
  Cleanup          : cockpit-238.2-1.el8.x86_64                            336/654
  Cleanup          : cockpit-storaged-238.2-1.el8.noarch                   337/654
  Cleanup          : cockpit-system-238.2-1.el8.noarch                     338/654
  Cleanup          : authselect-compat-1.2.2-2.el8.x86_64                  339/654
  Running scriptlet: open-vm-tools-11.2.0-2.el8.x86_64                     340/654
  Cleanup          : open-vm-tools-11.2.0-2.el8.x86_64                     340/654
  Running scriptlet: open-vm-tools-11.2.0-2.el8.x86_64                     340/654
  Running scriptlet: cockpit-ws-238.2-1.el8.x86_64                         341/654
  Cleanup          : cockpit-ws-238.2-1.el8.x86_64                         341/654
  Running scriptlet: cockpit-ws-238.2-1.el8.x86_64                         341/654
  Running scriptlet: openssh-server-8.0p1-5.el8.x86_64                     342/654
  Cleanup          : openssh-server-8.0p1-5.el8.x86_64                     342/654
  Running scriptlet: openssh-server-8.0p1-5.el8.x86_64                     342/654
  Running scriptlet: sssd-kcm-2.4.0-9.el8.x86_64                           343/654
  Cleanup          : sssd-kcm-2.4.0-9.el8.x86_64                           343/654
  Running scriptlet: sssd-kcm-2.4.0-9.el8.x86_64                           343/654
  Cleanup          : NetworkManager-tui-1:1.30.0-7.el8.x86_64              344/654
  Running scriptlet: kexec-tools-2.0.20-46.el8.x86_64                      345/654
  Cleanup          : kexec-tools-2.0.20-46.el8.x86_64                      345/654
  Running scriptlet: kexec-tools-2.0.20-46.el8.x86_64                      345/654
  Obsoleting       : python3-libstoragemgmt-clibs-1.8.7-1.el8.x86_64       346/654
  Running scriptlet: libstoragemgmt-1.8.7-1.el8.x86_64                     347/654
  Cleanup          : libstoragemgmt-1.8.7-1.el8.x86_64                     347/654
  Running scriptlet: libstoragemgmt-1.8.7-1.el8.x86_64                     347/654
  Cleanup          : python3-libstoragemgmt-1.8.7-1.el8.noarch             348/654
  Running scriptlet: device-mapper-multipath-0.8.4-10.el8.x86_64           349/654
  Cleanup          : device-mapper-multipath-0.8.4-10.el8.x86_64           349/654
  Running scriptlet: device-mapper-multipath-0.8.4-10.el8.x86_64           349/654
  Running scriptlet: binutils-2.30-93.el8.x86_64                           350/654
  Cleanup          : binutils-2.30-93.el8.x86_64                           350/654
  Running scriptlet: binutils-2.30-93.el8.x86_64                           350/654
  Running scriptlet: vdo-6.2.4.14-14.el8.x86_64                            351/654
  Cleanup          : vdo-6.2.4.14-14.el8.x86_64                            351/654
  Running scriptlet: vdo-6.2.4.14-14.el8.x86_64                            351/654
  Cleanup          : python3-lxml-4.2.3-2.el8.x86_64                       352/654
  Cleanup          : device-mapper-multipath-libs-0.8.4-10.el8.x86_64      353/654
  Running scriptlet: device-mapper-multipath-libs-0.8.4-10.el8.x86_64      353/654
  Cleanup          : sudo-1.8.29-7.el8.x86_64                              354/654
  Running scriptlet: chrony-3.5-2.el8.x86_64                               355/654
  Cleanup          : chrony-3.5-2.el8.x86_64                               355/654
  Running scriptlet: chrony-3.5-2.el8.x86_64                               355/654
  Cleanup          : openssh-clients-8.0p1-5.el8.x86_64                    356/654
  Cleanup          : realmd-0.16.3-22.el8.x86_64                           357/654
  Cleanup          : kernel-tools-4.18.0-305.3.1.el8.x86_64                358/654
  Cleanup          : buildah-1.19.7-1.module_el8.4.0+781+acf4c33b.x86_64   359/654
  Cleanup          : dracut-network-049-135.git20210121.el8.x86_64         360/654
  Cleanup          : dhcp-client-12:4.3.6-44.0.1.el8.x86_64                361/654
  Cleanup          : dhcp-libs-12:4.3.6-44.0.1.el8.x86_64                  362/654
  Running scriptlet: kmod-kvdo-6.2.4.26-77.el8.x86_64                      363/654
  Cleanup          : kmod-kvdo-6.2.4.26-77.el8.x86_64                      363/654
  Running scriptlet: kmod-kvdo-6.2.4.26-77.el8.x86_64                      363/654
  Cleanup          : kpatch-dnf-0.2-3.el8.noarch                           364/654
  Cleanup          : centos-logos-85.5-1.el8.x86_64                        365/654
  Running scriptlet: centos-logos-85.5-1.el8.x86_64                        365/654
  Cleanup          : sos-4.0-11.el8.noarch                                 366/654
  Cleanup          : dracut-squash-049-135.git20210121.el8.x86_64          367/654
  Cleanup          : python3-firewall-0.8.2-6.el8.noarch                   368/654
  Cleanup          : python3-nftables-1:0.9.3-18.el8.x86_64                369/654
  Cleanup          : python3-syspurpose-1.28.13-2.el8.x86_64               370/654
  Cleanup          : dracut-config-rescue-049-135.git20210121.el8.x86_64   371/654
  Cleanup          : cockpit-podman-29-2.module_el8.4.0+781+acf4c33b.noa   372/654
  Cleanup          : podman-3.0.1-6.module_el8.4.0+781+acf4c33b.x86_64     373/654
  Cleanup          : bind-export-libs-32:9.11.26-3.el8.x86_64              374/654
  Running scriptlet: bind-export-libs-32:9.11.26-3.el8.x86_64              374/654
  Cleanup          : openssh-8.0p1-5.el8.x86_64                            375/654
  Cleanup          : openssl-1:1.1.1g-15.el8_3.x86_64                      376/654
  Cleanup          : lshw-B.02.19.2-5.el8.x86_64                           377/654
  Cleanup          : vim-enhanced-2:8.0.1763-15.el8.x86_64                 378/654
  Cleanup          : nmap-ncat-2:7.70-5.el8.x86_64                         379/654
  Running scriptlet: nmap-ncat-2:7.70-5.el8.x86_64                         379/654
  Cleanup          : iproute-5.9.0-4.el8.x86_64                            380/654
  Cleanup          : bpftool-4.18.0-305.3.1.el8.x86_64                     381/654
  Cleanup          : iptables-1.8.4-17.el8.x86_64                          382/654
  Running scriptlet: iptables-1.8.4-17.el8.x86_64                          382/654
  Cleanup          : libibverbs-32.0-4.el8.x86_64                          383/654
  Running scriptlet: libibverbs-32.0-4.el8.x86_64                          383/654
  Cleanup          : fontconfig-2.13.1-3.el8.x86_64                        384/654
  Cleanup          : conmon-2:2.0.26-1.module_el8.4.0+781+acf4c33b.x86_6   385/654
  Cleanup          : runc-1.0.0-70.rc92.module_el8.4.0+673+eabfc99d.x86_   386/654
  Cleanup          : criu-3.15-1.module_el8.4.0+641+6116a774.x86_64        387/654
  Cleanup          : python3-perf-4.18.0-305.3.1.el8.x86_64                388/654
  Cleanup          : strace-5.7-2.el8.x86_64                               389/654
  Running scriptlet: libgomp-8.4.1-1.el8.x86_64                            390/654
  Cleanup          : libgomp-8.4.1-1.el8.x86_64                            390/654
  Running scriptlet: libgomp-8.4.1-1.el8.x86_64                            390/654
  Cleanup          : plymouth-0.9.4-9.20200615git1e36e30.el8.x86_64        391/654
  Running scriptlet: plymouth-0.9.4-9.20200615git1e36e30.el8.x86_64        391/654
  Cleanup          : plymouth-core-libs-0.9.4-9.20200615git1e36e30.el8.x   392/654
  Running scriptlet: plymouth-core-libs-0.9.4-9.20200615git1e36e30.el8.x   392/654
  Running scriptlet: nftables-1:0.9.3-18.el8.x86_64                        393/654
  Cleanup          : nftables-1:0.9.3-18.el8.x86_64                        393/654
  Running scriptlet: nftables-1:0.9.3-18.el8.x86_64                        393/654
  Cleanup          : util-linux-user-2.32.1-27.el8.x86_64                  394/654
  Running scriptlet: mcelog-3:173-0.el8.x86_64                             395/654
  Cleanup          : mcelog-3:173-0.el8.x86_64                             395/654
  Running scriptlet: mcelog-3:173-0.el8.x86_64                             395/654
  Running scriptlet: authselect-1.2.2-2.el8.x86_64                         396/654
  Cleanup          : authselect-1.2.2-2.el8.x86_64                         396/654
  Cleanup          : tpm2-tools-4.1.1-2.el8.x86_64                         397/654
  Cleanup          : tpm2-tss-2.3.2-3.el8.x86_64                           398/654
  Running scriptlet: tpm2-tss-2.3.2-3.el8.x86_64                           398/654
  Cleanup          : man-db-2.7.6.1-17.el8.x86_64                          399/654
  Cleanup          : NetworkManager-team-1:1.30.0-7.el8.x86_64             400/654
  Running scriptlet: NetworkManager-1:1.30.0-7.el8.x86_64                  401/654
  Cleanup          : NetworkManager-1:1.30.0-7.el8.x86_64                  401/654
  Running scriptlet: NetworkManager-1:1.30.0-7.el8.x86_64                  401/654
  Cleanup          : NetworkManager-libnm-1:1.30.0-7.el8.x86_64            402/654
  Running scriptlet: NetworkManager-libnm-1:1.30.0-7.el8.x86_64            402/654
  Cleanup          : grub2-tools-extra-1:2.02-99.el8.x86_64                403/654
  Cleanup          : authselect-libs-1.2.2-2.el8.x86_64                    404/654
  Cleanup          : rdma-core-32.0-4.el8.x86_64                           405/654
  Cleanup          : udisks2-lvm2-2.9.0-6.el8.x86_64                       406/654
  Cleanup          : libblockdev-lvm-2.24-5.el8.x86_64                     407/654
  Running scriptlet: lvm2-8:2.03.11-5.el8.x86_64                           408/654
  Cleanup          : lvm2-8:2.03.11-5.el8.x86_64                           408/654
  Running scriptlet: lvm2-8:2.03.11-5.el8.x86_64                           408/654
  Cleanup          : lvm2-libs-8:2.03.11-5.el8.x86_64                      409/654
  Running scriptlet: device-mapper-event-8:1.02.175-5.el8.x86_64           410/654
  Cleanup          : device-mapper-event-8:1.02.175-5.el8.x86_64           410/654
  Cleanup          : device-mapper-persistent-data-0.8.5-4.el8.x86_64      411/654
  Cleanup          : device-mapper-event-libs-8:1.02.175-5.el8.x86_64      412/654
  Cleanup          : elfutils-debuginfod-client-0.182-3.el8.x86_64         413/654
  Cleanup          : python3-psutil-5.4.3-10.el8.x86_64                    414/654
  Cleanup          : libX11-1.6.8-4.el8.x86_64                             415/654
  Cleanup          : container-selinux-2:2.158.0-1.module_el8.4.0+781+ac   416/654
  Running scriptlet: container-selinux-2:2.158.0-1.module_el8.4.0+781+ac   416/654
  Cleanup          : sssd-2.4.0-9.el8.x86_64                               417/654
  Cleanup          : rpm-plugin-selinux-4.14.3-13.el8.x86_64               418/654
  Cleanup          : selinux-policy-targeted-3.14.3-67.el8.noarch          419/654
  Running scriptlet: selinux-policy-targeted-3.14.3-67.el8.noarch          419/654
  Cleanup          : selinux-policy-3.14.3-67.el8.noarch                   420/654
  Running scriptlet: selinux-policy-3.14.3-67.el8.noarch                   420/654
  Cleanup          : plymouth-scripts-0.9.4-9.20200615git1e36e30.el8.x86   421/654
  Cleanup          : containers-common-1:1.2.2-8.module_el8.4.0+781+acf4   422/654
  Cleanup          : python3-sssdconfig-2.4.0-9.el8.noarch                 423/654
  Cleanup          : cockpit-packagekit-238.2-1.el8.noarch                 424/654
  Cleanup          : grub2-pc-modules-1:2.02-99.el8.noarch                 425/654
  Cleanup          : yum-4.4.2-11.el8.noarch                               426/654
  Running scriptlet: dnf-4.4.2-11.el8.noarch                               427/654
  Cleanup          : dnf-4.4.2-11.el8.noarch                               427/654
  Running scriptlet: dnf-4.4.2-11.el8.noarch                               427/654
  Cleanup          : gsettings-desktop-schemas-3.32.0-5.el8.x86_64         428/654
  Cleanup          : dnf-plugins-core-4.0.18-4.el8.noarch                  429/654
  Cleanup          : python3-dnf-plugins-core-4.0.18-4.el8.noarch          430/654
  Cleanup          : python3-dnf-4.4.2-11.el8.noarch                       431/654
  Cleanup          : centos-linux-release-8.4-1.2105.el8.noarch            432/654
  Cleanup          : centos-linux-repos-8-2.el8.noarch                     433/654
  Cleanup          : setroubleshoot-plugins-3.3.13-1.el8.noarch            434/654
  Cleanup          : centos-gpg-keys-1:8-2.el8.noarch                      435/654
  Cleanup          : dnf-data-4.4.2-11.el8.noarch                          436/654
  Cleanup          : libX11-common-1.6.8-4.el8.noarch                      437/654
  Cleanup          : hwdata-0.314-8.8.el8.noarch                           438/654
  Cleanup          : podman-catatonit-3.0.1-6.module_el8.4.0+781+acf4c33   439/654
  Cleanup          : dhcp-common-12:4.3.6-44.0.1.el8.noarch                440/654
  Cleanup          : firewalld-filesystem-0.8.2-6.el8.noarch               441/654
  Cleanup          : linux-firmware-20201218-102.git05789708.el8.noarch    442/654
  Cleanup          : libertas-usb8388-firmware-2:20201218-102.git0578970   443/654
  Cleanup          : iwl7260-firmware-1:25.30.13.0-102.el8.1.noarch        444/654
  Cleanup          : iwl6050-firmware-41.28.5.1-102.el8.1.noarch           445/654
  Cleanup          : iwl6000g2b-firmware-18.168.6.1-102.el8.1.noarch       446/654
  Cleanup          : iwl6000g2a-firmware-18.168.6.1-102.el8.1.noarch       447/654
  Cleanup          : iwl6000-firmware-9.221.4.1-102.el8.1.noarch           448/654
  Cleanup          : iwl5150-firmware-8.24.2.2-102.el8.1.noarch            449/654
  Cleanup          : iwl5000-firmware-8.83.5.1_1-102.el8.1.noarch          450/654
  Cleanup          : iwl3160-firmware-1:25.30.13.0-102.el8.1.noarch        451/654
  Cleanup          : iwl2030-firmware-18.168.6.1-102.el8.1.noarch          452/654
  Cleanup          : iwl2000-firmware-18.168.6.1-102.el8.1.noarch          453/654
  Cleanup          : iwl135-firmware-18.168.6.1-102.el8.1.noarch           454/654
  Cleanup          : iwl105-firmware-18.168.6.1-102.el8.1.noarch           455/654
  Cleanup          : iwl1000-firmware-1:39.31.5.1-102.el8.1.noarch         456/654
  Cleanup          : iwl100-firmware-39.31.5.1-102.el8.1.noarch            457/654
  Cleanup          : emacs-filesystem-1:26.1-5.el8.noarch                  458/654
  Cleanup          : alsa-sof-firmware-1.6.1-2.el8.noarch                  459/654
  Cleanup          : NetworkManager-config-server-1:1.30.0-7.el8.noarch    460/654
  Cleanup          : man-pages-overrides-8.3.0.2-2.el8.noarch              461/654
  Cleanup          : sssd-ipa-2.4.0-9.el8.x86_64                           462/654
  Cleanup          : sssd-ad-2.4.0-9.el8.x86_64                            463/654
  Cleanup          : libsmbclient-4.13.3-3.el8.x86_64                      464/654
  Cleanup          : sssd-common-pac-2.4.0-9.el8.x86_64                    465/654
  Running scriptlet: libwbclient-4.13.3-3.el8.x86_64                       466/654
  Cleanup          : libwbclient-4.13.3-3.el8.x86_64                       466/654
  Cleanup          : samba-client-libs-4.13.3-3.el8.x86_64                 467/654
  Cleanup          : samba-common-libs-4.13.3-3.el8.x86_64                 468/654
  Cleanup          : python3-hawkey-0.55.0-7.el8.x86_64                    469/654
  Cleanup          : python3-libdnf-0.55.0-7.el8.x86_64                    470/654
  Cleanup          : libdnf-0.55.0-7.el8.x86_64                            471/654
  Cleanup          : sssd-ldap-2.4.0-9.el8.x86_64                          472/654
  Cleanup          : sssd-proxy-2.4.0-9.el8.x86_64                         473/654
  Cleanup          : bind-utils-32:9.11.26-3.el8.x86_64                    474/654
  Cleanup          : cockpit-bridge-238.2-1.el8.x86_64                     475/654
  Cleanup          : cups-libs-1:2.2.6-38.el8.x86_64                       476/654
  Cleanup          : sssd-krb5-2.4.0-9.el8.x86_64                          477/654
  Cleanup          : bind-libs-32:9.11.26-3.el8.x86_64                     478/654
  Cleanup          : bind-libs-lite-32:9.11.26-3.el8.x86_64                479/654
  Cleanup          : adcli-0.8.2-9.el8.x86_64                              480/654
  Running scriptlet: adcli-0.8.2-9.el8.x86_64                              480/654
  Cleanup          : sssd-krb5-common-2.4.0-9.el8.x86_64                   481/654
  Running scriptlet: sssd-common-2.4.0-9.el8.x86_64                        482/654
  Cleanup          : sssd-common-2.4.0-9.el8.x86_64                        482/654
  Running scriptlet: sssd-common-2.4.0-9.el8.x86_64                        482/654
  Running scriptlet: sssd-client-2.4.0-9.el8.x86_64                        483/654
  Cleanup          : sssd-client-2.4.0-9.el8.x86_64                        483/654
  Running scriptlet: sssd-client-2.4.0-9.el8.x86_64                        483/654
  Cleanup          : librepo-1.12.0-3.el8.x86_64                           484/654
  Cleanup          : setroubleshoot-server-3.3.24-3.el8.x86_64             485/654
  Running scriptlet: setroubleshoot-server-3.3.24-3.el8.x86_64             485/654
  Cleanup          : python3-libxml2-2.9.7-9.el8.x86_64                    486/654
  Running scriptlet: polkit-0.115-11.el8.x86_64                            487/654
  Cleanup          : polkit-0.115-11.el8.x86_64                            487/654
  Running scriptlet: polkit-0.115-11.el8.x86_64                            487/654
  Cleanup          : python3-rpm-4.14.3-13.el8.x86_64                      488/654
  Cleanup          : rpm-build-libs-4.14.3-13.el8.x86_64                   489/654
  Running scriptlet: rpm-build-libs-4.14.3-13.el8.x86_64                   489/654
  Cleanup          : libstdc++-8.4.1-1.el8.x86_64                          490/654
  Running scriptlet: libstdc++-8.4.1-1.el8.x86_64                          490/654
  Cleanup          : fuse-overlayfs-1.4.0-2.module_el8.4.0+673+eabfc99d.   491/654
  Cleanup          : rpm-plugin-systemd-inhibit-4.14.3-13.el8.x86_64       492/654
  Cleanup          : libldb-2.2.0-2.el8.x86_64                             493/654
  Cleanup          : sqlite-3.26.0-13.el8.x86_64                           494/654
  Cleanup          : libsss_nss_idmap-2.4.0-9.el8.x86_64                   495/654
  Running scriptlet: libsss_nss_idmap-2.4.0-9.el8.x86_64                   495/654
  Cleanup          : libsolv-0.7.16-2.el8.x86_64                           496/654
  Cleanup          : libsss_certmap-2.4.0-9.el8.x86_64                     497/654
  Running scriptlet: libsss_certmap-2.4.0-9.el8.x86_64                     497/654
  Cleanup          : slirp4netns-1.1.8-1.module_el8.4.0+641+6116a774.x86   498/654
  Cleanup          : libtevent-0.10.2-2.el8.x86_64                         499/654
  Cleanup          : python3-gpg-1.13.1-7.el8.x86_64                       500/654
  Cleanup          : python3-unbound-1.7.3-15.el8.x86_64                   501/654
  Running scriptlet: unbound-libs-1.7.3-15.el8.x86_64                      502/654
  Cleanup          : unbound-libs-1.7.3-15.el8.x86_64                      502/654
  Running scriptlet: unbound-libs-1.7.3-15.el8.x86_64                      502/654
  Cleanup          : libmodulemd-2.9.4-2.el8.x86_64                        503/654
  Cleanup          : gpgme-1.13.1-7.el8.x86_64                             504/654
  Cleanup          : python3-libcomps-0.1.11-5.el8.x86_64                  505/654
  Cleanup          : libcomps-0.1.11-5.el8.x86_64                          506/654
  Cleanup          : libxml2-2.9.7-9.el8.x86_64                            507/654
  Cleanup          : numactl-libs-2.0.12-11.el8.x86_64                     508/654
  Running scriptlet: numactl-libs-2.0.12-11.el8.x86_64                     508/654
  Cleanup          : libdrm-2.4.103-1.el8.x86_64                           509/654
  Cleanup          : udisks2-iscsi-2.9.0-6.el8.x86_64                      510/654
  Running scriptlet: udisks2-2.9.0-6.el8.x86_64                            511/654
  Cleanup          : udisks2-2.9.0-6.el8.x86_64                            511/654
  Running scriptlet: udisks2-2.9.0-6.el8.x86_64                            511/654
  Running scriptlet: iscsi-initiator-utils-6.2.1.2-1.gita8fcb37.el8.x86_   512/654
  Cleanup          : iscsi-initiator-utils-6.2.1.2-1.gita8fcb37.el8.x86_   512/654
  Running scriptlet: iscsi-initiator-utils-6.2.1.2-1.gita8fcb37.el8.x86_   512/654
  Cleanup          : libblockdev-crypto-2.24-5.el8.x86_64                  513/654
  Cleanup          : nss-3.53.1-17.el8_3.x86_64                            514/654
  Cleanup          : libblockdev-fs-2.24-5.el8.x86_64                      515/654
  Cleanup          : nss-softokn-3.53.1-17.el8_3.x86_64                    516/654
  Cleanup          : e2fsprogs-1.45.6-1.el8.x86_64                         517/654
  Cleanup          : xfsprogs-5.0.0-8.el8.x86_64                           518/654
  Running scriptlet: xfsprogs-5.0.0-8.el8.x86_64                           518/654
  Running scriptlet: iscsi-initiator-utils-iscsiuio-6.2.1.2-1.gita8fcb37   519/654
  Cleanup          : iscsi-initiator-utils-iscsiuio-6.2.1.2-1.gita8fcb37   519/654
  Running scriptlet: iscsi-initiator-utils-iscsiuio-6.2.1.2-1.gita8fcb37   519/654
  Cleanup          : libblockdev-part-2.24-5.el8.x86_64                    520/654
  Cleanup          : libblockdev-2.24-5.el8.x86_64                         521/654
  Cleanup          : nss-sysinit-3.53.1-17.el8_3.x86_64                    522/654
  Cleanup          : libblockdev-swap-2.24-5.el8.x86_64                    523/654
  Cleanup          : polkit-libs-0.115-11.el8.x86_64                       524/654
  Running scriptlet: polkit-libs-0.115-11.el8.x86_64                       524/654
  Cleanup          : libss-1.45.6-1.el8.x86_64                             525/654
  Running scriptlet: libss-1.45.6-1.el8.x86_64                             525/654
  Cleanup          : libblockdev-loop-2.24-5.el8.x86_64                    526/654
  Cleanup          : libblockdev-mdraid-2.24-5.el8.x86_64                  527/654
  Running scriptlet: mdadm-4.1-15.el8.x86_64                               528/654
  Cleanup          : mdadm-4.1-15.el8.x86_64                               528/654
  Running scriptlet: mdadm-4.1-15.el8.x86_64                               528/654
  Cleanup          : libblockdev-utils-2.24-5.el8.x86_64                   529/654
  Running scriptlet: parted-3.2-38.el8.x86_64                              530/654
  Cleanup          : parted-3.2-38.el8.x86_64                              530/654
  Running scriptlet: parted-3.2-38.el8.x86_64                              530/654
  Cleanup          : nss-softokn-freebl-3.53.1-17.el8_3.x86_64             531/654
  Cleanup          : nss-util-3.53.1-17.el8_3.x86_64                       532/654
  Cleanup          : nspr-4.25.0-2.el8_2.x86_64                            533/654
  Running scriptlet: nspr-4.25.0-2.el8_2.x86_64                            533/654
  Cleanup          : quota-1:4.04-12.el8.x86_64                            534/654
  Cleanup          : e2fsprogs-libs-1.45.6-1.el8.x86_64                    535/654
  Running scriptlet: e2fsprogs-libs-1.45.6-1.el8.x86_64                    535/654
  Cleanup          : nvme-cli-1.12-3.el8.0.1.x86_64                        536/654
  Cleanup          : tcpdump-14:4.9.3-1.el8.x86_64                         537/654
  Cleanup          : libsss_idmap-2.4.0-9.el8.x86_64                       538/654
  Running scriptlet: libsss_idmap-2.4.0-9.el8.x86_64                       538/654
  Cleanup          : libsss_sudo-2.4.0-9.el8.x86_64                        539/654
  Running scriptlet: libsss_sudo-2.4.0-9.el8.x86_64                        539/654
  Cleanup          : vim-minimal-2:8.0.1763-15.el8.x86_64                  540/654
  Cleanup          : rsyslog-gssapi-8.1911.0-7.el8.x86_64                  541/654
  Cleanup          : sssd-nfs-idmap-2.4.0-9.el8.x86_64                     542/654
  Cleanup          : libnfsidmap-1:2.3.3-41.el8.x86_64                     543/654
  Cleanup          : fstrm-0.6.0-3.el8.1.x86_64                            544/654
  Cleanup          : json-c-0.13.1-0.4.el8.x86_64                          545/654
  Cleanup          : containernetworking-plugins-0.9.1-1.module_el8.4.0+   546/654
  Cleanup          : ethtool-2:5.8-5.el8.x86_64                            547/654
  Cleanup          : rsyslog-relp-8.1911.0-7.el8.x86_64                    548/654
  Cleanup          : librelp-1.2.16-1.el8.x86_64                           549/654
  Running scriptlet: librelp-1.2.16-1.el8.x86_64                           549/654
  Cleanup          : libslirp-4.3.1-1.module_el8.4.0+575+63b40ad7.x86_64   550/654
  Cleanup          : libipa_hbac-2.4.0-9.el8.x86_64                        551/654
  Running scriptlet: libipa_hbac-2.4.0-9.el8.x86_64                        551/654
  Cleanup          : rsyslog-gnutls-8.1911.0-7.el8.x86_64                  552/654
  Running scriptlet: rsyslog-8.1911.0-7.el8.x86_64                         553/654
  Cleanup          : rsyslog-8.1911.0-7.el8.x86_64                         553/654
  Running scriptlet: rsyslog-8.1911.0-7.el8.x86_64                         553/654
  Cleanup          : libfastjson-0.99.8-2.el8.x86_64                       554/654
  Running scriptlet: libfastjson-0.99.8-2.el8.x86_64                       554/654
  Cleanup          : libudisks2-2.9.0-6.el8.x86_64                         555/654
  Cleanup          : glib2-2.56.4-9.el8.x86_64                             556/654
  Cleanup          : pcre-8.42-4.el8.x86_64                                557/654
  Cleanup          : libtalloc-2.3.1-2.el8.x86_64                          558/654
  Cleanup          : libndp-1.7-5.el8.x86_64                               559/654
  Running scriptlet: libndp-1.7-5.el8.x86_64                               559/654
  Cleanup          : vim-common-2:8.0.1763-15.el8.x86_64                   560/654
  Cleanup          : virt-what-1.18-6.el8.x86_64                           561/654
  Cleanup          : hdparm-9.54-3.el8.x86_64                              562/654
  Cleanup          : unzip-6.0-44.el8.x86_64                               563/654
  Cleanup          : libsss_autofs-2.4.0-9.el8.x86_64                      564/654
  Cleanup          : kernel-tools-libs-4.18.0-305.3.1.el8.x86_64           565/654
  Running scriptlet: kernel-tools-libs-4.18.0-305.3.1.el8.x86_64           565/654
  Cleanup          : lsscsi-0.32-2.el8.x86_64                              566/654
  Cleanup          : python3-bind-32:9.11.26-3.el8.noarch                  567/654
  Cleanup          : samba-common-4.13.3-3.el8.noarch                      568/654
  Cleanup          : policycoreutils-python-utils-2.9-14.el8.noarch        569/654
  Cleanup          : python3-policycoreutils-2.9-14.el8.noarch             570/654
  Cleanup          : bind-license-32:9.11.26-3.el8.noarch                  571/654
  Cleanup          : vim-filesystem-2:8.0.1763-15.el8.noarch               572/654
  Cleanup          : quota-nls-1:4.04-12.el8.noarch                        573/654
  Running scriptlet: policycoreutils-2.9-14.el8.x86_64                     574/654
  Cleanup          : policycoreutils-2.9-14.el8.x86_64                     574/654
  Cleanup          : platform-python-pip-9.0.3-19.el8.noarch               575/654
  Cleanup          : dbus-tools-1:1.12.8-12.el8.x86_64                     576/654
  Cleanup          : dbus-libs-1:1.12.8-12.el8.x86_64                      577/654
  Running scriptlet: dbus-libs-1:1.12.8-12.el8.x86_64                      577/654
  Running scriptlet: dbus-daemon-1:1.12.8-12.el8.x86_64                    578/654
  Cleanup          : dbus-daemon-1:1.12.8-12.el8.x86_64                    578/654
  Running scriptlet: dbus-daemon-1:1.12.8-12.el8.x86_64                    578/654
  Cleanup          : systemd-pam-239-45.el8.x86_64                         579/654
  Cleanup          : libfdisk-2.32.1-27.el8.x86_64                         580/654
  Running scriptlet: libfdisk-2.32.1-27.el8.x86_64                         580/654
  Cleanup          : libmount-2.32.1-27.el8.x86_64                         581/654
  Running scriptlet: libmount-2.32.1-27.el8.x86_64                         581/654
  Cleanup          : ca-certificates-2020.2.41-80.0.el8_2.noarch           582/654
  Cleanup          : shadow-utils-2:4.6-12.el8.x86_64                      583/654
  Cleanup          : libblkid-2.32.1-27.el8.x86_64                         584/654
  Running scriptlet: libblkid-2.32.1-27.el8.x86_64                         584/654
  Cleanup          : systemd-libs-239-45.el8.x86_64                        585/654
  Cleanup          : pam-1.3.1-14.el8.x86_64                               586/654
  Running scriptlet: pam-1.3.1-14.el8.x86_64                               586/654
  Cleanup          : util-linux-2.32.1-27.el8.x86_64                       587/654
  Cleanup          : python3-libs-3.6.8-37.el8.x86_64                      588/654
  Cleanup          : libtirpc-1.1.4-4.el8.x86_64                           589/654
  Running scriptlet: libtirpc-1.1.4-4.el8.x86_64                           589/654
  Cleanup          : platform-python-3.6.8-37.el8.x86_64                   590/654
  Running scriptlet: platform-python-3.6.8-37.el8.x86_64                   590/654
  Cleanup          : kmod-25-17.el8.x86_64                                 591/654
  Cleanup          : kmod-libs-25-17.el8.x86_64                            592/654
  Running scriptlet: kmod-libs-25-17.el8.x86_64                            592/654
  Cleanup          : libdb-utils-5.3.28-40.el8.x86_64                      593/654
  Cleanup          : libcurl-7.61.1-18.el8.x86_64                          594/654
  Cleanup          : libssh-0.9.4-2.el8.x86_64                             595/654
  Cleanup          : krb5-libs-1.18.2-8.el8.x86_64                         596/654
  Cleanup          : openldap-2.4.46-16.el8.x86_64                         597/654
  Cleanup          : curl-7.61.1-18.el8.x86_64                             598/654
  Cleanup          : rpm-4.14.3-13.el8.x86_64                              599/654
  Cleanup          : rpm-libs-4.14.3-13.el8.x86_64                         600/654
  Running scriptlet: rpm-libs-4.14.3-13.el8.x86_64                         600/654
  Cleanup          : libdb-5.3.28-40.el8.x86_64                            601/654
  Running scriptlet: libdb-5.3.28-40.el8.x86_64                            601/654
  Cleanup          : coreutils-8.30-8.el8.x86_64                           602/654
  Cleanup          : openssl-libs-1:1.1.1g-15.el8_3.x86_64                 603/654
  Running scriptlet: openssl-libs-1:1.1.1g-15.el8_3.x86_64                 603/654
  Cleanup          : gnutls-3.6.14-7.el8_3.x86_64                          604/654
  Cleanup          : crypto-policies-20210209-1.gitbfb6bed.el8_3.noarch    605/654
  Cleanup          : crypto-policies-scripts-20210209-1.gitbfb6bed.el8_3   606/654
  Cleanup          : grubby-8.40-41.el8.x86_64                             607/654
  Running scriptlet: grub2-tools-1:2.02-99.el8.x86_64                      608/654
  Cleanup          : grub2-tools-1:2.02-99.el8.x86_64                      608/654
  Cleanup          : dracut-049-135.git20210121.el8.x86_64                 609/654
  Cleanup          : kpartx-0.8.4-10.el8.x86_64                            610/654
  Cleanup          : os-prober-1.74-6.el8.x86_64                           611/654
  Cleanup          : systemd-udev-239-45.el8.x86_64                        612/654
  Running scriptlet: systemd-udev-239-45.el8.x86_64                        612/654
  Cleanup          : grub2-tools-minimal-1:2.02-99.el8.x86_64              613/654
  Cleanup          : elfutils-libs-0.182-3.el8.x86_64                      614/654
  Cleanup          : elfutils-default-yama-scope-0.182-3.el8.noarch        615/654
  Cleanup          : device-mapper-libs-8:1.02.175-5.el8.x86_64            616/654
  Cleanup          : device-mapper-8:1.02.175-5.el8.x86_64                 617/654
  Cleanup          : dbus-1:1.12.8-12.el8.x86_64                           618/654
  Running scriptlet: systemd-239-45.el8.x86_64                             619/654
  Cleanup          : systemd-239-45.el8.x86_64                             619/654
  Cleanup          : libsmartcols-2.32.1-27.el8.x86_64                     620/654
  Running scriptlet: libsmartcols-2.32.1-27.el8.x86_64                     620/654
  Cleanup          : sqlite-libs-3.26.0-13.el8.x86_64                      621/654
  Cleanup          : libuuid-2.32.1-27.el8.x86_64                          622/654
  Running scriptlet: libuuid-2.32.1-27.el8.x86_64                          622/654
  Cleanup          : iptables-libs-1.8.4-17.el8.x86_64                     623/654
  Cleanup          : lua-libs-5.3.4-11.el8.x86_64                          624/654
  Cleanup          : libcom_err-1.45.6-1.el8.x86_64                        625/654
  Running scriptlet: libcom_err-1.45.6-1.el8.x86_64                        625/654
  Cleanup          : libgcrypt-1.8.5-4.el8.x86_64                          626/654
  Running scriptlet: libgcrypt-1.8.5-4.el8.x86_64                          626/654
  Running scriptlet: nettle-3.4.1-2.el8.x86_64                             627/654
  Cleanup          : nettle-3.4.1-2.el8.x86_64                             627/654
  Running scriptlet: nettle-3.4.1-2.el8.x86_64                             627/654
  Cleanup          : ncurses-6.1-7.20180224.el8.x86_64                     628/654
  Cleanup          : chkconfig-1.13-2.el8.x86_64                           629/654
  Cleanup          : libsepol-2.9-2.el8.x86_64                             630/654
  Running scriptlet: libsepol-2.9-2.el8.x86_64                             630/654
  Cleanup          : libcap-ng-0.7.9-5.el8.x86_64                          631/654
  Cleanup          : libcap-2.26-4.el8.x86_64                              632/654
  Cleanup          : libxcrypt-4.1.1-4.el8.x86_64                          633/654
  Cleanup          : elfutils-libelf-0.182-3.el8.x86_64                    634/654
  Cleanup          : keyutils-libs-1.5.10-6.el8.x86_64                     635/654
  Cleanup          : file-5.33-16.el8_3.1.x86_64                           636/654
  Cleanup          : file-libs-5.33-16.el8_3.1.x86_64                      637/654
  Cleanup          : which-2.21-12.el8.x86_64                              638/654
  Cleanup          : dmidecode-1:3.2-8.el8.x86_64                          639/654
  Running scriptlet: coreutils-common-8.30-8.el8.x86_64                    640/654
  Cleanup          : coreutils-common-8.30-8.el8.x86_64                    640/654
  Cleanup          : grub2-common-1:2.02-99.el8.noarch                     641/654
  Cleanup          : libssh-config-0.9.4-2.el8.noarch                      642/654
  Cleanup          : python3-pip-wheel-9.0.3-19.el8.noarch                 643/654
  Cleanup          : dbus-common-1:1.12.8-12.el8.noarch                    644/654
  Cleanup          : lz4-libs-1.8.3-2.el8.x86_64                           645/654
  Cleanup          : ncurses-libs-6.1-7.20180224.el8.x86_64                646/654
  Cleanup          : bash-4.4.19-14.el8.x86_64                             647/654
  Running scriptlet: bash-4.4.19-14.el8.x86_64                             647/654
  Cleanup          : glibc-2.28-151.el8.x86_64                             648/654
  Cleanup          : glibc-langpack-en-2.28-151.el8.x86_64                 649/654
  Cleanup          : glibc-common-2.28-151.el8.x86_64                      650/654
  Cleanup          : tzdata-2021a-1.el8.noarch                             651/654
  Cleanup          : filesystem-3.8-3.el8.x86_64                           652/654
  Cleanup          : ncurses-base-6.1-7.20180224.el8.noarch                653/654
  Cleanup          : libgcc-8.4.1-1.el8.x86_64                             654/654
  Running scriptlet: libgcc-8.4.1-1.el8.x86_64                             654/654
  Running scriptlet: filesystem-3.8-6.el8.x86_64                           654/654
  Running scriptlet: crypto-policies-scripts-20210617-1.gitc776d3e.el8.n   654/654
  Running scriptlet: ca-certificates-2021.2.50-80.0.el8_4.noarch           654/654
  Running scriptlet: kernel-core-4.18.0-348.7.1.el8_5.x86_64               654/654
  Running scriptlet: kmod-kvdo-6.2.5.72-81.el8.x86_64                      654/654
  Running scriptlet: container-selinux-2:2.167.0-1.module_el8.5.0+911+f1   654/654

  Running scriptlet: authselect-libs-1.2.2-3.el8.x86_64                    654/654
  Running scriptlet: nss-3.67.0-7.el8_5.x86_64                             654/654
  Running scriptlet: centos-logos-85.8-2.el8.x86_64                        654/654
  Running scriptlet: sssd-common-2.5.2-2.el8_5.3.x86_64                    654/654
  Running scriptlet: libwbclient-4.14.5-7.el8_5.x86_64                     654/654
  Running scriptlet: tuned-2.16.0-1.el8.noarch                             654/654
  Running scriptlet: authselect-compat-1.2.2-3.el8.x86_64                  654/654
  Running scriptlet: microcode_ctl-4:20210608-1.el8.x86_64                 654/654
  Running scriptlet: libgcc-8.4.1-1.el8.x86_64                             654/654
  Running scriptlet: glibc-common-2.28-164.el8.x86_64                      654/654
  Running scriptlet: systemd-239-51.el8_5.2.x86_64                         654/654
  Running scriptlet: systemd-udev-239-51.el8_5.2.x86_64                    654/654
  Running scriptlet: glib2-2.56.4-156.el8.x86_64                           654/654
  Running scriptlet: vim-common-2:8.0.1763-16.el8.x86_64                   654/654
  Running scriptlet: fontconfig-2.13.1-4.el8.x86_64                        654/654
  Running scriptlet: man-db-2.7.6.1-18.el8.x86_64                          654/654
  Verifying        : grub2-tools-efi-1:2.02-106.el8.x86_64                   1/654
  Verifying        : kernel-4.18.0-348.7.1.el8_5.x86_64                      2/654
  Verifying        : kernel-core-4.18.0-348.7.1.el8_5.x86_64                 3/654
  Verifying        : kernel-modules-4.18.0-348.7.1.el8_5.x86_64              4/654
  Verifying        : libbpf-0.4.0-1.el8.x86_64                               5/654
  Verifying        : authselect-compat-1.2.2-3.el8.x86_64                    6/654
  Verifying        : authselect-compat-1.2.2-2.el8.x86_64                    7/654
  Verifying        : bind-libs-32:9.11.26-6.el8.x86_64                       8/654
  Verifying        : bind-libs-32:9.11.26-3.el8.x86_64                       9/654
  Verifying        : bind-libs-lite-32:9.11.26-6.el8.x86_64                 10/654
  Verifying        : bind-libs-lite-32:9.11.26-3.el8.x86_64                 11/654
  Verifying        : bind-license-32:9.11.26-6.el8.noarch                   12/654
  Verifying        : bind-license-32:9.11.26-3.el8.noarch                   13/654
  Verifying        : bind-utils-32:9.11.26-6.el8.x86_64                     14/654
  Verifying        : bind-utils-32:9.11.26-3.el8.x86_64                     15/654
  Verifying        : buildah-1.22.3-2.module_el8.5.0+911+f19012f9.x86_64    16/654
  Verifying        : buildah-1.19.7-1.module_el8.4.0+781+acf4c33b.x86_64    17/654
  Verifying        : cockpit-packagekit-251.1-1.el8.noarch                  18/654
  Verifying        : cockpit-packagekit-238.2-1.el8.noarch                  19/654
  Verifying        : cockpit-podman-33-1.module_el8.5.0+890+6b136101.noa    20/654
  Verifying        : cockpit-podman-29-2.module_el8.4.0+781+acf4c33b.noa    21/654
  Verifying        : cockpit-storaged-251.1-1.el8.noarch                    22/654
  Verifying        : cockpit-storaged-238.2-1.el8.noarch                    23/654
  Verifying        : conmon-2:2.0.29-1.module_el8.5.0+890+6b136101.x86_6    24/654
  Verifying        : conmon-2:2.0.26-1.module_el8.4.0+781+acf4c33b.x86_6    25/654
  Verifying        : container-selinux-2:2.167.0-1.module_el8.5.0+911+f1    26/654
  Verifying        : container-selinux-2:2.158.0-1.module_el8.4.0+781+ac    27/654
  Verifying        : containernetworking-plugins-1.0.0-1.module_el8.5.0+    28/654
  Verifying        : containernetworking-plugins-0.9.1-1.module_el8.4.0+    29/654
  Verifying        : containers-common-2:1-2.module_el8.5.0+890+6b136101    30/654
  Verifying        : containers-common-1:1.2.2-8.module_el8.4.0+781+acf4    31/654
  Verifying        : criu-3.15-3.module_el8.5.0+890+6b136101.x86_64         32/654
  Verifying        : criu-3.15-1.module_el8.4.0+641+6116a774.x86_64         33/654
  Verifying        : fstrm-0.6.1-2.el8.x86_64                               34/654
  Verifying        : fstrm-0.6.0-3.el8.1.x86_64                             35/654
  Verifying        : fuse-overlayfs-1.7.1-1.module_el8.5.0+890+6b136101.    36/654
  Verifying        : fuse-overlayfs-1.4.0-2.module_el8.4.0+673+eabfc99d.    37/654
  Verifying        : libX11-1.6.8-5.el8.x86_64                              38/654
  Verifying        : libX11-1.6.8-4.el8.x86_64                              39/654
  Verifying        : libX11-common-1.6.8-5.el8.noarch                       40/654
  Verifying        : libX11-common-1.6.8-4.el8.noarch                       41/654
  Verifying        : libblockdev-2.24-7.el8.x86_64                          42/654
  Verifying        : libblockdev-2.24-5.el8.x86_64                          43/654
  Verifying        : libblockdev-crypto-2.24-7.el8.x86_64                   44/654
  Verifying        : libblockdev-crypto-2.24-5.el8.x86_64                   45/654
  Verifying        : libblockdev-fs-2.24-7.el8.x86_64                       46/654
  Verifying        : libblockdev-fs-2.24-5.el8.x86_64                       47/654
  Verifying        : libblockdev-loop-2.24-7.el8.x86_64                     48/654
  Verifying        : libblockdev-loop-2.24-5.el8.x86_64                     49/654
  Verifying        : libblockdev-lvm-2.24-7.el8.x86_64                      50/654
  Verifying        : libblockdev-lvm-2.24-5.el8.x86_64                      51/654
  Verifying        : libblockdev-mdraid-2.24-7.el8.x86_64                   52/654
  Verifying        : libblockdev-mdraid-2.24-5.el8.x86_64                   53/654
  Verifying        : libblockdev-part-2.24-7.el8.x86_64                     54/654
  Verifying        : libblockdev-part-2.24-5.el8.x86_64                     55/654
  Verifying        : libblockdev-swap-2.24-7.el8.x86_64                     56/654
  Verifying        : libblockdev-swap-2.24-5.el8.x86_64                     57/654
  Verifying        : libblockdev-utils-2.24-7.el8.x86_64                    58/654
  Verifying        : libblockdev-utils-2.24-5.el8.x86_64                    59/654
  Verifying        : libdrm-2.4.106-2.el8.x86_64                            60/654
  Verifying        : libdrm-2.4.103-1.el8.x86_64                            61/654
  Verifying        : libfastjson-0.99.9-1.el8.x86_64                        62/654
  Verifying        : libfastjson-0.99.8-2.el8.x86_64                        63/654
  Verifying        : librelp-1.9.0-1.el8.x86_64                             64/654
  Verifying        : librelp-1.2.16-1.el8.x86_64                            65/654
  Verifying        : libslirp-4.4.0-1.module_el8.5.0+890+6b136101.x86_64    66/654
  Verifying        : libslirp-4.3.1-1.module_el8.4.0+575+63b40ad7.x86_64    67/654
  Verifying        : libudisks2-2.9.0-7.el8.x86_64                          68/654
  Verifying        : libudisks2-2.9.0-6.el8.x86_64                          69/654
  Verifying        : man-pages-overrides-8.5.0.1-1.el8.noarch               70/654
  Verifying        : man-pages-overrides-8.3.0.2-2.el8.noarch               71/654
  Verifying        : nmap-ncat-2:7.70-6.el8.x86_64                          72/654
  Verifying        : nmap-ncat-2:7.70-5.el8.x86_64                          73/654
  Verifying        : nspr-4.32.0-1.el8_4.x86_64                             74/654
  Verifying        : nspr-4.25.0-2.el8_2.x86_64                             75/654
  Verifying        : nss-3.67.0-7.el8_5.x86_64                              76/654
  Verifying        : nss-3.53.1-17.el8_3.x86_64                             77/654
  Verifying        : nss-softokn-3.67.0-7.el8_5.x86_64                      78/654
  Verifying        : nss-softokn-3.53.1-17.el8_3.x86_64                     79/654
  Verifying        : nss-softokn-freebl-3.67.0-7.el8_5.x86_64               80/654
  Verifying        : nss-softokn-freebl-3.53.1-17.el8_3.x86_64              81/654
  Verifying        : nss-sysinit-3.67.0-7.el8_5.x86_64                      82/654
  Verifying        : nss-sysinit-3.53.1-17.el8_3.x86_64                     83/654
  Verifying        : nss-util-3.67.0-7.el8_5.x86_64                         84/654
  Verifying        : nss-util-3.53.1-17.el8_3.x86_64                        85/654
  Verifying        : open-vm-tools-11.2.5-2.el8.x86_64                      86/654
  Verifying        : open-vm-tools-11.2.0-2.el8.x86_64                      87/654
  Verifying        : plymouth-0.9.4-10.20200615git1e36e30.el8.x86_64        88/654
  Verifying        : plymouth-0.9.4-9.20200615git1e36e30.el8.x86_64         89/654
  Verifying        : plymouth-core-libs-0.9.4-10.20200615git1e36e30.el8.    90/654
  Verifying        : plymouth-core-libs-0.9.4-9.20200615git1e36e30.el8.x    91/654
  Verifying        : plymouth-scripts-0.9.4-10.20200615git1e36e30.el8.x8    92/654
  Verifying        : plymouth-scripts-0.9.4-9.20200615git1e36e30.el8.x86    93/654
  Verifying        : podman-3.3.1-9.module_el8.5.0+988+b1f0b741.x86_64      94/654
  Verifying        : podman-3.0.1-6.module_el8.4.0+781+acf4c33b.x86_64      95/654
  Verifying        : podman-catatonit-3.3.1-9.module_el8.5.0+988+b1f0b74    96/654
  Verifying        : podman-catatonit-3.0.1-6.module_el8.4.0+781+acf4c33    97/654
  Verifying        : python3-bind-32:9.11.26-6.el8.noarch                   98/654
  Verifying        : python3-bind-32:9.11.26-3.el8.noarch                   99/654
  Verifying        : python3-lxml-4.2.3-3.el8.x86_64                       100/654
  Verifying        : python3-lxml-4.2.3-2.el8.x86_64                       101/654
  Verifying        : python3-psutil-5.4.3-11.el8.x86_64                    102/654
  Verifying        : python3-psutil-5.4.3-10.el8.x86_64                    103/654
  Verifying        : python3-unbound-1.7.3-17.el8.x86_64                   104/654
  Verifying        : python3-unbound-1.7.3-15.el8.x86_64                   105/654
  Verifying        : rsyslog-8.2102.0-5.el8.x86_64                         106/654
  Verifying        : rsyslog-8.1911.0-7.el8.x86_64                         107/654
  Verifying        : rsyslog-gnutls-8.2102.0-5.el8.x86_64                  108/654
  Verifying        : rsyslog-gnutls-8.1911.0-7.el8.x86_64                  109/654
  Verifying        : rsyslog-gssapi-8.2102.0-5.el8.x86_64                  110/654
  Verifying        : rsyslog-gssapi-8.1911.0-7.el8.x86_64                  111/654
  Verifying        : rsyslog-relp-8.2102.0-5.el8.x86_64                    112/654
  Verifying        : rsyslog-relp-8.1911.0-7.el8.x86_64                    113/654
  Verifying        : runc-1.0.2-1.module_el8.5.0+911+f19012f9.x86_64       114/654
  Verifying        : runc-1.0.0-70.rc92.module_el8.4.0+673+eabfc99d.x86_   115/654
  Verifying        : setroubleshoot-plugins-3.3.14-1.el8.noarch            116/654
  Verifying        : setroubleshoot-plugins-3.3.13-1.el8.noarch            117/654
  Verifying        : setroubleshoot-server-3.3.24-4.el8.x86_64             118/654
  Verifying        : setroubleshoot-server-3.3.24-3.el8.x86_64             119/654
  Verifying        : slirp4netns-1.1.8-1.module_el8.5.0+890+6b136101.x86   120/654
  Verifying        : slirp4netns-1.1.8-1.module_el8.4.0+641+6116a774.x86   121/654
  Verifying        : tcpdump-14:4.9.3-2.el8.x86_64                         122/654
  Verifying        : tcpdump-14:4.9.3-1.el8.x86_64                         123/654
  Verifying        : udisks2-2.9.0-7.el8.x86_64                            124/654
  Verifying        : udisks2-2.9.0-6.el8.x86_64                            125/654
  Verifying        : udisks2-iscsi-2.9.0-7.el8.x86_64                      126/654
  Verifying        : udisks2-iscsi-2.9.0-6.el8.x86_64                      127/654
  Verifying        : udisks2-lvm2-2.9.0-7.el8.x86_64                       128/654
  Verifying        : udisks2-lvm2-2.9.0-6.el8.x86_64                       129/654
  Verifying        : unbound-libs-1.7.3-17.el8.x86_64                      130/654
  Verifying        : unbound-libs-1.7.3-15.el8.x86_64                      131/654
  Verifying        : vim-common-2:8.0.1763-16.el8.x86_64                   132/654
  Verifying        : vim-common-2:8.0.1763-15.el8.x86_64                   133/654
  Verifying        : vim-enhanced-2:8.0.1763-16.el8.x86_64                 134/654
  Verifying        : vim-enhanced-2:8.0.1763-15.el8.x86_64                 135/654
  Verifying        : vim-filesystem-2:8.0.1763-16.el8.noarch               136/654
  Verifying        : vim-filesystem-2:8.0.1763-15.el8.noarch               137/654
  Verifying        : NetworkManager-1:1.32.10-4.el8.x86_64                 138/654
  Verifying        : NetworkManager-1:1.30.0-7.el8.x86_64                  139/654
  Verifying        : NetworkManager-config-server-1:1.32.10-4.el8.noarch   140/654
  Verifying        : NetworkManager-config-server-1:1.30.0-7.el8.noarch    141/654
  Verifying        : NetworkManager-libnm-1:1.32.10-4.el8.x86_64           142/654
  Verifying        : NetworkManager-libnm-1:1.30.0-7.el8.x86_64            143/654
  Verifying        : NetworkManager-team-1:1.32.10-4.el8.x86_64            144/654
  Verifying        : NetworkManager-team-1:1.30.0-7.el8.x86_64             145/654
  Verifying        : NetworkManager-tui-1:1.32.10-4.el8.x86_64             146/654
  Verifying        : NetworkManager-tui-1:1.30.0-7.el8.x86_64              147/654
  Verifying        : adcli-0.8.2-12.el8.x86_64                             148/654
  Verifying        : adcli-0.8.2-9.el8.x86_64                              149/654
  Verifying        : alsa-sof-firmware-1.8-1.el8.noarch                    150/654
  Verifying        : alsa-sof-firmware-1.6.1-2.el8.noarch                  151/654
  Verifying        : authselect-1.2.2-3.el8.x86_64                         152/654
  Verifying        : authselect-1.2.2-2.el8.x86_64                         153/654
  Verifying        : authselect-libs-1.2.2-3.el8.x86_64                    154/654
  Verifying        : authselect-libs-1.2.2-2.el8.x86_64                    155/654
  Verifying        : bash-4.4.20-2.el8.x86_64                              156/654
  Verifying        : bash-4.4.19-14.el8.x86_64                             157/654
  Verifying        : bind-export-libs-32:9.11.26-6.el8.x86_64              158/654
  Verifying        : bind-export-libs-32:9.11.26-3.el8.x86_64              159/654
  Verifying        : binutils-2.30-108.el8_5.1.x86_64                      160/654
  Verifying        : binutils-2.30-93.el8.x86_64                           161/654
  Verifying        : bpftool-4.18.0-348.7.1.el8_5.x86_64                   162/654
  Verifying        : bpftool-4.18.0-305.3.1.el8.x86_64                     163/654
  Verifying        : ca-certificates-2021.2.50-80.0.el8_4.noarch           164/654
  Verifying        : ca-certificates-2020.2.41-80.0.el8_2.noarch           165/654
  Verifying        : centos-gpg-keys-1:8-3.el8.noarch                      166/654
  Verifying        : centos-gpg-keys-1:8-2.el8.noarch                      167/654
  Verifying        : centos-linux-release-8.5-1.2111.el8.noarch            168/654
  Verifying        : centos-linux-release-8.4-1.2105.el8.noarch            169/654
  Verifying        : centos-linux-repos-8-3.el8.noarch                     170/654
  Verifying        : centos-linux-repos-8-2.el8.noarch                     171/654
  Verifying        : centos-logos-85.8-2.el8.x86_64                        172/654
  Verifying        : centos-logos-85.5-1.el8.x86_64                        173/654
  Verifying        : chkconfig-1.19.1-1.el8.x86_64                         174/654
  Verifying        : chkconfig-1.13-2.el8.x86_64                           175/654
  Verifying        : chrony-4.1-1.el8.x86_64                               176/654
  Verifying        : chrony-3.5-2.el8.x86_64                               177/654
  Verifying        : cockpit-251.1-1.el8.x86_64                            178/654
  Verifying        : cockpit-238.2-1.el8.x86_64                            179/654
  Verifying        : cockpit-bridge-251.1-1.el8.x86_64                     180/654
  Verifying        : cockpit-bridge-238.2-1.el8.x86_64                     181/654
  Verifying        : cockpit-system-251.1-1.el8.noarch                     182/654
  Verifying        : cockpit-system-238.2-1.el8.noarch                     183/654
  Verifying        : cockpit-ws-251.1-1.el8.x86_64                         184/654
  Verifying        : cockpit-ws-238.2-1.el8.x86_64                         185/654
  Verifying        : coreutils-8.30-12.el8.x86_64                          186/654
  Verifying        : coreutils-8.30-8.el8.x86_64                           187/654
  Verifying        : coreutils-common-8.30-12.el8.x86_64                   188/654
  Verifying        : coreutils-common-8.30-8.el8.x86_64                    189/654
  Verifying        : crypto-policies-20210617-1.gitc776d3e.el8.noarch      190/654
  Verifying        : crypto-policies-20210209-1.gitbfb6bed.el8_3.noarch    191/654
  Verifying        : crypto-policies-scripts-20210617-1.gitc776d3e.el8.n   192/654
  Verifying        : crypto-policies-scripts-20210209-1.gitbfb6bed.el8_3   193/654
  Verifying        : cups-libs-1:2.2.6-40.el8.x86_64                       194/654
  Verifying        : cups-libs-1:2.2.6-38.el8.x86_64                       195/654
  Verifying        : curl-7.61.1-22.el8.x86_64                             196/654
  Verifying        : curl-7.61.1-18.el8.x86_64                             197/654
  Verifying        : dbus-1:1.12.8-14.el8.x86_64                           198/654
  Verifying        : dbus-1:1.12.8-12.el8.x86_64                           199/654
  Verifying        : dbus-common-1:1.12.8-14.el8.noarch                    200/654
  Verifying        : dbus-common-1:1.12.8-12.el8.noarch                    201/654
  Verifying        : dbus-daemon-1:1.12.8-14.el8.x86_64                    202/654
  Verifying        : dbus-daemon-1:1.12.8-12.el8.x86_64                    203/654
  Verifying        : dbus-libs-1:1.12.8-14.el8.x86_64                      204/654
  Verifying        : dbus-libs-1:1.12.8-12.el8.x86_64                      205/654
  Verifying        : dbus-tools-1:1.12.8-14.el8.x86_64                     206/654
  Verifying        : dbus-tools-1:1.12.8-12.el8.x86_64                     207/654
  Verifying        : device-mapper-8:1.02.177-10.el8.x86_64                208/654
  Verifying        : device-mapper-8:1.02.175-5.el8.x86_64                 209/654
  Verifying        : device-mapper-event-8:1.02.177-10.el8.x86_64          210/654
  Verifying        : device-mapper-event-8:1.02.175-5.el8.x86_64           211/654
  Verifying        : device-mapper-event-libs-8:1.02.177-10.el8.x86_64     212/654
  Verifying        : device-mapper-event-libs-8:1.02.175-5.el8.x86_64      213/654
  Verifying        : device-mapper-libs-8:1.02.177-10.el8.x86_64           214/654
  Verifying        : device-mapper-libs-8:1.02.175-5.el8.x86_64            215/654
  Verifying        : device-mapper-multipath-0.8.4-17.el8.x86_64           216/654
  Verifying        : device-mapper-multipath-0.8.4-10.el8.x86_64           217/654
  Verifying        : device-mapper-multipath-libs-0.8.4-17.el8.x86_64      218/654
  Verifying        : device-mapper-multipath-libs-0.8.4-10.el8.x86_64      219/654
  Verifying        : device-mapper-persistent-data-0.9.0-4.el8.x86_64      220/654
  Verifying        : device-mapper-persistent-data-0.8.5-4.el8.x86_64      221/654
  Verifying        : dhcp-client-12:4.3.6-45.el8.x86_64                    222/654
  Verifying        : dhcp-client-12:4.3.6-44.0.1.el8.x86_64                223/654
  Verifying        : dhcp-common-12:4.3.6-45.el8.noarch                    224/654
  Verifying        : dhcp-common-12:4.3.6-44.0.1.el8.noarch                225/654
  Verifying        : dhcp-libs-12:4.3.6-45.el8.x86_64                      226/654
  Verifying        : dhcp-libs-12:4.3.6-44.0.1.el8.x86_64                  227/654
  Verifying        : dmidecode-1:3.2-10.el8.x86_64                         228/654
  Verifying        : dmidecode-1:3.2-8.el8.x86_64                          229/654
  Verifying        : dnf-4.7.0-4.el8.noarch                                230/654
  Verifying        : dnf-4.4.2-11.el8.noarch                               231/654
  Verifying        : dnf-data-4.7.0-4.el8.noarch                           232/654
  Verifying        : dnf-data-4.4.2-11.el8.noarch                          233/654
  Verifying        : dnf-plugins-core-4.0.21-3.el8.noarch                  234/654
  Verifying        : dnf-plugins-core-4.0.18-4.el8.noarch                  235/654
  Verifying        : dracut-049-191.git20210920.el8.x86_64                 236/654
  Verifying        : dracut-049-135.git20210121.el8.x86_64                 237/654
  Verifying        : dracut-config-rescue-049-191.git20210920.el8.x86_64   238/654
  Verifying        : dracut-config-rescue-049-135.git20210121.el8.x86_64   239/654
  Verifying        : dracut-network-049-191.git20210920.el8.x86_64         240/654
  Verifying        : dracut-network-049-135.git20210121.el8.x86_64         241/654
  Verifying        : dracut-squash-049-191.git20210920.el8.x86_64          242/654
  Verifying        : dracut-squash-049-135.git20210121.el8.x86_64          243/654
  Verifying        : e2fsprogs-1.45.6-2.el8.x86_64                         244/654
  Verifying        : e2fsprogs-1.45.6-1.el8.x86_64                         245/654
  Verifying        : e2fsprogs-libs-1.45.6-2.el8.x86_64                    246/654
  Verifying        : e2fsprogs-libs-1.45.6-1.el8.x86_64                    247/654
  Verifying        : elfutils-debuginfod-client-0.185-1.el8.x86_64         248/654
  Verifying        : elfutils-debuginfod-client-0.182-3.el8.x86_64         249/654
  Verifying        : elfutils-default-yama-scope-0.185-1.el8.noarch        250/654
  Verifying        : elfutils-default-yama-scope-0.182-3.el8.noarch        251/654
  Verifying        : elfutils-libelf-0.185-1.el8.x86_64                    252/654
  Verifying        : elfutils-libelf-0.182-3.el8.x86_64                    253/654
  Verifying        : elfutils-libs-0.185-1.el8.x86_64                      254/654
  Verifying        : elfutils-libs-0.182-3.el8.x86_64                      255/654
  Verifying        : emacs-filesystem-1:26.1-7.el8.noarch                  256/654
  Verifying        : emacs-filesystem-1:26.1-5.el8.noarch                  257/654
  Verifying        : ethtool-2:5.8-7.el8.x86_64                            258/654
  Verifying        : ethtool-2:5.8-5.el8.x86_64                            259/654
  Verifying        : file-5.33-20.el8.x86_64                               260/654
  Verifying        : file-5.33-16.el8_3.1.x86_64                           261/654
  Verifying        : file-libs-5.33-20.el8.x86_64                          262/654
  Verifying        : file-libs-5.33-16.el8_3.1.x86_64                      263/654
  Verifying        : filesystem-3.8-6.el8.x86_64                           264/654
  Verifying        : filesystem-3.8-3.el8.x86_64                           265/654
  Verifying        : firewalld-0.9.3-7.el8.noarch                          266/654
  Verifying        : firewalld-0.8.2-6.el8.noarch                          267/654
  Verifying        : firewalld-filesystem-0.9.3-7.el8.noarch               268/654
  Verifying        : firewalld-filesystem-0.8.2-6.el8.noarch               269/654
  Verifying        : fontconfig-2.13.1-4.el8.x86_64                        270/654
  Verifying        : fontconfig-2.13.1-3.el8.x86_64                        271/654
  Verifying        : glib2-2.56.4-156.el8.x86_64                           272/654
  Verifying        : glib2-2.56.4-9.el8.x86_64                             273/654
  Verifying        : glibc-2.28-164.el8.x86_64                             274/654
  Verifying        : glibc-2.28-151.el8.x86_64                             275/654
  Verifying        : glibc-common-2.28-164.el8.x86_64                      276/654
  Verifying        : glibc-common-2.28-151.el8.x86_64                      277/654
  Verifying        : glibc-langpack-en-2.28-164.el8.x86_64                 278/654
  Verifying        : glibc-langpack-en-2.28-151.el8.x86_64                 279/654
  Verifying        : gnutls-3.6.16-4.el8.x86_64                            280/654
  Verifying        : gnutls-3.6.14-7.el8_3.x86_64                          281/654
  Verifying        : gpgme-1.13.1-9.el8.x86_64                             282/654
  Verifying        : gpgme-1.13.1-7.el8.x86_64                             283/654
  Verifying        : grub2-common-1:2.02-106.el8.noarch                    284/654
  Verifying        : grub2-common-1:2.02-99.el8.noarch                     285/654
  Verifying        : grub2-pc-1:2.02-106.el8.x86_64                        286/654
  Verifying        : grub2-pc-1:2.02-99.el8.x86_64                         287/654
  Verifying        : grub2-pc-modules-1:2.02-106.el8.noarch                288/654
  Verifying        : grub2-pc-modules-1:2.02-99.el8.noarch                 289/654
  Verifying        : grub2-tools-1:2.02-106.el8.x86_64                     290/654
  Verifying        : grub2-tools-1:2.02-99.el8.x86_64                      291/654
  Verifying        : grub2-tools-extra-1:2.02-106.el8.x86_64               292/654
  Verifying        : grub2-tools-extra-1:2.02-99.el8.x86_64                293/654
  Verifying        : grub2-tools-minimal-1:2.02-106.el8.x86_64             294/654
  Verifying        : grub2-tools-minimal-1:2.02-99.el8.x86_64              295/654
  Verifying        : grubby-8.40-42.el8.x86_64                             296/654
  Verifying        : grubby-8.40-41.el8.x86_64                             297/654
  Verifying        : gsettings-desktop-schemas-3.32.0-6.el8.x86_64         298/654
  Verifying        : gsettings-desktop-schemas-3.32.0-5.el8.x86_64         299/654
  Verifying        : hdparm-9.54-4.el8.x86_64                              300/654
  Verifying        : hdparm-9.54-3.el8.x86_64                              301/654
  Verifying        : hwdata-0.314-8.10.el8.noarch                          302/654
  Verifying        : hwdata-0.314-8.8.el8.noarch                           303/654
  Verifying        : iproute-5.12.0-4.el8.x86_64                           304/654
  Verifying        : iproute-5.9.0-4.el8.x86_64                            305/654
  Verifying        : iptables-1.8.4-20.el8.x86_64                          306/654
  Verifying        : iptables-1.8.4-17.el8.x86_64                          307/654
  Verifying        : iptables-ebtables-1.8.4-20.el8.x86_64                 308/654
  Verifying        : iptables-ebtables-1.8.4-17.el8.x86_64                 309/654
  Verifying        : iptables-libs-1.8.4-20.el8.x86_64                     310/654
  Verifying        : iptables-libs-1.8.4-17.el8.x86_64                     311/654
  Verifying        : iscsi-initiator-utils-6.2.1.4-4.git095f59c.el8.x86_   312/654
  Verifying        : iscsi-initiator-utils-6.2.1.2-1.gita8fcb37.el8.x86_   313/654
  Verifying        : iscsi-initiator-utils-iscsiuio-6.2.1.4-4.git095f59c   314/654
  Verifying        : iscsi-initiator-utils-iscsiuio-6.2.1.2-1.gita8fcb37   315/654
  Verifying        : iwl100-firmware-39.31.5.1-103.el8.1.noarch            316/654
  Verifying        : iwl100-firmware-39.31.5.1-102.el8.1.noarch            317/654
  Verifying        : iwl1000-firmware-1:39.31.5.1-103.el8.1.noarch         318/654
  Verifying        : iwl1000-firmware-1:39.31.5.1-102.el8.1.noarch         319/654
  Verifying        : iwl105-firmware-18.168.6.1-103.el8.1.noarch           320/654
  Verifying        : iwl105-firmware-18.168.6.1-102.el8.1.noarch           321/654
  Verifying        : iwl135-firmware-18.168.6.1-103.el8.1.noarch           322/654
  Verifying        : iwl135-firmware-18.168.6.1-102.el8.1.noarch           323/654
  Verifying        : iwl2000-firmware-18.168.6.1-103.el8.1.noarch          324/654
  Verifying        : iwl2000-firmware-18.168.6.1-102.el8.1.noarch          325/654
  Verifying        : iwl2030-firmware-18.168.6.1-103.el8.1.noarch          326/654
  Verifying        : iwl2030-firmware-18.168.6.1-102.el8.1.noarch          327/654
  Verifying        : iwl3160-firmware-1:25.30.13.0-103.el8.1.noarch        328/654
  Verifying        : iwl3160-firmware-1:25.30.13.0-102.el8.1.noarch        329/654
  Verifying        : iwl5000-firmware-8.83.5.1_1-103.el8.1.noarch          330/654
  Verifying        : iwl5000-firmware-8.83.5.1_1-102.el8.1.noarch          331/654
  Verifying        : iwl5150-firmware-8.24.2.2-103.el8.1.noarch            332/654
  Verifying        : iwl5150-firmware-8.24.2.2-102.el8.1.noarch            333/654
  Verifying        : iwl6000-firmware-9.221.4.1-103.el8.1.noarch           334/654
  Verifying        : iwl6000-firmware-9.221.4.1-102.el8.1.noarch           335/654
  Verifying        : iwl6000g2a-firmware-18.168.6.1-103.el8.1.noarch       336/654
  Verifying        : iwl6000g2a-firmware-18.168.6.1-102.el8.1.noarch       337/654
  Verifying        : iwl6000g2b-firmware-18.168.6.1-103.el8.1.noarch       338/654
  Verifying        : iwl6000g2b-firmware-18.168.6.1-102.el8.1.noarch       339/654
  Verifying        : iwl6050-firmware-41.28.5.1-103.el8.1.noarch           340/654
  Verifying        : iwl6050-firmware-41.28.5.1-102.el8.1.noarch           341/654
  Verifying        : iwl7260-firmware-1:25.30.13.0-103.el8.1.noarch        342/654
  Verifying        : iwl7260-firmware-1:25.30.13.0-102.el8.1.noarch        343/654
  Verifying        : json-c-0.13.1-2.el8.x86_64                            344/654
  Verifying        : json-c-0.13.1-0.4.el8.x86_64                          345/654
  Verifying        : kernel-tools-4.18.0-348.7.1.el8_5.x86_64              346/654
  Verifying        : kernel-tools-4.18.0-305.3.1.el8.x86_64                347/654
  Verifying        : kernel-tools-libs-4.18.0-348.7.1.el8_5.x86_64         348/654
  Verifying        : kernel-tools-libs-4.18.0-305.3.1.el8.x86_64           349/654
  Verifying        : kexec-tools-2.0.20-57.el8_5.1.x86_64                  350/654
  Verifying        : kexec-tools-2.0.20-46.el8.x86_64                      351/654
  Verifying        : keyutils-libs-1.5.10-9.el8.x86_64                     352/654
  Verifying        : keyutils-libs-1.5.10-6.el8.x86_64                     353/654
  Verifying        : kmod-25-18.el8.x86_64                                 354/654
  Verifying        : kmod-25-17.el8.x86_64                                 355/654
  Verifying        : kmod-kvdo-6.2.5.72-81.el8.x86_64                      356/654
  Verifying        : kmod-kvdo-6.2.4.26-77.el8.x86_64                      357/654
  Verifying        : kmod-libs-25-18.el8.x86_64                            358/654
  Verifying        : kmod-libs-25-17.el8.x86_64                            359/654
  Verifying        : kpartx-0.8.4-17.el8.x86_64                            360/654
  Verifying        : kpartx-0.8.4-10.el8.x86_64                            361/654
  Verifying        : kpatch-0.9.2-5.el8.noarch                             362/654
  Verifying        : kpatch-0.9.2-3.el8.noarch                             363/654
  Verifying        : kpatch-dnf-0.2-5.el8.noarch                           364/654
  Verifying        : kpatch-dnf-0.2-3.el8.noarch                           365/654
  Verifying        : krb5-libs-1.18.2-14.el8.x86_64                        366/654
  Verifying        : krb5-libs-1.18.2-8.el8.x86_64                         367/654
  Verifying        : libblkid-2.32.1-28.el8.x86_64                         368/654
  Verifying        : libblkid-2.32.1-27.el8.x86_64                         369/654
  Verifying        : libcap-2.26-5.el8.x86_64                              370/654
  Verifying        : libcap-2.26-4.el8.x86_64                              371/654
  Verifying        : libcap-ng-0.7.11-1.el8.x86_64                         372/654
  Verifying        : libcap-ng-0.7.9-5.el8.x86_64                          373/654
  Verifying        : libcom_err-1.45.6-2.el8.x86_64                        374/654
  Verifying        : libcom_err-1.45.6-1.el8.x86_64                        375/654
  Verifying        : libcomps-0.1.16-2.el8.x86_64                          376/654
  Verifying        : libcomps-0.1.11-5.el8.x86_64                          377/654
  Verifying        : libcurl-7.61.1-22.el8.x86_64                          378/654
  Verifying        : libcurl-7.61.1-18.el8.x86_64                          379/654
  Verifying        : libdb-5.3.28-42.el8_4.x86_64                          380/654
  Verifying        : libdb-5.3.28-40.el8.x86_64                            381/654
  Verifying        : libdb-utils-5.3.28-42.el8_4.x86_64                    382/654
  Verifying        : libdb-utils-5.3.28-40.el8.x86_64                      383/654
  Verifying        : libdnf-0.63.0-3.el8.x86_64                            384/654
  Verifying        : libdnf-0.55.0-7.el8.x86_64                            385/654
  Verifying        : libertas-usb8388-firmware-2:20210702-103.gitd79c267   386/654
  Verifying        : libertas-usb8388-firmware-2:20201218-102.git0578970   387/654
  Verifying        : libfdisk-2.32.1-28.el8.x86_64                         388/654
  Verifying        : libfdisk-2.32.1-27.el8.x86_64                         389/654
  Verifying        : libgcc-8.5.0-4.el8_5.x86_64                           390/654
  Verifying        : libgcc-8.4.1-1.el8.x86_64                             391/654
  Verifying        : libgcrypt-1.8.5-6.el8.x86_64                          392/654
  Verifying        : libgcrypt-1.8.5-4.el8.x86_64                          393/654
  Verifying        : libgomp-8.5.0-4.el8_5.x86_64                          394/654
  Verifying        : libgomp-8.4.1-1.el8.x86_64                            395/654
  Verifying        : libibverbs-35.0-1.el8.x86_64                          396/654
  Verifying        : libibverbs-32.0-4.el8.x86_64                          397/654
  Verifying        : libipa_hbac-2.5.2-2.el8_5.3.x86_64                    398/654
  Verifying        : libipa_hbac-2.4.0-9.el8.x86_64                        399/654
  Verifying        : libldb-2.3.0-2.el8.x86_64                             400/654
  Verifying        : libldb-2.2.0-2.el8.x86_64                             401/654
  Verifying        : libmodulemd-2.13.0-1.el8.x86_64                       402/654
  Verifying        : libmodulemd-2.9.4-2.el8.x86_64                        403/654
  Verifying        : libmount-2.32.1-28.el8.x86_64                         404/654
  Verifying        : libmount-2.32.1-27.el8.x86_64                         405/654
  Verifying        : libndp-1.7-6.el8.x86_64                               406/654
  Verifying        : libndp-1.7-5.el8.x86_64                               407/654
  Verifying        : libnfsidmap-1:2.3.3-46.el8.x86_64                     408/654
  Verifying        : libnfsidmap-1:2.3.3-41.el8.x86_64                     409/654
  Verifying        : librepo-1.14.0-2.el8.x86_64                           410/654
  Verifying        : librepo-1.12.0-3.el8.x86_64                           411/654
  Verifying        : libsepol-2.9-3.el8.x86_64                             412/654
  Verifying        : libsepol-2.9-2.el8.x86_64                             413/654
  Verifying        : libsmartcols-2.32.1-28.el8.x86_64                     414/654
  Verifying        : libsmartcols-2.32.1-27.el8.x86_64                     415/654
  Verifying        : libsmbclient-4.14.5-7.el8_5.x86_64                    416/654
  Verifying        : libsmbclient-4.13.3-3.el8.x86_64                      417/654
  Verifying        : libsolv-0.7.19-1.el8.x86_64                           418/654
  Verifying        : libsolv-0.7.16-2.el8.x86_64                           419/654
  Verifying        : libss-1.45.6-2.el8.x86_64                             420/654
  Verifying        : libss-1.45.6-1.el8.x86_64                             421/654
  Verifying        : libssh-0.9.4-3.el8.x86_64                             422/654
  Verifying        : libssh-0.9.4-2.el8.x86_64                             423/654
  Verifying        : libssh-config-0.9.4-3.el8.noarch                      424/654
  Verifying        : libssh-config-0.9.4-2.el8.noarch                      425/654
  Verifying        : libsss_autofs-2.5.2-2.el8_5.3.x86_64                  426/654
  Verifying        : libsss_autofs-2.4.0-9.el8.x86_64                      427/654
  Verifying        : libsss_certmap-2.5.2-2.el8_5.3.x86_64                 428/654
  Verifying        : libsss_certmap-2.4.0-9.el8.x86_64                     429/654
  Verifying        : libsss_idmap-2.5.2-2.el8_5.3.x86_64                   430/654
  Verifying        : libsss_idmap-2.4.0-9.el8.x86_64                       431/654
  Verifying        : libsss_nss_idmap-2.5.2-2.el8_5.3.x86_64               432/654
  Verifying        : libsss_nss_idmap-2.4.0-9.el8.x86_64                   433/654
  Verifying        : libsss_sudo-2.5.2-2.el8_5.3.x86_64                    434/654
  Verifying        : libsss_sudo-2.4.0-9.el8.x86_64                        435/654
  Verifying        : libstdc++-8.5.0-4.el8_5.x86_64                        436/654
  Verifying        : libstdc++-8.4.1-1.el8.x86_64                          437/654
  Verifying        : libstoragemgmt-1.9.1-1.el8.x86_64                     438/654
  Verifying        : libstoragemgmt-1.8.7-1.el8.x86_64                     439/654
  Verifying        : libtalloc-2.3.2-1.el8.x86_64                          440/654
  Verifying        : libtalloc-2.3.1-2.el8.x86_64                          441/654
  Verifying        : libtevent-0.11.0-0.el8.x86_64                         442/654
  Verifying        : libtevent-0.10.2-2.el8.x86_64                         443/654
  Verifying        : libtirpc-1.1.4-5.el8.x86_64                           444/654
  Verifying        : libtirpc-1.1.4-4.el8.x86_64                           445/654
  Verifying        : libuuid-2.32.1-28.el8.x86_64                          446/654
  Verifying        : libuuid-2.32.1-27.el8.x86_64                          447/654
  Verifying        : libwbclient-4.14.5-7.el8_5.x86_64                     448/654
  Verifying        : libwbclient-4.13.3-3.el8.x86_64                       449/654
  Verifying        : libxcrypt-4.1.1-6.el8.x86_64                          450/654
  Verifying        : libxcrypt-4.1.1-4.el8.x86_64                          451/654
  Verifying        : libxml2-2.9.7-9.el8_4.2.x86_64                        452/654
  Verifying        : libxml2-2.9.7-9.el8.x86_64                            453/654
  Verifying        : linux-firmware-20210702-103.gitd79c2677.el8.noarch    454/654
  Verifying        : linux-firmware-20201218-102.git05789708.el8.noarch    455/654
  Verifying        : lshw-B.02.19.2-6.el8.x86_64                           456/654
  Verifying        : lshw-B.02.19.2-5.el8.x86_64                           457/654
  Verifying        : lsscsi-0.32-3.el8.x86_64                              458/654
  Verifying        : lsscsi-0.32-2.el8.x86_64                              459/654
  Verifying        : lua-libs-5.3.4-12.el8.x86_64                          460/654
  Verifying        : lua-libs-5.3.4-11.el8.x86_64                          461/654
  Verifying        : lvm2-8:2.03.12-10.el8.x86_64                          462/654
  Verifying        : lvm2-8:2.03.11-5.el8.x86_64                           463/654
  Verifying        : lvm2-libs-8:2.03.12-10.el8.x86_64                     464/654
  Verifying        : lvm2-libs-8:2.03.11-5.el8.x86_64                      465/654
  Verifying        : lz4-libs-1.8.3-3.el8_4.x86_64                         466/654
  Verifying        : lz4-libs-1.8.3-2.el8.x86_64                           467/654
  Verifying        : man-db-2.7.6.1-18.el8.x86_64                          468/654
  Verifying        : man-db-2.7.6.1-17.el8.x86_64                          469/654
  Verifying        : mcelog-3:175-1.el8.x86_64                             470/654
  Verifying        : mcelog-3:173-0.el8.x86_64                             471/654
  Verifying        : mdadm-4.2-rc2.el8.x86_64                              472/654
  Verifying        : mdadm-4.1-15.el8.x86_64                               473/654
  Verifying        : microcode_ctl-4:20210608-1.el8.x86_64                 474/654
  Verifying        : microcode_ctl-4:20210216-1.el8.x86_64                 475/654
  Verifying        : ncurses-6.1-9.20180224.el8.x86_64                     476/654
  Verifying        : ncurses-6.1-7.20180224.el8.x86_64                     477/654
  Verifying        : ncurses-base-6.1-9.20180224.el8.noarch                478/654
  Verifying        : ncurses-base-6.1-7.20180224.el8.noarch                479/654
  Verifying        : ncurses-libs-6.1-9.20180224.el8.x86_64                480/654
  Verifying        : ncurses-libs-6.1-7.20180224.el8.x86_64                481/654
  Verifying        : nettle-3.4.1-7.el8.x86_64                             482/654
  Verifying        : nettle-3.4.1-2.el8.x86_64                             483/654
  Verifying        : nftables-1:0.9.3-21.el8.x86_64                        484/654
  Verifying        : nftables-1:0.9.3-18.el8.x86_64                        485/654
  Verifying        : numactl-libs-2.0.12-13.el8.x86_64                     486/654
  Verifying        : numactl-libs-2.0.12-11.el8.x86_64                     487/654
  Verifying        : nvme-cli-1.14-3.el8.x86_64                            488/654
  Verifying        : nvme-cli-1.12-3.el8.0.1.x86_64                        489/654
  Verifying        : openldap-2.4.46-18.el8.x86_64                         490/654
  Verifying        : openldap-2.4.46-16.el8.x86_64                         491/654
  Verifying        : openssh-8.0p1-10.el8.x86_64                           492/654
  Verifying        : openssh-8.0p1-5.el8.x86_64                            493/654
  Verifying        : openssh-clients-8.0p1-10.el8.x86_64                   494/654
  Verifying        : openssh-clients-8.0p1-5.el8.x86_64                    495/654
  Verifying        : openssh-server-8.0p1-10.el8.x86_64                    496/654
  Verifying        : openssh-server-8.0p1-5.el8.x86_64                     497/654
  Verifying        : openssl-1:1.1.1k-5.el8_5.x86_64                       498/654
  Verifying        : openssl-1:1.1.1g-15.el8_3.x86_64                      499/654
  Verifying        : openssl-libs-1:1.1.1k-5.el8_5.x86_64                  500/654
  Verifying        : openssl-libs-1:1.1.1g-15.el8_3.x86_64                 501/654
  Verifying        : os-prober-1.74-9.el8.x86_64                           502/654
  Verifying        : os-prober-1.74-6.el8.x86_64                           503/654
  Verifying        : pam-1.3.1-15.el8.x86_64                               504/654
  Verifying        : pam-1.3.1-14.el8.x86_64                               505/654
  Verifying        : parted-3.2-39.el8.x86_64                              506/654
  Verifying        : parted-3.2-38.el8.x86_64                              507/654
  Verifying        : pcre-8.42-6.el8.x86_64                                508/654
  Verifying        : pcre-8.42-4.el8.x86_64                                509/654
  Verifying        : platform-python-3.6.8-41.el8.x86_64                   510/654
  Verifying        : platform-python-3.6.8-37.el8.x86_64                   511/654
  Verifying        : platform-python-pip-9.0.3-20.el8.noarch               512/654
  Verifying        : platform-python-pip-9.0.3-19.el8.noarch               513/654
  Verifying        : policycoreutils-2.9-16.el8.x86_64                     514/654
  Verifying        : policycoreutils-2.9-14.el8.x86_64                     515/654
  Verifying        : policycoreutils-python-utils-2.9-16.el8.noarch        516/654
  Verifying        : policycoreutils-python-utils-2.9-14.el8.noarch        517/654
  Verifying        : polkit-0.115-12.el8.x86_64                            518/654
  Verifying        : polkit-0.115-11.el8.x86_64                            519/654
  Verifying        : polkit-libs-0.115-12.el8.x86_64                       520/654
  Verifying        : polkit-libs-0.115-11.el8.x86_64                       521/654
  Verifying        : python3-dnf-4.7.0-4.el8.noarch                        522/654
  Verifying        : python3-dnf-4.4.2-11.el8.noarch                       523/654
  Verifying        : python3-dnf-plugins-core-4.0.21-3.el8.noarch          524/654
  Verifying        : python3-dnf-plugins-core-4.0.18-4.el8.noarch          525/654
  Verifying        : python3-firewall-0.9.3-7.el8.noarch                   526/654
  Verifying        : python3-firewall-0.8.2-6.el8.noarch                   527/654
  Verifying        : python3-gpg-1.13.1-9.el8.x86_64                       528/654
  Verifying        : python3-gpg-1.13.1-7.el8.x86_64                       529/654
  Verifying        : python3-hawkey-0.63.0-3.el8.x86_64                    530/654
  Verifying        : python3-hawkey-0.55.0-7.el8.x86_64                    531/654
  Verifying        : python3-libcomps-0.1.16-2.el8.x86_64                  532/654
  Verifying        : python3-libcomps-0.1.11-5.el8.x86_64                  533/654
  Verifying        : python3-libdnf-0.63.0-3.el8.x86_64                    534/654
  Verifying        : python3-libdnf-0.55.0-7.el8.x86_64                    535/654
  Verifying        : python3-libs-3.6.8-41.el8.x86_64                      536/654
  Verifying        : python3-libs-3.6.8-37.el8.x86_64                      537/654
  Verifying        : python3-libstoragemgmt-1.9.1-1.el8.x86_64             538/654
  Verifying        : python3-libstoragemgmt-1.8.7-1.el8.noarch             539/654
  Verifying        : python3-libstoragemgmt-clibs-1.8.7-1.el8.x86_64       540/654
  Verifying        : python3-libxml2-2.9.7-9.el8_4.2.x86_64                541/654
  Verifying        : python3-libxml2-2.9.7-9.el8.x86_64                    542/654
  Verifying        : python3-nftables-1:0.9.3-21.el8.x86_64                543/654
  Verifying        : python3-nftables-1:0.9.3-18.el8.x86_64                544/654
  Verifying        : python3-perf-4.18.0-348.7.1.el8_5.x86_64              545/654
  Verifying        : python3-perf-4.18.0-305.3.1.el8.x86_64                546/654
  Verifying        : python3-pip-wheel-9.0.3-20.el8.noarch                 547/654
  Verifying        : python3-pip-wheel-9.0.3-19.el8.noarch                 548/654
  Verifying        : python3-policycoreutils-2.9-16.el8.noarch             549/654
  Verifying        : python3-policycoreutils-2.9-14.el8.noarch             550/654
  Verifying        : python3-rpm-4.14.3-19.el8.x86_64                      551/654
  Verifying        : python3-rpm-4.14.3-13.el8.x86_64                      552/654
  Verifying        : python3-sssdconfig-2.5.2-2.el8_5.3.noarch             553/654
  Verifying        : python3-sssdconfig-2.4.0-9.el8.noarch                 554/654
  Verifying        : python3-syspurpose-1.28.21-3.el8.x86_64               555/654
  Verifying        : python3-syspurpose-1.28.13-2.el8.x86_64               556/654
  Verifying        : quota-1:4.04-14.el8.x86_64                            557/654
  Verifying        : quota-1:4.04-12.el8.x86_64                            558/654
  Verifying        : quota-nls-1:4.04-14.el8.noarch                        559/654
  Verifying        : quota-nls-1:4.04-12.el8.noarch                        560/654
  Verifying        : rdma-core-35.0-1.el8.x86_64                           561/654
  Verifying        : rdma-core-32.0-4.el8.x86_64                           562/654
  Verifying        : realmd-0.16.3-23.el8.x86_64                           563/654
  Verifying        : realmd-0.16.3-22.el8.x86_64                           564/654
  Verifying        : rpm-4.14.3-19.el8.x86_64                              565/654
  Verifying        : rpm-4.14.3-13.el8.x86_64                              566/654
  Verifying        : rpm-build-libs-4.14.3-19.el8.x86_64                   567/654
  Verifying        : rpm-build-libs-4.14.3-13.el8.x86_64                   568/654
  Verifying        : rpm-libs-4.14.3-19.el8.x86_64                         569/654
  Verifying        : rpm-libs-4.14.3-13.el8.x86_64                         570/654
  Verifying        : rpm-plugin-selinux-4.14.3-19.el8.x86_64               571/654
  Verifying        : rpm-plugin-selinux-4.14.3-13.el8.x86_64               572/654
  Verifying        : rpm-plugin-systemd-inhibit-4.14.3-19.el8.x86_64       573/654
  Verifying        : rpm-plugin-systemd-inhibit-4.14.3-13.el8.x86_64       574/654
  Verifying        : samba-client-libs-4.14.5-7.el8_5.x86_64               575/654
  Verifying        : samba-client-libs-4.13.3-3.el8.x86_64                 576/654
  Verifying        : samba-common-4.14.5-7.el8_5.noarch                    577/654
  Verifying        : samba-common-4.13.3-3.el8.noarch                      578/654
  Verifying        : samba-common-libs-4.14.5-7.el8_5.x86_64               579/654
  Verifying        : samba-common-libs-4.13.3-3.el8.x86_64                 580/654
  Verifying        : selinux-policy-3.14.3-80.el8_5.2.noarch               581/654
  Verifying        : selinux-policy-3.14.3-67.el8.noarch                   582/654
  Verifying        : selinux-policy-targeted-3.14.3-80.el8_5.2.noarch      583/654
  Verifying        : selinux-policy-targeted-3.14.3-67.el8.noarch          584/654
  Verifying        : shadow-utils-2:4.6-14.el8.x86_64                      585/654
  Verifying        : shadow-utils-2:4.6-12.el8.x86_64                      586/654
  Verifying        : sos-4.1-5.el8.noarch                                  587/654
  Verifying        : sos-4.0-11.el8.noarch                                 588/654
  Verifying        : sqlite-3.26.0-15.el8.x86_64                           589/654
  Verifying        : sqlite-3.26.0-13.el8.x86_64                           590/654
  Verifying        : sqlite-libs-3.26.0-15.el8.x86_64                      591/654
  Verifying        : sqlite-libs-3.26.0-13.el8.x86_64                      592/654
  Verifying        : sssd-2.5.2-2.el8_5.3.x86_64                           593/654
  Verifying        : sssd-2.4.0-9.el8.x86_64                               594/654
  Verifying        : sssd-ad-2.5.2-2.el8_5.3.x86_64                        595/654
  Verifying        : sssd-ad-2.4.0-9.el8.x86_64                            596/654
  Verifying        : sssd-client-2.5.2-2.el8_5.3.x86_64                    597/654
  Verifying        : sssd-client-2.4.0-9.el8.x86_64                        598/654
  Verifying        : sssd-common-2.5.2-2.el8_5.3.x86_64                    599/654
  Verifying        : sssd-common-2.4.0-9.el8.x86_64                        600/654
  Verifying        : sssd-common-pac-2.5.2-2.el8_5.3.x86_64                601/654
  Verifying        : sssd-common-pac-2.4.0-9.el8.x86_64                    602/654
  Verifying        : sssd-ipa-2.5.2-2.el8_5.3.x86_64                       603/654
  Verifying        : sssd-ipa-2.4.0-9.el8.x86_64                           604/654
  Verifying        : sssd-kcm-2.5.2-2.el8_5.3.x86_64                       605/654
  Verifying        : sssd-kcm-2.4.0-9.el8.x86_64                           606/654
  Verifying        : sssd-krb5-2.5.2-2.el8_5.3.x86_64                      607/654
  Verifying        : sssd-krb5-2.4.0-9.el8.x86_64                          608/654
  Verifying        : sssd-krb5-common-2.5.2-2.el8_5.3.x86_64               609/654
  Verifying        : sssd-krb5-common-2.4.0-9.el8.x86_64                   610/654
  Verifying        : sssd-ldap-2.5.2-2.el8_5.3.x86_64                      611/654
  Verifying        : sssd-ldap-2.4.0-9.el8.x86_64                          612/654
  Verifying        : sssd-nfs-idmap-2.5.2-2.el8_5.3.x86_64                 613/654
  Verifying        : sssd-nfs-idmap-2.4.0-9.el8.x86_64                     614/654
  Verifying        : sssd-proxy-2.5.2-2.el8_5.3.x86_64                     615/654
  Verifying        : sssd-proxy-2.4.0-9.el8.x86_64                         616/654
  Verifying        : strace-5.7-3.el8.x86_64                               617/654
  Verifying        : strace-5.7-2.el8.x86_64                               618/654
  Verifying        : sudo-1.8.29-7.el8_4.1.x86_64                          619/654
  Verifying        : sudo-1.8.29-7.el8.x86_64                              620/654
  Verifying        : systemd-239-51.el8_5.2.x86_64                         621/654
  Verifying        : systemd-239-45.el8.x86_64                             622/654
  Verifying        : systemd-libs-239-51.el8_5.2.x86_64                    623/654
  Verifying        : systemd-libs-239-45.el8.x86_64                        624/654
  Verifying        : systemd-pam-239-51.el8_5.2.x86_64                     625/654
  Verifying        : systemd-pam-239-45.el8.x86_64                         626/654
  Verifying        : systemd-udev-239-51.el8_5.2.x86_64                    627/654
  Verifying        : systemd-udev-239-45.el8.x86_64                        628/654
  Verifying        : tpm2-tools-4.1.1-5.el8.x86_64                         629/654
  Verifying        : tpm2-tools-4.1.1-2.el8.x86_64                         630/654
  Verifying        : tpm2-tss-2.3.2-4.el8.x86_64                           631/654
  Verifying        : tpm2-tss-2.3.2-3.el8.x86_64                           632/654
  Verifying        : tuned-2.16.0-1.el8.noarch                             633/654
  Verifying        : tuned-2.15.0-2.el8.noarch                             634/654
  Verifying        : tzdata-2021e-1.el8.noarch                             635/654
  Verifying        : tzdata-2021a-1.el8.noarch                             636/654
  Verifying        : unzip-6.0-45.el8_4.x86_64                             637/654
  Verifying        : unzip-6.0-44.el8.x86_64                               638/654
  Verifying        : util-linux-2.32.1-28.el8.x86_64                       639/654
  Verifying        : util-linux-2.32.1-27.el8.x86_64                       640/654
  Verifying        : util-linux-user-2.32.1-28.el8.x86_64                  641/654
  Verifying        : util-linux-user-2.32.1-27.el8.x86_64                  642/654
  Verifying        : vdo-6.2.5.74-14.el8.x86_64                            643/654
  Verifying        : vdo-6.2.4.14-14.el8.x86_64                            644/654
  Verifying        : vim-minimal-2:8.0.1763-16.el8.x86_64                  645/654
  Verifying        : vim-minimal-2:8.0.1763-15.el8.x86_64                  646/654
  Verifying        : virt-what-1.18-12.el8.x86_64                          647/654
  Verifying        : virt-what-1.18-6.el8.x86_64                           648/654
  Verifying        : which-2.21-16.el8.x86_64                              649/654
  Verifying        : which-2.21-12.el8.x86_64                              650/654
  Verifying        : xfsprogs-5.0.0-9.el8.x86_64                           651/654
  Verifying        : xfsprogs-5.0.0-8.el8.x86_64                           652/654
  Verifying        : yum-4.7.0-4.el8.noarch                                653/654
  Verifying        : yum-4.4.2-11.el8.noarch                               654/654

Upgraded:
  NetworkManager-1:1.32.10-4.el8.x86_64
  NetworkManager-config-server-1:1.32.10-4.el8.noarch
  NetworkManager-libnm-1:1.32.10-4.el8.x86_64
  NetworkManager-team-1:1.32.10-4.el8.x86_64
  NetworkManager-tui-1:1.32.10-4.el8.x86_64
  adcli-0.8.2-12.el8.x86_64
  alsa-sof-firmware-1.8-1.el8.noarch
  authselect-1.2.2-3.el8.x86_64
  authselect-compat-1.2.2-3.el8.x86_64
  authselect-libs-1.2.2-3.el8.x86_64
  bash-4.4.20-2.el8.x86_64
  bind-export-libs-32:9.11.26-6.el8.x86_64
  bind-libs-32:9.11.26-6.el8.x86_64
  bind-libs-lite-32:9.11.26-6.el8.x86_64
  bind-license-32:9.11.26-6.el8.noarch
  bind-utils-32:9.11.26-6.el8.x86_64
  binutils-2.30-108.el8_5.1.x86_64
  bpftool-4.18.0-348.7.1.el8_5.x86_64
  buildah-1.22.3-2.module_el8.5.0+911+f19012f9.x86_64
  ca-certificates-2021.2.50-80.0.el8_4.noarch
  centos-gpg-keys-1:8-3.el8.noarch
  centos-linux-release-8.5-1.2111.el8.noarch
  centos-linux-repos-8-3.el8.noarch
  centos-logos-85.8-2.el8.x86_64
  chkconfig-1.19.1-1.el8.x86_64
  chrony-4.1-1.el8.x86_64
  cockpit-251.1-1.el8.x86_64
  cockpit-bridge-251.1-1.el8.x86_64
  cockpit-packagekit-251.1-1.el8.noarch
  cockpit-podman-33-1.module_el8.5.0+890+6b136101.noarch
  cockpit-storaged-251.1-1.el8.noarch
  cockpit-system-251.1-1.el8.noarch
  cockpit-ws-251.1-1.el8.x86_64
  conmon-2:2.0.29-1.module_el8.5.0+890+6b136101.x86_64
  container-selinux-2:2.167.0-1.module_el8.5.0+911+f19012f9.noarch
  containernetworking-plugins-1.0.0-1.module_el8.5.0+890+6b136101.x86_64
  containers-common-2:1-2.module_el8.5.0+890+6b136101.noarch
  coreutils-8.30-12.el8.x86_64
  coreutils-common-8.30-12.el8.x86_64
  criu-3.15-3.module_el8.5.0+890+6b136101.x86_64
  crypto-policies-20210617-1.gitc776d3e.el8.noarch
  crypto-policies-scripts-20210617-1.gitc776d3e.el8.noarch
  cups-libs-1:2.2.6-40.el8.x86_64
  curl-7.61.1-22.el8.x86_64
  dbus-1:1.12.8-14.el8.x86_64
  dbus-common-1:1.12.8-14.el8.noarch
  dbus-daemon-1:1.12.8-14.el8.x86_64
  dbus-libs-1:1.12.8-14.el8.x86_64
  dbus-tools-1:1.12.8-14.el8.x86_64
  device-mapper-8:1.02.177-10.el8.x86_64
  device-mapper-event-8:1.02.177-10.el8.x86_64
  device-mapper-event-libs-8:1.02.177-10.el8.x86_64
  device-mapper-libs-8:1.02.177-10.el8.x86_64
  device-mapper-multipath-0.8.4-17.el8.x86_64
  device-mapper-multipath-libs-0.8.4-17.el8.x86_64
  device-mapper-persistent-data-0.9.0-4.el8.x86_64
  dhcp-client-12:4.3.6-45.el8.x86_64
  dhcp-common-12:4.3.6-45.el8.noarch
  dhcp-libs-12:4.3.6-45.el8.x86_64
  dmidecode-1:3.2-10.el8.x86_64
  dnf-4.7.0-4.el8.noarch
  dnf-data-4.7.0-4.el8.noarch
  dnf-plugins-core-4.0.21-3.el8.noarch
  dracut-049-191.git20210920.el8.x86_64
  dracut-config-rescue-049-191.git20210920.el8.x86_64
  dracut-network-049-191.git20210920.el8.x86_64
  dracut-squash-049-191.git20210920.el8.x86_64
  e2fsprogs-1.45.6-2.el8.x86_64
  e2fsprogs-libs-1.45.6-2.el8.x86_64
  elfutils-debuginfod-client-0.185-1.el8.x86_64
  elfutils-default-yama-scope-0.185-1.el8.noarch
  elfutils-libelf-0.185-1.el8.x86_64
  elfutils-libs-0.185-1.el8.x86_64
  emacs-filesystem-1:26.1-7.el8.noarch
  ethtool-2:5.8-7.el8.x86_64
  file-5.33-20.el8.x86_64
  file-libs-5.33-20.el8.x86_64
  filesystem-3.8-6.el8.x86_64
  firewalld-0.9.3-7.el8.noarch
  firewalld-filesystem-0.9.3-7.el8.noarch
  fontconfig-2.13.1-4.el8.x86_64
  fstrm-0.6.1-2.el8.x86_64
  fuse-overlayfs-1.7.1-1.module_el8.5.0+890+6b136101.x86_64
  glib2-2.56.4-156.el8.x86_64
  glibc-2.28-164.el8.x86_64
  glibc-common-2.28-164.el8.x86_64
  glibc-langpack-en-2.28-164.el8.x86_64
  gnutls-3.6.16-4.el8.x86_64
  gpgme-1.13.1-9.el8.x86_64
  grub2-common-1:2.02-106.el8.noarch
  grub2-pc-1:2.02-106.el8.x86_64
  grub2-pc-modules-1:2.02-106.el8.noarch
  grub2-tools-1:2.02-106.el8.x86_64
  grub2-tools-extra-1:2.02-106.el8.x86_64
  grub2-tools-minimal-1:2.02-106.el8.x86_64
  grubby-8.40-42.el8.x86_64
  gsettings-desktop-schemas-3.32.0-6.el8.x86_64
  hdparm-9.54-4.el8.x86_64
  hwdata-0.314-8.10.el8.noarch
  iproute-5.12.0-4.el8.x86_64
  iptables-1.8.4-20.el8.x86_64
  iptables-ebtables-1.8.4-20.el8.x86_64
  iptables-libs-1.8.4-20.el8.x86_64
  iscsi-initiator-utils-6.2.1.4-4.git095f59c.el8.x86_64
  iscsi-initiator-utils-iscsiuio-6.2.1.4-4.git095f59c.el8.x86_64
  iwl100-firmware-39.31.5.1-103.el8.1.noarch
  iwl1000-firmware-1:39.31.5.1-103.el8.1.noarch
  iwl105-firmware-18.168.6.1-103.el8.1.noarch
  iwl135-firmware-18.168.6.1-103.el8.1.noarch
  iwl2000-firmware-18.168.6.1-103.el8.1.noarch
  iwl2030-firmware-18.168.6.1-103.el8.1.noarch
  iwl3160-firmware-1:25.30.13.0-103.el8.1.noarch
  iwl5000-firmware-8.83.5.1_1-103.el8.1.noarch
  iwl5150-firmware-8.24.2.2-103.el8.1.noarch
  iwl6000-firmware-9.221.4.1-103.el8.1.noarch
  iwl6000g2a-firmware-18.168.6.1-103.el8.1.noarch
  iwl6000g2b-firmware-18.168.6.1-103.el8.1.noarch
  iwl6050-firmware-41.28.5.1-103.el8.1.noarch
  iwl7260-firmware-1:25.30.13.0-103.el8.1.noarch
  json-c-0.13.1-2.el8.x86_64
  kernel-tools-4.18.0-348.7.1.el8_5.x86_64
  kernel-tools-libs-4.18.0-348.7.1.el8_5.x86_64
  kexec-tools-2.0.20-57.el8_5.1.x86_64
  keyutils-libs-1.5.10-9.el8.x86_64
  kmod-25-18.el8.x86_64
  kmod-kvdo-6.2.5.72-81.el8.x86_64
  kmod-libs-25-18.el8.x86_64
  kpartx-0.8.4-17.el8.x86_64
  kpatch-0.9.2-5.el8.noarch
  kpatch-dnf-0.2-5.el8.noarch
  krb5-libs-1.18.2-14.el8.x86_64
  libX11-1.6.8-5.el8.x86_64
  libX11-common-1.6.8-5.el8.noarch
  libblkid-2.32.1-28.el8.x86_64
  libblockdev-2.24-7.el8.x86_64
  libblockdev-crypto-2.24-7.el8.x86_64
  libblockdev-fs-2.24-7.el8.x86_64
  libblockdev-loop-2.24-7.el8.x86_64
  libblockdev-lvm-2.24-7.el8.x86_64
  libblockdev-mdraid-2.24-7.el8.x86_64
  libblockdev-part-2.24-7.el8.x86_64
  libblockdev-swap-2.24-7.el8.x86_64
  libblockdev-utils-2.24-7.el8.x86_64
  libcap-2.26-5.el8.x86_64
  libcap-ng-0.7.11-1.el8.x86_64
  libcom_err-1.45.6-2.el8.x86_64
  libcomps-0.1.16-2.el8.x86_64
  libcurl-7.61.1-22.el8.x86_64
  libdb-5.3.28-42.el8_4.x86_64
  libdb-utils-5.3.28-42.el8_4.x86_64
  libdnf-0.63.0-3.el8.x86_64
  libdrm-2.4.106-2.el8.x86_64
  libertas-usb8388-firmware-2:20210702-103.gitd79c2677.el8.noarch
  libfastjson-0.99.9-1.el8.x86_64
  libfdisk-2.32.1-28.el8.x86_64
  libgcc-8.5.0-4.el8_5.x86_64
  libgcrypt-1.8.5-6.el8.x86_64
  libgomp-8.5.0-4.el8_5.x86_64
  libibverbs-35.0-1.el8.x86_64
  libipa_hbac-2.5.2-2.el8_5.3.x86_64
  libldb-2.3.0-2.el8.x86_64
  libmodulemd-2.13.0-1.el8.x86_64
  libmount-2.32.1-28.el8.x86_64
  libndp-1.7-6.el8.x86_64
  libnfsidmap-1:2.3.3-46.el8.x86_64
  librelp-1.9.0-1.el8.x86_64
  librepo-1.14.0-2.el8.x86_64
  libsepol-2.9-3.el8.x86_64
  libslirp-4.4.0-1.module_el8.5.0+890+6b136101.x86_64
  libsmartcols-2.32.1-28.el8.x86_64
  libsmbclient-4.14.5-7.el8_5.x86_64
  libsolv-0.7.19-1.el8.x86_64
  libss-1.45.6-2.el8.x86_64
  libssh-0.9.4-3.el8.x86_64
  libssh-config-0.9.4-3.el8.noarch
  libsss_autofs-2.5.2-2.el8_5.3.x86_64
  libsss_certmap-2.5.2-2.el8_5.3.x86_64
  libsss_idmap-2.5.2-2.el8_5.3.x86_64
  libsss_nss_idmap-2.5.2-2.el8_5.3.x86_64
  libsss_sudo-2.5.2-2.el8_5.3.x86_64
  libstdc++-8.5.0-4.el8_5.x86_64
  libstoragemgmt-1.9.1-1.el8.x86_64
  libtalloc-2.3.2-1.el8.x86_64
  libtevent-0.11.0-0.el8.x86_64
  libtirpc-1.1.4-5.el8.x86_64
  libudisks2-2.9.0-7.el8.x86_64
  libuuid-2.32.1-28.el8.x86_64
  libwbclient-4.14.5-7.el8_5.x86_64
  libxcrypt-4.1.1-6.el8.x86_64
  libxml2-2.9.7-9.el8_4.2.x86_64
  linux-firmware-20210702-103.gitd79c2677.el8.noarch
  lshw-B.02.19.2-6.el8.x86_64
  lsscsi-0.32-3.el8.x86_64
  lua-libs-5.3.4-12.el8.x86_64
  lvm2-8:2.03.12-10.el8.x86_64
  lvm2-libs-8:2.03.12-10.el8.x86_64
  lz4-libs-1.8.3-3.el8_4.x86_64
  man-db-2.7.6.1-18.el8.x86_64
  man-pages-overrides-8.5.0.1-1.el8.noarch
  mcelog-3:175-1.el8.x86_64
  mdadm-4.2-rc2.el8.x86_64
  microcode_ctl-4:20210608-1.el8.x86_64
  ncurses-6.1-9.20180224.el8.x86_64
  ncurses-base-6.1-9.20180224.el8.noarch
  ncurses-libs-6.1-9.20180224.el8.x86_64
  nettle-3.4.1-7.el8.x86_64
  nftables-1:0.9.3-21.el8.x86_64
  nmap-ncat-2:7.70-6.el8.x86_64
  nspr-4.32.0-1.el8_4.x86_64
  nss-3.67.0-7.el8_5.x86_64
  nss-softokn-3.67.0-7.el8_5.x86_64
  nss-softokn-freebl-3.67.0-7.el8_5.x86_64
  nss-sysinit-3.67.0-7.el8_5.x86_64
  nss-util-3.67.0-7.el8_5.x86_64
  numactl-libs-2.0.12-13.el8.x86_64
  nvme-cli-1.14-3.el8.x86_64
  open-vm-tools-11.2.5-2.el8.x86_64
  openldap-2.4.46-18.el8.x86_64
  openssh-8.0p1-10.el8.x86_64
  openssh-clients-8.0p1-10.el8.x86_64
  openssh-server-8.0p1-10.el8.x86_64
  openssl-1:1.1.1k-5.el8_5.x86_64
  openssl-libs-1:1.1.1k-5.el8_5.x86_64
  os-prober-1.74-9.el8.x86_64
  pam-1.3.1-15.el8.x86_64
  parted-3.2-39.el8.x86_64
  pcre-8.42-6.el8.x86_64
  platform-python-3.6.8-41.el8.x86_64
  platform-python-pip-9.0.3-20.el8.noarch
  plymouth-0.9.4-10.20200615git1e36e30.el8.x86_64
  plymouth-core-libs-0.9.4-10.20200615git1e36e30.el8.x86_64
  plymouth-scripts-0.9.4-10.20200615git1e36e30.el8.x86_64
  podman-3.3.1-9.module_el8.5.0+988+b1f0b741.x86_64
  podman-catatonit-3.3.1-9.module_el8.5.0+988+b1f0b741.x86_64
  policycoreutils-2.9-16.el8.x86_64
  policycoreutils-python-utils-2.9-16.el8.noarch
  polkit-0.115-12.el8.x86_64
  polkit-libs-0.115-12.el8.x86_64
  python3-bind-32:9.11.26-6.el8.noarch
  python3-dnf-4.7.0-4.el8.noarch
  python3-dnf-plugins-core-4.0.21-3.el8.noarch
  python3-firewall-0.9.3-7.el8.noarch
  python3-gpg-1.13.1-9.el8.x86_64
  python3-hawkey-0.63.0-3.el8.x86_64
  python3-libcomps-0.1.16-2.el8.x86_64
  python3-libdnf-0.63.0-3.el8.x86_64
  python3-libs-3.6.8-41.el8.x86_64
  python3-libstoragemgmt-1.9.1-1.el8.x86_64
  python3-libxml2-2.9.7-9.el8_4.2.x86_64
  python3-lxml-4.2.3-3.el8.x86_64
  python3-nftables-1:0.9.3-21.el8.x86_64
  python3-perf-4.18.0-348.7.1.el8_5.x86_64
  python3-pip-wheel-9.0.3-20.el8.noarch
  python3-policycoreutils-2.9-16.el8.noarch
  python3-psutil-5.4.3-11.el8.x86_64
  python3-rpm-4.14.3-19.el8.x86_64
  python3-sssdconfig-2.5.2-2.el8_5.3.noarch
  python3-syspurpose-1.28.21-3.el8.x86_64
  python3-unbound-1.7.3-17.el8.x86_64
  quota-1:4.04-14.el8.x86_64
  quota-nls-1:4.04-14.el8.noarch
  rdma-core-35.0-1.el8.x86_64
  realmd-0.16.3-23.el8.x86_64
  rpm-4.14.3-19.el8.x86_64
  rpm-build-libs-4.14.3-19.el8.x86_64
  rpm-libs-4.14.3-19.el8.x86_64
  rpm-plugin-selinux-4.14.3-19.el8.x86_64
  rpm-plugin-systemd-inhibit-4.14.3-19.el8.x86_64
  rsyslog-8.2102.0-5.el8.x86_64
  rsyslog-gnutls-8.2102.0-5.el8.x86_64
  rsyslog-gssapi-8.2102.0-5.el8.x86_64
  rsyslog-relp-8.2102.0-5.el8.x86_64
  runc-1.0.2-1.module_el8.5.0+911+f19012f9.x86_64
  samba-client-libs-4.14.5-7.el8_5.x86_64
  samba-common-4.14.5-7.el8_5.noarch
  samba-common-libs-4.14.5-7.el8_5.x86_64
  selinux-policy-3.14.3-80.el8_5.2.noarch
  selinux-policy-targeted-3.14.3-80.el8_5.2.noarch
  setroubleshoot-plugins-3.3.14-1.el8.noarch
  setroubleshoot-server-3.3.24-4.el8.x86_64
  shadow-utils-2:4.6-14.el8.x86_64
  slirp4netns-1.1.8-1.module_el8.5.0+890+6b136101.x86_64
  sos-4.1-5.el8.noarch
  sqlite-3.26.0-15.el8.x86_64
  sqlite-libs-3.26.0-15.el8.x86_64
  sssd-2.5.2-2.el8_5.3.x86_64
  sssd-ad-2.5.2-2.el8_5.3.x86_64
  sssd-client-2.5.2-2.el8_5.3.x86_64
  sssd-common-2.5.2-2.el8_5.3.x86_64
  sssd-common-pac-2.5.2-2.el8_5.3.x86_64
  sssd-ipa-2.5.2-2.el8_5.3.x86_64
  sssd-kcm-2.5.2-2.el8_5.3.x86_64
  sssd-krb5-2.5.2-2.el8_5.3.x86_64
  sssd-krb5-common-2.5.2-2.el8_5.3.x86_64
  sssd-ldap-2.5.2-2.el8_5.3.x86_64
  sssd-nfs-idmap-2.5.2-2.el8_5.3.x86_64
  sssd-proxy-2.5.2-2.el8_5.3.x86_64
  strace-5.7-3.el8.x86_64
  sudo-1.8.29-7.el8_4.1.x86_64
  systemd-239-51.el8_5.2.x86_64
  systemd-libs-239-51.el8_5.2.x86_64
  systemd-pam-239-51.el8_5.2.x86_64
  systemd-udev-239-51.el8_5.2.x86_64
  tcpdump-14:4.9.3-2.el8.x86_64
  tpm2-tools-4.1.1-5.el8.x86_64
  tpm2-tss-2.3.2-4.el8.x86_64
  tuned-2.16.0-1.el8.noarch
  tzdata-2021e-1.el8.noarch
  udisks2-2.9.0-7.el8.x86_64
  udisks2-iscsi-2.9.0-7.el8.x86_64
  udisks2-lvm2-2.9.0-7.el8.x86_64
  unbound-libs-1.7.3-17.el8.x86_64
  unzip-6.0-45.el8_4.x86_64
  util-linux-2.32.1-28.el8.x86_64
  util-linux-user-2.32.1-28.el8.x86_64
  vdo-6.2.5.74-14.el8.x86_64
  vim-common-2:8.0.1763-16.el8.x86_64
  vim-enhanced-2:8.0.1763-16.el8.x86_64
  vim-filesystem-2:8.0.1763-16.el8.noarch
  vim-minimal-2:8.0.1763-16.el8.x86_64
  virt-what-1.18-12.el8.x86_64
  which-2.21-16.el8.x86_64
  xfsprogs-5.0.0-9.el8.x86_64
  yum-4.7.0-4.el8.noarch
Installed:
  grub2-tools-efi-1:2.02-106.el8.x86_64
  kernel-4.18.0-348.7.1.el8_5.x86_64
  kernel-core-4.18.0-348.7.1.el8_5.x86_64
  kernel-modules-4.18.0-348.7.1.el8_5.x86_64
  libbpf-0.4.0-1.el8.x86_64

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
httpd-devel.x86_64 : Development interfaces for the Apache HTTP server
httpd-filesystem.noarch : The basic directory layout for the Apache HTTP server
httpd-manual.noarch : Documentation for the Apache HTTP server
httpd-tools.x86_64 : Tools for use with the Apache HTTP Server
keycloak-httpd-client-install.noarch : Tools to configure Apache HTTPD as Keycloak
                                     : client
librdkafka.i686 : The Apache Kafka C library
librdkafka.x86_64 : The Apache Kafka C library
mod_auth_gssapi.x86_64 : A GSSAPI Authentication module for Apache
mod_auth_mellon.x86_64 : A SAML 2.0 authentication module for the Apache Httpd
                       : Server
mod_dav_svn.x86_64 : Apache httpd module for Subversion server
mod_fcgid.x86_64 : FastCGI interface module for Apache 2
mod_http2.x86_64 : module implementing HTTP/2 for Apache 2
mod_intercept_form_submit.x86_64 : Apache module to intercept login form submission
                                 : and run PAM authentication
mod_ldap.x86_64 : LDAP authentication modules for the Apache HTTP Server
mod_lookup_identity.x86_64 : Apache module to retrieve additional information about
                           : the authenticated user
mod_md.x86_64 : Certificate provisioning using ACME for the Apache HTTP Server
mod_proxy_html.x86_64 : HTML and XML content filters for the Apache HTTP Server
mod_security.x86_64 : Security module for the Apache HTTP Server
mod_session.x86_64 : Session interface for the Apache HTTP Server
mod_ssl.x86_64 : SSL/TLS module for the Apache HTTP Server
pcp-export-pcp2spark.x86_64 : Performance Co-Pilot tools for exporting PCP metrics
                            : to Apache Spark
python3-keycloak-httpd-client-install.noarch : Tools to configure Apache HTTPD as
                                             : Keycloak client
python3-mod_wsgi.x86_64 : A WSGI interface for Python web applications in Apache
python38-mod_wsgi.x86_64 : A WSGI interface for Python web applications in Apache
python39-mod_wsgi.x86_64 : A WSGI interface for Python web applications in Apache
[osboxes@osboxes ~]$ sudo dnf install apache http server
Last metadata expiration check: 7:50:30 ago on Tue Feb  8 01:51:14 2022.
No match for argument: apache
No match for argument: http
No match for argument: server
Error: Unable to find a match: apache http server
[osboxes@osboxes ~]$ sudo dnf install httpd
Last metadata expiration check: 7:50:50 ago on Tue Feb  8 01:51:14 2022.
Dependencies resolved.
===================================================================================
 Package            Arch   Version                                 Repo       Size
===================================================================================
Installing:
 httpd              x86_64 2.4.37-43.module_el8.5.0+1022+b541f3b1  appstream 1.4 M
Installing dependencies:
 apr                x86_64 1.6.3-12.el8                            appstream 129 k
 apr-util           x86_64 1.6.1-6.el8                             appstream 105 k
 centos-logos-httpd noarch 85.8-2.el8                              baseos     75 k
 httpd-filesystem   noarch 2.4.37-43.module_el8.5.0+1022+b541f3b1  appstream  39 k
 httpd-tools        x86_64 2.4.37-43.module_el8.5.0+1022+b541f3b1  appstream 107 k
 mod_http2          x86_64 1.15.7-3.module_el8.4.0+778+c970deab    appstream 154 k
Installing weak dependencies:
 apr-util-bdb       x86_64 1.6.1-6.el8                             appstream  25 k
 apr-util-openssl   x86_64 1.6.1-6.el8                             appstream  27 k
Enabling module streams:
 httpd                     2.4

Transaction Summary
===================================================================================
Install  9 Packages

Total download size: 2.1 M
Installed size: 5.6 M
Is this ok [y/N]: y
Downloading Packages:
(1/9): apr-util-bdb-1.6.1-6.el8.x86_64.rpm          11 kB/s |  25 kB     00:02
(2/9): apr-util-1.6.1-6.el8.x86_64.rpm              37 kB/s | 105 kB     00:02
(3/9): apr-util-openssl-1.6.1-6.el8.x86_64.rpm      43 kB/s |  27 kB     00:00
(4/9): apr-1.6.3-12.el8.x86_64.rpm                  46 kB/s | 129 kB     00:02
(5/9): httpd-filesystem-2.4.37-43.module_el8.5.0+1  21 kB/s |  39 kB     00:01
(6/9): httpd-2.4.37-43.module_el8.5.0+1022+b541f3b 568 kB/s | 1.4 MB     00:02
(7/9): mod_http2-1.15.7-3.module_el8.4.0+778+c970d 185 kB/s | 154 kB     00:00
(8/9): centos-logos-httpd-85.8-2.el8.noarch.rpm    135 kB/s |  75 kB     00:00
(9/9): httpd-tools-2.4.37-43.module_el8.5.0+1022+b  34 kB/s | 107 kB     00:03
-----------------------------------------------------------------------------------
Total                                              349 kB/s | 2.1 MB     00:06
Running transaction check
Transaction check succeeded.
Running transaction test
Transaction test succeeded.
Running transaction
  Preparing        :                                                           1/1
  Installing       : apr-1.6.3-12.el8.x86_64                                   1/9
  Running scriptlet: apr-1.6.3-12.el8.x86_64                                   1/9
  Installing       : apr-util-bdb-1.6.1-6.el8.x86_64                           2/9
  Installing       : apr-util-openssl-1.6.1-6.el8.x86_64                       3/9
  Installing       : apr-util-1.6.1-6.el8.x86_64                               4/9
  Running scriptlet: apr-util-1.6.1-6.el8.x86_64                               4/9
  Installing       : httpd-tools-2.4.37-43.module_el8.5.0+1022+b541f3b1.x86_   5/9
  Installing       : centos-logos-httpd-85.8-2.el8.noarch                      6/9
  Running scriptlet: httpd-filesystem-2.4.37-43.module_el8.5.0+1022+b541f3b1   7/9
  Installing       : httpd-filesystem-2.4.37-43.module_el8.5.0+1022+b541f3b1   7/9
  Installing       : mod_http2-1.15.7-3.module_el8.4.0+778+c970deab.x86_64     8/9
  Installing       : httpd-2.4.37-43.module_el8.5.0+1022+b541f3b1.x86_64       9/9
  Running scriptlet: httpd-2.4.37-43.module_el8.5.0+1022+b541f3b1.x86_64       9/9
  Verifying        : apr-1.6.3-12.el8.x86_64                                   1/9
  Verifying        : apr-util-1.6.1-6.el8.x86_64                               2/9
  Verifying        : apr-util-bdb-1.6.1-6.el8.x86_64                           3/9
  Verifying        : apr-util-openssl-1.6.1-6.el8.x86_64                       4/9
  Verifying        : httpd-2.4.37-43.module_el8.5.0+1022+b541f3b1.x86_64       5/9
  Verifying        : httpd-filesystem-2.4.37-43.module_el8.5.0+1022+b541f3b1   6/9
  Verifying        : httpd-tools-2.4.37-43.module_el8.5.0+1022+b541f3b1.x86_   7/9
  Verifying        : mod_http2-1.15.7-3.module_el8.4.0+778+c970deab.x86_64     8/9
  Verifying        : centos-logos-httpd-85.8-2.el8.noarch                      9/9

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
drwxr-xr-x. 105 root root     8192 Feb  8 09:42 .
dr-xr-xr-x.  17 root root      224 Feb  8 03:40 ..
-rw-------.   1 root root        0 Jun 19  2021 .pwd.lock
-rw-r--r--.   1 root root      208 Jun 19  2021 .updated
-rw-r--r--.   1 root root     4536 Jul 14  2021 DIR_COLORS
-rw-r--r--.   1 root root     5214 Jul 14  2021 DIR_COLORS.256color
-rw-r--r--.   1 root root     4618 Jul 14  2021 DIR_COLORS.lightbgcolor
-rw-r--r--.   1 root root       94 May 10  2019 GREP_COLORS
drwxr-xr-x.   7 root root      134 Feb  8 03:46 NetworkManager
drwxr-xr-x.   2 root root       48 Jun 19  2021 PackageKit
drwxr-xr-x.   6 root root       70 Jun 22  2021 X11
-rw-r--r--.   1 root root       16 Jun 19  2021 adjtime
-rw-r--r--.   1 root root     1529 May 15  2020 aliases
drwxr-xr-x.   2 root root      194 Jul 28  2021 alternatives
-rw-r--r--.   1 root root      541 Nov  8  2019 anacrontab
-rw-r--r--.   1 root root        1 May 11  2019 at.deny
drwxr-x---.   4 root root      100 Jun 19  2021 audit
drwxr-xr-x.   3 root root      228 Feb  8 09:35 authselect
drwxr-xr-x.   2 root root      135 Feb  8 03:49 bash_completion.d
-rw-r--r--.   1 root root     3019 May 15  2020 bashrc
-rw-r--r--.   1 root root      535 Apr 21  2021 bindresvport.blacklist
drwxr-xr-x.   2 root root        6 Dec 21 15:08 binfmt.d
-rw-r--r--.   1 root root       30 Nov  9 19:17 centos-release
-rw-r--r--.   1 root root       42 Nov  9 19:17 centos-release-upstream
drwxr-xr-x.   2 root root        6 Jul 28  2021 chkconfig.d
-rw-r--r--.   1 root root     1085 Jun 24  2021 chrony.conf
-rw-r-----.   1 root chrony    540 May 12  2021 chrony.keys
drwxr-xr-x.   2 root root       26 Dec 21 15:14 cifs-utils
drwxr-xr-x.   3 root root       19 Jun 19  2021 cni
drwxr-xr-x.   4 root root       42 Sep 10 10:51 cockpit
drwxr-xr-x.   6 root root      139 Feb  8 03:48 containers
drwxr-xr-x.   2 root root       39 Jun 19  2021 cron.d
drwxr-xr-x.   2 root root       23 Jun 19  2021 cron.daily
-rw-r--r--.   1 root root        0 Nov  8  2019 cron.deny
drwxr-xr-x.   2 root root       22 Jun 19  2021 cron.hourly
drwxr-xr-x.   2 root root        6 Jan 12  2021 cron.monthly
drwxr-xr-x.   2 root root        6 Jan 12  2021 cron.weekly
-rw-r--r--.   1 root root      451 Jan 12  2021 crontab
drwxr-xr-x.   6 root root       81 Jun 19  2021 crypto-policies
-rw-------.   1 root root        0 Jun 19  2021 crypttab
-rw-r--r--.   1 root root     1629 May 15  2020 csh.cshrc
-rw-r--r--.   1 root root     1078 May 15  2020 csh.login
drwxr-xr-x.   4 root root       78 May  8  2021 dbus-1
drwxr-xr-x.   3 root root       16 Jun 19  2021 dconf
drwxr-xr-x.   2 root root       33 Feb  8 03:48 default
drwxr-xr-x.   2 root root       40 Feb  8 03:46 depmod.d
drwxr-x---.   3 root root       45 Jul 22  2021 dhcp
drwxr-xr-x.   8 root root      128 Sep 17 15:05 dnf
-rw-r--r--.   1 root root      117 Sep 29 17:07 dracut.conf
drwxr-xr-x.   2 root root        6 Sep 29 17:07 dracut.conf.d
-rw-r--r--.   1 root root        0 May 15  2020 environment
-rw-r--r--.   1 root root     1362 Aug 24 19:13 ethertypes
-rw-r--r--.   1 root root        0 Sep 10  2018 exports
lrwxrwxrwx.   1 root root       56 Nov 12 13:19 favicon.png -> /usr/share/icons/hicolor/16x16/apps/fedora-logo-icon.png
-rw-r--r--.   1 root root       66 Sep 10  2018 filesystems
drwxr-x---.   8 root root      149 Feb  8 03:48 firewalld
drwxr-xr-x.   3 root root       38 Jun 19  2021 fonts
-rw-r--r--.   1 root root       20 Jan 13  2021 fprintd.conf
-rw-r--r--.   1 root root      709 Jun 19  2021 fstab
-rw-r--r--.   1 root root       38 May 11  2019 fuse.conf
drwxr-xr-x.   2 root root       25 Jun 29  2021 gcrypt
drwxr-xr-x.   2 root root        6 May 15  2020 gnupg
drwxr-xr-x.   4 root root       40 Jun 19  2021 groff
-rw-r--r--.   1 root root      691 Feb  8 09:42 group
-rw-r--r--.   1 root root      678 Jun 19  2021 group-
drwx------.   2 root root      288 Feb  8 03:44 grub.d
lrwxrwxrwx.   1 root root       22 Nov  8 01:39 grub2.cfg -> ../boot/grub2/grub.cfg
----------.   1 root root      555 Feb  8 09:42 gshadow
----------.   1 root root      544 Jun 19  2021 gshadow-
drwxr-xr-x.   3 root root       20 Aug 26 01:12 gss
-rw-r--r--.   1 root root        9 Sep 10  2018 host.conf
-rw-r--r--.   1 root root        8 Feb  5 02:10 hostname
-rw-r--r--.   1 root root      158 Sep 10  2018 hosts
drwxr-xr-x.   5 root root      105 Feb  8 09:42 httpd
-rw-r--r--.   1 root root     4849 Jul 30  2021 idmapd.conf
lrwxrwxrwx.   1 root root       11 Jul 28  2021 init.d -> rc.d/init.d
-rw-r--r--.   1 root root      490 Dec 21 15:08 inittab
-rw-r--r--.   1 root root      942 Sep 10  2018 inputrc
drwxr-xr-x.   2 root root      159 Oct 15 13:08 iproute2
drwxr-xr-x.   2 root root       52 Aug 10 01:10 iscsi
-rw-r--r--.   1 root root       23 Nov  9 19:17 issue
drwxr-xr-x.   2 root root       27 Jun 19  2021 issue.d
-rw-r--r--.   1 root root       22 Nov  9 19:17 issue.net
drwxr-xr-x.   4 root root       33 Dec 21 15:15 kdump
-rw-r--r--.   1 root root     8550 Feb  8 03:47 kdump.conf
drwxr-xr-x.   4 root root       41 Dec 21 15:08 kernel
-rw-r--r--.   1 root root      812 Aug 26 01:05 krb5.conf
drwxr-xr-x.   2 root root       55 Feb  8 03:48 krb5.conf.d
-rw-r--r--.   1 root root    21275 Feb  8 09:42 ld.so.cache
-rw-r--r--.   1 root root       28 Aug 24 19:10 ld.so.conf
drwxr-xr-x.   2 root root      129 Feb  8 03:44 ld.so.conf.d
-rw-r-----.   1 root root      191 Nov  4  2019 libaudit.conf
drwxr-xr-x.   3 root root       20 Jul  2  2021 libblockdev
drwxr-xr-x.   2 root root      246 May 17  2021 libibverbs.d
drwxr-xr-x.   2 root root       35 Jun 19  2021 libnl
drwxr-xr-x.   6 root root       70 Jun 19  2021 libreport
drwxr-xr-x.   2 root root       62 May  5  2021 libssh
-rw-r--r--.   1 root root     2391 Jul 23  2015 libuser.conf
-rw-r--r--.   1 root root       19 Jun 19  2021 locale.conf
lrwxrwxrwx.   1 root root       38 Feb  5 02:29 localtime -> ../usr/share/zoneinfo/America/New_York
-rw-r--r--.   1 root root     2512 Aug 18 15:04 login.defs
-rw-r--r--.   1 root root      438 Feb 19  2018 logrotate.conf
drwxr-xr-x.   2 root root      188 Feb  8 09:42 logrotate.d
drwxr-xr-x.   3 root root       43 Apr 22  2021 lsm
drwxr-xr-x.   6 root root      100 Feb  8 03:46 lvm
-r--r--r--.   1 root root       33 Jun 19  2021 machine-id
-rw-r--r--.   1 root root      111 May  8  2021 magic
-rw-r--r--.   1 root root      272 May 11  2017 mailcap
-rw-r--r--.   1 root root     5122 Dec 21 15:15 makedumpfile.conf.sample
-rw-r--r--.   1 root root     5165 Jun 29  2021 man_db.conf
drwxr-xr-x.   2 root root      181 Feb  5 05:47 mc
drwxr-xr-x.   3 root root       41 Aug  9  2021 mcelog
drwxr-xr-x.   3 root root       32 Jul 24  2021 microcode_ctl
-rw-r--r--.   1 root root    60352 May 11  2017 mime.types
-rw-r--r--.   1 root root     1108 Jun 24  2021 mke2fs.conf
drwxr-xr-x.   2 root root       93 Feb  8 03:48 modprobe.d
drwxr-xr-x.   2 root root        6 Dec 21 15:08 modules-load.d
-rw-r--r--.   1 root root        0 Sep 10  2018 motd
drwxr-xr-x.   2 root root       21 Jun 19  2021 motd.d
lrwxrwxrwx.   1 root root       19 Jun 19  2021 mtab -> ../proc/self/mounts
drwxr-xr-x.   2 root root        6 Jul 28  2021 multipath
-rw-r--r--.   1 root root     9450 May 11  2019 nanorc
-rw-r--r--.   1 root root      767 Apr 21  2021 netconfig
-rw-r--r--.   1 root root       58 Sep 10  2018 networks
drwx------.   3 root root       66 Aug 24 19:25 nftables
lrwxrwxrwx.   1 root root       29 Feb  8 09:35 nsswitch.conf -> /etc/authselect/nsswitch.conf
-rw-r--r--.   1 root root     2197 Mar 11  2021 nsswitch.conf.bak
drwxr-xr-x.   2 root root       35 Sep 17 15:06 nvme
drwxr-xr-x.   2 root root        6 Dec 16  2020 oddjob
-rw-r--r--.   1 root root     4922 Dec 16  2020 oddjobd.conf
drwxr-xr-x.   2 root root       70 Jun 19  2021 oddjobd.conf.d
drwxr-xr-x.   3 root root       36 Aug 10 01:12 openldap
drwxr-xr-x.   2 root root        6 Jun 22  2021 opt
lrwxrwxrwx.   1 root root       21 Nov  9 19:17 os-release -> ../usr/lib/os-release
drwxr-xr-x.   2 root root     4096 Feb  8 09:35 pam.d
-rw-r--r--.   1 root root     1594 Feb  8 09:42 passwd
-rw-r--r--.   1 root root     1541 Jun 19  2021 passwd-
-rw-r--r--.   1 root root     2872 May 14  2019 pinforc
drwxr-xr-x.   3 root root       21 Jun 19  2021 pkcs11
drwxr-xr-x.   8 root root       88 Jun 22  2021 pki
drwxr-xr-x.   2 root root       28 Aug 24 19:28 plymouth
drwxr-xr-x.   5 root root       52 Jun 22  2021 pm
drwxr-xr-x.   5 root root       72 Jun  2  2021 polkit-1
drwxr-xr-x.   2 root root        6 Jan 19  2021 popt.d
drwxr-xr-x.   2 root root       24 Feb  8 03:43 prelink.conf.d
-rw-r--r--.   1 root root      233 Sep 10  2018 printcap
-rw-r--r--.   1 root root     2123 May 15  2020 profile
drwxr-xr-x.   2 root root     4096 Feb  8 03:41 profile.d
-rw-r--r--.   1 root root     6568 Sep 10  2018 protocols
drwxr-xr-x.  10 root root      127 Jul 28  2021 rc.d
lrwxrwxrwx.   1 root root       13 Dec 21 15:08 rc.local -> rc.d/rc.local
lrwxrwxrwx.   1 root root       10 Jul 28  2021 rc0.d -> rc.d/rc0.d
lrwxrwxrwx.   1 root root       10 Jul 28  2021 rc1.d -> rc.d/rc1.d
lrwxrwxrwx.   1 root root       10 Jul 28  2021 rc2.d -> rc.d/rc2.d
lrwxrwxrwx.   1 root root       10 Jul 28  2021 rc3.d -> rc.d/rc3.d
lrwxrwxrwx.   1 root root       10 Jul 28  2021 rc4.d -> rc.d/rc4.d
lrwxrwxrwx.   1 root root       10 Jul 28  2021 rc5.d -> rc.d/rc5.d
lrwxrwxrwx.   1 root root       10 Jul 28  2021 rc6.d -> rc.d/rc6.d
drwxr-xr-x.   3 root root       38 May 17  2021 rdma
lrwxrwxrwx.   1 root root       14 Nov  9 19:17 redhat-release -> centos-release
-rw-r--r--.   1 root root       54 Feb  7 22:08 resolv.conf
drwxr-xr-x.   3 root root       24 Jun 19  2021 rhsm
-rw-r--r--.   1 root root     1634 Aug  1  2018 rpc
drwxr-xr-x.   2 root root       25 Feb  8 03:48 rpm
-rw-r--r--.   1 root root     3186 Aug 10 12:43 rsyslog.conf
drwxr-xr-x.   2 root root        6 Aug 10 12:46 rsyslog.d
drwxr-xr-x.   2 root root       35 Dec 21 15:14 rwtab.d
drwxr-xr-x.   2 root root       61 Feb  8 03:44 samba
drwxr-xr-x.   2 root root        6 May 15  2020 sasl2
drwxr-xr-x.   7 root root     4096 May  7  2021 security
drwxr-xr-x.   3 root root       57 Dec 21 15:17 selinux
-rw-r--r--.   1 root root   692252 May 15  2020 services
-rw-r--r--.   1 root root      216 Sep 30 01:08 sestatus.conf
drwxr-xr-x.   2 root root       33 Sep 30 01:08 setroubleshoot
----------.   1 root root      972 Feb  8 09:42 shadow
----------.   1 root root      950 Jun 19  2021 shadow-
-rw-r--r--.   1 root root       44 Sep 10  2018 shells
drwxr-xr-x.   2 root root       62 Jun 22  2021 skel
drwxr-xr-x.   3 root root       74 Jun 19  2021 smartmontools
drwxr-xr-x.   6 root root       86 Feb  8 03:46 sos
drwxr-xr-x.   3 root root      245 Jul 13  2021 ssh
drwxr-xr-x.   2 root root       19 Feb  8 03:43 ssl
drwx------.   4 sssd sssd       31 Dec 21 15:14 sssd
-rw-r--r--.   1 root root       21 Jun 19  2021 subgid
-rw-r--r--.   1 root root        0 Sep 10  2018 subgid-
-rw-r--r--.   1 root root       21 Jun 19  2021 subuid
-rw-r--r--.   1 root root        0 Sep 10  2018 subuid-
-rw-r-----.   1 root root     3181 Oct 25 10:27 sudo-ldap.conf
-rw-r-----.   1 root root     1786 Oct 25 10:27 sudo.conf
-r--r-----.   1 root root     4328 Oct 25 10:27 sudoers
drwxr-x---.   2 root root        6 Oct 25 10:29 sudoers.d
drwxr-xr-x.   5 root root     4096 Feb  8 09:42 sysconfig
-rw-r--r--.   1 root root      449 Dec 21 15:08 sysctl.conf
drwxr-xr-x.   2 root root       28 Dec 21 15:08 sysctl.d
lrwxrwxrwx.   1 root root       14 Nov  9 19:17 system-release -> centos-release
-rw-r--r--.   1 root root       23 Nov  9 19:17 system-release-cpe
drwxr-xr-x.   4 root root      150 Feb  8 03:44 systemd
-rw-r-----.   1 root tss      7046 Nov 16  2020 tcsd.conf
drwxr-xr-x.   2 root root        6 Jun  1  2021 terminfo
drwxr-xr-x.   2 root root        6 Dec 21 15:08 tmpfiles.d
-rw-r--r--.   1 root root      375 Aug 24 19:20 trusted-key.key
drwxr-xr-x.   3 root root      136 Feb  8 03:48 tuned
drwxr-xr-x.   4 root root       68 Feb  8 09:37 udev
drwxr-xr-x.   2 root root       60 Feb  8 03:47 udisks2
drwxr-xr-x.   2 root root       45 Feb  8 03:47 unbound
-rw-r--r--.   1 root root      587 May 11  2019 updatedb.conf
-rw-r--r--.   1 root root     1523 May 11  2019 usb_modeswitch.conf
-rw-r--r--.   1 root root       28 Jun 19  2021 vconsole.conf
-rw-r--r--.   1 root root     1982 Sep 22 07:10 vimrc
-rw-r--r--.   1 root root     1204 Sep 22 07:10 virc
drwxr-xr-x.   4 root root      208 Feb  8 03:48 vmware-tools
-rw-r--r--.   1 root root     4925 Apr 26  2020 wgetrc
-rw-r--r--.   1 root root      642 Dec  9  2016 xattr.conf
drwxr-xr-x.   4 root root       38 Jun 22  2021 xdg
drwxr-xr-x.   2 root root        6 Jun 22  2021 xinetd.d
drwxr-xr-x.   2 root root       57 Feb  8 03:48 yum
lrwxrwxrwx.   1 root root       12 Sep 17 15:05 yum.conf -> dnf/dnf.conf
drwxr-xr-x.   2 root root     4096 Feb  8 03:48 yum.repos.d
[root@osboxes etc]# cat httpd
cat: httpd: Is a directory
[root@osboxes etc]# cd httpd
[root@osboxes httpd]# ls -al
total 12
drwxr-xr-x.   5 root root  105 Feb  8 09:42 .
drwxr-xr-x. 105 root root 8192 Feb  8 09:42 ..
drwxr-xr-x.   2 root root   37 Feb  8 09:42 conf
drwxr-xr-x.   2 root root   82 Feb  8 09:42 conf.d
drwxr-xr-x.   2 root root  226 Feb  8 09:42 conf.modules.d
lrwxrwxrwx.   1 root root   19 Nov 11 23:58 logs -> ../../var/log/httpd
lrwxrwxrwx.   1 root root   29 Nov 11 23:58 modules -> ../../usr/lib64/httpd/modules
lrwxrwxrwx.   1 root root   10 Nov 11 23:58 run -> /run/httpd
lrwxrwxrwx.   1 root root   19 Nov 11 23:58 state -> ../../var/lib/httpd
[root@osboxes httpd]# cat conf.d
cat: conf.d: Is a directory
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
# UserDir: The name of the directory that is appended onto a users: File name too long
[root@osboxes httpd]# # directory if a ~user request is received.
[root@osboxes httpd]# #
[root@osboxes httpd]# # The path to the end user account 'public_html' directory must be
[root@osboxes httpd]# # accessible to the webserver userid.  This usually means that ~userid
[root@osboxes httpd]# # must have permissions of 711, ~userid/public_html must have permissions
[root@osboxes httpd]# # of 755, and documents contained therein must be world-readable.
[root@osboxes httpd]# # Otherwise, the client will only receive a "403 Forbidden" message.
[root@osboxes httpd]# #
[root@osboxes httpd]# <IfModule mod_userdir.c>
-bash: syntax error near unexpected token `newline'
[root@osboxes httpd]#     #
[root@osboxes httpd]#     # UserDir is disabled by default since it can confirm the presence
[root@osboxes httpd]#     # of a username on the system (depending on home directory
[root@osboxes httpd]#     # permissions).
[root@osboxes httpd]#     #
[root@osboxes httpd]#     UserDir disabled
-bash: UserDir: command not found
[root@osboxes httpd]#
[root@osboxes httpd]#     #
[root@osboxes httpd]#     # To enable requests to /~user/ to serve the user's public_html
[root@osboxes httpd]#     # directory, remove the "UserDir disabled" line above, and uncomment
[root@osboxes httpd]#     # the following line instead:
[root@osboxes httpd]#     #
[root@osboxes httpd]#     #UserDir public_html
[root@osboxes httpd]# </IfModule>
-bash: syntax error near unexpected token `newline'
[root@osboxes httpd]#
[root@osboxes httpd]# #
[root@osboxes httpd]# # Control access to UserDir directories.  The following is an example
[root@osboxes httpd]# # for a site where these directories are restricted to read-only.
[root@osboxes httpd]# #
[root@osboxes httpd]# <Directory "/home/*/public_html">
-bash: syntax error near unexpected token `newline'
[root@osboxes httpd]#     AllowOverride FileInfo AuthConfig Limit Indexes
-bash: AllowOverride: command not found
[root@osboxes httpd]#     Options MultiViews Indexes SymLinksIfOwnerMatch IncludesNoExec
-bash: Options: command not found
[root@osboxes httpd]#     Require method GET POST OPTIONS
-bash: Require: command not found
[root@osboxes httpd]# </Directory>
-bash: syntax error near unexpected token `newline'
[root@osboxes httpd]#
[root@osboxes httpd]# [root@osboxes conf.d]# cd /var/log/httpd
-bash: [root@osboxes: command not found
[root@osboxes httpd]# [root@osboxes httpd]# ls -al
-bash: [root@osboxes: command not found
[root@osboxes httpd]# total 8
-bash: total: command not found
[root@osboxes httpd]# drwx------.  2 root root   41 Feb  8 09:53 .
-bash: drwx------.: command not found
[root@osboxes httpd]# drwxr-xr-x. 10 root root 4096 Feb  8 09:42 ..
-bash: drwxr-xr-x.: command not found
[root@osboxes httpd]# -rw-r--r--.  1 root root    0 Feb  8 09:53 access_log
-bash: -rw-r--r--.: command not found
[root@osboxes httpd]# -rw-r--r--.  1 root root 1039 Feb  8 09:53 error_log
-bash: -rw-r--r--.: command not found
[root@osboxes httpd]# [root@osboxes httpd]# cat access_log
-bash: [root@osboxes: command not found
[root@osboxes httpd]# [root@osboxes httpd]# vim access_log
-bash: [root@osboxes: command not found
[root@osboxes httpd]# [root@osboxes httpd]# file access_log
-bash: [root@osboxes: command not found
[root@osboxes httpd]# access_log: empty
-bash: access_log:: command not found
[root@osboxes httpd]# [root@osboxes httpd]#
