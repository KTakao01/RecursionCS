/*
 * https://recursionist.io/dashboard/problems/418
 * n が 3 の倍数のときは Fizz、5 の倍数のときは Buzz、15 の倍数のときは FizzBuzz、それ以外は n を返すようなラムダを作成し、テストケースを出力してください。
 */
import java.util.function.*;

class MyClass {
    public static void main(String[] args) {
        Function<Integer, String> fizzBuzz = (x) -> {
            StringBuilder results = new StringBuilder();
            for (int i = 1; i <= x; i++) { 
                if (i % 3 == 0 && i % 5 == 0) {
                    results.append("FizzBuzz");
                } else if (i % 3 == 0) { 
                    results.append("Fizz");
                } else if (i % 5 == 0) {
                    results.append("Buzz");
                } else {
                    results.append(i); 
                }
                if (i != x) {
                    results.append("-"); 
                }
            }
            return results.toString(); 
        };
        System.out.println(fizzBuzz.apply(9));
        System.out.println(fizzBuzz.apply(20)); 
    }
}
