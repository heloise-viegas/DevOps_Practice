<!--linux skill up challenge-->

#### Day 2 - Linux Skill Up Challenge

**Aim** : Basic Navigation

PS: This was posted a month ago on Reddit. Not sure how this works but my goal is just to learn so I took it up .

**Tools Used** :
- Oracle Virtual Machine - CentOS Distribution
- Windows Subsystem for Linux - Ubuntu Distribution

**Procedure :**
1) Connect to VM server via Ubuntu</br>
   Open Ubuntu terminal.</br>
   Commands Used :</br>
   i)ssh user_Id_of_VM@ipAddress_of_VM</br>
   
Operations performed on Remote Server</br>
[osboxes@osboxes ~]$ pwd</br>
/home/osboxes</br>
[osboxes@osboxes ~]$ mkdir test</br>
[osboxes@osboxes ~]$ ls</br>
demo.txt  test</br>
[osboxes@osboxes ~]$ mv demo.txt test</br>
[osboxes@osboxes ~]$ ls</br>
test</br>
[osboxes@osboxes ~]$ cd test</br>
[osboxes@osboxes test]$ pwd</br>
/home/osboxes/test</br>
[osboxes@osboxes test]$ ls -a</br>
.  ..  demo.txt</br>
[osboxes@osboxes test]$ cat > places.txt</br>
USA Dubai India UK 
[osboxes@osboxes test]$ cat places.txt</br>
USA Dubai India UK</br>
[osboxes@osboxes test]$ grep -i "bai" places.txt</br>
Dubai</br>
[osboxes@osboxes test]$ grep -i "i" places.txt | tr a-z A-Z</br>
DUBAI INDIA</br>
[osboxes@osboxes test]$ pushd ~/</br>
~ ~/test</br>
[osboxes@osboxes ~]$ pwd</br>
/home/osboxes</br>
[osboxes@osboxes ~]$ pushed test</br>
-bash: pushed: command not found</br>
[osboxes@osboxes ~]$ pushd ~/test</br>
~/test ~ ~/test</br>
[osboxes@osboxes test]$ popd</br>
~ ~/test</br>
[osboxes@osboxes ~]$ pwd</br>
/home/osboxes</br>
[osboxes@osboxes ~]$ dirs -v</br>
 0  ~
 1  ~/test</br>
[osboxes@osboxes ~]$ pwd</br>
/home/osboxes</br>
[osboxes@osboxes ~]$ ls</br>
test</br>
[osboxes@osboxes ~]$ man pushd</br>
[osboxes@osboxes ~]$ pushd test</br>
~/test ~ ~/test</br>
[osboxes@osboxes test]$ mkdir insidetest</br>
[osboxes@osboxes test]$ pushd insidetest</br>
~/test/insidetest ~/test ~ ~/test</br>
[osboxes@osboxes insidetest]$ pwd</br>
/home/osboxes/test/insidetest</br>
[osboxes@osboxes insidetest]$ mkdir final</br>
[osboxes@osboxes insidetest]$ pushd ~</br>
~ ~/test/insidetest ~/test ~ ~/test</br>
[osboxes@osboxes ~]$ pwd</br>
/home/osboxes</br>
[osboxes@osboxes ~]$ dir -v</br>
test</br>
[osboxes@osboxes ~]$ popd</br>
~/test/insidetest ~/test ~ ~/test</br>
[osboxes@osboxes insidetest]$ dir -v</br>
final</br>
[osboxes@osboxes insidetest]$ popd</br>
~/test ~ ~/test</br>
[osboxes@osboxes test]$ dir -v</br>
demo.txt  insidetest  places.txt</br>
[osboxes@osboxes test]$ popd</br>
~ ~/test</br>
[osboxes@osboxes ~]$ popd</br>
~/test</br>
[osboxes@osboxes test]$ popd</br>
-bash: popd: directory stack empty</br>
[osboxes@osboxes test]$ pwd</br>
/home/osboxes/test</br>
[osboxes@osboxes test]$ pushd final</br>
-bash: pushd: final: No such file or directory</br>
[osboxes@osboxes test]$ pushd test</br>
-bash: pushd: test: No such file or directory</br>
[osboxes@osboxes test]$ popd</br>
-bash: popd: directory stack empty</br>
[osboxes@osboxes test]$ pushd insidetest</br>
~/test/insidetest ~/test</br>
[osboxes@osboxes insidetest]$ popd</br>
~/test</br>
[osboxes@osboxes test]$ popd</br>
-bash: popd: directory stack empty</br>
[osboxes@osboxes test]$ pushd insidetest/final</br>
~/test/insidetest/final ~/test</br>
[osboxes@osboxes final]$ popd</br>
~/test</br>
[osboxes@osboxes test]$ popd</br>
-bash: popd: directory stack empty</br>



**Resources**:-</br>
1)https://www.reddit.com/r/linuxupskillchallenge/comments/raesp8/day_2_basic_navigation/</br>
2)https://www.tecmint.com/pushd-and-popd-linux-filesystem-navigation/
