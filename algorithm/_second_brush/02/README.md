今日二刷七题:
1. 有效的括号（栈，成双成对出现，最后需要判断栈里面还有没剩余括号）
2. 合并有序列表（注意使用头虚拟节点，这里可以递归，也可以内部for循环）
3. 括号生成器（回溯或者动态规划，回溯用数组装要塞的括号，递归n*2次即可）
    1. 动态规划公式 str = "("+x+")"+y
    2. result[i]=result[j]+result[i-j-1]
    3. 这里result[0]=[]string{""}
4. 合并k个升序列表（用pre记住上一个合并好的列表，一个个合并即可，借助合并有序列表的代码）
5. 下一个排列（从尾往前找第一个小于上一个的元素，记住索引A，然后从尾往前找第一个大于该索引的元素B，交互A与B，然后翻转A+1之后的数据）
6. 最长有效括号（使用栈，栈头部填充-1，防止第一个元素是右括号，然后遍历记住左括号的下标，每来一个完整的括号则计算长度最大值，右括号多于左括号时会把-1弹出去，右括号索引排在栈定）
7. 搜索旋转排序数组（双指针i=0，j=len(nums)-1，计算出mid，判断target和`nums[mid]`是否相等,不相等情况下分别判断左有序还是右有序）
    1. 注意：这里判断左侧有序，如果用`nums[i]`判断，则需要用`nums[mid]`>=`nums[i]`
    2. 如果用`nums[j]`判断，则直接使用`nums[mid]`>`nums[j]`即可
