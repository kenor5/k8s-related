# Errors encounter

1. **[network: error getting ClusterInformation: Get "https://10.96.0.1:443/apis/crd.projectcalico.org/v1/clusterinformations/default": x509: certificate signed by unknown authority (possibly because of "crypto/rsa: verification error" while trying to verify cand](network: error getting ClusterInformation: Get "https://10.96.0.1:443/apis/crd.projectcalico.org/v1/clusterinformations/default": x509: certificate signed by unknown authority (possibly because of "crypto/rsa: verification error" while trying to verify cand)**

  使用flannel，但是环境没有请理干净，dns起不来

  ls /etc/cni/net.d, 把除了fannel之外的删除碗

  [https://www.cnblogs.com/huiyichanmian/p/15760579.html](https://www.cnblogs.com/huiyichanmian/p/15760579.html)



2. Unable to connect to the server: x509: certificate signed by unknown authority (possibly because of "crypto/rsa: verification error" while trying to verify candidate authority certificate "kubernetes")

  [https://blog.csdn.net/woay2008/article/details/93250137](https://blog.csdn.net/woay2008/article/details/93250137)

  删除config，重建一下

  ```Plain Text
rm -rf $HOME/.kube
mkdir -p $HOME/.kube
sudo cp -i /etc/kubernetes/admin.conf $HOME/.kube/config
sudo chown $(id -u):$(id -g) $HOME/.kube/config
```


3. It seems like the kubelet isn't running or healthy

  当时没有解决，重装了集群配置好网络等，自己就active了

  

4. 集群装好之后，任务分配不到worker节点上

  先安装一个metrics[https://www.cnblogs.com/binghe001/p/12821804.html](https://www.cnblogs.com/binghe001/p/12821804.html)

  运行的时候报错“ensure CRDs are installed first”

5. 自定义调度器时，调度器一直重启

  Liveness probe failed: HTTP probe failed with statuscode: 400

  协议搞错了，把https写成http了

6. unable install v0.0

  需要在mod文件中更改版本重定向

  replace (

      k8s.io/api => k8s.io/api v0.25.7

      k8s.io/apiextensions-apiserver => k8s.io/apiextensions-apiserver v0.25.7

      k8s.io/apimachinery => k8s.io/apimachinery v0.25.7

      k8s.io/apiserver => k8s.io/apiserver v0.25.7

      k8s.io/cli-runtime => k8s.io/cli-runtime v0.25.7

      k8s.io/client-go => k8s.io/client-go v0.25.7

      k8s.io/cloud-provider => k8s.io/cloud-provider v0.25.7

      k8s.io/cluster-bootstrap => k8s.io/cluster-bootstrap v0.25.7

      k8s.io/code-generator => k8s.io/code-generator v0.25.7

      k8s.io/component-base => k8s.io/component-base v0.25.7

      k8s.io/component-helpers => k8s.io/component-helpers v0.25.7

      k8s.io/controller-manager => k8s.io/controller-manager v0.25.7

      k8s.io/cri-api => k8s.io/cri-api v0.25.7

      k8s.io/csi-translation-lib => k8s.io/csi-translation-lib v0.25.7

      k8s.io/kube-aggregator => k8s.io/kube-aggregator v0.25.7

      k8s.io/kube-controller-manager => k8s.io/kube-controller-manager v0.25.7

      k8s.io/kube-proxy => k8s.io/kube-proxy v0.25.7

      k8s.io/kube-scheduler => k8s.io/kube-scheduler v0.25.7

      k8s.io/kubectl => k8s.io/kubectl v0.25.7

      k8s.io/kubelet => k8s.io/kubelet v0.25.7

      k8s.io/kubernetes => k8s.io/kubernetes v1.25.7

      k8s.io/legacy-cloud-providers => k8s.io/legacy-cloud-providers v0.25.7

      k8s.io/metrics => k8s.io/metrics v0.25.7

      k8s.io/mount-utils => k8s.io/mount-utils v0.25.7

      k8s.io/pod-security-admission => k8s.io/pod-security-admission v0.25.7

      k8s.io/sample-apiserver => k8s.io/sample-apiserver v0.25.7

  )

