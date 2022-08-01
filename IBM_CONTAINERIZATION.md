EXAMPLE 1: HOSTING A JS APP ON VM
- Structure for VM 
  - Hardware layer
  - Host OS
  - Hyperviser
  - Linux VM ---needed to run the js app 
   - VM has its own guest OS and other required liberaries along with the JS app
- To scale the app multiple VMs are setup each having its own guest OS and libraries and JS application
- DRAWBACKS
  - Memory : even a small js app requires a separate guest OS, thus initial setup for OS etc may require more memory than the actual application.
  - Cost : similarly uneccesary cost will go up.
  - Not Portable : if application is developed on Windows machine then it may not run on a Linux/Mac OS

VS

 HOSTING A JS APP USING A CONTAINER
- CONTAINERIZATION is a 3 step process 
  - creat a manifest/YAML file
  - build the imagee
  - run the container. This container contains all the dependencies/binaries/libraries required to run the JS application
- Structure for VM 
  - Hardware layer
  - Host OS
  - Container Runtime eg. Docker Engine
  - Container with the required libraries and JS app
- To scale the app multiple containers are setup each having its own libraries and JS application
- ADVANTAGE over VM
  - multiple guest VMs arent required to scale the app
  - the conatiner will run irrespective of the OS
  
  
  
 EXAMPLE 2 : Suppose we want to use a third party api via Python, i.e JS intereacts with Python which then connects to the API
 via VM
 - To reuse the exsisting VMs we can run python on them. Hence each VM will have JS and Python running
 - Drawback : Cant independently scale only JS VM's since VM has both Python and JS on it
 - Solution : Free up any of the existing VM and run only Python on it. This isn't the ideal way in cloud native.
   
 via Container
 - simply add a new container with python 
 - addtional free resources will be shared by all the running containers
 - if some containers aren't using CPU/memory then these resources are freed so that other containers running on the machine can use them
 
 ![docker_ibm](https://user-images.githubusercontent.com/37453877/182145646-87265508-3d05-4a47-a1db-1b6775873c16.JPG)
 
 - Resources : https://youtu.be/0qotVMX-J5s
