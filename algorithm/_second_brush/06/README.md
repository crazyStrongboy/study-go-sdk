今日二刷十题:
1. 最小栈（核心点：栈里面存元素本身值和比前面所有数据的最小值）
2. 相交链表（先求得两个链表长度，长度长的先走（long-short）步数，然后再一起next，相交即有交点，否则返回nil）
3. 多数元素(摩尔投票法)
   1. 分别用两个变量记录元素elem和数量count
   2. 遇见相等的元素，则count++
   3. 遇见不相等的元素，则count--,如果count==0,那么更换元素
   4. 最终留下来的elem就是占一半数以上的元素
4. 打家劫舍(动态规划dp[i]=max(dp[i-2]+nums[i],dp[i-1]))
5. 岛屿数量(污染周围的岛屿，将临近为1的岛屿置为'2')
6. 反转链表（用pre记住head,cur=head.next）
7. 课程表
   1. 用map1记住完成一个课程可以解锁哪些课程，
   2. 用map2记住每个课程的深度，依赖一门课，则深度+1
   3. 优先从深度为0的课程开始学习,完成课程则从map1中获取可以解锁哪些课程，分别给这些课程深度-1
   4. 当课程深度为0时表示可以学习
8. 前缀树（[26]*Trie, isEnd）
9. 数组中第k大元素(快排法)
   1. 第k大元素则是第index（index=len(nums)-k）个元素
   2. 随机取mid=rand.Intn(right-left+1)+left,进行移位，使得count+1右边的元素都大于nums[mid]
   3. 这里要注意需要先交换nums[mid]和nums[right]位置，这样left-->right-1都可以被遍历到
   4. count计数从left-1开始，如果发现nums[i]<nums[right],那么count++,并交换nums[count]与nums[i]的位置
   5. count+1与index比较，相等则返回nums[count+1]
   6. count+1<index,则在右边检索，search(nums,count+2,right,index)
   7. count+1>index,则在左边检索, search(nums,left,count,index)
10. 最大正方形
    1. 定义sideLength=0
    2. 先将byte二维数组转化成dp int二维数组,只要有一个'1'则sideLength=1
    3. 然后递推公式 dp[i][j] = min(min(dp[i-1][j-1],dp[i-1][j]),dp[i][j-1])
    4. 取sideLength=max(sideLength,dp[i][j])即可