package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

// Complete the climbingLeaderboard function below.
func climbingLeaderboard(scores []int32, alice []int32) []int32 {
	ac := removeDuplicates(scores)
	rs := []int32{}
	for i := 0; i < len(alice); i++ {
		index := findIndex(ac, alice[i], 0, int32(len(ac)-1))
		rs = append(rs, index+1)
	}

	return rs
}

func findIndex(ac []int32, in int32, l, h int32) int32 {
	var m int32
	if in < ac[h] {
		return h + 1
	}
	if in > ac[l] {
		return 0
	}

	for l < h {
		m = (l + h) / 2

		if in < ac[m] {
			l = m + 1
			continue
		}

		h = m
	}
	if ac[l] == in {
		return l
	}
	if ac[l] < in || ac[l] > in {
		return l
	}

	return -1
}

func removeDuplicates(el []int32) []int32 {
	e := map[int32]bool{}
	rs := []int32{}

	for v := range el {
		if e[el[v]] == true {
		} else {
			e[el[v]] = true
			rs = append(rs, el[v])
		}
	}

	return rs
}

func main() {
	filePath := "/Users/namnh/workspace/go/src/github.com/namnhce/my-go-practices/climbing-the-leader-board/input.txt"
	output := "/Users/namnh/workspace/go/src/github.com/namnhce/my-go-practices/climbing-the-leader-board/output.txt"

	f, err := os.Open(filePath)
	checkError(err)

	reader := bufio.NewReaderSize(f, 1024*1024*1024)

	stdout, err := os.Create(output)
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 1024*1024)

	scoresCount, err := strconv.ParseInt(readLine(reader), 10, 64)
	checkError(err)
	scoresTemp := strings.Split(readLine(reader), " ")

	var scores []int32

	for i := 0; i < int(scoresCount); i++ {
		scoresItemTemp, err := strconv.ParseInt(scoresTemp[i], 10, 64)
		if err != nil {
		}
		checkError(err)
		scoresItem := int32(scoresItemTemp)
		scores = append(scores, scoresItem)
	}

	aliceCount, err := strconv.ParseInt(readLine(reader), 10, 64)
	checkError(err)

	aliceTemp := strings.Split(readLine(reader), " ")

	var alice []int32

	for i := 0; i < int(aliceCount); i++ {
		aliceItemTemp, err := strconv.ParseInt(aliceTemp[i], 10, 64)
		checkError(err)
		aliceItem := int32(aliceItemTemp)
		alice = append(alice, aliceItem)
	}

	result := climbingLeaderboard(scores, alice)

	for i, resultItem := range result {
		fmt.Fprintf(writer, "%d", resultItem)

		if i != len(result)-1 {
			fmt.Fprintf(writer, "\n")
		}
	}

	fmt.Fprintf(writer, "\n")

	writer.Flush()
}

func readLine(reader *bufio.Reader) string {
	str, _, err := reader.ReadLine()
	if err == io.EOF {
		return ""
	}

	return strings.TrimRight(string(str), "\r\n")
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

#include <bits/stdc++.h>
#define pb push_back
#define sqr(x) (x)*(x)
#define sz(a) int(a.size())
#define reset(a,b) memset(a,b,sizeof(a))
#define oo 1000000007

// using namespace std;

// typedef pair<int,int> pii;
// typedef long long ll;

// const int maxn=200007;

// int n,m,a[maxn],r[maxn];

// int main(){
// //    freopen("input.txt","r",stdin);
//     cin>>n;
//     for(int i=1; i<=n; ++i) cin>>a[i];
//     cin>>m;
//     r[1]=1;
//     for(int i=2; i<=n; ++i){
//         if(a[i]==a[i-1])
//			 r[i]=r[i-1];
//         else r[i]=r[i-1]+1;
//     }
//     int x=n;
//     r[0] = 0;
//     for(int i=1; i<=m; ++i){
//         int v;
//         cin>>v;
//         while(x>=1 && a[x] < v) --x;
//         if(x>0 && a[x]==v) cout<<r[x]<<endl;
//         else cout<<r[x]+1<<endl;
//     }
// }

