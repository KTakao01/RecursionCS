//https://recursionist.io/dashboard/problems/556
//整数と単位（mm, m, km）が文字列のリスト l が与えられるので、map を使って m に変換し、filter を使って 100m 以上のリストを返す over100m という関数を作成してください。
import java.util.Arrays;
import java.util.function.Function;

class Solution{
    public static double[] over100m(String[] l){
        // 関数を完成させてください
        Function<String,Double> changeMeter = x -> {
            int len = x.length();
            if (x.contains("mm")) { 
                x = x.substring(0,len-2);
                double num = Double.parseDouble(x);
                return num / 1000;
            } else if (x.contains("km")) {
                x = x.substring(0,len-2);
                double num = Double.parseDouble(x);
                return num * 1000;
            } else {
                x = x.substring(0,len-1);
                double num = Double.parseDouble(x);
                return num;            
            }
        };

        return Arrays.stream(l)
            .mapToDouble(changeMeter::apply)
            .filter(x -> x >= 100)
            .toArray();
    }
}

