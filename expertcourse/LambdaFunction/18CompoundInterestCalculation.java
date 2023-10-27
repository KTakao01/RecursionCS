//https://recursionist.io/dashboard/problems/443
/*複利計算とは、投資した金額に年利が加算されるというプロセスを数年間繰り返した場合に、最終的に得られる収益を計算することです。ここでは、投資した金額（初期値）と、複数年間の年利が格納された配列が与えられるので、reduce 関数を使用して最終的な収益を計算してください。
例えば、初期値が 1000 円で、配列が [3,5,2,5,4] の場合、年利が 3%, 5%, 2%, 5%, 4% と変化していくので、3 年後の最終的な収益は 100 * (1 + 0.03) * (1 + 0.05) * (1 + 0.02) * (1 + 0.05) * (1 + 0.04) = 120 円となります。
*/

import java.util.Arrays;

class Solution{
    public static int calculateFinalMoney(int[] interests, int capital){
        // 関数を完成させてください
        double capitalGain = Arrays.stream(interests)
            .asDoubleStream() 
            .reduce(capital, (total, x) -> total * (1 + x / 100.0));

        return (int)capitalGain;
    }
}

