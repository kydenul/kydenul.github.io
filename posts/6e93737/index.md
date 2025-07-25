# Apache Kafka 深度解析：分布式流处理平台完全指南


{{< admonition type=abstract title="导语" open=true >}}
在现代分布式系统架构中，**数据流处理**已成为企业数字化转型的核心驱动力。Apache Kafka 作为业界领先的分布式流处理平台，以其卓越的**高吞吐量**、**低延迟**和**高可用性**特性，在大数据生态系统中扮演着至关重要的角色。

无论是**实时数据管道**构建、**微服务架构**中的消息传递，还是 **大规模日志收集**与**流式计算**，Kafka 都能提供稳定可靠的解决方案。本文将带您深入探索 Kafka 的核心架构设计、技术原理以及在企业级应用中的最佳实践，助您构建更加健壮的数据驱动型应用系统。
{{< /admonition >}}

<!--more-->

## 发布与订阅消息系统

**数据（消息）的发送者（发布者）不会直接把消息发送给接收者**，这是**发布与订阅消息系统**的一个特点。

发布者以某种方式对消息进行分类，接收者（订阅者）通过订阅它们来接收特定类型的消息。发布与订阅系统一般 会有一个 broker，也就是发布消息的地方。

---

## Kafka

Kafka，是一款基于发布与订阅模式的消息系统，一般被称为“分布式提交日志” 或 “分布式流式平台”。

文件系统或数据库提交日志，是为了在保存事务的持久化记录，通过重放这些日志可以重建系统状态。
同样地，**Kafka 的数据是按照一定的顺序持久化保存的，并且可以按需读取**。另外，**Kafka 的数据分布在整个系统中，具备数据故障恢复能力和性能伸缩能力**。

### 信息和批次

为了提高效率，消息会被分批次写入 Kafka。

---

#### 消息

Kafka 的数据单元被称为**消息**，它有字节数组组成。
对于 Kafka 来说，**消息里的数据没有特殊的格式或含义**。

消息可以有一个可选的元数据，也就是**键**。
键也是一个字节数组，与消息一样，对 Kafka 来说没有特殊含义。

*当需要以一种可控的方式将消息写入不同的分区时，需要用到键*。
最简单的例子，就是为键生成一个一致性哈希值，然后用哈希值对主题分区数进行取模，为消息选取分区。这样可以保证具有相同键的消息总是会被写到相同的分区中（前提是分区数量没有发生变化）。

---

#### 批次

**批次**包含了一组属于同一 Topic 和 Partition 的消息。

批次越大，单位时间内处理的消息就越多，单于单条消息来说，其传输时间就越长。
消息批次会被压缩，这样可以提升数据的传输和存储性能，但需要做更多的计算处理.

---

### 模式

消息模式（schema）有很多可选项，例如 JavaScript Object Notation（JSON）、Extensible Markup Language（XML）、Apache Avro等等。

**数据格式一致性**对 Kafka 来说非常重要，它消除了消息读写操作之间的耦合性。

---

### 主题和分区

Kafka 的消息通过 **主题（Topic）** 进行分类。
主题可以被分为若干个 **分区（Partition）**，一个分区就是一个提交日志。
**消息会以追加的方式被写入分区，然后按照先入先出的顺序读取；需要注意的是，由于一个主题一般包含几个分区，因此无法在整个主题范围内保证消息的顺序，但可以保证信息在单个分区内的有序性**。

**Kafka 通过分区来实现数据的冗余和伸缩，分区可以位于不同的服务器上，分区也可以被复制（相同分区的多个副本可以保证在多台服务器上，以防止其中一台服务器发放故障）**。

主题就好比数据库的表或文件系统的文件夹。

**流**，是一组从生产者移动到消费者的数据。因此，通常会使用 **流（steam）** 来描述 Kafka 这类系统中的数据，即把一个主题的数据看作一个流，不管它有多少个分区。

*流式处理有别于离线处理框架（如 Hadoop）处理数据的方式，后者被用来在未来某个时刻处理大量的数据。*

---

### 生产者和消费者

Kafka 的客户端就是Kafka系统的用户，其被分为两种基本类型：生产者（Producer）和消费者（Consumer）。
*除此之外，还有其他高级客户端API --- 用于数据集成的 Kafka Connect API、用于流式处理的 Kafka Streams，这些高级客户端 API 使用生产者和消费者作为内部组件，提供了更高级的功能。*

- **生产者** 创建消息。在默认情况下，生产者会把信息均衡地分布到对应主题的所有分区中。
*可以通过键和分区器，实现将特定消息写入指定的分区。*

- **消费者** 读取消息。消费者会订阅一个或多个主题，并按照消息写入分区顺序读取。
消费者通过检查消息的偏移量来区分已经读取过的消息。

- **偏移量**，是一个不断递增的整数值，也是一种元数据。在创建消息时，Kafka 会把它添加到消息里。在给定的分区中，每一条消息的偏移量都是唯一的，越往后消息的偏移量越大（不一定严格单调递增）。

消费者可以是**消费者群组**的一部分，属于同一群组的一个或多个消费者共同读取一个主题。
群组可以保证每个分区只被该群组里的一个消费者读取，这种消费者与分区之间的映射关系，通常被称为消费者对分区的**所有权关系**。

