#include<bits/stdc++.h>
using namespace std;
unsigned int getBlkno(string &input_data,int idx){
    string s = input_data.substr(idx*4,4);
    unsigned int res = 0;
    for(int i=0;i<4;i++){
        res += s[i]<<(24-i*8);
    }
    return res;
}
unsigned int Murmurhash3(string input_data,unsigned int seed){
    unsigned int h = seed;
    //声明常量
    const unsigned int c1 = 0xcc9e2d51;		// 3,432,918,353
    const unsigned int c2 = 0x1b873593;		// 461,845,907

    const int r1 = 15;
    const int r2 = 13;

    const int m = 5;
    const int n = 0xe6546b64;	//3,864,292,196
    //分块处理
    int blkNums = input_data.size()/4;
    //1.一个块一个块地处理,这是第一部分地工作
    for(int i=0;i<blkNums;i++){
        unsigned int K = getBlkno(input_data,i);
        K *= c1;
        K = _rotl(K,r1);
        K *= c2;
        K = _rotl(K,r2);
        h = h*m + n;
    }
    //2.处理剩余量
    string remaining_bytes = input_data.substr(blkNums*4);
    unsigned int k = 0;
    switch (remaining_bytes.size()){
        case 3:k^=remaining_bytes[2]<<16;
        case 2:k^=remaining_bytes[1]<<8;
        case 1:k^=remaining_bytes[0];
    }
    k = k * c1;
    k = _rotl(k,r1);
    k = k * c2;
    h ^= k;
    h ^= input_data.size();
    //3.加强雪崩测试
    h ^= h >> 16;
	h *= 0x85ebca6b;	// 2,246,822,507
	h ^= h >> 13;
	h *= 0xc2b2ae35;	// 3,266,489,909
	h ^= h >> 16;
    return h;
}

int main(){
    string input_data = "Jack";
    int seed = 123;
    cout<<"Murmurhash3(Jack,123) = "<<Murmurhash3(input_data,seed);
    system("pause");
}
/***************************************
 * Written By I am Jack(JackTan)
 * 
 * Murmurhash3 Test Programming
 **/