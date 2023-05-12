# how to run an application on k8s

# 1. build your docker image

```Dockerfile
FROM codenvy/cpp_gcc
RUN mkdir /home/user/app; mkdir /home/user/app/include;mkdir /home/user/app/bin
ADD include/* /home/user/app/include/
ADD *.cpp /home/user/app/
ADD Makefile /home/user/app/
WORKDIR /home/user/app
RUN make all
```


A docker file contains several part, `FROM`follows the base image you work on. In the example above, we use `codenvy/cpp_cpp`, which is an image has `gcc`.

In next line, `RUN` follows a list of command which is seprated by `;` it is recommended to write these command in a single line.

`ADD [file] [path]`add file to path in image

`WORKDIR`is your working directory



 **After that, you are to build docker image with command `docker build -t [image name] [working dir]`**,you can look up your images with `docker images`



# 2. write yaml file

yaml file is used to specify configuration such as images name, pod name, resources your image need ,etc. A representative example is as follow

```YAML
kind: Deployment
apiVersion: apps/v1
metadata:
  name: ${ProJ}
spec:
  replicas: 1
  selector:
    matchLabels:
      app: ${ProJ}
  template:
    metadata:
      labels:
        app: ${ProJ}
    spec:
      containers:
      - name: ${ProJ}
        image: ${ImageName}
        imagePullPolicy: IfNotPresent
        command: [ "/bin/sh", "-c", "--" ]
        args: ["./csob ${FileName}"]

        volumeMounts:
        - mountPath: /nfs   #容器的挂载点,也就是在容器里访问PV的路径
          name: statics-datadir                        #被挂载卷的名称
      volumes:
      - name: statics-datadir  #共享存储卷名称,把下面的PVC声明一个卷叫做statics-datadir，再把这个卷挂载到上面的容器目录
        persistentVolumeClaim:   #类型是PVC
          claimName: nfs-pvc-orderbook  #指定要绑定的PVC，前面已经创建好了
```




# 3. apply your application with `kubectl apply -f [yaml file name]`

