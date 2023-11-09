//https://recursionist.io/dashboard/problems/525
class Assertion {
    public static void run(boolean b) throws Exception {
        if (!b) throw new Exception("Assertion Error");
        else System.out.println("Passed the assertion test.");
    }
}

class MyClass {
    public static String fizzBuzz(int n) {
        String output = "";
        if (n <= 0) output = "";
        else if (n % 15 == 0) output = "fizzBuzz";
        else if (n % 3 == 0) output = "Fizz";
        else if (n % 5 == 0) output = "Buzz";
        else output = "not covered";

        //System.out.println(output);
        return output;
    }

    public static void main(String[] args) throws Exception {
        // ここから書きましょう
        Assertion.run(fizzBuzz(7) == "not covered");
        Assertion.run(fizzBuzz(5) == "Buzz");
        Assertion.run(fizzBuzz(3) == "Fizz");
        Assertion.run(fizzBuzz(15) == "fizzBuzz");
        Assertion.run(fizzBuzz(0) == "");
        Assertion.run(fizzBuzz(-3) == "");
    }
}
