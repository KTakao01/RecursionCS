// テストの問題は、記述したテストケースをテストした結果が判断されます。
// 通常のコーディング問題のように入力が自動で与えられるわけではありません。
// まずは、問題で指定されているfibonacciのテストを作成してみましょう。

// なお、右上の各ボタンは、下記の通りです。
// 実行：記述したコードを実行します。デバッグに使用します。
// テスト：提出時と同じ環境で実行されます。提出前の最終チェックに使用します。
// 提出：問題文中のテストケースが実行され、出力が正しいか確認されます。

class MyClass {
    // 問題文中で指定がない場合、テスト対象の関数は修正する必要はありません。
    public static long fibonacci(int n) {
        long[] cache = new long[n + 1];
        cache[0] = 0;
        cache[1] = 1;

        for (int i = 2; i <= n; i++) {
            cache[i] = cache[i - 1] + cache[i - 2];
        }

        return cache[n];
    }

    // テストで使用する関数です。変更の必要はありません。
    public static void unitTestCheck(boolean predicate) {
        if (predicate) {
            System.out.println("The test passed!!");
        } else {
            System.out.println("ERROR! The test failed!!");
        }
    }

    // unitTests()が実行され、テストの結果がチェックされます。、
    public static void unitTests() {
        // ① 問題文にある一つ目のテストケースを記述しましょう。
        unitTestCheck(fibonacci(3) == 2);

        // ② 上記を参考に他のテストケースを作成しましょう。
        unitTestCheck(fibonacci(4) == 3);
        unitTestCheck(fibonacci(6) == 8);
        unitTestCheck(fibonacci(9) == 34);
        unitTestCheck(fibonacci(10) == 55);
        // テストケースを記述することが目的の問題や、
        // 指定したテストケースが全て合格するように関数を修正することが目的の問題もあります。

        // ④ テストボタンを押して正しく出力されているかを確認しましょう。

        // ⑤ テストが確認できたら、提出ボタンを押して完了です。

    }
}
