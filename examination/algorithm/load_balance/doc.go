package load_balance

// 常见负载均衡算法:
// 1.负载均衡（Load balancing）是一种将网络流量分布到多个服务器或计算资源上的技术，以提高系统性能、利用率和可靠性。下面是几种常见的负载均衡算法：
// 2.轮询（Round Robin）算法：将请求按顺序分配给后端服务器，每个请求依次轮流分配到不同的服务器上。适用于服务器性能相近且无状态的情况。
// 3.最少连接（Least Connections）算法：将请求分配给当前连接数最少的服务器，以实现负载均衡。适用于处理长连接或请求处理时间不均匀的情况。
// 4.IP哈希（IP Hash）算法：将请求根据客户端IP地址的哈希值分配给特定的服务器。对于相同的IP地址，始终将请求发送到同一台服务器上，适用于需要保持会话一致性的应用。
// 5.加权轮询（Weighted Round Robin）算法：为每个服务器分配一个权重值，根据权重比例决定发送请求的概率。权重越高的服务器获取到的请求数量越多。
// 6.加权最少连接（Weighted Least Connections）算法：结合了加权和最少连接算法，通过权重值和当前连接数来决定请求的分配。权重越高的服务器拥有更多的处理能力。
// 7.随机（Random）算法：随机选择一个服务器来处理请求，以实现负载均衡。适用于服务器性能相近且无状态的情况。
// 8.响应时间（Response Time）算法：根据服务器的平均响应时间来进行负载均衡，将请求发送给响应时间最短的服务器。适用于需要考虑服务器性能差异的情况。

// nginx支持负载均衡算法:
// 1.轮询算法
// 2.加权算法
// 3.IP哈希算法
// 4.URL哈希算法
// 5.响应时间算法
