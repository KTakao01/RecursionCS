//https://recursionist.io/dashboard/problems/440
import java.util.Arrays;
import java.util.stream.*;

class Solution{
    public static int[] ageCheck(int[] ages){
        // 関数を完成させてください
        //Arrays.stream()メソッド プリミティブ型
        IntStream intStream = Arrays.stream(ages);
        return intStream.filter(x -> x >= 18).toArray();

    }
}