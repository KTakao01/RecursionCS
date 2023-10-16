package main

import (
	"sort"
)

func shipmentVolumePackages(packages []int32) int32 {
	// 関数を完成させてください
	var queue []int32
	var ans int32
	for len(packages) != 1 {
		queue = []int32{}
		sort.Slice(packages, func(i, j int) bool { return packages[i] < packages[j] })

		node := packages[0] + packages[1]
		ans += node

		packages = packages[2:]
		queue = append(queue, node)
		queue = append(queue, packages...)
		packages = queue
	}
	return ans
}

//データ構造を何使うか決める。すたっく？きゅー？後ろから扱うメリットない。キュー。
//1.配列を昇順にソートして配列の最初の小さい２数を足す。足した数をキューにいれる
//2.残りの要素をそのままキューに入れる。
//3.1.を繰り返したい。しかし、キューが足りない。packagesにキューをコピーする。キューをリセットする。
//
