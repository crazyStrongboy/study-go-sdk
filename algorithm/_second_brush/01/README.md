今日二刷十题:
1. 两数之和(hash表解决,这里可以一次遍历解决,要求不能重复选数,可以一边遍历,一边填充hash表)
2. 合并链表，两数相加(取模/取余)
3. 无重复字符最长子串(滑动窗口)
4. 寻找两个有序列表的中位数(k = ((len1+len2+1)/2 + (len1+len2+2)/2))/2 )
   1. 和求topK一样的思路,感觉就像二分法，分别从数组里面先淘汰一半，这里边界处理需要注意
   2. 边界1: 其中有个数组遍历完了，直接取另一个数组的start+k-1索引即可
   3. 边界2: 中位数前只剩一个数，则分别取两个数组start位置的值,比较大小，小的取胜
   4. 边界3: 取mid时，需要考虑start+k/2-1与len-1的大小,取小点的，防止越界
   5. 边界4: 进行下一轮寻找时,需要比较k-k/2与k-(len-start)的大小,大的进行传递，防止扣除的数不够
5. 最长回文串(分别从i或者i,i+1向两边辐射)
6. 正则表达式匹配(动态规划,初始化考虑空串匹配,向前递推考虑'*'、'.',然后考虑带有'*'的前面的字符串出现多少次(0次、1次、n次))
7. 盛最多的水(双指针,从头和尾分别向中间靠拢,按高度小的优先进行缩减)
8. 电话号码(回溯)
9. 三数之和(下次做优先考虑双指针,在满足条件的情况下对两边进行缩减,hash表写起来边界不好处理)
10. 删除链表中倒数第N个节点(快慢指针，快指针优先走n+1步,然后slow.next=slow.next.next,这里要注意用虚拟头节点,好处理头节点被移除这个边界)