{{< figure src="/posts/kafka/images/消费者群组从主题读取消息.jpg" title="" >}}

---

### broker 和集群

**一台单独的 Kafka 服务器被称为 broker**。根据硬件配置及其性能特征的不同，单个 broker 可以轻松处理数千个分区和没秒百万级的消息量。

- broker 会接收来自生产者的消息，为其设置偏移量，并提交到磁盘保存。
- broker 会为消费者提供服务，对读取分区的请求做出响应，并返回已经发布的消息。

broker 组成了**集群**。
每个集群都有一个充当了**集群控制器**角色的 broker（自动从活动的集群成员中选举出来），它负责管理工作，包括为 broker 分配分区和监控 broker。
在集群中，一个分区从属于一个broker，这个broker被称为分区的**首领**。一个被分配给其他broker的分区副本叫做这个分区的“跟随者”。
分区复制提供了分区的消息冗余，如果一个 broker 发生故障，则其中的一个跟随者可以接管它的领导权。
所有想要发布消息的生产者必须连接到首领，但消费者可以从首领或者跟随者那里读取消息。

{{< figure src="/posts/kafka/images/集群中的分区复制.jpg" title="" >}}

**保留消息**（在一定期限内）是 Kafka 的一个重要特性。
broker 默认的保留策略如下：要么保留一段时间（例如 7 天），要么保留消息总量达到一定的字节数（例如 1GB）。
当消息数量达到这些上限时，就消息就会过期并被删除 => **在任意时刻，可用消息总量都不会超过配置参数所指定的大小**

---

### 多集群

随着 broker 数量的增加，最好使用多个集群，原因如下：

- 数据类型分离
- 安全需求隔离
- 多数据中心（灾难恢复）

**Kafka 的消息复制机制只能在单个集群中而不能在多个集群之间进行**。
Kafka 提供了一个叫做 **MirrorMaker** 的工具，它可以将数据复制到其他集群中。

MirrorMaker 的核心组件包括一个消费者和一个生产者，它们之间通过队列相连：

- 消费者会从一个集群读取消息
- 生产者会把消息发送到另一个集群

{{< figure src="/posts/kafka/images/多数据中心架构.jpg" title="多数据中心架构" >}}

---

### 数据生态系统

{{< figure src="/posts/kafka/images/大型数据生态系统.jpg" title="大型数据生态系统" >}}

#### 应用场景

##### 活动跟踪

Kafka 最初的应用场景是**跟踪网站用户的活动**。

网站用户与前端应用程序发生交互，前端应用程序生成与用户活动相关的消息。
这些消息既可以是一些静态信息，比如页面访问次数和点击量，也可以一些复杂的操作，比如修改用户资料。
这些消息会被发布到一个或多个主题上，并会被后端应用程读取。
这样就可以生成报告，为机器学习系统提供数据，更新搜索结果，或者实现更多其他功能。

---

##### 传递消息

Kafka 的另一个基本用途是传递消息。

应用程序向用户发送通知（如邮件）就是通过消息传递来实的。
这些应用程序组件可以生成消息，而无须关心消息的格式以及消息是如何被发送出去的。
一个公共应用程序会负责读取并处理如下这些消息。

- 格式化消息（也就是所谓的**装饰**）。
- 将多条消息放在同一个通知里发送。
- 根据用户配置的首选项来发送消息。

---

##### 指标和日志记录

Kafka 也可以用来收集应用程序以及系统的指标和日志。Kafka 的多生产者特性在这个时候就派上用场了。

应用程序定期把指标发布到 Kafka 主题上，监控系统或告警系统会读取这些消息。
Kafka 也可以被用在离线处理系统（如 Hadoop）中，进行较长时间片段的数据分析，比如年度增长走势预测。
我们也可以把日志消息发布到 Kafka 主题上，然后再路由给专门的日志搜索系统（如 Elasticsearch）或安全分析应用程序。
更改目标系统（如日志存储系统）不会影响前端应用程序或聚合方法，这是Kafka 的另一个优点。

---

##### 提交日志

Kafka 的基本概念源自提交日志。

我们可以把数据库的更新发布到 Kafka，然后应用程序会通过监控事件流来接收数据库的实时更新。
这种变更日志流也可以用于把数据库的更新复制到远程系统，或者将多个应用程序的更新合并到一个单独的数据库。
持久化的数据为变更日志提供了缓冲，也就是说，如果消费者应用程序发生故障，则可以通过重放这些日志来恢复系统状态。
另外，可以用紧凑型主题更长时间地保留数据，因为我们只为一个键保留了一条最新的变更数据。

---

##### 流式处理

流式处理是另一个包含多种类型应用程序的领域。

虽然可以认为大部分 Kafka 应用程序是基于流式处理，但真正的流式处理通常是指提供了类似 map/reduce（Hadoop）处理功能的应用程序。
**Hadoop 通常依赖较长时间片段的数据聚合，可以是几小时或几天。**
**流式处理采用实时的方式处理消息，速度几乎与生成消息一样快**。
开发人员可以通过用流式处理框架开发小型应用程序来处理 Kafka 消息，执行一些常见的任务，比如指标计数、对消息进行分区或使用多个数据源的数据来转换消息，等等。


---

> Author: [kyden](https://github.com/kydenul)  
> URL: http://kydenul.github.io/posts/6e93737/  

