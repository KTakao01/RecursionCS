//https://recursionist.io/dashboard/problems/531
import java.util.ArrayList;
import java.util.List;
import java.util.function.*;
import java.util.stream.*;

class MyClass {
    public static ArrayList<ArrayList<Integer>> getDivisorCombinationList(int n) {
        ArrayList<ArrayList<Integer>> divisor = new ArrayList<>();
        for (int i = 1; i <= Math.ceil(Math.sqrt(n)); i++) {
            ArrayList<Integer> combination = new ArrayList<>();
            if (n % i != 0) continue;
            combination.add(i);
            combination.add(n / i);
            divisor.add(combination);
        }
        return divisor;
    }

    public static boolean divisorCheck(int k, int num) {
        // getDivisorCombinationList関数の整合性をチェックしてください
        //numの約数の組み合わせはk個ある
        int count = 0;
        for (int i = 0 ; i < getDivisorCombinationList(num).size() ; i++) {
                count++;
        }
        return count == k;
    }

    public static void unitTestCheck(boolean predicate) {
        if (predicate) {
            System.out.println("The test passed!!");
        } else {
            System.out.println("ERROR! The test failed!!");
        }
    }

    public static void unitTests(int[] inputs, int[] outputs) {
        // ここから書いてください
        IntStream.range(0,inputs.length).forEach(i -> unitTestCheck(divisorCheck(inputs[i],outputs[i])));
    }

    public static void main(String[] args){
        unitTests(new int[]{0,2,8,15,18,30,45}, new int[]{1,10,120,720,1260,5040,25200});
    }
}
