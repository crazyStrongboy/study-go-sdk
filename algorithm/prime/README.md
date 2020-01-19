### 求出所有的质数（从2开始）

整体思路如下：
![prime](http://images.hcyhj.cn/blogimages/prime/prime.png)

例如：
   
从源头开始，每次投递一个自然数到`prime`为2的通道中，发现3是第一个不被2整除的，所以取出3作为下一个通道的`prime`，过滤掉被3整除的所有数字，然后取第一个不被3整除的数作为下一个通道的`prime`,以此类推。