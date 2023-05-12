## 关于目录

.
├── kubernetes_notes.pdf                k8s 学习记录，包括部署过程、基本操作
├── openfaas                            serverless 工具
├── doc                                 文档
│   ├── deploy.md                       部署应用的步骤
│   ├── errors.md                       学习时遇到的错误
│   └── scheduler.md                    scheduler 相关的文档，包括原理，拓展点等
├── README.md       
├── scheduler                           customize scheduler 的简单实现
│   ├── extender                        scheduler extender
│   └── framework                       scheduler framework
└── yamls                               yaml 文件
    ├── kubemark                        kubemark 是 k8s 自带的性能测试工具
    ├── nfs                             将本地磁盘挂载为 NFS 
    └── scheduler-extender              部署 scheduler extender 的yaml文件
