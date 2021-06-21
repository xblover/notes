// 算法P（打印前500个素数的表）
// 算法包含两个不同部分：步骤P1-P8在机器内部准备500个素数的表，
// 步骤P9-P11按照固定格式打印结果

//P1. [开始建表.]置PRIME[1]<-2, n<-3, j<-1.（在下面的步骤中，n将遍历可能是素数的奇数，j记录当前已经找到多少个素数）
//P2. [n是素数.]置j<-j+1, PRIME[j]<-n.
//P3. [已找到500个?]如果j=500,转到P9.
//P4. [推进n.]置n<-n+2.
//P5. [k<-2.]置k<-2.（PRIME[k]将遍历可能是n的素因数的数.）
//P6. [PRIME[k]\n?]用PRIME[k]除n，令q为商，r为余数.如果r=0（因此n不是素数），转到P4.
//P7. [PRIME[k]足够大?]如果q<=PRIME[k], 转到P2.（此时n必定是素数.见习题11？）
//P8. [推进k.]k增加1，然后转到P6.
//P9. [打印标题.]至此我们做好了打印素数表的准备.打印标题行，置m<-1.
//P10. [打印行.]以适当格式输出含PRIME[m],PRIME[50+m],...,PRIME[450+m]的行.
//P11. [已打印50行?]m增加1.如果m<=50,返回P10；否则算法总结. |||



//程序P（打印前500个素数的表）

%演示程序... 素数表
L		IS		500					//欲求素数的数量
t		IS		$255				//临时存储
n		GREG	0					//候选素数
q		GREG	0					//商
r		GREG	0					//余数
jj		GREG	0					//PRIME[j]的下标
kk		GREG	0					//PRIME[k]的下标
pk		GREG	0					//PRIME[k]的值
mm		IS		kk					//输出行的下标
		LOC		Data_Segment
PRIME1	WYDE	2					//PRIME[1]=2
		LOC		PRIME1+2*L	
ptop	GREG	@					//PRIME[501]的地址
j0		GREG	PRIME1+2-@			//jj的初值
BUF		OCTA	0					//十进制数字串的生成位置

		LOC		#100
Main	SET		n,3					//P1.开始建表.	n<-3.
		SET		jj,j0				//j<-1.
2H		STWU	n,ptop,jj			//P2.n是素数.	PRIME[j+1]<-n.
		INCL	jj,2				//j<j+1.
3H		BZ		jj,2F				//P3.已找到500个?
4H		



0H		GREG	#2030303030000000	//" 0000",0,0,0

