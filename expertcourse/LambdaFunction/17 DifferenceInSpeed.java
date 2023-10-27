//https://recursionist.io/dashboard/problems/555
/*下記のような文字列として表された配列 pointsData が与えられます。速度 10m/s 以上の要素の速度 [m/s] を配列として返す upper10mpsList という関数を作成してください。

["x1-x2-s"]: x1m 地点から x2m 地点まで s 秒かかる
*/

import java.util.Arrays;
import java.util.function.*;

class Solution{
    public static String[] upper10mpsList(String[] pointsData){
        // 関数を完成させてください
        Function<String,Integer> calcVelocity = x -> {
            String[] data = x.split("-");
            double length1 = Double.parseDouble(data[0]);
            double length2 = Double.parseDouble(data[1]);
            double time    = Double.parseDouble(data[2]);
            return (int)(Math.abs(length2 - length1) / time); 
        };

        return Arrays.stream(pointsData)
            .map(calcVelocity)
            .peek(x -> System.out.println("Velocity: " + x)) 
            .filter(x -> x>=10)
            .peek(x -> System.out.println("Filtered: " + x))
            .map(String::valueOf)
            .peek(x -> System.out.println("Mapped: " + x)) 
            .toArray(String[]::new);

    }
}


