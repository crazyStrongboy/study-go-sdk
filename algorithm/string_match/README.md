### 字符串匹配
1. BF算法，Brute Force 的缩写，中文叫作暴力匹配算法

   依次进行检查，时间复杂度O（n*m）----->主串长度为n，匹配串长度为m。

2. RK算法，Rabin-Karp 算法，是由它的两位发明者 Rabin 和 Karp 的名字来命名的。

   对没一个字符进行hash运算，然后将m个字符hash值相加，与匹配串的总值进行比较，相等则证明匹配上了，时间复杂度为O（n）。
   
3. BM算法

4. KMP算法