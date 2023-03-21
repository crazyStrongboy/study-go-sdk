二刷不会的题汇总：
1. 下一个排列（从尾往前找第一个小于上一个的元素，记住索引A，然后从尾往前找第一个大于该索引的元素B，交互A与B，然后翻转A+1之后的数据）
2. 最长有效括号（使用栈，栈头部填充-1，防止第一个元素是右括号，然后遍历记住左括号的下标，每来一个完整的括号则计算长度最大值，右括号多于左括号时会把-1弹出去，右括号索引排在栈定）
3. 搜索旋转排序数组（双指针i=0，j=len(nums)-1，计算出mid，判断target和`nums[mid]`是否相等,不相等情况下分别判断左有序还是右有序）
    1. 注意：这里判断左侧有序，如果用`nums[i]`判断，则需要用`nums[mid]`>=`nums[i]`
    2. 如果用`nums[j]`判断，则直接使用`nums[mid]`>`nums[j]`即可
4. 单词搜索 （回溯，将上下左右依次回溯即可，回溯前将选过的位置置为'.'，防止重复选择）
5. 二叉树展开为链表
   1. 前序遍历的反向遍历法
   2. 用全局pre记住上一个节点即可
6. 排序列表(这里进行二分，然后一一合并有序列表即可)
   1. 二分时使用快慢指针，fast比slow走快两倍，这样slow就是mid了，在分别合并（head,mid）与（mid,tail）即可 
   2. 边界处理时注意，head.Next == tail时，返回head前需要将head.Next置空
7. 只出现一次的数字
   1. 第一种方式是数组取模
   2. 第二种方式是异或，x^0=x  x^x=0
8. 最长连续序列（hash表方式）
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
11. 多数元素(摩尔投票法)
    1. 分别用两个变量记录元素elem和数量count
    2. 遇见相等的元素，则count++
    3. 遇见不相等的元素，则count--,如果count==0,那么更换元素
    4. 最终留下来的elem就是占一半数以上的元素
12. 寻找重复数(将数组看成链表，然后快慢指针去解决这个问题，slow=nums[slow],fast=nums[nums[fast]])
13. 戳气球(这里把问题转换成添加气球,这样可以保证两边的值都是已有的)
    1. 具体公式 dp[i][j]=max(dp[i][j],dp[i][k]+nums[i]*nums[k]*nums[j]+dp[k][j])
    2. 这里需要给nums首尾都需要添加一个1，题意也表明边界都是乘以1的，所以可以在入口就直接补上去
    3. 这里遍历时要注意：i从最大值开始遍历,n-3，j从i+2处遍历,也就是n-1，k从i+1处遍历,也就是n-2,k必须要在i和j中间
14. 比特位计数
    1. 最简单的办法是累计i=i&(i-1),直到i为0，就知道有多少个1
    2. 第二个办法nums[i>>1]+i&1,去掉最低位，看有多少个1，然后加上末尾是否为1
    3. 第三个办法nums[i&(i-1)]+1,去掉最高位，直接补上1
15. 前k个高频元素（最简单的办法是用个二维数组进行排序，这样时间复杂度和空间复杂度都比较高）
     1. 先用map搜集出所有的元素的数量
     2. 利用最小堆的原理永远只保留三个元素
16. 和为k的子数组
    1. 第一种方式：暴力破解，两层for循环
    2. 第二种方式：前缀+map形式，注意map里面存的前缀累计count，不能是1
17. 最短无序连续子数组
     1. 第一种方式：拷贝一个出来进行排序，然后两个数组分别从首尾像中间推进，第一个不相等的即为left和right，
        注意边界处理，从左边开始，如果一个不相等的都找不到，则返回最大值
     2. 第二种方式： 
        1. max=math.MinInt32,从左边开始往右边找,如果小于max则right = i,大于则max=nums[i] (保证right之后都不小于right)
        2. min=math.MaxInt32,从右边往左边找,如果大于min，则left = i,小于则min=nums[i] (保证left之前都不大于left)
        3. 计算right-left即可(这里可以保证[0:left] [left:right] [right:]三段是递增的关系)