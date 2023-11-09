//https://recursionist.io/dashboard/problems/530
class MyClass {
    public static int factorial(int n) {
        if (n <= 0) return 1;
        return n * factorial(n - 1);
    }

    public static boolean factorialCheck(int number, int result) {
        // factorial関数の整合性をチェックしてください
        return factorial(number) == result;
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
        for (int i = 0 ; i < inputs.length; i++) {
           unitTestCheck(factorialCheck(inputs[i],outputs[i]));
        }
    }

    public static void main(String[] args) {
        unitTests(new int[]{1,2,3,4,5,6,7,8},new int[]{1,1,6,24,130,720,5040,42000});
    }
}