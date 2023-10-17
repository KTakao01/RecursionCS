<?php
// https://recursionist.io/dashboard/course/3/lesson/480
// 二分木 root が与えられるので、階層ごとに左->右で配列で返す、levelOrderTraversal という関数を作成してください。
class BinaryTree{
    public $data;
    public ?BinaryTree $left;
    public ?BinaryTree $right;

    public function __construct($data, $left = null, $right = null){
        $this->data = $data;
        $this->left = $left;
        $this->right = $right;
    }
}

function levelOrderTraversal(?BinaryTree $root): array{
    // 関数を完成させてください
    $queue = [$root];
    $results = [];
    while (count($queue) !== 0 ) {
        $curr = array_shift($queue);
        if ($curr === 'null'){
            array_push($results,'null');
            continue;
        }

        array_push($results,$curr->data);
        if ($curr->left !== null){
            array_push($queue,$curr->left);
        } else {
            array_push($queue,'null');
        }
        if ($curr->right !== null){
            array_push($queue,$curr->right);
        } else {
            array_push($queue,'null');
        }
    }
    while (count($results) > 0 && $results[count($results)-1] === 'null')array_pop($results);
    return $results;
}
