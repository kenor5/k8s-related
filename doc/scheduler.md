# how to write a scheduler

## 4 ways to customize a scheduler

1. directly clone the official kube-scheduler source code, modify the code directly.

  drawback: requires a lot of extra effort to keep up with the upstream scheduler changes

2. overriden by pod's  `spec.schedulerName`,

  drawback: multiple schedulers coexist may cause conflict when serverial pod is schedulered to same node.

3. scheduler extender, a configurable webhook that contains  `filter`  and  `priority`

4. scheduling framework. introduced in kubernetes v1.15. add a set of plug-in APIs to the existing scheduler, which keep the core of the scheduler simple and easy to maintain. And `scheduler extensions` have been deprecated in v1.16, so scheduling framework is the core way to customize the scheduler

## how does a scheduler work

5. the default scheduler is tarted with the specified parameters(config file in `/etc/kubernetes/manifests/kube-schdueler.yaml`)

6. watch apiserver and put Pods with empty `spec.nodeName` into the scheduler's internal scheduling queue

7. pop out a pod and schedule it

8. retrieve the "hard requirements(cpu/ memory/nodeAffinity)", and then **filter nodes** that donot fit

9. retrueve the "soft requirements" and applies some default "soft policies(e.g. pods tend to be more clustered or spread out on nodes)", and finally, it **gives a score to each node** and selects the final one with the highest score.

10. communicates with apiserver(send a binding call) and then sets the pod's `spec.nodeName`property to indicate the node to which the pod will be dispatched

## monitoring resources usage

- `kubectl top [pods/nodes] [podname/node name]`display the usage of cpu and memory



# 传统调度策略

 *源码*

[kubernetes/selector_spread.go at master · kubernetes/kubernetes](https://github.com/kubernetes/kubernetes/blob/master/pkg/scheduler/framework/plugins/selectorspread/selector_spread.go)



**可能存在的问题**

- **如果不指定硬性资源的话，会全部分配到一个节点上，并且会占用资源过多**

- 外部碎片



![image.png](https://flowus.cn/preview/2916e9f3-0fd8-4a63-9a89-e267d4f23360)

实际测试：三台机器，每台8G内存，先起三个3G内存的应用，再起一个6G的，会显示pending

![image.png](https://flowus.cn/preview/f3e4b2a9-b053-4231-96b8-5fd451ffbfd1)



如果我们优先把一个 node 占满，再去往另一个 pod 上调度，就能在一定程度上解决这个问题。 具体算法就是如下。

![image.png](https://flowus.cn/preview/d4c7d279-7433-4eb8-b345-ea05d53a9507)


# scheduler plugin

### the meaning of different extention points

![image.png](https://flowus.cn/preview/ec27c1e0-fefa-41cf-b733-59abfdd830c0)

### PreEnqueue

These plugins are called prior to adding Pods to the internal active queue, where Pods are marked as ready for scheduling.

Only when all PreEnqueue plugins return `Success`, the Pod is allowed to enter the active queue. Otherwise, it's placed in the internal unschedulable Pods list, and doesn't get an `Unschedulable` condition.

For more details about how internal scheduler queues work, read [Scheduling queue in kube-scheduler](https://github.com/kubernetes/community/blob/f03b6d5692bd979f07dd472e7b6836b2dad0fd9b/contributors/devel/sig-scheduling/scheduler_queues.md).

### QueueSort

These plugins are used to sort Pods in the scheduling queue. A queue sort plugin essentially provides a `Less(Pod1, Pod2)` function. Only one queue sort plugin may be enabled at a time.

### PreFilter

These plugins are used to pre-process info about the Pod, or to check certain conditions that the cluster or the Pod must meet. If a PreFilter plugin returns an error, the scheduling cycle is aborted.

### Filter

These plugins are used to filter out nodes that cannot run the Pod. For each node, the scheduler will call filter plugins in their configured order. If any filter plugin marks the node as infeasible, the remaining plugins will not be called for that node. Nodes may be evaluated concurrently.

### PostFilter

These plugins are called after Filter phase, but only when no feasible nodes were found for the pod. Plugins are called in their configured order. If any postFilter plugin marks the node as `Schedulable`, the remaining plugins will not be called. A typical PostFilter implementation is preemption, which tries to make the pod schedulable by preempting other Pods.

### PreScore

These plugins are used to perform "pre-scoring" work, which generates a sharable state for Score plugins to use. If a PreScore plugin returns an error, the scheduling cycle is aborted.

### Score

These plugins are used to rank nodes that have passed the filtering phase. The scheduler will call each scoring plugin for each node. There will be a well defined range of integers representing the minimum and maximum scores. After the [NormalizeScore](https://kubernetes.io/docs/concepts/scheduling-eviction/scheduling-framework/#normalize-scoring) phase, the scheduler will combine node scores from all plugins according to the configured plugin weights.

### NormalizeScore

These plugins are used to modify scores before the scheduler computes a final ranking of Nodes. A plugin that registers for this extension point will be called with the [Score](https://kubernetes.io/docs/concepts/scheduling-eviction/scheduling-framework/#scoring) results from the same plugin. This is called once per plugin per scheduling cycle.

For example, suppose a plugin `BlinkingLightScorer` ranks Nodes based on how many blinking lights they have.

