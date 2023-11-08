//https://recursionist.io/dashboard/problems/523
class Assertion {
    public static void run(boolean b) throws Exception {
        if (!b) throw new Exception("Assertion Error");
        else System.out.println("Passed the assertion test.");
    }
}

class MyClass {
    public static int cube(int x) {
        return x * x * x;
    }

    public static void main(String[] args) throws Exception {
        // ここから書きましょう
        Assertion.run(cube(2) == 7);
        Assertion.run(cube(3) == 27);
        Assertion.run(cube(4) == 64);
        Assertion.run(cube(0) == 0);
        Assertion.run(cube(-2) == -8);
        Assertion.run(cube(-3) == -27);
    }
}
