package main

// 并查集
// https://blog.csdn.net/xiaolinnulidushu/article/details/104123088

type union struct {
	parent []int //  parent[i]表示i这个元素指向的父亲节点
	rank   []int // 以i为根节点的树的高度
}

func initUnion(u *union, size int) {

	for i := 0; i < size; i++ {
		u.parent[i] = i // 初始的时候每个元素都指向自己
		u.rank[i] = 1   // 初始的时候以该元素为根节点的树的高度就为1
	}

}

// 查询的时候我们是查询该元素所在的根节点
// 方式一的代码实现
func (u *union) find_method1(p int) int {
	for u.parent[p] != p { // 说明还没到达根节点
		// 当前节点遍历时将节点指向父亲的父亲节点
		u.parent[p] = u.parent[u.parent[p]]
		p = u.parent[p]
	}

	return p
}

// 查询的时候我们是查询该元素所在的根节点
// 方式二的代码实现，迭代版本的实现
func (u *union) find_method2_1(p int) int {
	curp := p
	for u.parent[p] != p { // 说明还没到达根节点
		curp = u.parent[p]
	}
	u.parent[p] = curp
	return curp
}

func (u *union) find_method2_2(p int) int {
	for u.parent[p] != p { // 说明还没到达根节点
		// 递归更新p的父亲节点，这种方式是将所有路径的节点的父亲一次性更新为根节点
		u.parent[p] = u.find_method2_2(u.parent[p])

	}
	return u.parent[p]
}

func (u *union) isConnected(p, q int) bool {
	return u.find(p) == u.find(q)
}

// 合并两个元素
// 合并的时间复杂度就由find来决定啦，不再是O(n)的时间复杂度
func (u *union) union(p, q int) {
	pRoot := u.find(p)
	qRoot := u.find(q)

	if pRoot == qRoot {
		return
	}

	// 这里合并的时候判断树的高度，将其高度低的节点指向高度高的根节点
	if u.rank[pRoot] < u.rank[qRoot] {
		u.parent[pRoot] = qRoot

	} else if u.rank[pRoot] > u.rank[qRoot] {
		u.parent[qRoot] = pRoot
	} else {
		// rank[pRoot] == rank[qRoot] 相等的情况下，随便指向，但是高度需要维护
		u.parent[pRoot] = qRoot
		u.rank[qRoot]++
	}

}

func (union *union) find(p int) int {
	return union.find_method1(p)

}
