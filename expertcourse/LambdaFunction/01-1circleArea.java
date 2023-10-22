//https://recursionist.io/dashboard/problems/417

// ラムダ関連の問題は、記述したテストケースが実行されることで出力が正しいかを判断します。
// 通常のコーディング問題のように入力が自動で与えられるわけではありません。
// まずは、問題で指定されているcircleAreaのラムダを作成してみましょう。

// なお、右上の各ボタンは、下記の通りです。
// 実行：記述したコードを実行します。デバッグに使用します。
// テスト：提出時と同じ環境で実行されます。提出前の最終チェックに使用します。
// 提出：問題文中のテストケースが実行され、出力が正しいか確認されます。
import java.util.function.Function; 
class MyClass{
    public static void main(String[] args){
        Function<Integer, Double> circleArea = (x) -> x*x*3.14;
        // ① MyClass内で circleArea ラムダを作成してみましょう。
        /*Function<Integer, Double> circleArea = new Function<>(){
            public Integer apply(Integer x){
                return x*x*3.14
            }
        }
        */

        // ② 問題文にある一つ目のテストケースをMyClassに記述しましょう。
        // コンソール出力することでテストケースが正しい値であるかが判定されます。
        System.out.println(circleArea.apply(1));

        // ③ 上記を参考に他のテストケースを作成しましょう。
        // 問題文中のテストケースを全て満たすことで合格します。
        System.out.println(circleArea.apply(5));
        // ④ テストボタンを押して正しく出力されているかを確認しましょう。
        // メソッド名、文字列のスペースなどは注意深く確認しましょう。

        // ⑤ テストが確認できたら、提出ボタンを押して完了です。
    }
}